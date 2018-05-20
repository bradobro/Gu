package au_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/bradobro/au"
)

func assertTrue(t au.T, actual bool, msg string) {
	if !actual {
		t.Errorf("Expected true: %s", msg)
	}
}

func assertEquals(t au.T, actual, expected string) {
	if actual != expected {
		t.Errorf("Actual %#v\nExpected %#v", actual, expected)
	}
}

func testStringContains(t au.T, actual, contains string) {
	if !strings.Contains(actual, contains) {
		t.Errorf("%#v\ndoes not contain %#v", actual, contains)
	}
}

/*CustomT wraps a *testing.T in ways that test packages can be tested.
 */
type CustomT struct {
	t          au.T
	NoFail     bool // if true, t.Fail* and t.Error* only mark a flag
	Failed     bool
	FailedFast bool
	Writer     io.Writer
}

func NewT(t au.T, noFail bool, writer io.Writer) *CustomT {
	ct := &CustomT{t: t, NoFail: noFail, Writer: writer}
	return ct
}

func newTestT(t au.T) (*CustomT, *bytes.Buffer) {
	var buf bytes.Buffer
	ct := NewT(t, true, &buf)
	return ct, &buf
}

func (t *CustomT) Error(args ...interface{}) {
	t.Log(args...)
	t.Fail()
}

func (t *CustomT) Errorf(format string, args ...interface{}) {
	t.Logf(format, args...)
	t.Fail()
}

func (t *CustomT) Fail() {
	if t.NoFail {
		t.Failed = true
	} else {
		t.t.Fail()
	}
}

func (t *CustomT) FailNow() {
	if t.NoFail {
		t.Failed = true
		t.FailedFast = true
	} else {
		t.t.FailNow()
	}
}

func (t *CustomT) Log(args ...interface{}) {
	if t.Writer == nil {
		t.t.Log(args...)
	} else {
		fmt.Fprint(t.Writer, args...)
	}
}

func (t *CustomT) Name() string {
	if gt, ok := t.t.(*testing.T); ok {
		return gt.Name()
	}
	return ""
}

func (t *CustomT) Logf(format string, args ...interface{}) {
	if t.Writer == nil {
		t.t.Log(args...)
	} else {
		fmt.Fprintf(t.Writer, format, args...)
	}
}

func TestAssert(t *testing.T) {
	ct, _ := newTestT(t)

	tt := au.NewAsserter(false, 4, au.VerbosityInsane)
	x := 5
	y := x + 1

	tt.Assert(ct, au.Equal, x, x)
	tt.Assert(ct, au.Unequal, x, y)
	assertTrue(t, !(ct.Failed || ct.FailedFast), "these tests shouldn't have failed")

	tt.Assert(ct, au.Equal, x, y)
	assertTrue(t, ct.Failed, "this test should have failed")
	assertTrue(t, !ct.FailedFast, "this test should not have failed fast")

	ct.Failed = false
	tt.FailFast = true
	tt.Assert(ct, au.Equal, x, y)
	assertTrue(t, ct.Failed, "this test should have failed")
	assertTrue(t, ct.FailedFast, "this test should have failed fast")
}

// func TestTweakStackDepth(t *testing.T) {
// 	a := au.NewAsserter(t, false, 25, au.VerbosityInsane)
// 	a.Assert(au.Equal, 2, 5)
// }
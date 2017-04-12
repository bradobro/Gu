// This file is automatically generated by qtc from "tablepage.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:1
package templates

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

// Table page template. Implements BasePage methods.
//

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:4
type TablePage struct {
	Rows []string
}

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:10
func (p *TablePage) StreamTitle(qw422016 *qt422016.Writer) {
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:10
	qw422016.N().S(`
	This is table page
`)
//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:12
}

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:12
func (p *TablePage) WriteTitle(qq422016 qtio422016.Writer) {
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:12
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:12
	p.StreamTitle(qw422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:12
	qt422016.ReleaseWriter(qw422016)
//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:12
}

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:12
func (p *TablePage) Title() string {
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:12
	qb422016 := qt422016.AcquireByteBuffer()
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:12
	p.WriteTitle(qb422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:12
	qs422016 := string(qb422016.B)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:12
	qt422016.ReleaseByteBuffer(qb422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:12
	return qs422016
//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:12
}

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:15
func (p *TablePage) StreamBody(qw422016 *qt422016.Writer) {
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:15
	qw422016.N().S(`
	<h1>Table page</h1>

	`)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:18
	p.streamform(qw422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:18
	qw422016.N().S(`

	`)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:20
	if len(p.Rows) == 0 {
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:20
		qw422016.N().S(`
		No rows. Click <a href="/table?rowsCount=5">here</a>.
	`)
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:22
	} else {
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:22
		qw422016.N().S(`
		<table>
			`)
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:24
		streamemitRows(qw422016, p.Rows)
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:24
		qw422016.N().S(`
		</table>
	`)
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:26
	}
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:26
	qw422016.N().S(`
`)
//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:27
}

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:27
func (p *TablePage) WriteBody(qq422016 qtio422016.Writer) {
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:27
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:27
	p.StreamBody(qw422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:27
	qt422016.ReleaseWriter(qw422016)
//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:27
}

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:27
func (p *TablePage) Body() string {
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:27
	qb422016 := qt422016.AcquireByteBuffer()
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:27
	p.WriteBody(qb422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:27
	qs422016 := string(qb422016.B)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:27
	qt422016.ReleaseByteBuffer(qb422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:27
	return qs422016
//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:27
}

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:29
func streamemitRows(qw422016 *qt422016.Writer, rows []string) {
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:29
	qw422016.N().S(`
	<tr>
		<th>#</th>
		<th>value</th>
	</tr>

	`)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:35
	for n, r := range rows {
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:35
		qw422016.N().S(`
		`)
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:36
		if r == "bingo" {
			//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:36
			qw422016.N().S(`
			<tr><td colspan="2"><h1>BINGO!</h1></td></tr>
			`)
			//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:38
			return
			//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:39
		} else if n == 42 {
			//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:39
			qw422016.N().S(`
			<tr><td colspan="2">42 rows already generated</td></tr>
			`)
			//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:41
			break
			//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:42
		}
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:42
		qw422016.N().S(`

		<tr style="background: `)
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:44
		if n&1 == 1 {
			//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:44
			qw422016.N().S(`white`)
			//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:44
		} else {
			//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:44
			qw422016.N().S(`#ddd`)
			//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:44
		}
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:44
		qw422016.N().S(`">
			<td>`)
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:45
		qw422016.N().D(n + 1)
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:45
		qw422016.N().S(`</td>
			<td>`)
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:46
		qw422016.E().S(r)
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:46
		qw422016.N().S(`</td>
		</tr>
	`)
		//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:48
	}
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:48
	qw422016.N().S(`

	<tr><td colspan="2">No bingo found</td></tr>
`)
//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:51
}

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:51
func writeemitRows(qq422016 qtio422016.Writer, rows []string) {
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:51
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:51
	streamemitRows(qw422016, rows)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:51
	qt422016.ReleaseWriter(qw422016)
//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:51
}

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:51
func emitRows(rows []string) string {
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:51
	qb422016 := qt422016.AcquireByteBuffer()
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:51
	writeemitRows(qb422016, rows)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:51
	qs422016 := string(qb422016.B)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:51
	qt422016.ReleaseByteBuffer(qb422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:51
	return qs422016
//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:51
}

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:53
func (p *TablePage) streamform(qw422016 *qt422016.Writer) {
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:53
	qw422016.N().S(`
	<form>
		Rows: <input type="text" name="rowsCount" value="`)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:55
	qw422016.N().D(len(p.Rows))
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:55
	qw422016.N().S(`"/><br/>
		<input type="submit" value="Generate!"/>
	</form>
`)
//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:58
}

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:58
func (p *TablePage) writeform(qq422016 qtio422016.Writer) {
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:58
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:58
	p.streamform(qw422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:58
	qt422016.ReleaseWriter(qw422016)
//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:58
}

//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:58
func (p *TablePage) form() string {
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:58
	qb422016 := qt422016.AcquireByteBuffer()
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:58
	p.writeform(qb422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:58
	qs422016 := string(qb422016.B)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:58
	qt422016.ReleaseByteBuffer(qb422016)
	//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:58
	return qs422016
//line vendor/github.com/valyala/quicktemplate/examples/basicserver/templates/tablepage.qtpl:58
}

// Code generated by qtc from "error.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:1
package templates

//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:2
type ErrorPage struct {
	*BasePage
	Message string
}

//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
func (p *ErrorPage) StreamTitle(qw422016 *qt422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
	qw422016.N().S(`Error`)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
}

//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
func (p *ErrorPage) WriteTitle(qq422016 qtio422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
	p.StreamTitle(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
}

//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
func (p *ErrorPage) Title() string {
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
	p.WriteTitle(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
	qs422016 := string(qb422016.B)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
	return qs422016
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:8
}

//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:10
func (p *ErrorPage) StreamNavbarMargin(qw422016 *qt422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:10
}

//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:10
func (p *ErrorPage) WriteNavbarMargin(qq422016 qtio422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:10
	p.StreamNavbarMargin(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:10
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:10
}

//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:10
func (p *ErrorPage) NavbarMargin() string {
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:10
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:10
	p.WriteNavbarMargin(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:10
	qs422016 := string(qb422016.B)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:10
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:10
	return qs422016
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:10
}

//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:12
func (p *ErrorPage) StreamBody(qw422016 *qt422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:12
	qw422016.N().S(`
<section class="hero is-large">
    <div class="hero-body">
        <div class="container">
            <h1 class="title">Uh Oh Sisters!!!!1!!!11!!!!!</h1>
            <h2 class="subtitle">`)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:17
	qw422016.E().S(p.Message)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:17
	qw422016.N().S(`</h2>
        </div>
    </div>
</section>
`)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:21
}

//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:21
func (p *ErrorPage) WriteBody(qq422016 qtio422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:21
	p.StreamBody(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:21
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:21
}

//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:21
func (p *ErrorPage) Body() string {
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:21
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:21
	p.WriteBody(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:21
	qs422016 := string(qb422016.B)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:21
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:21
	return qs422016
//line /home/max/git/jaunty/redux/internal/web/templates/error.qtpl:21
}
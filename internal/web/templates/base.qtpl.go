// Code generated by qtc from "base.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line base.qtpl:1
package templates

//line base.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line base.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line base.qtpl:1
type Page interface {
//line base.qtpl:1
	Title() string
//line base.qtpl:1
	StreamTitle(qw422016 *qt422016.Writer)
//line base.qtpl:1
	WriteTitle(qq422016 qtio422016.Writer)
//line base.qtpl:1
	Meta() string
//line base.qtpl:1
	StreamMeta(qw422016 *qt422016.Writer)
//line base.qtpl:1
	WriteMeta(qq422016 qtio422016.Writer)
//line base.qtpl:1
	NavbarLogin() string
//line base.qtpl:1
	StreamNavbarLogin(qw422016 *qt422016.Writer)
//line base.qtpl:1
	WriteNavbarLogin(qq422016 qtio422016.Writer)
//line base.qtpl:1
	Body() string
//line base.qtpl:1
	StreamBody(qw422016 *qt422016.Writer)
//line base.qtpl:1
	WriteBody(qq422016 qtio422016.Writer)
//line base.qtpl:1
}

//line base.qtpl:9
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
//line base.qtpl:9
	qw422016.N().S(`
<!DOCTYPE html>
<html>
    <head>
        <title>`)
//line base.qtpl:13
	p.StreamTitle(qw422016)
//line base.qtpl:13
	qw422016.N().S(` &bull; Jaunty</title>

        <link rel="stylesheet" href="/static/jaunty.css">

        `)
//line base.qtpl:17
	p.StreamMeta(qw422016)
//line base.qtpl:17
	qw422016.N().S(`
    </head>
    <body>
        <nav class="navbar is-jaunty mb-5">
            <div class="container">
                <div class="navbar-brand is-size-5">
                    <a class="navbar-item" href="/">ᕕ(՞ᗜ՞)ᕗ</a>
                </div>

                <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false">
                    <span aria-hidden="true"></span>
                    <span aria-hidden="true"></span>
                    <span aria-hidden="true"></span>
                </a>

                <div class="navbar-menu">
                    <div class="navbar-start">
                        <a class="navbar-item" href="/join">
                            Join
                        </a>
                    </div>

                    <div class="navbar-end">
                        `)
//line base.qtpl:40
	p.StreamNavbarLogin(qw422016)
//line base.qtpl:40
	qw422016.N().S(`
                    </div>
                </div>
            </div>
        </nav>

        `)
//line base.qtpl:46
	p.StreamBody(qw422016)
//line base.qtpl:46
	qw422016.N().S(`
    </body>
</html>
`)
//line base.qtpl:49
}

//line base.qtpl:49
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
//line base.qtpl:49
	qw422016 := qt422016.AcquireWriter(qq422016)
//line base.qtpl:49
	StreamPageTemplate(qw422016, p)
//line base.qtpl:49
	qt422016.ReleaseWriter(qw422016)
//line base.qtpl:49
}

//line base.qtpl:49
func PageTemplate(p Page) string {
//line base.qtpl:49
	qb422016 := qt422016.AcquireByteBuffer()
//line base.qtpl:49
	WritePageTemplate(qb422016, p)
//line base.qtpl:49
	qs422016 := string(qb422016.B)
//line base.qtpl:49
	qt422016.ReleaseByteBuffer(qb422016)
//line base.qtpl:49
	return qs422016
//line base.qtpl:49
}

//line base.qtpl:51
type User struct {
	Username  string
	Snowflake string
	Avatar    string
}

//line base.qtpl:58
type BasePage struct {
	User *User
}

//line base.qtpl:61
func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
//line base.qtpl:61
	qw422016.N().S(`Jaunty`)
//line base.qtpl:61
}

//line base.qtpl:61
func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
//line base.qtpl:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line base.qtpl:61
	p.StreamTitle(qw422016)
//line base.qtpl:61
	qt422016.ReleaseWriter(qw422016)
//line base.qtpl:61
}

//line base.qtpl:61
func (p *BasePage) Title() string {
//line base.qtpl:61
	qb422016 := qt422016.AcquireByteBuffer()
//line base.qtpl:61
	p.WriteTitle(qb422016)
//line base.qtpl:61
	qs422016 := string(qb422016.B)
//line base.qtpl:61
	qt422016.ReleaseByteBuffer(qb422016)
//line base.qtpl:61
	return qs422016
//line base.qtpl:61
}

//line base.qtpl:62
func (p *BasePage) StreamNavbarLogin(qw422016 *qt422016.Writer) {
//line base.qtpl:62
	qw422016.N().S(`
    `)
//line base.qtpl:63
	if p.User == nil {
//line base.qtpl:63
		qw422016.N().S(`
        <a class="navbar-item" href="/login">Log in</a>
    `)
//line base.qtpl:65
	} else {
//line base.qtpl:65
		qw422016.N().S(`
        <a class="navbar-item" href="/dashboard">`)
//line base.qtpl:66
		qw422016.E().S(p.User.Username)
//line base.qtpl:66
		qw422016.N().S(`</a>
    `)
//line base.qtpl:67
	}
//line base.qtpl:67
	qw422016.N().S(`
`)
//line base.qtpl:68
}

//line base.qtpl:68
func (p *BasePage) WriteNavbarLogin(qq422016 qtio422016.Writer) {
//line base.qtpl:68
	qw422016 := qt422016.AcquireWriter(qq422016)
//line base.qtpl:68
	p.StreamNavbarLogin(qw422016)
//line base.qtpl:68
	qt422016.ReleaseWriter(qw422016)
//line base.qtpl:68
}

//line base.qtpl:68
func (p *BasePage) NavbarLogin() string {
//line base.qtpl:68
	qb422016 := qt422016.AcquireByteBuffer()
//line base.qtpl:68
	p.WriteNavbarLogin(qb422016)
//line base.qtpl:68
	qs422016 := string(qb422016.B)
//line base.qtpl:68
	qt422016.ReleaseByteBuffer(qb422016)
//line base.qtpl:68
	return qs422016
//line base.qtpl:68
}

//line base.qtpl:70
func (p *BasePage) StreamBody(qw422016 *qt422016.Writer) {
//line base.qtpl:70
}

//line base.qtpl:70
func (p *BasePage) WriteBody(qq422016 qtio422016.Writer) {
//line base.qtpl:70
	qw422016 := qt422016.AcquireWriter(qq422016)
//line base.qtpl:70
	p.StreamBody(qw422016)
//line base.qtpl:70
	qt422016.ReleaseWriter(qw422016)
//line base.qtpl:70
}

//line base.qtpl:70
func (p *BasePage) Body() string {
//line base.qtpl:70
	qb422016 := qt422016.AcquireByteBuffer()
//line base.qtpl:70
	p.WriteBody(qb422016)
//line base.qtpl:70
	qs422016 := string(qb422016.B)
//line base.qtpl:70
	qt422016.ReleaseByteBuffer(qb422016)
//line base.qtpl:70
	return qs422016
//line base.qtpl:70
}

//line base.qtpl:71
func (p *BasePage) StreamMeta(qw422016 *qt422016.Writer) {
//line base.qtpl:71
}

//line base.qtpl:71
func (p *BasePage) WriteMeta(qq422016 qtio422016.Writer) {
//line base.qtpl:71
	qw422016 := qt422016.AcquireWriter(qq422016)
//line base.qtpl:71
	p.StreamMeta(qw422016)
//line base.qtpl:71
	qt422016.ReleaseWriter(qw422016)
//line base.qtpl:71
}

//line base.qtpl:71
func (p *BasePage) Meta() string {
//line base.qtpl:71
	qb422016 := qt422016.AcquireByteBuffer()
//line base.qtpl:71
	p.WriteMeta(qb422016)
//line base.qtpl:71
	qs422016 := string(qb422016.B)
//line base.qtpl:71
	qt422016.ReleaseByteBuffer(qb422016)
//line base.qtpl:71
	return qs422016
//line base.qtpl:71
}

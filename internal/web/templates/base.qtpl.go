// Code generated by qtc from "base.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
package templates

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
type Page interface {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	Title() string
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	StreamTitle(qw422016 *qt422016.Writer)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	WriteTitle(qq422016 qtio422016.Writer)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	Meta() string
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	StreamMeta(qw422016 *qt422016.Writer)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	WriteMeta(qq422016 qtio422016.Writer)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	NavbarLogin() string
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	StreamNavbarLogin(qw422016 *qt422016.Writer)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	WriteNavbarLogin(qq422016 qtio422016.Writer)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	NavbarMargin() string
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	StreamNavbarMargin(qw422016 *qt422016.Writer)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	WriteNavbarMargin(qq422016 qtio422016.Writer)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	Body() string
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	StreamBody(qw422016 *qt422016.Writer)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
	WriteBody(qq422016 qtio422016.Writer)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:1
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:10
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:10
	qw422016.N().S(`
<!DOCTYPE html>
<html>
    <head>
        <title>`)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:14
	p.StreamTitle(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:14
	qw422016.N().S(` &bull; Jaunty</title>

        <link rel="stylesheet" href="/static/jaunty.css">

        `)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:18
	p.StreamMeta(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:18
	qw422016.N().S(`
    </head>
    <body>
        <nav class="navbar is-jaunty `)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:21
	p.StreamNavbarMargin(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:21
	qw422016.N().S(`">
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
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:41
	p.StreamNavbarLogin(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:41
	qw422016.N().S(`
                    </div>
                </div>
            </div>
        </nav>

        `)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:47
	p.StreamBody(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:47
	qw422016.N().S(`
    </body>
</html>
`)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:50
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:50
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:50
	StreamPageTemplate(qw422016, p)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:50
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:50
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:50
func PageTemplate(p Page) string {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:50
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:50
	WritePageTemplate(qb422016, p)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:50
	qs422016 := string(qb422016.B)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:50
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:50
	return qs422016
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:50
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:52
type User struct {
	Username  string
	Snowflake string
	Avatar    string
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:59
type BasePage struct {
	User *User
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
	qw422016.N().S(`Jaunty`)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
	p.StreamTitle(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
func (p *BasePage) Title() string {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
	p.WriteTitle(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
	qs422016 := string(qb422016.B)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
	return qs422016
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:62
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:63
func (p *BasePage) StreamNavbarLogin(qw422016 *qt422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:63
	qw422016.N().S(`
    `)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:64
	if p.User == nil {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:64
		qw422016.N().S(`
        <a class="navbar-item" href="/login">Log in</a>
    `)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:66
	} else {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:66
		qw422016.N().S(`
        <a class="navbar-item" href="/dashboard">`)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:67
		qw422016.E().S(p.User.Username)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:67
		qw422016.N().S(`</a>
    `)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:68
	}
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:68
	qw422016.N().S(`
`)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:69
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:69
func (p *BasePage) WriteNavbarLogin(qq422016 qtio422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:69
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:69
	p.StreamNavbarLogin(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:69
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:69
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:69
func (p *BasePage) NavbarLogin() string {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:69
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:69
	p.WriteNavbarLogin(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:69
	qs422016 := string(qb422016.B)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:69
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:69
	return qs422016
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:69
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:71
func (p *BasePage) StreamNavbarMargin(qw422016 *qt422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:71
	qw422016.N().S(`
mb-5
`)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:73
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:73
func (p *BasePage) WriteNavbarMargin(qq422016 qtio422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:73
	p.StreamNavbarMargin(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:73
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:73
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:73
func (p *BasePage) NavbarMargin() string {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:73
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:73
	p.WriteNavbarMargin(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:73
	qs422016 := string(qb422016.B)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:73
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:73
	return qs422016
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:73
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:75
func (p *BasePage) StreamBody(qw422016 *qt422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:75
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:75
func (p *BasePage) WriteBody(qq422016 qtio422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:75
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:75
	p.StreamBody(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:75
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:75
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:75
func (p *BasePage) Body() string {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:75
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:75
	p.WriteBody(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:75
	qs422016 := string(qb422016.B)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:75
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:75
	return qs422016
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:75
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:76
func (p *BasePage) StreamMeta(qw422016 *qt422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:76
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:76
func (p *BasePage) WriteMeta(qq422016 qtio422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:76
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:76
	p.StreamMeta(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:76
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:76
}

//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:76
func (p *BasePage) Meta() string {
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:76
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:76
	p.WriteMeta(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:76
	qs422016 := string(qb422016.B)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:76
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:76
	return qs422016
//line /home/max/git/jaunty/redux/internal/web/templates/base.qtpl:76
}

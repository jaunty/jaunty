// Code generated by qtc from "join.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:1
package templates

//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:2
type JoinPage struct {
	*BasePage
}

//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
func (p *JoinPage) StreamTitle(qw422016 *qt422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
	qw422016.N().S(`Join`)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
}

//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
func (p *JoinPage) WriteTitle(qq422016 qtio422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
	p.StreamTitle(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
}

//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
func (p *JoinPage) Title() string {
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
	p.WriteTitle(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
	qs422016 := string(qb422016.B)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
	return qs422016
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:7
}

//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:9
func (p *JoinPage) StreamBody(qw422016 *qt422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:9
	qw422016.N().S(`
    `)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:10
	if p.BasePage.User != nil {
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:10
		qw422016.N().S(`
        <div class="container">
            <div class="columns is-centered">
                <div class="column is-half">
                    <h1 class="title">Check Yourself Before You Whitelist Yourself</h1>
                    <hr>
                </div>
            </div>

            <div class="columns is-centered">
                <div class="column is-one-third">
                    <form action="" method="POST">
                        <div class="field has-addons has-addons-centered">
                            <div class="control">
                                <input id="username" name="username" class="input" type="text" placeholder="Minecraft Username">
                            </div>
                            <div class="control">
                                <input for="username" type="submit" class="button is-jaunty" value="Submit">
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    `)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:34
	} else {
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:34
		qw422016.N().S(`
        <div class="modal is-active">
            <div class="modal-background"></div>
            <div class="modal-content box">
                <h1 class="title">This page requires you to be logged in!</h1>
                <h2 class="subtitle">Click <a href="/login">here</a> to do that.</h2>
            </div>
        </div>
    `)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:42
	}
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:42
	qw422016.N().S(`
`)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:43
}

//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:43
func (p *JoinPage) WriteBody(qq422016 qtio422016.Writer) {
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:43
	p.StreamBody(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:43
	qt422016.ReleaseWriter(qw422016)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:43
}

//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:43
func (p *JoinPage) Body() string {
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:43
	qb422016 := qt422016.AcquireByteBuffer()
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:43
	p.WriteBody(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:43
	qs422016 := string(qb422016.B)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:43
	qt422016.ReleaseByteBuffer(qb422016)
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:43
	return qs422016
//line /home/max/git/jaunty/redux/internal/web/templates/join.qtpl:43
}

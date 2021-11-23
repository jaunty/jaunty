package web

import (
	"net/http"

	"github.com/jaunty/jaunty/internal/web/templates"
)

func (s *Server) basePage(r *http.Request) *templates.BasePage {
	sess := s.getSession(r)
	bp := new(templates.BasePage)

	if !sess.isNew() {
		bp.User = &templates.User{
			Username:  sess.getUsername(),
			Snowflake: sess.getSnowflake(),
			Avatar:    sess.getAvatar(),
		}
	}

	return bp
}

func (s *Server) serveError(w http.ResponseWriter, r *http.Request, msg string) {
	templates.WritePageTemplate(w, &templates.ErrorPage{
		BasePage: s.basePage(r),
		Message:  msg,
	})
}

package web

import (
	"net/http"
	"strings"

	"github.com/jaunty/jaunty/internal/web/templates"
)

func (s *Server) basePage(r *http.Request) *templates.BasePage {
	sess := s.getSession(r)

	pth := r.URL.Path
	if !strings.HasPrefix(pth, "/") {
		pth = "/" + pth
	}

	bp := &templates.BasePage{
		Path: pth,
	}

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

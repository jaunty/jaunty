package web

import (
	"net/http"

	"github.com/jaunty/jaunty/internal/web/templates"
)

func (s *Server) recoverFunc(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.ErrorPage{
		BasePage: s.basePage(r),
		Message:  "Internal Server Error",
	})
}

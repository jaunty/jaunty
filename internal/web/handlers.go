package web

import (
	"net/http"

	"github.com/jaunty/jaunty/internal/web/templates"
)

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.IndexPage{})
}

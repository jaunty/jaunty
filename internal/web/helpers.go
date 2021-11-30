package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jaunty/jaunty/internal/web/templates"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

func (s *Server) recoverFunc(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.ErrorPage{
		BasePage: s.basePage(r),
		Message:  "Internal Server Error",
	})
}

func respondJSON(ctx context.Context, w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		ctxlog.Error(ctx, "unable to marshal json", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body) //nolint:errcheck
}

func getRedirect(to string, r *http.Request) string {
	pth := r.URL.Path
	if !strings.HasPrefix(pth, "/") {
		pth = "/" + pth
	}

	return fmt.Sprintf("%s?redirect=%s", to, pth)
}

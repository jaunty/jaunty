package web

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/holedaemon/web/middleware"
	"github.com/jaunty/jaunty/internal/database/models"
	"github.com/jaunty/jaunty/internal/web/templates"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

func (s *Server) authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sess := s.getSession(r)
		tx := middleware.TxFromContext(ctx)

		user, err := models.Users(qm.Where("sf = ?", sess.getSnowflake())).One(ctx, tx)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				ctxlog.Info(ctx, "user without an account tried to access an authenticated page", zap.String("sf", sess.getSnowflake()))
				s.serveError(w, r, "The page you're trying to access requires that you're logged in!!")
				return
			}

			ctxlog.Error(ctx, "error getting user from database", zap.Error(err))
			s.serveError(w, r, "Unable to query user from database")
			return
		}

		if user.Banned {
			templates.WritePageTemplate(w, &templates.BannedPage{
				BasePage: s.basePage(r),
			})
			return
		}

		if sess.isNew() {
			s.serveError(w, r, "The page you're trying to access requires that you're logged in!!")
			return
		}

		next.ServeHTTP(w, r)
	})
}

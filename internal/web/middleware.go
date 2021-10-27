package web

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/jaunty/jaunty/internal/database/models"
	"github.com/jaunty/jaunty/internal/web/templates"
	"github.com/rs/xid"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

// Credits to Zik
// https://github.com/hortbot/hortbot/blob/master/internal/web/mid/mid.go

type requestIDKey struct{}

type txKey struct{}

const requestIDHeader = "X-Request-ID"

func requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var id xid.ID
		requestID := r.Header.Get(requestIDHeader)

		if requestID != "" {
			var err error
			id, err = xid.FromString(requestID)
			if err != nil {
				old := requestID
				id = xid.New()
				requestID = id.String()

				ctxlog.Debug(ctx, "replacing request ID", zap.String("old", old), zap.String("new", requestID))
			}
		} else {
			id = xid.New()
			requestID = id.String()
		}

		w.Header().Set(requestIDHeader, requestID)
		ctx = context.WithValue(ctx, requestIDKey{}, id)
		ctx = ctxlog.With(ctx, zap.String("requestID", requestID))

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// Logger adds a logger to a request chain.
func logger(l *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := ctxlog.WithLogger(r.Context(), l)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func (s *Server) authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sess := s.getSession(r)
		tx := txFromCtx(ctx)

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

func (s *Server) transaction(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tx, err := s.db.BeginTx(ctx, nil)
		if err != nil {
			ctxlog.Error(ctx, "error starting transaction", zap.Error(err))
			s.serveError(w, r, "Unable to start database transaction")
			return
		}

		defer func() {
			if err := tx.Rollback(); err != nil {
				if !errors.Is(err, sql.ErrTxDone) {
					ctxlog.Error(ctx, "error rolling back transaction", zap.Error(err))
					s.serveError(w, r, "Unable to rollback transaction")
					return
				}
			}
		}()

		ctx = context.WithValue(ctx, txKey{}, tx)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

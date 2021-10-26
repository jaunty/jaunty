package web

import (
	"context"
	"net/http"

	"github.com/jaunty/jaunty/internal/web/templates"
	"github.com/rs/xid"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

// Credits to Zik
// https://github.com/hortbot/hortbot/blob/master/internal/web/mid/mid.go

type requestIDKey struct{}

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
		sess := s.getSession(r)

		if sess.isNew() {
			templates.WritePageTemplate(w, &templates.ErrorPage{
				BasePage: s.basePage(r),
				Message:  "The page you're trying to access requires that you're logged in!!",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
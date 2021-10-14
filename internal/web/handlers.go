package web

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
	"github.com/jaunty/jaunty/internal/web/templates"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

const authTimeout = time.Hour

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.IndexPage{})
}

func (s *Server) authDiscord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	state := uuid.Must(uuid.NewV4()).String()

	redirect := r.URL.Query().Get("redirect")
	redir, err := url.QueryUnescape(redirect)
	if err != nil {
		redir = "/"
	}

	if err := s.redis.SetStringWithExpiration(ctx, state, redir, authTimeout); err != nil {
		ctxlog.Error(ctx, "error setting state in redis", zap.Error(err))
		http.Error(w, "Error caching state", http.StatusInternalServerError)
		return
	}

	url := s.discord.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (s *Server) authDiscordCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	state := r.FormValue("state")
	if state == "" {
		http.Error(w, "Wrong state returned", http.StatusBadRequest)
		return
	}

	_, err := s.redis.FetchString(ctx, state)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			http.Error(w, "States do not match", http.StatusBadRequest)
			return
		}

		ctxlog.Error(ctx, "error getting state from redis", zap.Error(err))
		http.Error(w, "Error retrieving state from cache", http.StatusInternalServerError)
		return
	}

	_, err = s.discord.Exchange(ctx, r.FormValue("code"))
	if err != nil {
		ctxlog.Error(ctx, "error exchanging code", zap.Error(err))
		http.Error(w, "Error exchanging OAuth2 code", http.StatusInternalServerError)
		return
	}
}

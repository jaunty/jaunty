package web

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
	"github.com/jaunty/jaunty/internal/database/models"
	"github.com/jaunty/jaunty/internal/database/modelsx"
	"github.com/jaunty/jaunty/internal/pkg/api/mojang"
	"github.com/jaunty/jaunty/internal/web/templates"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

const authTimeout = time.Hour

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	s.writePageTemplate(w, r, &templates.IndexPage{})
}

func (s *Server) join(w http.ResponseWriter, r *http.Request) {
	s.writePageTemplate(w, r, &templates.JoinPage{})
}

func (s *Server) postJoin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		s.writeError(w, r, "Unable to parse given form.")
		return
	}

	un := r.FormValue("username")
	if un == "" {
		s.writeError(w, r, "Minecraft username cannot be blank")
		return
	}

	uid, err := s.mojang.FetchUUIDByUsername(ctx, un)
	if err != nil {
		if errors.Is(err, mojang.ErrNotFound) {
			s.writeError(w, r, "Username does not exist according to Mojang")
			return
		}

		ctxlog.Error(ctx, "error getting uuid by username from Mojang", zap.Error(err))
		s.writeError(w, r, "Unable to convert Minecraft username into UUID")
		return
	}

	exists, err := models.Whitelists(qm.Where("uuid = ?", uid)).Exists(ctx, s.db)
	if err != nil {
		ctxlog.Error(ctx, "error getting whitelist from database", zap.Error(err))
		s.writeError(w, r, "Error checking request's existence in the database")
		return
	}

	if exists {
		s.writeError(w, r, "A whitelist request already exists for the given account")
		return
	}

	sess := s.getSession(r)

	wh := models.Whitelist{
		SF:   sess.getSnowflake(),
		UUID: uid,
	}

	if err := wh.Insert(ctx, s.db, boil.Infer()); err != nil {
		ctxlog.Error(ctx, "error creating whitelist request", zap.Error(err))
		s.writeError(w, r, "Error inserting whitelist request into the database")
		return
	}

	s.writePageTemplate(w, r, &templates.NewRequestPage{
		Username: un,
	})
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
		s.writeError(w, r, "Error caching state")
		return
	}

	url := s.discord.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (s *Server) authDiscordCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	state := r.FormValue("state")
	if state == "" {
		s.writeError(w, r, "Unexpected state was returned")
		return
	}

	_, err := s.redis.FetchString(ctx, state)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			s.writeError(w, r, "Given state does not match")
			return
		}

		ctxlog.Error(ctx, "error getting state from redis", zap.Error(err))
		s.writeError(w, r, "Unable to retrieve state from Redis")
		return
	}

	token, err := s.discord.Exchange(ctx, r.FormValue("code"))
	if err != nil {
		ctxlog.Error(ctx, "error exchanging code", zap.Error(err))
		s.writeError(w, r, "Error exchanging OAuth2 code for access token")
		return
	}

	user, err := s.discord.GetCurrentUser(ctx, token.AccessToken)
	if err != nil {
		ctxlog.Error(ctx, "error getting user from discord", zap.Error(err))
		s.writeError(w, r, "Error getting user from Discord's API")
		return
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		ctxlog.Error(ctx, "error beginning transaction", zap.Error(err))
		s.writeError(w, r, "Error starting database transaction")
		return
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			if !errors.Is(err, sql.ErrTxDone) {
				ctxlog.Error(ctx, "error rolling back transaction", zap.Error(err))
			}
		}
	}()

	if err := modelsx.UpsertToken(ctx, tx, user.ID, token); err != nil {
		ctxlog.Error(ctx, "error upserting token", zap.Error(err))
		fmt.Println(err)
		s.writeError(w, r, "Error upserting OAuth2 token in the database")
		return
	}

	if err := tx.Commit(); err != nil {
		ctxlog.Error(ctx, "error committing transaction", zap.Error(err))
		s.writeError(w, r, "Error committing database transaction")
		return
	}

	sess := s.getSession(r)
	sess.clear()
	sess.setSnowflake(user.ID)
	sess.setUsername(user.Username + "#" + user.Discriminator)
	sess.setAvatar(user.Avatar.String)

	if err := sess.save(w, r); err != nil {
		ctxlog.Error(ctx, "error saving session", zap.Error(err))
		s.writeError(w, r, "Error saving session")
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) destroyAuth(w http.ResponseWriter, r *http.Request) {
	if err := s.destroySession(w, r); err != nil {
		s.writeError(w, r, "Error destroying session???")
		return
	}

	s.writePageTemplate(w, r, &templates.IndexPage{})
}

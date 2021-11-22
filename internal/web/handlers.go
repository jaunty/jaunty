package web

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/disaccord/beelzebub/flies/guild"
	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
	"github.com/holedaemon/web/middleware"
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
	templates.WritePageTemplate(w, &templates.IndexPage{
		BasePage: s.basePage(r),
	})
}

func (s *Server) join(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.JoinPage{
		BasePage: s.basePage(r),
	})
}

func (s *Server) postJoin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		s.serveError(w, r, "Unable to parse given form.")
		return
	}

	un := r.FormValue("username")
	if un == "" {
		s.serveError(w, r, "Minecraft username cannot be blank")
		return
	}

	uid, err := s.mojang.FetchUUIDByUsername(ctx, un)
	if err != nil {
		if errors.Is(err, mojang.ErrNotFound) {
			s.serveError(w, r, "Username does not exist according to Mojang")
			return
		}

		ctxlog.Error(ctx, "error getting uuid by username from Mojang", zap.Error(err))
		s.serveError(w, r, "Unable to convert Minecraft username into UUID")
		return
	}

	tx := middleware.TxFromContext(ctx)

	exists, err := models.Whitelists(qm.Where("uuid = ?", uid)).Exists(ctx, tx)
	if err != nil {
		ctxlog.Error(ctx, "error getting whitelist from database", zap.Error(err))
		s.serveError(w, r, "Error checking request's existence in the database")
		return
	}

	if exists {
		s.serveError(w, r, "A whitelist request already exists for the given account")
		return
	}

	sess := s.getSession(r)

	count, err := models.Whitelists(qm.Where("sf = ?", sess.getSnowflake())).Count(ctx, tx)
	if err != nil {
		ctxlog.Error(ctx, "error counting requests in database", zap.Error(err))
		s.serveError(w, r, "Error counting your requests in the database")
		return
	}

	if count >= int64(s.maxRequests) {
		s.serveError(w, r, "You have reached the maximum number of requests allowed.")
		return
	}

	wh := models.Whitelist{
		SF:   sess.getSnowflake(),
		UUID: uid,
	}

	if err := wh.Insert(ctx, tx, boil.Infer()); err != nil {
		ctxlog.Error(ctx, "error creating whitelist request", zap.Error(err))
		s.serveError(w, r, "Error inserting whitelist request into the database")
		return
	}

	if err := tx.Commit(); err != nil {
		ctxlog.Error(ctx, "error committing transaction", zap.Error(err))
		s.serveError(w, r, "Error committing transaction")
		return
	}

	if err := s.SendWhitelistNotification(ctx, sess.getUsername(), un); err != nil {
		ctxlog.Error(ctx, "error sending notification to channel", zap.Error(err))
	}

	templates.WritePageTemplate(w, &templates.NewRequestPage{
		BasePage: s.basePage(r),
		Username: un,
	})
}

func (s *Server) dashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sess := s.getSession(r)
	tx := middleware.TxFromContext(ctx)

	wr, err := models.Whitelists(qm.Where("sf = ?", sess.getSnowflake())).All(ctx, tx)
	if err != nil {
		ctxlog.Error(ctx, "error getting whitelist requests", zap.Error(err), zap.String("sf", sess.getSnowflake()))
		s.serveError(w, r, "Unable to fetch whitelist requests")
		return
	}

	names := make(map[string]string)
	for _, wr := range wr {
		name, err := s.mojang.FetchUsernameByUUID(ctx, wr.UUID)
		if err != nil {
			ctxlog.Error(ctx, "error getting username by uuid", zap.Error(err))
			s.serveError(w, r, "Error getting username by UUID.")
			return
		}

		names[wr.UUID] = name
	}

	if err := tx.Commit(); err != nil {
		ctxlog.Error(ctx, "error committing transaction", zap.Error(err))
		s.serveError(w, r, "Error committing transaction")
		return
	}

	templates.WritePageTemplate(w, &templates.DashboardPage{
		BasePage:      s.basePage(r),
		Requests:      wr,
		ResolvedUUIDs: names,
	})
}

func (s *Server) accountDelete(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.AccountDeletePage{
		BasePage: s.basePage(r),
	})
}

func (s *Server) postAccountDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sess := s.getSession(r)
	sf := sess.getSnowflake()
	tx := middleware.TxFromContext(ctx)

	user, err := models.Users(qm.Where("sf = ?", sf)).One(ctx, tx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.serveError(w, r, "User doesn't exist")
			return
		}

		ctxlog.Error(ctx, "error querying user from database", zap.Error(err))
		s.serveError(w, r, "Unable to query user from database")
		return
	}

	if err := user.Delete(ctx, tx); err != nil {
		ctxlog.Error(ctx, "error removing user from database", zap.Error(err))
		s.serveError(w, r, "Unable to delete user from database")
		return
	}

	if err := tx.Commit(); err != nil {
		ctxlog.Error(ctx, "error committing transaction", zap.Error(err))
		s.serveError(w, r, "Error committing transaction")
		return
	}

	if err := s.SendSiteNotification(ctx, "Account Deleted", sess.getUsername()); err != nil {
		ctxlog.Error(ctx, "error sending notification", zap.Error(err))
	}

	if err := s.destroySession(w, r); err != nil {
		ctxlog.Error(ctx, "error destroying session", zap.Error(err))
	}

	gf := s.discord.Guild(s.guildID)
	err = gf.RemoveMember(ctx, sf, "User has deleted their Jaunty account.")
	if err != nil {
		ctxlog.Error(ctx, "error kicking user from the discord", zap.Error(err))
		templates.WritePageTemplate(w, &templates.SuccessPage{
			BasePage:  s.basePage(r),
			PageTitle: "Account Deleted",
			Header:    "Account has been deleted",
			SubHeader: "For some reason I was not able to kick you from Discord, you should be nice and leave.",
		})
		return
	}

	templates.WritePageTemplate(w, &templates.SuccessPage{
		BasePage:  s.basePage(r),
		PageTitle: "Account Deleted",
		Header:    "Account has been deleted",
		SubHeader: "See You Space Cowboy or something idk I didn't read the book",
	})
}

func (s *Server) requestCancel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		s.serveError(w, r, "Request ID is blank, doofus")
		return
	}

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		s.serveError(w, r, "Unable to parse request ID into int64")
		return
	}

	templates.WritePageTemplate(w, &templates.CancelRequestPage{
		BasePage:  s.basePage(r),
		RequestID: i,
	})
}

func (s *Server) postRequestCancel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		ctxlog.Error(ctx, "error parsing form", zap.Error(err))
		s.serveError(w, r, "Unable to parse HTML form")
		return
	}

	id := r.FormValue("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		s.serveError(w, r, "Unable to parse request ID into int64")
		return
	}

	tx := middleware.TxFromContext(ctx)

	sess := s.getSession(r)
	req, err := models.Whitelists(qm.Where("id = ? AND sf = ?", i, sess.getSnowflake())).One(ctx, tx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.serveError(w, r, "Request does not exist, or it's not yours to delete")
			return
		}

		ctxlog.Error(ctx, "error getting whitelist request from database", zap.Error(err))
		s.serveError(w, r, "Unable to query your request from the database")
		return
	}

	if err := req.Delete(ctx, tx); err != nil {
		ctxlog.Error(ctx, "error deleting whitelist request", zap.Error(err))
		s.serveError(w, r, "Unable to remove request from the database")
		return
	}

	if err := tx.Commit(); err != nil {
		ctxlog.Error(ctx, "error committing transaction", zap.Error(err))
		s.serveError(w, r, "Error committing transaction")
		return
	}

	templates.WritePageTemplate(w, &templates.SuccessPage{
		BasePage:  s.basePage(r),
		PageTitle: "Request Deleted",
		Header:    "Your request has been deleted",
		SubHeader: "Later, idiot!!",
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
		s.serveError(w, r, "Error caching state")
		return
	}

	url := s.oauth2.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (s *Server) authDiscordCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	state := r.FormValue("state")
	if state == "" {
		s.serveError(w, r, "Unexpected state was returned")
		return
	}

	_, err := s.redis.FetchString(ctx, state)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			s.serveError(w, r, "Given state does not match")
			return
		}

		ctxlog.Error(ctx, "error getting state from redis", zap.Error(err))
		s.serveError(w, r, "Unable to retrieve state from Redis")
		return
	}

	token, err := s.oauth2.Exchange(ctx, r.FormValue("code"))
	if err != nil {
		ctxlog.Error(ctx, "error exchanging code", zap.Error(err))
		s.serveError(w, r, "Error exchanging OAuth2 code for access token")
		return
	}

	uf := s.discord.UserWithToken(fmt.Sprintf("Bearer %s", token.AccessToken))
	user, err := uf.Get(ctx)
	if err != nil {
		ctxlog.Error(ctx, "error getting user from discord", zap.Error(err))
		s.serveError(w, r, "Error getting user from Discord's API")
		return
	}

	gf := s.discord.Guild(s.guildID)
	_, err = gf.AddMember(ctx, user.ID, &guild.AddMemberOptions{
		AccessToken: token.AccessToken,
	})
	if err != nil {
		ctxlog.Error(ctx, "error adding user to discord", zap.Error(err))
		s.serveError(w, r, "Error adding user to the Discord server.")
		return
	}

	tx := middleware.TxFromContext(ctx)

	exists, err := models.Users(qm.Where("sf = ?", user.ID)).Exists(ctx, tx)
	if err != nil {
		ctxlog.Error(ctx, "error querying user from database", zap.Error(err))
		s.serveError(w, r, "Unable to query user from database")
		return

	}

	if !exists {
		u := &models.User{
			SF: user.ID,
		}

		if err := u.Insert(ctx, tx, boil.Infer()); err != nil {
			ctxlog.Error(ctx, "error creating new user", zap.Error(err))
			s.serveError(w, r, "Error creating new user account")
			return
		}
	}

	if err := modelsx.UpsertToken(ctx, tx, user.ID, token); err != nil {
		ctxlog.Error(ctx, "error upserting token", zap.Error(err))
		s.serveError(w, r, "Error upserting OAuth2 token in the database")
		return
	}

	if err := tx.Commit(); err != nil {
		ctxlog.Error(ctx, "error committing transaction", zap.Error(err))
		s.serveError(w, r, "Error committing transaction")
		return
	}

	sess := s.getSession(r)
	sess.clear()
	sess.setSnowflake(user.ID)
	sess.setUsername(
		fmt.Sprintf("%s#%s", user.Username, user.Discriminator),
	)

	if err := sess.save(w, r); err != nil {
		ctxlog.Error(ctx, "error saving session", zap.Error(err))
		s.serveError(w, r, "Error saving session")
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) destroyAuth(w http.ResponseWriter, r *http.Request) {
	if err := s.destroySession(w, r); err != nil {
		s.serveError(w, r, "Error destroying session???")
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

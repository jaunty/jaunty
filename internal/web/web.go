package web

import (
	"context"
	"database/sql"
	"embed"
	"io/fs"
	"net/http"
	"time"

	"github.com/disaccord/beelzebub"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/holedaemon/web"
	"github.com/holedaemon/web/middleware"
	"github.com/jaunty/jaunty/internal/pkg/api/mojang"
	"github.com/jaunty/jaunty/internal/pkg/rcon"
	"github.com/jaunty/jaunty/internal/web/templates"
	"github.com/patrickmn/go-cache"
	"github.com/zikaeroh/ctxlog"
	"golang.org/x/oauth2"
)

const (
	cacheExpir = time.Hour * 1
)

//go:embed static
var static embed.FS

var staticDir fs.FS

func init() {
	var err error
	staticDir, err = fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}
}

// Options configures a Server.
type Options struct {
	Addr        string
	SessionKey  []byte
	MaxRequests int
	PublicKey   string

	GuildID               string
	WhitelistChannelID    string
	NotificationChannelID string
	ApprovedRoleID        string

	DB      *sql.DB
	RCON    *rcon.Client
	OAuth2  *oauth2.Config
	Discord *beelzebub.Devil
	Mojang  *mojang.Client
}

// Server is responsible for serving the website and auth server.
type Server struct {
	addr        string
	maxRequests int
	publicKey   string

	guildID               string
	whitelistChannelID    string
	notificationChannelID string
	approvedRoleID        string

	interactionHandlers map[string]intHandler

	store   sessions.Store
	discord *beelzebub.Devil
	oauth2  *oauth2.Config
	mojang  *mojang.Client

	cache *cache.Cache

	rcon *rcon.Client
	db   *sql.DB
}

// New creates a new Server.
func New(opts *Options) (*Server, error) {
	s := &Server{
		addr:        opts.Addr,
		maxRequests: opts.MaxRequests,
		publicKey:   opts.PublicKey,

		whitelistChannelID:    opts.WhitelistChannelID,
		notificationChannelID: opts.NotificationChannelID,
		guildID:               opts.GuildID,
		approvedRoleID:        opts.ApprovedRoleID,

		interactionHandlers: make(map[string]intHandler),

		db: opts.DB,

		discord: opts.Discord,
		oauth2:  opts.OAuth2,
		mojang:  opts.Mojang,
		rcon:    opts.RCON,

		cache: cache.New(cacheExpir, cacheExpir),

		store: sessions.NewCookieStore(opts.SessionKey),
	}

	s.interactionHandlers["whitelist-approve"] = s.handlerApproveWhitelist
	s.interactionHandlers["whitelist-reject"] = s.handlerRejectWhitelist

	return s, nil
}

func (s *Server) router(ctx context.Context) *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger(ctxlog.FromContext(ctx)),
		middleware.Recoverer(s.recoverFunc),
		middleware.RequestID,
		middleware.Transaction(s.db),
	)

	r.Get("/", s.index)

	r.Group(func(r chi.Router) {
		r.Use(s.authenticator)

		r.Get("/join", s.join)
		r.Post("/join", s.postJoin)

		r.Get("/dashboard", s.dashboard)
		r.Get("/dashboard/account/delete", s.accountDelete)
		r.Post("/dashboard/account/delete", s.postAccountDelete)
		r.Get("/dashboard/request/delete", s.requestDelete)
		r.Post("/dashboard/request/delete", s.postRequestDelete)
	})

	r.Get("/login", s.authDiscord)
	r.Get("/auth", s.authDiscord)
	r.Get("/auth/callback", s.authDiscordCallback)
	r.Get("/auth/destroy", s.destroyAuth)
	r.Get("/logout", s.destroyAuth)
	r.Post("/webhook", s.postWebhook)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		templates.WritePageTemplate(w, &templates.ErrorPage{
			BasePage: s.basePage(r),
			Message:  "Whatever you're looking for ain't here",
		})
	})

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticDir))))

	return r
}

// Start runs a Server.
func (s *Server) Start(ctx context.Context) error {
	return web.Start(ctx, &web.Options{
		Addr:    s.addr,
		Service: "jaunty",
		Router:  s.router(ctx),
	})
}

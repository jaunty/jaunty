package web

import (
	"context"
	"database/sql"
	"embed"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/jaunty/jaunty/internal/pkg/api/discord"
	"github.com/jaunty/jaunty/internal/pkg/redisx"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
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
	Addr       string
	SessionKey []byte

	DB      *sql.DB
	Redis   *redisx.Redis
	Discord *discord.Client
}

// Server is responsible for serving the website and auth server.
type Server struct {
	addr string

	store sessions.Store

	db      *sql.DB
	discord *discord.Client
	redis   *redisx.Redis
}

// New creates a new Server.
func New(opts *Options) (*Server, error) {
	s := &Server{
		addr:    opts.Addr,
		db:      opts.DB,
		discord: opts.Discord,
		redis:   opts.Redis,
		store:   sessions.NewCookieStore(opts.SessionKey),
	}

	return s, nil
}

func (s *Server) router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", s.index)

	r.Get("/login", s.authDiscord)
	r.Get("/auth", s.authDiscord)
	r.Get("/auth/callback", s.authDiscordCallback)

	r.Get("/auth/destroy", s.destroyAuth)
	r.Get("/logout", s.destroyAuth)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticDir))))

	return r
}

// Start runs a Server.
func (s *Server) Start(ctx context.Context) error {
	srv := &http.Server{
		Handler: s.router(),
		Addr:    s.addr,
	}

	go func() {
		<-ctx.Done()

		if err := srv.Shutdown(context.Background()); err != nil {
			ctxlog.Error(ctx, "error shutting down server", zap.Error(err))
		}
	}()

	ctxlog.Info(ctx, "starting server", zap.String("addr", s.addr))
	return srv.ListenAndServe()
}

package web

import (
	"context"
	"database/sql"
	"embed"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
	"github.com/jaunty/jaunty/internal/pkg/api/discord"
	"github.com/jaunty/jaunty/internal/pkg/api/mojang"
	"github.com/jaunty/jaunty/internal/pkg/redisx"
	"github.com/jaunty/jaunty/internal/web/templates"
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
	Mojang  *mojang.Client
}

// Server is responsible for serving the website and auth server.
type Server struct {
	addr string

	store sessions.Store

	db      *sql.DB
	discord *discord.Client
	mojang  *mojang.Client
	redis   *redisx.Redis
}

// New creates a new Server.
func New(opts *Options) (*Server, error) {
	s := &Server{
		addr:    opts.Addr,
		db:      opts.DB,
		discord: opts.Discord,
		mojang:  opts.Mojang,
		redis:   opts.Redis,
		store:   sessions.NewCookieStore(opts.SessionKey),
	}

	return s, nil
}

func (s *Server) basePage(r *http.Request) *templates.BasePage {
	sess := s.getSession(r)
	bp := new(templates.BasePage)

	if !sess.isNew() {
		bp.User = &templates.User{
			Username:  sess.getUsername(),
			Snowflake: sess.getSnowflake(),
			Avatar:    sess.getAvatar(),
		}
	}

	return bp
}

func (s *Server) router(ctx context.Context) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(requestID)
	r.Use(logger(ctxlog.FromContext(ctx)))

	r.Get("/", s.index)

	r.Group(func(r chi.Router) {
		r.Use(s.authenticator)

		r.Get("/join", s.join)
		r.Post("/join", s.postJoin)

		r.Get("/dashboard", s.dashboard)
		r.Get("/dashboard/account/delete", s.accountDelete)
		r.Post("/dashboard/account/delete", s.postAccountDelete)
		r.Get("/dashboard/request/cancel", s.requestCancel)
		r.Post("/dashboard/request/cancel", s.postRequestCancel)

		r.Get("/auth/destroy", s.destroyAuth)
		r.Get("/logout", s.destroyAuth)
	})

	r.Get("/login", s.authDiscord)
	r.Get("/auth", s.authDiscord)
	r.Get("/auth/callback", s.authDiscordCallback)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		templates.WritePageTemplate(w, &templates.ErrorPage{
			BasePage: s.basePage(r),
			Message:  "Whatever you're looking for ain't here",
		})
	})

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticDir))))

	return r
}

func (s *Server) serveError(w http.ResponseWriter, r *http.Request, msg string) {
	templates.WritePageTemplate(w, &templates.ErrorPage{
		BasePage: s.basePage(r),
		Message:  msg,
	})
}

// Start runs a Server.
func (s *Server) Start(ctx context.Context) error {
	srv := &http.Server{
		Handler: s.router(ctx),
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

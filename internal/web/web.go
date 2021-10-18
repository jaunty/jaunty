package web

import (
	"context"
	"embed"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
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

type Options struct {
	Addr string
}

// Server is responsible for serving the website and auth server.
type Server struct {
	Addr string
}

// New creates a new Server.
func New(opts *Options) (*Server, error) {
	s := &Server{
		Addr: opts.Addr,
	}

	return s, nil
}

func (s *Server) router() *chi.Mux {
	r := chi.NewRouter()

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticDir))))

	return r
}

// Start runs a Server.
func (s *Server) Start(ctx context.Context) error {
	srv := &http.Server{
		Handler: s.router(),
		Addr:    s.Addr,
	}

	go func() {
		if err := srv.Shutdown(context.Background()); err != nil {
			ctxlog.Error(ctx, "error shutting down server", zap.Error(err))
		}
	}()

	ctxlog.Info(ctx, "starting server", zap.String("addr", s.Addr))
	return srv.ListenAndServe()
}

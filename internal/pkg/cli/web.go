package cli

import (
	"context"

	"github.com/jaunty/jaunty/internal/web"
	"github.com/zikaeroh/ctxlog"
)

// Web is a subcommand that runs the web server.
type Web struct {
	Addr string `help:"Address to listen for connections on." short:"a" default:":8080" env:"JAUNTY_WEB_ADDR"`
}

// Run executes the Web subcommand.
func (w *Web) Run(ctx context.Context, debug bool) error {
	l := ctxlog.New(debug)
	ctx = ctxlog.WithLogger(ctx, l)

	opts := &web.Options{
		Addr: w.Addr,
	}

	srv, err := web.New(opts)
	if err != nil {
		return err
	}

	return srv.Start(ctx)
}

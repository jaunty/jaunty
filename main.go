package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/alecthomas/kong"
	"github.com/jaunty/jaunty/internal/pkg/cli"
)

func main() {
	app := cli.CLI{}

	ctx, stop := signal.NotifyContext(context.Background(), os.Kill)
	defer stop()

	ktx := kong.Parse(&app,
		kong.Name("jaunty"),
		kong.Description("CLI application for running Jaunty."),
		kong.UsageOnError(),
		kong.Configuration(cli.EnvLoader, ".env"),
		kong.BindTo(ctx, (*context.Context)(nil)),
	)

	err := ktx.Run(app.Debug)
	ktx.FatalIfErrorf(err)
}

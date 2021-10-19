package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/alecthomas/kong"
	"github.com/jaunty/jaunty/internal/pkg/cli"
)

func main() {
	cli := cli.CLI{}

	ctx, stop := signal.NotifyContext(context.Background(), os.Kill)
	defer stop()

	ktx := kong.Parse(&cli,
		kong.Name("jaunty"),
		kong.Description("CLI application for running Jaunty."),
		kong.UsageOnError(),
		kong.BindTo(ctx, (*context.Context)(nil)),
	)

	err := ktx.Run(cli.Debug)
	ktx.FatalIfErrorf(err)
}

package cli

import (
	"context"

	"github.com/jaunty/jaunty/internal/pkg/api/discord"
	"github.com/jaunty/jaunty/internal/pkg/dbx"
	"github.com/jaunty/jaunty/internal/pkg/redisx"
	"github.com/jaunty/jaunty/internal/web"
	"github.com/zikaeroh/ctxlog"
	"golang.org/x/oauth2"
)

var endpoint = oauth2.Endpoint{
	TokenURL: "https://discord.com/api/oauth2/token",
	AuthURL:  "https://discord.com/api/oauth2/authorize",
}

// Web is a subcommand that runs the web server.
type Web struct {
	Addr       string `help:"Address to listen for connections on." env:"JAUNTY_WEB_ADDR"`
	SessionKey string `help:"Secret key for session encryption." required:"" env:"JAUNTY_SESSION_KEY"`
	DSN        string `help:"PostgreSQL DSN." required:"" env:"JAUNTY_DB_DSN"`
	Redis      string `help:"Address to Redis server." required:"" env:"JAUNTY_REDIS_ADDR"`

	BotToken string `help:"Discord bot OAuth2 token for API access" required:"" env:"JAUNTY_BOT_TOKEN"`

	ClientID     string   `help:"Discord OAuth2 client ID." required:"" env:"JAUNTY_OAUTH2_CLIENT_ID"`
	ClientSecret string   `help:"Discord OAuth2 client secret." required:"" env:"JAUNTY_OAUTH2_CLIENT_SECRET"`
	Scopes       []string `help:"Discord OAuth2 scopes." required:"" env:"JAUNTY_OAUTH2_SCOPES"`
	RedirectURI  string   `help:"Discord OAuth2 redirect uri" required:"" env:"JAUNTY_OAUTH2_REDIRECT_URI"`
}

// Run executes the Web subcommand.
func (w *Web) Run(ctx context.Context, debug bool) error {
	l := ctxlog.New(debug)
	ctx = ctxlog.WithLogger(ctx, l)

	rdb := redisx.Open(w.Redis)
	if err := rdb.Ping(ctx); err != nil {
		return err
	}

	db, err := dbx.Open(ctx, w.DSN)
	if err != nil {
		return err
	}

	defer db.Close()

	oa := &oauth2.Config{
		ClientID:     w.ClientID,
		ClientSecret: w.ClientSecret,
		Endpoint:     endpoint,
		RedirectURL:  w.RedirectURI,
		Scopes:       w.Scopes,
	}

	dsc, err := discord.New(&discord.Options{
		BotToken: w.BotToken,
		Redis:    rdb,
		OAuth2:   oa,
	})
	if err != nil {
		return err
	}

	opts := &web.Options{
		Addr:       w.Addr,
		SessionKey: []byte(w.SessionKey),
		DB:         db,
		Redis:      rdb,

		Discord: dsc,
	}

	srv, err := web.New(opts)
	if err != nil {
		return err
	}

	return srv.Start(ctx)
}

package cli

import (
	"context"
	"strings"

	"github.com/disaccord/beelzebub"
	"github.com/jaunty/jaunty/internal/pkg/api/mojang"
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

	GuildID            string `help:"Guild ID for the associated Discord" required:"" env:"JAUNTY_GUILD_ID"`
	WhitelistChannelID string `help:"Channel ID for the whitelist notifications channel" required:"" env:"JAUNTY_WHITELIST_CHANNEL_ID"`

	MaxRequests int `help:"Maximum whitelist requests per user." default:"2" env:"JAUNTY_MAX_REQUESTS"`

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

	var token string
	if !strings.HasPrefix(token, "Bot ") {
		token = "Bot " + w.BotToken
	} else {
		token = w.BotToken
	}

	dsc, err := beelzebub.New(token)
	if err != nil {
		return err
	}

	moj, err := mojang.New(&mojang.Options{
		Redis: rdb,
	})
	if err != nil {
		return err
	}

	opts := &web.Options{
		Addr:               w.Addr,
		SessionKey:         []byte(w.SessionKey),
		DB:                 db,
		Redis:              rdb,
		MaxRequests:        w.MaxRequests,
		GuildID:            w.GuildID,
		WhitelistChannelID: w.WhitelistChannelID,

		Discord: dsc,
		OAuth2:  oa,
		Mojang:  moj,
	}

	srv, err := web.New(opts)
	if err != nil {
		return err
	}

	return srv.Start(ctx)
}

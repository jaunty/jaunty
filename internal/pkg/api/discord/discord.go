package discord

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaunty/jaunty/internal/pkg/httpx"
	"github.com/jaunty/jaunty/internal/pkg/redisx"
	"golang.org/x/oauth2"
)

const (
	contentType = "application/json"
	accept      = "application/json"
	userAgent   = "DiscordBot (https://github.com/jaunty/jaunty, v0.0.0)"

	root = "https://discord.com/api/v9/"
)

// Client interacts with Discord's REST API.
// It implements the api.Client interface.
type Client struct {
	cli   *http.Client
	rdb   *redisx.Redis
	oauth *oauth2.Config
}

// Options configures a Client.
type Options struct {
	Client *http.Client
	Redis  *redisx.Redis
	OAuth2 *oauth2.Config
}

// New creates a new Client with the given options.
func New(opts *Options) (*Client, error) {
	cli := &Client{
		cli:   opts.Client,
		rdb:   opts.Redis,
		oauth: opts.OAuth2,
	}

	if cli.cli == nil {
		cli.cli = http.DefaultClient
	}

	if cli.rdb == nil {
		return nil, fmt.Errorf("discord: missing redis client")
	}

	if cli.oauth == nil {
		return nil, fmt.Errorf("discord: missing oauth2 config")
	}

	return cli, nil
}

// Do performs an HTTP request against Discord's API.
func (c *Client) Do(ctx context.Context, url string, opts ...httpx.RequestOption) (*http.Response, error) {
	req, err := httpx.NewRequest(ctx, url, opts...)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", accept)
	req.Header.Set("User-Agent", userAgent)

	return c.cli.Do(req)
}

// FetchValue fetches a value from the Client's caching backend.
func (c *Client) FetchValue(ctx context.Context, key string) (interface{}, error) {
	return c.rdb.FetchValue(ctx, key)
}

// SetVaue sets a value in the CLient's caching backend.
func (c *Client) SetValue(ctx context.Context, key string, val interface{}) error {
	return c.rdb.SetValue(ctx, key, val)
}

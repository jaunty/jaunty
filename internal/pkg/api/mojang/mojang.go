package mojang

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaunty/jaunty/internal/pkg/httpx"
	"github.com/jaunty/jaunty/internal/pkg/redisx"
)

const (
	root        = "https://api.mojang.com"
	contentType = "application/json"
	accept      = "application/json"
)

// Options configure a Client.
type Options struct {
	Redis  *redisx.Redis
	Client *http.Client
}

// Client interacts with Mojang's API for Minecraft.
type Client struct {
	rdb *redisx.Redis
	cli *http.Client
}

// New creates a new Client.
func New(opts *Options) (*Client, error) {
	c := &Client{
		rdb: opts.Redis,
		cli: opts.Client,
	}

	if opts.Client == nil {
		c.cli = http.DefaultClient
	}

	if opts.Redis == nil {
		return nil, fmt.Errorf("mojang: missing redis client")
	}

	return c, nil
}

func (c *Client) Do(ctx context.Context, uri string, opts ...httpx.RequestOption) (*http.Response, error) {
	u := fmt.Sprintf("%s/%s", root, uri)

	req, err := httpx.NewRequest(ctx, u, opts...)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", accept)

	return c.cli.Do(req)
}

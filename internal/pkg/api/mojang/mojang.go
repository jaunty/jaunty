package mojang

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jaunty/jaunty/internal/pkg/httpx"
)

const (
	root        = "https://api.mojang.com"
	contentType = "application/json"
	accept      = "application/json"
)

// Client interacts with Mojang's API for Minecraft.
type Client struct {
	cli *http.Client
}

// New creates a new Client.
func New() *Client {
	return &Client{
		cli: http.DefaultClient,
	}
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

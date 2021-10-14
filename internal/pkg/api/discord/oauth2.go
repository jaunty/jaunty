package discord

import (
	"context"

	"golang.org/x/oauth2"
)

// AuthCodeURL returns a URL to Discord's consent page
// that asks for permissions for the required scopes explicitly.
func (c *Client) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return c.oauth.AuthCodeURL(state, opts...)
}

// Exchange an authorization code into a token.
func (c *Client) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return c.oauth.Exchange(ctx, code, opts...)
}

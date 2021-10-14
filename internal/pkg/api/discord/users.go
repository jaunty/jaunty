package discord

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/holedaemon/tumult"
	"github.com/jaunty/jaunty/internal/pkg/httpx"
)

// GetCurrentUser fetches a Discord user by their associated OAuth2 access token.
func (c *Client) GetCurrentUser(ctx context.Context, token string) (*tumult.User, error) {
	res, err := c.Do(ctx, "/users/@me",
		httpx.WithHeaders(httpx.Headers("Authorization", fmt.Sprintf("Bearer %s", token))),
	)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		e := new(Error)
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		}

		return nil, e
	}

	u := new(tumult.User)
	if err := json.NewDecoder(res.Body).Decode(&u); err != nil {
		return nil, err
	}

	return u, nil
}

// GetUser fetches a Discord user by their snowflake.
func (c *Client) GetUser(ctx context.Context, sf string) (*tumult.User, error) {
	res, err := c.Do(ctx,
		fmt.Sprintf("/users/%s", sf),
	)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		e := new(Error)
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		}

		return nil, e
	}

	u := new(tumult.User)
	if err := json.NewDecoder(res.Body).Decode(&u); err != nil {
		return nil, err
	}

	return u, nil
}

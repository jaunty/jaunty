package mojang

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jaunty/jaunty/internal/pkg/httpx"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

// ErrNotFound is returned when user does not exist per Mojang.
var ErrNotFound = errors.New("mojang: user not found")

type getUUIDResponse struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// FetchUUIDByUsernames fetches a Minecraft UUID by the player's username.
func (c *Client) FetchUUIDByUsername(ctx context.Context, username string) (string, error) {
	uuid, err := c.rdb.FetchString(ctx, username)
	if err == nil {
		return uuid, nil
	} else {
		ctxlog.Error(ctx, "error retrieving string from redis, ignoring", zap.Error(err))
	}

	res, err := c.Do(ctx, fmt.Sprintf("/users/profiles/minecraft/%s", username))
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusNoContent {
		return "", ErrNotFound
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return "", fmt.Errorf("%w: HTTP %d", httpx.ErrStatusCode, res.StatusCode)
	}

	var rq *getUUIDResponse
	if err := json.NewDecoder(res.Body).Decode(&rq); err != nil {
		return "", err
	}

	if err := c.rdb.SetString(ctx, username, rq.ID); err != nil {
		return "", err
	}

	return rq.ID, nil
}

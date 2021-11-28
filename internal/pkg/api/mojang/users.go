package mojang

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/jaunty/jaunty/internal/pkg/httpx"
)

// ErrNotFound is returned when user does not exist per Mojang.
var ErrNotFound = errors.New("mojang: user not found")

type getUUIDResponse struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// FetchUUIDByUsername fetches a Minecraft UUID by the player's username.
func (c *Client) FetchUUIDByUsername(ctx context.Context, username string) (string, error) {
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

	return rq.ID, nil
}

type getUsernameResponse struct {
	Name        string  `json:"name"`
	ChangedToAt mojTime `json:"changedToAt"`
}

type mojTime struct {
	time.Time
}

func (m *mojTime) UnmarshalJSON(b []byte) error {
	i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}

	if i == 0 {
		m.Time = time.Time{}
		return nil
	}

	t := time.Unix(i/1000, i*(1e+6))
	m.Time = t
	return nil
}

// FetchUsernameByUUID retrieves a username by the associated UUID.
func (c *Client) FetchUsernameByUUID(ctx context.Context, uuid string) (string, error) {
	res, err := c.Do(ctx, fmt.Sprintf("/user/profiles/%s/names", uuid))
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	var names []*getUsernameResponse
	if err := json.NewDecoder(res.Body).Decode(&names); err != nil {
		return "", err
	}

	for i, n := range names {
		if n.ChangedToAt.IsZero() {
			names[i] = names[len(names)-1]
		}
	}

	sort.Slice(names, func(i, j int) bool {
		return names[j].ChangedToAt.Before(names[i].ChangedToAt.Time)
	})

	return names[0].Name, nil
}

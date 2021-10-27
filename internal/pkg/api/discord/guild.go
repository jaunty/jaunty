package discord

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jaunty/jaunty/internal/pkg/httpx"
)

// DeleteGuildMember kicks a member from a guild.
func (c *Client) DeleteGuildMemberWithReason(ctx context.Context, guildID, userID, reason string) error {
	u := fmt.Sprintf("%s/guilds/%s/members/%s", root, guildID, userID)

	opts := make([]httpx.RequestOption, 0)

	if reason != "" {
		opts = append(opts, httpx.WithHeaders(
			httpx.Headers("X-Audit-Log-Reason", reason),
		))
	}

	res, err := c.Do(ctx, u, opts...)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		e := new(Error)
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return err
		}

		return e
	}

	return nil
}

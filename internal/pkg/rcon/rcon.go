package rcon

import (
	"fmt"
	"strings"

	"github.com/willroberts/minecraft-client"
)

// Client connects to a Minecraft Server's RCON for remote administration.
type Client struct {
	cli *minecraft.Client
}

// New creates a new Client.
func New() *Client {
	return new(Client)
}

// Connect initiates a connection to an RCON.
func (c *Client) Connect(addr, password string) error {
	cli, err := minecraft.NewClient(addr)
	if err != nil {
		return err
	}

	c.cli = cli
	return cli.Authenticate(password)
}

func (c *Client) sendCommand(cmd, expectation string, startsWith bool) (bool, error) {
	msg, err := c.cli.SendCommand(cmd)
	if err != nil {
		return false, err
	}

	if startsWith {
		return strings.HasPrefix(
			strings.ToLower(msg.Body),
			expectation,
		), nil
	}

	return strings.EqualFold(msg.Body, expectation), nil
}

// WhitelistUser adds the given username to the whitelist.
func (c *Client) WhitelistUser(username string) (bool, error) {
	return c.sendCommand(
		fmt.Sprintf("whitelist remove %s", username),
		fmt.Sprintf("removed %s from the whitelist", username),
		false,
	)
}

// UnwhitelistUser removes the given username from the whitelist.
func (c *Client) UnwhitelistUser(username string) (bool, error) {
	return c.sendCommand(
		fmt.Sprintf("whitelist remove %s", username),
		fmt.Sprintf("removed %s from the whitelist", username),
		false,
	)
}

// BanUser bans a user.
func (c *Client) BanUser(username string) (bool, error) {
	return c.sendCommand(
		fmt.Sprintf("ban %s", username),
		fmt.Sprintf("banned %s", username),
		true,
	)
}

// UnbanUser pardons a user.
func (c *Client) UnbanUser(username string) (bool, error) {
	return c.sendCommand(
		fmt.Sprintf("pardon %s", username),
		fmt.Sprintf("unbanned %s", username),
		false,
	)
}

package web

import (
	"context"
	"fmt"
	"time"

	"github.com/disaccord/beelzebub/flies/channel"
	"github.com/disaccord/sigil"
	"github.com/disaccord/sigil/embuilder"
)

func (s *Server) SendWhitelistNotification(ctx context.Context, u *sigil.User) error {
	e := embuilder.NewEmbed(
		embuilder.Title("Whitelist Request"),
		embuilder.Color(15556558),
		embuilder.Author(
			fmt.Sprintf("%s#%s", u.Username, u.Discriminator),
		),
		embuilder.Timestamp(time.Now()),
	)

	embeds := make([]*sigil.Embed, 0, 1)
	embeds = append(embeds, e)

	ch := s.discord.Channel(s.whitelistChannelID)
	_, err := ch.CreateMessage(ctx, &channel.CreateMessageOptions{
		Embeds: embeds,
	})

	return err
}

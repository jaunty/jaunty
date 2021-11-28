package web

import (
	"context"
	"fmt"
	"time"

	"github.com/disaccord/beelzebub/flies/channel"
	"github.com/disaccord/sigil"
	"github.com/disaccord/sigil/cmbuilder"
	"github.com/disaccord/sigil/embuilder"
	"github.com/jaunty/jaunty/internal/database/models"
)

func (s *Server) SendSiteNotification(ctx context.Context, field ...string) error {
	if len(field)%2 != 0 {
		panic("web: notifier: field values must be an even number")
	}

	fields := make([]*sigil.EmbedField, 0, len(field))
	skip := false
	for i, f := range field {
		if skip {
			skip = false
			continue
		}

		fields = append(fields, &sigil.EmbedField{
			Name:  f,
			Value: field[i+1],
		})
		skip = true
	}

	e := embuilder.NewEmbed(
		embuilder.Title("Site Notification"),
		embuilder.Color(15433780),
		embuilder.Timestamp(time.Now()),
		embuilder.Fields(fields...),
	)

	ch := s.discord.Channel(s.notificationChannelID)
	_, err := ch.CreateMessage(ctx, &channel.CreateMessageOptions{
		Embeds: []*sigil.Embed{e},
	})

	return err
}

func (s *Server) SendWhitelistNotification(ctx context.Context, wh models.Whitelist, du, mcu string) error {
	e := embuilder.NewEmbed(
		embuilder.Title("Whitelist Request"),
		embuilder.Color(15556558),
		embuilder.Timestamp(time.Now()),
		embuilder.Field("Discord Username", du),
		embuilder.Field("Minecraft Username", mcu),
	)

	row := cmbuilder.NewActionRow(
		cmbuilder.Buttons(
			cmbuilder.NewButton(
				cmbuilder.CustomButtonID(fmt.Sprintf("whitelist-approve:%d", wh.ID)),
				cmbuilder.Style(sigil.ButtonStyleSuccess),
				cmbuilder.Label("Approve"),
			),
			cmbuilder.NewButton(
				cmbuilder.CustomButtonID(fmt.Sprintf("whitelist-reject:%d", wh.ID)),
				cmbuilder.Style(sigil.ButtonStyleDanger),
				cmbuilder.Label("Reject"),
			),
		),
	)

	ch := s.discord.Channel(s.whitelistChannelID)
	_, err := ch.CreateMessage(ctx, &channel.CreateMessageOptions{
		Embeds:     []*sigil.Embed{e},
		Components: []*sigil.ActionRow{row},
	})

	return err
}

package web

import (
	"context"
	"time"

	"github.com/disaccord/beelzebub/flies/channel"
	"github.com/disaccord/sigil"
	"github.com/disaccord/sigil/cmbuilder"
	"github.com/disaccord/sigil/embuilder"
)

func (s *Server) SendSiteNotification(ctx context.Context, field ...string) error {
	if len(field)%2 != 0 {
		panic("web: notifier: field values must be an even number")
	}

	fields := make([]*sigil.EmbedField, 0, len(field))
	for i, f := range field {
		if i == len(field)+1 {
			break
		}

		fields = append(fields, &sigil.EmbedField{
			Name:  f,
			Value: field[i+1],
		})
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

func (s *Server) SendWhitelistNotification(ctx context.Context, du, mcu string) error {
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
				cmbuilder.CustomButtonID("whitelist-request-approve"),
				cmbuilder.Style(sigil.ButtonStyleSuccess),
				cmbuilder.Label("Approve"),
			),
			cmbuilder.NewButton(
				cmbuilder.CustomButtonID("whitelist-request-deny"),
				cmbuilder.Style(sigil.ButtonStyleDanger),
				cmbuilder.Label("Deny"),
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

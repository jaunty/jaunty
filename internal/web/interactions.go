package web

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/disaccord/sigil"
	"github.com/disaccord/sigil/embuilder"
	"github.com/holedaemon/web/middleware"
	"github.com/jaunty/jaunty/internal/database/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

type intHandler func(context.Context, *sigil.Interaction, ...string) (*sigil.InteractionResponse, error)

func (s *Server) interactionError(msg string) *sigil.InteractionResponse {
	e := embuilder.NewEmbed(
		embuilder.Title("Uh Oh Sisters"),
		embuilder.Author("Almighty Server"),
		embuilder.Field("Situation Report", msg),
		embuilder.Color(14880305),
	)

	return &sigil.InteractionResponse{
		Type: sigil.InteractionCallbackTypeUpdateMessage,
		Data: &sigil.Message{
			Embeds: []*sigil.Embed{e},
		},
	}
}

func (s *Server) handlerApproveWhitelist(ctx context.Context, i *sigil.Interaction, args ...string) (*sigil.InteractionResponse, error) {
	if len(args) == 0 {
		return s.interactionError("Request ID was not sent."), nil
	}

	tx := middleware.TxFromContext(ctx)

	reqID := args[0]
	wr, err := models.Whitelists(qm.Where("id = ?", reqID)).One(ctx, tx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return s.interactionError("Request does not exist."), nil
		}
		ctxlog.Error(ctx, "error querying whitelist request", zap.Error(err))
		return s.interactionError("Unable to query whitelist"), nil
	}

	wr.WhitelistStatus = models.WhitelistStatusApproved
	if err := wr.Update(ctx, tx, boil.Infer()); err != nil {
		ctxlog.Error(ctx, "error updating request", zap.Error(err))
		return s.interactionError("Unable to update request"), nil
	}

	if err := tx.Commit(); err != nil {
		ctxlog.Error(ctx, "error committing transaction", zap.Error(err))
		return s.interactionError("Unable to commit transaction"), nil
	}

	ru, err := s.mojang.FetchUsernameByUUID(ctx, wr.UUID)
	if err != nil {
		ctxlog.Error(ctx, "error fetching minecraft username by uuid", zap.Error(err))
		return s.interactionError("Unable to fetch Minecraft username"), nil
	}

	whitelisted, err := s.rcon.WhitelistUser(ru)
	if err != nil {
		ctxlog.Error(ctx, "unable to whitelist user", zap.Error(err))
		return s.interactionError("Error whitelisting user"), nil
	}

	if !whitelisted {
		return s.interactionError("Whitelisting failed according to flimsy checking"), nil
	}

	g := s.discord.Guild(s.guildID)
	err = g.AddMemberRole(ctx, wr.SF, s.approvedRoleID, "Member has been approved.")
	if err != nil {
		ctxlog.Error(ctx, "error adding role to member", zap.Error(err), zap.String("sf", wr.SF))
		return s.interactionError("Unable to add role to member"), nil
	}

	user, err := s.fetchDiscordUser(ctx, wr.SF)
	if err != nil {
		ctxlog.Error(ctx, "error fetching user", zap.Error(err))
		return s.interactionError("Unable to fetch user information"), nil
	}

	e := embuilder.NewEmbed(
		embuilder.Title("Request Approved"),
		embuilder.Field("User", fmt.Sprintf("%s#%s", user.Username, user.Discriminator)),
		embuilder.Color(1680963),
	)

	return &sigil.InteractionResponse{
		Type: sigil.InteractionCallbackTypeUpdateMessage,
		Data: &sigil.Message{
			Embeds: []*sigil.Embed{e},
		},
	}, nil
}

func (s *Server) handlerRejectWhitelist(ctx context.Context, i *sigil.Interaction, args ...string) (*sigil.InteractionResponse, error) {
	if len(args) == 0 {
		return s.interactionError("Request ID was not sent."), nil
	}

	tx := middleware.TxFromContext(ctx)

	reqID := args[0]
	wr, err := models.Whitelists(qm.Where("id = ?", reqID)).One(ctx, tx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return s.interactionError("Request does not exist"), nil
		}

		ctxlog.Error(ctx, "error querying request", zap.Error(err))
		return s.interactionError("Unable to query request."), nil
	}

	wr.WhitelistStatus = models.WhitelistStatusRejected
	if err := wr.Update(ctx, tx, boil.Infer()); err != nil {
		ctxlog.Error(ctx, "unable to update whitelist request", zap.Error(err))
		return s.interactionError("Unable to update request"), nil
	}

	if err := tx.Commit(); err != nil {
		ctxlog.Error(ctx, "error committing transaction", zap.Error(err))
		return s.interactionError("Unable to commit transaction"), nil
	}

	ru, err := s.fetchMojangUsernameByUUID(ctx, wr.UUID)
	if err != nil {
		ctxlog.Error(ctx, "error fetching username by UUID", zap.String("UUID", wr.UUID), zap.Error(err))
		return s.interactionError("Unable to fetch username by UUID"), nil
	}

	dewhitelisted, err := s.rcon.UnwhitelistUser(ru)
	if err != nil {
		ctxlog.Error(ctx, "error unwhitelisting user", zap.Error(err), zap.String("user", ru))
		return s.interactionError("Unable to remove user from the whitelist"), nil
	}

	if !dewhitelisted {
		return s.interactionError(fmt.Sprintf("Removing %s from the whitelist failed according to flimsy checking", ru)), nil
	}

	user, err := s.fetchDiscordUser(ctx, wr.SF)
	if err != nil {
		ctxlog.Error(ctx, "error fetching Discord user", zap.Error(err))
		return s.interactionError("Unable to fetch user data"), nil
	}

	e := embuilder.NewEmbed(
		embuilder.Title("Request Rejected"),
		embuilder.Field("User", fmt.Sprintf("%s#%s", user.Username, user.Discriminator)),
		embuilder.Color(14880305),
	)

	return &sigil.InteractionResponse{
		Type: sigil.InteractionCallbackTypeUpdateMessage,
		Data: &sigil.Message{
			Embeds: []*sigil.Embed{e},
		},
	}, nil
}

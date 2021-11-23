package web

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/disaccord/sigil"
	"github.com/disaccord/sigil/embuilder"
	"github.com/holedaemon/web/middleware"
	"github.com/jaunty/jaunty/internal/database/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

func (s *Server) updateMessage(msg string) *sigil.InteractionResponse {
	return &sigil.InteractionResponse{
		Type: sigil.InteractionCallbackTypeChannelMessageWithSource,
		Data: &sigil.Message{
			Content: msg,
		},
	}
}

func (s *Server) modifyWhitelist(ctx context.Context, sf, kind string) (*sigil.InteractionResponse, error) {
	switch kind {
	case models.WhitelistStatusApproved, models.WhitelistStatusPending, models.WhitelistStatusCancelled, models.WhitelistStatusRejected:
		tx := middleware.TxFromContext(ctx)
		w, err := models.Whitelists(
			qm.Where("sf = ?", sf),
		).One(ctx, tx)
		if err != nil {
			ctxlog.Error(ctx, "error querying whitelist from db", zap.Error(err))
			un, err := s.mojang.FetchUsernameByUUID(ctx, w.UUID)
			if err != nil {
				return nil, err
			}

			return s.updateMessage(fmt.Sprintf("Error updating whitelist request for %s :(", un)), nil
		}

		un, err := s.mojang.FetchUsernameByUUID(ctx, w.UUID)
		if err != nil {
			return nil, err
		}

		w.WhitelistStatus = kind
		if err := w.Update(ctx, tx, boil.Infer()); err != nil {
			ctxlog.Error(ctx, "error updating whitelist", zap.Error(err))

			return s.updateMessage(fmt.Sprintf("Error updating whitelist request for %s :(", un)), nil
		}

		ue := embuilder.NewEmbed(
			embuilder.Title("Whitelist Request"),
			embuilder.Field("Minecraft User", un),
			embuilder.Field("Status", strings.ToTitle(kind)),
		)

		switch w.WhitelistStatus {
		case models.WhitelistStatusApproved:
			ue.Color = 1754923
		case models.WhitelistStatusCancelled:
			ue.Color = 15263020
		case models.WhitelistStatusRejected:
			ue.Color = 13048378
		case models.WhitelistStatusPending:
			ue.Color = 13078554
		}

		resp := &sigil.InteractionResponse{
			Type: sigil.InteractionCallbackTypeChannelMessageWithSource,
			Data: &sigil.Message{
				Embeds: []*sigil.Embed{ue},
			},
		}
		return resp, nil
	default:
		return s.updateMessage("Unknown status type given ???"), nil
	}
}

var errMemberIDMissing = errors.New("web: member ID not given in interaction data")

func (s *Server) ApproveWhitelistHook(ctx context.Context, i *sigil.Interaction) (*sigil.InteractionResponse, error) {
	if i.Member.User.ID == "" {
		return nil, errMemberIDMissing
	}

	return s.modifyWhitelist(ctx, i.Member.User.ID, models.WhitelistStatusApproved)
}

func (s *Server) RejectWhitelistHook(ctx context.Context, i *sigil.Interaction) (*sigil.InteractionResponse, error) {
	if i.Member.User.ID == "" {
		return nil, errMemberIDMissing
	}

	return s.modifyWhitelist(ctx, i.Member.User.ID, models.WhitelistStatusRejected)
}

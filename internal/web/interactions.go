package web

import (
	"context"

	"github.com/disaccord/sigil"
	"github.com/holedaemon/web/middleware"
)

func (s *Server) ApproveWhitelistHook(ctx context.Context, i *sigil.Interaction) (*sigil.InteractionResponse, error) {
	tx := middleware.TxFromContext(ctx)

	if i.Member.User.ID == "" {
		panic("web: ApproveWhitelistHook received a blank member id")
	}
}

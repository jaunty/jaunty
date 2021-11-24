package web

import (
	"context"
	"fmt"

	"github.com/disaccord/sigil"
	"github.com/holedaemon/web/middleware"
	"github.com/jaunty/jaunty/internal/database/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

func (s *Server) ApproveWhitelistHook(ctx context.Context, i *sigil.Interaction) (*sigil.InteractionResponse, error) {
	tx := middleware.TxFromContext(ctx)

	if i.Member.User.ID == "" {
		panic("web: ApproveWhitelistHook received a blank member id")
	}

	sf := i.Member.User.ID
	fullName := fmt.Sprintf("%s#%s", i.Member.User.Username, i.Member.User.Discriminator)

	count, err := models.Whitelists(qm.Where("sf = ?", sf)).Count(ctx, tx)
	if err != nil {
		ctxlog.Error(ctx, "error counting whitelist requests for sf", zap.Error(err), zap.String("sf", sf))
		return followupError(ctx, "Unable to count whitelist requests for %s", fullName), nil
	}

	if count == 0 {
		return followupError(ctx, "%s does not have a whitelist request in the database.", fullName), nil
	}

	if count > 1 {

	}
}

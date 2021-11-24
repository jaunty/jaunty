package modelsx

import (
	"context"
	"strings"

	"github.com/jaunty/jaunty/internal/database/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// ModifyWhitelistStatus modifies the status of a whitelist request.
func ModifyWhitelistStatus(ctx context.Context, exec boil.ContextExecutor, sf, uuid string, status string) (*models.Whitelist, error) {
	w, err := models.Whitelists(
		qm.Where("sf = ?", sf),
		qm.Where("uuid = ?", uuid),
	).One(ctx, exec)
	if err != nil {
		return nil, err
	}

	if strings.EqualFold(w.WhitelistStatus, status) {
		return w, nil
	}

	w.WhitelistStatus = status
	if err := w.Update(ctx, exec, boil.Infer()); err != nil {
		return nil, err
	}

	return w, nil
}

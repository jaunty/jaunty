package modelsx

import (
	"context"

	"github.com/jaunty/jaunty/internal/database/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// DeleteUser deletes a User's account.
func DeleteUser(ctx context.Context, exec boil.ContextTransactor, sf string) error {
	user, err := models.Users(
		qm.Where("sf = ?", sf),
	).One(ctx, exec)
	if err != nil {
		return err
	}

	return user.Delete(ctx, exec)
}

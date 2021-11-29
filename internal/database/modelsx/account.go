package modelsx

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// DeleteAccount deletes a User, their requests, and Discord OAuth2 tokens.
func DeleteAccount(ctx context.Context, exec boil.ContextTransactor, sf string) error {
	if err := DeleteUser(ctx, exec, sf); err != nil {
		return err
	}

	if err := DeleteRequests(ctx, exec, sf); err != nil {
		return err
	}

	if err := DeleteTokens(ctx, exec, sf); err != nil {
		return err
	}

	return nil
}

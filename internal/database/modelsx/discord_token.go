package modelsx

import (
	"context"

	"github.com/jaunty/jaunty/internal/database/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/oauth2"
)

var tokenConflictColumns = []string{"sf"}

var tokenUpdateColumns = boil.Whitelist("access_token", "token_type", "refresh_token", "expiry")

// UpsertToken upserts a Discord token.
func UpsertToken(ctx context.Context, exec boil.ContextExecutor, sf string, tok *oauth2.Token) error {
	dt := &models.DiscordToken{
		SF:           sf,
		AccessToken:  tok.AccessToken,
		TokenType:    tok.TokenType,
		Expiry:       tok.Expiry,
		RefreshToken: tok.RefreshToken,
	}

	return dt.Upsert(ctx, exec, true, tokenConflictColumns, tokenUpdateColumns, boil.Infer())
}

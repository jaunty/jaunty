// Package dbx provides extensions for the database/sql package.
package dbx

import (
	"context"
	"database/sql"
	"time"

	"github.com/jaunty/jaunty/internal/pkg/dbx/driver"
	"github.com/zikaeroh/ctxlog"
	"go.uber.org/zap"
)

const (
	MaxAttempts = 10
	Timeout     = 3 * time.Second
)

// Open opens a connection to a SQL DB and pings the server to verify the connection is alive.
func Open(ctx context.Context, dsn string) (*sql.DB, error) {
	var err error

	db, err := sql.Open(driver.Driver, dsn)
	if err != nil {
		return nil, err
	}

	connected := false

	for i := 0; i < MaxAttempts; i++ {
		err = db.PingContext(ctx)
		if err != nil {
			ctxlog.Error(ctx, "error connecting to postgres db", zap.Int("attempt", i), zap.Int("max_attempts", MaxAttempts), zap.Duration("retry", Timeout))
			time.Sleep(Timeout)
			continue
		}

		connected = true
	}

	if !connected {
		return nil, err
	}

	return db, nil
}

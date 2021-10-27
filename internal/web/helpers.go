package web

import (
	"context"
	"database/sql"
)

func txFromCtx(ctx context.Context) *sql.Tx {
	tr := ctx.Value(txKey{})
	if tr == nil {
		panic("web: tx not present in ctx")
	}

	tx, ok := tr.(*sql.Tx)
	if !ok {
		panic("web: tx is not of type *sql.Tx")
	}

	return tx
}

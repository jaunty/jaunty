package redisx

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"
)

func TestRedis(t *testing.T) {
	ctx := context.Background()

	r := Open("localhost:6379")
	err := r.Ping(ctx)
	assert.NilError(t, err, "pinging")

	err = r.SetString(ctx, "hello", "there")
	assert.NilError(t, err, "setting string")

	str, err := r.FetchString(ctx, "hello")
	assert.NilError(t, err, "getting string")

	assert.Assert(t, str == "there")
}

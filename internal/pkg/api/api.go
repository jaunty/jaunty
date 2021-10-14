// Package api provides a base API client for child packages.
package api

import (
	"context"
	"net/http"

	"github.com/jaunty/jaunty/internal/pkg/httpx"
)

// Client is an HTTP Client intended to consume REST APIs,
// with a caching backend.
type Client interface {
	Do(context.Context, string, ...httpx.RequestOption) (*http.Response, error)
	FetchValue(context.Context, string) (interface{}, error)
	SetValue(context.Context, string, interface{}) error
}

// Package httpx provides extensions to the http package.
package httpx

import (
	"errors"
	"net/http"
	"time"
)

var ErrStatusCode = errors.New("httpx: unexpected status code returned")

func init() {
	http.DefaultClient = &http.Client{
		Timeout: time.Second * 10,
	}
}

// Package httpx provides extensions to the http package.
package httpx

import (
	"net/http"
	"time"
)

func init() {
	http.DefaultClient = &http.Client{
		Timeout: time.Second * 10,
	}
}

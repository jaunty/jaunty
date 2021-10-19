package httpx

import (
	"io"
	"net/http"
	"net/url"
)

type requestOptions struct {
	Headers http.Header
	Query   url.Values
	Body    io.Reader
}

type RequestOption func(*requestOptions)

func WithQuery(q url.Values) RequestOption {
	return func(ro *requestOptions) {
		ro.Query = q
	}
}

func WithBody(r io.Reader) RequestOption {
	return func(ro *requestOptions) {
		ro.Body = r
	}
}

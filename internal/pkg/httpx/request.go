package httpx

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type requestOptions struct {
	Headers http.Header
	Query   url.Values
	Body    io.Reader
	Method  string
}

type RequestOption func(*requestOptions)

// WithQuery sets a request's query parameters.
func WithQuery(q url.Values) RequestOption {
	return func(ro *requestOptions) {
		ro.Query = q
	}
}

// WithBody sets a request's body.
func WithBody(r io.Reader) RequestOption {
	return func(ro *requestOptions) {
		ro.Body = r
	}
}

// WithHeaders sets a request's headers.
func WithHeaders(h http.Header) RequestOption {
	return func(ro *requestOptions) {
		ro.Headers = h
	}
}

// WithMethod sets a request's method.
func WithMethod(m string) RequestOption {
	return func(ro *requestOptions) {
		ro.Method = m
	}
}

// NewRequest creates a new HTTP request with the given options.
// If opts.Method is blank, the method is automatically set to GET.
func NewRequest(ctx context.Context, url string, opts ...RequestOption) (*http.Request, error) {
	ro := new(requestOptions)
	for _, o := range opts {
		o(ro)
	}

	if ro.Method == "" {
		ro.Method = http.MethodGet
	}

	req, err := http.NewRequestWithContext(ctx, ro.Method, url, ro.Body)
	if err != nil {
		return nil, err
	}

	if ro.Headers != nil {
		req.Header = ro.Headers
	}

	if ro.Query != nil {
		req.URL.RawQuery = ro.Query.Encode()
	}

	return req, nil
}

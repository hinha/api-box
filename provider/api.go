package provider

import (
	"context"
	"net/http"
)

// APIContext used by API handler to modify it's request
type APIContext interface {
	// Request returns `*http.Request`.
	Request() *http.Request

	// RealIP returns the client's network address based on `X-Forwarded-For`
	// or `X-Real-IP` request header.
	// The behavior can be configured using `Echo#IPExtractor`.
	RealIP() string

	// Path returns the registered path for the handler.
	Path() string

	// JSON sends a JSON response with status code.
	JSON(code int, i interface{}) error
}

// APIHandler handling api request from client
type APIHandler interface {
	Handle(context APIContext)
	Method() string
	Path() string
}

// APIEngine ...
type APIEngine interface {
	Run() error
	InjectAPI(handler APIHandler)
	Shutdown(ctx context.Context) error
}

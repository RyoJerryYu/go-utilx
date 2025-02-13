package httpx

import "net/http"

// Client interface abstracts the complexity of http.Client,
// limiting XClient to only use the Do method when calling http.Client.
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

// ClientDecorator is a function type that wraps a Client with additional functionality.
// It can be used to add middleware-like behavior to the client.
type ClientDecorator func(Client) Client

// ClientFunc is an adapter to allow the use of ordinary functions as HTTP clients.
// If f is a function with the appropriate signature, ClientFunc(f) is a Client that calls f.
type ClientFunc func(req *http.Request) (*http.Response, error)

func (f ClientFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

var _ Client = ClientFunc(nil)

type ClientOption func(*http.Client) *http.Client

// NewClient creates a new http.Client with the given options.
// It automatically adds OpenTelemetry instrumentation as the first option.
//
// Example:
//
//	client := NewClient(
//	    WithBearerAuth("token"),
//	    WithTimeout(5 * time.Second),
//	)
func NewClient(opts ...ClientOption) *http.Client {
	cli := &http.Client{}
	opts = append([]ClientOption{WithOtel()}, opts...)
	for _, opt := range opts {
		cli = opt(cli)
	}
	return cli
}

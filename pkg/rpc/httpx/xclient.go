package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// XClient is an enhanced HTTP client that wraps http.Client with additional functionality.
// It provides a more convenient API and supports middleware-like decorators.
type XClient struct {
	inner Client
}

// NewXClient creates a new XClient with default http.Client and options.
// By default, it includes OpenTelemetry instrumentation and error handling for non-2xx responses.
//
// Example:
//
//	client := NewXClient(
//	    WithBearerAuth("token"),
//	    WithUnwrapTransportError(),
//	)
func NewXClient(opts ...XClientOption) *XClient {
	return NewXClientFromHttp(&http.Client{}, opts...)
}

// NewXClientFromHttp creates a new XClient from an existing http.Client.
// It applies the given options to the client while preserving its original configuration.
//
// The default options (OpenTelemetry and error handling) are still applied unless
// WithoutDefaultOption() is used.
func NewXClientFromHttp(httpCli *http.Client, opts ...XClientOption) *XClient {
	c := xClientConfig{}
	for _, opt := range opts {
		opt.apply(&c)
	}

	if !c.withoutDefaultOption {
		c.clientOptions = append(c.clientOptions, WithOtel())
		c.clientDecorators = append(c.clientDecorators, WithReturnErrorIfNot2xx())
	}

	for _, opt := range c.clientOptions {
		httpCli = opt(httpCli)
	}

	var cliCore Client = httpCli
	for _, opt := range c.clientDecorators {
		cliCore = opt(cliCore)
	}
	return &XClient{
		inner: cliCore,
	}
}

func NewXClientFromInterface(httpCli Client, opts ...XClientOption) *XClient {
	c := xClientConfig{}
	for _, opt := range opts {
		opt.apply(&c)
	}

	if !c.withoutDefaultOption {
		c.clientOptions = append(c.clientOptions, WithOtel())
		c.clientDecorators = append(c.clientDecorators, WithReturnErrorIfNot2xx())
	}

	var cliCore Client = httpCli
	for _, opt := range c.clientDecorators {
		cliCore = opt(cliCore)
	}
	return &XClient{
		inner: cliCore,
	}
}

func (c *XClient) Do(req *http.Request, opts ...XRequestOption) (*http.Response, error) {
	cfg := xRequestOpts{}
	for _, opt := range opts {
		opt(&cfg)
	}
	cfg.applyRequest(req)
	return c.inner.Do(req)
}

func (c *XClient) Head(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *XClient) Get(ctx context.Context, url string, opts ...XRequestOption) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req, opts...)
}

func (c *XClient) Post(ctx context.Context, url string, contentType string, body io.Reader, opts ...XRequestOption) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req, opts...)
}

func (c *XClient) GetBytes(ctx context.Context, url string, opts ...XRequestOption) ([]byte, error) {
	resp, err := c.Get(ctx, url, opts...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// GetJSON performs a GET request and unmarshals the response JSON into the provided interface.
// The response parameter must be a pointer to a type that can be unmarshaled from JSON.
//
// Example:
//
//	var response struct {
//	    Name string `json:"name"`
//	    Age  int    `json:"age"`
//	}
//	err := client.GetJSON(ctx, "https://api.example.com/user/1", &response)
func (c *XClient) GetJSON(ctx context.Context, url string, response any, opts ...XRequestOption) error {
	resp, err := c.Get(ctx, url, opts...)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(response)
}

func (c *XClient) PostForm(ctx context.Context, url string, data url.Values, opts ...XRequestOption) (*http.Response, error) {
	return c.Post(ctx, url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()), opts...)
}

// PostJSON performs a POST request with a JSON body and unmarshals the response into
// the provided interface. Both request and response parameters must be compatible with
// JSON marshaling/unmarshaling.
//
// Example:
//
//	request := struct {
//	    Name string `json:"name"`
//	}{Name: "John"}
//	var response struct {
//	    ID   int    `json:"id"`
//	    Name string `json:"name"`
//	}
//	err := client.PostJSON(ctx, "https://api.example.com/users", request, &response)
func (c *XClient) PostJSON(ctx context.Context, url string, request any, response any, opts ...XRequestOption) error {
	buf := bytes.NewBuffer(nil)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return err
	}

	resp, err := c.Post(ctx, url, "application/json", buf, opts...)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(response)
}

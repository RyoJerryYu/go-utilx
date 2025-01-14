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

type XClient struct {
	inner Client
}

func NewXClient(opts ...XClientOption) *XClient {
	return NewXClientFromHttp(&http.Client{}, opts...)
}

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

func (c *XClient) Do(req *http.Request) (*http.Response, error) {
	return c.inner.Do(req)
}

func (c *XClient) Head(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *XClient) Get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *XClient) Post(ctx context.Context, url string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req)

}

func (c *XClient) PostForm(ctx context.Context, url string, data url.Values) (*http.Response, error) {
	return c.Post(ctx, url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

func (c *XClient) PostJSON(ctx context.Context, url string, request any, response any) error {
	buf := bytes.NewBuffer(nil)
	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return err
	}

	resp, err := c.Post(ctx, url, "application/json", buf)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(response)
}

func (c *XClient) XGet(ctx context.Context, url string, opts ...XRequestOption) ([]byte, error) {
	cfg := xRequestOpts{}
	for _, opt := range opts {
		opt(&cfg)
	}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	cfg.applyRequest(req)

	req = req.WithContext(ctx)

	// Send the request and handle the response
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cfg.applyResponseBody(respBody)

	return respBody, nil
}

func (c *XClient) XPost(ctx context.Context, url string, contentType string, body io.Reader, opts ...XRequestOption) ([]byte, error) {
	cfg := xRequestOpts{}
	for _, opt := range opts {
		opt(&cfg)
	}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)

	cfg.applyRequest(req)

	req = req.WithContext(ctx)

	// Send the request and handle the response
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cfg.applyResponseBody(respBody)

	return respBody, nil
}

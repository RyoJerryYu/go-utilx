package httpx

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

/**
XClientOption is a type that allow additional options for http client,
that would apply to all requests.
*/

type xClientConfig struct {
	withoutDefaultOption bool
	clientOptions        []ClientOption    // clientOptions 主要是用来装饰 http.Client 内部，如 Transport
	clientDecorators     []ClientDecorator // clientDecorators 主要是通过 Do 方法用来装饰 http.Client 外部
}

type XClientOption interface {
	apply(*xClientConfig)
}

// a ClientOption is a XClientOption.
func (f ClientOption) apply(config *xClientConfig) {
	config.clientOptions = append(config.clientOptions, f)
}

// a ClientDecorator is a XClientOption.
func (f ClientDecorator) apply(config *xClientConfig) {
	config.clientDecorators = append(config.clientDecorators, f)
}

type XClientOptionFunc func(*xClientConfig)

// a XClientOptionFunc is a XClientOption.
func (f XClientOptionFunc) apply(config *xClientConfig) {
	f(config)
}

func WithoutDefaultOption() XClientOption {
	return XClientOptionFunc(func(scc *xClientConfig) {
		scc.withoutDefaultOption = true
	})
}

func WithOtel(opts ...otelhttp.Option) ClientOption {
	return func(httpCli *http.Client) *http.Client {
		defaultOpts := []otelhttp.Option{
			otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
				return r.Method + " " + r.URL.Host
			}),
		}

		opts = append(defaultOpts, opts...)

		httpCli.Transport = otelhttp.NewTransport(httpCli.Transport, opts...)
		return httpCli
	}
}

func WithBearerAuth(token string) ClientOption {
	return func(cli *http.Client) *http.Client {
		cli.Transport = NewAuthTransport(cli.Transport, "Bearer", token)
		return cli
	}
}

func WithUnwrapTransportError() ClientDecorator {
	return func(inner Client) Client {
		return ClientFunc(func(req *http.Request) (*http.Response, error) {
			resp, err := inner.Do(req)
			if unwrapedErr := errors.Unwrap(err); unwrapedErr != nil {
				return resp, unwrapedErr
			}

			return resp, err
		})
	}
}

type XError struct {
	Response *http.Response
	Method   string
	Code     int
	Body     []byte
}

func (e *XError) Error() string {
	return fmt.Sprintf("httpx %s error %d: %s", e.Method, e.Code, e.Body)
}

func WithReturnErrorIfNot2xx() ClientDecorator {
	return func(inner Client) Client {
		return ClientFunc(func(req *http.Request) (*http.Response, error) {
			resp, err := inner.Do(req)
			if err != nil {
				return resp, err
			}

			if code := resp.StatusCode; code < 200 || code >= 300 {
				defer resp.Body.Close() // 约定， Do 方法返回 error 时 Response.Body 一定要关闭

				respBody, err := io.ReadAll(resp.Body)
				if err != nil {
					return resp, err
				}

				return resp, &XError{
					Response: resp,
					Method:   resp.Request.Method,
					Code:     code,
					Body:     respBody,
				}
			}

			return resp, nil
		})
	}
}

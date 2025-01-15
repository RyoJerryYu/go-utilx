package httpx

import "net/http"

// 屏蔽 http.Client 复杂性，限制 XClient 最终都只能从 Do 方法调用 http.Client
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type ClientDecorator func(Client) Client

type ClientFunc func(req *http.Request) (*http.Response, error)

func (f ClientFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

var _ Client = ClientFunc(nil)

type ClientOption func(*http.Client) *http.Client

func NewClient(opts ...ClientOption) *http.Client {
	cli := &http.Client{}
	opts = append([]ClientOption{WithOtel()}, opts...)
	for _, opt := range opts {
		cli = opt(cli)
	}
	return cli
}

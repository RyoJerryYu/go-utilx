package httpx

import (
	"net/http"
	"net/url"
)

/**
XRequestOption is a type that allow additional options for one request.
*/

type xRequestOpts struct {
	applyReqs []func(*http.Request)
}

type XRequestOption func(*xRequestOpts)

func WithQuerys(queryParams url.Values) XRequestOption {
	return func(sgc *xRequestOpts) {
		sgc.addApplyRequest(func(r *http.Request) {
			r.URL.RawQuery = queryParams.Encode()
		})
	}
}

func WithQuery(key string, value string) XRequestOption {
	return func(sgc *xRequestOpts) {
		sgc.addApplyRequest(func(r *http.Request) {
			q := r.URL.Query()
			q.Add(key, value)
			r.URL.RawQuery = q.Encode()
		})
	}
}

func WithHeaders(headers map[string]string) XRequestOption {
	return func(sgc *xRequestOpts) {
		sgc.addApplyRequest(func(r *http.Request) {
			for k, v := range headers {
				r.Header.Set(k, v)
			}
		})
	}
}

func WithHeader(key string, value string) XRequestOption {
	return func(sgc *xRequestOpts) {
		sgc.addApplyRequest(func(r *http.Request) {
			r.Header.Set(key, value)
		})
	}
}

func (c *xRequestOpts) addApplyRequest(apply func(*http.Request)) {
	c.applyReqs = append(c.applyReqs, apply)
}

func (c *xRequestOpts) applyRequest(req *http.Request) {
	for _, apply := range c.applyReqs {
		apply(req)
	}
}

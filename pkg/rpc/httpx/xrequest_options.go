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

// WithQuerys adds multiple query parameters to the request URL.
// If the URL already has query parameters, they will be preserved.
//
// Example:
//
//	params := url.Values{}
//	params.Set("page", "1")
//	params.Set("limit", "10")
//	resp, err := client.Get(ctx, "https://api.example.com/users",
//	    WithQuerys(params))
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

// WithHeaders adds multiple headers to the request.
// If a header already exists, its value will be replaced.
//
// Example:
//
//	resp, err := client.Get(ctx, "https://api.example.com/users",
//	    WithHeaders(map[string]string{
//	        "X-API-Key": "secret",
//	        "Accept": "application/json",
//	    }))
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

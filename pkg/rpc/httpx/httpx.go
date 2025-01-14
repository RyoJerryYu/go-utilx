package httpx

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

var defaultCli = NewXClient()

// // MustNewJsonRpcCliWithHTTP 返回 HTTP transport 的 client，不需要 close
// func MustNewJsonRpcCliWithHTTP(endpoint string) *rpc.Client {
// 	cli, err := rpc.DialHTTPWithClient(endpoint, NewHttpClient())
// 	if err != nil {
// 		panic(err)
// 	}

// 	return cli
// }

func PostJSON(ctx context.Context, url string, request any, response any) error {
	return defaultCli.PostJSON(ctx, url, request, response)
}

func XGet(ctx context.Context, url string, opts ...XRequestOption) ([]byte, error) {
	return defaultCli.XGet(ctx, url, opts...)
}

func BuildURL(baseURL string, queryParams map[string]string) (string, error) {
	// Parse the base URL to get the URL object
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	// Add the query parameters to the URL object
	q := u.Query()
	for k, v := range queryParams {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	// Convert the URL object to a string
	return strings.TrimRight(u.String(), "?"), nil
}

func HeaderToMap(header http.Header) map[string]string {
	headers := make(map[string]string)
	for k, v := range header {
		headers[k] = v[0]
	}
	return headers
}

func MapToHeader(headers map[string]string) http.Header {
	header := make(http.Header)
	for k, v := range headers {
		header.Set(k, v)
	}
	return header
}

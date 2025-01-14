package httpx

import "net/http"

type AuthTransport struct {
	Base      http.RoundTripper
	AuthType  string
	AuthToken string
}

func NewAuthTransport(base http.RoundTripper, authType string, authToken string) http.RoundTripper {
	return &AuthTransport{
		Base:      base,
		AuthType:  authType,
		AuthToken: authToken,
	}
}

func (t *AuthTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}

func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", t.AuthType+" "+t.AuthToken)
	return t.base().RoundTrip(req)
}

package oauth2x

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/RyoJerryYu/go-utilx/pkg/rpc/httpx"
	"golang.org/x/oauth2"
)

/**
单测复制自： https://github.com/golang/oauth2/blob/master/transport_test.go
用来保证 OAuth2Transport 必定是 oauth2.Transport 的完全兼容
*/

type tokenSource struct{ token *oauth2.Token }

func (t *tokenSource) Token() (*oauth2.Token, error) {
	return t.token, nil
}

func TestTransportNilTokenSource(t *testing.T) {
	tr := &OAuth2Core{}
	server := newMockServer(func(w http.ResponseWriter, r *http.Request) {})
	defer server.Close()
	client := wrapOAuth2Core(tr)
	resp, err := client.Get(context.Background(), server.URL)
	if err == nil {
		t.Errorf("got no errors, want an error with nil token source")
	}
	if resp != nil {
		t.Errorf("Response = %v; want nil", resp)
	}
}

type readCloseCounter struct {
	CloseCount int
	ReadErr    error
}

func (r *readCloseCounter) Read(b []byte) (int, error) {
	return 0, r.ReadErr
}

func (r *readCloseCounter) Close() error {
	r.CloseCount++
	return nil
}

func TestTransportCloseRequestBody(t *testing.T) {
	tr := &OAuth2Core{}
	server := newMockServer(func(w http.ResponseWriter, r *http.Request) {})
	defer server.Close()
	client := wrapOAuth2Core(tr)
	body := &readCloseCounter{
		ReadErr: errors.New("readCloseCounter.Read not implemented"),
	}
	resp, err := client.Post(context.Background(), server.URL, "application/json", body)
	if err == nil {
		t.Errorf("got no errors, want an error with nil token source")
	}
	if resp != nil {
		t.Errorf("Response = %v; want nil", resp)
	}
	if expected := 1; body.CloseCount != expected {
		t.Errorf("Body was closed %d times, expected %d", body.CloseCount, expected)
	}
}

func TestTransportCloseRequestBodySuccess(t *testing.T) {
	tr := &OAuth2Core{
		Source: oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: "abc",
		}),
	}
	server := newMockServer(func(w http.ResponseWriter, r *http.Request) {})
	defer server.Close()
	client := wrapOAuth2Core(tr)
	body := &readCloseCounter{
		ReadErr: io.EOF,
	}
	resp, err := client.Post(context.Background(), server.URL, "application/json", body)
	if err != nil {
		t.Errorf("got error %v; expected none", err)
	}
	if resp == nil {
		t.Errorf("Response is nil; expected non-nil")
	}
	if expected := 1; body.CloseCount != expected {
		t.Errorf("Body was closed %d times, expected %d", body.CloseCount, expected)
	}
}

func TestTransportTokenSource(t *testing.T) {
	ts := &tokenSource{
		token: &oauth2.Token{
			AccessToken: "abc",
		},
	}
	tr := &OAuth2Core{
		Source: ts,
	}
	server := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Header.Get("Authorization"), "Bearer abc"; got != want {
			t.Errorf("Authorization header = %q; want %q", got, want)
		}
	})
	defer server.Close()
	client := wrapOAuth2Core(tr)
	res, err := client.Get(context.Background(), server.URL)
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()
}

// Test for case-sensitive token types, per https://github.com/golang/oauth2/issues/113
func TestTransportTokenSourceTypes(t *testing.T) {
	const val = "abc"
	tests := []struct {
		key  string
		val  string
		want string
	}{
		{key: "bearer", val: val, want: "Bearer abc"},
		{key: "mac", val: val, want: "MAC abc"},
		{key: "basic", val: val, want: "Basic abc"},
	}
	for _, tc := range tests {
		ts := &tokenSource{
			token: &oauth2.Token{
				AccessToken: tc.val,
				TokenType:   tc.key,
			},
		}
		tr := &OAuth2Core{
			Source: ts,
		}
		server := newMockServer(func(w http.ResponseWriter, r *http.Request) {
			if got, want := r.Header.Get("Authorization"), tc.want; got != want {
				t.Errorf("Authorization header (%q) = %q; want %q", val, got, want)
			}
		})
		defer server.Close()
		client := wrapOAuth2Core(tr)
		res, err := client.Get(context.Background(), server.URL)
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()
	}
}

func TestTokenValidNoAccessToken(t *testing.T) {
	token := &oauth2.Token{}
	if token.Valid() {
		t.Errorf("got valid with no access token; want invalid")
	}
}

func TestExpiredWithExpiry(t *testing.T) {
	token := &oauth2.Token{
		Expiry: time.Now().Add(-5 * time.Hour),
	}
	if token.Valid() {
		t.Errorf("got valid with expired token; want invalid")
	}
}

func newMockServer(handler func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(handler))
}

func wrapOAuth2Core(tr *OAuth2Core) *httpx.XClient {
	return httpx.NewXClientFromInterface(tr, httpx.WithoutDefaultOption())
}

package oauth2x

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/RyoJerryYu/go-utilx/pkg/rpc/httpx"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestTransportTriggerOnRrefreshTokenChange(t *testing.T) {
	ctx := context.Background()
	currentToken := oauth2.Token{
		RefreshToken: "refresh_token",
		AccessToken:  "access_token",
		TokenType:    "Bearer",
		Expiry:       time.Now().Add(-time.Hour),
	}

	newToken := oauth2.Token{
		RefreshToken: "new_refresh_token",
		AccessToken:  "access_token",
		TokenType:    "Bearer",
		Expiry:       time.Now().Add(time.Hour),
	}

	server := newMockServer(func(w http.ResponseWriter, r *http.Request) {})
	defer server.Close()

	triggered := false

	client := httpx.NewXClient(WithOAuth2Http(
		ctx,
		&currentToken,
		&tokenSource{token: &newToken},
		WithOnRefreshTokenChange(func(ctx context.Context, newToken *oauth2.Token) error {
			require.Equal(t, "new_refresh_token", newToken.RefreshToken)
			triggered = true
			return nil
		}),
	))

	res, err := client.Get(ctx, server.URL)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, 200, res.StatusCode)
	require.True(t, triggered)
	res.Body.Close()
}

type errorTokenSource struct{}

func (e *errorTokenSource) Token() (*oauth2.Token, error) {
	return nil, errors.New("error when refresh token")
}

var ErrAuthenticationInvalidInTest = errors.New("authentication invalid")

func TestTransportReturnSpecificErrorWhenRefreshError(t *testing.T) {
	ctx := context.Background()
	currentToken := oauth2.Token{
		RefreshToken: "refresh_token",
		AccessToken:  "access_token",
		TokenType:    "Bearer",
		Expiry:       time.Now().Add(-time.Hour),
	}

	server := newMockServer(func(w http.ResponseWriter, r *http.Request) {})
	defer server.Close()

	triggered := false

	client := httpx.NewXClient(WithOAuth2Http(
		ctx,
		&currentToken,
		&errorTokenSource{},
		WithOnRefreshTokenChange(func(ctx context.Context, newToken *oauth2.Token) error {
			triggered = true
			return nil
		}),
		WithAuthError(ErrAuthenticationInvalidInTest),
	))

	res, err := client.Get(ctx, server.URL)
	require.Error(t, err)

	require.Equal(t, ErrAuthenticationInvalidInTest, err)
	require.Nil(t, res)
	require.False(t, triggered)
}

func TestTransportReturnSpecificErrorWhenCallOnChangeError(t *testing.T) {
	ctx := context.Background()
	currentToken := oauth2.Token{
		RefreshToken: "refresh_token",
		AccessToken:  "access_token",
		TokenType:    "Bearer",
		Expiry:       time.Now().Add(-time.Hour),
	}
	newToken := oauth2.Token{
		RefreshToken: "new_refresh_token",
		AccessToken:  "access_token",
		TokenType:    "Bearer",
		Expiry:       time.Now().Add(time.Hour),
	}

	server := newMockServer(func(w http.ResponseWriter, r *http.Request) {})
	defer server.Close()

	triggered := false

	client := httpx.NewXClient(WithOAuth2Http(
		ctx,
		&currentToken,
		&tokenSource{token: &newToken},
		WithOnRefreshTokenChange(func(ctx context.Context, newToken *oauth2.Token) error {
			triggered = true
			return errors.New("error when call on refresh token change")
		}),
		WithAuthError(ErrAuthenticationInvalidInTest),
	))

	res, err := client.Get(ctx, server.URL)
	require.Error(t, err)

	require.Equal(t, ErrAuthenticationInvalidInTest, err)
	require.Nil(t, res)
	require.True(t, triggered)
}

func TestTransportReturnSpecificErrorWhenAuthHeaderInvalid(t *testing.T) {
	ctx := context.Background()
	currentToken := oauth2.Token{
		RefreshToken: "refresh_token",
		AccessToken:  "access_token",
		TokenType:    "Bearer",
		Expiry:       time.Now().Add(-time.Hour),
	}

	server := newMockServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
	})
	defer server.Close()

	triggered := false

	client := httpx.NewXClient(WithOAuth2Http(
		ctx,
		&currentToken,
		&tokenSource{token: &currentToken},
		WithOnRefreshTokenChange(func(ctx context.Context, newToken *oauth2.Token) error {
			triggered = true
			return errors.New("error when call on refresh token change")
		}),
		WithAuthError(ErrAuthenticationInvalidInTest),
	))

	res, err := client.Get(ctx, server.URL)
	require.Error(t, err)

	require.Equal(t, ErrAuthenticationInvalidInTest, err)
	require.Nil(t, res)
	require.False(t, triggered)
}

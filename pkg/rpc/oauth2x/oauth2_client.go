package oauth2x

/**
OAuth2Core 改造自 oauth2.Transport , 做了如下修改：
1. 不在 transport 层面处理 token 的刷新，而是在 client 层面处理
2. 增加了 OnRefreshTokenChange hook 用来处理 token 刷新后的逻辑
3. RefreshTokenFailed/AuthorizationFailed 时返回特定错误

代码在 https://github.com/golang/oauth2/blob/master/transport.go 基础上有所修改
*/

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/RyoJerryYu/go-utilx/pkg/rpc/httpx"
	"golang.org/x/oauth2"
)

// OAuth2Core is an httpx.Client that handles OAuth2 authentication automatically.
// It is based on oauth2.Transport but with the following improvements:
// 1. Token refresh is handled at the client level instead of transport level
// 2. Provides hooks for token refresh events
// 3. Returns specific errors for authentication failures
//
// Example usage:
//
//	core := &OAuth2Core{
//	    Source: oauth2.StaticTokenSource(initialToken),
//	    OnRefreshTokenChange: func(ctx context.Context, newToken *oauth2.Token) error {
//	        // Save the new token
//	        return saveToken(newToken)
//	    },
//	}
type OAuth2Core struct {
	// Source supplies the token to add to outgoing requests'
	// Authorization headers.
	Source oauth2.TokenSource

	// Inner is the wrapped httpx.Client used to make HTTP requests.
	// If nil, http.DefaultClient is used.
	Inner httpx.Client

	Ctx                      context.Context
	ErrAuthenticationInvalid error  // error to return when RefreshTokenFailed / AuthorizationFailed
	CurrentRefreshToken      string // used to compare with the new token
	OnRefreshTokenChange     func(ctx context.Context, newToken *oauth2.Token) error
	OnAuthError              func(ctx context.Context, oldToken *oauth2.Token, refreshErr error)
	OnRecordError            func(ctx context.Context, err error)
}

// Do authorizes and authenticates the request with an access token.
// It will:
// 1. Get a valid token from the TokenSource
// 2. Trigger OnRefreshTokenChange if the refresh token has changed
// 3. Add the token to the request's Authorization header
// 4. Execute the request
// 5. Handle authentication errors (401/403) by returning ErrAuthenticationInvalid
//
// Note: When returning an error, the response body will be closed automatically.
func (t *OAuth2Core) Do(req *http.Request) (*http.Response, error) {
	reqBodyClosed := false
	if req.Body != nil {
		defer func() {
			if !reqBodyClosed {
				req.Body.Close()
			}
		}()
	}

	if t.Source == nil {
		return nil, errors.New("oauth2: Transport's Source is nil")
	}
	token, err := t.Source.Token()
	if err != nil {
		return nil, t.returnAuthError(token, err) // 返回特定错误
	}
	err = t.doIfRefreshTokenChanged(token)
	if err != nil {
		return nil, t.returnAuthError(token, err) // 返回特定错误
	}

	req2 := t.cloneRequest(req) // per RoundTripper contract
	token.SetAuthHeader(req2)

	// req.Body is assumed to be closed by the base RoundTripper.
	reqBodyClosed = true
	resp, err := t.base().Do(req2)
	if err != nil {
		return nil, err // 这里不需要特殊处理，因为 err 不是认证错误
	}

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		defer resp.Body.Close() // golang HTTP 规范要求 RoundTrip 中返回 err 时，resp.Body 需要关闭
		respBuf, err := io.ReadAll(resp.Body)
		if err != nil {
			t.recordError(t.Ctx, err)
		}

		err = fmt.Errorf("oauth2: Authorization failed, resp: %s", string(respBuf))
		return nil, t.returnAuthError(token, err) // 返回特定错误
	}

	return resp, nil
}

func (t *OAuth2Core) base() httpx.Client {
	if t.Inner != nil {
		return t.Inner
	}
	return http.DefaultClient
}

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Header map.
func (t *OAuth2Core) cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}

func (t *OAuth2Core) doIfRefreshTokenChanged(newToken *oauth2.Token) error {
	if t.CurrentRefreshToken == "" || t.CurrentRefreshToken != newToken.RefreshToken {
		t.CurrentRefreshToken = newToken.RefreshToken
		if t.OnRefreshTokenChange != nil {
			return t.OnRefreshTokenChange(t.Ctx, newToken)
		}
	}
	return nil
}

func (t *OAuth2Core) returnAuthError(oldToken *oauth2.Token, err error) error {
	if t.OnAuthError != nil {
		t.OnAuthError(t.Ctx, oldToken, err)
	}

	// no authenticationInvalid error registered, return raw error
	if t.ErrAuthenticationInvalid == nil {
		return err
	}

	t.recordError(t.Ctx, err)
	return t.ErrAuthenticationInvalid
}

func (t *OAuth2Core) recordError(ctx context.Context, err error) {
	if t.OnRecordError != nil {
		t.OnRecordError(ctx, err)
	}
}

// WithOAuth2Http creates a ClientDecorator that adds OAuth2 authentication to an HTTP client.
// It will use the provided token source for authentication and handle token refresh automatically.
//
// Example:
//
//	client := httpx.NewXClient(
//	    WithOAuth2Http(ctx, currentToken, tokenSource,
//	        WithOnRefreshTokenChange(func(ctx context.Context, newToken *oauth2.Token) error {
//	            return saveNewToken(newToken)
//	        }),
//	        WithAuthError(ErrCustomAuth),
//	    ),
//	)
func WithOAuth2Http(
	ctx context.Context,
	currentToken *oauth2.Token,
	tokenSource oauth2.TokenSource,
	opts ...OAuth2HttpOption,
) httpx.ClientDecorator {
	refreshToken := ""
	if currentToken != nil {
		refreshToken = currentToken.RefreshToken
	}

	return func(cc httpx.Client) httpx.Client {
		oAuth2Core := &OAuth2Core{
			Source:              tokenSource,
			Inner:               cc,
			Ctx:                 ctx,
			CurrentRefreshToken: refreshToken,
		}
		for _, opt := range opts {
			opt(oAuth2Core)
		}
		return oAuth2Core
	}
}

type OAuth2HttpOption func(t *OAuth2Core)

// WithOnRefreshTokenChange sets a callback that is triggered when the refresh token changes.
// This is useful for persisting the new token for future use.
// If the callback returns an error, the request will fail with ErrAuthenticationInvalid
// (if configured) or the original error.
//
// Example:
//
//	WithOnRefreshTokenChange(func(ctx context.Context, newToken *oauth2.Token) error {
//	    return db.SaveToken(newToken)
//	})
func WithOnRefreshTokenChange(onRefreshTokenChange func(ctx context.Context, newToken *oauth2.Token) error) OAuth2HttpOption {
	return func(t *OAuth2Core) {
		t.OnRefreshTokenChange = onRefreshTokenChange
	}
}

// WithOnAuthError sets a callback that is triggered when any authentication error occurs.
// This includes token refresh failures, OnRefreshTokenChange errors, and 401/403 responses.
// The callback receives the token that was used in the failed request and the error that occurred.
//
// This is useful for logging authentication failures or updating metrics.
// The callback is called before returning ErrAuthenticationInvalid (if configured).
//
// Example:
//
//	WithOnAuthError(func(ctx context.Context, oldToken *oauth2.Token, err error) {
//	    log.Printf("Auth failed for user %v: %v", oldToken.AccessToken, err)
//	    metrics.AuthFailures.Inc()
//	})
func WithOnAuthError(onRefreshTokenFailed func(ctx context.Context, oldToken *oauth2.Token, refreshErr error)) OAuth2HttpOption {
	return func(t *OAuth2Core) {
		t.OnAuthError = onRefreshTokenFailed
	}
}

// WithAuthError configures the error that will be returned for all authentication failures.
// When set, this error will be returned instead of the original error for:
// - Token refresh failures
// - OnRefreshTokenChange callback errors
// - 401/403 responses from the server
//
// If not set, the original error will be returned.
//
// Example:
//
//	var ErrAuthExpired = errors.New("authentication expired")
//	client := NewXClient(WithOAuth2Http(ctx, token, source,
//	    WithAuthError(ErrAuthExpired),
//	))
func WithAuthError(err error) OAuth2HttpOption {
	return func(t *OAuth2Core) {
		t.ErrAuthenticationInvalid = err
	}
}

// WithRecordError sets a callback for recording internal errors that occur during
// authentication but don't affect the request outcome.
// Currently this is only used for errors that occur while reading the response body
// of a failed authentication attempt.
//
// Example:
//
//	WithRecordError(func(ctx context.Context, err error) {
//	    log.Printf("Internal oauth error: %v", err)
//	})
func WithRecordError(onRecordError func(ctx context.Context, err error)) OAuth2HttpOption {
	return func(t *OAuth2Core) {
		t.OnRecordError = onRecordError
	}
}

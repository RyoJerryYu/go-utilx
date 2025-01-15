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

// OAuth2Core is an httpx.Client
// It could handle oauth2 token refresh and auth automatically
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

// Do authorizes and authenticates the request with an
// access token from Transport's Source.
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

func WithOnRefreshTokenChange(onRefreshTokenChange func(ctx context.Context, newToken *oauth2.Token) error) OAuth2HttpOption {
	return func(t *OAuth2Core) {
		t.OnRefreshTokenChange = onRefreshTokenChange
	}
}

func WithOnAuthError(onRefreshTokenFailed func(ctx context.Context, oldToken *oauth2.Token, refreshErr error)) OAuth2HttpOption {
	return func(t *OAuth2Core) {
		t.OnAuthError = onRefreshTokenFailed
	}
}

func WithAuthError(err error) OAuth2HttpOption {
	return func(t *OAuth2Core) {
		t.ErrAuthenticationInvalid = err
	}
}

func WithRecordError(onRecordError func(ctx context.Context, err error)) OAuth2HttpOption {
	return func(t *OAuth2Core) {
		t.OnRecordError = onRecordError
	}
}

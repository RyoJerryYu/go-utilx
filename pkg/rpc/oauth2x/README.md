# OAuth2X

OAuth2X provides an enhanced OAuth2 client that improves upon the standard `oauth2.Transport`. It offers better control over token refresh and authentication failures.

## Key Features

- **Client-Level Token Refresh**: Token refresh is handled at the client level instead of transport level
- **Token Change Hooks**: Callbacks for token refresh events, making it easy to persist new tokens
- **Error Handling**: Customizable error handling for authentication failures
- **Flexible Configuration**: Multiple hooks for different authentication scenarios
- **Compatible API**: Fully compatible with the standard `oauth2.Transport`

## Basic Usage

```go
// Create a client with OAuth2 support
client := httpx.NewXClient(
    WithOAuth2Http(ctx, currentToken, tokenSource,
        WithOnRefreshTokenChange(func(ctx context.Context, newToken *oauth2.Token) error {
            // Save the new token when it changes
            return db.SaveToken(newToken)
        }),
        WithAuthError(ErrCustomAuth),
    ),
)

// Make requests as usual - token handling is automatic
var response struct {
    Data string `json:"data"`
}
err := client.GetJSON(ctx, "https://api.example.com/protected", &response)
```

## Architecture

OAuth2X is built on top of the `httpx.Client` interface and integrates seamlessly with the HTTPX package. It handles:

1. Automatic token refresh when needed
2. Token persistence through callbacks
3. Custom error handling for authentication failures
4. Proper cleanup of resources

## Options

### OAuth2 Options

- `WithOnRefreshTokenChange(func)`: Callback when refresh token changes
- `WithOnAuthError(func)`: Callback for authentication failures
- `WithAuthError(error)`: Custom error for auth failures
- `WithRecordError(func)`: Callback for internal errors

## Error Handling

Authentication failures can occur in several scenarios:
- Token refresh fails
- Token persistence fails
- Server returns 401/403 response

You can handle these cases by:
1. Setting a custom error with `WithAuthError`
2. Registering callbacks with `WithOnAuthError`
3. Logging internal errors with `WithRecordError`

```go
client := httpx.NewXClient(
    WithOAuth2Http(ctx, token, source,
        WithAuthError(ErrAuthExpired),
        WithOnAuthError(func(ctx context.Context, token *oauth2.Token, err error) {
            log.Printf("Auth failed: %v", err)
        }),
    ),
)
```

## Migration from oauth2.Transport

OAuth2X is designed to be a drop-in replacement for `oauth2.Transport`. The main differences are:

1. Token refresh happens at the client level
2. You can hook into token refresh events
3. You have more control over error handling
4. Response bodies are automatically closed on errors 
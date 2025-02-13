# HTTPX

HTTPX provides an enhanced HTTP client `XClient` that improves upon the standard `http.Client`.

## Key Features

- **Middleware Support**: Add behaviors to all requests through `XClientOption` and per-request through `XRequestOption`
- **Context-First**: Every request requires a context, making it easier to handle timeouts and cancellation
- **OpenTelemetry Integration**: Built-in support for distributed tracing
- **Error Handling**: Automatic handling of non-2xx responses with detailed error information
- **Flexible Configuration**: Supports both client-level and request-level configuration

## Basic Usage

```go
// Create a new client with options
client := httpx.NewXClient(
    WithBearerAuth("your-token"),
    WithTimeout(5 * time.Second),
)

// Make a GET request with query parameters
var response struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
err := client.GetJSON(ctx, "https://api.example.com/users",
    &response,
    WithQuery("id", "123"),
    WithHeader("X-Custom", "value"),
)

// Make a POST request with JSON body
request := struct {
    Name string `json:"name"`
}{Name: "John"}
err = client.PostJSON(ctx, "https://api.example.com/users", 
    request, &response)
```

## Architecture

`XClient` wraps `http.Client` through the `Client` interface, which only depends on the `Do` method. This design:

1. Simplifies testing and mocking
2. Allows for easy composition of client behaviors
3. Makes it easier to add middleware-like functionality

## Options

### Client Options

- `WithOtel()`: Adds OpenTelemetry instrumentation
- `WithBearerAuth(token)`: Adds Bearer token authentication
- `WithReturnErrorIfNot2xx()`: Returns errors for non-2xx responses

### Request Options

- `WithQuery(key, value)`: Adds a query parameter
- `WithQuerys(values)`: Adds multiple query parameters
- `WithHeader(key, value)`: Adds a header
- `WithHeaders(headers)`: Adds multiple headers

## Error Handling

When using `WithReturnErrorIfNot2xx()`, non-2xx responses return an `XError` containing:
- Original response
- HTTP method
- Status code
- Response body

```go
if xerr, ok := err.(*XError); ok {
    log.Printf("Request failed: %d %s", xerr.Code, xerr.Body)
}
```

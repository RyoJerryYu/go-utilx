# HTTPX

This package provide a http client `XClient` that improve the client in http package.

`XClient` wrap the `http.Client` , and allow modifying client by `XClientOption` ,
and allow modifying each request by `XRequestOption` .

`XClient` force every call provide a context, that would make much easier to work around 
such as otel.

`XClient` wrap the `http.Client` by `Client` interface, and only depend on it's `Do` method.

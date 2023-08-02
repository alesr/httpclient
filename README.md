# HTTP Client

This package is designed to simplify the configuration of timeouts, connection options, and other settings for Go's HTTP client. It supports a variety of options, including dialer timeout, keep-alive duration, TLS handshake timeout, response header timeout, idle connection timeout, maximum idle connections, and HTTP/2 configuration.

## Installation

You can import this package into your Go project using go get:

```bash
go get github.com/alesr/httpclient
```

## Usage

First, import the package into your Go file:

```go
import "github.com/alesr/httpclient"
```

You can create a new HTTP client with default settings as follows:

```go
client := httpclient.New()
```

You can modify the default configuration using the options provided by the library:

```go
client := httpclient.New(
    httpclient.WithTimeout(10 * time.Second),
    httpclient.WithDialerTimeout(5 * time.Second),
    httpclient.WithDialerKeepAlive(15 * time.Second),
    httpclient.WithTLSHandshakeTimeout(5 * time.Second),
    httpclient.WithResponseHeaderTimeout(10 * time.Second),
    httpclient.WithIdleConnTimeout(15 * time.Second),
    httpclient.WithMaxIdleConns(50),
    httpclient.WithForceHTTP2Disabled(),
)
```

Each option modifies a different part of the HTTP client configuration:

- `WithTimeout`: sets the request timeout duration.
- `WithDialerTimeout`: sets the dialer timeout duration.
- `WithDialerKeepAlive`: sets the dialer keep-alive duration.
- `WithTLSHandshakeTimeout`: sets the TLS handshake timeout duration.
- `WithResponseHeaderTimeout`: sets the response header timeout duration.
- `WithIdleConnTimeout`: sets the idle connection timeout duration.
- `WithMaxIdleConns`: sets the maximum number of idle connections.
- `WithForceHTTP2Disabled`: disables HTTP2 for the transport of the HTTP client.

After creating the client, you can use it to make HTTP requests as you would with the standard `net/http` package:

```go
resp, err := client.Get("http://example.com")
// handle error, use response...
```


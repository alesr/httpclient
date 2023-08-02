package httpclient

import (
	"net"
	"net/http"
	"time"
)

// Client struct holds the configurations for the HTTP client.
type Client struct {
	Timeout               time.Duration // Time limit for requests made by this Client.
	DialerTimeout         time.Duration // Maximum amount of time a dial will wait for a connect to complete.
	DialerKeepAlive       time.Duration // Interval between keep-alive probes for an active network connection.
	TLSHandshakeTimeout   time.Duration // Time spent waiting for a TLS handshake.
	ResponseHeaderTimeout time.Duration // Time to wait for a server's response headers.
	IdleConnTimeout       time.Duration // Maximum amount of time an idle (keep-alive) connection will remain idle before closing itself.
	MaxIdleConns          int           // Maximum number of idle (keep-alive) connections across all hosts.
	ForceAttemptHTTP2     bool          // If true, HTTP/2 is enabled for this transport.
}

// optionFunc defines the type of function that can be used to modify the configuration.
type optionFunc func(*Client)

// WithTimeout sets the request timeout duration.
func WithTimeout(timeout time.Duration) optionFunc {
	return func(c *Client) {
		c.Timeout = timeout
	}
}

// WithDialerTimeout sets the dialer timeout duration.
func WithDialerTimeout(timeout time.Duration) optionFunc {
	return func(c *Client) {
		c.DialerTimeout = timeout
	}
}

// WithDialerKeepAlive sets the dialer keep-alive duration.
func WithDialerKeepAlive(keepAlive time.Duration) optionFunc {
	return func(c *Client) {
		c.DialerKeepAlive = keepAlive
	}
}

// WithTLSHandshakeTimeout sets the TLS handshake timeout duration.
func WithTLSHandshakeTimeout(tlsHandshakeTimeout time.Duration) optionFunc {
	return func(c *Client) {
		c.TLSHandshakeTimeout = tlsHandshakeTimeout
	}
}

// WithResponseHeaderTimeout sets the response header timeout duration.
func WithResponseHeaderTimeout(responseHeaderTimeout time.Duration) optionFunc {
	return func(c *Client) {
		c.ResponseHeaderTimeout = responseHeaderTimeout
	}
}

// WithIdleConnTimeout sets the idle connection timeout duration.
func WithIdleConnTimeout(idleConnTimeout time.Duration) optionFunc {
	return func(c *Client) {
		c.IdleConnTimeout = idleConnTimeout
	}
}

// WithMaxIdleConns sets the maximum number of idle connections.
func WithMaxIdleConns(maxIdleConns int) optionFunc {
	return func(c *Client) {
		c.MaxIdleConns = maxIdleConns
	}
}

// WithForceHTTP2Disabled disables HTTP2 for the transport of the HTTP client.
func WithForceHTTP2Disabled() optionFunc {
	return func(c *Client) {
		c.ForceAttemptHTTP2 = false
	}
}

// New creates a new HTTP client with the provided options.
// If no options are provided, a client with default settings is returned.
func New(options ...optionFunc) *http.Client {
	c := Client{
		Timeout:               15 * time.Second,
		DialerTimeout:         5 * time.Second,
		DialerKeepAlive:       15 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		IdleConnTimeout:       15 * time.Second,
		MaxIdleConns:          50,
		ForceAttemptHTTP2:     true,
	}

	for _, option := range options {
		option(&c)
	}

	dialer := &net.Dialer{
		Timeout:   c.DialerTimeout,
		KeepAlive: c.DialerKeepAlive,
	}

	transport := http.Transport{
		DialContext:           dialer.DialContext,
		TLSHandshakeTimeout:   c.TLSHandshakeTimeout,
		ResponseHeaderTimeout: c.ResponseHeaderTimeout,
		IdleConnTimeout:       c.IdleConnTimeout,
		MaxIdleConns:          c.MaxIdleConns,
		ForceAttemptHTTP2:     c.ForceAttemptHTTP2,
	}

	return &http.Client{
		Timeout:   c.Timeout,
		Transport: &transport,
	}
}

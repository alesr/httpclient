package httpclient

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Parallel()

	t.Run("Default client", func(t *testing.T) {
		t.Parallel()

		httpClient := New()

		transport := httpClient.Transport.(*http.Transport)

		require.NotNil(t, httpClient)

		assert.Equal(t, 15*time.Second, httpClient.Timeout)
		assert.Equal(t, 15*time.Second, transport.IdleConnTimeout)
		assert.Equal(t, 10*time.Second, transport.ResponseHeaderTimeout)
		assert.Equal(t, 5*time.Second, transport.TLSHandshakeTimeout)
		assert.Equal(t, 50, transport.MaxIdleConns)
		assert.Equal(t, true, transport.ForceAttemptHTTP2)
	})

	t.Run("WithTimeout", func(t *testing.T) {
		t.Parallel()

		httpClient := New(WithTimeout(1 * time.Second))

		assert.Equal(t, 1*time.Second, httpClient.Timeout)
	})

	t.Run("WithTLSHandshakeTimeout", func(t *testing.T) {
		t.Parallel()

		httpClient := New(WithTLSHandshakeTimeout(1 * time.Second))

		transport := httpClient.Transport.(*http.Transport)

		assert.Equal(t, 1*time.Second, transport.TLSHandshakeTimeout)
	})

	t.Run("WithResponseHeaderTimeout", func(t *testing.T) {
		t.Parallel()

		httpClient := New(WithResponseHeaderTimeout(1 * time.Second))

		transport := httpClient.Transport.(*http.Transport)

		assert.Equal(t, 1*time.Second, transport.ResponseHeaderTimeout)
	})

	t.Run("WithIdleConnTimeout", func(t *testing.T) {
		t.Parallel()

		httpClient := New(WithIdleConnTimeout(1 * time.Second))

		transport := httpClient.Transport.(*http.Transport)

		assert.Equal(t, 1*time.Second, transport.IdleConnTimeout)
	})

	t.Run("WithMaxIdleConns", func(t *testing.T) {
		t.Parallel()

		httpClient := New(WithMaxIdleConns(100))

		transport := httpClient.Transport.(*http.Transport)

		assert.Equal(t, 100, transport.MaxIdleConns)
	})

	t.Run("WithForceHTTP2Disabled", func(t *testing.T) {
		t.Parallel()

		httpClient := New(WithForceHTTP2Disabled())

		transport := httpClient.Transport.(*http.Transport)

		assert.Equal(t, false, transport.ForceAttemptHTTP2)
	})

	t.Run("WithCustomTransport", func(t *testing.T) {
		t.Parallel()

		customTransport := &http.Transport{
			MaxIdleConns:    200,
			IdleConnTimeout: 30 * time.Second,
		}

		httpClient := New(WithTransport(customTransport))

		assert.Equal(t, customTransport, httpClient.Transport)
	})
}

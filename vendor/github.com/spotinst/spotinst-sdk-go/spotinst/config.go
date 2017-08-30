package spotinst

import (
	"fmt"
	"net/http"

	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
	"github.com/spotinst/spotinst-sdk-go/spotinst/log"
)

const (
	// SDKVersion is the current version of the SDK.
	SDKVersion = "2.4.6"

	// SDKName is the name of the SDK.
	SDKName = "spotinst-sdk-go"

	// DefaultAddress is the default address of the Spotinst API.
	// It is used e.g. when initializing a new Client without a specific address.
	DefaultAddress = "api.spotinst.io"

	// DefaultScheme is the default protocol scheme to use when making HTTP
	// calls.
	DefaultScheme = "https"

	// DefaultContentType is the default content type to use when making HTTP
	// calls.
	DefaultContentType = "application/json"

	// DefaultUserAgent is the default user agent to use when making HTTP
	// calls.
	DefaultUserAgent = SDKName + "/" + SDKVersion

	// DefaultMaxRetries is the number of retries for a single request after
	// the client will give up and return an error. It is zero by default, so
	// retry is disabled by default.
	DefaultMaxRetries = 0

	// DefaultGzipEnabled specifies if gzip compression is enabled by default.
	DefaultGzipEnabled = false
)

// clientConfig is used to configure the creation of a client.
type clientConfig struct {
	// address is the address of the API server.
	address string

	// scheme is the URI scheme for the API server.
	scheme string

	// httpClient is the client to use. Default will be used if not provided.
	httpClient *http.Client

	// credentials provides synchronous safe retrieval of Spotinst credentials.
	credentials *credentials.Credentials

	// userAgent is the user agent to use when making HTTP calls.
	userAgent string

	// contentType is the content type to use when making HTTP calls.
	contentType string

	// errorf logs to the error log.
	errorlog log.Logger

	// infof logs informational messages.
	infolog log.Logger

	// tracef logs to the trace log.
	tracelog log.Logger
}

// ClientOption is a function that configures a Client.
// It is used in NewClient.
type ClientOption func(*clientConfig)

// SetAddress defines the address of the Spotinst API.
func SetAddress(addr string) ClientOption {
	return func(c *clientConfig) {
		c.address = addr
	}
}

// SetScheme defines the scheme for the address of the Spotinst API.
func SetScheme(scheme string) ClientOption {
	return func(c *clientConfig) {
		c.scheme = scheme
	}
}

// SetHttpClient defines the HTTP client.
func SetHttpClient(client *http.Client) ClientOption {
	return func(c *clientConfig) {
		c.httpClient = client
	}
}

// SetCredentials defines the credentials.
func SetCredentials(creds *credentials.Credentials) ClientOption {
	return func(c *clientConfig) {
		c.credentials = creds
	}
}

// SetUserAgent defines the user agent.
func SetUserAgent(ua string) ClientOption {
	return func(c *clientConfig) {
		c.userAgent = fmt.Sprintf("%s+%s", ua, c.userAgent)
	}
}

// SetContentType defines the content type.
func SetContentType(ct string) ClientOption {
	return func(c *clientConfig) {
		c.contentType = ct
	}
}

// SetErrorLog sets the logger for critical messages like nodes joining
// or leaving the cluster or failing requests. It is nil by default.
func SetErrorLog(logger log.Logger) ClientOption {
	return func(c *clientConfig) {
		c.errorlog = logger
	}
}

// SetInfoLog sets the logger for informational messages, e.g. requests
// and their response times. It is nil by default.
func SetInfoLog(logger log.Logger) ClientOption {
	return func(c *clientConfig) {
		c.infolog = logger
	}
}

// SetTraceLog specifies the log.Logger to use for output of HTTP requests
// and responses which is helpful during debugging. It is nil by default.
func SetTraceLog(logger log.Logger) ClientOption {
	return func(c *clientConfig) {
		c.tracelog = logger
	}
}

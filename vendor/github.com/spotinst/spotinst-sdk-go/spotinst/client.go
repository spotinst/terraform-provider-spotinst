package spotinst

import (
	"context"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
)

// DefaultTransport returns a new http.Transport with similar default
// values to http.DefaultTransport. Do not use this for transient transports as
// it can leak file descriptors over time. Only use this for transports that
// will be re-used for the same host(s).
func DefaultTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,
		DisableKeepAlives:   false,
		MaxIdleConnsPerHost: 1,
	}
}

// DefaultHttpClient returns a new http.Client with similar default values to
// http.Client, but with a non-shared Transport, idle connections disabled, and
// KeepAlives disabled.
func DefaultHttpClient() *http.Client {
	return &http.Client{
		Transport: DefaultTransport(),
	}
}

// defaultConfig returns a default configuration for the client. By default this
// will pool and reuse idle connections to API. If you have a long-lived
// client object, this is the desired behavior and should make the most efficient
// use of the connections to API.
func defaultConfig() *clientConfig {
	return &clientConfig{
		address:     DefaultAddress,
		scheme:      DefaultScheme,
		userAgent:   DefaultUserAgent,
		contentType: DefaultContentType,
		httpClient:  DefaultHttpClient(),
		credentials: credentials.NewChainCredentials(
			new(credentials.EnvProvider),
			new(credentials.FileProvider),
		),
	}
}

// Client provides a client to the API.
type Client struct {
	config              *clientConfig
	GroupService        GroupService
	MultaiService       MultaiService
	HealthCheckService  HealthCheckService
	SubscriptionService SubscriptionService
}

// NewClient returns a new client.
func NewClient(opts ...ClientOption) *Client {
	config := defaultConfig()
	for _, o := range opts {
		o(config)
	}

	client := &Client{config: config}
	client.GroupService = &GroupServiceOp{client}
	client.MultaiService = &MultaiServiceOp{client}
	client.HealthCheckService = &HealthCheckServiceOp{client}
	client.SubscriptionService = &SubscriptionServiceOp{client}

	return client
}

// newRequest is used to create a new request.
func (c *Client) newRequest(ctx context.Context, method, path string) *request {
	req := &request{
		context: ctx,
		config:  c.config,
		method:  method,
		url: &url.URL{
			Scheme: c.config.scheme,
			Host:   c.config.address,
			Path:   path,
		},
		params: make(map[string][]string),
		header: make(http.Header),
	}
	return req
}

// doRequest runs a request with our client.
func (c *Client) doRequest(r *request) (time.Duration, *http.Response, error) {
	creds, err := c.config.credentials.Get()
	if err != nil {
		c.errorf("Failed to retrieve credentials: %s", err)
		return 0, nil, err
	}
	c.tracef("Credentials retrieved from provider: %s", creds.ProviderName)
	if creds.Token != "" {
		r.header.Set("Authorization", "Bearer "+creds.Token)
	}
	if creds.Account != "" {
		r.params.Set("accountId", creds.Account)
	}
	req, err := r.toHTTP()
	if err != nil {
		return 0, nil, err
	}
	c.dumpRequest(req)
	start := time.Now()
	resp, err := c.config.httpClient.Do(req)
	diff := time.Now().Sub(start)
	c.dumpResponse(resp)
	return diff, resp, err
}

// errorf logs to the error log.
func (c *Client) errorf(format string, args ...interface{}) {
	if c.config.errorlog != nil {
		c.config.errorlog.Printf(format, args...)
	}
}

// infof logs informational messages.
func (c *Client) infof(format string, args ...interface{}) {
	if c.config.infolog != nil {
		c.config.infolog.Printf(format, args...)
	}
}

// tracef logs to the trace log.
func (c *Client) tracef(format string, args ...interface{}) {
	if c.config.tracelog != nil {
		c.config.tracelog.Printf(format, args...)
	}
}

// dumpRequest dumps the given HTTP request to the trace log.
func (c *Client) dumpRequest(r *http.Request) {
	if c.config.tracelog != nil && r != nil {
		out, err := httputil.DumpRequestOut(r, true)
		if err == nil {
			c.tracef("%s\n", string(out))
		}
	}
}

// dumpResponse dumps the given HTTP response to the trace log.
func (c *Client) dumpResponse(resp *http.Response) {
	if c.config.tracelog != nil && resp != nil {
		out, err := httputil.DumpResponse(resp, true)
		if err == nil {
			c.tracef("%s\n", string(out))
		}
	}
}

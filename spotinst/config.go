package spotinst

import (
	"errors"
	"fmt"
	stdlog "log"
	"strings"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform-plugin-sdk/meta"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup"
	"github.com/spotinst/spotinst-sdk-go/service/healthcheck"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance"
	"github.com/spotinst/spotinst-sdk-go/service/mrscaler"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/subscription"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
	"github.com/spotinst/spotinst-sdk-go/spotinst/log"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"github.com/terraform-providers/terraform-provider-spotinst/version"
)

var ErrNoValidCredentials = errors.New("\n\nNo valid credentials found " +
	"for Spotinst Provider.\nPlease see https://www.terraform.io/docs/" +
	"providers/spotinst/index.html\nfor more information on providing " +
	"credentials for Spotinst Provider.")

type Config struct {
	Token   string
	Account string

	terraformVersion string
}

type Client struct {
	elastigroup     elastigroup.Service
	healthCheck     healthcheck.Service
	subscription    subscription.Service
	multai          multai.Service
	mrscaler        mrscaler.Service
	ocean           ocean.Service
	managedInstance managedinstance.Service
}

// Client configures and returns a fully initialized Spotinst client.
func (c *Config) Client() (*Client, error) {
	stdlog.Println("[INFO] Configuring a new Spotinst client")

	// Create a new session.
	sess, err := c.getSession()
	if err != nil {
		return nil, err
	}

	// Create a new client.
	client := &Client{
		elastigroup:     elastigroup.New(sess),
		healthCheck:     healthcheck.New(sess),
		subscription:    subscription.New(sess),
		multai:          multai.New(sess),
		mrscaler:        mrscaler.New(sess),
		ocean:           ocean.New(sess),
		managedInstance: managedinstance.New(sess),
	}

	stdlog.Println("[INFO] Spotinst client configured")
	return client, nil
}

func (c *Config) getSession() (*session.Session, error) {
	config := spotinst.DefaultConfig()

	// HTTP options.
	{
		config.WithHTTPClient(cleanhttp.DefaultPooledClient())
		config.WithUserAgent(c.getUserAgent())
	}

	// Credentials.
	{
		v, err := c.getCredentials()
		if err != nil {
			return nil, err
		}
		config.WithCredentials(v)
	}

	// Logging.
	{
		config.WithLogger(log.LoggerFunc(func(format string, args ...interface{}) {
			stdlog.Printf(fmt.Sprintf("[DEBUG] [spotinst-sdk-go] %s", format), args...)
		}))
	}

	return session.New(config), nil
}

func (c *Config) getUserAgent() string {
	agents := []struct {
		Product string
		Version string
		Comment []string
	}{
		{Product: "HashiCorp", Version: "1.0"},
		{Product: "Terraform", Version: c.terraformVersion, Comment: []string{"+https://www.terraform.io"}},
		{Product: "Terraform Plugin SDK", Version: meta.SDKVersionString()},
		{Product: "Terraform Provider Spotinst", Version: "v2-" + version.String()},
	}

	var ua string
	for _, agent := range agents {
		pv := fmt.Sprintf("%s/%s", agent.Product, agent.Version)
		if len(agent.Comment) > 0 {
			pv += fmt.Sprintf(" (%s)", strings.Join(agent.Comment, "; "))
		}
		if len(ua) > 0 {
			ua += " "
		}
		ua += pv
	}

	return ua
}

func (c *Config) getCredentials() (*credentials.Credentials, error) {
	var providers []credentials.Provider
	var static *credentials.StaticProvider

	if c.Token != "" || c.Account != "" {
		static = &credentials.StaticProvider{
			Value: credentials.Value{
				Token:   c.Token,
				Account: c.Account,
			},
		}
	}
	if static != nil {
		providers = append(providers, static)
	}

	providers = append(providers,
		new(credentials.EnvProvider),
		new(credentials.FileProvider))

	creds := credentials.NewChainCredentials(providers...)

	if _, err := creds.Get(); err != nil {
		stdlog.Printf("[ERROR] Failed to instantiate Spotinst client: %v", err)
		return nil, ErrNoValidCredentials
	}

	return creds, nil
}

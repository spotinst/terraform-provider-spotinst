package spotinst

import (
	"errors"
	"fmt"
	"github.com/terraform-providers/terraform-provider-spotinst/version"
	stdlog "log"
	"strings"

	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup"
	"github.com/spotinst/spotinst-sdk-go/service/healthcheck"
	"github.com/spotinst/spotinst-sdk-go/service/mrscaler"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/subscription"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
	"github.com/spotinst/spotinst-sdk-go/spotinst/log"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
)

var ErrNoValidCredentials = errors.New("\n\nNo valid credentials found " +
	"for Spotinst Provider.\nPlease see https://www.terraform.io/docs/" +
	"providers/spotinst/index.html\nfor more information on providing " +
	"credentials for Spotinst Provider.")

type Config struct {
	Token   string
	Account string
}

type Client struct {
	elastigroup  elastigroup.Service
	healthCheck  healthcheck.Service
	subscription subscription.Service
	multai       multai.Service
	mrscaler     mrscaler.Service
	ocean        ocean.Service
}

// Validate returns an error in case of invalid configuration.
func (c *Config) Validate() error {
	return nil
}

// Client returns a new client for accessing Spotinst.
func (c *Config) Client() (*Client, error) {
	config := spotinst.DefaultConfig()
	config.WithLogger(newStdLogger("DEBUG"))
	config.WithUserAgent("HashiCorp-Terraform/" + terraform.VersionString() + ",spotinst-provider/v2-" + version.GetShortVersion())

	var static *credentials.StaticProvider
	if c.Token != "" || c.Account != "" {
		static = &credentials.StaticProvider{
			Value: credentials.Value{
				Token:   c.Token,
				Account: c.Account,
			},
		}
	}

	providers := []credentials.Provider{}

	if static != nil {
		providers = append(providers, static)
	}

	providers = append(providers, new(credentials.EnvProvider), new(credentials.FileProvider))

	creds := credentials.NewChainCredentials(providers...)

	if _, err := creds.Get(); err != nil {
		stdlog.Printf("[ERROR] Failed to instantiate Spotinst client: %v", err)
		return nil, ErrNoValidCredentials
	}
	config.WithCredentials(creds)

	// Create a new session.
	sess := session.New(config)

	// Create a new client.
	client := &Client{
		elastigroup:  elastigroup.New(sess),
		healthCheck:  healthcheck.New(sess),
		subscription: subscription.New(sess),
		multai:       multai.New(sess),
		mrscaler:     mrscaler.New(sess),
		ocean:        ocean.New(sess),
	}
	stdlog.Println("[INFO] Spotinst client configured")

	return client, nil
}

func newStdLogger(level string) log.Logger {
	return log.LoggerFunc(func(format string, args ...interface{}) {
		stdlog.Printf(fmt.Sprintf("[%s] %s", strings.ToUpper(level), format), args...)
	})
}

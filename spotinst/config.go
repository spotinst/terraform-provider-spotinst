package spotinst

import (
	"errors"
	"fmt"
	stdlog "log"
	"strings"

	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
	"github.com/spotinst/spotinst-sdk-go/spotinst/log"
)

var ErrNoValidCredentials = errors.New("\n\nNo valid credentials found " +
	"for Spotinst Provider.\nPlease see https://www.terraform.io/docs/" +
	"providers/spotinst/index.html\nfor more information on providing " +
	"credentials for Spotinst Provider.")

type Config struct {
	Token   string
	Account string
}

// Validate returns an error in case of invalid configuration.
func (c *Config) Validate() error {
	return nil
}

// Client returns a new client for accessing Spotinst.
func (c *Config) Client() (*spotinst.Client, error) {
	// Set default client options.
	clientOpts := []spotinst.ClientOption{
		spotinst.SetUserAgent("HashiCorp-Terraform/" + terraform.VersionString()),
		spotinst.SetTraceLog(newStdLogger("TRACE")),
		spotinst.SetErrorLog(newStdLogger("ERROR")),
	}

	// Set user credentials.
	providers := []credentials.Provider{
		new(credentials.EnvProvider),
		new(credentials.FileProvider),
	}

	var static *credentials.StaticProvider
	if c.Token != "" || c.Account != "" {
		static = &credentials.StaticProvider{
			Value: credentials.Value{
				Token:   c.Token,
				Account: c.Account,
			},
		}
		// Static provider should be placed between Env and File providers.
		providers = append(providers[:1], append([]credentials.Provider{static}, providers[1:]...)...)
	}
	creds := credentials.NewChainCredentials(providers...)

	if _, err := creds.Get(); err != nil {
		stdlog.Printf("[ERROR] Failed to instantiate Spotinst client: %v", err)
		return nil, ErrNoValidCredentials
	}
	clientOpts = append(clientOpts, spotinst.SetCredentials(creds))

	// Create a new client.
	client := spotinst.NewClient(clientOpts...)
	stdlog.Println("[INFO] Spotinst client configured")

	return client, nil
}

func newStdLogger(level string) log.Logger {
	return log.LoggerFunc(func(format string, args ...interface{}) {
		stdlog.Printf(fmt.Sprintf("[%s] %s", strings.ToUpper(level), format), args...)
	})
}

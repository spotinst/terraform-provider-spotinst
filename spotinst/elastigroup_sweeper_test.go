package spotinst

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// getProviderClient returns a spotinst client setup with the correct cloud provider configs
func getProviderClient(provider string) (interface{}, error) {
	enabled := "SPOTINST_ENABLED"
	token := "SPOTINST_TOKEN_" + strings.ToUpper(provider)
	account := "SPOTINST_ACCOUNT_" + strings.ToUpper(provider)
	if os.Getenv(enabled) == "true" && os.Getenv(token) == "" && (os.Getenv(account) == "") {
		return nil, fmt.Errorf("must provide environment variables SPOTINST_TOKEN_AWS and SPOTINST_ACCOUNT_AWS")
	}

	conf := &Config{
		Enabled: os.Getenv(enabled),
		Token:   os.Getenv(token),
		Account: os.Getenv(account),
	}

	// configures a default client for the given provider
	client, err := conf.Client()
	if err != nil {
		return nil, fmt.Errorf("error getting Spotinst client")
	}

	return client, nil
}

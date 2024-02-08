package spotinst

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var TestAccProviders map[string]*schema.Provider

var testAccProviderGCP *schema.Provider
var testAccProviderAWS *schema.Provider
var testAccProviderAzure *schema.Provider
var testAccProviderAzureV3 *schema.Provider

func init() {
	testAccProviderGCP = Provider()
	testAccProviderAWS = Provider()
	testAccProviderAzure = Provider()
	testAccProviderAzureV3 = Provider()

	testAccProviderGCP.ConfigureContextFunc = providerConfigureGCP
	testAccProviderAWS.ConfigureContextFunc = providerConfigureAWS
	testAccProviderAzure.ConfigureContextFunc = providerConfigureAzure
	testAccProviderAzureV3.ConfigureContextFunc = providerConfigureAzure

	TestAccProviders = map[string]*schema.Provider{
		"gcp":     testAccProviderGCP,
		"aws":     testAccProviderAWS,
		"azure":   testAccProviderAzure,
		"azurev3": testAccProviderAzureV3,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	_ = Provider()
}

func testAccPreCheck(t *testing.T, provider string) {
	tokens := map[string]string{
		"enabled": os.Getenv("SPOTINST_ENABLED"),
		"gcp":     os.Getenv("SPOTINST_TOKEN_GCP"),
		"aws":     os.Getenv("SPOTINST_TOKEN_AWS"),
		"azure":   os.Getenv("SPOTINST_TOKEN_AZURE"),
	}
	if tokens["enabled"] == "true" && tokens[provider] == "" {
		t.Fatal(ErrNoValidCredentials.Error())
	}
}

func providerConfigureGCP(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		Enabled: os.Getenv("SPOTINST_ENABLED"),
		Token:   os.Getenv("SPOTINST_TOKEN_GCP"),
		Account: os.Getenv("SPOTINST_ACCOUNT_GCP"),
	}

	return config.Client()
}

func providerConfigureAWS(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		Enabled: os.Getenv("SPOTINST_ENABLED"),
		Token:   os.Getenv("SPOTINST_TOKEN_AWS"),
		Account: os.Getenv("SPOTINST_ACCOUNT_AWS"),
	}

	return config.Client()
}

func providerConfigureAzure(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		Enabled: os.Getenv("SPOTINST_ENABLED"),
		Token:   os.Getenv("SPOTINST_TOKEN_AZURE"),
		Account: os.Getenv("SPOTINST_ACCOUNT_AZURE"),
	}

	return config.Client()
}

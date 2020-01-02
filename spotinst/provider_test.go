package spotinst

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var TestAccProviders map[string]terraform.ResourceProvider

var testAccProviderGCP *schema.Provider
var testAccProviderAWS *schema.Provider
var testAccProviderAzure *schema.Provider

func init() {
	testAccProviderGCP = Provider().(*schema.Provider)
	testAccProviderAWS = Provider().(*schema.Provider)
	testAccProviderAzure = Provider().(*schema.Provider)

	testAccProviderGCP.ConfigureFunc = providerConfigureGCP
	testAccProviderAWS.ConfigureFunc = providerConfigureAWS
	testAccProviderAzure.ConfigureFunc = providerConfigureAzure

	TestAccProviders = map[string]terraform.ResourceProvider{
		"gcp":   testAccProviderGCP,
		"aws":   testAccProviderAWS,
		"azure": testAccProviderAzure,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	_ = Provider()
}

func testAccPreCheck(t *testing.T, provider string) {
	tokens := map[string]string{
		"gcp":   os.Getenv("SPOTINST_TOKEN_GCP"),
		"aws":   os.Getenv("SPOTINST_TOKEN_AWS"),
		"azure": os.Getenv("SPOTINST_TOKEN_AZURE"),
	}

	if tokens[provider] == "" {
		t.Fatal(ErrNoValidCredentials.Error())
	}
}

func providerConfigureGCP(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:   os.Getenv("SPOTINST_TOKEN_GCP"),
		Account: os.Getenv("SPOTINST_ACCOUNT_GCP"),
	}

	return config.Client()
}

func providerConfigureAWS(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:   os.Getenv("SPOTINST_TOKEN_AWS"),
		Account: os.Getenv("SPOTINST_ACCOUNT_AWS"),
	}

	return config.Client()
}

func providerConfigureAzure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:   os.Getenv("SPOTINST_TOKEN_AZURE"),
		Account: os.Getenv("SPOTINST_ACCOUNT_AZURE"),
	}

	return config.Client()
}

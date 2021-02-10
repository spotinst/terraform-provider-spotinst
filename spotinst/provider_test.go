package spotinst

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var TestAccProviders map[string]terraform.ResourceProvider

var testAccProviderGCP *schema.Provider
var testAccProviderAWS *schema.Provider
var testAccProviderAzure *schema.Provider
var testAccProviderAzureV3 *schema.Provider

func init() {
	testAccProviderGCP = Provider().(*schema.Provider)
	testAccProviderAWS = Provider().(*schema.Provider)
	testAccProviderAzure = Provider().(*schema.Provider)
	testAccProviderAzureV3 = Provider().(*schema.Provider)

	testAccProviderGCP.ConfigureFunc = providerConfigureGCP
	testAccProviderAWS.ConfigureFunc = providerConfigureAWS
	testAccProviderAzure.ConfigureFunc = providerConfigureAzure
	testAccProviderAzureV3.ConfigureFunc = providerConfigureAzureV3

	TestAccProviders = map[string]terraform.ResourceProvider{
		"gcp":     testAccProviderGCP,
		"aws":     testAccProviderAWS,
		"azure":   testAccProviderAzure,
		"azureV3": testAccProviderAzureV3,
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
		"gcp":     os.Getenv("SPOTINST_TOKEN_GCP"),
		"aws":     os.Getenv("SPOTINST_TOKEN_AWS"),
		"azure":   os.Getenv("SPOTINST_TOKEN_AZURE"),
		"azureV3": "eeab5e1e5e9b5dcbb1aba6d7023d2ae981c6b48dd13784439bb6061f8beb053a",
	}
	fmt.Printf(tokens["azureV3"])
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

func providerConfigureAzureV3(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:   "eeab5e1e5e9b5dcbb1aba6d7023d2ae981c6b48dd13784439bb6061f8beb053a",
		Account: "act-e929c6e7",
	}

	return config.Client()
}

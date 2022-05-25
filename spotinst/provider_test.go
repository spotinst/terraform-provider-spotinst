package spotinst

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviderGCP *schema.Provider
var testAccProviderAWS *schema.Provider
var testAccProviderAzure *schema.Provider
var testAccProviderAzureV3 *schema.Provider

var TestAccProviders map[string]func() (*schema.Provider, error)

func testAccProviderGCPFunc() (*schema.Provider, error) {
	return testAccProviderGCP, nil
}

func testAccProviderAWSFunc() (*schema.Provider, error) {
	return testAccProviderAWS, nil
}

func testAccProviderAzureFunc() (*schema.Provider, error) {
	return testAccProviderAzure, nil
}

func testAccProviderAzureV3Func() (*schema.Provider, error) {
	return testAccProviderAzureV3, nil
}

func init() {
	testAccProviderGCP = Provider()
	testAccProviderAWS = Provider()
	testAccProviderAzure = Provider()
	testAccProviderAzureV3 = Provider()

	testAccProviderGCP.ConfigureContextFunc = providerConfigureGCP
	testAccProviderAWS.ConfigureContextFunc = providerConfigureAWS
	testAccProviderAzure.ConfigureContextFunc = providerConfigureAzure
	testAccProviderAzureV3.ConfigureContextFunc = providerConfigureAzure

	TestAccProviders = map[string]func() (*schema.Provider, error){
		"gcp":     testAccProviderGCPFunc,
		"aws":     testAccProviderAWSFunc,
		"azure":   testAccProviderAzureFunc,
		"azurev3": testAccProviderAzureV3Func,
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
		"gcp":   os.Getenv("SPOTINST_TOKEN_GCP"),
		"aws":   os.Getenv("SPOTINST_TOKEN_AWS"),
		"azure": os.Getenv("SPOTINST_TOKEN_AZURE"),
	}

	if tokens[provider] == "" {
		t.Fatal(ErrNoValidCredentials.Error())
	}
}

func providerConfigureGCP(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		Token:   os.Getenv("SPOTINST_TOKEN_GCP"),
		Account: os.Getenv("SPOTINST_ACCOUNT_GCP"),
	}

	res, diagnostics := config.ClientV2()
	return res, diagnostics
}

func providerConfigureAWS(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		Token:   os.Getenv("SPOTINST_TOKEN_AWS"),
		Account: os.Getenv("SPOTINST_ACCOUNT_AWS"),
	}

	res, diagnostics := config.ClientV2()
	return res, diagnostics
}

func providerConfigureAzure(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		Token:   os.Getenv("SPOTINST_TOKEN_AZURE"),
		Account: os.Getenv("SPOTINST_ACCOUNT_AZURE"),
	}

	res, diagnostics := config.ClientV2()
	return res, diagnostics
}

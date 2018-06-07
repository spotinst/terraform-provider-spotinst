package spotinst

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

var TestAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

var testProviders map[string]terraform.ResourceProvider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	TestAccProviders = map[string]terraform.ResourceProvider{
		"spotinst": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func TestAccPreCheck(t *testing.T) {
	c := map[string]string{
		string(commons.ProviderToken): os.Getenv(credentials.EnvCredentialsVarToken),
	}
	if c[string(commons.ProviderToken)] == "" {
		t.Fatal(ErrNoValidCredentials.Error())
	}
}

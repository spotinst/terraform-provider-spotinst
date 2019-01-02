package spotinst

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			string(commons.ProviderToken): &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(credentials.EnvCredentialsVarToken, ""),
				Description: "Spotinst Personal API Access Token",
			},

			string(commons.ProviderAccount): &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(credentials.EnvCredentialsVarAccount, ""),
				Description: "Spotinst Account ID",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			string(commons.ElastigroupAwsResourceName):          resourceSpotinstElastigroupAws(),
			string(commons.SubscriptionResourceName):            resourceSpotinstSubscription(),
			string(commons.ElastigroupAWSBeanstalkResourceName): resourceSpotinstElastigroupAWSBeanstalk(),
			string(commons.OceanAWSResourceName):                resourceSpotinstOceanAWS(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:   d.Get(string(commons.ProviderToken)).(string),
		Account: d.Get(string(commons.ProviderAccount)).(string),
	}
	if err := config.Validate(); err != nil {
		return nil, err
	}
	return config.Client()
}

package spotinst

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			string(commons.ProviderToken): {
				Type:     schema.TypeString,
				Optional: true,
				//DefaultFunc: schema.EnvDefaultFunc(credentials.EnvCredentialsVarToken, ""),
				Description: "Spotinst Personal API Access Token",
			},

			string(commons.ProviderAccount): {
				Type:     schema.TypeString,
				Optional: true,
				//DefaultFunc: schema.EnvDefaultFunc(credentials.EnvCredentialsVarAccount, ""),
				Description: "Spotinst Account ID",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			// Elastigroup.
			string(commons.ElastigroupAWSResourceName):          resourceSpotinstElastigroupAWS(),
			string(commons.ElastigroupGCPResourceName):          resourceSpotinstElastigroupGCP(),
			string(commons.ElastigroupGKEResourceName):          resourceSpotinstElastigroupGKE(),
			string(commons.ElastigroupAWSBeanstalkResourceName): resourceSpotinstElastigroupAWSBeanstalk(),
			string(commons.ElastigroupAzureResourceName):        resourceSpotinstElastigroupAzure(),
			string(commons.SubscriptionResourceName):            resourceSpotinstSubscription(),
			string(commons.MRScalerAWSResourceName):             resourceSpotinstMRScalerAWS(),

			// Ocean.
			string(commons.OceanAWSResourceName):                 resourceSpotinstOceanAWS(),
			string(commons.OceanAWSLaunchSpecResourceName):       resourceSpotinstOceanAWSLaunchSpec(),
			string(commons.OceanGKEImportResourceName):           resourceSpotinstOceanGKEImport(),
			string(commons.OceanGKELaunchSpecResourceName):       resourceSpotinstOceanGKELaunchSpec(),
			string(commons.OceanGKELaunchSpecImportResourceName): resourceSpotinstOceanGKELaunchSpecImport(),
			string(commons.OceanECSResourceName):                 resourceSpotinstOceanECS(),
			string(commons.OceanECSLaunchSpecResourceName):       resourceSpotinstOceanECSLaunchSpec(),

			// Multai.
			string(commons.MultaiBalancerResourceName):    resourceSpotinstMultaiBalancer(),
			string(commons.MultaiDeploymentResourceName):  resourceSpotinstMultaiDeployment(),
			string(commons.MultaiListenerResourceName):    resourceSpotinstMultaiListener(),
			string(commons.MultaiRoutingRuleResourceName): resourceSpotinstMultaiRoutingRule(),
			string(commons.MultaiTargetResourceName):      resourceSpotinstMultaiTarget(),
			string(commons.MultaiTargetSetResourceName):   resourceSpotinstMultaiTargetSet(),

			// Managed Instance.
			string(commons.ManagedInstanceAWSResourceName): resourceSpotinstMangedInstanceAWS(),
		},
	}

	p.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := p.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}

	return p
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	config := Config{
		Token:            d.Get(string(commons.ProviderToken)).(string),
		Account:          d.Get(string(commons.ProviderAccount)).(string),
		terraformVersion: terraformVersion,
	}

	return config.Client()
}

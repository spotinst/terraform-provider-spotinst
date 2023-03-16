package spotinst

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
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

			string(commons.ProviderFeatureFlags): {
				Type:     schema.TypeString,
				Optional: true,
				//DefaultFunc: schema.EnvDefaultFunc(featureflag.EnvVar, ""),
				Description: "Spotinst SDK Feature Flags",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			// Elastigroup.
			string(commons.ElastigroupAWSResourceName):          resourceSpotinstElastigroupAWS(),
			string(commons.ElastigroupGCPResourceName):          resourceSpotinstElastigroupGCP(),
			string(commons.ElastigroupGKEResourceName):          resourceSpotinstElastigroupGKE(),
			string(commons.ElastigroupAWSBeanstalkResourceName): resourceSpotinstElastigroupAWSBeanstalk(),
			string(commons.ElastigroupAzureResourceName):        resourceSpotinstElastigroupAzure(),
			string(commons.ElastigroupAzureV3ResourceName):      resourceSpotinstElastigroupAzureV3(),
			string(commons.SubscriptionResourceName):            resourceSpotinstSubscription(),
			string(commons.MRScalerAWSResourceName):             resourceSpotinstMRScalerAWS(),

			// Ocean.
			string(commons.OceanAWSResourceName):                   resourceSpotinstOceanAWS(),
			string(commons.OceanAWSLaunchSpecResourceName):         resourceSpotinstOceanAWSLaunchSpec(),
			string(commons.OceanGKEImportResourceName):             resourceSpotinstOceanGKEImport(),
			string(commons.OceanGKELaunchSpecResourceName):         resourceSpotinstOceanGKELaunchSpec(),
			string(commons.OceanGKELaunchSpecImportResourceName):   resourceSpotinstOceanGKELaunchSpecImport(),
			string(commons.OceanECSResourceName):                   resourceSpotinstOceanECS(),
			string(commons.OceanECSLaunchSpecResourceName):         resourceSpotinstOceanECSLaunchSpec(),
			string(commons.OceanAKSResourceName):                   resourceSpotinstOceanAKS(),
			string(commons.OceanAKSVirtualNodeGroupResourceName):   resourceSpotinstOceanAKSVirtualNodeGroup(),
			string(commons.OceanAKSNPResourceName):                 resourceSpotinstOceanAKSNP(),
			string(commons.OceanAKSNPVirtualNodeGroupResourceName): resourceSpotinstOceanAKSNPVirtualNodeGroup(),

			// Multai.
			string(commons.MultaiBalancerResourceName):    resourceSpotinstMultaiBalancer(),
			string(commons.MultaiDeploymentResourceName):  resourceSpotinstMultaiDeployment(),
			string(commons.MultaiListenerResourceName):    resourceSpotinstMultaiListener(),
			string(commons.MultaiRoutingRuleResourceName): resourceSpotinstMultaiRoutingRule(),
			string(commons.MultaiTargetResourceName):      resourceSpotinstMultaiTarget(),
			string(commons.MultaiTargetSetResourceName):   resourceSpotinstMultaiTargetSet(),

			// Managed Instance.
			string(commons.ManagedInstanceAWSResourceName): resourceSpotinstMangedInstanceAWS(),

			// HealthCheck
			string(commons.HealthCheckResourceName): resourceSpotinstHealthCheck(),

			// SuspendProcesses
			string(commons.SuspendProcessesResourceName): resourceSpotinstElastigroupSuspendProcesses(),

			// ExtendedResourceDefinition
			string(commons.OceanAWSExtendedResourceDefinitionResourceName): resourceSpotinstOceanAWSExtendedResourceDefinition(),

			// Data Integration
			string(commons.DataIntegrationResourceName): resourceSpotinstDataIntegration(),

			// Stateful
			string(commons.StatefulNodeAzureResourceName): resourceSpotinstStatefulNodeAzureV3(),

			// Ocean Spark
			string(commons.OceanSparkResourceName): resourceSpotinstOceanSpark(),

			// Ocean Spark Virtual Node Group
			string(commons.OceanSparkVirtualNodeGroupResourceName): resourceSpotinstOceanSparkVirtualNodeGroup(),
		},
	}

	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
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

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, diag.Diagnostics) {
	config := Config{
		Token:            d.Get(string(commons.ProviderToken)).(string),
		Account:          d.Get(string(commons.ProviderAccount)).(string),
		FeatureFlags:     d.Get(string(commons.ProviderFeatureFlags)).(string),
		terraformVersion: terraformVersion,
	}

	return config.Client()
}

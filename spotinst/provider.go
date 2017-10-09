package spotinst

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(credentials.EnvCredentialsVarToken, ""),
				Description: "Spotinst Personal API Access Token",
			},

			"account": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(credentials.EnvCredentialsVarAccount, ""),
				Description: "Spotinst Account ID",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"spotinst_aws_group":           resourceSpotinstAWSGroup(), // deprecated
			"spotinst_group_aws":           resourceSpotinstAWSGroup(),
			"spotinst_mrscaler":            resourceSpotinstAWSMrScaler(),
			"spotinst_subscription":        resourceSpotinstSubscription(),
			"spotinst_healthcheck":         resourceSpotinstHealthCheck(),
			"spotinst_multai_deployment":   resourceSpotinstMultaiDeployment(),
			"spotinst_multai_balancer":     resourceSpotinstMultaiBalancer(),
			"spotinst_multai_listener":     resourceSpotinstMultaiListener(),
			"spotinst_multai_routing_rule": resourceSpotinstMultaiRoutingRule(),
			"spotinst_multai_target_set":   resourceSpotinstMultaiTargetSet(),
			"spotinst_multai_target":       resourceSpotinstMultaiTarget(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:   d.Get("token").(string),
		Account: d.Get("account").(string),
	}
	if err := config.Validate(); err != nil {
		return nil, err
	}
	return config.Client()
}

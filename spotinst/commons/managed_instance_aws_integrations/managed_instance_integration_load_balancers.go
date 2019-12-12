package managed_instance_aws_integrations

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

func SetupLoadBalancers(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[LoadBalancers] = commons.NewGenericField(
		commons.ManagedInstanceAWS,
		LoadBalancers,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(LoadBalancerName): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(Arn): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(BalancerID): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(TargetSetID): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(AzAwareness): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(AutoWeight): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(Type): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			mi := egWrapper.GetManagedInstance()
			if v, ok := resourceData.GetOk(string(LoadBalancers)); ok {
				if balancersConfig, err := expandLoadBalancers(v); err != nil {
					return err
				} else if len(balancersConfig.LoadBalancers) > 0 {
					mi.Integration.SetLoadBalancersConfig(balancersConfig)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			mi := egWrapper.GetManagedInstance()

			if v, ok := resourceData.GetOk(string(LoadBalancers)); ok {
				if balancersConfig, err := expandLoadBalancers(v); err != nil {
					return err
				} else if len(balancersConfig.LoadBalancers) > 0 {
					mi.Integration.SetLoadBalancersConfig(balancersConfig)
				}
			}
			return nil
		},
		nil,
	)

}

//
////-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
////            Utils
////-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func expandLoadBalancers(data interface{}) (*aws.LoadBalancersConfig, error) {
	list := data.(*schema.Set).List()
	balancers := make([]*aws.LoadBalancer, 0, len(list))
	lbConfig := &aws.LoadBalancersConfig{}

	for _, item := range list {
		m := item.(map[string]interface{})
		loadBalancer := &aws.LoadBalancer{}

		if v, ok := m[string(LoadBalancerName)].(string); ok && v != "" {
			loadBalancer.SetName(spotinst.String(v))
		}

		if v, ok := m[string(Arn)].(string); ok && v != "" {
			loadBalancer.SetArn(spotinst.String(v))
		}

		if v, ok := m[string(BalancerID)].(string); ok && v != "" {
			loadBalancer.SetBalancerId(spotinst.String(v))
		}

		if v, ok := m[string(TargetSetID)].(string); ok && v != "" {
			loadBalancer.SetTargetSetId(spotinst.String(v))
		}

		if v, ok := m[string(AzAwareness)].(bool); ok && v {
			loadBalancer.SetZoneAwareness(spotinst.Bool(v))
		}

		if v, ok := m[string(AutoWeight)].(bool); ok && v {
			loadBalancer.SetAutoWeight(spotinst.Bool(v))
		}

		if v, ok := m[string(Type)].(string); ok && v != "" {
			loadBalancer.SetType(spotinst.String(v))
		}
		balancers = append(balancers, loadBalancer)
	}

	lbConfig.SetLoadBalancers(balancers)

	return lbConfig, nil
}

type CreateBalancerObjFunc func(id string) (*aws.LoadBalancer, error)

var TargetGroupArnRegex = regexp.MustCompile(`arn:aws:elasticloadbalancing:.*:\d{12}:targetgroup/(.*)/.*`)

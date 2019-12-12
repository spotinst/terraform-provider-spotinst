package elastigroup_azure_load_balancer

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[LoadBalancer] = commons.NewGenericField(
		commons.ElastigroupAzureLoadBalancers,
		LoadBalancer,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Type): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(BalancerID): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(TargetSetID): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(AutoWeight): {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(LoadBalancer)); ok {
				if balancers, err := expandAzureGroupLoadBalancers(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(balancers)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []*azure.LoadBalancer = nil

			if v, ok := resourceData.GetOk(string(LoadBalancer)); ok {
				if balancers, err := expandAzureGroupLoadBalancers(v); err != nil {
					return err
				} else {
					value = balancers
				}
			}
			elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(value)
			return nil
		},
		nil,
	)

}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func flattenAzureGroupLoadBalancers(loadBalancers []*azure.LoadBalancer) []interface{} {
	result := make([]interface{}, 0, len(loadBalancers))
	for _, balancer := range loadBalancers {
		m := make(map[string]interface{})
		m[string(Type)] = spotinst.StringValue(balancer.Type)
		m[string(BalancerID)] = spotinst.StringValue(balancer.BalancerID)
		m[string(TargetSetID)] = spotinst.StringValue(balancer.TargetSetID)
		m[string(AutoWeight)] = spotinst.BoolValue(balancer.AutoWeight)

		result = append(result, m)
	}
	return result
}

func expandAzureGroupLoadBalancers(data interface{}) ([]*azure.LoadBalancer, error) {
	list := data.(*schema.Set).List()
	balancers := make([]*azure.LoadBalancer, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		loadBalancer := &azure.LoadBalancer{}

		if v, ok := m[string(Type)].(string); ok && v != "" {
			loadBalancer.SetType(spotinst.String(v))
		}

		if v, ok := m[string(BalancerID)].(string); ok && v != "" {
			loadBalancer.SetBalancerId(spotinst.String(v))
		}

		if v, ok := m[string(TargetSetID)].(string); ok && v != "" {
			loadBalancer.SetTargetSetId(spotinst.String(v))
		}

		if v, ok := m[string(AutoWeight)].(bool); ok {
			loadBalancer.SetAutoWeight(spotinst.Bool(v))
		}

		balancers = append(balancers, loadBalancer)
	}
	return balancers, nil
}

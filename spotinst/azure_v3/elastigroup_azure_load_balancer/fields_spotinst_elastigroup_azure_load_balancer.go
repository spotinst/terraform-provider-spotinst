package elastigroup_azure_load_balancer

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[LoadBalancer] = commons.NewGenericField(
		commons.ElastigroupAzureLoadBalancer,
		LoadBalancer,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(BackendPoolNames): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString},
					},
					string(Name): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(SKU): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(ResourceGroupName): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(Type): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup != nil && elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil && elastigroup.Compute.LaunchSpecification.LoadBalancersConfig != nil && elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers != nil {
				loadBalancers := elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers
				result = flattenLoadBalancers(loadBalancers)
			}

			if result != nil {
				if err := resourceData.Set(string(LoadBalancer), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(LoadBalancer), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()

			if value, ok := resourceData.GetOk(string(LoadBalancer)); ok {
				if loadBalancers, err := expandLoadBalancers(value); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(loadBalancers)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []*azurev3.LoadBalancer = nil

			if v, ok := resourceData.GetOk(string(LoadBalancer)); ok {
				if loadBalancers, err := expandLoadBalancers(v); err != nil {
					return err
				} else {
					value = loadBalancers
				}
			}
			if len(value) == 0 {
				elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(nil)
			} else {
				elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(value)
			}

			return nil
		},
		nil,
	)
}

func flattenLoadBalancers(loadBalancers []*azurev3.LoadBalancer) []interface{} {
	result := make([]interface{}, 0, len(loadBalancers))

	for _, loadBalancer := range loadBalancers {
		m := make(map[string]interface{})
		if loadBalancer.BackendPoolNames != nil {
			m[string(BackendPoolNames)] = spotinst.StringSlice(loadBalancer.BackendPoolNames)
		}
		m[string(SKU)] = spotinst.StringValue(loadBalancer.SKU)
		m[string(Name)] = spotinst.StringValue(loadBalancer.Name)
		m[string(ResourceGroupName)] = spotinst.StringValue(loadBalancer.ResourceGroupName)
		m[string(Type)] = spotinst.StringValue(loadBalancer.Type)

		result = append(result, m)
	}
	return result
}

func expandLoadBalancers(data interface{}) ([]*azurev3.LoadBalancer, error) {
	list := data.(*schema.Set).List()
	loadBalancers := make([]*azurev3.LoadBalancer, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		loadBalancer := &azurev3.LoadBalancer{}

		if v, ok := attr[string(SKU)].(string); ok && v != "" {
			loadBalancer.SetSKU(spotinst.String(v))
		}
		if v, ok := attr[string(Name)].(string); ok && v != "" {
			loadBalancer.SetName(spotinst.String(v))
		}
		if v, ok := attr[string(ResourceGroupName)].(string); ok && v != "" {
			loadBalancer.SetResourceGroupName(spotinst.String(v))
		}
		if v, ok := attr[string(Type)].(string); ok && v != "" {
			loadBalancer.SetType(spotinst.String(v))
		}
		if v, ok := attr[string(BackendPoolNames)]; ok {
			if backendPoolNames, err := expandBackendPoolNames(v); err != nil {
				return nil, err
			} else {
				loadBalancer.SetBackendPoolNames(backendPoolNames)
			}
		}

		loadBalancers = append(loadBalancers, loadBalancer)
	}
	return loadBalancers, nil
}

func expandBackendPoolNames(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if backendPoolName, ok := v.(string); ok && backendPoolName != "" {
			result = append(result, backendPoolName)
		}
	}
	return result, nil
}

package stateful_node_azure_load_balancer

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[LoadBalancer] = commons.NewGenericField(
		commons.StatefulNodeAzureLoadBalancer,
		LoadBalancer,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(BackendPoolNames): {
						Type:     schema.TypeList,
						Required: true,
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
						Computed: true,
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
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var result []interface{} = nil
			if st != nil && st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.LoadBalancersConfig != nil && st.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers != nil {
				loadBalancers := st.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers
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
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []*azure.LoadBalancer = nil

			if v, ok := resourceData.GetOk(string(LoadBalancer)); ok {
				var loadBalancers []*azure.LoadBalancer

				if st != nil && st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.LoadBalancersConfig != nil && st.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers != nil {
					loadBalancers = st.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers

					if loadBalancers, err := expandLoadBalancers(v, loadBalancers); err != nil {
						return err
					} else {
						value = loadBalancers
					}

					st.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(value)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []*azure.LoadBalancer = nil

			if v, ok := resourceData.GetOk(string(LoadBalancer)); ok {
				var loadBalancers []*azure.LoadBalancer

				if st != nil && st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.LoadBalancersConfig != nil && st.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers != nil {
					loadBalancers = st.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers

					if loadBalancers, err := expandLoadBalancers(v, loadBalancers); err != nil {
						return err
					} else {
						value = loadBalancers
					}

					st.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(value)
				}
			}

			return nil
		},
		nil,
	)
}

func flattenLoadBalancers(loadBalancers []*azure.LoadBalancer) []interface{} {
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

func expandLoadBalancers(data interface{}, loadBalancers []*azure.LoadBalancer) ([]*azure.LoadBalancer, error) {
	list := data.(*schema.Set).List()

	if len(list) > 0 {
		loadBalancers = make([]*azure.LoadBalancer, 0, len(list))

		for _, v := range list {
			attr, ok := v.(map[string]interface{})

			if !ok {
				continue
			}

			loadBalancer := &azure.LoadBalancer{}

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

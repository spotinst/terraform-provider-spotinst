package ocean_aks_load_balancers

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[LoadBalancer] = commons.NewGenericField(
		commons.OceanAKSLoadBalancers,
		LoadBalancer,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(BackendPoolNames): {
						Type:     schema.TypeList,
						Optional: true,
						Computed: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(LoadBalancerSKU): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},

					string(Name): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},

					string(ResourceGroupName): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},

					string(Type): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification.LoadBalancersConfig != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification.LoadBalancersConfig.LoadBalancers != nil {
				loadBalancers := cluster.VirtualNodeGroupTemplate.LaunchSpecification.LoadBalancersConfig.LoadBalancers
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
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value []*azure.LoadBalancer = nil

			if v, ok := resourceData.GetOk(string(LoadBalancer)); ok {
				//create new image object in case cluster did not get it from previous import step.
				var loadBalancers []*azure.LoadBalancer

				if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification.LoadBalancersConfig != nil {

					if cluster.VirtualNodeGroupTemplate.LaunchSpecification.LoadBalancersConfig.LoadBalancers != nil {
						loadBalancers = cluster.VirtualNodeGroupTemplate.LaunchSpecification.LoadBalancersConfig.LoadBalancers
					}

					if loadBalancers, err := expandLoadBalancers(v, loadBalancers); err != nil {
						return err
					} else {
						value = loadBalancers
					}

					cluster.VirtualNodeGroupTemplate.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(value)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value []*azure.LoadBalancer = nil

			if v, ok := resourceData.GetOk(string(LoadBalancer)); ok {
				//create new image object in case cluster did not get it from previous import step.
				var loadBalancers []*azure.LoadBalancer

				if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification.LoadBalancersConfig != nil {

					if cluster.VirtualNodeGroupTemplate.LaunchSpecification.LoadBalancersConfig.LoadBalancers != nil {
						loadBalancers = cluster.VirtualNodeGroupTemplate.LaunchSpecification.LoadBalancersConfig.LoadBalancers
					}

					if loadBalancers, err := expandLoadBalancers(v, loadBalancers); err != nil {
						return err
					} else {
						value = loadBalancers
					}

					cluster.VirtualNodeGroupTemplate.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(value)
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
		m[string(LoadBalancerSKU)] = spotinst.StringValue(loadBalancer.LoadBalancerSKU)
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

			if v, ok := attr[string(LoadBalancerSKU)].(string); ok && v != "" {
				loadBalancer.SetLoadBalancerSKU(spotinst.String(v))
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
					loadBalancer.SeBackendPoolNames(backendPoolNames)
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

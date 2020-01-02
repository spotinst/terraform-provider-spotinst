package ocean_gke_instance_types

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Whitelist] = commons.NewGenericField(
		commons.OceanGKEInstanceTypes,
		Whitelist,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []string
			if cluster.Compute != nil && cluster.Compute.InstanceTypes != nil &&
				cluster.Compute.InstanceTypes.Whitelist != nil {
				result = append(result, cluster.Compute.InstanceTypes.Whitelist...)
			}
			if err := resourceData.Set(string(Whitelist), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Whitelist), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(Whitelist)); ok {
				instances := v.([]interface{})
				instanceTypes := make([]string, len(instances))
				for i, j := range instances {
					instanceTypes[i] = j.(string)
				}
				cluster.Compute.InstanceTypes.SetWhitelist(instanceTypes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var instanceTypes []string = nil
			if v, ok := resourceData.GetOk(string(Whitelist)); ok {
				instances := v.([]interface{})
				instanceTypes = make([]string, len(instances))
				for i, v := range instances {
					instanceTypes[i] = v.(string)
				}
			}
			cluster.Compute.InstanceTypes.SetWhitelist(instanceTypes)
			return nil
		},
		nil,
	)

}

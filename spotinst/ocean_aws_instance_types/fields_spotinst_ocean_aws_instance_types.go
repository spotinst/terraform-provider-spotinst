package ocean_aws_instance_types

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
		commons.OceanAWSInstanceTypes,
		Whitelist,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []string = nil
			if cluster.Compute != nil && cluster.Compute.InstanceTypes != nil &&
				cluster.Compute.InstanceTypes.Whitelist != nil {
				result = cluster.Compute.InstanceTypes.Whitelist
			}
			if err := resourceData.Set(string(Whitelist), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Whitelist), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(Whitelist)); ok {
				if whitelist, err := expandInstanceTypeList(v); err != nil {
					return err
				} else {
					cluster.Compute.InstanceTypes.SetWhitelist(whitelist)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(Whitelist)); ok {
				if whitelist, err := expandInstanceTypeList(v); err != nil {
					return err
				} else {
					cluster.Compute.InstanceTypes.SetWhitelist(whitelist)
				}
			} else {
				cluster.Compute.InstanceTypes.SetWhitelist(nil)
			}

			return nil
		},
		nil,
	)

	fieldsMap[Blacklist] = commons.NewGenericField(
		commons.OceanAWSInstanceTypes,
		Blacklist,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []string = nil
			if cluster.Compute != nil && cluster.Compute.InstanceTypes != nil &&
				cluster.Compute.InstanceTypes.Blacklist != nil {
				result = cluster.Compute.InstanceTypes.Blacklist
			}
			if err := resourceData.Set(string(Blacklist), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Blacklist), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(Blacklist)); ok {
				if whitelist, err := expandInstanceTypeList(v); err != nil {
					return err
				} else {
					cluster.Compute.InstanceTypes.SetBlacklist(whitelist)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(Blacklist)); ok {
				if whitelist, err := expandInstanceTypeList(v); err != nil {
					return err
				} else {
					cluster.Compute.InstanceTypes.SetBlacklist(whitelist)
				}
			} else {
				cluster.Compute.InstanceTypes.SetBlacklist(nil)
			}

			return nil
		},
		nil,
	)
}

func expandInstanceTypeList(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if instanceTypeList, ok := v.(string); ok && instanceTypeList != "" {
			result = append(result, instanceTypeList)
		}
	}
	return result, nil
}

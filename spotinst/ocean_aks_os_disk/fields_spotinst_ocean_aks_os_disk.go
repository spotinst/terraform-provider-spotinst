package ocean_aks_os_disk

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[OSDisk] = commons.NewGenericField(
		commons.OceanAKSLaunchSpecification,
		OSDisk,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(SizeGB): {
						Type:     schema.TypeInt,
						Required: true,
					},
					string(Type): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value interface{} = nil

			if cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil &&
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.OSDisk != nil {
				value = flattenOSDisk(cluster.VirtualNodeGroupTemplate.LaunchSpecification.OSDisk)
			}
			if err := resourceData.Set(string(OSDisk), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OSDisk), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *azure.OSDisk = nil

			if v, ok := resourceData.GetOk(string(OSDisk)); ok {
				if osDisk, err := expandOSDisk(v); err != nil {
					return err
				} else {
					value = osDisk
				}
			}
			cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetOSDisk(value)

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *azure.OSDisk = nil

			if v, ok := resourceData.GetOk(string(OSDisk)); ok {
				if osDisk, err := expandOSDisk(v); err != nil {
					return err
				} else {
					value = osDisk
				}
			}
			cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetOSDisk(value)
			return nil
		},
		nil,
	)
}

func flattenOSDisk(osd *azure.OSDisk) interface{} {
	osDisk := make(map[string]interface{})
	osDisk[string(SizeGB)] = spotinst.IntValue(osd.SizeGB)
	osDisk[string(Type)] = spotinst.StringValue(osd.Type)
	return []interface{}{osDisk}
}

func expandOSDisk(data interface{}) (*azure.OSDisk, error) {
	if list := data.([]interface{}); len(list) > 0 {
		osDisk := &azure.OSDisk{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})
			var sizeGB *int = nil
			var osType *string = nil

			if v, ok := m[string(SizeGB)].(int); ok && v > 0 {
				sizeGB = spotinst.Int(v)
			}
			osDisk.SetSizeGB(sizeGB)

			if v, ok := m[string(Type)].(string); ok && v != "" {
				osType = spotinst.String(v)
			}
			osDisk.SetType(osType)

		}

		return osDisk, nil
	}

	return nil, nil
}

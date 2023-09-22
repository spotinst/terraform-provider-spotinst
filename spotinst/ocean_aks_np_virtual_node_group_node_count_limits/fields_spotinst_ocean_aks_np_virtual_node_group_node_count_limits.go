package ocean_aks_np_virtual_node_group_node_count_limits

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[MinCount] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroupNodeCountLimits,
		MinCount,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			var value *int = nil
			if virtualNodeGroup != nil && virtualNodeGroup.NodeCountLimits != nil && virtualNodeGroup.NodeCountLimits.MinCount != nil {
				value = virtualNodeGroup.NodeCountLimits.MinCount
			} else {
				value = spotinst.Int(-1)
			}
			if err := resourceData.Set(string(MinCount), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MinCount), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.Get(string(MinCount)).(int); ok && v > -1 {
				virtualNodeGroup.NodeCountLimits.SetMinCount(spotinst.Int(v))
			} else {
				virtualNodeGroup.NodeCountLimits.SetMinCount(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.Get(string(MinCount)).(int); ok && v > -1 {
				virtualNodeGroup.NodeCountLimits.SetMinCount(spotinst.Int(v))
			} else {
				virtualNodeGroup.NodeCountLimits.SetMinCount(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[MaxCount] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroupNodeCountLimits,
		MaxCount,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			var value *int = nil
			if virtualNodeGroup != nil && virtualNodeGroup.NodeCountLimits != nil && virtualNodeGroup.NodeCountLimits.MaxCount != nil {
				value = virtualNodeGroup.NodeCountLimits.MaxCount
			} else {
				value = spotinst.Int(-1)
			}
			if err := resourceData.Set(string(MaxCount), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MaxCount), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.Get(string(MaxCount)).(int); ok && v > -1 {
				virtualNodeGroup.NodeCountLimits.SetMaxCount(spotinst.Int(v))
			} else {
				virtualNodeGroup.NodeCountLimits.SetMaxCount(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.Get(string(MaxCount)).(int); ok && v > -1 {
				virtualNodeGroup.NodeCountLimits.SetMaxCount(spotinst.Int(v))
			} else {
				virtualNodeGroup.NodeCountLimits.SetMaxCount(nil)
			}
			return nil
		},
		nil,
	)
}

package ocean_aks_np_virtual_node_group_strategy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[SpotPercentage] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroupStrategy,
		SpotPercentage,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			var value *int = nil
			if virtualNodeGroup != nil && virtualNodeGroup.Strategy != nil {
				value = virtualNodeGroup.Strategy.SpotPercentage
			} else {
				value = spotinst.Int(-1)
			}
			if value != nil {
				if err := resourceData.Set(string(SpotPercentage), spotinst.IntValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SpotPercentage), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.Get(string(SpotPercentage)).(int); ok && v >= 0 {
				virtualNodeGroup.Strategy.SetSpotPercentage(spotinst.Int(v))
			} else {
				virtualNodeGroup.Strategy.SetSpotPercentage(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.Get(string(SpotPercentage)).(int); ok && v > -1 {
				virtualNodeGroup.Strategy.SetSpotPercentage(spotinst.Int(v))
			} else {
				virtualNodeGroup.Strategy.SetSpotPercentage(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[FallbackToOnDemand] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroupStrategy,
		FallbackToOnDemand,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			var value *bool = nil
			if virtualNodeGroup.Strategy != nil && virtualNodeGroup.Strategy.FallbackToOD != nil {
				value = virtualNodeGroup.Strategy.FallbackToOD
			}
			if value != nil {
				if err := resourceData.Set(string(FallbackToOnDemand), spotinst.BoolValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(FallbackToOnDemand), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()

			if v, ok := resourceData.GetOkExists(string(FallbackToOnDemand)); ok && v != nil {
				ftod := v.(bool)
				fallback := spotinst.Bool(ftod)
				virtualNodeGroup.Strategy.SetFallbackToOD(fallback)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			var fallback *bool = nil
			if v, ok := resourceData.GetOkExists(string(FallbackToOnDemand)); ok && v != nil {
				ftod := v.(bool)
				fallback = spotinst.Bool(ftod)
			}
			virtualNodeGroup.Strategy.SetFallbackToOD(fallback)
			return nil
		},
		nil,
	)
}

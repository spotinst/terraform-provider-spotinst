package ocean_aks_np_virtual_node_group_node_pool_properties

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[MaxPodsPerNode] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroupNodePoolProperties,
		MaxPodsPerNode,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			var value *int = nil
			if virtualNodeGroup != nil && virtualNodeGroup.NodePoolProperties != nil && virtualNodeGroup.NodePoolProperties.MaxPodsPerNode != nil {
				value = virtualNodeGroup.NodePoolProperties.MaxPodsPerNode
			} else {
				value = spotinst.Int(-1)
			}
			if err := resourceData.Set(string(MaxPodsPerNode), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MaxPodsPerNode), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.Get(string(MaxPodsPerNode)).(int); ok && v > -1 {
				virtualNodeGroup.NodePoolProperties.SetMaxPodsPerNode(spotinst.Int(v))
			} else {
				virtualNodeGroup.NodePoolProperties.SetMaxPodsPerNode(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.Get(string(MaxPodsPerNode)).(int); ok && v > -1 {
				virtualNodeGroup.NodePoolProperties.SetMaxPodsPerNode(spotinst.Int(v))
			} else {
				virtualNodeGroup.NodePoolProperties.SetMaxPodsPerNode(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[EnableNodePublicIP] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroupNodePoolProperties,
		EnableNodePublicIP,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			var value *bool = nil
			if virtualNodeGroup.NodePoolProperties != nil && virtualNodeGroup.NodePoolProperties.EnableNodePublicIP != nil {
				value = virtualNodeGroup.NodePoolProperties.EnableNodePublicIP
			}
			if value != nil {
				if err := resourceData.Set(string(EnableNodePublicIP), spotinst.BoolValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(EnableNodePublicIP), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.GetOkExists(string(EnableNodePublicIP)); ok && v != nil {
				publicIp := v.(bool)
				enableIp := spotinst.Bool(publicIp)
				virtualNodeGroup.NodePoolProperties.SetEnableNodePublicIP(enableIp)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			var enableIp *bool = nil
			if v, ok := resourceData.GetOkExists(string(EnableNodePublicIP)); ok && v != nil {
				publicIp := v.(bool)
				enableIp = spotinst.Bool(publicIp)
			}
			virtualNodeGroup.NodePoolProperties.SetEnableNodePublicIP(enableIp)
			return nil
		},
		nil,
	)

	fieldsMap[OsDiskSizeGB] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroupNodePoolProperties,
		OsDiskSizeGB,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			var value *int = nil
			if virtualNodeGroup != nil && virtualNodeGroup.NodePoolProperties != nil && virtualNodeGroup.NodePoolProperties.OsDiskSizeGB != nil {
				value = virtualNodeGroup.NodePoolProperties.OsDiskSizeGB
			} else {
				value = spotinst.Int(-1)
			}
			if err := resourceData.Set(string(OsDiskSizeGB), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OsDiskSizeGB), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.Get(string(OsDiskSizeGB)).(int); ok && v > 0 {
				virtualNodeGroup.NodePoolProperties.SetOsDiskSizeGB(spotinst.Int(v))
			} else {
				virtualNodeGroup.NodePoolProperties.SetOsDiskSizeGB(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.Get(string(OsDiskSizeGB)).(int); ok && v > 0 {
				virtualNodeGroup.NodePoolProperties.SetOsDiskSizeGB(spotinst.Int(v))
			} else {
				virtualNodeGroup.NodePoolProperties.SetOsDiskSizeGB(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[OsDiskType] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroupNodePoolProperties,
		OsDiskType,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if err := resourceData.Set(string(OsDiskType), spotinst.StringValue(virtualNodeGroup.NodePoolProperties.OsDiskType)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(OsDiskType), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.GetOkExists(string(OsDiskType)); ok {
				virtualNodeGroup.NodePoolProperties.SetOsDiskType(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.GetOkExists(string(OsDiskType)); ok {
				virtualNodeGroup.NodePoolProperties.SetOsDiskType(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[OsType] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroupNodePoolProperties,
		OsType,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if err := resourceData.Set(string(OsType), spotinst.StringValue(virtualNodeGroup.NodePoolProperties.OsType)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(OsType), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.GetOkExists(string(OsType)); ok {
				virtualNodeGroup.NodePoolProperties.SetOsType(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.GetOkExists(string(OsType)); ok {
				virtualNodeGroup.NodePoolProperties.SetOsType(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[OsSKU] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroupNodePoolProperties,
		OsSKU,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if err := resourceData.Set(string(OsSKU), spotinst.StringValue(virtualNodeGroup.NodePoolProperties.OsSKU)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(OsSKU), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.GetOkExists(string(OsSKU)); ok {
				virtualNodeGroup.NodePoolProperties.SetOsSKU(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.GetOkExists(string(OsSKU)); ok {
				virtualNodeGroup.NodePoolProperties.SetOsSKU(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[KubernetesVersion] = commons.NewGenericField(
		commons.OceanAKSNPVirtualNodeGroupNodePoolProperties,
		KubernetesVersion,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if err := resourceData.Set(string(KubernetesVersion), spotinst.StringValue(virtualNodeGroup.NodePoolProperties.KubernetesVersion)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(KubernetesVersion), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.GetOkExists(string(KubernetesVersion)); ok {
				virtualNodeGroup.NodePoolProperties.SetKubernetesVersion(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			vngWrapper := resourceObject.(*commons.VirtualNodeGroupAKSNPWrapper)
			virtualNodeGroup := vngWrapper.GetVirtualNodeGroup()
			if v, ok := resourceData.GetOkExists(string(KubernetesVersion)); ok {
				virtualNodeGroup.NodePoolProperties.SetKubernetesVersion(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)
}

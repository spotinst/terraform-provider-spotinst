package ocean_aks_np_node_pool_properties

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[MaxPodsPerNode] = commons.NewGenericField(
		commons.OceanAKSNPProperties,
		MaxPodsPerNode,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *int = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.NodePoolProperties != nil && cluster.VirtualNodeGroupTemplate.NodePoolProperties.MaxPodsPerNode != nil {
				value = cluster.VirtualNodeGroupTemplate.NodePoolProperties.MaxPodsPerNode
			} else {
				value = spotinst.Int(-1)
			}
			if err := resourceData.Set(string(MaxPodsPerNode), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MaxPodsPerNode), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(MaxPodsPerNode)).(int); ok && v > -1 {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetMaxPodsPerNode(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetMaxPodsPerNode(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(MaxPodsPerNode)).(int); ok && v > -1 {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetMaxPodsPerNode(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetMaxPodsPerNode(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[EnableNodePublicIP] = commons.NewGenericField(
		commons.OceanAKSNPProperties,
		EnableNodePublicIP,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *bool = nil
			if cluster.VirtualNodeGroupTemplate.NodePoolProperties != nil && cluster.VirtualNodeGroupTemplate.NodePoolProperties.EnableNodePublicIP != nil {
				value = cluster.VirtualNodeGroupTemplate.NodePoolProperties.EnableNodePublicIP
			}
			if value != nil {
				if err := resourceData.Set(string(EnableNodePublicIP), spotinst.BoolValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(EnableNodePublicIP), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOk(string(EnableNodePublicIP)); ok && v != nil {
				publicIp := v.(bool)
				enableIp := spotinst.Bool(publicIp)
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetEnableNodePublicIP(enableIp)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var enableIp *bool = nil
			if v, ok := resourceData.GetOk(string(EnableNodePublicIP)); ok && v != nil {
				publicIp := v.(bool)
				enableIp = spotinst.Bool(publicIp)
			}
			cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetEnableNodePublicIP(enableIp)
			return nil
		},
		nil,
	)

	fieldsMap[OsDiskSizeGB] = commons.NewGenericField(
		commons.OceanAKSNPProperties,
		OsDiskSizeGB,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *int = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.NodePoolProperties != nil && cluster.VirtualNodeGroupTemplate.NodePoolProperties.OsDiskSizeGB != nil {
				value = cluster.VirtualNodeGroupTemplate.NodePoolProperties.OsDiskSizeGB
			} else {
				value = spotinst.Int(-1)
			}
			if err := resourceData.Set(string(OsDiskSizeGB), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OsDiskSizeGB), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(OsDiskSizeGB)).(int); ok && v > 0 {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetOsDiskSizeGB(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetOsDiskSizeGB(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.Get(string(OsDiskSizeGB)).(int); ok && v > 0 {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetOsDiskSizeGB(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetOsDiskSizeGB(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[OsDiskType] = commons.NewGenericField(
		commons.OceanAKSNPProperties,
		OsDiskType,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if err := resourceData.Set(string(OsDiskType), spotinst.StringValue(cluster.VirtualNodeGroupTemplate.NodePoolProperties.OsDiskType)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(OsDiskType), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOk(string(OsDiskType)); ok {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetOsDiskType(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOk(string(OsDiskType)); ok {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetOsDiskType(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[OsType] = commons.NewGenericField(
		commons.OceanAKSNPProperties,
		OsType,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if err := resourceData.Set(string(OsType), spotinst.StringValue(cluster.VirtualNodeGroupTemplate.NodePoolProperties.OsType)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(OsType), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOk(string(OsType)); ok {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetOsType(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOk(string(OsType)); ok {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetOsType(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[OsSKU] = commons.NewGenericField(
		commons.OceanAKSNPProperties,
		OsSKU,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if err := resourceData.Set(string(OsSKU), spotinst.StringValue(cluster.VirtualNodeGroupTemplate.NodePoolProperties.OsSKU)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(OsSKU), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOk(string(OsSKU)); ok {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetOsSKU(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOk(string(OsSKU)); ok {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetOsSKU(spotinst.String(v.(string)))
			} else {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetOsSKU(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[KubernetesVersion] = commons.NewGenericField(
		commons.OceanAKSNPProperties,
		KubernetesVersion,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if err := resourceData.Set(string(KubernetesVersion), spotinst.StringValue(cluster.VirtualNodeGroupTemplate.NodePoolProperties.KubernetesVersion)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(KubernetesVersion), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOk(string(KubernetesVersion)); ok {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetKubernetesVersion(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOk(string(KubernetesVersion)); ok {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetKubernetesVersion(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)
}

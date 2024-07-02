package ocean_aks_np_node_pool_properties

import (
	"fmt"

	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"

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
			//Default:  false,
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

	fieldsMap[PodSubnetIDs] = commons.NewGenericField(
		commons.OceanAKSNP,
		PodSubnetIDs,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value []string = nil
			if cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.NodePoolProperties != nil {
				value = cluster.VirtualNodeGroupTemplate.NodePoolProperties.PodSubnetIDs
			}
			if err := resourceData.Set(string(PodSubnetIDs), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PodSubnetIDs), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.GetOk(string(PodSubnetIDs)); ok {
				if PodSubnetIds, err := expandSubnetList(value); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetPodSubnetIDs(PodSubnetIds)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.GetOk(string(PodSubnetIDs)); ok {
				if PodSubnetIds, err := expandSubnetList(value); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetPodSubnetIDs(PodSubnetIds)
				}
			} else {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetPodSubnetIDs(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[VnetSubnetIDs] = commons.NewGenericField(
		commons.OceanAKSNP,
		VnetSubnetIDs,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value []string = nil
			if cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.NodePoolProperties != nil {
				value = cluster.VirtualNodeGroupTemplate.NodePoolProperties.VnetSubnetIDs
			}
			if err := resourceData.Set(string(VnetSubnetIDs), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(VnetSubnetIDs), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.GetOk(string(VnetSubnetIDs)); ok {
				if vnetSubnetIds, err := expandSubnetList(value); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetVnetSubnetIDs(vnetSubnetIds)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.GetOk(string(VnetSubnetIDs)); ok {
				if vnetSubnetIds, err := expandSubnetList(value); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetVnetSubnetIDs(vnetSubnetIds)
				}
			} else {
				cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetVnetSubnetIDs(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[LinuxOSConfig] = commons.NewGenericField(
		commons.OceanAKSNPProperties,
		LinuxOSConfig,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Sysctls): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(VmMaxMapCount): {
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value []interface{} = nil
			if cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.NodePoolProperties != nil {
				if cluster.VirtualNodeGroupTemplate.NodePoolProperties.LinuxOSConfig != nil {
					value = flattenLinuxOSConfig(cluster.VirtualNodeGroupTemplate.NodePoolProperties.LinuxOSConfig)
				}
			}
			if len(value) > 0 {
				if err := resourceData.Set(string(LinuxOSConfig), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(LinuxOSConfig), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.GetOk(string(LinuxOSConfig)); ok {
				if config, err := expandLinuxOSConfig(value); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetLinuxOSConfig(config)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var linuxConfig *azure_np.LinuxOSConfig = nil
			if v, ok := resourceData.GetOk(string(LinuxOSConfig)); ok {
				if config, err := expandLinuxOSConfig(v); err != nil {
					return err
				} else {
					linuxConfig = config
				}
			}
			cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetLinuxOSConfig(linuxConfig)
			return nil
		},
		nil,
	)
}

func flattenLinuxOSConfig(linuxConfig *azure_np.LinuxOSConfig) []interface{} {
	var out []interface{}

	if linuxConfig != nil {
		result := make(map[string]interface{})

		if linuxConfig.Sysctls != nil {
			result[string(Sysctls)] = flattenSysctls(linuxConfig.Sysctls)
		}
		out = append(out, result)
	}
	return out
}

func flattenSysctls(sysctls *azure_np.Sysctls) []interface{} {
	var out []interface{}

	if sysctls != nil {
		result := make(map[string]interface{})

		if sysctls.VmMaxMapCount != nil {
			result[string(VmMaxMapCount)] = spotinst.IntValue(sysctls.VmMaxMapCount)
		}
		out = append(out, result)
	}
	return out
}

func expandSubnetList(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if subnetIds, ok := v.(string); ok && subnetIds != "" {
			result = append(result, subnetIds)
		}
	}
	return result, nil
}

func expandLinuxOSConfig(data interface{}) (*azure_np.LinuxOSConfig, error) {
	if list := data.([]interface{}); len(list) > 0 {
		linuxConfig := &azure_np.LinuxOSConfig{}
		if list[0] != nil {
			m := list[0].(map[string]interface{})
			if v, ok := m[string(Sysctls)]; ok {
				sysctls, err := expandSysctls(v)
				if err != nil {
					return nil, err
				}
				if sysctls != nil {
					linuxConfig.SetSysctls(sysctls)
				} else {
					linuxConfig.SetSysctls(nil)
				}
			}
		}
		return linuxConfig, nil
	}
	return nil, nil
}

func expandSysctls(data interface{}) (*azure_np.Sysctls, error) {
	if list := data.([]interface{}); len(list) > 0 {
		sysctls := &azure_np.Sysctls{}
		if list[0] != nil {
			m := list[0].(map[string]interface{})
			if v, ok := m[string(VmMaxMapCount)].(int); ok {
				sysctls.SetVmMaxMapCount(spotinst.Int(v))
			}
		}
		return sysctls, nil
	}
	return nil, nil
}

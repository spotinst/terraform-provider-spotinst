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

	fieldsMap[LocalDnsProfile] = commons.NewGenericField(
		commons.OceanAKSNPProperties,
		LocalDnsProfile,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Mode): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(VnetDNSOverrides): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: dnsOverrideSettingsSchema(),
						},
					},
					string(KubeDNSOverrides): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: dnsOverrideSettingsSchema(),
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
				if cluster.VirtualNodeGroupTemplate.NodePoolProperties.LocalDnsProfile != nil {
					value = flattenLocalDnsProfile(cluster.VirtualNodeGroupTemplate.NodePoolProperties.LocalDnsProfile)
				}
			}
			if len(value) > 0 {
				if err := resourceData.Set(string(LocalDnsProfile), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(LocalDnsProfile), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.GetOk(string(LocalDnsProfile)); ok {
				if profile, err := expandLocalDnsProfile(value); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetLocalDnsProfile(profile)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var profile *azure_np.LocalDnsProfile = nil
			if v, ok := resourceData.GetOk(string(LocalDnsProfile)); ok {
				if p, err := expandLocalDnsProfile(v); err != nil {
					return err
				} else {
					profile = p
				}
			}
			cluster.VirtualNodeGroupTemplate.NodePoolProperties.SetLocalDnsProfile(profile)
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

func flattenLocalDnsProfile(profile *azure_np.LocalDnsProfile) []interface{} {
	var out []interface{}

	if profile != nil {
		result := make(map[string]interface{})

		if profile.Mode != nil {
			result[string(Mode)] = spotinst.StringValue(profile.Mode)
		}
		if profile.VnetDNSOverrides != nil {
			result[string(VnetDNSOverrides)] = flattenDNSOverrides(profile.VnetDNSOverrides)
		}
		if profile.KubeDNSOverrides != nil {
			result[string(KubeDNSOverrides)] = flattenDNSOverrides(profile.KubeDNSOverrides)
		}
		out = append(out, result)
	}
	return out
}

func flattenDNSOverrides(overrides map[string]*azure_np.DNSOverrideSettings) []interface{} {
	result := make([]interface{}, 0, len(overrides))

	for zone, dnsSettings := range overrides {
		if dnsSettings == nil {
			continue
		}
		item := make(map[string]interface{})
		item[string(DNSZone)] = zone

		if dnsSettings.QueryLogging != nil {
			item[string(QueryLogging)] = spotinst.StringValue(dnsSettings.QueryLogging)
		}
		if dnsSettings.Protocol != nil {
			item[string(Protocol)] = spotinst.StringValue(dnsSettings.Protocol)
		}
		if dnsSettings.ForwardDestination != nil {
			item[string(ForwardDestination)] = spotinst.StringValue(dnsSettings.ForwardDestination)
		}
		if dnsSettings.ForwardPolicy != nil {
			item[string(ForwardPolicy)] = spotinst.StringValue(dnsSettings.ForwardPolicy)
		}
		if dnsSettings.MaxConcurrent != nil {
			item[string(MaxConcurrent)] = spotinst.IntValue(dnsSettings.MaxConcurrent)
		}
		if dnsSettings.CacheDurationInSeconds != nil {
			item[string(CacheDurationInSeconds)] = spotinst.IntValue(dnsSettings.CacheDurationInSeconds)
		}
		if dnsSettings.ServeStaleDurationInSeconds != nil {
			item[string(ServeStaleDurationInSeconds)] = spotinst.IntValue(dnsSettings.ServeStaleDurationInSeconds)
		}
		if dnsSettings.ServeStale != nil {
			item[string(ServeStale)] = spotinst.StringValue(dnsSettings.ServeStale)
		}
		result = append(result, item)
	}
	return result
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

func expandLocalDnsProfile(data interface{}) (*azure_np.LocalDnsProfile, error) {
	list := data.([]interface{})
	if len(list) == 0 || list[0] == nil {
		return nil, nil
	}
	m := list[0].(map[string]interface{})
	profile := &azure_np.LocalDnsProfile{}

	if v, ok := m[string(Mode)].(string); ok && v != "" {
		profile.SetMode(spotinst.String(v))
	} else {
		profile.SetMode(nil)
	}

	if v, ok := m[string(VnetDNSOverrides)]; ok {
		overrides, err := expandDNSOverrides(v)
		if err != nil {
			return nil, err
		}
		if len(overrides) > 0 {
			profile.SetVnetDNSOverrides(overrides)
		} else {
			profile.SetVnetDNSOverrides(nil)
		}
	} else {
		profile.SetVnetDNSOverrides(nil)
	}

	if v, ok := m[string(KubeDNSOverrides)]; ok {
		overrides, err := expandDNSOverrides(v)
		if err != nil {
			return nil, err
		}
		if len(overrides) > 0 {
			profile.SetKubeDNSOverrides(overrides)
		} else {
			profile.SetKubeDNSOverrides(nil)
		}
	} else {
		profile.SetKubeDNSOverrides(nil)
	}

	return profile, nil
}

func expandDNSOverrides(data interface{}) (map[string]*azure_np.DNSOverrideSettings, error) {
	list := data.([]interface{})
	result := make(map[string]*azure_np.DNSOverrideSettings, len(list))

	for _, v := range list {
		m, ok := v.(map[string]interface{})
		if !ok || m == nil {
			continue
		}

		zone, ok := m[string(DNSZone)].(string)
		if !ok || zone == "" {
			continue
		}

		settings := &azure_np.DNSOverrideSettings{}

		if v, ok := m[string(QueryLogging)].(string); ok && v != "" {
			settings.SetQueryLogging(spotinst.String(v))
		}
		if v, ok := m[string(Protocol)].(string); ok && v != "" {
			settings.SetProtocol(spotinst.String(v))
		}
		if v, ok := m[string(ForwardDestination)].(string); ok && v != "" {
			settings.SetForwardDestination(spotinst.String(v))
		}
		if v, ok := m[string(ForwardPolicy)].(string); ok && v != "" {
			settings.SetForwardPolicy(spotinst.String(v))
		}
		if v, ok := m[string(MaxConcurrent)].(int); ok && v > 0 {
			settings.SetMaxConcurrent(spotinst.Int(v))
		}
		if v, ok := m[string(CacheDurationInSeconds)].(int); ok && v > 0 {
			settings.SetCacheDurationInSeconds(spotinst.Int(v))
		}
		if v, ok := m[string(ServeStaleDurationInSeconds)].(int); ok && v > 0 {
			settings.SetServeStaleDurationInSeconds(spotinst.Int(v))
		}
		if v, ok := m[string(ServeStale)].(string); ok && v != "" {
			settings.SetServeStale(spotinst.String(v))
		}

		result[zone] = settings
	}

	return result, nil
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

func dnsOverrideSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		string(DNSZone): {
			Type:     schema.TypeString,
			Required: true,
		},
		string(QueryLogging): {
			Type:     schema.TypeString,
			Optional: true,
		},
		string(Protocol): {
			Type:     schema.TypeString,
			Optional: true,
		},
		string(ForwardDestination): {
			Type:     schema.TypeString,
			Optional: true,
		},
		string(ForwardPolicy): {
			Type:     schema.TypeString,
			Optional: true,
		},
		string(MaxConcurrent): {
			Type:     schema.TypeInt,
			Optional: true,
		},
		string(CacheDurationInSeconds): {
			Type:     schema.TypeInt,
			Optional: true,
		},
		string(ServeStaleDurationInSeconds): {
			Type:     schema.TypeInt,
			Optional: true,
		},
		string(ServeStale): {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

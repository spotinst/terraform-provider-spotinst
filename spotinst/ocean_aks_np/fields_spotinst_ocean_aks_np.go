package ocean_aks_np

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanAKSNP,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if err := resourceData.Set(string(Name), spotinst.StringValue(cluster.Name)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOk(string(Name)); ok {
				cluster.SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOk(string(Name)); ok {
				cluster.SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[ControllerClusterID] = commons.NewGenericField(
		commons.OceanAKSNP,
		ControllerClusterID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if err := resourceData.Set(string(ControllerClusterID), spotinst.StringValue(cluster.ControllerClusterID)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(ControllerClusterID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOk(string(ControllerClusterID)); ok {
				cluster.SetControllerClusterId(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[AKSClusterName] = commons.NewGenericField(
		commons.OceanAKSNP,
		AKSClusterName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *string = nil
			if cluster.AKS != nil && cluster.AKS.ClusterName != nil {
				value = cluster.AKS.ClusterName
			}
			if err := resourceData.Set(string(AKSClusterName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AKSClusterName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.Get(string(AKSClusterName)).(string); ok && value != "" {
				cluster.AKS.SetClusterName(spotinst.String(value))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[AKSResourceGroupName] = commons.NewGenericField(
		commons.OceanAKSNP,
		AKSResourceGroupName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *string = nil
			if cluster.AKS != nil && cluster.AKS.ResourceGroupName != nil {
				value = cluster.AKS.ResourceGroupName
			}
			if err := resourceData.Set(string(AKSResourceGroupName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AKSResourceGroupName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.Get(string(AKSResourceGroupName)).(string); ok && value != "" {
				cluster.AKS.SetResourceGroupName(spotinst.String(value))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[AKSRegion] = commons.NewGenericField(
		commons.OceanAKSNP,
		AKSRegion,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *string = nil
			if cluster.AKS != nil && cluster.AKS.Region != nil {
				value = cluster.AKS.Region
			}
			if err := resourceData.Set(string(AKSRegion), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AKSRegion), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.Get(string(AKSRegion)).(string); ok && value != "" {
				cluster.AKS.SetRegion(spotinst.String(value))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[AKSInfrastructureResourceGroupName] = commons.NewGenericField(
		commons.OceanAKSNP,
		AKSInfrastructureResourceGroupName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *string = nil
			if cluster.AKS != nil && cluster.AKS.InfrastructureResourceGroupName != nil {
				value = cluster.AKS.InfrastructureResourceGroupName
			}
			if err := resourceData.Set(string(AKSInfrastructureResourceGroupName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AKSInfrastructureResourceGroupName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.Get(string(AKSInfrastructureResourceGroupName)).(string); ok && value != "" {
				cluster.AKS.SetInfrastructureResourceGroupName(spotinst.String(value))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[AvailabilityZones] = commons.NewGenericField(
		commons.OceanAKSNP,
		AvailabilityZones,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value []string = nil
			if cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.AvailabilityZones != nil {
				value = cluster.VirtualNodeGroupTemplate.AvailabilityZones
			}
			if err := resourceData.Set(string(AvailabilityZones), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AvailabilityZones), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				if zones, err := expandZones(value); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.SetAvailabilityZones(zones)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if value, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				if zones, err := expandZones(value); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.SetAvailabilityZones(zones)
				}
			} else {
				cluster.VirtualNodeGroupTemplate.SetAvailabilityZones(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[UpdatePolicy] = commons.NewGenericField(
		commons.OceanAKSNP,
		UpdatePolicy,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ShouldRoll): {
						Type:     schema.TypeBool,
						Required: true,
					},
					string(ConditionedRoll): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(RollConfig): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(BatchSizePercentage): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  -1,
								},
								string(VngIDs): {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								string(BatchMinHealthyPercentage): {
									Type:     schema.TypeInt,
									Optional: true,
								},
								string(RespectPDB): {
									Type:     schema.TypeBool,
									Optional: true,
								},
								string(Comment): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(NodePoolNames): {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								string(RespectRestrictScaleDown): {
									Type:     schema.TypeBool,
									Optional: true,
								},
								string(NodeNames): {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
							},
						},
					},
				},
			},
		},
		nil, nil, nil, nil,
	)

	fieldsMap[VNG_Template_Scheduling] = commons.NewGenericField(
		commons.OceanAKSNP,
		VNG_Template_Scheduling,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ShutdownHours): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(ShutdownHoursIsEnabled): {
									Type:     schema.TypeBool,
									Optional: true,
								},

								string(TimeWindows): {
									Type:     schema.TypeList,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
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
			var result []interface{} = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.Scheduling != nil {
				result = flattenScheduling(cluster.VirtualNodeGroupTemplate.Scheduling)
			}

			if len(result) > 0 {
				if err := resourceData.Set(string(VNG_Template_Scheduling), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(VNG_Template_Scheduling), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *azure_np.Scheduling = nil
			if v, ok := resourceData.GetOkExists(string(VNG_Template_Scheduling)); ok {
				if scheduling, err := expandScheduling(v); err != nil {
					return err
				} else {
					value = scheduling
				}
			}
			cluster.VirtualNodeGroupTemplate.SetScheduling(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *azure_np.Scheduling = nil
			if v, ok := resourceData.GetOk(string(VNG_Template_Scheduling)); ok {
				if scheduling, err := expandScheduling(v); err != nil {
					return err
				} else {
					value = scheduling
				}
			}
			cluster.VirtualNodeGroupTemplate.SetScheduling(value)
			return nil
		},
		nil,
	)
}

func expandZones(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if zones, ok := v.(string); ok && zones != "" {
			result = append(result, zones)
		}
	}
	return result, nil
}

func flattenScheduling(scheduling *azure_np.Scheduling) []interface{} {
	var out []interface{}

	if scheduling != nil {
		result := make(map[string]interface{})
		if scheduling.ShutdownHours != nil {
			result[string(ShutdownHours)] = flattenShutdownHours(scheduling.ShutdownHours)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

func flattenShutdownHours(shutdownHours *azure_np.ShutdownHours) []interface{} {
	result := make(map[string]interface{})
	result[string(ShutdownHoursIsEnabled)] = spotinst.BoolValue(shutdownHours.IsEnabled)
	if len(shutdownHours.TimeWindows) > 0 {
		result[string(TimeWindows)] = shutdownHours.TimeWindows
	}
	return []interface{}{result}
}

func expandScheduling(data interface{}) (*azure_np.Scheduling, error) {
	scheduling := &azure_np.Scheduling{}
	if list := data.([]interface{}); len(list) > 0 {
		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(ShutdownHours)]; ok {
				shutdownHours, err := expandShutdownHours(v)
				if err != nil {
					return nil, err
				}
				if shutdownHours != nil {
					if scheduling.ShutdownHours == nil {
						scheduling.SetShutdownHours(&azure_np.ShutdownHours{})
					}
					scheduling.SetShutdownHours(shutdownHours)
				}
			}
		}
		return scheduling, nil
	}
	return nil, nil
}

func expandShutdownHours(data interface{}) (*azure_np.ShutdownHours, error) {
	shutDownHours := &azure_np.ShutdownHours{}
	if list := data.([]interface{}); len(list) > 0 && list[0] != nil {
		m := list[0].(map[string]interface{})

		var isEnabled = spotinst.Bool(false)
		if v, ok := m[string(ShutdownHoursIsEnabled)].(bool); ok {
			isEnabled = spotinst.Bool(v)
		}
		shutDownHours.SetIsEnabled(isEnabled)

		var timeWindows []string = nil
		if v, ok := m[string(TimeWindows)].([]interface{}); ok && len(v) > 0 {
			timeWindowList := make([]string, 0, len(v))
			for _, timeWindow := range v {
				if v, ok := timeWindow.(string); ok && len(v) > 0 {
					timeWindowList = append(timeWindowList, v)
				}
			}
			timeWindows = timeWindowList
		}
		shutDownHours.SetTimeWindows(timeWindows)

		return shutDownHours, nil
	}
	return nil, nil
}

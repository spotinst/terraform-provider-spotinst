package ocean_aks

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[ACDIdentifier] = commons.NewGenericField(
		commons.OceanAKS,
		ACDIdentifier,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern), string(ACDIdentifier))
		},
		nil,
	)

	fieldsMap[ControllerClusterID] = commons.NewGenericField(
		commons.OceanAKS,
		ControllerClusterID,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if err := resourceData.Set(string(ControllerClusterID), spotinst.StringValue(cluster.ControllerClusterID)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(ControllerClusterID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
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

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanAKS,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if err := resourceData.Set(string(Name), spotinst.StringValue(cluster.Name)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(Name)); ok {
				cluster.SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(Name)); ok {
				cluster.SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[AKSName] = commons.NewGenericField(
		commons.OceanAKS,
		AKSName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if cluster.AKS != nil && cluster.AKS.Name != nil {
				value = cluster.AKS.Name
			}
			if err := resourceData.Set(string(AKSName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AKSName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if value, ok := resourceData.Get(string(AKSName)).(string); ok && value != "" {
				cluster.AKS.SetName(spotinst.String(resourceData.Get(string(AKSName)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[AKSResourceGroupName] = commons.NewGenericField(
		commons.OceanAKS,
		AKSResourceGroupName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
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
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[Zones] = commons.NewGenericField(
		commons.OceanAKS,
		Zones,
		&schema.Schema{
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString},
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value []string = nil
			if cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.Zones != nil {
				value = cluster.VirtualNodeGroupTemplate.Zones
			}
			if err := resourceData.Set(string(Zones), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Zones), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if value, ok := resourceData.GetOk(string(Zones)); ok && value != nil {
				if zones, err := expandZones(value); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.SetZones(zones)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if value, ok := resourceData.GetOk(string(Zones)); ok && value != nil {
				if zones, err := expandZones(value); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.SetZones(zones)
				}
			}
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

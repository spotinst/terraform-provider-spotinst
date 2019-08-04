package ocean_ecs

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanECS,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value *string = nil
			if cluster.Name != nil {
				value = cluster.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			cluster.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			cluster.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Region] = commons.NewGenericField(
		commons.OceanECS,
		Region,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value *string = nil
			if cluster.Region != nil {
				value = cluster.Region
			}
			if err := resourceData.Set(string(Region), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Region), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			cluster.SetRegion(spotinst.String(resourceData.Get(string(Region)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			cluster.SetRegion(spotinst.String(resourceData.Get(string(Region)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[ClusterName] = commons.NewGenericField(
		commons.OceanECS,
		ClusterName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value *string = nil
			if cluster.ClusterName != nil {
				value = cluster.ClusterName
			}
			if err := resourceData.Set(string(ClusterName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ClusterName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			cluster.SetClusterName(spotinst.String(resourceData.Get(string(ClusterName)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			cluster.SetClusterName(spotinst.String(resourceData.Get(string(ClusterName)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[MaxSize] = commons.NewGenericField(
		commons.OceanECS,
		MaxSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value *int = nil
			if cluster.Capacity != nil && cluster.Capacity.Maximum != nil {
				value = cluster.Capacity.Maximum
			}
			if err := resourceData.Set(string(MaxSize), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MaxSize), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.GetOk(string(MaxSize)); ok {
				cluster.Capacity.SetMaximum(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.GetOk(string(MaxSize)); ok {
				cluster.Capacity.SetMaximum(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[MinSize] = commons.NewGenericField(
		commons.OceanECS,
		MinSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value *int = nil
			if cluster.Capacity != nil && cluster.Capacity.Minimum != nil {
				value = cluster.Capacity.Minimum
			}
			if err := resourceData.Set(string(MinSize), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MinSize), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.GetOk(string(MinSize)); ok {
				cluster.Capacity.SetMinimum(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.GetOk(string(MinSize)); ok {
				cluster.Capacity.SetMinimum(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[DesiredCapacity] = commons.NewGenericField(
		commons.OceanECS,
		DesiredCapacity,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value *int = nil
			if cluster.Capacity != nil && cluster.Capacity.Target != nil {
				value = cluster.Capacity.Target
			}
			if err := resourceData.Set(string(DesiredCapacity), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DesiredCapacity), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.GetOk(string(DesiredCapacity)); ok {
				cluster.Capacity.SetTarget(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.GetOk(string(DesiredCapacity)); ok {
				cluster.Capacity.SetTarget(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[SubnetIds] = commons.NewGenericField(
		commons.OceanECS,
		SubnetIds,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value []string = nil
			if cluster.Compute != nil && cluster.Compute.SubnetIDs != nil {
				value = cluster.Compute.SubnetIDs
			}
			if err := resourceData.Set(string(SubnetIds), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SubnetIds), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if value, ok := resourceData.GetOk(string(SubnetIds)); ok && value != nil {
				if subnetIds, err := expandSubnetIDs(value); err != nil {
					return err
				} else {
					cluster.Compute.SetSubnetIDs(subnetIds)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if value, ok := resourceData.GetOk(string(SubnetIds)); ok && value != nil {
				if subnetIds, err := expandSubnetIDs(value); err != nil {
					return err
				} else {
					cluster.Compute.SetSubnetIDs(subnetIds)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[UpdatePolicy] = commons.NewGenericField(
		commons.OceanECS,
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

					string(RollConfig): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(BatchSizePercentage): {
									Type:     schema.TypeInt,
									Required: true,
								},
							},
						},
					},
				},
			},
		},
		nil, nil, nil, nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandSubnetIDs(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if subnetID, ok := v.(string); ok && subnetID != "" {
			result = append(result, subnetID)
		}
	}
	return result, nil
}

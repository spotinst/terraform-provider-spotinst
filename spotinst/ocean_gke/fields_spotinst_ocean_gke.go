package ocean_gke

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanGKE,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
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
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			cluster.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			cluster.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[MasterLocation] = commons.NewGenericField(
		commons.OceanGKE,
		MasterLocation,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(MasterLocation)); ok && v != nil {
				if cluster.GKE == nil {
					cluster.SetGKE(&gcp.GKE{})
				}
				cluster.GKE.SetMasterLocation(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[ClusterName] = commons.NewGenericField(
		commons.OceanGKE,
		ClusterName,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if cluster.GKE != nil && cluster.GKE.ClusterName != nil {
				value = cluster.GKE.ClusterName
			}
			if err := resourceData.Set(string(ClusterName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ClusterName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(ClusterName)); ok && v != nil {
				if cluster.GKE == nil {
					cluster.SetGKE(&gcp.GKE{})
				}
				cluster.GKE.SetClusterName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if v, ok := resourceData.GetOk(string(ClusterName)); ok && v != nil {
				if cluster.GKE == nil {
					cluster.SetGKE(&gcp.GKE{})
				}
				value = spotinst.String(v.(string))
			}
			cluster.GKE.SetClusterName(value)
			return nil
		},
		nil,
	)

	fieldsMap[ControllerClusterID] = commons.NewGenericField(
		commons.OceanGKE,
		ControllerClusterID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if cluster.ControllerClusterID != nil {
				value = cluster.ControllerClusterID
			}
			if err := resourceData.Set(string(ControllerClusterID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ControllerClusterID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			cluster.SetControllerClusterId(spotinst.String(resourceData.Get(string(ControllerClusterID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			cluster.SetControllerClusterId(spotinst.String(resourceData.Get(string(ControllerClusterID)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[MaxSize] = commons.NewGenericField(
		commons.OceanGKE,
		MaxSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
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
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(MaxSize)); ok {
				cluster.Capacity.SetMaximum(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(MaxSize)); ok {
				cluster.Capacity.SetMaximum(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[MinSize] = commons.NewGenericField(
		commons.OceanGKE,
		MinSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
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
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(MinSize)); ok {
				cluster.Capacity.SetMinimum(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(MinSize)); ok {
				cluster.Capacity.SetMinimum(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[DesiredCapacity] = commons.NewGenericField(
		commons.OceanGKE,
		DesiredCapacity,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
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
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(DesiredCapacity)); ok {
				cluster.Capacity.SetTarget(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(DesiredCapacity)); ok {
				cluster.Capacity.SetTarget(spotinst.Int(v.(int)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[SubnetName] = commons.NewGenericField(
		commons.OceanGKE,
		SubnetName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if cluster.Compute != nil && cluster.Compute.SubnetName != nil {
				value = cluster.Compute.SubnetName
			}
			if err := resourceData.Set(string(SubnetName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SubnetName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			cluster.Compute.SetSubnetName(spotinst.String(resourceData.Get(string(SubnetName)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			cluster.Compute.SetSubnetName(spotinst.String(resourceData.Get(string(SubnetName)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[AvailabilityZones] = commons.NewGenericField(
		commons.OceanGKELaunchConfiguration,
		AvailabilityZones,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
			//DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []string = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.AvailabilityZones != nil {
				values := cluster.Compute.AvailabilityZones
				for _, zones := range values {
					result = append(result, zones)
				}
			}
			if err := resourceData.Set(string(AvailabilityZones), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AvailabilityZones), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []string
			if v, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				azList := v.([]interface{})
				result = make([]string, len(azList))
				for i, j := range azList {
					result[i] = j.(string)
				}
				cluster.Compute.SetAvailabilityZones(result)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []string
			if v, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				azList := v.([]interface{})
				result = make([]string, len(azList))
				for i, j := range azList {
					result[i] = j.(string)
				}
				cluster.Compute.SetAvailabilityZones(result)
			}
			return nil
		},
		nil,
	)

}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

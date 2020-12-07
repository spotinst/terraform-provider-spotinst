package ocean_ecs_optimize_images

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[PerformAt] = commons.NewGenericField(
		commons.OceanECSOptimizeImages,
		PerformAt,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value = ""
			if cluster.Compute != nil && cluster.Compute.OptimizeImages != nil &&
				cluster.Compute.OptimizeImages.PerformAt != nil {
				oi := cluster.Compute.OptimizeImages
				value = spotinst.StringValue(oi.PerformAt)
			}
			if err := resourceData.Set(string(PerformAt), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PerformAt), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.Get(string(PerformAt)).(string); ok && v != "" {
				cluster.Compute.OptimizeImages.SetPerformAt(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var performAt *string = nil
			if v, ok := resourceData.Get(string(PerformAt)).(string); ok && v != "" {
				performAt = spotinst.String(v)
			}
			cluster.Compute.OptimizeImages.SetPerformAt(performAt)
			return nil
		},
		nil,
	)

	fieldsMap[TimeWindows] = commons.NewGenericField(
		commons.OceanECSOptimizeImages,
		TimeWindows,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var result []string
			if cluster.Compute != nil && cluster.Compute.OptimizeImages != nil &&
				cluster.Compute.OptimizeImages.TimeWindows != nil {
				result = append(result, cluster.Compute.OptimizeImages.TimeWindows...)
			}
			if err := resourceData.Set(string(TimeWindows), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(TimeWindows), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.GetOk(string(TimeWindows)); ok {
				tw := v.([]interface{})
				timeWindows := make([]string, len(tw))
				for i, j := range tw {
					timeWindows[i] = j.(string)
				}
				cluster.Compute.OptimizeImages.SetTimeWindows(timeWindows)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var timeWindows []string = nil
			if v, ok := resourceData.GetOk(string(TimeWindows)); ok {
				tw := v.([]interface{})
				timeWindows = make([]string, len(tw))
				for i, v := range tw {
					timeWindows[i] = v.(string)
				}
			}
			cluster.Compute.OptimizeImages.SetTimeWindows(timeWindows)
			return nil
		},
		nil,
	)

	fieldsMap[ShouldOptimizeECSAMI] = commons.NewGenericField(
		commons.OceanECSOptimizeImages,
		ShouldOptimizeECSAMI,
		&schema.Schema{
			Type:     schema.TypeBool,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value *bool = nil
			if cluster.Compute != nil && cluster.Compute.OptimizeImages != nil &&
				cluster.Compute.OptimizeImages.ShouldOptimizeECSAMI != nil {
				value = cluster.Compute.OptimizeImages.ShouldOptimizeECSAMI
			}
			if err := resourceData.Set(string(ShouldOptimizeECSAMI), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShouldOptimizeECSAMI), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.GetOkExists(string(ShouldOptimizeECSAMI)); ok {
				cluster.Compute.OptimizeImages.SetShouldOptimizeECSAMI(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var shouldOptimizeECSAMI *bool = nil
			if v, ok := resourceData.GetOkExists(string(ShouldOptimizeECSAMI)); ok {
				soea := v.(bool)
				shouldOptimizeECSAMI = spotinst.Bool(soea)
			}
			cluster.Compute.OptimizeImages.SetShouldOptimizeECSAMI(shouldOptimizeECSAMI)
			return nil
		},
		nil,
	)
}

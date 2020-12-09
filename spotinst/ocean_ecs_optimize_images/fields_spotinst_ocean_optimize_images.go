package ocean_ecs_optimize_images

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[OptimizeImages] = commons.NewGenericField(
		commons.OceanECSOptimizeImages,
		OptimizeImages,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(PerformAt): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(TimeWindows): {
						Type:     schema.TypeList,
						MinItems: 1,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					string(ShouldOptimizeECSAMI): {
						Type:     schema.TypeBool,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var result []interface{} = nil
			if cluster.Compute != nil && cluster.Compute.OptimizeImages != nil {
				compute := cluster.Compute
				result = flattenOptimizeImages(compute.OptimizeImages)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(OptimizeImages), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OptimizeImages), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.GetOk(string(OptimizeImages)); ok {
				if v, err := expandOptimizeImages(v); err != nil {
					return err
				} else {
					cluster.Compute.SetOptimizeImages(v)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value *aws.ECSOptimizeImages = nil
			if v, ok := resourceData.GetOk(string(OptimizeImages)); ok {
				if otimizeImages, err := expandOptimizeImages(v); err != nil {
					return err
				} else {
					value = otimizeImages
				}
			}
			cluster.Compute.SetOptimizeImages(value)
			return nil
		},
		nil,
	)
}

func flattenOptimizeImages(oi *aws.ECSOptimizeImages) []interface{} {
	optimizeImages := make(map[string]interface{})

	optimizeImages[string(PerformAt)] = spotinst.StringValue(oi.PerformAt)
	optimizeImages[string(TimeWindows)] = spotinst.StringSlice(oi.TimeWindows)
	optimizeImages[string(ShouldOptimizeECSAMI)] = spotinst.BoolValue(oi.ShouldOptimizeECSAMI)

	return []interface{}{optimizeImages}
}

func expandOptimizeImages(data interface{}) (*aws.ECSOptimizeImages, error) {
	if list := data.([]interface{}); len(list) > 0 {
		oi := &aws.ECSOptimizeImages{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(PerformAt)].(string); ok && v != "" {
				oi.SetPerformAt(spotinst.String(v))
			}

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
			oi.SetTimeWindows(timeWindows)

			if v, ok := m[string(ShouldOptimizeECSAMI)].(bool); ok {
				oi.SetShouldOptimizeECSAMI(spotinst.Bool(v))
			}
		}
		return oi, nil
	}

	return nil, nil
}

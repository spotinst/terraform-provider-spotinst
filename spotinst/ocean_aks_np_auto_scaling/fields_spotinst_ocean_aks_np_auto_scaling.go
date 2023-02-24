package ocean_aks_np_auto_scaling

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[AutoScaler] = commons.NewGenericField(
		commons.OceanAKSNPAutoScaler,
		AutoScaler,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(AutoscaleIsEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
						Computed: true,
					},
					string(ResourceLimits): {
						Type:     schema.TypeList,
						Optional: true,
						Computed: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(MaxVCPU): {
									Type:     schema.TypeInt,
									Optional: true,
								},
								string(MaxMemoryGib): {
									Type:     schema.TypeInt,
									Optional: true,
									Computed: true,
								},
							},
						},
					},
					string(Down): {
						Type:     schema.TypeList,
						Optional: true,
						Computed: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(MaxScaleDownPercentage): {
									Type:     schema.TypeFloat,
									Optional: true,
									Computed: true,
								},
							},
						},
					},
					string(Headroom): {
						Type:     schema.TypeList,
						Optional: true,
						Computed: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Automatic): {
									Type:     schema.TypeList,
									Optional: true,
									Computed: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(IsEnabled): {
												Type:     schema.TypeBool,
												Optional: true,
												Computed: true,
											},
											string(Percentage): {
												Type:     schema.TypeInt,
												Optional: true,
												Computed: true,
											},
										},
									},
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

			if cluster != nil && cluster.AutoScaler != nil {
				result = flattenAutoScaler(cluster.AutoScaler)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(AutoScaler), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AutoScaler), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *azure_np.AutoScaler = nil

			if v, ok := resourceData.GetOk(string(AutoScaler)); ok {
				if autoScaler, err := expandAutoScaler(v); err != nil {
					return err
				} else {
					value = autoScaler
				}
			}
			cluster.SetAutoScaler(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *azure_np.AutoScaler = nil

			if v, ok := resourceData.GetOk(string(AutoScaler)); ok {
				if autoScaler, err := expandAutoScaler(v); err != nil {
					return err
				} else {
					value = autoScaler
				}
			}
			cluster.SetAutoScaler(value)
			return nil
		},
		nil,
	)
}

func expandAutoScaler(data interface{}) (*azure_np.AutoScaler, error) {
	if list := data.([]interface{}); len(list) > 0 {
		autoScaler := &azure_np.AutoScaler{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(AutoscaleIsEnabled)].(bool); ok {
				autoScaler.SetIsEnabled(spotinst.Bool(v))
			}

			if v, ok := m[string(ResourceLimits)]; ok && v != nil {

				resLimits, err := expandResourceLimits(v)
				if err != nil {
					return nil, err
				}
				if resLimits != nil {
					autoScaler.SetResourceLimits(resLimits)
				} else {
					log.Printf("resLimits == nil")
					autoScaler.ResourceLimits = nil
				}
			}

			if v, ok := m[string(Down)]; ok {
				down, err := expandDown(v)
				if err != nil {
					return nil, err
				}
				if down != nil {
					autoScaler.SetDown(down)
				} else {
					autoScaler.Down = nil
				}
			}

			if v, ok := m[string(Headroom)]; ok {
				headroom, err := expandHeadroom(v)
				if err != nil {
					return nil, err
				}
				if headroom != nil {
					autoScaler.SetHeadroom(headroom)
				} else {
					autoScaler.Headroom = nil
				}
			}
		}

		return autoScaler, nil
	}

	return nil, nil
}

func expandResourceLimits(data interface{}) (*azure_np.ResourceLimits, error) {
	resLimits := &azure_np.ResourceLimits{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return nil, nil
	}

	m := list[0].(map[string]interface{})

	if v, ok := m[string(MaxMemoryGib)].(int); ok && v >= 0 {
		resLimits.SetMaxMemoryGib(spotinst.Int(v))
	}

	if v, ok := m[string(MaxVCPU)].(int); ok && v >= 0 {
		resLimits.SetMaxVcpu(spotinst.Int(v))
	}

	return resLimits, nil
}

func expandDown(data interface{}) (*azure_np.Down, error) {
	down := &azure_np.Down{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return nil, nil
	}

	m := list[0].(map[string]interface{})

	if v, ok := m[string(MaxScaleDownPercentage)].(float64); ok && v >= 0 {
		down.SetMaxScaleDownPercentage(spotinst.Float64(v))
	}

	return down, nil
}

func expandHeadroom(data interface{}) (*azure_np.Headroom, error) {
	headroom := &azure_np.Headroom{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return nil, nil
	}

	m := list[0].(map[string]interface{})

	if v, ok := m[string(Automatic)]; ok {
		automatic, err := expandAutomatic(v)
		if err != nil {
			return nil, err
		}
		if automatic != nil {
			headroom.SetAutomatic(automatic)
		} else {
			headroom.Automatic = nil
		}
	}

	return headroom, nil
}

func expandAutomatic(data interface{}) (*azure_np.Automatic, error) {
	automatic := &azure_np.Automatic{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return nil, nil
	}

	m := list[0].(map[string]interface{})

	if v, ok := m[string(Percentage)].(int); ok && v >= 0 {
		automatic.SetPercentage(spotinst.Int(v))
	}

	if v, ok := m[string(IsEnabled)].(bool); ok {
		automatic.SetIsEnabled(spotinst.Bool(v))
	}

	return automatic, nil
}

func flattenAutoScaler(autoScaler *azure_np.AutoScaler) []interface{} {
	result := make(map[string]interface{})
	result[string(AutoscaleIsEnabled)] = spotinst.BoolValue(autoScaler.IsEnabled)

	if autoScaler.Headroom != nil {
		result[string(Headroom)] = flattenHeadroom(autoScaler.Headroom)
	}

	if autoScaler.Down != nil {
		result[string(Down)] = flattenDown(autoScaler.Down)
	}

	if autoScaler.ResourceLimits != nil {
		result[string(ResourceLimits)] = flattenResourceLimits(autoScaler.ResourceLimits)
	}

	return []interface{}{result}
}

func flattenHeadroom(headroom *azure_np.Headroom) []interface{} {
	result := make(map[string]interface{})

	if headroom.Automatic != nil {
		result[string(Automatic)] = flattenAutomatic(headroom.Automatic)
	}

	return []interface{}{result}
}

func flattenDown(autoScaleDown *azure_np.Down) []interface{} {
	down := make(map[string]interface{})
	down[string(MaxScaleDownPercentage)] = spotinst.Float64Value(autoScaleDown.MaxScaleDownPercentage)

	return []interface{}{down}
}

func flattenResourceLimits(autoScaleResourceLimits *azure_np.ResourceLimits) []interface{} {
	resourceLimits := make(map[string]interface{})
	resourceLimits[string(MaxVCPU)] = spotinst.IntValue(autoScaleResourceLimits.MaxVCPU)
	resourceLimits[string(MaxMemoryGib)] = spotinst.IntValue(autoScaleResourceLimits.MaxMemoryGib)

	return []interface{}{resourceLimits}
}

func flattenAutomatic(autoScaleAutomatic *azure_np.Automatic) []interface{} {
	automatic := make(map[string]interface{})
	automatic[string(IsEnabled)] = spotinst.BoolValue(autoScaleAutomatic.IsEnabled)
	automatic[string(Percentage)] = spotinst.IntValue(autoScaleAutomatic.Percentage)

	return []interface{}{automatic}
}

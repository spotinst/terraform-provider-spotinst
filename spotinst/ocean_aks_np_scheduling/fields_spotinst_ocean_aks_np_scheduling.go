package ocean_aks_np_scheduling

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Scheduling] = commons.NewGenericField(
		commons.OceanAKSNPScheduling,
		Scheduling,
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
									Required: true,
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
			if cluster != nil && cluster.Scheduling != nil {
				result = flattenScheduling(cluster.Scheduling)
			}

			if len(result) > 0 {
				if err := resourceData.Set(string(Scheduling), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Scheduling), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *azure_np.Scheduling = nil
			if v, ok := resourceData.GetOkExists(string(Scheduling)); ok {
				if scheduling, err := expandScheduling(v); err != nil {
					return err
				} else {
					value = scheduling
				}
			}
			cluster.SetScheduling(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *azure_np.Scheduling = nil
			if v, ok := resourceData.GetOk(string(Scheduling)); ok {
				if scheduling, err := expandScheduling(v); err != nil {
					return err
				} else {
					value = scheduling
				}
			}
			cluster.SetScheduling(value)
			return nil
		},
		nil,
	)

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

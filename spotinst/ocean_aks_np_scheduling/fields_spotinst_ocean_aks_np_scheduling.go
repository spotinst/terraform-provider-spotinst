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
					string(Tasks): {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(TasksIsEnabled): {
									Type:     schema.TypeBool,
									Required: true,
								},

								string(TaskType): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(CronExpression): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(Parameters): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(ParametersClusterRoll): {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														string(BatchMinHealthyPercentage): {
															Type:     schema.TypeInt,
															Optional: true,
															Default:  -1,
														},
														string(BatchSizePercentage): {
															Type:     schema.TypeInt,
															Optional: true,
															Default:  -1,
														},
														string(Comment): {
															Type:     schema.TypeString,
															Optional: true,
														},
														string(RespectPdb): {
															Type:     schema.TypeBool,
															Optional: true,
														},
														string(RespectRestrictScaleDown): {
															Type:     schema.TypeBool,
															Optional: true,
														},
														string(VngIDs): {
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
		if scheduling.Tasks != nil {
			result[string(Tasks)] = flattenTasks(scheduling.Tasks)
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
			if v, ok := m[string(Tasks)]; ok {
				tasks, err := expandtasks(v)
				if err != nil {
					return nil, err
				}
				if tasks != nil {
					scheduling.SetTasks(tasks)
				} else {
					scheduling.SetTasks(nil)
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

func expandtasks(data interface{}) ([]*azure_np.Tasks, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		tasks := make([]*azure_np.Tasks, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			task := &azure_np.Tasks{}

			if v, ok := m[string(TasksIsEnabled)].(bool); ok {
				task.SetIsEnabled(spotinst.Bool(v))
			}

			if v, ok := m[string(TaskType)].(string); ok && v != "" {
				task.SetTaskType(spotinst.String(v))
			}

			if v, ok := m[string(CronExpression)].(string); ok && v != "" {
				task.SetCronExpression(spotinst.String(v))
			}

			if v, ok := m[string(Parameters)]; ok {
				parameters, err := expandParameters(v)
				if err != nil {
					return nil, err
				}
				if parameters != nil {
					task.SetParameters(parameters)
				} else {
					task.SetParameters(nil)
				}
			}

			tasks = append(tasks, task)
		}

		return tasks, nil
	}
	return nil, nil
}
func expandParameters(data interface{}) (*azure_np.Parameters, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		parameter := &azure_np.Parameters{}
		m := list[0].(map[string]interface{})

		if v, ok := m[string(ParametersClusterRoll)]; ok {
			expandClusterRoll, err := expandParameterClusterRoll(v)
			if err != nil {
				return nil, err
			}
			if expandClusterRoll != nil {
				parameter.SetClusterRoll(expandClusterRoll)
			} else {
				parameter.ClusterRoll = nil
			}
		}

		return parameter, nil
	}

	return nil, nil
}
func expandParameterClusterRoll(data interface{}) (*azure_np.ParameterClusterRoll, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		parameterClusterRoll := &azure_np.ParameterClusterRoll{}
		m := list[0].(map[string]interface{})

		if v, ok := m[string(BatchMinHealthyPercentage)].(int); ok {
			if v == -1 {
				parameterClusterRoll.SetBatchMinHealthyPercentage(nil)
			} else {
				parameterClusterRoll.SetBatchMinHealthyPercentage(spotinst.Int(v))
			}
		}

		if v, ok := m[string(BatchSizePercentage)].(int); ok {
			if v == -1 {
				parameterClusterRoll.SetBatchSizePercentage(nil)
			} else {
				parameterClusterRoll.SetBatchSizePercentage(spotinst.Int(v))
			}
		}

		if v, ok := m[string(Comment)].(string); ok && v != "" {
			parameterClusterRoll.SetComment(spotinst.String(v))
		} else {
			parameterClusterRoll.SetComment(nil)
		}

		var isRespectPdb = spotinst.Bool(false)
		if v, ok := m[string(RespectPdb)].(bool); ok {
			isRespectPdb = spotinst.Bool(v)
		}
		parameterClusterRoll.SetRespectPdb(isRespectPdb)

		var isRespectRestrictScaleDown = spotinst.Bool(false)
		if v, ok := m[string(RespectRestrictScaleDown)].(bool); ok {
			isRespectRestrictScaleDown = spotinst.Bool(v)
		}
		parameterClusterRoll.SetRespectRestrictScaleDown(isRespectRestrictScaleDown)

		if v, ok := m[string(VngIDs)]; ok {
			parameterClusterRoll.VngIds = expandListVNG(v)
		}

		return parameterClusterRoll, nil
	}

	return nil, nil
}
func expandListVNG(data interface{}) []string {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if ls, ok := v.(string); ok && ls != "" {
			result = append(result, ls)
		}
	}

	return result
}
func flattenTasks(tasks []*azure_np.Tasks) []interface{} {
	result := make([]interface{}, 0, len(tasks))

	for _, task := range tasks {
		m := make(map[string]interface{})
		m[string(TasksIsEnabled)] = spotinst.BoolValue(task.IsEnabled)
		m[string(TaskType)] = spotinst.StringValue(task.TaskType)
		m[string(CronExpression)] = spotinst.StringValue(task.CronExpression)
		if task.Parameters != nil {
			m[string(Parameters)] = flattenParameters(task.Parameters)
		}
		result = append(result, m)
	}

	return result
}
func flattenParameters(parameters *azure_np.Parameters) []interface{} {
	result := make(map[string]interface{})

	if parameters.ClusterRoll != nil {
		result[string(ParametersClusterRoll)] = flattenParameterClusterRoll(parameters.ClusterRoll)
	}

	return []interface{}{result}
}
func flattenParameterClusterRoll(clusterRoll *azure_np.ParameterClusterRoll) []interface{} {
	result := make(map[string]interface{})
	value := spotinst.Int(-1)
	result[string(BatchMinHealthyPercentage)] = value
	result[string(BatchSizePercentage)] = value

	if clusterRoll.BatchMinHealthyPercentage != nil {
		result[string(BatchMinHealthyPercentage)] = spotinst.IntValue(clusterRoll.BatchMinHealthyPercentage)
	}
	if clusterRoll.BatchSizePercentage != nil {
		result[string(BatchSizePercentage)] = spotinst.IntValue(clusterRoll.BatchSizePercentage)
	}
	result[string(Comment)] = spotinst.StringValue(clusterRoll.Comment)
	result[string(RespectPdb)] = spotinst.BoolValue(clusterRoll.RespectPdb)
	result[string(RespectRestrictScaleDown)] = spotinst.BoolValue(clusterRoll.RespectRestrictScaleDown)
	result[string(VngIDs)] = spotinst.StringSlice(clusterRoll.VngIds)

	return []interface{}{result}
}

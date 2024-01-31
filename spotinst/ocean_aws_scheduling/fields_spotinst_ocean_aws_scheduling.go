package ocean_aws_scheduling

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ScheduledTask] = commons.NewGenericField(
		commons.OceanAWSScheduling,
		ScheduledTask,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
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
											string(AmiAutoUpdate): {
												Type:     schema.TypeList,
												Optional: true,
												MaxItems: 1,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														string(ApplyRoll): {
															Type:     schema.TypeBool,
															Optional: true,
														},
														string(AmiAutoUpdateClusterRoll): {
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
																},
															},
														},
														string(MinorVersion): {
															Type:     schema.TypeBool,
															Optional: true,
														},
														string(Patch): {
															Type:     schema.TypeBool,
															Optional: true,
														},
													},
												},
											},
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
													},
												},
											},
										},
									},
								},
							},
						},
					},
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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster != nil && cluster.Scheduling != nil {
				result = flattenScheduledTasks(cluster.Scheduling)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(ScheduledTask), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ScheduledTask), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var scheduling *aws.Scheduling = nil
			if v, ok := resourceData.GetOk(string(ScheduledTask)); ok {
				if scheduledTask, err := expandScheduledTasks(v); err != nil {
					return err
				} else {
					scheduling = scheduledTask
				}
			}
			cluster.SetScheduling(scheduling)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var scheduling *aws.Scheduling = nil
			if v, ok := resourceData.GetOk(string(ScheduledTask)); ok {
				if scheduledTask, err := expandScheduledTasks(v); err != nil {
					return err
				} else {
					scheduling = scheduledTask
				}
			}
			cluster.SetScheduling(scheduling)
			return nil
		},

		nil,
	)

}

func flattenScheduledTasks(scheduling *aws.Scheduling) []interface{} {
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

func flattenShutdownHours(shutdownHours *aws.ShutdownHours) []interface{} {
	result := make(map[string]interface{})
	result[string(ShutdownHoursIsEnabled)] = spotinst.BoolValue(shutdownHours.IsEnabled)

	if len(shutdownHours.TimeWindows) > 0 {
		result[string(TimeWindows)] = shutdownHours.TimeWindows
	}

	return []interface{}{result}
}

func flattenTasks(tasks []*aws.Task) []interface{} {
	result := make([]interface{}, 0, len(tasks))

	for _, task := range tasks {
		m := make(map[string]interface{})
		m[string(TasksIsEnabled)] = spotinst.BoolValue(task.IsEnabled)
		m[string(TaskType)] = spotinst.StringValue(task.Type)
		m[string(CronExpression)] = spotinst.StringValue(task.CronExpression)
		if task.Parameter != nil {
			m[string(Parameters)] = flattenParameters(task.Parameter)
		}
		result = append(result, m)
	}

	return result
}

func flattenParameters(parameters *aws.Parameter) []interface{} {
	result := make(map[string]interface{})

	if parameters.AmiAutoUpdate != nil {
		result[string(AmiAutoUpdate)] = flattenParametersAmiAutoUpdate(parameters.AmiAutoUpdate)
	}

	if parameters.ClusterRoll != nil {
		result[string(ParametersClusterRoll)] = flattenParameterClusterRoll(parameters.ClusterRoll)
	}

	return []interface{}{result}
}

func flattenParametersAmiAutoUpdate(amiAutoUpdate *aws.AmiAutoUpdate) []interface{} {
	result := make(map[string]interface{})

	result[string(ApplyRoll)] = spotinst.BoolValue(amiAutoUpdate.ApplyRoll)
	if amiAutoUpdate.AmiAutoUpdateClusterRoll != nil {
		result[string(AmiAutoUpdateClusterRoll)] = flattenAmiUpdateClusterRoll(amiAutoUpdate.AmiAutoUpdateClusterRoll)
	}
	result[string(MinorVersion)] = spotinst.BoolValue(amiAutoUpdate.MinorVersion)
	result[string(Patch)] = spotinst.BoolValue(amiAutoUpdate.Patch)

	return []interface{}{result}
}

func flattenAmiUpdateClusterRoll(clusterRoll *aws.AmiAutoUpdateClusterRoll) []interface{} {
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

	return []interface{}{result}
}

func flattenParameterClusterRoll(clusterRoll *aws.ParameterClusterRoll) []interface{} {
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

	return []interface{}{result}
}

func expandScheduledTasks(data interface{}) (*aws.Scheduling, error) {
	if list := data.([]interface{}); (list != nil) || (len(list) > 0 && list[0] != nil) {
		scheduling := &aws.Scheduling{}
		m := list[0].(map[string]interface{})
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

		if v, ok := m[string(ShutdownHours)]; ok {
			shutdownHours, err := expandShutdownHours(v)
			if err != nil {
				return nil, err
			}
			if shutdownHours != nil {
				scheduling.SetShutdownHours(shutdownHours)
			} else {
				scheduling.SetShutdownHours(nil)
			}
		}

		return scheduling, nil
	}
	return nil, nil
}

func expandShutdownHours(data interface{}) (*aws.ShutdownHours, error) {
	if list := data.([]interface{}); len(list) > 0 && list[0] != nil {
		shutdownHours := &aws.ShutdownHours{}
		m := list[0].(map[string]interface{})

		var isEnabled = spotinst.Bool(false)
		if v, ok := m[string(ShutdownHoursIsEnabled)].(bool); ok {
			isEnabled = spotinst.Bool(v)
		}
		shutdownHours.SetIsEnabled(isEnabled)

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
		shutdownHours.SetTimeWindows(timeWindows)

		return shutdownHours, nil
	}

	return nil, nil
}

func expandParameters(data interface{}) (*aws.Parameter, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		parameter := &aws.Parameter{}
		m := list[0].(map[string]interface{})

		if v, ok := m[string(AmiAutoUpdate)]; ok {
			amiAutoUpdate, err := expandAmiAutoUpdate(v)
			if err != nil {
				return nil, err
			}
			if amiAutoUpdate != nil {
				parameter.SetAmiAutoUpdate(amiAutoUpdate)
			} else {
				parameter.AmiAutoUpdate = nil
			}
		}

		if v, ok := m[string(ParametersClusterRoll)]; ok {
			expandClusRoll, err := expandParameterClusterRoll(v)
			if err != nil {
				return nil, err
			}
			if expandClusRoll != nil {
				parameter.SetClusterRoll(expandClusRoll)
			} else {
				parameter.ClusterRoll = nil
			}
		}

		return parameter, nil
	}

	return nil, nil
}

func expandAmiAutoUpdate(data interface{}) (*aws.AmiAutoUpdate, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		amiAutoUpdate := &aws.AmiAutoUpdate{}
		m := list[0].(map[string]interface{})

		var isApplyRoll = spotinst.Bool(false)
		if v, ok := m[string(ApplyRoll)].(bool); ok {
			isApplyRoll = spotinst.Bool(v)
		}
		amiAutoUpdate.SetApplyRoll(isApplyRoll)

		if v, ok := m[string(AmiAutoUpdateClusterRoll)]; ok {
			expandClusRoll, err := expandAmiAutoUpdateClusterRoll(v)
			if err != nil {
				return nil, err
			}
			if expandClusRoll != nil {
				amiAutoUpdate.SetClusterRoll(expandClusRoll)
			} else {
				amiAutoUpdate.AmiAutoUpdateClusterRoll = nil
			}
		}

		var isMinorVersion = spotinst.Bool(false)
		if v, ok := m[string(MinorVersion)].(bool); ok {
			isMinorVersion = spotinst.Bool(v)
		}
		amiAutoUpdate.SetMinorVersion(isMinorVersion)

		var isPatch = spotinst.Bool(false)
		if v, ok := m[string(Patch)].(bool); ok {
			isPatch = spotinst.Bool(v)
		}
		amiAutoUpdate.SetPatch(isPatch)

		return amiAutoUpdate, nil
	}

	return nil, nil
}

func expandParameterClusterRoll(data interface{}) (*aws.ParameterClusterRoll, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		parameterClusterRoll := &aws.ParameterClusterRoll{}
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

		var isComment = spotinst.String("")
		if v, ok := m[string(Comment)].(string); ok {
			isComment = spotinst.String(v)
		}
		parameterClusterRoll.SetComment(isComment)

		var isRespectPdb = spotinst.Bool(false)
		if v, ok := m[string(RespectPdb)].(bool); ok {
			isRespectPdb = spotinst.Bool(v)
		}
		parameterClusterRoll.SetRespectPdb(isRespectPdb)

		return parameterClusterRoll, nil
	}

	return nil, nil
}

func expandAmiAutoUpdateClusterRoll(data interface{}) (*aws.AmiAutoUpdateClusterRoll, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		amiAutoUpdateClusterRoll := &aws.AmiAutoUpdateClusterRoll{}
		m := list[0].(map[string]interface{})

		if v, ok := m[string(BatchMinHealthyPercentage)].(int); ok {
			if v == -1 {
				amiAutoUpdateClusterRoll.SetBatchMinHealthyPercentage(nil)
			} else {
				amiAutoUpdateClusterRoll.SetBatchMinHealthyPercentage(spotinst.Int(v))
			}
		}

		if v, ok := m[string(BatchSizePercentage)].(int); ok {

			if v == -1 {
				amiAutoUpdateClusterRoll.SetBatchSizePercentage(nil)
			} else {
				amiAutoUpdateClusterRoll.SetBatchSizePercentage(spotinst.Int(v))
			}
		}

		var isComment = spotinst.String("")
		if v, ok := m[string(Comment)].(string); ok {
			isComment = spotinst.String(v)
		}
		amiAutoUpdateClusterRoll.SetComment(isComment)

		var isRespectPdb = spotinst.Bool(false)
		if v, ok := m[string(RespectPdb)].(bool); ok {
			isRespectPdb = spotinst.Bool(v)
		}
		amiAutoUpdateClusterRoll.SetRespectPdb(isRespectPdb)

		return amiAutoUpdateClusterRoll, nil
	}

	return nil, nil
}

func expandtasks(data interface{}) ([]*aws.Task, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		tasks := make([]*aws.Task, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			task := &aws.Task{}

			if v, ok := m[string(TasksIsEnabled)].(bool); ok {
				task.SetIsEnabled(spotinst.Bool(v))
			}

			if v, ok := m[string(TaskType)].(string); ok && v != "" {
				task.SetType(spotinst.String(v))
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
					task.SetParameter(parameters)
				} else {
					task.SetParameter(nil)
				}
			}

			tasks = append(tasks, task)
		}

		return tasks, nil
	}
	return nil, nil
}

package ocean_gke_import_scheduling

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ScheduledTask] = commons.NewGenericField(
		commons.OceanGKEImportScheduling,
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
								string(TaskParameters): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(ClusterRoll): {
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
															Default:  false,
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
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster != nil && cluster.Scheduling != nil {
				scheduling := cluster.Scheduling
				result = flattenScheduledTasks(scheduling)
			}

			if result != nil {
				if err := resourceData.Set(string(ScheduledTask), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ScheduledTask), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(ScheduledTask)); ok {
				if scheduling, err := expandScheduledTasks(v); err != nil {
					return err
				} else {
					cluster.SetScheduling(scheduling)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var scheduling *gcp.Scheduling = nil
			if v, ok := resourceData.GetOk(string(ScheduledTask)); ok {
				if interfaces, err := expandScheduledTasks(v); err != nil {
					return err
				} else {
					scheduling = interfaces
				}
			}
			cluster.SetScheduling(scheduling)
			return nil
		},

		nil,
	)

}

func expandShutdownHours(data interface{}) (*gcp.ShutdownHours, error) {
	if list := data.([]interface{}); len(list) > 0 && list[0] != nil {
		runner := &gcp.ShutdownHours{}
		m := list[0].(map[string]interface{})

		var isEnabled = spotinst.Bool(false)
		if v, ok := m[string(ShutdownHoursIsEnabled)].(bool); ok {
			isEnabled = spotinst.Bool(v)
		}
		runner.SetIsEnabled(isEnabled)

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
		runner.SetTimeWindows(timeWindows)

		return runner, nil
	}

	return nil, nil
}

func flattenScheduledTasks(scheduling *gcp.Scheduling) []interface{} {
	var out []interface{}

	if scheduling != nil {
		result := make(map[string]interface{})

		if scheduling.ShutdownHours != nil {
			result[string(ShutdownHours)] = flattenShutdownHours(scheduling.ShutdownHours)
		}

		if len(scheduling.Tasks) > 0 {
			result[string(Tasks)] = flattenTasks(scheduling.Tasks)
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func flattenShutdownHours(shutdownHours *gcp.ShutdownHours) []interface{} {
	result := make(map[string]interface{})
	result[string(ShutdownHoursIsEnabled)] = spotinst.BoolValue(shutdownHours.IsEnabled)

	if shutdownHours.TimeWindows != nil {
		result[string(TimeWindows)] = shutdownHours.TimeWindows
	}

	return []interface{}{result}
}

func flattenTasks(tasks []*gcp.Task) []interface{} {
	result := make([]interface{}, 0, len(tasks))
	for _, task := range tasks {
		m := make(map[string]interface{})
		m[string(TasksIsEnabled)] = spotinst.BoolValue(task.IsEnabled)
		m[string(TaskType)] = spotinst.StringValue(task.Type)
		m[string(CronExpression)] = spotinst.StringValue(task.CronExpression)
		if task.Parameters != nil {
			m[string(TaskParameters)] = flattenParameters(task.Parameters)
		}
		result = append(result, m)
	}
	return result
}
func flattenParameters(parameters *gcp.Parameters) []interface{} {
	result := make(map[string]interface{})

	if parameters.ClusterRoll != nil {
		result[string(ClusterRoll)] = flattenParameterClusterRoll(parameters.ClusterRoll)
	}

	return []interface{}{result}
}
func flattenParameterClusterRoll(clusterRoll *gcp.ClusterRoll) []interface{} {
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

func expandScheduledTasks(data interface{}) (*gcp.Scheduling, error) {
	if list := data.([]interface{}); len(list) > 0 {
		scheduling := &gcp.Scheduling{}
		if list != nil && list[0] != nil {
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
		}
		return scheduling, nil
	}
	return nil, nil

}

func expandtasks(data interface{}) ([]*gcp.Task, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		tasks := make([]*gcp.Task, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			task := &gcp.Task{}

			if v, ok := m[string(TasksIsEnabled)].(bool); ok {
				task.SetIsEnabled(spotinst.Bool(v))
			}

			if v, ok := m[string(TaskType)].(string); ok && v != "" {
				task.SetType(spotinst.String(v))
			}

			if v, ok := m[string(CronExpression)].(string); ok && v != "" {
				task.SetCronExpression(spotinst.String(v))
			}
			if v, ok := m[string(TaskParameters)]; ok {
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

func expandParameters(data interface{}) (*gcp.Parameters, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		parameters := &gcp.Parameters{}
		list := data.([]interface{})
		m := list[0].(map[string]interface{})
		if v, ok := m[string(ClusterRoll)]; ok {
			clusterRoll, err := expandClusterRoll(v)
			if err != nil {
				return nil, err
			}
			if clusterRoll != nil {
				parameters.SetClusterRoll(clusterRoll)
			} else {
				parameters.SetClusterRoll(nil)
			}
		}
		return parameters, nil

	}
	return nil, nil
}

func expandClusterRoll(data interface{}) (*gcp.ClusterRoll, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		clusterRoll := &gcp.ClusterRoll{}
		m := list[0].(map[string]interface{})
		if v, ok := m[string(BatchMinHealthyPercentage)].(int); ok {
			if v == -1 {
				clusterRoll.SetBatchMinHealthyPercentage(nil)
			} else {
				clusterRoll.SetBatchMinHealthyPercentage(spotinst.Int(v))
			}
		}
		if v, ok := m[string(BatchSizePercentage)].(int); ok {
			if v == -1 {
				clusterRoll.SetBatchSizePercentage(nil)
			} else {
				clusterRoll.SetBatchSizePercentage(spotinst.Int(v))
			}
		}
		if v, ok := m[string(Comment)].(string); ok && v != "" {
			clusterRoll.SetComment(spotinst.String(v))
		}
		if v, ok := m[string(RespectPdb)].(bool); ok {
			clusterRoll.SetRespectPdb(spotinst.Bool(v))
		}
		return clusterRoll, nil
	}
	return nil, nil
}

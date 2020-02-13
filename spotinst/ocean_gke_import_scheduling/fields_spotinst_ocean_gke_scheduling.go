package ocean_gke_import_scheduling

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ScheduledTask] = commons.NewGenericField(
		commons.OceanGKEImportScheduling,
		ScheduledTask,
		&schema.Schema{
			Type:     schema.TypeSet,
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

								string(BatchSizePercentage): {
									Type:     schema.TypeInt,
									Optional: true,
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

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandShutdownHours(data interface{}) (*gcp.ShutdownHours, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
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
		var timeWindowList []string = nil
		for _, timeWindow := range shutdownHours.TimeWindows {
			timeWindowList = append(timeWindowList, timeWindow)
		}
		result[string(TimeWindows)] = timeWindowList
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
		m[string(BatchSizePercentage)] = spotinst.IntValue(task.BatchSizePercentage)
		result = append(result, m)
	}
	return result
}

func expandScheduledTasks(data interface{}) (*gcp.Scheduling, error) {
	scheduling := &gcp.Scheduling{}
	list := data.(*schema.Set).List()
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(Tasks)]; ok {
			tasks, err := expandtasks(v)
			if err != nil {
				return nil, err
			}
			if tasks != nil {
				scheduling.SetTasks(tasks)
			}
		}
		if v, ok := m[string(ShutdownHours)]; ok {
			shutdownHours, err := expandShutdownHours(v)
			if err != nil {
				return nil, err
			}
			if shutdownHours != nil {
				if scheduling.ShutdownHours == nil {
					scheduling.SetShutdownHours(&gcp.ShutdownHours{})
				}
				scheduling.SetShutdownHours(shutdownHours)
			}
		}
	}

	return scheduling, nil
}

func expandtasks(data interface{}) ([]*gcp.Task, error) {
	list := data.([]interface{})
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

		if v, ok := m[string(BatchSizePercentage)].(int); ok {
			task.SetBatchSizePercentage(spotinst.Int(v))
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

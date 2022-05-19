package ocean_gke_launch_spec_scheduling

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[SchedulingTask] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		SchedulingTask,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(IsEnabled): {
						Type:     schema.TypeBool,
						Required: true,
					},
					string(CronExpression): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(TaskType): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(TaskHeadroom): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{

								string(CPUPerUnit): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(GPUPerUnit): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(MemoryPerUnit): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(NumOfUnits): {
									Type:     schema.TypeInt,
									Required: true,
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec.LaunchSpecScheduling != nil && launchSpec.LaunchSpecScheduling.Tasks != nil {
				tasks := launchSpec.LaunchSpecScheduling.Tasks
				result = flattenTasks(tasks)
			}
			if result != nil {
				if err := resourceData.Set(string(SchedulingTask), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SchedulingTask), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(SchedulingTask)); ok {
				if tasks, err := expandTasks(value); err != nil {
					return err
				} else {
					launchSpec.LaunchSpecScheduling.SetTasks(tasks)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value []*gcp.GKELaunchSpecTask = nil

			if v, ok := resourceData.GetOk(string(SchedulingTask)); ok {
				if tasks, err := expandTasks(v); err != nil {
					return err
				} else {
					value = tasks
				}
			}
			launchSpec.LaunchSpecScheduling.SetTasks(value)
			return nil
		},
		nil,
	)
}

func flattenTasks(tasks []*gcp.GKELaunchSpecTask) []interface{} {
	result := make([]interface{}, 0, len(tasks))

	for _, task := range tasks {
		m := make(map[string]interface{})
		m[string(IsEnabled)] = spotinst.BoolValue(task.IsEnabled)
		m[string(CronExpression)] = spotinst.StringValue(task.CronExpression)
		m[string(TaskType)] = spotinst.StringValue(task.TaskType)

		if task.Config != nil && task.Config.TaskHeadrooms != nil {
			m[string(TaskHeadroom)] = flattenTaskHeadroom(task.Config.TaskHeadrooms)
		}

		result = append(result, m)
	}

	return result
}

func flattenTaskHeadroom(headrooms []*gcp.GKELaunchSpecTaskHeadroom) []interface{} {
	result := make([]interface{}, 0, len(headrooms))

	for _, headroom := range headrooms {
		m := make(map[string]interface{})
		m[string(CPUPerUnit)] = spotinst.IntValue(headroom.CPUPerUnit)
		m[string(GPUPerUnit)] = spotinst.IntValue(headroom.GPUPerUnit)
		m[string(NumOfUnits)] = spotinst.IntValue(headroom.NumOfUnits)
		m[string(MemoryPerUnit)] = spotinst.IntValue(headroom.MemoryPerUnit)

		result = append(result, m)
	}

	return result
}

func expandTasks(data interface{}) ([]*gcp.GKELaunchSpecTask, error) {
	list := data.(*schema.Set).List()
	tasks := make([]*gcp.GKELaunchSpecTask, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		task := &gcp.GKELaunchSpecTask{}

		if !ok {
			continue
		}

		if v, ok := attr[string(IsEnabled)].(bool); ok {
			task.SetIsEnabled(spotinst.Bool(v))
		}

		if v, ok := attr[string(CronExpression)].(string); ok && v != "" {
			task.SetCronExpression(spotinst.String(v))
		}

		if v, ok := attr[string(TaskType)].(string); ok && v != "" {
			task.SetTaskType(spotinst.String(v))
		}

		if v, ok := attr[string(TaskHeadroom)]; ok {
			if config, err := expandTaskHeadroom(v); err != nil {
				return nil, err
			} else {
				task.SetTaskConfig(config)
			}
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func expandTaskHeadroom(data interface{}) (*gcp.GKETaskConfig, error) {
	list := data.(*schema.Set).List()
	headrooms := make([]*gcp.GKELaunchSpecTaskHeadroom, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		headroom := &gcp.GKELaunchSpecTaskHeadroom{}

		if !ok {
			continue
		}

		if v, ok := attr[string(CPUPerUnit)].(int); ok {
			headroom.SetCPUPerUnit(spotinst.Int(v))
		}

		if v, ok := attr[string(GPUPerUnit)].(int); ok {
			headroom.SetGPUPerUnit(spotinst.Int(v))
		}

		if v, ok := attr[string(NumOfUnits)].(int); ok {
			headroom.SetNumOfUnits(spotinst.Int(v))
		}

		if v, ok := attr[string(MemoryPerUnit)].(int); ok {
			headroom.SetMemoryPerUnit(spotinst.Int(v))
		}

		headrooms = append(headrooms, headroom)
	}

	taskConfig := &gcp.GKETaskConfig{
		TaskHeadrooms: headrooms,
	}

	return taskConfig, nil
}

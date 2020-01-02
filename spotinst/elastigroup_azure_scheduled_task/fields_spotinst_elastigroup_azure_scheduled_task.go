package elastigroup_azure_scheduled_task

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ScheduledTask] = commons.NewGenericField(
		commons.ElastigroupAzureScheduledTask,
		ScheduledTask,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(IsEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  true,
					},

					string(TaskType): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(CronExpression): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(ScaleTargetCapacity): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(ScaleMinCapacity): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(ScaleMaxCapacity): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(BatchSizePercentage): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(GracePeriod): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(Adjustment): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(AdjustmentPercentage): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(ScheduledTask)); ok {
				if tasks, err := expandAzureGroupScheduledTasks(v); err != nil {
					return err
				} else {
					elastigroup.Scheduling.SetTasks(tasks)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []*azure.ScheduledTask = nil
			if v, ok := resourceData.GetOk(string(ScheduledTask)); ok {
				if interfaces, err := expandAzureGroupScheduledTasks(v); err != nil {
					return err
				} else {
					value = interfaces
				}
			}
			elastigroup.Scheduling.SetTasks(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func flattenAzureGroupScheduledTasks(tasks []*azure.ScheduledTask) []interface{} {
	result := make([]interface{}, 0, len(tasks))
	for _, t := range tasks {
		m := make(map[string]interface{})
		m[string(IsEnabled)] = spotinst.BoolValue(t.IsEnabled)
		m[string(TaskType)] = spotinst.StringValue(t.TaskType)
		m[string(CronExpression)] = spotinst.StringValue(t.CronExpression)

		if t.ScaleTargetCapacity != nil {
			m[string(ScaleTargetCapacity)] = strconv.Itoa(spotinst.IntValue(t.ScaleTargetCapacity))
		}
		if t.ScaleMinCapacity != nil {
			m[string(ScaleMinCapacity)] = strconv.Itoa(spotinst.IntValue(t.ScaleMinCapacity))
		}
		if t.ScaleMaxCapacity != nil {
			m[string(ScaleMaxCapacity)] = strconv.Itoa(spotinst.IntValue(t.ScaleMaxCapacity))
		}
		if t.BatchSizePercentage != nil {
			m[string(BatchSizePercentage)] = strconv.Itoa(spotinst.IntValue(t.BatchSizePercentage))
		}
		if t.GracePeriod != nil {
			m[string(GracePeriod)] = strconv.Itoa(spotinst.IntValue(t.GracePeriod))
		}
		if t.Adjustment != nil {
			m[string(Adjustment)] = strconv.Itoa(spotinst.IntValue(t.Adjustment))
		}
		if t.AdjustmentPercentage != nil {
			m[string(AdjustmentPercentage)] = strconv.Itoa(spotinst.IntValue(t.AdjustmentPercentage))
		}
		result = append(result, m)
	}
	return result
}

func expandAzureGroupScheduledTasks(data interface{}) ([]*azure.ScheduledTask, error) {
	list := data.(*schema.Set).List()
	tasks := make([]*azure.ScheduledTask, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		task := &azure.ScheduledTask{}

		if v, ok := m[string(IsEnabled)].(bool); ok {
			task.SetIsEnabled(spotinst.Bool(v))
		}

		if v, ok := m[string(TaskType)].(string); ok && v != "" {
			task.SetTaskType(spotinst.String(v))
		}

		if v, ok := m[string(CronExpression)].(string); ok && v != "" {
			task.SetCronExpression(spotinst.String(v))
		}

		if v, ok := m[string(BatchSizePercentage)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetBatchSizePercentage(spotinst.Int(intVal))
			}
		}

		if v, ok := m[string(GracePeriod)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetGracePeriod(spotinst.Int(intVal))
			}
		}

		if v, ok := m[string(ScaleTargetCapacity)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetScaleTargetCapacity(spotinst.Int(intVal))
			}
		}

		if v, ok := m[string(ScaleMinCapacity)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetScaleMinCapacity(spotinst.Int(intVal))
			}
		}

		if v, ok := m[string(ScaleMaxCapacity)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetScaleMaxCapacity(spotinst.Int(intVal))
			}
		}

		if v, ok := m[string(Adjustment)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetAdjustment(spotinst.Int(intVal))
			}
		}

		if v, ok := m[string(AdjustmentPercentage)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetAdjustmentPercentage(spotinst.Int(intVal))
			}
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

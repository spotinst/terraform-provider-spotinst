package elastigroup_azure_scheduling

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"strconv"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[SchedulingTask] = commons.NewGenericField(
		commons.ElastigroupAzureScheduling,
		SchedulingTask,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(IsEnabled): {
						Type:     schema.TypeBool,
						Required: true,
					},
					string(Type): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(CronExpression): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(ScaleMaxCapacity): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(ScaleMinCapacity): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(ScaleTargetCapacity): {
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
					string(BatchSizePercentage): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(GracePeriod): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{} = nil
			if elastigroup.Scheduling != nil && elastigroup.Scheduling.Tasks != nil {
				tasks := elastigroup.Scheduling.Tasks
				value = flattenAzureSchedulingTasks(tasks)
			}
			if value != nil {
				if err := resourceData.Set(string(SchedulingTask), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SchedulingTask), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(SchedulingTask)); ok {
				if tasks, err := expandAzureSchedulingTasks(v); err != nil {
					return err
				} else {
					elastigroup.Scheduling.SetTasks(tasks)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []*azurev3.Tasks = nil
			if v, ok := resourceData.GetOk(string(SchedulingTask)); ok {
				if tasks, err := expandAzureSchedulingTasks(v); err != nil {
					return err
				} else {
					value = tasks
				}
			}
			elastigroup.Scheduling.SetTasks(value)
			return nil
		},
		nil,
	)
}

func flattenAzureSchedulingTasks(tasks []*azurev3.Tasks) []interface{} {
	result := make([]interface{}, 0, len(tasks))
	for _, t := range tasks {
		m := make(map[string]interface{})
		m[string(IsEnabled)] = spotinst.BoolValue(t.IsEnabled)
		m[string(Type)] = spotinst.StringValue(t.Type)
		m[string(CronExpression)] = spotinst.StringValue(t.CronExpression)

		if t.Adjustment != nil {
			m[string(Adjustment)] = strconv.Itoa(spotinst.IntValue(t.Adjustment))
		}
		if t.AdjustmentPercentage != nil {
			m[string(AdjustmentPercentage)] = strconv.Itoa(spotinst.IntValue(t.AdjustmentPercentage))
		}
		if t.GracePeriod != nil {
			m[string(GracePeriod)] = strconv.Itoa(spotinst.IntValue(t.GracePeriod))
		}
		if t.ScaleMaxCapacity != nil {
			m[string(ScaleMaxCapacity)] = strconv.Itoa(spotinst.IntValue(t.ScaleMaxCapacity))
		}
		if t.ScaleMinCapacity != nil {
			m[string(ScaleMinCapacity)] = strconv.Itoa(spotinst.IntValue(t.ScaleMinCapacity))
		}
		if t.ScaleTargetCapacity != nil {
			m[string(ScaleTargetCapacity)] = strconv.Itoa(spotinst.IntValue(t.ScaleTargetCapacity))
		}
		if t.BatchSizePercentage != nil {
			m[string(BatchSizePercentage)] = strconv.Itoa(spotinst.IntValue(t.BatchSizePercentage))
		}
		result = append(result, m)
	}
	return result
}

func expandAzureSchedulingTasks(data interface{}) ([]*azurev3.Tasks, error) {
	list := data.([]interface{})
	tasks := make([]*azurev3.Tasks, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		task := &azurev3.Tasks{}

		if v, ok := m[string(IsEnabled)].(bool); ok {
			task.SetIsEnabled(spotinst.Bool(v))
		}

		if v, ok := m[string(Type)].(string); ok && v != "" {
			task.SetType(spotinst.String(v))
		}

		if v, ok := m[string(CronExpression)].(string); ok && v != "" {
			task.SetCronExpression(spotinst.String(v))
		}

		if v, ok := m[string(ScaleMaxCapacity)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetScaleMaxCapacity(spotinst.Int(intVal))
			}
		}

		if v, ok := m[string(ScaleMinCapacity)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetScaleMinCapacity(spotinst.Int(intVal))
			}
		}

		if v, ok := m[string(ScaleTargetCapacity)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetScaleTargetCapacity(spotinst.Int(intVal))
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

		if v, ok := m[string(GracePeriod)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetGracePeriod(spotinst.Int(intVal))
			}
		}

		if v, ok := m[string(BatchSizePercentage)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetBatchSizePercentage(spotinst.Int(intVal))
			}
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

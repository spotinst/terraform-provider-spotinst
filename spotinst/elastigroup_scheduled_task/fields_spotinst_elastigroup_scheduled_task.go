package elastigroup_scheduled_task

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ScheduledTask] = commons.NewGenericField(
		commons.ElastigroupScheduledTask,
		ScheduledTask,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(IsEnabled): &schema.Schema{
						Type:     schema.TypeBool,
						Optional: true,
						Default:  true,
					},

					string(TaskType): &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},

					string(Frequency): &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},

					string(CronExpression): &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},

					string(ScaleTargetCapacity): &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(ScaleMinCapacity): &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(ScaleMaxCapacity): &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(BatchSizePercentage): &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(GracePeriod): &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(TargetCapacity): &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(MinCapacity): &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(MaxCapacity): &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value []interface{} = nil
			if elastigroup.Scheduling != nil && elastigroup.Scheduling.Tasks != nil {
				tasks := elastigroup.Scheduling.Tasks
				value = flattenAWSGroupScheduledTasks(tasks)
			}
			if value != nil {
				if err := resourceData.Set(string(ScheduledTask), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ScheduledTask), err)
				}
			} else {
				if err := resourceData.Set(string(ScheduledTask), []*aws.Task{}); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ScheduledTask), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.GetOk(string(ScheduledTask)); ok {
				if tasks, err := expandAWSGroupScheduledTasks(v); err != nil {
					return err
				} else {
					elastigroup.Scheduling.SetTasks(tasks)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value []*aws.Task = nil
			if v, ok := resourceData.GetOk(string(ScheduledTask)); ok {
				if interfaces, err := expandAWSGroupScheduledTasks(v); err != nil {
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
func flattenAWSGroupScheduledTasks(tasks []*aws.Task) []interface{} {
	result := make([]interface{}, 0, len(tasks))
	for _, t := range tasks {
		m := make(map[string]interface{})
		m[string(IsEnabled)] = spotinst.BoolValue(t.IsEnabled)
		m[string(TaskType)] = spotinst.StringValue(t.Type)
		m[string(CronExpression)] = spotinst.StringValue(t.CronExpression)
		m[string(Frequency)] = spotinst.StringValue(t.Frequency)
		m[string(ScaleTargetCapacity)] = spotinst.IntValue(t.ScaleTargetCapacity)
		m[string(ScaleMinCapacity)] = spotinst.IntValue(t.ScaleMinCapacity)
		m[string(ScaleMaxCapacity)] = spotinst.IntValue(t.ScaleMaxCapacity)
		m[string(BatchSizePercentage)] = spotinst.IntValue(t.BatchSizePercentage)
		m[string(GracePeriod)] = spotinst.IntValue(t.GracePeriod)
		m[string(TargetCapacity)] = spotinst.IntValue(t.TargetCapacity)
		m[string(MinCapacity)] = spotinst.IntValue(t.MinCapacity)
		m[string(MaxCapacity)] = spotinst.IntValue(t.MaxCapacity)
		result = append(result, m)
	}
	return result
}

func expandAWSGroupScheduledTasks(data interface{}) ([]*aws.Task, error) {
	list := data.(*schema.Set).List()
	tasks := make([]*aws.Task, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		task := &aws.Task{}

		if v, ok := m[string(IsEnabled)].(bool); ok {
			task.SetIsEnabled(spotinst.Bool(v))
		}

		if v, ok := m[string(TaskType)].(string); ok && v != "" {
			task.SetType(spotinst.String(v))
		}

		if v, ok := m[string(Frequency)].(string); ok && v != "" {
			task.SetFrequency(spotinst.String(v))
		}

		if v, ok := m[string(CronExpression)].(string); ok && v != "" {
			task.SetCronExpression(spotinst.String(v))
		}

		if v, ok := m[string(BatchSizePercentage)].(int); ok && v > 0 {
			task.SetBatchSizePercentage(spotinst.Int(v))
		}

		if v, ok := m[string(GracePeriod)].(int); ok && v > 0 {
			task.SetGracePeriod(spotinst.Int(v))
		}

		if spotinst.StringValue(task.Type) != TaskTypeStatefulUpdateCapacity {
			if v, ok := m[string(ScaleTargetCapacity)].(int); ok && v >= 0 {
				task.SetScaleTargetCapacity(spotinst.Int(v))
			}

			if v, ok := m[string(ScaleMinCapacity)].(int); ok && v >= 0 {
				task.SetScaleMinCapacity(spotinst.Int(v))
			}

			if v, ok := m[string(ScaleMaxCapacity)].(int); ok && v >= 0 {
				task.SetScaleMaxCapacity(spotinst.Int(v))
			}
		}

		if spotinst.StringValue(task.Type) == TaskTypeStatefulUpdateCapacity {
			if v, ok := m[string(TargetCapacity)].(int); ok && v >= 0 {
				task.SetTargetCapacity(spotinst.Int(v))
			}

			if v, ok := m[string(MinCapacity)].(int); ok && v >= 0 {
				task.SetMinCapacity(spotinst.Int(v))
			}

			if v, ok := m[string(MaxCapacity)].(int); ok && v >= 0 {
				task.SetMaxCapacity(spotinst.Int(v))
			}
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

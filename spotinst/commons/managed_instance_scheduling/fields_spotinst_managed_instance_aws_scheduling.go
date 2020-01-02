package managed_instance_scheduling

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ScheduledTask] = commons.NewGenericField(
		commons.ManagedInstanceAWSScheduling,
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

					string(Frequency): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(CronExpression): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(StartTime): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value []interface{} = nil
			if managedInstance.Scheduling != nil && managedInstance.Scheduling.Tasks != nil {
				tasks := managedInstance.Scheduling.Tasks
				value = flattenAWSManagedInstanceScheduledTasks(tasks)
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
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.GetOk(string(ScheduledTask)); ok {
				if tasks, err := expandAWSManagedInstanceScheduledTasks(v); err != nil {
					return err
				} else {
					managedInstance.Scheduling.SetTasks(tasks)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value []*aws.Task = nil
			if v, ok := resourceData.GetOk(string(ScheduledTask)); ok {
				if interfaces, err := expandAWSManagedInstanceScheduledTasks(v); err != nil {
					return err
				} else {
					value = interfaces
				}
			}
			managedInstance.Scheduling.SetTasks(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func flattenAWSManagedInstanceScheduledTasks(tasks []*aws.Task) []interface{} {
	result := make([]interface{}, 0, len(tasks))
	for _, t := range tasks {
		m := make(map[string]interface{})
		m[string(IsEnabled)] = spotinst.BoolValue(t.IsEnabled)
		m[string(TaskType)] = spotinst.StringValue(t.Type)
		m[string(CronExpression)] = spotinst.StringValue(t.CronExpression)
		m[string(StartTime)] = spotinst.StringValue(t.StartTime)
		m[string(Frequency)] = spotinst.StringValue(t.Frequency)

		result = append(result, m)
	}
	return result
}

func expandAWSManagedInstanceScheduledTasks(data interface{}) ([]*aws.Task, error) {
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

		if v, ok := m[string(StartTime)].(string); ok && v != "" {
			task.SetStartTime(spotinst.String(v))
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

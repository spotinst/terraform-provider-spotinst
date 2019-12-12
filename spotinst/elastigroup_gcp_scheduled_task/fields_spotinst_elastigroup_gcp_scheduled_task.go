package elastigroup_gcp_scheduled_task

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ScheduledTask] = commons.NewGenericField(
		commons.ElastigroupGCPScheduledTask,
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
						Optional: true,
					},

					string(TargetCapacity): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(MinCapacity): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(MaxCapacity): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{} = nil
			if elastigroup.Scheduling != nil && elastigroup.Scheduling.Tasks != nil {
				tasks := elastigroup.Scheduling.Tasks
				value = flattenGCPGroupScheduledTasks(tasks)
			}
			if value != nil {
				if err := resourceData.Set(string(ScheduledTask), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ScheduledTask), err)
				}
			} else {
				if err := resourceData.Set(string(ScheduledTask), []*gcp.Task{}); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ScheduledTask), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(ScheduledTask)); ok {
				if tasks, err := expandGCPGroupScheduledTasks(v); err != nil {
					return err
				} else {
					elastigroup.Scheduling.SetTasks(tasks)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []*gcp.Task = nil
			if v, ok := resourceData.GetOk(string(ScheduledTask)); ok {
				if interfaces, err := expandGCPGroupScheduledTasks(v); err != nil {
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
func flattenGCPGroupScheduledTasks(tasks []*gcp.Task) []interface{} {
	result := make([]interface{}, 0, len(tasks))
	for _, t := range tasks {
		m := make(map[string]interface{})
		m[string(IsEnabled)] = spotinst.BoolValue(t.IsEnabled)
		m[string(TaskType)] = spotinst.StringValue(t.Type)
		m[string(CronExpression)] = spotinst.StringValue(t.CronExpression)

		if t.TargetCapacity != nil {
			m[string(TargetCapacity)] = strconv.Itoa(spotinst.IntValue(t.TargetCapacity))
		}
		if t.MinCapacity != nil {
			m[string(MinCapacity)] = strconv.Itoa(spotinst.IntValue(t.MinCapacity))
		}
		if t.MaxCapacity != nil {
			m[string(MaxCapacity)] = strconv.Itoa(spotinst.IntValue(t.MaxCapacity))
		}
		result = append(result, m)
	}
	return result
}

func expandGCPGroupScheduledTasks(data interface{}) ([]*gcp.Task, error) {
	list := data.(*schema.Set).List()
	tasks := make([]*gcp.Task, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		task := &gcp.Task{}

		if v, ok := m[string(IsEnabled)].(bool); ok {
			task.SetIsEnabled(spotinst.Bool(v))
		}

		if v, ok := m[string(TaskType)].(string); ok && v != "" {
			task.SetType(spotinst.String(v))
		}

		if v, ok := m[string(CronExpression)].(string); ok && v != "" {
			task.SetCronExpression(spotinst.String(v))
		}

		if v, ok := m[string(TargetCapacity)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetTargetCapacity(spotinst.Int(intVal))
			}
		}

		if v, ok := m[string(MinCapacity)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetMinCapacity(spotinst.Int(intVal))
			}
		}

		if v, ok := m[string(MaxCapacity)].(string); ok && v != "" {
			if intVal, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				task.SetMaxCapacity(spotinst.Int(intVal))
			}
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

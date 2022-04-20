package stateful_node_azure_scheduling

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Task] = commons.NewGenericField(
		commons.StatefulNodeAzureScheduling,
		Task,
		&schema.Schema{
			Type:     schema.TypeSet,
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
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value []interface{} = nil
			if statefulNode.Scheduling != nil && statefulNode.Scheduling.Tasks != nil {
				tasks := statefulNode.Scheduling.Tasks
				value = flattenStatefulNodeAzureTasks(tasks)
			}
			if value != nil {
				if err := resourceData.Set(string(Task), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Task), err)
				}
			} else {
				if err := resourceData.Set(string(Task), []*azure.Task{}); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Task), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(Task)); ok {
				if tasks, err := expandStatefulNodeAzureTasks(v); err != nil {
					return err
				} else {
					statefulNode.Scheduling.SetTasks(tasks)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value []*azure.Task = nil
			if v, ok := resourceData.GetOk(string(Task)); ok {
				if interfaces, err := expandStatefulNodeAzureTasks(v); err != nil {
					return err
				} else {
					value = interfaces
				}
			}
			statefulNode.Scheduling.SetTasks(value)
			return nil
		},
		nil,
	)
}

func flattenStatefulNodeAzureTasks(tasks []*azure.Task) []interface{} {
	result := make([]interface{}, 0, len(tasks))
	for _, t := range tasks {
		m := make(map[string]interface{})
		m[string(IsEnabled)] = spotinst.BoolValue(t.IsEnabled)
		m[string(Type)] = spotinst.StringValue(t.Type)
		m[string(CronExpression)] = spotinst.StringValue(t.CronExpression)

		result = append(result, m)
	}
	return result
}

func expandStatefulNodeAzureTasks(data interface{}) ([]*azure.Task, error) {
	list := data.(*schema.Set).List()
	tasks := make([]*azure.Task, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		task := &azure.Task{}

		if v, ok := m[string(IsEnabled)].(bool); ok {
			task.SetIsEnabled(spotinst.Bool(v))
		}

		if v, ok := m[string(Type)].(string); ok && v != "" {
			task.SetType(spotinst.String(v))
		}

		if v, ok := m[string(CronExpression)].(string); ok && v != "" {
			task.SetCronExpression(spotinst.String(v))
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

package ocean_aws_scheduling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"log"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ScheduledTask] = commons.NewGenericField(
		commons.OceanAWSScheduling,
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
			//clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			//
			//log.Printf("#################### 1")
			//cluster := clusterWrapper.GetCluster()
			//
			//log.Printf("#################### 2")
			////var shutdownHoursResult []interface{} = nil
			////var taskResult []interface{} = nil
			//
			//if cluster.Scheduling != nil {
			//	scheduling := cluster.Scheduling
			//	log.Printf("#################### 3")
			//	log.Printf("#################### scheduling %s", scheduling)
			//	if scheduling.Tasks != nil {
			//		log.Printf("#################### 8")
			//		taskResult := flattenTasks(scheduling.Tasks)
			//
			//		log.Printf("#################### 9")
			//		if taskResult != nil {
			//			if err := resourceData.Set(string(Tasks), taskResult); err != nil {
			//				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tasks), err)
			//			}
			//		}
			//	}
			//	log.Printf("#################### 7")
			//	if scheduling.ShutdownHours != nil {
			//		log.Printf("#################### 4")
			//		shutdownHoursResult := flattenShutdownHours(scheduling.ShutdownHours)
			//		log.Printf("#################### 5")
			//		log.Printf("#################### shutdownHoursResult %s", shutdownHoursResult)
			//
			//		if shutdownHoursResult != nil {
			//			log.Printf("#################### 51")
			//			log.Printf("#################### resourceData %s", resourceData)
			//			if err := resourceData.Set(string(ShutdownHours), shutdownHoursResult); err != nil {
			//				log.Printf("#################### 52")
			//				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShutdownHours), err)
			//			}
			//		}
			//		log.Printf("#################### 6")
			//	}
			//	log.Printf("#################### 10")
			//} else {
			//	if err := resourceData.Set(string(ScheduledTask), []*aws.Scheduling{}); err != nil {
			//		log.Printf("#################### 11")
			//		return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ScheduledTask), err)
			//
			//	}
			//}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			log.Printf("#################### 7")

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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var scheduling *aws.Scheduling = nil
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
func flattenScheduledTasks(scheduling *aws.Scheduling) []interface{} {
	result := make(map[string]interface{})

	log.Printf("#################### 11")
	log.Printf("################### result %s", result)
	if scheduling.ShutdownHours != nil {
		log.Printf("#################### 12")
		log.Printf("#################### scheduling.ShutdownHours %s", scheduling.ShutdownHours)
		result[string(ShutdownHours)] = flattenShutdownHours(scheduling.ShutdownHours)
		log.Printf("#################### result %s", result)
	}
	if scheduling.Tasks != nil {
		log.Printf("#################### 13")
		log.Printf("#################### scheduling.Tasks %s", scheduling.Tasks)
		result[string(Tasks)] = flattenTasks(scheduling.Tasks)
		log.Printf("#################### result %s", result)
	}

	return []interface{}{result}
}

func flattenShutdownHours(shutdownHours *aws.ShutdownHours) []interface{} {
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

func flattenTasks(tasks []*aws.Task) []interface{} {
	result := make([]interface{}, 0, len(tasks))
	for _, task := range tasks {
		m := make(map[string]interface{})
		m[string(TasksIsEnabled)] = spotinst.BoolValue(task.IsEnabled)
		m[string(TaskType)] = spotinst.StringValue(task.Type)
		m[string(CronExpression)] = spotinst.StringValue(task.CronExpression)
		result = append(result, m)
	}
	return result
}

func expandScheduledTasks(data interface{}) (*aws.Scheduling, error) {
	scheduling := &aws.Scheduling{}
	log.Printf("############### 21")
	list := data.(*schema.Set).List()
	log.Printf("############### 22")
	if list != nil && list[0] != nil {
		log.Printf("############### 23")
		m := list[0].(map[string]interface{})
		log.Printf("############### 24")

		if v, ok := m[string(Tasks)]; ok {
			log.Printf("############### 25")
			tasks, err := expandtasks(v)
			log.Printf("############### 26")
			if err != nil {
				return nil, err
			}
			if tasks != nil {
				scheduling.SetTasks(tasks)
			}
		}

		log.Printf("############### 27")
		if v, ok := m[string(ShutdownHours)]; ok {
			log.Printf("############### 28")
			shutdownHours, err := expandShutdownHours(v)
			log.Printf("############### 29")
			if err != nil {
				return nil, err
			}
			if shutdownHours != nil {
				if scheduling.ShutdownHours == nil {
					scheduling.SetShutdownHours(&aws.ShutdownHours{})
				}
				scheduling.SetShutdownHours(shutdownHours)
			}
		}
		log.Printf("############### 30")
	}

	return scheduling, nil
}

func expandShutdownHours(data interface{}) (*aws.ShutdownHours, error) {
	if list := data.([]interface{}); list != nil && len(list) > 0 && list[0] != nil {
		runner := &aws.ShutdownHours{}
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

func expandtasks(data interface{}) ([]*aws.Task, error) {
	log.Printf("############### 31")
	log.Printf("############### data %s", data)
	list := data.([]interface{})
	log.Printf("############### 32")
	log.Printf("############### list %s", list)
	tasks := make([]*aws.Task, 0, len(list))
	log.Printf("############### 33")
	for _, item := range list {
		m := item.(map[string]interface{})
		task := &aws.Task{}

		log.Printf("############### 34")
		if v, ok := m[string(TasksIsEnabled)].(bool); ok {
			task.SetIsEnabled(spotinst.Bool(v))
		}

		log.Printf("############### 35")
		if v, ok := m[string(TaskType)].(string); ok && v != "" {
			task.SetType(spotinst.String(v))
		}

		log.Printf("############### 36")
		if v, ok := m[string(CronExpression)].(string); ok && v != "" {
			task.SetCronExpression(spotinst.String(v))
		}

		log.Printf("############### 37")
		tasks = append(tasks, task)
	}

	return tasks, nil
}

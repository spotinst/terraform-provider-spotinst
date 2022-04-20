package stateful_node_azure_health

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[HealthCheckTypes] = commons.NewGenericField(
		commons.StatefulNodeAzureHealth,
		HealthCheckTypes,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var result []string
			if statefulNode.Health != nil && statefulNode.Health.HealthCheckTypes != nil {
				result = append(result, statefulNode.Health.HealthCheckTypes...)
				if err := resourceData.Set(string(HealthCheckTypes), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(HealthCheckTypes), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(HealthCheckTypes)); ok {
				checkTypes := v.([]interface{})
				HealthCheckTypes := make([]string, len(checkTypes))
				for i, j := range checkTypes {
					HealthCheckTypes[i] = j.(string)
				}
				statefulNode.Health.SetHealthCheckTypes(HealthCheckTypes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(HealthCheckTypes)); ok {
				checkTypes := v.([]interface{})
				HealthCheckTypes := make([]string, len(checkTypes))
				for i, j := range checkTypes {
					HealthCheckTypes[i] = j.(string)
				}
				statefulNode.Health.SetHealthCheckTypes(HealthCheckTypes)
			}
			return nil
		},
		nil,
	)

	fieldsMap[GracePeriod] = commons.NewGenericField(
		commons.StatefulNodeAzureHealth,
		GracePeriod,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *int = nil
			if statefulNode.Health != nil && statefulNode.Health.GracePeriod != nil {
				value = statefulNode.Health.GracePeriod
			}
			if err := resourceData.Set(string(GracePeriod), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(GracePeriod), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(GracePeriod)).(int); ok && v >= 0 {
				statefulNode.Health.SetGracePeriod(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *int = nil
			if v, ok := resourceData.Get(string(GracePeriod)).(int); ok && v >= 0 {
				value = spotinst.Int(v)
			}
			statefulNode.Health.SetGracePeriod(value)
			return nil
		},
		nil,
	)

	fieldsMap[UnhealthyDuration] = commons.NewGenericField(
		commons.StatefulNodeAzureHealth,
		UnhealthyDuration,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *int = nil
			if statefulNode.Health != nil && statefulNode.Health.UnhealthyDuration != nil {
				value = statefulNode.Health.UnhealthyDuration
			}
			if err := resourceData.Set(string(UnhealthyDuration), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UnhealthyDuration), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(UnhealthyDuration)).(int); ok && v >= 0 {
				statefulNode.Health.SetUnhealthyDuration(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *int = nil
			if v, ok := resourceData.Get(string(UnhealthyDuration)).(int); ok && v >= 0 {
				value = spotinst.Int(v)
			}
			statefulNode.Health.SetUnhealthyDuration(value)
			return nil
		},
		nil,
	)

	fieldsMap[AutoHealing] = commons.NewGenericField(
		commons.StatefulNodeAzureHealth,
		AutoHealing,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *bool = nil
			if statefulNode.Health != nil && statefulNode.Health.AutoHealing != nil {
				value = statefulNode.Health.AutoHealing
			}
			if err := resourceData.Set(string(AutoHealing), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AutoHealing), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(AutoHealing)).(bool); ok {
				statefulNode.Health.SetAutoHealing(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(AutoHealing)).(bool); ok {
				statefulNode.Health.SetAutoHealing(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)
}

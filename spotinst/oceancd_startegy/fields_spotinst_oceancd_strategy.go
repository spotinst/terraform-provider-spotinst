package oceancd_startegy

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanCDStrategy,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			strategyWrapper := resourceObject.(*commons.OceanCDStrategyWrapper)
			strategy := strategyWrapper.GetStrategy()
			var value *string = nil
			if strategy.Name != nil {
				value = strategy.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			strategyWrapper := resourceObject.(*commons.OceanCDStrategyWrapper)
			strategy := strategyWrapper.GetStrategy()
			if value, ok := resourceData.Get(string(Name)).(string); ok && value != "" {
				strategy.SetName(spotinst.String(value))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)
}

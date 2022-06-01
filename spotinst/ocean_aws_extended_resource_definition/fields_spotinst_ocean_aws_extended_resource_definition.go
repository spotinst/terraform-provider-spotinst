package ocean_aws_extended_resource_definition

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[ExtendedResourceName] = commons.NewGenericField(
		commons.OceanAWSExtendedResourceDefinition,
		ExtendedResourceName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			erdWrapper := resourceObject.(*commons.ExtendedResourceDefinitionWrapper)
			erd := erdWrapper.GetOceanAWSExtendedResourceDefinition()
			var value *string = nil
			if erd.Name != nil {
				value = erd.Name
			}
			if err := resourceData.Set(string(ExtendedResourceName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ExtendedResourceName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			erdWrapper := resourceObject.(*commons.ExtendedResourceDefinitionWrapper)
			erd := erdWrapper.GetOceanAWSExtendedResourceDefinition()
			erd.SetName(spotinst.String(resourceData.Get(string(ExtendedResourceName)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			erdWrapper := resourceObject.(*commons.ExtendedResourceDefinitionWrapper)
			erd := erdWrapper.GetOceanAWSExtendedResourceDefinition()
			erd.SetName(spotinst.String(resourceData.Get(string(ExtendedResourceName)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Mapping] = commons.NewGenericField(
		commons.OceanAWSExtendedResourceDefinition,
		Mapping,
		&schema.Schema{
			Type:     schema.TypeMap,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			erdWrapper := resourceObject.(*commons.ExtendedResourceDefinitionWrapper)
			erd := erdWrapper.GetOceanAWSExtendedResourceDefinition()
			var value map[string]interface{} = nil
			if erd.Mapping != nil {
				value = erd.Mapping
			}
			if err := resourceData.Set(string(Mapping), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Mapping), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			erdWrapper := resourceObject.(*commons.ExtendedResourceDefinitionWrapper)
			erd := erdWrapper.GetOceanAWSExtendedResourceDefinition()
			if v, ok := resourceData.Get(string(Mapping)).(map[string]interface{}); ok {
				erd.SetMapping(v)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			erdWrapper := resourceObject.(*commons.ExtendedResourceDefinitionWrapper)
			erd := erdWrapper.GetOceanAWSExtendedResourceDefinition()
			if v, ok := resourceData.Get(string(Mapping)).(map[string]interface{}); ok {
				erd.SetMapping(v)
			}
			return nil
		},

		nil,
	)

}

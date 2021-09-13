package elastigroup_azure_launchspecification

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[CustomData] = commons.NewGenericField(
		commons.ElastigroupAWSStrategy,
		CustomData,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.CustomData != nil {
				value = elastigroup.Compute.LaunchSpecification.CustomData
			}
			if err := resourceData.Set(string(CustomData), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CustomData), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(CustomData)).(string); ok && v != "" {
				customData := spotinst.String(v)
				elastigroup.Compute.LaunchSpecification.SetCustomData(customData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var customData *string = nil
			if v, ok := resourceData.Get(string(CustomData)).(string); ok && v != "" {
				customData = spotinst.String(v)
			}
			elastigroup.Compute.LaunchSpecification.SetCustomData(customData)
			return nil
		},
		nil,
	)
}

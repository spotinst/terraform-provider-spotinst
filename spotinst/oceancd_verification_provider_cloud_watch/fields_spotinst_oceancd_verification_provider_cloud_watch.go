package oceancd_verification_provider_cloud_watch

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[CloudWatch] = commons.NewGenericField(
		commons.OceanCDVerificationProviderCloudWatch,
		CloudWatch,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(IAmArn): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var result []interface{} = nil

			if verificationProvider != nil && verificationProvider.CloudWatch != nil {
				result = flattenCloudWatch(verificationProvider.CloudWatch)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(CloudWatch), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CloudWatch), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var value *oceancd.CloudWatch = nil

			if v, ok := resourceData.GetOk(string(CloudWatch)); ok {
				if cloudWatch, err := expandCloudWatch(v); err != nil {
					return err
				} else {
					value = cloudWatch
				}
			}
			verificationProvider.SetCloudWatch(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var value *oceancd.CloudWatch = nil

			if v, ok := resourceData.GetOk(string(CloudWatch)); ok {
				if cloudWatch, err := expandCloudWatch(v); err != nil {
					return err
				} else {
					value = cloudWatch
				}
			}
			verificationProvider.SetCloudWatch(value)
			return nil
		},
		nil,
	)
}

func expandCloudWatch(data interface{}) (*oceancd.CloudWatch, error) {

	cloudwatch := &oceancd.CloudWatch{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return cloudwatch, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(IAmArn)].(string); ok && v != "" {
		cloudwatch.SetIAmArn(spotinst.String(v))
	} else {
		cloudwatch.SetIAmArn(nil)
	}

	return cloudwatch, nil
}

func flattenCloudWatch(cloudwatch *oceancd.CloudWatch) []interface{} {
	var out []interface{}

	if cloudwatch != nil {
		result := make(map[string]interface{})

		if cloudwatch.IAmArn != nil {
			result[string(IAmArn)] = spotinst.StringValue(cloudwatch.IAmArn)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

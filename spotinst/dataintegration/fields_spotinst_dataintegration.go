package dataintegration

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/dataintegration/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[DataIntegrationName] = commons.NewGenericField(
		commons.DataIntegration,
		DataIntegrationName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			var value *string = nil
			if di.Name != nil {
				value = di.Name
			}
			if err := resourceData.Set(string(DataIntegrationName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DataIntegrationName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			di.SetName(spotinst.String(resourceData.Get(string(DataIntegrationName)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			di.SetName(spotinst.String(resourceData.Get(string(DataIntegrationName)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[S3] = commons.NewGenericField(
		commons.DataIntegration,
		S3,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(BucketName): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(SubDir): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			var value []interface{} = nil
			if di.Config != nil && di.Vendor != nil {
				value = flattenS3(di.Config)
			}
			if err := resourceData.Set(string(S3), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(S3), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			if value, ok := resourceData.GetOk(string(S3)); ok {
				if config, err := expandS3(value); err != nil {
					return err
				} else {
					di.SetConfig(config)
					di.SetVendor(spotinst.String("s3"))
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			var value *aws.Config = nil
			if v, ok := resourceData.GetOk(string(S3)); ok {
				if config, err := expandS3(v); err != nil {
					return err
				} else {
					value = config
				}
			}
			di.SetConfig(value)
			di.SetVendor(spotinst.String("s3"))
			return nil
		},

		nil,
	)

	fieldsMap[Status] = commons.NewGenericField(
		commons.DataIntegration,
		Status,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			var value *string = nil
			if di.Status != nil {
				value = di.Status
			}
			if err := resourceData.Set(string(Status), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Status), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			di.SetStatus(spotinst.String(resourceData.Get(string(Status)).(string)))
			return nil
		},
		nil,
	)
}

func flattenS3(config *aws.Config) []interface{} {
	m := make(map[string]interface{})
	m[string(BucketName)] = spotinst.StringValue(config.BucketName)
	m[string(SubDir)] = spotinst.StringValue(config.SubDir)

	return []interface{}{m}
}

func expandS3(data interface{}) (*aws.Config, error) {
	if list := data.([]interface{}); len(list) > 0 {
		s3 := &aws.Config{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(BucketName)].(string); ok && v != "" {
				s3.SetBucketName(spotinst.String(v))
			}

			if v, ok := m[string(SubDir)].(string); ok && v != "" {
				s3.SetSubDir(spotinst.String(v))
			}
		}
		return s3, nil
	}
	return nil, nil
}

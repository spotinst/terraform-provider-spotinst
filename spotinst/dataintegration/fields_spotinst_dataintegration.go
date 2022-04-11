package dataintegration

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

	fieldsMap[Vendor] = commons.NewGenericField(
		commons.DataIntegration,
		Vendor,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			var value *string = nil
			if di.Vendor != nil {
				value = di.Vendor
			}
			if err := resourceData.Set(string(Vendor), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Vendor), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			di.SetVendor(spotinst.String(resourceData.Get(string(Vendor)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			di.SetVendor(spotinst.String(resourceData.Get(string(Vendor)).(string)))
			return nil
		},

		nil,
	)

	fieldsMap[BucketName] = commons.NewGenericField(
		commons.DataIntegration,
		BucketName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			var value *string = nil
			if di.Config != nil && di.Config.BucketName != nil {
				value = di.Config.BucketName
			}
			if err := resourceData.Set(string(BucketName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BucketName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			if v, ok := resourceData.GetOk(string(BucketName)); ok {
				di.Config.SetBucketName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			if v, ok := resourceData.GetOk(string(BucketName)); ok {
				di.Config.SetBucketName(spotinst.String(v.(string)))
			}
			return nil
		},

		nil,
	)

	fieldsMap[Subdir] = commons.NewGenericField(
		commons.DataIntegration,
		Subdir,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			var value *string = nil
			if di.Config != nil && di.Config.BucketName != nil {
				value = di.Config.BucketName
			}
			if err := resourceData.Set(string(Subdir), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Subdir), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			if v, ok := resourceData.GetOk(string(Subdir)); ok {
				di.Config.SetBucketName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			diWrapper := resourceObject.(*commons.DataIntegrationWrapper)
			di := diWrapper.GetDataIntegration()
			if v, ok := resourceData.GetOk(string(Subdir)); ok {
				di.Config.SetBucketName(spotinst.String(v.(string)))
			}
			return nil
		},

		nil,
	)

}

package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/mrscaler"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

const (
	MRScalerAWSResourceName ResourceName = "spotinst_mrscaler_aws"
)

var MRScalerAWSResource *MRScalerAWSTerraformResource

type MRScalerAWSTerraformResource struct {
	GenericResource // embedding
}

type MRScalerAWSWrapper struct {
	mrscaler *mrscaler.Scaler
}

func NewMRScalerAWSResource(fieldsMap map[FieldName]*GenericField) *MRScalerAWSTerraformResource {
	return &MRScalerAWSTerraformResource{
		GenericResource: GenericResource{
			resourceName: MRScalerAWSResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *MRScalerAWSTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*mrscaler.Scaler, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	mrsWrapper := NewMRScalerAWSWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(mrsWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return mrsWrapper.GetMRScalerAWS(), nil
}

func (res *MRScalerAWSTerraformResource) OnRead(
	mrscaler *mrscaler.Scaler,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	mrsWrapper := NewMRScalerAWSWrapper()
	mrsWrapper.SetMRScalerAWS(mrscaler)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(mrsWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *MRScalerAWSTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *mrscaler.Scaler, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	mrsWrapper := NewMRScalerAWSWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(mrsWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, mrsWrapper.GetMRScalerAWS(), nil
}

func NewMRScalerAWSWrapper() *MRScalerAWSWrapper {
	return &MRScalerAWSWrapper{
		mrscaler: &mrscaler.Scaler{
			Compute: &mrscaler.Compute{
				InstanceGroups: &mrscaler.InstanceGroups{
					TaskGroup: &mrscaler.InstanceGroup{},
				},
			},
			Scaling:     &mrscaler.Scaling{},
			CoreScaling: &mrscaler.Scaling{},
			Strategy:    &mrscaler.Strategy{},
			Scheduling:  &mrscaler.Scheduling{},
		},
	}
}

func (mrsWrapper *MRScalerAWSWrapper) GetMRScalerAWS() *mrscaler.Scaler {
	return mrsWrapper.mrscaler
}

func (mrsWrapper *MRScalerAWSWrapper) SetMRScalerAWS(mrscaler *mrscaler.Scaler) {
	mrsWrapper.mrscaler = mrscaler
}

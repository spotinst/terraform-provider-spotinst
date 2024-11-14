package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
)

const (
	OceanAWSRollResourceName ResourceName = "spotinst_ocean_aws_roll"
)

var OceanAWSRollResource *OceanAWSRollTerraformResource

type OceanAWSRollTerraformResource struct {
	GenericResource
}

type OceanAWSRollWrapper struct {
	roll *aws.RollSpec
}

func NewOceanAWSRollResource(fieldsMap map[FieldName]*GenericField) *OceanAWSRollTerraformResource {
	return &OceanAWSRollTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanAWSRollResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanAWSRollTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.RollSpec, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	rollWrapper := NewOceanAWSRollWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(rollWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return rollWrapper.GetRoll(), nil
}

func (res *OceanAWSRollTerraformResource) OnRead(
	roll *aws.RollSpec,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	rollWrapper := NewOceanAWSRollWrapper()
	rollWrapper.SetRoll(roll)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(rollWrapper, resourceData, meta); err != nil {
			return err
		}
	}

	return nil
}

/*func (res *OceanAWSRollTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *aws.RollSpec, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	rollWrapper := NewOceanAWSRollWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(rollWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, rollWrapper.GetRoll(), nil
}*/

func NewOceanAWSRollWrapper() *OceanAWSRollWrapper {
	return &OceanAWSRollWrapper{
		roll: &aws.RollSpec{},
	}
}

func (rollWrapper *OceanAWSRollWrapper) GetRoll() *aws.RollSpec {
	return rollWrapper.roll
}

func (rollWrapper *OceanAWSRollWrapper) SetRoll(roll *aws.RollSpec) {
	rollWrapper.roll = roll
}

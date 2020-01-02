package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	OceanECSLaunchSpecResourceName ResourceName = "spotinst_ocean_ecs_launch_spec"
)

var OceanECSLaunchSpecResource *OceanECSLaunchSpecTerraformResource

type OceanECSLaunchSpecTerraformResource struct {
	GenericResource // embedding
}

type ECSLaunchSpecWrapper struct {
	launchSpec *aws.ECSLaunchSpec
}

func NewOceanECSLaunchSpecResource(fieldsMap map[FieldName]*GenericField) *OceanECSLaunchSpecTerraformResource {
	return &OceanECSLaunchSpecTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanECSLaunchSpecResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanECSLaunchSpecTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.ECSLaunchSpec, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	launchSpecWrapper := NewLaunchSpecECSWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(launchSpecWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return launchSpecWrapper.GetLaunchSpec(), nil
}

func (res *OceanECSLaunchSpecTerraformResource) OnRead(
	launchSpec *aws.ECSLaunchSpec,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	launchSpecWrapper := NewLaunchSpecECSWrapper()
	launchSpecWrapper.SetLaunchSpec(launchSpec)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(launchSpecWrapper, resourceData, meta); err != nil {
			return err
		}
	}

	return nil
}

func (res *OceanECSLaunchSpecTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *aws.ECSLaunchSpec, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	launchSpecWrapper := NewLaunchSpecECSWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(launchSpecWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, launchSpecWrapper.GetLaunchSpec(), nil
}

func NewLaunchSpecECSWrapper() *ECSLaunchSpecWrapper {
	return &ECSLaunchSpecWrapper{
		launchSpec: &aws.ECSLaunchSpec{
			AutoScale: &aws.ECSAutoScale{},
		},
	}
}

func (launchSpecWrapper *ECSLaunchSpecWrapper) GetLaunchSpec() *aws.ECSLaunchSpec {
	return launchSpecWrapper.launchSpec
}

func (launchSpecWrapper *ECSLaunchSpecWrapper) SetLaunchSpec(launchSpec *aws.ECSLaunchSpec) {
	launchSpecWrapper.launchSpec = launchSpec
}

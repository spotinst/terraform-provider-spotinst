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
	OceanAWSLaunchSpecResourceName ResourceName = "spotinst_ocean_aws_launch_spec"
)

var OceanAWSLaunchSpecResource *OceanAWSLaunchSpecTerraformResource

type OceanAWSLaunchSpecTerraformResource struct {
	GenericResource // embedding
}

type LaunchSpecWrapper struct {
	launchSpec *aws.LaunchSpec
}

func NewOceanAWSLaunchSpecResource(fieldsMap map[FieldName]*GenericField) *OceanAWSLaunchSpecTerraformResource {
	return &OceanAWSLaunchSpecTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanAWSLaunchSpecResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanAWSLaunchSpecTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.LaunchSpec, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	launchSpecWrapper := NewLaunchSpecWrapper()

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

func (res *OceanAWSLaunchSpecTerraformResource) OnRead(
	launchSpec *aws.LaunchSpec,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	launchSpecWrapper := NewLaunchSpecWrapper()
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

func (res *OceanAWSLaunchSpecTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *aws.LaunchSpec, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	launchSpecWrapper := NewLaunchSpecWrapper()
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

func NewLaunchSpecWrapper() *LaunchSpecWrapper {
	return &LaunchSpecWrapper{
		launchSpec: &aws.LaunchSpec{
			AutoScale: &aws.AutoScale{},
		},
	}
}

func (launchSpecWrapper *LaunchSpecWrapper) GetLaunchSpec() *aws.LaunchSpec {
	return launchSpecWrapper.launchSpec
}

func (launchSpecWrapper *LaunchSpecWrapper) SetLaunchSpec(launchSpec *aws.LaunchSpec) {
	launchSpecWrapper.launchSpec = launchSpec
}

package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	OceanGKELaunchSpecImportResourceName ResourceName = "spotinst_ocean_gke_launch_spec_import"
)

var OceanGKELaunchSpecImportResource *OceanGKELaunchSpecImportTerraformResource

type OceanGKELaunchSpecImportTerraformResource struct {
	GenericResource // embedding
}

type GKELaunchSpecImportWrapper struct {
	launchSpec *gcp.LaunchSpec
}

func NewOceanGKELaunchSpecImportResource(fieldsMap map[FieldName]*GenericField) *OceanGKELaunchSpecImportTerraformResource {
	return &OceanGKELaunchSpecImportTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanGKELaunchSpecImportResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanGKELaunchSpecImportTerraformResource) OnCreate(
	importedLaunchSpec *gcp.LaunchSpec,
	resourceData *schema.ResourceData,
	meta interface{}) (*gcp.LaunchSpec, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	launchSpecWrapper := NewGKELaunchSpecImportWrapper()

	if importedLaunchSpec != nil {
		launchSpecWrapper.SetLaunchSpec(importedLaunchSpec)
	}

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

func (res *OceanGKELaunchSpecImportTerraformResource) OnRead(
	launchSpec *gcp.LaunchSpec,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	launchSpecWrapper := NewGKELaunchSpecImportWrapper()
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

func (res *OceanGKELaunchSpecImportTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *gcp.LaunchSpec, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	launchSpecWrapper := NewGKELaunchSpecImportWrapper()
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

func NewGKELaunchSpecImportWrapper() *GKELaunchSpecImportWrapper {
	return &GKELaunchSpecImportWrapper{
		launchSpec: &gcp.LaunchSpec{},
	}
}

func (launchSpecImportWrapper *GKELaunchSpecImportWrapper) GetLaunchSpec() *gcp.LaunchSpec {
	return launchSpecImportWrapper.launchSpec
}

func (launchSpecImportWrapper *GKELaunchSpecImportWrapper) SetLaunchSpec(launchSpecImport *gcp.LaunchSpec) {
	launchSpecImportWrapper.launchSpec = launchSpecImport
}

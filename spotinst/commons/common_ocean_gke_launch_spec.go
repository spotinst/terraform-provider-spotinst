package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
)

const (
	OceanGKELaunchSpecResourceName ResourceName = "spotinst_ocean_gke_launch_spec"
)

var OceanGKELaunchSpecResource *OceanGKELaunchSpecTerraformResource

type OceanGKELaunchSpecTerraformResource struct {
	GenericResource
}

type LaunchSpecGKEWrapper struct {
	launchSpec *gcp.LaunchSpec
}

func NewOceanGKELaunchSpecResource(fieldsMap map[FieldName]*GenericField) *OceanGKELaunchSpecTerraformResource {
	return &OceanGKELaunchSpecTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanGKELaunchSpecResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanGKELaunchSpecTerraformResource) OnCreate(
	importedLaunchSpec *gcp.LaunchSpec,
	resourceData *schema.ResourceData,
	meta interface{}) (*gcp.LaunchSpec, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	launchSpecWrapper := NewGKELaunchSpecWrapper()

	if importedLaunchSpec != nil {
		buildEmptyGKELaunchSpecRequirements(importedLaunchSpec)
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

func (res *OceanGKELaunchSpecTerraformResource) OnRead(
	launchSpec *gcp.LaunchSpec,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	launchSpecWrapper := NewGKELaunchSpecWrapper()
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

func (res *OceanGKELaunchSpecTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *gcp.LaunchSpec, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	launchSpecWrapper := NewGKELaunchSpecWrapper()
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

func NewGKELaunchSpecWrapper() *LaunchSpecGKEWrapper {
	return &LaunchSpecGKEWrapper{
		launchSpec: &gcp.LaunchSpec{
			Metadata: []*gcp.Metadata{},
			Labels:   []*gcp.Label{},
			Taints:   []*gcp.Taint{},
			AutoScale: &gcp.AutoScale{
				Headrooms: []*gcp.AutoScaleHeadroom{},
			},
			ShieldedInstanceConfig: &gcp.ShieldedInstanceConfig{},
			Storage:                &gcp.Storage{},
			LaunchSpecScheduling:   &gcp.GKELaunchSpecScheduling{},
		},
	}
}

func (launchSpecWrapper *LaunchSpecGKEWrapper) GetLaunchSpec() *gcp.LaunchSpec {
	return launchSpecWrapper.launchSpec
}

func (launchSpecWrapper *LaunchSpecGKEWrapper) SetLaunchSpec(launchSpec *gcp.LaunchSpec) {
	launchSpecWrapper.launchSpec = launchSpec
}

func buildEmptyGKELaunchSpecRequirements(launchSpec *gcp.LaunchSpec) {
	if launchSpec == nil {
		return
	}

	if launchSpec.AutoScale == nil {
		launchSpec.SetAutoScale(&gcp.AutoScale{})
	}

	if launchSpec.AutoScale.Headrooms == nil {
		launchSpec.AutoScale.SetHeadrooms([]*gcp.AutoScaleHeadroom{})
	}

	if launchSpec.LaunchSpecScheduling == nil {
		launchSpec.SetScheduling(&gcp.GKELaunchSpecScheduling{})
	}
}

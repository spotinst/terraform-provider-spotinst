package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
)

const (
	OceanCDVerificationTemplateResourceName ResourceName = "spotinst_oceancd_verification_template"
)

var OceanCDVerificationTemplateResource *OceanCDVerificationTemplateTerraformResource

type OceanCDVerificationTemplateTerraformResource struct {
	GenericResource
}

type OceanCDVerificationTemplateWrapper struct {
	VerificationTemplate *oceancd.VerificationTemplate
}

func NewOceanCDVerificationTemplateResource(fieldsMap map[FieldName]*GenericField) *OceanCDVerificationTemplateTerraformResource {
	return &OceanCDVerificationTemplateTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanCDVerificationTemplateResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanCDVerificationTemplateTerraformResource) OnRead(
	VerificationTemplate *oceancd.VerificationTemplate,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	oceancdVPWrapper := NewOceanCDVerificationTemplateWrapper()
	oceancdVPWrapper.SetVerificationTemplate(VerificationTemplate)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(oceancdVPWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *OceanCDVerificationTemplateTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*oceancd.VerificationTemplate, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	oceancdVPWrapper := NewOceanCDVerificationTemplateWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(oceancdVPWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return oceancdVPWrapper.GetVerificationTemplate(), nil
}

func (res *OceanCDVerificationTemplateTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *oceancd.VerificationTemplate, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	oceancdVTWrapper := NewOceanCDVerificationTemplateWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(oceancdVTWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, oceancdVTWrapper.GetVerificationTemplate(), nil
}

func NewOceanCDVerificationTemplateWrapper() *OceanCDVerificationTemplateWrapper {
	return &OceanCDVerificationTemplateWrapper{
		VerificationTemplate: &oceancd.VerificationTemplate{
			Args:    []*oceancd.Args{},
			Metrics: []*oceancd.Metrics{},
		},
	}
}

func (oceancdVTWrapper *OceanCDVerificationTemplateWrapper) GetVerificationTemplate() *oceancd.VerificationTemplate {
	return oceancdVTWrapper.VerificationTemplate
}

func (oceancdVTWrapper *OceanCDVerificationTemplateWrapper) SetVerificationTemplate(VerificationTemplate *oceancd.VerificationTemplate) {
	oceancdVTWrapper.VerificationTemplate = VerificationTemplate
}

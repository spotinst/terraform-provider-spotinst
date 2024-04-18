package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
)

const (
	OceanCDVerificationProviderResourceName ResourceName = "spotinst_oceancd_verification_provider"
)

var OceanCDVerificationProviderResource *OceanCDVerificationProviderTerraformResource

type OceanCDVerificationProviderTerraformResource struct {
	GenericResource
}

type OceanCDVerificationProviderWrapper struct {
	verificationProvider *oceancd.VerificationProvider
}

func NewOceanCDVerificationProviderResource(fieldsMap map[FieldName]*GenericField) *OceanCDVerificationProviderTerraformResource {
	return &OceanCDVerificationProviderTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanCDVerificationProviderResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanCDVerificationProviderTerraformResource) OnRead(
	verificationProvider *oceancd.VerificationProvider,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	oceancdVPWrapper := NewOceanCDVerificationProviderWrapper()
	oceancdVPWrapper.SetVerificationProvider(verificationProvider)

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

func (res *OceanCDVerificationProviderTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*oceancd.VerificationProvider, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	oceancdVPWrapper := NewOceanCDVerificationProviderWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(oceancdVPWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return oceancdVPWrapper.GetVerificationProvider(), nil
}

func (res *OceanCDVerificationProviderTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *oceancd.VerificationProvider, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	oceancdVPWrapper := NewOceanCDVerificationProviderWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(oceancdVPWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, oceancdVPWrapper.GetVerificationProvider(), nil
}

func NewOceanCDVerificationProviderWrapper() *OceanCDVerificationProviderWrapper {
	return &OceanCDVerificationProviderWrapper{
		verificationProvider: &oceancd.VerificationProvider{
			CloudWatch: &oceancd.CloudWatch{},
			DataDog:    &oceancd.DataDog{},
			Jenkins:    &oceancd.Jenkins{},
			NewRelic:   &oceancd.NewRelic{},
			Prometheus: &oceancd.Prometheus{},
		},
	}
}

func (oceancdVPWrapper *OceanCDVerificationProviderWrapper) GetVerificationProvider() *oceancd.VerificationProvider {
	return oceancdVPWrapper.verificationProvider
}

func (oceancdVPWrapper *OceanCDVerificationProviderWrapper) SetVerificationProvider(verificationProvider *oceancd.VerificationProvider) {
	oceancdVPWrapper.verificationProvider = verificationProvider
}

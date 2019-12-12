package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	MultaiDeploymentResourceName ResourceName = "spotinst_multai_deployment"
)

var MultaiDeploymentResource *MultaiDeploymentTerraformResource

type MultaiDeploymentTerraformResource struct {
	GenericResource // embedding
}

type MultaiDeploymentWrapper struct {
	deployment *multai.Deployment
}

func NewMultaiDeploymentResource(fieldMap map[FieldName]*GenericField) *MultaiDeploymentTerraformResource {
	return &MultaiDeploymentTerraformResource{
		GenericResource: GenericResource{
			resourceName: MultaiDeploymentResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

func (res *MultaiDeploymentTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*multai.Deployment, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	mlbWrapper := NewMultaiDeploymentWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(mlbWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return mlbWrapper.GetMultaiDeployment(), nil
}

func (res *MultaiDeploymentTerraformResource) OnRead(
	deployment *multai.Deployment,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	mlbWrapper := NewMultaiDeploymentWrapper()
	mlbWrapper.SetMultaiDeployment(deployment)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(mlbWrapper, resourceData, meta); err != nil {
			return err
		}
	}

	return nil
}

func (res *MultaiDeploymentTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *multai.Deployment, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	mlbWrapper := NewMultaiDeploymentWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(mlbWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, mlbWrapper.GetMultaiDeployment(), nil
}

func NewMultaiDeploymentWrapper() *MultaiDeploymentWrapper {
	return &MultaiDeploymentWrapper{
		deployment: &multai.Deployment{},
	}
}

func (mlbWrapper *MultaiDeploymentWrapper) GetMultaiDeployment() *multai.Deployment {
	return mlbWrapper.deployment
}

func (mlbWrapper *MultaiDeploymentWrapper) SetMultaiDeployment(deployment *multai.Deployment) {
	mlbWrapper.deployment = deployment
}

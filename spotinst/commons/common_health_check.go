package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/healthcheck"
)

const (
	HealthCheckResourceName ResourceName = "spotinst_health_check"
)

var HealthCheckResource *HealthCheckTerraformResource

type HealthCheckTerraformResource struct {
	GenericResource
}

type HealthCheckWrapper struct {
	healthCheck *healthcheck.HealthCheck
}

// NewHealthCheckResource creates a new HealthCheck resource
func NewHealthCheckResource(fieldMap map[FieldName]*GenericField) *HealthCheckTerraformResource {
	return &HealthCheckTerraformResource{
		GenericResource: GenericResource{
			resourceName: HealthCheckResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

// OnCreate is called when creating a new resource block and returns a new HealthCheck or an error.
func (res *HealthCheckTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*healthcheck.HealthCheck, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	hcWrapper := NewHealthCheckWrapper()
	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(hcWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return hcWrapper.GetHealthCheck(), nil
}

// OnRead is called when reading an existing resource and throws an error if it is unable to do so.
func (res *HealthCheckTerraformResource) OnRead(
	healthCheck *healthcheck.HealthCheck,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	hcWrapper := NewHealthCheckWrapper()
	hcWrapper.SetHealthCheck(healthCheck)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(hcWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

// OnUpdate is called when updating an existing resource and returns
// an healthCheck with a bool indicating if had been updated, or an error.
func (res *HealthCheckTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *healthcheck.HealthCheck, error) {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	hcWrapper := NewHealthCheckWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(hcWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, hcWrapper.GetHealthCheck(), nil
}

// NewElastigroupGCPWrapper avoids parameter collisions and returns a HealthCheck.
// Spotinst HealthCheck must have a wrapper struct.
// the wrapper struct is intended to help reflect the field states into the HealthCheck object properly.
func NewHealthCheckWrapper() *HealthCheckWrapper {
	return &HealthCheckWrapper{
		healthCheck: &healthcheck.HealthCheck{},
	}
}

// GetElastigroup returns a wrapped elastigroup
func (hcWrapper *HealthCheckWrapper) GetHealthCheck() *healthcheck.HealthCheck {
	return hcWrapper.healthCheck
}

// SetElastigroup applies elastigroup fields to the elastigroup wrapper.
func (hcWrapper *HealthCheckWrapper) SetHealthCheck(healthCheck *healthcheck.HealthCheck) {
	hcWrapper.healthCheck = healthCheck
}

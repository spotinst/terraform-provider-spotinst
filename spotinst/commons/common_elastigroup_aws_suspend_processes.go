package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
)

const (
	SuspendProcessesResourceName ResourceName = "spotinst_elastigroup_aws_suspension"
)

var SuspendProcessesResource *SuspendProcessesTerraformResource

type SuspendProcessesTerraformResource struct {
	GenericResource
}

type SuspendProcessesWrapper struct {
	GroupID          *string
	SuspendProcesses *aws.SuspendProcesses
}

func NewSuspendProcessesResource(fieldMap map[FieldName]*GenericField) *SuspendProcessesTerraformResource {
	return &SuspendProcessesTerraformResource{
		GenericResource: GenericResource{
			resourceName: SuspendProcessesResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

// OnCreate is called when creating a new resource block and returns a new SuspendProcesses or an error.
func (res *SuspendProcessesTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.SuspendProcesses, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	spWrapper := NewSuspendProcessesWrapper()
	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(spWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return spWrapper.GetSuspendProcesses().SuspendProcesses, nil
}

// OnRead is called when reading an existing resource and throws an error if it is unable to do so.
func (res *SuspendProcessesTerraformResource) OnRead(
	suspendProcesses *aws.SuspendProcesses,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	spWrapper := NewSuspendProcessesWrapper()
	spWrapper.SetSuspendProcesses(suspendProcesses)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(spWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

// OnUpdate is called when updating an existing resource and returns
// a suspedProcesses with a bool indicating if had been updated, or an error.
func (res *SuspendProcessesTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *aws.SuspendProcesses, error) {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	spWrapper := NewSuspendProcessesWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(spWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, spWrapper.GetSuspendProcesses().SuspendProcesses, nil
}

// NewsuspendProcessesWrapper avoids parameter collisions and returns a SuspendProcesses.
// Spotinst SuspendProcesses must have a wrapper struct.
// the wrapper struct is intended to help reflect the field states into the SuspendProcesses object properly.
func NewSuspendProcessesWrapper() *SuspendProcessesWrapper {
	return &SuspendProcessesWrapper{
		SuspendProcesses: &aws.SuspendProcesses{},
	}
}

// GetSuspendProcesses returns a wrapped SuspendProcesses
func (spWrapper *SuspendProcessesWrapper) GetSuspendProcesses() *SuspendProcessesWrapper {
	return spWrapper
}

// SetSuspendProcesses applies SuspendProcesses fields to the SuspendProcesses wrapper.
func (spWrapper *SuspendProcessesWrapper) SetSuspendProcesses(suspendProcesses *aws.SuspendProcesses) {
	spWrapper.SuspendProcesses = suspendProcesses
}

package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	ManagedInstanceAWSResourceName ResourceName = "spotinst_managed_instance_aws"
)

var ManagedInstanceResource *ManagedInstanceTerraformResource

type ManagedInstanceTerraformResource struct {
	GenericResource // embedding
}

type MangedInstanceAWSWrapper struct {
	mangedInstance *aws.ManagedInstance

	// Load balancer states
	StatusElbUpdated bool
	StatusTgUpdated  bool
	StatusMlbUpdated bool

	// Block devices states
	StatusEphemeralBlockDeviceUpdated bool
	StatusEbsBlockDeviceUpdated       bool
}

func NewManagedInstanceResource(fieldsMap map[FieldName]*GenericField) *ManagedInstanceTerraformResource {
	return &ManagedInstanceTerraformResource{
		GenericResource: GenericResource{
			resourceName: ManagedInstanceAWSResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *ManagedInstanceTerraformResource) OnRead(
	managedInstance *aws.ManagedInstance,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}
	miWrapper := NewManagedInstanceWrapper()
	miWrapper.SetManagedInstance(managedInstance)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(miWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *ManagedInstanceTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.ManagedInstance, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}
	miWrapper := NewManagedInstanceWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(miWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return miWrapper.GetManagedInstance(), nil
}

func (res *ManagedInstanceTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *aws.ManagedInstance, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}
	miWrapper := NewManagedInstanceWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(miWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, miWrapper.GetManagedInstance(), nil
}

func NewManagedInstanceWrapper() *MangedInstanceAWSWrapper {
	return &MangedInstanceAWSWrapper{
		mangedInstance: &aws.ManagedInstance{
			Strategy:    &aws.Strategy{},
			HealthCheck: &aws.HealthCheck{},
			Scheduling:  &aws.Scheduling{},
			Integration: &aws.Integration{},
			Compute: &aws.Compute{
				LaunchSpecification: &aws.LaunchSpecification{
					InstanceTypes: &aws.InstanceTypes{},
				},
			},
		},
	}
}

func (miWrapper *MangedInstanceAWSWrapper) GetManagedInstance() *aws.ManagedInstance {
	return miWrapper.mangedInstance
}

func (miWrapper *MangedInstanceAWSWrapper) SetManagedInstance(mangedInstance *aws.ManagedInstance) {
	miWrapper.mangedInstance = mangedInstance
}

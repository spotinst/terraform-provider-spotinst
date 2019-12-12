package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	ElastigroupAzureResourceName ResourceName = "spotinst_elastigroup_azure"
)

var ElastigroupAzureResource *ElastigroupAzureTerraformResource

type ElastigroupAzureTerraformResource struct {
	GenericResource // embedding
}

type ElastigroupAzureWrapper struct {
	elastigroup *azure.Group
}

func NewElastigroupAzureResource(fieldsMap map[FieldName]*GenericField) *ElastigroupAzureTerraformResource {
	return &ElastigroupAzureTerraformResource{
		GenericResource: GenericResource{
			resourceName: ElastigroupAzureResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *ElastigroupAzureTerraformResource) OnRead(
	elastigroup *azure.Group,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	egWrapper := NewElastigroupAzureWrapper()
	egWrapper.SetElastigroup(elastigroup)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(egWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *ElastigroupAzureTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*azure.Group, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	egWrapper := NewElastigroupAzureWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(egWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return egWrapper.GetElastigroup(), nil
}

func (res *ElastigroupAzureTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *azure.Group, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	egWrapper := NewElastigroupAzureWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(egWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, egWrapper.GetElastigroup(), nil
}

// Spotinst elastigroup must have a wrapper struct.
// Reason is that there are multiple fields who share the same elastigroup API object
// e.g. LoadBalancersConfig fields and BlockDeviceMapping fields
// Wrapper struct intended to help reflecting these fields state properly into the elastigroup object.
func NewElastigroupAzureWrapper() *ElastigroupAzureWrapper {
	return &ElastigroupAzureWrapper{
		elastigroup: &azure.Group{
			Scaling:     &azure.Scaling{},
			Scheduling:  &azure.Scheduling{},
			Integration: &azure.Integration{},
			Compute: &azure.Compute{
				LaunchSpecification: &azure.LaunchSpecification{
					LoadBalancersConfig: &azure.LoadBalancersConfig{},
				},
				VMSizes: &azure.VMSizes{},
			},
			Capacity: &azure.Capacity{},
			Strategy: &azure.Strategy{},
		},
	}
}

func (egWrapper *ElastigroupAzureWrapper) GetElastigroup() *azure.Group {
	return egWrapper.elastigroup
}

func (egWrapper *ElastigroupAzureWrapper) SetElastigroup(elastigroup *azure.Group) {
	egWrapper.elastigroup = elastigroup
}

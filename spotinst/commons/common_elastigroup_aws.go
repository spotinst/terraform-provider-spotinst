package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	ElastigroupAWSResourceName ResourceName = "spotinst_elastigroup_aws"
)

var ElastigroupResource *ElastigroupTerraformResource

type ElastigroupTerraformResource struct {
	GenericResource // embedding
}

type ElastigroupWrapper struct {
	elastigroup *aws.Group

	// Load balancer states
	StatusElbUpdated bool
	StatusTgUpdated  bool
	StatusMlbUpdated bool

	// Block devices states
	StatusEphemeralBlockDeviceUpdated bool
	StatusEbsBlockDeviceUpdated       bool
}

func NewElastigroupResource(fieldsMap map[FieldName]*GenericField) *ElastigroupTerraformResource {
	return &ElastigroupTerraformResource{
		GenericResource: GenericResource{
			resourceName: ElastigroupAWSResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *ElastigroupTerraformResource) OnRead(
	elastigroup *aws.Group,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	egWrapper := NewElastigroupWrapper()
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

func (res *ElastigroupTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.Group, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	egWrapper := NewElastigroupWrapper()

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

func (res *ElastigroupTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *aws.Group, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	egWrapper := NewElastigroupWrapper()
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
func NewElastigroupWrapper() *ElastigroupWrapper {
	return &ElastigroupWrapper{
		elastigroup: &aws.Group{
			Scaling:     &aws.Scaling{},
			Scheduling:  &aws.Scheduling{},
			Integration: &aws.Integration{},
			Compute: &aws.Compute{
				LaunchSpecification: &aws.LaunchSpecification{
					LoadBalancersConfig: &aws.LoadBalancersConfig{},
				},
				InstanceTypes: &aws.InstanceTypes{},
			},
			Capacity: &aws.Capacity{},
			Strategy: &aws.Strategy{
				Persistence: &aws.Persistence{},
			},
		},
	}
}

func (egWrapper *ElastigroupWrapper) GetElastigroup() *aws.Group {
	return egWrapper.elastigroup
}

func (egWrapper *ElastigroupWrapper) SetElastigroup(elastigroup *aws.Group) {
	egWrapper.elastigroup = elastigroup
}

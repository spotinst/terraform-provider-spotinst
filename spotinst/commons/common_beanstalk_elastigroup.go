package commons

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"log"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	BeanstalkElastigroupResourceName ResourceName = "spotinst_beanstalk_elastigroup"
)

var ElasticBeanstalkResource *ElasticBeanstalkTerraformResource

type ElasticBeanstalkTerraformResource struct {
	GenericResource // embedding
}

type BeanStalkElastigroupWrapper struct {
	elastigroup              *aws.Group
	BeanstalkEnvironmentName string

	// Load balancer states
	StatusElbUpdated bool
	StatusTgUpdated  bool
	StatusMlbUpdated bool

	// Block devices states
	StatusEphemeralBlockDeviceUpdated bool
	StatusEbsBlockDeviceUpdated       bool
}

func NewElasticBeanstalkResource(fieldsMap map[FieldName]*GenericField) *ElasticBeanstalkTerraformResource {
	return &ElasticBeanstalkTerraformResource{
		GenericResource: GenericResource{
			resourceName: BeanstalkElastigroupResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *ElasticBeanstalkTerraformResource) OnCreate(
	importedGroup *aws.Group,
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.Group, error) {

	buildEmptyElastigroupCapacity(importedGroup)
	buildEmptyElastigroupInstanceTypes(importedGroup)

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	beanstalkGroupWrapper := NewBeanstalkElastigroupWrapper()
	beanstalkGroupWrapper.SetBeanstalkElastigroup(importedGroup)

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)

		if err := field.onCreate(beanstalkGroupWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return beanstalkGroupWrapper.GetBeanstalkElastigroup(), nil
}

func (res *ElasticBeanstalkTerraformResource) OnRead(
	elastigroup *aws.Group,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	beanstalkWrapper := NewBeanstalkElastigroupWrapper()
	beanstalkWrapper.SetBeanstalkElastigroup(elastigroup)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(beanstalkWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *ElasticBeanstalkTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *aws.Group, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}
	beanstalkWrapper := NewBeanstalkElastigroupWrapper()
	hasChanged := false

	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(beanstalkWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, beanstalkWrapper.GetBeanstalkElastigroup(), nil
}

func (res *ElasticBeanstalkTerraformResource) MaintenanceState(
	resourceData *schema.ResourceData,
	meta interface{}) (string, error) {
	op := "NONE"

	if res.fields.fieldsMap["maintenance"] != nil {
		op = resourceData.Get("maintenance").(string)
		return op, nil
	}
	return op, nil
}

// Spotinst elastigroup must have a wrapper struct.
// Reason is that there are multiple fields who share the same elastigroup API object
// e.g. LoadBalancersConfig fields and BlockDeviceMapping fields
// Wrapper struct intended to help reflecting these fields state properly into the elastigroup object.
func NewBeanstalkElastigroupWrapper() *BeanStalkElastigroupWrapper {
	return &BeanStalkElastigroupWrapper{
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

func (egWrapper *BeanStalkElastigroupWrapper) GetBeanstalkElastigroup() *aws.Group {
	return egWrapper.elastigroup
}

func (egWrapper *BeanStalkElastigroupWrapper) SetBeanstalkElastigroup(elastigroup *aws.Group) {
	egWrapper.elastigroup = elastigroup
}

func buildEmptyElastigroupCompute(group *aws.Group) {
	if group != nil && group.Compute == nil {
		group.SetCompute(&aws.Compute{})
	}
}

func buildEmptyElastigroupCapacity(group *aws.Group) {
	if group != nil && group.Capacity == nil {
		group.SetCapacity(&aws.Capacity{})
	}
}

func buildEmptyElastigroupInstanceTypes(group *aws.Group) {
	buildEmptyElastigroupCompute(group)

	if group.Compute.InstanceTypes == nil {
		group.Compute.SetInstanceTypes(&aws.InstanceTypes{})
	}
}

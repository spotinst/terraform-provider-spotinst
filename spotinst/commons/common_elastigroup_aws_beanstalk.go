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
	ElastigroupAWSBeanstalkResourceName ResourceName = "spotinst_elastigroup_aws_beanstalk"
)

var ElastigroupAWSBeanstalkResource *ElastigroupAWSBeanstalkTerraformResource

type ElastigroupAWSBeanstalkTerraformResource struct {
	GenericResource // embedding
}

type ElastigroupAWSBeanstalkWrapper struct {
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

func NewElastigroupAWSBeanstalkResource(fieldsMap map[FieldName]*GenericField) *ElastigroupAWSBeanstalkTerraformResource {
	return &ElastigroupAWSBeanstalkTerraformResource{
		GenericResource: GenericResource{
			resourceName: ElastigroupAWSBeanstalkResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *ElastigroupAWSBeanstalkTerraformResource) OnCreate(
	importedGroup *aws.Group,
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.Group, error) {

	buildEmptyElastigroupCapacity(importedGroup)
	buildEmptyElastigroupInstanceTypes(importedGroup)
	buildEmptyElastigroupScheduling(importedGroup)

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	beanstalkGroupWrapper := NewElastigroupAWSBeanstalkWrapper()
	beanstalkGroupWrapper.SetElastigroupAWSBeanstalk(importedGroup)

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)

		if err := field.onCreate(beanstalkGroupWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return beanstalkGroupWrapper.GetElastigroupAWSBeanstalk(), nil
}

func (res *ElastigroupAWSBeanstalkTerraformResource) OnRead(
	elastigroup *aws.Group,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	beanstalkWrapper := NewElastigroupAWSBeanstalkWrapper()
	beanstalkWrapper.SetElastigroupAWSBeanstalk(elastigroup)

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

func (res *ElastigroupAWSBeanstalkTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *aws.Group, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}
	beanstalkWrapper := NewElastigroupAWSBeanstalkWrapper()
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

	return hasChanged, beanstalkWrapper.GetElastigroupAWSBeanstalk(), nil
}

func (res *ElastigroupAWSBeanstalkTerraformResource) MaintenanceState(
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
func NewElastigroupAWSBeanstalkWrapper() *ElastigroupAWSBeanstalkWrapper {
	return &ElastigroupAWSBeanstalkWrapper{
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

func (egWrapper *ElastigroupAWSBeanstalkWrapper) GetElastigroupAWSBeanstalk() *aws.Group {
	return egWrapper.elastigroup
}

func (egWrapper *ElastigroupAWSBeanstalkWrapper) SetElastigroupAWSBeanstalk(elastigroup *aws.Group) {
	egWrapper.elastigroup = elastigroup
}

func buildEmptyElastigroupCompute(group *aws.Group) {
	if group != nil && group.Compute == nil {
		group.SetCompute(&aws.Compute{})
	}
}

func buildEmptyElastigroupScheduling(group *aws.Group) {
	if group != nil && group.Scheduling == nil {
		group.SetScheduling(&aws.Scheduling{})
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

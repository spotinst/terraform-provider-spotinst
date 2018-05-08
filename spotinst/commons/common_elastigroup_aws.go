package commons

import (
	"fmt"
	"log"
	"sync"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	ElastigroupAwsResourceName ResourceName = "spotinst_elastigroup_aws"
)

var SpotinstElastigroup *ElastigroupResource

type ElastigroupResource struct {
	GenericResource // embedding

	mux         sync.Mutex
	elastigroup *aws.Group
}

func NewElastigroupResource(fieldsMap map[FieldName]*GenericField) *ElastigroupResource {
	return &ElastigroupResource{
		GenericResource: GenericResource{
			resourceName: ElastigroupAwsResourceName,
			fields: NewGenericFields(fieldsMap),
		},
	}
}

func (res *ElastigroupResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	egGroup := res.GetElastigroup()
	for _, field := range res.fields.fieldsMap {
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(egGroup, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *ElastigroupResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	var hasChanged = false
	egGroup := res.GetElastigroup()
	for _, field := range res.fields.fieldsMap {
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(egGroup, resourceData, meta); err != nil {
				return false, err
			}
			hasChanged = true
		}
	}

	return hasChanged, nil
}

func (res *ElastigroupResource) GetElastigroup() *aws.Group {
	if res.elastigroup == nil {
		res.mux.Lock()
		defer res.mux.Unlock()
		if res.elastigroup == nil {
			res.elastigroup = &aws.Group{
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
			}
		}
	}
	return res.elastigroup
}

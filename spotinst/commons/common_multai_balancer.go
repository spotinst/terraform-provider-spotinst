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
	MultaiBalancerResourceName ResourceName = "spotinst_multai_balancer"
)

var MultaiBalancerResource *MultaiBalancerTerraformResource

type MultaiBalancerTerraformResource struct {
	GenericResource // embedding
}

type MultaiBalancerWrapper struct {
	balancer *multai.LoadBalancer
}

func NewMultaiBalancerResource(fieldMap map[FieldName]*GenericField) *MultaiBalancerTerraformResource {
	return &MultaiBalancerTerraformResource{
		GenericResource: GenericResource{
			resourceName: MultaiBalancerResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

func (res *MultaiBalancerTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*multai.LoadBalancer, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	mlbWrapper := NewMultaiBalancerWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(mlbWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return mlbWrapper.GetMultaiBalancer(), nil
}

func (res *MultaiBalancerTerraformResource) OnRead(
	balancer *multai.LoadBalancer,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	mlbWrapper := NewMultaiBalancerWrapper()
	mlbWrapper.SetMultaiBalancer(balancer)

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

func (res *MultaiBalancerTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *multai.LoadBalancer, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	mlbWrapper := NewMultaiBalancerWrapper()
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

	return hasChanged, mlbWrapper.GetMultaiBalancer(), nil
}

func NewMultaiBalancerWrapper() *MultaiBalancerWrapper {
	return &MultaiBalancerWrapper{
		balancer: &multai.LoadBalancer{},
	}
}

func (mlbWrapper *MultaiBalancerWrapper) GetMultaiBalancer() *multai.LoadBalancer {
	return mlbWrapper.balancer
}

func (mlbWrapper *MultaiBalancerWrapper) SetMultaiBalancer(balancer *multai.LoadBalancer) {
	mlbWrapper.balancer = balancer
}

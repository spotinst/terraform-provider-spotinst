package commons

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"log"
)

const (
	StatefulNodeAzureV3ResourceName ResourceName = "spotinst_stateful_node_azure_v3"
)

var StatefulNodeAzureV3Resource *StatefulNodeAzureV3TerraformResource

type StatefulNodeAzureV3TerraformResource struct {
	GenericResource
}

type StatefulNodeAzureV3Wrapper struct {
	statefulNode *azurev3.StatefulNode
}

func NewStatefulNodeAzureV3Resource(fieldsMap map[FieldName]*GenericField) *StatefulNodeAzureV3TerraformResource {
	return &StatefulNodeAzureV3TerraformResource{
		GenericResource: GenericResource{
			resourceName: StatefulNodeAzureV3ResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *StatefulNodeAzureV3TerraformResource) OnRead(
	statefulNode *azurev3.StatefulNode,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	snWrapper := NewStatefulNodeAzureV3Wrapper()
	snWrapper.SetStatefulNode(statefulNode)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(snWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *StatefulNodeAzureV3TerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*azurev3.StatefulNode, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	snWrapper := NewStatefulNodeAzureV3Wrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(snWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return snWrapper.GetStatefulNode(), nil
}

func (res *StatefulNodeAzureV3TerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *azurev3.StatefulNode, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	snWrapper := NewStatefulNodeAzureV3Wrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(snWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, snWrapper.GetStatefulNode(), nil
}

func NewStatefulNodeAzureV3Wrapper() *StatefulNodeAzureV3Wrapper {
	return &StatefulNodeAzureV3Wrapper{
		statefulNode: &azurev3.StatefulNode{
			Strategy: &azurev3.Strategy{},
			Compute: &azurev3.Compute{
				LaunchSpecification: &azurev3.LaunchSpecification{},
				VMSizes:             &azurev3.VMSizes{},
			},
			Scheduling:  &azurev3.Scheduling{},
			Persistence: &azurev3.Persistence{},
			Health:      &azurev3.Health{},
		},
	}
}

func (snWrapper *StatefulNodeAzureV3Wrapper) GetStatefulNode() *azurev3.StatefulNode {
	return snWrapper.statefulNode
}

func (snWrapper *StatefulNodeAzureV3Wrapper) SetStatefulNode(statefulNode *azurev3.StatefulNode) {
	snWrapper.statefulNode = statefulNode
}

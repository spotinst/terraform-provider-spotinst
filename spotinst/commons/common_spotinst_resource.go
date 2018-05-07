package commons

import (
	"time"
	"math/rand"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"sync"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Init
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func init() {
	rand.Seed(time.Now().UnixNano())

	// Remove timestamp from provider logger, use the timestamp from the terraform logger
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
var ElastigroupResource *GenericResource


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//             Types
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
type hasFieldChange func(resourceData *schema.ResourceData, meta interface{}) bool
type onFieldRead func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error
type onFieldCreate func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error
type onFieldUpdate func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error

type TerraformData struct {
	ResourceData *schema.ResourceData
	Meta         interface{}
}

type GenericResource struct {
	// use interface{} to keep the generics between all Spotinst API resources
	elastigroup *aws.Group
	mux         sync.Mutex

	fields       *GenericFields
	resourceName string

	terraformData *TerraformData
}


type GenericField struct {
	resourceAffinity ResourceAffinity
	fieldNameStr     string
	fieldName        FieldName
	schema           *schema.Schema
	onRead           onFieldRead
	onCreate         onFieldCreate
	onUpdate         onFieldUpdate
	hasChangeCustom  hasFieldChange
}

type GenericFields struct {
	fieldsMap map[FieldName]*GenericField
	schemaMap map[string]*schema.Schema
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//          Constructors
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func NewGenericField(
	resourceAffinity ResourceAffinity,
	fieldName FieldName,
	schema *schema.Schema,
	onRead onFieldRead,
	onCreate onFieldCreate,
	onUpdate onFieldUpdate,
	hasChangeCustom hasFieldChange) *GenericField {

	return &GenericField{
		resourceAffinity: resourceAffinity,
		fieldNameStr:     string(fieldName),
		fieldName:        fieldName,
		schema:           schema,
		onRead:           onRead,
		onCreate:         onCreate,
		onUpdate:         onUpdate,
		hasChangeCustom:  hasChangeCustom,
	}
}

func NewGenericFields(fieldsMap map[FieldName]*GenericField) *GenericFields {
	var schemaMap = make(map[string]*schema.Schema)

	for _, field := range fieldsMap {
		schemaMap[field.fieldNameStr] = field.schema
	}

	return &GenericFields{
		fieldsMap: fieldsMap,
		schemaMap: schemaMap,
	}
}

func NewGenericResource(
	resourceName string,
	fieldsMap map[FieldName]*GenericField) *GenericResource {

	fields := NewGenericFields(fieldsMap)
	return &GenericResource{
		resourceName: resourceName,
		fields: fields,
	}
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//      Methods: GenericField
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func (field *GenericField) GetSchema() *schema.Schema {
	return field.schema
}

func (field *GenericField) hasFieldChange(resourceData *schema.ResourceData, meta interface{}) bool {
	if field.hasChangeCustom != nil {
		return field.hasChangeCustom(resourceData, meta)
	}
	return resourceData.HasChange(field.fieldNameStr)
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//   Methods: GenericResource
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func (res *GenericResource) OnRead(
	elastigroup *aws.Group) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	resourceData := res.terraformData.ResourceData
	meta := res.terraformData.Meta
	for _, field := range res.fields.fieldsMap {
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(elastigroup, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *GenericResource) OnCreate(
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

func (res *GenericResource) OnUpdate(
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

func (res *GenericResource) GetElastigroup() *aws.Group {
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
		//res.mux.Unlock()
	}
	return res.elastigroup
}

func (res *GenericResource) GetField(fieldName FieldName) *GenericField {
	if res.fields != nil && res.fields.fieldsMap != nil {
		return res.fields.fieldsMap[fieldName]
	}
	return nil
}

func (res *GenericResource) GetSchemaMap() map[string]*schema.Schema {
	if res.fields == nil || res.fields.schemaMap == nil || len(res.fields.schemaMap) == 0 {
		log.Printf("[ERROR] Resource schema is nil or empty")
		return nil
	}
	return res.fields.schemaMap
}

func (res *GenericResource) GetName() string {
	return res.resourceName
}

func (res *GenericResource) GetTerraformData() *TerraformData {
	return res.terraformData
}

func (res *GenericResource) SetTerraformData(data *TerraformData) {
	res.terraformData = data
}
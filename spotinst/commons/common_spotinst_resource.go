package commons

import (
	"time"
	"math/rand"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Init
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func init() {
	rand.Seed(time.Now().UnixNano())
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
var ElastigroupRepo *GenericApiResource
var LaunchConfigurationRepo *GenericCachedResource

//var ResourcesCache, _ = cache.NewLRUCache()


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//             Types
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
type hasFieldChange func(resourceData *schema.ResourceData, meta interface{}) bool
type onFieldRead func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error
type onFieldCreate func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error
type onFieldUpdate func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error

type GenericApiResource struct {
	// use interface{} to keep the generics between all Spotinst API resources
	elastigroup *aws.Group
	fields *GenericFields
}

type GenericCachedResource struct {
	fields *GenericFields
}

type GenericField struct {
	fieldNameStr    string
	fieldName       FieldName
	schema          *schema.Schema
	onRead          onFieldRead
	onCreate        onFieldCreate
	onUpdate        onFieldUpdate
	hasChangeCustom hasFieldChange
	dirty           bool
}

type GenericFields struct {
	fieldsMap map[FieldName]*GenericField
	schemaMap map[string]*schema.Schema
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//          Constructors
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func NewGenericField(
	fieldName FieldName,
	schema *schema.Schema,
	onRead onFieldRead,
	onCreate onFieldCreate,
	onUpdate onFieldUpdate,
	hasChangeCustom hasFieldChange) *GenericField {

	return &GenericField{
		fieldNameStr:    string(fieldName),
		fieldName:       fieldName,
		schema:          schema,
		onRead:          onRead,
		onCreate:        onCreate,
		onUpdate:        onUpdate,
		hasChangeCustom: hasChangeCustom,
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

func NewGenericApiResource(
	fieldsMap map[FieldName]*GenericField) *GenericApiResource {

	fields := NewGenericFields(fieldsMap)
	return &GenericApiResource{
		fields: fields,
	}
}

func NewGenericCachedResource(
	fieldsMap map[FieldName]*GenericField) *GenericCachedResource {

	fields := NewGenericFields(fieldsMap)
	return &GenericCachedResource{
		fields: fields,
	}
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//      Methods: GenericField
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func (field *GenericField) GetSchema() *schema.Schema {
	return field.schema
}

func (field *GenericField) Touch() {
	field.dirty = true
}

func (field *GenericField) IsDirty() bool {
	return field.dirty
}

func (field *GenericField) hasFieldChange(resourceData *schema.ResourceData, meta interface{}) bool {
	if field.hasChangeCustom != nil {
		return field.hasChangeCustom(resourceData, meta)
	}
	return resourceData.HasChange(field.fieldNameStr)
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//   Methods: GenericApiResource
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func (res *GenericApiResource) OnRead(
	elastigroup *aws.Group,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("api resource fields are nil or empty, cannot read")
	}

	for _, field := range res.fields.fieldsMap {
		if err := field.onRead(elastigroup, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *GenericApiResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("api resource fields are nil or empty, cannot create")
	}

	for _, field := range res.fields.fieldsMap {
		if field.hasFieldChange(resourceData, meta) {
			if err := field.onCreate(res.elastigroup, resourceData, meta); err != nil {
				return err
			}
		}
	}
	return nil
}

func (res *GenericApiResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, fmt.Errorf("api resource fields are nil or empty, cannot update")
	}

	var hasChanged = false
	for _, field := range res.fields.fieldsMap {
		if field.hasFieldChange(resourceData, meta) {
			if err := field.onUpdate(res.elastigroup, resourceData, meta); err != nil {
				return false, err
			}
			hasChanged = true
		}
	}

	return hasChanged, nil
}

func (res *GenericApiResource) GetElastigroup() *aws.Group {
	if res.elastigroup == nil {
		res.elastigroup = &aws.Group{
			Scaling:     &aws.Scaling{},
			Scheduling:  &aws.Scheduling{},
			Integration: &aws.Integration{},
			Compute: &aws.Compute{
				LaunchSpecification: &aws.LaunchSpecification{},
			},
		}
	}
	return res.elastigroup
}

func (res *GenericApiResource) GetField(fieldName FieldName) *GenericField {
	if res.fields != nil && res.fields.fieldsMap != nil {
		return res.fields.fieldsMap[fieldName]
	}
	return nil
}

func (res *GenericApiResource) GetSchemaMap() map[string]*schema.Schema {
	if res.fields == nil || res.fields.schemaMap == nil || len(res.fields.schemaMap) == 0 {
		log.Printf("[ERROR] Resource schema is nil or empty")
		return nil
	}
	return res.fields.schemaMap
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
// Methods: GenericCachedResource
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func (res *GenericCachedResource) OnRead(
	elastigroup *aws.Group,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("cached resource fields are nil or empty, cannot read")
	}

	for _, field := range res.fields.fieldsMap {
		if err := field.onRead(elastigroup, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *GenericCachedResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("cached resource fields are nil or empty, cannot create")
	}

	egGroup := ElastigroupRepo.GetElastigroup()
	for _, field := range res.fields.fieldsMap {
		if field.hasFieldChange(resourceData, meta) {
			if err := field.onCreate(egGroup, resourceData, meta); err != nil {
				return err
			}
		}
	}
	return nil
}

func (res *GenericCachedResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, fmt.Errorf("cached resource fields are nil or empty, cannot update")
	}

	var hasChanged = false
	egGroup := ElastigroupRepo.GetElastigroup()
	for _, field := range res.fields.fieldsMap {
		if field.hasFieldChange(resourceData, meta) {
			if err := field.onUpdate(egGroup, resourceData, meta); err != nil {
				return false, err
			}
			hasChanged = true
		}
	}
	return hasChanged, nil
}

func (res *GenericCachedResource) GetSchemaMap() map[string]*schema.Schema {
	if res.fields == nil || res.fields.schemaMap == nil || len(res.fields.schemaMap) == 0 {
		log.Printf("[ERROR] Resource schema is nil or empty")
		return nil
	}
	return res.fields.schemaMap
}
package commons

import (
	"log"
	"math/rand"
	"time"

	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
//             Types
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
type hasFieldChange func(resourceData *schema.ResourceData, meta interface{}) bool
type onFieldRead func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error
type onFieldCreate func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error
type onFieldUpdate func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error

type GenericResource struct {
	fields       *GenericFields
	resourceName ResourceName
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
	return string(res.resourceName)
}

func ToJson(object interface{}) (string, error) {
	if bytes, err := json.MarshalIndent(object, "", "  "); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

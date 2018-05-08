package commons

import (
	"time"
	"math/rand"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
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

type TerraformData struct {
	ResourceData *schema.ResourceData
	Meta         interface{}
}

type GenericResource struct {
	fields       *GenericFields
	resourceName ResourceName

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
	resourceObject interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	resourceData := res.terraformData.ResourceData
	meta := res.terraformData.Meta
	for _, field := range res.fields.fieldsMap {
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(resourceObject, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
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
	return string(res.resourceName)
}

func (res *GenericResource) GetTerraformData() *TerraformData {
	return res.terraformData
}

func (res *GenericResource) SetTerraformData(data *TerraformData) {
	res.terraformData = data
}
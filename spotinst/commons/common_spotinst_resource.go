package commons

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// GenerateSecureRandomInt generates a secure random integer between min and max using crypto/rand
func GenerateSecureRandomInt(min, max int64) (int64, error) {
	if min >= max {
		return 0, fmt.Errorf("invalid range: min must be less than max")
	}

	// Calculate the range
	rangeSize := max - min + 1

	// Generate a secure random number in the range 0 to rangeSize-1
	nBig, err := rand.Int(rand.Reader, big.NewInt(rangeSize))
	if err != nil {
		return 0, err
	}

	// Add the min to shift the range to min...max
	return nBig.Int64() + min, nil
}

func init() {
	// Example usage of GenerateSecureRandomInt within init
	randomNumber, err := GenerateSecureRandomInt(1, 100)
	if err != nil {
		log.Fatalf("Failed to generate secure random number: %v", err)
	}
	log.Printf("Secure random number: %d", randomNumber)

	// Remove timestamp from provider logger, use the timestamp from the Terraform logger.
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

type (
	hasFieldChange func(resourceData *schema.ResourceData, meta interface{}) bool
	onFieldRead    func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error
	onFieldCreate  func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error
	onFieldUpdate  func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error
)

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

func NewGenericField(
	resourceAffinity ResourceAffinity,
	fieldName FieldName,
	schema *schema.Schema,
	onRead onFieldRead,
	onCreate onFieldCreate,
	onUpdate onFieldUpdate,
	hasChangeCustom hasFieldChange,
) *GenericField {

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

func (field *GenericField) GetSchema() *schema.Schema {
	return field.schema
}

func (field *GenericField) hasFieldChange(resourceData *schema.ResourceData, meta interface{}) bool {
	if field.hasChangeCustom != nil {
		return field.hasChangeCustom(resourceData, meta)
	}
	return resourceData.HasChange(field.fieldNameStr)
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

func ToJson(object interface{}) (string, error) {
	if bytes, err := json.MarshalIndent(object, "", "  "); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}

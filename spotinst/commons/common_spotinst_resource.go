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
var ElastigroupResource *GenericApiResource
var ElastigroupLaunchConfigurationResource *GenericCachedResource
var ElastigroupStrategyResource *GenericCachedResource
var ElastigroupInstanceTypesResource *GenericCachedResource


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

type GenericApiResource struct {
	// use interface{} to keep the generics between all Spotinst API resources
	elastigroup *aws.Group
	mux         sync.Mutex

	fields       *GenericFields
	resourceName string

	terraformData *TerraformData
}

type GenericCachedResource struct {
	fields *GenericFields
	resourceName string

	terraformData *TerraformData
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
	resourceName string,
	fieldsMap map[FieldName]*GenericField) *GenericApiResource {

	fields := NewGenericFields(fieldsMap)
	return &GenericApiResource{
		resourceName: resourceName,
		fields: fields,
	}
}

func NewGenericCachedResource(
	resourceName string,
	fieldsMap map[FieldName]*GenericField) *GenericCachedResource {

	fields := NewGenericFields(fieldsMap)
	return &GenericCachedResource{
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
// Read method relies on the parent resource to fetch the group and then trigger onRead()
// on all other cached resources
func (res *GenericApiResource) OnRead(
	elastigroup *aws.Group) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("api resource fields are nil or empty, cannot read")
	}

	resourceData := res.terraformData.ResourceData
	meta := res.terraformData.Meta
	for _, field := range res.fields.fieldsMap {
		log.Printf(string(ResourceFieldOnRead), res.resourceName, field.fieldNameStr)
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

	egGroup := res.GetElastigroup()
	for _, field := range res.fields.fieldsMap {
		log.Printf(string(ResourceFieldOnCreate), res.resourceName, field.fieldNameStr)
		if err := field.onCreate(egGroup, resourceData, meta); err != nil {
			return err
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
	egGroup := res.GetElastigroup()
	for _, field := range res.fields.fieldsMap {
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), res.resourceName, field.fieldNameStr)
			if err := field.onUpdate(egGroup, resourceData, meta); err != nil {
				return false, err
			}
			hasChanged = true
		}
	}

	return hasChanged, nil
}

func (res *GenericApiResource) GetElastigroup() *aws.Group {
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
				Strategy: &aws.Strategy{},
			}
		}
		//res.mux.Unlock()
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

func (res *GenericApiResource) GetName() string {
	return res.resourceName
}

func (res *GenericApiResource) GetTerraformData() *TerraformData {
	return res.terraformData
}

func (res *GenericApiResource) SetTerraformData(data *TerraformData) {
	res.terraformData = data
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
// Methods: GenericCachedResource
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func (res *GenericCachedResource) OnRead(
	elastigroup *aws.Group) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("cached resource fields are nil or empty, cannot read")
	}

	resourceData := res.terraformData.ResourceData
	meta := res.terraformData.Meta
	for _, field := range res.fields.fieldsMap {
		log.Printf(string(ResourceFieldOnRead), res.resourceName, field.fieldNameStr)
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

	egGroup := ElastigroupResource.GetElastigroup()
	for _, field := range res.fields.fieldsMap {
		log.Printf(string(ResourceFieldOnCreate), res.resourceName, field.fieldNameStr)
		if err := field.onCreate(egGroup, resourceData, meta); err != nil {
			return err
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
	egGroup := ElastigroupResource.GetElastigroup()
	for _, field := range res.fields.fieldsMap {
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), res.resourceName, field.fieldNameStr)
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

func (res *GenericCachedResource) GetName() string {
	return res.resourceName
}

func (res *GenericCachedResource) GetTerraformData() *TerraformData {
	return res.terraformData
}

func (res *GenericCachedResource) SetTerraformData(data *TerraformData) {
	res.terraformData = data
}
package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/subscription"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	SubscriptionResourceName ResourceName = "spotinst_subscription"
)

var SubscriptionResource *SubscriptionTerraformResource

type SubscriptionTerraformResource struct {
	GenericResource // embedding
}

func NewSubscriptionResource(
	fieldsMap map[FieldName]*GenericField) *SubscriptionTerraformResource {

	return &SubscriptionTerraformResource{
		GenericResource: GenericResource{
			resourceName: SubscriptionResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *SubscriptionTerraformResource) OnRead(
	subscription *subscription.Subscription,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(subscription, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *SubscriptionTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*subscription.Subscription, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	sub := NewSubscription()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(sub, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return sub, nil
}

func (res *SubscriptionTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *subscription.Subscription, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	sub := NewSubscription()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(sub, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, sub, nil
}

func NewSubscription() *subscription.Subscription {
	return &subscription.Subscription{}
}

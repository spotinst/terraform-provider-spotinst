package commons

import (
	"sync"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/subscription"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	SubscriptionResourceName ResourceName = "spotinst_subscription"
)

var SpotinstSubscription *SubscriptionResource

type SubscriptionResource struct {
	GenericResource // embedding

	mux          sync.Mutex
	subscription *subscription.Subscription
}

func NewSubscriptionResource(
	fieldsMap map[FieldName]*GenericField) *SubscriptionResource {

	return &SubscriptionResource{
		GenericResource: GenericResource{
			resourceName: SubscriptionResourceName,
			fields: NewGenericFields(fieldsMap),
		},
	}
}

func (res *SubscriptionResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	egGroup := res.GetSubscription()
	for _, field := range res.fields.fieldsMap {
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(egGroup, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *SubscriptionResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	var hasChanged = false
	egGroup := res.GetSubscription()
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

func (res *SubscriptionResource) GetSubscription() *subscription.Subscription {
	if res.subscription == nil {
		res.mux.Lock()
		defer res.mux.Unlock()
		if res.subscription == nil {
			res.subscription = &subscription.Subscription{}
		}
	}
	return res.subscription
}

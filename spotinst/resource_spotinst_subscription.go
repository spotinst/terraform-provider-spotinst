package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/subscription"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

	subscriptionPackage "github.com/terraform-providers/terraform-provider-spotinst/spotinst/subscription"
)

func resourceSpotinstSubscription() *schema.Resource {
	setupSubscription()
	return &schema.Resource{
		Create: resourceSpotinstSubscriptionCreate,
		Update: resourceSpotinstSubscriptionUpdate,
		Read:   resourceSpotinstSubscriptionRead,
		Delete: resourceSpotinstSubscriptionDelete,

		Schema: commons.SubscriptionResource.GetSchemaMap(),
	}
}

func setupSubscription() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	subscriptionPackage.Setup(fieldsMap)

	commons.SubscriptionResource = commons.NewSubscriptionResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstSubscriptionDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.SubscriptionResource.GetName(), id)

	input := &subscription.DeleteSubscriptionInput{SubscriptionID: spotinst.String(id)}
	if _, err := meta.(*Client).subscription.Delete(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to delete subscription: %s", err)
	}

	resourceData.SetId("")
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstSubscriptionRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.SubscriptionResource.GetName(), id)

	client := meta.(*Client)
	input := &subscription.ReadSubscriptionInput{SubscriptionID: spotinst.String(resourceData.Id())}
	subResponse, err := client.subscription.Read(context.Background(), input)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to read subscription: %s", err)
	}

	// If nothing was found, then return no state.
	sub := subResponse.Subscription
	if sub == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.SubscriptionResource.OnRead(sub, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> Subscription read successfully: %s <===", id)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstSubscriptionCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.SubscriptionResource.GetName())

	sub, err := commons.SubscriptionResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	subscriptionId, err := createSubscription(sub, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(subscriptionId))
	log.Printf("===> Subscription created successfully: %s <===", resourceData.Id())

	return resourceSpotinstSubscriptionRead(resourceData, meta)
}

func createSubscription(subObj *subscription.Subscription, spotinstClient *Client) (*string, error) {
	input := &subscription.CreateSubscriptionInput{Subscription: subObj}
	resp, err := spotinstClient.subscription.Create(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create subscription: %s", err)
	}
	return resp.Subscription.ID, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstSubscriptionUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.SubscriptionResource.GetName(), id)

	shouldUpdate, sub, err := commons.SubscriptionResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		sub.SetId(spotinst.String(id))
		if err := updateSubscription(sub, resourceData, meta); err != nil {
			return err
		}
	}

	log.Printf("===> Subscription updated successfully: %s <===", id)
	return resourceSpotinstSubscriptionRead(resourceData, meta)
}

func updateSubscription(sub *subscription.Subscription, resourceData *schema.ResourceData, meta interface{}) error {
	input := &subscription.UpdateSubscriptionInput{
		Subscription: sub,
	}

	if json, err := commons.ToJson(sub); err != nil {
		return err
	} else {
		log.Printf("===> Subscrption update configuration: %s", json)
	}

	if _, err := meta.(*Client).subscription.Update(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] failed to update subscription %s: %s", resourceData.Id(), err)
	}
	return nil
}

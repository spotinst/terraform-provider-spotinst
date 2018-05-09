package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/subscription"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
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

		Schema: commons.SpotinstSubscription.GetSchemaMap(),
	}
}

func setupSubscription() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	subscriptionPackage.Setup(fieldsMap)

	commons.SpotinstSubscription = commons.NewSubscriptionResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstSubscriptionDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.SpotinstSubscription.GetName(), id)

	input := &subscription.DeleteSubscriptionInput{ SubscriptionID: spotinst.String(id)}
	if _, err := meta.(*Client).subscription.Delete(context.Background(), input); err != nil {
		return fmt.Errorf("failed to delete subscription: %s", err)
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
		commons.SpotinstSubscription.GetName(), id)

	client := meta.(*Client)
	input := &subscription.ReadSubscriptionInput{SubscriptionID: spotinst.String(resourceData.Id())}
	subResponse, err := client.subscription.Read(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to read subscription: %s", err)
	}

	// If nothing was found, then return no state.
	if subResponse.Subscription == nil {
		resourceData.SetId("")
		return nil
	}

	commons.SpotinstSubscription.SetTerraformData(
		&commons.TerraformData{
			ResourceData: resourceData,
			Meta:         meta,
		})

	commons.SpotinstSubscription.OnRead(subResponse)
	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstSubscriptionCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.SpotinstSubscription.GetName())

	err := commons.SpotinstSubscription.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	subObj := commons.SpotinstSubscription.GetSubscription()
	subscriptionId, err := createSubscription(subObj, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(subscriptionId))
	log.Printf("Subscription created successfully: %v", resourceData.Id())

	return resourceSpotinstSubscriptionRead(resourceData, meta)
}

func createSubscription(subObj *subscription.Subscription, spotinstClient *Client) (*string, error) {
	input := &subscription.CreateSubscriptionInput{Subscription: subObj}
	resp, err := spotinstClient.subscription.Create(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to create subscription: %s", err)
	}
	return resp.Subscription.ID, nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstSubscriptionUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.SpotinstSubscription.GetName(), id)

	shouldUpdate, err := commons.SpotinstSubscription.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		subObj := commons.SpotinstSubscription.GetSubscription()
		subObj.SetId(spotinst.String(id))
		updateSubscription(subObj, resourceData, meta)
	}

	return resourceSpotinstSubscriptionRead(resourceData, meta)
}

func updateSubscription(subObj *subscription.Subscription, resourceData *schema.ResourceData, meta interface{}) error {
	var input *subscription.UpdateSubscriptionInput

	input = &subscription.UpdateSubscriptionInput{
		Subscription: subObj,
	}

	log.Printf("Subscrption update configuration: %s", stringutil.Stringify(subObj))

	if _, err := meta.(*Client).subscription.Update(context.Background(), input); err != nil {
		return fmt.Errorf("failed to update subscription %s: %s", resourceData.Id(), err)
	}
	return nil
}
package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/notificationcenter"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/notification_center"
	"log"
)

func resourceSpotinstNotificationCenter() *schema.Resource {
	setupNotificationCenter()
	return &schema.Resource{
		CreateContext: resourceSpotinstNotificationCenterCreate,
		UpdateContext: resourceSpotinstNotificationCenterUpdate,
		ReadContext:   resourceSpotinstNotificationCenterRead,
		DeleteContext: resourceSpotinstNotificationCenterDelete,

		Schema: commons.NotificationCenterResource.GetSchemaMap(),
	}
}

func setupNotificationCenter() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	notification_center.Setup(fieldsMap)
	commons.NotificationCenterResource = commons.NewNotificationCenterResource(fieldsMap)
}

func resourceSpotinstNotificationCenterRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.NotificationCenterResource.GetName(), id)

	client := meta.(*Client)
	input := &notificationcenter.ReadNotificationCenterPolicyInput{PolicyId: spotinst.String(resourceData.Id())}
	ncResponse, err := client.notificationCenter.ReadNotificationCenterPolicy(context.Background(), input)
	if err != nil {
		return diag.Errorf("[ERROR] Failed to read notification center policy: %s", err)
	}

	nc := ncResponse.NotificationCenter
	if nc == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.NotificationCenterResource.OnRead(nc, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> Notification Center Policy read successfully: %s <===", id)
	return nil
}

func resourceSpotinstNotificationCenterCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.NotificationCenterResource.GetName())

	nc, err := commons.NotificationCenterResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	policyId, err := createNotificationCenter(nc, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}
	resourceData.SetId(spotinst.StringValue(policyId))
	log.Printf("===> Notification Center Policy created successfully: %s <===", resourceData.Id())

	return resourceSpotinstNotificationCenterRead(ctx, resourceData, meta)
}

func createNotificationCenter(nc *notificationcenter.NotificationCenter, client *Client) (*string, error) {
	input := nc
	resp, err := client.notificationCenter.CreateNotificationCenterPolicy(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Failed to create notification center policy: %s", err)
	}
	return resp.NotificationCenter.ID, nil
}

func resourceSpotinstNotificationCenterUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.NotificationCenterResource.GetName(), id)

	shouldUpdate, nc, err := commons.NotificationCenterResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		nc.SetID(spotinst.String(id))
		if err := updateNotificationCenter(nc, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}
	log.Printf("===> Notification Center Policy updated successfully: %s <===", id)
	return resourceSpotinstNotificationCenterRead(ctx, resourceData, meta)
}

func updateNotificationCenter(nc *notificationcenter.NotificationCenter, resourceData *schema.ResourceData, meta interface{}) error {
	input := &notificationcenter.UpdateNotificationCenterPolicyInput{
		NotificationCenter: nc,
	}

	if json, err := commons.ToJson(nc); err != nil {
		return err
	} else {
		log.Printf("===> Notification Center Policy update configuration: %s", json)
	}

	if _, err := meta.(*Client).notificationCenter.UpdateNotificationCenterPolicy(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update notification center policy %s: %s", resourceData.Id(), err)
	}
	return nil
}

func resourceSpotinstNotificationCenterDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.NotificationCenterResource.GetName(), id)

	input := &notificationcenter.DeleteNotificationCenterPolicyInput{PolicyId: spotinst.String(id)}
	if _, err := meta.(*Client).notificationCenter.DeleteNotificationCenterPolicy(context.Background(), input); err != nil {
		return diag.Errorf("[ERROR] Failed to delete notification center policy: %s", err)
	}
	resourceData.SetId("")
	return nil
}

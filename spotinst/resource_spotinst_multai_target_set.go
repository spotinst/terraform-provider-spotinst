package spotinst

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/multai_target_set"
)

func resourceSpotinstMultaiTargetSet() *schema.Resource {
	setupMultaiTargetSetResource()

	return &schema.Resource{
		Create: resourceSpotinstMultaiTargetSetCreate,
		Read:   resourceSpotinstMultaiTargetSetRead,
		Update: resourceSpotinstMultaiTargetSetUpdate,
		Delete: resourceSpotinstMultaiTargetSetDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.MultaiTargetSetResource.GetSchemaMap(),
	}
}

func setupMultaiTargetSetResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	multai_target_set.Setup(fieldsMap)

	commons.MultaiTargetSetResource = commons.NewMultaiTargetSetResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiTargetSetCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.MultaiTargetSetResource.GetName())

	targetSet, err := commons.MultaiTargetSetResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	targetSetId, err := createTargetSet(targetSet, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(targetSetId))
	log.Printf("===> Target Set created successfully: %s <===", resourceData.Id())

	return resourceSpotinstMultaiTargetSetRead(resourceData, meta)
}

func createTargetSet(targetSet *multai.TargetSet, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(targetSet); err != nil {
		return nil, err
	} else {
		log.Printf("===> Target Set create configuration: %s", json)
	}

	input := &multai.CreateTargetSetInput{TargetSet: targetSet}

	var resp *multai.CreateTargetSetOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.multai.CreateTargetSet(context.Background(), input)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create target set: %s", err)
	}

	return resp.TargetSet.ID, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiTargetSetRead(resourceData *schema.ResourceData, meta interface{}) error {
	targetSetId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.MultaiTargetSetResource.GetName(), targetSetId)

	input := &multai.ReadTargetSetInput{TargetSetID: spotinst.String(targetSetId)}
	resp, err := meta.(*Client).multai.ReadTargetSet(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to read target set: %s", err)
	}

	// If nothing was found, return no state
	targetSetResponse := resp.TargetSet
	if targetSetResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.MultaiTargetSetResource.OnRead(targetSetResponse, resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Target Set read successfully: %s <===", targetSetId)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiTargetSetUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	targetSetId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.MultaiTargetSetResource.GetName(), targetSetId)

	shouldUpdate, targetSet, err := commons.MultaiTargetSetResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		targetSet.SetId(spotinst.String(targetSetId))
		if err := updateTargetSet(targetSet, resourceData, meta); err != nil {
			return err
		}
	}

	log.Printf("===> Target Set updated successfully: %s <===", targetSetId)
	return resourceSpotinstMultaiTargetSetRead(resourceData, meta)
}

func updateTargetSet(targetSet *multai.TargetSet, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &multai.UpdateTargetSetInput{TargetSet: targetSet}
	targetSetId := resourceData.Id()

	if json, err := commons.ToJson(targetSet); err != nil {
		return err
	} else {
		log.Printf("===> Target Set update configuration: %s", json)
	}

	if _, err := meta.(*Client).multai.UpdateTargetSet(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update target set [%v]: %v", targetSetId, err)
	}

	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiTargetSetDelete(resourceData *schema.ResourceData, meta interface{}) error {
	targetSetId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.MultaiTargetSetResource.GetName(), targetSetId)

	if err := deleteTargetSet(resourceData, meta); err != nil {
		return err
	}

	err := awaitTargetSetDeleted(spotinst.String(targetSetId), meta.(*Client))
	if err != nil {
		return fmt.Errorf("[ERROR] Timed out when waiting for the target set to delete. error: %v", err)
	}

	log.Printf("===> Target Set deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteTargetSet(resourceData *schema.ResourceData, meta interface{}) error {
	targetSetId := resourceData.Id()
	input := &multai.DeleteTargetSetInput{TargetSetID: spotinst.String(targetSetId)}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Target Set delete configuration: %s", json)
	}

	if _, err := meta.(*Client).multai.DeleteTargetSet(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete target set: %s", err)
	}
	return nil
}

func awaitTargetSetDeleted(targetSetId *string, client *Client) error {
	input := &multai.ReadTargetSetInput{TargetSetID: spotinst.String(*targetSetId)}
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		resp, err := client.multai.ReadTargetSet(context.Background(), input)
		if err == nil && resp != nil && resp.TargetSet != nil {
			return resource.RetryableError(fmt.Errorf("===> waiting for target set to delete <==="))
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

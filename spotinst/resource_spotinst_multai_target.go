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
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/multai_target"
)

func resourceSpotinstMultaiTarget() *schema.Resource {
	setupMultaiTargetResource()

	return &schema.Resource{
		Create: resourceSpotinstMultaiTargetCreate,
		Read:   resourceSpotinstMultaiTargetRead,
		Update: resourceSpotinstMultaiTargetUpdate,
		Delete: resourceSpotinstMultaiTargetDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.MultaiTargetResource.GetSchemaMap(),
	}
}

func setupMultaiTargetResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	multai_target.Setup(fieldsMap)

	commons.MultaiTargetResource = commons.NewMultaiTargetResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiTargetCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.MultaiTargetResource.GetName())

	target, err := commons.MultaiTargetResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	targetId, err := createTarget(target, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(targetId))
	log.Printf("===> Target created successfully: %s <===", resourceData.Id())

	return resourceSpotinstMultaiTargetRead(resourceData, meta)
}

func createTarget(target *multai.Target, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(target); err != nil {
		return nil, err
	} else {
		log.Printf("===> Target create configuration: %s", json)
	}

	input := &multai.CreateTargetInput{Target: target}

	var resp *multai.CreateTargetOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.multai.CreateTarget(context.Background(), input)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create target: %s", err)
	}

	return resp.Target.ID, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiTargetRead(resourceData *schema.ResourceData, meta interface{}) error {
	targetId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.MultaiTargetResource.GetName(), targetId)

	input := &multai.ReadTargetInput{TargetID: spotinst.String(targetId)}
	resp, err := meta.(*Client).multai.ReadTarget(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to read target: %s", err)
	}

	// If nothing was found, return no state
	targetResponse := resp.Target
	if targetResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.MultaiTargetResource.OnRead(targetResponse, resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Target read successfully: %s <===", targetId)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiTargetUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	targetId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.MultaiTargetResource.GetName(), targetId)

	shouldUpdate, target, err := commons.MultaiTargetResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		target.SetId(spotinst.String(targetId))
		if err := updateTarget(target, resourceData, meta); err != nil {
			return err
		}
	}

	log.Printf("===> Target updated successfully: %s <===", targetId)
	return resourceSpotinstMultaiTargetRead(resourceData, meta)
}

func updateTarget(target *multai.Target, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &multai.UpdateTargetInput{Target: target}
	targetId := resourceData.Id()

	if json, err := commons.ToJson(target); err != nil {
		return err
	} else {
		log.Printf("===> Target update configuration: %s", json)
	}

	if _, err := meta.(*Client).multai.UpdateTarget(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update target [%v]: %v", targetId, err)
	}

	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiTargetDelete(resourceData *schema.ResourceData, meta interface{}) error {
	targetId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.MultaiTargetResource.GetName(), targetId)

	if err := deleteTarget(resourceData, meta); err != nil {
		return err
	}

	err := awaitTargetDeleted(spotinst.String(targetId), meta.(*Client))
	if err != nil {
		return fmt.Errorf("[ERROR] Timed out when waiting for the target to delete. error: %v", err)
	}

	log.Printf("===> Target deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteTarget(resourceData *schema.ResourceData, meta interface{}) error {
	targetId := resourceData.Id()
	input := &multai.DeleteTargetInput{TargetID: spotinst.String(targetId)}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Target delete configuration: %s", json)
	}

	if _, err := meta.(*Client).multai.DeleteTarget(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete target: %s", err)
	}
	return nil
}

func awaitTargetDeleted(targetId *string, client *Client) error {
	input := &multai.ReadTargetInput{TargetID: spotinst.String(*targetId)}
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		resp, err := client.multai.ReadTarget(context.Background(), input)
		if err == nil && resp != nil && resp.Target != nil {
			return resource.RetryableError(fmt.Errorf("===> waiting for target to delete <==="))
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

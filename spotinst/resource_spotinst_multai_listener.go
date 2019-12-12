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
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/multai_listener"
)

func resourceSpotinstMultaiListener() *schema.Resource {
	setupMultaiListenerResource()

	return &schema.Resource{
		Create: resourceSpotinstMultaiListenerCreate,
		Read:   resourceSpotinstMultaiListenerRead,
		Update: resourceSpotinstMultaiListenerUpdate,
		Delete: resourceSpotinstMultaiListenerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.MultaiListenerResource.GetSchemaMap(),
	}
}

func setupMultaiListenerResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	multai_listener.Setup(fieldsMap)

	commons.MultaiListenerResource = commons.NewMultaiListenerResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiListenerCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.MultaiListenerResource.GetName())

	listener, err := commons.MultaiListenerResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	listenerId, err := createListener(listener, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(listenerId))
	log.Printf("===> Listener created successfully: %s <===", resourceData.Id())

	return resourceSpotinstMultaiListenerRead(resourceData, meta)
}

func createListener(listener *multai.Listener, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(listener); err != nil {
		return nil, err
	} else {
		log.Printf("===> Listener create configuration: %s", json)
	}

	input := &multai.CreateListenerInput{Listener: listener}

	var resp *multai.CreateListenerOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.multai.CreateListener(context.Background(), input)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create listener: %s", err)
	}

	return resp.Listener.ID, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiListenerRead(resourceData *schema.ResourceData, meta interface{}) error {
	listenerId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.MultaiListenerResource.GetName(), listenerId)

	input := &multai.ReadListenerInput{ListenerID: spotinst.String(listenerId)}
	resp, err := meta.(*Client).multai.ReadListener(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to read listener: %s", err)
	}

	// If nothing was found, return no state
	listenerResponse := resp.Listener
	if listenerResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.MultaiListenerResource.OnRead(listenerResponse, resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Listener read successfully: %s <===", listenerId)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiListenerUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	listenerId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.MultaiListenerResource.GetName(), listenerId)

	shouldUpdate, listener, err := commons.MultaiListenerResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		listener.SetId(spotinst.String(listenerId))
		if err := updateListener(listener, resourceData, meta); err != nil {
			return err
		}
	}

	log.Printf("===> Listener updated successfully: %s <===", listenerId)
	return resourceSpotinstMultaiListenerRead(resourceData, meta)
}

func updateListener(listener *multai.Listener, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &multai.UpdateListenerInput{Listener: listener}
	listenerId := resourceData.Id()

	if json, err := commons.ToJson(listener); err != nil {
		return err
	} else {
		log.Printf("===> Listener update configuration: %s", json)
	}

	if _, err := meta.(*Client).multai.UpdateListener(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update listener [%v]: %v", listenerId, err)
	}

	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiListenerDelete(resourceData *schema.ResourceData, meta interface{}) error {
	listenerId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.MultaiListenerResource.GetName(), listenerId)

	if err := deleteListener(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Listener deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteListener(resourceData *schema.ResourceData, meta interface{}) error {
	listenerId := resourceData.Id()
	input := &multai.DeleteListenerInput{ListenerID: spotinst.String(listenerId)}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Listener delete configuration: %s", json)
	}

	if _, err := meta.(*Client).multai.DeleteListener(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete listener: %s", err)
	}
	return nil
}

package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/dataintegration/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/dataintegration"
	"log"
	"time"
)

func resourceSpotinstDataIntegration() *schema.Resource {
	setupDataIntegrationResource()

	return &schema.Resource{
		Create: resourceSpotinstDataIntegrationCreate,
		Update: resourceSpotinstDataIntegrationUpdate,
		Read:   resourceSpotinstDataIntegrationRead,
		Delete: resourceSpotinstDataIntegrationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.DataIntegrationResource.GetSchemaMap(),
	}
}

func setupDataIntegrationResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)
	dataintegration.Setup(fieldsMap)

	commons.DataIntegrationResource = commons.NewDataIntegrationResource(fieldsMap)
}

const ErrCodeDataIntegrationNotFound = "DATA_INTEGRATION_DOESNT_EXIST"

func resourceSpotinstDataIntegrationRead(resourceData *schema.ResourceData, meta interface{}) error {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.DataIntegrationResource.GetName(), resourceId)

	input := &aws.ReadDataIntegrationInput{DataIntegrationId: spotinst.String(resourceId)}
	resp, err := meta.(*Client).dataIntegration.CloudProviderAWS().ReadDataIntegration(context.Background(), input)
	if err != nil {
		// If the DataIntegration was not found, return nil so that we can show
		// that the DataIntegration does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeDataIntegrationNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}
		return fmt.Errorf("failed to read data integration: %s", err)
	}

	// If nothing was found, then return no state.
	DataIntegrationResponse := resp.DataIntegration
	if DataIntegrationResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.DataIntegrationResource.OnRead(DataIntegrationResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> DataIntegration read successfully: %s <===", resourceId)
	return nil
}

func resourceSpotinstDataIntegrationCreate(resourceData *schema.ResourceData, meta interface{}) error {

	log.Printf(string(commons.ResourceOnCreate), commons.DataIntegrationResource.GetName())

	DataIntegration, err := commons.DataIntegrationResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}
	DataIntegrationId, err := createDataIntegration(resourceData, DataIntegration, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(DataIntegrationId))

	log.Printf("===> DataIntegration created successfully: %s <===", resourceData.Id())

	return resourceSpotinstDataIntegrationRead(resourceData, meta)

}

func createDataIntegration(resourceData *schema.ResourceData, di *aws.DataIntegration, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(di); err != nil {
		return nil, err
	} else {
		log.Printf("===> DataIntegration create configuration: %s", json)
	}
	var resp *aws.CreateDataIntegrationOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		input := &aws.CreateDataIntegrationInput{DataIntegration: di}
		r, err := spotinstClient.dataIntegration.CloudProviderAWS().CreateDataIntegration(context.Background(), input)
		if err != nil {

			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create DataIntegration: %s", err)
	}
	return resp.DataIntegration.Id, nil

}

func resourceSpotinstDataIntegrationUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.DataIntegrationResource.GetName(), resourceId)

	shouldUpdate, di, err := commons.DataIntegrationResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		di.SetId(spotinst.String(resourceId))
		if err := updateDataIntegrationResource(di, resourceData, meta); err != nil {
			return err
		}
	}
	log.Printf("===> DataIntegration updated successfully: %s <===", resourceId)
	return resourceSpotinstDataIntegrationRead(resourceData, meta)
}

func updateDataIntegrationResource(di *aws.DataIntegration, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &aws.UpdateDataIntegrationInput{
		DataIntegration: di,
	}
	diId := resourceData.Id()

	if json, err := commons.ToJson(di); err != nil {
		return err
	} else {
		log.Printf("===> DataIntegration update configuration: %s", json)
	}

	if _, err := meta.(*Client).dataIntegration.CloudProviderAWS().UpdateDataIntegration(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update DataIntegration [%v]: %v", diId, err)
	}
	return nil
}

func resourceSpotinstDataIntegrationDelete(resourceData *schema.ResourceData, meta interface{}) error {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete), commons.DataIntegrationResource.GetName(), resourceId)

	if err := deleteDataIntegration(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> DataIntegration deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteDataIntegration(resourceData *schema.ResourceData, meta interface{}) error {
	diId := resourceData.Id()
	input := &aws.DeleteDataIntegrationInput{
		DataIntegrationId: spotinst.String(diId),
	}
	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> DataIntegration delete configuration: %s", json)
	}

	if _, err := meta.(*Client).dataIntegration.CloudProviderAWS().DeleteDataIntegration(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete DataIntegration: %s", err)
	}
	return nil
}

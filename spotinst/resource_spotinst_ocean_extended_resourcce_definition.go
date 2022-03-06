package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/extended_resource_definition"
)

func resourceSpotinstExtendedResourceDefinition() *schema.Resource {
	setupExtendedResourceDefinitionResource()

	return &schema.Resource{
		Create: resourceSpotinstExtendedResourceDefinitionCreate,
		Update: resourceSpotinstExtendedResourceDefinitionUpdate,
		Read:   resourceSpotinstExtendedResourceDefinitionRead,
		Delete: resourceSpotinstExtendedResourceDefinitionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.ExtendedResourceDefinitionResource.GetSchemaMap(),
	}
}

func setupExtendedResourceDefinitionResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)
	extended_resource_definition.Setup(fieldsMap)

	commons.ExtendedResourceDefinitionResource = commons.NewExtendedResourceDefinitionResource(fieldsMap)
}

const ErrCodeExtendedResourceDefinitionNotFound = "EXTENDED_RESOURCE_DEFINITION_DOESNT_EXIST"

func resourceSpotinstExtendedResourceDefinitionRead(resourceData *schema.ResourceData, meta interface{}) error {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.ExtendedResourceDefinitionResource.GetName(), resourceId)

	input := &aws.ReadExtendedResourceDefinitionInput{ExtendedResourceDefinitionID: spotinst.String(resourceId)}
	resp, err := meta.(*Client).ocean.CloudProviderAWS().ReadExtendedResourceDefinition(context.Background(), input)
	if err != nil {
		// If the ExtendedResourceDefinition was not found, return nil so that we can show
		// that the ExtendedResourceDefinition does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeExtendedResourceDefinitionNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}
		return fmt.Errorf("failed to read extended resource definition: %s", err)
	}

	// If nothing was found, then return no state.
	ExtendedResourceDefinitionResponse := resp.ExtendedResourceDefinition
	if ExtendedResourceDefinitionResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.ExtendedResourceDefinitionResource.OnRead(ExtendedResourceDefinitionResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> ExtendedResourceDefinition read successfully: %s <===", resourceId)
	return nil
}

func resourceSpotinstExtendedResourceDefinitionCreate(resourceData *schema.ResourceData, meta interface{}) error {

	log.Printf(string(commons.ResourceOnCreate), commons.ExtendedResourceDefinitionResource.GetName())

	extendedResourceDefinition, err := commons.ExtendedResourceDefinitionResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}
	extendedResourceDefinitionId, err := createExtendedResourceDefinition(resourceData, extendedResourceDefinition, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(extendedResourceDefinitionId))

	log.Printf("===> ExtendedResourceDefinition created successfully: %s <===", resourceData.Id())

	return resourceSpotinstExtendedResourceDefinitionRead(resourceData, meta)

}

func createExtendedResourceDefinition(resourceData *schema.ResourceData, erd *aws.ExtendedResourceDefinition, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(erd); err != nil {
		return nil, err
	} else {
		log.Printf("===> ExtendedResourceDefinition create configuration: %s", json)
	}
	var resp *aws.CreateExtendedResourceDefinitionOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		input := &aws.CreateExtendedResourceDefinitionInput{ExtendedResourceDefinition: erd}
		r, err := spotinstClient.ocean.CloudProviderAWS().CreateExtendedResourceDefinition(context.Background(), input)
		if err != nil {
			// Checks whether we should retry the ExtendedResourceDefinition creation.
			if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
				for _, err := range errs {
					if err.Code == "InvalidParameterValue" &&
						strings.Contains(err.Message, "Invalid IAM Instance Profile") {
						return resource.RetryableError(err)
					}
				}
			}
			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create ExtendedResourceDefinition: %s", err)
	}
	return resp.ExtendedResourceDefinition.ID, nil

}

func resourceSpotinstExtendedResourceDefinitionUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.ExtendedResourceDefinitionResource.GetName(), resourceId)

	shouldUpdate, erd, err := commons.ExtendedResourceDefinitionResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		erd.SetId(spotinst.String(resourceId))
		if err := updateExtendedResourceDefinition(erd, resourceData, meta); err != nil {
			return err
		}
	}
	log.Printf("===> ExtendedResourceDefinition updated successfully: %s <===", resourceId)
	return resourceSpotinstExtendedResourceDefinitionRead(resourceData, meta)
}

func updateExtendedResourceDefinition(erd *aws.ExtendedResourceDefinition, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &aws.UpdateExtendedResourceDefinitionInput{
		ExtendedResourceDefinition: erd,
	}
	erdId := resourceData.Id()

	if json, err := commons.ToJson(erd); err != nil {
		return err
	} else {
		log.Printf("===> ExtendedResourceDefinition update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().UpdateExtendedResourceDefinition(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update ExtendedResourceDefinition [%v]: %v", erdId, err)
	}
	return nil
}

func resourceSpotinstExtendedResourceDefinitionDelete(resourceData *schema.ResourceData, meta interface{}) error {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete), commons.ExtendedResourceDefinitionResource.GetName(), resourceId)

	if err := deleteExtendedResourceDefinition(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> ExtendedResourceDefinition deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteExtendedResourceDefinition(resourceData *schema.ResourceData, meta interface{}) error {
	erdId := resourceData.Id()
	input := &aws.DeleteExtendedResourceDefinitionInput{
		ExtendedResourceDefinitionID: spotinst.String(erdId),
	}
	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> ExtendedResourceDefinition delete configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().DeleteExtendedResourceDefinition(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete ExtendedResourceDefinition: %s", err)
	}
	return nil
}

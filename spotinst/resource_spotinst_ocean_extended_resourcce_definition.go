package spotinst

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aws_extended_resource_definition"
)

func resourceSpotinstOceanAWSExtendedResourceDefinition() *schema.Resource {
	setupOceanAWSExtendedResourceDefinitionResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstOceanAWSExtendedResourceDefinitionCreate,
		UpdateContext: resourceSpotinstOceanAWSExtendedResourceDefinitionUpdate,
		ReadContext:   resourceSpotinstOceanAWSExtendedResourceDefinitionRead,
		DeleteContext: resourceSpotinstOceanAWSExtendedResourceDefinitionDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.OceanAWSExtendedResourceDefinitionResource.GetSchemaMap(),
	}
}

func setupOceanAWSExtendedResourceDefinitionResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)
	ocean_aws_extended_resource_definition.Setup(fieldsMap)

	commons.OceanAWSExtendedResourceDefinitionResource = commons.NewOceanAWSExtendedResourceDefinitionResource(fieldsMap)
}

const ErrCodeExtendedResourceDefinitionNotFound = "EXTENDED_RESOURCE_DEFINITION_DOESNT_EXIST"

func resourceSpotinstOceanAWSExtendedResourceDefinitionRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanAWSExtendedResourceDefinitionResource.GetName(), resourceId)

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
		return diag.Errorf("failed to read extended resource definition: %s", err)
	}

	// If nothing was found, then return no state.
	ExtendedResourceDefinitionResponse := resp.ExtendedResourceDefinition
	if ExtendedResourceDefinitionResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanAWSExtendedResourceDefinitionResource.OnRead(ExtendedResourceDefinitionResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> ExtendedResourceDefinition read successfully: %s <===", resourceId)
	return nil
}

func resourceSpotinstOceanAWSExtendedResourceDefinitionCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {

	log.Printf(string(commons.ResourceOnCreate), commons.OceanAWSExtendedResourceDefinitionResource.GetName())

	extendedResourceDefinition, err := commons.OceanAWSExtendedResourceDefinitionResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	extendedResourceDefinitionId, err := createOceanAWSExtendedResourceDefinition(resourceData, extendedResourceDefinition, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(extendedResourceDefinitionId))

	log.Printf("===> ExtendedResourceDefinition created successfully: %s <===", resourceData.Id())

	return resourceSpotinstOceanAWSExtendedResourceDefinitionRead(ctx, resourceData, meta)

}

func createOceanAWSExtendedResourceDefinition(resourceData *schema.ResourceData, erd *aws.ExtendedResourceDefinition, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(erd); err != nil {
		return nil, err
	} else {
		log.Printf("===> ExtendedResourceDefinition create configuration: %s", json)
	}
	var resp *aws.CreateExtendedResourceDefinitionOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &aws.CreateExtendedResourceDefinitionInput{ExtendedResourceDefinition: erd}
		r, err := spotinstClient.ocean.CloudProviderAWS().CreateExtendedResourceDefinition(context.Background(), input)
		if err != nil {

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

func resourceSpotinstOceanAWSExtendedResourceDefinitionUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.OceanAWSExtendedResourceDefinitionResource.GetName(), resourceId)

	shouldUpdate, erd, err := commons.OceanAWSExtendedResourceDefinitionResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		erd.SetId(spotinst.String(resourceId))
		if err := updateOceanAWSExtendedResourceDefinition(erd, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}
	log.Printf("===> ExtendedResourceDefinition updated successfully: %s <===", resourceId)
	return resourceSpotinstOceanAWSExtendedResourceDefinitionRead(ctx, resourceData, meta)
}

func updateOceanAWSExtendedResourceDefinition(erd *aws.ExtendedResourceDefinition, resourceData *schema.ResourceData, meta interface{}) error {
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

func resourceSpotinstOceanAWSExtendedResourceDefinitionDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete), commons.OceanAWSExtendedResourceDefinitionResource.GetName(), resourceId)

	if err := deleteOceanAWSExtendedResourceDefinition(resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> ExtendedResourceDefinition deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteOceanAWSExtendedResourceDefinition(resourceData *schema.ResourceData, meta interface{}) error {
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

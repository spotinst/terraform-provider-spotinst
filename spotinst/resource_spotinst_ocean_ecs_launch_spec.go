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
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_ecs_launch_spec"
)

func resourceSpotinstOceanECSLaunchSpec() *schema.Resource {
	setupOceanECSLaunchSpecResource()

	return &schema.Resource{
		Create: resourceSpotinstOceanECSLaunchSpecCreate,
		Read:   resourceSpotinstOceanECSLaunchSpecRead,
		Update: resourceSpotinstOceanECSLaunchSpecUpdate,
		Delete: resourceSpotinstOceanECSLaunchSpecDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.OceanECSLaunchSpecResource.GetSchemaMap(),
	}
}

func setupOceanECSLaunchSpecResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)
	ocean_ecs_launch_spec.Setup(fieldsMap)

	commons.OceanECSLaunchSpecResource = commons.NewOceanECSLaunchSpecResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstOceanECSLaunchSpecCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate), commons.OceanECSLaunchSpecResource.GetName())

	launchSpec, err := commons.OceanECSLaunchSpecResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	launchSpecId, err := createECSLaunchSpec(launchSpec, meta.(*Client))
	if err != nil {
		return err
	}
	resourceData.SetId(spotinst.StringValue(launchSpecId))

	return resourceSpotinstOceanECSLaunchSpecRead(resourceData, meta)
}

func createECSLaunchSpec(launchSpec *aws.ECSLaunchSpec, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(launchSpec); err != nil {
		return nil, err
	} else {
		log.Printf("===> LaunchSpec create configuration: %s", json)
	}

	input := &aws.CreateECSLaunchSpecInput{LaunchSpec: launchSpec}

	var resp *aws.CreateECSLaunchSpecOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.ocean.CloudProviderAWS().CreateECSLaunchSpec(context.Background(), input)
		if err != nil {
			// Checks whether we should retry launchSpec creation.
			if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
				for _, err := range errs {
					if err.Code == "InvalidParameterValue" &&
						strings.Contains(err.Message, "Invalid IAM Instance Profile") {
						return resource.NonRetryableError(err)
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
		return nil, fmt.Errorf("[ERROR] failed to create launchSpec: %s", err)
	}
	return resp.LaunchSpec.ID, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const ErrCodeECSLaunchSpecNotFound = "CANT_GET_OCEAN_ECS_LAUNCH_SPEC"

func resourceSpotinstOceanECSLaunchSpecRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanECSLaunchSpecResource.GetName(), id)

	input := &aws.ReadECSLaunchSpecInput{LaunchSpecID: spotinst.String(id)}
	resp, err := meta.(*Client).ocean.CloudProviderAWS().ReadECSLaunchSpec(context.Background(), input)

	if err != nil {
		// If the launchSpec was not found, return nil so that we can show
		// that it does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeECSLaunchSpecNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return fmt.Errorf("failed to read launchSpec: %s", err)
	}

	// if nothing was found, return no state
	launchSpecResponse := resp.LaunchSpec
	if launchSpecResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanECSLaunchSpecResource.OnRead(launchSpecResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> launchSpec read successfully: %s <===", id)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstOceanECSLaunchSpecUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.OceanECSLaunchSpecResource.GetName(), id)
	shouldUpdate, launchSpec, err := commons.OceanECSLaunchSpecResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		launchSpec.SetId(spotinst.String(id))
		if err := updateECSLaunchSpec(launchSpec, resourceData, meta); err != nil {
			return err
		}
	}
	log.Printf("===> launchSpec updated successfully: %s <===", id)
	return resourceSpotinstOceanECSLaunchSpecRead(resourceData, meta)
}

func updateECSLaunchSpec(launchSpec *aws.ECSLaunchSpec, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &aws.UpdateECSLaunchSpecInput{
		LaunchSpec: launchSpec,
	}

	launchSpecId := resourceData.Id()

	if json, err := commons.ToJson(launchSpec); err != nil {
		return err
	} else {
		log.Printf("===> launchSpec update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().UpdateECSLaunchSpec(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update launchSpec [%v]: %v", launchSpecId, err)
	}

	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstOceanECSLaunchSpecDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanECSLaunchSpecResource.GetName(), id)

	if err := deleteECSLaunchSpec(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> launchSpec deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteECSLaunchSpec(resourceData *schema.ResourceData, meta interface{}) error {
	launchSpecId := resourceData.Id()
	input := &aws.DeleteECSLaunchSpecInput{
		LaunchSpecID: spotinst.String(launchSpecId),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> launchSpec delete configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().DeleteECSLaunchSpec(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete launchSpecId: %s", err)
	}
	return nil
}

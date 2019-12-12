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
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_aws_launch_spec"
)

func resourceSpotinstOceanAWSLaunchSpec() *schema.Resource {
	setupOceanAWSLaunchSpecResource()

	return &schema.Resource{
		Create: resourceSpotinstOceanAWSLaunchSpecCreate,
		Read:   resourceSpotinstOceanAWSLaunchSpecRead,
		Update: resourceSpotinstOceanAWSLaunchSpecUpdate,
		Delete: resourceSpotinstOceanAWSLaunchSpecDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.OceanAWSLaunchSpecResource.GetSchemaMap(),
	}
}

func setupOceanAWSLaunchSpecResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)
	ocean_aws_launch_spec.Setup(fieldsMap)

	commons.OceanAWSLaunchSpecResource = commons.NewOceanAWSLaunchSpecResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstOceanAWSLaunchSpecCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate), commons.OceanAWSLaunchSpecResource.GetName())

	launchSpec, err := commons.OceanAWSLaunchSpecResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	launchSpecId, err := createLaunchSpec(launchSpec, meta.(*Client))
	if err != nil {
		return err
	}
	resourceData.SetId(spotinst.StringValue(launchSpecId))

	return resourceSpotinstOceanAWSLaunchSpecRead(resourceData, meta)
}

func createLaunchSpec(launchSpec *aws.LaunchSpec, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(launchSpec); err != nil {
		return nil, err
	} else {
		log.Printf("===> LaunchSpec create configuration: %s", json)
	}

	input := &aws.CreateLaunchSpecInput{LaunchSpec: launchSpec}

	var resp *aws.CreateLaunchSpecOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.ocean.CloudProviderAWS().CreateLaunchSpec(context.Background(), input)
		if err != nil {
			// Checks whether we should retry launchSpec creation.
			if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
				for _, err := range errs {
					if err.Code == "InvalidParamterValue" &&
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
const ErrCodeLaunchSpecNotFound = "CANT_GET_OCEAN_LAUNCH_SPEC"

func resourceSpotinstOceanAWSLaunchSpecRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanAWSLaunchSpecResource.GetName(), id)

	input := &aws.ReadLaunchSpecInput{LaunchSpecID: spotinst.String(id)}
	resp, err := meta.(*Client).ocean.CloudProviderAWS().ReadLaunchSpec(context.Background(), input)

	if err != nil {
		// If the launchSpec was not found, return nil so that we can show
		// that it does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeLaunchSpecNotFound {
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

	if err := commons.OceanAWSLaunchSpecResource.OnRead(launchSpecResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> launchSpec read successfully: %s <===", id)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstOceanAWSLaunchSpecUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.OceanAWSLaunchSpecResource.GetName(), id)

	shouldUpdate, launchSpec, err := commons.OceanAWSLaunchSpecResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		launchSpec.SetId(spotinst.String(id))
		if err := updateLaunchSpec(launchSpec, resourceData, meta); err != nil {
			return err
		}
	}
	log.Printf("===> launchSpec updated successfully: %s <===", id)
	return resourceSpotinstOceanAWSLaunchSpecRead(resourceData, meta)
}

func updateLaunchSpec(launchSpec *aws.LaunchSpec, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &aws.UpdateLaunchSpecInput{
		LaunchSpec: launchSpec,
	}

	launchSpecId := resourceData.Id()

	if json, err := commons.ToJson(launchSpec); err != nil {
		return err
	} else {
		log.Printf("===> launchSpec update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().UpdateLaunchSpec(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update launchSpec [%v]: %v", launchSpecId, err)
	}

	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstOceanAWSLaunchSpecDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanAWSLaunchSpecResource.GetName(), id)

	if err := deleteLaunchSpec(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> launchSpec deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteLaunchSpec(resourceData *schema.ResourceData, meta interface{}) error {
	launchSpecId := resourceData.Id()
	input := &aws.DeleteLaunchSpecInput{
		LaunchSpecID: spotinst.String(launchSpecId),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> launchSpec delete configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().DeleteLaunchSpec(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete launchSpecId: %s", err)
	}
	return nil
}

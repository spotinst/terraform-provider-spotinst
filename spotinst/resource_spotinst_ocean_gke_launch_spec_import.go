package spotinst

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_gke_launch_spec_import"

	"log"

	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

func resourceSpotinstOceanGKELaunchSpecImport() *schema.Resource {
	setupOceanGKELaunchSpecImportResource()

	return &schema.Resource{
		Create: resourceSpotinstOceanGKELaunchSpecImportCreate,
		Read:   resourceSpotinstOceanGKELaunchSpecImportRead,
		Update: resourceSpotinstOceanGKELaunchSpecImportUpdate,
		Delete: resourceSpotinstOceanGKELaunchSpecImportDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.OceanGKELaunchSpecImportResource.GetSchemaMap(),
	}
}

func setupOceanGKELaunchSpecImportResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	ocean_gke_launch_spec_import.Setup(fieldsMap)

	commons.OceanGKELaunchSpecImportResource = commons.NewOceanGKELaunchSpecImportResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstOceanGKELaunchSpecImportCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate), commons.OceanGKELaunchSpecImportResource.GetName())

	importedLaunchSpec, err := importOceanGKELaunchSpec(resourceData, meta)

	if err != nil {
		return err
	}

	launchSpec, err := commons.OceanGKELaunchSpecImportResource.OnCreate(importedLaunchSpec, resourceData, meta.(*Client))

	if err != nil {
		return err
	}

	launchSpecId, err := createGKELaunchSpecImport(launchSpec, meta.(*Client))

	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(launchSpecId))

	return resourceSpotinstOceanGKELaunchSpecImportRead(resourceData, meta)
}

func createGKELaunchSpecImport(launchSpec *gcp.LaunchSpec, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(launchSpec); err != nil {
		return nil, err
	} else {
		log.Printf("===> LaunchSpec GKE create configuration: %s", json)
	}

	input := &gcp.CreateLaunchSpecInput{LaunchSpec: launchSpec}

	if out, err := spotinstClient.ocean.CloudProviderGCP().CreateLaunchSpec(context.Background(), input); err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create launchSpec: %s", err)
	} else {
		return out.LaunchSpec.ID, nil
	}
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstOceanGKELaunchSpecImportRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanGKELaunchSpecImportResource.GetName(), id)

	input := &gcp.ReadLaunchSpecInput{LaunchSpecID: spotinst.String(id)}
	resp, err := meta.(*Client).ocean.CloudProviderGCP().ReadLaunchSpec(context.Background(), input)

	if err != nil {
		// If the launchSpec was not found, return nil so that we can show
		// that it does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGKELaunchSpecNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return fmt.Errorf("failed to read GKE launchSpec: %s", err)
	}

	// if nothing was found, return no state
	launchSpecResponse := resp.LaunchSpec
	if launchSpecResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanGKELaunchSpecImportResource.OnRead(launchSpecResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> launchSpec GKE read successfully: %s <===", id)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstOceanGKELaunchSpecImportUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.OceanGKELaunchSpecImportResource.GetName(), id)

	shouldUpdate, launchSpec, err := commons.OceanGKELaunchSpecImportResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		launchSpec.SetId(spotinst.String(id))
		if err := updateGKELaunchSpecImport(launchSpec, resourceData, meta); err != nil {
			return err
		}
	}
	log.Printf("===> launchSpec GKE updated successfully: %s <===", id)
	return resourceSpotinstOceanGKELaunchSpecImportRead(resourceData, meta)
}

func updateGKELaunchSpecImport(launchSpec *gcp.LaunchSpec, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &gcp.UpdateLaunchSpecInput{
		LaunchSpec: launchSpec,
	}

	launchSpecId := resourceData.Id()

	if json, err := commons.ToJson(launchSpec); err != nil {
		return err
	} else {
		log.Printf("===> launchSpec GKE update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderGCP().UpdateLaunchSpec(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update launchSpec GKE [%v]: %v", launchSpecId, err)
	}

	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstOceanGKELaunchSpecImportDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanGKELaunchSpecImportResource.GetName(), id)

	if err := deleteGKELaunchSpecImport(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> launchSpec GKE deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteGKELaunchSpecImport(resourceData *schema.ResourceData, meta interface{}) error {
	launchSpecId := resourceData.Id()
	input := &gcp.DeleteLaunchSpecInput{
		LaunchSpecID: spotinst.String(launchSpecId),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> launchSpec GKE delete configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderGCP().DeleteLaunchSpec(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete launchSpecId: %s", err)
	}
	return nil
}

//region Import Ocean GKE Launch Spec
func importOceanGKELaunchSpec(resourceData *schema.ResourceData, meta interface{}) (*gcp.LaunchSpec, error) {
	input := &gcp.ImportOceanGKELaunchSpecInput{
		OceanId:      spotinst.String(resourceData.Get("ocean_id").(string)),
		NodePoolName: spotinst.String(resourceData.Get("node_pool_name").(string)),
	}

	resp, err := meta.(*Client).ocean.CloudProviderGCP().ImportOceanGKELaunchSpec(context.Background(), input)

	if err != nil {
		// If the group was not found, return nil so that we can show
		// that the group is gone.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGroupNotFound {
					resourceData.SetId("")
					return nil, err
				}
			}
		}
		// Some other error, report it.
		return nil, fmt.Errorf("ocean GKE: import failed to read group: %s", err)
	}

	return resp.LaunchSpec, err
}

//endregion

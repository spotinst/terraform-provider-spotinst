package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_azure"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_azure_health_check"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_azure_image"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_azure_launch_configuration"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_azure_load_balancer"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_azure_login"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_azure_network"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_azure_scheduled_task"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_azure_strategy"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_azure_vm_sizes"
	"log"
	"time"
)

func resourceSpotinstElastigroupAzure() *schema.Resource {
	setupElastigroupAzureResource()

	return &schema.Resource{
		Create: resourceSpotinstElastigroupAzureCreate,
		Read:   resourceSpotinstElastigroupAzureRead,
		Update: resourceSpotinstElastigroupAzureUpdate,
		Delete: resourceSpotinstElastigroupAzureDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.ElastigroupAzureResource.GetSchemaMap(),
	}
}

func setupElastigroupAzureResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	elastigroup_azure.Setup(fieldsMap)
	elastigroup_azure_health_check.Setup(fieldsMap)
	elastigroup_azure_image.Setup(fieldsMap)
	elastigroup_azure_launch_configuration.Setup(fieldsMap)
	elastigroup_azure_load_balancer.Setup(fieldsMap)
	elastigroup_azure_login.Setup(fieldsMap)
	elastigroup_azure_network.Setup(fieldsMap)
	elastigroup_azure_strategy.Setup(fieldsMap)
	elastigroup_azure_vm_sizes.Setup(fieldsMap)
	elastigroup_azure_scheduled_task.Setup(fieldsMap)

	commons.ElastigroupAzureResource = commons.NewElastigroupAzureResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstElastigroupAzureCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.ElastigroupAzureResource.GetName())

	elastigroup, err := commons.ElastigroupAzureResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	groupId, err := createAzureGroup(elastigroup, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(groupId))

	log.Printf("===> Elastigroup created successfully: %s <===", resourceData.Id())

	return resourceSpotinstElastigroupAzureRead(resourceData, meta)
}

func createAzureGroup(group *azure.Group, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(group); err != nil {
		return nil, err
	} else {
		log.Printf("===> Group create configuration: %s", json)
	}

	input := &azure.CreateGroupInput{Group: group}

	var resp *azure.CreateGroupOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.elastigroup.CloudProviderAzure().Create(context.Background(), input)
		if err != nil {
			log.Printf("error: %v", err)
			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create group: %s", err)
	}
	return resp.Group.ID, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstElastigroupAzureRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceFieldOnRead),
		commons.ElastigroupAzureResource.GetName(), id)

	input := &azure.ReadGroupInput{GroupID: spotinst.String(id)}
	resp, err := meta.(*Client).elastigroup.CloudProviderAzure().Read(context.Background(), input)
	if err != nil {
		// If the group was not found, return nil so that we can show
		// that the group does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGroupNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return fmt.Errorf("failed to read group: %s", err)
	}

	// If nothing was found, then return no state.
	groupResponse := resp.Group
	if groupResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.ElastigroupAzureResource.OnRead(groupResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> Elastigroup read successfully: %s <===", id)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstElastigroupAzureUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.ElastigroupAzureResource.GetName(), id)

	shouldUpdate, elastigroup, err := commons.ElastigroupAzureResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		elastigroup.SetId(spotinst.String(id))
		if err := updateAzureGroup(elastigroup, resourceData, meta); err != nil {
			return err
		}
	}

	log.Printf("===> Elastigroup updated successfully: %s <===", id)
	return resourceSpotinstElastigroupAzureRead(resourceData, meta)
}

func updateAzureGroup(elastigroup *azure.Group, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &azure.UpdateGroupInput{
		Group: elastigroup,
	}

	groupId := resourceData.Id()

	if json, err := commons.ToJson(elastigroup); err != nil {
		return err
	} else {
		log.Printf("===> Group update configuration: %s", json)
	}

	if _, err := meta.(*Client).elastigroup.CloudProviderAzure().Update(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update group [%v]: %v", groupId, err)
	}
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstElastigroupAzureDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.ElastigroupAzureResource.GetName(), id)

	if err := deleteAzureGroup(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Elastigroup deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteAzureGroup(resourceData *schema.ResourceData, meta interface{}) error {
	groupId := resourceData.Id()
	input := &azure.DeleteGroupInput{
		GroupID: spotinst.String(groupId),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Group delete configuration: %s", json)
	}

	if _, err := meta.(*Client).elastigroup.CloudProviderAzure().Delete(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete group: %s", err)
	}
	return nil
}

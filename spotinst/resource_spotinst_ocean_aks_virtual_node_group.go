package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_virtual_node_group"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_virtual_node_group_auto_scaling"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_virtual_node_group_launch_specification"
)

func resourceSpotinstOceanAKSVirtualNodeGroup() *schema.Resource {
	setupOceanAKSVirtualNodeGroupResource()

	return &schema.Resource{
		Create: resourceSpotinstOceanAKSVirtualNodeGroupCreate,
		Read:   resourceSpotinstOceanAKSVirtualNodeGroupRead,
		Update: resourceSpotinstOceanAKSVirtualNodeGroupUpdate,
		Delete: resourceSpotinstOceanAKSVirtualNodeGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.OceanAKSVirtualNodeGroupResource.GetSchemaMap(),
	}
}

func setupOceanAKSVirtualNodeGroupResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	ocean_aks_virtual_node_group.Setup(fieldsMap)
	ocean_aks_virtual_node_group_auto_scaling.Setup(fieldsMap)
	ocean_aks_virtual_node_group_launch_specification.Setup(fieldsMap)

	commons.OceanAKSVirtualNodeGroupResource = commons.NewOceanAKSVirtualNodeGroupResource(fieldsMap)
}

// region Create

func resourceSpotinstOceanAKSVirtualNodeGroupCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate), commons.OceanAKSVirtualNodeGroupResource.GetName())

	virtualNodeGroup, err := commons.OceanAKSVirtualNodeGroupResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	virtualNodeGroupID, err := createAKSVirtualNodeGroup(context.TODO(), virtualNodeGroup, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(virtualNodeGroupID))
	log.Printf("ocean/aks: virtual node group created successfully: %s", resourceData.Id())

	return resourceSpotinstOceanAKSVirtualNodeGroupRead(resourceData, meta)
}

func createAKSVirtualNodeGroup(ctx context.Context, virtualNodeGroup *azure.VirtualNodeGroup, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(virtualNodeGroup); err != nil {
		return nil, err
	} else {
		log.Printf("ocean/aks: virtual node group create configuration: %s", json)
	}

	input := &azure.CreateVirtualNodeGroupInput{
		VirtualNodeGroup: virtualNodeGroup,
	}

	output, err := spotinstClient.ocean.CloudProviderAzure().CreateVirtualNodeGroup(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("ocean/aks: failed to create cluster: %v", err)
	}

	return output.VirtualNodeGroup.ID, nil
}

// endregion

// region Read

const ErrCodeAKSVirtualNodeGroupNotFound = "CANT_GET_OCEAN_LAUNCH_SPEC"

func resourceSpotinstOceanAKSVirtualNodeGroupRead(resourceData *schema.ResourceData, meta interface{}) error {
	virtualNodeGroupID := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanAKSVirtualNodeGroupResource.GetName(), virtualNodeGroupID)

	virtualNodeGroup, err := readAKSVirtualNodeGroup(context.TODO(), virtualNodeGroupID, meta.(*Client))
	if err != nil {
		return err
	}

	// If nothing was found, return no state.
	if virtualNodeGroup == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanAKSVirtualNodeGroupResource.OnRead(virtualNodeGroup, resourceData, meta); err != nil {
		return err
	}

	log.Printf("ocean/aks: virtual node group read successfully: %s", virtualNodeGroupID)
	return nil
}

func readAKSVirtualNodeGroup(ctx context.Context, virtualNodeGroupID string, spotinstClient *Client) (*azure.VirtualNodeGroup, error) {
	input := &azure.ReadVirtualNodeGroupInput{
		VirtualNodeGroupID: spotinst.String(virtualNodeGroupID),
	}

	output, err := spotinstClient.ocean.CloudProviderAzure().ReadVirtualNodeGroup(ctx, input)
	if err != nil {
		// If the virtual node group was not found, return nil so that we can
		// show that it does not exist.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeAKSVirtualNodeGroupNotFound {
					return nil, nil
				}
			}
		}

		// Some other error, report it.
		return nil, fmt.Errorf("ocean/aks: failed to read virtual node group: %v", err)
	}

	return output.VirtualNodeGroup, nil
}

// endregion

// region Update

func resourceSpotinstOceanAKSVirtualNodeGroupUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	virtualNodeGroupID := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.OceanAKSVirtualNodeGroupResource.GetName(), virtualNodeGroupID)

	shouldUpdate, virtualNodeGroup, err := commons.OceanAKSVirtualNodeGroupResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		virtualNodeGroup.SetId(spotinst.String(virtualNodeGroupID))
		if err = updateAKSVirtualNodeGroup(context.TODO(), virtualNodeGroup, meta.(*Client)); err != nil {
			return err
		}
	}

	log.Printf("ocean/aks: virtual node group updated successfully: %s", virtualNodeGroupID)
	return resourceSpotinstOceanAKSVirtualNodeGroupRead(resourceData, meta)
}

func updateAKSVirtualNodeGroup(ctx context.Context, virtualNodeGroup *azure.VirtualNodeGroup, spotinstClient *Client) error {
	input := &azure.UpdateVirtualNodeGroupInput{
		VirtualNodeGroup: virtualNodeGroup,
	}

	if json, err := commons.ToJson(virtualNodeGroup); err != nil {
		return err
	} else {
		log.Printf("ocean/aks: virtual node group update configuration: %s", json)
	}

	if _, err := spotinstClient.ocean.CloudProviderAzure().UpdateVirtualNodeGroup(ctx, input); err != nil {
		return fmt.Errorf("ocean/aks: failed to update virtual node group: %v", err)
	}

	return nil
}

// endregion

// region Delete

func resourceSpotinstOceanAKSVirtualNodeGroupDelete(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanAKSVirtualNodeGroupResource.GetName(), resourceData.Id())

	if err := deleteAKSVirtualNodeGroup(context.TODO(), resourceData, meta.(*Client)); err != nil {
		return err
	}

	log.Printf("ocean/aks: virtual node group deleted successfully: %s", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteAKSVirtualNodeGroup(ctx context.Context, resourceData *schema.ResourceData, spotinstClient *Client) error {
	input := &azure.DeleteVirtualNodeGroupInput{
		VirtualNodeGroupID: spotinst.String(resourceData.Id()),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("ocean/aks: virtual node group delete configuration: %s", json)
	}

	if _, err := spotinstClient.ocean.CloudProviderAzure().DeleteVirtualNodeGroup(ctx, input); err != nil {
		return fmt.Errorf("ocean/aks: failed to delete virtual node group: %v", err)
	}

	return nil
}

// endregion

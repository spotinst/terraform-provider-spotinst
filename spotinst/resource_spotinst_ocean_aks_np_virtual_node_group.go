package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_np_virtual_node_group"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_np_virtual_node_group_auto_scale"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_np_virtual_node_group_node_count_limits"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_np_virtual_node_group_node_pool_properties"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_np_virtual_node_group_strategy"
	"log"
)

func resourceSpotinstOceanAKSNPVirtualNodeGroup() *schema.Resource {
	setupOceanAKSNPVirtualNodeGroupResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstOceanAKSNPVirtualNodeGroupCreate,
		ReadContext:   resourceSpotinstOceanAKSNPVirtualNodeGroupRead,
		UpdateContext: resourceSpotinstOceanAKSNPVirtualNodeGroupUpdate,
		DeleteContext: resourceSpotinstOceanAKSNPVirtualNodeGroupDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.OceanAKSNPVirtualNodeGroupResource.GetSchemaMap(),
	}
}

func setupOceanAKSNPVirtualNodeGroupResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	ocean_aks_np_virtual_node_group.Setup(fieldsMap)
	ocean_aks_np_virtual_node_group_auto_scale.Setup(fieldsMap)
	ocean_aks_np_virtual_node_group_node_pool_properties.Setup(fieldsMap)
	ocean_aks_np_virtual_node_group_node_count_limits.Setup(fieldsMap)
	ocean_aks_np_virtual_node_group_strategy.Setup(fieldsMap)

	commons.OceanAKSNPVirtualNodeGroupResource = commons.NewOceanAKSNPVirtualNodeGroupResource(fieldsMap)
}

// region Create

func resourceSpotinstOceanAKSNPVirtualNodeGroupCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate), commons.OceanAKSNPVirtualNodeGroupResource.GetName())

	virtualNodeGroup, err := commons.OceanAKSNPVirtualNodeGroupResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	virtualNodeGroupID, err := createAKSNPVirtualNodeGroup(context.TODO(), virtualNodeGroup, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(virtualNodeGroupID))
	log.Printf("ocean/aks: virtual node group created successfully: %s", resourceData.Id())

	return resourceSpotinstOceanAKSNPVirtualNodeGroupRead(ctx, resourceData, meta)
}

func createAKSNPVirtualNodeGroup(ctx context.Context, virtualNodeGroup *azure_np.VirtualNodeGroup, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(virtualNodeGroup); err != nil {
		return nil, err
	} else {
		log.Printf("ocean/aks-np: virtual node group create configuration: %s", json)
	}

	input := &azure_np.CreateVirtualNodeGroupInput{
		VirtualNodeGroup: virtualNodeGroup,
	}

	output, err := spotinstClient.ocean.CloudProviderAzureNP().CreateVirtualNodeGroup(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("ocean/aks: failed to create cluster: %v", err)
	}

	return output.VirtualNodeGroup.ID, nil
}

// endregion

// region Read

const ErrCodeAKSNPVirtualNodeGroupNotFound = "CANT_GET_OCEAN_LAUNCH_SPEC"

func resourceSpotinstOceanAKSNPVirtualNodeGroupRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	virtualNodeGroupID := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanAKSVirtualNodeGroupResource.GetName(), virtualNodeGroupID)

	virtualNodeGroup, err := readAKSNPVirtualNodeGroup(context.TODO(), virtualNodeGroupID, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	// If nothing was found, return no state.
	if virtualNodeGroup == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanAKSNPVirtualNodeGroupResource.OnRead(virtualNodeGroup, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("ocean/aks-np: virtual node group read successfully: %s", virtualNodeGroupID)
	return nil
}

func readAKSNPVirtualNodeGroup(ctx context.Context, virtualNodeGroupID string, spotinstClient *Client) (*azure_np.VirtualNodeGroup, error) {
	input := &azure_np.ReadVirtualNodeGroupInput{
		VirtualNodeGroupID: spotinst.String(virtualNodeGroupID),
	}

	output, err := spotinstClient.ocean.CloudProviderAzureNP().ReadVirtualNodeGroup(ctx, input)
	if err != nil {
		// If the virtual node group was not found, return nil so that we can
		// show that it does not exist.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeAKSNPVirtualNodeGroupNotFound {
					return nil, nil
				}
			}
		}

		// Some other error, report it.
		return nil, fmt.Errorf("ocean/aks-np: failed to read virtual node group: %v", err)
	}

	return output.VirtualNodeGroup, nil
}

// endregion

// region Update

func resourceSpotinstOceanAKSNPVirtualNodeGroupUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	virtualNodeGroupID := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.OceanAKSNPVirtualNodeGroupResource.GetName(), virtualNodeGroupID)

	shouldUpdate, virtualNodeGroup, err := commons.OceanAKSNPVirtualNodeGroupResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		virtualNodeGroup.SetId(spotinst.String(virtualNodeGroupID))
		if err = updateAKSNPVirtualNodeGroup(context.TODO(), virtualNodeGroup, meta.(*Client)); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("ocean/aks: virtual node group updated successfully: %s", virtualNodeGroupID)
	return resourceSpotinstOceanAKSNPVirtualNodeGroupRead(ctx, resourceData, meta)
}

func updateAKSNPVirtualNodeGroup(ctx context.Context, virtualNodeGroup *azure_np.VirtualNodeGroup, spotinstClient *Client) error {
	input := &azure_np.UpdateVirtualNodeGroupInput{
		VirtualNodeGroup: virtualNodeGroup,
	}

	if json, err := commons.ToJson(virtualNodeGroup); err != nil {
		return err
	} else {
		log.Printf("ocean/aks-np: virtual node group update configuration: %s", json)
	}

	if _, err := spotinstClient.ocean.CloudProviderAzureNP().UpdateVirtualNodeGroup(ctx, input); err != nil {
		return fmt.Errorf("ocean/aks-np: failed to update virtual node group: %v", err)
	}

	return nil
}

// endregion

// region Delete

func resourceSpotinstOceanAKSNPVirtualNodeGroupDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanAKSNPVirtualNodeGroupResource.GetName(), resourceData.Id())

	if err := deleteAKSNPVirtualNodeGroup(context.TODO(), resourceData, meta.(*Client)); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("ocean/aks-np: virtual node group deleted successfully: %s", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteAKSNPVirtualNodeGroup(ctx context.Context, resourceData *schema.ResourceData, spotinstClient *Client) error {
	input := &azure_np.DeleteVirtualNodeGroupInput{
		VirtualNodeGroupID: spotinst.String(resourceData.Id()),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("ocean/aks-np: virtual node group delete configuration: %s", json)
	}

	if _, err := spotinstClient.ocean.CloudProviderAzureNP().DeleteVirtualNodeGroup(ctx, input); err != nil {
		return fmt.Errorf("ocean/aks-np: failed to delete virtual node group: %v", err)
	}

	return nil
}

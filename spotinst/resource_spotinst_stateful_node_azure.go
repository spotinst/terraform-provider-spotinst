package spotinst

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/azure_v3/stateful_node_azure"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/azure_v3/stateful_node_azure_extension"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/azure_v3/stateful_node_azure_health"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/azure_v3/stateful_node_azure_image"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/azure_v3/stateful_node_azure_launch_spec"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/azure_v3/stateful_node_azure_load_balancer"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/azure_v3/stateful_node_azure_login"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/azure_v3/stateful_node_azure_network"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/azure_v3/stateful_node_azure_persistence"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/azure_v3/stateful_node_azure_scheduling"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/azure_v3/stateful_node_azure_secret"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/azure_v3/stateful_node_azure_strategy"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/azure_v3/stateful_node_azure_vm_sizes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceSpotinstStatefulNodeAzureV3() *schema.Resource {
	setupStatefulNodeAzureV3Resource()

	return &schema.Resource{
		CreateContext: resourceSpotinstStatefulNodeAzureV3Create,
		ReadContext:   resourceSpotinstStatefulNodeAzureV3Read,
		UpdateContext: resourceSpotinstStatefulNodeAzureV3Update,
		DeleteContext: resourceSpotinstStatefulNodeAzureV3Delete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.StatefulNodeAzureV3Resource.GetSchemaMap(),
	}
}

func setupStatefulNodeAzureV3Resource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	stateful_node_azure.Setup(fieldsMap)
	stateful_node_azure_strategy.Setup(fieldsMap)
	stateful_node_azure_launch_spec.Setup(fieldsMap)
	stateful_node_azure_image.Setup(fieldsMap)
	stateful_node_azure_network.Setup(fieldsMap)
	stateful_node_azure_login.Setup(fieldsMap)
	stateful_node_azure_load_balancer.Setup(fieldsMap)
	stateful_node_azure_extension.Setup(fieldsMap)
	stateful_node_azure_secret.Setup(fieldsMap)
	stateful_node_azure_vm_sizes.Setup(fieldsMap)
	stateful_node_azure_persistence.Setup(fieldsMap)
	stateful_node_azure_scheduling.Setup(fieldsMap)
	stateful_node_azure_health.Setup(fieldsMap)

	commons.StatefulNodeAzureV3Resource = commons.NewStatefulNodeAzureV3Resource(fieldsMap)
}

func resourceSpotinstStatefulNodeAzureV3Create(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.StatefulNodeAzureV3Resource.GetName())

	statefulNode, err := commons.StatefulNodeAzureV3Resource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if importVMConfig, ok := resourceData.GetOk(string(stateful_node_azure.ImportVM)); ok {
		importVMStatefulNodeInput, err := expandStatefulNodeAzureImportVMConfig(importVMConfig, statefulNode)
		if err != nil {
			return diag.Errorf("stateful node/azure: failed expanding import vm configuration: %v", err)
		}

		statefulNodeId, err := createAzureV3StatefulNodeImportVM(importVMStatefulNodeInput, meta.(*Client))
		if err != nil {
			return diag.FromErr(err)
		}

		resourceData.SetId(spotinst.StringValue(statefulNodeId))
		log.Printf("===> Stateful node using import vm created successfully: %s <===", resourceData.Id())

	} else {
		statefulNodeId, err := createAzureV3StatefulNode(statefulNode, meta.(*Client))
		if err != nil {
			return diag.FromErr(err)
		}

		resourceData.SetId(spotinst.StringValue(statefulNodeId))
		log.Printf("===> Stateful node created successfully: %s <===", resourceData.Id())
	}

	return resourceSpotinstStatefulNodeAzureV3Read(ctx, resourceData, meta)
}

func expandStatefulNodeAzureImportVMConfig(data interface{}, statefulNode *azure.StatefulNode) (*azure.ImportVMStatefulNodeInput, error) {
	spec := &azure.ImportVMStatefulNodeInput{
		StatefulNodeImport: &azure.StatefulNodeImport{
			StatefulNode: statefulNode,
		},
	}

	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(stateful_node_azure.ImportVMOriginalVMName)].(string); ok && v != "" {
			spec.StatefulNodeImport.OriginalVMName = spotinst.String(v)
		}

		if v, ok := m[string(stateful_node_azure.ImportVMResourceGroupName)].(string); ok && v != "" {
			spec.StatefulNodeImport.ResourceGroupName = spotinst.String(v)
		}

		if v, ok := m[string(stateful_node_azure.ImportVMDrainingTimeout)].(int); ok && v >= 0 {
			spec.StatefulNodeImport.DrainingTimeout = spotinst.Int(v)
		}

		if v, ok := m[string(stateful_node_azure.ImportVMResourcesRetentionTime)].(int); ok && v >= 0 {
			spec.StatefulNodeImport.ResourcesRetentionTime = spotinst.Int(v)
		}
	}

	return spec, nil
}

func createAzureV3StatefulNodeImportVM(importVMStatefulNodeInput *azure.ImportVMStatefulNodeInput, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(importVMStatefulNodeInput); err != nil {
		return nil, err
	} else {
		log.Printf("===> Stateful node import vm create configuration: %s", json)
	}

	var resp *azure.ImportVMStatefulNodeOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.statefulNode.CloudProviderAzure().ImportVM(context.Background(), importVMStatefulNodeInput)
		if err != nil {
			log.Printf("error: %v", err)
			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create stateful node using import vm: %s", err)
	}
	return resp.StatefulNodeImport.StatefulNode.ID, nil
}

func createAzureV3StatefulNode(statefulNode *azure.StatefulNode, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(statefulNode); err != nil {
		return nil, err
	} else {
		log.Printf("===> Stateful node create configuration: %s", json)
	}

	var resp *azure.CreateStatefulNodeOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &azure.CreateStatefulNodeInput{StatefulNode: statefulNode}
		r, err := spotinstClient.statefulNode.CloudProviderAzure().Create(context.Background(), input)
		if err != nil {
			log.Printf("error: %v", err)
			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create stateful node: %s", err)
	}
	return resp.StatefulNode.ID, nil
}

func resourceSpotinstStatefulNodeAzureV3Read(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceFieldOnRead),
		commons.StatefulNodeAzureV3Resource.GetName(), id)

	input := &azure.ReadStatefulNodeInput{ID: spotinst.String(id)}
	resp, err := meta.(*Client).statefulNode.CloudProviderAzure().Read(context.Background(), input)
	if err != nil {
		// If the stateful node was not found, return nil so that we can show
		// that the stateful node does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGroupNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return diag.Errorf("failed to read stateful node: %s", err)
	}

	// If nothing was found, then return no state.
	statefulNodeResponse := resp.StatefulNode
	if statefulNodeResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.StatefulNodeAzureV3Resource.OnRead(statefulNodeResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> Stateful node read successfully: %s <===", id)
	return nil
}

func resourceSpotinstStatefulNodeAzureV3Update(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.StatefulNodeAzureV3Resource.GetName(), id)

	shouldUpdate, statefulNode, err := commons.StatefulNodeAzureV3Resource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		statefulNode.SetID(spotinst.String(id))
		if err := updateAzureV3StatefulNode(statefulNode, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("===> Stateful node updated successfully: %s <===", id)
	return resourceSpotinstStatefulNodeAzureV3Read(ctx, resourceData, meta)
}

func updateAzureV3StatefulNode(statefulNode *azure.StatefulNode, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &azure.UpdateStatefulNodeInput{
		StatefulNode: statefulNode,
	}

	statefulNodeId := resourceData.Id()
	var shouldUpdateState = false
	var shouldDetachDataDisk = false
	var shouldAttachDataDisk = false
	if updateState, ok := resourceData.GetOk(string(stateful_node_azure.UpdateState)); ok {
		list := updateState.([]interface{})
		if len(list) > 0 && list[0] != nil {
			shouldUpdateState = true
		}
	}

	if attachDataDisk, ok := resourceData.GetOk(string(stateful_node_azure.AttachDataDisk)); ok {
		list := attachDataDisk.([]interface{})
		if len(list) > 0 && list[0] != nil {
			shouldAttachDataDisk = true
		}
	}

	if detachDataDisk, ok := resourceData.GetOk(string(stateful_node_azure.DetachDataDisk)); ok {
		list := detachDataDisk.([]interface{})
		if len(list) > 0 && list[0] != nil {
			shouldDetachDataDisk = true
		}
	}

	if json, err := commons.ToJson(statefulNode); err != nil {
		return err
	} else {
		log.Printf("===> Stateful node update configuration: %s", json)
	}

	if _, err := meta.(*Client).statefulNode.CloudProviderAzure().Update(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update stateful node [%v]: %v", statefulNodeId, err)
	} else if shouldUpdateState {
		if err := updateStateAzureV3StatefulNode(resourceData, meta); err != nil {
			log.Printf("[ERROR] Stateful node [%v] state update failed, error: %v", statefulNodeId, err)
			return err
		}
	} else {
		log.Printf("onUpdate() -> Field [%v] is missing, skipping state update for stateful node",
			string(stateful_node_azure.UpdateState))
	}

	if shouldAttachDataDisk {
		if err := attachDataDiskAzureV3StatefulNode(resourceData, meta); err != nil {
			log.Printf("[ERROR] Stateful node [%v] attach data disk failed, error: %v", statefulNodeId, err)
			return err
		}
	} else {
		log.Printf("onUpdate() -> Field [%v] is missing, skipping attach data disk for stateful node",
			string(stateful_node_azure.AttachDataDisk))
	}

	if shouldDetachDataDisk {
		if err := detachDataDiskAzureV3StatefulNode(resourceData, meta); err != nil {
			log.Printf("[ERROR] Stateful node [%v] detach data disk failed, error: %v", statefulNodeId, err)
			return err
		}
	} else {
		log.Printf("onUpdate() -> Field [%v] is missing, skipping detach data disk for stateful node",
			string(stateful_node_azure.DetachDataDisk))
	}

	return nil
}

func updateStateAzureV3StatefulNode(resourceData *schema.ResourceData, meta interface{}) error {
	statefulNodeID := resourceData.Id()

	updateState, ok := resourceData.GetOk(string(stateful_node_azure.UpdateState))
	if !ok {
		return fmt.Errorf("stateful node/azure: missing update_state for stateful node %q", statefulNodeID)
	}

	list := updateState.([]interface{})
	if len(list) > 0 && list[0] != nil {
		updateStatefulNodeStateSchema := list[0].(map[string]interface{})
		if updateStatefulNodeStateSchema == nil {
			return fmt.Errorf("stateful node/azure: missing update state configuration, "+
				"skipping update state for stateful node %q", statefulNodeID)
		}

		updateStateSpec, err := expandStatefulNodeAzureUpdateStateConfig(updateStatefulNodeStateSchema, statefulNodeID)
		if err != nil {
			return fmt.Errorf("stateful node/azure: failed expanding state update "+
				"configuration for stateful node %q, error: %v", statefulNodeID, err)
		}

		updateStateJSON, err := commons.ToJson(updateStatefulNodeStateSchema)
		if err != nil {
			return fmt.Errorf("stateful node/azure: failed marshaling state update "+
				"configuration for stateful node %q, error: %v", statefulNodeID, err)
		}

		log.Printf("onUpdate() -> Updating stateful node [%v] with configuration %s", statefulNodeID, updateStateJSON)
		updateStateInput := &azure.UpdateStatefulNodeStateInput{ID: updateStateSpec.ID,
			StatefulNodeState: updateStateSpec.StatefulNodeState}
		if _, err = meta.(*Client).statefulNode.CloudProviderAzure().UpdateState(context.TODO(),
			updateStateInput); err != nil {
			return fmt.Errorf("onUpdate() -> State update failed for stateful node [%v], error: %v",
				statefulNodeID, err)
		}
		log.Printf("onUpdate() -> Successfully updated state for stateful node [%v]", statefulNodeID)
	}

	return nil
}

func attachDataDiskAzureV3StatefulNode(resourceData *schema.ResourceData, meta interface{}) error {
	statefulNodeID := resourceData.Id()

	attachDataDisk, ok := resourceData.GetOk(string(stateful_node_azure.AttachDataDisk))
	if !ok {
		return fmt.Errorf("stateful node/azure: missing attach_data_disk for stateful node %q", statefulNodeID)
	}

	list := attachDataDisk.([]interface{})
	if len(list) > 0 && list[0] != nil {
		attachDataDiskStatefulNodeSchema := list[0].(map[string]interface{})
		if attachDataDiskStatefulNodeSchema == nil {
			return fmt.Errorf("stateful node/azure: missing attach data disk configuration, "+
				"skipping attach data disk for stateful node %q", statefulNodeID)
		}

		attachDataDiskSpec, err := expandStatefulNodeAzureAttachDataDiskConfig(attachDataDiskStatefulNodeSchema, statefulNodeID)
		if err != nil {
			return fmt.Errorf("stateful node/azure: failed expanding attach data disk "+
				"configuration for stateful node %q, error: %v", statefulNodeID, err)
		}

		updateStateJSON, err := commons.ToJson(attachDataDiskStatefulNodeSchema)
		if err != nil {
			return fmt.Errorf("stateful node/azure: failed marshaling attach data disk "+
				"configuration for stateful node %q, error: %v", statefulNodeID, err)
		}

		log.Printf("onUpdate() -> Updating stateful node [%v] with configuration %s", statefulNodeID, updateStateJSON)
		attachDataDiskInput := &azure.AttachStatefulNodeDataDiskInput{
			ID:                        attachDataDiskSpec.ID,
			DataDiskName:              attachDataDiskSpec.DataDiskName,
			DataDiskResourceGroupName: attachDataDiskSpec.DataDiskResourceGroupName,
			StorageAccountType:        attachDataDiskSpec.StorageAccountType,
			SizeGB:                    attachDataDiskSpec.SizeGB,
			LUN:                       attachDataDiskSpec.LUN,
			Zone:                      attachDataDiskSpec.Zone}
		if _, err = meta.(*Client).statefulNode.CloudProviderAzure().AttachDataDisk(context.TODO(),
			attachDataDiskInput); err != nil {
			return fmt.Errorf("onUpdate() -> Attach data disk failed for stateful node [%v], error: %v",
				statefulNodeID, err)
		}
		log.Printf("onUpdate() -> Successfully attached data disk for stateful node [%v]", statefulNodeID)
	}

	return nil
}

func detachDataDiskAzureV3StatefulNode(resourceData *schema.ResourceData, meta interface{}) error {
	statefulNodeID := resourceData.Id()

	detachDataDisk, ok := resourceData.GetOk(string(stateful_node_azure.DetachDataDisk))
	if !ok {
		return fmt.Errorf("stateful node/azure: missing detach_data_disk for stateful node %q", statefulNodeID)
	}

	list := detachDataDisk.([]interface{})
	if len(list) > 0 && list[0] != nil {
		detachDataDiskStatefulNodeSchema := list[0].(map[string]interface{})
		if detachDataDiskStatefulNodeSchema == nil {
			return fmt.Errorf("stateful node/azure: missing detach data disk configuration, "+
				"skipping detach data disk for stateful node %q", statefulNodeID)
		}

		detachDataDiskSpec, err := expandStatefulNodeAzureDetachDataDiskConfig(detachDataDiskStatefulNodeSchema, statefulNodeID)
		if err != nil {
			return fmt.Errorf("stateful node/azure: failed expanding detach data disk "+
				"configuration for stateful node %q, error: %v", statefulNodeID, err)
		}

		updateStateJSON, err := commons.ToJson(detachDataDiskStatefulNodeSchema)
		if err != nil {
			return fmt.Errorf("stateful node/azure: failed marshaling detach data disk "+
				"configuration for stateful node %q, error: %v", statefulNodeID, err)
		}

		log.Printf("onUpdate() -> Updating stateful node [%v] with configuration %s", statefulNodeID, updateStateJSON)
		detachDataDiskInput := &azure.DetachStatefulNodeDataDiskInput{
			ID:                        detachDataDiskSpec.ID,
			DataDiskName:              detachDataDiskSpec.DataDiskName,
			DataDiskResourceGroupName: detachDataDiskSpec.DataDiskResourceGroupName,
			ShouldDeallocate:          detachDataDiskSpec.ShouldDeallocate}
		if _, err = meta.(*Client).statefulNode.CloudProviderAzure().DetachDataDisk(context.TODO(),
			detachDataDiskInput); err != nil {
			return fmt.Errorf("onUpdate() -> detach data disk failed for stateful node [%v], error: %v",
				statefulNodeID, err)
		}
		log.Printf("onUpdate() -> Successfully detached data disk for stateful node [%v]", statefulNodeID)
	}

	return nil
}

func expandStatefulNodeAzureUpdateStateConfig(data interface{}, statefulNodeID string) (*azure.UpdateStatefulNodeStateInput, error) {
	spec := &azure.UpdateStatefulNodeStateInput{
		ID: spotinst.String(statefulNodeID),
	}

	if data != nil {
		m := data.(map[string]interface{})

		if v, ok := m[string(stateful_node_azure.State)].(string); ok && v != "" {
			spec.StatefulNodeState = spotinst.String(v)
		}
	}

	return spec, nil
}

func expandStatefulNodeAzureAttachDataDiskConfig(data interface{},
	statefulNodeID string) (*azure.AttachStatefulNodeDataDiskInput, error) {
	spec := &azure.AttachStatefulNodeDataDiskInput{
		ID: spotinst.String(statefulNodeID),
	}

	if data != nil {
		m := data.(map[string]interface{})

		if v, ok := m[string(stateful_node_azure.AttachDataDiskName)].(string); ok && v != "" {
			spec.DataDiskName = spotinst.String(v)
		}

		if v, ok := m[string(stateful_node_azure.AttachDataDiskResourceGroupName)].(string); ok && v != "" {
			spec.DataDiskResourceGroupName = spotinst.String(v)
		}

		if v, ok := m[string(stateful_node_azure.AttachStorageAccountType)].(string); ok && v != "" {
			spec.StorageAccountType = spotinst.String(v)
		}

		if v, ok := m[string(stateful_node_azure.AttachSizeGB)].(int); ok && v > 0 {
			spec.SizeGB = spotinst.Int(v)
		}

		if v, ok := m[string(stateful_node_azure.AttachLUN)].(int); ok && v >= 0 {
			spec.LUN = spotinst.Int(v)
		}

		if v, ok := m[string(stateful_node_azure.AttachZone)].(string); ok && v != "" {
			spec.Zone = spotinst.String(v)
		}
	}

	return spec, nil
}

func expandStatefulNodeAzureDetachDataDiskConfig(data interface{}, statefulNodeID string) (*azure.DetachStatefulNodeDataDiskInput, error) {
	spec := &azure.DetachStatefulNodeDataDiskInput{
		ID: spotinst.String(statefulNodeID),
	}

	if data != nil {
		m := data.(map[string]interface{})

		if v, ok := m[string(stateful_node_azure.DetachDataDiskName)].(string); ok && v != "" {
			spec.DataDiskName = spotinst.String(v)
		}

		if v, ok := m[string(stateful_node_azure.DetachDataDiskResourceGroupName)].(string); ok && v != "" {
			spec.DataDiskResourceGroupName = spotinst.String(v)
		}

		if v, ok := m[string(stateful_node_azure.DetachShouldDeallocate)].(bool); ok {
			spec.ShouldDeallocate = spotinst.Bool(v)
		}

		if v, ok := m[string(stateful_node_azure.DetachTTLInHours)].(int); ok && v > 0 {
			spec.TTLInHours = spotinst.Int(v)
		}

	}

	return spec, nil
}

func resourceSpotinstStatefulNodeAzureV3Delete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.StatefulNodeAzureV3Resource.GetName(), id)

	if err := deleteAzureV3StatefulNode(resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Stateful node deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteAzureV3StatefulNode(resourceData *schema.ResourceData, meta interface{}) error {
	statefulNodeId := resourceData.Id()
	if deleteConfig, ok := resourceData.GetOk(string(stateful_node_azure.Delete)); ok {
		deleteStatefulNodeAzureInput, err := expandStatefulNodeAzureDeleteConfig(deleteConfig, statefulNodeId)
		if err != nil {
			return fmt.Errorf("stateful node/azure: failed expanding delete configuration: %v", err)
		}

		if _, err := meta.(*Client).statefulNode.CloudProviderAzure().Delete(context.Background(), deleteStatefulNodeAzureInput); err != nil {
			return fmt.Errorf("[ERROR] onDelete() -> Failed to delete stateful node: %s", err)
		}

		return nil
	}

	return fmt.Errorf("stateful node/azure: missing delete configuration")

}

func expandStatefulNodeAzureDeleteConfig(data interface{}, statefulNodeID string) (*azure.DeleteStatefulNodeInput, error) {
	spec := &azure.DeleteStatefulNodeInput{
		ID: spotinst.String(statefulNodeID),
		DeallocationConfig: &azure.DeallocationConfig{
			NetworkDeallocationConfig:  &azure.ResourceDeallocationConfig{},
			DiskDeallocationConfig:     &azure.ResourceDeallocationConfig{},
			SnapshotDeallocationConfig: &azure.ResourceDeallocationConfig{},
			PublicIPDeallocationConfig: &azure.ResourceDeallocationConfig{},
		},
	}

	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(stateful_node_azure.ShouldTerminateVm)].(bool); ok {
			spec.DeallocationConfig.ShouldTerminateVM = spotinst.Bool(v)
		}

		if v, ok := m[string(stateful_node_azure.ShouldDeallocateNetwork)].(bool); ok {
			spec.DeallocationConfig.NetworkDeallocationConfig.ShouldDeallocate = spotinst.Bool(v)
		}

		if v, ok := m[string(stateful_node_azure.NetworkTTLInHours)].(int); ok && v >= 0 {
			spec.DeallocationConfig.NetworkDeallocationConfig.TTLInHours = spotinst.Int(v)
		}

		if v, ok := m[string(stateful_node_azure.ShouldDeallocateDisk)].(bool); ok {
			spec.DeallocationConfig.DiskDeallocationConfig.ShouldDeallocate = spotinst.Bool(v)
		}

		if v, ok := m[string(stateful_node_azure.DiskTTLInHours)].(int); ok && v >= 0 {
			spec.DeallocationConfig.DiskDeallocationConfig.TTLInHours = spotinst.Int(v)
		}

		if v, ok := m[string(stateful_node_azure.ShouldDeallocateSnapshot)].(bool); ok {
			spec.DeallocationConfig.SnapshotDeallocationConfig.ShouldDeallocate = spotinst.Bool(v)
		}

		if v, ok := m[string(stateful_node_azure.SnapshotTTLInHours)].(int); ok && v >= 0 {
			spec.DeallocationConfig.SnapshotDeallocationConfig.TTLInHours = spotinst.Int(v)
		}

		if v, ok := m[string(stateful_node_azure.ShouldDeallocatePublicIP)].(bool); ok {
			spec.DeallocationConfig.PublicIPDeallocationConfig.ShouldDeallocate = spotinst.Bool(v)
		}

		if v, ok := m[string(stateful_node_azure.PublicIPTTLInHours)].(int); ok && v >= 0 {
			spec.DeallocationConfig.PublicIPDeallocationConfig.TTLInHours = spotinst.Int(v)
		}

		if v, ok := m[string(stateful_node_azure.ShouldDeregisterFromLb)].(bool); ok {
			spec.DeallocationConfig.ShouldDeregisterFromLb = spotinst.Bool(v)
		}
	}

	return spec, nil
}

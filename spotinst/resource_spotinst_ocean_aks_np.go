package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"

	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_np_scheduling"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_np_vm_sizes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_np"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_np_auto_scale"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_np_auto_scaler"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_np_health"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_np_node_count_limits"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_np_node_pool_properties"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aks_np_strategy"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceSpotinstOceanAKSNP() *schema.Resource {
	setupClusterAKSNPResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstClusterAKSNPCreate,
		ReadContext:   resourceSpotinstClusterAKSNPRead,
		UpdateContext: resourceSpotinstClusterAKSNPUpdate,
		DeleteContext: resourceSpotinstClusterAKSNPDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.OceanAKSNPResource.GetSchemaMap(),
	}
}

func setupClusterAKSNPResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	ocean_aks_np.Setup(fieldsMap)
	ocean_aks_np_auto_scaler.Setup(fieldsMap)
	ocean_aks_np_strategy.Setup(fieldsMap)
	ocean_aks_np_health.Setup(fieldsMap)
	ocean_aks_np_node_pool_properties.Setup(fieldsMap)
	ocean_aks_np_node_count_limits.Setup(fieldsMap)
	ocean_aks_np_auto_scale.Setup(fieldsMap)
	ocean_aks_np_scheduling.Setup(fieldsMap)
	ocean_aks_np_vm_sizes.Setup(fieldsMap)

	commons.OceanAKSNPResource = commons.NewOceanAKSNPResource(fieldsMap)
}

// region Create

func resourceSpotinstClusterAKSNPCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate), commons.OceanAKSNPResource.GetName())

	cluster, err := commons.OceanAKSNPResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	clusterID, err := createAKSNPCluster(cluster, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(clusterID))
	log.Printf("ocean/aks: AKS cluster created successfully: %s", resourceData.Id())

	return resourceSpotinstClusterAKSNPRead(ctx, resourceData, meta)
}

func createAKSNPCluster(cluster *azure_np.Cluster, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(cluster); err != nil {
		return nil, err
	} else {
		log.Printf("ocean/aks: cluster configuration: %s", json)
	}

	input := &azure_np.CreateClusterInput{
		Cluster: cluster,
	}

	output, err := spotinstClient.ocean.CloudProviderAzureNP().CreateCluster(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("ocean/aks: failed to create cluster: %v", err)
	}

	return output.Cluster.ID, nil
}

// endregion

// region Read

func resourceSpotinstClusterAKSNPRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	clusterID := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanAKSNPResource.GetName(), clusterID)

	cluster, err := readAKSNPCluster(context.TODO(), clusterID, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	// If nothing was found, return no state.
	if cluster == nil {
		resourceData.SetId("")
		return nil
	}

	// Expose the controller cluster identifier.
	if cluster.ControllerClusterID != nil {
		_ = resourceData.Set(string(ocean_aks_np.ControllerClusterID),
			spotinst.StringValue(cluster.ControllerClusterID))
	}

	if err := commons.OceanAKSNPResource.OnRead(cluster, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("ocean/aks: cluster read successfully: %s", clusterID)
	return nil
}

func readAKSNPCluster(ctx context.Context, clusterID string, spotinstClient *Client) (*azure_np.Cluster, error) {
	input := &azure_np.ReadClusterInput{
		ClusterID: spotinst.String(clusterID),
	}

	output, err := spotinstClient.ocean.CloudProviderAzureNP().ReadCluster(ctx, input)
	if err != nil {
		// If the cluster was not found, return nil so that we can show that it
		// does not exist.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeClusterNotFound {
					return nil, nil
				}
			}
		}

		// Some other error, report it.
		return nil, fmt.Errorf("ocean/aks: failed to read cluster: %v", err)
	}

	return output.Cluster, nil
}

// endregion

// region Update

func resourceSpotinstClusterAKSNPUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	clusterID := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.OceanAKSNPResource.GetName(), clusterID)

	shouldUpdate, changesRequiredRoll, cluster, err := commons.OceanAKSNPResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		cluster.SetId(spotinst.String(clusterID))
		if err := updateAKSNPCluster(cluster, resourceData, meta.(*Client), changesRequiredRoll); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("ocean/aks: cluster updated successfully: %s", clusterID)
	return resourceSpotinstClusterAKSNPRead(ctx, resourceData, meta)
}

func updateAKSNPCluster(cluster *azure_np.Cluster, resourceData *schema.ResourceData, spotinstClient *Client, changesRequiredRoll bool) error {
	input := &azure_np.UpdateClusterInput{
		Cluster: cluster,
	}

	var shouldRoll = false
	var conditionedRoll = false
	clusterID := resourceData.Id()
	if updatePolicy, exists := resourceData.GetOkExists(string(ocean_aks_np.UpdatePolicy)); exists {
		list := updatePolicy.([]interface{})
		if len(list) > 0 && list[0] != nil {
			m := list[0].(map[string]interface{})

			if roll, ok := m[string(ocean_aks_np.ShouldRoll)].(bool); ok && roll {
				shouldRoll = roll
			}
			if condRoll, ok := m[string(ocean_aks_np.ConditionedRoll)].(bool); ok && condRoll {
				conditionedRoll = condRoll
			}
		}
	}

	if json, err := commons.ToJson(cluster); err != nil {
		return err
	} else {
		log.Printf("ocean/aks: cluster update configuration: %s", json)
	}

	if _, err := spotinstClient.ocean.CloudProviderAzureNP().UpdateCluster(context.TODO(), input); err != nil {
		return fmt.Errorf("ocean/aks: failed to update cluster: %v", err)
	} else if shouldRoll {
		if !conditionedRoll || changesRequiredRoll {
			if err := rollOceanAKSCluster(resourceData, spotinstClient); err != nil {
				log.Printf("[ERROR] Cluster [%v] roll failed, error: %v", clusterID, err)
				return err
			}
		}
	} else {
		log.Printf("onRoll() -> Field [%v] is false, skipping cluster roll", string(ocean_aks_np.ShouldRoll))
	}

	return nil
}

func rollOceanAKSCluster(resourceData *schema.ResourceData, meta interface{}) error {
	clusterID := resourceData.Id()

	updatePolicy, exists := resourceData.GetOkExists(string(ocean_aks_np.UpdatePolicy))
	if !exists {
		return fmt.Errorf("ocean/aksnp: missing update policy for cluster %q", clusterID)
	}

	list := updatePolicy.([]interface{})
	if len(list) > 0 && list[0] != nil {
		updateClusterSchema := list[0].(map[string]interface{})

		rollConfig, ok := updateClusterSchema[string(ocean_aks_np.RollConfig)]
		if !ok || rollConfig == nil {
			return fmt.Errorf("ocean/aksnp: missing roll configuration, "+
				"skipping roll for cluster %q", clusterID)
		}

		rollSpec, err := expandOceanAKSClusterRollConfig(rollConfig, clusterID)
		if err != nil {
			return fmt.Errorf("ocean/aks: failed expanding roll "+
				"configuration for cluster %q, error: %v", clusterID, err)
		}

		rollJSON, err := commons.ToJson(rollConfig)
		if err != nil {
			return fmt.Errorf("ocean/aks: failed marshaling roll "+
				"configuration for cluster %q, error: %v", clusterID, err)
		}

		log.Printf("onRoll() -> Rolling cluster [%v] with configuration %s", clusterID, rollJSON)
		rollInput := &azure.CreateRollInput{Roll: rollSpec}
		if _, err = meta.(*Client).ocean.CloudProviderAzure().CreateRoll(context.TODO(), rollInput); err != nil {
			return fmt.Errorf("onRoll() -> Roll failed for cluster [%v], error: %v", clusterID, err)
		}
		log.Printf("onRoll() -> Successfully rolled cluster [%v]", clusterID)
	}

	return nil
}
func expandOceanAKSClusterRollConfig(data interface{}, clusterID string) (*azure.RollSpec, error) {
	list := data.([]interface{})
	spec := &azure.RollSpec{
		ClusterID: spotinst.String(clusterID),
	}

	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(ocean_aks_np.BatchSizePercentage)].(int); ok {
			spec.BatchSizePercentage = spotinst.Int(v)
		}

		if v, ok := m[string(ocean_aks_np.VngIDs)].([]string); ok {
			spec.VngIds = expandList(v)
		}

		if v, ok := m[string(ocean_aks_np.BatchMinHealthyPercentage)].(int); ok && v > 0 {
			spec.BatchMinHealthyPercentage = spotinst.Int(v)
		}

		if v, ok := m[string(ocean_aks_np.RespectPDB)].(bool); ok {
			spec.RespectPDB = spotinst.Bool(v)
		}

		if v, ok := m[string(ocean_aks_np.Comment)].(string); ok {
			spec.Comment = spotinst.String(v)
		}

		if v, ok := m[string(ocean_aks_np.NodePoolNames)].(bool); ok {
			spec.NodePoolNames = expandList(v)
		}
		if v, ok := m[string(ocean_aks_np.RespectRestrictScaleDown)].(bool); ok {
			spec.RespectRestrictScaleDown = spotinst.Bool(v)
		}
		if v, ok := m[string(ocean_aks_np.NodeNames)].([]string); ok {
			spec.NodeNames = expandList(v)
		}

	}

	return spec, nil
}

func expandList(data interface{}) []string {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if ls, ok := v.(string); ok && ls != "" {
			result = append(result, ls)
		}
	}

	return result
}

// endregion

// region Delete

func resourceSpotinstClusterAKSNPDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	clusterID := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete), commons.OceanAKSNPResource.GetName(), clusterID)

	if err := deleteAKSNPCluster(clusterID, meta.(*Client)); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("ocean/aks: cluster deleted successfully: %s", clusterID)
	resourceData.SetId("")

	return nil
}

func deleteAKSNPCluster(clusterID string, spotinstClient *Client) error {
	input := &azure_np.DeleteClusterInput{
		ClusterID: spotinst.String(clusterID),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("ocean/aks: cluster delete configuration: %s", json)
	}

	if _, err := spotinstClient.ocean.CloudProviderAzureNP().DeleteCluster(context.TODO(), input); err != nil {
		return fmt.Errorf("ocean/aks: failed to delete cluster: %v", err)
	}

	return nil
}

// endregion

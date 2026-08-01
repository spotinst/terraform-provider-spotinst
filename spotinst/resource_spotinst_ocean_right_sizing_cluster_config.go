package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/cluster_right_sizing"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_right_sizing_cluster_config"
)

func resourceSpotinstOceanRightSizingClusterConfig() *schema.Resource {
	setupOceanRightSizingClusterConfigResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstOceanRightSizingClusterConfigCreate,
		ReadContext:   resourceSpotinstOceanRightSizingClusterConfigRead,
		UpdateContext: resourceSpotinstOceanRightSizingClusterConfigUpdate,
		DeleteContext: resourceSpotinstOceanRightSizingClusterConfigDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: commons.OceanRightSizingClusterConfigResource.GetSchemaMap(),
	}
}

func setupOceanRightSizingClusterConfigResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)
	ocean_right_sizing_cluster_config.Setup(fieldsMap)

	commons.OceanRightSizingClusterConfigResource = commons.NewOceanRightSizingClusterConfigResource(fieldsMap)
}

func resourceSpotinstOceanRightSizingClusterConfigCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate), commons.OceanRightSizingClusterConfigResource.GetName())

	input, err := commons.OceanRightSizingClusterConfigResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	oceanId := spotinst.StringValue(input.OceanId)
	clusterIdentifier := spotinst.StringValue(input.ClusterIdentifier)

	if _, err := createOceanRightSizingClusterConfig(input, meta.(*Client)); err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(buildOceanRightSizingClusterConfigID(oceanId, clusterIdentifier))

	log.Printf("===> Ocean right sizing cluster config created successfully: %s <===", resourceData.Id())
	return resourceSpotinstOceanRightSizingClusterConfigRead(ctx, resourceData, meta)
}

func createOceanRightSizingClusterConfig(input *cluster_right_sizing.PostClusterConfigurationInput, spotinstClient *Client) (*cluster_right_sizing.PostClusterConfigurationOutput, error) {
	if json, err := commons.ToJson(input); err != nil {
		return nil, err
	} else {
		log.Printf("===> Ocean right sizing cluster config create configuration: %s", json)
	}

	output, err := spotinstClient.ocean.ClusterRightSizing().PostClusterConfiguration(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create ocean right sizing cluster config: %s", err)
	}

	return output, nil
}

func resourceSpotinstOceanRightSizingClusterConfigRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceID := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanRightSizingClusterConfigResource.GetName(), resourceID)

	oceanID := resourceData.Get(string(ocean_right_sizing_cluster_config.OceanId)).(string)
	clusterIdentifier := resourceData.Get(string(ocean_right_sizing_cluster_config.ClusterIdentifier)).(string)

	if oceanID == "" || clusterIdentifier == "" {
		oceanID, clusterIdentifier = parseOceanRightSizingClusterConfigID(resourceID)
		if oceanID == "" || clusterIdentifier == "" {
			return diag.Errorf("failed to parse resource ID for spotinst_ocean_right_sizing_cluster_config: %s", resourceID)
		}
	}

	input := &cluster_right_sizing.ReadClusterConfigurationInput{
		OceanId:           spotinst.String(oceanID),
		ClusterIdentifier: spotinst.String(clusterIdentifier),
	}

	output, err := meta.(*Client).ocean.ClusterRightSizing().ReadClusterConfiguration(context.Background(), input)
	if err != nil {
		return diag.FromErr(err)
	}

	if output == nil || output.ClusterConfiguration == nil {
		resourceData.SetId("")
		return nil
	}

	if err := resourceData.Set(string(ocean_right_sizing_cluster_config.OceanId), oceanID); err != nil {
		return diag.FromErr(err)
	}
	if err := resourceData.Set(string(ocean_right_sizing_cluster_config.ClusterIdentifier), clusterIdentifier); err != nil {
		return diag.FromErr(err)
	}

	if err := commons.OceanRightSizingClusterConfigResource.OnRead(output.ClusterConfiguration, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Ocean right sizing cluster config read successfully: %s <===", resourceID)
	return nil
}

func resourceSpotinstOceanRightSizingClusterConfigUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnUpdate), commons.OceanRightSizingClusterConfigResource.GetName(), resourceData.Id())

	input, err := commons.OceanRightSizingClusterConfigResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	oceanId := spotinst.StringValue(input.OceanId)
	clusterIdentifier := spotinst.StringValue(input.ClusterIdentifier)

	if _, err := createOceanRightSizingClusterConfig(input, meta.(*Client)); err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(buildOceanRightSizingClusterConfigID(oceanId, clusterIdentifier))

	log.Printf("===> Ocean right sizing cluster config updated successfully: %s <===", resourceData.Id())
	return resourceSpotinstOceanRightSizingClusterConfigRead(ctx, resourceData, meta)
}

func resourceSpotinstOceanRightSizingClusterConfigDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[WARN] Delete is not supported for spotinst_ocean_right_sizing_cluster_config, removing from state only")
	resourceData.SetId("")
	return nil
}

func buildOceanRightSizingClusterConfigID(oceanID, clusterIdentifier string) string {
	return fmt.Sprintf("%s:%s", oceanID, clusterIdentifier)
}

func parseOceanRightSizingClusterConfigID(id string) (string, string) {
	parts := strings.SplitN(id, ":", 2)
	if len(parts) != 2 {
		return "", ""
	}
	return parts[0], parts[1]
}

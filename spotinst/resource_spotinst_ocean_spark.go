package spotinst

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/spark"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"

	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_spark"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_spark_compute"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_spark_ingress"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_spark_log_collection"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_spark_webhook"
)

const (
	ErrCodeResourceDoesNotExist = "RESOURCE_DOES_NOT_EXIST"
)

func resourceSpotinstOceanSpark() *schema.Resource {
	setupSparkClusterResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstSparkClusterCreate,
		ReadContext:   resourceSpotinstSparkClusterRead,
		UpdateContext: resourceSpotinstSparkClusterUpdate,
		DeleteContext: resourceSpotinstSparkClusterDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: commons.OceanSparkResource.GetSchemaMap(),
	}
}

func setupSparkClusterResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	ocean_spark.Setup(fieldsMap)
	ocean_spark_ingress.Setup(fieldsMap)
	ocean_spark_webhook.Setup(fieldsMap)
	ocean_spark_compute.Setup(fieldsMap)
	ocean_spark_log_collection.Setup(fieldsMap)

	commons.OceanSparkResource = commons.NewOceanSparkResource(fieldsMap)
}

func resourceSpotinstSparkClusterCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OceanSparkResource.GetName())

	cluster, err := commons.OceanSparkResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	clusterID, err := createSparkCluster(cluster, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(clusterID))

	log.Printf("===> Cluster created successfully: %s <===", resourceData.Id())
	return resourceSpotinstSparkClusterRead(ctx, resourceData, meta)
}

func createSparkCluster(cluster *spark.Cluster, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(cluster); err != nil {
		return nil, err
	} else {
		log.Printf("===> Cluster create configuration: %s", json)
	}

	createClusterRequest := &spark.CreateClusterRequest{
		OceanClusterID: cluster.OceanClusterID,
		Config:         cluster.Config,
	}

	var resp *spark.CreateClusterOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &spark.CreateClusterInput{Cluster: createClusterRequest}
		r, err := spotinstClient.ocean.Spark().CreateCluster(context.Background(), input)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create cluster: %s", err)
	}

	return resp.Cluster.ID, nil
}

func resourceSpotinstSparkClusterRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.OceanSparkResource.GetName(), id)

	input := &spark.ReadClusterInput{ClusterID: spotinst.String(id)}
	resp, err := meta.(*Client).ocean.Spark().ReadCluster(ctx, input)
	if err != nil {
		// If the cluster was not found, return nil so that we can show
		// that the cluster does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeResourceDoesNotExist {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return diag.Errorf("failed to read cluster: %s", err)
	}

	// if nothing was found, return no state
	clusterResponse := resp.Cluster
	if clusterResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanSparkResource.OnRead(clusterResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Cluster read successfully: %s <===", id)
	return nil
}

func resourceSpotinstSparkClusterUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.OceanSparkResource.GetName(), id)

	shouldUpdate, cluster, err := commons.OceanSparkResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		cluster.ID = spotinst.String(id)
		if err := updateSparkCluster(cluster, meta); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("===> Cluster updated successfully: %s <===", id)
	return resourceSpotinstSparkClusterRead(ctx, resourceData, meta)
}

func updateSparkCluster(cluster *spark.Cluster, meta interface{}) error {
	updateClusterRequest := &spark.UpdateClusterRequest{
		Config: cluster.Config,
	}

	var input = &spark.UpdateClusterInput{
		ClusterID: cluster.ID,
		Cluster:   updateClusterRequest,
	}

	if json, err := commons.ToJson(updateClusterRequest); err != nil {
		return err
	} else {
		log.Printf("===> Cluster update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.Spark().UpdateCluster(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update cluster [%v]: %v", cluster.ID, err)
	}

	return nil
}

func resourceSpotinstSparkClusterDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanSparkResource.GetName(), id)

	if err := deleteSparkCluster(ctx, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Cluster deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")

	return nil
}

func deleteSparkCluster(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) error {
	clusterID := resourceData.Id()
	input := &spark.DeleteClusterInput{
		ClusterID: spotinst.String(clusterID),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Cluster delete configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.Spark().DeleteCluster(ctx, input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete cluster: %s", err)
	}

	return nil
}

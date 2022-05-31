package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_gke"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_gke_auto_scaling"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_gke_instance_types"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_gke_network_interface"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_gke_strategy"
)

func resourceSpotinstOceanGKE() *schema.Resource {
	setupClusterGKEResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstClusterGKECreate,
		ReadContext:   resourceSpotinstClusterGKERead,
		UpdateContext: resourceSpotinstClusterGKEUpdate,
		DeleteContext: resourceSpotinstClusterGKEDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: commons.OceanGKEResource.GetSchemaMap(),
	}
}

func setupClusterGKEResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	ocean_gke.Setup(fieldsMap)
	ocean_gke_auto_scaling.Setup(fieldsMap)
	ocean_gke_instance_types.Setup(fieldsMap)
	ocean_gke_network_interface.Setup(fieldsMap)
	ocean_gke_strategy.Setup(fieldsMap)

	commons.OceanGKEResource = commons.NewOceanGKEResource(fieldsMap)
}

func resourceSpotinstClusterGKECreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OceanGKEResource.GetName())

	cluster, err := commons.OceanGKEResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	clusterID, err := createGKECluster(cluster, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(clusterID))

	log.Printf("===> Elastigroup created successfully: %s <===", resourceData.Id())
	return resourceSpotinstClusterGKERead(ctx, resourceData, meta)
}

func createGKECluster(cluster *gcp.Cluster, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(cluster); err != nil {
		return nil, err
	} else {
		log.Printf("===> Cluster create configuration: %s", json)
	}

	var resp *gcp.CreateClusterOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &gcp.CreateClusterInput{Cluster: cluster}
		r, err := spotinstClient.ocean.CloudProviderGCP().CreateCluster(context.Background(), input)
		if err != nil {

			// Some other error, report it.
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

func resourceSpotinstClusterGKERead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.OceanGKEResource.GetName(), id)

	input := &gcp.ReadClusterInput{ClusterID: spotinst.String(id)}
	resp, err := meta.(*Client).ocean.CloudProviderGCP().ReadCluster(context.Background(), input)

	if err != nil {
		// If the cluster was not found, return nil so that we can show
		// that the group does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeClusterNotFound {
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

	if err := commons.OceanGKEResource.OnRead(clusterResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> Cluster read successfully: %s <===", id)
	return nil
}

func resourceSpotinstClusterGKEUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.OceanGKEResource.GetName(), id)

	shouldUpdate, cluster, err := commons.OceanGKEResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		cluster.SetId(spotinst.String(id))
		if err := updateGKECluster(cluster, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}
	log.Printf("===> Cluster updated successfully: %s <===", id)
	return resourceSpotinstClusterGKERead(ctx, resourceData, meta)
}

func updateGKECluster(cluster *gcp.Cluster, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &gcp.UpdateClusterInput{
		Cluster: cluster,
	}

	clusterID := resourceData.Id()

	if json, err := commons.ToJson(cluster); err != nil {
		return err
	} else {
		log.Printf("===> Cluster update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderGCP().UpdateCluster(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update cluster [%v]: %v", clusterID, err)
	}

	return nil
}

func resourceSpotinstClusterGKEDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanGKEResource.GetName(), id)

	if err := deleteGKECluster(resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Cluster deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteGKECluster(resourceData *schema.ResourceData, meta interface{}) error {
	clusterID := resourceData.Id()
	input := &gcp.DeleteClusterInput{
		ClusterID: spotinst.String(clusterID),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Cluster delete configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderGCP().DeleteCluster(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete cluster: %s", err)
	}
	return nil
}

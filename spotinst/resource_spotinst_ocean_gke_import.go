package spotinst

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_gke_import"
)

func resourceSpotinstOceanGKEImport() *schema.Resource {
	setupClusterGKEImportResource()

	return &schema.Resource{
		Create: resourceSpotinstClusterGKEImportCreate,
		Read:   resourceSpotinstClusterGKEImportRead,
		Update: resourceSpotinstClusterGKEImportUpdate,
		Delete: resourceSpotinstClusterGKEImportDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: commons.OceanGKEImportResource.GetSchemaMap(),
	}
}

func setupClusterGKEImportResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	ocean_gke_import.Setup(fieldsMap)

	commons.OceanGKEImportResource = commons.NewOceanGKEImportResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Import Ocean GKE
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func importOceanGKECluster(resourceData *schema.ResourceData, meta interface{}) (*gcp.Cluster, error) {
	input := &gcp.ImportOceanGKEClusterInput{
		ClusterName: spotinst.String(resourceData.Get("cluster_name").(string)),
		Location:    spotinst.String(resourceData.Get("location").(string)),
	}

	resp, err := meta.(*Client).ocean.CloudProviderGCP().ImportOceanGKECluster(context.Background(), input)

	if err != nil {
		// If the group was not found, return nil so that we can show
		// that the group is gone.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeClusterNotFound {
					resourceData.SetId("")
					return nil, err
				}
			}
		}
		// Some other error, report it.
		return nil, fmt.Errorf("ocean GKE: import failed to read group: %s", err)
	}

	return resp.Cluster, err
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstClusterGKEImportCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OceanGKEImportResource.GetName())

	importedCluster, err := importOceanGKECluster(resourceData, meta.(*Client))
	if err != nil {
		return err
	}

	cluster, err := commons.OceanGKEImportResource.OnCreate(importedCluster, resourceData, meta)
	if err != nil {
		return err
	}

	clusterId, err := createGKEImportedCluster(cluster, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(clusterId))

	log.Printf("===> GKE imported cluster created successfully: %s <===", resourceData.Id())
	return resourceSpotinstClusterGKEImportRead(resourceData, meta)
}

func createGKEImportedCluster(cluster *gcp.Cluster, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(cluster); err != nil {
		return nil, err
	} else {
		log.Printf("===> Cluster GKE import configuration: %s", json)
	}

	input := &gcp.CreateClusterInput{Cluster: cluster}

	var resp *gcp.CreateClusterOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.ocean.CloudProviderGCP().CreateCluster(context.Background(), input)
		if err != nil {

			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create GKE imported cluster: %s", err)
	}
	return resp.Cluster.ID, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstClusterGKEImportRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.OceanGKEImportResource.GetName(), id)

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
		return fmt.Errorf("failed to read GKE cluster: %s", err)
	}

	// if nothing was found, return no state
	clusterResponse := resp.Cluster
	if clusterResponse == nil {
		resourceData.SetId("")
		return nil
	}

	// Expose the controller cluster identifier.
	if clusterResponse.ControllerClusterID != nil {
		if err := resourceData.Set("cluster_controller_id", *clusterResponse.ControllerClusterID); err != nil {
			log.Printf("[ERROR] Failed to set cluster_controller_id")
		}
	}

	if err := commons.OceanGKEImportResource.OnRead(clusterResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> GKE cluster read successfully: %s <===", id)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstClusterGKEImportUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.OceanGKEImportResource.GetName(), id)

	shouldUpdate, cluster, err := commons.OceanGKEImportResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		cluster.SetId(spotinst.String(id))
		if err := updateGKEImportCluster(cluster, resourceData, meta); err != nil {
			return err
		}
	}
	log.Printf("===> GLE Cluster updated successfully: %s <===", id)
	return resourceSpotinstClusterGKEImportRead(resourceData, meta)
}

func updateGKEImportCluster(cluster *gcp.Cluster, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &gcp.UpdateClusterInput{
		Cluster: cluster,
	}

	clusterId := resourceData.Id()

	if json, err := commons.ToJson(cluster); err != nil {
		return err
	} else {
		log.Printf("===> GKE Cluster update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderGCP().UpdateCluster(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update GKE cluster [%v]: %v", clusterId, err)
	}

	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstClusterGKEImportDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanGKEImportResource.GetName(), id)

	if err := deleteGKEImportCluster(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> GKE Cluster deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteGKEImportCluster(resourceData *schema.ResourceData, meta interface{}) error {
	clusterId := resourceData.Id()
	input := &gcp.DeleteClusterInput{
		ClusterID: spotinst.String(clusterId),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> GKE Cluster delete configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderGCP().DeleteCluster(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete GKE cluster: %s", err)
	}
	return nil
}

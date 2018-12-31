package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_aws"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_aws_auto_scaling"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_aws_instance_types"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_aws_launch_configuration"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_cluster_aws_strategy"
	"log"
	"strings"
	"time"
)

func resourceSpotinstOceanAWS() *schema.Resource {
	setupClusterAWSResource()

	return &schema.Resource{
		Create: resourceSpotinstClusterAWSCreate,
		Read:   resourceSpotinstClusterAWSRead,
		Update: resourceSpotinstClusterAWSUpdate,
		Delete: resourceSpotinstClusterAWSDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: commons.OceanResource.GetSchemaMap(),
	}
}

func setupClusterAWSResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	ocean_aws.Setup(fieldsMap)
	ocean_aws_auto_scaling.Setup(fieldsMap)
	ocean_aws_instance_types.Setup(fieldsMap)
	ocean_aws_launch_configuration.Setup(fieldsMap)
	ocean_cluster_aws_strategy.Setup(fieldsMap)

	commons.OceanResource = commons.NewOceanAWSResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstClusterAWSCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OceanResource.GetName())

	cluster, err := commons.OceanResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	clusterId, err := createCluster(cluster, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(clusterId))

	log.Printf("===> Elastigroup created successfully: %s <===", resourceData.Id())
	return resourceSpotinstClusterAWSRead(resourceData, meta)
}

func createCluster(cluster *aws.Cluster, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(cluster); err != nil {
		return nil, err
	} else {
		log.Printf("===> Cluster create configuration: %s", json)
	}

	input := &aws.CreateClusterInput{Cluster: cluster}

	var resp *aws.CreateClusterOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.ocean.CloudProviderAWS().CreateCluster(context.Background(), input)
		if err != nil {
			// Checks whether we should retry cluster creation.
			if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
				for _, err := range errs {
					if err.Code == "InvalidParamterValue" &&
						strings.Contains(err.Message, "Invalid IAM Instance Profile") {
						return resource.NonRetryableError(err)
					}
				}
			}
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
	return nil, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const ErrCodeClusterNotFound = "CLUSTER_DOESNT_EXIST"

func resourceSpotinstClusterAWSRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.OceanResource.GetName(), id)

	input := &aws.ReadClusterInput{ClusterID: spotinst.String(id)}
	resp, err := meta.(*Client).ocean.CloudProviderAWS().ReadCluster(context.Background(), input)

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
		return fmt.Errorf("failed to read cluster: %s", err)
	}

	// if nothing was found, return no state
	clusterResponse := resp.Cluster
	if clusterResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanResource.OnRead(clusterResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> Cluster read successfully: %s <===", id)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstClusterAWSUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.OceanResource.GetName(), id)

	shouldUpdate, cluster, err := commons.OceanResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		cluster.SetId(spotinst.String(id))
		if err := updateCluster(cluster, resourceData, meta); err != nil {
			return err
		}
	}
	log.Printf("===> Cluster updated successfully: %s <===", id)
	return resourceSpotinstClusterAWSRead(resourceData, meta)
}

func updateCluster(cluster *aws.Cluster, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &aws.UpdateClusterInput{
		Cluster: cluster,
	}

	clusterId := resourceData.Id()

	if json, err := commons.ToJson(cluster); err != nil {
		return err
	} else {
		log.Printf("===> Cluster update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().UpdateCluster(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update cluster [%v]: %v", clusterId, err)
	}

	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstClusterAWSDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanResource.GetName(), id)

	if err := deleteCluster(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Cluster deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteCluster(resourceData *schema.ResourceData, meta interface{}) error {
	clusterId := resourceData.Id()
	input := &aws.DeleteClusterInput{
		ClusterID: spotinst.String(clusterId),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Cluster delete configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().DeleteCluster(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete cluster: %s", err)
	}
	return nil
}

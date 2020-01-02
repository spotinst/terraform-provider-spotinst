package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_aws"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_aws_auto_scaling"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_aws_instance_types"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_aws_launch_configuration"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_aws_strategy"
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
		Schema: commons.OceanAWSResource.GetSchemaMap(),
	}
}

func setupClusterAWSResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	ocean_aws.Setup(fieldsMap)
	ocean_aws_auto_scaling.Setup(fieldsMap)
	ocean_aws_instance_types.Setup(fieldsMap)
	ocean_aws_launch_configuration.Setup(fieldsMap)
	ocean_aws_strategy.Setup(fieldsMap)

	commons.OceanAWSResource = commons.NewOceanAWSResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstClusterAWSCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OceanAWSResource.GetName())

	cluster, err := commons.OceanAWSResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	clusterId, err := createAWSCluster(resourceData, cluster, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(clusterId))

	log.Printf("===> Cluster created successfully: %s <===", resourceData.Id())
	return resourceSpotinstClusterAWSRead(resourceData, meta)
}

func createAWSCluster(resourceData *schema.ResourceData, cluster *aws.Cluster, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(cluster); err != nil {
		return nil, err
	} else {
		log.Printf("===> Cluster create configuration: %s", json)
	}

	input := &aws.CreateClusterInput{Cluster: cluster}
	if v, ok := resourceData.Get(string(ocean_aws_launch_configuration.IAMInstanceProfile)).(string); ok && v != "" {
		// Wait for IAM instance profile to be ready.
		time.Sleep(10 * time.Second)
	}
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
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const ErrCodeClusterNotFound = "CLUSTER_DOESNT_EXIST"

func resourceSpotinstClusterAWSRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.OceanAWSResource.GetName(), id)

	input := &aws.ReadClusterInput{ClusterID: spotinst.String(id)}
	resp, err := meta.(*Client).ocean.CloudProviderAWS().ReadCluster(context.Background(), input)

	if err != nil {
		// If the cluster was not found, return nil so that we can show
		// that the cluster does not exist
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

	if err := commons.OceanAWSResource.OnRead(clusterResponse, resourceData, meta); err != nil {
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
		commons.OceanAWSResource.GetName(), id)

	shouldUpdate, cluster, err := commons.OceanAWSResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		cluster.SetId(spotinst.String(id))
		if err := updateAWSCluster(cluster, resourceData, meta); err != nil {
			return err
		}
	}
	log.Printf("===> Cluster updated successfully: %s <===", id)
	return resourceSpotinstClusterAWSRead(resourceData, meta)
}

func updateAWSCluster(cluster *aws.Cluster, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &aws.UpdateClusterInput{
		Cluster: cluster,
	}

	var shouldRoll = false
	clusterId := resourceData.Id()
	if updatePolicy, exists := resourceData.GetOkExists(string(ocean_aws.UpdatePolicy)); exists {
		list := updatePolicy.([]interface{})
		if len(list) > 0 && list[0] != nil {
			m := list[0].(map[string]interface{})

			if roll, ok := m[string(ocean_aws.ShouldRoll)].(bool); ok && roll {
				shouldRoll = roll
			}
		}
	}

	if json, err := commons.ToJson(cluster); err != nil {
		return err
	} else {
		log.Printf("===> Cluster update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().UpdateCluster(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update cluster [%v]: %v", clusterId, err)
	} else if shouldRoll {
		if err := rollCluster(resourceData, meta); err != nil {
			log.Printf("[ERROR] Cluster [%v] roll failed, error: %v", clusterId, err)
			return err
		}
	} else {
		log.Printf("onRoll() -> Field [%v] is false, skipping cluster roll", string(ocean_aws.ShouldRoll))
	}

	return nil
}

func rollCluster(resourceData *schema.ResourceData, meta interface{}) error {
	var errResult error = nil
	clusterId := resourceData.Id()

	if updatePolicy, exists := resourceData.GetOkExists(string(ocean_aws.UpdatePolicy)); exists {
		list := updatePolicy.([]interface{})
		if len(list) > 0 && list[0] != nil {
			updateClusterSchema := list[0].(map[string]interface{})
			if rollConfig, ok := updateClusterSchema[string(ocean_aws.RollConfig)]; !ok || rollConfig == nil {
				errResult = fmt.Errorf("[ERROR] onRoll() -> Field [%v] is missing, skipping roll for cluster [%v]", string(ocean_aws.RollConfig), clusterId)
			} else {
				if rollClusterInput, err := expandOceanRollConfig(rollConfig, spotinst.String(clusterId)); err != nil {
					errResult = fmt.Errorf("[ERROR] onRoll() -> Failed expanding roll configuration for cluster [%v], error: %v", clusterId, err)
				} else {
					if json, err := commons.ToJson(rollConfig); err != nil {
						return err
					} else {
						log.Printf("onRoll() -> Rolling cluster [%v] with configuration %s", clusterId, json)
						rollClusterInput.Roll.ClusterID = spotinst.String(clusterId)
						_, err := meta.(*Client).ocean.CloudProviderAWS().Roll(context.Background(), rollClusterInput)
						if err != nil {
							return fmt.Errorf("onRoll() -> Roll failed for cluster [%v], error: %v", clusterId, err)
						} else {
							log.Printf("onRoll() -> Successfully rolled cluster [%v]", clusterId)
						}
					}
				}
			}
		}
	} else {
		errResult = fmt.Errorf("[ERROR] onRoll() -> Missing update policy for cluster [%v]", clusterId)
	}

	if errResult != nil {
		return errResult
	}
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstClusterAWSDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanAWSResource.GetName(), id)

	if err := deleteAWSCluster(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Cluster deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteAWSCluster(resourceData *schema.ResourceData, meta interface{}) error {
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

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func expandOceanRollConfig(data interface{}, clusterId *string) (*aws.RollClusterInput, error) {
	i := &aws.RollClusterInput{Roll: &aws.Roll{ClusterID: clusterId}}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(ocean_aws.BatchSizePercentage)].(int); ok {
			i.Roll.BatchSizePercentage = spotinst.Int(v)
		}

	}
	return i, nil
}

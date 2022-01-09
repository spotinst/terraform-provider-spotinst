package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_ecs_optimize_images"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aws"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_ecs"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_ecs_autoscaler"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_ecs_instance_types"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_ecs_launch_specification"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_ecs_scheduling"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_ecs_strategy"
)

func resourceSpotinstOceanECS() *schema.Resource {
	setupClusterECSResource()

	return &schema.Resource{
		Create: resourceSpotinstClusterECSCreate,
		Read:   resourceSpotinstClusterECSRead,
		Update: resourceSpotinstClusterECSUpdate,
		Delete: resourceSpotinstClusterECSDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: commons.OceanECSResource.GetSchemaMap(),
	}
}

func setupClusterECSResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)
	ocean_ecs.Setup(fieldsMap)
	ocean_ecs_instance_types.Setup(fieldsMap)
	ocean_ecs_launch_specification.Setup(fieldsMap)
	ocean_ecs_autoscaler.Setup(fieldsMap)
	ocean_ecs_strategy.Setup(fieldsMap)
	ocean_ecs_scheduling.Setup(fieldsMap)
	ocean_ecs_optimize_images.Setup(fieldsMap)

	commons.OceanECSResource = commons.NewOceanECSResource(fieldsMap)
}

func resourceSpotinstClusterECSCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OceanECSResource.GetName())

	cluster, err := commons.OceanECSResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	clusterID, err := createECSCluster(resourceData, cluster, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(clusterID))

	log.Printf("===> Cluster created successfully: %s <===", resourceData.Id())
	return resourceSpotinstClusterECSRead(resourceData, meta)
}

func createECSCluster(resourceData *schema.ResourceData, cluster *aws.ECSCluster, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(cluster); err != nil {
		return nil, err
	} else {
		log.Printf("===> Cluster create configuration: %s", json)
	}

	if v, ok := resourceData.Get(string(ocean_ecs_launch_specification.IamInstanceProfile)).(string); ok && v != "" {
		// Wait for IAM instance profile to be ready.
		time.Sleep(10 * time.Second)
	}

	var resp *aws.CreateECSClusterOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		input := &aws.CreateECSClusterInput{Cluster: cluster}
		r, err := spotinstClient.ocean.CloudProviderAWS().CreateECSCluster(context.Background(), input)
		if err != nil {
			// Checks whether we should retry cluster creation.
			if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
				for _, err := range errs {
					if err.Code == "InvalidParameterValue" &&
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

const ErrCodeECSClusterNotFound = "CANT_GET_OCEAN_CLUSTER"

func resourceSpotinstClusterECSRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.OceanECSResource.GetName(), id)

	input := &aws.ReadECSClusterInput{ClusterID: spotinst.String(id)}
	resp, err := meta.(*Client).ocean.CloudProviderAWS().ReadECSCluster(context.Background(), input)

	if err != nil {
		// If the cluster was not found, return nil so that we can show
		// that the cluster does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeECSClusterNotFound {
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

	if err := commons.OceanECSResource.OnRead(clusterResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> Cluster read successfully: %s <===", id)
	return nil
}

func resourceSpotinstClusterECSUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.OceanAWSResource.GetName(), id)

	shouldUpdate, changesRequiredRoll, tagsChanged, cluster, err := commons.OceanECSResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		cluster.SetId(spotinst.String(id))
		if err := updateECSCluster(cluster, resourceData, meta, changesRequiredRoll, tagsChanged); err != nil {
			return err
		}
	}
	log.Printf("===> Cluster updated successfully: %s <===", id)
	return resourceSpotinstClusterECSRead(resourceData, meta)
}

func updateECSCluster(cluster *aws.ECSCluster, resourceData *schema.ResourceData, meta interface{}, changesRequiredRoll bool, tagsChanged bool) error {
	var input = &aws.UpdateECSClusterInput{
		Cluster: cluster,
	}

	var shouldRoll = false
	var conditionedRoll = false
	var autoApplyTags = false
	clusterID := resourceData.Id()
	if updatePolicy, exists := resourceData.GetOkExists(string(ocean_ecs.UpdatePolicy)); exists {
		list := updatePolicy.([]interface{})
		if len(list) > 0 && list[0] != nil {
			m := list[0].(map[string]interface{})

			if roll, ok := m[string(ocean_ecs.ShouldRoll)].(bool); ok && roll {
				shouldRoll = roll
			}

			if condRoll, ok := m[string(ocean_ecs.ConditionedRoll)].(bool); ok && condRoll {
				conditionedRoll = condRoll
			}

			if aat, ok := m[string(ocean_ecs.AutoApplyTags)].(bool); ok && aat {
				autoApplyTags = aat
			}
		}
	}

	if json, err := commons.ToJson(cluster); err != nil {
		return err
	} else {
		log.Printf("===> Cluster update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().UpdateECSCluster(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update cluster [%v]: %v", clusterID, err)
	} else if shouldRoll {
		if !conditionedRoll || changesRequiredRoll || (!autoApplyTags && tagsChanged) {
			if err := rollECSCluster(resourceData, meta); err != nil {
				log.Printf("[ERROR] Cluster [%v] roll failed, error: %v", clusterID, err)
				return err
			}
		}
	} else {
		log.Printf("onRoll() -> Field [%v] is false, skipping cluster roll", string(ocean_aws.ShouldRoll))
	}

	return nil
}

func rollECSCluster(resourceData *schema.ResourceData, meta interface{}) error {
	var errResult error = nil
	clusterID := resourceData.Id()

	if updatePolicy, exists := resourceData.GetOkExists(string(ocean_ecs.UpdatePolicy)); exists {
		list := updatePolicy.([]interface{})
		if len(list) > 0 && list[0] != nil {
			updateClusterSchema := list[0].(map[string]interface{})
			if rollConfig, ok := updateClusterSchema[string(ocean_ecs.RollConfig)]; !ok || rollConfig == nil {
				errResult = fmt.Errorf("[ERROR] onRoll() -> Field [%v] is missing, skipping roll for cluster [%v]", string(ocean_ecs.RollConfig), clusterID)
			} else {
				if rollClusterInput, err := expandECSOceanRollConfig(rollConfig, spotinst.String(clusterID)); err != nil {
					errResult = fmt.Errorf("[ERROR] onRoll() -> Failed expanding roll configuration for cluster [%v], error: %v", clusterID, err)
				} else {
					if json, err := commons.ToJson(rollConfig); err != nil {
						return err
					} else {
						log.Printf("onRoll() -> Rolling cluster [%v] with configuration %s", clusterID, json)
						rollClusterInput.Roll.ClusterID = spotinst.String(clusterID)
						_, err := meta.(*Client).ocean.CloudProviderAWS().RollECS(context.Background(), rollClusterInput)
						if err != nil {
							return fmt.Errorf("onRoll() -> Roll failed for cluster [%v], error: %v", clusterID, err)
						} else {
							log.Printf("onRoll() -> Successfully rolled cluster [%v]", clusterID)
						}
					}
				}
			}
		}
	} else {
		errResult = fmt.Errorf("[ERROR] onRoll() -> Missing update policy for cluster [%v]", clusterID)
	}

	if errResult != nil {
		return errResult
	}
	return nil
}

func resourceSpotinstClusterECSDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanECSResource.GetName(), id)

	if err := deleteECSCluster(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Cluster deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteECSCluster(resourceData *schema.ResourceData, meta interface{}) error {
	clusterID := resourceData.Id()
	input := &aws.DeleteECSClusterInput{
		ClusterID: spotinst.String(clusterID),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Cluster delete configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().DeleteECSCluster(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete cluster: %s", err)
	}
	return nil
}

func expandECSOceanRollConfig(data interface{}, clusterID *string) (*aws.ECSRollClusterInput, error) {
	i := &aws.ECSRollClusterInput{Roll: &aws.ECSRoll{ClusterID: clusterID}}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(ocean_ecs.BatchSizePercentage)].(int); ok {
			i.Roll.BatchSizePercentage = spotinst.Int(v)
		}

	}
	return i, nil
}

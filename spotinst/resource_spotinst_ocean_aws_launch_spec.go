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
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aws_launch_spec"
)

func resourceSpotinstOceanAWSLaunchSpec() *schema.Resource {
	setupOceanAWSLaunchSpecResource()

	return &schema.Resource{
		Create: resourceSpotinstOceanAWSLaunchSpecCreate,
		Read:   resourceSpotinstOceanAWSLaunchSpecRead,
		Update: resourceSpotinstOceanAWSLaunchSpecUpdate,
		Delete: resourceSpotinstOceanAWSLaunchSpecDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.OceanAWSLaunchSpecResource.GetSchemaMap(),
	}
}

func setupOceanAWSLaunchSpecResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)
	ocean_aws_launch_spec.Setup(fieldsMap)

	commons.OceanAWSLaunchSpecResource = commons.NewOceanAWSLaunchSpecResource(fieldsMap)
}

func resourceSpotinstOceanAWSLaunchSpecCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate), commons.OceanAWSLaunchSpecResource.GetName())

	launchSpec, err := commons.OceanAWSLaunchSpecResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	launchSpecId, err := createLaunchSpec(resourceData, launchSpec, meta.(*Client))
	if err != nil {
		return err
	}
	resourceData.SetId(spotinst.StringValue(launchSpecId))

	return resourceSpotinstOceanAWSLaunchSpecRead(resourceData, meta)
}

func createLaunchSpec(resourceData *schema.ResourceData, launchSpec *aws.LaunchSpec, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(launchSpec); err != nil {
		return nil, err
	} else {
		log.Printf("===> LaunchSpec create configuration: %s", json)
	}

	var resp *aws.CreateLaunchSpecOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		input := &aws.CreateLaunchSpecInput{LaunchSpec: launchSpec}
		if createOptions, exists := resourceData.GetOkExists(string(ocean_aws_launch_spec.CreateOptions)); exists {
			list := createOptions.([]interface{})
			if len(list) > 0 && list[0] != nil {
				m := list[0].(map[string]interface{})
				if initialNodes, ok := m[string(ocean_aws_launch_spec.InitialNodes)].(int); ok && initialNodes > 0 {
					input.InitialNodes = spotinst.Int(initialNodes)
				}
			}
		}
		r, err := spotinstClient.ocean.CloudProviderAWS().CreateLaunchSpec(context.Background(), input)
		if err != nil {
			// Checks whether we should retry launchSpec creation.
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
		return nil, fmt.Errorf("[ERROR] failed to create launchSpec: %s", err)
	}
	return resp.LaunchSpec.ID, nil
}

const ErrCodeLaunchSpecNotFound = "CANT_GET_OCEAN_LAUNCH_SPEC"

func resourceSpotinstOceanAWSLaunchSpecRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanAWSLaunchSpecResource.GetName(), id)

	input := &aws.ReadLaunchSpecInput{LaunchSpecID: spotinst.String(id)}
	resp, err := meta.(*Client).ocean.CloudProviderAWS().ReadLaunchSpec(context.Background(), input)

	if err != nil {
		// If the launchSpec was not found, return nil so that we can show
		// that it does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeLaunchSpecNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return fmt.Errorf("failed to read launchSpec: %s", err)
	}

	// if nothing was found, return no state
	launchSpecResponse := resp.LaunchSpec
	if launchSpecResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanAWSLaunchSpecResource.OnRead(launchSpecResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> launchSpec read successfully: %s <===", id)
	return nil
}

func resourceSpotinstOceanAWSLaunchSpecUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.OceanAWSLaunchSpecResource.GetName(), id)

	shouldUpdate, launchSpec, err := commons.OceanAWSLaunchSpecResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		launchSpec.SetId(spotinst.String(id))
		if err := updateLaunchSpec(launchSpec, resourceData, meta); err != nil {
			return err
		}
	}
	log.Printf("===> launchSpec updated successfully: %s <===", id)
	return resourceSpotinstOceanAWSLaunchSpecRead(resourceData, meta)
}

func updateLaunchSpec(launchSpec *aws.LaunchSpec, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &aws.UpdateLaunchSpecInput{
		LaunchSpec: launchSpec,
	}

	launchSpecId := resourceData.Id()
	oceanId := resourceData.Get(string(ocean_aws_launch_spec.OceanID))
	var shouldRoll = false
	if updatePolicy, exists := resourceData.GetOkExists(string(ocean_aws_launch_spec.UpdatePolicy)); exists {
		list := updatePolicy.([]interface{})
		if len(list) > 0 && list[0] != nil {
			m := list[0].(map[string]interface{})

			if roll, ok := m[string(ocean_aws_launch_spec.ShouldRoll)].(bool); ok && roll {
				shouldRoll = roll
			}
		}
	}

	if json, err := commons.ToJson(launchSpec); err != nil {
		return err
	} else {
		log.Printf("===> launchSpec update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().UpdateLaunchSpec(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update launchSpec [%v]: %v", launchSpecId, err)
	} else if shouldRoll {
		if err := rollLaunchSpecAwsCluster(resourceData, meta); err != nil {
			log.Printf("[ERROR] Cluster [%v] roll failed, error: %v", oceanId, err)
			return err
		}
	} else {
		log.Printf("onRoll() -> Field [%v] is false, skipping cluster roll", string(ocean_aws_launch_spec.ShouldRoll))
	}

	return nil
}

func rollLaunchSpecAwsCluster(resourceData *schema.ResourceData, meta interface{}) error {
	var errResult error = nil
	launchSpecId := resourceData.Id()
	clusterID := resourceData.Get(string(ocean_aws_launch_spec.OceanID)).(string)

	if updatePolicy, exists := resourceData.GetOkExists(string(ocean_aws_launch_spec.UpdatePolicy)); exists {
		list := updatePolicy.([]interface{})
		if len(list) > 0 && list[0] != nil {
			updateClusterSchema := list[0].(map[string]interface{})
			if rollConfig, ok := updateClusterSchema[string(ocean_aws_launch_spec.RollConfig)]; !ok || rollConfig == nil {
				errResult = fmt.Errorf("[ERROR] onRoll() -> Field [%v] is missing, skipping roll for cluster [%v]", string(ocean_aws_launch_spec.RollConfig), clusterID)
			} else {
				if rollClusterInput, err := expandLaunchSpecAWSRollConfig(rollConfig, &clusterID, launchSpecId); err != nil {
					errResult = fmt.Errorf("[ERROR] onRoll() -> Failed expanding roll configuration for cluster [%v], error: %v", clusterID, err)
				} else {
					if json, err := commons.ToJson(rollConfig); err != nil {
						return err
					} else {
						log.Printf("onRoll() -> Rolling cluster [%v] with configuration %s", clusterID, json)
						rollClusterInput.Roll.ClusterID = &clusterID
						_, err := meta.(*Client).ocean.CloudProviderAWS().CreateRoll(context.Background(), rollClusterInput)
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

func resourceSpotinstOceanAWSLaunchSpecDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanAWSLaunchSpecResource.GetName(), id)

	if err := deleteLaunchSpec(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> launchSpec deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteLaunchSpec(resourceData *schema.ResourceData, meta interface{}) error {
	launchSpecId := resourceData.Id()
	input := &aws.DeleteLaunchSpecInput{
		LaunchSpecID: spotinst.String(launchSpecId),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> launchSpec delete configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().DeleteLaunchSpec(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete launchSpecId: %s", err)
	}
	return nil
}

func expandLaunchSpecAWSRollConfig(data interface{}, clusterID *string, launchSpecID string) (*aws.CreateRollInput, error) {
	i := &aws.CreateRollInput{Roll: &aws.RollSpec{ClusterID: clusterID}}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(ocean_aws_launch_spec.BatchSizePercentage)].(int); ok {
			i.Roll.BatchSizePercentage = spotinst.Int(v)
		}

		lsResult := make([]string, 0, 1)
		lsResult = append(lsResult, launchSpecID)

		i.Roll.LaunchSpecIDs = lsResult

	}
	return i, nil
}

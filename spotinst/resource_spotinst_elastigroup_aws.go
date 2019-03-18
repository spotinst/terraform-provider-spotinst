package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws_block_devices"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws_instance_types"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws_integrations"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws_launch_configuration"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws_network_interface"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws_scaling_policies"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws_scheduled_task"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws_stateful"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws_strategy"
)

func resourceSpotinstElastigroupAws() *schema.Resource {
	setupElastigroupResource()

	return &schema.Resource{
		Create: resourceSpotinstElastigroupAwsCreate,
		Read:   resourceSpotinstElastigroupAwsRead,
		Update: resourceSpotinstElastigroupAwsUpdate,
		Delete: resourceSpotinstElastigroupAwsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.ElastigroupResource.GetSchemaMap(),
	}
}

func setupElastigroupResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	elastigroup_aws.Setup(fieldsMap)
	elastigroup_aws_block_devices.Setup(fieldsMap)
	elastigroup_aws_instance_types.Setup(fieldsMap)
	elastigroup_aws_integrations.Setup(fieldsMap)
	elastigroup_aws_launch_configuration.Setup(fieldsMap)
	elastigroup_aws_network_interface.Setup(fieldsMap)
	elastigroup_aws_scaling_policies.Setup(fieldsMap)
	elastigroup_aws_scheduled_task.Setup(fieldsMap)
	elastigroup_aws_stateful.Setup(fieldsMap)
	elastigroup_aws_strategy.Setup(fieldsMap)

	commons.ElastigroupResource = commons.NewElastigroupResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstElastigroupAwsDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.ElastigroupResource.GetName(), id)

	if err := deleteGroup(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Elastigroup deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteGroup(resourceData *schema.ResourceData, meta interface{}) error {
	groupId := resourceData.Id()
	input := &aws.DeleteGroupInput{
		GroupID: spotinst.String(groupId),
	}

	if statefulDeallocation, exists := resourceData.GetOkExists(string(elastigroup_aws_stateful.StatefulDeallocation)); exists {
		list := statefulDeallocation.([]interface{})
		if list != nil && len(list) > 0 && list[0] != nil {
			m := list[0].(map[string]interface{})

			var result = &aws.StatefulDeallocation{}
			if shouldDeleteImage, ok := m[string(elastigroup_aws_stateful.ShouldDeleteImages)].(bool); ok && shouldDeleteImage {
				result.ShouldDeleteImages = spotinst.Bool(shouldDeleteImage)
			}

			if shouldDeleteNetworkInterfaces, ok := m[string(elastigroup_aws_stateful.ShouldDeleteNetworkInterfaces)].(bool); ok && shouldDeleteNetworkInterfaces {
				result.ShouldDeleteNetworkInterfaces = spotinst.Bool(shouldDeleteNetworkInterfaces)
			}

			if shouldDeleteSnapshots, ok := m[string(elastigroup_aws_stateful.ShouldDeleteSnapshots)].(bool); ok && shouldDeleteSnapshots {
				result.ShouldDeleteSnapshots = spotinst.Bool(shouldDeleteSnapshots)
			}

			if shouldDeleteVolumes, ok := m[string(elastigroup_aws_stateful.ShouldDeleteVolumes)].(bool); ok && shouldDeleteVolumes {
				result.ShouldDeleteVolumes = spotinst.Bool(shouldDeleteVolumes)
			}

			input.StatefulDeallocation = result
		}
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Group delete configuration: %s", json)
	}

	if _, err := meta.(*Client).elastigroup.CloudProviderAWS().Delete(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete group: %s", err)
	}
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
// ErrCodeGroupNotFound for service response error code "GROUP_DOESNT_EXIST".
const ErrCodeGroupNotFound = "GROUP_DOESNT_EXIST"

func resourceSpotinstElastigroupAwsRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.ElastigroupResource.GetName(), id)

	input := &aws.ReadGroupInput{GroupID: spotinst.String(id)}
	resp, err := meta.(*Client).elastigroup.CloudProviderAWS().Read(context.Background(), input)
	if err != nil {
		// If the group was not found, return nil so that we can show
		// that the group does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGroupNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return fmt.Errorf("failed to read group: %s", err)
	}

	// If nothing was found, then return no state.
	groupResponse := resp.Group
	if groupResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.ElastigroupResource.OnRead(groupResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> Elastigroup read successfully: %s <===", id)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstElastigroupAwsCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.ElastigroupResource.GetName())

	elastigroup, err := commons.ElastigroupResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	groupId, err := createGroup(elastigroup, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(groupId))

	if capacity, ok := resourceData.GetOkExists(string(elastigroup_aws.WaitForCapacity)); ok {
		if *elastigroup.Capacity.Target < capacity.(int) {

			return fmt.Errorf("[ERROR] Your target healthy capacity must be less than or equal to your desired capcity")
		}
		if timeout, ok := resourceData.GetOkExists(string(elastigroup_aws.WaitForCapacityTimeout)); ok {
			err := awaitReady(groupId, timeout.(int), capacity.(int), meta.(*Client))
			if err != nil {
				return fmt.Errorf("[ERROR] Timed out when creating group: %s", err)
			}
		}
	}

	log.Printf("===> Elastigroup created successfully: %s <===", resourceData.Id())

	return resourceSpotinstElastigroupAwsRead(resourceData, meta)
}

func createGroup(group *aws.Group, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(group); err != nil {
		return nil, err
	} else {
		log.Printf("===> Group create configuration: %s", json)
	}

	input := &aws.CreateGroupInput{Group: group}

	var resp *aws.CreateGroupOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.elastigroup.CloudProviderAWS().Create(context.Background(), input)
		if err != nil {
			// Checks whether we should retry the group creation.
			if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
				for _, err := range errs {
					if err.Code == "InvalidParameterValue" &&
						strings.Contains(err.Message, "Invalid IAM Instance Profile") {
						return resource.RetryableError(err)
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
		return nil, fmt.Errorf("[ERROR] failed to create group: %s", err)
	}
	return resp.Group.ID, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstElastigroupAwsUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.ElastigroupResource.GetName(), id)

	shouldUpdate, elastigroup, err := commons.ElastigroupResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		elastigroup.SetId(spotinst.String(id))
		if err := updateGroup(elastigroup, resourceData, meta); err != nil {
			return err
		}
	}

	log.Printf("===> Elastigroup updated successfully: %s <===", id)
	return resourceSpotinstElastigroupAwsRead(resourceData, meta)
}

func updateGroup(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &aws.UpdateGroupInput{
		Group: elastigroup,
	}

	var shouldRoll = false
	groupId := resourceData.Id()
	if updatePolicy, exists := resourceData.GetOkExists(string(elastigroup_aws.UpdatePolicy)); exists {
		list := updatePolicy.([]interface{})
		if list != nil && len(list) > 0 && list[0] != nil {
			m := list[0].(map[string]interface{})
			if resumeStateful, ok := m[string(elastigroup_aws.ShouldResumeStateful)].(bool); ok && resumeStateful {
				log.Printf("Resuming paused stateful instances on group [%v]...", groupId)
				input.ShouldResumeStateful = spotinst.Bool(resumeStateful)
			}

			if autoApplyTags, ok := m[string(elastigroup_aws.AutoApplyTags)].(bool); ok && autoApplyTags {
				log.Printf("Updating tags without rolling the group")
				input.AutoApplyTags = spotinst.Bool(autoApplyTags)
			}

			if roll, ok := m[string(elastigroup_aws.ShouldRoll)].(bool); ok && roll {
				shouldRoll = roll
			}

		}
	}

	if json, err := commons.ToJson(elastigroup); err != nil {
		return err
	} else {
		log.Printf("===> Group update configuration: %s", json)
	}

	if _, err := meta.(*Client).elastigroup.CloudProviderAWS().Update(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update group [%v]: %v", groupId, err)
	} else if shouldRoll {
		if err := rollGroup(resourceData, meta); err != nil {
			log.Printf("[ERROR] Group [%v] roll failed, error: %v", groupId, err)
			return err
		}
	} else {
		log.Printf("onRoll() -> Field [%v] is false, skipping group roll", string(elastigroup_aws.ShouldRoll))
		if capacity, ok := resourceData.GetOkExists(string(elastigroup_aws.WaitForCapacity)); ok {
			if target, ok := resourceData.GetOkExists(string(elastigroup_aws.DesiredCapacity)); ok {
				if target.(int) < capacity.(int) {
					return fmt.Errorf("[ERROR] You've asked to wait for a healthy capacity that is above your desired capacity")
				}

				if timeout, ok := resourceData.GetOkExists(string(elastigroup_aws.WaitForCapacityTimeout)); ok {
					err := awaitReady(spotinst.String(groupId), timeout.(int), capacity.(int), meta.(*Client))
					if err != nil {
						return fmt.Errorf("[ERROR] Timed out when updating group: %s", err)
					}
				}
			}
		}
	}
	return nil
}

func rollGroup(resourceData *schema.ResourceData, meta interface{}) error {
	var errResult error = nil
	groupId := resourceData.Id()

	if updatePolicy, exists := resourceData.GetOkExists(string(elastigroup_aws.UpdatePolicy)); exists {
		list := updatePolicy.([]interface{})
		if list != nil && len(list) > 0 && list[0] != nil {
			updateGroupSchema := list[0].(map[string]interface{})
			if rollConfig, ok := updateGroupSchema[string(elastigroup_aws.RollConfig)]; !ok || rollConfig == nil {
				errResult = fmt.Errorf("[ERROR] onRoll() -> Field [%v] is missing, skipping roll for group [%v]", string(elastigroup_aws.RollConfig), groupId)
			} else {
				if rollGroupInput, err := expandElastigroupRollConfig(rollConfig, spotinst.String(groupId)); err != nil {
					errResult = fmt.Errorf("[ERROR] onRoll() -> Failed expanding roll configuration for group [%v], error: %v", groupId, err)
				} else {
					if json, err := commons.ToJson(rollConfig); err != nil {
						return err
					} else {
						log.Printf("onRoll() -> Rolling group [%v] with configuration %s", groupId, json)
						// we want the outer retry timeout to equal the inner retry timeout, or 5 mins if undefined
						rto := spotinst.IntValue(getRollTimeout(rollConfig))
						if rto == 0 {
							rto = 300
						}
						errResult = resource.Retry(time.Duration(rto)*time.Second, func() *resource.RetryError {
							rollGroupInput.GroupID = spotinst.String(groupId)
							rollOut, err := meta.(*Client).elastigroup.CloudProviderAWS().Roll(context.Background(), rollGroupInput)
							if err != nil {

								// checks whether to retry role
								if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
									for _, err := range errs {
										if strings.Contains(err.Code, "CANT_ROLL_CAPACITY_BELOW_MINIMUM") {
											time.Sleep(time.Minute)
											return resource.RetryableError(err)
										}
									}
								}
								// Some other error, report it.
								return resource.NonRetryableError(err)
							}

							awaitErr := awaitReadyRoll(groupId, rollConfig, rollOut, meta.(*Client))
							if awaitErr != nil {
								waitErr := fmt.Errorf("[ERROR] Timed out when waiting for minimum roll %%: %s", err)
								return resource.NonRetryableError(waitErr)
							} else {
								log.Printf("onRoll() -> Successfully rolled group [%v]", groupId)
							}
							return nil
						})
					}
				}
			}
		}
	} else {
		errResult = fmt.Errorf("[ERROR] onRoll() -> Missing update policy for group [%v]", groupId)
	}

	if errResult != nil {
		return errResult
	}
	return nil
}

func awaitReady(groupId *string, timeout int, capacity int, client *Client) error {
	if capacity == 0 || timeout == 0 {
		return nil
	}
	input := &aws.GetInstanceHealthinessInput{GroupID: spotinst.String(*groupId)}
	err := resource.Retry(time.Second*time.Duration(timeout), func() *resource.RetryError {
		numHealthy := 0
		status, err := client.elastigroup.CloudProviderAWS().GetInstanceHealthiness(context.Background(), input)
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("[ERROR] awaitReady() -> getInstanceHealthiness [%v] API call failed, error: %v", groupId, err))
		}

		for _, item := range status.Instances {
			if *item.HealthStatus == "HEALTHY" {
				numHealthy += 1
			}
		}

		if numHealthy < capacity {
			log.Printf("===> waiting for %d more healthy instances <===\n", capacity-numHealthy)
			err = fmt.Errorf("===> waiting for %d more healthy instances <===", capacity-numHealthy)
			return resource.RetryableError(err)
		}

		log.Printf("awaitReady() -> Target number of health instances reached [%v]", *groupId)
		return nil
	})

	if err != nil {
		return fmt.Errorf("[ERROR] Instances not ready: %s", err)
	}

	return nil
}

func awaitReadyRoll(groupId string, rollConfig interface{}, rollOut *aws.RollGroupOutput, client *Client) error {
	pctTimeout := spotinst.IntValue(getRollTimeout(rollConfig))
	pctComplete := spotinst.IntValue(getRollMinPct(rollConfig))
	rollId := spotinst.StringValue(getRollStatus(rollOut))

	if pctTimeout > 0 && pctComplete > 0 {
		if rollId != "" {
			deployStatusInput := &aws.DeploymentStatusInput{GroupID: spotinst.String(groupId), RollID: spotinst.String(rollId)}
			err := resource.Retry(time.Second*time.Duration(pctTimeout), func() *resource.RetryError {
				if rollStatus, err := client.elastigroup.CloudProviderAWS().DeploymentStatus(context.Background(), deployStatusInput); err != nil {
					return resource.NonRetryableError(fmt.Errorf("[ERROR] awaitReadyRoll() -> Roll group status [%v] API call failed, error: %v", groupId, err))
				} else {
					if spotinst.IntValue(rollStatus.RollGroupStatus[0].Progress.Value) < pctComplete {
						log.Printf("===> waiting for at least %d%% of batches to complete, currently %d%% <===\n", pctComplete, spotinst.IntValue(rollStatus.RollGroupStatus[0].Progress.Value))
						return resource.RetryableError(fmt.Errorf("===> roll at %v%% complete <===", spotinst.IntValue(rollStatus.RollGroupStatus[0].Progress.Value)))
					}
				}
				return nil
			})
			if err != nil {
				return fmt.Errorf("[ERROR] Did not reach target deployment amount. Message: %s", err)
			}
			log.Printf("awaitReadyRoll() -> Target deployment percentage reached [%v]", groupId)
		}
	}
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Fields Expand
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandElastigroupRollConfig(data interface{}, groupID *string) (*aws.RollGroupInput, error) {
	i := &aws.RollGroupInput{GroupID: groupID}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(elastigroup_aws.BatchSizePercentage)].(int); ok { // Required value
			i.BatchSizePercentage = spotinst.Int(v)
		}

		if v, ok := m[string(elastigroup_aws.GracePeriod)].(int); ok && v != -1 { // Default value set to -1
			i.GracePeriod = spotinst.Int(v)
		}

		if v, ok := m[string(elastigroup_aws.HealthCheckType)].(string); ok && v != "" { // Default value ""
			i.HealthCheckType = spotinst.String(v)
		}
	}
	return i, nil
}

func getRollTimeout(data interface{}) *int {
	var timeout *int
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(elastigroup_aws.WaitForRollTimeout)].(int); ok {
			timeout = spotinst.Int(v)
		}
	}
	return timeout
}

func getRollMinPct(data interface{}) *int {
	var minPct *int
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(elastigroup_aws.WaitForRollPct)].(int); ok {
			minPct = spotinst.Int(v)
		}
	}
	return minPct
}

func getRollStatus(rollOut *aws.RollGroupOutput) *string {
	for item := range rollOut.RollGroupStatus {
		rs := strings.ToUpper(spotinst.StringValue(rollOut.RollGroupStatus[item].RollStatus))
		if rs == "IN_PROGRESS" || rs == "STARTING" {
			return rollOut.RollGroupStatus[item].RollID
		}
	}
	return nil
}

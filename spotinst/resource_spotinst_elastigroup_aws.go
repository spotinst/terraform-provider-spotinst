package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/elastigroup_aws"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/elastigroup_aws_block_devices"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/elastigroup_aws_instance_types"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/elastigroup_aws_integrations"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/elastigroup_aws_launch_configuration"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/elastigroup_aws_network_interface"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/elastigroup_aws_scaling_policies"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/elastigroup_aws_scheduled_task"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/elastigroup_aws_stateful"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/elastigroup_aws_strategy"
)

func resourceSpotinstElastigroupAWS() *schema.Resource {
	setupElastigroupResource()

	return &schema.Resource{
		Create: resourceSpotinstElastigroupAWSCreate,
		Read:   resourceSpotinstElastigroupAWSRead,
		Update: resourceSpotinstElastigroupAWSUpdate,
		Delete: resourceSpotinstElastigroupAWSDelete,

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

func resourceSpotinstElastigroupAWSDelete(resourceData *schema.ResourceData, meta interface{}) error {
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
		if len(list) > 0 && list[0] != nil {
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

// ErrCodeGroupNotFound for service response error code "GROUP_DOESNT_EXIST".
const ErrCodeGroupNotFound = "GROUP_DOESNT_EXIST"

func resourceSpotinstElastigroupAWSRead(resourceData *schema.ResourceData, meta interface{}) error {
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

	updateCapitalSlice(resourceData, groupResponse)

	if err := commons.ElastigroupResource.OnRead(groupResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> Elastigroup read successfully: %s <===", id)
	return nil
}

func resourceSpotinstElastigroupAWSCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.ElastigroupResource.GetName())

	elastigroup, err := commons.ElastigroupResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	groupId, err := createGroup(resourceData, elastigroup, meta.(*Client))
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

	return resourceSpotinstElastigroupAWSRead(resourceData, meta)
}

func createGroup(resourceData *schema.ResourceData, group *aws.Group, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(group); err != nil {
		return nil, err
	} else {
		log.Printf("===> Group create configuration: %s", json)
	}

	if v, ok := resourceData.Get(string(elastigroup_aws_launch_configuration.IamInstanceProfile)).(string); ok && v != "" {
		// Wait for IAM instance profile to be ready.
		time.Sleep(10 * time.Second)
	}

	var resp *aws.CreateGroupOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		input := &aws.CreateGroupInput{Group: group}
		r, err := spotinstClient.elastigroup.CloudProviderAWS().Create(context.Background(), input)
		if err != nil {
			// Checks whether we should retry the group creation.
			if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
				for _, err := range errs {
					if err.Code == "InvalidParameterValue" &&
						strings.Contains(strings.ToLower(err.Message), "invalid iam instance profile") {
						return resource.RetryableError(err)
					}
					if err.Code == "CANT_CREATE_GROUP" &&
						strings.Contains(strings.ToLower(err.Message), "failed to create group") {
						return resource.RetryableError(err)
					}
					if err.Code == "CANT_VALIDATE_IMAGE" &&
						strings.Contains(strings.ToLower(err.Message), "can't validate ami") {
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

func resourceSpotinstElastigroupAWSUpdate(resourceData *schema.ResourceData, meta interface{}) error {
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
	return resourceSpotinstElastigroupAWSRead(resourceData, meta)
}

func updateGroup(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &aws.UpdateGroupInput{
		Group: elastigroup,
	}

	var shouldRoll = false
	groupId := resourceData.Id()
	if updatePolicy, exists := resourceData.GetOkExists(string(elastigroup_aws.UpdatePolicy)); exists {
		list := updatePolicy.([]interface{})
		if len(list) > 0 && list[0] != nil {
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

	if instanceActions, exists := resourceData.GetOkExists(string(elastigroup_aws_stateful.StatefulInstanceAction)); exists {
		actionList := instanceActions.([]interface{})
		if err := checkStatefulActionUniqueness(actionList); err != nil {
			log.Printf("[ERROR] Uniqueness check failed with error: %v", err)
			return err
		}

		ctx := context.TODO()
		svc := meta.(*Client).elastigroup.CloudProviderAWS()

		for _, action := range actionList {
			var (
				actionMap  = action.(map[string]interface{})
				actionType = actionMap[string(elastigroup_aws_stateful.ActionType)].(string)
				instanceID = actionMap[string(elastigroup_aws_stateful.StatefulInstanceID)].(string)
				err        error
			)
			switch strings.ToLower(actionType) {
			case "pause":
				err = pauseStatefulInstance(ctx, svc, resourceData.Id(), instanceID)
			case "resume":
				err = resumeStatefulInstance(ctx, svc, resourceData.Id(), instanceID)
			case "recycle":
				err = recycleStatefulInstance(ctx, svc, resourceData.Id(), instanceID)
			case "deallocate":
				err = deallocateStatefulInstance(ctx, svc, resourceData.Id(), instanceID)
			default:
				err = fmt.Errorf("unsupported action %q on instance %q", actionType, instanceID)
			}
			if err != nil {
				log.Printf("[ERROR] Stateful instance (%s) action failed with error: %v", instanceID, err)
				return err
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

func checkStatefulActionUniqueness(actionList []interface{}) error {
	seenIDs := make(map[string]struct{})
	for _, action := range actionList {
		actionMap := action.(map[string]interface{})
		instanceID := actionMap[string(elastigroup_aws_stateful.StatefulInstanceID)].(string)
		if _, seen := seenIDs[instanceID]; seen {
			return fmt.Errorf("multiple actions are not allowed for the same instance (%v)", instanceID)
		}
		seenIDs[instanceID] = struct{}{}
	}
	return nil
}

func pauseStatefulInstance(ctx context.Context, svc aws.Service, groupID, instanceID string) error {
	log.Printf("Pausing instance (%s)", instanceID)

	input := &aws.PauseStatefulInstanceInput{
		GroupID:            spotinst.String(groupID),
		StatefulInstanceID: spotinst.String(instanceID),
	}
	_, err := svc.PauseStatefulInstance(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to pause instance (%s): %v", instanceID, err)
	}

	log.Printf("Successfully paused instance (%s)", instanceID)
	return nil
}

func resumeStatefulInstance(ctx context.Context, svc aws.Service, groupID, instanceID string) error {
	log.Printf("Resuming instance (%s)", instanceID)

	input := &aws.ResumeStatefulInstanceInput{
		GroupID:            spotinst.String(groupID),
		StatefulInstanceID: spotinst.String(instanceID),
	}
	_, err := svc.ResumeStatefulInstance(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to resume instance (%s): %v", instanceID, err)
	}

	log.Printf("Successfully resumed instance (%s)", instanceID)
	return nil
}

func recycleStatefulInstance(ctx context.Context, svc aws.Service, groupID, instanceID string) error {
	log.Printf("Recycling instance (%s)", instanceID)

	input := &aws.RecycleStatefulInstanceInput{
		GroupID:            spotinst.String(groupID),
		StatefulInstanceID: spotinst.String(instanceID),
	}
	_, err := svc.RecycleStatefulInstance(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to recycle instance (%s): %v", instanceID, err)
	}

	log.Printf("Successfully recycled instance (%s)", instanceID)
	return nil
}

func deallocateStatefulInstance(ctx context.Context, svc aws.Service, groupID, instanceID string) error {
	log.Printf("Deallocating instance (%s)", instanceID)

	input := &aws.DeallocateStatefulInstanceInput{
		GroupID:            spotinst.String(groupID),
		StatefulInstanceID: spotinst.String(instanceID),
	}
	_, err := svc.DeallocateStatefulInstance(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to deallocate instance (%s): %v", instanceID, err)
	}

	log.Printf("Successfully deallocated instance (%s)", instanceID)
	return nil
}

func rollGroup(resourceData *schema.ResourceData, meta interface{}) error {
	ctx := context.Background()
	groupID := resourceData.Id()

	updatePolicy, exists := resourceData.GetOkExists(string(elastigroup_aws.UpdatePolicy))
	if !exists {
		return fmt.Errorf("[ERROR] onRoll() -> Missing update policy for group [%v]", groupID)
	}

	list := updatePolicy.([]interface{})
	if len(list) == 0 || list[0] == nil {
		return nil
	}

	updateGroupSchema := list[0].(map[string]interface{})
	rollConfig, ok := updateGroupSchema[string(elastigroup_aws.RollConfig)]
	if !ok || rollConfig == nil {
		return fmt.Errorf("[ERROR] onRoll() -> Field [%v] is missing, skipping roll for group [%v]", string(elastigroup_aws.RollConfig), groupID)
	}

	rollGroupInput, err := expandElastigroupRollConfig(rollConfig, spotinst.String(groupID))
	if err != nil {
		return fmt.Errorf("[ERROR] onRoll() -> Failed expanding roll configuration for group [%v], error: %v", groupID, err)
	}

	json, err := commons.ToJson(rollConfig)
	if err != nil {
		return fmt.Errorf("[ERROR] onRoll() -> Failed marshaling roll configuration for group [%v], error: %v", groupID, err)
	}
	log.Printf("onRoll() -> Rolling group [%v] with configuration %s", groupID, json)

	retryTimeout := spotinst.IntValue(getRollTimeout(rollConfig))
	if retryTimeout == 0 {
		retryTimeout = 300
	}

	var rollECS bool
	if v, ok := resourceData.GetOk(string(elastigroup_aws_integrations.IntegrationEcs)); ok && v != "" {
		rollECS = true
	}

	svc := meta.(*Client).elastigroup.CloudProviderAWS()

	retryFn := func() *resource.RetryError {
		var rollOut *aws.RollGroupOutput
		var err error

		// Start the roll.
		if rollECS {
			rollOut, err = svc.RollECS(ctx, convertToECSRollInput(rollGroupInput))
		} else {
			rollOut, err = svc.Roll(ctx, rollGroupInput)
		}
		if err != nil {
			// Check whether to retry.
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

		// Wait for the roll completion.
		err = awaitReadyRoll(ctx, groupID, rollConfig, rollECS, rollOut, meta.(*Client))
		if err != nil {
			err = fmt.Errorf("[ERROR] Timed out when waiting for minimum roll percentage: %v", err)
			return resource.NonRetryableError(err)
		}

		log.Printf("onRoll() -> Successfully rolled group [%v]", groupID)
		return nil
	}

	return resource.Retry(time.Duration(retryTimeout)*time.Second, retryFn)
}

func convertToECSRollInput(rollGroupInput *aws.RollGroupInput) *aws.RollECSGroupInput {
	r := &aws.RollECSGroupInput{GroupID: rollGroupInput.GroupID}
	r.Roll = &aws.RollECSWrapper{}
	r.Roll.BatchSizePercentage = rollGroupInput.BatchSizePercentage

	return r
}

func awaitReady(groupId *string, timeout int, capacity int, client *Client) error {
	if capacity == 0 || timeout == 0 {
		return nil
	}

	err := resource.Retry(time.Second*time.Duration(timeout), func() *resource.RetryError {
		input := &aws.GetInstanceHealthinessInput{GroupID: spotinst.String(*groupId)}
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

func awaitReadyRoll(ctx context.Context, groupID string, rollConfig interface{}, rollECS bool, rollOut *aws.RollGroupOutput, client *Client) error {
	log.Printf("awaitReadyRoll() Waiting for deployment of group: %s", groupID)

	pctTimeout := spotinst.IntValue(getRollTimeout(rollConfig))
	pctComplete := spotinst.Float64Value(getRollMinPct(rollConfig))
	rollID := spotinst.StringValue(getRollStatus(rollOut))

	if pctTimeout <= 0 || pctComplete <= 0 {
		return fmt.Errorf("invalid timeout/complete durations: timeout=%d, complete=%f", pctTimeout, pctComplete)
	}
	if rollID == "" {
		return fmt.Errorf("invalid roll id: %s", rollID)
	}

	svc := client.elastigroup.CloudProviderAWS()
	err := resource.Retry(time.Second*time.Duration(pctTimeout), func() *resource.RetryError {
		var rollStatus *aws.RollGroupOutput
		var rollErr error

		deployStatusInput := &aws.DeploymentStatusInput{
			GroupID: spotinst.String(groupID),
			RollID:  spotinst.String(rollID),
		}
		if rollECS {
			rollStatus, rollErr = svc.DeploymentStatusECS(ctx, deployStatusInput)
		} else {
			rollStatus, rollErr = svc.DeploymentStatus(ctx, deployStatusInput)
		}

		if rollErr != nil {
			return resource.NonRetryableError(fmt.Errorf("call to roll status of group %q failed: %v", groupID, rollErr))
		}

		if spotinst.Float64Value(rollStatus.RollGroupStatus[0].Progress.Value) < pctComplete {
			log.Printf("awaitReadyRoll() Waiting for at least %f%% of batches to complete, current status: %f%%",
				pctComplete, spotinst.Float64Value(rollStatus.RollGroupStatus[0].Progress.Value))

			return resource.RetryableError(fmt.Errorf("roll at %v%% complete",
				spotinst.Float64Value(rollStatus.RollGroupStatus[0].Progress.Value)))
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("did not reach target deployment amount: %v", err)
	}

	log.Printf("awaitReadyRoll() Target deployment percentage reached for group: %s", groupID)
	return nil
}

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

		if v, ok := m[string(elastigroup_aws.Strategy)]; ok {
			strategy, err := expandStrategy(v)
			if err != nil {
				return nil, err
			}
			if strategy != nil {
				i.Strategy = strategy
			}
		}

	}
	return i, nil
}

func expandStrategy(data interface{}) (*aws.RollStrategy, error) {
	if list := data.([]interface{}); len(list) > 0 {
		strategy := &aws.RollStrategy{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(elastigroup_aws.Action)].(string); ok && v != "" { // Default value ""
				strategy.Action = spotinst.String(v)
			}

			if v, ok := m[string(elastigroup_aws.ShouldDrainInstances)].(bool); ok {
				strategy.ShouldDrainInstances = spotinst.Bool(v)
			}

			if v, ok := m[string(elastigroup_aws.BatchMinHealthyPercentage)].(int); ok && v >= 0 {
				strategy.BatchMinHealthyPercentage = spotinst.Int(v)
			}

			if v, ok := m[string(elastigroup_aws.OnFailure)]; ok {
				onFailure, err := expandOnFailure(v)
				if err != nil {
					return nil, err
				}
				if onFailure != nil {
					strategy.OnFailure = onFailure
				}
			}

		}

		return strategy, nil
	}

	return nil, nil
}

func expandOnFailure(data interface{}) (*aws.OnFailure, error) {
	if list := data.([]interface{}); len(list) > 0 {
		onFailure := &aws.OnFailure{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(elastigroup_aws.ActionType)].(string); ok && v != "" { // Default value ""
				onFailure.ActionType = spotinst.String(v)
			}

			if v, ok := m[string(elastigroup_aws.ShouldHandleAllBatches)].(bool); ok {
				onFailure.ShouldHandleAllBatches = spotinst.Bool(v)
			}

			if v, ok := m[string(elastigroup_aws.BatchNum)].(int); ok && v > 0 {
				onFailure.BatchNum = spotinst.Int(v)
			}

			if v, ok := m[string(elastigroup_aws.DrainingTimeout)].(int); ok && v >= 0 {
				onFailure.DrainingTimeout = spotinst.Int(v)
			}

			if v, ok := m[string(elastigroup_aws.ShouldDecrementTargetCapacity)].(bool); ok {
				onFailure.ShouldDecrementTargetCapacity = spotinst.Bool(v)
			}
		}

		return onFailure, nil
	}

	return nil, nil
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

func getRollMinPct(data interface{}) *float64 {
	var minPct *float64
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(elastigroup_aws.WaitForRollPct)].(float64); ok {
			minPct = spotinst.Float64(v)
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

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func updateCapitalSlice(resourceData *schema.ResourceData, groupResponse *aws.Group) {
	if groupResponse.Compute != nil && groupResponse.Compute.LaunchSpecification != nil && groupResponse.Compute.LaunchSpecification.BlockDeviceMappings != nil {
		blockDeviceMappings := groupResponse.Compute.LaunchSpecification.BlockDeviceMappings
		v := resourceData.Get(string(elastigroup_aws_block_devices.EbsBlockDevice))
		list := v.(*schema.Set).List()

		for index, item := range list {
			m := item.(map[string]interface{})

			if v, ok := m[string(elastigroup_aws_block_devices.VolumeType)].(string); ok && v != "" {
				if isUpper(v) == false {
					if blockDeviceMappings[index] != nil {
						if blockDeviceMappings[index].EBS != nil {
							vol := blockDeviceMappings[index].EBS.VolumeType
							if vol != nil {
								*vol = strings.ToLower(*vol)
								blockDeviceMappings[index].EBS.SetVolumeType(vol)
							}
						}
					}
				}
			}
		}
	}
}

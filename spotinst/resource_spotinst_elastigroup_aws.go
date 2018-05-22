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
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_launch_configuration"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_instance_types"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_strategy"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_block_devices"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_network_interface"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_scaling_policies"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_scheduled_task"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_stateful"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_integrations"
)

func resourceSpotinstElastigroupAws() *schema.Resource {
	setupElastigroup()
	return &schema.Resource{
		Create: resourceSpotinstElastigroupAwsCreate,
		Read:   resourceSpotinstElastigroupAwsRead,
		Update: resourceSpotinstElastigroupAwsUpdate,
		Delete: resourceSpotinstElastigroupAwsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.SpotinstElastigroup.GetSchemaMap(),
	}
}

func setupElastigroup() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	elastigroup_aws.Setup(fieldsMap)
	elastigroup_block_devices.Setup(fieldsMap)
	elastigroup_instance_types.Setup(fieldsMap)
	elastigroup_integrations.Setup(fieldsMap)
	elastigroup_launch_configuration.Setup(fieldsMap)
	elastigroup_network_interface.Setup(fieldsMap)
	elastigroup_scaling_policies.Setup(fieldsMap)
	elastigroup_scheduled_task.Setup(fieldsMap)
	elastigroup_stateful.Setup(fieldsMap)
	elastigroup_strategy.Setup(fieldsMap)

	commons.SpotinstElastigroup = commons.NewElastigroupResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstElastigroupAwsDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.SpotinstElastigroup.GetName(), id)

	input := &aws.DeleteGroupInput{GroupID: spotinst.String(id)}
	if _, err := meta.(*Client).elastigroup.CloudProviderAWS().Delete(context.Background(), input); err != nil {
		return fmt.Errorf("failed to delete group: %s", err)
	}
	resourceData.SetId("")
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
		commons.SpotinstElastigroup.GetName(), id)

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

	commons.SpotinstElastigroup.SetTerraformData(
		&commons.TerraformData{
			ResourceData: resourceData,
			Meta:         meta,
		})

	if err := commons.SpotinstElastigroup.OnRead(groupResponse); err != nil {
		return err
	}
	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstElastigroupAwsCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.SpotinstElastigroup.GetName())

	if err := commons.SpotinstElastigroup.OnCreate(resourceData, meta); err != nil {
		return err
	}

	group := commons.SpotinstElastigroup.GetElastigroup()
	groupId, err := createGroup(group, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(groupId))
	log.Printf("AWSGroup created successfully: %s", resourceData.Id())

	return resourceSpotinstElastigroupAwsRead(resourceData, meta)
}

func createGroup(group *aws.Group, spotinstClient *Client) (*string, error) {
	log.Printf("Group create configuration: %s", stringutil.Stringify(group))
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
		return nil, fmt.Errorf("failed to create group: %s", err)
	}
	return resp.Group.ID, nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstElastigroupAwsUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.SpotinstElastigroup.GetName(), id)

	shouldUpdate, err := commons.SpotinstElastigroup.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		elastigroup := commons.SpotinstElastigroup.GetElastigroup()
		elastigroup.SetId(spotinst.String(id))
		if err := updateGroup(elastigroup, resourceData, meta); err != nil {
			return err
		}
	}

	return resourceSpotinstElastigroupAwsRead(resourceData, meta)
}

func updateGroup(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &aws.UpdateGroupInput{
		Group: elastigroup,
	}

	var shouldRoll = false
	groupId := resourceData.Id()
	if updatePolicy, exists := resourceData.GetOkExists(string(elastigroup_aws.UpdatePolicy)); exists {
		list := updatePolicy.(*schema.Set).List()
		m := list[0].(map[string]interface{})

		if resumeStateful, ok := m[string(elastigroup_aws.ShouldResumeStateful)].(bool); ok && resumeStateful {
			log.Printf("Resuming paused stateful instances on group [%v]...", groupId)
			input.ShouldResumeStateful = spotinst.Bool(resumeStateful)
		}

		if roll, ok := m[string(elastigroup_aws.ShouldRoll)].(bool); ok && roll {
			shouldRoll = roll
		}
	}

	log.Printf("Group update configuration: %s", stringutil.Stringify(elastigroup))

	if _, err := meta.(*Client).elastigroup.CloudProviderAWS().Update(context.Background(), input); err != nil {
		return fmt.Errorf("failed to update group [%v]: %v", groupId, err)
	} else if shouldRoll {
		if err := rollGroup(resourceData, meta); err != nil {
			log.Printf("[ERROR] Group [%v] roll failed, error: %v", groupId, err)
			return err
		}
	} else {
		log.Printf("onRoll() -> Field [%v] is false, skipping group roll", string(elastigroup_aws.ShouldRoll))
	}
	return nil
}

func rollGroup(resourceData *schema.ResourceData, meta interface{}) error {
	var errResult error = nil
	groupId := resourceData.Id()

	if updatePolicy, exists := resourceData.GetOkExists(string(elastigroup_aws.UpdatePolicy)); exists {
		list := updatePolicy.(*schema.Set).List()
		updateGroupSchema := list[0].(map[string]interface{})

		if rollConfig, ok := updateGroupSchema[string(elastigroup_aws.RollConfig)].(*schema.Set); !ok || rollConfig == nil {
			errResult = fmt.Errorf("[ERROR] onRoll() -> Field [%v] is missing, skipping roll for group [%v]", string(elastigroup_aws.RollConfig), groupId)
		} else {
			if rollGroupInput, err := expandAWSGroupRollConfig(rollConfig, groupId); err != nil {
				errResult = fmt.Errorf("[ERROR] onRoll() -> Failed expanding roll configuration for group [%v], error: %v", groupId, err)
			} else {
				log.Printf("onRoll() -> Rolling group [%v] with configuration %v...", groupId, stringutil.Stringify(rollConfig))
				if _, err := meta.(*Client).elastigroup.CloudProviderAWS().Roll(context.Background(), rollGroupInput); err != nil {
					errResult = fmt.Errorf("[ERROR] onRoll() -> Roll group [%v] API call failed, error: %v", groupId, err)
				}
				log.Printf("onRoll() -> Successfully rolled group [%v]", groupId)
			}
		}
	} else {
		errResult = fmt.Errorf("[ERROR] onRoll() -> Missig update policy for group [%v]", groupId)
	}

	if errResult != nil {
		return errResult
	}
	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Fields Expand
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupRollConfig(data interface{}, groupID string) (*aws.RollGroupInput, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &aws.RollGroupInput{GroupID: spotinst.String(groupID)}

	if v, ok := m[string(elastigroup_aws.BatchSizePercentage)].(int); ok { // Required value
		i.BatchSizePercentage = spotinst.Int(v)
	}

	if v, ok := m[string(elastigroup_aws.GracePeriod)].(int); ok && v != -1 { // Default value set to -1
		i.GracePeriod = spotinst.Int(v)
	}

	if v, ok := m[string(elastigroup_aws.HealthCheckType)].(string); ok && v != "" { // Default value ""
		i.HealthCheckType = spotinst.String(v)
	}
	return i, nil
}
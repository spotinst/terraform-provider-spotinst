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
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/aws_elastigroup"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
)

func resourceSpotinstAWSElastigroup() *schema.Resource {
	aws_elastigroup.SetupAwsElastigroupResource()

	return &schema.Resource{
		Create: onSpotinstAWSElastigroupCreate,
		Read:   onSpotinstAWSElastigroupRead,
		Update: onSpotinstAWSElastigroupUpdate,
		Delete: onSpotinstAWSElastigroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.ElastigroupRepo.GetSchemaMap(),
	}
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstAWSElastigroupDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Deleting group: %s", d.Id())
	input := &aws.DeleteGroupInput{GroupID: spotinst.String(d.Id())}
	if _, err := meta.(*Client).elastigroup.CloudProviderAWS().Delete(context.Background(), input); err != nil {
		return fmt.Errorf("failed to delete group: %s", err)
	}
	d.SetId("")
	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
// ErrCodeGroupNotFound for service response error code "GROUP_DOESNT_EXIST".
const ErrCodeGroupNotFound = "GROUP_DOESNT_EXIST"

// Read from the lowest resource to the Spotinst elastigroup
// Since read is based on the group id, an API call is mandatory as the top down approach
// This is the reason only on read method we are triggering all the related repositories
func onSpotinstAWSElastigroupRead(resourceData *schema.ResourceData, meta interface{}) error {
	input := &aws.ReadGroupInput{GroupID: spotinst.String(resourceData.Id())}
	resp, err := meta.(*Client).elastigroup.CloudProviderAWS().Read(context.Background(), input)
	if err != nil {
		// If the group was not found, return nil so that we can show
		// that the group is gone.
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
	if resp.Group == nil {
		resourceData.SetId("")
		return nil
	}

	commons.ElastigroupRepo.OnRead(resp.Group, resourceData, meta)
	commons.LaunchConfigurationRepo.OnRead(resp.Group, resourceData, meta)

	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstAWSElastigroupCreate(resourceData *schema.ResourceData, meta interface{}) error {
	err := commons.ElastigroupRepo.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	var groupId *string
	group := commons.ElastigroupRepo.GetElastigroup()
	groupId, err = createGroup(group, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(groupId))
	log.Printf("[INFO] AWSGroup created successfully: %s", resourceData.Id())

	return onSpotinstAWSElastigroupRead(resourceData, meta)
}

func createGroup(group *aws.Group, spotinstClient *Client) (*string, error) {
	log.Printf("[DEBUG] Group create configuration: %s", stringutil.Stringify(group))
	input := &aws.CreateGroupInput{Group: group}

	var resp *aws.CreateGroupOutput
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		resp, err := spotinstClient.elastigroup.CloudProviderAWS().Create(context.Background(), input)
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
func onSpotinstAWSElastigroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	group := &aws.Group{}
	group.SetId(spotinst.String(d.Id()))

	shouldUpdate, err := commons.ElastigroupRepo.OnUpdate(d, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		updateGroup()
	}

	return onSpotinstAWSElastigroupRead(d, meta)
}

func updateGroup() {
	var shouldResumeStateful bool
	var input *aws.UpdateGroupInput

	if _, exist := d.GetOkExists("should_resume_stateful"); exist {
		log.Print("[DEBUG] Resuming paused stateful instances on group if any exist")
		shouldResumeStateful = d.Get("should_resume_stateful").(bool)
	}

	input = &aws.UpdateGroupInput{
		Group:                group,
		ShouldResumeStateful: spotinst.Bool(shouldResumeStateful),
	}

	log.Printf("[DEBUG] Group update configuration: %s", stringutil.Stringify(group))

	if _, err := client.elastigroup.CloudProviderAWS().Update(context.Background(), input); err != nil {
		return fmt.Errorf("failed to update group %s: %s", d.Id(), err)
	} else {
		// On Update Success, roll if required.
		if rc, ok := d.GetOk("roll_config"); ok {
			list := rc.(*schema.Set).List()
			m := list[0].(map[string]interface{})
			if sr, ok := m["should_roll"].(bool); ok && sr != false {
				log.Printf("[DEBUG] User has chosen to roll this group: %s", d.Id())
				if roll, err := expandAWSGroupRollConfig(rc, d.Id()); err != nil {
					log.Printf("[ERROR] Failed to expand roll configuration for group %s: %s", d.Id(), err)
					return err
				} else {
					log.Printf("[DEBUG] Sending roll request to the Spotinst API...")
					if _, err := client.elastigroup.CloudProviderAWS().Roll(context.Background(), roll); err != nil {
						log.Printf("[ERROR] Failed to roll group: %s", err)
					}
				}
			} else {
				log.Printf("[DEBUG] User has chosen not to roll this group: %s", d.Id())
			}
		}
	}
}

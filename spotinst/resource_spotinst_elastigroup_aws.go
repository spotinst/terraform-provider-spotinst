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
)

func resourceSpotinstElastigroupAws() *schema.Resource {
	elastigroup_aws.SetupAwsElastigroupResource()

	return &schema.Resource{
		Create: onSpotinstElastigroupCreateAws,
		Read:   onSpotinstElastigroupReadAws,
		Update: onSpotinstElastigroupUpdateAws,
		Delete: onSpotinstElastigroupDeleteAws,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.ElastigroupResource.GetSchemaMap(),
	}
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstElastigroupDeleteAws(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnDelete),
		commons.ElastigroupResource.GetName(), resourceData.Id())
	input := &aws.DeleteGroupInput{GroupID: spotinst.String(resourceData.Id())}
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

// Read from the lowest resource to the Spotinst elastigroup
// Since read is based on the group id, an API call is mandatory as the top down approach
// This is the reason only on read method we are triggering all the related repositories
func onSpotinstElastigroupReadAws(resourceData *schema.ResourceData, meta interface{}) error {
	return readGroup(resourceData, meta, true)
}

func readGroup(resourceData *schema.ResourceData, meta interface{}, shouldCascade bool) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.ElastigroupResource.GetName(), id)

	input := &aws.ReadGroupInput{GroupID: spotinst.String(id)}
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
	groupResponse := resp.Group
	if groupResponse == nil {
		resourceData.SetId("")
		return nil
	}

	commons.ElastigroupResource.SetTerraformData(
		&commons.TerraformData{
			ResourceData: resourceData,
			Meta:         meta,
		})

	if shouldCascade {
		cascadeGroupRead(groupResponse)
	}
	return nil
}

// We should cascade the group from parent to children only it we've initiated the elastigroup api resource
// read method. Reason is that Terraform interpolation didn't take place to trigger the read method on
// all the elastigroup cached resources
func cascadeGroupRead(elastigroup *aws.Group) {
	commons.ElastigroupResource.OnRead(elastigroup)
	commons.ElastigroupLaunchConfigurationResource.OnRead(elastigroup)
	commons.ElastigroupStrategyResource.OnRead(elastigroup)
	commons.ElastigroupInstanceTypesResource.OnRead(elastigroup)
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstElastigroupCreateAws(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.ElastigroupResource.GetName())

	err := commons.ElastigroupResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	var groupId *string
	group := commons.ElastigroupResource.GetElastigroup()
	groupId, err = createGroup(group, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(groupId))
	log.Printf("AWSGroup created successfully: %s", resourceData.Id())

	return readGroup(resourceData, meta, false)
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
func onSpotinstElastigroupUpdateAws(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.ElastigroupResource.GetName(), id)

	shouldUpdate, err := commons.ElastigroupResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		elastigroup := commons.ElastigroupResource.GetElastigroup()
		elastigroup.SetId(spotinst.String(resourceData.Id()))
		updateGroup(elastigroup, resourceData, meta)
	}

	return readGroup(resourceData, meta, true)
}

func updateGroup(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
	var shouldResumeStateful bool
	var input *aws.UpdateGroupInput

	if _, exist := resourceData.GetOkExists(string(elastigroup_aws.ShouldResumeStateful)); exist {
		shouldResumeStateful = resourceData.Get(string(elastigroup_aws.ShouldResumeStateful)).(bool)
		if shouldResumeStateful {
			log.Print("Resuming paused stateful instances on group...")
		}
	}

	input = &aws.UpdateGroupInput{
		Group:                elastigroup,
		ShouldResumeStateful: spotinst.Bool(shouldResumeStateful),
	}

	log.Printf("Group update configuration: %s", stringutil.Stringify(elastigroup))

	if _, err := meta.(*Client).elastigroup.CloudProviderAWS().Update(context.Background(), input); err != nil {
		return fmt.Errorf("failed to update group %s: %s", resourceData.Id(), err)
	} else {
		// On Update Success, roll if required.
		//if rc, ok := resourceData.GetOk(string(elastigroup_aws.RollConfig)); ok {
		//	list := rc.(*schema.Set).List()
		//	m := list[0].(map[string]interface{})
		//	if sr, ok := m["should_roll"].(bool); ok && sr != false {
		//		log.Printf("[DEBUG] User has chosen to roll this group: %s", resourceData.Id())
		//		if roll, err := expandAWSGroupRollConfig(rc, resourceData.Id()); err != nil {
		//			log.Printf("[ERROR] Failed to expand roll configuration for group %s: %s", d.Id(), err)
		//			return err
		//		} else {
		//			log.Printf("[DEBUG] Sending roll request to the Spotinst API...")
		//			if _, err := client.elastigroup.CloudProviderAWS().Roll(context.Background(), roll); err != nil {
		//				log.Printf("[ERROR] Failed to roll group: %s", err)
		//			}
		//		}
		//	} else {
		//		log.Printf("[DEBUG] User has chosen not to roll this group: %s", d.Id())
		//	}
		//}
	}
	return nil
}

package spotinst

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"log"

	"context"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/beanstalk_elastigroup"
	"strings"
	"time"
)

func resourceSpotinstBeanstalkElastigroup() *schema.Resource {
	setupElasticBeanstalk()
	return &schema.Resource{
		Create: resourceSpotinstAWSBeanstalkGroupCreate,
		Read:   resourceSpotinstAWSBeanstalkGroupRead,
		Update: resourceSpotinstAWSBeanstalkGroupUpdate,
		Delete: resourceSpotinstAWSBeanstalkGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.ElasticBeanstalkResource.GetSchemaMap(),
	}
}

func setupElasticBeanstalk() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	beanstalk_elastigroup.Setup(fieldsMap)

	commons.ElasticBeanstalkResource = commons.NewElasticBeanstalkResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//     Import Beanstalk Group
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func importBeanstalkGroup(resourceData *schema.ResourceData, meta interface{}) (*aws.Group, error) {
	input := &aws.ImportBeanstalkInput{
		EnvironmentName: spotinst.String(resourceData.Get("beanstalk_environment_name").(string)),
		Region:          spotinst.String(resourceData.Get("region").(string))}

	resp, err := meta.(*Client).elastigroup.CloudProviderAWS().ImportBeanstalkEnv(context.Background(), input)

	if err != nil {
		// If the group was not found, return nil so that we can show
		// that the group is gone.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGroupNotFound {
					resourceData.SetId("")
					return nil, err
				}
			}
		}
		// Some other error, report it.
		return nil, fmt.Errorf("BEANSTALK:IMPORT failed to read group: %s", err)
	}

	return resp.Group, err
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstAWSBeanstalkGroupCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Print(string(commons.ResourceOnCreate),
		commons.ElasticBeanstalkResource.GetName())

	beanstalkGroup, err := importBeanstalkGroup(resourceData, meta.(*Client))
	if err != nil {
		return err
	}

	tempGroup, err := commons.ElasticBeanstalkResource.OnCreate(beanstalkGroup, resourceData, meta)
	if err != nil {
		return err
	}

	groupId, err := createBeanstalkGroup(tempGroup, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(groupId))
	log.Printf("===> AWSBeanstalkGroup created successfully: %s <===", resourceData.Id())
	return resourceSpotinstAWSBeanstalkGroupRead(resourceData, meta)
}

func createBeanstalkGroup(beanstalkGroup *aws.Group, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(beanstalkGroup); err != nil {
		return nil, err
	} else {
		log.Printf("===> Group create configuration: %s", json)
	}

	input := &aws.CreateGroupInput{Group: beanstalkGroup}

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

			// If there's some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("BEANSTALK:Create failed to create group: %s", err)
	}
	return resp.Group.ID, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstAWSBeanstalkGroupRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.ElasticBeanstalkResource.GetName(), id)

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
		return fmt.Errorf("BEANSTALK:READ failed to read group: %s", err)
	}

	// If nothing was found, then return no state.
	groupResponse := resp.Group
	if groupResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.ElasticBeanstalkResource.OnRead(groupResponse, resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Elastigroup read successfully: %s <===", id)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstAWSBeanstalkGroupUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.ElasticBeanstalkResource.GetName(), id)

	shouldUpdate, beanstalkElastigroup, err := commons.ElasticBeanstalkResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		beanstalkElastigroup.SetId(spotinst.String(id))
		if err := updateGroup(beanstalkElastigroup, resourceData, meta); err != nil {
			return err
		}
	}

	log.Printf("===> Beanstalk Elastigroup updated successfully: %s <===", id)
	return resourceSpotinstAWSBeanstalkGroupRead(resourceData, meta)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstAWSBeanstalkGroupDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Deleting group: %s", d.Id())
	input := &aws.DeleteGroupInput{GroupID: spotinst.String(d.Id())}
	if _, err := meta.(*Client).elastigroup.CloudProviderAWS().Delete(context.Background(), input); err != nil {
		return fmt.Errorf("failed to delete group: %s", err)
	}
	d.SetId("")
	return nil
}

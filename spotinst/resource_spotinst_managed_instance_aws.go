package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons/managed_instance_aws"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons/managed_instance_aws_compute"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons/managed_instance_aws_compute_instance_type"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons/managed_instance_aws_compute_launchspecification"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons/managed_instance_aws_integrations"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons/managed_instance_healthcheck"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons/managed_instance_persistence"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons/managed_instance_scheduling"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons/managed_instance_strategy"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons/managed_instances_aws_compute_launchspecification_networkinterfaces"
)

func resourceSpotinstMangedInstanceAWS() *schema.Resource {
	setupMangedInstanceResource()

	return &schema.Resource{
		Create: resourceSpotinstManagedInstanceAWSCreate,
		Read:   resourceSpotinstManagedInstanceAWSRead,
		Update: resourceSpotinstManagedInstanceAWSUpdate,
		Delete: resourceSpotinstManagedInstanceAWSDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.ManagedInstanceResource.GetSchemaMap(),
	}
}

func setupMangedInstanceResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	managed_instance_aws.Setup(fieldsMap)
	managed_instance_strategy.Setup(fieldsMap)
	managed_instance_persistence.Setup(fieldsMap)
	managed_instance_healthcheck.Setup(fieldsMap)
	managed_instance_aws_compute.Setup(fieldsMap)
	managed_instance_aws_integrations.Setup(fieldsMap)
	managed_instance_scheduling.Setup(fieldsMap)
	managed_instances_aws_compute_launchspecification_networkinterfaces.Setup(fieldsMap)
	managed_instance_aws_compute_launchspecification.Setup(fieldsMap)
	managed_instance_aws_compute_instance_type.Setup(fieldsMap)

	commons.ManagedInstanceResource = commons.NewManagedInstanceResource(fieldsMap)
}

////-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
////          Read
////-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const ErrCodeManagedInstanceDoesntExist = "MANAGED_INSTANCE_DOESNT_EXIST"

func resourceSpotinstManagedInstanceAWSRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.ManagedInstanceResource.GetName(), id)

	input := &aws.ReadManagedInstanceInput{ManagedInstanceID: spotinst.String(id)}
	resp, err := meta.(*Client).managedInstance.CloudProviderAWS().Read(context.Background(), input)
	if err != nil {
		// If the group was not found, return nil so that we can show
		// that the group does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeManagedInstanceDoesntExist {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return fmt.Errorf("failed to read ManagedInstance: %s", err)
	}

	// If nothing was found, then return no state.
	managedInstanceResponse := resp.ManagedInstance
	if managedInstanceResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.ManagedInstanceResource.OnRead(managedInstanceResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> ManagedInstance read successfully: %s <===", id)
	return nil
}

////-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
////            Create
////-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstManagedInstanceAWSCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.ManagedInstanceResource.GetName())

	mangedInstance, err := commons.ManagedInstanceResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	ManagedInstanceId, err := createManagedInstance(resourceData, mangedInstance, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(ManagedInstanceId))

	log.Printf("===> ManagedInstance created successfully: %s <===", resourceData.Id())

	return resourceSpotinstManagedInstanceAWSRead(resourceData, meta)
}

func createManagedInstance(resourceData *schema.ResourceData, mangedInstance *aws.ManagedInstance, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(mangedInstance); err != nil {
		return nil, err
	} else {
		log.Printf("===> ManagedInstance create configuration: %s", json)
	}
	if v, ok := resourceData.Get(string(managed_instance_aws_compute_launchspecification.IamInstanceProfile)).(string); ok && v != "" {
		time.Sleep(5 * time.Second)
	}
	input := &aws.CreateManagedInstanceInput{ManagedInstance: mangedInstance}

	var resp *aws.CreateManagedInstanceOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.managedInstance.CloudProviderAWS().Create(context.Background(), input)
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
		return nil, fmt.Errorf("[ERROR] failed to create ManagedInstance: %s", err)
	}
	return resp.ManagedInstance.ID, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstManagedInstanceAWSUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.ManagedInstanceResource.GetName(), id)

	shouldUpdate, managedInstance, err := commons.ManagedInstanceResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		managedInstance.SetId(spotinst.String(id))
		if err := updateAWSManagedInstance(managedInstance, resourceData, meta); err != nil {
			return err
		}
	}

	log.Printf("===> ManagedInstance updated successfully: %s <===", id)
	return resourceSpotinstManagedInstanceAWSRead(resourceData, meta)
}

func updateAWSManagedInstance(managedInstance *aws.ManagedInstance, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &aws.UpdateManagedInstanceInput{
		ManagedInstance: managedInstance,
	}

	groupId := resourceData.Id()

	if json, err := commons.ToJson(managedInstance); err != nil {
		return err
	} else {
		log.Printf("===> ManagedInstance update configuration: %s", json)
	}

	if _, err := meta.(*Client).managedInstance.CloudProviderAWS().Update(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update ManagedInstance [%v]: %v", groupId, err)
	}
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//           Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstManagedInstanceAWSDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.ManagedInstanceResource.GetName(), id)

	if err := deleteManagedInstance(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> ManagedInstance deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteManagedInstance(resourceData *schema.ResourceData, meta interface{}) error {
	managedInstanceId := resourceData.Id()
	input := &aws.DeleteManagedInstanceInput{
		ManagedInstanceID: spotinst.String(managedInstanceId),
	}
	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> ManagedInstance delete configuration: %s", json)
	}

	if _, err := meta.(*Client).managedInstance.CloudProviderAWS().Delete(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete ManagedInstance: %s", err)
	}
	return nil
}

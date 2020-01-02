package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/mrscaler"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/mrscaler_aws"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/mrscaler_aws_cluster"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/mrscaler_aws_instance_groups"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/mrscaler_aws_scaling_policies"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/mrscaler_aws_scheduled_task"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/mrscaler_aws_strategy"
)

func resourceSpotinstMRScalerAWS() *schema.Resource {
	setupMRScalerAWSResource()

	return &schema.Resource{
		Create: resourceSpotinstMRScalerAWSCreate,
		Read:   resourceSpotinstMRScalerAWSRead,
		Update: resourceSpotinstMRScalerAWSUpdate,
		Delete: resourceSpotinstMRScalerAWSDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.MRScalerAWSResource.GetSchemaMap(),
	}
}

func setupMRScalerAWSResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	mrscaler_aws.Setup(fieldsMap)
	mrscaler_aws_instance_groups.Setup(fieldsMap)
	mrscaler_aws_strategy.Setup(fieldsMap)
	mrscaler_aws_cluster.Setup(fieldsMap)
	mrscaler_aws_scaling_policies.Setup(fieldsMap)
	mrscaler_aws_scheduled_task.Setup(fieldsMap)

	commons.MRScalerAWSResource = commons.NewMRScalerAWSResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMRScalerAWSCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.MRScalerAWSResource.GetName())

	scaler, err := commons.MRScalerAWSResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	scalerId, err := createScaler(scaler, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(scalerId))

	log.Printf("===> MRScaler created successfully: %s <===", resourceData.Id())

	return resourceSpotinstMRScalerAWSRead(resourceData, meta)
}

func createScaler(scaler *mrscaler.Scaler, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(scaler); err != nil {
		return nil, err
	} else {
		log.Printf("===> Scaler create configuration: %s", json)
	}

	input := &mrscaler.CreateScalerInput{Scaler: scaler}

	var resp *mrscaler.CreateScalerOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.mrscaler.Create(context.Background(), input)
		if err != nil {
			// Checks whether we should retry the scaler creation.
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
		return nil, fmt.Errorf("[ERROR] failed to create scaler: %s", err)
	}
	return resp.Scaler.ID, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMRScalerAWSRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	time.Sleep(10 * time.Second)
	log.Printf(string(commons.ResourceOnRead),
		commons.MRScalerAWSResource.GetName(), id)

	input := &mrscaler.ReadScalerInput{ScalerID: spotinst.String(id)}
	resp, err := meta.(*Client).mrscaler.Read(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to read mr scaler: %s", err)
	}

	// If nothing was found, then return no state.
	scalerResponse := resp.Scaler
	if scalerResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if exist := resourceData.Get(string(mrscaler_aws.ExposeClusterID)).(bool); exist {
		if err := exposeMrScalerClusterId(resourceData, meta); err != nil {
			return err
		}
	}

	if err := commons.MRScalerAWSResource.OnRead(scalerResponse, resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> MRScaler read successfully: %s <===", id)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMRScalerAWSUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.MRScalerAWSResource.GetName(), id)

	shouldUpdate, scaler, err := commons.MRScalerAWSResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		scaler.SetId(spotinst.String(id))
		if err := updateScaler(scaler, resourceData, meta); err != nil {
			return err
		}
	}

	log.Printf("===> MRScaler updated successfully: %s <===", id)
	return resourceSpotinstMRScalerAWSRead(resourceData, meta)
}

func updateScaler(scaler *mrscaler.Scaler, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &mrscaler.UpdateScalerInput{
		Scaler: scaler,
	}

	scalerId := resourceData.Id()

	if json, err := commons.ToJson(scaler); err != nil {
		return err
	} else {
		log.Printf("===> Scaler update configuration: %s", json)
	}

	if _, err := meta.(*Client).mrscaler.Update(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update group [%v]: %v", scalerId, err)
	}
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func resourceSpotinstMRScalerAWSDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.MRScalerAWSResource.GetName(), id)

	if err := deleteScaler(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> MRScaler deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteScaler(resourceData *schema.ResourceData, meta interface{}) error {
	scalerId := resourceData.Id()
	input := &mrscaler.DeleteScalerInput{
		ScalerID: spotinst.String(scalerId),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Scaler delete configuration: %s", json)
	}

	if _, err := meta.(*Client).mrscaler.Delete(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete scaler: %s", err)
	}
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func exposeMrScalerClusterId(resourceData *schema.ResourceData, meta interface{}) error {
	spotinstClient := meta.(*Client)
	input := &mrscaler.ScalerClusterStatusInput{ScalerID: spotinst.String(resourceData.Id())}
	resp, err := spotinstClient.mrscaler.ReadScalerCluster(context.Background(), input)

	if err != nil {
		return fmt.Errorf("failed reading cloned cluster id of mr scaler : %s", err)
	}

	if resp.ScalerClusterId != nil {
		if err = resourceData.Set(string(mrscaler_aws.OutputClusterID), resp.ScalerClusterId); err != nil {
			return err
		}
	}

	return nil
}

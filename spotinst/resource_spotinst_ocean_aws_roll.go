package spotinst

import (
	"context"
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aws_roll_config"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceSpotinstOceanAWSRoll() *schema.Resource {
	setupClusterAWSRollResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstClusterAWSRollCreate,
		ReadContext:   resourceSpotinstClusterAWSRollRead,
		UpdateContext: schema.NoopContext,
		DeleteContext: schema.NoopContext,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.OceanAWSRollResource.GetSchemaMap(),
	}
}

func setupClusterAWSRollResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	ocean_aws_roll_config.Setup(fieldsMap)

	commons.OceanAWSRollResource = commons.NewOceanAWSRollResource(fieldsMap)
}

// region Create

func resourceSpotinstClusterAWSRollCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate), commons.OceanAWSRollResource.GetName())

	rollSpec, err := commons.OceanAWSRollResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	rollId, err := createOceanAWSRoll(rollSpec, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(rollId))
	log.Printf("ocean/aws: Ocean AWS Roll created successfully: %s", resourceData.Id())

	return resourceSpotinstClusterAWSRollRead(ctx, resourceData, meta)
}

func createOceanAWSRoll(rollSpec *aws.RollSpec, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(rollSpec); err != nil {
		return nil, err
	} else {
		log.Printf("ocean/aws: roll configuration: %s", json)
	}

	input := &aws.CreateRollInput{
		Roll: rollSpec,
	}

	output, err := spotinstClient.ocean.CloudProviderAWS().CreateRoll(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("ocean/aws: failed to create the roll: %v", err)
	}

	return output.Roll.ID, nil
}

// endregion

// region Read

func resourceSpotinstClusterAWSRollRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rollId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanAWSRollResource.GetName(), rollId)
	rollSpec, err := commons.OceanAWSRollResource.OnCreate(resourceData, meta)

	input := &aws.ReadRollInput{
		RollID:    spotinst.String(rollId),
		ClusterID: rollSpec.ClusterID,
	}
	output, err := meta.(*Client).ocean.CloudProviderAWS().ReadRoll(ctx, input)
	if err != nil {
		return diag.FromErr(err)
	}

	// If nothing was found, return no state.
	RollResponse := output.Roll
	if RollResponse == nil {
		resourceData.SetId("")
		return nil
	}

	//if RollResponse != nil && *RollResponse.Status != "STOPPED" {
	if *RollResponse.Status == "COMPLETED" || *RollResponse.Status == "IN_PROGRESS" {
		status := *RollResponse.Status
		for status != "STOPPED" || status != "COMPLETED" || status != "FAILED" {
			time.Sleep(10 * time.Second)
			readResponse, err := meta.(*Client).ocean.CloudProviderAWS().ReadRoll(ctx, input)
			if err != nil {
				return diag.FromErr(err)
			}
			RollResponse = readResponse.Roll
			if *RollResponse.Status == "STOPPED" || *RollResponse.Status == "COMPLETED" || *RollResponse.Status == "FAILED" {
				break
			}
		}
		if RollResponse.Status != nil && (*RollResponse.Status == "FAILED" || *RollResponse.Status == "STOPPED") {
			return diag.FromErr(fmt.Errorf("roll status is %s", *RollResponse.Status))
		}
	}

	if err := commons.OceanAWSRollResource.OnRead(RollResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("ocean/aws: roll read successfully: %s", rollId)
	return nil
}

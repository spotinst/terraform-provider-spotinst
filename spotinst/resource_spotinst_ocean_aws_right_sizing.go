package spotinst

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceSpotinstOceanAWSRightSizingRule() *schema.Resource {
	setupOceanAWSRrightSizingRuleResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstOceanAWSRightSizingRuleCreate,
		UpdateContext: resourceSpotinstOceanAWSRightSizingRuleUpdate,
		ReadContext:   resourceSpotinstOceanAWSRightSizingRuleRead,
		DeleteContext: resourceSpotinstOceanAWSRightSizingRuleDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.OceanAWSRightSizingRuleResource.GetSchemaMap(),
	}
}

func setupOceanAWSRrightSizingRuleResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)
	ocean_aws_right_sizing_rule.Setup(fieldsMap)

	commons.OceanAWSRightSizingRuleResource = commons.NewOceanAWSRightSizingRuleResource(fieldsMap)
}

func resourceSpotinstOceanAWSRightSizingRuleRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanAWSRightSizingRuleResource.GetName(), resourceId)

	input := &aws.ReadRightSizingRuleInput{RuleName: spotinst.String(resourceId)}
	resp, err := meta.(*Client).ocean.CloudProviderAWS().ReadRightSizingRule(context.Background(), input)
	if err != nil {
		return diag.FromErr(err)
	}

	// If nothing was found, then return no state.
	RightSizingRuleResponse := resp.RightSizingRule
	if RightSizingRuleResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanAWSRightSizingRuleResource.OnRead(RightSizingRuleResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> RightSizing Rule read successfully: %s <===", resourceId)
	return nil
}

func resourceSpotinstOceanAWSRightSizingRuleCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {

	log.Printf(string(commons.ResourceOnCreate), commons.OceanAWSRightSizingRuleResource.GetName())

	rightSizingRule, err := commons.OceanAWSRightSizingRuleResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	rightSizingRuleName, err := createOceanAWSRightSizingRule(resourceData, rightSizingRule, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(rightSizingRuleName))

	log.Printf("===> RightSizing rule created successfully: %s <===", resourceData.Id())

	return resourceSpotinstOceanAWSRightSizingRuleRead(ctx, resourceData, meta)

}

func createOceanAWSRightSizingRule(resourceData *schema.ResourceData, rsr *aws.RightSizingRule, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(rsr); err != nil {
		return nil, err
	} else {
		log.Printf("===> RightSizing Rule create configuration: %s", json)
	}
	var resp *aws.CreateRightSizingRuleOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &aws.CreateRightSizingRuleInput{RightSizingRule: rsr}
		rsr, err := spotinstClient.ocean.CloudProviderAWS().CreateRightSizingRule(context.Background(), input)
		if err != nil {

			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = rsr
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create RightSizing Rule: %s", err)
	}
	return resp.RightSizingRule.Name, nil

}

func resourceSpotinstOceanAWSRightSizingRuleUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.OceanAWSRightSizingRuleResource.GetName(), resourceId)

	shouldUpdate, rsr, err := commons.OceanAWSRightSizingRuleResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		rsr.SetName(spotinst.String(resourceId))
		if err := updateOceanAWSRightSizingRule(rsr, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}
	log.Printf("===> ExtendedResourceDefinition updated successfully: %s <===", resourceId)
	return resourceSpotinstOceanAWSExtendedResourceDefinitionRead(ctx, resourceData, meta)
}

func updateOceanAWSRightSizingRule(rsr *aws.RightSizingRule, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &aws.UpdateRightSizingRuleInput{
		RightSizingRule: rsr,
	}
	erdId := resourceData.Id()

	if json, err := commons.ToJson(erd); err != nil {
		return err
	} else {
		log.Printf("===> ExtendedResourceDefinition update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().UpdateExtendedResourceDefinition(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update ExtendedResourceDefinition [%v]: %v", erdId, err)
	}
	return nil
}

func resourceSpotinstOceanAWSExtendedResourceDefinitionDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete), commons.OceanAWSExtendedResourceDefinitionResource.GetName(), resourceId)

	if err := deleteOceanAWSExtendedResourceDefinition(resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> ExtendedResourceDefinition deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteOceanAWSExtendedResourceDefinition(resourceData *schema.ResourceData, meta interface{}) error {
	erdId := resourceData.Id()
	input := &aws.DeleteExtendedResourceDefinitionInput{
		ExtendedResourceDefinitionID: spotinst.String(erdId),
	}
	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> ExtendedResourceDefinition delete configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().DeleteExtendedResourceDefinition(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete ExtendedResourceDefinition: %s", err)
	}
	return nil
}

package spotinst

import (
	"context"
	"fmt"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aws_right_sizing_rule"
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
	rightSizingRule, err := commons.OceanAWSRightSizingRuleResource.OnCreate(resourceData, meta)

	input := &aws.ReadRightSizingRuleInput{
		RuleName: spotinst.String(resourceId),
		OceanId:  rightSizingRule.OceanId,
	}
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

	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &aws.CreateRightSizingRuleInput{RightSizingRule: rsr}
		_, err := spotinstClient.ocean.CloudProviderAWS().CreateRightSizingRule(context.Background(), input)
		if err != nil {

			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create RightSizing Rule: %s", err)
	}
	return rsr.Name, nil

}

func resourceSpotinstOceanAWSRightSizingRuleUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.OceanAWSRightSizingRuleResource.GetName(), resourceId)

	shouldUpdate, rsr, err := commons.OceanAWSRightSizingRuleResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		if err := updateOceanAWSRightSizingRule(rsr, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}
	log.Printf("===> RightSizing Rule updated successfully: %s <===", resourceId)

	if rsr.Name != nil {
		resourceData.SetId(spotinst.StringValue(rsr.Name))
	}

	return resourceSpotinstOceanAWSRightSizingRuleRead(ctx, resourceData, meta)
}

func updateOceanAWSRightSizingRule(rsr *aws.RightSizingRule, resourceData *schema.ResourceData, meta interface{}) error {
	resourceId := resourceData.Id()
	rsrOnCreate, _ := commons.OceanAWSRightSizingRuleResource.OnCreate(resourceData, meta)

	rsr.OceanId = rsrOnCreate.OceanId

	var shouldAttachWorkloads = false

	if attachWorkloads, ok := resourceData.GetOk(string(ocean_aws_right_sizing_rule.AttachWorkloads)); ok {
		list := attachWorkloads.(*schema.Set).List()
		if len(list) > 0 && list[0] != nil {
			shouldAttachWorkloads = true
		}
	}

	var input = &aws.UpdateRightSizingRuleInput{
		RuleName:        spotinst.String(resourceId),
		RightSizingRule: rsr,
	}

	if json, err := commons.ToJson(rsr); err != nil {
		return err
	} else {
		log.Printf("===> RightSizingRule update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().UpdateRightSizingRule(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update Right Sizing Rule [%v]: %v", resourceId, err)
	}

	if shouldAttachWorkloads {
		if err := attachWorkloadsToRule(resourceData, meta, input.RightSizingRule.OceanId); err != nil {
			log.Printf("[ERROR] Attach Workloads for Right Sizing Rule [%v] failed, error: %v", resourceData, err)
			return err
		}
	} else {
		log.Printf("onUpdate() -> Field [%v] is missing, skipping attach workloads for right sizing rule",
			string(ocean_aws_right_sizing_rule.AttachWorkloads))
	}

	return nil
}

func attachWorkloadsToRule(resourceData *schema.ResourceData, meta interface{}, oceanId *string) error {
	ruleName := resourceData.Id()

	attachWorkloads, ok := resourceData.GetOk(string(ocean_aws_right_sizing_rule.AttachWorkloads))
	if !ok {
		return fmt.Errorf("missing attach_workloads for ocean aws right sizing rule %q", ruleName)
	}

	list := attachWorkloads.(*schema.Set).List()
	if len(list) > 0 && list[0] != nil {
		attachWorkloadsSchema := list[0].(map[string]interface{})
		if attachWorkloadsSchema == nil {
			return fmt.Errorf("missing attach workloads configuration, "+
				"skipping attach workloads for right sizing rule %q", ruleName)
		}

		attachWorkloadsSpec, err := expandAttachWorkloadsConfig(attachWorkloadsSchema, ruleName, oceanId)
		if err != nil {
			return fmt.Errorf("failed expanding attach workloads "+
				"configuration for right sizing rule %q, error: %v", ruleName, err)
		}

		updateStateJSON, err := commons.ToJson(attachWorkloadsSchema)
		if err != nil {
			return fmt.Errorf("failed marshaling attach workloads "+
				"configuration for right sizing rule %q, error: %v", ruleName, err)
		}

		log.Printf("onUpdate() -> Updating right sizing rule [%v] with configuration %s", ruleName, updateStateJSON)
		attachWorkloadsInput := &aws.RightSizingAttachDetachInput{
			RuleName:   attachWorkloadsSpec.RuleName,
			OceanId:    attachWorkloadsSpec.OceanId,
			Namespaces: attachWorkloadsSpec.Namespaces,
		}
		if _, err = meta.(*Client).ocean.CloudProviderAWS().AttachWorkloadsToRule(context.TODO(),
			attachWorkloadsInput); err != nil {
			return fmt.Errorf("onUpdate() -> Attach workloads failed for right sizing rule [%v], error: %v",
				ruleName, err)
		}
		log.Printf("onUpdate() -> Successfully attached workloads for right sizing rule [%v]", ruleName)
	}

	return nil
}

func expandAttachWorkloadsConfig(data interface{},
	ruleName string, oceanId *string) (*aws.RightSizingAttachDetachInput, error) {
	spec := &aws.RightSizingAttachDetachInput{
		OceanId:  oceanId,
		RuleName: spotinst.String(ruleName),
	}

	if data != nil {
		m := data.(map[string]interface{})

		if v, ok := m[string(ocean_aws_right_sizing_rule.Namespaces)]; ok {
			namespaces, err := expandNamespaces(v)
			if err != nil {
				return nil, err
			}

			if namespaces != nil {
				spec.Namespaces = namespaces
			}
		} else {
			spec.Namespaces = nil
		}
	}

	return spec, nil
}

func expandNamespaces(data interface{}) ([]*aws.Namespace, error) {
	list := data.(*schema.Set).List()
	namespaces := make([]*aws.Namespace, 0, len(list))

	for _, item := range list {
		attr := item.(map[string]interface{})

		namespace := &aws.Namespace{}

		if v, ok := attr[string(ocean_aws_right_sizing_rule.NamespaceName)].(string); ok && v != "" {
			namespace.NamespaceName = spotinst.String(v)
		}

		if v, ok := attr[string(ocean_aws_right_sizing_rule.Workloads)]; ok {
			workloads, err := expandWorkloads(v)
			if err != nil {
				return nil, err
			}

			if workloads != nil {
				namespace.Workloads = workloads
			}
		} else {
			namespace.Workloads = nil
		}

		if v, ok := attr[string(ocean_aws_right_sizing_rule.Labels)]; ok {
			labels, err := expandLabels(v)
			if err != nil {
				return nil, err
			}

			if labels != nil {
				namespace.Labels = labels
			}
		} else {
			namespace.Labels = nil
		}

		namespaces = append(namespaces, namespace)
	}
	return namespaces, nil
}

func expandWorkloads(data interface{}) ([]*aws.Workload, error) {
	list := data.(*schema.Set).List()
	workloads := make([]*aws.Workload, 0, len(list))

	for _, item := range list {
		attr := item.(map[string]interface{})

		workload := &aws.Workload{}

		if v, ok := attr[string(ocean_aws_right_sizing_rule.WorkloadName)].(string); ok && v != "" {
			workload.Name = spotinst.String(v)
		}

		if v, ok := attr[string(ocean_aws_right_sizing_rule.WorkloadType)].(string); ok && v != "" {
			workload.WorkloadType = spotinst.String(v)
		}

		if v, ok := attr[string(ocean_aws_right_sizing_rule.RegexName)].(string); ok && v != "" {
			workload.RegexName = spotinst.String(v)
		}

		workloads = append(workloads, workload)
	}
	return workloads, nil
}

func expandLabels(data interface{}) ([]*aws.Label, error) {
	list := data.(*schema.Set).List()
	labels := make([]*aws.Label, 0, len(list))

	for _, item := range list {
		attr := item.(map[string]interface{})

		label := &aws.Label{}

		if v, ok := attr[string(ocean_aws_right_sizing_rule.Key)].(string); ok && v != "" {
			label.Key = spotinst.String(v)
		}

		if v, ok := attr[string(ocean_aws_right_sizing_rule.Key)].(string); ok && v != "" {
			label.Value = spotinst.String(v)
		}

		labels = append(labels, label)

	}
	return labels, nil
}

func resourceSpotinstOceanAWSRightSizingRuleDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete), commons.OceanAWSRightSizingRuleResource.GetName(), resourceId)

	rightSizingRule, _ := commons.OceanAWSRightSizingRuleResource.OnCreate(resourceData, meta)
	if err := deleteOceanAWSRightSizingRule(resourceData, rightSizingRule, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> RightSizingRule deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteOceanAWSRightSizingRule(resourceData *schema.ResourceData, rsr *aws.RightSizingRule, meta interface{}) error {
	ruleName := resourceData.Id()
	input := &aws.DeleteRightSizingRuleInput{
		RuleNames: []string{ruleName},
		OceanId:   rsr.OceanId,
	}
	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> RightSizingRule delete configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderAWS().DeleteRightSizingRules(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete RightSizingRule: %s", err)
	}
	return nil
}

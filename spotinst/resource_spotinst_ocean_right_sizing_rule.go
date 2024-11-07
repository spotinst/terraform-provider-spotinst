package spotinst

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/right_sizing"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_right_sizing_rule"
)

func resourceSpotinstOceanRightSizingRule() *schema.Resource {
	setupOceanRightSizingRuleResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstOceanRightSizingRuleCreate,
		UpdateContext: resourceSpotinstOceanRightSizingRuleUpdate,
		ReadContext:   resourceSpotinstOceanRightSizingRuleRead,
		DeleteContext: resourceSpotinstOceanRightSizingRuleDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.OceanRightSizingRuleResource.GetSchemaMap(),
	}
}

func setupOceanRightSizingRuleResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)
	ocean_right_sizing_rule.Setup(fieldsMap)

	commons.OceanRightSizingRuleResource = commons.NewOceanRightSizingRuleResource(fieldsMap)
}

func resourceSpotinstOceanRightSizingRuleRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanRightSizingRuleResource.GetName(), resourceId)
	rightSizingRule, err := commons.OceanRightSizingRuleResource.OnCreate(resourceData, meta)

	input := &right_sizing.ReadRightsizingRuleInput{
		RuleName: spotinst.String(resourceId),
		OceanId:  rightSizingRule.OceanId,
	}
	resp, err := meta.(*Client).ocean.RightSizing().ReadRightsizingRule(context.Background(), input)
	if err != nil {
		return diag.FromErr(err)
	}

	// If nothing was found, then return no state.
	RightSizingRuleResponse := resp.RightsizingRule
	if RightSizingRuleResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanRightSizingRuleResource.OnRead(RightSizingRuleResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> RightSizing Rule read successfully: %s <===", resourceId)
	return nil
}

func resourceSpotinstOceanRightSizingRuleCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {

	log.Printf(string(commons.ResourceOnCreate), commons.OceanRightSizingRuleResource.GetName())

	rightSizingRule, err := commons.OceanRightSizingRuleResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	rightSizingRuleName, err := createOceanRightSizingRule(resourceData, rightSizingRule, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(rightSizingRuleName))

	log.Printf("===> RightSizing rule created successfully: %s <===", resourceData.Id())

	return resourceSpotinstOceanRightSizingRuleRead(ctx, resourceData, meta)

}

func createOceanRightSizingRule(resourceData *schema.ResourceData, rsr *right_sizing.RightsizingRule, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(rsr); err != nil {
		return nil, err
	} else {
		log.Printf("===> RightSizing Rule create configuration: %s", json)
	}

	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &right_sizing.CreateRightsizingRuleInput{RightsizingRule: rsr}
		_, err := spotinstClient.ocean.RightSizing().CreateRightsizingRule(context.Background(), input)
		if err != nil {

			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create RightSizing Rule: %s", err)
	}
	return rsr.RuleName, nil

}

func resourceSpotinstOceanRightSizingRuleUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.OceanRightSizingRuleResource.GetName(), resourceId)

	shouldUpdate, rsr, err := commons.OceanRightSizingRuleResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		if err := updateOceanRightSizingRule(rsr, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}
	log.Printf("===> RightSizing Rule updated successfully: %s <===", resourceId)

	if rsr.RuleName != nil {
		resourceData.SetId(spotinst.StringValue(rsr.RuleName))
	}

	return resourceSpotinstOceanRightSizingRuleRead(ctx, resourceData, meta)
}

func updateOceanRightSizingRule(rsr *right_sizing.RightsizingRule, resourceData *schema.ResourceData, meta interface{}) error {
	resourceId := resourceData.Id()
	rsrOnCreate, _ := commons.OceanRightSizingRuleResource.OnCreate(resourceData, meta)

	rsr.OceanId = rsrOnCreate.OceanId

	var shouldAttachWorkloads = false
	var shouldDetachWorkloads = false

	var oceanId = spotinst.StringValue(rsr.OceanId)

	if attachWorkloads, ok := resourceData.GetOk(string(ocean_right_sizing_rule.AttachRightSizingRule)); ok {
		list := attachWorkloads.(*schema.Set).List()
		if len(list) > 0 && list[0] != nil {
			shouldAttachWorkloads = true
		}
	}

	if detachWorkloads, ok := resourceData.GetOk(string(ocean_right_sizing_rule.DetachRightSizingRule)); ok {
		list := detachWorkloads.(*schema.Set).List()
		if len(list) > 0 && list[0] != nil {
			shouldDetachWorkloads = true
		}
	}

	var input = &right_sizing.UpdateRightsizingRuleInput{
		RuleName:        spotinst.String(resourceId),
		RightsizingRule: rsr,
	}

	if json, err := commons.ToJson(rsr); err != nil {
		return err
	} else {
		log.Printf("===> RightSizingRule update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.RightSizing().UpdateRightsizingRule(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update Right Sizing Rule [%v]: %v", resourceId, err)
	}

	if shouldAttachWorkloads {
		if err := attachRightSizingRule(resourceData, meta, &oceanId); err != nil {
			log.Printf("[ERROR] Attach Workloads for Right Sizing Rule [%v] failed, error: %v", resourceData, err)
			return err
		}
	} else {
		log.Printf("onUpdate() -> Field [%v] is missing, skipping attach workloads for right sizing rule",
			string(ocean_right_sizing_rule.AttachRightSizingRule))
	}

	if shouldDetachWorkloads {
		if err := detachRightSizingRule(resourceData, meta, &oceanId); err != nil {
			log.Printf("[ERROR] Attach Workloads for Right Sizing Rule [%v] failed, error: %v", resourceData, err)
			return err
		}
	} else {
		log.Printf("onUpdate() -> Field [%v] is missing, skipping detach workloads for right sizing rule",
			string(ocean_right_sizing_rule.DetachRightSizingRule))
	}

	return nil
}

func attachRightSizingRule(resourceData *schema.ResourceData, meta interface{}, oceanId *string) error {
	ruleName := resourceData.Id()

	attachWorkloads, ok := resourceData.GetOk(string(ocean_right_sizing_rule.AttachRightSizingRule))
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

		attachWorkloadsSpec, err := expandAttachWorkloadsConfig(attachWorkloadsSchema, ruleName, oceanId, meta)
		if err != nil {
			return fmt.Errorf("failed expanding attach workloads "+
				"configuration for right sizing rule %q, error: %v", ruleName, err)
		}

		updateStateJSON, err := commons.ToJson(attachWorkloadsSpec)
		if err != nil {
			return fmt.Errorf("failed marshaling attach workloads "+
				"configuration for right sizing rule %q, error: %v", ruleName, err)
		}

		if len(attachWorkloadsSpec.Namespaces) > 0 {
			log.Printf("onUpdate() -> Updating right sizing rule [%v] with configuration %s", ruleName, updateStateJSON)
			attachRuleInput := &right_sizing.RightSizingAttachDetachInput{
				RuleName:   attachWorkloadsSpec.RuleName,
				OceanId:    attachWorkloadsSpec.OceanId,
				Namespaces: attachWorkloadsSpec.Namespaces,
			}
			if _, err = meta.(*Client).ocean.RightSizing().AttachRightSizingRule(context.TODO(),
				attachRuleInput); err != nil {
				return fmt.Errorf("onUpdate() -> Attach workloads failed for right sizing rule [%v], error: %v",
					ruleName, err)
			}
			log.Printf("onUpdate() -> Successfully attached workloads for right sizing rule [%v]", ruleName)
		}
	}

	return nil
}

func detachRightSizingRule(resourceData *schema.ResourceData, meta interface{}, oceanId *string) error {
	ruleName := resourceData.Id()

	detachWorkloads, ok := resourceData.GetOk(string(ocean_right_sizing_rule.DetachRightSizingRule))
	if !ok {
		return fmt.Errorf("missing detach_workloads for ocean aws right sizing rule %q", ruleName)
	}

	list := detachWorkloads.(*schema.Set).List()
	if len(list) > 0 && list[0] != nil {
		detachWorkloadsSchema := list[0].(map[string]interface{})
		if detachWorkloadsSchema == nil {
			return fmt.Errorf("missing detach workloads configuration, "+
				"skipping detach workloads for right sizing rule %q", ruleName)
		}

		detachWorkloadsSpec, err := expandDetachWorkloadsConfig(detachWorkloadsSchema, ruleName, oceanId)
		if err != nil {
			return fmt.Errorf("failed expanding attach workloads "+
				"configuration for right sizing rule %q, error: %v", ruleName, err)
		}

		updateStateJSON, err := commons.ToJson(detachWorkloadsSpec)
		if err != nil {
			return fmt.Errorf("failed marshaling detach workloads "+
				"configuration for right sizing rule %q, error: %v", ruleName, err)
		}

		log.Printf("onUpdate() -> Updating right sizing rule [%v] with configuration %s", ruleName, updateStateJSON)
		detachWorkloadsInput := &right_sizing.RightSizingAttachDetachInput{
			RuleName:   detachWorkloadsSpec.RuleName,
			OceanId:    detachWorkloadsSpec.OceanId,
			Namespaces: detachWorkloadsSpec.Namespaces,
		}
		if _, err = meta.(*Client).ocean.RightSizing().DetachRightSizingRule(context.TODO(),
			detachWorkloadsInput); err != nil {
			return fmt.Errorf("onUpdate() -> Detach workloads failed for right sizing rule [%v], error: %v",
				ruleName, err)
		}
		log.Printf("onUpdate() -> Successfully detached workloads for right sizing rule [%v]", ruleName)
	}

	return nil
}

func expandAttachWorkloadsConfig(data interface{},
	ruleName string, oceanId *string, meta interface{}) (*right_sizing.RightSizingAttachDetachInput, error) {
	spec := &right_sizing.RightSizingAttachDetachInput{
		OceanId:  oceanId,
		RuleName: spotinst.String(ruleName),
	}

	if data != nil {
		m := data.(map[string]interface{})

		if v, ok := m[string(ocean_right_sizing_rule.Namespaces)]; ok {
			namespaces, err := expandNamespaces(v, ruleName, oceanId, meta)
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

func expandNamespaces(data interface{}, ruleName string, oceanId *string, meta interface{}) ([]*right_sizing.Namespace, error) {
	list := data.(*schema.Set).List()
	namespaces := make([]*right_sizing.Namespace, 0, len(list))

	input := &right_sizing.ReadRightsizingRuleAttachedWorkloadsInput{
		RuleName: spotinst.String(ruleName),
		OceanId:  oceanId,
	}
	resp, err := meta.(*Client).ocean.RightSizing().ReadRightsizingRuleAttachedWorkloads(context.Background(), input)
	log.Print(resp)
	log.Print(err)

	for _, item := range list {
		attr := item.(map[string]interface{})

		namespace := &right_sizing.Namespace{}

		if v, ok := attr[string(ocean_right_sizing_rule.NamespaceName)].(string); ok && v != "" {
			namespace.NamespaceName = spotinst.String(v)
		}

		if v, ok := attr[string(ocean_right_sizing_rule.Workloads)]; ok {
			workloads, err := expandWorkloads(v, namespace.NamespaceName, resp)
			if err != nil {
				return nil, err
			}

			if workloads != nil {
				namespace.Workloads = workloads
			}
		} else {
			namespace.Workloads = nil
		}

		if v, ok := attr[string(ocean_right_sizing_rule.Labels)]; ok {
			labels, err := expandLabels(v, namespace.NamespaceName, resp)
			if err != nil {
				return nil, err
			}

			if labels != nil {
				namespace.Labels = labels
			}
		} else {
			namespace.Labels = nil
		}

		if len(namespace.Labels) == 0 && len(namespace.Workloads) == 0 {
			continue
		} else {
			namespaces = append(namespaces, namespace)
		}
	}
	return namespaces, nil
}

func expandWorkloads(data interface{}, namespaceName *string, response *right_sizing.ReadRightsizingRuleAttachedWorkloadsOutput) ([]*right_sizing.Workload, error) {
	list := data.(*schema.Set).List()
	workloads := make([]*right_sizing.Workload, 0, len(list))

	// Fetching details of workload already attached .
	responseDataWorkload := response.RightsizingRuleAttachedWorkloads.RightsizingRuleWorkloads
	responseDataRegex := response.RightsizingRuleAttachedWorkloads.RightsizingRuleRegex

	for _, item := range list {
		attr := item.(map[string]interface{})

		workload := &right_sizing.Workload{}

		if v, ok := attr[string(ocean_right_sizing_rule.WorkloadName)].(string); ok && v != "" {
			workload.Name = spotinst.String(v)
		}

		if v, ok := attr[string(ocean_right_sizing_rule.WorkloadType)].(string); ok && v != "" {
			workload.WorkloadType = spotinst.String(v)
		}

		if v, ok := attr[string(ocean_right_sizing_rule.RegexName)].(string); ok && v != "" {
			workload.RegexName = spotinst.String(v)
		}

		//During first workload API call responseDataWorkload, responseDataRegex and responseDataLabel will be empty list, hence not validating user's workload again GET attachedWorkload response in IF block.
		//In case of responseDataWorkload, responseDataRegex empty list not going to validate user's workload again GET attachedWorkload response in IF block.
		if len(responseDataWorkload) == 0 && len(responseDataRegex) == 0 {
			workloads = append(workloads, workload)
		} else {
			if v, ok := attr[string(ocean_right_sizing_rule.WorkloadName)].(string); ok && v != "" {

				// Creating list[map] so that we can check whether workload requested by user is already attached on not.
				var list []map[string]interface{}
				for _, obj := range responseDataWorkload {
					item := map[string]interface{}{
						"Name":         spotinst.StringValue(obj.Name),
						"WorkloadType": spotinst.StringValue(obj.Type),
						"Namespace":    spotinst.StringValue(obj.Namespace),
					}
					list = append(list, item)
				}
				userWorkload := map[string]interface{}{
					"Name":         spotinst.StringValue(workload.Name),
					"WorkloadType": spotinst.StringValue(workload.WorkloadType),
					"Namespace":    spotinst.StringValue(namespaceName),
				}
				alreadyExist := false
				for _, item := range list {
					if reflect.DeepEqual(item, userWorkload) {
						alreadyExist = true
						break
					}
				}
				if alreadyExist {
					break
				} else {
					workloads = append(workloads, workload)
				}
			}
			if v, ok := attr[string(ocean_right_sizing_rule.RegexName)].(string); ok && v != "" {
				// Creating list[map] so that we can check whether workload requested by user is already attached on not.
				var list []map[string]interface{}
				for _, obj := range responseDataRegex {
					item := map[string]interface{}{
						"Name":         spotinst.StringValue(obj.Name),
						"WorkloadType": spotinst.StringValue(obj.WorkloadType),
						"Namespace":    spotinst.StringValue(obj.Namespace),
					}
					list = append(list, item)
				}
				userWorkload := map[string]interface{}{
					"Name":         spotinst.StringValue(workload.RegexName),
					"WorkloadType": spotinst.StringValue(workload.WorkloadType),
					"Namespace":    spotinst.StringValue(namespaceName),
				}
				alreadyExist := false
				for _, item := range list {
					if reflect.DeepEqual(item, userWorkload) {
						alreadyExist = true
						break
					}
				}
				if alreadyExist {
					break
				} else {
					workloads = append(workloads, workload)
				}
			}
		}
	}
	return workloads, nil
}

func expandLabels(data interface{}, namespaceName *string, response *right_sizing.ReadRightsizingRuleAttachedWorkloadsOutput) ([]*right_sizing.Label, error) {
	list := data.(*schema.Set).List()
	labels := make([]*right_sizing.Label, 0, len(list))

	// Fetching details of workload using labels already attached .
	responseDataLabels := response.RightsizingRuleAttachedWorkloads.RightsizingRuleLabels

	for _, item := range list {
		attr := item.(map[string]interface{})

		label := &right_sizing.Label{}

		if v, ok := attr[string(ocean_right_sizing_rule.Key)].(string); ok && v != "" {
			label.Key = spotinst.String(v)
		}

		if v, ok := attr[string(ocean_right_sizing_rule.Value)].(string); ok && v != "" {
			label.Value = spotinst.String(v)
		}

		//In case of responseDataLabel empty list not going to validate user's workload again GET attachedWorkload response in IF block.
		if len(responseDataLabels) == 0 {
			labels = append(labels, label)
		} else {
			if v, ok := attr[string(ocean_right_sizing_rule.Key)].(string); ok && v != "" {

				// Creating list[map] so that we can check whether workload having requested label by user is already attached on not.
				var list []map[string]interface{}
				for _, obj := range responseDataLabels {
					item := map[string]interface{}{
						"Key":       spotinst.StringValue(obj.Key),
						"Value":     spotinst.StringValue(obj.Value),
						"Namespace": spotinst.StringValue(obj.Namespace),
					}
					list = append(list, item)
				}
				userWorkload := map[string]interface{}{
					"Key":       spotinst.StringValue(label.Key),
					"Value":     spotinst.StringValue(label.Value),
					"Namespace": spotinst.StringValue(namespaceName),
				}
				alreadyExist := false
				for _, item := range list {
					if reflect.DeepEqual(item, userWorkload) {
						alreadyExist = true
						break
					}
				}
				if alreadyExist {
					break
				} else {
					labels = append(labels, label)
				}
			}
		}
	}
	return labels, nil
}

func expandDetachWorkloadsConfig(data interface{},
	ruleName string, oceanId *string) (*right_sizing.RightSizingAttachDetachInput, error) {
	spec := &right_sizing.RightSizingAttachDetachInput{
		OceanId:  oceanId,
		RuleName: spotinst.String(ruleName),
	}

	if data != nil {
		m := data.(map[string]interface{})

		if v, ok := m[string(ocean_right_sizing_rule.NamespacesForDetach)]; ok {
			namespaces, err := expandNamespacesForDetach(v)
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

func expandNamespacesForDetach(data interface{}) ([]*right_sizing.Namespace, error) {
	list := data.(*schema.Set).List()
	namespaces := make([]*right_sizing.Namespace, 0, len(list))

	for _, item := range list {
		attr := item.(map[string]interface{})

		namespace := &right_sizing.Namespace{}

		if v, ok := attr[string(ocean_right_sizing_rule.NamespaceNameForDetach)].(string); ok && v != "" {
			namespace.NamespaceName = spotinst.String(v)
		}

		if v, ok := attr[string(ocean_right_sizing_rule.WorkloadsForDetach)]; ok {
			workloads, err := expandWorkloadsForDetach(v)
			if err != nil {
				return nil, err
			}

			if workloads != nil {
				namespace.Workloads = workloads
			}
		} else {
			namespace.Workloads = nil
		}

		if v, ok := attr[string(ocean_right_sizing_rule.LabelsForDetach)]; ok {
			labels, err := expandLabelsForDetach(v)
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

func expandWorkloadsForDetach(data interface{}) ([]*right_sizing.Workload, error) {
	list := data.(*schema.Set).List()
	workloads := make([]*right_sizing.Workload, 0, len(list))

	for _, item := range list {
		attr := item.(map[string]interface{})

		workload := &right_sizing.Workload{}

		if v, ok := attr[string(ocean_right_sizing_rule.WorkloadNameForDetach)].(string); ok && v != "" {
			workload.Name = spotinst.String(v)
		}

		if v, ok := attr[string(ocean_right_sizing_rule.WorkloadTypeForDetach)].(string); ok && v != "" {
			workload.WorkloadType = spotinst.String(v)
		}

		if v, ok := attr[string(ocean_right_sizing_rule.RegexNameForDetach)].(string); ok && v != "" {
			workload.RegexName = spotinst.String(v)
		}

		workloads = append(workloads, workload)
	}
	return workloads, nil
}

func expandLabelsForDetach(data interface{}) ([]*right_sizing.Label, error) {
	list := data.(*schema.Set).List()
	labels := make([]*right_sizing.Label, 0, len(list))

	for _, item := range list {
		attr := item.(map[string]interface{})

		label := &right_sizing.Label{}

		if v, ok := attr[string(ocean_right_sizing_rule.KeyForDetach)].(string); ok && v != "" {
			label.Key = spotinst.String(v)
		}

		if v, ok := attr[string(ocean_right_sizing_rule.ValueForDetach)].(string); ok && v != "" {
			label.Value = spotinst.String(v)
		}

		labels = append(labels, label)

	}
	return labels, nil
}

func resourceSpotinstOceanRightSizingRuleDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete), commons.OceanRightSizingRuleResource.GetName(), resourceId)

	rightSizingRule, _ := commons.OceanRightSizingRuleResource.OnCreate(resourceData, meta)
	if err := deleteOceanRightSizingRule(resourceData, rightSizingRule, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> RightSizingRule deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteOceanRightSizingRule(resourceData *schema.ResourceData, rsr *right_sizing.RightsizingRule, meta interface{}) error {
	ruleName := resourceData.Id()
	input := &right_sizing.DeleteRightsizingRuleInput{
		RuleNames: []string{ruleName},
		OceanId:   rsr.OceanId,
	}
	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> RightSizingRule delete configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.RightSizing().DeleteRightsizingRules(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete RightSizingRule: %s", err)
	}
	return nil
}

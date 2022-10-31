package spotinst

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/multai_routing_rule"
)

func resourceSpotinstMultaiRoutingRule() *schema.Resource {
	setupMultaiRoutingRuleResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstMultaiRoutingRuleCreate,
		ReadContext:   resourceSpotinstMultaiRoutingRuleRead,
		UpdateContext: resourceSpotinstMultaiRoutingRuleUpdate,
		DeleteContext: resourceSpotinstMultaiRoutingRuleDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.MultaiRoutingRuleResource.GetSchemaMap(),
	}
}

func setupMultaiRoutingRuleResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	multai_routing_rule.Setup(fieldsMap)

	commons.MultaiRoutingRuleResource = commons.NewMultaiRoutingRuleResource(fieldsMap)
}

func resourceSpotinstMultaiRoutingRuleCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.MultaiRoutingRuleResource.GetName())

	routingRule, err := commons.MultaiRoutingRuleResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	routingRuleId, err := createRoutingRule(routingRule, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(routingRuleId))
	log.Printf("===> Routing Rule created successfully: %s <===", resourceData.Id())

	return resourceSpotinstMultaiRoutingRuleRead(ctx, resourceData, meta)
}

func createRoutingRule(routingRule *multai.RoutingRule, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(routingRule); err != nil {
		return nil, err
	} else {
		log.Printf("===> Routing Rule create configuration: %s", json)
	}

	var resp *multai.CreateRoutingRuleOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &multai.CreateRoutingRuleInput{RoutingRule: routingRule}
		r, err := spotinstClient.multai.CreateRoutingRule(context.Background(), input)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create routing rule: %s", err)
	}

	return resp.RoutingRule.ID, nil
}

func resourceSpotinstMultaiRoutingRuleRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	routingRuleId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.MultaiRoutingRuleResource.GetName(), routingRuleId)

	input := &multai.ReadRoutingRuleInput{RoutingRuleID: spotinst.String(routingRuleId)}
	resp, err := meta.(*Client).multai.ReadRoutingRule(context.Background(), input)
	if err != nil {
		return diag.Errorf("failed to read routing rule: %s", err)
	}

	// If nothing was found, return no state
	routingResponse := resp.RoutingRule
	if routingResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.MultaiRoutingRuleResource.OnRead(routingResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Routing Rule read successfully: %s <===", routingRuleId)
	return nil
}

func resourceSpotinstMultaiRoutingRuleUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	routingRuleId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.MultaiRoutingRuleResource.GetName(), routingRuleId)

	shouldUpdate, routingRule, err := commons.MultaiRoutingRuleResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		routingRule.SetId(spotinst.String(routingRuleId))
		if err := updateRoutingRule(routingRule, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("===> Routing Rule updated successfully: %s <===", routingRuleId)
	return resourceSpotinstMultaiRoutingRuleRead(ctx, resourceData, meta)
}

func updateRoutingRule(routingRule *multai.RoutingRule, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &multai.UpdateRoutingRuleInput{RoutingRule: routingRule}
	routingRuleId := resourceData.Id()

	if json, err := commons.ToJson(routingRule); err != nil {
		return err
	} else {
		log.Printf("===> Routing Rule update configuration: %s", json)
	}

	if _, err := meta.(*Client).multai.UpdateRoutingRule(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update routing rule [%v]: %v", routingRuleId, err)
	}

	return nil
}

func resourceSpotinstMultaiRoutingRuleDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	routingRuleId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.MultaiRoutingRuleResource.GetName(), routingRuleId)

	if err := deleteRoutingRule(resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Routing Rule deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteRoutingRule(resourceData *schema.ResourceData, meta interface{}) error {
	routingRuleId := resourceData.Id()
	input := &multai.DeleteRoutingRuleInput{RoutingRuleID: spotinst.String(routingRuleId)}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Routing Rule delete configuration: %s", json)
	}

	if _, err := meta.(*Client).multai.DeleteRoutingRule(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete routing rule: %s", err)
	}
	return nil
}

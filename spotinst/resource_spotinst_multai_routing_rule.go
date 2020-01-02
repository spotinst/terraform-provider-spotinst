package spotinst

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/multai_routing_rule"
)

func resourceSpotinstMultaiRoutingRule() *schema.Resource {
	setupMultaiRoutingRuleResource()

	return &schema.Resource{
		Create: resourceSpotinstMultaiRoutingRuleCreate,
		Read:   resourceSpotinstMultaiRoutingRuleRead,
		Update: resourceSpotinstMultaiRoutingRuleUpdate,
		Delete: resourceSpotinstMultaiRoutingRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.MultaiRoutingRuleResource.GetSchemaMap(),
	}
}

func setupMultaiRoutingRuleResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	multai_routing_rule.Setup(fieldsMap)

	commons.MultaiRoutingRuleResource = commons.NewMultaiRoutingRuleResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiRoutingRuleCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.MultaiRoutingRuleResource.GetName())

	routingRule, err := commons.MultaiRoutingRuleResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	routingRuleId, err := createRoutingRule(routingRule, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(routingRuleId))
	log.Printf("===> Routing Rule created successfully: %s <===", resourceData.Id())

	return resourceSpotinstMultaiRoutingRuleRead(resourceData, meta)
}

func createRoutingRule(routingRule *multai.RoutingRule, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(routingRule); err != nil {
		return nil, err
	} else {
		log.Printf("===> Routing Rule create configuration: %s", json)
	}

	input := &multai.CreateRoutingRuleInput{RoutingRule: routingRule}

	var resp *multai.CreateRoutingRuleOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
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

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiRoutingRuleRead(resourceData *schema.ResourceData, meta interface{}) error {
	routingRuleId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.MultaiRoutingRuleResource.GetName(), routingRuleId)

	input := &multai.ReadRoutingRuleInput{RoutingRuleID: spotinst.String(routingRuleId)}
	resp, err := meta.(*Client).multai.ReadRoutingRule(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to read routing rule: %s", err)
	}

	// If nothing was found, return no state
	routingResponse := resp.RoutingRule
	if routingResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.MultaiRoutingRuleResource.OnRead(routingResponse, resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Routing Rule read successfully: %s <===", routingRuleId)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiRoutingRuleUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	routingRuleId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.MultaiRoutingRuleResource.GetName(), routingRuleId)

	shouldUpdate, routingRule, err := commons.MultaiRoutingRuleResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		routingRule.SetId(spotinst.String(routingRuleId))
		if err := updateRoutingRule(routingRule, resourceData, meta); err != nil {
			return err
		}
	}

	log.Printf("===> Routing Rule updated successfully: %s <===", routingRuleId)
	return resourceSpotinstMultaiRoutingRuleRead(resourceData, meta)
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

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiRoutingRuleDelete(resourceData *schema.ResourceData, meta interface{}) error {
	routingRuleId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.MultaiRoutingRuleResource.GetName(), routingRuleId)

	if err := deleteRoutingRule(resourceData, meta); err != nil {
		return err
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

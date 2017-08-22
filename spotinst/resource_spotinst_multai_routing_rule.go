package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
)

func resourceSpotinstMultaiRoutingRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpotinstMultaiRoutingRuleCreate,
		Update: resourceSpotinstMultaiRoutingRuleUpdate,
		Read:   resourceSpotinstMultaiRoutingRuleRead,
		Delete: resourceSpotinstMultaiRoutingRuleDelete,

		Schema: map[string]*schema.Schema{
			"balancer_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"listener_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"route": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},

			"strategy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  spotinst.StrategyRoundRobin.String(),
			},

			"middleware_ids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"target_set_ids": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceSpotinstMultaiRoutingRuleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*spotinst.Client)
	rule, err := buildBalancerRoutingRuleOpts(d, meta)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Routing rule create configuration: %s",
		stringutil.Stringify(rule))
	input := &spotinst.CreateRoutingRuleInput{
		RoutingRule: rule,
	}
	resp, err := client.MultaiService.BalancerService().CreateRoutingRule(context.Background(), input)
	if err != nil {
		return fmt.Errorf("Error creating routing rule: %s", err)
	}
	d.SetId(spotinst.StringValue(resp.RoutingRule.ID))
	log.Printf("[INFO] Routing rule created successfully: %s", d.Id())
	return resourceSpotinstMultaiRoutingRuleRead(d, meta)
}

func resourceSpotinstMultaiRoutingRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*spotinst.Client)
	input := &spotinst.ReadRoutingRuleInput{
		RoutingRuleID: spotinst.String(d.Id()),
	}
	resp, err := client.MultaiService.BalancerService().ReadRoutingRule(context.Background(), input)
	if err != nil {
		return fmt.Errorf("Error retrieving routing rule: %s", err)
	}
	if rr := resp.RoutingRule; rr != nil {
		d.Set("balancer_id", rr.BalancerID)
		d.Set("listener_id", rr.ListenerID)
		d.Set("route", rr.Route)
		d.Set("priority", rr.Priority)
		d.Set("strategy", rr.Strategy)
		d.Set("middleware_ids", rr.MiddlewareIDs)
		d.Set("target_set_ids", rr.TargetSetIDs)
		d.Set("tags", flattenTags(rr.Tags))
	} else {
		d.SetId("")
	}
	return nil
}

func resourceSpotinstMultaiRoutingRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*spotinst.Client)
	rule := &spotinst.RoutingRule{ID: spotinst.String(d.Id())}
	update := false

	if d.HasChange("listener_id") {
		rule.ListenerID = spotinst.String(d.Get("listener_id").(string))
		update = true
	}

	if d.HasChange("route") {
		rule.Route = spotinst.String(d.Get("route").(string))
		update = true
	}

	if d.HasChange("priority") {
		rule.Priority = spotinst.Int(d.Get("priority").(int))
		update = true
	}

	if d.HasChange("strategy") {
		rule.Strategy = spotinst.String(d.Get("strategy").(string))
		update = true
	}

	if d.HasChange("middleware_ids") {
		if v, ok := d.GetOk("middleware_ids"); ok {
			if mddlwrs := expandBalancerMiddlewareIds(v); len(mddlwrs) > 0 {
				rule.MiddlewareIDs = mddlwrs
				update = true
			}
		}
	}

	if d.HasChange("target_set_ids") {
		if v, ok := d.GetOk("target_set_ids"); ok {
			if sets := expandBalancerTargetSetIds(v); len(sets) > 0 {
				rule.TargetSetIDs = sets
				update = true
			}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			if tags, err := expandTags(v); err != nil {
				return err
			} else {
				rule.Tags = tags
				update = true
			}
		}
	}

	if update {
		log.Printf("[DEBUG] Routing rule update configuration: %s",
			stringutil.Stringify(rule))
		input := &spotinst.UpdateRoutingRuleInput{
			RoutingRule: rule,
		}
		if _, err := client.MultaiService.BalancerService().UpdateRoutingRule(context.Background(), input); err != nil {
			return fmt.Errorf("Error updating routing rule %s: %s", d.Id(), err)
		}
	}

	return resourceSpotinstMultaiRoutingRuleRead(d, meta)
}

func resourceSpotinstMultaiRoutingRuleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*spotinst.Client)
	log.Printf("[INFO] Deleting routing rule: %s", d.Id())
	input := &spotinst.DeleteRoutingRuleInput{
		RoutingRuleID: spotinst.String(d.Id()),
	}
	if _, err := client.MultaiService.BalancerService().DeleteRoutingRule(context.Background(), input); err != nil {
		return fmt.Errorf("Error deleting routing rule: %s", err)
	}
	d.SetId("")
	return nil
}

func buildBalancerRoutingRuleOpts(d *schema.ResourceData, meta interface{}) (*spotinst.RoutingRule, error) {
	rule := &spotinst.RoutingRule{
		BalancerID: spotinst.String(d.Get("balancer_id").(string)),
		ListenerID: spotinst.String(d.Get("listener_id").(string)),
		Route:      spotinst.String(d.Get("route").(string)),
		Priority:   spotinst.Int(d.Get("priority").(int)),
		Strategy:   spotinst.String(d.Get("strategy").(string)),
	}
	if v, ok := d.GetOk("middleware_ids"); ok {
		if mddlwrs := expandBalancerMiddlewareIds(v); len(mddlwrs) > 0 {
			rule.MiddlewareIDs = mddlwrs
		}
	}
	if v, ok := d.GetOk("target_set_ids"); ok {
		if sets := expandBalancerTargetSetIds(v); len(sets) > 0 {
			rule.TargetSetIDs = sets
		}
	}
	if v, ok := d.GetOk("tags"); ok {
		if tags, err := expandTags(v); err != nil {
			return nil, err
		} else {
			rule.Tags = tags
		}
	}
	return rule, nil
}

func expandBalancerMiddlewareIds(data interface{}) []string {
	list := data.([]interface{})
	ids := make([]string, len(list))
	for i, item := range list {
		ids[i] = item.(string)
	}
	return ids
}

func expandBalancerTargetSetIds(data interface{}) []string {
	list := data.([]interface{})
	ids := make([]string, len(list))
	for i, item := range list {
		ids[i] = item.(string)
	}
	return ids
}

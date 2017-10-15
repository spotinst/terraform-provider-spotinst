package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/subscription"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
)

func resourceSpotinstSubscription() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpotinstSubscriptionCreate,
		Update: resourceSpotinstSubscriptionUpdate,
		Read:   resourceSpotinstSubscriptionRead,
		Delete: resourceSpotinstSubscriptionDelete,

		Schema: map[string]*schema.Schema{
			"resource_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"event_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				StateFunc: func(v interface{}) string {
					value := v.(string)
					return strings.ToUpper(value)
				},
			},

			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"format": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceSpotinstSubscriptionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	newSubscription, err := buildSubscriptionOpts(d, meta)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Subscription create configuration: %s", stringutil.Stringify(newSubscription))
	input := &subscription.CreateSubscriptionInput{Subscription: newSubscription}
	resp, err := client.subscription.Create(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to create subscription: %s", err)
	}
	d.SetId(spotinst.StringValue(resp.Subscription.ID))
	log.Printf("[INFO] Subscription created successfully: %s", d.Id())
	return resourceSpotinstSubscriptionRead(d, meta)
}

func resourceSpotinstSubscriptionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	input := &subscription.ReadSubscriptionInput{SubscriptionID: spotinst.String(d.Id())}
	resp, err := client.subscription.Read(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to read subscription: %s", err)
	}

	// If nothing was found, then return no state.
	if resp.Subscription == nil {
		d.SetId("")
		return nil
	}

	s := resp.Subscription
	d.Set("resource_id", s.ResourceID)
	d.Set("event_type", s.EventType)
	d.Set("protocol", s.Protocol)
	d.Set("endpoint", s.Endpoint)
	d.Set("format", s.Format)

	return nil
}

func resourceSpotinstSubscriptionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	sub := &subscription.Subscription{}
	sub.SetId(spotinst.String(d.Id()))
	update := false

	if d.HasChange("resource_id") {
		sub.SetResourceId(spotinst.String(d.Get("resource_id").(string)))
		update = true
	}

	if d.HasChange("event_type") {
		sub.SetEventType(spotinst.String(d.Get("event_type").(string)))
		update = true
	}

	if d.HasChange("protocol") {
		sub.SetProtocol(spotinst.String(d.Get("protocol").(string)))
		update = true
	}

	if d.HasChange("endpoint") {
		sub.SetEndpoint(spotinst.String(d.Get("endpoint").(string)))
		update = true
	}

	if d.HasChange("format") {
		sub.SetFormat(d.Get("format").(map[string]interface{}))
		update = true
	}

	if update {
		log.Printf("[DEBUG] Subscription update configuration: %s", stringutil.Stringify(sub))
		input := &subscription.UpdateSubscriptionInput{Subscription: sub}
		if _, err := client.subscription.Update(context.Background(), input); err != nil {
			return fmt.Errorf("failed to update subscription %s: %s", d.Id(), err)
		}
	}

	return resourceSpotinstSubscriptionRead(d, meta)
}

func resourceSpotinstSubscriptionDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}

// buildSubscriptionOpts builds the Spotinst Subscription options.
func buildSubscriptionOpts(d *schema.ResourceData, meta interface{}) (*subscription.Subscription, error) {
	subscription := &subscription.Subscription{}
	subscription.SetResourceId(spotinst.String(d.Get("resource_id").(string)))
	subscription.SetEventType(spotinst.String(strings.ToUpper(d.Get("event_type").(string))))
	subscription.SetProtocol(spotinst.String(d.Get("protocol").(string)))
	subscription.SetEndpoint(spotinst.String(d.Get("endpoint").(string)))
	subscription.SetFormat(d.Get("format").(map[string]interface{}))
	return subscription, nil
}

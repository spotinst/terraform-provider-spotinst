package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
)

func resourceSpotinstMultaiBalancer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpotinstMultaiBalancerCreate,
		Update: resourceSpotinstMultaiBalancerUpdate,
		Read:   resourceSpotinstMultaiBalancerRead,
		Delete: resourceSpotinstMultaiBalancerDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"dns_cname_aliases": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"connection_timeouts": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"idle": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"draining": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceSpotinstMultaiBalancerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	balancer, err := buildBalancerOpts(d, meta)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Balancer create configuration: %s",
		stringutil.Stringify(balancer))
	input := &multai.CreateLoadBalancerInput{
		Balancer: balancer,
	}
	resp, err := client.multai.CreateLoadBalancer(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to create balancer: %s", err)
	}
	d.SetId(spotinst.StringValue(resp.Balancer.ID))
	log.Printf("[INFO] Balancer created successfully: %s", d.Id())
	return resourceSpotinstMultaiBalancerRead(d, meta)
}

func resourceSpotinstMultaiBalancerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	input := &multai.ReadLoadBalancerInput{
		BalancerID: spotinst.String(d.Id()),
	}
	resp, err := client.multai.ReadLoadBalancer(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to read balabcer: %s", err)
	}

	// If nothing was found, then return no state.
	if resp.Balancer == nil {
		d.SetId("")
		return nil
	}

	b := resp.Balancer
	d.Set("name", b.Name)
	d.Set("dns_cname_aliases", b.DNSCNAMEAliases)
	d.Set("tags", flattenTags(b.Tags))
	d.Set("connection_timeouts", flattenBalancerTimeouts(b.Timeouts))

	return nil
}

func resourceSpotinstMultaiBalancerUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	balancer := &multai.LoadBalancer{ID: spotinst.String(d.Id())}
	update := false

	if d.HasChange("connection_timeouts") {
		if v, ok := d.GetOk("connection_timeouts"); ok {
			if timeouts, err := expandBalancerTimeouts(v); err != nil {
				return err
			} else {
				balancer.Timeouts = timeouts
				update = true
			}
		}
	}

	if d.HasChange("dns_cname_aliases") {
		if v, ok := d.GetOk("dns_cname_aliases"); ok {
			if dnsCnameAliases, err := expandBalancerDnsAliases(v); err != nil {
				return err
			} else {
				balancer.DNSCNAMEAliases = dnsCnameAliases
				update = true
			}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			if tags, err := expandTags(v); err != nil {
				return err
			} else {
				balancer.Tags = tags
				update = true
			}
		}
	}

	if update {
		log.Printf("[DEBUG] Balancer update configuration: %s",
			stringutil.Stringify(balancer))
		input := &multai.UpdateLoadBalancerInput{
			Balancer: balancer,
		}
		if _, err := client.multai.UpdateLoadBalancer(context.Background(), input); err != nil {
			return fmt.Errorf("failed to update balancer %s: %s", d.Id(), err)
		}
	}

	return resourceSpotinstMultaiBalancerRead(d, meta)
}

func resourceSpotinstMultaiBalancerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	log.Printf("[INFO] Deleting balancer: %s", d.Id())
	input := &multai.DeleteLoadBalancerInput{
		BalancerID: spotinst.String(d.Id()),
	}
	if _, err := client.multai.DeleteLoadBalancer(context.Background(), input); err != nil {
		return fmt.Errorf("failed to delete balancer: %s", err)
	}
	d.SetId("")
	return nil
}

func buildBalancerOpts(d *schema.ResourceData, meta interface{}) (*multai.LoadBalancer, error) {
	balancer := &multai.LoadBalancer{
		Name: spotinst.String(d.Get("name").(string)),
	}
	if v, ok := d.GetOk("connection_timeouts"); ok {
		if timeouts, err := expandBalancerTimeouts(v); err != nil {
			return nil, err
		} else {
			balancer.Timeouts = timeouts
		}
	}
	if v, ok := d.GetOk("dns_cname_aliases"); ok {
		if aliases, err := expandBalancerDnsAliases(v); err != nil {
			return nil, err
		} else {
			balancer.DNSCNAMEAliases = aliases
		}
	}
	if v, ok := d.GetOk("tags"); ok {
		if tags, err := expandTags(v); err != nil {
			return nil, err
		} else {
			balancer.Tags = tags
		}
	}
	return balancer, nil
}

func expandBalancerDnsAliases(data interface{}) ([]string, error) {
	list := data.([]interface{})
	aliases := make([]string, len(list))
	for i, item := range list {
		aliases[i] = item.(string)
	}
	log.Printf("[DEBUG] DNS CNAME liases configuration: %s", stringutil.Stringify(aliases))
	return aliases, nil
}

func expandBalancerTimeouts(data interface{}) (*multai.Timeouts, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	timeouts := &multai.Timeouts{}
	if v, ok := m["idle"].(int); ok {
		timeouts.Idle = spotinst.Int(v)
	}
	if v, ok := m["draining"].(int); ok {
		timeouts.Draining = spotinst.Int(v)
	}
	log.Printf("[DEBUG] Timeouts configuration: %s", stringutil.Stringify(timeouts))
	return timeouts, nil
}

func flattenBalancerTimeouts(timeouts *multai.Timeouts) []interface{} {
	out := make(map[string]interface{})
	out["idle"] = spotinst.IntValue(timeouts.Idle)
	out["draining"] = spotinst.IntValue(timeouts.Draining)
	return []interface{}{out}
}

func expandTags(data interface{}) ([]*multai.Tag, error) {
	list := data.(map[string]interface{})
	tags := make([]*multai.Tag, 0, len(list))
	for k, v := range list {
		tag := &multai.Tag{
			Key:   spotinst.String(k),
			Value: spotinst.String(v.(string)),
		}
		log.Printf("[DEBUG] Tags configuration: %s", stringutil.Stringify(tag))
		tags = append(tags, tag)
	}
	return tags, nil
}

func flattenTags(tags []*multai.Tag) map[string]string {
	out := make(map[string]string)
	for _, t := range tags {
		out[spotinst.StringValue(t.Key)] = spotinst.StringValue(t.Value)
	}
	return out
}

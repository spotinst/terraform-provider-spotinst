package spotinst

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
)

func resourceSpotinstMultaiTargetSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpotinstMultaiTargetSetCreate,
		Update: resourceSpotinstMultaiTargetSetUpdate,
		Read:   resourceSpotinstMultaiTargetSetRead,
		Delete: resourceSpotinstMultaiTargetSetDelete,

		Schema: map[string]*schema.Schema{
			"balancer_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"deployment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				StateFunc: func(v interface{}) string {
					value := v.(string)
					return strings.ToUpper(value)
				},
			},

			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"health_check": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							StateFunc: func(v interface{}) string {
								value := v.(string)
								return strings.ToUpper(value)
							},
						},

						"path": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
						},

						"interval": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"healthy_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"unhealthy_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
				Set: hashBalancerTargetSetHealthCheck,
			},

			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceSpotinstMultaiTargetSetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	set, err := buildBalancerTargetSetOpts(d, meta)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Target set create configuration: %s",
		stringutil.Stringify(set))
	input := &multai.CreateTargetSetInput{
		TargetSet: set,
	}
	resp, err := client.multai.CreateTargetSet(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to create target set: %s", err)
	}
	d.SetId(spotinst.StringValue(resp.TargetSet.ID))
	log.Printf("[INFO] Target set created successfully: %s", d.Id())
	return resourceSpotinstMultaiTargetSetRead(d, meta)
}

func resourceSpotinstMultaiTargetSetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	input := &multai.ReadTargetSetInput{
		TargetSetID: spotinst.String(d.Id()),
	}
	resp, err := client.multai.ReadTargetSet(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to read target set: %s", err)
	}

	// If nothing was found, then return no state.
	if resp.TargetSet == nil {
		d.SetId("")
		return nil
	}

	s := resp.TargetSet
	d.Set("balancer_id", s.BalancerID)
	d.Set("deployment_id", s.DeploymentID)
	d.Set("name", s.Name)
	d.Set("protocol", s.Protocol)
	d.Set("port", s.Port)
	d.Set("weight", s.Weight)
	d.Set("health_check", flattenBalancerTargetSetHealthCheck(s.HealthCheck))
	d.Set("tags", flattenTags(s.Tags))

	return nil
}

func resourceSpotinstMultaiTargetSetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	set := &multai.TargetSet{ID: spotinst.String(d.Id())}
	update := false

	if d.HasChange("name") {
		set.Name = spotinst.String(d.Get("name").(string))
		update = true
	}

	if d.HasChange("protocol") {
		set.Protocol = spotinst.String(d.Get("protocol").(string))
		update = true
	}

	if d.HasChange("port") {
		set.Port = spotinst.Int(d.Get("port").(int))
		update = true
	}

	if d.HasChange("weight") {
		set.Weight = spotinst.Int(d.Get("weight").(int))
		update = true
	}

	if d.HasChange("health_check") {
		if v, ok := d.GetOk("health_check"); ok {
			if hc, err := expandBalancerTargetSetHealthCheck(v); err != nil {
				return err
			} else {
				set.HealthCheck = hc
				update = true
			}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			if tags, err := expandTags(v); err != nil {
				return err
			} else {
				set.Tags = tags
				update = true
			}
		}
	}

	if update {
		log.Printf("[DEBUG] Target set update configuration: %s",
			stringutil.Stringify(set))
		input := &multai.UpdateTargetSetInput{
			TargetSet: set,
		}
		if _, err := client.multai.UpdateTargetSet(context.Background(), input); err != nil {
			return fmt.Errorf("failed to update target set %s: %s", d.Id(), err)
		}
	}

	return resourceSpotinstMultaiTargetSetRead(d, meta)
}

func resourceSpotinstMultaiTargetSetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	log.Printf("[INFO] Deleting target set: %s", d.Id())
	input := &multai.DeleteTargetSetInput{
		TargetSetID: spotinst.String(d.Id()),
	}
	if _, err := client.multai.DeleteTargetSet(context.Background(), input); err != nil {
		return fmt.Errorf("failed to delete target set: %s", err)
	}
	d.SetId("")
	return nil
}

func buildBalancerTargetSetOpts(d *schema.ResourceData, meta interface{}) (*multai.TargetSet, error) {
	set := &multai.TargetSet{
		BalancerID:   spotinst.String(d.Get("balancer_id").(string)),
		DeploymentID: spotinst.String(d.Get("deployment_id").(string)),
		Name:         spotinst.String(d.Get("name").(string)),
		Protocol:     spotinst.String(strings.ToUpper(d.Get("protocol").(string))),
		Port:         spotinst.Int(d.Get("port").(int)),
		Weight:       spotinst.Int(d.Get("weight").(int)),
	}
	if v, ok := d.GetOk("health_check"); ok {
		if hc, err := expandBalancerTargetSetHealthCheck(v); err != nil {
			return nil, err
		} else {
			set.HealthCheck = hc
		}
	}
	if v, ok := d.GetOk("tags"); ok {
		if tags, err := expandTags(v); err != nil {
			return nil, err
		} else {
			set.Tags = tags
		}
	}
	return set, nil
}

func expandBalancerTargetSetHealthCheck(data interface{}) (*multai.TargetSetHealthCheck, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	hc := new(multai.TargetSetHealthCheck)
	if v, ok := m["protocol"].(string); ok {
		hc.Protocol = spotinst.String(strings.ToUpper(v))
	}
	if v, ok := m["path"].(string); ok {
		hc.Path = spotinst.String(v)
	}
	if v, ok := m["port"].(int); ok && v > 0 {
		hc.Port = spotinst.Int(v)
	}
	if v, ok := m["interval"].(int); ok {
		hc.Interval = spotinst.Int(v)
	}
	if v, ok := m["timeout"].(int); ok {
		hc.Timeout = spotinst.Int(v)
	}
	if v, ok := m["healthy_threshold"].(int); ok {
		hc.HealthyThresholdCount = spotinst.Int(v)
	}
	if v, ok := m["unhealthy_threshold"].(int); ok {
		hc.UnhealthyThresholdCount = spotinst.Int(v)
	}
	log.Printf("[DEBUG] Target set health check configuration: %s",
		stringutil.Stringify(hc))
	return hc, nil
}

func flattenBalancerTargetSetHealthCheck(hc *multai.TargetSetHealthCheck) []interface{} {
	out := make(map[string]interface{})
	out["protocol"] = spotinst.StringValue(hc.Protocol)
	out["path"] = spotinst.StringValue(hc.Path)
	out["port"] = spotinst.IntValue(hc.Port)
	out["interval"] = spotinst.IntValue(hc.Interval)
	out["timeout"] = spotinst.IntValue(hc.Timeout)
	out["healthy_threshold"] = spotinst.IntValue(hc.HealthyThresholdCount)
	out["unhealthy_threshold"] = spotinst.IntValue(hc.UnhealthyThresholdCount)
	return []interface{}{out}
}

func hashBalancerTargetSetHealthCheck(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["protocol"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", m["path"].(string)))
	buf.WriteString(fmt.Sprintf("%d-", m["port"].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m["interval"].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m["timeout"].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m["healthy_threshold"].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m["unhealthy_threshold"].(int)))
	return hashcode.String(buf.String())
}

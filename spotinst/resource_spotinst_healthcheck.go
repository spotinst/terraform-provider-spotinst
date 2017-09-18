package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/healthcheck"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
)

func resourceSpotinstHealthCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpotinstHealthCheckCreate,
		Update: resourceSpotinstHealthCheckUpdate,
		Read:   resourceSpotinstHealthCheckRead,
		Delete: resourceSpotinstHealthCheckDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"resource_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"check": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"endpoint": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"interval": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"healthy": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"unhealthy": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"threshold": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"healthy": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"unhealthy": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
				Deprecated: "Attribute `threshold` is deprecated. Use `check.healthy` and `check.unhealthy` instead.",
			},

			"proxy": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceSpotinstHealthCheckCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	newHealthCheck, err := buildHealthCheckOpts(d, meta)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] HealthCheck create configuration: %#v", newHealthCheck)
	input := &healthcheck.CreateHealthCheckInput{HealthCheck: newHealthCheck}
	resp, err := client.healthCheck.Create(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to create health check: %s", err)
	}
	d.SetId(spotinst.StringValue(resp.HealthCheck.ID))
	log.Printf("[INFO] HealthCheck created successfully: %s", d.Id())
	return resourceSpotinstHealthCheckRead(d, meta)
}

func resourceSpotinstHealthCheckRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	input := &healthcheck.ReadHealthCheckInput{HealthCheckID: spotinst.String(d.Id())}
	resp, err := client.healthCheck.Read(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to read health check: %s", err)
	}
	if hc := resp.HealthCheck; hc != nil {
		d.Set("name", hc.Name)
		d.Set("resource_id", hc.ResourceID)

		// Set the check.
		check := make([]map[string]interface{}, 0, 1)
		check = append(check, map[string]interface{}{
			"protocol":  hc.Check.Protocol,
			"endpoint":  hc.Check.Endpoint,
			"port":      hc.Check.Port,
			"interval":  hc.Check.Interval,
			"timeout":   hc.Check.Timeout,
			"healthy":   hc.Check.Healthy,
			"unhealthy": hc.Check.Unhealthy,
		})
		d.Set("check", check)

		// TODO: This can be removed later; for backward compatibility only.
		// Set the threshold.
		threshold := make([]map[string]interface{}, 0, 1)
		threshold = append(threshold, map[string]interface{}{
			"healthy":   hc.Check.Healthy,
			"unhealthy": hc.Check.Unhealthy,
		})
		d.Set("threshold", threshold)

		// Set the proxy.
		proxy := make([]map[string]interface{}, 0, 1)
		proxy = append(proxy, map[string]interface{}{
			"addr": hc.ProxyAddr,
			"port": hc.ProxyPort,
		})
		d.Set("proxy", proxy)
	} else {
		d.SetId("")
	}
	return nil
}

func resourceSpotinstHealthCheckUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	healthCheck := &healthcheck.HealthCheck{}
	healthCheck.SetId(spotinst.String(d.Id()))
	update := false

	if d.HasChange("name") {
		healthCheck.SetName(spotinst.String(d.Get("name").(string)))
		update = true
	}

	if d.HasChange("resource_id") {
		healthCheck.SetResourceId(spotinst.String(d.Get("resource_id").(string)))
		update = true
	}

	if d.HasChange("check") {
		if v, ok := d.GetOk("check"); ok {
			if check, err := expandHealthCheckConfig(v); err != nil {
				return err
			} else {
				healthCheck.SetCheck(check)
				update = true
			}
		}
	}

	// TODO: This can be removed later; for backward compatibility only.
	if d.HasChange("threshold") {
		if v, ok := d.GetOk("threshold"); ok {
			if healthy, unhealthy, err := expandHealthCheckThreshold(v); err != nil {
				return err
			} else {
				if healthCheck.Check == nil {
					healthCheck.Check = &healthcheck.Check{}
				}
				if healthy != nil {
					healthCheck.Check.SetHealthy(healthy)
					update = true
				}
				if unhealthy != nil {
					healthCheck.Check.SetUnhealthy(unhealthy)
					update = true
				}
			}
		}
	}

	if d.HasChange("proxy") {
		if v, ok := d.GetOk("proxy"); ok {
			if addr, port, err := expandHealthCheckProxy(v); err != nil {
				return err
			} else {
				if addr != nil {
					healthCheck.SetProxyAddr(addr)
					update = true
				}
				if port != nil {
					healthCheck.SetProxyPort(port)
					update = true
				}
			}
		}
	}

	if update {
		log.Printf("[DEBUG] HealthCheck update configuration: %s", stringutil.Stringify(healthCheck))
		input := &healthcheck.UpdateHealthCheckInput{HealthCheck: healthCheck}
		if _, err := client.healthCheck.Update(context.Background(), input); err != nil {
			return fmt.Errorf("failed to update health check %s: %s", d.Id(), err)
		}
	}

	return resourceSpotinstHealthCheckRead(d, meta)
}

func resourceSpotinstHealthCheckDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	log.Printf("[INFO] Deleting health check: %s", d.Id())
	input := &healthcheck.DeleteHealthCheckInput{HealthCheckID: spotinst.String(d.Id())}
	if _, err := client.healthCheck.Delete(context.Background(), input); err != nil {
		return fmt.Errorf("failed to delete health check: %s", err)
	}
	d.SetId("")
	return nil
}

// buildHealthCheckOpts builds the Spotinst HealthCheck options.
func buildHealthCheckOpts(d *schema.ResourceData, meta interface{}) (*healthcheck.HealthCheck, error) {
	healthCheck := &healthcheck.HealthCheck{}
	healthCheck.SetName(spotinst.String(d.Get("name").(string)))
	healthCheck.SetResourceId(spotinst.String(d.Get("resource_id").(string)))

	if v, ok := d.GetOk("check"); ok {
		if check, err := expandHealthCheckConfig(v); err != nil {
			return nil, err
		} else {
			healthCheck.SetCheck(check)
		}
	}

	// TODO: This can be removed later; for backward compatibility only.
	if v, ok := d.GetOk("threshold"); ok {
		if healthy, unhealthy, err := expandHealthCheckThreshold(v); err != nil {
			return nil, err
		} else {
			if healthCheck.Check == nil {
				healthCheck.Check = &healthcheck.Check{}
			}
			if healthy != nil {
				healthCheck.Check.SetHealthy(healthy)
			}
			if unhealthy != nil {
				healthCheck.Check.SetUnhealthy(unhealthy)
			}
		}
	}

	if v, ok := d.GetOk("proxy"); ok {
		if addr, port, err := expandHealthCheckProxy(v); err != nil {
			return nil, err
		} else {
			if addr != nil {
				healthCheck.SetProxyAddr(addr)
			}
			if port != nil {
				healthCheck.SetProxyPort(port)
			}
		}
	}

	return healthCheck, nil
}

// expandHealthCheckConfig expands the Check block.
func expandHealthCheckConfig(data interface{}) (*healthcheck.Check, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	check := &healthcheck.Check{}

	if v, ok := m["protocol"].(string); ok && v != "" {
		check.SetProtocol(spotinst.String(v))
	}

	if v, ok := m["endpoint"].(string); ok && v != "" {
		check.SetEndpoint(spotinst.String(v))
	}

	if v, ok := m["port"].(int); ok && v >= 0 {
		check.SetPort(spotinst.Int(v))
	}

	if v, ok := m["interval"].(int); ok && v >= 0 {
		check.SetInterval(spotinst.Int(v))
	}

	if v, ok := m["timeout"].(int); ok && v >= 0 {
		check.SetTimeout(spotinst.Int(v))
	}

	if v, ok := m["healthy"].(int); ok && v >= 0 {
		check.SetHealthy(spotinst.Int(v))
	}

	if v, ok := m["unhealthy"].(int); ok && v >= 0 {
		check.SetUnhealthy(spotinst.Int(v))
	}

	log.Printf("[DEBUG] HealthCheck check configuration: %s", stringutil.Stringify(check))
	return check, nil
}

// expandHealthCheckProxy expands the Proxy block.
func expandHealthCheckProxy(data interface{}) (*string, *int, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})

	var addr *string
	var port *int

	if v, ok := m["addr"].(string); ok && v != "" {
		addr = spotinst.String(v)
	}

	if v, ok := m["port"].(int); ok && v > 0 {
		port = spotinst.Int(v)
	}

	return addr, port, nil
}

// expandHealthCheckThreshold expands the Threshold block.
func expandHealthCheckThreshold(data interface{}) (*int, *int, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	var healthy, unhealthy *int

	if v, ok := m["healthy"].(int); ok && v >= 0 {
		healthy = spotinst.Int(v)
	}

	if v, ok := m["unhealthy"].(int); ok && v >= 0 {
		unhealthy = spotinst.Int(v)
	}

	return healthy, unhealthy, nil
}

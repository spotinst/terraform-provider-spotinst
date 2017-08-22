package spotinst

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
)

func resourceSpotinstMultaiListener() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpotinstMultaiListenerCreate,
		Update: resourceSpotinstMultaiListenerUpdate,
		Read:   resourceSpotinstMultaiListenerRead,
		Delete: resourceSpotinstMultaiListenerDelete,

		Schema: map[string]*schema.Schema{
			"balancer_id": &schema.Schema{
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

			"tls_config": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_ids": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},

						"min_version": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							StateFunc: func(v interface{}) string {
								value := v.(string)
								return strings.ToUpper(value)
							},
						},

						"max_version": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							StateFunc: func(v interface{}) string {
								value := v.(string)
								return strings.ToUpper(value)
							},
						},

						"session_tickets_disabled": &schema.Schema{
							Type:     schema.TypeBool,
							Required: true,
						},

						"prefer_server_cipher_suites": &schema.Schema{
							Type:     schema.TypeBool,
							Required: true,
						},

						"cipher_suites": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
				Set: hashBalancerListenerTLSConfig,
			},

			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceSpotinstMultaiListenerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*spotinst.Client)
	listener, err := buildBalancerListenerOpts(d, meta)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Listener create configuration: %s",
		stringutil.Stringify(listener))
	input := &spotinst.CreateListenerInput{
		Listener: listener,
	}
	resp, err := client.MultaiService.BalancerService().CreateListener(context.Background(), input)
	if err != nil {
		return fmt.Errorf("Error creating listener: %s", err)
	}
	d.SetId(spotinst.StringValue(resp.Listener.ID))
	log.Printf("[INFO] Listener created successfully: %s", d.Id())
	return resourceSpotinstMultaiListenerRead(d, meta)
}

func resourceSpotinstMultaiListenerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*spotinst.Client)
	input := &spotinst.ReadListenerInput{
		ListenerID: spotinst.String(d.Id()),
	}
	resp, err := client.MultaiService.BalancerService().ReadListener(context.Background(), input)
	if err != nil {
		return fmt.Errorf("Error retrieving listener: %s", err)
	}
	if ln := resp.Listener; ln != nil {
		d.Set("balancer_id", ln.BalancerID)
		d.Set("protocol", ln.Protocol)
		d.Set("port", ln.Port)
		d.Set("tags", flattenTags(ln.Tags))
		if ln.TLSConfig != nil {
			d.Set("tls_config", flattenBalancerListenerTLSConfig(ln.TLSConfig))
		}
	} else {
		d.SetId("")
	}
	return nil
}

func resourceSpotinstMultaiListenerUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*spotinst.Client)
	listener := &spotinst.Listener{ID: spotinst.String(d.Id())}
	update := false

	if d.HasChange("protocol") {
		listener.Protocol = spotinst.String(d.Get("protocol").(string))
		update = true
	}

	if d.HasChange("port") {
		listener.Port = spotinst.Int(d.Get("port").(int))
		update = true
	}

	if d.HasChange("tls_config") {
		if v, ok := d.GetOk("tls_config"); ok {
			if config, err := expandBalancerListenerTLSConfig(v); err != nil {
				return err
			} else {
				listener.TLSConfig = config
				update = true
			}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			if tags, err := expandTags(v); err != nil {
				return err
			} else {
				listener.Tags = tags
				update = true
			}
		}
	}

	if update {
		log.Printf("[DEBUG] Listener update configuration: %s",
			stringutil.Stringify(listener))
		input := &spotinst.UpdateListenerInput{
			Listener: listener,
		}
		if _, err := client.MultaiService.BalancerService().UpdateListener(context.Background(), input); err != nil {
			return fmt.Errorf("Error updating listener %s: %s", d.Id(), err)
		}
	}

	return resourceSpotinstMultaiListenerRead(d, meta)
}

func resourceSpotinstMultaiListenerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*spotinst.Client)
	log.Printf("[INFO] Deleting listener: %s", d.Id())
	input := &spotinst.DeleteListenerInput{
		ListenerID: spotinst.String(d.Id()),
	}
	if _, err := client.MultaiService.BalancerService().DeleteListener(context.Background(), input); err != nil {
		return fmt.Errorf("Error deleting listener: %s", err)
	}
	d.SetId("")
	return nil
}

func buildBalancerListenerOpts(d *schema.ResourceData, meta interface{}) (*spotinst.Listener, error) {
	listener := &spotinst.Listener{
		BalancerID: spotinst.String(d.Get("balancer_id").(string)),
		Protocol:   spotinst.String(strings.ToUpper(d.Get("protocol").(string))),
		Port:       spotinst.Int(d.Get("port").(int)),
	}
	if v, ok := d.GetOk("tls_config"); ok {
		if config, err := expandBalancerListenerTLSConfig(v); err != nil {
			return nil, err
		} else {
			listener.TLSConfig = config
		}
	}
	if v, ok := d.GetOk("tags"); ok {
		if tags, err := expandTags(v); err != nil {
			return nil, err
		} else {
			listener.Tags = tags
		}
	}
	return listener, nil
}

func expandBalancerListenerTLSConfig(data interface{}) (*spotinst.TLSConfig, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	config := new(spotinst.TLSConfig)
	if v, ok := m["certificate_ids"].([]interface{}); ok {
		out := make([]string, len(v))
		for i, j := range v {
			out[i] = j.(string)
		}
		config.CertificateIDs = out
	}
	if v, ok := m["min_version"].(string); ok {
		config.MinVersion = spotinst.String(strings.ToUpper(v))
	}
	if v, ok := m["max_version"].(string); ok {
		config.MaxVersion = spotinst.String(strings.ToUpper(v))
	}
	if v, ok := m["session_tickets_disabled"].(bool); ok {
		config.SessionTicketsDisabled = spotinst.Bool(v)
	}
	if v, ok := m["prefer_server_cipher_suites"].(bool); ok {
		config.PreferServerCipherSuites = spotinst.Bool(v)
	}
	if v, ok := m["cipher_suites"].([]interface{}); ok {
		out := make([]string, len(v))
		for i, j := range v {
			out[i] = j.(string)
		}
		config.CipherSuites = out
	}
	log.Printf("[DEBUG] Listener TLS configuration: %s", stringutil.Stringify(config))
	return config, nil
}

func flattenBalancerListenerTLSConfig(config *spotinst.TLSConfig) []interface{} {
	out := make(map[string]interface{})
	out["certificate_ids"] = config.CertificateIDs
	out["min_version"] = spotinst.StringValue(config.MinVersion)
	out["max_version"] = spotinst.StringValue(config.MaxVersion)
	out["session_tickets_disabled"] = spotinst.BoolValue(config.SessionTicketsDisabled)
	out["prefer_server_cipher_suites"] = spotinst.BoolValue(config.PreferServerCipherSuites)
	out["cipher_suites"] = config.CipherSuites
	return []interface{}{out}
}

func hashBalancerListenerTLSConfig(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["min_version"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["max_version"].(string))))
	buf.WriteString(fmt.Sprintf("%t-", m["session_tickets_disabled"].(bool)))
	buf.WriteString(fmt.Sprintf("%t-", m["prefer_server_cipher_suites"].(bool)))
	if ids, ok := m["certificate_ids"].([]string); ok {
		for _, id := range ids {
			buf.WriteString(fmt.Sprintf("%s-", id))
		}
	}
	if css, ok := m["cipher_suites"].([]string); ok {
		for _, cs := range css {
			buf.WriteString(fmt.Sprintf("%s-", cs))
		}
	}
	return hashcode.String(buf.String())
}

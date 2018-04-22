package spotinst
//
//import (
//	"context"
//	"fmt"
//	"log"
//
//	"github.com/hashicorp/terraform/helper/schema"
//	"github.com/spotinst/spotinst-sdk-go/service/multai"
//	"github.com/spotinst/spotinst-sdk-go/spotinst"
//	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
//	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
//)
//
//func resourceSpotinstMultaiTarget() *schema.Resource {
//	return &schema.Resource{
//		Create: resourceSpotinstMultaiTargetCreate,
//		Update: resourceSpotinstMultaiTargetUpdate,
//		Read:   resourceSpotinstMultaiTargetRead,
//		Delete: resourceSpotinstMultaiTargetDelete,
//
//		Schema: map[string]*schema.Schema{
//			"balancer_id": &schema.Schema{
//				Type:     schema.TypeString,
//				Required: true,
//			},
//
//			"target_set_id": &schema.Schema{
//				Type:     schema.TypeString,
//				Required: true,
//			},
//
//			"name": &schema.Schema{
//				Type:     schema.TypeString,
//				Optional: true,
//			},
//
//			"host": &schema.Schema{
//				Type:     schema.TypeString,
//				Required: true,
//			},
//
//			"port": &schema.Schema{
//				Type:     schema.TypeInt,
//				Optional: true,
//			},
//
//			"weight": &schema.Schema{
//				Type:     schema.TypeInt,
//				Required: true,
//			},
//
//			"tags": &schema.Schema{
//				Type:     schema.TypeMap,
//				Optional: true,
//			},
//		},
//	}
//}
//
//func resourceSpotinstMultaiTargetCreate(d *schema.ResourceData, meta interface{}) error {
//	client := meta.(*Client)
//	target, err := buildBalancerTargetOpts(d, meta)
//	if err != nil {
//		return err
//	}
//	log.Printf("[DEBUG] Target create configuration: %s",
//		stringutil.Stringify(target))
//	input := &multai.CreateTargetInput{
//		Target: target,
//	}
//	resp, err := client.multai.CreateTarget(context.Background(), input)
//	if err != nil {
//		return fmt.Errorf("failed to create target: %s", err)
//	}
//	d.SetId(spotinst.StringValue(resp.Target.ID))
//	log.Printf("[INFO] Target created successfully: %s", d.Id())
//	return resourceSpotinstMultaiTargetRead(d, meta)
//}
//
//// ErrCodeTargetNotFound for service response error code "FAILED_TO_GET_TARGET".
//const ErrCodeTargetNotFound = "FAILED_TO_GET_TARGET"
//
//func resourceSpotinstMultaiTargetRead(d *schema.ResourceData, meta interface{}) error {
//	input := &multai.ReadTargetInput{
//		TargetID: spotinst.String(d.Id()),
//	}
//	resp, err := meta.(*Client).multai.ReadTarget(context.Background(), input)
//	if err != nil {
//		// If the group was not found, return nil so that we can show
//		// that the group is gone.
//		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
//			for _, err := range errs {
//				if err.Code == ErrCodeTargetNotFound {
//					d.SetId("")
//					return nil
//				}
//			}
//		}
//
//		// Some other error, report it.
//		return fmt.Errorf("failed to read target: %s", err)
//	}
//
//	// If nothing was found, then return no state.
//	if resp.Target == nil {
//		d.SetId("")
//		return nil
//	}
//
//	t := resp.Target
//	d.Set("balancer_id", t.BalancerID)
//	d.Set("target_set_id", t.TargetSetID)
//	d.Set("name", t.Name)
//	d.Set("host", t.Host)
//	d.Set("port", t.Port)
//	d.Set("weight", t.Weight)
//	d.Set("tags", flattenTags(t.Tags))
//
//	return nil
//}
//
//func resourceSpotinstMultaiTargetUpdate(d *schema.ResourceData, meta interface{}) error {
//	client := meta.(*Client)
//	target := &multai.Target{ID: spotinst.String(d.Id())}
//	update := false
//
//	if d.HasChange("name") {
//		target.Name = spotinst.String(d.Get("name").(string))
//		update = true
//	}
//
//	if d.HasChange("host") {
//		target.Host = spotinst.String(d.Get("host").(string))
//		update = true
//	}
//
//	if d.HasChange("port") {
//		target.Port = spotinst.Int(d.Get("port").(int))
//		update = true
//	}
//
//	if d.HasChange("weight") {
//		target.Weight = spotinst.Int(d.Get("weight").(int))
//		update = true
//	}
//
//	if d.HasChange("tags") {
//		if v, ok := d.GetOk("tags"); ok {
//			if tags, err := expandTags(v); err != nil {
//				return err
//			} else {
//				target.Tags = tags
//				update = true
//			}
//		}
//	}
//
//	if update {
//		log.Printf("[DEBUG] Target update configuration: %s",
//			stringutil.Stringify(target))
//		input := &multai.UpdateTargetInput{
//			Target: target,
//		}
//		if _, err := client.multai.UpdateTarget(context.Background(), input); err != nil {
//			return fmt.Errorf("failed to update target %s: %s", d.Id(), err)
//		}
//	}
//
//	return resourceSpotinstMultaiTargetRead(d, meta)
//}
//
//func resourceSpotinstMultaiTargetDelete(d *schema.ResourceData, meta interface{}) error {
//	client := meta.(*Client)
//	log.Printf("[INFO] Deleting target: %s", d.Id())
//	input := &multai.DeleteTargetInput{
//		TargetID: spotinst.String(d.Id()),
//	}
//	if _, err := client.multai.DeleteTarget(context.Background(), input); err != nil {
//		return fmt.Errorf("failed to delete target: %s", err)
//	}
//	d.SetId("")
//	return nil
//}
//
//func buildBalancerTargetOpts(d *schema.ResourceData, meta interface{}) (*multai.Target, error) {
//	target := &multai.Target{
//		BalancerID:  spotinst.String(d.Get("balancer_id").(string)),
//		TargetSetID: spotinst.String(d.Get("target_set_id").(string)),
//		Name:        spotinst.String(d.Get("name").(string)),
//		Host:        spotinst.String(d.Get("host").(string)),
//		Port:        spotinst.Int(d.Get("port").(int)),
//		Weight:      spotinst.Int(d.Get("weight").(int)),
//	}
//	if v, ok := d.GetOk("tags"); ok {
//		if tags, err := expandTags(v); err != nil {
//			return nil, err
//		} else {
//			target.Tags = tags
//		}
//	}
//	return target, nil
//}

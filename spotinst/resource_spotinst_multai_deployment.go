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

func resourceSpotinstMultaiDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpotinstMultaiDeploymentCreate,
		Update: resourceSpotinstMultaiDeploymentUpdate,
		Read:   resourceSpotinstMultaiDeploymentRead,
		Delete: resourceSpotinstMultaiDeploymentDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"tags": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceSpotinstMultaiDeploymentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	deployment, err := buildDeploymentOpts(d, meta)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Deployment create configuration: %s",
		stringutil.Stringify(deployment))
	input := &multai.CreateDeploymentInput{
		Deployment: deployment,
	}
	resp, err := client.multai.CreateDeployment(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to create deployment: %s", err)
	}
	d.SetId(spotinst.StringValue(resp.Deployment.ID))
	log.Printf("[INFO] Deployment created successfully: %s", d.Id())
	return resourceSpotinstMultaiDeploymentRead(d, meta)
}

func resourceSpotinstMultaiDeploymentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	input := &multai.ReadDeploymentInput{
		DeploymentID: spotinst.String(d.Id()),
	}
	resp, err := client.multai.ReadDeployment(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to read deployment: %s", err)
	}
	if b := resp.Deployment; b != nil {
		d.Set("name", b.Name)
		d.Set("tags", flattenTags(b.Tags))
	} else {
		d.SetId("")
	}
	return nil
}

func resourceSpotinstMultaiDeploymentUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	deployment := &multai.Deployment{ID: spotinst.String(d.Id())}
	update := false

	if d.HasChange("name") {
		deployment.Name = spotinst.String(d.Get("name").(string))
		update = true
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			if tags, err := expandTags(v); err != nil {
				return err
			} else {
				deployment.Tags = tags
				update = true
			}
		}
	}

	if update {
		log.Printf("[DEBUG] Deployment update configuration: %s",
			stringutil.Stringify(deployment))
		input := &multai.UpdateDeploymentInput{
			Deployment: deployment,
		}
		if _, err := client.multai.UpdateDeployment(context.Background(), input); err != nil {
			return fmt.Errorf("failed to update deployment %s: %s", d.Id(), err)
		}
	}

	return resourceSpotinstMultaiDeploymentRead(d, meta)
}

func resourceSpotinstMultaiDeploymentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	log.Printf("[INFO] Deleting deployment: %s", d.Id())
	input := &multai.DeleteDeploymentInput{
		DeploymentID: spotinst.String(d.Id()),
	}
	if _, err := client.multai.DeleteDeployment(context.Background(), input); err != nil {
		return fmt.Errorf("failed to delete deployment: %s", err)
	}
	d.SetId("")
	return nil
}

func buildDeploymentOpts(d *schema.ResourceData, meta interface{}) (*multai.Deployment, error) {
	deployment := &multai.Deployment{
		Name: spotinst.String(d.Get("name").(string)),
	}
	if v, ok := d.GetOk("tags"); ok {
		if tags, err := expandTags(v); err != nil {
			return nil, err
		} else {
			deployment.Tags = tags
		}
	}
	return deployment, nil
}

package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
)

// NOTE:
// Currently the values set to default -1 are ones which have to be allowed to
// be optional as they are either optional or not always set. These are all
// target min and max values. Each of these must be validated to be != -1 upon
// creation and update. IOPS has default 0 as it cannot be otherwise anyway.

//region Resource

func resourceSpotinstAWSBeanstalkElastigroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpotinstAWSBeanstalkGroupCreate,
		Read:   resourceSpotinstAWSBeanstalkGroupRead,
		Update: resourceSpotinstAWSBeanstalkGroupUpdate,
		Delete: resourceSpotinstAWSBeanstalkGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"region": {
				Type:     schema.TypeString,
				Required: true,
			},

			"product": {
				Type:     schema.TypeString,
				Required: true,
			},

			"minimum": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"maximum": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"target": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"beanstalk_environment_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"spot_instance_types": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

//endregion

//region CRUD methods

func resourceSpotinstAWSBeanstalkGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	newGroup, err := importBeanstalkGroup(d, meta)

	buildEmptyElastigroupCapacity(newGroup)
	buildEmptyElastigroupInstanceTypes(newGroup)
	newGroup.Compute.SetProduct(spotinst.String(d.Get("product").(string)))
	newGroup.SetName(spotinst.String(d.Get("name").(string)))

	if v, ok := d.GetOkExists("minimum"); ok {
		newGroup.Capacity.SetMinimum(spotinst.Int(v.(int)))
	}

	if v, ok := d.GetOkExists("maximum"); ok {
		newGroup.Capacity.SetMaximum(spotinst.Int(v.(int)))
	}

	if v, ok := d.GetOkExists("target"); ok {
		newGroup.Capacity.SetTarget(spotinst.Int(v.(int)))
	}

	if v, ok := d.GetOk("spot_instance_types"); ok {
		types := expandElastigroupInstanceTypesList(v.([]interface{}))
		newGroup.Compute.InstanceTypes.SetSpot(types)
	}

	log.Printf("[DEBUG] Group create configuration: %s", stringutil.Stringify(newGroup))

	input := &aws.CreateGroupInput{Group: newGroup}
	resp, err := client.elastigroup.CloudProviderAWS().Create(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to create group: %s", err)
	}

	d.SetId(spotinst.StringValue(resp.Group.ID))
	log.Printf("[INFO] AWSBeanstalkGroup created successfully: %s", d.Id())
	return resourceSpotinstAWSBeanstalkGroupRead(d, meta)
}

func resourceSpotinstAWSBeanstalkGroupRead(d *schema.ResourceData, meta interface{}) error {
	input := &aws.ReadGroupInput{GroupID: spotinst.String(d.Id())}
	resp, err := meta.(*Client).elastigroup.CloudProviderAWS().Read(context.Background(), input)

	if err != nil {
		// If the group was not found, return nil so that we can show
		// that the group is gone.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGroupNotFound {
					d.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return fmt.Errorf("failed to read group: %s", err)
	}

	// If nothing was found, then return no state.
	if resp.Group == nil {
		d.SetId("")
		return nil
	}

	g := resp.Group
	d.Set("name", g.Name)
	d.Set("product", g.Compute.Product)

	if g.Capacity != nil {
		d.Set("minimum", g.Capacity.Minimum)
		d.Set("maximum", g.Capacity.Maximum)
		d.Set("target", g.Capacity.Target)
	}

	if g.Compute != nil {
		if g.Compute.InstanceTypes != nil {
			d.Set("spot_instance_types", g.Compute.InstanceTypes.Spot)
		}
	}

	return nil
}

func importBeanstalkGroup(d *schema.ResourceData, meta interface{}) (*aws.Group, error) {
	input := &aws.ImportBeanstalkInput{
		EnvironmentName: spotinst.String(d.Get("beanstalk_environment_name").(string)),
		Region:          spotinst.String(d.Get("region").(string))}

	resp, err := meta.(*Client).elastigroup.CloudProviderAWS().ImportBeanstalkEnv(context.Background(), input)

	if err != nil {
		// If the group was not found, return nil so that we can show
		// that the group is gone.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGroupNotFound {
					d.SetId("")
					return nil, err
				}
			}
		}

		// Some other error, report it.
		return nil, fmt.Errorf("failed to read group: %s", err)
	}

	return resp.Group, err
}

func resourceSpotinstAWSBeanstalkGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	group := &aws.Group{}
	group.SetId(spotinst.String(d.Id()))
	update := false
	disallowedFieldUpdate := false
	disallowedField := ""

	if d.HasChange("region") {
		disallowedField = "region"
		disallowedFieldUpdate = true
	}

	if d.HasChange("product") {
		disallowedField = "product"
		disallowedFieldUpdate = true
	}

	if d.HasChange("beanstalk_environment_name") {
		disallowedField = "beanstalk_environment_name"
		disallowedFieldUpdate = true
	}

	if disallowedFieldUpdate == true {
		return fmt.Errorf("field %s is immutable - revert to previous value to proceed", disallowedField)
	}

	if d.HasChange("name") {
		group.SetName(spotinst.String(d.Get("name").(string)))
		update = true
	}

	if d.HasChange("minimum") || d.HasChange("maximum") || d.HasChange("target") {

		if group.Capacity == nil {
			newCapacity := &aws.Capacity{}
			group.SetCapacity(newCapacity)
		}

		if v, ok := d.GetOkExists("minimum"); ok {
			group.Capacity.SetMinimum(spotinst.Int(v.(int)))
		}

		if v, ok := d.GetOkExists("maximum"); ok {
			group.Capacity.SetMaximum(spotinst.Int(v.(int)))
		}

		if v, ok := d.GetOkExists("target"); ok {
			group.Capacity.SetTarget(spotinst.Int(v.(int)))
		}

		update = true
	}

	if d.HasChange("spot_instance_types") {
		buildEmptyElastigroupInstanceTypes(group)
		if v, ok := d.GetOk("spot_instance_types"); ok {
			types := expandElastigroupInstanceTypesList(v.([]interface{}))
			group.Compute.InstanceTypes.SetSpot(types)
			update = true
		}
	}

	if update {
		log.Printf("[DEBUG] Group update configuration: %s", stringutil.Stringify(group))
		input := &aws.UpdateGroupInput{Group: group}
		if _, err := client.elastigroup.CloudProviderAWS().Update(context.Background(), input); err != nil {
			return fmt.Errorf("failed to update group %s: %s", d.Id(), err)
		}
	}

	return resourceSpotinstAWSBeanstalkGroupRead(d, meta)
}

func resourceSpotinstAWSBeanstalkGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	log.Printf("[INFO] Deleting group: %s", d.Id())
	input := &aws.DeleteGroupInput{GroupID: spotinst.String(d.Id())}
	if _, err := client.elastigroup.CloudProviderAWS().Delete(context.Background(), input); err != nil {
		return fmt.Errorf("failed to delete group: %s", err)
	}
	d.SetId("")
	return nil
}

//endregion

//region Expand methods
func expandElastigroupInstanceTypesList(instanceTypes []interface{}) []string {
	types := make([]string, 0, len(instanceTypes))
	for _, str := range instanceTypes {
		if typ, ok := str.(string); ok {
			types = append(types, typ)
		}
	}
	return types
}

//endregion

//region Build Empty
func buildEmptyElastigroupCompute(group *aws.Group) {
	if group.Compute == nil {
		group.SetCompute(&aws.Compute{})
	}
}

func buildEmptyElastigroupCapacity(group *aws.Group) {
	if group.Capacity == nil {
		group.SetCapacity(&aws.Capacity{})
	}
}

func buildEmptyElastigroupInstanceTypes(group *aws.Group) {

	buildEmptyElastigroupCompute(group)

	if group.Compute.InstanceTypes == nil {
		group.Compute.SetInstanceTypes(&aws.InstanceTypes{})
	}
}

//endregion

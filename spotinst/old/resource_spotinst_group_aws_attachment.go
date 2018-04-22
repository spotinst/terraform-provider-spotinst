package spotinst
//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"log"
//
//	"github.com/hashicorp/errwrap"
//	"github.com/hashicorp/terraform/helper/resource"
//	"github.com/hashicorp/terraform/helper/schema"
//	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
//	"github.com/spotinst/spotinst-sdk-go/spotinst"
//)
//
//func resourceSpotinstAWSGroupAttachment() *schema.Resource {
//	return &schema.Resource{
//		Create: resourceSpotinstAWSGroupAttachmentCreate,
//		Read:   resourceSpotinstAWSGroupAttachmentRead,
//		Delete: resourceSpotinstAWSGroupAttachmentDelete,
//
//		Schema: map[string]*schema.Schema{
//			"group_id": {
//				Type:     schema.TypeString,
//				ForceNew: true,
//				Required: true,
//			},
//
//			"elb": {
//				Type:     schema.TypeString,
//				ForceNew: true,
//				Optional: true,
//			},
//
//			"alb_target_group_arn": {
//				Type:     schema.TypeString,
//				ForceNew: true,
//				Optional: true,
//			},
//		},
//	}
//}
//
//func resourceSpotinstAWSGroupAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
//	client := meta.(*Client)
//	service := client.elastigroup.CloudProviderAWS()
//	groupID := d.Get("group_id").(string)
//
//	// Retrieve the group properties to get list of associated load balancers.
//	input := &aws.ReadGroupInput{GroupID: spotinst.String(groupID)}
//	output, err := service.Read(context.Background(), input)
//	if err != nil {
//		return fmt.Errorf("failed to read group: %s", err)
//	}
//
//	// If nothing was found, then return no state.
//	if output.Group == nil {
//		d.SetId("")
//		return nil
//	}
//
//	var lbs []*aws.LoadBalancer
//	var update bool
//
//	if lc := output.Group.Compute.LaunchSpecification; lc != nil {
//		if lc.LoadBalancersConfig != nil {
//			lbs = lc.LoadBalancersConfig.LoadBalancers
//		}
//	}
//
//	if v, ok := d.GetOk("elb"); ok {
//		log.Printf("[INFO] Registering group %q with ELB %q", groupID, v.(string))
//		lbs = append(lbs, &aws.LoadBalancer{
//			Type: spotinst.String(loadBalanacerTypeClassic),
//			Name: spotinst.String(v.(string)),
//		})
//		update = true
//	}
//
//	if v, ok := d.GetOk("alb_target_group_arn"); ok {
//		log.Printf("[INFO] Registering group %q with ALB Target Group %q", groupID, v.(string))
//		lbs = append(lbs, &aws.LoadBalancer{
//			Type: spotinst.String(loadBalanacerTypeTargetGroup),
//			Name: spotinst.String(v.(string)),
//			Arn:  spotinst.String(v.(string)),
//		})
//		update = true
//	}
//
//	if update {
//		input := &aws.UpdateGroupInput{
//			Group: &aws.Group{
//				ID: spotinst.String(groupID),
//				Compute: &aws.Compute{
//					LaunchSpecification: &aws.LaunchSpecification{
//						LoadBalancersConfig: &aws.LoadBalancersConfig{},
//					},
//				},
//			},
//		}
//		input.Group.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(lbs)
//		b, err := json.MarshalIndent(input, "", "  ")
//		if err == nil {
//			log.Printf("[DEBUG] Update group payload: %s", string(b))
//		}
//		if _, err := service.Update(context.Background(), input); err != nil {
//			return errwrap.Wrapf(fmt.Sprintf("failed to update group %q: {{err}}", groupID), err)
//		}
//	}
//
//	d.SetId(resource.PrefixedUniqueId(fmt.Sprintf("%s-", groupID)))
//	return resourceSpotinstAWSGroupAttachmentRead(d, meta)
//}
//
//func resourceSpotinstAWSGroupAttachmentRead(d *schema.ResourceData, meta interface{}) error {
//	client := meta.(*Client)
//	service := client.elastigroup.CloudProviderAWS()
//	groupID := d.Get("group_id").(string)
//
//	// Retrieve the group properties to get list of associated load balancers.
//	input := &aws.ReadGroupInput{GroupID: spotinst.String(groupID)}
//	output, err := service.Read(context.Background(), input)
//	if err != nil {
//		return fmt.Errorf("failed to read group: %s", err)
//	}
//
//	// If nothing was found, then return no state.
//	if output.Group == nil {
//		d.SetId("")
//		return nil
//	}
//
//	var lbs []*aws.LoadBalancer
//	if lc := output.Group.Compute.LaunchSpecification; lc != nil {
//		if lc.LoadBalancersConfig != nil {
//			lbs = lc.LoadBalancersConfig.LoadBalancers
//		}
//	}
//
//	if v, ok := d.GetOk("elb"); ok {
//		found := false
//		for _, lb := range lbs {
//			if spotinst.StringValue(lb.Name) == v.(string) {
//				d.Set("elb", v.(string))
//				found = true
//				break
//			}
//		}
//		if !found {
//			log.Printf("[WARN] Association for %q was not found in group assocation", v.(string))
//			d.SetId("")
//		}
//	}
//
//	if v, ok := d.GetOk("alb_target_group_arn"); ok {
//		found := false
//		for _, lb := range lbs {
//			if spotinst.StringValue(lb.Arn) == v.(string) {
//				d.Set("alb_target_group_arn", v.(string))
//				found = true
//				break
//			}
//		}
//		if !found {
//			log.Printf("[WARN] Association for %q was not found in group assocation", v.(string))
//			d.SetId("")
//		}
//	}
//
//	return nil
//}
//
//func resourceSpotinstAWSGroupAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
//	client := meta.(*Client)
//	service := client.elastigroup.CloudProviderAWS()
//	groupID := d.Get("group_id").(string)
//
//	// Retrieve the group properties to get list of associated load balancers.
//	input := &aws.ReadGroupInput{GroupID: spotinst.String(groupID)}
//	output, err := service.Read(context.Background(), input)
//	if err != nil {
//		return fmt.Errorf("failed to read group: %s", err)
//	}
//
//	// If nothing was found, then return no state.
//	if output.Group == nil {
//		d.SetId("")
//		return nil
//	}
//
//	var oldList, newList []*aws.LoadBalancer
//	var update bool
//
//	if lc := output.Group.Compute.LaunchSpecification; lc != nil {
//		if lc.LoadBalancersConfig != nil {
//			oldList = lc.LoadBalancersConfig.LoadBalancers
//		}
//	}
//
//	if v, ok := d.GetOk("elb"); ok {
//		found := false
//		for _, lb := range oldList {
//			if spotinst.StringValue(lb.Name) == v.(string) {
//				found = true
//			} else {
//				newList = append(newList, lb)
//			}
//		}
//		if found {
//			log.Printf("[INFO] Deleting ELB %q association from group %q", v.(string), groupID)
//			update = true
//		}
//	}
//
//	if v, ok := d.GetOk("alb_target_group_arn"); ok {
//		found := false
//		for _, lb := range oldList {
//			if spotinst.StringValue(lb.Arn) == v.(string) {
//				found = true
//			} else {
//				newList = append(newList, lb)
//			}
//		}
//		if found {
//			log.Printf("[INFO] Deleting ALB Target Group %q association from group %q", v.(string), groupID)
//			update = true
//		}
//	}
//
//	if update {
//		input := &aws.UpdateGroupInput{
//			Group: &aws.Group{
//				ID: spotinst.String(groupID),
//				Compute: &aws.Compute{
//					LaunchSpecification: &aws.LaunchSpecification{
//						LoadBalancersConfig: &aws.LoadBalancersConfig{},
//					},
//				},
//			},
//		}
//		input.Group.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(newList)
//		b, err := json.MarshalIndent(input, "", "  ")
//		if err == nil {
//			log.Printf("[DEBUG] Update group payload: %s", string(b))
//		}
//		if _, err := service.Update(context.Background(), input); err != nil {
//			return errwrap.Wrapf(fmt.Sprintf("failed to update group %q: {{err}}", groupID), err)
//		}
//	}
//
//	return nil
//}

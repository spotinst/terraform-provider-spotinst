package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/multai_balancer"
)

func resourceSpotinstMultaiBalancer() *schema.Resource {
	setupMultaiBalancerResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstMultaiBalancerCreate,
		ReadContext:   resourceSpotinstMultaiBalancerRead,
		UpdateContext: resourceSpotinstMultaiBalancerUpdate,
		DeleteContext: resourceSpotinstMultaiBalancerDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.MultaiBalancerResource.GetSchemaMap(),
	}
}

func setupMultaiBalancerResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	multai_balancer.Setup(fieldsMap)

	commons.MultaiBalancerResource = commons.NewMultaiBalancerResource(fieldsMap)
}

func resourceSpotinstMultaiBalancerCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.MultaiBalancerResource.GetName())

	balancer, err := commons.MultaiBalancerResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	balancerId, err := createBalancer(balancer, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(balancerId))
	log.Printf("===> Balancer created successfully: %s <===", resourceData.Id())

	return resourceSpotinstMultaiBalancerRead(ctx, resourceData, meta)
}

func createBalancer(balancer *multai.LoadBalancer, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(balancer); err != nil {
		return nil, err
	} else {
		log.Printf("===> Balancer create configuration: %s", json)
	}

	var resp *multai.CreateLoadBalancerOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &multai.CreateLoadBalancerInput{Balancer: balancer}
		r, err := spotinstClient.multai.CreateLoadBalancer(context.Background(), input)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create balancer: %s", err)
	}

	return resp.Balancer.ID, nil
}

func resourceSpotinstMultaiBalancerRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	balancerId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.MultaiBalancerResource.GetName(), balancerId)

	input := &multai.ReadLoadBalancerInput{BalancerID: spotinst.String(balancerId)}
	resp, err := meta.(*Client).multai.ReadLoadBalancer(context.Background(), input)
	if err != nil {
		return diag.Errorf("failed to read balancer: %s", err)
	}

	// If nothing was found, return no state
	balResponse := resp.Balancer
	if balResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.MultaiBalancerResource.OnRead(balResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Balancer read successfully: %s <===", balancerId)
	return nil
}

func resourceSpotinstMultaiBalancerUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	balancerId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.MultaiBalancerResource.GetName(), balancerId)

	shouldUpdate, balancer, err := commons.MultaiBalancerResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		balancer.SetId(spotinst.String(balancerId))
		if err := updateBalancer(balancer, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("===> Balancer updated successfully: %s <===", balancerId)
	return resourceSpotinstMultaiBalancerRead(ctx, resourceData, meta)
}

func updateBalancer(balancer *multai.LoadBalancer, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &multai.UpdateLoadBalancerInput{Balancer: balancer}
	balancerId := resourceData.Id()

	if json, err := commons.ToJson(balancer); err != nil {
		return err
	} else {
		log.Printf("===> Balancer update configuration: %s", json)
	}

	if _, err := meta.(*Client).multai.UpdateLoadBalancer(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update balancer [%v]: %v", balancerId, err)
	}

	return nil
}

func resourceSpotinstMultaiBalancerDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	balancerId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.MultaiBalancerResource.GetName(), balancerId)

	if err := deleteBalancer(resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Balancer deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteBalancer(resourceData *schema.ResourceData, meta interface{}) error {
	balancerId := resourceData.Id()
	input := &multai.DeleteLoadBalancerInput{BalancerID: spotinst.String(balancerId)}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Balancer delete configuration: %s", json)
	}

	if _, err := meta.(*Client).multai.DeleteLoadBalancer(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete balancer: %s", err)
	}
	return nil
}

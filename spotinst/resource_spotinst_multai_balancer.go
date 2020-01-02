package spotinst

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/multai_balancer"
)

func resourceSpotinstMultaiBalancer() *schema.Resource {
	setupMultaiBalancerResource()

	return &schema.Resource{
		Create: resourceSpotinstMultaiBalancerCreate,
		Read:   resourceSpotinstMultaiBalancerRead,
		Update: resourceSpotinstMultaiBalancerUpdate,
		Delete: resourceSpotinstMultaiBalancerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.MultaiBalancerResource.GetSchemaMap(),
	}
}

func setupMultaiBalancerResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	multai_balancer.Setup(fieldsMap)

	commons.MultaiBalancerResource = commons.NewMultaiBalancerResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiBalancerCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.MultaiBalancerResource.GetName())

	balancer, err := commons.MultaiBalancerResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	balancerId, err := createBalancer(balancer, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(balancerId))
	log.Printf("===> Balancer created successfully: %s <===", resourceData.Id())

	return resourceSpotinstMultaiBalancerRead(resourceData, meta)
}

func createBalancer(balancer *multai.LoadBalancer, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(balancer); err != nil {
		return nil, err
	} else {
		log.Printf("===> Balancer create configuration: %s", json)
	}

	input := &multai.CreateLoadBalancerInput{Balancer: balancer}

	var resp *multai.CreateLoadBalancerOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
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

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiBalancerRead(resourceData *schema.ResourceData, meta interface{}) error {
	balancerId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.MultaiBalancerResource.GetName(), balancerId)

	input := &multai.ReadLoadBalancerInput{BalancerID: spotinst.String(balancerId)}
	resp, err := meta.(*Client).multai.ReadLoadBalancer(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to read balancer: %s", err)
	}

	// If nothing was found, return no state
	balResponse := resp.Balancer
	if balResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.MultaiBalancerResource.OnRead(balResponse, resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Balancer read successfully: %s <===", balancerId)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiBalancerUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	balancerId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.MultaiBalancerResource.GetName(), balancerId)

	shouldUpdate, balancer, err := commons.MultaiBalancerResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		balancer.SetId(spotinst.String(balancerId))
		if err := updateBalancer(balancer, resourceData, meta); err != nil {
			return err
		}
	}

	log.Printf("===> Balancer updated successfully: %s <===", balancerId)
	return resourceSpotinstMultaiBalancerRead(resourceData, meta)
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

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiBalancerDelete(resourceData *schema.ResourceData, meta interface{}) error {
	balancerId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.MultaiBalancerResource.GetName(), balancerId)

	if err := deleteBalancer(resourceData, meta); err != nil {
		return err
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

package spotinst

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/oceancd_rollout_spec"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/oceancd_rollout_spec_spot_deployment"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/oceancd_rollout_spec_strategy"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/oceancd_rollout_spec_traffic"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceSpotinstOceanCDRolloutSpec() *schema.Resource {
	setupOceanCDRolloutSpec()

	return &schema.Resource{
		CreateContext: resourceSpotinstOceanCDRolloutSpecCreate,
		ReadContext:   resourceSpotinstOceanCDRolloutSpecRead,
		UpdateContext: resourceSpotinstOceanCDRolloutSpecUpdate,
		DeleteContext: resourceSpotinstOceanCDRolloutSpecDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.OceanCDRolloutSpecResource.GetSchemaMap(),
	}
}

func setupOceanCDRolloutSpec() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)
	oceancd_rollout_spec.Setup(fieldsMap)
	oceancd_rollout_spec_spot_deployment.Setup(fieldsMap)
	oceancd_rollout_spec_strategy.Setup(fieldsMap)
	oceancd_rollout_spec_traffic.Setup(fieldsMap)

	commons.OceanCDRolloutSpecResource = commons.NewOceanCDRolloutSpecResource(fieldsMap)
}

func resourceSpotinstOceanCDRolloutSpecCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OceanCDRolloutSpecResource.GetName())

	RolloutSpec, err := commons.OceanCDRolloutSpecResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	vpname, err := createRolloutSpec(RolloutSpec, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(vpname))

	log.Printf("===> RolloutSpec created successfully: %s <===", resourceData.Id())

	return resourceSpotinstOceanCDRolloutSpecRead(ctx, resourceData, meta)
}

func createRolloutSpec(RolloutSpec *oceancd.RolloutSpec, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(RolloutSpec); err != nil {
		return nil, err
	} else {
		log.Printf("===> RolloutSpec create configuration: %s", json)
	}

	var resp *oceancd.CreateRolloutSpecOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &oceancd.CreateRolloutSpecInput{RolloutSpec: RolloutSpec}
		r, err := spotinstClient.oceancd.CreateRolloutSpec(context.Background(), input)
		if err != nil {
			log.Printf("error: %v", err)
			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create RolloutSpec: %s", err)
	}
	return resp.RolloutSpec.Name, nil
}

//end region

// region read
func resourceSpotinstOceanCDRolloutSpecRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanCDRolloutSpecResource.GetName(), name)

	RolloutSpec, err := readOceanCDRolloutSpec(context.TODO(), name, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	// If nothing was found, return no state.
	if RolloutSpec == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanCDRolloutSpecResource.OnRead(RolloutSpec, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("ocean/aks: rolloutSpec read successfully: %s", name)
	return nil
}

func readOceanCDRolloutSpec(ctx context.Context, name string, spotinstClient *Client) (*oceancd.RolloutSpec, error) {
	input := &oceancd.ReadRolloutSpecInput{
		RolloutSpecName: spotinst.String(name),
	}

	output, err := spotinstClient.oceancd.ReadRolloutSpec(ctx, input)
	if err != nil {
		// If the rolloutSpec was not found, return nil so that we can show that it
		// does not exist.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeClusterNotFound {
					return nil, nil
				}
			}
		}

		// Some other error, report it.
		return nil, fmt.Errorf("oceancd: failed to read rolloutSpec: %v", err)
	}

	return output.RolloutSpec, nil
}

// endregion

// region Update

func resourceSpotinstOceanCDRolloutSpecUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.OceanCDRolloutSpecResource.GetName(), name)

	shouldUpdate, RolloutSpec, err := commons.OceanCDRolloutSpecResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		RolloutSpec.SetName(spotinst.String(name))
		if err := updateOceanCDRolloutSpec(RolloutSpec, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("===> rolloutSpec updated successfully: %s <===", name)
	return resourceSpotinstOceanCDRolloutSpecRead(ctx, resourceData, meta)
}

func updateOceanCDRolloutSpec(RolloutSpec *oceancd.RolloutSpec, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &oceancd.PatchRolloutSpecInput{
		RolloutSpec: RolloutSpec,
	}

	name := resourceData.Id()

	if json, err := commons.ToJson(RolloutSpec); err != nil {
		return err
	} else {
		log.Printf("===> stratgey update configuration: %s", json)
	}

	if _, err := meta.(*Client).oceancd.PatchRolloutSpec(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update rolloutSpec [%v]: %v", name, err)
	}
	return nil
}

//end region

//region Delete

func resourceSpotinstOceanCDRolloutSpecDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanCDRolloutSpecResource.GetName(), name)

	if err := deleteOceanCDRolloutSpec(resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> rolloutSpec deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteOceanCDRolloutSpec(resourceData *schema.ResourceData, meta interface{}) error {
	name := resourceData.Id()
	input := &oceancd.DeleteRolloutSpecInput{
		RolloutSpecName: spotinst.String(name),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> rolloutSpec delete configuration: %s", json)
	}

	if _, err := meta.(*Client).oceancd.DeleteRolloutSpec(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete rolloutSpec: %s", err)
	}
	return nil
}

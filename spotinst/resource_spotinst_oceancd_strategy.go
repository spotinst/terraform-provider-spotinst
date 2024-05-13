package spotinst

import (
	"context"
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/oceancd_verification_provider"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/oceancd_verification_provider_cloud_watch"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/oceancd_verification_provider_datadog"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/oceancd_verification_provider_jenkins"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/oceancd_verification_provider_new_relic"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/oceancd_verification_provider_prometheus"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceSpotinstOceanCDStrategy() *schema.Resource {
	setupOceanCDStrategy()

	return &schema.Resource{
		CreateContext: resourceSpotinstOceanCDStrategyCreate,
		ReadContext:   resourceSpotinstOceanCDStrategyRead,
		UpdateContext: resourceSpotinstOceanCDStrategyUpdate,
		DeleteContext: resourceSpotinstOceanCDStrategyDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.OceanCDStrategyResource.GetSchemaMap(),
	}
}

func setupOceanCDStrategy() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	oceancd_verification_provider.Setup(fieldsMap)
	oceancd_verification_provider_cloud_watch.Setup(fieldsMap)
	oceancd_verification_provider_datadog.Setup(fieldsMap)
	oceancd_verification_provider_jenkins.Setup(fieldsMap)
	oceancd_verification_provider_new_relic.Setup(fieldsMap)
	oceancd_verification_provider_prometheus.Setup(fieldsMap)

	commons.OceanCDStrategyResource = commons.NewOceanCDStrategyResource(fieldsMap)
}

func resourceSpotinstOceanCDStrategyCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OceanCDStrategyResource.GetName())

	Strategy, err := commons.OceanCDStrategyResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	vpname, err := createStrategy(Strategy, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(vpname))

	log.Printf("===> Strategy created successfully: %s <===", resourceData.Id())

	return resourceSpotinstOceanCDStrategyRead(ctx, resourceData, meta)
}

func createStrategy(Strategy *oceancd.Strategy, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(Strategy); err != nil {
		return nil, err
	} else {
		log.Printf("===> Strategy create configuration: %s", json)
	}

	var resp *oceancd.CreateStrategyOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &oceancd.CreateStrategyInput{Strategy: Strategy}
		r, err := spotinstClient.oceancd.CreateStrategy(context.Background(), input)
		if err != nil {
			log.Printf("error: %v", err)
			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create Strategy: %s", err)
	}
	return resp.Strategy.Name, nil
}

//end region

//region read
func resourceSpotinstOceanCDStrategyRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanCDStrategyResource.GetName(), name)

	Strategy, err := readOceanCDStrategy(context.TODO(), name, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	// If nothing was found, return no state.
	if Strategy == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanCDStrategyResource.OnRead(Strategy, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("ocean/aks: strategy read successfully: %s", name)
	return nil
}

func readOceanCDStrategy(ctx context.Context, name string, spotinstClient *Client) (*oceancd.Strategy, error) {
	input := &oceancd.ReadStrategyInput{
		StrategyName: spotinst.String(name),
	}

	output, err := spotinstClient.oceancd.ReadStrategy(ctx, input)
	if err != nil {
		// If the strategy was not found, return nil so that we can show that it
		// does not exist.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeClusterNotFound {
					return nil, nil
				}
			}
		}

		// Some other error, report it.
		return nil, fmt.Errorf("oceancd: failed to read strategy: %v", err)
	}

	return output.Strategy, nil
}

// endregion

//region Update

func resourceSpotinstOceanCDStrategyUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.OceanCDStrategyResource.GetName(), name)

	shouldUpdate, Strategy, err := commons.OceanCDStrategyResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		Strategy.SetName(spotinst.String(name))
		if err := updateOceanCDStrategy(Strategy, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("===> strategy updated successfully: %s <===", name)
	return resourceSpotinstOceanCDStrategyRead(ctx, resourceData, meta)
}

func updateOceanCDStrategy(Strategy *oceancd.Strategy, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &oceancd.UpdateStrategyInput{
		Strategy: Strategy,
	}

	name := resourceData.Id()

	if json, err := commons.ToJson(Strategy); err != nil {
		return err
	} else {
		log.Printf("===> stratgey update configuration: %s", json)
	}

	if _, err := meta.(*Client).oceancd.UpdateStrategy(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update strategy [%v]: %v", name, err)
	}
	return nil
}

//end region

//region Delete

func resourceSpotinstOceanCDStrategyDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanCDStrategyResource.GetName(), name)

	if err := deleteOceanCDStrategy(resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> strategy deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteOceanCDStrategy(resourceData *schema.ResourceData, meta interface{}) error {
	name := resourceData.Id()
	input := &oceancd.DeleteStrategyInput{
		StrategyName: spotinst.String(name),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> strategy delete configuration: %s", json)
	}

	if _, err := meta.(*Client).oceancd.DeleteStrategy(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete strategy: %s", err)
	}
	return nil
}

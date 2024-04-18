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

func resourceSpotinstOceanCDVerificationProvider() *schema.Resource {
	setupOceanCDVerificationProvider()

	return &schema.Resource{
		CreateContext: resourceSpotinstOceanCDVerificationProviderCreate,
		ReadContext:   resourceSpotinstOceanCDVerificationProviderRead,
		UpdateContext: resourceSpotinstOceanCDVerificationProviderUpdate,
		DeleteContext: resourceSpotinstOceanCDVerificationProviderDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.OceanCDVerificationProviderResource.GetSchemaMap(),
	}
}

func setupOceanCDVerificationProvider() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	oceancd_verification_provider.Setup(fieldsMap)
	oceancd_verification_provider_cloud_watch.Setup(fieldsMap)
	oceancd_verification_provider_datadog.Setup(fieldsMap)
	oceancd_verification_provider_jenkins.Setup(fieldsMap)
	oceancd_verification_provider_new_relic.Setup(fieldsMap)
	oceancd_verification_provider_prometheus.Setup(fieldsMap)

	commons.OceanCDVerificationProviderResource = commons.NewOceanCDVerificationProviderResource(fieldsMap)
}

func resourceSpotinstOceanCDVerificationProviderCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OceanCDVerificationProviderResource.GetName())

	verificationProvider, err := commons.OceanCDVerificationProviderResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	vpname, err := createVerificationProvider(verificationProvider, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(vpname))

	log.Printf("===> Verification Provider created successfully: %s <===", resourceData.Id())

	return resourceSpotinstOceanCDVerificationProviderRead(ctx, resourceData, meta)
}

func createVerificationProvider(verificationProvider *oceancd.VerificationProvider, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(verificationProvider); err != nil {
		return nil, err
	} else {
		log.Printf("===> Verification Provider create configuration: %s", json)
	}

	var resp *oceancd.CreateVerificationProviderOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &oceancd.CreateVerificationProviderInput{VerificationProvider: verificationProvider}
		r, err := spotinstClient.oceancd.CreateVerificationProvider(context.Background(), input)
		if err != nil {
			log.Printf("error: %v", err)
			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create Verification Provider: %s", err)
	}
	return resp.VerificationProvider.Name, nil
}

//end region

//region read
func resourceSpotinstOceanCDVerificationProviderRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanCDVerificationProviderResource.GetName(), name)

	verificationProvider, err := readOceanCDVerificationProvider(context.TODO(), name, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	// If nothing was found, return no state.
	if verificationProvider == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanCDVerificationProviderResource.OnRead(verificationProvider, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("ocean/aks: verification provider read successfully: %s", name)
	return nil
}

func readOceanCDVerificationProvider(ctx context.Context, name string, spotinstClient *Client) (*oceancd.VerificationProvider, error) {
	input := &oceancd.ReadVerificationProviderInput{
		Name: spotinst.String(name),
	}

	output, err := spotinstClient.oceancd.ReadVerificationProvider(ctx, input)
	if err != nil {
		// If the cluster was not found, return nil so that we can show that it
		// does not exist.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeClusterNotFound {
					return nil, nil
				}
			}
		}

		// Some other error, report it.
		return nil, fmt.Errorf("ocean/aks: failed to read verification provider: %v", err)
	}

	return output.VerificationProvider, nil
}

// endregion

//region Update

func resourceSpotinstOceanCDVerificationProviderUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.OceanCDVerificationProviderResource.GetName(), name)

	shouldUpdate, verificationProvider, err := commons.OceanCDVerificationProviderResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		verificationProvider.SetName(spotinst.String(name))
		if err := updateOceanCDVerificationProvider(verificationProvider, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("===> Verification Provider updated successfully: %s <===", name)
	return resourceSpotinstOceanCDVerificationProviderRead(ctx, resourceData, meta)
}

func updateOceanCDVerificationProvider(verificationProvider *oceancd.VerificationProvider, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &oceancd.UpdateVerificationProviderInput{
		VerificationProvider: verificationProvider,
	}

	name := resourceData.Id()

	if json, err := commons.ToJson(verificationProvider); err != nil {
		return err
	} else {
		log.Printf("===> Verification Provider update configuration: %s", json)
	}

	if _, err := meta.(*Client).oceancd.UpdateVerificationProvider(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update Verification Provider [%v]: %v", name, err)
	}
	return nil
}

//end region

//region Delete

func resourceSpotinstOceanCDVerificationProviderDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanCDVerificationProviderResource.GetName(), name)

	if err := deleteOceanCDVerificationProvider(resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Verification Provider deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteOceanCDVerificationProvider(resourceData *schema.ResourceData, meta interface{}) error {
	name := resourceData.Id()
	input := &oceancd.DeleteVerificationProviderInput{
		Name: spotinst.String(name),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Verification Provider delete configuration: %s", json)
	}

	if _, err := meta.(*Client).oceancd.DeleteVerificationProvider(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete verification provider: %s", err)
	}
	return nil
}

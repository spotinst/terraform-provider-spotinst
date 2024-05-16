package spotinst

import (
	"context"
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/oceancd_verification_template"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/oceancd_verification_template_args"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/oceancd_verification_template_metrics"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceSpotinstOceanCDVerificationTemplate() *schema.Resource {
	setupOceanCDVerificationTemplate()

	return &schema.Resource{
		CreateContext: resourceSpotinstOceanCDVerificationTemplateCreate,
		ReadContext:   resourceSpotinstOceanCDVerificationTemplateRead,
		UpdateContext: resourceSpotinstOceanCDVerificationTemplateUpdate,
		DeleteContext: resourceSpotinstOceanCDVerificationTemplateDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.OceanCDVerificationTemplateResource.GetSchemaMap(),
	}
}

func setupOceanCDVerificationTemplate() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	oceancd_verification_template.Setup(fieldsMap)
	oceancd_verification_template_args.Setup(fieldsMap)
	oceancd_verification_template_metrics.Setup(fieldsMap)

	commons.OceanCDVerificationTemplateResource = commons.NewOceanCDVerificationTemplateResource(fieldsMap)
}

func resourceSpotinstOceanCDVerificationTemplateCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OceanCDVerificationTemplateResource.GetName())

	VerificationTemplate, err := commons.OceanCDVerificationTemplateResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	vtname, err := createVerificationTemplate(VerificationTemplate, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(vtname))

	log.Printf("===> Verification Template created successfully: %s <===", resourceData.Id())

	return resourceSpotinstOceanCDVerificationTemplateRead(ctx, resourceData, meta)
}

func createVerificationTemplate(VerificationTemplate *oceancd.VerificationTemplate, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(VerificationTemplate); err != nil {
		return nil, err
	} else {
		log.Printf("===> Verification Template create configuration: %s", json)
	}

	var resp *oceancd.CreateVerificationTemplateOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &oceancd.CreateVerificationTemplateInput{VerificationTemplate: VerificationTemplate}
		r, err := spotinstClient.oceancd.CreateVerificationTemplate(context.Background(), input)
		if err != nil {
			log.Printf("error: %v", err)
			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create Verification Template: %s", err)
	}
	return resp.VerificationTemplate.Name, nil
}

//end region

// region read
func resourceSpotinstOceanCDVerificationTemplateRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanCDVerificationTemplateResource.GetName(), name)

	VerificationTemplate, err := readOceanCDVerificationTemplate(context.TODO(), name, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	// If nothing was found, return no state.
	if VerificationTemplate == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanCDVerificationTemplateResource.OnRead(VerificationTemplate, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("ocean/aks: verification template read successfully: %s", name)
	return nil
}

func readOceanCDVerificationTemplate(ctx context.Context, name string, spotinstClient *Client) (*oceancd.VerificationTemplate, error) {
	input := &oceancd.ReadVerificationTemplateInput{
		Name: spotinst.String(name),
	}

	output, err := spotinstClient.oceancd.ReadVerificationTemplate(ctx, input)
	if err != nil {
		// If the Verification Template was not found, return nil so that we can show that it
		// does not exist.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeClusterNotFound {
					return nil, nil
				}
			}
		}

		// Some other error, report it.
		return nil, fmt.Errorf("oceancd: failed to read verification template: %v", err)
	}

	return output.VerificationTemplate, nil
}

// endregion

//region Update

func resourceSpotinstOceanCDVerificationTemplateUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.OceanCDVerificationTemplateResource.GetName(), name)

	shouldUpdate, VerificationTemplate, err := commons.OceanCDVerificationTemplateResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		VerificationTemplate.SetName(spotinst.String(name))
		if err := updateOceanCDVerificationTemplate(VerificationTemplate, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("===> Verification Template updated successfully: %s <===", name)
	return resourceSpotinstOceanCDVerificationTemplateRead(ctx, resourceData, meta)
}

func updateOceanCDVerificationTemplate(VerificationTemplate *oceancd.VerificationTemplate, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &oceancd.UpdateVerificationTemplateInput{
		VerificationTemplate: VerificationTemplate,
	}

	name := resourceData.Id()

	if json, err := commons.ToJson(VerificationTemplate); err != nil {
		return err
	} else {
		log.Printf("===> Verification Template update configuration: %s", json)
	}

	if _, err := meta.(*Client).oceancd.UpdateVerificationTemplate(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update Verification Template [%v]: %v", name, err)
	}
	return nil
}

//end region

//region Delete

func resourceSpotinstOceanCDVerificationTemplateDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanCDVerificationTemplateResource.GetName(), name)

	if err := deleteOceanCDVerificationTemplate(resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Verification Template deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteOceanCDVerificationTemplate(resourceData *schema.ResourceData, meta interface{}) error {
	name := resourceData.Id()
	input := &oceancd.DeleteVerificationTemplateInput{
		Name: spotinst.String(name),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Verification Template delete configuration: %s", json)
	}

	if _, err := meta.(*Client).oceancd.DeleteVerificationTemplate(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete verification template: %s", err)
	}
	return nil
}

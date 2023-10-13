package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/organization"
	organizationPackage "github.com/spotinst/terraform-provider-spotinst/spotinst/organization_policy"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func resourceOrgPolicy() *schema.Resource {
	setupOrgPolicy()
	return &schema.Resource{
		CreateContext: resourceOrgPolicyCreate,
		UpdateContext: resourceOrgPolicyUpdate,
		ReadContext:   resourceOrgPolicyRead,
		DeleteContext: resourceOrgPolicyDelete,

		Schema: commons.OrgPolicyResource.GetSchemaMap(),
	}
}

func setupOrgPolicy() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	organizationPackage.Setup(fieldsMap)

	commons.OrgPolicyResource = commons.NewOrgPolicyResource(fieldsMap)
}

func resourceOrgPolicyDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OrgPolicyResource.GetName(), id)

	input := &organization.DeletePolicyInput{PolicyID: spotinst.String(id)}
	if _, err := meta.(*Client).organization.DeletePolicy(context.Background(), input); err != nil {
		return diag.Errorf("[ERROR] Failed to delete policy: %s", err)
	}

	resourceData.SetId("")
	return nil
}

func resourceOrgPolicyRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.OrgPolicyResource.GetName(), id)

	client := meta.(*Client)
	input := &organization.ReadPolicyInput{PolicyID: spotinst.String(resourceData.Id())}
	policyResponse, err := client.organization.ReadPolicy(context.Background(), input)
	if err != nil {
		return diag.Errorf("[ERROR] Failed to read policy: %s", err)
	}

	// If nothing was found, then return no state.
	policy := policyResponse.Policy
	if policy == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OrgPolicyResource.OnRead(policy, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("===> Policy read successfully: %s <===", id)
	return nil
}

func resourceOrgPolicyCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OrgPolicyResource.GetName())

	policy, err := commons.OrgPolicyResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	policyId, err := createPolicy(policy, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(policyId))
	log.Printf("===> Policy created successfully: %s <===", resourceData.Id())

	return resourceOrgPolicyRead(ctx, resourceData, meta)
}

func createPolicy(policyObj *organization.Policy, spotinstClient *Client) (*string, error) {
	input := &organization.CreatePolicyInput{
		Policy: policyObj,
	}
	resp, err := spotinstClient.organization.CreatePolicy(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create policy: %s", err)
	}
	return resp.Policy.PolicyID, nil
}

func resourceOrgPolicyUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.OrgPolicyResource.GetName(), id)

	shouldUpdate, policy, err := commons.OrgPolicyResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		policy.PolicyID = spotinst.String(id)
		if err := updatePolicy(policy, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("===> Policy updated successfully: %s <===", id)
	return resourceOrgPolicyRead(ctx, resourceData, meta)
}

func updatePolicy(policy *organization.Policy, resourceData *schema.ResourceData, meta interface{}) error {
	input := &organization.UpdatePolicyInput{
		Policy: policy,
	}
	if json, err := commons.ToJson(policy); err != nil {
		return err
	} else {
		log.Printf("===> policy update configuration: %s", json)
	}

	if _, err := meta.(*Client).organization.UpdatePolicy(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] failed to update policy %s: %s", resourceData.Id(), err)
	}
	return nil
}

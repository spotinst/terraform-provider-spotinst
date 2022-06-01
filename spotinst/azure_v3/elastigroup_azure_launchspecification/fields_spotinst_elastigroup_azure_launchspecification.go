package elastigroup_azure_launchspecification

import (
	"fmt"

	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[CustomData] = commons.NewGenericField(
		commons.ElastigroupAzureLaunchSpecification,
		CustomData,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.CustomData != nil {
				value = elastigroup.Compute.LaunchSpecification.CustomData
			}
			if err := resourceData.Set(string(CustomData), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CustomData), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(CustomData)).(string); ok && v != "" {
				customData := spotinst.String(v)
				elastigroup.Compute.LaunchSpecification.SetCustomData(customData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var customData *string = nil
			if v, ok := resourceData.Get(string(CustomData)).(string); ok && v != "" {
				customData = spotinst.String(v)
			}
			elastigroup.Compute.LaunchSpecification.SetCustomData(customData)
			return nil
		},
		nil,
	)

	fieldsMap[ManagedServiceIdentity] = commons.NewGenericField(
		commons.ElastigroupAzureLaunchSpecification,
		ManagedServiceIdentity,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ManagedServiceIdentityResourceGroupName): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(ManagedServiceIdentityName): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{} = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.ManagedServiceIdentities != nil {
				value = flattenManagedServiceIdentities(elastigroup.Compute.LaunchSpecification.ManagedServiceIdentities)
			}
			if err := resourceData.Set(string(ManagedServiceIdentity), value); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(ManagedServiceIdentity), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(ManagedServiceIdentity)); ok {
				if msis, err := expandManagedServiceIdentities(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetManagedServiceIdentities(msis)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []*azurev3.ManagedServiceIdentity = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil {
				if v, ok := resourceData.GetOk(string(ManagedServiceIdentity)); ok {
					if msis, err := expandManagedServiceIdentities(v); err != nil {
						return err
					} else {
						value = msis
					}
				}
				elastigroup.Compute.LaunchSpecification.SetManagedServiceIdentities(value)
			}
			return nil
		},
		nil,
	)
}

func expandManagedServiceIdentities(data interface{}) ([]*azurev3.ManagedServiceIdentity, error) {
	list := data.(*schema.Set).List()
	msis := make([]*azurev3.ManagedServiceIdentity, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		msis = append(msis, &azurev3.ManagedServiceIdentity{
			ResourceGroupName: spotinst.String(attr[string(ManagedServiceIdentityResourceGroupName)].(string)),
			Name:              spotinst.String(attr[string(ManagedServiceIdentityName)].(string)),
		})
	}
	return msis, nil
}

func flattenManagedServiceIdentities(msis []*azurev3.ManagedServiceIdentity) []interface{} {
	result := make([]interface{}, 0, len(msis))
	for _, msi := range msis {
		m := make(map[string]interface{})
		m[string(ManagedServiceIdentityResourceGroupName)] = spotinst.StringValue(msi.ResourceGroupName)
		m[string(ManagedServiceIdentityName)] = spotinst.StringValue(msi.Name)
		result = append(result, m)
	}
	return result
}

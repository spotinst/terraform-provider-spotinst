package elastigroup_azure_launchspecification

import (
	"errors"
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

	fieldsMap[Tags] = commons.NewGenericField(
		commons.ElastigroupAzure,
		Tags,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(TagKey): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(TagValue): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Tags != nil {
				tags := elastigroup.Compute.LaunchSpecification.Tags
				result = flattenTags(tags)
			}
			if result != nil {
				if err := resourceData.Set(string(Tags), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tags), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var tagsToAdd []*azurev3.Tags = nil
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					tagsToAdd = tags
				}
			}
			elastigroup.Compute.LaunchSpecification.SetTags(tagsToAdd)
			return nil
		},
		nil,
	)

	fieldsMap[ShutdownScript] = commons.NewGenericField(
		commons.ElastigroupAzureLaunchSpecification,
		ShutdownScript,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup != nil && elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil && elastigroup.Compute.LaunchSpecification.ShutdownScript != nil {
				value = elastigroup.Compute.LaunchSpecification.ShutdownScript
			}
			if err := resourceData.Set(string(ShutdownScript), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShutdownScript), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(ShutdownScript)).(string); ok && v != "" {
				resourceGroupName := spotinst.String(v)
				elastigroup.Compute.LaunchSpecification.SetShutdownScript(resourceGroupName)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if v, ok := resourceData.Get(string(ShutdownScript)).(string); ok && v != "" {
				shutDown := spotinst.String(v)
				value = shutDown
			}
			elastigroup.Compute.LaunchSpecification.SetShutdownScript(value)
			return nil
		},
		nil,
	)

	fieldsMap[UserData] = commons.NewGenericField(
		commons.ElastigroupAzureLaunchSpecification,
		UserData,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup != nil && elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil && elastigroup.Compute.LaunchSpecification.UserData != nil {
				value = elastigroup.Compute.LaunchSpecification.UserData
			}
			if err := resourceData.Set(string(UserData), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UserData), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData := spotinst.String(v)
				elastigroup.Compute.LaunchSpecification.SetUserData(userData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData := spotinst.String(v)
				value = userData
			}
			elastigroup.Compute.LaunchSpecification.SetUserData(value)
			return nil
		},
		nil,
	)

	fieldsMap[ProximityPlacementGroups] = commons.NewGenericField(
		commons.ElastigroupAzureLaunchSpecification,
		ProximityPlacementGroups,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(PPGName): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(PPGResourceGroupName): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{}
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil && elastigroup.Compute.LaunchSpecification.ProximityPlacementGroups != nil {
				value = flattenProximityPlacementGroups(elastigroup.Compute.LaunchSpecification.ProximityPlacementGroups)
			}
			if value != nil {
				if err := resourceData.Set(string(ProximityPlacementGroups), value); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(ProximityPlacementGroups), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(ProximityPlacementGroups)); ok {
				if ppgs, err := expandProximityPlacementGroups(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetProximityPlacementGroups(ppgs)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []*azurev3.ProximityPlacementGroups = nil

			if v, ok := resourceData.GetOk(string(ProximityPlacementGroups)); ok {
				if ppgs, err := expandProximityPlacementGroups(v); err != nil {
					return err
				} else {
					value = ppgs
				}
				elastigroup.Compute.LaunchSpecification.SetProximityPlacementGroups(value)
			} else {
				elastigroup.Compute.LaunchSpecification.SetProximityPlacementGroups(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[OsDisk] = commons.NewGenericField(
		commons.ElastigroupAzureLaunchSpecification,
		OsDisk,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(OsDiskSizeGB): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
					string(OsDiskType): {
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
				elastigroup.Compute.LaunchSpecification.OsDisk != nil {
				value = flattenOSDisk(elastigroup.Compute.LaunchSpecification.OsDisk)
			}
			if value != nil {
				if err := resourceData.Set(string(OsDisk), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OsDisk), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(OsDisk)); ok {
				if osDisk, err := expandOSDisk(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetOsDisk(osDisk)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *azurev3.OsDisk = nil
			if v, ok := resourceData.GetOk(string(OsDisk)); ok {
				if osDisk, err := expandOSDisk(v); err != nil {
					return err
				} else {
					value = osDisk
				}
				elastigroup.Compute.LaunchSpecification.SetOsDisk(value)
			} else {
				elastigroup.Compute.LaunchSpecification.SetOsDisk(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[DataDisk] = commons.NewGenericField(
		commons.ElastigroupAzureLaunchSpecification,
		DataDisk,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(DataDiskSizeGB): {
						Type:     schema.TypeInt,
						Required: true,
					},
					string(DataDiskLUN): {
						Type:     schema.TypeInt,
						Required: true,
					},
					string(DataDiskType): {
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

			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil && elastigroup.Compute.LaunchSpecification.DataDisks != nil {
				dataDisks := elastigroup.Compute.LaunchSpecification.DataDisks
				value = flattenDataDisks(dataDisks)
			}
			if value != nil {
				if err := resourceData.Set(string(DataDisk), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DataDisk), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(DataDisk)); ok {
				if value, err := expandDataDisks(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetDataDisks(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []*azurev3.DataDisks = nil

			if v, ok := resourceData.GetOk(string(DataDisk)); ok {
				if dataDisks, err := expandDataDisks(v); err != nil {
					return err
				} else {
					value = dataDisks
				}
				elastigroup.Compute.LaunchSpecification.SetDataDisks(value)
			} else {
				elastigroup.Compute.LaunchSpecification.SetDataDisks(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[BootDiagnostics] = commons.NewGenericField(
		commons.ElastigroupAzureLaunchSpecification,
		BootDiagnostics,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(BootDiagnosticsIsEnabled): {
						Type:     schema.TypeBool,
						Required: true,
					},
					string(BootDiagnosticsType): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(BootDiagnosticsStorageURL): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value interface{} = nil

			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.BootDiagnostics != nil {
				value = flattenBootDiagnostics(elastigroup.Compute.LaunchSpecification.BootDiagnostics)
			}
			if err := resourceData.Set(string(BootDiagnostics), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BootDiagnostics), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *azurev3.BootDiagnostics = nil

			if v, ok := resourceData.GetOk(string(BootDiagnostics)); ok {
				if bd, err := expandBootDiagnostics(v); err != nil {
					return err
				} else {
					value = bd
				}
			}
			elastigroup.Compute.LaunchSpecification.SetBootDiagnostics(value)

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *azurev3.BootDiagnostics = nil

			if v, ok := resourceData.GetOk(string(BootDiagnostics)); ok {
				if bd, err := expandBootDiagnostics(v); err != nil {
					return err
				} else {
					value = bd
				}
			}
			elastigroup.Compute.LaunchSpecification.SetBootDiagnostics(value)
			return nil
		},
		nil,
	)

	fieldsMap[VmNamePrefix] = commons.NewGenericField(
		commons.ElastigroupAzureLaunchSpecification,
		VmNamePrefix,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup != nil && elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil && elastigroup.Compute.LaunchSpecification.VmNamePrefix != nil {
				value = elastigroup.Compute.LaunchSpecification.VmNamePrefix
			}
			if err := resourceData.Set(string(VmNamePrefix), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(VmNamePrefix), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(VmNamePrefix)).(string); ok && v != "" {
				vmNamePrefix := spotinst.String(v)
				elastigroup.Compute.LaunchSpecification.SetVmNamePrefix(vmNamePrefix)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if v, ok := resourceData.Get(string(VmNamePrefix)).(string); ok && v != "" {
				vmNamePrefix := spotinst.String(v)
				value = vmNamePrefix
			}
			elastigroup.Compute.LaunchSpecification.SetVmNamePrefix(value)
			return nil
		},
		nil,
	)

	fieldsMap[Security] = commons.NewGenericField(
		commons.ElastigroupAzureLaunchSpecification,
		Security,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(SecureBootEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(SecurityType): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(VTpmEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(ConfidentialOsDiskEncryption): {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value interface{} = nil

			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Security != nil {
				value = flattenSecurity(elastigroup.Compute.LaunchSpecification.Security)
			}
			if err := resourceData.Set(string(Security), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Security), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *azurev3.Security = nil

			if v, ok := resourceData.GetOk(string(Security)); ok {
				if security, err := expandSecurity(v); err != nil {
					return err
				} else {
					value = security
				}
			}
			elastigroup.Compute.LaunchSpecification.SetSecurity(value)

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *azurev3.Security = nil

			if v, ok := resourceData.GetOk(string(Security)); ok {
				if security, err := expandSecurity(v); err != nil {
					return err
				} else {
					value = security
				}
			}
			elastigroup.Compute.LaunchSpecification.SetSecurity(value)
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

func flattenTags(tags []*azurev3.Tags) []interface{} {
	result := make([]interface{}, 0, len(tags))
	for _, tag := range tags {
		m := make(map[string]interface{})
		m[string(TagKey)] = spotinst.StringValue(tag.TagKey)
		m[string(TagValue)] = spotinst.StringValue(tag.TagValue)

		result = append(result, m)
	}
	return result
}

func expandTags(data interface{}) ([]*azurev3.Tags, error) {
	list := data.(*schema.Set).List()
	tags := make([]*azurev3.Tags, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(TagKey)]; !ok {
			return nil, errors.New("invalid tag attributes: key missing")
		}

		if _, ok := attr[string(TagValue)]; !ok {
			return nil, errors.New("invalid tag attributes: value missing")
		}
		tag := &azurev3.Tags{
			TagKey:   spotinst.String(attr[string(TagKey)].(string)),
			TagValue: spotinst.String(attr[string(TagValue)].(string)),
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func expandProximityPlacementGroups(data interface{}) ([]*azurev3.ProximityPlacementGroups, error) {
	list := data.([]interface{})
	ppgs := make([]*azurev3.ProximityPlacementGroups, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		ProximityPlacementGroup := &azurev3.ProximityPlacementGroups{}
		if v, ok := attr[string(PPGName)].(string); ok && v != "" {
			ProximityPlacementGroup.SetName(spotinst.String(v))
		}
		if v, ok := attr[string(PPGResourceGroupName)].(string); ok && v != "" {
			ProximityPlacementGroup.SetResourceGroupName(spotinst.String(v))
		}
		ppgs = append(ppgs, ProximityPlacementGroup)
	}
	return ppgs, nil
}

func flattenProximityPlacementGroups(ppgs []*azurev3.ProximityPlacementGroups) []interface{} {
	var result []interface{}
	for _, ppg := range ppgs {
		m := make(map[string]interface{})
		m[string(PPGResourceGroupName)] = spotinst.StringValue(ppg.ResourceGroupName)
		m[string(PPGName)] = spotinst.StringValue(ppg.Name)
		result = append(result, m)
	}
	return result
}

func flattenOSDisk(osd *azurev3.OsDisk) []interface{} {
	osDisk := make(map[string]interface{})
	osDisk[string(OsDiskSizeGB)] = spotinst.IntValue(osd.SizeGB)
	osDisk[string(OsDiskType)] = spotinst.StringValue(osd.Type)
	return []interface{}{osDisk}
}

func expandOSDisk(data interface{}) (*azurev3.OsDisk, error) {
	osDisk := &azurev3.OsDisk{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return osDisk, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(OsDiskSizeGB)].(int); ok {
		if v == -1 {
			osDisk.SetSizeGB(nil)
		} else {
			osDisk.SetSizeGB(spotinst.Int(v))
		}
	}
	if v, ok := m[string(OsDiskType)].(string); ok && v != "" {
		osDisk.SetType(spotinst.String(v))
	} else {
		osDisk.SetType(nil)
	}
	return osDisk, nil
}

func flattenDataDisks(dataDisks []*azurev3.DataDisks) []interface{} {
	var result []interface{}

	for _, disk := range dataDisks {
		m := make(map[string]interface{})
		m[string(DataDiskSizeGB)] = spotinst.IntValue(disk.SizeGB)
		m[string(DataDiskLUN)] = spotinst.IntValue(disk.Lun)
		m[string(DataDiskType)] = spotinst.StringValue(disk.Type)
		result = append(result, m)
	}
	return result
}

func expandDataDisks(data interface{}) ([]*azurev3.DataDisks, error) {
	list := data.([]interface{})
	dd := make([]*azurev3.DataDisks, 0, len(list))
	for _, m := range list {
		attr, ok := m.(map[string]interface{})
		if !ok {
			continue
		}

		dataDisk := &azurev3.DataDisks{}
		if v, ok := attr[string(DataDiskSizeGB)].(int); ok && v > 0 {
			dataDisk.SetSizeGB(spotinst.Int(v))
		}
		if v, ok := attr[string(DataDiskLUN)].(int); ok && v >= 0 {
			dataDisk.SetLun(spotinst.Int(v))
		}
		if v, ok := attr[string(DataDiskType)].(string); ok && v != "" {
			dataDisk.SetType(spotinst.String(v))
		}
		dd = append(dd, dataDisk)
	}
	return dd, nil
}

func flattenBootDiagnostics(bd *azurev3.BootDiagnostics) interface{} {
	bootDiagnostic := make(map[string]interface{})
	bootDiagnostic[string(BootDiagnosticsIsEnabled)] = spotinst.BoolValue(bd.IsEnabled)
	bootDiagnostic[string(BootDiagnosticsType)] = spotinst.StringValue(bd.Type)
	bootDiagnostic[string(BootDiagnosticsStorageURL)] = spotinst.StringValue(bd.StorageUri)
	return []interface{}{bootDiagnostic}
}

func expandBootDiagnostics(data interface{}) (*azurev3.BootDiagnostics, error) {
	if list := data.([]interface{}); len(list) > 0 {
		bootDiagnostic := &azurev3.BootDiagnostics{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})
			var enabled *bool = nil
			var bsType *string = nil
			var storageURL *string = nil

			if v, ok := m[string(BootDiagnosticsIsEnabled)].(bool); ok {
				enabled = spotinst.Bool(v)
			}
			bootDiagnostic.SetIsEnabled(enabled)

			if v, ok := m[string(BootDiagnosticsType)].(string); ok && v != "" {
				bsType = spotinst.String(v)
			}
			bootDiagnostic.SetType(bsType)

			if v, ok := m[string(BootDiagnosticsStorageURL)].(string); ok && v != "" {
				storageURL = spotinst.String(v)
			}
			bootDiagnostic.SetStorageUri(storageURL)

		}

		return bootDiagnostic, nil
	}

	return nil, nil
}

func flattenSecurity(secure *azurev3.Security) interface{} {
	security := make(map[string]interface{})

	security[string(SecureBootEnabled)] = spotinst.BoolValue(secure.SecureBootEnabled)
	security[string(SecurityType)] = spotinst.StringValue(secure.SecurityType)
	security[string(VTpmEnabled)] = spotinst.BoolValue(secure.VTpmEnabled)
	security[string(ConfidentialOsDiskEncryption)] = spotinst.BoolValue(secure.ConfidentialOsDiskEncryption)

	return []interface{}{security}
}

func expandSecurity(data interface{}) (*azurev3.Security, error) {
	if list := data.([]interface{}); len(list) > 0 {
		security := &azurev3.Security{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(SecureBootEnabled)].(bool); ok {
				security.SetSecureBootEnabled(spotinst.Bool(v))
			}
			if v, ok := m[string(SecurityType)].(string); ok && v != "" {
				security.SetSecurityType(spotinst.String(v))
			}
			if v, ok := m[string(VTpmEnabled)].(bool); ok {
				security.SetVTpmEnabled(spotinst.Bool(v))
			}
			if v, ok := m[string(ConfidentialOsDiskEncryption)].(bool); ok {
				security.SetConfidentialOsDiskEncryption(spotinst.Bool(v))
			}
		}
		return security, nil
	}
	return nil, nil
}

package stateful_node_azure_launch_spec

import (
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[CustomData] = commons.NewGenericField(
		commons.StatefulNodeAzureLaunchSpecification,
		CustomData,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil
			if st != nil && st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.CustomData != nil {
				value = st.Compute.LaunchSpecification.CustomData
			}
			if err := resourceData.Set(string(CustomData), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CustomData), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(CustomData)).(string); ok && v != "" {
				customData := spotinst.String(base64Encode(v))
				st.Compute.LaunchSpecification.SetCustomData(customData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil
			if v, ok := resourceData.Get(string(CustomData)).(string); ok && v != "" {
				customData := spotinst.String(base64Encode(v))
				value = customData
			}
			st.Compute.LaunchSpecification.SetCustomData(value)
			return nil
		},
		nil,
	)

	fieldsMap[ShutdownScript] = commons.NewGenericField(
		commons.StatefulNodeAzureExtensions,
		ShutdownScript,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil
			if st != nil && st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.ShutdownScript != nil {
				value = st.Compute.LaunchSpecification.ShutdownScript
			}
			if err := resourceData.Set(string(ShutdownScript), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShutdownScript), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(ShutdownScript)).(string); ok && v != "" {
				resourceGroupName := spotinst.String(v)
				st.Compute.LaunchSpecification.SetShutdownScript(resourceGroupName)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil
			if v, ok := resourceData.Get(string(ShutdownScript)).(string); ok && v != "" {
				shutDown := spotinst.String(v)
				value = shutDown
			}
			st.Compute.LaunchSpecification.SetShutdownScript(value)
			return nil
		},
		nil,
	)

	fieldsMap[UserData] = commons.NewGenericField(
		commons.StatefulNodeAzureExtensions,
		UserData,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil
			if st != nil && st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.UserData != nil {
				value = st.Compute.LaunchSpecification.UserData
			}
			if err := resourceData.Set(string(UserData), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UserData), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData := spotinst.String(v)
				st.Compute.LaunchSpecification.SetUserData(userData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData := spotinst.String(v)
				value = userData
			}
			st.Compute.LaunchSpecification.SetUserData(value)
			return nil
		},
		nil,
	)

	fieldsMap[Tag] = commons.NewGenericField(
		commons.StatefulNodeAzureLaunchSpecification,
		Tag,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(TagKey): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(TagValue): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []interface{} = nil

			if st.Compute != nil && st.Compute.LaunchSpecification != nil &&
				st.Compute.LaunchSpecification.Tags != nil {
				value = flattenTags(st.Compute.LaunchSpecification.Tags)
			}
			if value != nil {
				if err := resourceData.Set(string(Tag), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tag), err)
				}
			} else {
				if err := resourceData.Set(string(Tag), []*azure.Tag{}); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tag), err)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(Tag)); ok {
				if tags, err := expandTags(v); err != nil {
					return err
				} else {
					st.Compute.LaunchSpecification.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []*azure.Tag = nil
			if v, ok := resourceData.GetOk(string(Tag)); ok {
				if tasks, err := expandTags(v); err != nil {
					return err
				} else {
					value = tasks
				}
			}
			st.Compute.LaunchSpecification.SetTags(value)
			return nil
		},
		nil,
	)

	fieldsMap[ManagedServiceIdentities] = commons.NewGenericField(
		commons.StatefulNodeAzureLaunchSpecification,
		ManagedServiceIdentities,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Name): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(ResourceGroupName): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []interface{}
			if st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.ManagedServiceIdentities != nil {
				value = flattenManagedServiceIdentities(st.Compute.LaunchSpecification.ManagedServiceIdentities)
			}

			if err := resourceData.Set(string(ManagedServiceIdentities), value); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(ManagedServiceIdentities), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(ManagedServiceIdentities)); ok {
				if msis, err := expandManagedServiceIdentities(v); err != nil {
					return err
				} else {
					st.Compute.LaunchSpecification.SetManagedServiceIdentities(msis)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []*azure.ManagedServiceIdentity
			if v, ok := resourceData.GetOk(string(ManagedServiceIdentities)); ok {
				if msis, err := expandManagedServiceIdentities(v); err != nil {
					return err
				} else {
					value = msis
				}

			}
			st.Compute.LaunchSpecification.SetManagedServiceIdentities(value)
			return nil
		},
		nil,
	)

	fieldsMap[ProximityPlacementGroups] = commons.NewGenericField(
		commons.StatefulNodeAzureLaunchSpecification,
		ProximityPlacementGroups,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
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
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []interface{}
			if st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.ProximityPlacementGroups != nil {
				value = flattenProximityPlacementGroups(st.Compute.LaunchSpecification.ProximityPlacementGroups)
			}
			if err := resourceData.Set(string(ProximityPlacementGroups), value); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(ProximityPlacementGroups), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(ProximityPlacementGroups)); ok {
				if ppgs, err := expandProximityPlacementGroups(v); err != nil {
					return err
				} else {
					st.Compute.LaunchSpecification.SetProximityPlacementGroups(ppgs)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []*azure.ProximityPlacementGroups = nil

			if v, ok := resourceData.GetOk(string(ProximityPlacementGroups)); ok {
				if ppgs, err := expandProximityPlacementGroups(v); err != nil {
					return err
				} else {
					value = ppgs
				}
				st.Compute.LaunchSpecification.SetProximityPlacementGroups(value)
			} else {
				st.Compute.LaunchSpecification.SetProximityPlacementGroups(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[OSDisk] = commons.NewGenericField(
		commons.StatefulNodeAzureLaunchSpecification,
		OSDisk,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(OSDiskSizeGB): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
					string(OSDiskType): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(OSCaching): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []interface{} = nil

			if st.Compute != nil && st.Compute.LaunchSpecification != nil &&
				st.Compute.LaunchSpecification.OSDisk != nil {
				value = flattenOSDisk(st.Compute.LaunchSpecification.OSDisk)
			}
			if value != nil {
				if err := resourceData.Set(string(OSDisk), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OSDisk), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(OSDisk)); ok {
				if osDisk, err := expandOSDisk(v); err != nil {
					return err
				} else {
					st.Compute.LaunchSpecification.SetOSDisk(osDisk)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *azure.OSDisk = nil
			if v, ok := resourceData.GetOk(string(OSDisk)); ok {
				if osDisk, err := expandOSDisk(v); err != nil {
					return err
				} else {
					value = osDisk
				}
				st.Compute.LaunchSpecification.SetOSDisk(value)
			} else {
				st.Compute.LaunchSpecification.SetOSDisk(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[DataDisk] = commons.NewGenericField(
		commons.StatefulNodeAzureLaunchSpecification,
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
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []interface{} = nil

			if st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.DataDisks != nil {
				dataDisks := st.Compute.LaunchSpecification.DataDisks
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
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := stWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(DataDisk)); ok {
				if value, err := expandDataDisks(v); err != nil {
					return err
				} else {
					statefulNode.Compute.LaunchSpecification.SetDataDisks(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []*azure.DataDisk = nil

			if v, ok := resourceData.GetOk(string(DataDisk)); ok {
				if dataDisks, err := expandDataDisks(v); err != nil {
					return err
				} else {
					value = dataDisks
				}
				st.Compute.LaunchSpecification.SetDataDisks(value)
			} else {
				st.Compute.LaunchSpecification.SetDataDisks(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[BootDiagnostics] = commons.NewGenericField(
		commons.StatefulNodeAzureLaunchSpecification,
		BootDiagnostics,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(BootDiagnosticsIsEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
						Computed: true,
					},
					string(BootDiagnosticsType): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
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
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value interface{} = nil

			if st.Compute != nil && st.Compute.LaunchSpecification != nil &&
				st.Compute.LaunchSpecification.BootDiagnostics != nil {
				value = flattenBootDiagnostics(st.Compute.LaunchSpecification.BootDiagnostics)
			}
			if err := resourceData.Set(string(BootDiagnostics), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BootDiagnostics), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *azure.BootDiagnostics = nil

			if v, ok := resourceData.GetOk(string(BootDiagnostics)); ok {
				if bd, err := expandBootDiagnostics(v); err != nil {
					return err
				} else {
					value = bd
				}
			}
			st.Compute.LaunchSpecification.SetBootDiagnostics(value)

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *azure.BootDiagnostics = nil

			if v, ok := resourceData.GetOk(string(BootDiagnostics)); ok {
				if bd, err := expandBootDiagnostics(v); err != nil {
					return err
				} else {
					value = bd
				}
			}
			st.Compute.LaunchSpecification.SetBootDiagnostics(value)
			return nil
		},
		nil,
	)

	fieldsMap[VMName] = commons.NewGenericField(
		commons.StatefulNodeAzureLaunchSpecification,
		VMName,
		&schema.Schema{
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{string(VMNamePrefix)},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil
			if st != nil && st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.VMName != nil {
				value = st.Compute.LaunchSpecification.VMName
			}
			if err := resourceData.Set(string(VMName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(VMName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(VMName)).(string); ok && v != "" {
				vmName := spotinst.String(v)
				st.Compute.LaunchSpecification.SetVMName(vmName)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil
			if v, ok := resourceData.Get(string(VMName)).(string); ok && v != "" {
				vmName := spotinst.String(v)
				value = vmName
			}
			st.Compute.LaunchSpecification.SetVMName(value)
			return nil
		},
		nil,
	)

	fieldsMap[VMNamePrefix] = commons.NewGenericField(
		commons.StatefulNodeAzureLaunchSpecification,
		VMNamePrefix,
		&schema.Schema{
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{string(VMName)},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil
			if st != nil && st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.VMNamePrefix != nil {
				value = st.Compute.LaunchSpecification.VMNamePrefix
			}
			if err := resourceData.Set(string(VMNamePrefix), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(VMNamePrefix), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(VMNamePrefix)).(string); ok && v != "" {
				vmNamePrefix := spotinst.String(v)
				st.Compute.LaunchSpecification.SetVMNamePrefix(vmNamePrefix)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil
			if v, ok := resourceData.Get(string(VMNamePrefix)).(string); ok && v != "" {
				vmNamePrefix := spotinst.String(v)
				value = vmNamePrefix
			}
			st.Compute.LaunchSpecification.SetVMNamePrefix(value)
			return nil
		},
		nil,
	)

	fieldsMap[LicenseType] = commons.NewGenericField(
		commons.StatefulNodeAzureLaunchSpecification,
		LicenseType,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil
			if st != nil && st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.LicenseType != nil {
				value = st.Compute.LaunchSpecification.LicenseType
			}
			if err := resourceData.Set(string(LicenseType), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(LicenseType), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(LicenseType)).(string); ok && v != "" {
				licenseType := spotinst.String(v)
				st.Compute.LaunchSpecification.SetLicenseType(licenseType)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil
			if v, ok := resourceData.Get(string(LicenseType)).(string); ok && v != "" {
				licenseType := spotinst.String(v)
				value = licenseType
			}
			st.Compute.LaunchSpecification.SetLicenseType(value)
			return nil
		},
		nil,
	)

	fieldsMap[Security] = commons.NewGenericField(
		commons.StatefulNodeAzureLaunchSpecification,
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
						Default:  nil,
					},
					string(SecurityType): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(VTpmEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  nil,
					},
					string(EncryptionAtHost): {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  nil,
					},
					string(ConfidentialOsDiskEncryption): {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  nil,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value interface{} = nil

			if st.Compute != nil && st.Compute.LaunchSpecification != nil &&
				st.Compute.LaunchSpecification.Security != nil {
				value = flattenSecurity(st.Compute.LaunchSpecification.Security)
			}
			if err := resourceData.Set(string(Security), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Security), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(Security)); ok {
				if security, err := expandSecurity(v); err != nil {
					return err
				} else if security != nil {
					st.Compute.LaunchSpecification.SetSecurity(security)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *azure.Security = nil
			if v, ok := resourceData.GetOk(string(Security)); ok {
				if security, err := expandSecurity(v); err != nil {
					return err
				} else {
					value = security
				}
			}
			st.Compute.LaunchSpecification.SetSecurity(value)
			return nil
		},
		nil,
	)
}

func expandTags(data interface{}) ([]*azure.Tag, error) {
	list := data.(*schema.Set).List()
	tags := make([]*azure.Tag, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		tag := &azure.Tag{}

		if v, ok := m[string(TagKey)].(string); ok && v != "" {
			tag.SetTagKey(spotinst.String(v))
		}

		if v, ok := m[string(TagValue)].(string); ok && v != "" {
			tag.SetTagValue(spotinst.String(v))
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func flattenTags(tags []*azure.Tag) []interface{} {
	result := make([]interface{}, 0, len(tags))
	for _, tag := range tags {
		m := make(map[string]interface{})
		m[string(TagKey)] = spotinst.StringValue(tag.TagKey)
		m[string(TagValue)] = spotinst.StringValue(tag.TagValue)

		result = append(result, m)
	}
	return result
}

func expandManagedServiceIdentities(data interface{}) ([]*azure.ManagedServiceIdentity, error) {
	list := data.(*schema.Set).List()
	msis := make([]*azure.ManagedServiceIdentity, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		serviceId := &azure.ManagedServiceIdentity{}
		if v, ok := attr[string(ResourceGroupName)].(string); ok && v != "" {
			serviceId.SetResourceGroupName(spotinst.String(v))
		}
		if v, ok := attr[string(Name)].(string); ok && v != "" {
			serviceId.SetName(spotinst.String(v))
		}
		msis = append(msis, serviceId)
	}
	return msis, nil
}

func flattenManagedServiceIdentities(msis []*azure.ManagedServiceIdentity) []interface{} {
	result := make([]interface{}, 0, len(msis))
	for _, msi := range msis {
		m := make(map[string]interface{})
		m[string(ResourceGroupName)] = spotinst.StringValue(msi.ResourceGroupName)
		m[string(Name)] = spotinst.StringValue(msi.Name)
		result = append(result, m)
	}
	return result
}

func expandProximityPlacementGroups(data interface{}) ([]*azure.ProximityPlacementGroups, error) {
	list := data.(*schema.Set).List()
	ppgs := make([]*azure.ProximityPlacementGroups, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		ProximityPlacementGroup := &azure.ProximityPlacementGroups{}
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

func flattenProximityPlacementGroups(ppgs []*azure.ProximityPlacementGroups) []interface{} {
	result := make([]interface{}, 0, len(ppgs))
	for _, ppg := range ppgs {
		m := make(map[string]interface{})
		m[string(PPGResourceGroupName)] = spotinst.StringValue(ppg.ResourceGroupName)
		m[string(PPGName)] = spotinst.StringValue(ppg.Name)
		result = append(result, m)
	}
	return result
}

func flattenOSDisk(osd *azure.OSDisk) []interface{} {
	osDisk := make(map[string]interface{})
	osDisk[string(OSDiskSizeGB)] = spotinst.IntValue(osd.SizeGB)
	osDisk[string(OSDiskType)] = spotinst.StringValue(osd.Type)
	osDisk[string(OSCaching)] = spotinst.StringValue(osd.Caching)
	return []interface{}{osDisk}
}

func expandOSDisk(data interface{}) (*azure.OSDisk, error) {
	osDisk := &azure.OSDisk{}
	list := data.(*schema.Set).List()
	if len(list) > 0 {
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(OSDiskSizeGB)].(int); ok {
				if v == -1 {
					osDisk.SetSizeGB(nil)
				} else {
					osDisk.SetSizeGB(spotinst.Int(v))
				}
			}

			if v, ok := m[string(OSDiskType)].(string); ok && v != "" {
				osDisk.SetType(spotinst.String(v))
			} else {
				osDisk.SetType(nil)
			}

			if v, ok := m[string(OSCaching)].(string); ok && v != "" {
				osDisk.SetCaching(spotinst.String(v))
			} else {
				osDisk.SetCaching(nil)
			}
		}
		return osDisk, nil
	}
	return nil, nil
}

func flattenDataDisks(dataDisks []*azure.DataDisk) []interface{} {
	var result []interface{}

	for _, disk := range dataDisks {
		m := make(map[string]interface{})
		m[string(DataDiskSizeGB)] = spotinst.IntValue(disk.SizeGB)
		m[string(DataDiskLUN)] = spotinst.IntValue(disk.LUN)
		m[string(DataDiskType)] = spotinst.StringValue(disk.Type)
		result = append(result, m)
	}
	return result
}

func expandDataDisks(data interface{}) ([]*azure.DataDisk, error) {
	list := data.([]interface{})
	dd := make([]*azure.DataDisk, 0, len(list))
	for _, m := range list {
		attr, ok := m.(map[string]interface{})
		if !ok {
			continue
		}

		dataDisk := &azure.DataDisk{}
		if v, ok := attr[string(DataDiskSizeGB)].(int); ok && v > 0 {
			dataDisk.SetSizeGB(spotinst.Int(v))
		}

		if v, ok := attr[string(DataDiskLUN)].(int); ok && v >= 0 {
			dataDisk.SetLUN(spotinst.Int(v))
		}

		if v, ok := attr[string(DataDiskType)].(string); ok && v != "" {
			dataDisk.SetType(spotinst.String(v))
		}

		dd = append(dd, dataDisk)
	}

	return dd, nil
}

func flattenBootDiagnostics(bd *azure.BootDiagnostics) interface{} {
	bootDiagnostic := make(map[string]interface{})
	bootDiagnostic[string(BootDiagnosticsIsEnabled)] = spotinst.BoolValue(bd.IsEnabled)
	bootDiagnostic[string(BootDiagnosticsType)] = spotinst.StringValue(bd.Type)
	bootDiagnostic[string(BootDiagnosticsStorageURL)] = spotinst.StringValue(bd.StorageURL)
	return []interface{}{bootDiagnostic}
}

func expandBootDiagnostics(data interface{}) (*azure.BootDiagnostics, error) {
	if list := data.([]interface{}); len(list) > 0 {
		bootDiagnostic := &azure.BootDiagnostics{}

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
			bootDiagnostic.SetStorageURL(storageURL)

		}

		return bootDiagnostic, nil
	}

	return nil, nil
}

func flattenSecurity(secure *azure.Security) interface{} {
	security := make(map[string]interface{})

	security[string(SecureBootEnabled)] = spotinst.BoolValue(secure.SecureBootEnabled)
	security[string(SecurityType)] = spotinst.StringValue(secure.SecurityType)
	security[string(VTpmEnabled)] = spotinst.BoolValue(secure.VTpmEnabled)
	security[string(EncryptionAtHost)] = spotinst.BoolValue(secure.EncryptionAtHost)
	security[string(ConfidentialOsDiskEncryption)] = spotinst.BoolValue(secure.ConfidentialOsDiskEncryption)

	return []interface{}{security}
}

func expandSecurity(data interface{}) (*azure.Security, error) {
	if list := data.([]interface{}); len(list) > 0 {
		security := &azure.Security{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, exists := m[string(SecureBootEnabled)]; exists && v != nil {
				if b, ok := v.(bool); ok {
					security.SetSecureBootEnabled(spotinst.Bool(b))
				}
			}

			if v, exists := m[string(EncryptionAtHost)]; exists && v != nil {
				if b, ok := v.(bool); ok {
					security.SetEncryptionAtHost(spotinst.Bool(b))
				}
			}

			if v, exists := m[string(ConfidentialOsDiskEncryption)]; exists && v != nil {
				if b, ok := v.(bool); ok {
					security.SetConfidentialOsDiskEncryption(spotinst.Bool(b))
				}
			}

			if v, ok := m[string(SecurityType)].(string); ok && v != "" {
				security.SetSecurityType(spotinst.String(v))
			}

			if v, exists := m[string(VTpmEnabled)]; exists && v != nil {
				if b, ok := v.(bool); ok {
					security.SetVTpmEnabled(spotinst.Bool(b))
				}
			}

			if v, exists := m[string(EncryptionAtHost)]; exists && v != nil {
				if b, ok := v.(bool); ok {
					security.SetEncryptionAtHost(spotinst.Bool(b))
				}
			}

			if v, exists := m[string(ConfidentialOsDiskEncryption)]; exists && v != nil {
				if b, ok := v.(bool); ok {
					security.SetConfidentialOsDiskEncryption(spotinst.Bool(b))
				}
			}
		}

		return security, nil
	}

	return nil, nil
}

func base64Encode(data string) string {
	// Check whether the data is already Base64 encoded; don't double-encode
	if isBase64Encoded(data) {
		return data
	}
	// data has not been encoded -> encode and return
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func isBase64Encoded(data string) bool {
	_, err := base64.StdEncoding.DecodeString(data)
	return err == nil
}

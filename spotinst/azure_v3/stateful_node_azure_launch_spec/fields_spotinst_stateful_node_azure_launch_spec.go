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
						//Computed: true,
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
			Computed: true,
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
			if st != nil && st.Compute != nil &&
				st.Compute.LaunchSpecification != nil &&
				st.Compute.LaunchSpecification.ManagedServiceIdentities != nil {
				if v, ok := resourceData.GetOk(string(ManagedServiceIdentities)); ok {
					if msis, err := expandManagedServiceIdentities(v); err != nil {
						return err
					} else {
						value = msis
					}
				}
				st.Compute.LaunchSpecification.SetManagedServiceIdentities(value)
			}
			return nil
		},
		nil,
	)

	fieldsMap[OSDisk] = commons.NewGenericField(
		commons.StatefulNodeAzureLaunchSpecification,
		OSDisk,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(OSDiskSizeGB): {
						Type:     schema.TypeInt,
						Optional: true,
						Computed: true,
					},
					string(OSDiskType): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value interface{} = nil

			if st.Compute != nil && st.Compute.LaunchSpecification != nil &&
				st.Compute.LaunchSpecification.OSDisk != nil {
				value = flattenOSDisk(st.Compute.LaunchSpecification.OSDisk)
			}
			if err := resourceData.Set(string(OSDisk), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OSDisk), err)
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
			}
			st.Compute.LaunchSpecification.SetOSDisk(value)

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
			}
			st.Compute.LaunchSpecification.SetOSDisk(value)
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
			Computed: true,
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
			}
			st.Compute.LaunchSpecification.SetDataDisks(value)
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
		msis = append(msis, &azure.ManagedServiceIdentity{
			ResourceGroupName: spotinst.String(attr[string(ResourceGroupName)].(string)),
			Name:              spotinst.String(attr[string(Name)].(string)),
		})
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

func flattenOSDisk(osd *azure.OSDisk) interface{} {
	osDisk := make(map[string]interface{})
	osDisk[string(OSDiskSizeGB)] = spotinst.IntValue(osd.SizeGB)
	osDisk[string(OSDiskType)] = spotinst.StringValue(osd.Type)
	return []interface{}{osDisk}
}

func expandOSDisk(data interface{}) (*azure.OSDisk, error) {
	if list := data.([]interface{}); len(list) > 0 {
		osDisk := &azure.OSDisk{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})
			var sizeGB *int = nil
			var osType *string = nil

			if v, ok := m[string(OSDiskSizeGB)].(int); ok && v > 0 {
				sizeGB = spotinst.Int(v)
			}
			osDisk.SetSizeGB(sizeGB)

			if v, ok := m[string(OSDiskType)].(string); ok && v != "" {
				osType = spotinst.String(v)
			}
			osDisk.SetType(osType)

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
		} else {
			dataDisk.SetSizeGB(nil)
		}

		if v, ok := attr[string(DataDiskLUN)].(int); ok && v >= 0 {
			dataDisk.SetLUN(spotinst.Int(v))
		} else {
			dataDisk.SetLUN(nil)
		}

		if v, ok := attr[string(DataDiskType)].(string); ok && v != "" {
			dataDisk.SetType(spotinst.String(v))
		} else {
			dataDisk.SetType(nil)
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

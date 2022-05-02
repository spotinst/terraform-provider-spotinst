package stateful_node_azure_launch_spec

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
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
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(TagKey): {
						Type:     schema.TypeString,
						Optional: true,
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
			if err := resourceData.Set(string(Tag), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tag), err)
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
			var value []*azurev3.Tag = nil
			if st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.Tags != nil {
				if v, ok := resourceData.GetOk(string(Tag)); ok {
					if tags, err := expandTags(v); err != nil {
						return err
					} else {
						value = tags
					}
				}
				st.Compute.LaunchSpecification.SetTags(value)
			}
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
			var value []*azurev3.ManagedServiceIdentity
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
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(SizeGB): {
						Type:     schema.TypeInt,
						Optional: true,
					},
					string(Type): {
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
			var value *azurev3.OSDisk = nil

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
			var value *azurev3.OSDisk = nil

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
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(SizeGB): {
						Type:     schema.TypeInt,
						Required: true,
					},
					string(LUN): {
						Type:     schema.TypeInt,
						Required: true,
					},
					string(Type): {
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
			var value []*azurev3.DataDisk = nil

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
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(IsEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(Type): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(StorageURL): {
						Type:     schema.TypeString,
						Optional: true,
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
			var value *azurev3.BootDiagnostics = nil

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
			var value *azurev3.BootDiagnostics = nil

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

func expandTags(data interface{}) ([]*azurev3.Tag, error) {
	list := data.(*schema.Set).List()
	tags := make([]*azurev3.Tag, 0, len(list))

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

		tag := &azurev3.Tag{
			TagKey:   spotinst.String(attr[string(TagKey)].(string)),
			TagValue: spotinst.String(attr[string(TagValue)].(string)),
		}

		tags = append(tags, tag)
	}
	return tags, nil
}

func flattenTags(tags []*azurev3.Tag) []interface{} {
	result := make([]interface{}, 0, len(tags))
	for _, tag := range tags {
		m := make(map[string]interface{})
		m[string(TagKey)] = spotinst.StringValue(tag.TagKey)
		m[string(TagValue)] = spotinst.StringValue(tag.TagValue)

		result = append(result, m)
	}
	return result
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
			ResourceGroupName: spotinst.String(attr[string(ResourceGroupName)].(string)),
			Name:              spotinst.String(attr[string(Name)].(string)),
		})
	}
	return msis, nil
}

func flattenManagedServiceIdentities(msis []*azurev3.ManagedServiceIdentity) []interface{} {
	result := make([]interface{}, 0, len(msis))
	for _, msi := range msis {
		m := make(map[string]interface{})
		m[string(ResourceGroupName)] = spotinst.StringValue(msi.ResourceGroupName)
		m[string(Name)] = spotinst.StringValue(msi.Name)
		result = append(result, m)
	}
	return result
}

func flattenOSDisk(osd *azurev3.OSDisk) interface{} {
	osDisk := make(map[string]interface{})
	osDisk[string(SizeGB)] = spotinst.IntValue(osd.SizeGB)
	osDisk[string(Type)] = spotinst.StringValue(osd.Type)
	return []interface{}{osDisk}
}

func expandOSDisk(data interface{}) (*azurev3.OSDisk, error) {
	if list := data.([]interface{}); len(list) > 0 {
		osDisk := &azurev3.OSDisk{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})
			var sizeGB *int = nil
			var osType *string = nil

			if v, ok := m[string(SizeGB)].(int); ok && v > 0 {
				sizeGB = spotinst.Int(v)
			}
			osDisk.SetSizeGB(sizeGB)

			if v, ok := m[string(Type)].(string); ok && v != "" {
				osType = spotinst.String(v)
			}
			osDisk.SetType(osType)

		}

		return osDisk, nil
	}

	return nil, nil
}

func flattenDataDisks(dataDisks []*azurev3.DataDisk) []interface{} {
	var result []interface{}

	for _, disk := range dataDisks {
		m := make(map[string]interface{})
		m[string(SizeGB)] = spotinst.IntValue(disk.SizeGB)
		m[string(LUN)] = spotinst.IntValue(disk.LUN)
		m[string(Type)] = spotinst.StringValue(disk.Type)
		result = append(result, m)
	}
	return result
}

func expandDataDisks(data interface{}) ([]*azurev3.DataDisk, error) {
	list := data.([]interface{})
	dd := make([]*azurev3.DataDisk, 0, len(list))
	for _, m := range list {
		attr, ok := m.(map[string]interface{})
		if !ok {
			continue
		}

		dataDisk := &azurev3.DataDisk{}
		if v, ok := attr[string(SizeGB)].(int); ok && v > 0 {
			dataDisk.SetSizeGB(spotinst.Int(v))
		} else {
			dataDisk.SetSizeGB(nil)
		}

		if v, ok := attr[string(LUN)].(int); ok && v >= 0 {
			dataDisk.SetLUN(spotinst.Int(v))
		} else {
			dataDisk.SetLUN(nil)
		}

		if v, ok := attr[string(Type)].(string); ok && v != "" {
			dataDisk.SetType(spotinst.String(v))
		} else {
			dataDisk.SetType(nil)
		}

		dd = append(dd, dataDisk)
	}

	return dd, nil
}

func flattenBootDiagnostics(bd *azurev3.BootDiagnostics) interface{} {
	bootDiagnostic := make(map[string]interface{})
	bootDiagnostic[string(IsEnabled)] = spotinst.BoolValue(bd.IsEnabled)
	bootDiagnostic[string(Type)] = spotinst.StringValue(bd.Type)
	bootDiagnostic[string(StorageURL)] = spotinst.StringValue(bd.StorageURL)
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

			if v, ok := m[string(IsEnabled)].(bool); ok {
				enabled = spotinst.Bool(v)
			}
			bootDiagnostic.SetIsEnabled(enabled)

			if v, ok := m[string(Type)].(string); ok && v != "" {
				bsType = spotinst.String(v)
			}
			bootDiagnostic.SetType(bsType)

			if v, ok := m[string(StorageURL)].(string); ok && v != "" {
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
	// data has not been encoded encode and return
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func isBase64Encoded(data string) bool {
	_, err := base64.StdEncoding.DecodeString(data)
	return err == nil
}

package elastigroup_azure_extension

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Extensions] = commons.NewGenericField(
		commons.ElastirgoupAzureExtensions,
		Extensions,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Publisher): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(APIVersion): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(MinorVersionAutoUpgrade): {
						Type:     schema.TypeBool,
						Required: true,
					},

					string(EnableAutomaticUpgrade): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(Name): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(Type): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(ProtectedSettings): {
						Type:     schema.TypeMap,
						Optional: true,
						Computed: true,
					},

					string(PublicSettings): {
						Type:     schema.TypeMap,
						Optional: true,
						Computed: true,
					},

					string(ProtectedSettingsFromKeyVault): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(SecretUrl): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(SourceVault): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			eg := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if eg != nil && eg.Compute != nil && eg.Compute.LaunchSpecification != nil && eg.Compute.LaunchSpecification.Extensions != nil {
				extensions := eg.Compute.LaunchSpecification.Extensions
				result = flattenExtensions(extensions)
			}
			if err := resourceData.Set(string(Extensions), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Extensions), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			eg := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Extensions)); ok {
				if ext, err := expandExtensions(v); err != nil {
					return err
				} else {
					eg.Compute.LaunchSpecification.SetExtensions(ext)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			eg := egWrapper.GetElastigroup()
			var value []*azurev3.Extensions = nil
			if v, ok := resourceData.GetOk(string(Extensions)); ok && v != nil {
				if ext, err := expandExtensions(v); err != nil {
					return err
				} else {
					value = ext
				}
			}
			eg.Compute.LaunchSpecification.SetExtensions(value)
			return nil
		},
		nil,
	)
}

func flattenExtensions(extensions []*azurev3.Extensions) []interface{} {
	result := make([]interface{}, 0, len(extensions))

	for _, extension := range extensions {
		m := make(map[string]interface{})
		m[string(APIVersion)] = spotinst.StringValue(extension.APIVersion)
		m[string(Name)] = spotinst.StringValue(extension.Name)
		m[string(Publisher)] = spotinst.StringValue(extension.Publisher)
		m[string(Type)] = spotinst.StringValue(extension.Type)
		m[string(MinorVersionAutoUpgrade)] = spotinst.BoolValue(extension.MinorVersionAutoUpgrade)
		m[string(EnableAutomaticUpgrade)] = spotinst.BoolValue(extension.EnableAutomaticUpgrade)
		m[string(ProtectedSettings)] = extension.ProtectedSettings
		m[string(PublicSettings)] = extension.PublicSettings
		if extension.ProtectedSettingsFromKeyVault != nil {
			m[string(ProtectedSettingsFromKeyVault)] = flattenProtectedSettingsFromKeyVault(extension.ProtectedSettingsFromKeyVault)
		}

		result = append(result, m)
	}
	return result
}

func flattenProtectedSettingsFromKeyVault(protectedSettings *azurev3.ProtectedSettingsFromKeyVault) []interface{} {
	result := make(map[string]interface{})

	result[string(SecretUrl)] = spotinst.StringValue(protectedSettings.SecretUrl)
	result[string(SourceVault)] = spotinst.StringValue(protectedSettings.SourceVault)

	return []interface{}{result}
}

func expandExtensions(data interface{}) ([]*azurev3.Extensions, error) {
	if list := data.([]interface{}); len(list) > 0 {
		extensions := make([]*azurev3.Extensions, 0, len(list))
		for _, v := range list {
			ext, ok := v.(map[string]interface{})
			if !ok {
				continue
			}
			extension := &azurev3.Extensions{}

			if v, ok := ext[string(APIVersion)].(string); ok && v != "" {
				extension.SetAPIVersion(spotinst.String(v))
			}
			if v, ok := ext[string(Name)].(string); ok && v != "" {
				extension.SetName(spotinst.String(v))
			}
			if v, ok := ext[string(Publisher)].(string); ok && v != "" {
				extension.SetPublisher(spotinst.String(v))
			}
			if v, ok := ext[string(Type)].(string); ok && v != "" {
				extension.SetType(spotinst.String(v))
			}
			if v, ok := ext[string(MinorVersionAutoUpgrade)].(bool); ok {
				extension.SetMinorVersionAutoUpgrade(spotinst.Bool(v))
			}
			if v, ok := ext[string(EnableAutomaticUpgrade)].(bool); ok {
				extension.SetEnableAutomaticUpgrade(spotinst.Bool(v))
			}
			if v, ok := ext[string(ProtectedSettings)].(map[string]interface{}); ok {
				extension.SetProtectedSettings(v)
			}
			if v, ok := ext[string(PublicSettings)].(map[string]interface{}); ok {
				extension.SetPublicSettings(v)
			}
			if v, ok := ext[string(ProtectedSettingsFromKeyVault)]; ok {
				settings, err := expandProtectedSettingsFromKeyVault(v)
				if err != nil {
					return nil, err
				}
				if settings != nil {
					extension.SetProtectedSettingsFromKeyVault(settings)
				} else {
					extension.SetProtectedSettingsFromKeyVault(nil)
				}
			}

			extensions = append(extensions, extension)
		}
		return extensions, nil
	}
	return nil, nil
}

func expandProtectedSettingsFromKeyVault(data interface{}) (*azurev3.ProtectedSettingsFromKeyVault, error) {
	settings := &azurev3.ProtectedSettingsFromKeyVault{}
	list := data.([]interface{})

	if list == nil || len(list) == 0 {
		return nil, nil
	}
	m := list[0].(map[string]interface{})
	if v, ok := m[string(SecretUrl)].(string); ok && v != "" {
		settings.SetSecretUrl(spotinst.String(v))
	}
	if v, ok := m[string(SourceVault)].(string); ok && v != "" {
		settings.SetSourceVault(spotinst.String(v))
	}
	return settings, nil
}

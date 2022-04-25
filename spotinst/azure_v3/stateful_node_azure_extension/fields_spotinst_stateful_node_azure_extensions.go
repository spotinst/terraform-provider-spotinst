package stateful_node_azure_extension

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"strings"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Extension] = commons.NewGenericField(
		commons.StatefulNodeAzureExtensions,
		Extension,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if strings.ToLower(old) == strings.ToLower(new) {
					return true
				}
				return false
			},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Publisher): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},

					string(APIVersion): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},

					string(MinorVersionAutoUpgrade): {
						Type:     schema.TypeBool,
						Optional: true,
						Computed: true,
					},

					string(Name): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},

					string(Type): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
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
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var result []interface{} = nil
			if st != nil && st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.Extensions != nil {
				extensions := st.Compute.LaunchSpecification.Extensions
				result = flattenExtensions(extensions)
			}

			if result != nil {
				if err := resourceData.Set(string(Extension), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Extension), err)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []*.Extension = nil

			if v, ok := resourceData.GetOk(string(Extension)); ok {
				var extensions []*azure.Extension

				if st != nil && st.Compute != nil && st.Compute.LaunchSpecification != nil {
					if st.Compute.LaunchSpecification.Extensions != nil {
						extensions = st.Compute.LaunchSpecification.Extensions
					}

					if ext, err := expandExtensions(v, extensions); err != nil {
						return err
					} else {
						value = ext
					}

					st.Compute.LaunchSpecification.SetExtensions(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.AKSstWrapper)
			st := stWrapper.Getst()
			var value []*azure.Extension = nil

			if v, ok := resourceData.GetOk(string(Extension)); ok {
				//create new image object in case st did not get it from previous import step.
				var extensions []*azure.Extension

				if st != nil && st.VirtualNodeGroupTemplate != nil && st.VirtualNodeGroupTemplate.LaunchSpecification != nil {

					if st.VirtualNodeGroupTemplate.LaunchSpecification.Extensions != nil {
						extensions = st.VirtualNodeGroupTemplate.LaunchSpecification.Extensions
					}

					if extensions, err := expandExtensions(v, extensions); err != nil {
						return err
					} else {
						value = extensions
					}

					st.VirtualNodeGroupTemplate.LaunchSpecification.SetExtensions(value)
				}
			}
			return nil
		},
		nil,
	)
}

func flattenExtensions(extensions []*azure.Extension) []interface{} {
	result := make([]interface{}, 0, len(extensions))

	for _, extension := range extensions {
		m := make(map[string]interface{})
		m[string(APIVersion)] = spotinst.StringValue(extension.APIVersion)
		m[string(Name)] = spotinst.StringValue(extension.Name)
		m[string(Publisher)] = spotinst.StringValue(extension.Publisher)
		m[string(Type)] = spotinst.StringValue(extension.Type)
		m[string(MinorVersionAutoUpgrade)] = spotinst.BoolValue(extension.MinorVersionAutoUpgrade)

		result = append(result, m)
	}
	return result
}

func expandExtensions(data interface{}, extensions []*azure.Extension) ([]*azure.Extension, error) {
	list := data.(*schema.Set).List()

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		extension := &azure.Extension{}

		if v, ok := attr[string(APIVersion)].(string); ok && v != "" {
			extension.SetAPIVersion(spotinst.String(v))
		}
		if v, ok := attr[string(Name)].(string); ok && v != "" {
			extension.SetName(spotinst.String(v))
		}
		if v, ok := attr[string(Publisher)].(string); ok && v != "" {
			extension.SetPublisher(spotinst.String(v))
		}
		if v, ok := attr[string(Type)].(string); ok && v != "" {
			extension.SetType(spotinst.String(v))
		}
		if v, ok := attr[string(MinorVersionAutoUpgrade)].(bool); ok {
			extension.SetMinorVersionAutoUpgrade(spotinst.Bool(v))
		}

		extensions = append(extensions, extension)
	}

	return extensions, nil
}

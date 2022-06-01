package ocean_aks_extensions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Extension] = commons.NewGenericField(
		commons.OceanAKSExtensions,
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
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification.Extensions != nil {
				extensions := cluster.VirtualNodeGroupTemplate.LaunchSpecification.Extensions
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
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value []*azure.Extension = nil

			if v, ok := resourceData.GetOk(string(Extension)); ok {
				var extensions []*azure.Extension

				if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil {
					if cluster.VirtualNodeGroupTemplate.LaunchSpecification.Extensions != nil {
						extensions = cluster.VirtualNodeGroupTemplate.LaunchSpecification.Extensions
					}

					if ext, err := expandExtensions(v, extensions); err != nil {
						return err
					} else {
						value = ext
					}

					cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetExtensions(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value []*azure.Extension = nil

			if v, ok := resourceData.GetOk(string(Extension)); ok {
				//create new image object in case cluster did not get it from previous import step.
				var extensions []*azure.Extension

				if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil {

					if cluster.VirtualNodeGroupTemplate.LaunchSpecification.Extensions != nil {
						extensions = cluster.VirtualNodeGroupTemplate.LaunchSpecification.Extensions
					}

					if extensions, err := expandExtensions(v, extensions); err != nil {
						return err
					} else {
						value = extensions
					}

					cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetExtensions(value)
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

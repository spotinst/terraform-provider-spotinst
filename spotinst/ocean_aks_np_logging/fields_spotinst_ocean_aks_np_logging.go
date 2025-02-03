package ocean_aks_np_logging

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Logging] = commons.NewGenericField(
		commons.OceanAKSNPLogging,
		Logging,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Export): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(AzureBlob): {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Id): {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var result []interface{} = nil
			if cluster != nil && cluster.Logging != nil {
				result = flattenLogging(cluster.Logging)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Logging), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Logging), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			if v, ok := resourceData.GetOk(string(Logging)); ok {
				if logging, err := expandLogging(v); err != nil {
					return err
				} else {
					cluster.SetLogging(logging)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *azure_np.Logging = nil

			if v, ok := resourceData.GetOk(string(Logging)); ok {
				if logging, err := expandLogging(v); err != nil {
					return err
				} else {
					value = logging
				}
			}
			cluster.SetLogging(value)
			return nil
		},
		nil,
	)
}

func flattenLogging(logging *azure_np.Logging) []interface{} {
	var out []interface{}

	if logging != nil {
		result := make(map[string]interface{})

		if logging.Export != nil {
			result[string(Export)] = flattenExport(logging.Export)
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func flattenExport(export *azure_np.Export) []interface{} {
	var out []interface{}

	if export != nil {
		result := make(map[string]interface{})

		if export.AzureBlob != nil {
			result[string(AzureBlob)] = flattenAzureBlob(export.AzureBlob)
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func flattenAzureBlob(azureBlob *azure_np.AzureBlob) []interface{} {
	var out []interface{}

	if azureBlob != nil {
		result := make(map[string]interface{})

		if azureBlob.Id != nil {
			result[string(Id)] = azureBlob.Id
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func expandLogging(data interface{}) (*azure_np.Logging, error) {
	logging := &azure_np.Logging{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return logging, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Export)]; ok {
		export, err := expandExport(v)
		if err != nil {
			return nil, err
		}
		if export != nil {
			logging.SetExport(export)
		} else {
			logging.Export = nil
		}
	}

	return logging, nil
}

func expandExport(data interface{}) (*azure_np.Export, error) {
	export := &azure_np.Export{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return export, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(AzureBlob)]; ok {
		azureBlob, err := expandAzureBlob(v)
		if err != nil {
			return nil, err
		}
		if azureBlob != nil {
			export.SetAzureBlob(azureBlob)
		} else {
			export.AzureBlob = nil
		}
	}

	return export, nil
}

func expandAzureBlob(data interface{}) (*azure_np.AzureBlob, error) {
	azureBlob := &azure_np.AzureBlob{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return azureBlob, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Id)].(string); ok && v != "" {
		azureBlob.SetId(spotinst.String(v))
	}

	return azureBlob, nil
}

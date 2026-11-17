package ocean_gke_import_launch_specification

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[RootVolumeType] = commons.NewGenericField(
		commons.OceanGKEImportLaunchSpecification,
		RootVolumeType,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result *string = nil
			if cluster != nil && cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil && cluster.Compute.LaunchSpecification.RootVolumeType != nil {
				result = cluster.Compute.LaunchSpecification.RootVolumeType
			}
			if result != nil {
				if err := resourceData.Set(string(RootVolumeType), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RootVolumeType), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var rootVolumeType *string = nil

			if cluster != nil && cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil {
				// get rootVolumeType from previous import step.
				rootVolumeType = cluster.Compute.LaunchSpecification.RootVolumeType

				// get rootVolumeType from user configuration.
				if v, ok := resourceData.GetOk(string(RootVolumeType)); ok {
					rootVolumeType = spotinst.String(v.(string))

					if rootVolumeType != nil {
						cluster.Compute.LaunchSpecification.SetRootVolumeType(rootVolumeType)
					}
				}

			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var rootVolumeType *string = nil
			if v, ok := resourceData.GetOk(string(RootVolumeType)); ok && v != "" {
				rootVolumeType = spotinst.String(v.(string))
			}
			cluster.Compute.LaunchSpecification.SetRootVolumeType(rootVolumeType)
			return nil
		},
		nil,
	)

	fieldsMap[DynamicVolumeSize] = commons.NewGenericField(
		commons.OceanGKEImportLaunchSpecification,
		DynamicVolumeSize,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(BaseSize): {
						Type:     schema.TypeInt,
						Required: true,
					},
					string(Resource): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(SizePerResourceUnit): {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			cluster := resourceObject.(*commons.GKEImportClusterWrapper).GetCluster()
			var result []interface{}
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil {
				result = flattenDynamicVolumeSize(cluster.Compute.LaunchSpecification.DynamicVolumeSize)
			}
			if err := resourceData.Set(string(DynamicVolumeSize), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DynamicVolumeSize), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			cluster := resourceObject.(*commons.GKEImportClusterWrapper).GetCluster()

			if v, ok := resourceData.GetOk(string(DynamicVolumeSize)); ok {
				dynamicVolumeSize, err := expandDynamicVolumeSize(v)
				if err != nil {
					return err
				}
				if cluster.Compute == nil {
					cluster.SetCompute(&gcp.Compute{})
				}
				if cluster.Compute.LaunchSpecification == nil {
					cluster.Compute.SetLaunchSpecification(&gcp.LaunchSpecification{})
				}
				cluster.Compute.LaunchSpecification.SetDynamicVolumeSize(dynamicVolumeSize)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			ls := resourceObject.(*commons.GKEImportClusterWrapper).GetCluster().Compute.LaunchSpecification
			var dynamicVolumeSize *gcp.ClusterDynamicVolumeSize
			if v, ok := resourceData.GetOk(string(DynamicVolumeSize)); ok {
				var err error
				dynamicVolumeSize, err = expandDynamicVolumeSize(v)
				if err != nil {
					return err
				}
			}
			ls.SetDynamicVolumeSize(dynamicVolumeSize)
			return nil
		},
		nil,
	)

	fieldsMap[ShieldedInstanceConfig] = commons.NewGenericField(
		commons.OceanGKEImportLaunchSpecification,
		ShieldedInstanceConfig,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(EnableSecureBoot): {
						Type:     schema.TypeBool,
						Computed: true,
						Optional: true,
					},
					string(EnableIntegrityMonitoring): {
						Type:     schema.TypeBool,
						Computed: true,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster != nil && cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.ShieldedInstanceConfig != nil {
				shieldedInstanceConfig := cluster.Compute.LaunchSpecification.ShieldedInstanceConfig
				result = flattenShieldedInstanceConfig(shieldedInstanceConfig)
			}
			if result != nil {
				if err := resourceData.Set(string(ShieldedInstanceConfig), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShieldedInstanceConfig), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *gcp.LaunchSpecShieldedInstanceConfig = nil

			if v, ok := resourceData.GetOk(string(ShieldedInstanceConfig)); ok {

				if cluster != nil && cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
					cluster.Compute.LaunchSpecification.ShieldedInstanceConfig != nil {
					value = cluster.Compute.LaunchSpecification.ShieldedInstanceConfig
				}

				if shieldedInstanceConfig, err := expandShieldedInstanceConfig(v); err != nil {
					return err
				} else if shieldedInstanceConfig != nil {
					value = shieldedInstanceConfig
				}

				cluster.Compute.LaunchSpecification.SetShieldedInstanceConfig(value)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *gcp.LaunchSpecShieldedInstanceConfig = nil

			if v, ok := resourceData.GetOk(string(ShieldedInstanceConfig)); ok {
				if shieldedInstanceConfig, err := expandShieldedInstanceConfig(v); err != nil {
					return err
				} else {
					value = shieldedInstanceConfig
				}
			}
			cluster.Compute.LaunchSpecification.SetShieldedInstanceConfig(value)
			return nil
		},
		nil,
	)

	fieldsMap[UseAsTemplateOnly] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		UseAsTemplateOnly,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			var value *bool = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.UseAsTemplateOnly != nil {

				value = cluster.Compute.LaunchSpecification.UseAsTemplateOnly
			}

			if err := resourceData.Set(string(UseAsTemplateOnly), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UseAsTemplateOnly), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.Get(string(UseAsTemplateOnly)).(bool); ok {
				cluster.Compute.LaunchSpecification.SetUseAsTemplateOnly(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEImportClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.Get(string(UseAsTemplateOnly)).(bool); ok {
				cluster.Compute.LaunchSpecification.SetUseAsTemplateOnly(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)
}

func expandDynamicVolumeSize(data interface{}) (*gcp.ClusterDynamicVolumeSize, error) {
	var dynamicVolumeSize *gcp.ClusterDynamicVolumeSize
	list := data.([]interface{})

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		d := &gcp.ClusterDynamicVolumeSize{}

		if v, ok := attr[string(BaseSize)].(int); ok {
			d.SetBaseSize(spotinst.Int(v))
		}

		if v, ok := attr[string(Resource)].(string); ok && v != "" {
			d.SetResource(spotinst.String(v))
		}

		if v, ok := attr[string(SizePerResourceUnit)].(int); ok {
			d.SetSizePerResourceUnit(spotinst.Int(v))
		}

		dynamicVolumeSize = d
	}

	return dynamicVolumeSize, nil
}

func flattenDynamicVolumeSize(dynamicVolumeSize *gcp.ClusterDynamicVolumeSize) []interface{} {
	var out []interface{}

	if dynamicVolumeSize != nil {
		result := make(map[string]interface{})

		if dynamicVolumeSize.BaseSize != nil {
			result[string(BaseSize)] = spotinst.IntValue(dynamicVolumeSize.BaseSize)
		}

		if dynamicVolumeSize.Resource != nil {
			result[string(Resource)] = spotinst.StringValue(dynamicVolumeSize.Resource)
		}

		if dynamicVolumeSize.SizePerResourceUnit != nil {
			result[string(SizePerResourceUnit)] = spotinst.IntValue(dynamicVolumeSize.SizePerResourceUnit)
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

func flattenShieldedInstanceConfig(shieldedInstanceConfig *gcp.LaunchSpecShieldedInstanceConfig) []interface{} {
	var out []interface{}

	if shieldedInstanceConfig != nil {
		result := make(map[string]interface{})

		if shieldedInstanceConfig.EnableSecureBoot != nil {
			result[string(EnableSecureBoot)] = spotinst.BoolValue(shieldedInstanceConfig.EnableSecureBoot)
		}
		if shieldedInstanceConfig.EnableIntegrityMonitoring != nil {
			result[string(EnableIntegrityMonitoring)] = spotinst.BoolValue(shieldedInstanceConfig.EnableIntegrityMonitoring)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

func expandShieldedInstanceConfig(data interface{}) (*gcp.LaunchSpecShieldedInstanceConfig, error) {
	shieldedInstanceConfig := &gcp.LaunchSpecShieldedInstanceConfig{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(EnableSecureBoot)].(bool); ok {
			shieldedInstanceConfig.SetEnableSecureBoot(spotinst.Bool(v))
		}

		if v, ok := m[string(EnableIntegrityMonitoring)].(bool); ok {
			shieldedInstanceConfig.SetEnableIntegrityMonitoring(spotinst.Bool(v))
		}
	}

	return shieldedInstanceConfig, nil
}

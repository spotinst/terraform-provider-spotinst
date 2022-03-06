package ocean_gke_import_launch_specification

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

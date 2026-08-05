package ocean_right_sizing_cluster_config

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/right_sizing_cluster_config"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

const resourceAffinity commons.ResourceAffinity = "Ocean_Right_Sizing_Cluster_Config"

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[OceanId] = commons.NewGenericField(
		resourceAffinity,
		OceanId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			wrapper := resourceObject.(*commons.RightSizingClusterConfigWrapper)
			input := wrapper.GetRightsizingClusterConfigurationInput()

			if v, ok := resourceData.GetOk(string(OceanId)); ok && v.(string) != "" {
				input.OceanId = spotinst.String(v.(string))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			wrapper := resourceObject.(*commons.RightSizingClusterConfigWrapper)
			input := wrapper.GetRightsizingClusterConfigurationInput()

			if v, ok := resourceData.GetOk(string(OceanId)); ok && v.(string) != "" {
				input.OceanId = spotinst.String(v.(string))
			}
			return nil
		},
		nil,
	)

	fieldsMap[ClusterIdentifier] = commons.NewGenericField(
		resourceAffinity,
		ClusterIdentifier,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			wrapper := resourceObject.(*commons.RightSizingClusterConfigWrapper)
			input := wrapper.GetRightsizingClusterConfigurationInput()

			if v, ok := resourceData.GetOk(string(ClusterIdentifier)); ok && v.(string) != "" {
				input.ClusterIdentifier = spotinst.String(v.(string))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			wrapper := resourceObject.(*commons.RightSizingClusterConfigWrapper)
			input := wrapper.GetRightsizingClusterConfigurationInput()

			if v, ok := resourceData.GetOk(string(ClusterIdentifier)); ok && v.(string) != "" {
				input.ClusterIdentifier = spotinst.String(v.(string))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Config] = commons.NewGenericField(
		resourceAffinity,
		Config,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(AdjustLimitOnDownsize): {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
					string(DownsideOnly): {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
					string(RecommendationsCpuPercentile): {
						Type:     schema.TypeInt,
						Optional: true,
					},
					string(RecommendationsMemoryPercentile): {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			wrapper := resourceObject.(*commons.RightSizingClusterConfigWrapper)
			config := wrapper.GetRightsizingClusterConfiguration()

			if err := resourceData.Set(string(Config), flattenRightsizingClusterConfiguration(config)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Config), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			wrapper := resourceObject.(*commons.RightSizingClusterConfigWrapper)
			input := wrapper.GetRightsizingClusterConfigurationInput()

			config, err := expandRightsizingClusterConfiguration(resourceData.Get(string(Config)))
			if err != nil {
				return err
			}
			input.Config = config
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			wrapper := resourceObject.(*commons.RightSizingClusterConfigWrapper)
			input := wrapper.GetRightsizingClusterConfigurationInput()

			config, err := expandRightsizingClusterConfiguration(resourceData.Get(string(Config)))
			if err != nil {
				return err
			}
			input.Config = config
			return nil
		},
		nil,
	)
}

func flattenRightsizingClusterConfiguration(config *right_sizing_cluster_config.RightsizingClusterConfiguration) []interface{} {
	if config == nil {
		return nil
	}

	m := make(map[string]interface{})

	if config.AdjustLimitOnDownsize != nil {
		m[string(AdjustLimitOnDownsize)] = spotinst.BoolValue(config.AdjustLimitOnDownsize)
	}
	if config.DownsideOnly != nil {
		m[string(DownsideOnly)] = spotinst.BoolValue(config.DownsideOnly)
	}
	if config.RecommendationsCpuPercentile != nil {
		m[string(RecommendationsCpuPercentile)] = spotinst.IntValue(config.RecommendationsCpuPercentile)
	}
	if config.RecommendationsMemoryPercentile != nil {
		m[string(RecommendationsMemoryPercentile)] = spotinst.IntValue(config.RecommendationsMemoryPercentile)
	}

	return []interface{}{m}
}

func expandRightsizingClusterConfiguration(data interface{}) (*right_sizing_cluster_config.RightsizingClusterConfiguration, error) {
	list, ok := data.([]interface{})
	if !ok || len(list) == 0 || list[0] == nil {
		return nil, nil
	}

	m, ok := list[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid config format")
	}

	config := &right_sizing_cluster_config.RightsizingClusterConfiguration{}

	if v, ok := m[string(AdjustLimitOnDownsize)].(bool); ok {
		config.SetAdjustLimitOnDownsize(spotinst.Bool(v))
	}
	if v, ok := m[string(DownsideOnly)].(bool); ok {
		config.SetDownsideOnly(spotinst.Bool(v))
	}
	if v, ok := m[string(RecommendationsCpuPercentile)].(int); ok {
		config.SetRecommendationsCpuPercentile(spotinst.Int(v))
	}
	if v, ok := m[string(RecommendationsMemoryPercentile)].(int); ok {
		config.SetRecommendationsMemoryPercentile(spotinst.Int(v))
	}

	return config, nil
}

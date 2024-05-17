package oceancd_rollout_spec_spot_deployment

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[SpotDeployment] = commons.NewGenericField(
		commons.OceanCDRolloutSpecSpotDeployment,
		SpotDeployment,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(SpotDeploymentName): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(Namespace): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(ClusterId): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			var result []interface{} = nil

			if rolloutSpec != nil && rolloutSpec.SpotDeployment != nil {
				result = flattenSpotDeployment(rolloutSpec.SpotDeployment)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(SpotDeployment), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SpotDeployment), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			var value *oceancd.SpotDeployment = nil

			if v, ok := resourceData.GetOkExists(string(SpotDeployment)); ok {
				if spotDeployment, err := expandSpotDeployment(v); err != nil {
					return err
				} else {
					value = spotDeployment
				}
			}
			rolloutSpec.SetSpotDeployment(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			var value *oceancd.SpotDeployment = nil
			if v, ok := resourceData.GetOkExists(string(SpotDeployment)); ok {
				if spotDeployment, err := expandSpotDeployment(v); err != nil {
					return err
				} else {
					value = spotDeployment
				}
			}
			rolloutSpec.SetSpotDeployment(value)
			return nil
		},
		nil,
	)
	fieldsMap[SpotDeployments] = commons.NewGenericField(
		commons.OceanCDRolloutSpecSpotDeployment,
		SpotDeployment,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(SpotDeploymentName): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(Namespace): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(ClusterId): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()

			var spotDeploymentsResults []interface{} = nil
			if rolloutSpec != nil && rolloutSpec.SpotDeployments != nil {
				spotDeployments := rolloutSpec.SpotDeployments
				spotDeploymentsResults = flattenSpotDeployments(spotDeployments)
			}

			if err := resourceData.Set(string(SpotDeployments), spotDeploymentsResults); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SpotDeployments), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			if value, ok := resourceData.GetOkExists(string(SpotDeployments)); ok {
				if spotDeployments, err := expandSpotDeployments(value); err != nil {
					return err
				} else {
					rolloutSpec.SetSpotDeployments(spotDeployments)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rolloutSpecWrapper := resourceObject.(*commons.OceanCDRolloutSpecWrapper)
			rolloutSpec := rolloutSpecWrapper.GetRolloutSpec()
			var result []*oceancd.SpotDeployment = nil
			if value, ok := resourceData.GetOkExists(string(SpotDeployments)); ok {
				if spotDeployments, err := expandSpotDeployments(value); err != nil {
					return err
				} else {
					result = spotDeployments
				}
			}

			if len(result) == 0 {
				rolloutSpec.SetSpotDeployments(nil)
			} else {
				rolloutSpec.SetSpotDeployments(result)
			}

			return nil
		},
		nil,
	)
}

func expandSpotDeployment(data interface{}) (*oceancd.SpotDeployment, error) {
	if list := data.([]interface{}); len(list) > 0 {
		spotDeployment := &oceancd.SpotDeployment{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(ClusterId)].(string); ok && v != "" {
				spotDeployment.SetClusterId(spotinst.String(v))
			} else {
				spotDeployment.SetClusterId(nil)
			}

			if v, ok := m[string(Namespace)].(string); ok && v != "" {
				spotDeployment.SetNamespace(spotinst.String(v))
			} else {
				spotDeployment.SetNamespace(nil)
			}

			if v, ok := m[string(SpotDeploymentName)].(string); ok && v != "" {
				spotDeployment.SetName(spotinst.String(v))
			} else {
				spotDeployment.SetName(nil)
			}
		}
		return spotDeployment, nil
	}
	return nil, nil
}

func expandSpotDeployments(data interface{}) ([]*oceancd.SpotDeployment, error) {

	list := data.(*schema.Set).List()

	spotDeployments := make([]*oceancd.SpotDeployment, 0, len(list))

	for _, v := range list {

		m := v.(map[string]interface{})

		spotDeployment := &oceancd.SpotDeployment{}

		if v, ok := m[string(SpotDeploymentName)].(string); ok && v != "" {
			spotDeployment.SetName(spotinst.String(v))
		} else {
			spotDeployment.SetName(nil)
		}

		if v, ok := m[string(Namespace)].(string); ok && v != "" {
			spotDeployment.SetNamespace(spotinst.String(v))
		} else {
			spotDeployment.SetNamespace(nil)
		}

		if v, ok := m[string(ClusterId)].(string); ok && v != "" {
			spotDeployment.SetClusterId(spotinst.String(v))
		} else {
			spotDeployment.SetClusterId(nil)
		}
		spotDeployments = append(spotDeployments, spotDeployment)

	}
	return spotDeployments, nil

}

func flattenSpotDeployments(spotDeployments []*oceancd.SpotDeployment) []interface{} {

	m := make([]interface{}, 0, len(spotDeployments))

	for _, spotDeployment := range spotDeployments {

		result := make(map[string]interface{})

		result[string(SpotDeploymentName)] = spotinst.StringValue(spotDeployment.Name)

		result[string(Namespace)] = spotinst.StringValue(spotDeployment.Namespace)

		result[string(ClusterId)] = spotinst.StringValue(spotDeployment.ClusterId)

		m = append(m, result)
	}
	return m
}

func flattenSpotDeployment(spotDeployment *oceancd.SpotDeployment) []interface{} {
	var response []interface{}

	if spotDeployment != nil {
		result := make(map[string]interface{})

		result[string(SpotDeploymentName)] = spotinst.StringValue(spotDeployment.Name)

		result[string(Namespace)] = spotinst.StringValue(spotDeployment.Namespace)

		result[string(ClusterId)] = spotinst.StringValue(spotDeployment.ClusterId)

		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}

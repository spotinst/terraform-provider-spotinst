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
			Type:          schema.TypeList,
			Optional:      true,
			ConflictsWith: []string{string(SpotDeployments)},
			MaxItems:      1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(SpotDeploymentName): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(SpotDeploymentNamespace): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(SpotDeploymentClusterId): {
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
		SpotDeployments,
		&schema.Schema{
			Type:          schema.TypeSet,
			Optional:      true,
			ConflictsWith: []string{string(SpotDeployment)},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(SpotDeploymentsName): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(SpotDeploymentsNamespace): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(SpotDeploymentsClusterId): {
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
			var spotDeploymentList []*oceancd.SpotDeployment = nil
			if value, ok := resourceData.GetOk(string(SpotDeployments)); ok {
				if args, err := expandSpotDeployments(value); err != nil {
					return err
				} else {
					spotDeploymentList = args
				}
			}
			rolloutSpec.SetSpotDeployments(spotDeploymentList)
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

			if v, ok := m[string(SpotDeploymentClusterId)].(string); ok && v != "" {
				spotDeployment.SetClusterId(spotinst.String(v))
			} else {
				spotDeployment.SetClusterId(nil)
			}

			if v, ok := m[string(SpotDeploymentNamespace)].(string); ok && v != "" {
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

		if v, ok := m[string(SpotDeploymentsName)].(string); ok && v != "" {
			spotDeployment.SetName(spotinst.String(v))
		} else {
			spotDeployment.SetName(nil)
		}

		if v, ok := m[string(SpotDeploymentsNamespace)].(string); ok && v != "" {
			spotDeployment.SetNamespace(spotinst.String(v))
		} else {
			spotDeployment.SetNamespace(nil)
		}

		if v, ok := m[string(SpotDeploymentsClusterId)].(string); ok && v != "" {
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

		result[string(SpotDeploymentsName)] = spotinst.StringValue(spotDeployment.Name)

		result[string(SpotDeploymentsNamespace)] = spotinst.StringValue(spotDeployment.Namespace)

		result[string(SpotDeploymentsClusterId)] = spotinst.StringValue(spotDeployment.ClusterId)

		m = append(m, result)
	}
	return m
}

func flattenSpotDeployment(spotDeployment *oceancd.SpotDeployment) []interface{} {
	var response []interface{}

	if spotDeployment != nil {
		result := make(map[string]interface{})

		result[string(SpotDeploymentName)] = spotinst.StringValue(spotDeployment.Name)

		result[string(SpotDeploymentNamespace)] = spotinst.StringValue(spotDeployment.Namespace)

		result[string(SpotDeploymentClusterId)] = spotinst.StringValue(spotDeployment.ClusterId)

		if len(result) > 0 {
			response = append(response, result)
		}
	}
	return response
}

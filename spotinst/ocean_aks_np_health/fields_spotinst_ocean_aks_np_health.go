package ocean_aks_np_health

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Health] = commons.NewGenericField(
		commons.OceanAKSNPHealth,
		Health,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(GracePeriod): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var result []interface{} = nil

			if cluster != nil && cluster.Health != nil {
				result = flattenHealth(cluster.Health)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Health), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Health), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *azure_np.Health = nil

			if v, ok := resourceData.GetOk(string(Health)); ok {
				if health, err := expandHealth(v); err != nil {
					return err
				} else {
					value = health
				}
			}
			cluster.SetHealth(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSNPClusterWrapper)
			cluster := clusterWrapper.GetNPCluster()
			var value *azure_np.Health = nil

			if v, ok := resourceData.GetOk(string(Health)); ok {
				if health, err := expandHealth(v); err != nil {
					return err
				} else {
					value = health
				}
			}
			cluster.SetHealth(value)
			return nil
		},
		nil,
	)
}

func expandHealth(data interface{}) (*azure_np.Health, error) {
	if list := data.([]interface{}); len(list) > 0 {
		health := &azure_np.Health{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(GracePeriod)].(int); ok && v > -1 {
				health.SetGracePeriod(spotinst.Int(v))
			} else {
				health.SetGracePeriod(nil)
			}
		}
		return health, nil
	}
	return nil, nil
}

func flattenHealth(health *azure_np.Health) []interface{} {
	var out []interface{}

	if health != nil {
		result := make(map[string]interface{})

		if health.GracePeriod != nil {
			result[string(GracePeriod)] = spotinst.IntValue(health.GracePeriod)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

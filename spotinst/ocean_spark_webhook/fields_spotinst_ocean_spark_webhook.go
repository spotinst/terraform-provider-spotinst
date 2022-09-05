package ocean_spark_webhook

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/spark"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Webhook] = commons.NewGenericField(
		commons.OceanSparkWebhook,
		Webhook,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(UseHostNetwork): {
						Type:     schema.TypeBool,
						Optional: true,
						Computed: true,
					},

					string(HostNetworkPorts): {
						Type:     schema.TypeList,
						Optional: true,
						Computed: true,
						Elem:     &schema.Schema{Type: schema.TypeInt},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster.Config != nil && cluster.Config.Webhook != nil {
				result = flattenWebhook(cluster.Config.Webhook)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Webhook), result); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(Webhook), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if value, ok := resourceData.GetOk(string(Webhook)); ok {
				if webhook, err := expandWebhook(value, false); err != nil {
					return err
				} else {
					if cluster.Config == nil {
						cluster.Config = &spark.Config{}
					}
					cluster.Config.SetWebhook(webhook)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *spark.WebhookConfig = nil
			if v, ok := resourceData.GetOk(string(Webhook)); ok {
				if webhook, err := expandWebhook(v, true); err != nil {
					return err
				} else {
					value = webhook
				}
			}
			if cluster.Config == nil {
				cluster.Config = &spark.Config{}
			}
			cluster.Config.SetWebhook(value)
			return nil
		},
		nil,
	)
}

func flattenWebhook(webhook *spark.WebhookConfig) []interface{} {
	if webhook == nil {
		return nil
	}
	result := make(map[string]interface{})
	result[string(UseHostNetwork)] = spotinst.BoolValue(webhook.UseHostNetwork)
	if webhook.HostNetworkPorts != nil {
		ports := make([]int, len(webhook.HostNetworkPorts))
		for i := range webhook.HostNetworkPorts {
			ports[i] = spotinst.IntValue(webhook.HostNetworkPorts[i])
		}
		result[string(HostNetworkPorts)] = ports
	}
	return []interface{}{result}
}

func expandWebhook(data interface{}, nullify bool) (*spark.WebhookConfig, error) {
	webhook := &spark.WebhookConfig{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return webhook, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(UseHostNetwork)].(bool); ok {
		webhook.SetUseHostNetwork(spotinst.Bool(v))
	}

	if v, ok := m[string(HostNetworkPorts)]; ok {
		ports, err := expandHostNetworkPorts(v)
		if err != nil {
			return nil, err
		}
		if ports != nil && len(ports) > 0 {
			webhook.SetHostNetworkPorts(ports)
		} else {
			if nullify {
				webhook.SetHostNetworkPorts(nil)
			}
		}
	}

	return webhook, nil
}

func expandHostNetworkPorts(data interface{}) ([]*int, error) {
	list := data.([]interface{})
	result := make([]*int, 0, len(list))
	for _, v := range list {
		if port, ok := v.(int); ok {
			result = append(result, spotinst.Int(port))
		}
	}

	return result, nil
}

package elastigroup_aws_integrations

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupRancher(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationRancher] = commons.NewGenericField(
		commons.ElastigroupAWSIntegrations,
		IntegrationRancher,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MasterHost): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(AccessKey): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(SecretKey): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(Version): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{} = nil
			if elastigroup.Integration != nil && elastigroup.Integration.Rancher != nil {
				value = flattenAWSGroupRancherIntegration(elastigroup.Integration.Rancher)
			}
			if value != nil {
				if err := resourceData.Set(string(IntegrationRancher), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationRancher), err)
				}
			} else {
				if err := resourceData.Set(string(IntegrationRancher), []*aws.RancherIntegration{}); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationRancher), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(IntegrationRancher)); ok {
				if integration, err := expandAWSGroupRancherIntegration(v); err != nil {
					return err
				} else {
					elastigroup.Integration.SetRancher(integration)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.RancherIntegration = nil
			if v, ok := resourceData.GetOk(string(IntegrationRancher)); ok {
				if integration, err := expandAWSGroupRancherIntegration(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetRancher(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func flattenAWSGroupRancherIntegration(integration *aws.RancherIntegration) []interface{} {
	result := make(map[string]interface{})
	result[string(MasterHost)] = spotinst.StringValue(integration.MasterHost)
	result[string(AccessKey)] = spotinst.StringValue(integration.AccessKey)
	result[string(SecretKey)] = spotinst.StringValue(integration.SecretKey)
	result[string(Version)] = spotinst.StringValue(integration.Version)
	return []interface{}{result}
}

func expandAWSGroupRancherIntegration(data interface{}) (*aws.RancherIntegration, error) {
	integration := &aws.RancherIntegration{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(MasterHost)].(string); ok && v != "" {
			integration.SetMasterHost(spotinst.String(v))
		}

		if v, ok := m[string(AccessKey)].(string); ok && v != "" {
			integration.SetAccessKey(spotinst.String(v))
		}

		if v, ok := m[string(SecretKey)].(string); ok && v != "" {
			integration.SetSecretKey(spotinst.String(v))
		}

		if v, ok := m[string(Version)].(string); ok && v != "" {
			integration.SetVersion(spotinst.String(v))
		}
	}
	return integration, nil
}

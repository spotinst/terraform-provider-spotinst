package elastigroup_integrations

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupRancher(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationRancher] = commons.NewGenericField(
		commons.ElastigroupIntegrations,
		IntegrationRancher,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MasterHost): &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},

					string(AccessKey): &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},

					string(SecretKey): &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
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
			elastigroup := resourceObject.(*aws.Group)
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
			elastigroup := resourceObject.(*aws.Group)
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
	return []interface{}{result}
}

func expandAWSGroupRancherIntegration(data interface{}) (*aws.RancherIntegration, error) {
	list := data.([]interface{})
	m := list[0].(map[string]interface{})
	i := &aws.RancherIntegration{}

	if v, ok := m[string(MasterHost)].(string); ok && v != "" {
		i.SetMasterHost(spotinst.String(v))
	}

	if v, ok := m[string(AccessKey)].(string); ok && v != "" {
		i.SetAccessKey(spotinst.String(v))
	}

	if v, ok := m[string(SecretKey)].(string); ok && v != "" {
		i.SetSecretKey(spotinst.String(v))
	}
	return i, nil
}
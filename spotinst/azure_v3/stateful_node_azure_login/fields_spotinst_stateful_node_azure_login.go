package stateful_node_azure_login

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Login] = commons.NewGenericField(
		commons.StatefulNodeAzureLogin,
		Login,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(UserName): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(Password): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					string(SSHPublicKey): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var result []interface{} = nil
			if statefulNode.Compute != nil && statefulNode.Compute.LaunchSpecification != nil && statefulNode.Compute.LaunchSpecification.Login != nil {
				result = flattenLogin(statefulNode.Compute.LaunchSpecification.Login)
			}

			if err := resourceData.Set(string(Login), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Login), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(Login)); ok {
				if login, err := expandLogin(v); err != nil {
					return err
				} else {
					statefulNode.Compute.LaunchSpecification.SetLogin(login)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(Login)); ok {
				if login, err := expandLogin(v); err != nil {
					return err
				} else {
					statefulNode.Compute.LaunchSpecification.SetLogin(login)
				}
			}
			return nil
		},
		nil,
	)
}

func flattenLogin(login *azure.Login) []interface{} {
	result := make(map[string]interface{})
	result[string(UserName)] = spotinst.StringValue(login.UserName)
	result[string(SSHPublicKey)] = spotinst.StringValue(login.SSHPublicKey)
	result[string(Password)] = spotinst.StringValue(login.Password)
	return []interface{}{result}
}

func expandLogin(data interface{}) (*azure.Login, error) {
	login := &azure.Login{}
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(UserName)].(string); ok && v != "" {
			login.SetUserName(spotinst.String(v))
		}

		if v, ok := m[string(SSHPublicKey)].(string); ok && v != "" {
			login.SetSSHPublicKey(spotinst.String(v))
		}

		if v, ok := m[string(Password)].(string); ok && v != "" {
			login.SetPassword(spotinst.String(v))
		}
	}
	return login, nil
}

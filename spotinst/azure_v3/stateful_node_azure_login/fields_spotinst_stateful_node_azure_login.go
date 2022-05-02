package stateful_node_azure_login

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
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
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(UserName): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(Password): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true, //TODO - check
					},
					string(SSHPublicKey): {
						Type:     schema.TypeInt,
						Optional: true,
						Computed: true, //TODO - check
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var result []interface{} = nil
			if statefulNode.Compute != nil && statefulNode.Compute.LaunchSpecification != nil && statefulNode.Compute.LaunchSpecification.Login != nil {
				login := statefulNode.Compute.LaunchSpecification.Login
				result = flattenLogin(login)
			}
			if result != nil {
				if err := resourceData.Set(string(Login), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Login), err)
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
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *azurev3.Login = nil

			if v, ok := resourceData.GetOk(string(Login)); ok {
				if login, err := expandLogin(v); err != nil {
					return err
				} else {
					value = login
				}
			}
			statefulNode.Compute.LaunchSpecification.SetLogin(value)
			return nil
		},
		nil,
	)
}

func flattenLogin(login *azurev3.Login) []interface{} {
	var out []interface{}

	if login != nil {
		result := make(map[string]interface{})

		if login.UserName != nil {
			result[string(UserName)] = spotinst.StringValue(login.UserName)
		}

		if login.Password != nil {
			result[string(Password)] = spotinst.StringValue(login.Password)
		}

		if login.SSHPublicKey != nil {
			result[string(SSHPublicKey)] = spotinst.StringValue(login.SSHPublicKey)
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

func expandLogin(data interface{}) (*azurev3.Login, error) {
	if list := data.([]interface{}); len(list) > 0 {
		login := &azurev3.Login{}

		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(UserName)].(string); ok && v != "" {
				login.SetUserName(spotinst.String(v))
			} else {
				login.SetUserName(nil)
			}

			if v, ok := m[string(Password)].(string); ok && v != "" {
				login.SetPassword(spotinst.String(v))
			} else {
				login.SetPassword(nil)
			}

			if v, ok := m[string(SSHPublicKey)].(string); ok && v != "" {
				login.SetSSHPublicKey(spotinst.String(v))
			} else {
				login.SetSSHPublicKey(nil)
			}
		}

		return login, nil
	}

	return nil, nil
}

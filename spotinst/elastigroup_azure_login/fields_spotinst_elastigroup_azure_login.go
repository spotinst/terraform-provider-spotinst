package elastigroup_azure_login

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Login] = commons.NewGenericField(
		commons.ElastigroupAzureLogin,
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
					},

					string(SSHPublicKey): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{} = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Login != nil {
				value = flattenAzureGroupLogin(elastigroup.Compute.LaunchSpecification.Login)
			}
			if err := resourceData.Set(string(Login), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Login), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Login)); ok {
				if login, err := expandAzureGroupLogin(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetLogin(login)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Login)); ok {
				if login, err := expandAzureGroupLogin(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetLogin(login)
				}
			}
			return nil
		},
		nil,
	)

}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func flattenAzureGroupLogin(login *azure.Login) []interface{} {
	result := make(map[string]interface{})
	result[string(UserName)] = spotinst.StringValue(login.UserName)
	result[string(SSHPublicKey)] = spotinst.StringValue(login.SSHPublicKey)
	result[string(Password)] = spotinst.StringValue(login.Password)
	return []interface{}{result}
}

func expandAzureGroupLogin(data interface{}) (*azure.Login, error) {
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

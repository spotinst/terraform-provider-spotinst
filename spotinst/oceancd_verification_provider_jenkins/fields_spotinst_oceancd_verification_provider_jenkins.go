package oceancd_verification_provider_jenkins

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Jenkins] = commons.NewGenericField(
		commons.OceanCDVerificationProviderJenkins,
		Jenkins,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ApiToken): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(BaseUrl): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(UserName): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var result []interface{} = nil

			if verificationProvider != nil && verificationProvider.Jenkins != nil {
				result = flattenJenkins(verificationProvider.Jenkins)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Jenkins), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Jenkins), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var value *oceancd.Jenkins = nil

			if v, ok := resourceData.GetOk(string(Jenkins)); ok {
				if jenkins, err := expandJenkins(v); err != nil {
					return err
				} else {
					value = jenkins
				}
			}
			verificationProvider.SetJenkins(value)
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var value *oceancd.Jenkins = nil

			if v, ok := resourceData.GetOk(string(Jenkins)); ok {
				if jenkins, err := expandJenkins(v); err != nil {
					return err
				} else {
					value = jenkins
				}
			}
			verificationProvider.SetJenkins(value)
			return nil
		},
		nil,
	)
}

func expandJenkins(data interface{}) (*oceancd.Jenkins, error) {
	jenkins := &oceancd.Jenkins{}
	list := data.(*schema.Set).List()
	if len(list) > 0 {
		if list != nil && list[0] != nil {
			result := list[0].(map[string]interface{})

			if v, ok := result[string(ApiToken)].(string); ok && v != "" {
				jenkins.SetApiToken(spotinst.String(v))
			} else {
				jenkins.SetApiToken(nil)
			}

			if v, ok := result[string(BaseUrl)].(string); ok && v != "" {
				jenkins.SetBaseUrl(spotinst.String(v))
			} else {
				jenkins.SetBaseUrl(nil)
			}

			if v, ok := result[string(UserName)].(string); ok && v != "" {
				jenkins.SetUserName(spotinst.String(v))
			} else {
				jenkins.SetUserName(nil)
			}
		}
		return jenkins, nil
	}
	return nil, nil
}

func flattenJenkins(jenkins_vp *oceancd.Jenkins) []interface{} {
	jenkins := make(map[string]interface{})
	jenkins[string(ApiToken)] = spotinst.StringValue(jenkins_vp.ApiToken)
	jenkins[string(BaseUrl)] = spotinst.StringValue(jenkins_vp.BaseUrl)
	jenkins[string(UserName)] = spotinst.StringValue(jenkins_vp.UserName)
	return []interface{}{jenkins}
}

package stateful_node_azure_login

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[SSHPublicKey] = commons.NewGenericField(
		commons.StatefulNodeAzureLogin,
		SSHPublicKey,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil

			if st != nil && st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.Login != nil && st.Compute.LaunchSpecification.Login.SSHPublicKey != nil {
				value = st.Compute.LaunchSpecification.Login.SSHPublicKey
			}
			if err := resourceData.Set(string(SSHPublicKey), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SSHPublicKey), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()

			if v, ok := resourceData.GetOk(string(SSHPublicKey)); ok {
				st.Compute.LaunchSpecification.Login.SetSSHPublicKey(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[UserName] = commons.NewGenericField(
		commons.StatefulNodeAzureLogin,
		UserName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil

			if st != nil && st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.Login != nil && st.Compute.LaunchSpecification.Login.UserName != nil {
				value = st.Compute.LaunchSpecification.Login.UserName
			}
			if err := resourceData.Set(string(UserName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UserName), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()

			if v, ok := resourceData.GetOk(string(UserName)); ok {
				st.Compute.LaunchSpecification.Login.SetUserName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[Password] = commons.NewGenericField(
		commons.StatefulNodeAzureLogin,
		Password,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value *string = nil

			if st != nil && st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.Login != nil && st.Compute.LaunchSpecification.Login.Password != nil {
				value = st.Compute.LaunchSpecification.Login.Password
			}
			if err := resourceData.Set(string(Password), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Password), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()

			if v, ok := resourceData.GetOk(string(Password)); ok {
				st.Compute.LaunchSpecification.Login.SetPassword(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)
}

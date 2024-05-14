package oceancd_verification_template_args

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[ArgName] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateArgs,
		ArgName,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if err := resourceData.Set(string(ArgName), spotinst.StringValue(verificationTemplate.Args[].Name)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(ArgName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(ArgName)); ok {
				verificationTemplate.Args[].SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(ArgName)); ok {
				verificationTemplate.Args[].SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Value] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateArgs,
		Value,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if err := resourceData.Set(string(Value), spotinst.StringValue(verificationTemplate.Args[].Value)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(Value), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(Value)); ok {
				verificationTemplate.Args[].SetValue(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(Value)); ok {
				verificationTemplate.Args[].SetValue(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[ValueFrom] = commons.NewGenericField(
		commons.OceanCDVerificationTemplate,
		ValueFrom,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(SetKeyRef): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Key): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(KeyName): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var result []interface{} = nil

			if verificationTemplate != nil && verificationTemplate.Args != nil &&
				verificationTemplate.Args[].ValueFrom != nil {
				result = flattenValueFrom(verificationTemplate.Args[].ValueFrom)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(ValueFrom), result); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(ValueFrom), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(ValueFrom)); ok {
				if valueFrom, err := expandValueFrom(v); err != nil {
					return err
				} else {
					verificationTemplate.Args[].SetValueFrom(valueFrom)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *oceancd.ValueFrom = nil

			if v, ok := resourceData.GetOk(string(ValueFrom)); ok {
				if valueFrom, err := expandValueFrom(v); err != nil {
					return err
				} else {
					value = valueFrom
				}
			}
			if verificationTemplate.Args == nil {
				verificationTemplate.Args = []*oceancd.Args{}
			}
			verificationTemplate.Args[].SetValueFrom(value)
			return nil
		},
		nil,
	)
}

func expandValueFrom(data interface{}) (*oceancd.ValueFrom, error) {
	valueFrom := &oceancd.ValueFrom{}
	list := data.([]interface{})
	if list == nil || len(list) == 0 {
		return nil, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(SetKeyRef)]; ok && v != nil {
		secretKeyRef, err := expandSecretKeyRef(v)
		if err != nil {
			return nil, err
		}
		if secretKeyRef != nil {
			valueFrom.SetSecretKeyRef(secretKeyRef)
		} else {
			valueFrom.SetSecretKeyRef(nil)
		}
	}
	return valueFrom, nil
}

func expandSecretKeyRef(data interface{}) (*oceancd.SecretKeyRef, error) {

	secretKeyRef := &oceancd.SecretKeyRef{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return secretKeyRef, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Key)].(string); ok && v != "" {
		secretKeyRef.SetKey(spotinst.String(v))
	}

	if v, ok := m[string(KeyName)].(string); ok && v != "" {
		secretKeyRef.SetKey(spotinst.String(v))
	}
	return secretKeyRef, nil
}

func flattenSecretKeyRef(secretKeyRef *oceancd.SecretKeyRef) []interface{} {
	result := make(map[string]interface{})
	result[string(Key)] = spotinst.StringValue(secretKeyRef.Key)
	result[string(KeyName)] = spotinst.StringValue(secretKeyRef.Name)

	return []interface{}{result}
}

func flattenValueFrom(valueFrom *oceancd.ValueFrom) []interface{} {
	result := make(map[string]interface{})

	if valueFrom.SecretKeyRef != nil {
		result[string(SetKeyRef)] = flattenSecretKeyRef(valueFrom.SecretKeyRef)
	}
	return []interface{}{result}
}

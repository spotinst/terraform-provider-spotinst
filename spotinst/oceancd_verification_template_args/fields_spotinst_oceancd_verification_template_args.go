package oceancd_verification_template_args

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Args] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateArgs,
		Args,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ArgName): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(Value): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(ValueFrom): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(SecretKeyRef): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Key): {
												Type:     schema.TypeString,
												Required: true,
											},

											string(Name): {
												Type:     schema.TypeString,
												Required: true,
											},
										},
									},
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

			var argsResults []interface{} = nil
			if verificationTemplate != nil && verificationTemplate.Args != nil {
				args := verificationTemplate.Args
				argsResults = flattenArgs(args)
			}

			if err := resourceData.Set(string(Args), argsResults); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Args), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if value, ok := resourceData.GetOkExists(string(Args)); ok {
				if args, err := expandArgs(value); err != nil {
					return err
				} else {
					verificationTemplate.SetArgs(args)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			virtualNodeGroupWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			virtualNodeGroup := virtualNodeGroupWrapper.GetVerificationTemplate()
			var argList []*oceancd.Args = nil
			if value, ok := resourceData.GetOk(string(Args)); ok {
				if args, err := expandArgs(value); err != nil {
					return err
				} else {
					argList = args
				}
			}
			virtualNodeGroup.SetArgs(argList)
			return nil
		},
		nil,
	)
}

func expandArgs(data interface{}) ([]*oceancd.Args, error) {
	list := data.(*schema.Set).List()
	args := make([]*oceancd.Args, 0, len(list))

	for _, v := range list {
		m := v.(map[string]interface{})
		arg := &oceancd.Args{}

		if v, ok := m[string(ArgName)].(string); ok && v != "" {
			arg.SetName(spotinst.String(v))
		}
		if v, ok := m[string(Value)].(string); ok && v != "" {
			arg.SetValue(spotinst.String(v))
		}

		if v, ok := m[string(ValueFrom)]; ok && v != nil {
			valueFrom, err := expandValueFrom(v)
			if err != nil {
				return nil, err
			}
			if valueFrom != nil {
				arg.SetValueFrom(valueFrom)
			} else {
				arg.SetValueFrom(nil)
			}
		}

		args = append(args, arg)
	}
	return args, nil
}

func expandValueFrom(data interface{}) (*oceancd.ValueFrom, error) {
	valueFrom := &oceancd.ValueFrom{}
	list := data.([]interface{})
	if list == nil || len(list) == 0 {
		return nil, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(SecretKeyRef)]; ok && v != nil {
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

	if v, ok := m[string(Name)].(string); ok && v != "" {
		secretKeyRef.SetName(spotinst.String(v))
	}
	return secretKeyRef, nil
}

func flattenSecretKeyRef(secretKeyRef *oceancd.SecretKeyRef) []interface{} {
	result := make(map[string]interface{})
	result[string(Key)] = spotinst.StringValue(secretKeyRef.Key)
	result[string(Name)] = spotinst.StringValue(secretKeyRef.Name)

	return []interface{}{result}
}

func flattenValueFrom(valueFrom *oceancd.ValueFrom) []interface{} {
	result := make(map[string]interface{})

	if valueFrom.SecretKeyRef != nil {
		result[string(SecretKeyRef)] = flattenSecretKeyRef(valueFrom.SecretKeyRef)
	}
	return []interface{}{result}
}

func flattenArgs(args []*oceancd.Args) []interface{} {
	result := make([]interface{}, 0, len(args))

	for _, arg := range args {
		m := make(map[string]interface{})

		m[string(ArgName)] = spotinst.StringValue(arg.Name)
		m[string(Value)] = spotinst.StringValue(arg.Value)

		if arg.ValueFrom != nil {
			m[string(ValueFrom)] = flattenValueFrom(arg.ValueFrom)
		}
		result = append(result, m)
	}
	return result
}

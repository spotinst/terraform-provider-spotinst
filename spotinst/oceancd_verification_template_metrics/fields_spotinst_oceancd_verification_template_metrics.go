package oceancd_verification_template_args

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[ConsecutiveErrorLimit] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		ConsecutiveErrorLimit,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *int = nil
			if verificationTemplate != nil && verificationTemplate.Metrics != nil && verificationTemplate.Metrics[].ConsecutiveErrorLimit != nil {
				value = verificationTemplate.Metrics[].ConsecutiveErrorLimit
			} else {
				value = spotinst.Int(4)
			}
			if err := resourceData.Set(string(ConsecutiveErrorLimit), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ConsecutiveErrorLimit), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(ConsecutiveErrorLimit)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetConsecutiveErrorLimit(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetConsecutiveErrorLimit(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(ConsecutiveErrorLimit)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetConsecutiveErrorLimit(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetConsecutiveErrorLimit(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[Count] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		Count,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *int = nil
			if verificationTemplate != nil && verificationTemplate.Metrics != nil && verificationTemplate.Metrics[].Count != nil {
				value = verificationTemplate.Metrics[].Count
			} else {
				value = spotinst.Int(1)
			}
			if err := resourceData.Set(string(Count), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Count), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(Count)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetCount(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetCount(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(Count)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetCount(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetCount(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[DryRun] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		DryRun,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *bool = nil
			if verificationTemplate.Metrics[] != nil && verificationTemplate.Metrics[].DryRun != nil {
				value = verificationTemplate.Metrics[].DryRun
			}
			if value != nil {
				if err := resourceData.Set(string(DryRun), spotinst.BoolValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DryRun), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(DryRun)); ok && v != nil {
				dryRuns := v.(bool)
				dryRun := spotinst.Bool(dryRuns)
				verificationTemplate.Metrics[].SetDryRun(dryRun)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var dryRun *bool = nil
			if v, ok := resourceData.GetOk(string(DryRun)); ok && v != nil {
				dryRuns := v.(bool)
				dryRun = spotinst.Bool(dryRuns)
			}
			verificationTemplate.Metrics[].SetDryRun(dryRun)
			return nil
		},
		nil,
	)

	fieldsMap[FailureCondition] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		FailureCondition,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if err := resourceData.Set(string(FailureCondition), spotinst.StringValue(verificationTemplate.Metrics[].FailureCondition)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(FailureCondition), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(FailureCondition)); ok {
				verificationTemplate.Metrics[].SetFailureCondition(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(FailureCondition)); ok {
				verificationTemplate.Metrics[].SetFailureCondition(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[FailureLimit] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		FailureLimit,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *int = nil
			if verificationTemplate != nil && verificationTemplate.Metrics != nil && verificationTemplate.Metrics[].FailureLimit != nil {
				value = verificationTemplate.Metrics[].FailureLimit
			} else {
				value = spotinst.Int(0)
			}
			if err := resourceData.Set(string(FailureLimit), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(FailureLimit), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(FailureLimit)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetFailureLimit(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetFailureLimit(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(FailureLimit)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetFailureLimit(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetFailureLimit(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[FailureLimit] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		FailureLimit,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			var value *int = nil
			if verificationTemplate != nil && verificationTemplate.Metrics != nil && verificationTemplate.Metrics[].FailureLimit != nil {
				value = verificationTemplate.Metrics[].FailureLimit
			} else {
				value = spotinst.Int(0)
			}
			if err := resourceData.Set(string(FailureLimit), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(FailureLimit), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(FailureLimit)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetFailureLimit(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetFailureLimit(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.Get(string(FailureLimit)).(int); ok && v >= 0 {
				verificationTemplate.Metrics[].SetFailureLimit(spotinst.Int(v))
			} else {
				verificationTemplate.Metrics[].SetFailureLimit(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[InitialDelay] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		InitialDelay,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if err := resourceData.Set(string(InitialDelay), spotinst.StringValue(verificationTemplate.Metrics[].InitialDelay)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(InitialDelay), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(InitialDelay)); ok {
				verificationTemplate.Metrics[].SetInitialDelay(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(InitialDelay)); ok {
				verificationTemplate.Metrics[].SetInitialDelay(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Interval] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		Interval,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if err := resourceData.Set(string(Interval), spotinst.StringValue(verificationTemplate.Metrics[].Interval)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(Interval), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(Interval)); ok {
				verificationTemplate.Metrics[].SetInterval(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(Interval)); ok {
				verificationTemplate.Metrics[].SetInterval(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if err := resourceData.Set(string(Name), spotinst.StringValue(verificationTemplate.Metrics[].Name)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(Name)); ok {
				verificationTemplate.Metrics[].SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(Name)); ok {
				verificationTemplate.Metrics[].SetName(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[SuccessCondition] = commons.NewGenericField(
		commons.OceanCDVerificationTemplateMetrics,
		SuccessCondition,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if err := resourceData.Set(string(SuccessCondition), spotinst.StringValue(verificationTemplate.Metrics[].SuccessCondition)); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(SuccessCondition), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(SuccessCondition)); ok {
				verificationTemplate.Metrics[].SetSuccessCondition(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationTemplateWrapper := resourceObject.(*commons.OceanCDVerificationTemplateWrapper)
			verificationTemplate := verificationTemplateWrapper.GetVerificationTemplate()
			if v, ok := resourceData.GetOk(string(SuccessCondition)); ok {
				verificationTemplate.Metrics[].SetSuccessCondition(spotinst.String(v.(string)))
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

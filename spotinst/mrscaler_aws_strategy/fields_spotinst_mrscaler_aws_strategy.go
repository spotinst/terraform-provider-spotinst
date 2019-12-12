package mrscaler_aws_strategy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/mrscaler"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Strategy] = commons.NewGenericField(
		commons.MRScalerAWSStrategy,
		Strategy,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
				v := val.(string)
				if v != "clone" && v != "wrap" && v != "new" {
					errs = append(errs, fmt.Errorf("%q must be one of: new, wrap, clone. Got: %v", key, v))
				}
				return
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(Strategy))
			return err
		},
		nil,
	)

	fieldsMap[ReleaseLabel] = commons.NewGenericField(
		commons.MRScalerAWSStrategy,
		ReleaseLabel,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()

			if strategy, ok := resourceData.Get(string(Strategy)).(string); ok && strategy != "" {
				if strategy == New {
					if scaler.Strategy.CreateNew == nil {
						scaler.Strategy.SetCreateNew(&mrscaler.CreateNew{})
					}
					if label, ok := resourceData.Get(string(ReleaseLabel)).(string); ok && label != "" {
						scaler.Strategy.CreateNew.SetReleaseLabel(spotinst.String(label))
					}
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(ReleaseLabel))
			return err
		},
		nil,
	)

	fieldsMap[Retries] = commons.NewGenericField(
		commons.MRScalerAWSStrategy,
		Retries,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()

			if strategy, ok := resourceData.Get(string(Strategy)).(string); ok && strategy != "" {
				switch strategy {
				case New:
					if scaler.Strategy.CreateNew == nil {
						scaler.Strategy.SetCreateNew(&mrscaler.CreateNew{})
					}
					if retries, ok := resourceData.Get(string(Retries)).(int); ok {
						scaler.Strategy.CreateNew.SetRetries(spotinst.Int(retries))
					}
				case Clone:
					if scaler.Strategy.Cloning == nil {
						scaler.Strategy.SetCloning(&mrscaler.Cloning{})
					}
					if retries, ok := resourceData.Get(string(Retries)).(int); ok {
						scaler.Strategy.Cloning.SetRetries(spotinst.Int(retries))
					}
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(Retries))
			return err
		},
		nil,
	)

	fieldsMap[ProvisioningTimeout] = commons.NewGenericField(
		commons.MRScalerAWSStrategy,
		ProvisioningTimeout,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Timeout): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(TimeoutAction): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if scaler.Strategy != nil && scaler.Strategy.ProvisioningTimeout != nil {
				pt := scaler.Strategy.ProvisioningTimeout
				result := make(map[string]interface{})
				result[string(TimeoutAction)] = spotinst.StringValue(pt.TimeoutAction)
				result[string(Timeout)] = spotinst.IntValue(pt.Timeout)
				if err := resourceData.Set(string(ProvisioningTimeout), []interface{}{result}); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ProvisioningTimeout), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(ProvisioningTimeout)); ok {
				if pt, err := expandProvisioningTimeout(v); err != nil {
					return err
				} else {
					scaler.Strategy.SetProvisioningTimeout(pt)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *mrscaler.ProvisioningTimeout = nil
			if v, ok := resourceData.GetOk(string(ProvisioningTimeout)); ok {
				if pt, err := expandProvisioningTimeout(v); err != nil {
					return err
				} else {
					value = pt
				}
			}
			scaler.Strategy.SetProvisioningTimeout(value)
			return nil
		},
		nil,
	)

}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Utilities
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func expandProvisioningTimeout(data interface{}) (*mrscaler.ProvisioningTimeout, error) {
	pt := &mrscaler.ProvisioningTimeout{}
	list := data.([]interface{})
	for _, item := range list {
		m := item.(map[string]interface{})

		if v, ok := m[string(Timeout)].(int); ok && v >= 15 {
			pt.SetTimeout(spotinst.Int(v))
		}

		if v, ok := m[string(TimeoutAction)].(string); ok && v != "" {
			pt.SetTimeoutAction(spotinst.String(v))
		}

	}
	return pt, nil
}

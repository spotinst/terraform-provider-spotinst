package mrscaler_aws_scaling_policies

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/mrscaler"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupCoreScalingPolicies(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[CoreScalingUpPolicy] = commons.NewGenericField(
		commons.MRScalerAWSCoreScalingPolicies,
		CoreScalingUpPolicy,
		upDownScalingPolicySchema(),
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var policiesResult []interface{} = nil
			if scaler.CoreScaling != nil && scaler.CoreScaling.Up != nil {
				scaleUpPolicies := scaler.CoreScaling.Up
				policiesResult = flattenMRScalerAWSScalingPolicy(scaleUpPolicies)
			}
			if err := resourceData.Set(string(CoreScalingUpPolicy), policiesResult); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CoreScalingUpPolicy), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(CoreScalingUpPolicy)); ok {
				if policies, err := expandMRScalerAWSScalingPolicies(v); err != nil {
					return err
				} else {
					scaler.CoreScaling.SetUp(policies)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value []*mrscaler.ScalingPolicy = nil
			if v, ok := resourceData.GetOk(string(CoreScalingUpPolicy)); ok && v != nil {
				if policies, err := expandMRScalerAWSScalingPolicies(v); err != nil {
					return err
				} else {
					value = policies
				}
			}
			scaler.CoreScaling.SetUp(value)
			return nil
		},
		nil,
	)

	fieldsMap[CoreScalingDownPolicy] = commons.NewGenericField(
		commons.MRScalerAWSCoreScalingPolicies,
		CoreScalingDownPolicy,
		upDownScalingPolicySchema(),
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var policiesResult []interface{} = nil
			if scaler.CoreScaling != nil && scaler.CoreScaling.Down != nil {
				scaleDownPolicies := scaler.CoreScaling.Down
				policiesResult = flattenMRScalerAWSScalingPolicy(scaleDownPolicies)
			}
			if err := resourceData.Set(string(CoreScalingDownPolicy), policiesResult); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CoreScalingDownPolicy), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(CoreScalingDownPolicy)); ok {
				if policies, err := expandMRScalerAWSScalingPolicies(v); err != nil {
					return err
				} else {
					scaler.CoreScaling.SetDown(policies)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()

			var value []*mrscaler.ScalingPolicy = nil
			if v, ok := resourceData.GetOk(string(CoreScalingDownPolicy)); ok && v != nil {
				if policies, err := expandMRScalerAWSScalingPolicies(v); err != nil {
					return err
				} else {
					value = policies
				}
			}
			scaler.CoreScaling.SetDown(value)
			return nil
		},
		nil,
	)
}

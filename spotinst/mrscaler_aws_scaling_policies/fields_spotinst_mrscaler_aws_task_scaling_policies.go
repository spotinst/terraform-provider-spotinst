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
func SetupTaskScalingPolicies(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[TaskScalingUpPolicy] = commons.NewGenericField(
		commons.MRScalerAWSTaskScalingPolicies,
		TaskScalingUpPolicy,
		upDownScalingPolicySchema(),
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var policiesResult []interface{} = nil
			if scaler.Scaling != nil && scaler.Scaling.Up != nil {
				scaleUpPolicies := scaler.Scaling.Up
				policiesResult = flattenMRScalerAWSScalingPolicy(scaleUpPolicies)
			}
			if err := resourceData.Set(string(TaskScalingUpPolicy), policiesResult); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(TaskScalingUpPolicy), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(TaskScalingUpPolicy)); ok {
				if policies, err := expandMRScalerAWSScalingPolicies(v); err != nil {
					return err
				} else {
					scaler.Scaling.SetUp(policies)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value []*mrscaler.ScalingPolicy = nil
			if v, ok := resourceData.GetOk(string(TaskScalingUpPolicy)); ok && v != nil {
				if policies, err := expandMRScalerAWSScalingPolicies(v); err != nil {
					return err
				} else {
					value = policies
				}
			}
			scaler.Scaling.SetUp(value)
			return nil
		},
		nil,
	)

	fieldsMap[TaskScalingDownPolicy] = commons.NewGenericField(
		commons.MRScalerAWSTaskScalingPolicies,
		TaskScalingDownPolicy,
		upDownScalingPolicySchema(),
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var policiesResult []interface{} = nil
			if scaler.Scaling != nil && scaler.Scaling.Down != nil {
				scaleDownPolicies := scaler.Scaling.Down
				policiesResult = flattenMRScalerAWSScalingPolicy(scaleDownPolicies)
			}
			if err := resourceData.Set(string(TaskScalingDownPolicy), policiesResult); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(TaskScalingDownPolicy), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(TaskScalingDownPolicy)); ok {
				if policies, err := expandMRScalerAWSScalingPolicies(v); err != nil {
					return err
				} else {
					scaler.Scaling.SetDown(policies)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value []*mrscaler.ScalingPolicy = nil
			if v, ok := resourceData.GetOk(string(TaskScalingDownPolicy)); ok && v != nil {
				if policies, err := expandMRScalerAWSScalingPolicies(v); err != nil {
					return err
				} else {
					value = policies
				}
			}
			scaler.Scaling.SetDown(value)
			return nil
		},
		nil,
	)
}

package mrscaler_aws_cluster

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

	fieldsMap[LogURI] = commons.NewGenericField(
		commons.MRScalerAWSCluster,
		LogURI,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Cluster != nil && scaler.Cluster.LogURI != nil {
				value = scaler.Cluster.LogURI
			}
			if err := resourceData.Set(string(LogURI), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(LogURI), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(LogURI)); ok && v != "" {
				if scaler.Cluster == nil {
					scaler.SetCluster(&mrscaler.Cluster{})
				}
				scaler.Cluster.SetLogURI(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(LogURI))
			return err
		},
		nil,
	)

	fieldsMap[AdditionalInfo] = commons.NewGenericField(
		commons.MRScalerAWSCluster,
		AdditionalInfo,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Cluster != nil && scaler.Cluster.AdditionalInfo != nil {
				value = scaler.Cluster.AdditionalInfo
			}
			if err := resourceData.Set(string(AdditionalInfo), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AdditionalInfo), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(AdditionalInfo)); ok && v != "" {
				if scaler.Cluster == nil {
					scaler.SetCluster(&mrscaler.Cluster{})
				}
				scaler.Cluster.SetAdditionalInfo(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(AdditionalInfo))
			return err
		},
		nil,
	)

	fieldsMap[JobFlowRole] = commons.NewGenericField(
		commons.MRScalerAWSCluster,
		JobFlowRole,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Cluster != nil && scaler.Cluster.JobFlowRole != nil {
				value = scaler.Cluster.JobFlowRole
			}
			if err := resourceData.Set(string(JobFlowRole), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(JobFlowRole), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(JobFlowRole)); ok && v != "" {
				if scaler.Cluster == nil {
					scaler.SetCluster(&mrscaler.Cluster{})
				}
				scaler.Cluster.SetJobFlowRole(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(JobFlowRole))
			return err
		},
		nil,
	)

	fieldsMap[SecurityConfig] = commons.NewGenericField(
		commons.MRScalerAWSCluster,
		SecurityConfig,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Cluster != nil && scaler.Cluster.SecurityConfiguration != nil {
				value = scaler.Cluster.SecurityConfiguration
			}
			if err := resourceData.Set(string(SecurityConfig), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SecurityConfig), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(SecurityConfig)); ok && v != "" {
				if scaler.Cluster == nil {
					scaler.SetCluster(&mrscaler.Cluster{})
				}
				scaler.Cluster.SetSecurityConfiguration(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(SecurityConfig))
			return err
		},
		nil,
	)

	fieldsMap[ServiceRole] = commons.NewGenericField(
		commons.MRScalerAWSCluster,
		ServiceRole,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Cluster != nil && scaler.Cluster.ServiceRole != nil {
				value = scaler.Cluster.ServiceRole
			}
			if err := resourceData.Set(string(ServiceRole), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ServiceRole), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(ServiceRole)); ok {
				if scaler.Cluster == nil {
					scaler.SetCluster(&mrscaler.Cluster{})
				}
				scaler.Cluster.SetServiceRole(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(ServiceRole))
			return err
		},
		nil,
	)

	fieldsMap[VisibleToAllUsers] = commons.NewGenericField(
		commons.MRScalerAWSCluster,
		VisibleToAllUsers,
		&schema.Schema{
			Type:       schema.TypeBool,
			Optional:   true,
			Deprecated: "This field has been removed from our API and is no longer functional.",
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[TerminationProtected] = commons.NewGenericField(
		commons.MRScalerAWSCluster,
		TerminationProtected,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *bool = nil
			if scaler.Cluster != nil && scaler.Cluster.TerminationProtected != nil {
				value = scaler.Cluster.TerminationProtected
			}
			if err := resourceData.Set(string(TerminationProtected), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(TerminationProtected), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOkExists(string(TerminationProtected)); ok {
				if scaler.Cluster == nil {
					scaler.SetCluster(&mrscaler.Cluster{})
				}
				scaler.Cluster.SetTerminationProtected(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOkExists(string(TerminationProtected)); ok {
				if scaler.Cluster == nil {
					scaler.SetCluster(&mrscaler.Cluster{})
				}
				scaler.Cluster.SetTerminationProtected(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[KeepJobFlowAlive] = commons.NewGenericField(
		commons.MRScalerAWSCluster,
		KeepJobFlowAlive,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *bool = nil
			if scaler.Cluster != nil && scaler.Cluster.KeepJobFlowAliveWhenNoSteps != nil {
				value = scaler.Cluster.KeepJobFlowAliveWhenNoSteps
			}
			if err := resourceData.Set(string(KeepJobFlowAlive), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(KeepJobFlowAlive), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOkExists(string(KeepJobFlowAlive)); ok {
				if scaler.Cluster == nil {
					scaler.SetCluster(&mrscaler.Cluster{})
				}
				scaler.Cluster.SetKeepJobFlowAliveWhenNoSteps(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(KeepJobFlowAlive))
			return err
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utilities
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func flattenMRScalerAWSCluster(cluster *mrscaler.Cluster) []interface{} {
	result := make(map[string]interface{})
	result[string(LogURI)] = spotinst.StringValue(cluster.LogURI)
	result[string(AdditionalInfo)] = spotinst.StringValue(cluster.AdditionalInfo)
	result[string(JobFlowRole)] = spotinst.StringValue(cluster.JobFlowRole)
	result[string(SecurityConfig)] = spotinst.StringValue(cluster.SecurityConfiguration)
	result[string(ServiceRole)] = spotinst.StringValue(cluster.ServiceRole)
	result[string(TerminationProtected)] = spotinst.BoolValue(cluster.TerminationProtected)
	result[string(KeepJobFlowAlive)] = spotinst.BoolValue(cluster.KeepJobFlowAliveWhenNoSteps)
	return []interface{}{result}
}

func expandMRScalerAWSCluster(data interface{}) (*mrscaler.Cluster, error) {
	cluster := &mrscaler.Cluster{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return cluster, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(LogURI)].(string); ok && v != "" {
		cluster.SetLogURI(spotinst.String(v))
	}

	if v, ok := m[string(AdditionalInfo)].(string); ok && v != "" {
		cluster.SetAdditionalInfo(spotinst.String(v))
	}

	if v, ok := m[string(JobFlowRole)].(string); ok && v != "" {
		cluster.SetJobFlowRole(spotinst.String(v))
	}

	if v, ok := m[string(SecurityConfig)].(string); ok && v != "" {
		cluster.SetSecurityConfiguration(spotinst.String(v))
	}

	if v, ok := m[string(ServiceRole)].(string); ok && v != "" {
		cluster.SetServiceRole(spotinst.String(v))
	}

	if v, ok := m[string(TerminationProtected)].(bool); ok {
		cluster.SetTerminationProtected(spotinst.Bool(v))
	}

	if v, ok := m[string(KeepJobFlowAlive)].(bool); ok {
		cluster.SetKeepJobFlowAliveWhenNoSteps(spotinst.Bool(v))
	}

	return cluster, nil
}

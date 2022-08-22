package elastigroup_aws_launch_configuration

import (
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[ImageId] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		ImageId,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.ImageID != nil {
				value = elastigroup.Compute.LaunchSpecification.ImageID
			}
			if err := resourceData.Set(string(ImageId), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ImageId), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(ImageId)).(string); ok && v != "" {
				elastigroup.Compute.LaunchSpecification.SetImageId(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(ImageId)).(string); ok && v != "" {
				elastigroup.Compute.LaunchSpecification.SetImageId(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Images] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		Images,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Id): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Images != nil {
				value = flattenImages(elastigroup.Compute.LaunchSpecification.Images)
			}
			if value != nil {
				if err := resourceData.Set(string(Images), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Images), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Images)); ok {
				if value, err := expandImages(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetImages(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Images)); ok {
				if value, err := expandImages(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetImages(value)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[IamInstanceProfile] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		IamInstanceProfile,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value = ""
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.IAMInstanceProfile != nil {

				lc := elastigroup.Compute.LaunchSpecification
				if lc.IAMInstanceProfile.Arn != nil {
					value = spotinst.StringValue(lc.IAMInstanceProfile.Arn)
				} else if lc.IAMInstanceProfile.Name != nil {
					value = spotinst.StringValue(lc.IAMInstanceProfile.Name)
				}
			}
			if err := resourceData.Set(string(IamInstanceProfile), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IamInstanceProfile), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(IamInstanceProfile)).(string); ok && v != "" {
				iamInstanceProf := &aws.IAMInstanceProfile{}

				if InstanceProfileArnRegex.MatchString(v) {
					iamInstanceProf.SetArn(spotinst.String(v))
				} else {
					iamInstanceProf.SetName(spotinst.String(v))
				}
				elastigroup.Compute.LaunchSpecification.SetIAMInstanceProfile(iamInstanceProf)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(IamInstanceProfile)).(string); ok && v != "" {
				iamInstanceProf := &aws.IAMInstanceProfile{}
				if InstanceProfileArnRegex.MatchString(v) {
					iamInstanceProf.SetArn(spotinst.String(v))
				} else {
					iamInstanceProf.SetName(spotinst.String(v))
				}
				elastigroup.Compute.LaunchSpecification.SetIAMInstanceProfile(iamInstanceProf)
			} else {
				elastigroup.Compute.LaunchSpecification.SetIAMInstanceProfile(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[KeyName] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		KeyName,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.KeyPair != nil {
				value = elastigroup.Compute.LaunchSpecification.KeyPair
			}
			if err := resourceData.Set(string(KeyName), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(KeyName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(KeyName)).(string); ok && v != "" {
				elastigroup.Compute.LaunchSpecification.SetKeyPair(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if v, ok := resourceData.Get(string(KeyName)).(string); ok && v != "" {
				value = spotinst.String(v)
			}
			elastigroup.Compute.LaunchSpecification.SetKeyPair(value)
			return nil
		},
		nil,
	)

	fieldsMap[SecurityGroups] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		SecurityGroups,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.SecurityGroupIDs != nil {
				value = elastigroup.Compute.LaunchSpecification.SecurityGroupIDs
			}
			if err := resourceData.Set(string(SecurityGroups), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SecurityGroups), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(SecurityGroups)).([]interface{}); ok {
				ids := make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
				elastigroup.Compute.LaunchSpecification.SetSecurityGroupIDs(ids)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(SecurityGroups)).([]interface{}); ok {
				ids := make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
				elastigroup.Compute.LaunchSpecification.SetSecurityGroupIDs(ids)
			}
			return nil
		},
		nil,
	)

	fieldsMap[UserData] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		UserData,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				// Sometimes the EC2 API responds with the equivalent, empty SHA1 sum
				if (old == "da39a3ee5e6b4b0d3255bfef95601890afd80709" && new == "") ||
					(old == "" && new == "da39a3ee5e6b4b0d3255bfef95601890afd80709") {
					return true
				}
				return false
			},
			StateFunc: Base64StateFunc,
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value = ""
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.UserData != nil {

				userData := elastigroup.Compute.LaunchSpecification.UserData
				userDataValue := spotinst.StringValue(userData)
				if userDataValue != "" {
					if isBase64Encoded(resourceData.Get(string(UserData)).(string)) {
						value = userDataValue
					} else {
						decodedUserData, _ := base64.StdEncoding.DecodeString(userDataValue)
						value = string(decodedUserData)
					}
				}
			}
			if err := resourceData.Set(string(UserData), Base64StateFunc(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UserData), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData := spotinst.String(base64Encode(v))
				elastigroup.Compute.LaunchSpecification.SetUserData(userData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var userData *string = nil
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData = spotinst.String(base64Encode(v))
			}
			elastigroup.Compute.LaunchSpecification.SetUserData(userData)
			return nil
		},
		nil,
	)

	fieldsMap[ShutdownScript] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		ShutdownScript,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				// Sometimes the EC2 API responds with the equivalent, empty SHA1 sum
				if (old == "da39a3ee5e6b4b0d3255bfef95601890afd80709" && new == "") ||
					(old == "" && new == "da39a3ee5e6b4b0d3255bfef95601890afd80709") {
					return true
				}
				return false
			},
			StateFunc: Base64StateFunc,
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value = ""
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.ShutdownScript != nil {

				shutdownScript := elastigroup.Compute.LaunchSpecification.ShutdownScript
				shutdownScriptValue := spotinst.StringValue(shutdownScript)
				if shutdownScriptValue != "" {
					if isBase64Encoded(resourceData.Get(string(ShutdownScript)).(string)) {
						value = shutdownScriptValue
					} else {
						decodedShutdownScript, _ := base64.StdEncoding.DecodeString(shutdownScriptValue)
						value = string(decodedShutdownScript)
					}
				}
			}
			if err := resourceData.Set(string(ShutdownScript), Base64StateFunc(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShutdownScript), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(ShutdownScript)).(string); ok && v != "" {
				shutdownScript := spotinst.String(base64Encode(v))
				elastigroup.Compute.LaunchSpecification.SetShutdownScript(shutdownScript)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var shutdownScript *string = nil
			if v, ok := resourceData.Get(string(ShutdownScript)).(string); ok && v != "" {
				shutdownScript = spotinst.String(base64Encode(v))
			}
			elastigroup.Compute.LaunchSpecification.SetShutdownScript(shutdownScript)
			return nil
		},
		nil,
	)

	fieldsMap[EnableMonitoring] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		EnableMonitoring,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *bool = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Monitoring != nil {
				value = elastigroup.Compute.LaunchSpecification.Monitoring
			}
			if err := resourceData.Set(string(EnableMonitoring), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(EnableMonitoring), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(EnableMonitoring)).(bool); ok {
				elastigroup.Compute.LaunchSpecification.SetMonitoring(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(EnableMonitoring)).(bool); ok {
				elastigroup.Compute.LaunchSpecification.SetMonitoring(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[EbsOptimized] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		EbsOptimized,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *bool = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.EBSOptimized != nil {
				value = elastigroup.Compute.LaunchSpecification.EBSOptimized
			}
			if err := resourceData.Set(string(EbsOptimized), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(EbsOptimized), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(EbsOptimized)).(bool); ok {
				elastigroup.Compute.LaunchSpecification.SetEBSOptimized(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(EbsOptimized)).(bool); ok {
				elastigroup.Compute.LaunchSpecification.SetEBSOptimized(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[PlacementTenancy] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		PlacementTenancy,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Tenancy != nil {
				value = elastigroup.Compute.LaunchSpecification.Tenancy
			}
			if err := resourceData.Set(string(PlacementTenancy), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PlacementTenancy), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(PlacementTenancy)).(string); ok && v != "" {
				tenancy := spotinst.String(v)
				elastigroup.Compute.LaunchSpecification.SetTenancy(tenancy)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var tenancy *string = nil
			if v, ok := resourceData.Get(string(PlacementTenancy)).(string); ok && v != "" {
				tenancy = spotinst.String(v)
			}
			elastigroup.Compute.LaunchSpecification.SetTenancy(tenancy)
			return nil
		},
		nil,
	)

	fieldsMap[CPUCredits] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		CPUCredits,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.CreditSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.CreditSpecification.CPUCredits != nil {
				value = elastigroup.Compute.LaunchSpecification.CreditSpecification.CPUCredits
			}
			if err := resourceData.Set(string(CPUCredits), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CPUCredits), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(CPUCredits)).(string); ok && v != "" {
				if elastigroup.Compute.LaunchSpecification.CreditSpecification == nil {
					elastigroup.Compute.LaunchSpecification.CreditSpecification = &aws.CreditSpecification{}
				}
				credits := spotinst.String(v)
				elastigroup.Compute.LaunchSpecification.CreditSpecification.SetCPUCredits(credits)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			credSpec := &aws.CreditSpecification{}
			if v, ok := resourceData.Get(string(CPUCredits)).(string); ok && v != "" {
				credSpec.SetCPUCredits(spotinst.String(v))
				elastigroup.Compute.LaunchSpecification.SetCreditSpecification(credSpec)
			} else {
				elastigroup.Compute.LaunchSpecification.SetCreditSpecification(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[MetadataOptions] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		MetadataOptions,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(HTTPTokens): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(HTTPPutResponseHopLimit): {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup != nil && elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.MetadataOptions != nil {
				result = flattenMetadataOptions(elastigroup.Compute.LaunchSpecification.MetadataOptions)
			}

			if result != nil {
				if err := resourceData.Set(string(MetadataOptions), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MetadataOptions), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(MetadataOptions)); ok {
				if metaDataOptions, err := expandMetadataOptions(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetMetadataOptions(metaDataOptions)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.MetadataOptions = nil
			if v, ok := resourceData.GetOk(string(MetadataOptions)); ok {
				if metaDataOptions, err := expandMetadataOptions(v); err != nil {
					return err
				} else {
					value = metaDataOptions
				}
			}
			elastigroup.Compute.LaunchSpecification.SetMetadataOptions(value)
			return nil
		},
		nil,
	)

	fieldsMap[CPUOptions] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		CPUOptions,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ThreadsPerCore): {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup != nil && elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.CPUOptions != nil {
				result = flattenCPUOptions(elastigroup.Compute.LaunchSpecification.CPUOptions)
			}

			if result != nil {
				if err := resourceData.Set(string(CPUOptions), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CPUOptions), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(CPUOptions)); ok {
				if cpuOptions, err := expandCPUOptions(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetCPUOptions(cpuOptions)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.CPUOptions = nil
			if v, ok := resourceData.GetOk(string(CPUOptions)); ok {
				if cpuOptions, err := expandCPUOptions(v); err != nil {
					return err
				} else {
					value = cpuOptions
				}
			}
			elastigroup.Compute.LaunchSpecification.SetCPUOptions(value)
			return nil
		},
		nil,
	)

	fieldsMap[ResourceTagSpecification] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		ResourceTagSpecification,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ShouldTagVolumes): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(ShouldTagAMIs): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(ShouldTagENIs): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(ShouldTagSnapshots): {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.ResourceTagSpecification != nil {
				resourceTagSpecification := elastigroup.Compute.LaunchSpecification.ResourceTagSpecification
				result = flattenResourceTagSpecification(resourceTagSpecification)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(ResourceTagSpecification), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ResourceTagSpecification), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(ResourceTagSpecification)); ok {
				if v, err := expandResourceTagSpecification(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetResourceTagSpecification(v)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.ResourceTagSpecification = nil
			if v, ok := resourceData.GetOk(string(ResourceTagSpecification)); ok {
				if resourceTagSpecification, err := expandResourceTagSpecification(v); err != nil {
					return err
				} else {
					value = resourceTagSpecification
				}
			}
			elastigroup.Compute.LaunchSpecification.SetResourceTagSpecification(value)
			return nil
		},
		nil,
	)

	fieldsMap[ITF] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		ITF,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(LoadBalancer): {
						Type:     schema.TypeSet,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(ListenerRule): {
									Type:     schema.TypeSet,
									Required: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(RuleARN): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(StaticTargetGroup): {
												Type:     schema.TypeList,
												MaxItems: 1,
												Optional: true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														string(ARN): {
															Type:     schema.TypeString,
															Required: true,
														},
														string(Percentage): {
															Type:     schema.TypeFloat,
															Required: true,
														},
													},
												},
											},
										},
									},
								},
								string(LoadBalancerARN): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},
					string(MigrationHealthinessThreshold): {
						Type:     schema.TypeInt,
						Optional: true,
					},
					string(FixedTargetGroups): {
						Type:     schema.TypeBool,
						Required: true,
					},
					string(WeightStrategy): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(TargetGroupConfig): {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(HealthCheckIntervalSeconds): {
									Type:     schema.TypeInt,
									Optional: true,
								},
								string(HealthCheckPath): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(HealthCheckPort): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(HealthCheckProtocol): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(HealthCheckTimeoutSeconds): {
									Type:     schema.TypeInt,
									Optional: true,
								},
								string(HealthyThresholdCount): {
									Type:     schema.TypeInt,
									Optional: true,
								},
								string(UnhealthyThresholdCount): {
									Type:     schema.TypeInt,
									Optional: true,
								},
								string(Port): {
									Type:     schema.TypeInt,
									Required: true,
								},
								string(Protocol): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(ProtocolVersion): {
									Type:     schema.TypeString,
									Optional: true,
								},
								string(Matcher): {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(HTTPCode): {
												Type:     schema.TypeString,
												Optional: true,
											},
											string(GRPCCode): {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},
								string(VPCID): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(Tags): {
									Type:     schema.TypeSet,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(TagKey): {
												Type:     schema.TypeString,
												Required: true,
											},
											string(TagValue): {
												Type:     schema.TypeString,
												Optional: true,
											},
										},
									},
								},
							},
						},
					},
					string(DefaultStaticTargetGroup): {
						Type:     schema.TypeList,
						MaxItems: 1,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(ARN): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(Percentage): {
									Type:     schema.TypeFloat,
									Required: true,
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.ITF != nil {
				itf := elastigroup.Compute.LaunchSpecification.ITF
				result = flattenITF(itf)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(ITF), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ITF), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.ITF = nil
			if v, ok := resourceData.GetOk(string(ITF)); ok && v != nil {
				if itf, err := expandITF(v); err != nil {
					return err
				} else {
					value = itf
				}
			}
			elastigroup.Compute.LaunchSpecification.SetITF(value)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.ITF = nil
			if v, ok := resourceData.GetOk(string(ITF)); ok && v != nil {
				if itf, err := expandITF(v); err != nil {
					return err
				} else {
					value = itf
				}
			}
			elastigroup.Compute.LaunchSpecification.SetITF(value)
			return nil
		},
		nil,
	)
}

var InstanceProfileArnRegex = regexp.MustCompile(`arn:aws:iam::\d{12}:instance-profile/?[a-zA-Z_0-9+=,.@\-_/]+`)

func Base64StateFunc(v interface{}) string {
	if isBase64Encoded(v.(string)) {
		return v.(string)
	} else {
		return base64Encode(v.(string))
	}
}

// base64Encode encodes data if the input isn't already encoded using
// base64.StdEncoding.EncodeToString. If the input is already base64 encoded,
// return the original input unchanged.
func base64Encode(data string) string {
	// Check whether the data is already Base64 encoded; don't double-encode
	if isBase64Encoded(data) {
		return data
	}
	// data has not been encoded encode and return
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func isBase64Encoded(data string) bool {
	_, err := base64.StdEncoding.DecodeString(data)
	return err == nil
}

func expandITF(data interface{}) (*aws.ITF, error) {
	itf := &aws.ITF{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return itf, nil
	}
	m := list[0].(map[string]interface{})
	if v, ok := m[string(FixedTargetGroups)].(bool); ok {
		itf.SetFixedTargetGroups(spotinst.Bool(v))
	}
	if v, ok := m[string(WeightStrategy)].(string); ok && v != "" {
		itf.SetWeightStrategy(spotinst.String(v))
	}
	if v, ok := m[string(MigrationHealthinessThreshold)].(int); ok && v >= 0 {
		itf.SetMigrationHealthinessThreshold(spotinst.Int(v))
	} else {
		itf.SetMigrationHealthinessThreshold(nil)
	}
	if v, ok := m[string(LoadBalancer)]; ok {

		loadBalancers, err := expandLoadBalancers(v)
		if err != nil {
			return nil, err
		}
		if loadBalancers != nil {
			itf.SetLoadBalancers(loadBalancers)
		}
	} else {
		itf.LoadBalancers = nil
	}
	if v, ok := m[string(TargetGroupConfig)]; ok {

		targetGroupConfig, err := expandTargetGroupConfig(v)
		if err != nil {
			return nil, err
		}
		if targetGroupConfig != nil {
			itf.SetTargetGroupConfig(targetGroupConfig)
		}
	} else {
		itf.TargetGroupConfig = nil
	}
	if v, ok := m[string(DefaultStaticTargetGroup)]; ok {

		staticTargetGroups, err := expandDefaultStaticTargetGroups(v)
		if err != nil {
			return nil, err
		}
		if staticTargetGroups != nil {
			itf.SetDefaultStaticTargetGroups(staticTargetGroups)
		}
	} else {
		itf.DefaultStaticTargetGroups = nil
	}

	return itf, nil
}

func flattenITF(itf *aws.ITF) []interface{} {
	result := make(map[string]interface{})

	result[string(WeightStrategy)] = spotinst.StringValue(itf.WeightStrategy)
	result[string(MigrationHealthinessThreshold)] = spotinst.IntValue(itf.MigrationHealthinessThreshold)
	result[string(FixedTargetGroups)] = spotinst.BoolValue(itf.FixedTargetGroups)
	if itf.LoadBalancers != nil {
		result[string(LoadBalancer)] = flattenLoadBalancers(itf.LoadBalancers)
	}
	if itf.TargetGroupConfig != nil {
		result[string(TargetGroupConfig)] = flattenTargetGroupConfig(itf.TargetGroupConfig)
	}
	if itf.DefaultStaticTargetGroups != nil {
		result[string(DefaultStaticTargetGroup)] = flattenDefaultStaticTargetGroups(itf.DefaultStaticTargetGroups)
	}

	return []interface{}{result}
}

func flattenTargetGroupConfig(targetGroupConfig *aws.TargetGroupConfig) []interface{} {
	result := make(map[string]interface{})

	result[string(VPCID)] = spotinst.StringValue(targetGroupConfig.VPCID)
	result[string(HealthCheckIntervalSeconds)] = spotinst.IntValue(targetGroupConfig.HealthCheckIntervalSeconds)
	result[string(HealthCheckPath)] = spotinst.StringValue(targetGroupConfig.HealthCheckPath)
	result[string(HealthCheckPort)] = spotinst.StringValue(targetGroupConfig.HealthCheckPort)
	result[string(HealthCheckProtocol)] = spotinst.StringValue(targetGroupConfig.HealthCheckProtocol)
	result[string(HealthCheckTimeoutSeconds)] = spotinst.IntValue(targetGroupConfig.HealthCheckTimeoutSeconds)
	result[string(HealthyThresholdCount)] = spotinst.IntValue(targetGroupConfig.HealthyThresholdCount)
	result[string(UnhealthyThresholdCount)] = spotinst.IntValue(targetGroupConfig.UnhealthyThresholdCount)
	result[string(Port)] = spotinst.IntValue(targetGroupConfig.Port)
	result[string(Protocol)] = spotinst.StringValue(targetGroupConfig.Protocol)
	result[string(ProtocolVersion)] = spotinst.StringValue(targetGroupConfig.ProtocolVersion)
	if targetGroupConfig.Matcher != nil {
		result[string(Matcher)] = flattenMatcher(targetGroupConfig.Matcher)
	}
	if targetGroupConfig.Tags != nil {
		result[string(Tags)] = flattenTags(targetGroupConfig.Tags)
	}

	return []interface{}{result}
}

func flattenLoadBalancers(loadBalancers []*aws.ITFLoadBalancer) []interface{} {
	result := make([]interface{}, 0, len(loadBalancers))

	for _, loadBalancer := range loadBalancers {
		m := make(map[string]interface{})
		m[string(LoadBalancerARN)] = spotinst.StringValue(loadBalancer.LoadBalancerARN)
		if loadBalancer.ListenerRules != nil {
			m[string(ListenerRule)] = flattenListenerRules(loadBalancer.ListenerRules)
		}
		result = append(result, m)
	}

	return result
}

func flattenDefaultStaticTargetGroups(StaticTargetGroups []*aws.StaticTargetGroup) []interface{} {
	result := make([]interface{}, 0, len(StaticTargetGroups))

	for _, staticTargetGroup := range StaticTargetGroups {
		m := make(map[string]interface{})
		m[string(ARN)] = spotinst.StringValue(staticTargetGroup.StaticTargetGroupARN)
		m[string(Percentage)] = spotinst.Float64Value(staticTargetGroup.Percentage)
		result = append(result, m)
	}

	return result
}

func flattenListenerRules(listenerRules []*aws.ListenerRule) []interface{} {
	result := make([]interface{}, 0, len(listenerRules))

	for _, listenerRule := range listenerRules {
		m := make(map[string]interface{})
		m[string(RuleARN)] = spotinst.StringValue(listenerRule.RuleARN)
		if listenerRule.StaticTargetGroups != nil {
			m[string(StaticTargetGroup)] = flattenStaticTargetGroups(listenerRule.StaticTargetGroups)
		}
		result = append(result, m)
	}

	return result
}

func flattenStaticTargetGroups(StaticTargetGroups []*aws.StaticTargetGroup) []interface{} {
	result := make([]interface{}, 0, len(StaticTargetGroups))

	for _, staticTargetGroup := range StaticTargetGroups {
		m := make(map[string]interface{})
		m[string(ARN)] = spotinst.StringValue(staticTargetGroup.StaticTargetGroupARN)
		m[string(Percentage)] = spotinst.Float64Value(staticTargetGroup.Percentage)
		result = append(result, m)
	}

	return result
}

func flattenMatcher(matcher *aws.Matcher) []interface{} {
	result := make(map[string]interface{})
	result[string(HTTPCode)] = spotinst.StringValue(matcher.HTTPCode)
	result[string(GRPCCode)] = spotinst.StringValue(matcher.GRPCCode)

	return []interface{}{result}
}

func flattenTags(tags []*aws.Tag) []interface{} {
	result := make([]interface{}, 0, len(tags))
	for _, tag := range tags {
		m := make(map[string]interface{})
		m[string(TagKey)] = spotinst.StringValue(tag.Key)
		m[string(TagValue)] = spotinst.StringValue(tag.Value)

		result = append(result, m)
	}
	return result
}

func expandDefaultStaticTargetGroups(data interface{}) ([]*aws.StaticTargetGroup, error) {
	list := data.([]interface{})
	staticTargetGroups := make([]*aws.StaticTargetGroup, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})

		if !ok {
			continue
		}

		staticTargetGroup := &aws.StaticTargetGroup{}

		if v, ok := attr[string(ARN)].(string); ok && v != "" {
			staticTargetGroup.SetStaticTargetGroupARN(spotinst.String(v))
		}
		if v, ok := attr[string(Percentage)].(float64); ok {
			staticTargetGroup.SetPercentage(spotinst.Float64(v))
		}

		staticTargetGroups = append(staticTargetGroups, staticTargetGroup)
	}
	return staticTargetGroups, nil
}

func expandTargetGroupConfig(data interface{}) (*aws.TargetGroupConfig, error) {
	targetGroupConfig := &aws.TargetGroupConfig{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return targetGroupConfig, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(VPCID)].(string); ok && v != "" {
		targetGroupConfig.SetVPCId(spotinst.String(v))
	}
	if v, ok := m[string(HealthCheckIntervalSeconds)].(int); ok && v >= 0 {
		targetGroupConfig.SetHealthCheckIntervalSeconds(spotinst.Int(v))
	} else {
		targetGroupConfig.SetHealthCheckIntervalSeconds(nil)
	}
	if v, ok := m[string(HealthCheckPath)].(string); ok && v != "" {
		targetGroupConfig.SetHealthCheckPath(spotinst.String(v))
	}
	if v, ok := m[string(HealthCheckPort)].(string); ok && v != "" {
		targetGroupConfig.SetHealthCheckPort(spotinst.String(v))
	} else {
		targetGroupConfig.SetHealthCheckPort(nil)
	}
	if v, ok := m[string(HealthCheckProtocol)].(string); ok && v != "" {
		targetGroupConfig.SetHealthCheckProtocol(spotinst.String(v))
	}
	if v, ok := m[string(HealthCheckTimeoutSeconds)].(int); ok && v >= 0 {
		targetGroupConfig.SetHealthCheckTimeoutSeconds(spotinst.Int(v))
	} else {
		targetGroupConfig.SetHealthCheckTimeoutSeconds(nil)
	}
	if v, ok := m[string(HealthyThresholdCount)].(int); ok && v >= 0 {
		targetGroupConfig.SetHealthyThresholdCount(spotinst.Int(v))
	} else {
		targetGroupConfig.SetHealthyThresholdCount(nil)
	}
	if v, ok := m[string(UnhealthyThresholdCount)].(int); ok && v >= 0 {
		targetGroupConfig.SetUnhealthyThresholdCount(spotinst.Int(v))
	} else {
		targetGroupConfig.SetUnhealthyThresholdCount(nil)
	}
	if v, ok := m[string(Port)].(int); ok && v >= 0 {
		targetGroupConfig.SetPort(spotinst.Int(v))
	} else {
		targetGroupConfig.SetPort(nil)
	}
	if v, ok := m[string(Protocol)].(string); ok && v != "" {
		targetGroupConfig.SetProtocol(spotinst.String(v))
	}
	if v, ok := m[string(ProtocolVersion)].(string); ok && v != "" {
		targetGroupConfig.SetProtocolVersion(spotinst.String(v))
	}
	if v, ok := m[string(Matcher)]; ok && v != nil {
		matcher, err := expandMatcher(v)
		if err != nil {
			return nil, err
		}
		if matcher != nil {
			targetGroupConfig.SetMatcher(matcher)
		} else {
			targetGroupConfig.Matcher = nil
		}
	}
	if v, ok := m[string(Tags)]; ok {
		tags, err := expandTags(v)
		if err != nil {
			return nil, err
		}
		if tags != nil {
			targetGroupConfig.SetTags(tags)
		} else {
			targetGroupConfig.SetTags(nil)
		}
	}

	return targetGroupConfig, nil
}

func expandTags(data interface{}) ([]*aws.Tag, error) {
	list := data.(*schema.Set).List()
	tags := make([]*aws.Tag, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(TagKey)]; !ok {
			return nil, errors.New("invalid tag attributes: key missing")
		}

		if _, ok := attr[string(TagValue)]; !ok {
			return nil, errors.New("invalid tag attributes: value missing")
		}
		tag := &aws.Tag{
			Key:   spotinst.String(attr[string(TagKey)].(string)),
			Value: spotinst.String(attr[string(TagValue)].(string)),
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func expandMatcher(data interface{}) (*aws.Matcher, error) {
	if list := data.([]interface{}); len(list) > 0 {
		matcher := &aws.Matcher{}
		if list != nil && list[0] != nil {

			m := list[0].(map[string]interface{})
			if v, ok := m[string(HTTPCode)].(string); ok && v != "" {
				matcher.SetHTTPCode(spotinst.String(v))
			} else {
				matcher.HTTPCode = nil
			}
			if v, ok := m[string(GRPCCode)].(string); ok && v != "" {
				matcher.SetGRPCCode(spotinst.String(v))
			} else {
				matcher.GRPCCode = nil
			}
		}
		return matcher, nil
	}

	return nil, nil
}

func expandLoadBalancers(data interface{}) ([]*aws.ITFLoadBalancer, error) {
	list := data.(*schema.Set).List()
	loadBalancers := make([]*aws.ITFLoadBalancer, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})

		if !ok {
			continue
		}

		loadBalancer := &aws.ITFLoadBalancer{}

		if v, ok := attr[string(LoadBalancerARN)].(string); ok && v != "" {
			loadBalancer.SetLoadBalancerARN(spotinst.String(v))
		}
		if v, ok := attr[string(ListenerRule)]; ok {
			listenerRules, err := expandListenerRules(v)
			if err != nil {
				return nil, err
			}
			if listenerRules != nil {
				loadBalancer.SetListenerRules(listenerRules)
			} else {
				loadBalancer.SetListenerRules(nil)
			}
		}

		loadBalancers = append(loadBalancers, loadBalancer)
	}

	return loadBalancers, nil
}

func expandListenerRules(data interface{}) ([]*aws.ListenerRule, error) {
	list := data.(*schema.Set).List()
	listenerRules := make([]*aws.ListenerRule, 0, len(list))

	for _, item := range list {
		attr, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		listenerRule := &aws.ListenerRule{}
		if v, ok := attr[string(RuleARN)].(string); ok && v != "" {
			listenerRule.SetRuleARN(spotinst.String(v))
		}
		if v, ok := attr[string(StaticTargetGroup)]; ok {

			staticTargetGroups, err := expandStaticTargetGroups(v)
			if err != nil {
				return nil, err
			}
			if staticTargetGroups != nil {
				listenerRule.SetStaticTargetGroups(staticTargetGroups)
			}
		} else {
			listenerRule.StaticTargetGroups = nil
		}
		listenerRules = append(listenerRules, listenerRule)
	}

	return listenerRules, nil
}

func expandStaticTargetGroups(data interface{}) ([]*aws.StaticTargetGroup, error) {
	list := data.([]interface{})
	staticTargetGroups := make([]*aws.StaticTargetGroup, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})

		if !ok {
			continue
		}

		staticTargetGroup := &aws.StaticTargetGroup{}

		if v, ok := attr[string(ARN)].(string); ok && v != "" {
			staticTargetGroup.SetStaticTargetGroupARN(spotinst.String(v))
		}
		if v, ok := attr[string(Percentage)].(float64); ok {
			staticTargetGroup.SetPercentage(spotinst.Float64(v))
		}

		staticTargetGroups = append(staticTargetGroups, staticTargetGroup)
	}
	return staticTargetGroups, nil
}

func expandMetadataOptions(data interface{}) (*aws.MetadataOptions, error) {
	metadataOptions := &aws.MetadataOptions{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return metadataOptions, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(HTTPTokens)].(string); ok && v != "" {
		metadataOptions.SetHTTPTokens(spotinst.String(v))
	}
	if v, ok := m[string(HTTPPutResponseHopLimit)].(int); ok && v >= 0 {
		metadataOptions.SetHTTPPutResponseHopLimit(spotinst.Int(v))
	} else {
		metadataOptions.SetHTTPPutResponseHopLimit(nil)
	}

	return metadataOptions, nil
}

func flattenMetadataOptions(metadataOptions *aws.MetadataOptions) []interface{} {
	result := make(map[string]interface{})
	result[string(HTTPTokens)] = spotinst.StringValue(metadataOptions.HTTPTokens)
	result[string(HTTPPutResponseHopLimit)] = spotinst.IntValue(metadataOptions.HTTPPutResponseHopLimit)

	return []interface{}{result}
}

func expandCPUOptions(data interface{}) (*aws.CPUOptions, error) {
	cpuOptions := &aws.CPUOptions{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return cpuOptions, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(ThreadsPerCore)].(int); ok && v >= 0 {
		cpuOptions.SetThreadsPerCore(spotinst.Int(v))
	} else {
		cpuOptions.SetThreadsPerCore(nil)
	}

	return cpuOptions, nil
}

func flattenCPUOptions(cpuOptions *aws.CPUOptions) []interface{} {
	result := make(map[string]interface{})
	result[string(ThreadsPerCore)] = spotinst.IntValue(cpuOptions.ThreadsPerCore)

	return []interface{}{result}
}

func expandResourceTagSpecification(data interface{}) (*aws.ResourceTagSpecification, error) {
	resourceTagSpecification := &aws.ResourceTagSpecification{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return resourceTagSpecification, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(ShouldTagVolumes)].(bool); ok {
		volumes := &aws.Volumes{}
		resourceTagSpecification.SetVolumes(volumes)
		resourceTagSpecification.Volumes.SetShouldTag(spotinst.Bool(v))

	}
	if v, ok := m[string(ShouldTagAMIs)].(bool); ok {
		anis := &aws.AMIs{}
		resourceTagSpecification.SetAMIs(anis)
		resourceTagSpecification.AMIs.SetShouldTag(spotinst.Bool(v))

	}
	if v, ok := m[string(ShouldTagENIs)].(bool); ok {
		enis := &aws.ENIs{}
		resourceTagSpecification.SetENIs(enis)
		resourceTagSpecification.ENIs.SetShouldTag(spotinst.Bool(v))
	}
	if v, ok := m[string(ShouldTagSnapshots)].(bool); ok {
		snapshots := &aws.Snapshots{}
		resourceTagSpecification.SetSnapshots(snapshots)
		resourceTagSpecification.Snapshots.SetShouldTag(spotinst.Bool(v))

	}

	return resourceTagSpecification, nil
}

func flattenResourceTagSpecification(resourceTagSpecification *aws.ResourceTagSpecification) []interface{} {
	result := make(map[string]interface{})
	if resourceTagSpecification != nil && resourceTagSpecification.Snapshots != nil {
		result[string(ShouldTagSnapshots)] = spotinst.BoolValue(resourceTagSpecification.Snapshots.ShouldTag)
	}
	if resourceTagSpecification != nil && resourceTagSpecification.ENIs != nil {
		result[string(ShouldTagENIs)] = spotinst.BoolValue(resourceTagSpecification.ENIs.ShouldTag)
	}
	if resourceTagSpecification != nil && resourceTagSpecification.AMIs != nil {
		result[string(ShouldTagAMIs)] = spotinst.BoolValue(resourceTagSpecification.AMIs.ShouldTag)
	}
	if resourceTagSpecification != nil && resourceTagSpecification.Volumes != nil {
		result[string(ShouldTagVolumes)] = spotinst.BoolValue(resourceTagSpecification.Volumes.ShouldTag)
	}

	return []interface{}{result}
}

func expandImages(data interface{}) (*aws.Images, error) {
	Images := &aws.Images{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return Images, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Id)].(int); ok && v >= 0 {
		Images.SetImageId(spotinst.Int(v))
	} else {
		Images.SetImageId(nil)
	}

	return cpuOptions, nil
}

func flattenImages(Images *aws.Images) []interface{} {
	result := make(map[string]interface{})
	result[string(ImageId)] = spotinst.IntValue(Images.ImageId)

	return []interface{}{result}
}

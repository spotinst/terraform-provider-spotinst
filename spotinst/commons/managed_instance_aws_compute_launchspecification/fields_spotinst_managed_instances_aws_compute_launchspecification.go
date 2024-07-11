package managed_instance_aws_compute_launchspecification

import (
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[EBSOptimized] = commons.NewGenericField(
		commons.ManagedInstanceAWSLaunchSpecification,
		EBSOptimized,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *bool = nil
			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.EBSOptimized != nil {
				value = managedInstance.Compute.LaunchSpecification.EBSOptimized
			}
			if err := resourceData.Set(string(EBSOptimized), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(EBSOptimized), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(EBSOptimized)).(bool); ok {
				managedInstance.Compute.LaunchSpecification.SetEBSOptimized(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(EBSOptimized)).(bool); ok {
				managedInstance.Compute.LaunchSpecification.SetEBSOptimized(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[EnableMonitoring] = commons.NewGenericField(
		commons.ManagedInstanceAWSLaunchSpecification,
		EnableMonitoring,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *bool = nil
			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.Monitoring != nil {
				value = managedInstance.Compute.LaunchSpecification.Monitoring
			}
			if err := resourceData.Set(string(EnableMonitoring), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(EnableMonitoring), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(EnableMonitoring)).(bool); ok {
				managedInstance.Compute.LaunchSpecification.SetMonitoring(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(EnableMonitoring)).(bool); ok {
				managedInstance.Compute.LaunchSpecification.SetMonitoring(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[PlacementTenancy] = commons.NewGenericField(
		commons.ManagedInstanceAWSLaunchSpecification,
		PlacementTenancy,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *string = nil
			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.Tenancy != nil {
				value = managedInstance.Compute.LaunchSpecification.Tenancy
			}
			if err := resourceData.Set(string(PlacementTenancy), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PlacementTenancy), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(PlacementTenancy)).(string); ok && v != "" {
				tenancy := spotinst.String(v)
				managedInstance.Compute.LaunchSpecification.SetTenancy(tenancy)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var tenancy *string = nil
			if v, ok := resourceData.Get(string(PlacementTenancy)).(string); ok && v != "" {
				tenancy = spotinst.String(v)
			}
			managedInstance.Compute.LaunchSpecification.SetTenancy(tenancy)
			return nil
		},
		nil,
	)
	fieldsMap[SecurityGroupIDs] = commons.NewGenericField(
		commons.ManagedInstanceAWSLaunchSpecification,
		SecurityGroupIDs,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value []string = nil
			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.SecurityGroupIDs != nil {
				value = managedInstance.Compute.LaunchSpecification.SecurityGroupIDs
			}
			if err := resourceData.Set(string(SecurityGroupIDs), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SecurityGroupIDs), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(SecurityGroupIDs)).([]interface{}); ok {
				ids := make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
				managedInstance.Compute.LaunchSpecification.SetSecurityGroupIDs(ids)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(SecurityGroupIDs)).([]interface{}); ok {
				ids := make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
				managedInstance.Compute.LaunchSpecification.SetSecurityGroupIDs(ids)
			}
			return nil
		},
		nil,
	)

	fieldsMap[ImageID] = commons.NewGenericField(
		commons.ManagedInstanceAWSLaunchSpecification,
		ImageID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *string = nil
			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.ImageID != nil {
				value = managedInstance.Compute.LaunchSpecification.ImageID
			}
			if err := resourceData.Set(string(ImageID), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ImageID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(ImageID)).(string); ok && v != "" {
				managedInstance.Compute.LaunchSpecification.SetImageId(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(ImageID)).(string); ok && v != "" {
				managedInstance.Compute.LaunchSpecification.SetImageId(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[KeyPair] = commons.NewGenericField(
		commons.ManagedInstanceAWSLaunchSpecification,
		KeyPair,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *string = nil
			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.KeyPair != nil {
				value = managedInstance.Compute.LaunchSpecification.KeyPair
			}
			if err := resourceData.Set(string(KeyPair), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(KeyPair), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(KeyPair)).(string); ok && v != "" {
				managedInstance.Compute.LaunchSpecification.SetKeyPair(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(KeyPair)).(string); ok && v != "" {
				managedInstance.Compute.LaunchSpecification.SetKeyPair(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[UserData] = commons.NewGenericField(
		commons.ManagedInstanceAWSLaunchSpecification,
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
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value = ""
			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.UserData != nil {

				userData := managedInstance.Compute.LaunchSpecification.UserData
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
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData := spotinst.String(base64Encode(v))
				managedInstance.Compute.LaunchSpecification.SetUserData(userData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var userData *string = nil
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData = spotinst.String(base64Encode(v))
			}
			managedInstance.Compute.LaunchSpecification.SetUserData(userData)
			return nil
		},
		nil,
	)

	fieldsMap[ShutdownScript] = commons.NewGenericField(
		commons.ManagedInstanceAWSLaunchSpecification,
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
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value = ""
			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.ShutdownScript != nil {

				shutdownScript := managedInstance.Compute.LaunchSpecification.ShutdownScript
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
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(ShutdownScript)).(string); ok && v != "" {
				shutdownScript := spotinst.String(base64Encode(v))
				managedInstance.Compute.LaunchSpecification.SetShutdownScript(shutdownScript)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var shutdownScript *string = nil
			if v, ok := resourceData.Get(string(ShutdownScript)).(string); ok && v != "" {
				shutdownScript = spotinst.String(base64Encode(v))
			}
			managedInstance.Compute.LaunchSpecification.SetShutdownScript(shutdownScript)
			return nil
		},
		nil,
	)

	fieldsMap[IAMInstanceProfile] = commons.NewGenericField(
		commons.ManagedInstanceAWSLaunchSpecification,
		IAMInstanceProfile,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value = ""
			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.IAMInstanceProfile != nil {

				lc := managedInstance.Compute.LaunchSpecification
				if lc.IAMInstanceProfile.Arn != nil {
					value = spotinst.StringValue(lc.IAMInstanceProfile.Arn)
				} else if lc.IAMInstanceProfile.Name != nil {
					value = spotinst.StringValue(lc.IAMInstanceProfile.Name)
				}
			}
			if err := resourceData.Set(string(IAMInstanceProfile), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IAMInstanceProfile), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(IAMInstanceProfile)).(string); ok && v != "" {
				iamInstanceProf := &aws.IAMInstanceProfile{}
				if InstanceProfileArnRegex.MatchString(v) {
					iamInstanceProf.SetArn(spotinst.String(v))
				} else {
					iamInstanceProf.SetName(spotinst.String(v))
				}
				managedInstance.Compute.LaunchSpecification.SetIAMInstanceProfile(iamInstanceProf)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(IAMInstanceProfile)).(string); ok && v != "" {
				iamInstanceProf := &aws.IAMInstanceProfile{}
				if InstanceProfileArnRegex.MatchString(v) {
					iamInstanceProf.SetArn(spotinst.String(v))
				} else {
					iamInstanceProf.SetName(spotinst.String(v))
				}
				managedInstance.Compute.LaunchSpecification.SetIAMInstanceProfile(iamInstanceProf)
			} else {
				managedInstance.Compute.LaunchSpecification.SetIAMInstanceProfile(nil)
			}
			return nil
		},
		nil,
	)
	fieldsMap[CPUCredits] = commons.NewGenericField(
		commons.ManagedInstanceAWSLaunchSpecification,
		CPUCredits,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *string = nil
			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.CreditSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.CreditSpecification.CPUCredits != nil {
				value = managedInstance.Compute.LaunchSpecification.CreditSpecification.CPUCredits
			}
			if err := resourceData.Set(string(CPUCredits), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CPUCredits), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(CPUCredits)).(string); ok && v != "" {
				if managedInstance.Compute.LaunchSpecification.CreditSpecification == nil {
					managedInstance.Compute.LaunchSpecification.CreditSpecification = &aws.CreditSpecification{}
				}
				credits := spotinst.String(v)
				managedInstance.Compute.LaunchSpecification.CreditSpecification.SetCPUCredits(credits)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(CPUCredits)).(string); ok && v != "" {
				if managedInstance.Compute.LaunchSpecification.CreditSpecification == nil {
					managedInstance.Compute.LaunchSpecification.CreditSpecification = &aws.CreditSpecification{}
				}
				credits := spotinst.String(v)
				managedInstance.Compute.LaunchSpecification.CreditSpecification.SetCPUCredits(credits)
			}
			return nil
		},
		nil,
	)

	fieldsMap[Tags] = commons.NewGenericField(
		commons.ManagedInstanceAWSLaunchSpecification,
		Tags,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(TagKey): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(TagValue): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var result []interface{} = nil
			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.Tags != nil {
				tags := managedInstance.Compute.LaunchSpecification.Tags
				result = flattenTags(tags)
			}
			if result != nil {
				if err := resourceData.Set(string(Tags), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tags), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					managedInstance.Compute.LaunchSpecification.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var tagsToAdd []*aws.Tag = nil
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					tagsToAdd = tags
				}
			}
			managedInstance.Compute.LaunchSpecification.SetTags(tagsToAdd)
			return nil
		},
		nil,
	)

	fieldsMap[BlockDeviceMappings] = commons.NewGenericField(
		commons.ManagedInstanceAWSLaunchSpecification,
		BlockDeviceMappings,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(DeviceName): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(EBS): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(DeleteOnTermination): {
									Type:     schema.TypeBool,
									Optional: true,
									Computed: true,
								},

								string(IOPS): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(VolumeSize): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(VolumeType): {
									Type:     schema.TypeString,
									Optional: true,
									Computed: true,
								},

								string(Throughput): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(Encrypted): {
									Type:     schema.TypeBool,
									Optional: true,
								},

								string(KmsKeyId): {
									Type:     schema.TypeString,
									Optional: true,
								},

								string(SnapshotId): {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var result []interface{} = nil
			if managedInstance != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.BlockDeviceMappings != nil {
				blockDeviceMappings := managedInstance.Compute.LaunchSpecification.BlockDeviceMappings
				result = flattenBlockDeviceMappings(blockDeviceMappings)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(BlockDeviceMappings), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BlockDeviceMappings), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.GetOk(string(BlockDeviceMappings)); ok {
				if v, err := expandBlockDeviceMappings(v); err != nil {
					return err
				} else {
					managedInstance.Compute.LaunchSpecification.SetBlockDeviceMappings(v)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value []*aws.BlockDeviceMapping = nil
			if v, ok := resourceData.GetOk(string(BlockDeviceMappings)); ok {
				if blockDeviceMappings, err := expandBlockDeviceMappings(v); err != nil {
					return err
				} else {
					value = blockDeviceMappings
				}
			}
			managedInstance.Compute.LaunchSpecification.SetBlockDeviceMappings(value)
			return nil
		},
		nil,
	)

	fieldsMap[ResourceTagSpecification] = commons.NewGenericField(
		commons.ManagedInstanceAWSLaunchSpecification,
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
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var result []interface{} = nil
			if managedInstance != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.ResourceTagSpecification != nil {
				resourceTagSpecification := managedInstance.Compute.LaunchSpecification.ResourceTagSpecification
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
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.GetOk(string(ResourceTagSpecification)); ok {
				if v, err := expandResourceTagSpecification(v); err != nil {
					return err
				} else {
					managedInstance.Compute.LaunchSpecification.SetResourceTagSpecification(v)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *aws.ResourceTagSpecification = nil
			if v, ok := resourceData.GetOk(string(ResourceTagSpecification)); ok {
				if resourceTagSpecification, err := expandResourceTagSpecification(v); err != nil {
					return err
				} else {
					value = resourceTagSpecification
				}
			}
			managedInstance.Compute.LaunchSpecification.SetResourceTagSpecification(value)
			return nil
		},
		nil,
	)

	fieldsMap[MetadataOptions] = commons.NewGenericField(
		commons.ManagedInstanceAWSLaunchSpecification,
		MetadataOptions,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(HttpTokens): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(HttpPutResponseHopLimit): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(InstanceMetadataTags): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var result []interface{} = nil
			if managedInstance != nil && managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.MetadataOptions != nil {
				result = flattenMetadataOptions(managedInstance.Compute.LaunchSpecification.MetadataOptions)
			}

			if result != nil {
				if err := resourceData.Set(string(MetadataOptions), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MetadataOptions), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.GetOk(string(MetadataOptions)); ok {
				if metaDataOptions, err := expandMetadataOptions(v); err != nil {
					return err
				} else {
					managedInstance.Compute.LaunchSpecification.SetMetadataOptions(metaDataOptions)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *aws.MetadataOptions = nil
			if v, ok := resourceData.GetOk(string(MetadataOptions)); ok {
				if metaDataOptions, err := expandMetadataOptions(v); err != nil {
					return err
				} else {
					value = metaDataOptions
				}
			}
			managedInstance.Compute.LaunchSpecification.SetMetadataOptions(value)
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

func expandBlockDeviceMappings(data interface{}) ([]*aws.BlockDeviceMapping, error) {
	list := data.([]interface{})
	bdms := make([]*aws.BlockDeviceMapping, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		bdm := &aws.BlockDeviceMapping{}

		if v, ok := attr[string(DeviceName)].(string); ok && v != "" {
			bdm.SetDeviceName(spotinst.String(v))
		}

		if r, ok := attr[string(EBS)]; ok {
			if ebs, err := expandEBS(r); err != nil {
				return nil, err
			} else {
				bdm.SetEBS(ebs)
			}
		}

		bdms = append(bdms, bdm)
	}

	return bdms, nil
}

func expandEBS(data interface{}) (*aws.EBS, error) {
	ebs := &aws.EBS{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return ebs, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(DeleteOnTermination)].(bool); ok {
		ebs.SetDeleteOnTermination(spotinst.Bool(v))
	}

	if v, ok := m[string(IOPS)].(int); ok && v > 0 {
		ebs.SetIOPS(spotinst.Int(v))
	}

	if v, ok := m[string(VolumeSize)].(int); ok && v > 0 {
		ebs.SetVolumeSize(spotinst.Int(v))
	}

	if v, ok := m[string(VolumeType)].(string); ok && v != "" {
		ebs.SetVolumeType(spotinst.String(v))
	}

	if v, ok := m[string(Throughput)].(int); ok && v > 0 {
		ebs.SetThroughput(spotinst.Int(v))
	}

	var encrypted = spotinst.Bool(false)
	if v, ok := m[string(Encrypted)].(bool); ok {
		encrypted = spotinst.Bool(v)
	}
	ebs.SetEncrypted(encrypted)

	if v, ok := m[string(KmsKeyId)].(string); ok && v != "" {
		ebs.SetKmsKeyId(spotinst.String(v))
	}

	if v, ok := m[string(SnapshotId)].(string); ok && v != "" {
		ebs.SetSnapshotId(spotinst.String(v))
	}

	return ebs, nil
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

func flattenBlockDeviceMappings(bdms []*aws.BlockDeviceMapping) []interface{} {
	result := make([]interface{}, 0, len(bdms))
	for _, bdm := range bdms {
		m := make(map[string]interface{})
		m[string(DeviceName)] = spotinst.StringValue(bdm.DeviceName)
		if bdm.EBS != nil {
			m[string(EBS)] = flattenEBS(bdm.EBS)
		}
		result = append(result, m)
	}
	return result

}

func flattenEBS(ebs *aws.EBS) []interface{} {
	e := make(map[string]interface{})
	e[string(DeleteOnTermination)] = spotinst.BoolValue(ebs.DeleteOnTermination)
	e[string(IOPS)] = spotinst.IntValue(ebs.IOPS)
	e[string(VolumeType)] = spotinst.StringValue(ebs.VolumeType)
	e[string(VolumeSize)] = spotinst.IntValue(ebs.VolumeSize)
	e[string(Throughput)] = spotinst.IntValue(ebs.Throughput)
	e[string(Encrypted)] = spotinst.BoolValue(ebs.Encrypted)
	e[string(KmsKeyId)] = spotinst.StringValue(ebs.KmsKeyId)
	e[string(SnapshotId)] = spotinst.StringValue(ebs.SnapshotID)
	return []interface{}{e}
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

func flattenMetadataOptions(metadataOptions *aws.MetadataOptions) []interface{} {
	result := make(map[string]interface{})
	result[string(HttpTokens)] = spotinst.StringValue(metadataOptions.HttpTokens)
	result[string(HttpPutResponseHopLimit)] = spotinst.IntValue(metadataOptions.HttpPutResponseHopLimit)
	result[string(InstanceMetadataTags)] = spotinst.StringValue(metadataOptions.InstanceMetadataTags)

	return []interface{}{result}
}

func expandMetadataOptions(data interface{}) (*aws.MetadataOptions, error) {
	metadataOptions := &aws.MetadataOptions{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return metadataOptions, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(HttpTokens)].(string); ok && v != "" {
		metadataOptions.SetHttpTokens(spotinst.String(v))
	}
	if v, ok := m[string(HttpPutResponseHopLimit)].(int); ok && v > 0 {
		metadataOptions.SetHttpPutResponseHopLimit(spotinst.Int(v))
	} else {
		metadataOptions.SetHttpPutResponseHopLimit(nil)
	}
	if v, ok := m[string(InstanceMetadataTags)].(string); ok && v != "" {
		metadataOptions.SetInstanceMetadataTags(spotinst.String(v))
	}

	return metadataOptions, nil
}

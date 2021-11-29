package ocean_ecs_launch_spec

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[OceanID] = commons.NewGenericField(
		commons.OceanECSLaunchSpec,
		OceanID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var value *string = nil
			if launchSpec.OceanID != nil {
				value = launchSpec.OceanID
			}
			if err := resourceData.Set(string(OceanID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OceanID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			launchSpec.SetOceanId(spotinst.String(resourceData.Get(string(OceanID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			launchSpec.SetOceanId(spotinst.String(resourceData.Get(string(OceanID)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanECSLaunchSpec,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var value *string = nil
			if launchSpec.Name != nil {
				value = launchSpec.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			launchSpec.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			launchSpec.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[ImageID] = commons.NewGenericField(
		commons.OceanECSLaunchSpec,
		ImageID,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var value *string = nil
			if launchSpec.ImageID != nil {
				value = launchSpec.ImageID
			}
			if err := resourceData.Set(string(ImageID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ImageID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(ImageID)).(string); ok && v != "" {
				launchSpec.SetImageId(spotinst.String(resourceData.Get(string(ImageID)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(ImageID)).(string); ok && v != "" {
				launchSpec.SetImageId(spotinst.String(resourceData.Get(string(ImageID)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[UserData] = commons.NewGenericField(
		commons.OceanECSLaunchSpec,
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
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var value = ""
			if launchSpec.UserData != nil {
				userData := launchSpec.UserData
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
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData := spotinst.String(base64Encode(v))
				launchSpec.SetUserData(userData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var userData *string = nil
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData = spotinst.String(base64Encode(v))
			}
			launchSpec.SetUserData(userData)
			return nil
		},
		nil,
	)

	fieldsMap[SecurityGroupIds] = commons.NewGenericField(
		commons.OceanECSLaunchSpec,
		SecurityGroupIds,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var value []string = nil
			if launchSpec.SecurityGroupIDs != nil {
				value = launchSpec.SecurityGroupIDs
			}
			if err := resourceData.Set(string(SecurityGroupIds), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SecurityGroupIds), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := clusterWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(SecurityGroupIds)).([]interface{}); ok {
				ids := make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
				launchSpec.SetSecurityGroupIDs(ids)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := clusterWrapper.GetLaunchSpec()
			var ids []string = nil
			if v, ok := resourceData.Get(string(SecurityGroupIds)).([]interface{}); ok && len(v) > 0 {
				ids = make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
			}
			launchSpec.SetSecurityGroupIDs(ids)
			return nil
		},
		nil,
	)

	fieldsMap[Attributes] = commons.NewGenericField(
		commons.OceanECSLaunchSpec,
		Attributes,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(AttributeKey): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(AttributeValue): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			Set: hashKV,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec.Attributes != nil {
				attributes := launchSpec.Attributes
				result = flattenAttributes(attributes)
			}
			if result != nil {
				if err := resourceData.Set(string(Attributes), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Attributes), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(Attributes)); ok {
				if attributes, err := expandAttributes(value); err != nil {
					return err
				} else {
					launchSpec.SetAttributes(attributes)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var attributeList []*aws.ECSAttribute = nil
			if value, ok := resourceData.GetOk(string(Attributes)); ok {
				if attributes, err := expandAttributes(value); err != nil {
					return err
				} else {
					attributeList = attributes
				}
			}
			launchSpec.SetAttributes(attributeList)
			return nil
		},
		nil,
	)

	fieldsMap[IamInstanceProfile] = commons.NewGenericField(
		commons.OceanECSLaunchSpec,
		IamInstanceProfile,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var value = ""
			if launchSpec.IAMInstanceProfile != nil {

				iam := launchSpec.IAMInstanceProfile
				if iam.ARN != nil {
					value = spotinst.StringValue(iam.ARN)
				} else if iam.Name != nil {
					value = spotinst.StringValue(iam.Name)
				}
			}
			if err := resourceData.Set(string(IamInstanceProfile), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IamInstanceProfile), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(IamInstanceProfile)).(string); ok && v != "" {
				iamInstanceProf := &aws.ECSIAMInstanceProfile{}
				if InstanceProfileArnRegex.MatchString(v) {
					iamInstanceProf.SetArn(spotinst.String(v))
				} else {
					iamInstanceProf.SetName(spotinst.String(v))
				}
				launchSpec.SetIAMInstanceProfile(iamInstanceProf)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(IamInstanceProfile)).(string); ok && v != "" {
				iamInstanceProf := &aws.ECSIAMInstanceProfile{}
				if InstanceProfileArnRegex.MatchString(v) {
					iamInstanceProf.SetArn(spotinst.String(v))
				} else {
					iamInstanceProf.SetName(spotinst.String(v))
				}
				launchSpec.SetIAMInstanceProfile(iamInstanceProf)
			} else {
				launchSpec.SetIAMInstanceProfile(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[AutoscaleHeadrooms] = commons.NewGenericField(
		commons.OceanECSLaunchSpec,
		AutoscaleHeadrooms,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(CPUPerUnit): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(MemoryPerUnit): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(NumOfUnits): {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec.AutoScale != nil && launchSpec.AutoScale.Headrooms != nil {
				headrooms := launchSpec.AutoScale.Headrooms
				result = flattenHeadrooms(headrooms)
			}
			if result != nil {
				if err := resourceData.Set(string(AutoscaleHeadrooms), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AutoscaleHeadrooms), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(AutoscaleHeadrooms)); ok {
				if headrooms, err := expandHeadrooms(value); err != nil {
					return err
				} else {
					launchSpec.AutoScale.SetHeadrooms(headrooms)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var headroomList []*aws.ECSAutoScaleHeadroom = nil
			if value, ok := resourceData.GetOk(string(AutoscaleHeadrooms)); ok {
				if expandedList, err := expandHeadrooms(value); err != nil {
					return err
				} else {
					headroomList = expandedList
				}
			}
			launchSpec.AutoScale.SetHeadrooms(headroomList)
			return nil
		},
		nil,
	)

	fieldsMap[Tags] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
		Tags,
		&schema.Schema{
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
						Required: true,
					},
				},
			},
			Set: hashKVTags,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec.Tags != nil {
				tags := launchSpec.Tags
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
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					launchSpec.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var tagsToAdd []*aws.Tag = nil
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					tagsToAdd = tags
				}
			}
			launchSpec.SetTags(tagsToAdd)
			return nil
		},
		nil,
	)

	fieldsMap[BlockDeviceMappings] = commons.NewGenericField(
		commons.OceanECSLaunchSpec,
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

					string(Ebs): {
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

								string(Encrypted): {
									Type:     schema.TypeBool,
									Optional: true,
									Computed: true,
								},

								string(IOPS): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(KMSKeyID): {
									Type:     schema.TypeString,
									Optional: true,
								},

								string(SnapshotID): {
									Type:     schema.TypeString,
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

								string(DynamicVolumeSize): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{

											string(BaseSize): {
												Type:     schema.TypeInt,
												Required: true,
											},

											string(Resource): {
												Type:     schema.TypeString,
												Required: true,
											},

											string(SizePerResourceUnit): {
												Type:     schema.TypeInt,
												Required: true,
											},
										},
									},
								},

								string(Throughput): {
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},

					string(NoDevice): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(VirtualName): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil

			if launchSpec != nil && launchSpec.BlockDeviceMappings != nil {
				result = flattenBlockDeviceMappings(launchSpec.BlockDeviceMappings)
			}

			if len(result) > 0 {
				if err := resourceData.Set(string(BlockDeviceMappings), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BlockDeviceMappings), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(BlockDeviceMappings)); ok {
				if v, err := expandBlockDeviceMappings(v); err != nil {
					return err
				} else {
					launchSpec.SetBlockDeviceMappings(v)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var value []*aws.ECSBlockDeviceMapping = nil

			if v, ok := resourceData.GetOk(string(BlockDeviceMappings)); ok {
				if blockDeviceMappings, err := expandBlockDeviceMappings(v); err != nil {
					return err
				} else {
					value = blockDeviceMappings
				}
			}
			launchSpec.SetBlockDeviceMappings(value)
			return nil
		},
		nil,
	)

	fieldsMap[InstanceTypes] = commons.NewGenericField(
		commons.OceanECSLaunchSpec,
		InstanceTypes,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			MinItems: 1,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value []string = nil
			if launchSpec.InstanceTypes != nil {
				value = launchSpec.InstanceTypes
			}
			if err := resourceData.Set(string(InstanceTypes), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(InstanceTypes), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(InstanceTypes)); ok {
				if instanceTypes, err := expandInstanceTypes(v); err != nil {
					return err
				} else {
					launchSpec.SetInstanceTypes(instanceTypes)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(InstanceTypes)); ok {
				if instanceTypes, err := expandInstanceTypes(v); err != nil {
					return err
				} else {
					launchSpec.SetInstanceTypes(instanceTypes)
				}
			} else {
				launchSpec.SetInstanceTypes(nil)
			}

			return nil
		},
		nil,
	)

	fieldsMap[RestrictScaleDown] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		RestrictScaleDown,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *bool = nil
			if launchSpec.RestrictScaleDown != nil {
				value = launchSpec.RestrictScaleDown
			}
			if value != nil {
				if err := resourceData.Set(string(RestrictScaleDown), spotinst.BoolValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RestrictScaleDown), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOkExists(string(RestrictScaleDown)); ok && v != nil {
				restrictScaleDown := spotinst.Bool(v.(bool))
				launchSpec.SetRestrictScaleDown(restrictScaleDown)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var restrictScaleDown *bool = nil
			if v, ok := resourceData.GetOkExists(string(RestrictScaleDown)); ok && v != nil {
				restrictScaleDown = spotinst.Bool(v.(bool))
			}
			launchSpec.SetRestrictScaleDown(restrictScaleDown)
			return nil
		},
		nil,
	)

	fieldsMap[SubnetIDs] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		SubnetIDs,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value []string = nil
			if launchSpec.SubnetIDs != nil {
				value = launchSpec.SubnetIDs
			}
			if err := resourceData.Set(string(SubnetIDs), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SubnetIDs), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(SubnetIDs)); ok {
				if subnetIDs, err := expandSubnetIDs(v); err != nil {
					return err
				} else {
					launchSpec.SetSubnetIDs(subnetIDs)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(SubnetIDs)); ok {
				if subnetIDs, err := expandSubnetIDs(v); err != nil {
					return err
				} else {
					launchSpec.SetSubnetIDs(subnetIDs)
				}
			} else {
				launchSpec.SetSubnetIDs(nil)
			}

			return nil
		},
		nil,
	)

	fieldsMap[SchedulingTask] = commons.NewGenericField(
		commons.OceanECSLaunchSpec,
		SchedulingTask,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(IsEnabled): {
						Type:     schema.TypeBool,
						Required: true,
					},
					string(CronExpression): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(TaskType): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(TaskHeadroom): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{

								string(CPUPerUnit): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(MemoryPerUnit): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(NumOfUnits): {
									Type:     schema.TypeInt,
									Required: true,
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec.LaunchSpecScheduling != nil && launchSpec.LaunchSpecScheduling.Tasks != nil {
				tasks := launchSpec.LaunchSpecScheduling.Tasks
				result = flattenTasks(tasks)
			}
			if result != nil {
				if err := resourceData.Set(string(SchedulingTask), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SchedulingTask), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(SchedulingTask)); ok {
				if tasks, err := expandTasks(value); err != nil {
					return err
				} else {
					launchSpec.LaunchSpecScheduling.SetTasks(tasks)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.ECSLaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value []*aws.ECSLaunchSpecTask = nil

			if v, ok := resourceData.GetOk(string(SchedulingTask)); ok {
				if tasks, err := expandTasks(v); err != nil {
					return err
				} else {
					value = tasks
				}
			}
			launchSpec.LaunchSpecScheduling.SetTasks(value)
			return nil
		},
		nil,
	)

}

func hashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(AttributeKey)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(AttributeValue)].(string)))
	return hashcode.String(buf.String())
}

func hashKVTags(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(LabelKey)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(LabelValue)].(string)))
	return hashcode.String(buf.String())
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

func expandAttributes(data interface{}) ([]*aws.ECSAttribute, error) {
	list := data.(*schema.Set).List()
	attributes := make([]*aws.ECSAttribute, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(AttributeKey)]; !ok {
			return nil, errors.New("invalid attribute: key missing")
		}

		if _, ok := attr[string(AttributeValue)]; !ok {
			return nil, errors.New("invalid attribute: value missing")
		}
		attribute := &aws.ECSAttribute{
			Key:   spotinst.String(attr[string(AttributeKey)].(string)),
			Value: spotinst.String(attr[string(AttributeValue)].(string)),
		}
		attributes = append(attributes, attribute)
	}
	return attributes, nil
}

func flattenAttributes(labels []*aws.ECSAttribute) []interface{} {
	result := make([]interface{}, 0, len(labels))
	for _, label := range labels {
		m := make(map[string]interface{})
		m[string(AttributeKey)] = spotinst.StringValue(label.Key)
		m[string(AttributeValue)] = spotinst.StringValue(label.Value)

		result = append(result, m)
	}
	return result
}

func expandHeadrooms(data interface{}) ([]*aws.ECSAutoScaleHeadroom, error) {
	list := data.(*schema.Set).List()
	headrooms := make([]*aws.ECSAutoScaleHeadroom, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		headroom := &aws.ECSAutoScaleHeadroom{
			CPUPerUnit:    spotinst.Int(attr[string(CPUPerUnit)].(int)),
			NumOfUnits:    spotinst.Int(attr[string(NumOfUnits)].(int)),
			MemoryPerUnit: spotinst.Int(attr[string(MemoryPerUnit)].(int)),
		}

		headrooms = append(headrooms, headroom)
	}
	return headrooms, nil
}

func flattenHeadrooms(headrooms []*aws.ECSAutoScaleHeadroom) []interface{} {
	result := make([]interface{}, 0, len(headrooms))

	for _, headroom := range headrooms {
		m := make(map[string]interface{})
		m[string(CPUPerUnit)] = spotinst.IntValue(headroom.CPUPerUnit)
		m[string(NumOfUnits)] = spotinst.IntValue(headroom.NumOfUnits)
		m[string(MemoryPerUnit)] = spotinst.IntValue(headroom.MemoryPerUnit)

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

func expandBlockDeviceMappings(data interface{}) ([]*aws.ECSBlockDeviceMapping, error) {

	list := data.([]interface{})
	bdms := make([]*aws.ECSBlockDeviceMapping, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		bdm := &aws.ECSBlockDeviceMapping{}

		if v, ok := attr[string(DeviceName)].(string); ok && v != "" {
			bdm.SetDeviceName(spotinst.String(v))
		}

		if r, ok := attr[string(Ebs)]; ok {
			if ebs, err := expandEbs(r); err != nil {
				return nil, err
			} else {
				bdm.SetEBS(ebs)
			}
		}

		if v, ok := attr[string(NoDevice)].(string); ok && v != "" {
			bdm.SetNoDevice(spotinst.String(v))
		}

		if v, ok := attr[string(VirtualName)].(string); ok && v != "" {
			bdm.SetVirtualName(spotinst.String(v))
		}
		bdms = append(bdms, bdm)
	}
	return bdms, nil
}

func expandEbs(data interface{}) (*aws.ECSEBS, error) {

	ebs := &aws.ECSEBS{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return ebs, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(DeleteOnTermination)].(bool); ok {
		ebs.SetDeleteOnTermination(spotinst.Bool(v))
	}

	if v, ok := m[string(Encrypted)].(bool); ok {
		ebs.SetEncrypted(spotinst.Bool(v))
	}

	if v, ok := m[string(IOPS)].(int); ok && v > 0 {
		ebs.SetIOPS(spotinst.Int(v))
	}

	if v, ok := m[string(KMSKeyID)].(string); ok && v != "" {
		ebs.SetKMSKeyId(spotinst.String(v))
	}

	if v, ok := m[string(SnapshotID)].(string); ok && v != "" {
		ebs.SetSnapshotId(spotinst.String(v))
	}

	if v, ok := m[string(VolumeSize)].(int); ok && v > 0 {
		ebs.SetVolumeSize(spotinst.Int(v))
	}

	if v, ok := m[string(VolumeType)].(string); ok && v != "" {
		ebs.SetVolumeType(spotinst.String(v))
	}

	if v, ok := m[string(DynamicVolumeSize)]; ok && v != nil {
		if dynamicVolumeSize, err := expandDynamicVolumeSize(v); err != nil {
			return nil, err
		} else {
			if dynamicVolumeSize != nil {
				ebs.SetDynamicVolumeSize(dynamicVolumeSize)
			}
		}
	}

	if v, ok := m[string(Throughput)].(int); ok && v > 0 {
		ebs.SetThroughput(spotinst.Int(v))
	}
	return ebs, nil
}

func expandDynamicVolumeSize(data interface{}) (*aws.ECSDynamicVolumeSize, error) {
	if list := data.([]interface{}); len(list) > 0 {
		dvs := &aws.ECSDynamicVolumeSize{}
		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(BaseSize)].(int); ok && v >= 0 {
				dvs.SetBaseSize(spotinst.Int(v))
			}

			if v, ok := m[string(Resource)].(string); ok && v != "" {
				dvs.SetResource(spotinst.String(v))
			}

			if v, ok := m[string(SizePerResourceUnit)].(int); ok && v >= 0 {
				dvs.SetSizePerResourceUnit(spotinst.Int(v))
			}
		}
		return dvs, nil
	}
	return nil, nil
}

func flattenBlockDeviceMappings(bdms []*aws.ECSBlockDeviceMapping) []interface{} {
	result := make([]interface{}, 0, len(bdms))

	for _, bdm := range bdms {
		m := make(map[string]interface{})
		m[string(DeviceName)] = spotinst.StringValue(bdm.DeviceName)
		if bdm.EBS != nil {
			m[string(Ebs)] = flattenEbs(bdm.EBS)
		}
		m[string(NoDevice)] = spotinst.StringValue(bdm.NoDevice)
		m[string(VirtualName)] = spotinst.StringValue(bdm.VirtualName)
		result = append(result, m)
	}
	return result

}

func flattenEbs(ebs *aws.ECSEBS) []interface{} {

	elasticBS := make(map[string]interface{})
	elasticBS[string(DeleteOnTermination)] = spotinst.BoolValue(ebs.DeleteOnTermination)
	elasticBS[string(Encrypted)] = spotinst.BoolValue(ebs.Encrypted)
	elasticBS[string(IOPS)] = spotinst.IntValue(ebs.IOPS)
	elasticBS[string(KMSKeyID)] = spotinst.StringValue(ebs.KMSKeyID)
	elasticBS[string(SnapshotID)] = spotinst.StringValue(ebs.SnapshotID)
	elasticBS[string(VolumeType)] = spotinst.StringValue(ebs.VolumeType)
	elasticBS[string(VolumeSize)] = spotinst.IntValue(ebs.VolumeSize)
	elasticBS[string(Throughput)] = spotinst.IntValue(ebs.Throughput)
	if ebs.DynamicVolumeSize != nil {
		elasticBS[string(DynamicVolumeSize)] = flattenDynamicVolumeSize(ebs.DynamicVolumeSize)
	}

	return []interface{}{elasticBS}
}

func flattenDynamicVolumeSize(dvs *aws.ECSDynamicVolumeSize) interface{} {

	DynamicVS := make(map[string]interface{})
	DynamicVS[string(BaseSize)] = spotinst.IntValue(dvs.BaseSize)
	DynamicVS[string(Resource)] = spotinst.StringValue(dvs.Resource)
	DynamicVS[string(SizePerResourceUnit)] = spotinst.IntValue(dvs.SizePerResourceUnit)

	return []interface{}{DynamicVS}
}

func expandInstanceTypes(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if instanceType, ok := v.(string); ok && instanceType != "" {
			result = append(result, instanceType)
		}
	}
	return result, nil
}

func expandSubnetIDs(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if subnetID, ok := v.(string); ok && subnetID != "" {
			result = append(result, subnetID)
		}
	}
	return result, nil
}

func flattenTasks(tasks []*aws.ECSLaunchSpecTask) []interface{} {
	result := make([]interface{}, 0, len(tasks))

	for _, task := range tasks {
		m := make(map[string]interface{})
		m[string(IsEnabled)] = spotinst.BoolValue(task.IsEnabled)
		m[string(CronExpression)] = spotinst.StringValue(task.CronExpression)
		m[string(TaskType)] = spotinst.StringValue(task.TaskType)

		if task.Config != nil && task.Config.TaskHeadrooms != nil {
			m[string(TaskHeadroom)] = flattenTaskHeadroom(task.Config.TaskHeadrooms)
		}

		result = append(result, m)
	}

	return result
}

func flattenTaskHeadroom(headrooms []*aws.ECSLaunchSpecTaskHeadroom) []interface{} {
	result := make([]interface{}, 0, len(headrooms))

	for _, headroom := range headrooms {
		m := make(map[string]interface{})
		m[string(CPUPerUnit)] = spotinst.IntValue(headroom.CPUPerUnit)
		m[string(NumOfUnits)] = spotinst.IntValue(headroom.NumOfUnits)
		m[string(MemoryPerUnit)] = spotinst.IntValue(headroom.MemoryPerUnit)

		result = append(result, m)
	}

	return result
}

func expandTasks(data interface{}) ([]*aws.ECSLaunchSpecTask, error) {
	list := data.(*schema.Set).List()
	tasks := make([]*aws.ECSLaunchSpecTask, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		task := &aws.ECSLaunchSpecTask{}

		if !ok {
			continue
		}

		if v, ok := attr[string(IsEnabled)].(bool); ok {
			task.SetIsEnabled(spotinst.Bool(v))
		}

		if v, ok := attr[string(CronExpression)].(string); ok && v != "" {
			task.SetCronExpression(spotinst.String(v))
		}

		if v, ok := attr[string(TaskType)].(string); ok && v != "" {
			task.SetTaskType(spotinst.String(v))
		}

		if v, ok := attr[string(TaskHeadroom)]; ok {
			if config, err := expandTaskHeadroom(v); err != nil {
				return nil, err
			} else {
				task.SetTaskConfig(config)
			}
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func expandTaskHeadroom(data interface{}) (*aws.ECSTaskConfig, error) {
	list := data.(*schema.Set).List()
	headrooms := make([]*aws.ECSLaunchSpecTaskHeadroom, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		headroom := &aws.ECSLaunchSpecTaskHeadroom{}

		if !ok {
			continue
		}

		if v, ok := attr[string(CPUPerUnit)].(int); ok {
			headroom.SetCPUPerUnit(spotinst.Int(v))
		}

		if v, ok := attr[string(NumOfUnits)].(int); ok {
			headroom.SetNumOfUnits(spotinst.Int(v))
		}

		if v, ok := attr[string(MemoryPerUnit)].(int); ok {
			headroom.SetMemoryPerUnit(spotinst.Int(v))
		}

		headrooms = append(headrooms, headroom)
	}

	taskConfig := &aws.ECSTaskConfig{
		TaskHeadrooms: headrooms,
	}

	return taskConfig, nil
}

package ocean_aws_launch_spec

import (
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[OceanID] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		OceanID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			launchSpec.SetOceanId(spotinst.String(resourceData.Get(string(OceanID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			launchSpec.SetOceanId(spotinst.String(resourceData.Get(string(OceanID)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[ImageID] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		ImageID,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(ImageID)); ok && value != nil {
				launchSpec.SetImageId(spotinst.String(resourceData.Get(string(ImageID)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var imageId *string = nil
			if value, ok := resourceData.GetOk(string(ImageID)); ok && value != nil {
				imageId = spotinst.String(resourceData.Get(string(ImageID)).(string))
			}
			launchSpec.SetImageId(imageId)
			return nil
		},
		nil,
	)

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(Name)); ok && value != nil {
				launchSpec.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(Name)); ok && value != nil {
				launchSpec.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[UserData] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData := spotinst.String(base64Encode(v))
				launchSpec.SetUserData(userData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var userData *string = nil
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData = spotinst.String(base64Encode(v))
			}
			launchSpec.SetUserData(userData)
			return nil
		},
		nil,
	)

	fieldsMap[Labels] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		Labels,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(LabelKey): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(LabelValue): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec.Labels != nil {
				labels := launchSpec.Labels
				result = flattenLabels(labels)
			}
			if result != nil {
				if err := resourceData.Set(string(Labels), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Labels), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(Labels)); ok {
				if labels, err := expandLabels(value); err != nil {
					return err
				} else {
					launchSpec.SetLabels(labels)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var labelList []*aws.Label = nil
			if value, ok := resourceData.GetOk(string(Labels)); ok {
				if labels, err := expandLabels(value); err != nil {
					return err
				} else {
					labelList = labels
				}
			}
			launchSpec.SetLabels(labelList)
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
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
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

	fieldsMap[ElasticIpPool] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		ElasticIpPool,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(TagSelector): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{

								string(TagSelectorKey): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(TagSelectorValue): {
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec.ElasticIPPool != nil {
				elasticIpPool := launchSpec.ElasticIPPool
				result = flattenElasticIpPool(elasticIpPool)
			}
			if result != nil {
				if err := resourceData.Set(string(ElasticIpPool), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ElasticIpPool), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(ElasticIpPool)); ok {
				if elasticIpPool, err := expandElasticIpPool(value); err != nil {
					return err
				} else {
					launchSpec.SetElasticIPPool(elasticIpPool)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *aws.ElasticIPPool = nil

			if v, ok := resourceData.GetOk(string(ElasticIpPool)); ok {
				if elasticIpPool, err := expandElasticIpPool(v); err != nil {
					return err
				} else {
					value = elasticIpPool
				}
			}
			launchSpec.SetElasticIPPool(value)
			return nil
		},
		nil,
	)

	fieldsMap[BlockDeviceMappings] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		BlockDeviceMappings,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(DeviceName): {
						Type:     schema.TypeString,
						Optional: true,
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value []*aws.BlockDeviceMapping = nil

			if v, ok := resourceData.GetOk(string(BlockDeviceMappings)); ok {
				if blockdevicemappings, err := expandBlockDeviceMappings(v); err != nil {
					return err
				} else {
					value = blockdevicemappings
				}
			}
			launchSpec.SetBlockDeviceMappings(value)
			return nil
		},
		nil,
	)

	fieldsMap[ResourceLimits] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
		ResourceLimits,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(MaxInstanceCount): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
					string(MinInstanceCount): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec.ResourceLimits != nil {
				resourceLimits := launchSpec.ResourceLimits
				result = flattenResourceLimits(resourceLimits)
			}
			if result != nil {
				if err := resourceData.Set(string(ResourceLimits), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ResourceLimits), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(ResourceLimits)); ok {
				if resourceLimits, err := expandResourceLimits(value); err != nil {
					return err
				} else {
					launchSpec.SetResourceLimits(resourceLimits)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *aws.ResourceLimits = nil

			if v, ok := resourceData.GetOk(string(ResourceLimits)); ok {
				if resourceLimits, err := expandResourceLimits(v); err != nil {
					return err
				} else {
					value = resourceLimits
				}
			}
			launchSpec.SetResourceLimits(value)
			return nil
		},
		nil,
	)

	fieldsMap[SecurityGroups] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		SecurityGroups,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value []string = nil
			if launchSpec.SecurityGroupIDs != nil {
				value = launchSpec.SecurityGroupIDs
			}
			if err := resourceData.Set(string(SecurityGroups), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SecurityGroups), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(SecurityGroups)).([]interface{}); ok {
				ids := make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
				launchSpec.SetSecurityGroupIDs(ids)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(SecurityGroups)).([]interface{}); ok {
				ids := make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
				launchSpec.SetSecurityGroupIDs(ids)
			}
			return nil
		},
		nil,
	)

	fieldsMap[Taints] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		Taints,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(TaintKey): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(TaintValue): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(Effect): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec.Labels != nil {
				taints := launchSpec.Taints
				result = flattenTaints(taints)
			}
			if result != nil {
				if err := resourceData.Set(string(Taints), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Taints), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(Taints)); ok {
				if labels, err := expandTaints(value); err != nil {
					return err
				} else {
					launchSpec.SetTaints(labels)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var taintList []*aws.Taint = nil
			if value, ok := resourceData.GetOk(string(Taints)); ok {
				if taints, err := expandTaints(value); err != nil {
					return err
				} else {
					taintList = taints
				}
			}
			launchSpec.SetTaints(taintList)
			return nil
		},
		nil,
	)

	fieldsMap[IamInstanceProfile] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		IamInstanceProfile,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(IamInstanceProfile)).(string); ok && v != "" {
				iamInstanceProf := &aws.IAMInstanceProfile{}
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(IamInstanceProfile)).(string); ok && v != "" {
				iamInstanceProf := &aws.IAMInstanceProfile{}
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

	fieldsMap[AutoscaleHeadroomsAutomatic] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		AutoscaleHeadroomsAutomatic,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(AutoHeadroomPercentage): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
						DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
							if old == "-1" && new == "null" {
								return true
							}
							return false
						},
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec != nil && launchSpec.AutoScale != nil && launchSpec.AutoScale.AutoHeadroomPercentage != nil {
				result = flattenAutoscaleHeadroomsAutomatic(launchSpec.AutoScale)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(AutoscaleHeadroomsAutomatic), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AutoscaleHeadroomsAutomatic), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(AutoscaleHeadroomsAutomatic)); ok {
				if err := expandAutoscaleHeadroomsAutomatic(v, launchSpec, false); err != nil {
					return err
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(AutoscaleHeadroomsAutomatic)); ok {
				if err := expandAutoscaleHeadroomsAutomatic(v, launchSpec, true); err != nil {
					return err
				}
			} else {
				launchSpec.AutoScale.SetAutoHeadroomPercentage(nil)
			}
			return nil
		},

		nil,
	)

	fieldsMap[AutoscaleHeadrooms] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
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

					string(GPUPerUnit): {
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var headroomList []*aws.AutoScaleHeadroom = nil
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
	fieldsMap[AutoscaleDown] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		AutoscaleDown,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MaxScaleDownPercentage): {
						Type:     schema.TypeFloat,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec.AutoScale != nil && launchSpec.AutoScale.Down != nil {
				down := launchSpec.AutoScale.Down
				result = flattenAutoscaleDown(down)
			}
			if result != nil {
				if err := resourceData.Set(string(AutoscaleDown), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AutoscaleDown), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(AutoscaleDown)); ok {
				if autoscaleDown, err := expandAutoscaleDown(value); err != nil {
					return err
				} else {
					launchSpec.AutoScale.SetAutoScalerDownVNG(autoscaleDown)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result *aws.AutoScalerDownVNG = nil
			if value, ok := resourceData.GetOk(string(AutoscaleDown)); ok {
				if autoscaleDown, err := expandAutoscaleDown(value); err != nil {
					return err
				} else {
					result = autoscaleDown
				}
			}
			launchSpec.AutoScale.SetAutoScalerDownVNG(result)
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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

	fieldsMap[InstanceTypes] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		InstanceTypes,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			MinItems: 1,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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

	fieldsMap[PreferredSpotTypes] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		PreferredSpotTypes,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value []string = nil
			if launchSpec.PreferredSpotTypes != nil {
				value = launchSpec.PreferredSpotTypes
			}
			if err := resourceData.Set(string(PreferredSpotTypes), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PreferredSpotTypes), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(PreferredSpotTypes)); ok {
				if preferredSpotTypes, err := expandPreferredSpotTypes(v); err != nil {
					return err
				} else {
					launchSpec.SetPreferredSpotTypes(preferredSpotTypes)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(PreferredSpotTypes)); ok {
				if preferredSpotTypes, err := expandInstanceTypes(v); err != nil {
					return err
				} else {
					launchSpec.SetPreferredSpotTypes(preferredSpotTypes)
				}
			} else {
				launchSpec.SetPreferredSpotTypes(nil)
			}

			return nil
		},
		nil,
	)

	fieldsMap[RootVolumeSize] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		RootVolumeSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *int = nil
			if launchSpec.RootVolumeSize != nil {
				value = launchSpec.RootVolumeSize
			}
			if err := resourceData.Set(string(RootVolumeSize), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RootVolumeSize), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(RootVolumeSize)).(int); ok && v > 0 {
				launchSpec.SetRootVolumeSize(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *int = nil
			if v, ok := resourceData.Get(string(RootVolumeSize)).(int); ok && v > 0 {
				value = spotinst.Int(v)
			}
			launchSpec.SetRootVolumeSize(value)
			return nil
		},
		nil,
	)

	fieldsMap[Strategy] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		Strategy,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(SpotPercentage): {
						Type:         schema.TypeInt,
						Optional:     true,
						Default:      -1,
						ValidateFunc: validation.IntAtLeast(-1),
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec.Strategy != nil {
				strategy := launchSpec.Strategy
				result = flattenStrategy(strategy)
			}
			if result != nil {
				if err := resourceData.Set(string(Strategy), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Strategy), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(Strategy)); ok {
				if strategy, err := expandStrategy(value); err != nil {
					return err
				} else {
					launchSpec.SetStrategy(strategy)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *aws.LaunchSpecStrategy = nil

			if v, ok := resourceData.GetOk(string(Strategy)); ok {
				if strategy, err := expandStrategy(v); err != nil {
					return err
				} else {
					value = strategy
				}
			}
			launchSpec.SetStrategy(value)
			return nil
		},
		nil,
	)

	fieldsMap[AssociatePublicIPAddress] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		AssociatePublicIPAddress,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *bool = nil
			if launchSpec.AssociatePublicIPAddress != nil {
				value = launchSpec.AssociatePublicIPAddress
			}
			if value != nil {
				if err := resourceData.Set(string(AssociatePublicIPAddress), spotinst.BoolValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AssociatePublicIPAddress), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOkExists(string(AssociatePublicIPAddress)); ok && v != nil {
				associatePublicIPAddress := spotinst.Bool(v.(bool))
				launchSpec.SetAssociatePublicIPAddress(associatePublicIPAddress)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var associatePublicIPAddress *bool = nil
			if v, ok := resourceData.GetOkExists(string(AssociatePublicIPAddress)); ok && v != nil {
				associatePublicIPAddress = spotinst.Bool(v.(bool))
			}
			launchSpec.SetAssociatePublicIPAddress(associatePublicIPAddress)
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOkExists(string(RestrictScaleDown)); ok && v != nil {
				restrictScaleDown := spotinst.Bool(v.(bool))
				launchSpec.SetRestrictScaleDown(restrictScaleDown)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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

	fieldsMap[CreateOptions] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		CreateOptions,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(InitialNodes): {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		nil, nil, nil, nil,
	)

	fieldsMap[UpdatePolicy] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		UpdatePolicy,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ShouldRoll): {
						Type:     schema.TypeBool,
						Required: true,
					},

					string(RollConfig): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(BatchSizePercentage): {
									Type:     schema.TypeInt,
									Required: true,
								},
							},
						},
					},
				},
			},
		},
		nil, nil, nil, nil,
	)

	fieldsMap[DeleteOptions] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		DeleteOptions,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ForceDelete): {
						Type:     schema.TypeBool,
						Required: true,
					},
					string(DeleteNodes): {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		nil, nil, nil, nil,
	)

	fieldsMap[SchedulingTask] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
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

								string(GPUPerUnit): {
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value []*aws.LaunchSpecTask = nil

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

	fieldsMap[SchedulingShutdownHours] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		SchedulingShutdownHours,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(IsEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(TimeWindows): {
						Type:     schema.TypeList,
						Elem:     &schema.Schema{Type: schema.TypeString},
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil

			if launchSpec != nil && launchSpec.LaunchSpecScheduling != nil &&
				launchSpec.LaunchSpecScheduling.ShutdownHours != nil {
				result = flattenShutdownHours(launchSpec.LaunchSpecScheduling.ShutdownHours)
			}

			if len(result) > 0 {
				if err := resourceData.Set(string(SchedulingShutdownHours), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SchedulingShutdownHours), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(SchedulingShutdownHours)); ok {
				if shutdownHours, err := expandShutdownHours(value); err != nil {
					return err
				} else {
					launchSpec.LaunchSpecScheduling.SetShutdownHours(shutdownHours)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *aws.LaunchSpecShutdownHours = nil

			if v, ok := resourceData.GetOk(string(SchedulingShutdownHours)); ok {
				if shutdownHours, err := expandShutdownHours(v); err != nil {
					return err
				} else {
					value = shutdownHours
				}
			}
			launchSpec.LaunchSpecScheduling.SetShutdownHours(value)
			return nil
		},
		nil,
	)

	fieldsMap[InstanceMetadataOptions] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		InstanceMetadataOptions,
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
						// Value mentioned below is used to set HTTPPutResponseHopLimit field to null when the customer doesn't want to set this param, as terraform set it 0 for integer type param by default
						Default: 1357997531,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec != nil && launchSpec.InstanceMetadataOptions != nil {
				result = flattenInstanceMetadataOptions(launchSpec.InstanceMetadataOptions)
			}

			if result != nil {
				if err := resourceData.Set(string(InstanceMetadataOptions), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(InstanceMetadataOptions), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(InstanceMetadataOptions)); ok {
				if metaDataOptions, err := expandInstanceMetadataOptions(v); err != nil {
					return err
				} else {
					launchSpec.SetLaunchspecInstanceMetadataOptions(metaDataOptions)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *aws.LaunchspecInstanceMetadataOptions = nil
			if v, ok := resourceData.GetOk(string(InstanceMetadataOptions)); ok {
				if metaDataOptions, err := expandInstanceMetadataOptions(v); err != nil {
					return err
				} else {
					value = metaDataOptions
				}
			}
			launchSpec.SetLaunchspecInstanceMetadataOptions(value)
			return nil
		},
		nil,
	)

	fieldsMap[Images] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		Images,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ImageId): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec.Images != nil {
				images := launchSpec.Images
				result = flattenImages(images)
			}
			if result != nil {
				if err := resourceData.Set(string(Images), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Images), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(Images)); ok {
				if images, err := expandImages(value); err != nil {
					return err
				} else {
					launchSpec.SetImages(images)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var imagesToAdd []*aws.Images = nil
			if value, ok := resourceData.GetOk(string(Images)); ok {
				if images, err := expandImages(value); err != nil {
					return err
				} else {
					imagesToAdd = images
				}
			}
			launchSpec.SetImages(imagesToAdd)
			return nil
		},
		nil,
	)
	fieldsMap[InstanceTypesFilters] = commons.NewGenericField(
		commons.OceanAWSInstanceTypes,
		InstanceTypesFilters,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(Categories): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(DiskTypes): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(ExcludeFamilies): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(ExcludeMetal): {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},

					string(Hypervisor): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(IncludeFamilies): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(IsEnaSupported): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(MaxGpu): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(MaxMemoryGiB): {
						Type:     schema.TypeFloat,
						Optional: true,
						Default:  -1,
					},

					string(MaxNetworkPerformance): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(MaxVcpu): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(MinEnis): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(MinGpu): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(MinMemoryGiB): {
						Type:     schema.TypeFloat,
						Optional: true,
						Default:  -1,
					},

					string(MinNetworkPerformance): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(MinVcpu): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(RootDeviceTypes): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(VirtualizationTypes): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil

			if launchSpec != nil && launchSpec.InstanceTypesFilters != nil {
				result = flattenInstanceTypesFilters(launchSpec.InstanceTypesFilters)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(InstanceTypesFilters), result); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(InstanceTypesFilters), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(InstanceTypesFilters)); ok {
				if instanceTypesFilters, err := expandInstanceTypesFilters(v, false); err != nil {
					return err
				} else {
					launchSpec.SetInstanceTypesFilters(instanceTypesFilters)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *aws.InstanceTypesFilters = nil

			if v, ok := resourceData.GetOk(string(InstanceTypesFilters)); ok {
				if instanceTypesFilters, err := expandInstanceTypesFilters(v, true); err != nil {
					return err
				} else {
					value = instanceTypesFilters
				}
			}
			if launchSpec.InstanceTypesFilters == nil {
				launchSpec.InstanceTypesFilters = &aws.InstanceTypesFilters{}
			}
			launchSpec.SetInstanceTypesFilters(value)
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

func expandLabels(data interface{}) ([]*aws.Label, error) {
	list := data.(*schema.Set).List()
	labels := make([]*aws.Label, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(LabelKey)]; !ok {
			return nil, errors.New("invalid label attributes: key missing")
		}

		if _, ok := attr[string(LabelValue)]; !ok {
			return nil, errors.New("invalid label attributes: value missing")
		}
		label := &aws.Label{
			Key:   spotinst.String(attr[string(LabelKey)].(string)),
			Value: spotinst.String(attr[string(LabelValue)].(string)),
		}
		labels = append(labels, label)
	}
	return labels, nil
}

func flattenLabels(labels []*aws.Label) []interface{} {
	result := make([]interface{}, 0, len(labels))
	for _, label := range labels {
		m := make(map[string]interface{})
		m[string(LabelKey)] = spotinst.StringValue(label.Key)
		m[string(LabelValue)] = spotinst.StringValue(label.Value)

		result = append(result, m)
	}
	return result
}

func expandTaints(data interface{}) ([]*aws.Taint, error) {
	list := data.(*schema.Set).List()
	taints := make([]*aws.Taint, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(TaintKey)]; !ok {
			return nil, errors.New("invalid taint attributes: key missing")
		}

		if _, ok := attr[string(TaintValue)]; !ok {
			return nil, errors.New("invalid taint attributes: value missing")
		}

		if _, ok := attr[string(Effect)]; !ok {
			return nil, errors.New("invalid taint attributes: effect missing")
		}

		taint := &aws.Taint{
			Key:    spotinst.String(attr[string(TaintKey)].(string)),
			Value:  spotinst.String(attr[string(TaintValue)].(string)),
			Effect: spotinst.String(attr[string(Effect)].(string)),
		}
		taints = append(taints, taint)
	}
	return taints, nil
}

func flattenTaints(taints []*aws.Taint) []interface{} {
	result := make([]interface{}, 0, len(taints))
	for _, taint := range taints {
		m := make(map[string]interface{})
		m[string(TaintKey)] = spotinst.StringValue(taint.Key)
		m[string(TaintValue)] = spotinst.StringValue(taint.Value)
		m[string(Effect)] = spotinst.StringValue(taint.Effect)

		result = append(result, m)
	}
	return result
}

func expandAutoscaleHeadroomsAutomatic(data interface{}, ls *aws.LaunchSpec, nullify bool) error {
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return nil
	}

	m := list[0].(map[string]interface{})

	if v, ok := m[string(AutoHeadroomPercentage)].(int); ok && v > -1 {
		ls.AutoScale.SetAutoHeadroomPercentage(spotinst.Int(v))
	} else if nullify {
		ls.AutoScale.SetAutoHeadroomPercentage(nil)
	}

	return nil
}

func expandHeadrooms(data interface{}) ([]*aws.AutoScaleHeadroom, error) {
	list := data.(*schema.Set).List()
	headrooms := make([]*aws.AutoScaleHeadroom, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		headroom := &aws.AutoScaleHeadroom{
			CPUPerUnit:    spotinst.Int(attr[string(CPUPerUnit)].(int)),
			GPUPerUnit:    spotinst.Int(attr[string(GPUPerUnit)].(int)),
			NumOfUnits:    spotinst.Int(attr[string(NumOfUnits)].(int)),
			MemoryPerUnit: spotinst.Int(attr[string(MemoryPerUnit)].(int)),
		}

		headrooms = append(headrooms, headroom)
	}
	return headrooms, nil
}

func flattenAutoscaleHeadroomsAutomatic(autoScale *aws.AutoScale) []interface{} {
	var out []interface{}

	if autoScale != nil {

		result := make(map[string]interface{})
		value := spotinst.Int(-1)
		if autoScale.AutoHeadroomPercentage != nil {
			value = autoScale.AutoHeadroomPercentage
		}
		result[string(AutoHeadroomPercentage)] = spotinst.IntValue(value)

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func flattenHeadrooms(headrooms []*aws.AutoScaleHeadroom) []interface{} {
	result := make([]interface{}, 0, len(headrooms))

	for _, headroom := range headrooms {
		m := make(map[string]interface{})
		m[string(CPUPerUnit)] = spotinst.IntValue(headroom.CPUPerUnit)
		m[string(GPUPerUnit)] = spotinst.IntValue(headroom.GPUPerUnit)
		m[string(NumOfUnits)] = spotinst.IntValue(headroom.NumOfUnits)
		m[string(MemoryPerUnit)] = spotinst.IntValue(headroom.MemoryPerUnit)

		result = append(result, m)
	}

	return result
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

func expandInstanceTypes(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if instanceTypes, ok := v.(string); ok && instanceTypes != "" {
			result = append(result, instanceTypes)
		}
	}
	return result, nil
}

func expandPreferredSpotTypes(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if preferredSpotTypes, ok := v.(string); ok && preferredSpotTypes != "" {
			result = append(result, preferredSpotTypes)
		}
	}
	return result, nil
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

func expandElasticIpPool(data interface{}) (*aws.ElasticIPPool, error) {
	elasticIpPool := &aws.ElasticIPPool{}
	list := data.(*schema.Set).List()

	if list == nil || list[0] == nil {
		return elasticIpPool, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(TagSelector)]; ok {
		tagSelector, err := expandTagSelector(v)
		if err != nil {
			return nil, err
		}
		if tagSelector != nil {
			elasticIpPool.SetTagSelector(tagSelector)
		} else {
			elasticIpPool.SetTagSelector(nil)
		}
	}
	return elasticIpPool, nil

}

func expandTagSelector(data interface{}) (*aws.TagSelector, error) {
	if list := data.([]interface{}); len(list) > 0 {
		tagSelector := &aws.TagSelector{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(TagSelectorKey)].(string); ok && v != "" {
				tagSelector.SetTagKey(spotinst.String(v))
			}

			if v, ok := m[string(TagSelectorValue)].(string); ok && v != "" {
				tagSelector.SetTagValue(spotinst.String(v))
			}
		}
		return tagSelector, nil
	}

	return nil, nil
}

func flattenElasticIpPool(elasticIpPool *aws.ElasticIPPool) []interface{} {
	var out []interface{}

	if elasticIpPool != nil {
		result := make(map[string]interface{})

		if elasticIpPool.TagSelector != nil {
			result[string(TagSelector)] = flattenTagSelector(elasticIpPool.TagSelector)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func expandBlockDeviceMappings(data interface{}) ([]*aws.BlockDeviceMapping, error) {

	list := data.([]interface{})
	bdms := make([]*aws.BlockDeviceMapping, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		bdm := &aws.BlockDeviceMapping{}

		if !ok {
			continue
		}

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

func expandEbs(data interface{}) (*aws.EBS, error) {

	ebs := &aws.EBS{}
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

func expandDynamicVolumeSize(data interface{}) (*aws.DynamicVolumeSize, error) {
	if list := data.([]interface{}); len(list) > 0 {
		dvs := &aws.DynamicVolumeSize{}
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

func flattenBlockDeviceMappings(bdms []*aws.BlockDeviceMapping) []interface{} {
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

func flattenEbs(ebs *aws.EBS) []interface{} {

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

func flattenDynamicVolumeSize(dvs *aws.DynamicVolumeSize) interface{} {

	DynamicVS := make(map[string]interface{})
	DynamicVS[string(BaseSize)] = spotinst.IntValue(dvs.BaseSize)
	DynamicVS[string(Resource)] = spotinst.StringValue(dvs.Resource)
	DynamicVS[string(SizePerResourceUnit)] = spotinst.IntValue(dvs.SizePerResourceUnit)

	return []interface{}{DynamicVS}
}

func flattenResourceLimits(resourceLimits *aws.ResourceLimits) []interface{} {
	var out []interface{}

	if resourceLimits != nil {
		result := make(map[string]interface{})

		value := spotinst.Int(-1)
		result[string(MinInstanceCount)] = value
		result[string(MaxInstanceCount)] = value

		if resourceLimits.MaxInstanceCount != nil {
			result[string(MaxInstanceCount)] = spotinst.IntValue(resourceLimits.MaxInstanceCount)
		}
		if resourceLimits.MinInstanceCount != nil {
			result[string(MinInstanceCount)] = spotinst.IntValue(resourceLimits.MinInstanceCount)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func flattenTagSelector(tagSelector *aws.TagSelector) []interface{} {
	m := make(map[string]interface{})
	m[string(TagSelectorKey)] = spotinst.StringValue(tagSelector.Key)
	m[string(TagSelectorValue)] = spotinst.StringValue(tagSelector.Value)

	return []interface{}{m}
}

func expandResourceLimits(data interface{}) (*aws.ResourceLimits, error) {
	if list := data.([]interface{}); len(list) > 0 {
		resLimits := &aws.ResourceLimits{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(MaxInstanceCount)].(int); ok && v >= 0 {
				resLimits.SetMaxInstanceCount(spotinst.Int(v))
			} else {
				resLimits.SetMaxInstanceCount(nil)
			}

			if v, ok := m[string(MinInstanceCount)].(int); ok && v >= 0 {
				resLimits.SetMinInstanceCount(spotinst.Int(v))
			} else {
				resLimits.SetMinInstanceCount(nil)
			}
		}
		return resLimits, nil
	}

	return nil, nil
}

func expandStrategy(data interface{}) (*aws.LaunchSpecStrategy, error) {
	if list := data.(*schema.Set).List(); len(list) > 0 {
		strategy := &aws.LaunchSpecStrategy{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(SpotPercentage)].(int); ok && v > -1 {
				strategy.SetSpotPercentage(spotinst.Int(v))
			} else {
				strategy.SetSpotPercentage(nil)
			}
		}
		return strategy, nil
	}
	return nil, nil
}

func flattenStrategy(strategy *aws.LaunchSpecStrategy) []interface{} {
	var out []interface{}

	if strategy != nil {
		result := make(map[string]interface{})

		if strategy.SpotPercentage != nil {
			result[string(SpotPercentage)] = spotinst.IntValue(strategy.SpotPercentage)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

func flattenTasks(tasks []*aws.LaunchSpecTask) []interface{} {
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

func flattenTaskHeadroom(headrooms []*aws.LaunchSpecTaskHeadroom) []interface{} {
	result := make([]interface{}, 0, len(headrooms))

	for _, headroom := range headrooms {
		m := make(map[string]interface{})
		m[string(CPUPerUnit)] = spotinst.IntValue(headroom.CPUPerUnit)
		m[string(GPUPerUnit)] = spotinst.IntValue(headroom.GPUPerUnit)
		m[string(NumOfUnits)] = spotinst.IntValue(headroom.NumOfUnits)
		m[string(MemoryPerUnit)] = spotinst.IntValue(headroom.MemoryPerUnit)

		result = append(result, m)
	}

	return result
}

func expandTasks(data interface{}) ([]*aws.LaunchSpecTask, error) {
	list := data.(*schema.Set).List()
	tasks := make([]*aws.LaunchSpecTask, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		task := &aws.LaunchSpecTask{}

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

func expandTaskHeadroom(data interface{}) (*aws.TaskConfig, error) {
	list := data.(*schema.Set).List()
	headrooms := make([]*aws.LaunchSpecTaskHeadroom, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		headroom := &aws.LaunchSpecTaskHeadroom{}

		if !ok {
			continue
		}

		if v, ok := attr[string(CPUPerUnit)].(int); ok {
			headroom.SetCPUPerUnit(spotinst.Int(v))
		}

		if v, ok := attr[string(GPUPerUnit)].(int); ok {
			headroom.SetGPUPerUnit(spotinst.Int(v))
		}

		if v, ok := attr[string(NumOfUnits)].(int); ok {
			headroom.SetNumOfUnits(spotinst.Int(v))
		}

		if v, ok := attr[string(MemoryPerUnit)].(int); ok {
			headroom.SetMemoryPerUnit(spotinst.Int(v))
		}

		headrooms = append(headrooms, headroom)
	}

	taskConfig := &aws.TaskConfig{
		TaskHeadrooms: headrooms,
	}

	return taskConfig, nil
}

func flattenShutdownHours(shutdownHours *aws.LaunchSpecShutdownHours) []interface{} {
	result := make(map[string]interface{})

	if shutdownHours != nil {

		result[string(IsEnabled)] = spotinst.BoolValue(shutdownHours.IsEnabled)

		if shutdownHours.TimeWindows != nil {
			result[string(TimeWindows)] = shutdownHours.TimeWindows
		}
	}

	return []interface{}{result}
}

func expandShutdownHours(data interface{}) (*aws.LaunchSpecShutdownHours, error) {
	shutdownHours := &aws.LaunchSpecShutdownHours{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return shutdownHours, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(IsEnabled)].(bool); ok {
		shutdownHours.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(TimeWindows)]; ok {
		timeWindows, err := expandTimeWindows(v)
		if err != nil {
			return nil, err
		}
		if timeWindows != nil {
			shutdownHours.SetTimeWindows(timeWindows)
		} else {
			shutdownHours.SetTimeWindows(nil)
		}
	}

	return shutdownHours, nil
}

func expandTimeWindows(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if timeWindow, ok := v.(string); ok && timeWindow != "" {
			result = append(result, timeWindow)
		}
	}

	return result, nil
}

func flattenAutoscaleDown(down *aws.AutoScalerDownVNG) []interface{} {
	var out []interface{}
	if down != nil {
		result := make(map[string]interface{})
		if down.MaxScaleDownPercentage != nil {
			result[string(MaxScaleDownPercentage)] = spotinst.Float64Value(down.MaxScaleDownPercentage)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

func expandAutoscaleDown(down interface{}) (*aws.AutoScalerDownVNG, error) {
	if list := down.([]interface{}); len(list) > 0 {
		autoscaleDown := &aws.AutoScalerDownVNG{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})
			if v, ok := m[string(MaxScaleDownPercentage)].(float64); ok {
				autoscaleDown.SetMaxScaleDownPercentage(spotinst.Float64(v))
			} else {
				autoscaleDown.SetMaxScaleDownPercentage(nil)
			}
		}
		return autoscaleDown, nil
	}
	return nil, nil
}

func flattenInstanceMetadataOptions(instanceMetadataOptions *aws.LaunchspecInstanceMetadataOptions) []interface{} {
	result := make(map[string]interface{})
	result[string(HTTPTokens)] = spotinst.StringValue(instanceMetadataOptions.HTTPTokens)
	result[string(HTTPPutResponseHopLimit)] = spotinst.IntValue(instanceMetadataOptions.HTTPPutResponseHopLimit)

	return []interface{}{result}
}
func expandInstanceMetadataOptions(data interface{}) (*aws.LaunchspecInstanceMetadataOptions, error) {
	instanceMetadataOptions := &aws.LaunchspecInstanceMetadataOptions{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return instanceMetadataOptions, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(HTTPTokens)].(string); ok && v != "" {
		instanceMetadataOptions.SetHTTPTokens(spotinst.String(v))
	}
	if v, ok := m[string(HTTPPutResponseHopLimit)].(int); ok {
		// Value(1357997531) mentioned below is used to set HTTPPutResponseHopLimit field to null when the customer doesn't want to set this param, as terraform set it 0 for integer type param by default.
		if v == 1357997531 {
			instanceMetadataOptions.SetHTTPPutResponseHopLimit(nil)
		} else {
			instanceMetadataOptions.SetHTTPPutResponseHopLimit(spotinst.Int(v))
		}
	}

	return instanceMetadataOptions, nil
}

func flattenImages(images []*aws.Images) []interface{} {
	result := make([]interface{}, 0, len(images))
	for _, image := range images {
		m := make(map[string]interface{})
		m[string(ImageId)] = spotinst.StringValue(image.ImageId)
		result = append(result, m)
	}
	return result
}

func expandImages(data interface{}) ([]*aws.Images, error) {
	list := data.([]interface{})
	images := make([]*aws.Images, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(ImageId)]; !ok {
			return nil, errors.New("invalid tag attributes: value missing")
		}
		image := &aws.Images{
			ImageId: spotinst.String(attr[string(ImageId)].(string)),
		}
		images = append(images, image)
	}
	return images, nil
}
func flattenInstanceTypesFilters(instanceTypesFilters *aws.InstanceTypesFilters) []interface{} {
	var out []interface{}

	if instanceTypesFilters != nil {
		result := make(map[string]interface{})
		value := spotinst.Int(-1)
		result[string(MaxGpu)] = value
		result[string(MinGpu)] = value
		result[string(MaxMemoryGiB)] = value
		result[string(MinMemoryGiB)] = value
		result[string(MaxVcpu)] = value
		result[string(MinVcpu)] = value
		result[string(MaxNetworkPerformance)] = value
		result[string(MinNetworkPerformance)] = value
		result[string(MinEnis)] = value

		if instanceTypesFilters.MaxGpu != nil {
			result[string(MaxGpu)] = spotinst.IntValue(instanceTypesFilters.MaxGpu)
		}
		if instanceTypesFilters.MinGpu != nil {
			result[string(MinGpu)] = spotinst.IntValue(instanceTypesFilters.MinGpu)
		}
		if instanceTypesFilters.MaxMemoryGiB != nil {
			result[string(MaxMemoryGiB)] = spotinst.Float64Value(instanceTypesFilters.MaxMemoryGiB)
		}
		if instanceTypesFilters.MinMemoryGiB != nil {
			result[string(MinMemoryGiB)] = spotinst.Float64Value(instanceTypesFilters.MinMemoryGiB)
		}
		if instanceTypesFilters.MaxVcpu != nil {
			result[string(MaxVcpu)] = spotinst.IntValue(instanceTypesFilters.MaxVcpu)
		}
		if instanceTypesFilters.MinVcpu != nil {
			result[string(MinVcpu)] = spotinst.IntValue(instanceTypesFilters.MinVcpu)
		}
		if instanceTypesFilters.MaxNetworkPerformance != nil {
			result[string(MaxNetworkPerformance)] = spotinst.IntValue(instanceTypesFilters.MaxNetworkPerformance)
		}
		if instanceTypesFilters.MinNetworkPerformance != nil {
			result[string(MinNetworkPerformance)] = spotinst.IntValue(instanceTypesFilters.MinNetworkPerformance)
		}
		if instanceTypesFilters.MinEnis != nil {
			result[string(MinEnis)] = spotinst.IntValue(instanceTypesFilters.MinEnis)
		}

		result[string(ExcludeMetal)] = spotinst.BoolValue(instanceTypesFilters.ExcludeMetal)

		if instanceTypesFilters.IsEnaSupported != nil {
			if *instanceTypesFilters.IsEnaSupported == true {
				b := "true"
				result[string(IsEnaSupported)] = b
			} else {
				b := "false"
				result[string(IsEnaSupported)] = b
			}
		}

		if instanceTypesFilters.Categories != nil {
			result[string(Categories)] = instanceTypesFilters.Categories
		}

		if instanceTypesFilters.DiskTypes != nil {
			result[string(DiskTypes)] = instanceTypesFilters.DiskTypes
		}

		if instanceTypesFilters.ExcludeFamilies != nil {
			result[string(ExcludeFamilies)] = instanceTypesFilters.ExcludeFamilies
		}
		if instanceTypesFilters.Hypervisor != nil {
			result[string(Hypervisor)] = instanceTypesFilters.Hypervisor
		}

		if instanceTypesFilters.IncludeFamilies != nil {
			result[string(IncludeFamilies)] = instanceTypesFilters.IncludeFamilies
		}

		if instanceTypesFilters.RootDeviceTypes != nil {
			result[string(RootDeviceTypes)] = instanceTypesFilters.RootDeviceTypes
		}

		if instanceTypesFilters.VirtualizationTypes != nil {
			result[string(VirtualizationTypes)] = instanceTypesFilters.VirtualizationTypes
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func expandInstanceTypesFilters(data interface{}, nullify bool) (*aws.InstanceTypesFilters, error) {
	instanceTypesFilters := &aws.InstanceTypesFilters{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return instanceTypesFilters, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Categories)]; ok {
		categories, err := expandInstanceTypeFiltersList(v)
		if err != nil {
			return nil, err
		}
		if categories != nil && len(categories) > 0 {
			instanceTypesFilters.SetCategories(categories)
		} else {
			if nullify {
				instanceTypesFilters.SetCategories(nil)
			}
		}
	}

	if v, ok := m[string(DiskTypes)]; ok {
		diskTypes, err := expandInstanceTypeFiltersList(v)
		if err != nil {
			return nil, err
		}
		if diskTypes != nil && len(diskTypes) > 0 {
			instanceTypesFilters.SetDiskTypes(diskTypes)
		} else {
			if nullify {
				instanceTypesFilters.SetDiskTypes(nil)
			}
		}
	}

	if v, ok := m[string(ExcludeFamilies)]; ok {
		excludeFamilies, err := expandInstanceTypeFiltersList(v)
		if err != nil {
			return nil, err
		}
		if excludeFamilies != nil && len(excludeFamilies) > 0 {
			instanceTypesFilters.SetExcludeFamilies(excludeFamilies)
		} else {
			if nullify {
				instanceTypesFilters.SetExcludeFamilies(nil)
			}
		}
	}

	if v, ok := m[string(Hypervisor)]; ok {
		hypervisor, err := expandInstanceTypeFiltersList(v)
		if err != nil {
			return nil, err
		}
		if hypervisor != nil && len(hypervisor) > 0 {
			instanceTypesFilters.SetHypervisor(hypervisor)
		} else {
			if nullify {
				instanceTypesFilters.SetHypervisor(nil)
			}
		}
	}

	if v, ok := m[string(ExcludeMetal)].(bool); ok {
		instanceTypesFilters.SetExcludeMetal(spotinst.Bool(v))
	}

	if v, ok := m[string(IncludeFamilies)]; ok {
		includeFamilies, err := expandInstanceTypeFiltersList(v)
		if err != nil {
			return nil, err
		}
		if includeFamilies != nil && len(includeFamilies) > 0 {
			instanceTypesFilters.SetIncludeFamilies(includeFamilies)
		} else {
			if nullify {
				instanceTypesFilters.SetIncludeFamilies(nil)
			}
		}
	}

	if v, ok := m[string(IsEnaSupported)].(string); ok {
		if v == "true" {
			instanceTypesFilters.SetIsEnaSupported(spotinst.Bool(true))
		} else if v == "false" {
			instanceTypesFilters.SetIsEnaSupported(spotinst.Bool(false))
		} else {
			instanceTypesFilters.SetIsEnaSupported(nil)
		}
	}

	if v, ok := m[string(MaxGpu)].(int); ok {
		if v == -1 {
			instanceTypesFilters.SetMaxGpu(nil)
		} else {
			instanceTypesFilters.SetMaxGpu(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MaxMemoryGiB)].(float64); ok {
		if v == -1 {
			instanceTypesFilters.SetMaxMemoryGiB(nil)
		} else {
			instanceTypesFilters.SetMaxMemoryGiB(spotinst.Float64(v))
		}
	}

	if v, ok := m[string(MaxNetworkPerformance)].(int); ok {
		if v == -1 {
			instanceTypesFilters.SetMaxNetworkPerformance(nil)
		} else {
			instanceTypesFilters.SetMaxNetworkPerformance(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MaxVcpu)].(int); ok {
		if v == -1 {
			instanceTypesFilters.SetMaxVcpu(nil)
		} else {
			instanceTypesFilters.SetMaxVcpu(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MinEnis)].(int); ok {
		if v == -1 {
			instanceTypesFilters.SetMinEnis(nil)
		} else {
			instanceTypesFilters.SetMinEnis(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MinGpu)].(int); ok {
		if v == -1 {
			instanceTypesFilters.SetMinGpu(nil)
		} else {
			instanceTypesFilters.SetMinGpu(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MinMemoryGiB)].(float64); ok {
		if v == -1 {
			instanceTypesFilters.SetMinMemoryGiB(nil)
		} else {
			instanceTypesFilters.SetMinMemoryGiB(spotinst.Float64(v))
		}
	}

	if v, ok := m[string(MinNetworkPerformance)].(int); ok {
		if v == -1 {
			instanceTypesFilters.SetMinNetworkPerformance(nil)
		} else {
			instanceTypesFilters.SetMinNetworkPerformance(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MinVcpu)].(int); ok {
		if v == -1 {
			instanceTypesFilters.SetMinVcpu(nil)
		} else {
			instanceTypesFilters.SetMinVcpu(spotinst.Int(v))
		}
	}

	if v, ok := m[string(RootDeviceTypes)]; ok {
		rootDevicetypes, err := expandInstanceTypeFiltersList(v)
		if err != nil {
			return nil, err
		}
		if rootDevicetypes != nil && len(rootDevicetypes) > 0 {
			instanceTypesFilters.SetRootDeviceTypes(rootDevicetypes)
		} else {
			if nullify {
				instanceTypesFilters.SetRootDeviceTypes(nil)
			}
		}
	}

	if v, ok := m[string(VirtualizationTypes)]; ok {
		virtualizationtypes, err := expandInstanceTypeFiltersList(v)
		if err != nil {
			return nil, err
		}
		if virtualizationtypes != nil && len(virtualizationtypes) > 0 {
			instanceTypesFilters.SetVirtualizationTypes(virtualizationtypes)
		} else {
			if nullify {
				instanceTypesFilters.SetVirtualizationTypes(nil)
			}
		}
	}

	return instanceTypesFilters, nil
}

func expandInstanceTypeFiltersList(data interface{}) ([]string, error) {
	list := data.(*schema.Set).List()
	result := make([]string, 0, len(list))

	for _, v := range list {
		if instanceTypeList, ok := v.(string); ok && instanceTypeList != "" {
			result = append(result, instanceTypeList)
		}
	}
	return result, nil
}

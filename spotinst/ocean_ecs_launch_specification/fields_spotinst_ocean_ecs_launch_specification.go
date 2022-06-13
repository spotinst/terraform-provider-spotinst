package ocean_ecs_launch_specification

import (
	"encoding/base64"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[ImageID] = commons.NewGenericField(
		commons.OceanECSLaunchSpecification,
		ImageID,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value = ""
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.ImageID != nil {
				value = spotinst.StringValue(cluster.Compute.LaunchSpecification.ImageID)
			}
			if err := resourceData.Set(string(ImageID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ImageID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.Get(string(ImageID)).(string); ok && v != "" {
				cluster.Compute.LaunchSpecification.SetImageId(spotinst.String(resourceData.Get(string(ImageID)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.Get(string(ImageID)).(string); ok && v != "" {
				cluster.Compute.LaunchSpecification.SetImageId(spotinst.String(resourceData.Get(string(ImageID)).(string)))
			} else {
				cluster.Compute.LaunchSpecification.SetImageId(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[IamInstanceProfile] = commons.NewGenericField(
		commons.OceanECSLaunchSpecification,
		IamInstanceProfile,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value = ""
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.IAMInstanceProfile != nil {

				lc := cluster.Compute.LaunchSpecification
				if lc.IAMInstanceProfile.ARN != nil {
					value = spotinst.StringValue(lc.IAMInstanceProfile.ARN)
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
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.Get(string(IamInstanceProfile)).(string); ok && v != "" {
				iamInstanceProf := &aws.ECSIAMInstanceProfile{}
				if InstanceProfileArnRegex.MatchString(v) {
					iamInstanceProf.SetArn(spotinst.String(v))
				} else {
					iamInstanceProf.SetName(spotinst.String(v))
				}
				cluster.Compute.LaunchSpecification.SetIAMInstanceProfile(iamInstanceProf)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.Get(string(IamInstanceProfile)).(string); ok && v != "" {
				iamInstanceProf := &aws.ECSIAMInstanceProfile{}
				if InstanceProfileArnRegex.MatchString(v) {
					iamInstanceProf.SetArn(spotinst.String(v))
				} else {
					iamInstanceProf.SetName(spotinst.String(v))
				}
				cluster.Compute.LaunchSpecification.SetIAMInstanceProfile(iamInstanceProf)
			} else {
				cluster.Compute.LaunchSpecification.SetIAMInstanceProfile(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[SecurityGroupIds] = commons.NewGenericField(
		commons.OceanECSLaunchSpecification,
		SecurityGroupIds,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value []string = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.SecurityGroupIDs != nil {
				value = cluster.Compute.LaunchSpecification.SecurityGroupIDs
			}
			if err := resourceData.Set(string(SecurityGroupIds), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SecurityGroupIds), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.Get(string(SecurityGroupIds)).([]interface{}); ok {
				ids := make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
				cluster.Compute.LaunchSpecification.SetSecurityGroupIDs(ids)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var ids []string = nil
			if v, ok := resourceData.Get(string(SecurityGroupIds)).([]interface{}); ok && len(v) > 0 {
				ids = make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
			}
			cluster.Compute.LaunchSpecification.SetSecurityGroupIDs(ids)
			return nil
		},
		nil,
	)

	fieldsMap[UserData] = commons.NewGenericField(
		commons.OceanECSLaunchSpecification,
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
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value = ""
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.UserData != nil {

				userData := cluster.Compute.LaunchSpecification.UserData
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
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData := spotinst.String(base64Encode(v))
				cluster.Compute.LaunchSpecification.SetUserData(userData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var userData *string = nil
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData = spotinst.String(base64Encode(v))
			}
			cluster.Compute.LaunchSpecification.SetUserData(userData)
			return nil
		},
		nil,
	)

	fieldsMap[KeyPair] = commons.NewGenericField(
		commons.OceanECSLaunchSpecification,
		KeyPair,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value *string = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.KeyPair != nil {
				value = cluster.Compute.LaunchSpecification.KeyPair
			}
			if err := resourceData.Set(string(KeyPair), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(KeyPair), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.Get(string(KeyPair)).(string); ok && v != "" {
				cluster.Compute.LaunchSpecification.SetKeyPair(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var key *string = nil
			if v, ok := resourceData.Get(string(KeyPair)).(string); ok && v != "" {
				key = spotinst.String(v)
			}
			cluster.Compute.LaunchSpecification.SetKeyPair(key)
			return nil
		},
		nil,
	)

	fieldsMap[AssociatePublicIpAddress] = commons.NewGenericField(
		commons.OceanECSLaunchSpecification,
		AssociatePublicIpAddress,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()

			var value *bool = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.AssociatePublicIPAddress != nil {

				value = cluster.Compute.LaunchSpecification.AssociatePublicIPAddress
			}

			if err := resourceData.Set(string(AssociatePublicIpAddress), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AssociatePublicIpAddress), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()

			if v, ok := resourceData.GetOkExists(string(AssociatePublicIpAddress)); ok {
				cluster.Compute.LaunchSpecification.SetAssociatePublicIPAddress(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()

			if v, ok := resourceData.GetOkExists(string(AssociatePublicIpAddress)); ok {
				cluster.Compute.LaunchSpecification.SetAssociatePublicIPAddress(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Monitoring] = commons.NewGenericField(
		commons.OceanECSLaunchSpecification,
		Monitoring,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()

			var value *bool = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.Monitoring != nil {

				value = cluster.Compute.LaunchSpecification.Monitoring
			}

			if err := resourceData.Set(string(Monitoring), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Monitoring), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()

			if v, ok := resourceData.GetOkExists(string(Monitoring)); ok {
				cluster.Compute.LaunchSpecification.SetMonitoring(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()

			if v, ok := resourceData.GetOkExists(string(Monitoring)); ok {
				cluster.Compute.LaunchSpecification.SetMonitoring(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[EBSOptimized] = commons.NewGenericField(
		commons.OceanECSLaunchSpecification,
		EBSOptimized,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()

			var value *bool = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.EBSOptimized != nil {

				value = cluster.Compute.LaunchSpecification.EBSOptimized
			}

			if err := resourceData.Set(string(EBSOptimized), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(EBSOptimized), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()

			if v, ok := resourceData.GetOkExists(string(EBSOptimized)); ok {
				cluster.Compute.LaunchSpecification.SetEBSOptimized(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()

			if v, ok := resourceData.GetOkExists(string(EBSOptimized)); ok {
				cluster.Compute.LaunchSpecification.SetEBSOptimized(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[UseAsTemplateOnly] = commons.NewGenericField(
		commons.OceanECSLaunchSpecification,
		UseAsTemplateOnly,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()

			var value *bool = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.UseAsTemplateOnly != nil {

				value = cluster.Compute.LaunchSpecification.UseAsTemplateOnly
			}

			if err := resourceData.Set(string(UseAsTemplateOnly), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UseAsTemplateOnly), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()

			if v, ok := resourceData.GetOkExists(string(UseAsTemplateOnly)); ok {
				cluster.Compute.LaunchSpecification.SetUseAsTemplateOnly(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()

			if v, ok := resourceData.GetOkExists(string(UseAsTemplateOnly)); ok {
				cluster.Compute.LaunchSpecification.SetUseAsTemplateOnly(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[BlockDeviceMappings] = commons.NewGenericField(
		commons.OceanECSLaunchSpecification,
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
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var result []interface{} = nil

			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.BlockDeviceMappings != nil {
				result = flattenBlockDeviceMappings(cluster.Compute.LaunchSpecification.BlockDeviceMappings)
			}

			if len(result) > 0 {
				if err := resourceData.Set(string(BlockDeviceMappings), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BlockDeviceMappings), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.GetOk(string(BlockDeviceMappings)); ok {
				if v, err := expandBlockDeviceMappings(v); err != nil {
					return err
				} else {
					cluster.Compute.LaunchSpecification.SetBlockDeviceMappings(v)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value []*aws.ECSBlockDeviceMapping = nil

			if v, ok := resourceData.GetOk(string(BlockDeviceMappings)); ok {
				if blockDeviceMappings, err := expandBlockDeviceMappings(v); err != nil {
					return err
				} else {
					value = blockDeviceMappings
				}
			}
			cluster.Compute.LaunchSpecification.SetBlockDeviceMappings(value)
			return nil
		},
		nil,
	)

	fieldsMap[InstanceMetadataOptions] = commons.NewGenericField(
		commons.OceanECSLaunchSpecification,
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
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var result []interface{} = nil
			if cluster != nil && cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.InstanceMetadataOptions != nil {
				result = flattenInstanceMetadataOptions(cluster.Compute.LaunchSpecification.InstanceMetadataOptions)
			}

			if result != nil {
				if err := resourceData.Set(string(InstanceMetadataOptions), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(InstanceMetadataOptions), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			if v, ok := resourceData.GetOk(string(InstanceMetadataOptions)); ok {
				if metaDataOptions, err := expandInstanceMetadataOptions(v); err != nil {
					return err
				} else {
					cluster.Compute.LaunchSpecification.SetInstanceMetadataOptions(metaDataOptions)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ECSClusterWrapper)
			cluster := clusterWrapper.GetECSCluster()
			var value *aws.ECSInstanceMetadataOptions = nil
			if v, ok := resourceData.GetOk(string(InstanceMetadataOptions)); ok {
				if metaDataOptions, err := expandInstanceMetadataOptions(v); err != nil {
					return err
				} else {
					value = metaDataOptions
				}
			}
			cluster.Compute.LaunchSpecification.SetInstanceMetadataOptions(value)
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

		if r, ok := attr[string(EBS)]; ok {
			if ebs, err := expandEBS(r); err != nil {
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

func expandEBS(data interface{}) (*aws.ECSEBS, error) {
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
			m[string(EBS)] = flattenEbs(bdm.EBS)
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
	vs := make(map[string]interface{})
	vs[string(BaseSize)] = spotinst.IntValue(dvs.BaseSize)
	vs[string(Resource)] = spotinst.StringValue(dvs.Resource)
	vs[string(SizePerResourceUnit)] = spotinst.IntValue(dvs.SizePerResourceUnit)
	return []interface{}{vs}
}

func expandInstanceMetadataOptions(data interface{}) (*aws.ECSInstanceMetadataOptions, error) {
	instanceMetadataOptions := &aws.ECSInstanceMetadataOptions{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return instanceMetadataOptions, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(HTTPTokens)].(string); ok && v != "" {
		instanceMetadataOptions.SetHTTPTokens(spotinst.String(v))
	}
	if v, ok := m[string(HTTPPutResponseHopLimit)].(int); ok && v >= 0 {
		instanceMetadataOptions.SetHTTPPutResponseHopLimit(spotinst.Int(v))
	} else {
		instanceMetadataOptions.SetHTTPPutResponseHopLimit(nil)
	}

	return instanceMetadataOptions, nil
}

func flattenInstanceMetadataOptions(instanceMetadataOptions *aws.ECSInstanceMetadataOptions) []interface{} {
	result := make(map[string]interface{})
	result[string(HTTPTokens)] = spotinst.StringValue(instanceMetadataOptions.HTTPTokens)
	result[string(HTTPPutResponseHopLimit)] = spotinst.IntValue(instanceMetadataOptions.HTTPPutResponseHopLimit)

	return []interface{}{result}
}

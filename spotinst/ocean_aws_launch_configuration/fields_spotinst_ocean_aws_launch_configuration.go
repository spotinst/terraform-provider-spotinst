package ocean_aws_launch_configuration

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
		commons.OceanAWSLaunchConfiguration,
		ImageID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.ImageID != nil {
				value = cluster.Compute.LaunchSpecification.ImageID
			}
			if err := resourceData.Set(string(ImageID), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ImageID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(ImageID)).(string); ok && v != "" {
				cluster.Compute.LaunchSpecification.SetImageId(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if v, ok := resourceData.Get(string(ImageID)).(string); ok && v != "" {
				value = spotinst.String(v)
			}
			cluster.Compute.LaunchSpecification.SetImageId(value)
			return nil
		},
		nil,
	)

	fieldsMap[RootVolumeSize] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
		RootVolumeSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *int = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.RootVolumeSize != nil {
				value = cluster.Compute.LaunchSpecification.RootVolumeSize
			}
			if err := resourceData.Set(string(RootVolumeSize), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RootVolumeSize), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(RootVolumeSize)).(int); ok && v > 0 {
				cluster.Compute.LaunchSpecification.SetRootVolumeSize(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *int = nil
			if v, ok := resourceData.Get(string(RootVolumeSize)).(int); ok && v > 0 {
				value = spotinst.Int(v)
			}
			cluster.Compute.LaunchSpecification.SetRootVolumeSize(value)
			return nil
		},
		nil,
	)

	fieldsMap[IAMInstanceProfile] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
		IAMInstanceProfile,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
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
			if err := resourceData.Set(string(IAMInstanceProfile), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IAMInstanceProfile), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(IAMInstanceProfile)).(string); ok && v != "" {
				iamInstanceProf := &aws.IAMInstanceProfile{}
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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(IAMInstanceProfile)).(string); ok && v != "" {
				iamInstanceProf := &aws.IAMInstanceProfile{}
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

	fieldsMap[KeyName] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
		KeyName,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.KeyPair != nil {
				value = cluster.Compute.LaunchSpecification.KeyPair
			}
			if err := resourceData.Set(string(KeyName), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(KeyName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(KeyName)).(string); ok && v != "" {
				cluster.Compute.LaunchSpecification.SetKeyPair(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var key *string = nil
			if v, ok := resourceData.Get(string(KeyName)).(string); ok && v != "" {
				key = spotinst.String(v)
			}
			cluster.Compute.LaunchSpecification.SetKeyPair(key)
			return nil
		},
		nil,
	)

	fieldsMap[SecurityGroups] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
		SecurityGroups,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value []string = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.SecurityGroupIDs != nil {
				value = cluster.Compute.LaunchSpecification.SecurityGroupIDs
			}
			if err := resourceData.Set(string(SecurityGroups), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SecurityGroups), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(SecurityGroups)).([]interface{}); ok {
				ids := make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
				cluster.Compute.LaunchSpecification.SetSecurityGroupIDs(ids)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(SecurityGroups)).([]interface{}); ok {
				ids := make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
				cluster.Compute.LaunchSpecification.SetSecurityGroupIDs(ids)
			}
			return nil
		},
		nil,
	)

	fieldsMap[UserData] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData := spotinst.String(base64Encode(v))
				cluster.Compute.LaunchSpecification.SetUserData(userData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var userData *string = nil
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData = spotinst.String(base64Encode(v))
			}
			cluster.Compute.LaunchSpecification.SetUserData(userData)
			return nil
		},
		nil,
	)

	fieldsMap[AssociatePublicIpAddress] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
		AssociatePublicIpAddress,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(AssociatePublicIpAddress)); ok {
				cluster.Compute.LaunchSpecification.SetAssociatePublicIPAddress(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(AssociatePublicIpAddress)); ok {
				cluster.Compute.LaunchSpecification.SetAssociatePublicIPAddress(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[AssociateIPv6Address] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
		AssociateIPv6Address,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			var value *bool = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.AssociateIPv6Address != nil {

				value = cluster.Compute.LaunchSpecification.AssociateIPv6Address
			}

			if err := resourceData.Set(string(AssociateIPv6Address), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AssociateIPv6Address), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(AssociateIPv6Address)); ok {
				cluster.Compute.LaunchSpecification.SetAssociateIPv6Address(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(AssociateIPv6Address)); ok {
				cluster.Compute.LaunchSpecification.SetAssociateIPv6Address(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[LoadBalancers] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
		LoadBalancers,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Arn): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(Name): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(Type): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			var balancerResults []interface{} = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil && cluster.Compute.LaunchSpecification.LoadBalancers != nil {
				balancers := cluster.Compute.LaunchSpecification.LoadBalancers
				balancerResults = flattenLoadBalancers(balancers)
			}

			if err := resourceData.Set(string(LoadBalancers), balancerResults); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(LoadBalancers), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if value, ok := resourceData.GetOk(string(LoadBalancers)); ok {
				if lb, err := expandLb(value); err != nil {
					return err
				} else {
					cluster.Compute.LaunchSpecification.SetLoadBalancers(lb)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			var result []*aws.LoadBalancer = nil
			if lbs := cluster.Compute.LaunchSpecification.LoadBalancers; len(lbs) > 0 {
				result = append(result, lbs...)
			}
			if value, ok := resourceData.GetOk(string(LoadBalancers)); ok {
				if lb, err := expandLb(value); err != nil {
					return err
				} else {
					result = append(result, lb...)
				}
			}

			cluster.Compute.LaunchSpecification.SetLoadBalancers(result)
			return nil
		},
		nil,
	)

	fieldsMap[Monitoring] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
		Monitoring,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(Monitoring)); ok {
				cluster.Compute.LaunchSpecification.SetMonitoring(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(Monitoring)); ok {
				cluster.Compute.LaunchSpecification.SetMonitoring(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[EBSOptimized] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
		EBSOptimized,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(EBSOptimized)); ok {
				cluster.Compute.LaunchSpecification.SetEBSOptimized(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(EBSOptimized)); ok {
				cluster.Compute.LaunchSpecification.SetEBSOptimized(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		nil,
	)
	fieldsMap[UseAsTemplateOnly] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
		UseAsTemplateOnly,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(UseAsTemplateOnly)); ok {
				cluster.Compute.LaunchSpecification.SetUseAsTemplateOnly(spotinst.Bool(v.(bool)))
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(UseAsTemplateOnly)); ok {
				cluster.Compute.LaunchSpecification.SetUseAsTemplateOnly(spotinst.Bool(v.(bool)))
			}

			return nil
		},
		nil,
	)

	fieldsMap[InstanceMetadataOptions] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *aws.InstanceMetadataOptions = nil
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
	fieldsMap[BlockDeviceMappings] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
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

								string(DynamicIops): {
									Type:     schema.TypeList,
									Optional: true,
									MaxItems: 1,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{

											string(IopsBaseSize): {
												Type:     schema.TypeInt,
												Required: true,
											},

											string(IopsResource): {
												Type:     schema.TypeString,
												Required: true,
											},

											string(IopsSizePerResourceUnit): {
												Type:     schema.TypeInt,
												Required: true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil

			if cluster != nil && cluster.Compute.LaunchSpecification.BlockDeviceMappings != nil {
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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(BlockDeviceMappings)); ok {
				if v, err := expandBlockDeviceMappings(v); err != nil {
					return err
				} else {
					cluster.Compute.LaunchSpecification.SetClusterBlockDeviceMappings(v)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value []*aws.ClusterBlockDeviceMappings = nil

			if v, ok := resourceData.GetOk(string(BlockDeviceMappings)); ok {
				if blockdevicemappings, err := expandBlockDeviceMappings(v); err != nil {
					return err
				} else {
					value = blockdevicemappings
				}
			}
			cluster.Compute.LaunchSpecification.SetClusterBlockDeviceMappings(value)
			return nil
		},
		nil,
	)

	fieldsMap[ResourceTagSpecification] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
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
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.ResourceTagSpecification != nil {
				resourceTagSpecification := cluster.Compute.LaunchSpecification.ResourceTagSpecification
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
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(ResourceTagSpecification)); ok {
				if v, err := expandResourceTagSpecification(v); err != nil {
					return err
				} else {
					cluster.Compute.LaunchSpecification.SetResourceTagSpecification(v)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *aws.ResourceTagSpecification = nil
			if v, ok := resourceData.GetOk(string(ResourceTagSpecification)); ok {
				if resourceTagSpecification, err := expandResourceTagSpecification(v); err != nil {
					return err
				} else {
					value = resourceTagSpecification
				}
			}
			cluster.Compute.LaunchSpecification.SetResourceTagSpecification(value)
			return nil
		},
		nil,
	)
	fieldsMap[HealthCheckUnhealthyDurationBeforeReplacement] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
		HealthCheckUnhealthyDurationBeforeReplacement,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *int = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.HealthCheckUnhealthyDurationBeforeReplacement != nil {
				value = cluster.Compute.LaunchSpecification.HealthCheckUnhealthyDurationBeforeReplacement
			}
			if err := resourceData.Set(string(HealthCheckUnhealthyDurationBeforeReplacement), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(HealthCheckUnhealthyDurationBeforeReplacement), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(HealthCheckUnhealthyDurationBeforeReplacement)).(int); ok && v > 0 {
				cluster.Compute.LaunchSpecification.SetHealthCheckUnhealthyDurationBeforeReplacement(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *int = nil
			if v, ok := resourceData.Get(string(HealthCheckUnhealthyDurationBeforeReplacement)).(int); ok && v > 0 {
				value = spotinst.Int(v)
			}
			cluster.Compute.LaunchSpecification.SetHealthCheckUnhealthyDurationBeforeReplacement(value)
			return nil
		},
		nil,
	)

}

func flattenResourceTagSpecification(resourceTagSpecification *aws.ResourceTagSpecification) []interface{} {
	result := make(map[string]interface{})

	if resourceTagSpecification != nil && resourceTagSpecification.Volumes != nil {
		result[string(ShouldTagVolumes)] = spotinst.BoolValue(resourceTagSpecification.Volumes.ShouldTag)
	}

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

	return resourceTagSpecification, nil
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

func expandLb(lb interface{}) ([]*aws.LoadBalancer, error) {
	list := lb.([]interface{})
	lbOutput := make([]*aws.LoadBalancer, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		singleLb := &aws.LoadBalancer{}
		if arn, ok := attr[string(Arn)]; ok && arn != "" {
			singleLb.SetArn(spotinst.String(arn.(string)))
		}
		if name, ok := attr[string(Name)]; ok && name != "" {
			singleLb.SetName(spotinst.String(name.(string)))
		}
		if typeIn, ok := attr[string(Type)]; ok && typeIn != "" {
			singleLb.SetType(spotinst.String(typeIn.(string)))
		}

		lbOutput = append(lbOutput, singleLb)
	}
	return lbOutput, nil
}

func flattenLoadBalancers(balancers []*aws.LoadBalancer) []interface{} {
	result := make([]interface{}, 0, len(balancers))
	for _, balancer := range balancers {
		m := make(map[string]interface{})
		m[string(Arn)] = spotinst.StringValue(balancer.Arn)
		m[string(Name)] = spotinst.StringValue(balancer.Name)
		m[string(Type)] = spotinst.StringValue(balancer.Type)

		result = append(result, m)
	}
	return result
}

func expandInstanceMetadataOptions(data interface{}) (*aws.InstanceMetadataOptions, error) {
	instanceMetadataOptions := &aws.InstanceMetadataOptions{}
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

func flattenInstanceMetadataOptions(instanceMetadataOptions *aws.InstanceMetadataOptions) []interface{} {
	result := make(map[string]interface{})
	result[string(HTTPTokens)] = spotinst.StringValue(instanceMetadataOptions.HTTPTokens)
	result[string(HTTPPutResponseHopLimit)] = spotinst.IntValue(instanceMetadataOptions.HTTPPutResponseHopLimit)

	return []interface{}{result}
}
func flattenBlockDeviceMappings(bdms []*aws.ClusterBlockDeviceMappings) []interface{} {
	result := make([]interface{}, 0, len(bdms))

	for _, bdm := range bdms {
		m := make(map[string]interface{})
		m[string(DeviceName)] = spotinst.StringValue(bdm.DeviceName)
		if bdm.EBS != nil {
			m[string(Ebs)] = flattenEbs(bdm.EBS)
		}

		result = append(result, m)
	}
	return result

}
func flattenEbs(ebs *aws.ClusterEBS) []interface{} {

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

	if ebs.DynamicIops != nil {
		elasticBS[string(DynamicIops)] = flattenDynamicIops(ebs.DynamicIops)
	}

	return []interface{}{elasticBS}
}
func flattenDynamicVolumeSize(dvs *aws.ClusterDynamicVolumeSize) interface{} {

	DynamicVS := make(map[string]interface{})
	DynamicVS[string(BaseSize)] = spotinst.IntValue(dvs.BaseSize)
	DynamicVS[string(Resource)] = spotinst.StringValue(dvs.Resource)
	DynamicVS[string(SizePerResourceUnit)] = spotinst.IntValue(dvs.SizePerResourceUnit)

	return []interface{}{DynamicVS}
}
func flattenDynamicIops(dvs *aws.ClusterDynamicIops) interface{} {

	dynamicIops := make(map[string]interface{})
	dynamicIops[string(IopsBaseSize)] = spotinst.IntValue(dvs.BaseSize)
	dynamicIops[string(IopsResource)] = spotinst.StringValue(dvs.Resource)
	dynamicIops[string(IopsSizePerResourceUnit)] = spotinst.IntValue(dvs.SizePerResourceUnit)

	return []interface{}{dynamicIops}
}
func expandBlockDeviceMappings(data interface{}) ([]*aws.ClusterBlockDeviceMappings, error) {

	list := data.([]interface{})
	bdms := make([]*aws.ClusterBlockDeviceMappings, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		bdm := &aws.ClusterBlockDeviceMappings{}

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
		bdms = append(bdms, bdm)
	}
	return bdms, nil
}
func expandEbs(data interface{}) (*aws.ClusterEBS, error) {

	ebs := &aws.ClusterEBS{}
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

	if v, ok := m[string(DynamicIops)]; ok && v != nil {
		if dynamicIops, err := expandDynamicIops(v); err != nil {
			return nil, err
		} else {
			if dynamicIops != nil {
				ebs.SetDynamicIops(dynamicIops)
			}
		}
	}

	return ebs, nil
}
func expandDynamicVolumeSize(data interface{}) (*aws.ClusterDynamicVolumeSize, error) {
	if list := data.([]interface{}); len(list) > 0 {
		dvs := &aws.ClusterDynamicVolumeSize{}
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

func expandDynamicIops(data interface{}) (*aws.ClusterDynamicIops, error) {
	if list := data.([]interface{}); len(list) > 0 {
		dvs := &aws.ClusterDynamicIops{}
		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(IopsBaseSize)].(int); ok && v >= 0 {
				dvs.SetBaseSize(spotinst.Int(v))
			}

			if v, ok := m[string(IopsResource)].(string); ok && v != "" {
				dvs.SetResource(spotinst.String(v))
			}

			if v, ok := m[string(IopsSizePerResourceUnit)].(int); ok && v >= 0 {
				dvs.SetSizePerResourceUnit(spotinst.Int(v))
			}
		}
		return dvs, nil
	}
	return nil, nil
}

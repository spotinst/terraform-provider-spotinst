package ocean_aws_launch_configuration

import (
	"encoding/base64"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ImageID] = commons.NewGenericField(
		commons.OceanAWSLaunchConfiguration,
		ImageID,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
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
			if v, ok := resourceData.Get(string(ImageID)).(string); ok && v != "" {
				cluster.Compute.LaunchSpecification.SetImageId(spotinst.String(v))
			}
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
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
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

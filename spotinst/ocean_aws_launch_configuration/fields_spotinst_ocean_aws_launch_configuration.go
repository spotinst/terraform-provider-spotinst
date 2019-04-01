package ocean_aws_launch_configuration

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"regexp"
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
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(ImageID)).(string); ok && v != "" {
				cluster.Compute.LaunchSpecification.SetImageId(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(ImageID)).(string); ok && v != "" {
				cluster.Compute.LaunchSpecification.SetImageId(spotinst.String(v))
			}
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
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(KeyName)).(string); ok && v != "" {
				cluster.Compute.LaunchSpecification.SetKeyPair(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
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
			StateFunc: HexStateFunc,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
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
			if err := resourceData.Set(string(UserData), HexStateFunc(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UserData), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData := spotinst.String(base64Encode(v))
				cluster.Compute.LaunchSpecification.SetUserData(userData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			var value *bool = nil
			if cluster.Compute != nil && cluster.Compute.LaunchSpecification != nil &&
				cluster.Compute.LaunchSpecification.AssociatePublicIpAddress != nil {

				value = cluster.Compute.LaunchSpecification.AssociatePublicIpAddress
			}

			if err := resourceData.Set(string(AssociatePublicIpAddress), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AssociatePublicIpAddress), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(AssociatePublicIpAddress)); ok {
				cluster.Compute.LaunchSpecification.SetAssociatePublicIpAddress(spotinst.Bool(v.(bool)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(AssociatePublicIpAddress)); ok {
				cluster.Compute.LaunchSpecification.SetAssociatePublicIpAddress(spotinst.Bool(v.(bool)))
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
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
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
			clusterWrapper := resourceObject.(*commons.ClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			var result []*aws.LoadBalancer = nil
			existingBalancers := cluster.Compute.LaunchSpecification.LoadBalancers

			if existingBalancers != nil && len(existingBalancers) > 0 {
				for _, balancer := range existingBalancers {
					result = append(result, balancer)
				}
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
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
var InstanceProfileArnRegex = regexp.MustCompile(`arn:aws:iam::\d{12}:instance-profile/?[a-zA-Z_0-9+=,.@\-_/]+`)

func HexStateFunc(v interface{}) string {
	switch s := v.(type) {
	case string:
		hash := sha1.Sum([]byte(s))
		return hex.EncodeToString(hash[:])
	default:
		return ""
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

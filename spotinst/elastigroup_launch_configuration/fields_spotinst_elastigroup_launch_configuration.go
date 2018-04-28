package elastigroup_launch_configuration

import (
	"fmt"
	"encoding/base64"
	"crypto/sha1"
	"encoding/hex"
	"regexp"

	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupSpotinstLaunchConfigurationResource() {
	fields := make(map[commons.FieldName]*commons.GenericField)
	var readFailurePattern = "launch configuration failed reading field %s - %#v"

	//if lspec.ShutdownScript != nil && spotinst.StringValue(lspec.ShutdownScript) != "" {
	//	decodedShutdownScript, _ := base64.StdEncoding.DecodeString(spotinst.StringValue(lspec.ShutdownScript))
	//	result["shutdown_script"] = string(decodedShutdownScript)
	//} else {
	//	result["shutdown_script"] = ""
	//}

	fields[ImageId] = commons.NewGenericField(
		ImageId,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.ImageID != nil {
				value = elastigroup.Compute.LaunchSpecification.ImageID
			}
			if err := resourceData.Set(string(ImageId), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(ImageId), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(ImageId)).(string); ok && v != "" {
				elastigroup.Compute.LaunchSpecification.SetImageId(spotinst.String(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(ImageId)).(string); ok && v != "" {
				elastigroup.Compute.LaunchSpecification.SetImageId(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fields[IamInstanceProfile] = commons.NewGenericField(
		IamInstanceProfile,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
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
				return fmt.Errorf(readFailurePattern, string(IamInstanceProfile), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
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
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
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

	fields[KeyName] = commons.NewGenericField(
		KeyName,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.KeyPair != nil {
				value = elastigroup.Compute.LaunchSpecification.KeyPair
			}
			if err := resourceData.Set(string(KeyName), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(KeyName), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(KeyName)).(string); ok && v != "" {
				elastigroup.Compute.LaunchSpecification.SetKeyPair(spotinst.String(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(KeyName)).(string); ok && v != "" {
				elastigroup.Compute.LaunchSpecification.SetKeyPair(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fields[SecurityGroups] = commons.NewGenericField(
		SecurityGroups,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value []string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.SecurityGroupIDs != nil {
				value = elastigroup.Compute.LaunchSpecification.SecurityGroupIDs
			}
			if err := resourceData.Set(string(SecurityGroups), value); err != nil {
				return fmt.Errorf(readFailurePattern, string(SecurityGroups), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(SecurityGroups)).([]interface{}); ok {
				ids := make([]string, len(v))
				for i, j := range v {
					ids[i] = j.(string)
				}
				elastigroup.Compute.LaunchSpecification.SetSecurityGroupIDs(ids)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
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

	fields[UserData] = commons.NewGenericField(
		UserData,
		&schema.Schema{
			Type:      schema.TypeString,
			Optional:  true,
			StateFunc: hexStateFunc,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value = ""
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.UserData != nil {

				userData := elastigroup.Compute.LaunchSpecification.UserData
				if spotinst.StringValue(userData) != "" {
					decodedUserData, _ := base64.StdEncoding.DecodeString(spotinst.StringValue(userData))
					value = string(decodedUserData)
				}
			}
			if err := resourceData.Set(string(UserData), value); err != nil {
				return fmt.Errorf(readFailurePattern, string(UserData), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData := spotinst.String(base64Encode(v))
				elastigroup.Compute.LaunchSpecification.SetUserData(userData)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var userData *string = nil
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData = spotinst.String(base64Encode(v))
			}
			elastigroup.Compute.LaunchSpecification.SetUserData(userData)
			return nil
		},
		nil,
	)

	fields[EnableMonitoring] = commons.NewGenericField(
		EnableMonitoring,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *bool = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Monitoring != nil {
				value = elastigroup.Compute.LaunchSpecification.Monitoring
			}
			if err := resourceData.Set(string(EnableMonitoring), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(EnableMonitoring), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(EnableMonitoring)).(bool); ok {
				elastigroup.Compute.LaunchSpecification.SetMonitoring(spotinst.Bool(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(EnableMonitoring)).(bool); ok {
				elastigroup.Compute.LaunchSpecification.SetMonitoring(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fields[EbsOptimized] = commons.NewGenericField(
		EbsOptimized,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *bool = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.EBSOptimized != nil {
				value = elastigroup.Compute.LaunchSpecification.EBSOptimized
			}
			if err := resourceData.Set(string(EbsOptimized), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(EbsOptimized), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(EbsOptimized)).(bool); ok {
				elastigroup.Compute.LaunchSpecification.SetEBSOptimized(spotinst.Bool(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(EbsOptimized)).(bool); ok {
				elastigroup.Compute.LaunchSpecification.SetEBSOptimized(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fields[PlacementTenancy] = commons.NewGenericField(
		PlacementTenancy,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Tenancy != nil {
				value = elastigroup.Compute.LaunchSpecification.Tenancy
			}
			if err := resourceData.Set(string(PlacementTenancy), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(PlacementTenancy), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(PlacementTenancy)).(string); ok && v != "" {
				tenancy := spotinst.String(v)
				elastigroup.Compute.LaunchSpecification.SetTenancy(tenancy)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var tenancy *string = nil
			if v, ok := resourceData.Get(string(PlacementTenancy)).(string); ok && v != "" {
				tenancy = spotinst.String(v)
			}
			elastigroup.Compute.LaunchSpecification.SetTenancy(tenancy)
			return nil
		},
		nil,
	)

	commons.ElastigroupLaunchConfigurationResource = commons.NewGenericCachedResource(
		string(commons.ElastigroupLaunchConfiguration),
		fields)
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
var InstanceProfileArnRegex = regexp.MustCompile(`arn:aws:iam::\d{12}:instance-profile/?[a-zA-Z_0-9+=,.@\-_/]+`)

func hexStateFunc(v interface{}) string {
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
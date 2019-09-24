package ocean_aws_launch_spec

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform/helper/hashcode"
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

	fieldsMap[OceanID] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		OceanID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			launchSpec.SetOceanId(spotinst.String(resourceData.Get(string(OceanID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			launchSpec.SetImageId(spotinst.String(resourceData.Get(string(ImageID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			launchSpec.SetImageId(spotinst.String(resourceData.Get(string(ImageID)).(string)))
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(UserData)).(string); ok && v != "" {
				userData := spotinst.String(base64Encode(v))
				launchSpec.SetUserData(userData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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

	fieldsMap[SecurityGroups] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		SecurityGroups,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var result []string = nil

			if launchSpec != nil && launchSpec.SecurityGroupIDs != nil {
				securityGroupIds := launchSpec.SecurityGroupIDs

				for _, securityGroupId := range securityGroupIds {
					securityGroupIdStr := spotinst.StringValue(securityGroupId)
					result = append(result, securityGroupIdStr)
				}
			}
			resourceData.Set(string(SecurityGroups), result)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var result []*string = nil

			if value, ok := resourceData.GetOk(string(SecurityGroups)); ok {
				if value != nil && len(value.([]interface{})) > 0 {
					for _, v := range value.([]interface{}) {
						result = append(result, spotinst.String(v.(string)))
					}
					launchSpec.SetSecurityGroupIDs(result)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error { //
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var securityGroupIds []*string = nil

			if value, ok := resourceData.GetOk(string(SecurityGroups)); ok {
				for _, v := range value.([]interface{}) {
					securityGroupIds = append(securityGroupIds, spotinst.String(v.(string)))
				}
			}

			launchSpec.SetSecurityGroupIDs(securityGroupIds)
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
			Set: hashKV,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
			Set: hashKV,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func hashKV(v interface{}) int {
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

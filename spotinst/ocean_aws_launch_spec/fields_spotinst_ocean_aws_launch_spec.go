package ocean_aws_launch_spec

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
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
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
	fieldsMap[SecurityGroups] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		SecurityGroups,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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

	fieldsMap[SubnetIDs] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		SubnetIDs,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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

	fieldsMap[RootVolumeSize] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		RootVolumeSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(RootVolumeSize)).(int); ok && v > 0 {
				launchSpec.SetRootVolumeSize(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var value *int = nil
			if v, ok := resourceData.Get(string(RootVolumeSize)).(int); ok && v > 0 {
				value = spotinst.Int(v)
			}
			launchSpec.SetRootVolumeSize(value)
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

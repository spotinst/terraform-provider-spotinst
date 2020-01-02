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
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
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
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func hashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(AttributeKey)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(AttributeValue)].(string)))
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

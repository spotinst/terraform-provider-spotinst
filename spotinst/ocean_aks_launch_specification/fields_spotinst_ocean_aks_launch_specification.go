package ocean_aks_launch_specification

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ResourceGroupName] = commons.NewGenericField(
		commons.OceanAKSLaunchSpecification,
		ResourceGroupName,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification.ResourceGroupName != nil {
				value = cluster.VirtualNodeGroupTemplate.LaunchSpecification.ResourceGroupName
			}
			if err := resourceData.Set(string(ResourceGroupName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ResourceGroupName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(ResourceGroupName)).(string); ok && v != "" {
				resourceGroupName := spotinst.String(v)
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetResourceGroupName(resourceGroupName)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if v, ok := resourceData.Get(string(ResourceGroupName)).(string); ok && v != "" {
				resourceGroupName := spotinst.String(v)
				value = resourceGroupName
			}
			cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetResourceGroupName(value)
			return nil
		},
		nil,
	)

	fieldsMap[CustomData] = commons.NewGenericField(
		commons.OceanAKSLaunchSpecification,
		CustomData,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				// check to make sure nil SHA isn't being passed from somewhere upstream
				if (old == "da39a3ee5e6b4b0d3255bfef95601890afd80709" && new == "") ||
					(old == "" && new == "da39a3ee5e6b4b0d3255bfef95601890afd80709") {
					return true
				}
				return false
			},
			StateFunc: Base64StateFunc,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification.CustomData != nil {
				value = cluster.VirtualNodeGroupTemplate.LaunchSpecification.CustomData
			}
			if err := resourceData.Set(string(CustomData), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CustomData), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(CustomData)).(string); ok && v != "" {
				customData := spotinst.String(base64Encode(v))
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetCustomData(customData)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil
			if v, ok := resourceData.Get(string(CustomData)).(string); ok && v != "" {
				customData := spotinst.String(base64Encode(v))
				value = customData
			}
			cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetCustomData(value)
			return nil
		},
		nil,
	)

	fieldsMap[MaxPods] = commons.NewGenericField(
		commons.OceanAKSLaunchSpecification,
		MaxPods,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *int = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification.MaxPods != nil {
				value = cluster.VirtualNodeGroupTemplate.LaunchSpecification.MaxPods
			}
			if err := resourceData.Set(string(MaxPods), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MaxPods), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(MaxPods)).(int); ok && v > 0 {
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetMaxPods(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetMaxPods(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.Get(string(MaxPods)).(int); ok && v > 0 {
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetMaxPods(spotinst.Int(v))
			} else {
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetMaxPods(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[ManagedServiceIdentity] = commons.NewGenericField(
		commons.OceanAKSLaunchSpecification,
		ManagedServiceIdentity,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ManagedServiceIdentityResourceGroupName): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(ManagedServiceIdentityName): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value []interface{}
			if cluster.VirtualNodeGroupTemplate != nil &&
				cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil &&
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.ManagedServiceIdentities != nil {
				value = flattenManagedServiceIdentities(cluster.VirtualNodeGroupTemplate.LaunchSpecification.ManagedServiceIdentities)
			}
			if err := resourceData.Set(string(ManagedServiceIdentity), value); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(ManagedServiceIdentity), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(ManagedServiceIdentity)); ok {
				if msis, err := expandManagedServiceIdentities(v); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetManagedServiceIdentities(msis)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value []*azure.ManagedServiceIdentity
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil &&
				cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil &&
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.ManagedServiceIdentities != nil {
				if v, ok := resourceData.GetOk(string(ManagedServiceIdentity)); ok {
					if msis, err := expandManagedServiceIdentities(v); err != nil {
						return err
					} else {
						value = msis
					}
				}
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetManagedServiceIdentities(value)
			}
			return nil
		},
		nil,
	)

	fieldsMap[Tag] = commons.NewGenericField(
		commons.OceanAKSLaunchSpecification,
		Tag,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(TagKey): {
						Type:     schema.TypeString,
						Optional: true,
					},
					string(TagValue): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value []interface{} = nil

			if cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil &&
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.Tags != nil {
				value = flattenTags(cluster.VirtualNodeGroupTemplate.LaunchSpecification.Tags)
			}
			if err := resourceData.Set(string(Tag), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tag), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(Tag)); ok {
				if tags, err := expandTags(v); err != nil {
					return err
				} else {
					cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value []*azure.Tag = nil
			if cluster != nil && cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification.Tags != nil {
				if v, ok := resourceData.GetOk(string(Tag)); ok {
					if tags, err := expandTags(v); err != nil {
						return err
					} else {
						value = tags
					}
				}
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.SetTags(value)
			}
			return nil
		},
		nil,
	)

}

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

func expandTags(data interface{}) ([]*azure.Tag, error) {
	list := data.(*schema.Set).List()
	tags := make([]*azure.Tag, 0, len(list))

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

		tag := &azure.Tag{
			Key:   spotinst.String(attr[string(TagKey)].(string)),
			Value: spotinst.String(attr[string(TagValue)].(string)),
		}

		tags = append(tags, tag)
	}
	return tags, nil
}

func flattenTags(tags []*azure.Tag) []interface{} {
	result := make([]interface{}, 0, len(tags))
	for _, tag := range tags {
		m := make(map[string]interface{})
		m[string(TagKey)] = spotinst.StringValue(tag.Key)
		m[string(TagValue)] = spotinst.StringValue(tag.Value)

		result = append(result, m)
	}
	return result
}

func expandManagedServiceIdentities(data interface{}) ([]*azure.ManagedServiceIdentity, error) {
	list := data.(*schema.Set).List()
	msis := make([]*azure.ManagedServiceIdentity, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		msis = append(msis, &azure.ManagedServiceIdentity{
			ResourceGroupName: spotinst.String(attr[string(ManagedServiceIdentityResourceGroupName)].(string)),
			Name:              spotinst.String(attr[string(ManagedServiceIdentityName)].(string)),
		})
	}
	return msis, nil
}

func flattenManagedServiceIdentities(msis []*azure.ManagedServiceIdentity) []interface{} {
	result := make([]interface{}, 0, len(msis))
	for _, msi := range msis {
		m := make(map[string]interface{})
		m[string(ManagedServiceIdentityResourceGroupName)] = spotinst.StringValue(msi.ResourceGroupName)
		m[string(ManagedServiceIdentityName)] = spotinst.StringValue(msi.Name)
		result = append(result, m)
	}
	return result
}

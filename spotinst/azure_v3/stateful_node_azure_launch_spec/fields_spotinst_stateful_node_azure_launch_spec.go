package stateful_node_azure_launch_spec

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Tags] = commons.NewGenericField(
		commons.StatefulNodeAzureLaunchSpecification,
		Tags,
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
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []interface{} = nil

			if st.Compute != nil && st.Compute.LaunchSpecification != nil &&
				st.Compute.LaunchSpecification.Tags != nil {
				value = flattenTags(st.Compute.LaunchSpecification.Tags)
			}
			if err := resourceData.Set(string(Tags), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tags), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(v); err != nil {
					return err
				} else {
					st.Compute.LaunchSpecification.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []*azurev3.Tag = nil
			if st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.Tags != nil {
				if v, ok := resourceData.GetOk(string(Tags)); ok {
					if tags, err := expandTags(v); err != nil {
						return err
					} else {
						value = tags
					}
				}
				st.Compute.LaunchSpecification.SetTags(value)
			}
			return nil
		},
		nil,
	)

	fieldsMap[ManagedServiceIdentities] = commons.NewGenericField(
		commons.StatefulNodeAzureLaunchSpecification,
		ManagedServiceIdentities,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Name): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(ResourceGroupName): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []interface{}
			if st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.ManagedServiceIdentities != nil {
				value = flattenManagedServiceIdentities(st.Compute.LaunchSpecification.ManagedServiceIdentities)
			}
			if err := resourceData.Set(string(ManagedServiceIdentities), value); err != nil {
				return fmt.Errorf(commons.FailureFieldReadPattern, string(ManagedServiceIdentities), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(ManagedServiceIdentities)); ok {
				if msis, err := expandManagedServiceIdentities(v); err != nil {
					return err
				} else {
					st.Compute.LaunchSpecification.SetManagedServiceIdentities(msis)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []*azurev3.ManagedServiceIdentity
			if st != nil && st.Compute != nil &&
				st.Compute.LaunchSpecification != nil &&
				st.Compute.LaunchSpecification.ManagedServiceIdentities != nil {
				if v, ok := resourceData.GetOk(string(ManagedServiceIdentities)); ok {
					if msis, err := expandManagedServiceIdentities(v); err != nil {
						return err
					} else {
						value = msis
					}
				}
				st.Compute.LaunchSpecification.SetManagedServiceIdentities(value)
			}
			return nil
		},
		nil,
	)

}

func expandTags(data interface{}) ([]*azurev3.Tag, error) {
	list := data.(*schema.Set).List()
	tags := make([]*azurev3.Tag, 0, len(list))

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

		tag := &azurev3.Tag{
			TagKey:   spotinst.String(attr[string(TagKey)].(string)),
			TagValue: spotinst.String(attr[string(TagValue)].(string)),
		}

		tags = append(tags, tag)
	}
	return tags, nil
}

func flattenTags(tags []*azurev3.Tag) []interface{} {
	result := make([]interface{}, 0, len(tags))
	for _, tag := range tags {
		m := make(map[string]interface{})
		m[string(TagKey)] = spotinst.StringValue(tag.TagKey)
		m[string(TagValue)] = spotinst.StringValue(tag.TagValue)

		result = append(result, m)
	}
	return result
}

func expandManagedServiceIdentities(data interface{}) ([]*azurev3.ManagedServiceIdentity, error) {
	list := data.(*schema.Set).List()
	msis := make([]*azurev3.ManagedServiceIdentity, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		msis = append(msis, &azurev3.ManagedServiceIdentity{
			ResourceGroupName: spotinst.String(attr[string(ResourceGroupName)].(string)),
			Name:              spotinst.String(attr[string(Name)].(string)),
		})
	}
	return msis, nil
}

func flattenManagedServiceIdentities(msis []*azurev3.ManagedServiceIdentity) []interface{} {
	result := make([]interface{}, 0, len(msis))
	for _, msi := range msis {
		m := make(map[string]interface{})
		m[string(ResourceGroupName)] = spotinst.StringValue(msi.ResourceGroupName)
		m[string(Name)] = spotinst.StringValue(msi.Name)
		result = append(result, m)
	}
	return result
}

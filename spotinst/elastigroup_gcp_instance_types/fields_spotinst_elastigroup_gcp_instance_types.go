package elastigroup_gcp_instance_types

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[OnDemand] = commons.NewGenericField(
		commons.ElastigroupGCPInstanceType,
		OnDemand,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.InstanceTypes != nil &&
				elastigroup.Compute.InstanceTypes.OnDemand != nil {
				value = elastigroup.Compute.InstanceTypes.OnDemand
			}
			if err := resourceData.Set(string(OnDemand), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OnDemand), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(OnDemand)).(string); ok && v != "" {
				elastigroup.Compute.InstanceTypes.SetOnDemand(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(OnDemand)).(string); ok && v != "" {
				elastigroup.Compute.InstanceTypes.SetOnDemand(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Preemptible] = commons.NewGenericField(
		commons.ElastigroupGCPInstanceType,
		Preemptible,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []string
			if elastigroup.Compute != nil && elastigroup.Compute.InstanceTypes != nil &&
				elastigroup.Compute.InstanceTypes.Preemptible != nil {
				result = append(result, elastigroup.Compute.InstanceTypes.Preemptible...)
			}
			if err := resourceData.Set(string(Preemptible), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Preemptible), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Preemptible)); ok {
				prempts := v.([]interface{})
				premptTypes := make([]string, len(prempts))
				for i, j := range prempts {
					premptTypes[i] = j.(string)
				}
				elastigroup.Compute.InstanceTypes.SetPreemptible(premptTypes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Preemptible)); ok {
				prempts := v.([]interface{})
				premptTypes := make([]string, len(prempts))
				for i, j := range prempts {
					premptTypes[i] = j.(string)
				}
				elastigroup.Compute.InstanceTypes.SetPreemptible(premptTypes)
			}
			return nil
		},
		nil,
	)

	fieldsMap[Custom] = commons.NewGenericField(
		commons.ElastigroupGCPInstanceType,
		Custom,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MemoryGiB): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(VCPU): {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup.Compute != nil && elastigroup.Compute.InstanceTypes != nil &&
				elastigroup.Compute.InstanceTypes.Custom != nil {
				customInstances := elastigroup.Compute.InstanceTypes.Custom
				result = flattenCustom(customInstances)
			}
			if result != nil {
				if err := resourceData.Set(string(Custom), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Custom), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Custom)); ok {
				if customInstances, err := expandCustom(v); err != nil {
					return err
				} else {
					elastigroup.Compute.InstanceTypes.SetCustom(customInstances)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Custom)); ok {
				if customInstances, err := expandCustom(v); err != nil {
					return err
				} else {
					elastigroup.Compute.InstanceTypes.SetCustom(customInstances)
				}
			}
			return nil
		},
		nil,
	)

}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func flattenCustom(customInstances []*gcp.CustomInstance) []interface{} {
	result := make([]interface{}, 0, len(customInstances))
	for _, instance := range customInstances {
		m := make(map[string]interface{})
		m[string(MemoryGiB)] = spotinst.IntValue(instance.MemoryGiB)
		m[string(VCPU)] = spotinst.IntValue(instance.VCPU)
		result = append(result, m)
	}

	return []interface{}{result}
}

func expandCustom(data interface{}) ([]*gcp.CustomInstance, error) {
	list := data.([]interface{})
	if list != nil && list[0] != nil {
		instances := make([]*gcp.CustomInstance, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			instance := &gcp.CustomInstance{}

			if v, ok := m[string(MemoryGiB)].(int); ok && v >= 10 {
				instance.SetMemoryGiB(spotinst.Int(v))
			}

			if v, ok := m[string(VCPU)].(int); ok && v > 0 {
				instance.SetVCPU(spotinst.Int(v))
			}
			instances = append(instances, instance)
		}
		return instances, nil
	}
	return nil, nil
}

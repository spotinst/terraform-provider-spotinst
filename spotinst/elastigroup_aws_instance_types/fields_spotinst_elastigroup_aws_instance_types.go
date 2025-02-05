package elastigroup_aws_instance_types

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[OnDemand] = commons.NewGenericField(
		commons.ElastigroupAWSInstanceType,
		OnDemand,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(OnDemand)).(string); ok && v != "" {
				elastigroup.Compute.InstanceTypes.SetOnDemand(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if v, ok := resourceData.GetOkExists(string(OnDemand)); ok && v != "" {
				if od, ok := v.(string); ok && od != "" {
					value = spotinst.String(od)
				}
			}
			elastigroup.Compute.InstanceTypes.SetOnDemand(value)
			return nil
		},
		nil,
	)

	fieldsMap[OnDemandTypes] = commons.NewGenericField(
		commons.ElastigroupAWSInstanceType,
		OnDemandTypes,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []string
			if elastigroup.Compute != nil && elastigroup.Compute.InstanceTypes != nil &&
				elastigroup.Compute.InstanceTypes.OnDemandTypes != nil {
				result = append(result, elastigroup.Compute.InstanceTypes.OnDemandTypes...)
			}
			if err := resourceData.Set(string(OnDemandTypes), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OnDemandTypes), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(OnDemandTypes)); ok {
				onDemand := v.([]interface{})
				onDemandTypes := make([]string, len(onDemand))
				for i, j := range onDemand {
					onDemandTypes[i] = j.(string)
				}
				elastigroup.Compute.InstanceTypes.SetOnDemandTypes(onDemandTypes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var onDemandTypes []string = nil
			if v, ok := resourceData.GetOk(string(OnDemandTypes)); ok {
				odTypes := v.([]interface{})
				onDemandTypes = make([]string, len(odTypes))
				for i, v := range odTypes {
					onDemandTypes[i] = v.(string)
				}
			}
			elastigroup.Compute.InstanceTypes.SetOnDemandTypes(onDemandTypes)
			return nil
		},
		nil,
	)

	fieldsMap[Spot] = commons.NewGenericField(
		commons.ElastigroupAWSInstanceType,
		Spot,
		&schema.Schema{
			Type:          schema.TypeList,
			Optional:      true,
			Elem:          &schema.Schema{Type: schema.TypeString},
			ConflictsWith: []string{string(ResourceRequirements)},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.InstanceTypes != nil &&
				elastigroup.Compute.InstanceTypes.Spot != nil {
				result = elastigroup.Compute.InstanceTypes.Spot
			}
			if err := resourceData.Set(string(Spot), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Spot), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if value, ok := resourceData.GetOk(string(Spot)); ok && value != nil {
				if spots, err := expandSpotInstanceTypes(value); err != nil {
					return err
				} else {
					elastigroup.Compute.InstanceTypes.SetSpot(spots)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []string = nil
			if v, ok := resourceData.GetOk(string(Spot)); ok && v != nil {
				if spots, err := expandSpotInstanceTypes(v); err != nil {
					return err
				} else {
					value = spots
				}
			}
			elastigroup.Compute.InstanceTypes.SetSpot(value)
			return nil
		},
		nil,
	)

	fieldsMap[PreferredSpot] = commons.NewGenericField(
		commons.ElastigroupAWSInstanceType,
		PreferredSpot,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []string
			if elastigroup.Compute != nil && elastigroup.Compute.InstanceTypes != nil &&
				elastigroup.Compute.InstanceTypes.PreferredSpot != nil {
				result = append(result, elastigroup.Compute.InstanceTypes.PreferredSpot...)
			}
			if err := resourceData.Set(string(PreferredSpot), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PreferredSpot), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(PreferredSpot)); ok {
				spots := v.([]interface{})
				spotTypes := make([]string, len(spots))
				for i, j := range spots {
					spotTypes[i] = j.(string)
				}
				elastigroup.Compute.InstanceTypes.SetPreferredSpot(spotTypes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var spotTypes []string = nil
			if v, ok := resourceData.GetOk(string(PreferredSpot)); ok {
				rawSpotTypes := v.([]interface{})
				spotTypes = make([]string, len(rawSpotTypes))
				for i, v := range rawSpotTypes {
					spotTypes[i] = v.(string)
				}
			}
			elastigroup.Compute.InstanceTypes.SetPreferredSpot(spotTypes)
			return nil
		},
		nil,
	)

	fieldsMap[InstanceTypeWeights] = commons.NewGenericField(
		commons.ElastigroupAWSInstanceType,
		InstanceTypeWeights,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(InstanceType): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(Weight): {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(InstanceTypeWeights)); ok && v != "" {
				if weights, err := expandAWSGroupInstanceTypeWeights(v); err != nil {
					return err
				} else {
					elastigroup.Compute.InstanceTypes.SetWeights(weights)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var weightsToAdd []*aws.InstanceTypeWeight = nil
			if v, ok := resourceData.GetOk(string(InstanceTypeWeights)); ok {
				if weights, err := expandAWSGroupInstanceTypeWeights(v); err != nil {
					return err
				} else {
					weightsToAdd = weights
				}
			}
			elastigroup.Compute.InstanceTypes.SetWeights(weightsToAdd)
			return nil
		},
		nil,
	)

	fieldsMap[ResourceRequirements] = commons.NewGenericField(
		commons.ElastigroupAWSInstanceType,
		ResourceRequirements,
		&schema.Schema{
			Type:          schema.TypeList,
			Optional:      true,
			ConflictsWith: []string{string(Spot)},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(RequiredGpuMaximum): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(RequiredGpuMinimum): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(RequiredMemoryMaximum): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(RequiredMemoryMinimum): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(RequiredVCpuMaximum): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(RequiredVCpuMinimum): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(ExcludedInstanceFamilies): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(ExcludedInstanceGenerations): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(ExcludedInstanceTypes): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil

			if elastigroup.Compute != nil && elastigroup.Compute.InstanceTypes != nil &&
				elastigroup.Compute.InstanceTypes.ResourceRequirements != nil {
				result = flattenResourceRequirements(elastigroup.Compute.InstanceTypes.ResourceRequirements)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(ResourceRequirements), result); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(ResourceRequirements), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(ResourceRequirements)); ok {
				if requirements, err := expandResourceRequirements(v, false); err != nil {
					return err
				} else {
					elastigroup.Compute.InstanceTypes.SetResourceRequirements(requirements)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.ResourceRequirements = nil

			if v, ok := resourceData.GetOk(string(ResourceRequirements)); ok {
				if requirements, err := expandResourceRequirements(v, true); err != nil {
					return err
				} else {
					value = requirements
				}
			}
			if elastigroup.Compute.InstanceTypes == nil {
				elastigroup.Compute.InstanceTypes = &aws.InstanceTypes{}
			}
			elastigroup.Compute.InstanceTypes.SetResourceRequirements(value)
			return nil
		},
		nil,
	)
}

func expandAWSGroupInstanceTypeWeights(data interface{}) ([]*aws.InstanceTypeWeight, error) {
	list := data.(*schema.Set).List()
	weights := make([]*aws.InstanceTypeWeight, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(InstanceType)]; !ok {
			return nil, errors.New("[ERROR] Invalid instance type weight: instance_type missing")
		}

		if _, ok := attr[string(Weight)]; !ok {
			return nil, errors.New("[ERROR] Invalid instance type weight: weight missing")
		}
		instanceWeight := &aws.InstanceTypeWeight{}
		instanceWeight.SetInstanceType(spotinst.String(attr[string(InstanceType)].(string)))
		instanceWeight.SetWeight(spotinst.Int(attr[string(Weight)].(int)))
		//log.Printf("Group instance type weight configuration: %s", stringutil.Stringify(instanceWeight))
		weights = append(weights, instanceWeight)
	}
	return weights, nil
}

func expandResourceRequirements(data interface{}, nullify bool) (*aws.ResourceRequirements, error) {
	requirements := &aws.ResourceRequirements{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return requirements, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(ExcludedInstanceFamilies)]; ok {
		instanceFamilies, err := expandResourceRequirementsList(v)
		if err != nil {
			return nil, err
		}
		if instanceFamilies != nil && len(instanceFamilies) > 0 {
			requirements.SetExcludedInstanceFamilies(instanceFamilies)
		} else {
			if nullify {
				requirements.SetExcludedInstanceFamilies(nil)
			}
		}
	}

	if v, ok := m[string(ExcludedInstanceGenerations)]; ok {
		instanceGenerations, err := expandResourceRequirementsList(v)
		if err != nil {
			return nil, err
		}
		if instanceGenerations != nil && len(instanceGenerations) > 0 {
			requirements.SetExcludedInstanceGenerations(instanceGenerations)
		} else {
			if nullify {
				requirements.SetExcludedInstanceGenerations(nil)
			}
		}
	}

	if v, ok := m[string(ExcludedInstanceTypes)]; ok {
		instanceTypes, err := expandResourceRequirementsList(v)
		if err != nil {
			return nil, err
		}
		if instanceTypes != nil && len(instanceTypes) > 0 {
			requirements.SetExcludedInstanceTypes(instanceTypes)
		} else {
			if nullify {
				requirements.SetExcludedInstanceTypes(nil)
			}
		}
	}

	requiredGpu := &aws.RequiredGpu{}
	requirements.SetRequiredGpu(requiredGpu)
	if v, ok := m[string(RequiredGpuMinimum)].(int); ok && v >= 1 {
		requirements.RequiredGpu.SetMinimum(spotinst.Int(v))
	} else {
		requirements.RequiredGpu.SetMinimum(nil)
	}

	if v, ok := m[string(RequiredGpuMaximum)].(int); ok && v >= 1 {
		requirements.RequiredGpu.SetMaximum(spotinst.Int(v))
	} else {
		requirements.RequiredGpu.SetMaximum(nil)
	}

	requiredVCpu := &aws.RequiredVCpu{}
	requirements.SetRequiredVCpu(requiredVCpu)
	if v, ok := m[string(RequiredVCpuMinimum)].(int); ok && v >= 1 {
		requirements.RequiredVCpu.SetMinimum(spotinst.Int(v))
	} else {
		requirements.RequiredVCpu.SetMinimum(nil)
	}

	if v, ok := m[string(RequiredVCpuMaximum)].(int); ok && v >= 1 {
		requirements.RequiredVCpu.SetMaximum(spotinst.Int(v))
	} else {
		requirements.RequiredVCpu.SetMaximum(nil)
	}

	requiredMemory := &aws.RequiredMemory{}
	requirements.SetRequiredMemory(requiredMemory)
	if v, ok := m[string(RequiredMemoryMinimum)].(int); ok && v >= 1 {
		requirements.RequiredMemory.SetMinimum(spotinst.Int(v))
	} else {
		requirements.RequiredMemory.SetMinimum(nil)
	}

	if v, ok := m[string(RequiredMemoryMaximum)].(int); ok && v >= 1 {
		requirements.RequiredMemory.SetMaximum(spotinst.Int(v))
	} else {
		requirements.RequiredMemory.SetMaximum(nil)
	}

	return requirements, nil
}

func expandResourceRequirementsList(data interface{}) ([]string, error) {
	list := data.(*schema.Set).List()
	result := make([]string, 0, len(list))

	for _, v := range list {
		if requirementsList, ok := v.(string); ok && requirementsList != "" {
			result = append(result, requirementsList)
		}
	}
	return result, nil
}

func expandSpotInstanceTypes(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if spotInstance, ok := v.(string); ok && spotInstance != "" {
			result = append(result, spotInstance)
		}
	}
	return result, nil
}

func flattenResourceRequirements(requirements *aws.ResourceRequirements) []interface{} {
	var out []interface{}

	if requirements != nil {
		result := make(map[string]interface{})
		value := spotinst.Int(-1)
		result[string(RequiredGpuMinimum)] = value
		result[string(RequiredGpuMaximum)] = value

		if requirements.RequiredGpu != nil {
			if requirements.RequiredGpu.Minimum != nil {
				result[string(RequiredGpuMinimum)] = spotinst.IntValue(requirements.RequiredGpu.Minimum)
			}
			if requirements.RequiredGpu.Maximum != nil {
				result[string(RequiredGpuMaximum)] = spotinst.IntValue(requirements.RequiredGpu.Maximum)
			}
		}
		result[string(RequiredMemoryMinimum)] = spotinst.IntValue(requirements.RequiredMemory.Minimum)
		result[string(RequiredMemoryMaximum)] = spotinst.IntValue(requirements.RequiredMemory.Maximum)
		result[string(RequiredVCpuMinimum)] = spotinst.IntValue(requirements.RequiredVCpu.Minimum)
		result[string(RequiredVCpuMaximum)] = spotinst.IntValue(requirements.RequiredVCpu.Maximum)

		if requirements.ExcludedInstanceFamilies != nil {
			result[string(ExcludedInstanceFamilies)] = requirements.ExcludedInstanceFamilies
		}

		if requirements.ExcludedInstanceGenerations != nil {
			result[string(ExcludedInstanceGenerations)] = requirements.ExcludedInstanceGenerations
		}

		if requirements.ExcludedInstanceTypes != nil {
			result[string(ExcludedInstanceTypes)] = requirements.ExcludedInstanceTypes
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

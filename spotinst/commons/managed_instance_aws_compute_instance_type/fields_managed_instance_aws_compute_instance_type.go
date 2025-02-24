package managed_instance_aws_compute_instance_type

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Product] = commons.NewGenericField(
		commons.ManagedInstanceAWSComputeInstanceType,
		Product,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *string = nil
			if managedInstance.Compute != nil && managedInstance.Compute.Product != nil {
				value = managedInstance.Compute.Product
			}
			if err := resourceData.Set(string(Product), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Product), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			managedInstance.Compute.SetProduct(spotinst.String(resourceData.Get(string(Product)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(Product))
			return err
		},
		nil,
	)

	fieldsMap[Types] = commons.NewGenericField(
		commons.ManagedInstanceAWSComputeInstanceType,
		Types,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var result []string
			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.InstanceTypes.Types != nil {
				result = append(result, managedInstance.Compute.LaunchSpecification.InstanceTypes.Types...)
			}
			if err := resourceData.Set(string(Types), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Types), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.GetOk(string(Types)); ok {
				instances := v.([]interface{})
				instanceTypes := make([]string, len(instances))
				for i, j := range instances {
					instanceTypes[i] = j.(string)
				}
				managedInstance.Compute.LaunchSpecification.InstanceTypes.SetInstanceTypes(instanceTypes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var instanceTypes []string = nil
			if v, ok := resourceData.GetOk(string(Types)); ok {
				instances := v.([]interface{})
				instanceTypes = make([]string, len(instances))
				for i, v := range instances {
					instanceTypes[i] = v.(string)
				}
			}
			managedInstance.Compute.LaunchSpecification.InstanceTypes.SetInstanceTypes(instanceTypes)
			return nil
		},
		nil,
	)

	fieldsMap[PreferredType] = commons.NewGenericField(
		commons.ManagedInstanceAWSComputeInstanceType,
		PreferredType,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *string = nil
			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.InstanceTypes.PreferredType != nil {
				value = managedInstance.Compute.LaunchSpecification.InstanceTypes.PreferredType
			}
			if err := resourceData.Set(string(PreferredType), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PreferredType), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.Get(string(PreferredType)).(string); ok && v != "" {
				tenancy := spotinst.String(v)
				if managedInstance.Compute.LaunchSpecification.InstanceTypes == nil {
					managedInstance.Compute.LaunchSpecification.InstanceTypes = new(aws.InstanceTypes)
				}
				managedInstance.Compute.LaunchSpecification.InstanceTypes.SetPreferredType(tenancy)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var tenancy *string = nil
			if v, ok := resourceData.Get(string(PreferredType)).(string); ok && v != "" {
				tenancy = spotinst.String(v)
			}
			if managedInstance.Compute.LaunchSpecification.InstanceTypes == nil {
				managedInstance.Compute.LaunchSpecification.InstanceTypes = new(aws.InstanceTypes)
			}
			managedInstance.Compute.LaunchSpecification.InstanceTypes.SetPreferredType(tenancy)
			return nil
		},
		nil,
	)

	fieldsMap[PreferredTypes] = commons.NewGenericField(
		commons.ManagedInstanceAWSComputeInstanceType,
		PreferredTypes,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var result []string = nil
			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification != nil &&
				managedInstance.Compute.LaunchSpecification.InstanceTypes.PreferredTypes != nil {
				result = managedInstance.Compute.LaunchSpecification.InstanceTypes.PreferredTypes
			}
			if err := resourceData.Set(string(PreferredTypes), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Types), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if value, ok := resourceData.GetOk(string(PreferredTypes)); ok && value != nil {
				if preferredTypes, err := expandStringList(value); err != nil {
					return err
				} else {
					managedInstance.Compute.LaunchSpecification.InstanceTypes.SetPreferredTypes(preferredTypes)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var preferredTypes []string = nil
			if value, ok := resourceData.GetOk(string(PreferredTypes)); ok && value != nil {
				if instances, err := expandStringList(value); err != nil {
					return err
				} else {
					preferredTypes = instances
				}
			}
			managedInstance.Compute.LaunchSpecification.InstanceTypes.SetPreferredTypes(preferredTypes)
			return nil
		},
		nil,
	)

	fieldsMap[ResourceRequirements] = commons.NewGenericField(
		commons.ManagedInstanceAWSComputeInstanceType,
		ResourceRequirements,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
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
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var result []interface{} = nil

			if managedInstance.Compute != nil && managedInstance.Compute.LaunchSpecification.InstanceTypes != nil &&
				managedInstance.Compute.LaunchSpecification.InstanceTypes.ResourceRequirements != nil {
				result = flattenResourceRequirements(managedInstance.Compute.LaunchSpecification.InstanceTypes.ResourceRequirements)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(ResourceRequirements), result); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(ResourceRequirements), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if v, ok := resourceData.GetOk(string(ResourceRequirements)); ok {
				if requirements, err := expandResourceRequirements(v, false); err != nil {
					return err
				} else {
					managedInstance.Compute.LaunchSpecification.InstanceTypes.SetResourceRequirements(requirements)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *aws.ResourceRequirements = nil

			if v, ok := resourceData.GetOk(string(ResourceRequirements)); ok {
				if requirements, err := expandResourceRequirements(v, true); err != nil {
					return err
				} else {
					value = requirements
				}
			}
			if managedInstance.Compute.LaunchSpecification.InstanceTypes == nil {
				managedInstance.Compute.LaunchSpecification.InstanceTypes = &aws.InstanceTypes{}
			}
			managedInstance.Compute.LaunchSpecification.InstanceTypes.SetResourceRequirements(value)
			return nil
		},
		nil,
	)

}

func expandResourceRequirements(data interface{}, nullify bool) (*aws.ResourceRequirements, error) {
	requirements := &aws.ResourceRequirements{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return requirements, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(ExcludedInstanceFamilies)]; ok {
		instanceFamilies, err := expandStringList(v)
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
		instanceGenerations, err := expandStringList(v)
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
		instanceTypes, err := expandStringList(v)
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

func expandStringList(data interface{}) ([]string, error) {
	list := data.(*schema.Set).List()
	result := make([]string, 0, len(list))

	for _, v := range list {
		if requirementsList, ok := v.(string); ok && requirementsList != "" {
			result = append(result, requirementsList)
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

package managed_instance_aws_compute_instance_type

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/managedinstance/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
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
			Required: true,
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

}

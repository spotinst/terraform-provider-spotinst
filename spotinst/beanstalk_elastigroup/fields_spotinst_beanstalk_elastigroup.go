package beanstalk_elastigroup

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.BeanstalkElastigroup,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		// onRead
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()
			var value *string = nil
			if beanstalkGroup.Name != nil {
				value = beanstalkGroup.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		// onCreate
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()
			beanstalkGroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		// onUpdate
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()
			beanstalkGroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Region] = commons.NewGenericField(
		commons.BeanstalkElastigroup,
		Region,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()
			var value *string = nil
			if beanstalkGroup.Region != nil {
				value = beanstalkGroup.Region
			}
			if err := resourceData.Set(string(Region), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Region), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()
			if v, ok := resourceData.GetOk(string(Region)); ok {
				beanstalkGroup.SetRegion(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(Region))
			return err
		},
		nil,
	)

	fieldsMap[Product] = commons.NewGenericField(
		commons.BeanstalkElastigroup,
		Product,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()
			var value *string = nil
			if beanstalkGroup.Compute != nil && beanstalkGroup.Compute.Product != nil {
				value = beanstalkGroup.Compute.Product
			}
			if err := resourceData.Set(string(Product), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Product), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()
			beanstalkGroup.Compute.SetProduct(spotinst.String(resourceData.Get(string(Product)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(Product))
			return err
		},
		nil,
	)

	fieldsMap[Minimum] = commons.NewGenericField(
		commons.BeanstalkElastigroup,
		Minimum,
		&schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()
			var value *int = nil
			if beanstalkGroup.Capacity != nil && beanstalkGroup.Capacity.Minimum != nil {
				value = beanstalkGroup.Capacity.Minimum
			}
			if err := resourceData.Set(string(Minimum), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Minimum), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()
			if v, ok := resourceData.Get(string(Minimum)).(int); ok && v >= 0 {
				beanstalkGroup.Capacity.SetMinimum(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()
			if v, ok := resourceData.Get(string(Minimum)).(int); ok && v >= 0 {
				beanstalkGroup.Capacity.SetMinimum(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Maximum] = commons.NewGenericField(
		commons.BeanstalkElastigroup,
		Maximum,
		&schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()
			var value *int = nil
			if beanstalkGroup.Capacity != nil && beanstalkGroup.Capacity.Maximum != nil {
				value = beanstalkGroup.Capacity.Maximum
			}
			if err := resourceData.Set(string(Maximum), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Maximum), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()

			if v, ok := resourceData.Get(string(Maximum)).(int); ok && v >= 0 {
				beanstalkGroup.Capacity.SetMaximum(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()
			if v, ok := resourceData.Get(string(Maximum)).(int); ok && v >= 0 {
				beanstalkGroup.Capacity.SetMaximum(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Target] = commons.NewGenericField(
		commons.BeanstalkElastigroup,
		Target,
		&schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()
			var value *int = nil
			if beanstalkGroup.Capacity != nil && beanstalkGroup.Capacity.Target != nil {
				value = beanstalkGroup.Capacity.Target
			}
			if err := resourceData.Set(string(Target), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Target), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()

			if v, ok := resourceData.Get(string(Target)).(int); ok && v >= 0 {
				beanstalkGroup.Capacity.SetTarget(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()
			if v, ok := resourceData.Get(string(Target)).(int); ok && v >= 0 {
				beanstalkGroup.Capacity.SetTarget(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[BeanstalkEnvironmentName] = commons.NewGenericField(
		commons.BeanstalkElastigroup,
		BeanstalkEnvironmentName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(BeanstalkEnvironmentName))
			return err
		},
		nil,
	)

	fieldsMap[SpotInstanceTypes] = commons.NewGenericField(
		commons.BeanstalkElastigroup,
		SpotInstanceTypes,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()

			var value []string = nil
			if beanstalkGroup.Compute != nil && beanstalkGroup.Compute.InstanceTypes != nil && beanstalkGroup.Compute.InstanceTypes.Spot != nil {
				value = beanstalkGroup.Compute.InstanceTypes.Spot
			}
			if err := resourceData.Set(string(SpotInstanceTypes), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SpotInstanceTypes), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()

			if v, ok := resourceData.GetOk(string(SpotInstanceTypes)); ok && v != nil {
				if spotTypes, err := expandElastigroupInstanceTypesList(v); err != nil {
					return err
				} else {
					beanstalkGroup.Compute.InstanceTypes.SetSpot(spotTypes)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			beanstalkWrapper := resourceObject.(*commons.BeanStalkElastigroupWrapper)
			beanstalkGroup := beanstalkWrapper.GetBeanstalkElastigroup()

			if v, ok := resourceData.GetOk(string(SpotInstanceTypes)); ok && v != nil {
				if spotTypes, err := expandElastigroupInstanceTypesList(v); err != nil {
					return err
				} else {
					beanstalkGroup.Compute.InstanceTypes.SetSpot(spotTypes)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[Maintenance] = commons.NewGenericField(
		commons.BeanstalkElastigroup,
		Maintenance,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.GetOk(string(Maintenance)); ok && v != nil {
				if v != "START" && v != "END" && v != "STATUS" {
					return fmt.Errorf("Error: Maintenace mode must be START, END, or STATUS")
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.GetOk(string(Maintenance)); ok && v != nil {
				if v != "START" && v != "END" && v != "STATUS" {
					return fmt.Errorf("Error: Maintenace mode must be START or END, or STATUS")
				}
			}
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Fields Expand
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandElastigroupInstanceTypesList(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))
	for _, str := range list {
		if typ, ok := str.(string); ok {
			result = append(result, typ)
		}
	}
	return result, nil
}

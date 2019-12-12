package elastigroup_azure

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.ElastigroupAzure,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Name != nil {
				value = elastigroup.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			elastigroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			elastigroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Region] = commons.NewGenericField(
		commons.ElastigroupAzure,
		Region,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Region != nil {
				value = elastigroup.Region
			}
			if err := resourceData.Set(string(Region), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Region), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Region)); ok {
				elastigroup.SetRegion(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Region)); ok {
				elastigroup.SetRegion(spotinst.String(v.(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Product] = commons.NewGenericField(
		commons.ElastigroupAzure,
		Product,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.Product != nil {
				value = elastigroup.Compute.Product
			}
			if err := resourceData.Set(string(Product), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Product), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			elastigroup.Compute.SetProduct(spotinst.String(resourceData.Get(string(Product)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(Product))
			return err
		},
		nil,
	)

	fieldsMap[ResourceGroupName] = commons.NewGenericField(
		commons.ElastigroupAzure,
		ResourceGroupName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.ResourceGroupName != nil {
				value = elastigroup.ResourceGroupName
			}
			if err := resourceData.Set(string(ResourceGroupName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ResourceGroupName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			elastigroup.SetResourceGroupName(spotinst.String(resourceData.Get(string(ResourceGroupName)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(ResourceGroupName))
			return err
		},
		nil,
	)

	fieldsMap[MinSize] = commons.NewGenericField(
		commons.ElastigroupAzure,
		MinSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if elastigroup.Capacity != nil && elastigroup.Capacity.Minimum != nil {
				value = elastigroup.Capacity.Minimum
			}
			if err := resourceData.Set(string(MinSize), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MinSize), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(MinSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMinimum(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(MinSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMinimum(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[MaxSize] = commons.NewGenericField(
		commons.ElastigroupAzure,
		MaxSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if elastigroup.Capacity != nil && elastigroup.Capacity.Maximum != nil {
				value = elastigroup.Capacity.Maximum
			}
			if err := resourceData.Set(string(MaxSize), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MaxSize), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(MaxSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMaximum(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(MaxSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMaximum(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[DesiredCapacity] = commons.NewGenericField(
		commons.ElastigroupAzure,
		DesiredCapacity,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if elastigroup.Capacity != nil && elastigroup.Capacity.Target != nil {
				value = elastigroup.Capacity.Target
			}
			if err := resourceData.Set(string(DesiredCapacity), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DesiredCapacity), err)
			}
			if v, ok := resourceData.Get(string(DesiredCapacity)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetTarget(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(DesiredCapacity)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetTarget(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(DesiredCapacity)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetTarget(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[UpdatePolicy] = commons.NewGenericField(
		commons.ElastigroupAzure,
		UpdatePolicy,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ShouldRoll): {
						Type:     schema.TypeBool,
						Required: true,
					},

					string(RollConfig): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(BatchSizePercentage): {
									Type:     schema.TypeInt,
									Required: true,
								},

								string(GracePeriod): {
									Type:     schema.TypeInt,
									Optional: true,
									Default:  -1,
								},

								string(HealthCheckType): {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		nil, nil, nil, nil,
	)

}

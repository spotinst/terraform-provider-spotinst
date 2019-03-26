package elastigroup_gcp

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[AvailabilityZones] = commons.NewGenericField(
		commons.ElastigroupGCP,
		AvailabilityZones,
		&schema.Schema{
			Type:       schema.TypeList,
			Optional:   true,
			Elem:       &schema.Schema{Type: schema.TypeString},
			Deprecated: "This field will soon be handled by Region in Subnets",
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []string
			if elastigroup.Compute != nil && elastigroup.Compute.AvailabilityZones != nil {
				values := elastigroup.Compute.AvailabilityZones
				for _, zones := range values {
					result = append(result, zones)
				}
			}
			if err := resourceData.Set(string(AvailabilityZones), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AvailabilityZones), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				zonesList := v.([]interface{})
				zones := make([]string, len(zonesList))
				for i, j := range zonesList {
					zones[i] = j.(string)
				}
				elastigroup.Compute.SetAvailabilityZones(zones)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				zonesList := v.([]interface{})
				zones := make([]string, len(zonesList))
				for i, j := range zonesList {
					zones[i] = j.(string)
				}
				elastigroup.Compute.SetAvailabilityZones(zones)
			}
			return nil
		},
		nil,
	)

	fieldsMap[Description] = commons.NewGenericField(
		commons.ElastigroupGCP,
		Description,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Description != nil {
				value = elastigroup.Description
			}
			if err := resourceData.Set(string(Description), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Description), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			elastigroup.SetDescription(spotinst.String(resourceData.Get(string(Description)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			elastigroup.SetDescription(spotinst.String(resourceData.Get(string(Description)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[HealthCheckGracePeriod] = commons.NewGenericField(
		commons.ElastigroupGCP,
		HealthCheckGracePeriod,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if elastigroup.Compute != nil && elastigroup.Compute.Health != nil &&
				elastigroup.Compute.Health.GracePeriod != nil {
				value = elastigroup.Compute.Health.GracePeriod
			}

			if err := resourceData.Set(string(HealthCheckGracePeriod), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(HealthCheckGracePeriod), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOkExists(string(HealthCheckGracePeriod)); ok && v != nil {
				if elastigroup.Compute.Health == nil {
					elastigroup.Compute.SetHealth(&gcp.Health{})
				}
				elastigroup.Compute.Health.SetGracePeriod(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if v, ok := resourceData.GetOkExists(string(HealthCheckGracePeriod)); ok && v != nil {
				if elastigroup.Compute.Health == nil {
					elastigroup.Compute.SetHealth(&gcp.Health{})
				}
				value = spotinst.Int(v.(int))
			}
			elastigroup.Compute.Health.SetGracePeriod(value)
			return nil
		},
		nil,
	)

	fieldsMap[MaxSize] = commons.NewGenericField(
		commons.ElastigroupGCP,
		MaxSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
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
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(MaxSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMaximum(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(MaxSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMaximum(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[MinSize] = commons.NewGenericField(
		commons.ElastigroupGCP,
		MinSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
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
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(MinSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMinimum(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(MinSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMinimum(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Name] = commons.NewGenericField(
		commons.ElastigroupGCP,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
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
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			elastigroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			elastigroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))

			return nil
		},
		nil,
	)

	fieldsMap[Subnets] = commons.NewGenericField(
		commons.ElastigroupGCP,
		Subnets,
		&schema.Schema{
			Type:             schema.TypeSet,
			Optional:         true,
			DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Region): &schema.Schema{
						Type:             schema.TypeString,
						Required:         true,
						DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
					},

					string(SubnetNames): &schema.Schema{
						Type:             schema.TypeList,
						Required:         true,
						Elem:             &schema.Schema{Type: schema.TypeString},
						DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Subnets)); ok {
				if subnets, err := expandSubnets(v); err != nil {
					return err
				} else {
					elastigroup.Compute.SetSubnets(subnets)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var subnetList []*gcp.Subnet = nil
			if value, ok := resourceData.GetOk(string(Subnets)); ok {
				if subnets, err := expandSubnets(value); err != nil {
					return err
				} else {
					subnetList = subnets
				}
			}
			elastigroup.Compute.SetSubnets(subnetList)
			return nil
		},
		nil,
	)

	fieldsMap[TargetCapacity] = commons.NewGenericField(
		commons.ElastigroupGCP,
		TargetCapacity,
		&schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *int = nil
			if elastigroup.Capacity != nil && elastigroup.Capacity.Target != nil {
				value = elastigroup.Capacity.Target
			}
			if err := resourceData.Set(string(TargetCapacity), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(TargetCapacity), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(TargetCapacity)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetTarget(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(TargetCapacity)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetTarget(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

}

// expandSubnets expands the list of subnet objects
func expandSubnets(data interface{}) ([]*gcp.Subnet, error) {
	list := data.(*schema.Set).List()
	out := make([]*gcp.Subnet, 0, len(list))

	for _, v := range list {
		elem := &gcp.Subnet{}
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		if v, ok := attr[string(Region)]; ok {
			elem.SetRegion(spotinst.String(v.(string)))
		}

		if v, ok := attr[string(SubnetNames)]; ok {
			subnetList := v.([]interface{})
			result := make([]string, len(subnetList))
			for i, j := range subnetList {
				result[i] = j.(string)
			}
			elem.SetSubnetNames(result)
		}
		out = append(out, elem)
	}
	return out, nil
}

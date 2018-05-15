package elastigroup_aws

import (
	"fmt"
	"strings"
	"errors"
	"bytes"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/hashicorp/terraform/helper/hashcode"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	
	fieldsMap[Name] = commons.NewGenericField(
		commons.ElastigroupAWS,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
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
			elastigroup := resourceObject.(*aws.Group)
			elastigroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			elastigroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Description] = commons.NewGenericField(
		commons.ElastigroupAWS,
		Description,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
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
			elastigroup := resourceObject.(*aws.Group)
			elastigroup.SetDescription(spotinst.String(resourceData.Get(string(Description)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			elastigroup.SetDescription(spotinst.String(resourceData.Get(string(Description)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Product] = commons.NewGenericField(
		commons.ElastigroupAWS,
		Product,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
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
			elastigroup := resourceObject.(*aws.Group)
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

	fieldsMap[MaxSize] = commons.NewGenericField(
		commons.ElastigroupAWS,
		MaxSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *int = nil
			if elastigroup.Capacity != nil && elastigroup.Capacity.Maximum != nil{
				value = elastigroup.Capacity.Maximum
			}
			if err := resourceData.Set(string(MaxSize), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MaxSize), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.Get(string(MaxSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMaximum(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.Get(string(MaxSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMaximum(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[MinSize] = commons.NewGenericField(
		commons.ElastigroupAWS,
		MinSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
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
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.Get(string(MinSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMinimum(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.Get(string(MinSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMinimum(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[DesiredCapacity] = commons.NewGenericField(
		commons.ElastigroupAWS,
		DesiredCapacity,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *int = nil
			if elastigroup.Capacity != nil && elastigroup.Capacity.Target != nil {
				value = elastigroup.Capacity.Target
			}
			if err := resourceData.Set(string(DesiredCapacity), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DesiredCapacity), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.Get(string(DesiredCapacity)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetTarget(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.Get(string(DesiredCapacity)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetTarget(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[CapacityUnit] = commons.NewGenericField(
		commons.ElastigroupAWS,
		CapacityUnit,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *string = nil
			if elastigroup.Capacity != nil && elastigroup.Capacity.Unit != nil {
				value = elastigroup.Capacity.Unit
			}
			if err := resourceData.Set(string(CapacityUnit), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CapacityUnit), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.Get(string(CapacityUnit)).(string); ok && v != "" {
				elastigroup.Capacity.SetUnit(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(CapacityUnit))
			return err
		},
		nil,
	)

	fieldsMap[HealthCheckGracePeriod] = commons.NewGenericField(
		commons.ElastigroupAWS,
		HealthCheckGracePeriod,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *int = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.HealthCheckGracePeriod != nil {
				value = elastigroup.Compute.LaunchSpecification.HealthCheckGracePeriod
			}
			if err := resourceData.Set(string(HealthCheckGracePeriod), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(HealthCheckGracePeriod), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.Get(string(HealthCheckGracePeriod)).(int); ok && v > 0 {
				elastigroup.Compute.LaunchSpecification.SetHealthCheckGracePeriod(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *int = nil
			if v, ok := resourceData.Get(string(HealthCheckGracePeriod)).(int); ok && v > 0 {
				value = spotinst.Int(v)
			}
			elastigroup.Compute.LaunchSpecification.SetHealthCheckGracePeriod(value)
			return nil
		},
		nil,
	)

	fieldsMap[HealthCheckType] = commons.NewGenericField(
		commons.ElastigroupAWS,
		HealthCheckType,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.HealthCheckType != nil {
				value = elastigroup.Compute.LaunchSpecification.HealthCheckType
			}
			if err := resourceData.Set(string(HealthCheckType), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(HealthCheckType), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.Get(string(HealthCheckType)).(string); ok && v != "" {
				elastigroup.Compute.LaunchSpecification.SetHealthCheckType(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *string = nil
			if v, ok := resourceData.Get(string(HealthCheckType)).(string); ok && v != "" {
				value = spotinst.String(v)
			}
			elastigroup.Compute.LaunchSpecification.SetHealthCheckType(value)
			return nil
		},
		nil,
	)

	fieldsMap[HealthCheckUnhealthyDurationBeforeReplacement] = commons.NewGenericField(
		commons.ElastigroupAWS,
		HealthCheckUnhealthyDurationBeforeReplacement,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *int = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.HealthCheckUnhealthyDurationBeforeReplacement != nil {
				value = elastigroup.Compute.LaunchSpecification.HealthCheckUnhealthyDurationBeforeReplacement
			}
			if err := resourceData.Set(string(HealthCheckUnhealthyDurationBeforeReplacement), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(HealthCheckUnhealthyDurationBeforeReplacement), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.Get(string(HealthCheckUnhealthyDurationBeforeReplacement)).(int); ok && v > 0 {
				elastigroup.Compute.LaunchSpecification.SetHealthCheckUnhealthyDurationBeforeReplacement(spotinst.Int(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *int = nil
			if v, ok := resourceData.Get(string(HealthCheckUnhealthyDurationBeforeReplacement)).(int); ok && v > 0 {
				value = spotinst.Int(v)
			}
			elastigroup.Compute.LaunchSpecification.SetHealthCheckUnhealthyDurationBeforeReplacement(value)
			return nil
		},
		nil,
	)

	fieldsMap[SubnetIds] = commons.NewGenericField(
		commons.ElastigroupAWS,
		SubnetIds,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value []string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.SubnetIDs != nil {
				value = elastigroup.Compute.SubnetIDs
			}
			if err := resourceData.Set(string(SubnetIds), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SubnetIds), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if value, ok := resourceData.GetOk(string(SubnetIds)); ok && value != nil {
				if subnetIds, err := expandSubnetIDs(value); err != nil {
					return err
				} else {
					elastigroup.Compute.SetSubnetIDs(subnetIds)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if value, ok := resourceData.GetOk(string(SubnetIds)); ok && value != nil {
				if subnetIds, err := expandSubnetIDs(value); err != nil {
					return err
				} else {
					elastigroup.Compute.SetSubnetIDs(subnetIds)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[AvailabilityZones] = commons.NewGenericField(
		commons.ElastigroupAWS,
		AvailabilityZones,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var zoneNames []string
			if elastigroup.Compute != nil && elastigroup.Compute.AvailabilityZones != nil {
				zones := elastigroup.Compute.AvailabilityZones
				for _, zone := range zones {
					zoneName := spotinst.StringValue(zone.Name)
					zoneNames = append(zoneNames, zoneName)
				}
			}
			if err := resourceData.Set(string(AvailabilityZones), zoneNames); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AvailabilityZones), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if value, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				if zones, err := expandAvailabilityZonesSlice(value); err != nil {
					return err
				} else {
					elastigroup.Compute.SetAvailabilityZones(zones)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if value, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				if zones, err := expandAvailabilityZonesSlice(value); err != nil {
					return err
				} else {
					elastigroup.Compute.SetAvailabilityZones(zones)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[ElasticLoadBalancers] = commons.NewGenericField(
		commons.ElastigroupAWS,
		ElasticLoadBalancers,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var balIds []string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.LoadBalancersConfig != nil &&
				elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers != nil {

				balancers := elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers
				for _, balancer := range balancers {
					balType := spotinst.StringValue(balancer.Type)
					if strings.ToUpper(balType) == string(BalancerTypeClassic) {
						balId := spotinst.StringValue(balancer.BalancerID)
						balIds = append(balIds, balId)
					}
				}
			}
			resourceData.Set(string(ElasticLoadBalancers), balIds)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if balIds, ok := resourceData.GetOk(string(ElasticLoadBalancers)); ok {
				var fn = func(id string) *aws.LoadBalancer {
					return &aws.LoadBalancer{
						Type:       spotinst.String(strings.ToUpper(string(BalancerTypeClassic))),
						BalancerID: spotinst.String(id),
					}
				}
				if err := expandBalancersContent(elastigroup, balIds, fn); err != nil {
					// Do not fail on group creation
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if balIds, ok := resourceData.GetOk(string(ElasticLoadBalancers)); ok {
				var fn = func(id string) *aws.LoadBalancer {
					return &aws.LoadBalancer{
						Type:       spotinst.String(strings.ToUpper(string(BalancerTypeClassic))),
						BalancerID: spotinst.String(id),
					}
				}
				if err := expandBalancersContent(elastigroup, balIds, fn); err != nil {
					// Do not fail on group update
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[TargetGroupArns] = commons.NewGenericField(
		commons.ElastigroupAWS,
		TargetGroupArns,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var tgArns []string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.LoadBalancersConfig != nil &&
				elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers != nil {

				balancers := elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers
				for _, balancer := range balancers {
					balType := spotinst.StringValue(balancer.Type)
					if balType == string(BalancerTypeTargetGroup) {
						arn := spotinst.StringValue(balancer.Arn)
						tgArns = append(tgArns, arn)
					}
				}
			}
			resourceData.Set(string(TargetGroupArns), tgArns)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if tgArns, ok := resourceData.GetOk(string(TargetGroupArns)); ok {
				var fn = func(id string) *aws.LoadBalancer {
					return &aws.LoadBalancer{
						Type: spotinst.String(strings.ToUpper(string(BalancerTypeTargetGroup))),
						Arn:  spotinst.String(id),
					}
				}
				if err := expandBalancersContent(elastigroup, tgArns, fn); err != nil {
					// Do not fail on group creation
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if tgArns, ok := resourceData.GetOk(string(TargetGroupArns)); ok {
				var fn = func(id string) *aws.LoadBalancer {
					return &aws.LoadBalancer{
						Type: spotinst.String(strings.ToUpper(string(BalancerTypeTargetGroup))),
						Arn:  spotinst.String(id),
					}
				}
				if err := expandBalancersContent(elastigroup, tgArns, fn); err != nil {
					// Do not fail on group update
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[MultaiTargetSetIds] = commons.NewGenericField(
		commons.ElastigroupAWS,
		MultaiTargetSetIds,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var tsIds []string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.LoadBalancersConfig != nil &&
				elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers != nil {

				balancers := elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers
				for _, balancer := range balancers {
					balType := spotinst.StringValue(balancer.Type)
					if balType == string(BalancerTypeMultaiTargetSet) {
						tsId := spotinst.StringValue(balancer.TargetSetID)
						tsIds = append(tsIds, tsId)
					}
				}
			}
			resourceData.Set(string(MultaiTargetSetIds), tsIds)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if multaiTsIds, ok := resourceData.GetOk(string(MultaiTargetSetIds)); ok {
				var fn = func(id string) *aws.LoadBalancer {
					return &aws.LoadBalancer{
						Type:        spotinst.String(strings.ToUpper(string(BalancerTypeMultaiTargetSet))),
						TargetSetID: spotinst.String(id),
					}
				}
				if err := expandBalancersContent(elastigroup, multaiTsIds, fn); err != nil {
					// Do not fail on group creation
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if multaiTsIds, ok := resourceData.GetOk(string(MultaiTargetSetIds)); ok {
				var fn = func(id string) *aws.LoadBalancer {
					return &aws.LoadBalancer{
						Type:        spotinst.String(strings.ToUpper(string(BalancerTypeMultaiTargetSet))),
						TargetSetID: spotinst.String(id),
					}
				}
				if err := expandBalancersContent(elastigroup, multaiTsIds, fn); err != nil {
					// Do not fail on group update
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[Tags] = commons.NewGenericField(
		commons.ElastigroupAWS,
		Tags,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(TagKey): &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},

					string(TagValue): &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			Set: hashKV,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var tagsSet *schema.Set = nil
			var tagsToAdd []interface{} = nil

			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Tags != nil {

				tags := elastigroup.Compute.LaunchSpecification.Tags
				tagsToAdd = make([]interface{}, 0, len(tags))
				for _, tag := range tags {
					tagToAdd := &aws.Tag{
						Key: tag.Key,
						Value: tag.Value,
					}
					tagsToAdd = append(tagsToAdd, tagToAdd)
				}

				tagHashFunc := func(item interface{}) int {
					tag := item.(*aws.Tag)
					return hashcode.String(spotinst.StringValue(tag.Key) + spotinst.StringValue(tag.Value))
				}
				tagsSet = schema.NewSet(tagHashFunc, tagsToAdd)
			}
			if err := resourceData.Set(string(Tags), tagsSet); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tags), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var tagsToAdd []*aws.Tag = nil
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					tagsToAdd = tags
				}
			}
			elastigroup.Compute.LaunchSpecification.SetTags(tagsToAdd)
			return nil
		},
		nil,
	)

	fieldsMap[ElasticIps] = commons.NewGenericField(
		commons.ElastigroupAWS,
		ElasticIps,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var result []string
			if elastigroup.Compute != nil && elastigroup.Compute.ElasticIPs != nil {
				elasticIps := elastigroup.Compute.ElasticIPs
				for _, elasticIp := range elasticIps {
					result = append(result, elasticIp)
				}
			}
			if err := resourceData.Set(string(ElasticIps), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ElasticIps), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if value, ok := resourceData.GetOk(string(ElasticIps)); ok {
				if eips, err := expandAWSGroupElasticIPs(value); err != nil {
					return err
				} else {
					elastigroup.Compute.SetElasticIPs(eips)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var result []string = nil
			if value, ok := resourceData.GetOk(string(ElasticIps)); ok {
				if eips, err := expandAWSGroupElasticIPs(value); err != nil {
					return err
				} else {
					result = eips
				}
			}
			elastigroup.Compute.SetElasticIPs(result)
			return nil
		},
		nil,
	)

	fieldsMap[Signal] = commons.NewGenericField(
		commons.ElastigroupAWS,
		Signal,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(SignalName): &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},

					string(SignalTimeout): &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var signalsToAdd = []*aws.Signal{}
			if elastigroup.Strategy != nil && elastigroup.Strategy.Signals != nil {
				signals := elastigroup.Strategy.Signals
				if signals != nil {
					signalsToAdd := make([]interface{}, 0, len(signals))
					for _, s := range signals {
						m := make(map[string]interface{})
						m[string(SignalName)] = strings.ToLower(spotinst.StringValue(s.Name))
						m[string(SignalTimeout)] = spotinst.IntValue(s.Timeout)
						signalsToAdd = append(signalsToAdd, m)
					}
				}
			}
			if err := resourceData.Set(string(Signal), signalsToAdd); err != nil {
				return fmt.Errorf("failed to set signals configuration: %#v", err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.GetOk(string(Signal)); ok {
				if signals, err := expandSignals(v); err != nil {
					return err
				} else {
					elastigroup.Strategy.SetSignals(signals)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var signalsToAdd []*aws.Signal = nil
			if v, ok := resourceData.GetOk(string(Signal)); ok {
				if signals, err := expandSignals(v); err != nil {
					return err
				} else {
					signalsToAdd = signals
				}
			}
			if elastigroup.Strategy == nil {
				elastigroup.SetStrategy(&aws.Strategy{})
			}
			elastigroup.Strategy.SetSignals(signalsToAdd)
			return nil
		},
		nil,
	)

	fieldsMap[UpdatePolicy] = commons.NewGenericField(
		commons.ElastigroupAWS,
		UpdatePolicy,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ShouldResumeStateful): &schema.Schema{
						Type:     schema.TypeBool,
						Required: true,
					},

					string(ShouldRoll): &schema.Schema{
						Type:     schema.TypeBool,
						Required: true,
					},

					string(RollConfig): &schema.Schema{
						Type:     schema.TypeSet,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(BatchSizePercentage): &schema.Schema{
									Type:     schema.TypeInt,
									Required: true,
								},

								string(GracePeriod): &schema.Schema{
									Type:     schema.TypeInt,
									Optional: true,
									Default:  -1,
								},

								string(HealthCheckType): &schema.Schema{
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

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Fields Expand
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAvailabilityZonesSlice(data interface{}) ([]*aws.AvailabilityZone, error) {
	list := data.([]interface{})
	zones := make([]*aws.AvailabilityZone, 0, len(list))
	for _, str := range list {
		if s, ok := str.(string); ok {
			parts := strings.Split(s, ":")
			zone := &aws.AvailabilityZone{}
			if len(parts) >= 1 && parts[0] != "" {
				zone.SetName(spotinst.String(parts[0]))
			}
			if len(parts) == 2 && parts[1] != "" {
				zone.SetSubnetId(spotinst.String(parts[1]))
			}
			if len(parts) == 3 && parts[2] != "" {
				zone.SetPlacementGroupName(spotinst.String(parts[2]))
			}
			zones = append(zones, zone)
		}
	}

	return zones, nil
}

func expandAWSGroupElasticIPs(data interface{}) ([]string, error) {
	list := data.([]interface{})
	eips := make([]string, 0, len(list))
	for _, str := range list {
		if eip, ok := str.(string); ok {
			eips = append(eips, eip)
		}
	}
	return eips, nil
}

func expandTags(data interface{}) ([]*aws.Tag, error) {
	list := data.(*schema.Set).List()
	tags := make([]*aws.Tag, 0, len(list))
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
		tag := &aws.Tag{
			Key:   spotinst.String(attr["key"].(string)),
			Value: spotinst.String(attr["value"].(string)),
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

type CreateBalancerObjFunc func(id string) *aws.LoadBalancer

func expandBalancersContent(elastigroup *aws.Group, ids interface{}, fn CreateBalancerObjFunc) error {
	if ids == nil {
		return nil
	}
	var balancers []*aws.LoadBalancer = nil
	list := ids.([]interface{})
	if elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers != nil {
		balancers = elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers
	} else {
		balancers = make([]*aws.LoadBalancer, 0, len(list))
	}
	for _, str := range list {
		if id, ok := str.(string); ok {
			lb := fn(id)
			balancers = append(balancers, lb)
		}
	}
	return nil
}

func expandSignals(data interface{}) ([]*aws.Signal, error) {
	list := data.(*schema.Set).List()
	signals := make([]*aws.Signal, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		signal := &aws.Signal{}

		if v, ok := m[string(SignalName)].(string); ok && v != "" {
			signal.SetName(spotinst.String(strings.ToUpper(v)))
		}

		if v, ok := m[string(SignalTimeout)].(int); ok && v > 0 {
			signal.SetTimeout(spotinst.Int(v))
		}
		signals = append(signals, signal)
	}

	return signals, nil
}

func expandSubnetIDs(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if subnetID, ok := v.(string); ok && subnetID != "" {
			result = append(result, subnetID)
		}
	}
	return result, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utilities
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func hashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(TagKey)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(TagValue)].(string)))
	return hashcode.String(buf.String())
}
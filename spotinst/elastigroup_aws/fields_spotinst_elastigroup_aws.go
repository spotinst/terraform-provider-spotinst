package elastigroup_aws

import (
	"fmt"
	"strings"
	"log"
	"errors"
	"bytes"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/hashicorp/terraform/helper/hashcode"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupAwsElastigroupResource() {
	fields := make(map[commons.FieldName]*commons.GenericField)
	var readFailurePattern = "elastigroup failed reading field %s - %#v"

	fields[Name] = commons.NewGenericField(
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *string = nil
			if elastigroup.Name != nil {
				value = elastigroup.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(readFailurePattern, string(Name), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		nil,
	)

	fields[Description] = commons.NewGenericField(
		Description,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *string = nil
			if elastigroup.Description!= nil {
				value = elastigroup.Description
			}
			if err := resourceData.Set(string(Description), value); err != nil {
				return fmt.Errorf(readFailurePattern, string(Description), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup.SetDescription(spotinst.String(resourceData.Get(string(Description)).(string)))
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup.SetDescription(spotinst.String(resourceData.Get(string(Description)).(string)))
			return nil
		},
		nil,
	)

	fields[Product] = commons.NewGenericField(
		Product,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.Product != nil {
				value = elastigroup.Compute.Product
			}
			if err := resourceData.Set(string(Product), value); err != nil {
				return fmt.Errorf(readFailurePattern, string(Product), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup.Compute.SetProduct(spotinst.String(resourceData.Get(string(Product)).(string)))
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup.Compute.SetProduct(spotinst.String(resourceData.Get(string(Product)).(string)))
			return nil
		},
		nil,
	)

	fields[MaxSize] = commons.NewGenericField(
		MaxSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *int = nil
			if elastigroup.Capacity != nil && elastigroup.Capacity.Maximum != nil{
				value = elastigroup.Capacity.Maximum
			}
			if err := resourceData.Set(string(MaxSize), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(MaxSize), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(MaxSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMaximum(spotinst.Int(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(MaxSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMaximum(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fields[MinSize] = commons.NewGenericField(
		MinSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *int = nil
			if elastigroup.Capacity != nil && elastigroup.Capacity.Minimum != nil {
				value = elastigroup.Capacity.Minimum
			}
			if err := resourceData.Set(string(MinSize), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(MinSize), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(MinSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMinimum(spotinst.Int(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(MinSize)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetMinimum(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fields[DesiredCapacity] = commons.NewGenericField(
		DesiredCapacity,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *int = nil
			if elastigroup.Capacity != nil && elastigroup.Capacity.Target != nil {
				value = elastigroup.Capacity.Target
			}
			if err := resourceData.Set(string(DesiredCapacity), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(DesiredCapacity), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(DesiredCapacity)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetTarget(spotinst.Int(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(DesiredCapacity)).(int); ok && v >= 0 {
				elastigroup.Capacity.SetTarget(spotinst.Int(v))
			}
			return nil
		},
		nil,
	)

	fields[CapacityUnit] = commons.NewGenericField(
		CapacityUnit,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *string = nil
			if elastigroup.Capacity != nil && elastigroup.Capacity.Unit != nil {
				value = elastigroup.Capacity.Unit
			}
			if err := resourceData.Set(string(CapacityUnit), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(CapacityUnit), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(CapacityUnit)).(string); ok && v != "" {
				elastigroup.Capacity.SetUnit(spotinst.String(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			// Do nothing
			return nil
		},
		nil,
	)

	fields[HealthCheckGracePeriod] = commons.NewGenericField(
		HealthCheckGracePeriod,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *int = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.HealthCheckGracePeriod != nil {
				value = elastigroup.Compute.LaunchSpecification.HealthCheckGracePeriod
			}
			if err := resourceData.Set(string(HealthCheckGracePeriod), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(HealthCheckGracePeriod), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(HealthCheckGracePeriod)).(int); ok && v > 0 {
				elastigroup.Compute.LaunchSpecification.SetHealthCheckGracePeriod(spotinst.Int(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *int = nil
			if v, ok := resourceData.Get(string(HealthCheckGracePeriod)).(int); ok && v > 0 {
				value = spotinst.Int(v)
			}
			elastigroup.Compute.LaunchSpecification.SetHealthCheckGracePeriod(value)
			return nil
		},
		nil,
	)

	fields[HealthCheckType] = commons.NewGenericField(
		HealthCheckType,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.HealthCheckType != nil {
				value = elastigroup.Compute.LaunchSpecification.HealthCheckType
			}
			if err := resourceData.Set(string(HealthCheckType), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(HealthCheckType), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(HealthCheckType)).(string); ok && v != "" {
				elastigroup.Compute.LaunchSpecification.SetHealthCheckType(spotinst.String(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *string = nil
			if v, ok := resourceData.Get(string(HealthCheckType)).(string); ok && v != "" {
				value = spotinst.String(v)
			}
			elastigroup.Compute.LaunchSpecification.SetHealthCheckType(value)
			return nil
		},
		nil,
	)

	fields[HealthCheckUnhealthyDurationBeforeReplacement] = commons.NewGenericField(
		HealthCheckUnhealthyDurationBeforeReplacement,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *int = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.HealthCheckUnhealthyDurationBeforeReplacement != nil {
				value = elastigroup.Compute.LaunchSpecification.HealthCheckUnhealthyDurationBeforeReplacement
			}
			if err := resourceData.Set(string(HealthCheckUnhealthyDurationBeforeReplacement), spotinst.IntValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(HealthCheckUnhealthyDurationBeforeReplacement), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(HealthCheckUnhealthyDurationBeforeReplacement)).(int); ok && v > 0 {
				elastigroup.Compute.LaunchSpecification.SetHealthCheckUnhealthyDurationBeforeReplacement(spotinst.Int(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *int = nil
			if v, ok := resourceData.Get(string(HealthCheckUnhealthyDurationBeforeReplacement)).(int); ok && v > 0 {
				value = spotinst.Int(v)
			}
			elastigroup.Compute.LaunchSpecification.SetHealthCheckUnhealthyDurationBeforeReplacement(value)
			return nil
		},
		nil,
	)

	fields[SubnetIds] = commons.NewGenericField(
		SubnetIds,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			// TODO: Complete when vpc_zone_identifier is implemented
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			// TODO: Complete when vpc_zone_identifier is implemented
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			// TODO: Complete when vpc_zone_identifier is implemented
			return nil
		},
		nil,
	)

	fields[AvailabilityZones] = commons.NewGenericField(
		AvailabilityZones,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var zoneNames []string
			if elastigroup.Compute != nil && elastigroup.Compute.AvailabilityZones != nil {
				zones := elastigroup.Compute.AvailabilityZones
				for _, zone := range zones {
					zoneName := spotinst.StringValue(zone.Name)
					zoneNames = append(zoneNames, zoneName)
				}
			}
			if err := resourceData.Set(string(AvailabilityZones), zoneNames); err != nil {
				return fmt.Errorf(readFailurePattern, string(AvailabilityZones), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if value, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				if zones, err := expandAWSGroupAvailabilityZonesSlice(value); err != nil {
					return err
				} else {
					elastigroup.Compute.SetAvailabilityZones(zones)
				}
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if value, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				if zones, err := expandAWSGroupAvailabilityZonesSlice(value); err != nil {
					return err
				} else {
					elastigroup.Compute.SetAvailabilityZones(zones)
				}
			}
			return nil
		},
		nil,
	)

	fields[ElasticLoadBalancers] = commons.NewGenericField(
		ElasticLoadBalancers,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
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
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if balIds, ok := resourceData.GetOk(string(ElasticLoadBalancers)); ok {
				if balancers, err := expandAWSGroupBalancers(balIds); err == nil {
					if elastigroup.Compute.LaunchSpecification.LoadBalancersConfig == nil {
						elastigroup.Compute.LaunchSpecification.LoadBalancersConfig = &aws.LoadBalancersConfig{}
					}
					elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(balancers)
				}
			}
			// Do not fail on group creation
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if elastigroup.Compute.LaunchSpecification.LoadBalancersConfig == nil {
				elastigroup.Compute.LaunchSpecification.SetLoadBalancersConfig(&aws.LoadBalancersConfig{})
			}

			if balIds, ok := resourceData.GetOk(string(ElasticLoadBalancers)); ok {
				if balancers, err := expandAWSGroupBalancers(balIds); err == nil {
					elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(balancers)
				}
			}
			// Do not fail on group update
			return nil
		},
		nil,
	)

	fields[TargetGroupArns] = commons.NewGenericField(
		TargetGroupArns,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var tarArns []string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.LoadBalancersConfig != nil &&
				elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers != nil {

				balancers := elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers
				for _, balancer := range balancers {
					balType := spotinst.StringValue(balancer.Type)
					if balType == string(BalancerTypeTargetGroup) {
						arn := spotinst.StringValue(balancer.Arn)
						tarArns = append(tarArns, arn)
					}
				}
			}
			resourceData.Set(string(TargetGroupArns), tarArns)
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fields[MultaiTargetSetIds] = commons.NewGenericField(
		MultaiTargetSetIds,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
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
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fields[Tags] = commons.NewGenericField(
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
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			//var tagsToAdd = make(map[string]interface{})
			var tagsSet *schema.Set = nil
			//var tagsToAdd []*aws.Tag = nil
			var tagsToAdd []interface{} = nil

			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Tags != nil {

				tags := elastigroup.Compute.LaunchSpecification.Tags
				//tagsToAdd = make([]*aws.Tag, 0, len(tags))
				tagsToAdd = make([]interface{}, 0, len(tags))
				//result := make(map[string]interface{})
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
				return fmt.Errorf(readFailurePattern, string(Tags), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandAWSGroupTags(value); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetTags(tags)
				}
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var tagsToAdd []*aws.Tag = nil
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandAWSGroupTags(value); err != nil {
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

	fields[LaunchConfiguration] = commons.NewGenericField(
		LaunchConfiguration,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			// Read is being handled by Terraform interpolation
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			// Create is being handled by Terraform interpolation
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			// Update is being handled by Terraform interpolation
			return nil
		},
		nil,
	)

	fields[Strategy] = commons.NewGenericField(
		Strategy,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			// Read is being handled by Terraform interpolation
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			// Creation is being handled by Terraform interpolation
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			// Strategy update is being handled by Terraform interpolation
			// Add strategy signal, if exists
			if v, ok := resourceData.GetOk(string(Signal)); ok {
				if signals, err := expandAWSGroupSignals(v); err != nil {
					return err
				} else {
					elastigroup.Strategy.SetSignals(signals)
				}
			}
			return nil
		},
		nil,
	)

	fields[Signal] = commons.NewGenericField(
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
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
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
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(Signal)).(int); ok && v > 0 {
				if signals, err := expandAWSGroupSignals(v); err != nil {
					return err
				} else {
					elastigroup.Strategy.SetSignals(signals)
				}
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var signalsToAdd []*aws.Signal = nil
			if v, ok := resourceData.GetOk(string(Signal)); ok {
				if signals, err := expandAWSGroupSignals(v); err != nil {
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

	fields[InstanceTypes] = commons.NewGenericField(
		InstanceTypes,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fields[EbsBlockDevice] = commons.NewGenericField(
		EbsBlockDevice,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fields[EphemeralBlockDevice] = commons.NewGenericField(
		EphemeralBlockDevice,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	//fields[UpdatePolicy] = commons.NewGenericField(
	//	UpdatePolicy,
	//	&schema.Schema{
	//		Type:     schema.TypeSet,
	//		Optional: true,
	//		Elem: &schema.Resource{
	//			Schema: map[string]*schema.Schema{
	//				"should_roll": &schema.Schema{
	//					Type:     schema.TypeBool,
	//					Required: true,
	//				},
	//
	//				"batch_size_percentage": &schema.Schema{
	//					Type:     schema.TypeInt,
	//					Required: true,
	//				},
	//
	//				"grace_period": &schema.Schema{
	//					Type:     schema.TypeInt,
	//					Optional: true,
	//					Default:  -1,
	//				},
	//
	//				"health_check_type": &schema.Schema{
	//					Type:     schema.TypeString,
	//					Optional: true,
	//				},
	//			},
	//		},
	//	},
	//	func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
	//		tags := elastigroup.Compute.LaunchSpecification.Tags
	//		if tags == nil {
	//			return nil
	//		}
	//
	//		var tagsToAdd = make(map[string]interface{})
	//		for _, tag := range tags {
	//			tagsToAdd[spotinst.StringValue(tag.Key)] = tag.Value
	//		}
	//
	//		if err := resourceData.Set(string(Tags), tagsToAdd); err != nil {
	//			return fmt.Errorf(readFailurePattern, string(Tags), err)
	//		}
	//		return nil
	//	},
	//	func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
	//		if value, ok := resourceData.GetOk(string(Tags)); ok {
	//			if tags, err := expandAWSGroupTags(value); err != nil {
	//				return err
	//			} else {
	//				elastigroup.Compute.LaunchSpecification.SetTags(tags)
	//			}
	//		}
	//		return nil
	//	},
	//	func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
	//		var tagsToAdd []*aws.Tag = nil
	//		if value, ok := resourceData.GetOk(string(Tags)); ok {
	//			if tags, err := expandAWSGroupTags(value); err != nil {
	//				return err
	//			} else {
	//				tagsToAdd = tags
	//			}
	//		}
	//		elastigroup.Compute.LaunchSpecification.SetTags(tagsToAdd)
	//		return nil
	//	},
	//	nil,
	//)


	commons.ElastigroupResource = commons.NewGenericApiResource(
		string(commons.ElastigroupAWS),
		fields)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Fields Expand
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupAvailabilityZonesSlice(data interface{}) ([]*aws.AvailabilityZone, error) {
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
			log.Printf("Group availability zone configuration: %s", stringutil.Stringify(zone))
			zones = append(zones, zone)
		}
	}

	return zones, nil
}

func expandAWSGroupTags(data interface{}) ([]*aws.Tag, error) {
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
		log.Printf("Group tag configuration: %s", stringutil.Stringify(tag))
		tags = append(tags, tag)
	}
	return tags, nil
}

func expandAWSGroupBalancers(balancerIds interface{}) ([]*aws.LoadBalancer, error) {
	var balancers []*aws.LoadBalancer = nil
	if balancerIds != nil {
		list := balancerIds.([]interface{})
		balancers = make([]*aws.LoadBalancer, 0, len(list))
		for _, str := range list {
			if balId, ok := str.(string); ok {
				log.Printf("Balancer id configuration: %s", stringutil.Stringify(balId))
				lb := &aws.LoadBalancer{
					Type:       spotinst.String(strings.ToUpper(string(BalancerTypeClassic))),
					BalancerID: spotinst.String(balId),
				}
				balancers = append(balancers, lb)
			}
		}
	}
	return balancers, nil
}

// expandAWSGroupSignals expands the Signal block.
func expandAWSGroupSignals(data interface{}) ([]*aws.Signal, error) {
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

		log.Printf("Group signal configuration: %s", stringutil.Stringify(signal))
		signals = append(signals, signal)
	}

	return signals, nil
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
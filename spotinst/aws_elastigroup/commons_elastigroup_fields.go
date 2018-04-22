package aws_elastigroup

import (
	"fmt"
	"strings"
	"log"
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupAwsElastigroupResource() {
	fields := make(map[commons.FieldName]*commons.GenericField)
	var readFailurePattern = "elastigroup failed reading field %s - %#v"
	var createFailurePattern = "elastigroup failed creating field %s - %#v"
	var updateFailurePattern = "elastigroup failed updating field %s - %#v"

	fields[Name] = commons.NewGenericField(
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if err := resourceData.Set(string(Name), elastigroup.Name); err != nil {
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
			if err := resourceData.Set(string(Description), elastigroup.Description); err != nil {
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

	fields[MaxSize] = commons.NewGenericField(
		MaxSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			value := spotinst.IntValue(elastigroup.Capacity.Maximum)
			if err := resourceData.Set(string(MaxSize), value); err != nil {
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
			value := spotinst.IntValue(elastigroup.Capacity.Minimum)
			if err := resourceData.Set(string(MinSize), value); err != nil {
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
			return nil
		},
		nil,
	)

	fields[DesiredCapacity] = commons.NewGenericField(
		DesiredCapacity,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			value := spotinst.IntValue(elastigroup.Capacity.Target)
			if err := resourceData.Set(string(DesiredCapacity), value); err != nil {
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
			value := spotinst.StringValue(elastigroup.Capacity.Unit)
			if err := resourceData.Set(string(CapacityUnit), value); err != nil {
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
			value := spotinst.IntValue(elastigroup.Compute.LaunchSpecification.HealthCheckGracePeriod)
			if err := resourceData.Set(string(HealthCheckGracePeriod), value); err != nil {
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
			value := spotinst.StringValue(elastigroup.Compute.LaunchSpecification.HealthCheckType)
			if err := resourceData.Set(string(HealthCheckType), value); err != nil {
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
			value := spotinst.IntValue(elastigroup.Compute.LaunchSpecification.HealthCheckUnhealthyDurationBeforeReplacement)
			if err := resourceData.Set(string(HealthCheckUnhealthyDurationBeforeReplacement), value); err != nil {
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
			zones := elastigroup.Compute.AvailabilityZones
			// TODO: convert to AZ schema
			value := convrertToZones(zones)
			if err := resourceData.Set(string(AvailabilityZones), value); err != nil {
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
			balancers := elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers
			var balIds []string
			for _, balancer, := range balancers {
				balType := spotinst.StringValue(balancer.Type)
				if balType == string(BalancerTypeClassic) {
					balId := spotinst.StringValue(balancer.BalancerID)
					balIds = append(balIds, balId)
				}
			}
			resourceData.Set(string(ElasticLoadBalancers), balIds)
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {

			if v, ok := resourceData.Get(string(ElasticLoadBalancers)); ok && v != nil {
				elastigroup.Compute.LaunchSpecification.SetHealthCheckUnhealthyDurationBeforeReplacement(spotinst.Int(v))
			}
			return nil

			if data == nil {
				log.Print("[ERROR] Cannot expand AWS group load balancers due to <nil> value")
				// Do not fail the terraform process
				return nil, nil
			}
			list := data.([]interface{})
			lbs := make([]*aws.LoadBalancer, 0, len(list))
			for _, item := range list {
				if item == nil {
					log.Print("[ERROR] Empty load balancer value, skipping creation")
					continue
				}
				m := item.(string)
				lb := &aws.LoadBalancer{}

				fields := strings.Split(m, ",")
				for _, field := range fields {
					kv := strings.Split(field, "=")
					if len(kv) == 2 {
						key := kv[0]
						val := spotinst.String(kv[1])
						switch key {
						case "type":
							lb.SetType(val)
						case "name":
							lb.SetName(val)
						case "arn":
							lb.SetArn(val)
						case "balancer_id":
							lb.SetBalancerId(val)
						case "target_set_id":
							lb.SetTargetSetId(val)
						case "auto_weight":
							if kv[1] == "true" {
								lb.SetAutoWeight(spotinst.Bool(true))
							}
						case "zone_awareness":
							if kv[1] == "true" {
								lb.SetZoneAwareness(spotinst.Bool(true))
							}
						}
					}
				}

				log.Printf("[DEBUG] Group load balancer configuration: %s", stringutil.Stringify(lb))
				lbs = append(lbs, lb)
			}

			return lbs, nil




			if lbs, err := expandAWSGroupLoadBalancers(v, nullify); err != nil {
				return nil, err
			} else {
				if elastigroup.Compute.LaunchSpecification.LoadBalancersConfig == nil {
					elastigroup.Compute.LaunchSpecification.LoadBalancersConfig = &aws.LoadBalancersConfig{}
				}
			 	elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(lbs)

			 	// TODO: create load balancer objects and create an array pre-assignment


			}

			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
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
			balancers := elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers
			var tarArns []string
			for _, balancer, := range balancers {
				balType := spotinst.StringValue(balancer.Type)
				if balType == string(BalancerTypeTargetGroup) {
					arn := spotinst.StringValue(balancer.Arn)
					tarArns = append(tarArns, arn)
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
			balancers := elastigroup.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers
			var tsIds []string
			for _, balancer, := range balancers {
				balType := spotinst.StringValue(balancer.Type)
				if balType == string(BalancerTypeMultaiTargetSet) {
					tsId := spotinst.StringValue(balancer.TargetSetID)
					tsIds = append(tsIds, tsId)
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
					"key": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},

					"value": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			Set: hashKV,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			tags := elastigroup.Compute.LaunchSpecification.Tags
			// TODO: convert tags to schema structure
			value := convertTags(tags)
			if err := resourceData.Set(string(Tags), value); err != nil {
				return fmt.Errorf(readFailurePattern, string(Tags), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
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
		nil,
	)

	fields[LaunchConfiguration] = commons.NewGenericField(
		LaunchConfiguration,
		&schema.Schema{
			Type:     schema.TypeString,
			MaxItems: 1,
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

	fields[InstanceTypes] = commons.NewGenericField(
		InstanceTypes,
		&schema.Schema{
			Type:     schema.TypeString,
			MaxItems: 1,
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
			MaxItems: 1,
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
			MaxItems: 1,
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

	commons.ElastigroupRepo = commons.NewGenericApiResource(fields)
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
			log.Printf("[DEBUG] Group availability zone configuration: %s", stringutil.Stringify(zone))
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
		if _, ok := attr["key"]; !ok {
			return nil, errors.New("invalid tag attributes: key missing")
		}

		if _, ok := attr["value"]; !ok {
			return nil, errors.New("invalid tag attributes: value missing")
		}
		tag := &aws.Tag{
			Key:   spotinst.String(attr["key"].(string)),
			Value: spotinst.String(attr["value"].(string)),
		}
		log.Printf("[DEBUG] Group tag configuration: %s", stringutil.Stringify(tag))
		tags = append(tags, tag)
	}
	return tags, nil
}
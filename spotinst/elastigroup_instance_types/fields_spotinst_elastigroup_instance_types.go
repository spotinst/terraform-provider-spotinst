package elastigroup_instance_types

import (
	"fmt"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"log"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
	"errors"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupSpotinstInstanceTypesResource() {
	fields := make(map[commons.FieldName]*commons.GenericField)
	var readFailurePattern = "instance types failed reading field %s - %#v"

	fields[OnDemand] = commons.NewGenericField(
		OnDemand,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.InstanceTypes != nil &&
				elastigroup.Compute.InstanceTypes.OnDemand != nil {
				value = elastigroup.Compute.InstanceTypes.OnDemand
			}
			if err := resourceData.Set(string(OnDemand), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(OnDemand), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(OnDemand)).(string); ok && v != "" {
				elastigroup.Compute.InstanceTypes.SetOnDemand(spotinst.String(v))
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.Get(string(OnDemand)).(string); ok && v != "" {
				elastigroup.Compute.InstanceTypes.SetOnDemand(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fields[Spot] = commons.NewGenericField(
		Spot,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value []string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.InstanceTypes != nil &&
				elastigroup.Compute.InstanceTypes.Spot != nil {
				value = elastigroup.Compute.InstanceTypes.Spot
			}
			if err := resourceData.Set(string(Spot), spotinst.StringSlice(value)); err != nil {
				return fmt.Errorf(readFailurePattern, string(Spot), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.GetOk(string(Spot)); ok {
				spots := v.([]interface{})
				spotTypes := make([]string, len(spots))
				for i, j := range spots {
					spotTypes[i] = j.(string)
				}
				elastigroup.Compute.InstanceTypes.SetSpot(spotTypes)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.GetOk(string(Spot)); ok {
				spotTypes := v.([]string)
				elastigroup.Compute.InstanceTypes.SetSpot(spotTypes)
			}
			return nil
		},
		nil,
	)

	fields[InstanceTypeWeights] = commons.NewGenericField(
		InstanceTypeWeights,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(InstanceType): &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},

					string(Weight): &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var instanceWeights []map[string]interface{} = nil
			if elastigroup.Compute != nil && elastigroup.Compute.InstanceTypes != nil &&
				elastigroup.Compute.InstanceTypes.Weights != nil {

				weights := elastigroup.Compute.InstanceTypes.Weights
				instanceWeights := []map[string]interface{}{}
				for _, t := range weights {
					instanceWeight := make(map[string]interface{})
					instanceWeight[string(InstanceType)] = spotinst.StringValue(t.InstanceType)
					instanceWeight[string(Weight)] = spotinst.IntValue(t.Weight)
					instanceWeights = append(instanceWeights, instanceWeight)
				}
			}
			if err := resourceData.Set(string(InstanceTypeWeights), instanceWeights); err != nil {
				return fmt.Errorf(readFailurePattern, string(InstanceTypeWeights), err)
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.GetOkExists(string(InstanceTypeWeights)); ok && v != "" {
				if weights, err := expandAWSGroupInstanceTypeWeights(v); err != nil {
					return err
				} else {
					elastigroup.Compute.InstanceTypes.SetWeights(weights)
				}
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var weightsToAdd []*aws.InstanceTypeWeight = nil
			if v, ok := resourceData.GetOk(string(InstanceTypeWeights)); ok {
				if weights, err := expandAWSGroupInstanceTypeWeights(v); err != nil {
					return err
				} else {
					weightsToAdd = weights
				}
			}

			if elastigroup.Compute == nil {
				elastigroup.SetCompute(&aws.Compute{})
			}
			if elastigroup.Compute.InstanceTypes == nil {
				elastigroup.Compute.SetInstanceTypes(&aws.InstanceTypes{})
			}
			elastigroup.Compute.InstanceTypes.SetWeights(weightsToAdd)
			return nil
		},
		nil,
	)

	commons.ElastigroupInstanceTypesResource = commons.NewGenericCachedResource(
		string(commons.ElastigroupInstanceType),
		fields)
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupInstanceTypeWeights(data interface{}) ([]*aws.InstanceTypeWeight, error) {
	list := data.(*schema.Set).List()
	weights := make([]*aws.InstanceTypeWeight, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(InstanceType)]; !ok {
			return nil, errors.New("invalid instance type weight: instance_type missing")
		}

		if _, ok := attr[string(Weight)]; !ok {
			return nil, errors.New("invalid instance type weight: weight missing")
		}
		instanceWeight := &aws.InstanceTypeWeight{}
		instanceWeight.SetInstanceType(spotinst.String(attr[string(InstanceType)].(string)))
		instanceWeight.SetWeight(spotinst.Int(attr[string(Weight)].(int)))
		log.Printf("Group instance type weight configuration: %s", stringutil.Stringify(instanceWeight))
		weights = append(weights, instanceWeight)
	}
	return weights, nil
}
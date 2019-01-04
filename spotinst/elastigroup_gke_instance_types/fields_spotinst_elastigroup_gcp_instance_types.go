package elastigroup_gke_instance_types

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

	fieldsMap[OnDemand] = commons.NewGenericField(
		commons.ElastigroupGKEInstanceType,
		OnDemand,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGKEWrapper)
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
			egWrapper := resourceObject.(*commons.ImportGKEWrapper)
			elastigroup := egWrapper.GetImport()
			if v, ok := resourceData.Get(string(OnDemand)).(string); ok && v != "" {
				if elastigroup.InstanceTypes == nil {
					elastigroup.SetInstanceTypes(&gcp.InstanceTypesGKE{OnDemand: spotinst.String(v)})
				} else {
					elastigroup.InstanceTypes.SetOnDemand(spotinst.String(v))
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGKEWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(OnDemand)).(string); ok && v != "" {
				elastigroup.Compute.InstanceTypes.SetOnDemand(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Preemptible] = commons.NewGenericField(
		commons.ElastigroupGCPInstanceType,
		Preemptible,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGKEWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []string
			if elastigroup.Compute != nil && elastigroup.Compute.InstanceTypes != nil &&
				elastigroup.Compute.InstanceTypes.Preemptible != nil {
				values := elastigroup.Compute.InstanceTypes.Preemptible
				for _, preemptible := range values {
					result = append(result, preemptible)
				}
			}
			if err := resourceData.Set(string(Preemptible), result); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Preemptible), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ImportGKEWrapper)
			elastigroup := egWrapper.GetImport()
			if v, ok := resourceData.GetOk(string(Preemptible)); ok {
				prempts := v.([]interface{})
				premptTypes := make([]string, len(prempts))
				for i, j := range prempts {
					premptTypes[i] = j.(string)
				}
				if elastigroup.InstanceTypes == nil {
					elastigroup.SetInstanceTypes(&gcp.InstanceTypesGKE{Preemptible: premptTypes})
				} else {
					elastigroup.InstanceTypes.SetPreemptible(premptTypes)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGKEWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Preemptible)); ok {
				prempts := v.([]interface{})
				premptTypes := make([]string, len(prempts))
				for i, j := range prempts {
					premptTypes[i] = j.(string)
				}
				elastigroup.Compute.InstanceTypes.SetPreemptible(premptTypes)
			}
			return nil
		},
		nil,
	)
}

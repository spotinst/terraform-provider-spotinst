package elastigroup_azure_vm_sizes

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[OnDemandSizes] = commons.NewGenericField(
		commons.ElastigroupAzureVMSizes,
		OnDemandSizes,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()

			var result []string
			if elastigroup.Compute != nil && elastigroup.Compute.VMSizes != nil &&
				elastigroup.Compute.VMSizes.OnDemandSizes != nil {
				result = append(result, elastigroup.Compute.VMSizes.OnDemandSizes...)
				if err := resourceData.Set(string(OnDemandSizes), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OnDemandSizes), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(OnDemandSizes)); ok {
				virtualMachines := v.([]interface{})
				onDemandSizes := make([]string, len(virtualMachines))
				for i, j := range virtualMachines {
					onDemandSizes[i] = j.(string)
				}
				elastigroup.Compute.VMSizes.SetOnDemandSizes(onDemandSizes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(OnDemandSizes)); ok {
				virtualMachines := v.([]interface{})
				onDemandSizes := make([]string, len(virtualMachines))
				for i, j := range virtualMachines {
					onDemandSizes[i] = j.(string)
				}
				elastigroup.Compute.VMSizes.SetOnDemandSizes(onDemandSizes)
			}
			return nil
		},
		nil,
	)

	fieldsMap[SpotSizes] = commons.NewGenericField(
		commons.ElastigroupAzureVMSizes,
		SpotSizes,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []string
			if elastigroup.Compute != nil && elastigroup.Compute.VMSizes != nil &&
				elastigroup.Compute.VMSizes.SpotSizes != nil {
				result = append(result, elastigroup.Compute.VMSizes.SpotSizes...)
				if err := resourceData.Set(string(SpotSizes), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SpotSizes), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(SpotSizes)); ok {
				virtualMachines := v.([]interface{})
				spotSizes := make([]string, len(virtualMachines))
				for i, j := range virtualMachines {
					spotSizes[i] = j.(string)
				}
				elastigroup.Compute.VMSizes.SetSpotSizes(spotSizes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(SpotSizes)); ok {
				virtualMachines := v.([]interface{})
				spotSizes := make([]string, len(virtualMachines))
				for i, j := range virtualMachines {
					spotSizes[i] = j.(string)
				}
				elastigroup.Compute.VMSizes.SetSpotSizes(spotSizes)
			}
			return nil
		},
		nil,
	)

}

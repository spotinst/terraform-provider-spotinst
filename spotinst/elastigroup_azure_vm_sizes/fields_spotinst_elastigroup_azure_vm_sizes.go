package elastigroup_azure_vm_sizes

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[OnDemand] = commons.NewGenericField(
		commons.ElastigroupAzureVMSizes,
		OnDemand,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()

			var result []string
			if elastigroup.Compute != nil && elastigroup.Compute.VMSizes != nil &&
				elastigroup.Compute.VMSizes.OnDemand != nil {
				result = append(result, elastigroup.Compute.VMSizes.OnDemand...)
				if err := resourceData.Set(string(OnDemand), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OnDemand), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(OnDemand)); ok {
				virtualMachines := v.([]interface{})
				onDemandSizes := make([]string, len(virtualMachines))
				for i, j := range virtualMachines {
					onDemandSizes[i] = j.(string)
				}
				elastigroup.Compute.VMSizes.SetOnDemand(onDemandSizes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(OnDemand)); ok {
				virtualMachines := v.([]interface{})
				onDemandSizes := make([]string, len(virtualMachines))
				for i, j := range virtualMachines {
					onDemandSizes[i] = j.(string)
				}
				elastigroup.Compute.VMSizes.SetOnDemand(onDemandSizes)
			}
			return nil
		},
		nil,
	)

	fieldsMap[LowPriority] = commons.NewGenericField(
		commons.ElastigroupAzureVMSizes,
		LowPriority,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []string
			if elastigroup.Compute != nil && elastigroup.Compute.VMSizes != nil &&
				elastigroup.Compute.VMSizes.LowPriority != nil {
				result = append(result, elastigroup.Compute.VMSizes.LowPriority...)
				if err := resourceData.Set(string(LowPriority), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(LowPriority), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(LowPriority)); ok {
				virtualMachines := v.([]interface{})
				lowPrioritySizes := make([]string, len(virtualMachines))
				for i, j := range virtualMachines {
					lowPrioritySizes[i] = j.(string)
				}
				elastigroup.Compute.VMSizes.SetLowPriority(lowPrioritySizes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(LowPriority)); ok {
				virtualMachines := v.([]interface{})
				lowPrioritySizes := make([]string, len(virtualMachines))
				for i, j := range virtualMachines {
					lowPrioritySizes[i] = j.(string)
				}
				elastigroup.Compute.VMSizes.SetLowPriority(lowPrioritySizes)
			}
			return nil
		},
		nil,
	)
}

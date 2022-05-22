package stateful_node_azure_vm_sizes

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[OnDemandSizes] = commons.NewGenericField(
		commons.StatefulNodeAzureVMSizes,
		OnDemandSizes,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()

			var result []string
			if statefulNode.Compute != nil && statefulNode.Compute.VMSizes != nil &&
				statefulNode.Compute.VMSizes.OnDemandSizes != nil {
				result = append(result, statefulNode.Compute.VMSizes.OnDemandSizes...)
				if err := resourceData.Set(string(OnDemandSizes), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OnDemandSizes), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(OnDemandSizes)); ok {
				virtualMachines := v.([]interface{})
				onDemandSizes := make([]string, len(virtualMachines))
				for i, j := range virtualMachines {
					onDemandSizes[i] = j.(string)
				}
				statefulNode.Compute.VMSizes.SetOnDemandSizes(onDemandSizes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(OnDemandSizes)); ok {
				virtualMachines := v.([]interface{})
				onDemandSizes := make([]string, len(virtualMachines))
				for i, j := range virtualMachines {
					onDemandSizes[i] = j.(string)
				}
				statefulNode.Compute.VMSizes.SetOnDemandSizes(onDemandSizes)
			}
			return nil
		},
		nil,
	)

	fieldsMap[SpotSizes] = commons.NewGenericField(
		commons.StatefulNodeAzureVMSizes,
		SpotSizes,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var result []string
			if statefulNode.Compute != nil && statefulNode.Compute.VMSizes != nil &&
				statefulNode.Compute.VMSizes.SpotSizes != nil {
				result = append(result, statefulNode.Compute.VMSizes.SpotSizes...)
				if err := resourceData.Set(string(SpotSizes), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SpotSizes), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(SpotSizes)); ok {
				virtualMachines := v.([]interface{})
				spotSizes := make([]string, len(virtualMachines))
				for i, j := range virtualMachines {
					spotSizes[i] = j.(string)
				}
				statefulNode.Compute.VMSizes.SetSpotSizes(spotSizes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(SpotSizes)); ok {
				virtualMachines := v.([]interface{})
				spotSizes := make([]string, len(virtualMachines))
				for i, j := range virtualMachines {
					spotSizes[i] = j.(string)
				}
				statefulNode.Compute.VMSizes.SetSpotSizes(spotSizes)
			}
			return nil
		},
		nil,
	)

	fieldsMap[PreferredSpotSizes] = commons.NewGenericField(
		commons.StatefulNodeAzureVMSizes,
		PreferredSpotSizes,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var result []string
			if statefulNode.Compute != nil && statefulNode.Compute.VMSizes != nil &&
				statefulNode.Compute.VMSizes.PreferredSpotSizes != nil {
				result = append(result, statefulNode.Compute.VMSizes.PreferredSpotSizes...)
				if err := resourceData.Set(string(PreferredSpotSizes), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PreferredSpotSizes), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(PreferredSpotSizes)); ok {
				virtualMachines := v.([]interface{})
				PreferredSpotSizes := make([]string, len(virtualMachines))
				for i, j := range virtualMachines {
					PreferredSpotSizes[i] = j.(string)
				}
				statefulNode.Compute.VMSizes.SetPreferredSpotSizes(PreferredSpotSizes)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(PreferredSpotSizes)); ok {
				virtualMachines := v.([]interface{})
				PreferredSpotSizes := make([]string, len(virtualMachines))
				for i, j := range virtualMachines {
					PreferredSpotSizes[i] = j.(string)
				}
				statefulNode.Compute.VMSizes.SetPreferredSpotSizes(PreferredSpotSizes)
			}
			return nil
		},
		nil,
	)

}

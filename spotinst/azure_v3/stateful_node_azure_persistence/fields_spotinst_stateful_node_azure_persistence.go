package stateful_node_azure_persistence

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ShouldPersistOSDisk] = commons.NewGenericField(
		commons.StatefulNodeAzurePersistence,
		ShouldPersistOSDisk,
		&schema.Schema{
			Type:     schema.TypeBool,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *bool = nil
			if statefulNode.Persistence != nil && statefulNode.Persistence.ShouldPersistOSDisk != nil {
				value = statefulNode.Persistence.ShouldPersistOSDisk
			}
			if err := resourceData.Set(string(ShouldPersistOSDisk), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShouldPersistOSDisk), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(ShouldPersistOSDisk)).(bool); ok {
				statefulNode.Persistence.SetShouldPersistOSDisk(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(ShouldPersistOSDisk)).(bool); ok {
				statefulNode.Persistence.SetShouldPersistOSDisk(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[OSDiskPersistenceMode] = commons.NewGenericField(
		commons.StatefulNodeAzurePersistence,
		OSDiskPersistenceMode,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *string = nil
			if statefulNode.Persistence != nil && statefulNode.Persistence.OSDiskPersistenceMode != nil {
				value = statefulNode.Persistence.OSDiskPersistenceMode
			}
			if err := resourceData.Set(string(OSDiskPersistenceMode), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OSDiskPersistenceMode), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(OSDiskPersistenceMode)).(string); ok && v != "" {
				statefulNode.Persistence.SetOSDiskPersistenceMode(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *string = nil
			if v, ok := resourceData.Get(string(OSDiskPersistenceMode)).(string); ok && v != "" {
				value = spotinst.String(v)
			}
			statefulNode.Persistence.SetOSDiskPersistenceMode(value)
			return nil
		},
		nil,
	)

	fieldsMap[ShouldPersistDataDisks] = commons.NewGenericField(
		commons.StatefulNodeAzurePersistence,
		ShouldPersistDataDisks,
		&schema.Schema{
			Type:     schema.TypeBool,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *bool = nil
			if statefulNode.Persistence != nil && statefulNode.Persistence.ShouldPersistDataDisks != nil {
				value = statefulNode.Persistence.ShouldPersistDataDisks
			}
			if err := resourceData.Set(string(ShouldPersistDataDisks), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShouldPersistDataDisks), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(ShouldPersistDataDisks)).(bool); ok {
				statefulNode.Persistence.SetShouldPersistDataDisks(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(ShouldPersistDataDisks)).(bool); ok {
				statefulNode.Persistence.SetShouldPersistDataDisks(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[DataDisksPersistenceMode] = commons.NewGenericField(
		commons.StatefulNodeAzurePersistence,
		DataDisksPersistenceMode,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *string = nil
			if statefulNode.Persistence != nil && statefulNode.Persistence.DataDisksPersistenceMode != nil {
				value = statefulNode.Persistence.DataDisksPersistenceMode
			}
			if err := resourceData.Set(string(DataDisksPersistenceMode), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DataDisksPersistenceMode), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(DataDisksPersistenceMode)).(string); ok && v != "" {
				statefulNode.Persistence.SetDataDisksPersistenceMode(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *string = nil
			if v, ok := resourceData.Get(string(DataDisksPersistenceMode)).(string); ok && v != "" {
				value = spotinst.String(v)
			}
			statefulNode.Persistence.SetDataDisksPersistenceMode(value)
			return nil
		},
		nil,
	)

	fieldsMap[ShouldPersistNetwork] = commons.NewGenericField(
		commons.StatefulNodeAzurePersistence,
		ShouldPersistNetwork,
		&schema.Schema{
			Type:     schema.TypeBool,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *bool = nil
			if statefulNode.Persistence != nil && statefulNode.Persistence.ShouldPersistNetwork != nil {
				value = statefulNode.Persistence.ShouldPersistNetwork
			}
			if err := resourceData.Set(string(ShouldPersistNetwork), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShouldPersistNetwork), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(ShouldPersistNetwork)).(bool); ok {
				statefulNode.Persistence.SetShouldPersistNetwork(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(ShouldPersistNetwork)).(bool); ok {
				statefulNode.Persistence.SetShouldPersistNetwork(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[ShouldPersistNetwork] = commons.NewGenericField(
		commons.StatefulNodeAzurePersistence,
		ShouldPersistNetwork,
		&schema.Schema{
			Type:     schema.TypeBool,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *bool = nil
			if statefulNode.Persistence != nil && statefulNode.Persistence.ShouldPersistNetwork != nil {
				value = statefulNode.Persistence.ShouldPersistNetwork
			}
			if err := resourceData.Set(string(ShouldPersistNetwork), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShouldPersistNetwork), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(ShouldPersistNetwork)).(bool); ok {
				statefulNode.Persistence.SetShouldPersistNetwork(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(ShouldPersistNetwork)).(bool); ok {
				statefulNode.Persistence.SetShouldPersistNetwork(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[ShouldPersistVM] = commons.NewGenericField(
		commons.StatefulNodeAzurePersistence,
		ShouldPersistVM,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			var value *bool = nil
			if statefulNode.Persistence != nil && statefulNode.Persistence.ShouldPersistVM != nil {
				value = statefulNode.Persistence.ShouldPersistVM
			}
			if err := resourceData.Set(string(ShouldPersistVM), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShouldPersistVM), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(ShouldPersistVM)).(bool); ok {
				statefulNode.Persistence.SetShouldPersistVM(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			snWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := snWrapper.GetStatefulNode()
			if v, ok := resourceData.Get(string(ShouldPersistVM)).(bool); ok {
				statefulNode.Persistence.SetShouldPersistVM(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

}

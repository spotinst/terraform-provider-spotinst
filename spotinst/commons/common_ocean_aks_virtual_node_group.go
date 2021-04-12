package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
)

const OceanAKSVirtualNodeGroupResourceName ResourceName = "spotinst_ocean_aks_virtual_node_group"

var OceanAKSVirtualNodeGroupResource *OceanAKSVirtualNodeGroupTerraformResource

type OceanAKSVirtualNodeGroupTerraformResource struct {
	GenericResource // embedding
}

type VirtualNodeGroupAKSWrapper struct {
	VirtualNodeGroup *azure.VirtualNodeGroup
}

func NewOceanAKSVirtualNodeGroupResource(fieldsMap map[FieldName]*GenericField) *OceanAKSVirtualNodeGroupTerraformResource {
	return &OceanAKSVirtualNodeGroupTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanAKSVirtualNodeGroupResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanAKSVirtualNodeGroupTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*azure.VirtualNodeGroup, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	launchSpecWrapper := NewVirtualNodeGroupAKSWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(launchSpecWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return launchSpecWrapper.GetVirtualNodeGroup(), nil
}

func (res *OceanAKSVirtualNodeGroupTerraformResource) OnRead(
	launchSpec *azure.VirtualNodeGroup,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	launchSpecWrapper := NewVirtualNodeGroupAKSWrapper()
	launchSpecWrapper.SetVirtualNodeGroup(launchSpec)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(launchSpecWrapper, resourceData, meta); err != nil {
			return err
		}
	}

	return nil
}

func (res *OceanAKSVirtualNodeGroupTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *azure.VirtualNodeGroup, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	launchSpecWrapper := NewVirtualNodeGroupAKSWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(launchSpecWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, launchSpecWrapper.GetVirtualNodeGroup(), nil
}

func NewVirtualNodeGroupAKSWrapper() *VirtualNodeGroupAKSWrapper {
	return &VirtualNodeGroupAKSWrapper{
		VirtualNodeGroup: &azure.VirtualNodeGroup{
			LaunchSpecification: &azure.VirtualNodeGroupLaunchSpecification{
				OSDisk: &azure.OSDisk{},
				Tags:   []*azure.Tag{},
			},
			AutoScale: &azure.VirtualNodeGroupAutoScale{
				Headrooms: []*azure.VirtualNodeGroupHeadroom{},
			},
			ResourceLimits: &azure.VirtualNodeGroupResourceLimits{},
			Taints:         []*azure.Taint{},
			Labels:         []*azure.Label{},
		},
	}
}

func (launchSpecWrapper *VirtualNodeGroupAKSWrapper) GetVirtualNodeGroup() *azure.VirtualNodeGroup {
	return launchSpecWrapper.VirtualNodeGroup
}

func (launchSpecWrapper *VirtualNodeGroupAKSWrapper) SetVirtualNodeGroup(launchSpec *azure.VirtualNodeGroup) {
	launchSpecWrapper.VirtualNodeGroup = launchSpec
}

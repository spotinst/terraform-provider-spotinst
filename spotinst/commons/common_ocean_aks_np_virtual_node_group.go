package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
)

const OceanAKSNPVirtualNodeGroupResourceName ResourceName = "spotinst_ocean_aks_np_virtual_node_group"

var OceanAKSNPVirtualNodeGroupResource *OceanAKSNPVirtualNodeGroupTerraformResource

type OceanAKSNPVirtualNodeGroupTerraformResource struct {
	GenericResource
}

type VirtualNodeGroupAKSNPWrapper struct {
	VirtualNodeGroup *azure_np.VirtualNodeGroup
}

func NewOceanAKSNPVirtualNodeGroupResource(fieldsMap map[FieldName]*GenericField) *OceanAKSNPVirtualNodeGroupTerraformResource {
	return &OceanAKSNPVirtualNodeGroupTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanAKSNPVirtualNodeGroupResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanAKSNPVirtualNodeGroupTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*azure_np.VirtualNodeGroup, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	vngWrapper := NewVirtualNodeGroupAKSNPWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(vngWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return vngWrapper.GetVirtualNodeGroup(), nil
}

func (res *OceanAKSNPVirtualNodeGroupTerraformResource) OnRead(
	vng *azure_np.VirtualNodeGroup,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	vngWrapper := NewVirtualNodeGroupAKSNPWrapper()
	vngWrapper.SetVirtualNodeGroup(vng)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(vngWrapper, resourceData, meta); err != nil {
			return err
		}
	}

	return nil
}

func (res *OceanAKSNPVirtualNodeGroupTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}, conditionParam []interface{}) (bool, bool, *azure_np.VirtualNodeGroup, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	vngWrapper := NewVirtualNodeGroupAKSNPWrapper()
	hasChanged := false
	changesRequiredRoll := false

	if len(conditionParam) > 0 {
		conditionedRollParams := make([]string, len(conditionParam))
		for i, v := range conditionParam {
			conditionedRollParams[i] = fmt.Sprint(v)
		}
		conditionedRollFieldsAKS = conditionedRollParams
	}

	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			if contains(conditionedRollFieldsAKS, field.fieldNameStr) {
				changesRequiredRoll = true
			}

			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(vngWrapper, resourceData, meta); err != nil {
				return false, false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, changesRequiredRoll, vngWrapper.GetVirtualNodeGroup(), nil
}

func NewVirtualNodeGroupAKSNPWrapper() *VirtualNodeGroupAKSNPWrapper {
	return &VirtualNodeGroupAKSNPWrapper{
		VirtualNodeGroup: &azure_np.VirtualNodeGroup{
			NodePoolProperties: &azure_np.NodePoolProperties{},
			NodeCountLimits:    &azure_np.NodeCountLimits{},
			Strategy:           &azure_np.Strategy{},
			AutoScale:          &azure_np.AutoScale{},
			VmSizes:            &azure_np.VmSizes{},
		},
	}
}

func (vngWrapper *VirtualNodeGroupAKSNPWrapper) GetVirtualNodeGroup() *azure_np.VirtualNodeGroup {
	return vngWrapper.VirtualNodeGroup
}

func (vngWrapper *VirtualNodeGroupAKSNPWrapper) SetVirtualNodeGroup(vng *azure_np.VirtualNodeGroup) {
	vngWrapper.VirtualNodeGroup = vng
}

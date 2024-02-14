package commons

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"log"
)

const OceanAKSNPResourceName ResourceName = "spotinst_ocean_aks_np"

var OceanAKSNPResource *OceanAKSNPTerraformResource

type OceanAKSNPTerraformResource struct {
	GenericResource
}

type AKSNPClusterWrapper struct {
	cluster *azure_np.Cluster
}

func NewOceanAKSNPResource(fieldsMap map[FieldName]*GenericField) *OceanAKSNPTerraformResource {
	return &OceanAKSNPTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanAKSNPResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanAKSNPTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*azure_np.Cluster, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	clusterWrapper := NewAKSNPClusterWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(clusterWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return clusterWrapper.GetNPCluster(), nil
}

func (res *OceanAKSNPTerraformResource) OnRead(
	cluster *azure_np.Cluster,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	clusterWrapper := NewAKSNPClusterWrapper()
	clusterWrapper.SetNPCluster(cluster)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(clusterWrapper, resourceData, meta); err != nil {
			return err
		}
	}

	return nil
}

func (res *OceanAKSNPTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, bool, *azure_np.Cluster, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	clusterWrapper := NewAKSNPClusterWrapper()
	hasChanged := false
	changesRequiredRoll := false

	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			if contains(conditionedRollFieldsAKS, field.fieldNameStr) {
				changesRequiredRoll = true
			}

			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(clusterWrapper, resourceData, meta); err != nil {
				return false, false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, changesRequiredRoll, clusterWrapper.GetNPCluster(), nil
}

func NewAKSNPClusterWrapper() *AKSNPClusterWrapper {
	return &AKSNPClusterWrapper{
		cluster: &azure_np.Cluster{
			AKS: &azure_np.AKS{},
			VirtualNodeGroupTemplate: &azure_np.VirtualNodeGroupTemplate{
				NodePoolProperties: &azure_np.NodePoolProperties{},
				NodeCountLimits:    &azure_np.NodeCountLimits{},
				Strategy:           &azure_np.Strategy{},
				AutoScale:          &azure_np.AutoScale{},
				VmSizes:            &azure_np.VmSizes{},
			},
		},
	}
}

func (clusterWrapper *AKSNPClusterWrapper) GetNPCluster() *azure_np.Cluster {
	return clusterWrapper.cluster
}

func (clusterWrapper *AKSNPClusterWrapper) SetNPCluster(cluster *azure_np.Cluster) {
	clusterWrapper.cluster = cluster
}

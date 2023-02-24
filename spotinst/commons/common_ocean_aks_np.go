package commons

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

	/*if importedCluster != nil {
		// This is the merge part of the import action
		// onCreate on every field is the override action on top of what returned from Spot API
		buildEmptyAKSNPClusterImportRequirements(importedCluster)
		clusterWrapper.SetNPCluster(importedCluster)
	}*/

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
	meta interface{}) (bool, *azure_np.Cluster, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	clusterWrapper := NewAKSNPClusterWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(clusterWrapper, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}

	return hasChanged, clusterWrapper.GetNPCluster(), nil
}

func NewAKSNPClusterWrapper() *AKSNPClusterWrapper {
	return &AKSNPClusterWrapper{
		cluster: &azure_np.Cluster{
			AKS: &azure_np.AKS{},
			VirtualNodeGroupTemplate: &azure_np.VirtualNodeGroupTemplate{
				NodePoolProperties: &azure_np.NodePoolProperties{},
				NodeCountLimits:    &azure_np.NodeCountLimits{},
				Strategy:           &azure_np.Strategy{},
				Tags:               []*azure_np.Tag{},
				Labels:             []*azure_np.Label{},
				Taints:             []*azure_np.Taint{},
			},
			Health: &azure_np.Health{},
			AutoScaler: &azure_np.AutoScaler{
				ResourceLimits: &azure_np.ResourceLimits{},
				Down:           &azure_np.Down{},
				Headroom: &azure_np.Headroom{
					Automatic: &azure_np.Automatic{},
				},
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

func buildEmptyAKSNPClusterImportRequirements(cluster *azure_np.Cluster) {
	if cluster == nil {
		return
	}

	if cluster.AutoScaler == nil {
		cluster.SetAutoScaler(&azure_np.AutoScaler{})
	}

	if cluster.Health == nil {
		cluster.SetHealth(&azure_np.Health{})
	}
}

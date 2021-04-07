package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
)

const OceanAKSResourceName ResourceName = "spotinst_ocean_aks"

var OceanAKSResource *OceanAKSTerraformResource

type OceanAKSTerraformResource struct {
	GenericResource // embedding
}

type AKSClusterWrapper struct {
	cluster *azure.Cluster
}

func NewOceanAKSResource(fieldsMap map[FieldName]*GenericField) *OceanAKSTerraformResource {
	return &OceanAKSTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanAKSResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanAKSTerraformResource) OnCreate(
	importedCluster *azure.Cluster,
	resourceData *schema.ResourceData,
	meta interface{}) (*azure.Cluster, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	clusterWrapper := NewAKSClusterWrapper()

	if importedCluster != nil {
		// This is the merge part of the import action
		// onCreate on every field is the override action on top of what returned from Spot API
		buildEmptyAKSClusterImportRequirements(importedCluster)
		clusterWrapper.SetCluster(importedCluster)
	}

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(clusterWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return clusterWrapper.GetCluster(), nil
}

func (res *OceanAKSTerraformResource) OnRead(
	cluster *azure.Cluster,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	clusterWrapper := NewAKSClusterWrapper()
	clusterWrapper.SetCluster(cluster)

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

func (res *OceanAKSTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *azure.Cluster, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	clusterWrapper := NewAKSClusterWrapper()
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

	return hasChanged, clusterWrapper.GetCluster(), nil
}

func NewAKSClusterWrapper() *AKSClusterWrapper {
	return &AKSClusterWrapper{
		cluster: &azure.Cluster{
			VirtualNodeGroupTemplate: &azure.VirtualNodeGroupTemplate{
				LaunchSpecification: &azure.LaunchSpecification{
					Login: &azure.Login{},
					Image: &azure.Image{
						MarketplaceImage: &azure.MarketplaceImage{},
					},
					Extensions: []*azure.Extension{},
					Network: &azure.Network{
						NetworkInterfaces: []*azure.NetworkInterface{},
					},
					LoadBalancersConfig: &azure.LoadBalancersConfig{
						LoadBalancers: []*azure.LoadBalancer{},
					},
					Tags: []*azure.Tag{},
				},
				VMSizes: &azure.VMSizes{},
			},
			Strategy: &azure.Strategy{},
			Health:   &azure.Health{},
			AutoScaler: &azure.AutoScaler{
				ResourceLimits: &azure.ResourceLimits{},
				Down:           &azure.Down{},
				Headroom: &azure.Headroom{
					Automatic: &azure.Automatic{},
				},
			},
		},
	}
}

func (clusterWrapper *AKSClusterWrapper) GetCluster() *azure.Cluster {
	return clusterWrapper.cluster
}

func (clusterWrapper *AKSClusterWrapper) SetCluster(cluster *azure.Cluster) {
	clusterWrapper.cluster = cluster
}

func buildEmptyAKSClusterImportRequirements(cluster *azure.Cluster) {
	if cluster == nil {
		return
	}

	if cluster.Strategy == nil {
		cluster.SetStrategy(&azure.Strategy{})
	}

	if cluster.AutoScaler == nil {
		cluster.SetAutoScaler(&azure.AutoScaler{})
	}

	if cluster.Health == nil {
		cluster.SetHealth(&azure.Health{})
	}
}

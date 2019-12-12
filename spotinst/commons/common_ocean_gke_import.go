package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	OceanGKEImportResourceName ResourceName = "spotinst_ocean_gke_import"
)

var OceanGKEImportResource *OceanGKEImportTerraformResource

type OceanGKEImportTerraformResource struct {
	GenericResource // embedding
}

type GKEImportClusterWrapper struct {
	cluster *gcp.Cluster
}

func NewOceanGKEImportResource(fieldsMap map[FieldName]*GenericField) *OceanGKEImportTerraformResource {
	return &OceanGKEImportTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanGKEImportResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanGKEImportTerraformResource) OnCreate(
	importedCluster *gcp.Cluster,
	resourceData *schema.ResourceData,
	meta interface{}) (*gcp.Cluster, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	clusterWrapper := NewGKEImportClusterWrapper()

	if importedCluster != nil {
		// This is the merge part of the import action
		// onCreate on every field is the override action on top of what returned from Spotinst API
		buildEmptyClusterImportRequirements(importedCluster)
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

func (res *OceanGKEImportTerraformResource) OnRead(
	cluster *gcp.Cluster,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	clusterWrapper := NewGKEImportClusterWrapper()
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

func (res *OceanGKEImportTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *gcp.Cluster, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	clusterWrapper := NewGKEImportClusterWrapper()
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

func NewGKEImportClusterWrapper() *GKEImportClusterWrapper {
	return &GKEImportClusterWrapper{
		cluster: &gcp.Cluster{
			Capacity: &gcp.Capacity{},
			Compute: &gcp.Compute{
				LaunchSpecification: &gcp.LaunchSpecification{},
				InstanceTypes:       &gcp.InstanceTypes{},
			},
		},
	}
}

func (clusterWrapper *GKEImportClusterWrapper) GetCluster() *gcp.Cluster {
	return clusterWrapper.cluster
}

func (clusterWrapper *GKEImportClusterWrapper) SetCluster(cluster *gcp.Cluster) {
	clusterWrapper.cluster = cluster
}

func buildEmptyClusterImportRequirements(cluster *gcp.Cluster) {
	if cluster == nil {
		return
	}

	if cluster.Compute == nil {
		cluster.SetCompute(&gcp.Compute{})
	}

	if cluster.Compute.InstanceTypes == nil {
		cluster.Compute.SetInstanceTypes(&gcp.InstanceTypes{})
	}
}

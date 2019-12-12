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
	OceanGKEResourceName ResourceName = "spotinst_ocean_gke"
)

var OceanGKEResource *OceanGKETerraformResource

type OceanGKETerraformResource struct {
	GenericResource // embedding
}

type GKEClusterWrapper struct {
	cluster *gcp.Cluster
}

func NewOceanGKEResource(fieldsMap map[FieldName]*GenericField) *OceanGKETerraformResource {
	return &OceanGKETerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanGKEResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanGKETerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*gcp.Cluster, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	clusterWrapper := NewGKEClusterWrapper()

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

func (res *OceanGKETerraformResource) OnRead(
	cluster *gcp.Cluster,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	clusterWrapper := NewGKEClusterWrapper()
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

func (res *OceanGKETerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *gcp.Cluster, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	clusterWrapper := NewGKEClusterWrapper()
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

func NewGKEClusterWrapper() *GKEClusterWrapper {
	return &GKEClusterWrapper{
		cluster: &gcp.Cluster{
			Capacity: &gcp.Capacity{},
			Compute: &gcp.Compute{
				LaunchSpecification: &gcp.LaunchSpecification{},
				InstanceTypes:       &gcp.InstanceTypes{},
			},
			Strategy: &gcp.Strategy{},
		},
	}
}

func (clusterWrapper *GKEClusterWrapper) GetCluster() *gcp.Cluster {
	return clusterWrapper.cluster
}

func (clusterWrapper *GKEClusterWrapper) SetCluster(cluster *gcp.Cluster) {
	clusterWrapper.cluster = cluster
}

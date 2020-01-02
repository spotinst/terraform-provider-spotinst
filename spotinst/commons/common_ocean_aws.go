package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Variables
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
const (
	OceanAWSResourceName ResourceName = "spotinst_ocean_aws"
)

var OceanAWSResource *OceanAWSTerraformResource

type OceanAWSTerraformResource struct {
	GenericResource // embedding
}

type AWSClusterWrapper struct {
	cluster *aws.Cluster
}

func NewOceanAWSResource(fieldsMap map[FieldName]*GenericField) *OceanAWSTerraformResource {
	return &OceanAWSTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanAWSResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *OceanAWSTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.Cluster, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	clusterWrapper := NewClusterWrapper()

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

func (res *OceanAWSTerraformResource) OnRead(
	cluster *aws.Cluster,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	clusterWrapper := NewClusterWrapper()
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

func (res *OceanAWSTerraformResource) OnUpdate(
	resourceData *schema.ResourceData,
	meta interface{}) (bool, *aws.Cluster, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	clusterWrapper := NewClusterWrapper()
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

func NewClusterWrapper() *AWSClusterWrapper {
	return &AWSClusterWrapper{
		cluster: &aws.Cluster{
			Capacity: &aws.Capacity{},
			Compute: &aws.Compute{
				LaunchSpecification: &aws.LaunchSpecification{},
				InstanceTypes:       &aws.InstanceTypes{},
			},
			Strategy: &aws.Strategy{},
		},
	}
}

func (clusterWrapper *AWSClusterWrapper) GetCluster() *aws.Cluster {
	return clusterWrapper.cluster
}

func (clusterWrapper *AWSClusterWrapper) SetCluster(cluster *aws.Cluster) {
	clusterWrapper.cluster = cluster
}

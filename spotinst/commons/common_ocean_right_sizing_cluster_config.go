package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/cluster_right_sizing"
)

const (
	OceanRightSizingClusterConfigResourceName ResourceName = "spotinst_ocean_right_sizing_cluster_config"
)

var OceanRightSizingClusterConfigResource *OceanRightSizingClusterConfigTerraformResource

type OceanRightSizingClusterConfigTerraformResource struct {
	GenericResource
}

type RightSizingClusterConfigWrapper struct {
	postClusterConfigurationInput *cluster_right_sizing.PostClusterConfigurationInput
	clusterConfiguration          *cluster_right_sizing.ClusterConfiguration
}

func NewOceanRightSizingClusterConfigResource(fieldMap map[FieldName]*GenericField) *OceanRightSizingClusterConfigTerraformResource {
	return &OceanRightSizingClusterConfigTerraformResource{
		GenericResource: GenericResource{
			resourceName: OceanRightSizingClusterConfigResourceName,
			fields:       NewGenericFields(fieldMap),
		},
	}
}

func (res *OceanRightSizingClusterConfigTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*cluster_right_sizing.PostClusterConfigurationInput, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	wrapper := NewRightSizingClusterConfigWrapper()
	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(wrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}

	return wrapper.GetPostClusterConfigurationInput(), nil
}

func (res *OceanRightSizingClusterConfigTerraformResource) OnRead(
	clusterConfiguration *cluster_right_sizing.ClusterConfiguration,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	wrapper := NewRightSizingClusterConfigWrapper()
	wrapper.SetClusterConfiguration(clusterConfiguration)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(wrapper, resourceData, meta); err != nil {
			return err
		}
	}

	return nil
}

func NewRightSizingClusterConfigWrapper() *RightSizingClusterConfigWrapper {
	return &RightSizingClusterConfigWrapper{
		postClusterConfigurationInput: &cluster_right_sizing.PostClusterConfigurationInput{},
		clusterConfiguration:          &cluster_right_sizing.ClusterConfiguration{},
	}
}

func (w *RightSizingClusterConfigWrapper) GetPostClusterConfigurationInput() *cluster_right_sizing.PostClusterConfigurationInput {
	return w.postClusterConfigurationInput
}

func (w *RightSizingClusterConfigWrapper) SetPostClusterConfigurationInput(input *cluster_right_sizing.PostClusterConfigurationInput) {
	w.postClusterConfigurationInput = input
}

func (w *RightSizingClusterConfigWrapper) GetClusterConfiguration() *cluster_right_sizing.ClusterConfiguration {
	return w.clusterConfiguration
}

func (w *RightSizingClusterConfigWrapper) SetClusterConfiguration(config *cluster_right_sizing.ClusterConfiguration) {
	w.clusterConfiguration = config
}

package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/right_sizing_cluster_config"
)

const (
	OceanRightSizingClusterConfigResourceName ResourceName = "spotinst_ocean_right_sizing_cluster_config"
)

var OceanRightSizingClusterConfigResource *OceanRightSizingClusterConfigTerraformResource

type OceanRightSizingClusterConfigTerraformResource struct {
	GenericResource
}

type RightSizingClusterConfigWrapper struct {
	RightsizingClusterConfigurationInput *right_sizing_cluster_config.RightsizingClusterConfigurationInput
	RightsizingClusterConfiguration      *right_sizing_cluster_config.RightsizingClusterConfiguration
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
	meta interface{}) (*right_sizing_cluster_config.RightsizingClusterConfigurationInput, error) {

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

	return wrapper.GetRightsizingClusterConfigurationInput(), nil
}

func (res *OceanRightSizingClusterConfigTerraformResource) OnRead(
	RightsizingClusterConfiguration *right_sizing_cluster_config.RightsizingClusterConfiguration,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	wrapper := NewRightSizingClusterConfigWrapper()
	wrapper.SetRightsizingClusterConfiguration(RightsizingClusterConfiguration)

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
		RightsizingClusterConfigurationInput: &right_sizing_cluster_config.RightsizingClusterConfigurationInput{},
		RightsizingClusterConfiguration:      &right_sizing_cluster_config.RightsizingClusterConfiguration{},
	}
}

func (w *RightSizingClusterConfigWrapper) GetRightsizingClusterConfigurationInput() *right_sizing_cluster_config.RightsizingClusterConfigurationInput {
	return w.RightsizingClusterConfigurationInput
}

func (w *RightSizingClusterConfigWrapper) SetRightsizingClusterConfigurationInput(input *right_sizing_cluster_config.RightsizingClusterConfigurationInput) {
	w.RightsizingClusterConfigurationInput = input
}

func (w *RightSizingClusterConfigWrapper) GetRightsizingClusterConfiguration() *right_sizing_cluster_config.RightsizingClusterConfiguration {
	return w.RightsizingClusterConfiguration
}

func (w *RightSizingClusterConfigWrapper) SetRightsizingClusterConfiguration(config *right_sizing_cluster_config.RightsizingClusterConfiguration) {
	w.RightsizingClusterConfiguration = config
}

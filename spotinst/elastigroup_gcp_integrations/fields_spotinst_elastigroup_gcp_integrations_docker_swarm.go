package elastigroup_gcp_integrations

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupDockerSwarm(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[IntegrationDockerSwarm] = commons.NewGenericField(
		commons.ElastigroupGCPIntegrations,
		IntegrationDockerSwarm,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MasterHost): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(MasterPort): {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(IntegrationDockerSwarm)); ok {
				if integration, err := expandGCPGroupDockerSwarmIntegration(v); err != nil {
					return err
				} else {
					elastigroup.Integration.SetDockerSwarm(integration)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *gcp.DockerSwarmIntegration = nil

			if v, ok := resourceData.GetOk(string(IntegrationDockerSwarm)); ok {
				if integration, err := expandGCPGroupDockerSwarmIntegration(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetDockerSwarm(value)
			return nil
		},

		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandGCPGroupDockerSwarmIntegration(data interface{}) (*gcp.DockerSwarmIntegration, error) {
	integration := &gcp.DockerSwarmIntegration{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return integration, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(MasterHost)].(string); ok && v != "" {
		integration.SetMasterHost(spotinst.String(v))
	}

	if v, ok := m[string(MasterPort)].(int); ok && v > 0 {
		integration.SetMasterPort(spotinst.Int(v))
	}

	return integration, nil
}

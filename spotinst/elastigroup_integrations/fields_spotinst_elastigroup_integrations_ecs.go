package elastigroup_integrations

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupEcs(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationEcs] = commons.NewGenericField(
		commons.ElastigroupIntegrations,
		IntegrationEcs,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ClusterName): &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},

					string(AutoscaleIsEnabled): &schema.Schema{
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(AutoscaleCooldown): &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(AutoscaleHeadroom): &schema.Schema{
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(CpuPerUnit): &schema.Schema{
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(MemoryPerUnit): &schema.Schema{
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(NumOfUnits): &schema.Schema{
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},

					string(AutoscaleDown): &schema.Schema{
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(EvaluationPeriods): &schema.Schema{
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			if v, ok := resourceData.GetOk(string(IntegrationEcs)); ok {
				if integration, err := expandAWSGroupEC2ContainerServiceIntegration(v); err != nil {
					return err
				} else {
					elastigroup.Integration.SetEC2ContainerService(integration)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			elastigroup := resourceObject.(*aws.Group)
			var value *aws.EC2ContainerServiceIntegration = nil
			if v, ok := resourceData.GetOk(string(IntegrationEcs)); ok {
				if integration, err := expandAWSGroupEC2ContainerServiceIntegration(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetEC2ContainerService(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupEC2ContainerServiceIntegration(data interface{}) (*aws.EC2ContainerServiceIntegration, error) {
	integration := &aws.EC2ContainerServiceIntegration{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return integration, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(ClusterName)].(string); ok && v != "" {
		integration.SetClusterName(spotinst.String(v))
	}

	if v, ok := m[string(AutoscaleIsEnabled)].(bool); ok {
		if integration.AutoScale == nil {
			integration.SetAutoScale(&aws.AutoScale{})
		}
		integration.AutoScale.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoscaleCooldown)].(int); ok && v > 0 {
		if integration.AutoScale == nil {
			integration.SetAutoScale(&aws.AutoScale{})
		}
		integration.AutoScale.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m[string(AutoscaleHeadroom)]; ok {
		headroom, err := expandAWSGroupAutoScaleHeadroom(v)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			if integration.AutoScale == nil {
				integration.SetAutoScale(&aws.AutoScale{})
			}
			integration.AutoScale.SetHeadroom(headroom)
		}
	}

	if v, ok := m[string(AutoscaleDown)]; ok {
		down, err := expandAWSGroupAutoScaleDown(v)
		if err != nil {
			return nil, err
		}
		if down != nil {
			if integration.AutoScale == nil {
				integration.SetAutoScale(&aws.AutoScale{})
			}
			integration.AutoScale.SetDown(down)
		}
	}
	return integration, nil
}

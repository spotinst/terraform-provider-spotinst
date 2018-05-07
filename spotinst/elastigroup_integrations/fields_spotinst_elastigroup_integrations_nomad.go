package elastigroup_integrations

import (
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupNomad(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationNomad] = commons.NewGenericField(
		commons.ElastigroupIntegrations,
		IntegrationNomad,
		&schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			//var value []interface{} = nil
			//if elastigroup.Integration != nil && elastigroup.Integration.Nomad != nil {
			//	value = flattenAWSGroupRancherIntegration(elastigroup.Integration.Rancher)
			//}
			//if value != nil {
			//	if err := resourceData.Set(string(IntegrationNomad), value); err != nil {
			//		return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationNomad), err)
			//	}
			//} else {
			//	if err := resourceData.Set(string(IntegrationNomad), []*aws.NomadIntegration{}); err != nil {
			//		return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IntegrationNomad), err)
			//	}
			//}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			if v, ok := resourceData.GetOk(string(IntegrationNomad)); ok {
				if integration, err := expandAWSGroupNomadIntegration(v, false); err != nil {
					return err
				} else {
					elastigroup.Integration.SetNomad(integration)
				}
			}
			return nil
		},
		func(elastigroup *aws.Group, resourceData *schema.ResourceData, meta interface{}) error {
			var value *aws.NomadIntegration = nil
			if v, ok := resourceData.GetOk(string(IntegrationNomad)); ok {
				if integration, err := expandAWSGroupNomadIntegration(v, true); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetNomad(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupNomadIntegration(data interface{}, nullify bool) (*aws.NomadIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &aws.NomadIntegration{}

	if v, ok := m[string(MasterHost)].(string); ok && v != "" {
		i.SetMasterHost(spotinst.String(v))
	}

	if v, ok := m[string(MasterPort)].(int); ok && v > 0 {
		i.SetMasterPort(spotinst.Int(v))
	}

	if v, ok := m[string(AclToken)].(string); ok && v != "" {
		i.SetAclToken(spotinst.String(v))
	} else if nullify {
		i.SetAclToken(nil)
	}

	if v, ok := m[string(AutoscaleIsEnabled)].(bool); ok {
		if i.AutoScale == nil {
			i.SetAutoScale(&aws.AutoScale{})
		}
		i.AutoScale.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoscaleCooldown)].(int); ok && v > 0 {
		if i.AutoScale == nil {
			i.SetAutoScale(&aws.AutoScale{})
		}
		i.AutoScale.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m[string(AutoscaleHeadroom)]; ok {
		headroom, err := expandAWSGroupAutoScaleHeadroom(v)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			if i.AutoScale == nil {
				i.SetAutoScale(&aws.AutoScale{})
			}
			i.AutoScale.SetHeadroom(headroom)
		}
	}

	if v, ok := m[string(AutoscaleDown)]; ok {
		down, err := expandAWSGroupAutoScaleDown(v)
		if err != nil {
			return nil, err
		}
		if down != nil {
			if i.AutoScale == nil {
				i.SetAutoScale(&aws.AutoScale{})
			}
			i.AutoScale.SetDown(down)
		}
	}

	if v, ok := m[string(AutoscaleConstraints)]; ok {
		consts, err := expandAWSGroupAutoScaleConstraints(v)
		if err != nil {
			return nil, err
		}
		if consts != nil {
			if i.AutoScale == nil {
				i.SetAutoScale(&aws.AutoScale{})
			}
			i.AutoScale.SetConstraints(consts)
		}
	}
	return i, nil
}
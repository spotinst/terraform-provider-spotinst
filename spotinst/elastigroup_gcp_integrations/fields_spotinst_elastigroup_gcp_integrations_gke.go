package elastigroup_gcp_integrations

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupGKE(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[IntegrationGKE] = commons.NewGenericField(
		commons.ElastigroupGCPIntegrations,
		IntegrationGKE,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(ClusterID): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(Location): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(AutoUpdate): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(AutoscaleIsEnabled): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(AutoscaleCooldown): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(AutoscaleIsAutoConfig): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(AutoscaleHeadroom): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(CpuPerUnit): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(MemoryPerUnit): {
									Type:     schema.TypeInt,
									Optional: true,
								},

								string(NumOfUnits): {
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},

					string(AutoscaleDown): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(EvaluationPeriods): {
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},

					string(AutoscaleLabels): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Key): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(Value): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
						Set: labelHashKV,
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
			if v, ok := resourceData.GetOk(string(IntegrationGKE)); ok {
				if integration, err := ExpandGCPGroupGKEIntegration(v); err != nil {
					return err
				} else {
					elastigroup.Integration.SetGKE(integration)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *gcp.GKEIntegration = nil
			if v, ok := resourceData.GetOk(string(IntegrationGKE)); ok {
				if integration, err := ExpandGCPGroupGKEIntegration(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetGKE(value)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func ExpandGCPGroupGKEIntegration(data interface{}) (*gcp.GKEIntegration, error) {
	integration := &gcp.GKEIntegration{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return integration, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(AutoUpdate)].(bool); ok {
		integration.SetAutoUpdate(spotinst.Bool(v))
	}

	if v, ok := m[string(Location)].(string); ok && v != "" {
		integration.SetLocation(spotinst.String(v))
	}

	if v, ok := m[string(ClusterID)].(string); ok && v != "" {
		integration.SetClusterID(spotinst.String(v))
	}

	if v, ok := m[string(AutoscaleIsEnabled)].(bool); ok {
		if integration.AutoScale == nil {
			integration.SetAutoScale(&gcp.AutoScaleGKE{})
		}
		integration.AutoScale.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoscaleCooldown)].(int); ok && v > 0 {
		if integration.AutoScale == nil {
			integration.SetAutoScale(&gcp.AutoScaleGKE{})
		}
		integration.AutoScale.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m[string(AutoscaleIsAutoConfig)].(bool); ok {
		if integration.AutoScale == nil {
			integration.SetAutoScale(&gcp.AutoScaleGKE{})
		}
		integration.AutoScale.SetIsAutoConfig(spotinst.Bool(v))
	}

	if v, ok := m[string(AutoscaleHeadroom)]; ok {
		headroom, err := expandGCPGroupAutoScaleHeadroom(v)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			if integration.AutoScale == nil {
				integration.SetAutoScale(&gcp.AutoScaleGKE{})
			}
			integration.AutoScale.SetHeadroom(headroom)
		}
	}

	if v, ok := m[string(AutoscaleDown)]; ok {
		down, err := expandGCPGroupAutoScaleDown(v)
		if err != nil {
			return nil, err
		}
		if down != nil {
			if integration.AutoScale == nil {
				integration.SetAutoScale(&gcp.AutoScaleGKE{})
			}
			integration.AutoScale.SetDown(down)
		}
	}

	if v, ok := m[string(AutoscaleLabels)]; ok {
		labels, err := expandGKEAutoScaleLabels(v)
		if err != nil {
			return nil, err
		}
		if labels != nil {
			if integration.AutoScale == nil {
				integration.SetAutoScale(&gcp.AutoScaleGKE{})
			}
			integration.AutoScale.SetLabels(labels)
		}
	}
	return integration, nil
}

func expandGKEAutoScaleLabels(data interface{}) ([]*gcp.AutoScaleLabel, error) {
	list := data.(*schema.Set).List()
	out := make([]*gcp.AutoScaleLabel, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(Key)]; !ok {
			return nil, errors.New("invalid GKE label: key missing")
		}

		if _, ok := attr[string(Value)]; !ok {
			return nil, errors.New("invalid GKE label: value missing")
		}
		c := &gcp.AutoScaleLabel{
			Key:   spotinst.String(attr[string(Key)].(string)),
			Value: spotinst.String(attr[string(Value)].(string)),
		}
		out = append(out, c)
	}
	return out, nil
}

func labelHashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(Key)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(Value)].(string)))
	return hashcode.String(buf.String())
}

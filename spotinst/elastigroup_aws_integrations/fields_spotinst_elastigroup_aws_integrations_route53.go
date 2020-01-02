package elastigroup_aws_integrations

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func SetupRoute53(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[IntegrationRoute53] = commons.NewGenericField(
		commons.ElastigroupAWSIntegrations,
		IntegrationRoute53,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Domains): {
						Type:     schema.TypeSet,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(HostedZoneId): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(SpotinstAcctID): {
									Type:     schema.TypeString,
									Optional: true,
								},

								string(RecordSets): {
									Type:     schema.TypeSet,
									Required: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(UsePublicIP): {
												Type:     schema.TypeBool,
												Optional: true,
											},

											string(Name): {
												Type:     schema.TypeString,
												Required: true,
											},
										},
									},
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
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(IntegrationRoute53)); ok {
				if integration, err := expandAWSGroupRoute53Integration(v); err != nil {
					return err
				} else {
					elastigroup.Integration.SetRoute53(integration)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.Route53Integration = nil

			if v, ok := resourceData.GetOk(string(IntegrationRoute53)); ok {
				if integration, err := expandAWSGroupRoute53Integration(v); err != nil {
					return err
				} else {
					value = integration
				}
			}
			elastigroup.Integration.SetRoute53(value)
			return nil
		},

		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupRoute53Integration(data interface{}) (*aws.Route53Integration, error) {
	integration := &aws.Route53Integration{}
	list := data.([]interface{})

	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(Domains)]; ok {
			domains, err := expandAWSGroupRoute53IntegrationDomains(v)

			if err != nil {
				return nil, err
			}
			integration.SetDomains(domains)
		}
	}
	return integration, nil
}

func expandAWSGroupRoute53IntegrationDomains(data interface{}) ([]*aws.Domain, error) {
	list := data.(*schema.Set).List()
	domains := make([]*aws.Domain, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		domain := &aws.Domain{}

		if !ok {
			continue
		}

		if v, ok := attr[string(HostedZoneId)].(string); ok && v != "" {
			domain.SetHostedZoneID(spotinst.String(v))
		}

		if v, ok := attr[string(SpotinstAcctID)].(string); ok && v != "" {
			domain.SetSpotinstAccountID(spotinst.String(v))
		}

		if r, ok := attr[string(RecordSets)]; ok {
			if recordSets, err := expandAWSGroupRoute53IntegrationDomainsRecordSets(r); err != nil {
				return nil, err
			} else {
				domain.SetRecordSets(recordSets)
			}
		}
		domains = append(domains, domain)
	}
	return domains, nil
}

func expandAWSGroupRoute53IntegrationDomainsRecordSets(data interface{}) ([]*aws.RecordSet, error) {
	list := data.(*schema.Set).List()
	recordSets := make([]*aws.RecordSet, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})

		if !ok {
			continue
		}

		if _, ok := attr[string(UsePublicIP)]; !ok {
			return nil, errors.New("invalid record set attributes: use_public_ip missing")
		}

		if _, ok := attr[string(Name)]; !ok {
			return nil, errors.New("invalid record set attributes: name missing")
		}

		recordSet := &aws.RecordSet{
			UsePublicIP: spotinst.Bool(attr[string(UsePublicIP)].(bool)),
			Name:        spotinst.String(attr[string(Name)].(string)),
		}

		recordSets = append(recordSets, recordSet)
	}
	return recordSets, nil
}

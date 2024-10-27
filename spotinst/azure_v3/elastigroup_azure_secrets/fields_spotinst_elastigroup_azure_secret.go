package elastigroup_azure_secrets

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Secret] = commons.NewGenericField(
		commons.ElastigroupAzureSecret,
		Secret,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(SourceVault): {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Name): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(ResourceGroupName): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},
					string(VaultCertificates): {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(CertificateStore): {
									Type:     schema.TypeString,
									Required: true,
								},
								string(CertificateURL): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup != nil && elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil && elastigroup.Compute.LaunchSpecification.Secrets != nil {
				secrets := elastigroup.Compute.LaunchSpecification.Secrets
				result = flattenSecrets(secrets)
			}

			if result != nil {
				if err := resourceData.Set(string(Secret), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Secret), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Secret)); ok {
				if value, err := expandSecrets(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetSecrets(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupAzureV3Wrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []*azurev3.Secrets = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil && elastigroup.Compute.LaunchSpecification.Secrets != nil {
				if v, ok := resourceData.GetOk(string(Secret)); ok {
					if secrets, err := expandSecrets(v); err != nil {
						return err
					} else {
						value = secrets
					}
					elastigroup.Compute.LaunchSpecification.SetSecrets(value)
				} else {
					elastigroup.Compute.LaunchSpecification.SetSecrets(nil)
				}
			}
			return nil
		},
		nil,
	)
}

func flattenSecrets(secret []*azurev3.Secrets) []interface{} {
	result := make([]interface{}, 0, len(secret))

	for _, sec := range secret {
		m := make(map[string]interface{})
		if sec.SourceVault != nil {
			m[string(SourceVault)] = flattenSourceVault(sec.SourceVault)
		}
		if sec.VaultCertificates != nil {
			m[string(VaultCertificates)] = flattenVaultCertificate(sec.VaultCertificates)
		}
		result = append(result, m)
	}

	return result
}

func flattenSourceVault(sourceVault *azurev3.SourceVault) []interface{} {
	result := make(map[string]interface{})

	result[string(Name)] = spotinst.StringValue(sourceVault.Name)
	result[string(ResourceGroupName)] = spotinst.StringValue(sourceVault.ResourceGroupName)

	return []interface{}{result}
}

func flattenVaultCertificate(vaultCert []*azurev3.VaultCertificates) []interface{} {
	result := make([]interface{}, 0, len(vaultCert))

	for _, VaultCertification := range vaultCert {
		m := make(map[string]interface{})
		m[string(CertificateURL)] = spotinst.StringValue(VaultCertification.CertificateUrl)
		m[string(CertificateStore)] = spotinst.StringValue(VaultCertification.CertificateStore)
		result = append(result, m)
	}

	return result
}

func expandSecrets(data interface{}) ([]*azurev3.Secrets, error) {
	list := data.(*schema.Set).List()
	sec := make([]*azurev3.Secrets, 0, len(list))

	for _, v := range list {
		se, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		secret := &azurev3.Secrets{}

		if v, ok := se[string(SourceVault)]; ok {
			sourceVault := &azurev3.SourceVault{}

			if secret.SourceVault != nil {
				sourceVault = secret.SourceVault
			}
			if sv, err := expandSourceVault(v, sourceVault); err != nil {
				return nil, err
			} else {
				if sv != nil {
					secret.SetSourceVault(sv)
				}
			}
		}

		if v, ok := se[string(VaultCertificates)]; ok {
			var vaultCer []*azurev3.VaultCertificates

			if secret.VaultCertificates != nil {
				vaultCer = secret.VaultCertificates
			}
			if vc, err := expandVaultCertificate(v, vaultCer); err != nil {
				return nil, err
			} else {
				secret.SetVaultCertificates(vc)
			}
		}
		sec = append(sec, secret)
	}
	return sec, nil
}

func expandSourceVault(data interface{}, sourceVault *azurev3.SourceVault) (*azurev3.SourceVault, error) {
	if list := data.([]interface{}); len(list) > 0 {
		if list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(Name)].(string); ok && v != "" {
				sourceVault.SetName(spotinst.String(v))
			}
			if v, ok := m[string(ResourceGroupName)].(string); ok && v != "" {
				sourceVault.SetResourceGroupName(spotinst.String(v))
			}
		}
		return sourceVault, nil
	}
	return nil, nil
}

func expandVaultCertificate(data interface{}, vaultCertificates []*azurev3.VaultCertificates) ([]*azurev3.VaultCertificates, error) {
	list := data.([]interface{})

	if len(list) == 0 && vaultCertificates == nil {
		return nil, nil
	}

	length := len(list) + len(vaultCertificates)
	newVaultCertificatesList := make([]*azurev3.VaultCertificates, 0, length)

	if len(vaultCertificates) > 0 {
		newVaultCertificatesList = append(newVaultCertificatesList, vaultCertificates[0])
	}

	for _, v := range list {
		adic, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		vaultCertificate := &azurev3.VaultCertificates{}

		if v, ok := adic[string(CertificateURL)].(string); ok && v != "" {
			vaultCertificate.SetCertificateUrl(spotinst.String(v))
		}
		if v, ok := adic[string(CertificateStore)].(string); ok && v != "" {
			vaultCertificate.SetCertificateStore(spotinst.String(v))
		}

		newVaultCertificatesList = append(newVaultCertificatesList, vaultCertificate)
	}

	return newVaultCertificatesList, nil
}

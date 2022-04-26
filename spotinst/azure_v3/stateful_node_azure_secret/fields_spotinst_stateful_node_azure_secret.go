package stateful_node_azure_secret

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Secret] = commons.NewGenericField(
		commons.StatefulNodeAzureSecret,
		Secret,
		&schema.Schema{
			Type:     schema.TypeList,
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
									Optional: true,
								},
								string(CertificateURL): {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := stWrapper.GetStatefulNode()
			var result []interface{} = nil
			if statefulNode != nil && statefulNode.Compute != nil && statefulNode.Compute.LaunchSpecification != nil && statefulNode.Compute.LaunchSpecification.Secrets != nil {
				secret := statefulNode.Compute.LaunchSpecification.Secrets
				result = flattenSecret(secret)
			}

			if result != nil {
				if err := resourceData.Set(string(Secret), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Secret), err)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			statefulNode := stWrapper.GetStatefulNode()
			if v, ok := resourceData.GetOk(string(Secret)); ok {
				if value, err := expandSecret(v); err != nil {
					return err
				} else {
					statefulNode.Compute.LaunchSpecification.SetSecrets(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			stWrapper := resourceObject.(*commons.StatefulNodeAzureV3Wrapper)
			st := stWrapper.GetStatefulNode()
			var value []*azurev3.Secret = nil
			if st.Compute != nil && st.Compute.LaunchSpecification != nil && st.Compute.LaunchSpecification.Secrets != nil {
				if v, ok := resourceData.GetOk(string(Secret)); ok {
					if secrets, err := expandSecret(v); err != nil {
						return err
					} else {
						value = secrets
					}
				}
				st.Compute.LaunchSpecification.SetSecrets(value)
			}
			return nil
		},
		nil,
	)
}

func flattenSecret(secret []*azurev3.Secret) []interface{} {
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

func flattenVaultCertificate(vaultCert []*azurev3.VaultCertificate) []interface{} {
	result := make([]interface{}, 0, len(vaultCert))

	for _, VaultCertification := range vaultCert {
		m := make(map[string]interface{})
		m[string(CertificateURL)] = spotinst.StringValue(VaultCertification.CertificateURL)
		m[string(CertificateStore)] = spotinst.StringValue(VaultCertification.CertificateStore)
		result = append(result, m)
	}

	return result
}

func expandSecret(data interface{}) ([]*azurev3.Secret, error) {
	list := data.(*schema.Set).List()
	sec := make([]*azurev3.Secret, 0, len(list))

	if len(list) > 0 {
		for _, v := range list {
			se, ok := v.(map[string]interface{})
			if !ok {
				continue
			}

			secret := &azurev3.Secret{}

			if v, ok := se[string(SourceVault)]; ok {
				// Create new securityGroup object in case cluster did not get it from previous import step.
				sourceVault := &azurev3.SourceVault{}

				if secret.SourceVault != nil {
					sourceVault = secret.SourceVault
				}

				if sourceVault, err := expandSourceVault(v, sourceVault); err != nil {
					return nil, err
				} else {
					if sourceVault != nil {
						secret.SetSourceVault(sourceVault)
					}
				}
			}

			if v, ok := se[string(VaultCertificates)]; ok {
				var vaultCer []*azurev3.VaultCertificate

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

func expandVaultCertificate(data interface{}, vaultCertificates []*azurev3.VaultCertificate) ([]*azurev3.VaultCertificate, error) {
	list := data.(*schema.Set).List()

	if len(list) == 0 && vaultCertificates == nil {
		return nil, nil
	}

	length := len(list) + len(vaultCertificates)
	newVaultCertificatesList := make([]*azurev3.VaultCertificate, 0, length)

	if len(vaultCertificates) > 0 {
		newVaultCertificatesList = append(newVaultCertificatesList, vaultCertificates[0])
	}

	for _, v := range list {
		adic, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		vaultCertificate := &azurev3.VaultCertificate{}

		if v, ok := adic[string(CertificateURL)].(string); ok && v != "" {
			vaultCertificate.SetCertificateURL(spotinst.String(v))
		}

		if v, ok := adic[string(CertificateStore)].(string); ok && v != "" {
			vaultCertificate.SetCertificateStore(spotinst.String(v))
		}

		newVaultCertificatesList = append(newVaultCertificatesList, vaultCertificate)
	}

	return newVaultCertificatesList, nil
}

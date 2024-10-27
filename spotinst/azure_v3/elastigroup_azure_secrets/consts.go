package elastigroup_azure_secrets

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Secret commons.FieldName = "secret"

	SourceVault       commons.FieldName = "source_vault"
	Name              commons.FieldName = "name"
	ResourceGroupName commons.FieldName = "resource_group_name"

	VaultCertificates commons.FieldName = "vault_certificates"
	CertificateURL    commons.FieldName = "certificate_url"
	CertificateStore  commons.FieldName = "certificate_store"
)

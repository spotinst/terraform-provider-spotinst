package credentials_gcp

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	AccountId               commons.FieldName = "account_id"
	Type                    commons.FieldName = "type"
	ProjectId               commons.FieldName = "project_id"
	PrivateKeyId            commons.FieldName = "private_key_id"
	PrivateKey              commons.FieldName = "private_key"
	ClientEmail             commons.FieldName = "client_email"
	ClientId                commons.FieldName = "client_id"
	AuthUri                 commons.FieldName = "auth_uri"
	TokenUri                commons.FieldName = "token_uri"
	AuthProviderX509CertUrl commons.FieldName = "auth_provider_x509_cert_url"
	ClientX509CertUrl       commons.FieldName = "client_x509_cert_url"
)

package multai_listener

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	BalancerID commons.FieldName = "balancer_id"
	Protocol   commons.FieldName = "protocol"
	Port       commons.FieldName = "port"
	TLSConfig  commons.FieldName = "tls_config"
	Tags       commons.FieldName = "tags"

	CertificateIDs           commons.FieldName = "certificate_ids"
	MinVersion               commons.FieldName = "min_version"
	MaxVersion               commons.FieldName = "max_version"
	SessionTicketsDisabled   commons.FieldName = "session_tickets_disabled"
	PreferServerCipherSuites commons.FieldName = "prefer_server_cipher_suites"
	CipherSuites             commons.FieldName = "cipher_suites"

	TagKey   commons.FieldName = "key"
	TagValue commons.FieldName = "value"
)

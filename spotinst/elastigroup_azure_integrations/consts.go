package elastigroup_azure_integrations

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	// - KUBERNETES ----------------------
	IntegrationKubernetes commons.FieldName = "integration_kubernetes"
	ClusterIdentifier     commons.FieldName = "cluster_identifier"
	// -----------------------------------

	// - MULTAI-RUNTIME ------------------
	IntegrationMultaiRuntime commons.FieldName = "integration_multai_runtime"
	DeploymentId             commons.FieldName = "deployment_id"
	// -----------------------------------
)

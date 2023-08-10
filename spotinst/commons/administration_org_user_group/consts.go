package administration_org_user_group

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Name        commons.FieldName = "name"
	Description commons.FieldName = "description"
	UserIds     commons.FieldName = "user_ids"
	Policies    commons.FieldName = "policies"
	AccountIds  commons.FieldName = "account_ids"
	PolicyId    commons.FieldName = "policy_id"
)

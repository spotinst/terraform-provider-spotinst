package oceancd_strategy_canary

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Canary    commons.FieldName = "canary"
	SetWeight commons.FieldName = "set_weight"
)

const (
	BackgroundVerification commons.FieldName = "background_verification"
)

const (
	Steps commons.FieldName = "steps"
	Name  commons.FieldName = "name"
)

const (
	Pause    commons.FieldName = "pause"
	Duration commons.FieldName = "duration"
)

const (
	SetCanaryScale     commons.FieldName = "set_canary_scale"
	MatchTrafficWeight commons.FieldName = "match_traffic_weight"
	Replicas           commons.FieldName = "replicas"
	Weight             commons.FieldName = "weight"
)

const (
	SetHeaderRoute commons.FieldName = "set_header_route"
	Match          commons.FieldName = "match"
	HeaderName     commons.FieldName = "header_name"
	HeaderValue    commons.FieldName = "header_value"
	Exact          commons.FieldName = "exact"
	Prefix         commons.FieldName = "prefix"
	Regex          commons.FieldName = "regex"
)

const (
	Verification  commons.FieldName = "regex"
	TemplateNames commons.FieldName = "template_names"
)

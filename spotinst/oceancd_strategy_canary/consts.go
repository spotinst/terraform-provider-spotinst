package oceancd_strategy_canary

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Name   commons.FieldName = "strategy_name"
	Canary commons.FieldName = "canary"
)

const (
	Steps     commons.FieldName = "steps"
	StepName  commons.FieldName = "step_name"
	SetWeight commons.FieldName = "set_weight"
)

const (
	BackgroundVerification commons.FieldName = "background_verification"
	BGTemplateNames        commons.FieldName = "template_names"
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
	SetHeaderRoute  commons.FieldName = "set_header_route"
	HeaderRouteName commons.FieldName = "header_route_name"
	Match           commons.FieldName = "match"
	HeaderName      commons.FieldName = "header_name"
	HeaderValue     commons.FieldName = "header_value"
	Exact           commons.FieldName = "exact"
	Prefix          commons.FieldName = "prefix"
	Regex           commons.FieldName = "regex"
)

const (
	Verification  commons.FieldName = "verification"
	TemplateNames commons.FieldName = "template_names"
)

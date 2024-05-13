package oceancd_strategy_rolling

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	Rolling commons.FieldName = "canary"
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
	Verification  commons.FieldName = "regex"
	TemplateNames commons.FieldName = "template_names"
)

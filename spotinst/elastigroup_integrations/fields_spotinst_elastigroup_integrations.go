package elastigroup_integrations

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	SetupEcs(fieldsMap)
	SetupNomad(fieldsMap)
	SetupRancher(fieldsMap)
	SetupKubernetes(fieldsMap)
	SetupMesosphere(fieldsMap)
	SetupCodeDeploy(fieldsMap)
	SetupMultaiRuntime(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupAutoScaleHeadroom(data interface{}) (*aws.AutoScaleHeadroom, error) {
	if list := data.([]interface{}); len(list) > 0 {
		headroom := &aws.AutoScaleHeadroom{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(CpuPerUnit)].(int); ok && v > 0 {
				headroom.SetCPUPerUnit(spotinst.Int(v))
			}

			if v, ok := m[string(MemoryPerUnit)].(int); ok && v > 0 {
				headroom.SetMemoryPerUnit(spotinst.Int(v))
			}

			if v, ok := m[string(NumOfUnits)].(int); ok && v > 0 {
				headroom.SetNumOfUnits(spotinst.Int(v))
			}
		}
		return headroom, nil
	}

	return nil, nil
}

func expandAWSGroupAutoScaleDown(data interface{}) (*aws.AutoScaleDown, error) {
	if list := data.([]interface{}); len(list) > 0 {
		autoScaleDown := &aws.AutoScaleDown{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(EvaluationPeriods)].(int); ok && v > 0 {
				autoScaleDown.SetEvaluationPeriods(spotinst.Int(v))
			}
		}
		return autoScaleDown, nil
	}

	return nil, nil
}

func expandAWSGroupAutoScaleConstraints(data interface{}) ([]*aws.AutoScaleConstraint, error) {
	list := data.(*schema.Set).List()
	out := make([]*aws.AutoScaleConstraint, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(Key)]; !ok {
			return nil, errors.New("invalid constraint attributes: key missing")
		}

		if _, ok := attr[string(Value)]; !ok {
			return nil, errors.New("invalid constraint attributes: value missing")
		}
		c := &aws.AutoScaleConstraint{
			Key:   spotinst.String(fmt.Sprintf("${%s}", attr[string(Key)].(string))),
			Value: spotinst.String(attr[string(Value)].(string)),
		}
		out = append(out, c)
	}
	return out, nil
}

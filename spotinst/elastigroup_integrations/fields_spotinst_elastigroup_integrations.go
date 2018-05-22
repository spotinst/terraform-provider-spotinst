package elastigroup_integrations

import (
	"fmt"
	"errors"

	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
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
	SetupElasticBeanstalk(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandAWSGroupAutoScaleHeadroom(data interface{}) (*aws.AutoScaleHeadroom, error) {
	if list := data.([]interface{}); len(list) > 0 {
		m := list[0].(map[string]interface{})
		i := &aws.AutoScaleHeadroom{}

		if v, ok := m[string(CpuPerUnit)].(int); ok && v > 0 {
			i.SetCPUPerUnit(spotinst.Int(v))
		}

		if v, ok := m[string(MemoryPerUnit)].(int); ok && v > 0 {
			i.SetMemoryPerUnit(spotinst.Int(v))
		}

		if v, ok := m[string(NumOfUnits)].(int); ok && v > 0 {
			i.SetNumOfUnits(spotinst.Int(v))
		}

		return i, nil
	}

	return nil, nil
}

func expandAWSGroupAutoScaleDown(data interface{}) (*aws.AutoScaleDown, error) {
	if list := data.([]interface{}); len(list) > 0 {
		m := list[0].(map[string]interface{})
		i := &aws.AutoScaleDown{}

		if v, ok := m[string(EvaluationPeriods)].(int); ok && v > 0 {
			i.SetEvaluationPeriods(spotinst.Int(v))
		}

		return i, nil
	}

	return nil, nil
}

func expandAWSGroupAutoScaleConstraints(data interface{}) ([]*aws.AutoScaleConstraint, error) {
	list := data.([]interface{})
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
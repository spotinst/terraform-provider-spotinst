package elastigroup_aws_integrations

import (
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
	SetupGitlab(fieldsMap)
	SetupRancher(fieldsMap)
	SetupKubernetes(fieldsMap)
	SetupMesosphere(fieldsMap)
	SetupCodeDeploy(fieldsMap)
	SetupMultaiRuntime(fieldsMap)
	SetupRoute53(fieldsMap)
	SetupDockerSwarm(fieldsMap)
	SetupElasticBeanstalk(fieldsMap)
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

			if v, ok := m[string(MaxScaleDownPercentage)].(int); ok && v > 0 {
				autoScaleDown.SetMaxScaleDownPercentage(spotinst.Int(v))
			}
		}
		return autoScaleDown, nil
	}

	return nil, nil
}

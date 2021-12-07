package ocean_aws_logging

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Logging] = commons.NewGenericField(
		commons.OceanAwsLogging,
		Logging,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Export): {
						Type:     schema.TypeSet,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(S3Id): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster != nil && cluster.Logging != nil {
				result = flattenLogging(cluster.Logging)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Logging), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Logging), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(Logging)); ok {
				if logging, err := expandOceanAWSLogging(v); err != nil {
					return err
				} else {
					cluster.SetLogging(logging)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AWSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *aws.Logging = nil

			if v, ok := resourceData.GetOk(string(Logging)); ok {
				if logging, err := expandOceanAWSLogging(v); err != nil {
					return err
				} else {
					value = logging
				}
			}
			cluster.SetLogging(value)
			return nil
		},
		nil,
	)
}

func flattenLogging(logging *aws.Logging) []interface{} {
	var out []interface{}

	if logging != nil {
		result := make(map[string]interface{})

		if logging.Export != nil {
			result[string(Export)] = flattenExport(logging.Export)
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func flattenExport(export *aws.Export) []interface{} {
	var out []interface{}

	if export != nil {
		result := make(map[string]interface{})

		if export.S3 != nil && export.S3.ID != nil {
			result[string(S3Id)] = export.S3.ID
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func expandOceanAWSLogging(data interface{}) (*aws.Logging, error) {
	logging := &aws.Logging{}
	list := data.(*schema.Set).List()

	if list == nil || list[0] == nil {
		return logging, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Export)]; ok {
		export, err := expandOceanAWSExport(v)
		if err != nil {
			return nil, err
		}
		if export != nil {
			logging.SetExport(export)
		} else {
			logging.Export = nil
		}
	}

	return logging, nil
}

func expandOceanAWSExport(data interface{}) (*aws.Export, error) {
	export := &aws.Export{}
	list := data.(*schema.Set).List()

	if list == nil || list[0] == nil {
		return export, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(S3Id)].(string); ok && v != "" {
		s3 := &aws.S3{}
		s3.SetId(spotinst.String(v))
		export.SetS3(s3)
	}

	return export, nil
}

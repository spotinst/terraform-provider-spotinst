package elastigroup_aws_logging

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Logging] = commons.NewGenericField(
		commons.ElastigroupAwsLogging,
		Logging,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Export): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(S3): {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											string(Id): {
												Type:     schema.TypeString,
												Required: true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []interface{} = nil
			if elastigroup != nil && elastigroup.Logging != nil {
				value = flattenLogging(elastigroup.Logging)
			}
			if len(value) > 0 {
				if err := resourceData.Set(string(Logging), value); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Logging), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Logging)); ok {
				if logging, err := expandElastigroupAWSLogging(v); err != nil {
					return err
				} else {
					elastigroup.SetLogging(logging)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *aws.Logging = nil

			if v, ok := resourceData.GetOk(string(Logging)); ok {
				if logging, err := expandElastigroupAWSLogging(v); err != nil {
					return err
				} else {
					value = logging
				}
			}
			elastigroup.SetLogging(value)
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

		if export.S3 != nil {
			result[string(S3)] = flattenS3(export.S3)
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func flattenS3(s3 *aws.S3) []interface{} {
	var out []interface{}

	if s3 != nil {
		result := make(map[string]interface{})

		if s3.ID != nil {
			result[string(Id)] = s3.ID
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func expandElastigroupAWSLogging(data interface{}) (*aws.Logging, error) {
	logging := &aws.Logging{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return logging, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Export)]; ok {
		export, err := expandElastigroupAWSExport(v)
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

func expandElastigroupAWSExport(data interface{}) (*aws.Export, error) {
	export := &aws.Export{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return export, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(S3)]; ok {
		s3, err := expandElastigroupAWSS3(v)
		if err != nil {
			return nil, err
		}
		if s3 != nil {
			export.SetS3(s3)
		} else {
			export.S3 = nil
		}
	}

	return export, nil
}

func expandElastigroupAWSS3(data interface{}) (*aws.S3, error) {
	s3 := &aws.S3{}
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return s3, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(Id)].(string); ok && v != "" {
		s3.SetId(spotinst.String(v))
	}

	return s3, nil
}

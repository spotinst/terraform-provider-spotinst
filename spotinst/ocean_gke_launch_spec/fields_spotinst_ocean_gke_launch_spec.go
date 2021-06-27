package ocean_gke_launch_spec

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	/*fieldsMap[OceanId] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		OceanId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var value *string = nil
			if ls != nil && ls.OceanID != nil {
				value = ls.OceanID
			}
			if err := resourceData.Set(string(OceanId), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OceanId), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(OceanId)).(string); ok && v != "" {
				ls.SetOceanId(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(OceanId)).(string); ok && v != "" {
				ls.SetOceanId(spotinst.String(v))
			}
			return nil
		},
		nil,
	)*/

	fieldsMap[OceanId] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		OceanId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
			//Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(OceanId))
			return err
		},
		nil,
	)

	fieldsMap[NodePoolName] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		NodePoolName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
			//Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(NodePoolName))
			return err
		},
		nil,
	)

	fieldsMap[SourceImage] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		SourceImage,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var value *string = nil

			if ls != nil && ls.SourceImage != nil {
				value = ls.SourceImage
			}

			if err := resourceData.Set(string(SourceImage), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SourceImage), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(SourceImage)).(string); ok && v != "" {
				ls.SetSourceImage(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			if v, ok := resourceData.Get(string(SourceImage)).(string); ok && v != "" {
				ls.SetSourceImage(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Metadata] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		Metadata,
		&schema.Schema{
			Type:     schema.TypeSet,
			Required: true,
			//Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MetadataKey): {
						Type:     schema.TypeString,
						Required: true,
						//Computed: true,
					},

					string(MetadataValue): {
						Type:     schema.TypeString,
						Required: true,
						//Computed: true,
					},
				},
			},
			Set: hashKV,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if ls != nil && ls.Metadata != nil {
				metadata := ls.Metadata
				result = flattenMetadata(metadata)
			}
			if result != nil {
				if err := resourceData.Set(string(Metadata), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Metadata), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var value []*gcp.Metadata = nil

			if v, ok := resourceData.GetOk(string(Metadata)); ok {
				var metadata []*gcp.Metadata

				if ls != nil {
					if ls.Metadata != nil {
						metadata = ls.Metadata
					}

					if metadata, err := expandMetadata(v, metadata); err != nil {
						return err
					} else {
						value = metadata
					}

					ls.SetMetadata(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var value []*gcp.Metadata = nil

			if v, ok := resourceData.GetOk(string(Metadata)); ok {
				var metadata []*gcp.Metadata

				if ls != nil {
					if ls.Metadata != nil {
						metadata = ls.Metadata
					}

					if metadata, err := expandMetadata(v, metadata); err != nil {
						return err
					} else {
						value = metadata
					}

					ls.SetMetadata(value)
				}
			}

			return nil
		},
		nil,
	)

	fieldsMap[Labels] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		Labels,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(LabelKey): {
						Type:     schema.TypeString,
						Required: true,
						//Computed: true,
					},

					string(LabelValue): {
						Type:     schema.TypeString,
						Required: true,
						//Computed: true,
					},
				},
			},
			Set: hashKV,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if ls != nil && ls.Labels != nil {
				labels := ls.Labels
				result = flattenLabels(labels)
			}
			if result != nil {
				if err := resourceData.Set(string(Labels), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Labels), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var value []*gcp.Label = nil

			if v, ok := resourceData.GetOk(string(Labels)); ok {
				var labels []*gcp.Label

				if ls != nil {
					if ls.Labels != nil {
						labels = ls.Labels
					}

					if labels, err := expandLabels(v, labels); err != nil {
						return err
					} else {
						value = labels
					}

					ls.SetLabels(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var value []*gcp.Label = nil

			if v, ok := resourceData.GetOk(string(Labels)); ok {
				var labels []*gcp.Label

				if ls != nil {
					if ls.Labels != nil {
						labels = ls.Labels
					}

					if labels, err := expandLabels(v, labels); err != nil {
						return err
					} else {
						value = labels
					}

					ls.SetLabels(value)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[Taints] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		Taints,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(TaintKey): {
						Type:     schema.TypeString,
						Required: true,
						//Computed: true,
					},

					string(TaintValue): {
						Type:     schema.TypeString,
						Required: true,
						//Computed: true,
					},

					string(TaintEffect): {
						Type:     schema.TypeString,
						Required: true,
						//Computed: true,
					},
				},
			},
			Set: hashKVTaints,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if ls != nil && ls.Taints != nil {
				taints := ls.Taints
				result = flattenTaints(taints)
			}
			if result != nil {
				if err := resourceData.Set(string(Taints), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Taints), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var value []*gcp.Taint = nil

			if v, ok := resourceData.GetOk(string(Taints)); ok {
				var taints []*gcp.Taint

				if ls != nil {
					if ls.Taints != nil {
						taints = ls.Taints
					}

					if taints, err := expandTaints(v, taints); err != nil {
						return err
					} else {
						value = taints
					}

					ls.SetTaints(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var value []*gcp.Taint = nil

			if v, ok := resourceData.GetOk(string(Taints)); ok {
				var taints []*gcp.Taint

				if ls != nil {
					if ls.Taints != nil {
						taints = ls.Taints
					}

					if taints, err := expandTaints(v, taints); err != nil {
						return err
					} else {
						value = taints
					}

					ls.SetTaints(value)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[AutoscaleHeadrooms] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		AutoscaleHeadrooms,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(CPUPerUnit): {
						Type:     schema.TypeInt,
						Optional: true,
						Computed: true,
					},

					string(GPUPerUnit): {
						Type:     schema.TypeInt,
						Optional: true,
						Computed: true,
					},

					string(MemoryPerUnit): {
						Type:     schema.TypeInt,
						Optional: true,
						Computed: true,
					},

					string(NumOfUnits): {
						Type:     schema.TypeInt,
						Required: true,
						//Computed: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec.AutoScale != nil && launchSpec.AutoScale.Headrooms != nil {
				headrooms := launchSpec.AutoScale.Headrooms
				result = flattenHeadrooms(headrooms)
			}
			if result != nil {
				if err := resourceData.Set(string(AutoscaleHeadrooms), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AutoscaleHeadrooms), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var value []*gcp.AutoScaleHeadroom = nil

			if v, ok := resourceData.GetOk(string(AutoscaleHeadrooms)); ok {
				var autoScaleHeadrooms []*gcp.AutoScaleHeadroom

				if ls != nil && ls.AutoScale != nil {
					if ls.AutoScale.Headrooms != nil {
						autoScaleHeadrooms = ls.AutoScale.Headrooms
					}

					if autoScaleHeadrooms, err := expandHeadrooms(v, autoScaleHeadrooms); err != nil {
						return err
					} else {
						value = autoScaleHeadrooms
					}

					ls.AutoScale.SetHeadrooms(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var value []*gcp.AutoScaleHeadroom = nil

			if v, ok := resourceData.GetOk(string(AutoscaleHeadrooms)); ok {
				var autoScaleHeadrooms []*gcp.AutoScaleHeadroom

				if ls != nil && ls.AutoScale != nil {
					if ls.AutoScale.Headrooms != nil {
						autoScaleHeadrooms = ls.AutoScale.Headrooms
					}

					if autoScaleHeadrooms, err := expandHeadrooms(v, autoScaleHeadrooms); err != nil {
						return err
					} else {
						value = autoScaleHeadrooms
					}

					ls.AutoScale.SetHeadrooms(value)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[RestrictScaleDown] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		RestrictScaleDown,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *bool = nil
			if launchSpec.RestrictScaleDown != nil {
				value = launchSpec.RestrictScaleDown
			}
			if value != nil {
				if err := resourceData.Set(string(RestrictScaleDown), spotinst.BoolValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RestrictScaleDown), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOkExists(string(RestrictScaleDown)); ok && v != nil {
				restrictScaleDown := spotinst.Bool(v.(bool))
				launchSpec.SetRestrictScaleDown(restrictScaleDown)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var restrictScaleDown *bool = nil
			if v, ok := resourceData.GetOkExists(string(RestrictScaleDown)); ok && v != nil {
				restrictScaleDown = spotinst.Bool(v.(bool))
			}
			launchSpec.SetRestrictScaleDown(restrictScaleDown)
			return nil
		},
		nil,
	)
}

func hashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(LabelKey)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(LabelValue)].(string)))
	return hashcode.String(buf.String())
}

func hashKVTaints(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(TaintKey)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(TaintValue)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(TaintEffect)].(string)))
	return hashcode.String(buf.String())
}

func expandTaints(data interface{}, taints []*gcp.Taint) ([]*gcp.Taint, error) {
	list := data.(*schema.Set).List()
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		taint := &gcp.Taint{}

		if v, ok := attr[string(TaintKey)].(string); ok && v != "" {
			taint.SetKey(spotinst.String(v))
		}

		if v, ok := attr[string(TaintValue)].(string); ok && v != "" {
			taint.SetValue(spotinst.String(v))
		}

		if v, ok := attr[string(TaintEffect)].(string); ok && v != "" {
			taint.SetValue(spotinst.String(v))
		}

		taints = append(taints, taint)
	}
	return taints, nil
}

func expandLabels(data interface{}, labels []*gcp.Label) ([]*gcp.Label, error) {
	list := data.(*schema.Set).List()
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		label := &gcp.Label{}

		if v, ok := attr[string(LabelKey)].(string); ok && v != "" {
			label.SetKey(spotinst.String(v))
		}

		if v, ok := attr[string(LabelValue)].(string); ok && v != "" {
			label.SetValue(spotinst.String(v))
		}

		labels = append(labels, label)
	}
	return labels, nil
}

func expandMetadata(data interface{}, metadata []*gcp.Metadata) ([]*gcp.Metadata, error) {
	list := data.(*schema.Set).List()
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		md := &gcp.Metadata{}

		if v, ok := attr[string(MetadataKey)].(string); ok && v != "" {
			md.SetKey(spotinst.String(v))
		}

		if v, ok := attr[string(MetadataValue)].(string); ok && v != "" {
			md.SetValue(spotinst.String(v))
		}

		metadata = append(metadata, md)
	}
	return metadata, nil
}

func flattenTaints(taints []*gcp.Taint) []interface{} {
	result := make([]interface{}, 0, len(taints))
	for _, taint := range taints {
		m := make(map[string]interface{})
		m[string(TaintKey)] = spotinst.StringValue(taint.Key)
		m[string(TaintValue)] = spotinst.StringValue(taint.Value)
		m[string(TaintEffect)] = spotinst.StringValue(taint.Effect)

		result = append(result, m)
	}
	return result
}

func flattenLabels(labels []*gcp.Label) []interface{} {
	result := make([]interface{}, 0, len(labels))
	for _, label := range labels {
		m := make(map[string]interface{})
		m[string(LabelKey)] = spotinst.StringValue(label.Key)
		m[string(LabelValue)] = spotinst.StringValue(label.Value)

		result = append(result, m)
	}
	return result
}

func flattenMetadata(metadata []*gcp.Metadata) []interface{} {
	result := make([]interface{}, 0, len(metadata))
	for _, metaObject := range metadata {
		m := make(map[string]interface{})
		m[string(MetadataKey)] = spotinst.StringValue(metaObject.Key)
		m[string(MetadataValue)] = spotinst.StringValue(metaObject.Value)

		result = append(result, m)
	}
	return result
}

func expandHeadrooms(data interface{}, autoScaleHeadrooms []*gcp.AutoScaleHeadroom) ([]*gcp.AutoScaleHeadroom, error) {
	list := data.(*schema.Set).List()
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		autoScaleHeadroom := &gcp.AutoScaleHeadroom{}

		if v, ok := attr[string(CPUPerUnit)].(int); ok && v >= 0 {
			autoScaleHeadroom.SetCPUPerUnit(spotinst.Int(v))
		}

		if v, ok := attr[string(GPUPerUnit)].(int); ok && v >= 0 {
			autoScaleHeadroom.SetGPUPerUnit(spotinst.Int(v))
		}

		if v, ok := attr[string(NumOfUnits)].(int); ok && v >= 0 {
			autoScaleHeadroom.SetNumOfUnits(spotinst.Int(v))
		}

		if v, ok := attr[string(MemoryPerUnit)].(int); ok && v >= 0 {
			autoScaleHeadroom.SetMemoryPerUnit(spotinst.Int(v))
		}

		autoScaleHeadrooms = append(autoScaleHeadrooms, autoScaleHeadroom)
	}
	return autoScaleHeadrooms, nil
}

func flattenHeadrooms(headrooms []*gcp.AutoScaleHeadroom) []interface{} {
	result := make([]interface{}, 0, len(headrooms))

	for _, headroom := range headrooms {
		m := make(map[string]interface{})
		m[string(CPUPerUnit)] = spotinst.IntValue(headroom.CPUPerUnit)
		m[string(GPUPerUnit)] = spotinst.IntValue(headroom.GPUPerUnit)
		m[string(NumOfUnits)] = spotinst.IntValue(headroom.NumOfUnits)
		m[string(MemoryPerUnit)] = spotinst.IntValue(headroom.MemoryPerUnit)

		result = append(result, m)
	}

	return result
}

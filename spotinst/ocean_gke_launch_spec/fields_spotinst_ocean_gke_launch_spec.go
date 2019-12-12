package ocean_gke_launch_spec

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[OceanId] = commons.NewGenericField(
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
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MetadataKey): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(MetadataValue): {
						Type:     schema.TypeString,
						Required: true,
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
			if value, ok := resourceData.GetOk(string(Metadata)); ok {
				if metadata, err := expandMetadata(value); err != nil {
					return err
				} else {
					ls.SetMetadata(metadata)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var metadataList []*gcp.Metadata
			if value, ok := resourceData.GetOk(string(Metadata)); ok {
				if metadata, err := expandMetadata(value); err != nil {
					return err
				} else {
					metadataList = metadata
				}
			}
			ls.SetMetadata(metadataList)
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
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(LabelKey): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(LabelValue): {
						Type:     schema.TypeString,
						Required: true,
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
			if value, ok := resourceData.GetOk(string(Labels)); ok {
				if labels, err := expandLabels(value); err != nil {
					return err
				} else {
					ls.SetLabels(labels)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var labelList []*gcp.Label = nil
			if value, ok := resourceData.GetOk(string(Labels)); ok {
				if labels, err := expandLabels(value); err != nil {
					return err
				} else {
					labelList = labels
				}
			}
			ls.SetLabels(labelList)
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
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(TaintKey): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(TaintValue): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(TaintEffect): {
						Type:     schema.TypeString,
						Required: true,
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
			if value, ok := resourceData.GetOk(string(Taints)); ok {
				if taints, err := expandTaints(value); err != nil {
					return err
				} else {
					ls.SetTaints(taints)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var taintsList []*gcp.Taint = nil
			if value, ok := resourceData.GetOk(string(Taints)); ok {
				if taints, err := expandTaints(value); err != nil {
					return err
				} else {
					taintsList = taints
				}
			}
			ls.SetTaints(taintsList)
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
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(CPUPerUnit): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(GPUPerUnit): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(MemoryPerUnit): {
						Type:     schema.TypeInt,
						Optional: true,
					},

					string(NumOfUnits): {
						Type:     schema.TypeInt,
						Required: true,
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
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(AutoscaleHeadrooms)); ok {
				if headrooms, err := expandHeadrooms(value); err != nil {
					return err
				} else {
					launchSpec.AutoScale.SetHeadrooms(headrooms)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			LaunchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := LaunchSpecWrapper.GetLaunchSpec()
			var headroomList []*gcp.AutoScaleHeadroom = nil
			if value, ok := resourceData.GetOk(string(AutoscaleHeadrooms)); ok {
				if expandedList, err := expandHeadrooms(value); err != nil {
					return err
				} else {
					headroomList = expandedList
				}
			}
			launchSpec.AutoScale.SetHeadrooms(headroomList)
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

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

func expandTaints(data interface{}) ([]*gcp.Taint, error) {
	list := data.(*schema.Set).List()
	taints := make([]*gcp.Taint, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(TaintKey)]; !ok {
			return nil, errors.New("invalid taint attributes: key missing")
		}

		if _, ok := attr[string(TaintValue)]; !ok {
			return nil, errors.New("invalid taint attributes: value missing")
		}

		if _, ok := attr[string(TaintEffect)]; !ok {
			return nil, errors.New("invalid taint attributes: effect missing")
		}

		taint := &gcp.Taint{
			Key:    spotinst.String(attr[string(TaintKey)].(string)),
			Value:  spotinst.String(attr[string(TaintValue)].(string)),
			Effect: spotinst.String(attr[string(TaintEffect)].(string)),
		}
		taints = append(taints, taint)
	}
	return taints, nil
}

func expandLabels(data interface{}) ([]*gcp.Label, error) {
	list := data.(*schema.Set).List()
	labels := make([]*gcp.Label, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(LabelKey)]; !ok {
			return nil, errors.New("invalid label attributes: key missing")
		}

		if _, ok := attr[string(LabelValue)]; !ok {
			return nil, errors.New("invalid label attributes: value missing")
		}
		label := &gcp.Label{
			Key:   spotinst.String(attr[string(LabelKey)].(string)),
			Value: spotinst.String(attr[string(LabelValue)].(string)),
		}
		labels = append(labels, label)
	}
	return labels, nil
}

func expandMetadata(data interface{}) ([]*gcp.Metadata, error) {
	list := data.(*schema.Set).List()
	metadata := make([]*gcp.Metadata, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(MetadataKey)]; !ok {
			return nil, errors.New("invalid metadata attributes: key missing")
		}

		if _, ok := attr[string(MetadataValue)]; !ok {
			return nil, errors.New("invalid metadata attributes: value missing")
		}
		metaObject := &gcp.Metadata{
			Key:   spotinst.String(attr[string(MetadataKey)].(string)),
			Value: spotinst.String(attr[string(MetadataValue)].(string)),
		}
		metadata = append(metadata, metaObject)
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

func expandHeadrooms(data interface{}) ([]*gcp.AutoScaleHeadroom, error) {
	list := data.(*schema.Set).List()
	headrooms := make([]*gcp.AutoScaleHeadroom, 0, len(list))

	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		headroom := &gcp.AutoScaleHeadroom{
			CPUPerUnit:    spotinst.Int(attr[string(CPUPerUnit)].(int)),
			GPUPerUnit:    spotinst.Int(attr[string(GPUPerUnit)].(int)),
			NumOfUnits:    spotinst.Int(attr[string(NumOfUnits)].(int)),
			MemoryPerUnit: spotinst.Int(attr[string(MemoryPerUnit)].(int)),
		}

		headrooms = append(headrooms, headroom)
	}
	return headrooms, nil
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

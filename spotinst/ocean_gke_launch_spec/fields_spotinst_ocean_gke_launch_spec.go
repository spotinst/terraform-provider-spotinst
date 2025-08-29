package ocean_gke_launch_spec

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

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

	fieldsMap[NodePoolName] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		NodePoolName,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
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
			Optional: true,
			Computed: true,
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
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MetadataKey): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},

					string(MetadataValue): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
				},
			},
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

				if ls != nil {
					if ls.Metadata != nil {
						value = ls.Metadata
					}

					if metadata, err := expandMetadata(v); err != nil {
						return err
					} else if metadata != nil {
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

				if ls != nil {
					if ls.Metadata != nil {
						value = ls.Metadata
					}

					if metadata, err := expandMetadata(v); err != nil {
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
						Optional: true,
						Computed: true,
					},

					string(LabelValue): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
				},
			},
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
						Optional: true,
						Computed: true,
					},

					string(TaintValue): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},

					string(TaintEffect): {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
				},
			},
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

	fieldsMap[AutoscaleHeadroomsAutomatic] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		AutoscaleHeadroomsAutomatic,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(AutoHeadroomPercentage): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
						DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
							if old == "-1" && new == "null" {
								return true
							}
							return false
						},
					},
				},
			},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if launchSpec != nil && launchSpec.AutoScale != nil && launchSpec.AutoScale.AutoHeadroomPercentage != nil {
				result = flattenAutoscaleHeadroomsAutomatic(launchSpec.AutoScale)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(AutoscaleHeadroomsAutomatic), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(AutoscaleHeadroomsAutomatic), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(AutoscaleHeadroomsAutomatic)); ok {
				if err := expandAutoscaleHeadroomsAutomatic(v, launchSpec, false); err != nil {
					return err
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(AutoscaleHeadroomsAutomatic)); ok {
				if err := expandAutoscaleHeadroomsAutomatic(v, launchSpec, true); err != nil {
					return err
				}
			} else {
				launchSpec.AutoScale.SetAutoHeadroomPercentage(nil)
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

	fieldsMap[RootVolumeSizeInGB] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		RootVolumeSizeInGB,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *int = nil

			if launchSpec.RootVolumeSizeInGB != nil {
				value = launchSpec.RootVolumeSizeInGB
			}
			if value != nil {
				if err := resourceData.Set(string(RootVolumeSizeInGB), spotinst.IntValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RootVolumeSizeInGB), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()

			if v, ok := resourceData.GetOkExists(string(RootVolumeSizeInGB)); ok && v != nil {
				rootVolumeSize := spotinst.Int(v.(int))
				launchSpec.SetRootVolumeSizeInGB(rootVolumeSize)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var rootVolumeSize *int = nil

			if v, ok := resourceData.GetOkExists(string(RootVolumeSizeInGB)); ok && v != nil {
				rootVolumeSize = spotinst.Int(v.(int))
			}
			launchSpec.SetRootVolumeSizeInGB(rootVolumeSize)
			return nil
		},
		nil,
	)

	fieldsMap[RootVolumeType] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		RootVolumeType,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *string = nil

			if launchSpec.RootVolumeType != nil {
				value = launchSpec.RootVolumeType
			}
			if value != nil {
				if err := resourceData.Set(string(RootVolumeType), spotinst.StringValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RootVolumeType), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()

			if v, ok := resourceData.GetOkExists(string(RootVolumeType)); ok && v != nil {
				rootVolumeType := spotinst.String(v.(string))
				launchSpec.SetRootVolumeType(rootVolumeType)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var rootVolumeType *string = nil

			if v, ok := resourceData.GetOkExists(string(RootVolumeType)); ok && v != nil {
				rootVolumeType = spotinst.String(v.(string))
			}
			launchSpec.SetRootVolumeType(rootVolumeType)
			return nil
		},
		nil,
	)
	fieldsMap[InstanceTypes] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		InstanceTypes,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value []string = nil
			if launchSpec != nil && launchSpec.InstanceTypes != nil {
				value = launchSpec.InstanceTypes
			}
			if err := resourceData.Set(string(InstanceTypes), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(InstanceTypes), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()

			if v, ok := resourceData.GetOk(string(InstanceTypes)); ok {

				instances := v.([]interface{})
				if len(instances) > 0 {
					instanceTypes := make([]string, len(instances))
					for i, j := range instances {
						instanceTypes[i] = j.(string)
					}
					launchSpec.SetInstanceTypes(instanceTypes)
					return nil
				}

				if launchSpec != nil {
					if launchSpec.InstanceTypes != nil {
						launchSpec.SetInstanceTypes(launchSpec.InstanceTypes)
					}
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()

			if v, ok := resourceData.GetOk(string(InstanceTypes)); ok {

				instances := v.([]interface{})
				if len(instances) > 0 {
					instanceTypes := make([]string, len(instances))
					for i, j := range instances {
						instanceTypes[i] = j.(string)
					}
					launchSpec.SetInstanceTypes(instanceTypes)
					return nil
				}

				if launchSpec != nil {
					if launchSpec.InstanceTypes != nil {
						launchSpec.SetInstanceTypes(launchSpec.InstanceTypes)
					}
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[ShieldedInstanceConfig] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		ShieldedInstanceConfig,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(EnableSecureBoot): {
						Type:     schema.TypeBool,
						Optional: true,
						Computed: true,
					},
					string(EnableIntegrityMonitoring): {
						Type:     schema.TypeBool,
						Optional: true,
						Computed: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if ls != nil && ls.ShieldedInstanceConfig != nil {
				shieldedInstanceConfig := ls.ShieldedInstanceConfig
				result = flattenShieldedInstanceConfig(shieldedInstanceConfig)
			}
			if result != nil {
				if err := resourceData.Set(string(ShieldedInstanceConfig), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShieldedInstanceConfig), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var value *gcp.ShieldedInstanceConfig = nil

			if v, ok := resourceData.GetOk(string(ShieldedInstanceConfig)); ok {

				if ls != nil {
					if ls.ShieldedInstanceConfig != nil {
						value = ls.ShieldedInstanceConfig
					}

					if shieldedInstanceConfig, err := expandShieldedInstanceConfig(v); err != nil {
						return err
					} else if shieldedInstanceConfig != nil {
						value = shieldedInstanceConfig
					}

					ls.SetShieldedInstanceConfig(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var value *gcp.ShieldedInstanceConfig = nil
			if v, ok := resourceData.GetOk(string(ShieldedInstanceConfig)); ok {
				if shieldedInstanceConfig, err := expandShieldedInstanceConfig(v); err != nil {
					return err
				} else {
					value = shieldedInstanceConfig
				}
			}
			ls.SetShieldedInstanceConfig(value)
			return nil
		},
		nil,
	)

	fieldsMap[Storage] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		Storage,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(LocalSSDCount): {
						Type:     schema.TypeInt,
						Optional: true,
						Computed: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if ls != nil && ls.Storage != nil {
				storage := ls.Storage
				result = flattenStorage(storage)
			}
			if result != nil {
				if err := resourceData.Set(string(Storage), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Storage), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var value *gcp.Storage = nil

			if v, ok := resourceData.GetOk(string(Storage)); ok {

				if ls != nil {
					if ls.Storage != nil {
						value = ls.Storage
					}

					if storage, err := expandStorage(v); err != nil {
						return err
					} else if storage != nil {
						value = storage
					}

					ls.SetStorage(value)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var value *gcp.Storage = nil
			if v, ok := resourceData.GetOk(string(Storage)); ok {
				if storage, err := expandStorage(v); err != nil {
					return err
				} else {
					value = storage
				}
			}
			ls.SetStorage(value)
			return nil
		},
		nil,
	)

	fieldsMap[ServiceAccount] = commons.NewGenericField(
		commons.OceanAWSLaunchSpec,
		ServiceAccount,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *string = nil

			if launchSpec.ServiceAccount != nil {
				value = launchSpec.ServiceAccount
			}
			if value != nil {
				if err := resourceData.Set(string(ServiceAccount), spotinst.StringValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ServiceAccount), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()

			if v, ok := resourceData.GetOkExists(string(ServiceAccount)); ok && v != nil {
				serviceAccount := spotinst.String(v.(string))
				launchSpec.SetServiceAccount(serviceAccount)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(ServiceAccount))
			return err
		},
		nil,
	)

	fieldsMap[Tags] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		Tags,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()

			var value []string = nil
			if launchSpec != nil && launchSpec.LaunchSpecTags != nil {
				value = launchSpec.LaunchSpecTags
			}
			if err := resourceData.Set(string(Tags), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tags), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []string
			if v, ok := resourceData.GetOk(string(Tags)); ok && v != nil {
				tagsList := v.([]interface{})
				result = make([]string, len(tagsList))
				if len(tagsList) > 0 {
					for i, j := range tagsList {
						result[i] = j.(string)
					}
				}
			}
			if launchSpec != nil {
				if launchSpec.LaunchSpecTags != nil {
					result = append(result, launchSpec.LaunchSpecTags...)
				}
			}
			launchSpec.SetLaunchSpecTags(result)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []string
			if v, ok := resourceData.GetOk(string(Tags)); ok && v != nil {
				tagsList := v.([]interface{})
				result = make([]string, len(tagsList))
				if len(tagsList) > 0 {
					for i, j := range tagsList {
						result[i] = j.(string)
					}
				}
			}
			launchSpec.SetLaunchSpecTags(result)
			return nil
		},
		nil,
	)

	fieldsMap[ResourceLimits] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		ResourceLimits,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MaxInstanceCount): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
					string(MinInstanceCount): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var result []interface{} = nil
			if ls != nil && ls.ResourceLimits != nil {
				resourceLimits := ls.ResourceLimits
				result = flattenResourceLimits(resourceLimits)
			}
			if result != nil {
				if err := resourceData.Set(string(ResourceLimits), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ResourceLimits), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			if value, ok := resourceData.GetOk(string(ResourceLimits)); ok {
				if resourceLimits, err := expandResourceLimits(value); err != nil {
					return err
				} else {
					ls.SetResourceLimits(resourceLimits)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			lsWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			ls := lsWrapper.GetLaunchSpec()
			var resourceLimits *gcp.ResourceLimits = nil
			if value, ok := resourceData.GetOk(string(ResourceLimits)); ok {
				if interfaces, err := expandResourceLimits(value); err != nil {
					return err
				} else {
					resourceLimits = interfaces
				}
			}
			ls.SetResourceLimits(resourceLimits)
			return nil
		},
		nil,
	)

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *string = nil

			if launchSpec.Name != nil {
				value = launchSpec.Name
			}
			if value != nil {
				if err := resourceData.Set(string(Name), spotinst.StringValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()

			if v, ok := resourceData.GetOkExists(string(Name)); ok && v != nil {
				name := spotinst.String(v.(string))
				launchSpec.SetName(name)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var name *string = nil
			if v, ok := resourceData.GetOkExists(string(Name)); ok && v != nil {
				name = spotinst.String(v.(string))
			}
			launchSpec.SetName(name)
			return nil
		},
		nil,
	)

	fieldsMap[UpdatePolicy] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		UpdatePolicy,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ShouldRoll): {
						Type:     schema.TypeBool,
						Required: true,
					},

					string(RollConfig): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(BatchSizePercentage): {
									Type:     schema.TypeInt,
									Required: true,
								},
							},
						},
					},
				},
			},
		},
		nil, nil, nil, nil,
	)

	fieldsMap[NetworkInterfaces] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		NetworkInterfaces,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Network): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(ProjectId): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(AccessConfigs): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(AccessConfigsName): {
									Type:     schema.TypeString,
									Optional: true,
								},

								string(Type): {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},

					string(AliasIPRanges): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(IPCidrRange): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(SubnetworkRangeName): {
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
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(NetworkInterfaces)); ok {
				if networks, err := expandLaunchSpecNetworkInterfaces(v); err != nil {
					return err
				} else {
					launchSpec.SetLaunchSpecNetworkInterfaces(networks)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value []*gcp.LaunchSpecNetworkInterfaces = nil
			if v, ok := resourceData.GetOk(string(NetworkInterfaces)); ok {
				if networks, err := expandLaunchSpecNetworkInterfaces(v); err != nil {
					return err
				} else {
					value = networks
				}
			}
			launchSpec.SetLaunchSpecNetworkInterfaces(value)
			return nil
		},
		nil,
	)

	fieldsMap[CreateOptions] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		CreateOptions,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(InitialNodes): {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		nil, nil, nil, nil,
	)

	fieldsMap[Filters] = commons.NewGenericField(
		commons.OceanGKELaunchSpec,
		Filters,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ExcludeFamilies): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(IncludeFamilies): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(MaxMemoryGiB): {
						Type:     schema.TypeFloat,
						Optional: true,
						Default:  -1,
					},

					string(MaxVcpu): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},

					string(MinMemoryGiB): {
						Type:     schema.TypeFloat,
						Optional: true,
						Default:  -1,
					},

					string(MinVcpu): {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var result []interface{} = nil

			if launchSpec != nil && launchSpec.Filters != nil {
				result = flattenFilters(launchSpec.Filters)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Filters), result); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(Filters), err)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			if v, ok := resourceData.GetOk(string(Filters)); ok {
				if filters, err := expandFilters(v, false); err != nil {
					return err
				} else {
					launchSpec.SetFilters(filters)
				}
			}
			return nil
		},

		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			launchSpecWrapper := resourceObject.(*commons.LaunchSpecGKEWrapper)
			launchSpec := launchSpecWrapper.GetLaunchSpec()
			var value *gcp.Filters = nil

			if v, ok := resourceData.GetOk(string(Filters)); ok {
				if filters, err := expandFilters(v, true); err != nil {
					return err
				} else {
					value = filters
				}
			}
			launchSpec.SetFilters(value)
			return nil
		},
		nil,
	)
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
			taint.SetEffect(spotinst.String(v))
		}

		v := checkTaintKeyExists(taint, taints)
		if v == false {
			taints = append(taints, taint)
		}
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

		v := checkLabelKeyExists(label, labels)
		if v == false {
			labels = append(labels, label)
		}
	}
	return labels, nil
}

func expandMetadata(data interface{}) ([]*gcp.Metadata, error) {
	var metadata []*gcp.Metadata
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

		v := checkMetadataUniqueness(md, metadata)
		if v == true {
			metadata = append(metadata, md)

		}

	}
	if len(metadata) > 0 {
		return metadata, nil
	} else {
		return nil, nil
	}
}

func expandShieldedInstanceConfig(data interface{}) (*gcp.ShieldedInstanceConfig, error) {
	var shieldedInstanceConfig *gcp.ShieldedInstanceConfig
	updated := 0
	list := data.(*schema.Set).List()
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		sic := &gcp.ShieldedInstanceConfig{}

		if v, ok := attr[string(EnableSecureBoot)].(bool); ok {
			updated = 1
			sic.SetEnableSecureBoot(spotinst.Bool(v))
		}

		if v, ok := attr[string(EnableIntegrityMonitoring)].(bool); ok {
			updated = 1
			sic.SetEnableIntegrityMonitoring(spotinst.Bool(v))
		}

		shieldedInstanceConfig = sic
	}
	if updated != 0 {
		return shieldedInstanceConfig, nil
	} else {
		return nil, nil
	}
}

func expandStorage(data interface{}) (*gcp.Storage, error) {
	var storage *gcp.Storage
	updated := 0
	list := data.(*schema.Set).List()
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		s := &gcp.Storage{}

		if v, ok := attr[string(LocalSSDCount)].(int); ok {
			updated = 1
			s.SetLocalSSDCount(spotinst.Int(v))
		}

		storage = s
	}
	if updated != 0 {
		return storage, nil
	} else {
		return nil, nil
	}
}

func expandResourceLimits(data interface{}) (*gcp.ResourceLimits, error) {
	var resourceLimits *gcp.ResourceLimits
	updated := 0
	list := data.([]interface{})
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		r := &gcp.ResourceLimits{}

		if v, ok := attr[string(MaxInstanceCount)].(int); ok && v >= 0 {
			updated = 1
			r.SetMaxInstanceCount(spotinst.Int(v))
		} else {
			r.SetMaxInstanceCount(nil)
		}

		if v, ok := attr[string(MinInstanceCount)].(int); ok && v >= 0 {
			updated = 1
			r.SetMinInstanceCount(spotinst.Int(v))
		} else {
			r.SetMinInstanceCount(nil)
		}

		resourceLimits = r
	}
	if updated != 0 {
		return resourceLimits, nil
	} else {
		return nil, nil
	}
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

func flattenShieldedInstanceConfig(shieldedInstanceConfig *gcp.ShieldedInstanceConfig) []interface{} {
	var out []interface{}

	if shieldedInstanceConfig != nil {
		result := make(map[string]interface{})

		if shieldedInstanceConfig.EnableSecureBoot != nil {
			result[string(EnableSecureBoot)] = spotinst.BoolValue(shieldedInstanceConfig.EnableSecureBoot)
		}
		if shieldedInstanceConfig.EnableIntegrityMonitoring != nil {
			result[string(EnableIntegrityMonitoring)] = spotinst.BoolValue(shieldedInstanceConfig.EnableIntegrityMonitoring)
		}
		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

func flattenStorage(storage *gcp.Storage) []interface{} {
	var out []interface{}

	if storage != nil {
		result := make(map[string]interface{})

		if storage.LocalSSDCount != nil {
			result[string(LocalSSDCount)] = spotinst.IntValue(storage.LocalSSDCount)
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

func flattenResourceLimits(resourceLimits *gcp.ResourceLimits) []interface{} {
	var out []interface{}

	if resourceLimits != nil {
		result := make(map[string]interface{})

		value := spotinst.Int(-1)
		result[string(MinInstanceCount)] = value
		result[string(MaxInstanceCount)] = value

		if resourceLimits.MaxInstanceCount != nil {
			result[string(MaxInstanceCount)] = spotinst.IntValue(resourceLimits.MaxInstanceCount)
		}

		if resourceLimits.MinInstanceCount != nil {
			result[string(MinInstanceCount)] = spotinst.IntValue(resourceLimits.MinInstanceCount)
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}
	return out
}

func expandAutoscaleHeadroomsAutomatic(data interface{}, ls *gcp.LaunchSpec, nullify bool) error {
	list := data.([]interface{})

	if list == nil || list[0] == nil {
		return nil
	}

	m := list[0].(map[string]interface{})

	if v, ok := m[string(AutoHeadroomPercentage)].(int); ok && v > -1 {
		ls.AutoScale.SetAutoHeadroomPercentage(spotinst.Int(v))
	} else if nullify {
		ls.AutoScale.SetAutoHeadroomPercentage(nil)
	}

	return nil
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

func flattenAutoscaleHeadroomsAutomatic(autoScale *gcp.AutoScale) []interface{} {
	var out []interface{}

	if autoScale != nil {

		result := make(map[string]interface{})

		value := spotinst.Int(-1)
		if autoScale.AutoHeadroomPercentage != nil {
			value = autoScale.AutoHeadroomPercentage
		}
		result[string(AutoHeadroomPercentage)] = spotinst.IntValue(value)

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
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

func checkMetadataUniqueness(md *gcp.Metadata, metadata []*gcp.Metadata) bool {
	for _, mdElement := range metadata {
		if spotinst.StringValue(mdElement.Key) == spotinst.StringValue(md.Key) {
			if spotinst.StringValue(mdElement.Value) == spotinst.StringValue(md.Value) {
				return false
			}
		}
	}
	return true
}

func checkLabelKeyExists(label *gcp.Label, labels []*gcp.Label) bool {
	for index, labelElement := range labels {
		if spotinst.StringValue(labelElement.Key) == spotinst.StringValue(label.Key) {
			labels[index].SetValue(spotinst.String(*label.Value))
			return true
		}
	}
	return false
}

func checkTaintKeyExists(taint *gcp.Taint, taints []*gcp.Taint) bool {
	for index, taintElement := range taints {
		if spotinst.StringValue(taintElement.Key) == spotinst.StringValue(taint.Key) {
			taints[index].SetValue(spotinst.String(*taint.Value))
			taints[index].SetEffect(spotinst.String(*taint.Effect))
			return true
		}
	}
	return false
}

// expandVNGNetworkInterface sets the values from the plan as objects
func expandLaunchSpecNetworkInterfaces(data interface{}) ([]*gcp.LaunchSpecNetworkInterfaces, error) {
	list := data.(*schema.Set).List()

	if list != nil && list[0] != nil {
		ifaces := make([]*gcp.LaunchSpecNetworkInterfaces, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			iface := &gcp.LaunchSpecNetworkInterfaces{}

			if v, ok := m[string(Network)].(string); ok && v != "" {
				iface.SetNetwork(spotinst.String(v))
			}

			if v, ok := m[string(ProjectId)].(string); ok && v != "" {
				iface.SetProjectId(spotinst.String(v))
			}

			if v, ok := m[string(AccessConfigs)]; ok {
				accessConfigs, err := expandVNGAccessConfigs(v)
				if err != nil {
					return nil, err
				}

				if accessConfigs != nil {
					iface.SetLaunchSpecAccessConfigs(accessConfigs)
				}
			} else {
				iface.LaunchSpecAccessConfigs = nil
			}

			if v, ok := m[string(AliasIPRanges)]; ok {
				aliasRange, err := expandVNGAliasIPRanges(v)
				if err != nil {
					return nil, err
				}
				if aliasRange != nil {
					iface.SetLaunchSpecAliasIPRanges(aliasRange)
				}
			} else {
				iface.LaunchSpecAliasIPRanges = nil
			}

			ifaces = append(ifaces, iface)
		}
		return ifaces, nil
	}
	return nil, nil
}

// expandAccessConfigs sets the values from the plan as objects
func expandVNGAccessConfigs(data interface{}) ([]*gcp.LaunchSpecAccessConfigs, error) {
	list := data.(*schema.Set).List()
	accessConfigs := make([]*gcp.LaunchSpecAccessConfigs, 0, len(list))

	for _, item := range list {
		attr := item.(map[string]interface{})

		accessConfig := &gcp.LaunchSpecAccessConfigs{}

		if v, ok := attr[string(AccessConfigsName)].(string); ok && v != "" {
			accessConfig.SetLaunchSpecAccessConfigsName(spotinst.String(v))
		}

		if v, ok := attr[string(Type)].(string); ok && v != "" {
			accessConfig.SetType(spotinst.String(v))
		}

		accessConfigs = append(accessConfigs, accessConfig)
	}
	return accessConfigs, nil
}

// expandAccessConfigs sets the values from the plan as objects
func expandVNGAliasIPRanges(data interface{}) ([]*gcp.LaunchSpecAliasIPRanges, error) {
	list := data.(*schema.Set).List()
	aliasRanges := make([]*gcp.LaunchSpecAliasIPRanges, 0, len(list))

	for _, item := range list {
		m := item.(map[string]interface{})
		aliasRange := &gcp.LaunchSpecAliasIPRanges{}

		if v, ok := m[string(SubnetworkRangeName)].(string); ok && v != "" {
			aliasRange.SetSubnetworkRangeName(spotinst.String(v))
		}

		if v, ok := m[string(IPCidrRange)].(string); ok && v != "" {
			aliasRange.SetIPCidrRange(spotinst.String(v))
		}

		aliasRanges = append(aliasRanges, aliasRange)
	}
	return aliasRanges, nil
}

func flattenLaunchSpecNetworkIntefaces(networkInterfaces []*gcp.LaunchSpecNetworkInterfaces) []interface{} {
	result := make([]interface{}, 0, len(networkInterfaces))

	for _, networkInterface := range networkInterfaces {
		m := make(map[string]interface{})
		m[string(Network)] = spotinst.StringValue(networkInterface.Network)
		m[string(ProjectId)] = spotinst.StringValue(networkInterface.ProjectID)

		if networkInterface.LaunchSpecAccessConfigs != nil {
			m[string(AccessConfigs)] = flattenLaunchSpecAccessConfigs(networkInterface.LaunchSpecAccessConfigs)
		}

		if networkInterface.LaunchSpecAliasIPRanges != nil {
			m[string(AliasIPRanges)] = flattenLaunchSpecAliasIPRanges(networkInterface.LaunchSpecAliasIPRanges)
		}

		result = append(result, m)
	}

	return result
}

func flattenLaunchSpecAccessConfigs(accessConfigs []*gcp.LaunchSpecAccessConfigs) []interface{} {
	result := make([]interface{}, 0, len(accessConfigs))

	for _, accessConfig := range accessConfigs {
		m := make(map[string]interface{})
		m[string(AccessConfigsName)] = spotinst.StringValue(accessConfig.LaunchSpecAccessConfigsName)
		m[string(Type)] = spotinst.StringValue(accessConfig.Type)

		result = append(result, m)
	}

	return result
}

func flattenLaunchSpecAliasIPRanges(aliasIPRanges []*gcp.LaunchSpecAliasIPRanges) []interface{} {
	result := make([]interface{}, 0, len(aliasIPRanges))

	for _, aliasIPRange := range aliasIPRanges {
		m := make(map[string]interface{})
		m[string(IPCidrRange)] = spotinst.StringValue(aliasIPRange.IPCidrRange)
		m[string(SubnetworkRangeName)] = spotinst.StringValue(aliasIPRange.SubnetworkRangeName)

		result = append(result, m)
	}

	return result
}

func flattenFilters(filters *gcp.Filters) []interface{} {
	var out []interface{}

	if filters != nil {
		result := make(map[string]interface{})
		value := spotinst.Int(-1)
		result[string(MaxMemoryGiB)] = value
		result[string(MinMemoryGiB)] = value
		result[string(MaxVcpu)] = value
		result[string(MinVcpu)] = value

		if filters.MaxMemoryGiB != nil {
			result[string(MaxMemoryGiB)] = spotinst.Float64Value(filters.MaxMemoryGiB)
		}
		if filters.MinMemoryGiB != nil {
			result[string(MinMemoryGiB)] = spotinst.Float64Value(filters.MinMemoryGiB)
		}
		if filters.MaxVcpu != nil {
			result[string(MaxVcpu)] = spotinst.IntValue(filters.MaxVcpu)
		}
		if filters.MinVcpu != nil {
			result[string(MinVcpu)] = spotinst.IntValue(filters.MinVcpu)
		}
		if filters.ExcludeFamilies != nil {
			result[string(ExcludeFamilies)] = filters.ExcludeFamilies
		}
		if filters.IncludeFamilies != nil {
			result[string(IncludeFamilies)] = filters.IncludeFamilies
		}

		if len(result) > 0 {
			out = append(out, result)
		}
	}

	return out
}

func expandFilters(data interface{}, nullify bool) (*gcp.Filters, error) {
	filters := &gcp.Filters{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return filters, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(ExcludeFamilies)]; ok {
		excludeFamilies, err := expandFiltersList(v)
		if err != nil {
			return nil, err
		}
		if excludeFamilies != nil && len(excludeFamilies) > 0 {
			filters.SetExcludeFamilies(excludeFamilies)
		} else {
			if nullify {
				filters.SetExcludeFamilies(nil)
			}
		}
	}

	if v, ok := m[string(IncludeFamilies)]; ok {
		includeFamilies, err := expandFiltersList(v)
		if err != nil {
			return nil, err
		}
		if includeFamilies != nil && len(includeFamilies) > 0 {
			filters.SetIncludeFamilies(includeFamilies)
		} else {
			if nullify {
				filters.SetIncludeFamilies(nil)
			}
		}
	}

	if v, ok := m[string(MaxMemoryGiB)].(float64); ok {
		if v == -1 {
			filters.SetMaxMemoryGiB(nil)
		} else {
			filters.SetMaxMemoryGiB(spotinst.Float64(v))
		}
	}

	if v, ok := m[string(MaxVcpu)].(int); ok {
		if v == -1 {
			filters.SetMaxVcpu(nil)
		} else {
			filters.SetMaxVcpu(spotinst.Int(v))
		}
	}

	if v, ok := m[string(MinMemoryGiB)].(float64); ok {
		if v == -1 {
			filters.SetMinMemoryGiB(nil)
		} else {
			filters.SetMinMemoryGiB(spotinst.Float64(v))
		}
	}

	if v, ok := m[string(MinVcpu)].(int); ok {
		if v == -1 {
			filters.SetMinVcpu(nil)
		} else {
			filters.SetMinVcpu(spotinst.Int(v))
		}
	}

	return filters, nil
}

func expandFiltersList(data interface{}) ([]string, error) {
	list := data.(*schema.Set).List()
	result := make([]string, 0, len(list))

	for _, v := range list {
		if filterList, ok := v.(string); ok && filterList != "" {
			result = append(result, filterList)
		}
	}
	return result, nil
}

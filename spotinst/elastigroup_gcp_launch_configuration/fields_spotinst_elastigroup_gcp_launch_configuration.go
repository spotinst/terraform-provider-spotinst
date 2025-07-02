package elastigroup_gcp_launch_configuration

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[BackendServices] = commons.NewGenericField(
		commons.ElastigroupGCPLaunchConfiguration,
		BackendServices,
		&schema.Schema{
			Type:             schema.TypeSet,
			Optional:         true,
			DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(ServiceName): {
						Type:             schema.TypeString,
						Required:         true,
						DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
					},

					string(LocationType): {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
					},

					string(Scheme): {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
					},

					string(NamedPorts): {
						Type:             schema.TypeSet,
						Optional:         true,
						DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Name): {
									Type:             schema.TypeString,
									Required:         true,
									DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
								},

								string(Ports): {
									Type:             schema.TypeList,
									Required:         true,
									DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
									Elem:             &schema.Schema{Type: schema.TypeString},
								},
							},
						},
					},
					string(BackendBalancing): {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(BackendBalancingMode): {
									Type:     schema.TypeString,
									Optional: true,
								},

								string(MaxRatePerInstance): {
									Type:     schema.TypeInt,
									Optional: true,
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
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(BackendServices)); ok {
				if services, err := expandServices(v); err != nil {
					return err
				} else {
					backendSvsCfg := &gcp.BackendServiceConfig{}
					backendSvsCfg.SetBackendServices(services)
					elastigroup.Compute.LaunchSpecification.SetBackendServiceConfig(backendSvsCfg)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result *gcp.BackendServiceConfig = nil
			if v, ok := resourceData.GetOk(string(BackendServices)); ok {
				var value []*gcp.BackendService
				if services, err := expandServices(v); err != nil {
					return err
				} else {
					result = &gcp.BackendServiceConfig{}
					value = services
				}
				result.SetBackendServices(value)
			}
			elastigroup.Compute.LaunchSpecification.SetBackendServiceConfig(result)
			return nil
		},
		nil,
	)

	fieldsMap[Labels] = commons.NewGenericField(
		commons.ElastigroupGCPLaunchConfiguration,
		Labels,
		&schema.Schema{
			Type:             schema.TypeSet,
			Optional:         true,
			DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(LabelKey): {
						Type:             schema.TypeString,
						Required:         true,
						DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
					},

					string(LabelValue): {
						Type:             schema.TypeString,
						Required:         true,
						DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Labels != nil {
				labels := elastigroup.Compute.LaunchSpecification.Labels
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
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if value, ok := resourceData.GetOk(string(Labels)); ok {
				if labels, err := expandLabels(value); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetLabels(labels)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var labelList []*gcp.Label = nil
			if value, ok := resourceData.GetOk(string(Labels)); ok {
				if labels, err := expandLabels(value); err != nil {
					return err
				} else {
					labelList = labels
				}
			}
			elastigroup.Compute.LaunchSpecification.SetLabels(labelList)
			return nil
		},
		nil,
	)

	fieldsMap[Metadata] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		Metadata,
		&schema.Schema{
			Type:             schema.TypeSet,
			Optional:         true,
			DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(MetadataKey): {
						Type:             schema.TypeString,
						Required:         true,
						DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
					},

					string(MetadataValue): {
						Type:             schema.TypeString,
						Required:         true,
						DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []interface{} = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Metadata != nil {
				metadata := elastigroup.Compute.LaunchSpecification.Metadata
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
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if value, ok := resourceData.GetOk(string(Metadata)); ok {
				if metadata, err := expandMetadata(value); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetMetadata(metadata)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var metadataList []*gcp.Metadata
			if value, ok := resourceData.GetOk(string(Metadata)); ok {
				if metadata, err := expandMetadata(value); err != nil {
					return err
				} else {
					metadataList = metadata
				}
			}
			elastigroup.Compute.LaunchSpecification.SetMetadata(metadataList)
			return nil
		},
		nil,
	)

	fieldsMap[Tags] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		Tags,
		&schema.Schema{
			Type:             schema.TypeList,
			Optional:         true,
			Elem:             &schema.Schema{Type: schema.TypeString},
			DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value []string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.Tags != nil {
				value = elastigroup.Compute.LaunchSpecification.Tags
			}
			if err := resourceData.Set(string(Tags), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tags), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []string
			if v, ok := resourceData.GetOk(string(Tags)); ok && v != nil {
				tagsList := v.([]interface{})
				result = make([]string, len(tagsList))
				for i, j := range tagsList {
					result[i] = j.(string)
				}
			}
			elastigroup.Compute.LaunchSpecification.SetTags(result)
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result []string
			if v, ok := resourceData.GetOk(string(Tags)); ok && v != nil {
				tagsList := v.([]interface{})
				result = make([]string, len(tagsList))
				for i, j := range tagsList {
					result[i] = j.(string)
				}
			}
			elastigroup.Compute.LaunchSpecification.SetTags(result)
			return nil
		},
		nil,
	)

	fieldsMap[StartupScript] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		StartupScript,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.StartupScript != nil {
				value = elastigroup.Compute.LaunchSpecification.StartupScript
			}
			if err := resourceData.Set(string(StartupScript), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(StartupScript), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(StartupScript)); ok && v != nil {
				elastigroup.Compute.LaunchSpecification.SetStartupScript(spotinst.String(resourceData.Get(string(StartupScript)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(StartupScript)); ok && v != nil {
				elastigroup.Compute.LaunchSpecification.SetStartupScript(spotinst.String(resourceData.Get(string(StartupScript)).(string)))
			} else {
				elastigroup.Compute.LaunchSpecification.SetStartupScript(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[ShutdownScript] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		ShutdownScript,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			// occasionally shutdown_script will be set to the hash value of a null string
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				if (old == "da39a3ee5e6b4b0d3255bfef95601890afd80709" && new == "") ||
					(old == "" && new == "da39a3ee5e6b4b0d3255bfef95601890afd80709") {
					return true
				}
				return commons.SuppressIfImportedFromGKE(k, old, new, d)
			},
			StateFunc: Base64StateFunc,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value = ""
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.ShutdownScript != nil {

				shutdownScript := elastigroup.Compute.LaunchSpecification.ShutdownScript
				shutdownScriptValue := spotinst.StringValue(shutdownScript)
				if shutdownScriptValue != "" {
					if isBase64Encoded(resourceData.Get(string(ShutdownScript)).(string)) {
						value = shutdownScriptValue
					} else {
						decodedShutdownScript, _ := base64.StdEncoding.DecodeString(shutdownScriptValue)
						value = string(decodedShutdownScript)
					}
				}
			}
			if err := resourceData.Set(string(ShutdownScript), Base64StateFunc(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShutdownScript), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(ShutdownScript)).(string); ok && v != "" {
				shutdownScript := spotinst.String(base64Encode(v))
				elastigroup.Compute.LaunchSpecification.SetShutdownScript(shutdownScript)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var shutdownScript *string = nil
			if v, ok := resourceData.Get(string(ShutdownScript)).(string); ok && v != "" {
				shutdownScript = spotinst.String(base64Encode(v))
			}
			elastigroup.Compute.LaunchSpecification.SetShutdownScript(shutdownScript)
			return nil
		},
		nil,
	)

	fieldsMap[InstanceNamePrefix] = commons.NewGenericField(
		commons.ElastigroupGCPLaunchConfiguration,
		InstanceNamePrefix,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.InstanceNamePrefix != nil {
				value = elastigroup.Compute.LaunchSpecification.InstanceNamePrefix
			}
			if err := resourceData.Set(string(InstanceNamePrefix), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(InstanceNamePrefix), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(InstanceNamePrefix)); ok && v != nil {
				elastigroup.Compute.LaunchSpecification.SetInstanceNamePrefix(spotinst.String(resourceData.Get(string(InstanceNamePrefix)).(string)))
			} else {
				elastigroup.Compute.LaunchSpecification.SetInstanceNamePrefix(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(InstanceNamePrefix)); ok && v != nil {
				elastigroup.Compute.LaunchSpecification.SetInstanceNamePrefix(spotinst.String(resourceData.Get(string(InstanceNamePrefix)).(string)))
			} else {
				elastigroup.Compute.LaunchSpecification.SetInstanceNamePrefix(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[ServiceAccount] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		ServiceAccount,
		&schema.Schema{
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.ServiceAccount != nil {
				value = elastigroup.Compute.LaunchSpecification.ServiceAccount
			}
			if err := resourceData.Set(string(ServiceAccount), spotinst.StringValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ServiceAccount), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(ServiceAccount)).(string); ok && v != "" {
				elastigroup.Compute.LaunchSpecification.SetServiceAccount(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			serviceAccount := ""
			if v, ok := resourceData.Get(string(ServiceAccount)).(string); ok && v != "" {
				serviceAccount = v
			}
			elastigroup.Compute.LaunchSpecification.SetServiceAccount(spotinst.String(serviceAccount))
			return nil
		},
		nil,
	)

	fieldsMap[IPForwarding] = commons.NewGenericField(
		commons.ElastigroupAWSLaunchConfiguration,
		IPForwarding,
		&schema.Schema{
			Type:             schema.TypeBool,
			Optional:         true,
			Default:          false,
			DiffSuppressFunc: commons.SuppressIfImportedFromGKE,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *bool = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.IPForwarding != nil {
				value = elastigroup.Compute.LaunchSpecification.IPForwarding
			}
			if err := resourceData.Set(string(IPForwarding), spotinst.BoolValue(value)); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(IPForwarding), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(IPForwarding)).(bool); ok {
				elastigroup.Compute.LaunchSpecification.SetIPForwarding(spotinst.Bool(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.Get(string(IPForwarding)).(bool); ok {
				elastigroup.Compute.LaunchSpecification.SetIPForwarding(spotinst.Bool(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[MinCpuPlatform] = commons.NewGenericField(
		commons.ElastigroupGCPLaunchConfiguration,
		MinCpuPlatform,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var value *string = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.MinCpuPlatform != nil {
				value = elastigroup.Compute.LaunchSpecification.MinCpuPlatform
			}
			if err := resourceData.Set(string(MinCpuPlatform), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(MinCpuPlatform), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(MinCpuPlatform)); ok && v != nil {
				elastigroup.Compute.LaunchSpecification.SetMinCpuPlatform(spotinst.String(resourceData.Get(string(MinCpuPlatform)).(string)))
			} else {
				elastigroup.Compute.LaunchSpecification.SetMinCpuPlatform(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(MinCpuPlatform)); ok && v != nil {
				elastigroup.Compute.LaunchSpecification.SetMinCpuPlatform(spotinst.String(resourceData.Get(string(MinCpuPlatform)).(string)))
			} else {
				elastigroup.Compute.LaunchSpecification.SetMinCpuPlatform(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[ShieldedInstanceConfig] = commons.NewGenericField(
		commons.ElastigroupGCPLaunchConfiguration,
		ShieldedInstanceConfig,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(EnableSecureBoot): {
						Type:     schema.TypeBool,
						Optional: true,
					},
					string(EnableIntegrityMonitoring): {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result interface{} = nil
			if elastigroup.Compute != nil && elastigroup.Compute.LaunchSpecification != nil &&
				elastigroup.Compute.LaunchSpecification.ShieldedInstanceConfig != nil {
				shieldedConfig := elastigroup.Compute.LaunchSpecification.ShieldedInstanceConfig
				result = flattenShieldedInstanceConfig(shieldedConfig)
			}
			if result != nil {
				if err := resourceData.Set(string(ShieldedInstanceConfig), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ShieldedInstanceConfig), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(ShieldedInstanceConfig)); ok {
				if shieldedInstanceConfig, err := expandShieldedInstanceConfig(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetShieldedInstanceConfig(shieldedInstanceConfig)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result *gcp.ShieldedInstanceConfig = nil
			if v, ok := resourceData.GetOk(string(ShieldedInstanceConfig)); ok {
				if value, err := expandShieldedInstanceConfig(v); err != nil {
					return err
				} else {
					result = value
				}
			}
			elastigroup.Compute.LaunchSpecification.SetShieldedInstanceConfig(result)
			return nil
		},
		nil,
	)

}

func Base64StateFunc(v interface{}) string {
	if isBase64Encoded(v.(string)) {
		return v.(string)
	} else {
		return base64Encode(v.(string))
	}
}

// base64Encode encodes data if the input isn't already encoded using
// base64.StdEncoding.EncodeToString. If the input is already base64 encoded,
// return the original input unchanged.
func base64Encode(data string) string {
	// Check whether the data is already Base64 encoded; don't double-encode
	if isBase64Encoded(data) {
		return data
	}
	// data has not been encoded encode and return
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func isBase64Encoded(data string) bool {
	_, err := base64.StdEncoding.DecodeString(data)
	return err == nil
}

// flattenLabels flattens the labels struct
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

// flattenMetadata flattens the metadata struct
func flattenMetadata(metadata []*gcp.Metadata) []interface{} {
	result := make([]interface{}, 0, len(metadata))
	for _, metaObject := range metadata {
		m := make(map[string]interface{})
		m[string(LabelKey)] = spotinst.StringValue(metaObject.Key)
		m[string(LabelValue)] = spotinst.StringValue(metaObject.Value)

		result = append(result, m)
	}
	return result
}

// expandLabels sets the values from the plan
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

// expandMetadata sets the values from the plan
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

// expandServices expands the Backend Services object.
func expandServices(data interface{}) ([]*gcp.BackendService, error) {
	list := data.(*schema.Set).List()
	out := make([]*gcp.BackendService, 0, len(list))

	for _, v := range list {
		elem := &gcp.BackendService{}
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		if v, ok := attr[string(ServiceName)]; ok {
			elem.SetBackendServiceName(spotinst.String(v.(string)))
		}

		if v, ok := attr[string(Scheme)].(string); ok && v != "" {
			elem.SetScheme(spotinst.String(v))
		}

		if v, ok := attr[string(LocationType)].(string); ok && v != "" {
			elem.SetLocationType(spotinst.String(v))
		}

		if v, ok := attr[string(NamedPorts)]; ok {
			namedPorts, err := expandNamedPorts(v)
			if err != nil {
				return nil, err
			}
			if namedPorts != nil {
				elem.SetNamedPorts(namedPorts)
			}
		}

		if v, ok := attr[string(BackendBalancing)]; ok {
			backendBalancing, err := expandBackendBalancing(v)
			if err != nil {
				return nil, err
			}
			if backendBalancing != nil {
				elem.SetBackendBalancing(backendBalancing)
			}
		}

		out = append(out, elem)
	}
	return out, nil
}

// expandNamedPorts expands the named port object
func expandNamedPorts(data interface{}) (*gcp.NamedPorts, error) {
	list := data.(*schema.Set).List()
	namedPorts := &gcp.NamedPorts{}

	for _, item := range list {
		m := item.(map[string]interface{})
		if v, ok := m[string(Name)].(string); ok && v != "" {
			namedPorts.SetName(spotinst.String(v))
		}

		if v, ok := m[string(Ports)]; ok && v != nil {
			portsList := v.([]interface{})
			result := make([]int, len(portsList))
			for i, j := range portsList {
				if intVal, err := strconv.Atoi(j.(string)); err != nil {
					return nil, err
				} else {
					result[i] = intVal
				}
			}
			namedPorts.SetPorts(result)
		}
	}
	return namedPorts, nil
}

func expandBackendBalancing(data interface{}) (*gcp.BackendBalancing, error) {
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return nil, nil
	}
	m := list[0].(map[string]interface{})
	backendBalancing := &gcp.BackendBalancing{}

	if v, ok := m[string(BackendBalancingMode)].(string); ok && v != "" {
		backendBalancing.SetBackendBalancingMode(spotinst.String(v))
	}

	if v, ok := m[string(MaxRatePerInstance)].(int); ok && v > 0 {
		backendBalancing.SetMaxRatePerInstance(spotinst.Int(v))
	}

	return backendBalancing, nil
}

func expandShieldedInstanceConfig(data interface{}) (*gcp.ShieldedInstanceConfig, error) {
	shieldedInstanceConfigs := &gcp.ShieldedInstanceConfig{}
	list := data.([]interface{})
	if list == nil || list[0] == nil {
		return shieldedInstanceConfigs, nil
	}
	m := list[0].(map[string]interface{})

	if v, ok := m[string(EnableSecureBoot)].(bool); ok {
		shieldedInstanceConfigs.SetEnableSecureBoot(spotinst.Bool(v))
	}

	if v, ok := m[string(EnableIntegrityMonitoring)].(bool); ok {
		shieldedInstanceConfigs.SetEnableIntegrityMonitoring(spotinst.Bool(v))
	}
	return shieldedInstanceConfigs, nil

}

func flattenShieldedInstanceConfig(shieldedInstanceConfig *gcp.ShieldedInstanceConfig) []interface{} {
	result := make(map[string]interface{})
	result[string(EnableSecureBoot)] = spotinst.BoolValue(shieldedInstanceConfig.EnableSecureBoot)
	result[string(EnableIntegrityMonitoring)] = spotinst.BoolValue(shieldedInstanceConfig.EnableIntegrityMonitoring)

	return []interface{}{result}
}

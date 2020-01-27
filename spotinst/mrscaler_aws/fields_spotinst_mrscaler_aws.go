package mrscaler_aws

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/mrscaler"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/mrscaler_aws_strategy"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[Name] = commons.NewGenericField(
		commons.MRScalerAWS,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Name != nil {
				value = scaler.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			scaler.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			scaler.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Description] = commons.NewGenericField(
		commons.MRScalerAWS,
		Description,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Description != nil {
				value = scaler.Description
			}
			if err := resourceData.Set(string(Description), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Description), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			scaler.SetDescription(spotinst.String(resourceData.Get(string(Description)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			scaler.SetDescription(spotinst.String(resourceData.Get(string(Description)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[Region] = commons.NewGenericField(
		commons.MRScalerAWS,
		Region,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Region != nil {
				value = scaler.Region
			}
			if err := resourceData.Set(string(Region), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Region), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(Region)); ok {
				scaler.SetRegion(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(Region))
			return err
		},
		nil,
	)

	fieldsMap[ClusterID] = commons.NewGenericField(
		commons.MRScalerAWS,
		ClusterID,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Strategy != nil {
				if scaler.Strategy.Wrapping != nil && scaler.Strategy.Wrapping.SourceClusterID != nil {
					value = scaler.Strategy.Wrapping.SourceClusterID
				}
				if scaler.Strategy.Cloning != nil && scaler.Strategy.Cloning.OriginClusterID != nil {
					value = scaler.Strategy.Cloning.OriginClusterID
				}
			}
			if err := resourceData.Set(string(ClusterID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ClusterID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()

			if v, ok := resourceData.GetOk(string(ClusterID)); ok && v != nil {
				if strategy, ok := resourceData.GetOk(string(Strategy)); ok && strategy != "" {
					switch strategy {
					case mrscaler_aws_strategy.Clone:
						if scaler.Strategy.Cloning == nil {
							scaler.Strategy.SetCloning(&mrscaler.Cloning{})
						}
						scaler.Strategy.Cloning.SetOriginClusterId(spotinst.String(v.(string)))
					case mrscaler_aws_strategy.Wrap:
						if scaler.Strategy.Wrapping == nil {
							scaler.Strategy.SetWrapping(&mrscaler.Wrapping{})
						}
						scaler.Strategy.Wrapping.SetSourceClusterId(spotinst.String(v.(string)))
					}
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(ClusterID))
			return err
		},
		nil,
	)

	fieldsMap[ExposeClusterID] = commons.NewGenericField(
		commons.MRScalerAWS,
		ExposeClusterID,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[OutputClusterID] = commons.NewGenericField(
		commons.MRScalerAWS,
		OutputClusterID,
		&schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[AvailabilityZones] = commons.NewGenericField(
		commons.MRScalerAWS,
		AvailabilityZones,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			// Skip
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if value, ok := resourceData.GetOk(string(AvailabilityZones)); ok {
				if zones, err := expandMRScalerAWSAvailabilityZones(value); err != nil {
					return err
				} else {
					scaler.Compute.SetAvailabilityZones(zones)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(AvailabilityZones))
			return err
		},
		nil,
	)

	fieldsMap[Tags] = commons.NewGenericField(
		commons.MRScalerAWS,
		Tags,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(TagKey): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(TagValue): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			Set: hashKV,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var result []interface{} = nil
			if scaler.Compute != nil &&
				scaler.Compute.Tags != nil {
				tags := scaler.Compute.Tags
				result = flattenTags(tags)
			}
			if result != nil {
				if err := resourceData.Set(string(Tags), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Tags), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if value, ok := resourceData.GetOk(string(Tags)); ok {
				if tags, err := expandTags(value); err != nil {
					return err
				} else {
					scaler.Compute.SetTags(tags)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(Tags))
			return err
		},
		nil,
	)

	fieldsMap[ConfigurationsFile] = commons.NewGenericField(
		commons.MRScalerAWS,
		ConfigurationsFile,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Bucket): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(Key): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			//Set: hashKV,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var result []interface{} = nil
			if scaler.Compute != nil && scaler.Compute.Configurations != nil &&
				scaler.Compute.Configurations.File != nil {
				s3FilePath := scaler.Compute.Configurations.File
				result = flattenS3File(s3FilePath)
			}
			if result != nil {
				if err := resourceData.Set(string(ConfigurationsFile), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ConfigurationsFile), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if value, ok := resourceData.GetOk(string(ConfigurationsFile)); ok {
				if s3FilePath, err := expandS3File(value); err != nil {
					return err
				} else {
					scaler.Compute.Configurations.SetFile(s3FilePath)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(ConfigurationsFile))
			return err
		},
		nil,
	)

	fieldsMap[ConfigurationsFile] = commons.NewGenericField(
		commons.MRScalerAWS,
		ConfigurationsFile,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Bucket): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(Key): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			//Set: hashKV,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var result []interface{} = nil
			if scaler.Compute != nil && scaler.Compute.Configurations != nil &&
				scaler.Compute.Configurations.File != nil {
				s3FilePath := scaler.Compute.Configurations.File
				result = flattenS3File(s3FilePath)
			}
			if result != nil {
				if err := resourceData.Set(string(ConfigurationsFile), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ConfigurationsFile), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if value, ok := resourceData.GetOk(string(ConfigurationsFile)); ok {
				if s3FilePath, err := expandS3File(value); err != nil {
					return err
				} else {
					scaler.Compute.SetConfigurations(&mrscaler.Configurations{File: s3FilePath})
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(ConfigurationsFile))
			return err
		},
		nil,
	)

	fieldsMap[BootstrapActionsFile] = commons.NewGenericField(
		commons.MRScalerAWS,
		BootstrapActionsFile,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Bucket): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(Key): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			//Set: hashBK,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var result []interface{} = nil
			if scaler.Compute != nil && scaler.Compute.BootstrapActions != nil &&
				scaler.Compute.BootstrapActions.File != nil {
				s3FilePath := scaler.Compute.BootstrapActions.File
				result = flattenS3File(s3FilePath)
			}
			if result != nil {
				if err := resourceData.Set(string(BootstrapActionsFile), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BootstrapActionsFile), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if value, ok := resourceData.GetOk(string(BootstrapActionsFile)); ok {
				if s3FilePath, err := expandS3File(value); err != nil {
					return err
				} else {
					scaler.Compute.SetBootstrapActions(&mrscaler.BootstrapActions{File: s3FilePath})
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(BootstrapActionsFile))
			return err
		},
		nil,
	)

	fieldsMap[StepsFile] = commons.NewGenericField(
		commons.MRScalerAWS,
		StepsFile,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Bucket): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(Key): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			//Set: hashBK,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var result []interface{} = nil
			if scaler.Compute != nil && scaler.Compute.Steps != nil &&
				scaler.Compute.Steps.File != nil {
				s3FilePath := scaler.Compute.Steps.File
				result = flattenS3File(s3FilePath)
			}
			if result != nil {
				if err := resourceData.Set(string(StepsFile), result); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(StepsFile), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if value, ok := resourceData.GetOk(string(StepsFile)); ok {
				if s3FilePath, err := expandS3File(value); err != nil {
					return err
				} else {
					scaler.Compute.SetSteps(&mrscaler.Steps{File: s3FilePath})
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(StepsFile))
			return err
		},
		nil,
	)

	fieldsMap[EBSRootVolumeSize] = commons.NewGenericField(
		commons.MRScalerAWS,
		EBSRootVolumeSize,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *int = nil
			if scaler.Compute != nil && scaler.Compute.EBSRootVolumeSize != nil {
				value = scaler.Compute.EBSRootVolumeSize
			}
			if err := resourceData.Set(string(EBSRootVolumeSize), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(EBSRootVolumeSize), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(EBSRootVolumeSize)); ok {
				scaler.Compute.SetEBSRootVolumeSize(spotinst.Int(v.(int)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(EBSRootVolumeSize))
			return err
		},
		nil,
	)

	fieldsMap[ManagedPrimarySecurityGroup] = commons.NewGenericField(
		commons.MRScalerAWS,
		ManagedPrimarySecurityGroup,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Compute != nil && scaler.Compute.ManagedPrimarySecurityGroup != nil {
				value = scaler.Compute.ManagedPrimarySecurityGroup
			}
			if err := resourceData.Set(string(ManagedPrimarySecurityGroup), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ManagedPrimarySecurityGroup), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(ManagedPrimarySecurityGroup)); ok {
				scaler.Compute.SetManagedPrimarySecurityGroup(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(ManagedPrimarySecurityGroup))
			return err
		},
		nil,
	)

	fieldsMap[ManagedReplicaSecurityGroup] = commons.NewGenericField(
		commons.MRScalerAWS,
		ManagedReplicaSecurityGroup,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Compute != nil && scaler.Compute.ManagedReplicaSecurityGroup != nil {
				value = scaler.Compute.ManagedReplicaSecurityGroup
			}
			if err := resourceData.Set(string(ManagedReplicaSecurityGroup), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ManagedReplicaSecurityGroup), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(ManagedReplicaSecurityGroup)); ok {
				scaler.Compute.SetManagedReplicaSecurityGroup(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(ManagedReplicaSecurityGroup))
			return err
		},
		nil,
	)

	fieldsMap[ServiceAccessSecurityGroup] = commons.NewGenericField(
		commons.MRScalerAWS,
		ServiceAccessSecurityGroup,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Compute != nil && scaler.Compute.ServiceAccessSecurityGroup != nil {
				value = scaler.Compute.ServiceAccessSecurityGroup
			}
			if err := resourceData.Set(string(ServiceAccessSecurityGroup), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ServiceAccessSecurityGroup), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(ServiceAccessSecurityGroup)); ok {
				scaler.Compute.SetServiceAccessSecurityGroup(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(ServiceAccessSecurityGroup))
			return err
		},
		nil,
	)

	fieldsMap[CustomAMIID] = commons.NewGenericField(
		commons.MRScalerAWS,
		CustomAMIID,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Compute != nil && scaler.Compute.CustomAMIID != nil {
				value = scaler.Compute.CustomAMIID
			}
			if err := resourceData.Set(string(CustomAMIID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(CustomAMIID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(CustomAMIID)); ok {
				scaler.Compute.SetCustomAMIID(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(CustomAMIID))
			return err
		},
		nil,
	)

	fieldsMap[RepoUpgradeOnBoot] = commons.NewGenericField(
		commons.MRScalerAWS,
		RepoUpgradeOnBoot,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Compute != nil && scaler.Compute.RepoUpgradeOnBoot != nil {
				value = scaler.Compute.RepoUpgradeOnBoot
			}
			if err := resourceData.Set(string(RepoUpgradeOnBoot), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RepoUpgradeOnBoot), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(RepoUpgradeOnBoot)); ok && v != "" {
				scaler.Compute.SetRepoUpgradeOnBoot(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(RepoUpgradeOnBoot))
			return err
		},
		nil,
	)

	fieldsMap[EC2KeyName] = commons.NewGenericField(
		commons.MRScalerAWS,
		EC2KeyName,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			var value *string = nil
			if scaler.Compute != nil && scaler.Compute.EC2KeyName != nil {
				value = scaler.Compute.EC2KeyName
			}
			if err := resourceData.Set(string(EC2KeyName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(EC2KeyName), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(EC2KeyName)); ok {
				scaler.Compute.SetEC2KeyName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(EC2KeyName))
			return err
		},
		nil,
	)

	fieldsMap[AddlPrimarySecurityGroups] = commons.NewGenericField(
		commons.MRScalerAWS,
		AddlPrimarySecurityGroups,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if value, ok := resourceData.GetOk(string(AddlPrimarySecurityGroups)); ok {
				if groups, err := expandGenericList(value); err != nil {
					return err
				} else {
					scaler.Compute.SetAdditionalPrimarySecurityGroups(groups)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(AddlPrimarySecurityGroups))
			return err
		},
		nil,
	)

	fieldsMap[AddlReplicaSecurityGroups] = commons.NewGenericField(
		commons.MRScalerAWS,
		AddlReplicaSecurityGroups,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if value, ok := resourceData.GetOk(string(AddlPrimarySecurityGroups)); ok {
				if groups, err := expandGenericList(value); err != nil {
					return err
				} else {
					scaler.Compute.SetAdditionalReplicaSecurityGroups(groups)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(AddlReplicaSecurityGroups))
			return err
		},
		nil,
	)

	fieldsMap[Applications] = commons.NewGenericField(
		commons.MRScalerAWS,
		Applications,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Args): {
						Type:     schema.TypeList,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},

					string(AppName): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(Version): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(Applications)); ok {
				if apps, err := expandApplications(v); err != nil {
					return err
				} else {
					scaler.Compute.SetApplications(apps)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			err := fmt.Errorf(string(commons.FieldUpdateNotAllowedPattern),
				string(Applications))
			return err
		},
		nil,
	)

	fieldsMap[InstanceWeights] = commons.NewGenericField(
		commons.MRScalerAWS,
		InstanceWeights,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{

					string(InstanceType): {
						Type:     schema.TypeString,
						Required: true,
					},

					string(WeightedCapacity): {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			if v, ok := resourceData.GetOk(string(InstanceWeights)); ok {
				if apps, err := expandInstanceWeights(v); err != nil {
					return err
				} else {
					scaler.Compute.SetInstanceWeights(apps)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			mrsWrapper := resourceObject.(*commons.MRScalerAWSWrapper)
			scaler := mrsWrapper.GetMRScalerAWS()
			values := []*mrscaler.InstanceWeight{}
			if v, ok := resourceData.GetOk(string(InstanceWeights)); ok {
				if weights, err := expandInstanceWeights(v); err != nil {
					return err
				} else {
					values = weights
				}
			}
			scaler.Compute.SetInstanceWeights(values)
			return nil
		},
		nil,
	)

}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//             Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func expandMRScalerAWSAvailabilityZones(data interface{}) ([]*mrscaler.AvailabilityZone, error) {
	list := data.([]interface{})
	zones := make([]*mrscaler.AvailabilityZone, 0, len(list))
	for _, str := range list {
		if s, ok := str.(string); ok {
			parts := strings.Split(s, ":")
			zone := &mrscaler.AvailabilityZone{}
			if len(parts) >= 1 && parts[0] != "" {
				zone.SetName(spotinst.String(parts[0]))
			}
			if len(parts) == 2 && parts[1] != "" {
				zone.SetSubnetId(spotinst.String(parts[1]))
			}

			zones = append(zones, zone)
		}
	}
	return zones, nil
}

func expandGenericList(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if item, ok := v.(string); ok && item != "" {
			result = append(result, item)
		}
	}
	return result, nil
}

func expandS3File(data interface{}) (*mrscaler.S3File, error) {
	list := data.(*schema.Set).List()
	s3FilePath := &mrscaler.S3File{}
	for _, item := range list {
		m := item.(map[string]interface{})

		if v, ok := m[string(Key)].(string); ok {
			s3FilePath.SetKey(spotinst.String(v))
		}

		if v, ok := m[string(Bucket)].(string); ok {
			s3FilePath.SetBucket(spotinst.String(v))
		}
	}
	return s3FilePath, nil
}

func expandTags(data interface{}) ([]*mrscaler.Tag, error) {
	list := data.(*schema.Set).List()
	tags := make([]*mrscaler.Tag, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr[string(TagKey)]; !ok {
			return nil, errors.New("invalid tag attributes: key missing")
		}

		if _, ok := attr[string(TagValue)]; !ok {
			return nil, errors.New("invalid tag attributes: value missing")
		}
		tag := &mrscaler.Tag{
			Key:   spotinst.String(attr[string(TagKey)].(string)),
			Value: spotinst.String(attr[string(TagValue)].(string)),
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func expandApplications(data interface{}) ([]*mrscaler.Application, error) {
	list := data.(*schema.Set).List()
	apps := make([]*mrscaler.Application, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		app := &mrscaler.Application{}

		if v, ok := m[string(AppName)].(string); ok && v != "" {
			app.SetName(spotinst.String(v))
		}

		if v, ok := m[string(Version)].(string); ok && v != "" {
			app.SetVersion(spotinst.String(v))
		}

		if v, ok := m[string(Args)]; ok {
			args, err := expandGenericList(v)
			if err != nil {
				return nil, err
			} else {
				app.SetArgs(args)
			}
		}

		apps = append(apps, app)
	}

	return apps, nil
}

func expandInstanceWeights(data interface{}) ([]*mrscaler.InstanceWeight, error) {
	list := data.(*schema.Set).List()
	weights := make([]*mrscaler.InstanceWeight, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		weight := &mrscaler.InstanceWeight{}

		if v, ok := m[string(InstanceType)].(string); ok && v != "" {
			weight.SetInstanceType(spotinst.String(v))
		}

		if v, ok := m[string(WeightedCapacity)].(int); ok {
			weight.SetWeightedCapacity(spotinst.Int(v))
		}

		weights = append(weights, weight)
	}

	return weights, nil
}

func flattenS3File(file *mrscaler.S3File) []interface{} {
	m := make(map[string]interface{})
	m[string(Bucket)] = spotinst.StringValue(file.Bucket)
	m[string(Key)] = spotinst.StringValue(file.Key)
	return []interface{}{m}
}

func flattenTags(tags []*mrscaler.Tag) []interface{} {
	result := make([]interface{}, 0, len(tags))
	for _, tag := range tags {
		m := make(map[string]interface{})
		m[string(TagKey)] = spotinst.StringValue(tag.Key)
		m[string(TagValue)] = spotinst.StringValue(tag.Value)

		result = append(result, m)
	}
	return result
}

func hashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m[string(TagKey)].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m[string(TagValue)].(string)))
	return hashcode.String(buf.String())
}

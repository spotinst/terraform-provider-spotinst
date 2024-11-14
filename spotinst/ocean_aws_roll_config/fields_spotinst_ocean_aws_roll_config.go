package ocean_aws_roll_config

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[LaunchSpecIDs] = commons.NewGenericField(
		commons.OceanAWSRoll,
		LaunchSpecIDs,
		&schema.Schema{
			Type:          schema.TypeList,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{string(InstanceIds)},
			Elem: &schema.Schema{
				Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			var value []string = nil
			if roll != nil && roll.LaunchSpecIDs != nil {
				value = roll.LaunchSpecIDs
			}
			if err := resourceData.Set(string(LaunchSpecIDs), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(LaunchSpecIDs), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			if value, ok := resourceData.GetOk(string(LaunchSpecIDs)); ok {
				if list, err := expandIdList(value); err != nil {
					return err
				} else {
					roll.SetLaunchSpecIDs(list)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			/*rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			if value, ok := resourceData.GetOk(string(LaunchSpecIDs)); ok {
				if list, err := expandIdList(value); err != nil {
					return err
				} else {
					roll.SetLaunchSpecIDs(list)
				}
			} else {
				roll.SetLaunchSpecIDs(nil)
			}*/
			return nil
		},
		nil,
	)

	fieldsMap[InstanceIds] = commons.NewGenericField(
		commons.OceanAWSRoll,
		InstanceIds,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &schema.Schema{
				Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			var value []string = nil
			if roll != nil && roll.InstanceIDs != nil {
				value = roll.InstanceIDs
			}
			if err := resourceData.Set(string(InstanceIds), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(InstanceIds), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			if value, ok := resourceData.GetOk(string(InstanceIds)); ok {
				if list, err := expandIdList(value); err != nil {
					return err
				} else {
					roll.SetInstanceIDs(list)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			/*rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			if value, ok := resourceData.GetOk(string(InstanceIds)); ok {
				if list, err := expandIdList(value); err != nil {
					return err
				} else {
					roll.SetInstanceIDs(list)
				}
			} else {
				roll.SetInstanceIDs(nil)
			}*/
			return nil
		},
		nil,
	)

	fieldsMap[Comment] = commons.NewGenericField(
		commons.OceanAWSRoll,
		Comment,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			var value *string = nil
			if roll != nil && roll.Comment != nil {
				value = roll.Comment
			}
			if err := resourceData.Set(string(Comment), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Comment), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			if value, ok := resourceData.Get(string(Comment)).(string); ok && value != "" {
				roll.SetComment(spotinst.String(value))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			/*rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			roll.SetComment(spotinst.String(resourceData.Get(string(Comment)).(string)))*/
			/*if value, ok := resourceData.GetOk(string(Comment)); ok {
				roll.SetComment(spotinst.String(value.(string)))
			}*/
			return nil
		},
		nil,
	)

	fieldsMap[BatchSizePercentage] = commons.NewGenericField(
		commons.OceanAWSRoll,
		BatchSizePercentage,
		&schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			var value *int = nil
			if roll != nil && roll.BatchSizePercentage != nil {
				value = roll.BatchSizePercentage
			}
			if value != nil {
				if err := resourceData.Set(string(BatchSizePercentage), spotinst.IntValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BatchSizePercentage), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			if v, ok := resourceData.Get(string(BatchSizePercentage)).(int); ok && v >= 0 {
				roll.SetBatchSizePercentage(spotinst.Int(v))
			} else {
				roll.SetBatchSizePercentage(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			/*rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			roll.SetBatchSizePercentage(spotinst.Int(resourceData.Get(string(BatchSizePercentage)).(int)))*/
			/*if v, ok := resourceData.Get(string(BatchSizePercentage)).(int); ok && v >= 0 {
				roll.SetBatchSizePercentage(spotinst.Int(v))
			} else {
				roll.SetBatchSizePercentage(nil)
			}*/
			return nil
		},
		nil,
	)

	fieldsMap[BatchMinHealthyPercentage] = commons.NewGenericField(
		commons.OceanAWSRoll,
		BatchMinHealthyPercentage,
		&schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			var value *int = nil
			if roll != nil && roll.BatchMinHealthyPercentage != nil {
				value = roll.BatchMinHealthyPercentage
			}
			if value != nil {
				if err := resourceData.Set(string(BatchMinHealthyPercentage), spotinst.IntValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(BatchMinHealthyPercentage), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			if v, ok := resourceData.Get(string(BatchMinHealthyPercentage)).(int); ok && v >= 0 {
				roll.SetBatchMinHealthyPercentage(spotinst.Int(v))
			} else {
				roll.SetBatchMinHealthyPercentage(nil)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			/*rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			roll.SetBatchMinHealthyPercentage(spotinst.Int(resourceData.Get(string(BatchMinHealthyPercentage)).(int)))*/
			/*if v, ok := resourceData.Get(string(BatchMinHealthyPercentage)).(int); ok && v >= 0 {
				roll.SetBatchMinHealthyPercentage(spotinst.Int(v))
			} else {
				roll.SetBatchMinHealthyPercentage(nil)
			}*/
			return nil
		},
		nil,
	)

	fieldsMap[RespectPDB] = commons.NewGenericField(
		commons.OceanAWSRoll,
		RespectPDB,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			var value *bool = nil
			if roll != nil && roll.RespectPDB != nil {
				value = roll.RespectPDB
			}
			if value != nil {
				if err := resourceData.Set(string(RespectPDB), spotinst.BoolValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(RespectPDB), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			if v, ok := resourceData.GetOkExists(string(RespectPDB)); ok && v != nil {
				pdb := v.(bool)
				respectPdb := spotinst.Bool(pdb)
				roll.SetRespectPDB(respectPdb)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			/*rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			roll.SetRespectPDB(spotinst.Bool(resourceData.Get(string(RespectPDB)).(bool)))*/
			/*var respectPdb *bool = nil
			if v, ok := resourceData.GetOkExists(string(RespectPDB)); ok && v != nil {
				pdb := v.(bool)
				respectPdb = spotinst.Bool(pdb)
			}
			roll.SetRespectPDB(respectPdb)*/
			return nil
		},
		nil,
	)

	fieldsMap[DisableLaunchSpecAutoScaling] = commons.NewGenericField(
		commons.OceanAWSRoll,
		DisableLaunchSpecAutoScaling,
		&schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			var value *bool = nil
			if roll != nil && roll.DisableLaunchSpecAutoScaling != nil {
				value = roll.DisableLaunchSpecAutoScaling
			}
			if value != nil {
				if err := resourceData.Set(string(DisableLaunchSpecAutoScaling), spotinst.BoolValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DisableLaunchSpecAutoScaling), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			if v, ok := resourceData.GetOkExists(string(DisableLaunchSpecAutoScaling)); ok && v != nil {
				autoScaling := v.(bool)
				disable := spotinst.Bool(autoScaling)
				roll.SetDisableLaunchSpecAutoScaling(disable)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			/*rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			roll.SetDisableLaunchSpecAutoScaling(spotinst.Bool(resourceData.Get(string(DisableLaunchSpecAutoScaling)).(bool)))*/
			/*var disable *bool = nil
			if v, ok := resourceData.GetOkExists(string(DisableLaunchSpecAutoScaling)); ok && v != nil {
				autoScaling := v.(bool)
				disable = spotinst.Bool(autoScaling)
			}
			roll.SetDisableLaunchSpecAutoScaling(disable)*/
			return nil
		},
		nil,
	)

	fieldsMap[OceanClusterId] = commons.NewGenericField(
		commons.OceanAWSRoll,
		OceanClusterId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			//			ForceNew: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			var value *string = nil
			if roll != nil && roll.ClusterID != nil {
				value = roll.ClusterID
			}
			if value != nil {
				if err := resourceData.Set(string(OceanClusterId), spotinst.StringValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(OceanClusterId), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			if v, ok := resourceData.GetOk(string(OceanClusterId)); ok && v != "" {
				roll.SetClusterID(spotinst.String(resourceData.Get(string(OceanClusterId)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			/*rollWrapper := resourceObject.(*commons.OceanAWSRollWrapper)
			roll := rollWrapper.GetRoll()
			if v, ok := resourceData.GetOk(string(OceanClusterId)); ok && v != "" {
				roll.SetClusterID(spotinst.String(resourceData.Get(string(OceanClusterId)).(string)))
			}*/
			return nil
		},
		nil,
	)
}

func expandIdList(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if ids, ok := v.(string); ok && ids != "" {
			result = append(result, ids)
		}
	}
	return result, nil
}

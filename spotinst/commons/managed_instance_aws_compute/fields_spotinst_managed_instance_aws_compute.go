package managed_instance_aws_compute

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[SubnetIDs] = commons.NewGenericField(
		commons.ManagedInstanceAWSCompute,
		SubnetIDs,
		&schema.Schema{
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value []string = nil
			if managedInstance.Compute != nil && managedInstance.Compute.SubnetIDs != nil {
				value = managedInstance.Compute.SubnetIDs
			}
			if err := resourceData.Set(string(SubnetIDs), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SubnetIDs), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if value, ok := resourceData.GetOk(string(SubnetIDs)); ok && value != nil {
				if subnetIds, err := expandSubnetIDs(value); err != nil {
					return err
				} else {
					managedInstance.Compute.SetSubnetIDs(subnetIds)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if value, ok := resourceData.GetOk(string(SubnetIDs)); ok && value != nil {
				if subnetIds, err := expandSubnetIDs(value); err != nil {
					return err
				} else {
					managedInstance.Compute.SetSubnetIDs(subnetIds)
				}
			}
			return nil
		},
		nil,
	)

	fieldsMap[VpcID] = commons.NewGenericField(
		commons.ManagedInstanceAWSCompute,
		VpcID,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *string = nil
			if managedInstance.Compute != nil && managedInstance.Compute.VpcID != nil {
				value = managedInstance.Compute.VpcID
			}
			if err := resourceData.Set(string(VpcID), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(VpcID), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			managedInstance.Compute.SetVpcId(spotinst.String(resourceData.Get(string(VpcID)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			managedInstance.Compute.SetVpcId(spotinst.String(resourceData.Get(string(VpcID)).(string)))
			return nil
		},
		nil,
	)

	fieldsMap[ElasticIP] = commons.NewGenericField(
		commons.ManagedInstanceAWSCompute,
		ElasticIP,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *string = nil
			if managedInstance.Compute != nil && managedInstance.Compute.ElasticIP != nil {
				value = managedInstance.Compute.ElasticIP
			}
			if err := resourceData.Set(string(ElasticIP), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ElasticIP), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if value, ok := resourceData.GetOk(string(ElasticIP)); ok && value != nil {
				managedInstance.Compute.SetElasticIP(spotinst.String(resourceData.Get(string(ElasticIP)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if value, ok := resourceData.GetOk(string(ElasticIP)); ok && value != nil {
				managedInstance.Compute.SetElasticIP(spotinst.String(resourceData.Get(string(ElasticIP)).(string)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[PrivateIP] = commons.NewGenericField(
		commons.ManagedInstanceAWSCompute,
		PrivateIP,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			var value *string = nil
			if managedInstance.Compute != nil && managedInstance.Compute.PrivateIP != nil {
				value = managedInstance.Compute.PrivateIP
			}
			if err := resourceData.Set(string(PrivateIP), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(PrivateIP), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if value, ok := resourceData.GetOk(string(PrivateIP)); ok && value != nil {
				managedInstance.Compute.SetPrivateIP(spotinst.String(resourceData.Get(string(PrivateIP)).(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			miWrapper := resourceObject.(*commons.MangedInstanceAWSWrapper)
			managedInstance := miWrapper.GetManagedInstance()
			if value, ok := resourceData.GetOk(string(PrivateIP)); ok && value != nil {
				managedInstance.Compute.SetPrivateIP(spotinst.String(resourceData.Get(string(PrivateIP)).(string)))
			}
			return nil
		},
		nil,
	)
}

func expandSubnetIDs(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if subnetID, ok := v.(string); ok && subnetID != "" {
			result = append(result, subnetID)
		}
	}

	return result, nil
}

package oceancd_verification_provider

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ClusterIDs] = commons.NewGenericField(
		commons.OceanCDVerificationProvider,
		ClusterIDs,
		&schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var value []string = nil
			if verificationProvider.ClusterIDs != nil {
				value = verificationProvider.ClusterIDs
			}
			if err := resourceData.Set(string(ClusterIDs), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ClusterIDs), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			if value, ok := resourceData.GetOk(string(ClusterIDs)); ok {
				if clusterIds, err := expandClusterIDs(value); err != nil {
					return err
				} else {
					verificationProvider.SetClusterIDs(clusterIds)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			if value, ok := resourceData.GetOk(string(ClusterIDs)); ok {
				if clusterIds, err := expandClusterIDs(value); err != nil {
					return err
				} else {
					verificationProvider.SetClusterIDs(clusterIds)
				}
			} else {
				verificationProvider.SetClusterIDs(nil)
			}
			return nil
		},
		nil,
	)

	fieldsMap[Name] = commons.NewGenericField(
		commons.OceanCDVerificationProvider,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			var value *string = nil
			if verificationProvider.Name != nil {
				value = verificationProvider.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			verificationProviderWrapper := resourceObject.(*commons.OceanCDVerificationProviderWrapper)
			verificationProvider := verificationProviderWrapper.GetVerificationProvider()
			if value, ok := resourceData.Get(string(Name)).(string); ok && value != "" {
				verificationProvider.SetName(spotinst.String(value))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)
}

func expandClusterIDs(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))

	for _, v := range list {
		if clusterIds, ok := v.(string); ok && clusterIds != "" {
			result = append(result, clusterIds)
		}
	}
	return result, nil
}

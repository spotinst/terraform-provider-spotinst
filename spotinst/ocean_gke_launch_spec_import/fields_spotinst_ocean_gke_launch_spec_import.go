package ocean_gke_launch_spec_import

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[OceanId] = commons.NewGenericField(
		commons.OceanGKELaunchSpecImport,
		OceanId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
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
		commons.OceanGKELaunchSpecImport,
		NodePoolName,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
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
}

package account

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Name] = commons.NewGenericField(
		commons.Account,
		Name,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			accountWrapper := resourceObject.(*commons.AccountWrapper)
			account := accountWrapper.GetAccount()
			var value *string = nil
			if account.Name != nil {
				value = account.Name
			}
			if err := resourceData.Set(string(Name), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Name), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			accountWrapper := resourceObject.(*commons.AccountWrapper)
			account := accountWrapper.GetAccount()
			account.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			accountWrapper := resourceObject.(*commons.AccountWrapper)
			account := accountWrapper.GetAccount()
			account.SetName(spotinst.String(resourceData.Get(string(Name)).(string)))
			return nil
		},
		nil,
	)

}

package commons

import (
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/common"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	AccountResourceName ResourceName = "spotinst_account"
)

var AccountResource *AccountTerraformResource

type AccountTerraformResource struct {
	GenericResource
}

type AccountWrapper struct {
	account *common.Account
}

func NewAccountResource(fieldsMap map[FieldName]*GenericField) *AccountTerraformResource {
	return &AccountTerraformResource{
		GenericResource: GenericResource{
			resourceName: AccountResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *AccountTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*common.Account, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	accountWrapper := NewCommonAccountWrapper()

	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(accountWrapper, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return accountWrapper.GetAccount(), nil
}

func (res *AccountTerraformResource) OnRead(
	account *common.Account,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	accountWrapper := NewCommonAccountWrapper()
	accountWrapper.SetAccount(account)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(accountWrapper, resourceData, meta); err != nil {
			return err
		}
	}

	return nil
}

func NewCommonAccountWrapper() *AccountWrapper {
	return &AccountWrapper{
		account: &common.Account{},
	}
}

func (accountWrapper *AccountWrapper) GetAccount() *common.Account {
	return accountWrapper.account
}

func (accountWrapper *AccountWrapper) SetAccount(account *common.Account) {
	accountWrapper.account = account
}

package commons

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/account/providers/aws"
)

const (
	AccountAWSResourceName ResourceName = "spotinst_account_aws"
)

var AccountAWSResource *AccountAWSTerraformResource

type AccountAWSTerraformResource struct {
	GenericResource
}

type AWSAccountWrapper struct {
	account *aws.Account
}

func NewAccountAWSResource(fieldsMap map[FieldName]*GenericField) *AccountAWSTerraformResource {
	return &AccountAWSTerraformResource{
		GenericResource: GenericResource{
			resourceName: AccountAWSResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *AccountAWSTerraformResource) OnCreate(
	resourceData *schema.ResourceData,
	meta interface{}) (*aws.Account, error) {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	accountWrapper := NewAccountWrapper()

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

func (res *AccountAWSTerraformResource) OnRead(
	account *aws.Account,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	accountWrapper := NewAccountWrapper()
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

func NewAccountWrapper() *AWSAccountWrapper {
	return &AWSAccountWrapper{
		account: &aws.Account{},
	}
}

func (accountWrapper *AWSAccountWrapper) GetAccount() *aws.Account {
	return accountWrapper.account
}

func (accountWrapper *AWSAccountWrapper) SetAccount(account *aws.Account) {
	accountWrapper.account = account
}

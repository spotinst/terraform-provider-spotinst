package commons

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/notificationcenter"
	"log"
)

const NotificationCenterResourceName ResourceName = "spotinst_notification_center"

var NotificationCenterResource *NotificationCenterTerraformResource

type NotificationCenterTerraformResource struct {
	GenericResource
}

type NotificationCenterWrapper struct {
	notificationCenter *notificationcenter.NotificationCenter
}

func NewNotificationCenterResource(fieldsMap map[FieldName]*GenericField) *NotificationCenterTerraformResource {
	return &NotificationCenterTerraformResource{
		GenericResource: GenericResource{
			resourceName: NotificationCenterResourceName,
			fields:       NewGenericFields(fieldsMap),
		},
	}
}

func (res *NotificationCenterTerraformResource) OnRead(
	notificationCenter *notificationcenter.NotificationCenter,
	resourceData *schema.ResourceData,
	meta interface{}) error {

	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return fmt.Errorf("resource fields are nil or empty, cannot read")
	}

	notificationWrapper := NewNotificationCenterWrapper()
	notificationWrapper.SetNotificationCenter(notificationCenter)

	for _, field := range res.fields.fieldsMap {
		if field.onRead == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnRead), field.resourceAffinity, field.fieldNameStr)
		if err := field.onRead(notificationWrapper, resourceData, meta); err != nil {
			return err
		}
	}
	return nil
}

func (res *NotificationCenterTerraformResource) OnCreate(resourceData *schema.ResourceData,
	meta interface{}) (*notificationcenter.NotificationCenter, error) {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return nil, fmt.Errorf("resource fields are nil or empty, cannot create")
	}

	nc := NewNotificationCenterWrapper()
	for _, field := range res.fields.fieldsMap {
		if field.onCreate == nil {
			continue
		}
		log.Printf(string(ResourceFieldOnCreate), field.resourceAffinity, field.fieldNameStr)
		if err := field.onCreate(nc, resourceData, meta); err != nil {
			return nil, err
		}
	}
	return nc.GetNotificationCenter(), nil
}

func (res *NotificationCenterTerraformResource) OnUpdate(resourceData *schema.ResourceData,
	meta interface{}) (bool, *notificationcenter.NotificationCenter, error) {
	if res.fields == nil || res.fields.fieldsMap == nil || len(res.fields.fieldsMap) == 0 {
		return false, nil, fmt.Errorf("resource fields are nil or empty, cannot update")
	}

	nc := NewNotificationCenterWrapper()
	hasChanged := false
	for _, field := range res.fields.fieldsMap {
		if field.onUpdate == nil {
			continue
		}
		if field.hasFieldChange(resourceData, meta) {
			log.Printf(string(ResourceFieldOnUpdate), field.resourceAffinity, field.fieldNameStr)
			if err := field.onUpdate(nc, resourceData, meta); err != nil {
				return false, nil, err
			}
			hasChanged = true
		}
	}
	return hasChanged, nc.GetNotificationCenter(), nil
}

func NewNotificationCenterWrapper() *NotificationCenterWrapper {
	return &NotificationCenterWrapper{
		notificationCenter: &notificationcenter.NotificationCenter{},
	}
}

func (ncWrapper *NotificationCenterWrapper) GetNotificationCenter() *notificationcenter.NotificationCenter {
	return ncWrapper.notificationCenter
}

func (ncWrapper *NotificationCenterWrapper) SetNotificationCenter(notificationCenter *notificationcenter.NotificationCenter) {
	ncWrapper.notificationCenter = notificationCenter
}

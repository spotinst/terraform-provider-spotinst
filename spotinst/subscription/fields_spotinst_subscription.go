package subscription

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/subscription"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[ResourceId] = commons.NewGenericField(
		commons.Subscription,
		ResourceId,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if err := resourceData.Set(string(ResourceId), sub.ResourceID); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(ResourceId), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if v, ok := resourceData.Get(string(ResourceId)).(string); ok {
				sub.SetResourceId(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if v, ok := resourceData.Get(string(ResourceId)).(string); ok {
				sub.SetResourceId(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[EventType] = commons.NewGenericField(
		commons.Subscription,
		EventType,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			StateFunc: func(v interface{}) string {
				value := v.(string)
				return strings.ToUpper(value)
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if err := resourceData.Set(string(EventType), sub.EventType); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(EventType), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if v, ok := resourceData.Get(string(EventType)).(string); ok {
				sub.SetEventType(spotinst.String(strings.ToUpper(v)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if v, ok := resourceData.Get(string(EventType)).(string); ok {
				sub.SetEventType(spotinst.String(strings.ToUpper(v)))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Protocol] = commons.NewGenericField(
		commons.Subscription,
		Protocol,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if err := resourceData.Set(string(Protocol), sub.Protocol); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Protocol), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if v, ok := resourceData.Get(string(Protocol)).(string); ok {
				sub.SetProtocol(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if v, ok := resourceData.Get(string(Protocol)).(string); ok {
				sub.SetProtocol(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Endpoint] = commons.NewGenericField(
		commons.Subscription,
		Endpoint,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if err := resourceData.Set(string(Endpoint), sub.Endpoint); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Endpoint), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if v, ok := resourceData.Get(string(Endpoint)).(string); ok {
				sub.SetEndpoint(spotinst.String(v))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if v, ok := resourceData.Get(string(Endpoint)).(string); ok {
				sub.SetEndpoint(spotinst.String(v))
			}
			return nil
		},
		nil,
	)

	fieldsMap[Format] = commons.NewGenericField(
		commons.Subscription,
		Format,
		&schema.Schema{
			Type:     schema.TypeMap,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if err := resourceData.Set(string(Format), sub.Format); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(Format), err)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if v, ok := resourceData.Get(string(Format)).(map[string]interface{}); ok {
				sub.SetFormat(v)
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			sub := resourceObject.(*subscription.Subscription)
			if v, ok := resourceData.Get(string(Format)).(map[string]interface{}); ok {
				sub.SetFormat(v)
			}
			return nil
		},
		nil,
	)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

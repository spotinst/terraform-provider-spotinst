package spotinst

import (
	"context"
	"fmt"
	"github.com/spotinst/spotinst-sdk-go/service/notificationcenter"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func createNotificationCenterResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.NotificationCenterResourceName), name)
}

func testNotificationCeneterDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.NotificationCenterResourceName) {
			continue
		}
		input := &notificationcenter.ReadNotificationCenterPolicyInput{PolicyId: spotinst.String(rs.Primary.ID)}
		resp, err := client.notificationCenter.ReadNotificationCenterPolicy(context.Background(), input)
		if err == nil && resp != nil && resp.NotificationCenter != nil {
			return fmt.Errorf("notification center still exists")
		}
	}
	return nil
}

func testCheckNotificationCenterExists(sub *notificationcenter.NotificationCenter, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &notificationcenter.ReadNotificationCenterPolicyInput{PolicyId: spotinst.String(rs.Primary.ID)}
		resp, err := client.notificationCenter.ReadNotificationCenterPolicy(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.NotificationCenter.PrivacyLevel) != rs.Primary.Attributes["privacy_level"] {
			return fmt.Errorf("Notification Center Policy not found: %+v,\n %+v\n", resp.NotificationCenter, rs.Primary.Attributes)
		}
		*sub = *resp.NotificationCenter
		return nil
	}
}

func createNotificationCenterTerraform(tfResource string, resourceName string) string {
	template := ""

	template += fmt.Sprintf(tfResource, resourceName)

	log.Printf("Terraform [%v] template:\n%v", resourceName, template)
	return template
}

// region Subscription: Http
func TestAccSpotinstNotificationCenter(t *testing.T) {
	notificationcenterName := "notification-center"
	ncResourceName := createNotificationCenterResourceName(notificationcenterName)

	var notificationcenter notificationcenter.NotificationCenter
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testNotificationCeneterDestroy,

		Steps: []resource.TestStep{
			{
				Config: createNotificationCenterTerraform(testNotificationCenter_Create, notificationcenterName),
				Check: resource.ComposeTestCheckFunc(
					testCheckNotificationCenterExists(&notificationcenter, ncResourceName),
					resource.TestCheckResourceAttr(ncResourceName, "name", "Notification-center"),
					resource.TestCheckResourceAttr(ncResourceName, "description", "Testing of notification center policy"),
					resource.TestCheckResourceAttr(ncResourceName, "is_active", "true"),
					resource.TestCheckResourceAttr(ncResourceName, "privacy_level", "public"),
					resource.TestCheckResourceAttr(ncResourceName, "registered_users.#", "1"),
					resource.TestCheckResourceAttr(ncResourceName, "registered_users.0.subscription_types.#", "1"),
					resource.TestCheckResourceAttr(ncResourceName, "registered_users.0.subscription_types.0", "email"),
					resource.TestCheckResourceAttr(ncResourceName, "registered_users.0.user_email", "TestAutomation_Admin_DO_NOT_DELETE@spot.io"),
					resource.TestCheckResourceAttr(ncResourceName, "subscriptions.#", "1"),
					resource.TestCheckResourceAttr(ncResourceName, "subscriptions.0.endpoint", "https://webhook.si"),
					resource.TestCheckResourceAttr(ncResourceName, "subscriptions.0.subscription_type", "webhook"),
					resource.TestCheckResourceAttr(ncResourceName, "compute_policy_config.#", "1"),
					resource.TestCheckResourceAttr(ncResourceName, "compute_policy_config.0.events.#", "1"),
					resource.TestCheckResourceAttr(ncResourceName, "compute_policy_config.0.events.0.event", "Maximum capacity reached"),
					resource.TestCheckResourceAttr(ncResourceName, "compute_policy_config.0.events.0.event_type", "WARN"),
					resource.TestCheckResourceAttr(ncResourceName, "compute_policy_config.0.dynamic_rules.#", "1"),
					resource.TestCheckResourceAttr(ncResourceName, "compute_policy_config.0.dynamic_rules.filter_conditions#", "1"),
					resource.TestCheckResourceAttr(ncResourceName, "compute_policy_config.0.dynamic_rules.filter_conditions.0.expression", "DO-NOT-DELETE"),
					resource.TestCheckResourceAttr(ncResourceName, "compute_policy_config.0.dynamic_rules.filter_conditions.0.identifier", "resource_name"),
					resource.TestCheckResourceAttr(ncResourceName, "compute_policy_config.0.dynamic_rules.filter_conditions.0.operator", "contains"),
				),
			},
			{
				Config: createNotificationCenterTerraform(testNotificationCenter_Update, notificationcenterName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(ncResourceName, "name", "Notification-center"),
					resource.TestCheckResourceAttr(ncResourceName, "description", "Update Testing of notification center policy"),
					resource.TestCheckResourceAttr(ncResourceName, "is_active", "true"),
					resource.TestCheckResourceAttr(ncResourceName, "privacy_level", "public"),
					resource.TestCheckResourceAttr(ncResourceName, "registered_users.#", "1"),
					resource.TestCheckResourceAttr(ncResourceName, "registered_users.0.subscription_types.#", "1"),
					resource.TestCheckResourceAttr(ncResourceName, "registered_users.0.subscription_types.0", "email"),
					resource.TestCheckResourceAttr(ncResourceName, "registered_users.0.user_email", "Automation@spotinst.com"),
					resource.TestCheckResourceAttr(ncResourceName, "subscriptions.#", "1"),
					resource.TestCheckResourceAttr(ncResourceName, "subscriptions.0.endpoint", "https://webhook.si"),
					resource.TestCheckResourceAttr(ncResourceName, "subscriptions.0.subscription_type", "webhook"),
					resource.TestCheckResourceAttr(ncResourceName, "compute_policy_config.#", "1"),
					resource.TestCheckResourceAttr(ncResourceName, "compute_policy_config.0.events.#", "1"),
					resource.TestCheckResourceAttr(ncResourceName, "compute_policy_config.0.events.0.event", "Stateful recycle finished"),
					resource.TestCheckResourceAttr(ncResourceName, "compute_policy_config.0.events.0.event_type", "INFO"),
					resource.TestCheckResourceAttr(ncResourceName, "compute_policy_config.0.should_include_all_resources", "true"),
				),
			},
		},
	})
}

const testNotificationCenter_Create = `
resource "` + string(commons.NotificationCenterResourceName) + `" "%v" {
  provider = "aws"
  policy_id="%v"
  name="Notification-center"
  description="Testing of notification center policy"
  is_active=true
  privacy_level="public"
  registered_users {
	user_email="TestAutomation_Admin_DO_NOT_DELETE@spot.io"
	subscription_types = ["email"]
  }
  subscriptions {
	endpoint="https://webhook.si"
	subscription_type="webhook"
  }
  compute_policy_config {
	events {
	  event="Maximum capacity reached"
	  event_type="WARN"
	}
	dynamic_rules {
		filter_conditions {
			expression="DO-NOT-DELETE"
			identifier="resource_name"
			operator="contains"
		}
	}
  }
}
`

const testNotificationCenter_Update = `
resource "` + string(commons.NotificationCenterResourceName) + `" "%v" {
  provider = "aws"
  policy_id="%v"
  name="Notification-center"
  description="Update Testing of notification center policy"
  is_active=true
  privacy_level="public"
  registered_users {
	user_email="Automation@spotinst.com"
	subscription_types = ["email"]
  }
  subscriptions {
	endpoint="https://webhook.si"
	subscription_type="webhook"
  }
  compute_policy_config {
	events {
	  event="Stateful recycle finished"
	  event_type="INFO"
	}
	should_include_all_resources=true
  }
}
`

// endregion

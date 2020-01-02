package spotinst

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/service/subscription"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

func createSubscriptionResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.SubscriptionResourceName), name)
}

func testSubscriptionDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.SubscriptionResourceName) {
			continue
		}
		input := &subscription.ReadSubscriptionInput{SubscriptionID: spotinst.String(rs.Primary.ID)}
		resp, err := client.subscription.Read(context.Background(), input)
		if err == nil && resp != nil && resp.Subscription != nil {
			return fmt.Errorf("subscription still exists")
		}
	}
	return nil
}

func testCheckSubscriptionExists(sub *subscription.Subscription, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &subscription.ReadSubscriptionInput{SubscriptionID: spotinst.String(rs.Primary.ID)}
		resp, err := client.subscription.Read(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.Subscription.EventType) != rs.Primary.Attributes["event_type"] {
			return fmt.Errorf("Subscription not found: %+v,\n %+v\n", resp.Subscription, rs.Primary.Attributes)
		}
		*sub = *resp.Subscription
		return nil
	}
}

func createSubscriptionTerraform(tfResource string, resourceName string, groupResourceId string, groupTerraform string) string {
	template := ""

	template += fmt.Sprintf(tfResource, resourceName, groupResourceId)
	template = groupTerraform + "\n" + template

	log.Printf("Terraform [%v] template:\n%v", resourceName, template)
	return template
}

// region Subscription: Http
func TestAccSpotinstSubscription_Http(t *testing.T) {
	subscriptionName := "subscription-http"
	subResourceName := createSubscriptionResourceName(subscriptionName)

	groupName := "eg-baseline"
	groupResourceName := createElastigroupResourceName(groupName)
	groupResourceId := "${" + groupResourceName + ".id}"
	groupTerraform := createElastigroupTerraform(&GroupConfigMetadata{groupName: groupName})

	var group aws.Group
	var sub subscription.Subscription
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testSubscriptionDestroy,

		Steps: []resource.TestStep{
			{
				Config: createSubscriptionTerraform(testSubscription_Http_Create, subscriptionName, groupResourceId, groupTerraform),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, groupResourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(groupResourceName, "availability_zones.#", "2"),
					resource.TestCheckResourceAttr(groupResourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(groupResourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(groupResourceName, "desired_capacity", "0"),
					resource.TestCheckResourceAttr(groupResourceName, "capacity_unit", "weight"),

					testCheckSubscriptionExists(&sub, subResourceName),
					resource.TestCheckResourceAttr(subResourceName, "event_type", "AWS_EC2_INSTANCE_LAUNCH"),
					resource.TestCheckResourceAttr(subResourceName, "format.%", "2"),
					resource.TestCheckResourceAttr(subResourceName, "format.customField", "first"),
					resource.TestCheckResourceAttr(subResourceName, "format.anotherCustomField", "second"),
					resource.TestCheckResourceAttr(subResourceName, "protocol", "http"),
					resource.TestCheckResourceAttr(subResourceName, "endpoint", "http://test.me"),
				),
			},
			{
				Config: createSubscriptionTerraform(testSubscription_Http_Update, subscriptionName, groupResourceId, groupTerraform),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, groupResourceName),
					testCheckElastigroupAttributes(&group, groupName),

					testCheckSubscriptionExists(&sub, subResourceName),
					resource.TestCheckResourceAttr(subResourceName, "event_type", "AWS_EC2_INSTANCE_TERMINATE"),
					resource.TestCheckResourceAttr(subResourceName, "format.%", "2"),
					resource.TestCheckResourceAttr(subResourceName, "format.customField", "first updated"),
					resource.TestCheckResourceAttr(subResourceName, "format.anotherCustomField", "second updated"),
					resource.TestCheckResourceAttr(subResourceName, "protocol", "http"),
					resource.TestCheckResourceAttr(subResourceName, "endpoint", "http://test.that"),
				),
			},
		},
	})
}

const testSubscription_Http_Create = `
resource "` + string(commons.SubscriptionResourceName) + `" "%v" {
  provider = "aws"
  resource_id="%v"
  event_type="AWS_EC2_INSTANCE_LAUNCH"

  format = {
		customField        = "first"
		anotherCustomField = "second"
  }

  protocol="http"
  endpoint="http://test.me"
}
`

const testSubscription_Http_Update = `
resource "` + string(commons.SubscriptionResourceName) + `" "%v" {
  provider = "aws"
  resource_id="%v"
  event_type="AWS_EC2_INSTANCE_TERMINATE"

  format = {
		customField        = "first updated"
		anotherCustomField = "second updated"
  }

  protocol="http"
  endpoint="http://test.that"
}
`

// endregion

// region Subscription: Https
func TestAccSpotinstSubscription_Https(t *testing.T) {
	subscriptionName := "subscription-https"
	subResourceName := createSubscriptionResourceName(subscriptionName)

	groupName := "eg-baseline"
	groupResourceName := createElastigroupResourceName(groupName)
	groupResourceId := "${" + groupResourceName + ".id}"
	groupTerraform := createElastigroupTerraform(&GroupConfigMetadata{groupName: groupName})

	var group aws.Group
	var sub subscription.Subscription
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testSubscriptionDestroy,

		Steps: []resource.TestStep{
			{
				Config: createSubscriptionTerraform(testSubscription_Https_Create, subscriptionName, groupResourceId, groupTerraform),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, groupResourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(groupResourceName, "availability_zones.#", "2"),
					resource.TestCheckResourceAttr(groupResourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(groupResourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(groupResourceName, "desired_capacity", "0"),
					resource.TestCheckResourceAttr(groupResourceName, "capacity_unit", "weight"),

					testCheckSubscriptionExists(&sub, subResourceName),
					resource.TestCheckResourceAttr(subResourceName, "event_type", "CANT_SCALE_UP_GROUP_MAX_CAPACITY"),
					resource.TestCheckResourceAttr(subResourceName, "format.%", "2"),
					resource.TestCheckResourceAttr(subResourceName, "format.customField", "first"),
					resource.TestCheckResourceAttr(subResourceName, "format.anotherCustomField", "second"),
					resource.TestCheckResourceAttr(subResourceName, "protocol", "https"),
					resource.TestCheckResourceAttr(subResourceName, "endpoint", "https://test.me"),
				),
			},
			{
				Config: createSubscriptionTerraform(testSubscription_Https_Update, subscriptionName, groupResourceId, groupTerraform),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, groupResourceName),
					testCheckElastigroupAttributes(&group, groupName),

					testCheckSubscriptionExists(&sub, subResourceName),
					resource.TestCheckResourceAttr(subResourceName, "event_type", "GROUP_UPDATED"),
					resource.TestCheckResourceAttr(subResourceName, "format.%", "2"),
					resource.TestCheckResourceAttr(subResourceName, "format.customField", "first updated"),
					resource.TestCheckResourceAttr(subResourceName, "format.anotherCustomField", "second updated"),
					resource.TestCheckResourceAttr(subResourceName, "protocol", "https"),
					resource.TestCheckResourceAttr(subResourceName, "endpoint", "https://test.that"),
				),
			},
		},
	})
}

const testSubscription_Https_Create = `
resource "` + string(commons.SubscriptionResourceName) + `" "%v" {
  provider = "aws"
  resource_id="%v"
  event_type="CANT_SCALE_UP_GROUP_MAX_CAPACITY"

  format = {
		customField        = "first"
		anotherCustomField = "second"
  }

  protocol="https"
  endpoint="https://test.me"
}
`

const testSubscription_Https_Update = `
resource "` + string(commons.SubscriptionResourceName) + `" "%v" {
  provider = "aws"
  resource_id="%v"
  event_type="GROUP_UPDATED"

  format = {
		customField        = "first updated"
		anotherCustomField = "second updated"
  }

  protocol="https"
  endpoint="https://test.that"
}
`

// endregion

// region Subscription: Email
func TestAccSpotinstSubscription_Email(t *testing.T) {
	subscriptionName := "subscription-email"
	subResourceName := createSubscriptionResourceName(subscriptionName)

	groupName := "eg-baseline"
	groupResourceName := createElastigroupResourceName(groupName)
	groupResourceId := "${" + groupResourceName + ".id}"
	groupTerraform := createElastigroupTerraform(&GroupConfigMetadata{groupName: groupName})

	var group aws.Group
	var sub subscription.Subscription
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testSubscriptionDestroy,

		Steps: []resource.TestStep{
			{
				Config: createSubscriptionTerraform(testSubscription_Email_Create, subscriptionName, groupResourceId, groupTerraform),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, groupResourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(groupResourceName, "availability_zones.#", "2"),
					resource.TestCheckResourceAttr(groupResourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(groupResourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(groupResourceName, "desired_capacity", "0"),
					resource.TestCheckResourceAttr(groupResourceName, "capacity_unit", "weight"),

					testCheckSubscriptionExists(&sub, subResourceName),
					resource.TestCheckResourceAttr(subResourceName, "event_type", "GROUP_ROLL_FINISHED"),
					resource.TestCheckResourceAttr(subResourceName, "format.%", "2"),
					resource.TestCheckResourceAttr(subResourceName, "format.customField", "first"),
					resource.TestCheckResourceAttr(subResourceName, "format.anotherCustomField", "second"),
					resource.TestCheckResourceAttr(subResourceName, "protocol", "email"),
					resource.TestCheckResourceAttr(subResourceName, "endpoint", "test@me.com"),
				),
			},
			{
				Config: createSubscriptionTerraform(testSubscription_Email_Update, subscriptionName, groupResourceId, groupTerraform),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, groupResourceName),
					testCheckElastigroupAttributes(&group, groupName),

					testCheckSubscriptionExists(&sub, subResourceName),
					resource.TestCheckResourceAttr(subResourceName, "event_type", "GROUP_ROLL_FAILED"),
					resource.TestCheckResourceAttr(subResourceName, "format.%", "2"),
					resource.TestCheckResourceAttr(subResourceName, "format.customField", "first updated"),
					resource.TestCheckResourceAttr(subResourceName, "format.anotherCustomField", "second updated"),
					resource.TestCheckResourceAttr(subResourceName, "protocol", "email"),
					resource.TestCheckResourceAttr(subResourceName, "endpoint", "test.update@me.com"),
				),
			},
		},
	})
}

const testSubscription_Email_Create = `
resource "` + string(commons.SubscriptionResourceName) + `" "%v" {
  provider = "aws"
  resource_id="%v"
  event_type="GROUP_ROLL_FINISHED"

  format = {
		customField        = "first"
		anotherCustomField = "second"
  }

  protocol="email"
  endpoint="test@me.com"
}
`

const testSubscription_Email_Update = `
resource "` + string(commons.SubscriptionResourceName) + `" "%v" {
  provider = "aws"
  resource_id="%v"
  event_type="GROUP_ROLL_FAILED"

  format = {
		customField        = "first updated"
		anotherCustomField = "second updated"
  }

  protocol="email"
  endpoint="test.update@me.com"
}
`

// endregion

// region Subscription: Email Json
func TestAccSpotinstSubscription_EmailJson(t *testing.T) {
	subscriptionName := "subscription-email-json"
	subResourceName := createSubscriptionResourceName(subscriptionName)

	groupName := "eg-baseline"
	groupResourceName := createElastigroupResourceName(groupName)
	groupResourceId := "${" + groupResourceName + ".id}"
	groupTerraform := createElastigroupTerraform(&GroupConfigMetadata{groupName: groupName})

	var group aws.Group
	var sub subscription.Subscription
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testSubscriptionDestroy,

		Steps: []resource.TestStep{
			{
				Config: createSubscriptionTerraform(testSubscription_EmailJson_Create, subscriptionName, groupResourceId, groupTerraform),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, groupResourceName),
					testCheckElastigroupAttributes(&group, groupName),
					resource.TestCheckResourceAttr(groupResourceName, "availability_zones.#", "2"),
					resource.TestCheckResourceAttr(groupResourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(groupResourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(groupResourceName, "desired_capacity", "0"),
					resource.TestCheckResourceAttr(groupResourceName, "capacity_unit", "weight"),

					testCheckSubscriptionExists(&sub, subResourceName),
					resource.TestCheckResourceAttr(subResourceName, "event_type", "AWS_EC2_INSTANCE_UNHEALTHY_IN_ELB"),
					resource.TestCheckResourceAttr(subResourceName, "format.%", "2"),
					resource.TestCheckResourceAttr(subResourceName, "format.customField", "first"),
					resource.TestCheckResourceAttr(subResourceName, "format.anotherCustomField", "second"),
					resource.TestCheckResourceAttr(subResourceName, "protocol", "email-json"),
					resource.TestCheckResourceAttr(subResourceName, "endpoint", "test@me.com"),
				),
			},
			{
				Config: createSubscriptionTerraform(testSubscription_EmailJson_Update, subscriptionName, groupResourceId, groupTerraform),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupExists(&group, groupResourceName),
					testCheckElastigroupAttributes(&group, groupName),

					testCheckSubscriptionExists(&sub, subResourceName),
					resource.TestCheckResourceAttr(subResourceName, "event_type", "AWS_EC2_INSTANCE_TERMINATED"),
					resource.TestCheckResourceAttr(subResourceName, "format.%", "2"),
					resource.TestCheckResourceAttr(subResourceName, "format.customField", "first updated"),
					resource.TestCheckResourceAttr(subResourceName, "format.anotherCustomField", "second updated"),
					resource.TestCheckResourceAttr(subResourceName, "protocol", "email-json"),
					resource.TestCheckResourceAttr(subResourceName, "endpoint", "test.update@me.com"),
				),
			},
		},
	})
}

const testSubscription_EmailJson_Create = `
resource "` + string(commons.SubscriptionResourceName) + `" "%v" {
  provider = "aws"
  resource_id="%v"
  event_type="AWS_EC2_INSTANCE_UNHEALTHY_IN_ELB"

  format = {
		customField        = "first"
		anotherCustomField = "second"
  }

  protocol="email-json"
  endpoint="test@me.com"
}
`

const testSubscription_EmailJson_Update = `
resource "` + string(commons.SubscriptionResourceName) + `" "%v" {
  provider = "aws"
  resource_id="%v"
  event_type="AWS_EC2_INSTANCE_TERMINATED"

  format = {
		customField        = "first updated"
		anotherCustomField = "second updated"
  }

  protocol="email-json"
  endpoint="test.update@me.com"
}
`

// endregion

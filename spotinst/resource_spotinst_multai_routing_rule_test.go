package spotinst

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

func TestAccSpotinstMultaiRoutingRule_Basic(t *testing.T) {
	var rule spotinst.RoutingRule
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSpotinstMultaiRoutingRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSpotinstMultaiRoutingRuleConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiRoutingRuleExists("spotinst_multai_routing_rule.foo", &rule),
					testAccCheckSpotinstMultaiRoutingRuleAttributes(&rule),
					resource.TestCheckResourceAttr("spotinst_multai_routing_rule.foo", "route", "Path(`/foo`)"),
				),
			},
		},
	})
}

func TestAccSpotinstMultaiRoutingRule_Updated(t *testing.T) {
	var rule spotinst.RoutingRule
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSpotinstMultaiRoutingRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSpotinstMultaiRoutingRuleConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiRoutingRuleExists("spotinst_multai_routing_rule.foo", &rule),
					testAccCheckSpotinstMultaiRoutingRuleAttributes(&rule),
					resource.TestCheckResourceAttr("spotinst_multai_routing_rule.foo", "route", "Path(`/foo`)"),
				),
			},
			{
				Config: testAccCheckSpotinstMultaiRoutingRuleConfigNewValue,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiRoutingRuleExists("spotinst_multai_routing_rule.foo", &rule),
					testAccCheckSpotinstMultaiRoutingRuleAttributesUpdated(&rule),
					resource.TestCheckResourceAttr("spotinst_multai_routing_rule.foo", "route", "Path(`/bar`)"),
				),
			},
		},
	})
}

func testAccCheckSpotinstMultaiRoutingRuleDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*spotinst.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "spotinst_multai_routing_rule" {
			continue
		}
		input := &spotinst.ReadRoutingRuleInput{
			RoutingRuleID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.MultaiService.BalancerService().ReadRoutingRule(context.Background(), input)
		if err == nil && resp != nil && resp.RoutingRule != nil {
			return fmt.Errorf("RoutingRule still exists")
		}
	}
	return nil
}

func testAccCheckSpotinstMultaiRoutingRuleAttributes(rule *spotinst.RoutingRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if p := spotinst.StringValue(rule.Route); p != "Path(`/foo`)" {
			return fmt.Errorf("Bad content: %s", p)
		}
		return nil
	}
}

func testAccCheckSpotinstMultaiRoutingRuleAttributesUpdated(rule *spotinst.RoutingRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if p := spotinst.StringValue(rule.Route); p != "Path(`/bar`)" {
			return fmt.Errorf("Bad content: %s", p)
		}
		return nil
	}
}

func testAccCheckSpotinstMultaiRoutingRuleExists(n string, rule *spotinst.RoutingRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No resource ID is rule")
		}
		client := testAccProvider.Meta().(*spotinst.Client)
		input := &spotinst.ReadRoutingRuleInput{
			RoutingRuleID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.MultaiService.BalancerService().ReadRoutingRule(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.RoutingRule.ID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("RoutingRule not found: %+v,\n %+v\n", resp.RoutingRule, rs.Primary.Attributes)
		}
		*rule = *resp.RoutingRule
		return nil
	}
}

const testAccCheckSpotinstMultaiRoutingRuleConfigBasic = `
resource "spotinst_multai_balancer" "foo" {
  name = "foo"

  connection_timeouts {
    idle     = 10
    draining = 10
  }

  tags {
    env = "prod"
    app = "web"
  }
}

resource "spotinst_multai_listener" "foo" {
  balancer_id = "${spotinst_multai_balancer.foo.id}"
  protocol    = "http"
  port        = 1338

  tags {
    env = "prod"
    app = "web"
  }
}

resource "spotinst_multai_routing_rule" "foo" {
  balancer_id    = "${spotinst_multai_balancer.foo.id}"
  listener_id    = "${spotinst_multai_listener.foo.id}"
  route          = "Path(\x60/foo\x60)"
  middleware_ids = []
  target_set_ids = []

  tags {
    env = "prod"
    app = "web"
  }
}`

const testAccCheckSpotinstMultaiRoutingRuleConfigNewValue = `
resource "spotinst_multai_balancer" "foo" {
  name = "foo"

  connection_timeouts {
    idle     = 10
    draining = 10
  }

  tags {
    env = "prod"
    app = "web"
  }
}

resource "spotinst_multai_listener" "foo" {
  balancer_id = "${spotinst_multai_balancer.foo.id}"
  protocol    = "http"
  port        = 1338

  tags {
    env = "prod"
    app = "web"
  }
}

resource "spotinst_multai_routing_rule" "foo" {
  balancer_id    = "${spotinst_multai_balancer.foo.id}"
  listener_id    = "${spotinst_multai_listener.foo.id}"
  route          = "Path(\x60/bar\x60)"
  middleware_ids = []
  target_set_ids = []

  tags {
    env = "prod"
    app = "web"
  }
}`

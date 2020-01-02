package spotinst

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

func createMultaiRoutingRuleResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.MultaiRoutingRuleResourceName), name)
}

func testAccCheckSpotinstMultaiRoutingRuleDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "spotinst_multai_routing_rule" {
			continue
		}
		input := &multai.ReadRoutingRuleInput{
			RoutingRuleID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.multai.ReadRoutingRule(context.Background(), input)
		if err == nil && resp != nil && resp.RoutingRule != nil {
			return fmt.Errorf("routing rule still exists")
		}
	}
	return nil
}

func testAccCheckSpotinstMultaiRoutingRuleExists(routingRule *multai.RoutingRule, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &multai.ReadRoutingRuleInput{
			RoutingRuleID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.multai.ReadRoutingRule(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.RoutingRule.ID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("routing rule not found: %+v,\n %+v\n", resp.RoutingRule, rs.Primary.Attributes)
		}
		*routingRule = *resp.RoutingRule
		return nil
	}
}

type RoutingRuleConfigMetadata struct {
	provider             string
	name                 string
	updateBaselineFields bool
}

func createRoutingRuleTerraform(lcm *RoutingRuleConfigMetadata) string {
	if lcm == nil {
		return ""
	}

	if lcm.provider == "" {
		lcm.provider = "aws"
	}

	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`

	if lcm.updateBaselineFields {
		format := testBaselineRoutingRuleConfig_Update
		template += fmt.Sprintf(format,
			lcm.name,
			lcm.provider,
		)
	} else {
		format := testBaselineRoutingRuleConfig_Create

		template += fmt.Sprintf(format,
			lcm.name,
			lcm.provider,
		)
	}

	log.Printf("Terraform [%v] template:\n%v", lcm.name, template)
	return template
}

func TestAccSpotinstMultaiRoutingRule_Baseline(t *testing.T) {
	routingRuleName := "routing-rule-baseline"
	resourceName := createMultaiRoutingRuleResourceName(routingRuleName)

	var routingRule multai.RoutingRule
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckSpotinstMultaiRoutingRuleDestroy,

		Steps: []resource.TestStep{
			{
				Config: createRoutingRuleTerraform(&RoutingRuleConfigMetadata{
					name: routingRuleName,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiRoutingRuleExists(&routingRule, resourceName),
					resource.TestCheckResourceAttr(resourceName, "middleware_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route", "Path(`/bar`)"),
					resource.TestCheckResourceAttr(resourceName, "strategy", "RANDOM"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags."+RoutingRuleTagsHash_Create+".key", "fakeKey"),
					resource.TestCheckResourceAttr(resourceName, "tags."+RoutingRuleTagsHash_Create+".value", "fakeVal"),
				),
			},
			{
				Config: createRoutingRuleTerraform(&RoutingRuleConfigMetadata{
					name:                 routingRuleName,
					updateBaselineFields: true,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiRoutingRuleExists(&routingRule, resourceName),
					resource.TestCheckResourceAttr(resourceName, "middleware_ids.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route", "Path(`/baz`)"),
					resource.TestCheckResourceAttr(resourceName, "strategy", "LEASTCONN"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags."+RoutingRuleTagsHash_Update+".key", "updated"),
					resource.TestCheckResourceAttr(resourceName, "tags."+RoutingRuleTagsHash_Update+".value", "updated"),
				),
			},
		},
	})
}

const (
	RoutingRuleTagsHash_Create = "2538041064"
	RoutingRuleTagsHash_Update = "1968254376"

	Path_Create = "\"Path(\x60/bar\x60)\""
	Path_Update = "\"Path(\x60/baz\x60)\""
)

const testBaselineRoutingRuleConfig_Create = `
resource "spotinst_multai_balancer" "foo" {
  provider = "aws"
  name = "test-acc-foo"

  connection_timeouts {
    idle     = 10
    draining = 10
  }
}

resource "spotinst_multai_target_set" "foo" {
  provider = "aws"
  balancer_id   = "${spotinst_multai_balancer.foo.id}"
  deployment_id = "dp-12345"
  name          = "test-acc-bar"
  protocol      = "http"
  port          = 1338
  weight        = 2

  health_check {
    protocol            = "http"
    path                = "/"
    port                = 3001
    interval            = 20
    timeout             = 5
    healthy_threshold   = 3
    unhealthy_threshold = 3
  }

  tags {
   key = "updated"
   value = "updated"
  }
}

resource "spotinst_multai_listener" "foo" {
  provider = "aws"
  balancer_id = "${spotinst_multai_balancer.foo.id}"
  protocol    = "http"
  port        = 1338

  tags {
    key = "prod"
    value = "web"
  }
}

resource "` + string(commons.MultaiRoutingRuleResourceName) + `" "%v" {
  provider = "%v"
  balancer_id = "${spotinst_multai_balancer.foo.id}"
  listener_id    = "${spotinst_multai_listener.foo.id}"
  route          = ` + Path_Create + `
  strategy       = "RANDOM"
  //middleware_ids = ["example"]
  target_set_ids = ["${spotinst_multai_target_set.foo.id}"]

  tags {
   key = "fakeKey"
   value = "fakeVal"
  }
}`

const testBaselineRoutingRuleConfig_Update = `
resource "spotinst_multai_balancer" "foo" {
  provider = "aws"
  name = "test-acc-foo"

  connection_timeouts {
    idle     = 10
    draining = 10
  }

  tags {
   key = "prod"
   value = "web"
  }
}

resource "spotinst_multai_target_set" "foo" {
  provider = "aws"
  balancer_id   = "${spotinst_multai_balancer.foo.id}"
  deployment_id = "dp-12345"
  name          = "test-acc-bar"
  protocol      = "http"
  port          = 1338
  weight        = 2

  health_check {
    protocol            = "http"
    path                = "/"
    port                = 3001
    interval            = 20
    timeout             = 5
    healthy_threshold   = 3
    unhealthy_threshold = 3
  }

  tags {
   key = "updated"
   value = "updated"
  }
}

resource "spotinst_multai_listener" "foo" {
  provider = "aws"
  balancer_id = "${spotinst_multai_balancer.foo.id}"
  protocol    = "http"
  port        = 1338

  tags {
    key = "prod"
    value = "web"
  }
}

resource "` + string(commons.MultaiRoutingRuleResourceName) + `" "%v" {
  provider = "%v"
  balancer_id = "${spotinst_multai_balancer.foo.id}"
  listener_id    = "${spotinst_multai_listener.foo.id}"
  route          = ` + Path_Update + `
  strategy       = "LEASTCONN"
  target_set_ids = ["${spotinst_multai_target_set.foo.id}"]

  tags {
   key = "updated"
   value = "updated"
  }
}`

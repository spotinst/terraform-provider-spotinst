package spotinst

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

func TestAccSpotinstMultaiBalancer_Basic(t *testing.T) {
	var balancer spotinst.Balancer
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSpotinstMultaiBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSpotinstMultaiBalancerConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiBalancerExists("spotinst_multai_balancer.foo", &balancer),
					testAccCheckSpotinstMultaiBalancerAttributes(&balancer),
					resource.TestCheckResourceAttr("spotinst_multai_balancer.foo", "name", "foo"),
				),
			},
		},
	})
}

func TestAccSpotinstMultaiBalancer_Updated(t *testing.T) {
	var balancer spotinst.Balancer
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSpotinstMultaiBalancerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSpotinstMultaiBalancerConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiBalancerExists("spotinst_multai_balancer.foo", &balancer),
					testAccCheckSpotinstMultaiBalancerAttributes(&balancer),
					resource.TestCheckResourceAttr("spotinst_multai_balancer.foo", "name", "foo"),
				),
			},
			{
				Config: testAccCheckSpotinstMultaiBalancerConfigNewValue,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiBalancerExists("spotinst_multai_balancer.foo", &balancer),
					testAccCheckSpotinstMultaiBalancerAttributesUpdated(&balancer),
					resource.TestCheckResourceAttr("spotinst_multai_balancer.foo", "name", "foo"),
				),
			},
		},
	})
}

func testAccCheckSpotinstMultaiBalancerDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*spotinst.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "spotinst_multai_balancer" {
			continue
		}
		input := &spotinst.ReadBalancerInput{
			BalancerID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.MultaiService.BalancerService().ReadBalancer(context.Background(), input)
		if err == nil && resp != nil && resp.Balancer != nil {
			return fmt.Errorf("Balancer still exists")
		}
	}
	return nil
}

func testAccCheckSpotinstMultaiBalancerAttributes(balancer *spotinst.Balancer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if p := spotinst.StringValue(balancer.Name); p != "foo" {
			return fmt.Errorf("Bad content: %s", p)
		}
		return nil
	}
}

func testAccCheckSpotinstMultaiBalancerAttributesUpdated(balancer *spotinst.Balancer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if p := spotinst.StringValue(balancer.Name); p != "foo" {
			return fmt.Errorf("Bad content: %s", p)
		}
		return nil
	}
}

func testAccCheckSpotinstMultaiBalancerExists(n string, balancer *spotinst.Balancer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No resource ID is set")
		}
		client := testAccProvider.Meta().(*spotinst.Client)
		input := &spotinst.ReadBalancerInput{
			BalancerID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.MultaiService.BalancerService().ReadBalancer(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.Balancer.ID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("Balancer not found: %+v,\n %+v\n", resp.Balancer, rs.Primary.Attributes)
		}
		*balancer = *resp.Balancer
		return nil
	}
}

const testAccCheckSpotinstMultaiBalancerConfigBasic = `
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
}`

const testAccCheckSpotinstMultaiBalancerConfigNewValue = `
resource "spotinst_multai_balancer" "foo" {
  name = "foo"

  connection_timeouts {
    idle     = 20
    draining = 20
  }

  tags {
    env = "prod"
    app = "web"
  }
}`

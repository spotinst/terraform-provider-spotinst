package spotinst
//
//import (
//	"context"
//	"fmt"
//	"testing"
//
//	"github.com/hashicorp/terraform/helper/resource"
//	"github.com/hashicorp/terraform/terraform"
//	"github.com/spotinst/spotinst-sdk-go/service/multai"
//	"github.com/spotinst/spotinst-sdk-go/spotinst"
//)
//
//func TestAccSpotinstMultaiTarget_Basic(t *testing.T) {
//	var target multai.Target
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckSpotinstMultaiTargetDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccCheckSpotinstMultaiTargetConfigBasic,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSpotinstMultaiTargetExists("spotinst_multai_target.foo", &target),
//					testAccCheckSpotinstMultaiTargetAttributes(&target),
//					resource.TestCheckResourceAttr("spotinst_multai_target.foo", "weight", "1"),
//				),
//			},
//		},
//	})
//}
//
//func TestAccSpotinstMultaiTarget_Updated(t *testing.T) {
//	var target multai.Target
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckSpotinstMultaiTargetDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccCheckSpotinstMultaiTargetConfigBasic,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSpotinstMultaiTargetExists("spotinst_multai_target.foo", &target),
//					testAccCheckSpotinstMultaiTargetAttributes(&target),
//					resource.TestCheckResourceAttr("spotinst_multai_target.foo", "weight", "1"),
//				),
//			},
//			{
//				Config: testAccCheckSpotinstMultaiTargetConfigNewValue,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSpotinstMultaiTargetExists("spotinst_multai_target.foo", &target),
//					testAccCheckSpotinstMultaiTargetAttributesUpdated(&target),
//					resource.TestCheckResourceAttr("spotinst_multai_target.foo", "weight", "2"),
//				),
//			},
//		},
//	})
//}
//
//func testAccCheckSpotinstMultaiTargetDestroy(s *terraform.State) error {
//	client := testAccProvider.Meta().(*Client)
//	for _, rs := range s.RootModule().Resources {
//		if rs.Type != "spotinst_multai_target" {
//			continue
//		}
//		input := &multai.ReadTargetInput{
//			TargetID: spotinst.String(rs.Primary.ID),
//		}
//		resp, err := client.multai.ReadTarget(context.Background(), input)
//		if err == nil && resp != nil && resp.Target != nil {
//			return fmt.Errorf("target still exists")
//		}
//	}
//	return nil
//}
//
//func testAccCheckSpotinstMultaiTargetAttributes(target *multai.Target) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		if p := spotinst.IntValue(target.Weight); p != 1 {
//			return fmt.Errorf("bad content: %d", p)
//		}
//		return nil
//	}
//}
//
//func testAccCheckSpotinstMultaiTargetAttributesUpdated(target *multai.Target) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		if p := spotinst.IntValue(target.Weight); p != 2 {
//			return fmt.Errorf("bad content: %d", p)
//		}
//		return nil
//	}
//}
//
//func testAccCheckSpotinstMultaiTargetExists(n string, target *multai.Target) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		rs, ok := s.RootModule().Resources[n]
//		if !ok {
//			return fmt.Errorf("not found: %s", n)
//		}
//		if rs.Primary.ID == "" {
//			return fmt.Errorf("no resource ID is target")
//		}
//		client := testAccProvider.Meta().(*Client)
//		input := &multai.ReadTargetInput{
//			TargetID: spotinst.String(rs.Primary.ID),
//		}
//		resp, err := client.multai.ReadTarget(context.Background(), input)
//		if err != nil {
//			return err
//		}
//		if spotinst.StringValue(resp.Target.ID) != rs.Primary.Attributes["id"] {
//			return fmt.Errorf("target not found: %+v,\n %+v\n", resp.Target, rs.Primary.Attributes)
//		}
//		*target = *resp.Target
//		return nil
//	}
//}
//
//const testAccCheckSpotinstMultaiTargetConfigBasic = `
//resource "spotinst_multai_balancer" "foo" {
//  name = "foo"
//
//  connection_timeouts {
//    idle     = 10
//    draining = 10
//  }
//
//  tags {
//    env = "prod"
//    app = "web"
//  }
//}
//
//resource "spotinst_multai_target_set" "foo" {
//  balancer_id   = "${spotinst_multai_balancer.foo.id}"
//  deployment_id = "dp-12345"
//  name          = "bar"
//  protocol      = "http"
//  port          = 1337
//  weight        = 1
//
//  health_check {
//    protocol            = "http"
//    path                = "/"
//    interval            = 30
//    timeout             = 10
//    healthy_threshold   = 2
//    unhealthy_threshold = 2
//  }
//
//  tags {
//    env = "prod"
//    app = "web"
//  }
//}
//
//resource "spotinst_multai_target" "foo" {
//  balancer_id   = "${spotinst_multai_balancer.foo.id}"
//  target_set_id = "${spotinst_multai_target_set.foo.id}"
//  host          = "172.0.0.10"
//  port          = 1337
//  weight        = 1
//
//  tags {
//    env = "prod"
//    app = "web"
//  }
//}`
//
//const testAccCheckSpotinstMultaiTargetConfigNewValue = `
//resource "spotinst_multai_balancer" "foo" {
//  name = "foo"
//
//  connection_timeouts {
//    idle     = 10
//    draining = 10
//  }
//
//  tags {
//    env = "prod"
//    app = "web"
//  }
//}
//
//resource "spotinst_multai_target_set" "foo" {
//  balancer_id   = "${spotinst_multai_balancer.foo.id}"
//  deployment_id = "dp-12345"
//  name          = "bar"
//  protocol      = "http"
//  port          = 1337
//  weight        = 1
//
//  health_check {
//    protocol            = "http"
//    path                = "/"
//    interval            = 30
//    timeout             = 10
//    healthy_threshold   = 2
//    unhealthy_threshold = 2
//  }
//
//  tags {
//    env = "prod"
//    app = "web"
//  }
//}
//
//resource "spotinst_multai_target" "foo" {
//  balancer_id   = "${spotinst_multai_balancer.foo.id}"
//  target_set_id = "${spotinst_multai_target_set.foo.id}"
//  host          = "172.0.0.10"
//  port          = 1337
//  weight        = 2
//
//  tags {
//    env = "prod"
//    app = "web"
//  }
//}`

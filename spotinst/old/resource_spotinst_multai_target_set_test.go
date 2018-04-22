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
//func TestAccSpotinstMultaiTargetSet_Basic(t *testing.T) {
//	var set multai.TargetSet
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckSpotinstMultaiTargetSetDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccCheckSpotinstMultaiTargetSetConfigBasic,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSpotinstMultaiTargetSetExists("spotinst_multai_target_set.foo", &set),
//					testAccCheckSpotinstMultaiTargetSetAttributes(&set),
//					resource.TestCheckResourceAttr("spotinst_multai_target_set.foo", "name", "foo"),
//				),
//			},
//		},
//	})
//}
//
//func TestAccSpotinstMultaiTargetSet_Updated(t *testing.T) {
//	var set multai.TargetSet
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckSpotinstMultaiTargetSetDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccCheckSpotinstMultaiTargetSetConfigBasic,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSpotinstMultaiTargetSetExists("spotinst_multai_target_set.foo", &set),
//					testAccCheckSpotinstMultaiTargetSetAttributes(&set),
//					resource.TestCheckResourceAttr("spotinst_multai_target_set.foo", "name", "foo"),
//				),
//			},
//			{
//				Config: testAccCheckSpotinstMultaiTargetSetConfigNewValue,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSpotinstMultaiTargetSetExists("spotinst_multai_target_set.foo", &set),
//					testAccCheckSpotinstMultaiTargetSetAttributesUpdated(&set),
//					resource.TestCheckResourceAttr("spotinst_multai_target_set.foo", "name", "bar"),
//				),
//			},
//		},
//	})
//}
//
//func testAccCheckSpotinstMultaiTargetSetDestroy(s *terraform.State) error {
//	client := testAccProvider.Meta().(*Client)
//	for _, rs := range s.RootModule().Resources {
//		if rs.Type != "spotinst_multai_target_set" {
//			continue
//		}
//		input := &multai.ReadTargetSetInput{
//			TargetSetID: spotinst.String(rs.Primary.ID),
//		}
//		resp, err := client.multai.ReadTargetSet(context.Background(), input)
//		if err == nil && resp != nil && resp.TargetSet != nil {
//			return fmt.Errorf("target set still exists")
//		}
//	}
//	return nil
//}
//
//func testAccCheckSpotinstMultaiTargetSetAttributes(set *multai.TargetSet) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		if p := spotinst.StringValue(set.Name); p != "foo" {
//			return fmt.Errorf("bad content: %s", p)
//		}
//		return nil
//	}
//}
//
//func testAccCheckSpotinstMultaiTargetSetAttributesUpdated(set *multai.TargetSet) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		if p := spotinst.StringValue(set.Name); p != "bar" {
//			return fmt.Errorf("bad content: %s", p)
//		}
//		return nil
//	}
//}
//
//func testAccCheckSpotinstMultaiTargetSetExists(n string, set *multai.TargetSet) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		rs, ok := s.RootModule().Resources[n]
//		if !ok {
//			return fmt.Errorf("not found: %s", n)
//		}
//		if rs.Primary.ID == "" {
//			return fmt.Errorf("no resource ID is set")
//		}
//		client := testAccProvider.Meta().(*Client)
//		input := &multai.ReadTargetSetInput{
//			TargetSetID: spotinst.String(rs.Primary.ID),
//		}
//		resp, err := client.multai.ReadTargetSet(context.Background(), input)
//		if err != nil {
//			return err
//		}
//		if spotinst.StringValue(resp.TargetSet.ID) != rs.Primary.Attributes["id"] {
//			return fmt.Errorf("target set not found: %+v,\n %+v\n", resp.TargetSet, rs.Primary.Attributes)
//		}
//		*set = *resp.TargetSet
//		return nil
//	}
//}
//
//const testAccCheckSpotinstMultaiTargetSetConfigBasic = `
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
//  name          = "foo"
//  protocol      = "http"
//  port          = 1337
//  weight        = 1
//
//  health_check {
//    protocol            = "http"
//    path                = "/"
//    interval            = "30"
//    timeout             = 10
//    healthy_threshold   = 2
//    unhealthy_threshold = 2
//  }
//
//  tags {
//    env = "prod"
//    app = "web"
//  }
//}`
//
//const testAccCheckSpotinstMultaiTargetSetConfigNewValue = `
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
//}`

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
//func TestAccSpotinstMultaiListener_Basic(t *testing.T) {
//	var listener multai.Listener
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckSpotinstMultaiListenerDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccCheckSpotinstMultaiListenerConfigBasic,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSpotinstMultaiListenerExists("spotinst_multai_listener.foo", &listener),
//					testAccCheckSpotinstMultaiListenerAttributes(&listener),
//					resource.TestCheckResourceAttr("spotinst_multai_listener.foo", "port", "1337"),
//				),
//			},
//		},
//	})
//}
//
//func TestAccSpotinstMultaiListener_Updated(t *testing.T) {
//	var listener multai.Listener
//	resource.Test(t, resource.TestCase{
//		PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckSpotinstMultaiListenerDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccCheckSpotinstMultaiListenerConfigBasic,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSpotinstMultaiListenerExists("spotinst_multai_listener.foo", &listener),
//					testAccCheckSpotinstMultaiListenerAttributes(&listener),
//					resource.TestCheckResourceAttr("spotinst_multai_listener.foo", "port", "1337"),
//				),
//			},
//			{
//				Config: testAccCheckSpotinstMultaiListenerConfigNewValue,
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckSpotinstMultaiListenerExists("spotinst_multai_listener.foo", &listener),
//					testAccCheckSpotinstMultaiListenerAttributesUpdated(&listener),
//					resource.TestCheckResourceAttr("spotinst_multai_listener.foo", "port", "1338"),
//				),
//			},
//		},
//	})
//}
//
//func testAccCheckSpotinstMultaiListenerDestroy(s *terraform.State) error {
//	client := testAccProvider.Meta().(*Client)
//	for _, rs := range s.RootModule().Resources {
//		if rs.Type != "spotinst_multai_listener" {
//			continue
//		}
//		input := &multai.ReadListenerInput{
//			ListenerID: spotinst.String(rs.Primary.ID),
//		}
//		resp, err := client.multai.ReadListener(context.Background(), input)
//		if err == nil && resp != nil && resp.Listener != nil {
//			return fmt.Errorf("listener still exists")
//		}
//	}
//	return nil
//}
//
//func testAccCheckSpotinstMultaiListenerAttributes(listener *multai.Listener) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		if p := spotinst.IntValue(listener.Port); p != 1337 {
//			return fmt.Errorf("bad content: %d", p)
//		}
//		return nil
//	}
//}
//
//func testAccCheckSpotinstMultaiListenerAttributesUpdated(listener *multai.Listener) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		if p := spotinst.IntValue(listener.Port); p != 1338 {
//			return fmt.Errorf("bad content: %d", p)
//		}
//		return nil
//	}
//}
//
//func testAccCheckSpotinstMultaiListenerExists(n string, listener *multai.Listener) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		rs, ok := s.RootModule().Resources[n]
//		if !ok {
//			return fmt.Errorf("not found: %s", n)
//		}
//		if rs.Primary.ID == "" {
//			return fmt.Errorf("no resource ID is listener")
//		}
//		client := testAccProvider.Meta().(*Client)
//		input := &multai.ReadListenerInput{
//			ListenerID: spotinst.String(rs.Primary.ID),
//		}
//		resp, err := client.multai.ReadListener(context.Background(), input)
//		if err != nil {
//			return err
//		}
//		if spotinst.StringValue(resp.Listener.ID) != rs.Primary.Attributes["id"] {
//			return fmt.Errorf("listener not found: %+v,\n %+v\n", resp.Listener, rs.Primary.Attributes)
//		}
//		*listener = *resp.Listener
//		return nil
//	}
//}
//
//const testAccCheckSpotinstMultaiListenerConfigBasic = `
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
//resource "spotinst_multai_listener" "foo" {
//  balancer_id = "${spotinst_multai_balancer.foo.id}"
//  protocol    = "http"
//  port        = 1337
//
//  tags {
//    env = "prod"
//    app = "web"
//  }
//}`
//
//const testAccCheckSpotinstMultaiListenerConfigNewValue = `
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
//resource "spotinst_multai_listener" "foo" {
//  balancer_id = "${spotinst_multai_balancer.foo.id}"
//  protocol    = "http"
//  port        = 1338
//
//  tags {
//    env = "prod"
//    app = "web"
//  }
//}`

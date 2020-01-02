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

func createMultaiListenerResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.MultaiListenerResourceName), name)
}

func testAccCheckSpotinstMultaiListenerDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "spotinst_multai_listener" {
			continue
		}
		input := &multai.ReadListenerInput{
			ListenerID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.multai.ReadListener(context.Background(), input)
		if err == nil && resp != nil && resp.Listener != nil {
			return fmt.Errorf("listener still exists")
		}
	}
	return nil
}

func testAccCheckSpotinstMultaiListenerExists(listener *multai.Listener, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &multai.ReadListenerInput{
			ListenerID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.multai.ReadListener(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.Listener.ID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("listener not found: %+v,\n %+v\n", resp.Listener, rs.Primary.Attributes)
		}
		*listener = *resp.Listener
		return nil
	}
}

type ListenerConfigMetadata struct {
	provider             string
	name                 string
	updateBaselineFields bool
}

func createListenerTerraform(lcm *ListenerConfigMetadata) string {
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
		format := testBaselineListenerConfig_Update
		template += fmt.Sprintf(format,
			lcm.name,
			lcm.provider,
			//lcm.name,
		)
	} else {
		format := testBaselineListenerConfig_Create

		template += fmt.Sprintf(format,
			lcm.name,
			lcm.provider,
			//lcm.name,
		)
	}

	log.Printf("Terraform [%v] template:\n%v", lcm.name, template)
	return template
}

func TestAccSpotinstMultaiListener_Baseline(t *testing.T) {
	listenerName := "test-acc-listener-baseline"
	resourceName := createMultaiListenerResourceName(listenerName)

	var listener multai.Listener
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckSpotinstMultaiListenerDestroy,

		Steps: []resource.TestStep{
			{
				Config: createListenerTerraform(&ListenerConfigMetadata{
					name: listenerName,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiListenerExists(&listener, resourceName),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "port", "1337"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags."+ListenerTagsHash_Create+".key", "fakeKey"),
					resource.TestCheckResourceAttr(resourceName, "tags."+ListenerTagsHash_Create+".value", "fakeVal"),
				),
			},
			{
				Config: createListenerTerraform(&ListenerConfigMetadata{
					name:                 listenerName,
					updateBaselineFields: true,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiListenerExists(&listener, resourceName),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "port", "1338"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags."+ListenerTagsHash_Update+".key", "updated"),
					resource.TestCheckResourceAttr(resourceName, "tags."+ListenerTagsHash_Update+".value", "updated"),
				),
			},
		},
	})
}

const (
	ListenerTagsHash_Create = "2538041064"
	ListenerTagsHash_Update = "1968254376"
)

const testBaselineListenerConfig_Create = `
resource "spotinst_multai_balancer" "foo" {
  provider = "aws"
  name = "foo"

  connection_timeouts {
    idle     = 10
    draining = 10
  }

  tags {
    key = "prod"
    value = "web"
  }
}

resource "` + string(commons.MultaiListenerResourceName) + `" "%v" {
  provider = "%v"
  balancer_id = "${spotinst_multai_balancer.foo.id}"
  protocol    = "http"
  port        = 1337

  tags {
   key = "fakeKey"
   value = "fakeVal"
  }
}`

const testBaselineListenerConfig_Update = `
resource "spotinst_multai_balancer" "foo" {
  provider = "aws"
  name = "foo"

  connection_timeouts {
    idle     = 10
    draining = 10
  }

  tags {
   key = "prod"
   value = "web"
  }
}

resource "` + string(commons.MultaiListenerResourceName) + `" "%v" {
  provider = "%v"
  balancer_id = "${spotinst_multai_balancer.foo.id}"
  protocol    = "http"
  port        = 1338

  //tls_config {
  //  certificate_ids             = ["ce-b7159e06c63d"]
  //  min_version                 = "TLS10"
  //  max_version                 = "TLS12"
  //  cipher_suites               = [""]
  //  prefer_server_cipher_suites = true
  //  session_tickets_disabled    = false
  //}

  tags {
   key = "updated"
   value = "updated"
  }
}`

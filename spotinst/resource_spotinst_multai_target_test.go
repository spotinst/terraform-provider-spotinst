package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

func init() {
	resource.AddTestSweepers("spotinst_multai_target", &resource.Sweeper{
		Name: "spotinst_multai_target",
		F:    testSweepMultaiTarget,
	})
}

func testSweepMultaiTarget(region string) error {
	client, err := getProviderClient("aws")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).multai
	input := &multai.ListTargetsInput{}
	if resp, err := conn.ListTargets(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of targets to sweep")
	} else {
		if len(resp.Targets) == 0 {
			log.Printf("[INFO] No target to sweep")
		}
		for _, tgt := range resp.Targets {
			if strings.Contains(spotinst.StringValue(tgt.Name), "test-acc-") {
				if _, err := conn.DeleteTarget(context.Background(), &multai.DeleteTargetInput{TargetSetID: tgt.ID}); err != nil {
					return fmt.Errorf("unable to delete target %v in sweep", spotinst.StringValue(tgt.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(tgt.ID))
				}
			}
		}
	}
	return nil
}

func createMultaiTargetResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.MultaiTargetResourceName), name)
}

func testAccCheckSpotinstMultaiTargetDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.MultaiTargetResourceName) {
			continue
		}
		input := &multai.ReadTargetInput{TargetID: spotinst.String(rs.Primary.ID)}

		resp, err := client.multai.ReadTarget(context.Background(), input)
		if err == nil && resp != nil && resp.Target != nil {
			return fmt.Errorf("target still exists")
		}
	}
	return nil
}

func testAccCheckSpotinstMultaiTargetAttributes(target *multai.Target, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(target.Name) != expectedName {
			return fmt.Errorf("bad content: %v", spotinst.StringValue(target.Name))
		}
		return nil
	}
}

func testAccCheckSpotinstMultaiTargetExists(target *multai.Target, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &multai.ReadTargetInput{
			TargetID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.multai.ReadTarget(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.Target.ID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("target not found: %+v,\n %+v\n", resp.Target, rs.Primary.Attributes)
		}
		*target = *resp.Target
		return nil
	}
}

type TargetConfigMetadata struct {
	provider             string
	name                 string
	updateBaselineFields bool
}

func createTargetTerraform(tcm *TargetConfigMetadata) string {
	if tcm == nil {
		return ""
	}

	if tcm.provider == "" {
		tcm.provider = "aws"
	}

	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`

	if tcm.updateBaselineFields {
		format := testBaselineTargetConfig_Update
		template += fmt.Sprintf(format,
			tcm.name,
			tcm.provider,
			tcm.name,
		)
	} else {
		format := testBaselineTargetConfig_Create

		template += fmt.Sprintf(format,
			tcm.name,
			tcm.provider,
			tcm.name,
		)
	}

	log.Printf("Terraform [%v] template:\n%v", tcm.name, template)
	return template
}

func TestAccSpotinstMultaiTarget_Baseline(t *testing.T) {
	targetName := "test-acc-target-baseline"
	resourceName := createMultaiTargetResourceName(targetName)

	var target multai.Target
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckSpotinstMultaiTargetDestroy,

		Steps: []resource.TestStep{
			{
				Config: createTargetTerraform(&TargetConfigMetadata{
					name: targetName,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiTargetExists(&target, resourceName),
					testAccCheckSpotinstMultaiTargetAttributes(&target, targetName),
					resource.TestCheckResourceAttr(resourceName, "port", "1337"),
					resource.TestCheckResourceAttr(resourceName, "host", "host"),
					resource.TestCheckResourceAttr(resourceName, "weight", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags."+TargetTagsHash_Create+".key", "fakeKey"),
					resource.TestCheckResourceAttr(resourceName, "tags."+TargetTagsHash_Create+".value", "fakeVal"),
				),
			},
			{
				Config: createTargetTerraform(&TargetConfigMetadata{
					name:                 targetName,
					updateBaselineFields: true,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiTargetExists(&target, resourceName),
					testAccCheckSpotinstMultaiTargetAttributes(&target, targetName),
					resource.TestCheckResourceAttr(resourceName, "port", "1338"),
					resource.TestCheckResourceAttr(resourceName, "host", "host-updated"),
					resource.TestCheckResourceAttr(resourceName, "weight", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags."+TargetTagsHash_Update+".key", "updated"),
					resource.TestCheckResourceAttr(resourceName, "tags."+TargetTagsHash_Update+".value", "updated"),
				),
			},
		},
	})
}

const (
	TargetTagsHash_Create = "2538041064"
	TargetTagsHash_Update = "1968254376"
)

const testBaselineTargetConfig_Create = `
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

resource "` + string(commons.MultaiTargetResourceName) + `" "%v" {
  provider = "%v"
  name     = "%v"
  balancer_id = "${spotinst_multai_balancer.foo.id}"
  target_set_id = "${spotinst_multai_target_set.foo.id}"
  port        = 1337
  host        = "host"
  weight      = 1

  tags {
   key = "fakeKey"
   value = "fakeVal"
  }
}`

const testBaselineTargetConfig_Update = `
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

resource "` + string(commons.MultaiTargetResourceName) + `" "%v" {
  provider = "%v"
  name = "%v"
  balancer_id = "${spotinst_multai_balancer.foo.id}"
  target_set_id = "${spotinst_multai_target_set.foo.id}"
  port        = 1338
  host        = "host-updated"
  weight      = 2 

  tags {
   key = "updated"
   value = "updated"
  }
}`

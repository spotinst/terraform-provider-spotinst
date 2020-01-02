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
	resource.AddTestSweepers("spotinst_multai_target_set", &resource.Sweeper{
		Name: "spotinst_multai_target_set",
		F:    testSweepMultaiTargetSet,
	})
}

func testSweepMultaiTargetSet(region string) error {
	client, err := getProviderClient("aws")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).multai
	input := &multai.ListTargetSetsInput{}
	if resp, err := conn.ListTargetSets(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of target sets to sweep")
	} else {
		if len(resp.TargetSets) == 0 {
			log.Printf("[INFO] No target sets to sweep")
		}
		for _, set := range resp.TargetSets {
			if strings.Contains(spotinst.StringValue(set.Name), "test-acc-") {
				if _, err := conn.DeleteTargetSet(context.Background(), &multai.DeleteTargetSetInput{TargetSetID: set.ID}); err != nil {
					return fmt.Errorf("unable to delete target set %v in sweep", spotinst.StringValue(set.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(set.ID))
				}
			}
		}
	}
	return nil
}

func createMultaiTargetSetResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.MultaiTargetSetResourceName), name)
}

func testAccCheckSpotinstMultaiTargetSetDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "spotinst_multai_target_set" {
			continue
		}
		input := &multai.ReadTargetSetInput{
			TargetSetID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.multai.ReadTargetSet(context.Background(), input)
		if err == nil && resp != nil && resp.TargetSet != nil {
			return fmt.Errorf("target set still exists")
		}
	}
	return nil
}

func testAccCheckSpotinstMultaiTargetSetAttributes(targetSet *multai.TargetSet, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(targetSet.Name) != expectedName {
			return fmt.Errorf("bad content: %v", spotinst.StringValue(targetSet.Name))
		}
		return nil
	}
}

func testAccCheckSpotinstMultaiTargetSetExists(targetSet *multai.TargetSet, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &multai.ReadTargetSetInput{
			TargetSetID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.multai.ReadTargetSet(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.TargetSet.ID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("target set not found: %+v,\n %+v\n", resp.TargetSet, rs.Primary.Attributes)
		}
		*targetSet = *resp.TargetSet
		return nil
	}
}

type TargetSetConfigMetadata struct {
	provider             string
	name                 string
	updateBaselineFields bool
}

func createTargetSetTerraform(tscm *TargetSetConfigMetadata) string {
	if tscm == nil {
		return ""
	}

	if tscm.provider == "" {
		tscm.provider = "aws"
	}

	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`

	if tscm.updateBaselineFields {
		format := testBaselineTargetSetConfig_Update
		template += fmt.Sprintf(format,
			tscm.name,
			tscm.provider,
			tscm.name,
		)
	} else {
		format := testBaselineTargetSetConfig_Create

		template += fmt.Sprintf(format,
			tscm.name,
			tscm.provider,
			tscm.name,
		)
	}

	log.Printf("Terraform [%v] template:\n%v", tscm.name, template)
	return template
}

func TestAccSpotinstMultaiTargetSet_Baseline(t *testing.T) {
	targetSetName := "test-acc-target-set-baseline"
	resourceName := createMultaiTargetSetResourceName(targetSetName)

	var targetSet multai.TargetSet
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckSpotinstMultaiTargetSetDestroy,

		Steps: []resource.TestStep{
			{
				Config: createTargetSetTerraform(&TargetSetConfigMetadata{
					name: targetSetName,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiTargetSetExists(&targetSet, resourceName),
					testAccCheckSpotinstMultaiTargetSetAttributes(&targetSet, targetSetName),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "port", "1337"),
					resource.TestCheckResourceAttr(resourceName, "weight", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_check.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_check."+TargetSetHealthCheck_Create+".protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "health_check."+TargetSetHealthCheck_Create+".path", "/"),
					resource.TestCheckResourceAttr(resourceName, "health_check."+TargetSetHealthCheck_Create+".port", "3000"),
					resource.TestCheckResourceAttr(resourceName, "health_check."+TargetSetHealthCheck_Create+".interval", "30"),
					resource.TestCheckResourceAttr(resourceName, "health_check."+TargetSetHealthCheck_Create+".timeout", "10"),
					resource.TestCheckResourceAttr(resourceName, "health_check."+TargetSetHealthCheck_Create+".healthy_threshold", "2"),
					resource.TestCheckResourceAttr(resourceName, "health_check."+TargetSetHealthCheck_Create+".unhealthy_threshold", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags."+TargetSetTagsHash_Create+".key", "fakeKey"),
					resource.TestCheckResourceAttr(resourceName, "tags."+TargetSetTagsHash_Create+".value", "fakeVal"),
				),
			},
			{
				Config: createTargetSetTerraform(&TargetSetConfigMetadata{
					name:                 targetSetName,
					updateBaselineFields: true,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiTargetSetExists(&targetSet, resourceName),
					testAccCheckSpotinstMultaiTargetSetAttributes(&targetSet, targetSetName),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "port", "1338"),
					resource.TestCheckResourceAttr(resourceName, "weight", "2"),
					resource.TestCheckResourceAttr(resourceName, "health_check.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health_check."+TargetSetHealthCheck_Update+".protocol", "HTTP"),
					resource.TestCheckResourceAttr(resourceName, "health_check."+TargetSetHealthCheck_Update+".path", "/"),
					resource.TestCheckResourceAttr(resourceName, "health_check."+TargetSetHealthCheck_Update+".port", "3001"),
					resource.TestCheckResourceAttr(resourceName, "health_check."+TargetSetHealthCheck_Update+".interval", "20"),
					resource.TestCheckResourceAttr(resourceName, "health_check."+TargetSetHealthCheck_Update+".timeout", "5"),
					resource.TestCheckResourceAttr(resourceName, "health_check."+TargetSetHealthCheck_Update+".healthy_threshold", "3"),
					resource.TestCheckResourceAttr(resourceName, "health_check."+TargetSetHealthCheck_Update+".unhealthy_threshold", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags."+TargetSetTagsHash_Update+".key", "updated"),
					resource.TestCheckResourceAttr(resourceName, "tags."+TargetSetTagsHash_Update+".value", "updated"),
				),
			},
		},
	})
}

const (
	TargetSetTagsHash_Create    = "2538041064"
	TargetSetTagsHash_Update    = "1968254376"
	TargetSetHealthCheck_Create = "3726815735"
	TargetSetHealthCheck_Update = "3181670027"
)

const testBaselineTargetSetConfig_Create = `
resource "spotinst_multai_balancer" "foo" {
  provider = "aws"
  name = "test-acc-foo"

  connection_timeouts {
    idle     = 10
    draining = 10
  }
}

resource "spotinst_multai_target_set" "%v" {
  provider      = "%v"
  name          = "%v"
  balancer_id   = "${spotinst_multai_balancer.foo.id}"
  deployment_id = "dp-12345"
  protocol      = "http"
  port          = 1337
  weight        = 1

  health_check {
    protocol            = "http"
    path                = "/"
    port                = 3000
    interval            = 30
    timeout             = 10
    healthy_threshold   = 2
    unhealthy_threshold = 2
  }

  tags {
   key = "fakeKey"
   value = "fakeVal"
  }
}`

const testBaselineTargetSetConfig_Update = `
resource "spotinst_multai_balancer" "foo" {
  provider = "aws"
  name = "test-acc-foo"

  connection_timeouts {
    idle     = 10
    draining = 10
  }
}

resource "spotinst_multai_target_set" "%v" {
  provider = "%v"
  name     = "%v"
  balancer_id   = "${spotinst_multai_balancer.foo.id}"
  deployment_id = "dp-12345"
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
}`

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
	resource.AddTestSweepers("spotinst_multai_balancer", &resource.Sweeper{
		Name: "spotinst_multai_balancer",
		F:    testSweepMultaiBalancer,
	})
}

func testSweepMultaiBalancer(region string) error {
	client, err := getProviderClient("aws")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).multai
	input := &multai.ListLoadBalancersInput{}
	if resp, err := conn.ListLoadBalancers(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of balancers to sweep")
	} else {
		if len(resp.Balancers) == 0 {
			log.Printf("[INFO] No balancers to sweep")
		}
		for _, bal := range resp.Balancers {
			if strings.Contains(spotinst.StringValue(bal.Name), "test-acc-") {
				if _, err := conn.DeleteLoadBalancer(context.Background(), &multai.DeleteLoadBalancerInput{BalancerID: bal.ID}); err != nil {
					return fmt.Errorf("unable to delete balancer %v in sweep", spotinst.StringValue(bal.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(bal.ID))
				}
			}
		}
	}
	return nil
}

func createMultaiBalancerResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.MultaiBalancerResourceName), name)
}

func testAccCheckSpotinstMultaiBalancerDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "spotinst_multai_balancer" {
			continue
		}
		input := &multai.ReadLoadBalancerInput{
			BalancerID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.multai.ReadLoadBalancer(context.Background(), input)
		if err == nil && resp != nil && resp.Balancer != nil {
			return fmt.Errorf("balancer still exists")
		}
	}
	return nil
}

func testAccCheckSpotinstMultaiBalancerAttributes(balancer *multai.LoadBalancer, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if p := spotinst.StringValue(balancer.Name); p != expectedName {
			return fmt.Errorf("bad content: %s", p)
		}
		return nil
	}
}

func testAccCheckSpotinstMultaiBalancerExists(balancer *multai.LoadBalancer, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &multai.ReadLoadBalancerInput{
			BalancerID: spotinst.String(rs.Primary.ID),
		}
		resp, err := client.multai.ReadLoadBalancer(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.Balancer.ID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("balancer not found: %+v,\n %+v\n", resp.Balancer, rs.Primary.Attributes)
		}
		*balancer = *resp.Balancer
		return nil
	}
}

type BalancerConfigMetadata struct {
	provider             string
	name                 string
	updateBaselineFields bool
}

func createBalancerTerraform(bcm *BalancerConfigMetadata) string {
	if bcm == nil {
		return ""
	}

	if bcm.provider == "" {
		bcm.provider = "aws"
	}

	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`

	if bcm.updateBaselineFields {
		format := testBaselineBalancerConfig_Update
		template += fmt.Sprintf(format,
			bcm.name,
			bcm.provider,
			bcm.name,
		)
	} else {
		format := testBaselineBalancerConfig_Create

		template += fmt.Sprintf(format,
			bcm.name,
			bcm.provider,
			bcm.name,
		)
	}

	log.Printf("Terraform [%v] template:\n%v", bcm.name, template)
	return template
}

func TestAccSpotinstMultaiBalancer_Baseline(t *testing.T) {
	balName := "test-acc-mlb-baseline"
	resourceName := createMultaiBalancerResourceName(balName)

	var balancer multai.LoadBalancer
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckSpotinstMultaiBalancerDestroy,

		Steps: []resource.TestStep{
			{
				Config: createBalancerTerraform(&BalancerConfigMetadata{
					name: balName,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiBalancerExists(&balancer, resourceName),
					testAccCheckSpotinstMultaiBalancerAttributes(&balancer, balName),
					resource.TestCheckResourceAttr(resourceName, "name", "test-acc-mlb-baseline"),
					resource.TestCheckResourceAttr(resourceName, "scheme", "internal"),
					resource.TestCheckResourceAttr(resourceName, "connection_timeouts.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connection_timeouts."+BalancerTimeoutsHash_Create+".idle", "10"),
					resource.TestCheckResourceAttr(resourceName, "connection_timeouts."+BalancerTimeoutsHash_Create+".draining", "10"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags."+BalancerTagsHash_Create+".key", "prod"),
					resource.TestCheckResourceAttr(resourceName, "tags."+BalancerTagsHash_Create+".value", "web"),
				),
			},
			{
				Config: createBalancerTerraform(&BalancerConfigMetadata{
					name:                 balName,
					updateBaselineFields: true,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpotinstMultaiBalancerExists(&balancer, resourceName),
					testAccCheckSpotinstMultaiBalancerAttributes(&balancer, balName),
					resource.TestCheckResourceAttr(resourceName, "name", "test-acc-mlb-baseline"),
					resource.TestCheckResourceAttr(resourceName, "scheme", "internet-facing"),
					resource.TestCheckResourceAttr(resourceName, "connection_timeouts."+BalancerTimeoutsHash_Update+".idle", "20"),
					resource.TestCheckResourceAttr(resourceName, "connection_timeouts."+BalancerTimeoutsHash_Update+".draining", "20"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags."+BalancerTagsHash_Update+".key", "updated"),
					resource.TestCheckResourceAttr(resourceName, "tags."+BalancerTagsHash_Update+".value", "updated"),
				),
			},
		},
	})
}

const (
	BalancerTimeoutsHash_Create = "4167278370"
	BalancerTimeoutsHash_Update = "1674004346"
	BalancerTagsHash_Create     = "2247434205"
	BalancerTagsHash_Update     = "1968254376"
)

const testBaselineBalancerConfig_Create = `
resource "` + string(commons.MultaiBalancerResourceName) + `" "%v" {
  provider = "%v"
  name     = "%v"
  scheme   = "internal"

  dns_cname_aliases = []

  connection_timeouts {
    idle     = 10
    draining = 10
  }

  tags {
    key   = "prod"
    value = "web"
  }
}`

const testBaselineBalancerConfig_Update = `
resource "` + string(commons.MultaiBalancerResourceName) + `" "%v" {
  provider = "%v"
  name     = "%v"
  scheme   = "internet-facing"

  dns_cname_aliases = []

  connection_timeouts {
    idle     = 20
    draining = 20
  }

  tags {
    key   = "updated"
    value = "updated"
  }
}`

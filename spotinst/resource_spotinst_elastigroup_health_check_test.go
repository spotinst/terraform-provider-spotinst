package spotinst

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/healthcheck"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

func createHealthCheckResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.HealthCheckResourceName), name)
}

func testHealthCheckDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.HealthCheckResourceName) {
			continue
		}
		input := &healthcheck.ReadHealthCheckInput{HealthCheckID: spotinst.String(rs.Primary.ID)}
		resp, err := client.healthCheck.Read(context.Background(), input)
		if err == nil && resp != nil && resp.HealthCheck != nil {
			return fmt.Errorf("healthCheck still exists")
		}
	}
	return nil
}

func testCheckHealthCheckAttributes(healthCheck *healthcheck.HealthCheck, expectedName string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		if spotinst.StringValue(healthCheck.Name) != expectedName {
			return fmt.Errorf("bad content: %v", healthCheck.ID)
		}
		return nil
	}
}

func testCheckHealthCheckExists(healthCheck *healthcheck.HealthCheck, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &healthcheck.ReadHealthCheckInput{HealthCheckID: spotinst.String(rs.Primary.ID)}
		resp, err := client.healthCheck.Read(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.HealthCheck.ID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("healthCheck not found: %+v,\n %+v\n", resp.HealthCheck, rs.Primary.Attributes)
		}
		*healthCheck = *resp.HealthCheck
		return nil
	}
}

type HealthCheckMetadata struct {
	provider             string
	name                 string
	variables            string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createHealthCheckTerraform(ccm *HealthCheckMetadata, formatToUse string) string {
	if ccm == nil {
		return ""
	}

	if ccm.provider == "" {
		ccm.provider = "aws"
	}

	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`

	if ccm.updateBaselineFields {
		format := testBaselineHealthCheckConfig_Update
		template += fmt.Sprintf(format,
			ccm.name,
			ccm.provider,
			ccm.fieldsToAppend,
		)
	} else {
		format := testBaselineHealthCheckConfig_Create
		template += fmt.Sprintf(format,
			ccm.name,
			ccm.provider,
			ccm.fieldsToAppend,
		)
	}

	if ccm.variables != "" {
		template = ccm.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", "healt_check_test", template)
	return template
}

// region HealthCheck: Baseline
func TestAccSpotinstHealthCheckBaseline(t *testing.T) {
	name := "test-acc-health_check_terraform_test"
	resourceName := createHealthCheckResourceName(name)

	var healthCheck healthcheck.HealthCheck
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testHealthCheckDestroy,

		Steps: []resource.TestStep{
			{
				Config: createHealthCheckTerraform(&HealthCheckMetadata{
					name: name,
				}, testBaselineHealthCheckConfig_Create),

				Check: resource.ComposeTestCheckFunc(
					testCheckHealthCheckExists(&healthCheck, resourceName),
					testCheckHealthCheckAttributes(&healthCheck, name),
					resource.TestCheckResourceAttr(resourceName, "name", "test-acc-health_check_terraform_test"),
					resource.TestCheckResourceAttr(resourceName, "proxy_address", "http://proxy.com"),
					resource.TestCheckResourceAttr(resourceName, "proxy_port", "6"),
					resource.TestCheckResourceAttr(resourceName, "check.0.endpoint", "http://endpoint.com"),
					resource.TestCheckResourceAttr(resourceName, "check.0.healthy", "2"),
					resource.TestCheckResourceAttr(resourceName, "check.0.interval", "11"),
					resource.TestCheckResourceAttr(resourceName, "check.0.port", "1336"),
					resource.TestCheckResourceAttr(resourceName, "check.0.protocol", "http"),
					resource.TestCheckResourceAttr(resourceName, "check.0.timeout", "12"),
					resource.TestCheckResourceAttr(resourceName, "check.0.unhealthy", "3"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createHealthCheckTerraform(&HealthCheckMetadata{
					name:                 name,
					updateBaselineFields: true}, testBaselineHealthCheckConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckHealthCheckExists(&healthCheck, resourceName),
					testCheckHealthCheckAttributes(&healthCheck, name),
					resource.TestCheckResourceAttr(resourceName, "name", "test-acc-health_check_terraform_test"),
					resource.TestCheckResourceAttr(resourceName, "proxy_address", "http://proxy.com"),
					resource.TestCheckResourceAttr(resourceName, "proxy_port", "7"),
					resource.TestCheckResourceAttr(resourceName, "check.0.endpoint", "http://endpoint.com"),
					resource.TestCheckResourceAttr(resourceName, "check.0.healthy", "3"),
					resource.TestCheckResourceAttr(resourceName, "check.0.interval", "15"),
					resource.TestCheckResourceAttr(resourceName, "check.0.port", "1335"),
					resource.TestCheckResourceAttr(resourceName, "check.0.protocol", "http"),
					resource.TestCheckResourceAttr(resourceName, "check.0.timeout", "13"),
					resource.TestCheckResourceAttr(resourceName, "check.0.unhealthy", "2"),
				),
			},
		},
	})
}

const testBaselineHealthCheckConfig_Create = `
resource "` + string(commons.HealthCheckResourceName) + `" "%v" {
  provider = "%v"
    resource_id = "sig-05d0a009"
    name = "test-acc-health_check_terraform_test"
    proxy_address = "http://proxy.com"
    proxy_port = "6"
  check {
    protocol = "http"
    port = "1336"
    endpoint = "http://endpoint.com"
    interval = "11"
    timeout = "12"
    unhealthy  = "3"
    healthy = "2"
  }
 %v
}
`

const testBaselineHealthCheckConfig_Update = `
resource "` + string(commons.HealthCheckResourceName) + `" "%v" {
  provider = "%v"
  resource_id = "sig-05d0a009"
  name = "test-acc-health_check_terraform_test"
  proxy_address = "http://proxy.com"
  proxy_port = "7"
  check {
   protocol = "http"
   port = "1335"
   endpoint = "http://endpoint.com"
   interval = "15"
   timeout = "13"
   unhealthy  = "2"
   healthy = "3"
  }
  %v
}
`

// endregion

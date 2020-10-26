package spotinst

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/mrscaler"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func init() {
	resource.AddTestSweepers("spotinst_mrscaler_aws", &resource.Sweeper{
		Name: "spotinst_mrscaler_aws",
		F:    testSweepMRScalerAWS,
	})
}

func testSweepMRScalerAWS(region string) error {
	client, err := getProviderClient("aws")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).mrscaler
	input := &mrscaler.ListScalersInput{}
	if resp, err := conn.List(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of clusters to sweep")
	} else {
		if len(resp.Scalers) == 0 {
			log.Printf("[INFO] No scalers to sweep")
		}
		for _, scaler := range resp.Scalers {
			if strings.Contains(spotinst.StringValue(scaler.Name), "test-acc-") {
				if _, err := conn.Delete(context.Background(), &mrscaler.DeleteScalerInput{ScalerID: scaler.ID}); err != nil {
					return fmt.Errorf("unable to delete scaler %v in sweep", spotinst.StringValue(scaler.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(scaler.ID))
				}
			}
		}
	}
	return nil
}

func createMRScalerAWSResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.MRScalerAWSResourceName), name)
}

func testMRScalerAWSDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.MRScalerAWSResourceName) {
			continue
		}
		input := &mrscaler.ReadScalerInput{ScalerID: spotinst.String(rs.Primary.ID)}
		resp, err := client.mrscaler.Read(context.Background(), input)
		if err == nil && resp != nil && resp.Scaler != nil {
			return fmt.Errorf("scaler still exists")
		}
	}
	return nil
}

func testCheckMRScalerAWSAttributes(scaler *mrscaler.Scaler, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(scaler.Name) != expectedName {
			return fmt.Errorf("bad content: %v", scaler.Name)
		}
		return nil
	}
}

func testCheckMRScalerAWSExists(scaler *mrscaler.Scaler, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &mrscaler.ReadScalerInput{ScalerID: spotinst.String(rs.Primary.ID)}
		resp, err := client.mrscaler.Read(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.Scaler.Name) != rs.Primary.Attributes["name"] {
			return fmt.Errorf("mrscaler not found: %+v,\n %+v\n", resp.Scaler, rs.Primary.Attributes)
		}
		*scaler = *resp.Scaler
		return nil
	}
}

type MRScalerAWSConfigMetaData struct {
	provider             string
	scalerName           string
	newCluster           bool
	updateBaselineFields bool
}

func createMRScalerAWSTerraform(mcm *MRScalerAWSConfigMetaData) string {
	// check if "make testacc' is being ran, and sleep. Causes timeouts when running "make test"
	if os.Getenv("TF_ACC") == "1" {
		time.Sleep(30 * time.Second)
	}
	if mcm == nil {
		return ""
	}

	if mcm.provider == "" {
		mcm.provider = "aws"
	}

	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`
	if mcm.updateBaselineFields {
		format := testMRScalerAWSBaseline_Update
		template += fmt.Sprintf(format,
			mcm.scalerName,
			mcm.provider,
		)
	} else {
		format := testMRScalerAWSBaseline_Create
		template += fmt.Sprintf(format,
			mcm.scalerName,
			mcm.provider,
		)
	}

	log.Printf("Terraform [%v] template:\n%v", mcm.scalerName, template)
	return template
}

// region MRScalerAWS: Baseline
func TestAccSpotinstNewMRScalerAWS_Baseline(t *testing.T) {
	scalerName := "mrscaler-baseline"
	resourceName := createMRScalerAWSResourceName(scalerName)

	var scaler mrscaler.Scaler
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testMRScalerAWSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createMRScalerAWSTerraform(&MRScalerAWSConfigMetaData{
					scalerName: scalerName,
					newCluster: true,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckMRScalerAWSExists(&scaler, resourceName),
					testCheckMRScalerAWSAttributes(&scaler, "scalerName"),
					resource.TestCheckResourceAttr(resourceName, "description", "test create"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "strategy", "new"),
					resource.TestCheckResourceAttr(resourceName, "release_label", "emr-5.0.3"),
					resource.TestCheckResourceAttr(resourceName, "retries", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.0", "us-east-1a:subnet-b3b150ec"),
					resource.TestCheckResourceAttr(resourceName, "provisioning_timeout.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "provisioning_timeout.0.timeout", "15"),
					resource.TestCheckResourceAttr(resourceName, "provisioning_timeout.0.timeout_action", "terminateAndRetry"),
					resource.TestCheckResourceAttr(resourceName, "job_flow_role", "EC2Access"),
					resource.TestCheckResourceAttr(resourceName, "service_role", "Core-Services-Admin-Role"),
					resource.TestCheckResourceAttr(resourceName, "master_instance_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "master_instance_types.0", "c3.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "master_lifecycle", "ON_DEMAND"),
					resource.TestCheckResourceAttr(resourceName, "master_ebs_block_device.2770074213.volumes_per_instance", "1"),
					resource.TestCheckResourceAttr(resourceName, "master_ebs_block_device.2770074213.volume_type", "gp2"),
					resource.TestCheckResourceAttr(resourceName, "master_ebs_block_device.2770074213.size_in_gb", "10"),
					resource.TestCheckResourceAttr(resourceName, "core_instance_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "core_instance_types.0", "c3.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "core_min_size", "1"),
					resource.TestCheckResourceAttr(resourceName, "core_max_size", "2"),
					resource.TestCheckResourceAttr(resourceName, "core_desired_capacity", "1"),
					resource.TestCheckResourceAttr(resourceName, "core_lifecycle", "ON_DEMAND"),
					resource.TestCheckResourceAttr(resourceName, "core_ebs_optimized", "false"),
					resource.TestCheckResourceAttr(resourceName, "core_ebs_block_device.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "core_ebs_block_device.2385947046.volumes_per_instance", "2"),
					//resource.TestCheckResourceAttr(resourceName, "core_ebs_block_device.2385947046.volume_type", "gp2"),
					//resource.TestCheckResourceAttr(resourceName, "core_ebs_block_device.2385947046.size_in_gb", "40"),
					resource.TestCheckResourceAttr(resourceName, "tags.664003903.key", "Creator"),
					resource.TestCheckResourceAttr(resourceName, "tags.664003903.value", "Terraform"),
					resource.TestCheckResourceAttr(resourceName, "instance_weights.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_weights.1628034056.instance_type", "m3.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_weights.1628034056.weighted_capacity", "4"),
					resource.TestCheckResourceAttr(resourceName, "keep_job_flow_alive", "true"),
					resource.TestCheckResourceAttr(resourceName, "task_instance_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "task_instance_types.0", "c3.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "task_min_size", "5"),
					resource.TestCheckResourceAttr(resourceName, "task_max_size", "5"),
					resource.TestCheckResourceAttr(resourceName, "task_desired_capacity", "5"),
					resource.TestCheckResourceAttr(resourceName, "task_lifecycle", "ON_DEMAND"),
					resource.TestCheckResourceAttr(resourceName, "task_ebs_optimized", "false"),
					resource.TestCheckResourceAttr(resourceName, "task_unit", "instance"),
					resource.TestCheckResourceAttr(resourceName, "task_ebs_block_device.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "task_ebs_block_device.3329897523.volumes_per_instance", "2"),
					resource.TestCheckResourceAttr(resourceName, "task_ebs_block_device.3329897523.volume_type", "gp2"),
					resource.TestCheckResourceAttr(resourceName, "task_ebs_block_device.3329897523.size_in_gb", "40"),
				),
			},
			{
				Config: createMRScalerAWSTerraform(&MRScalerAWSConfigMetaData{
					scalerName:           scalerName,
					newCluster:           true,
					updateBaselineFields: true,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckMRScalerAWSExists(&scaler, resourceName),
					testCheckMRScalerAWSAttributes(&scaler, "scalerName"),
					resource.TestCheckResourceAttr(resourceName, "description", "test update"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "strategy", "new"),
					resource.TestCheckResourceAttr(resourceName, "release_label", "emr-5.0.3"),
					resource.TestCheckResourceAttr(resourceName, "retries", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.0", "us-east-1a:subnet-b3b150ec"),
					resource.TestCheckResourceAttr(resourceName, "provisioning_timeout.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "provisioning_timeout.0.timeout", "20"),
					resource.TestCheckResourceAttr(resourceName, "provisioning_timeout.0.timeout_action", "terminate"),
					resource.TestCheckResourceAttr(resourceName, "job_flow_role", "EC2Access"),
					resource.TestCheckResourceAttr(resourceName, "service_role", "Core-Services-Admin-Role"),
					resource.TestCheckResourceAttr(resourceName, "master_instance_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "master_instance_types.0", "c3.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "master_lifecycle", "ON_DEMAND"),
					resource.TestCheckResourceAttr(resourceName, "master_ebs_block_device.2770074213.volumes_per_instance", "1"),
					resource.TestCheckResourceAttr(resourceName, "master_ebs_block_device.2770074213.volume_type", "gp2"),
					resource.TestCheckResourceAttr(resourceName, "master_ebs_block_device.2770074213.size_in_gb", "10"),
					resource.TestCheckResourceAttr(resourceName, "core_instance_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "core_instance_types.0", "c3.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "core_min_size", "1"),
					resource.TestCheckResourceAttr(resourceName, "core_max_size", "2"),
					resource.TestCheckResourceAttr(resourceName, "core_desired_capacity", "1"),
					resource.TestCheckResourceAttr(resourceName, "core_lifecycle", "ON_DEMAND"),
					resource.TestCheckResourceAttr(resourceName, "core_ebs_optimized", "true"),
					resource.TestCheckResourceAttr(resourceName, "core_ebs_block_device.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "core_ebs_block_device.1402340189.volumes_per_instance", "2"),
					//resource.TestCheckResourceAttr(resourceName, "core_ebs_block_device.1402340189.volume_type", "gp2"),
					//resource.TestCheckResourceAttr(resourceName, "core_ebs_block_device.1402340189.size_in_gb", "40"),
					resource.TestCheckResourceAttr(resourceName, "tags.664003903.key", "Creator"),
					resource.TestCheckResourceAttr(resourceName, "tags.664003903.value", "Terraform"),
					resource.TestCheckResourceAttr(resourceName, "instance_weights.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_weights.454376557.instance_type", "m2.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_weights.454376557.weighted_capacity", "5"),
					resource.TestCheckResourceAttr(resourceName, "keep_job_flow_alive", "true"),
					resource.TestCheckResourceAttr(resourceName, "task_instance_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "task_instance_types.0", "c3.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "task_min_size", "1"),
					resource.TestCheckResourceAttr(resourceName, "task_max_size", "3"),
					resource.TestCheckResourceAttr(resourceName, "task_desired_capacity", "2"),
					resource.TestCheckResourceAttr(resourceName, "task_lifecycle", "ON_DEMAND"),
					resource.TestCheckResourceAttr(resourceName, "task_ebs_optimized", "false"),
					resource.TestCheckResourceAttr(resourceName, "task_unit", "instance"),
					resource.TestCheckResourceAttr(resourceName, "task_ebs_block_device.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "task_ebs_block_device.1008334328.volumes_per_instance", "1"),
					resource.TestCheckResourceAttr(resourceName, "task_ebs_block_device.1008334328.volume_type", "gp2"),
					resource.TestCheckResourceAttr(resourceName, "task_ebs_block_device.1008334328.size_in_gb", "30"),
				),
			},
		},
	})
}

const testMRScalerAWSBaseline_Create = `
resource "` + string(commons.MRScalerAWSResourceName) + `" "%v" {
 provider = "%v"

 name               = "scalerName"
 description        = "test create"
 strategy           = "new"
 release_label      = "emr-5.0.3"
 retries            = 1
 region             = "us-east-1"
 availability_zones = ["us-east-1a:subnet-b3b150ec"]
 
 provisioning_timeout {
   timeout        = 15
   timeout_action = "terminateAndRetry"
  }

 job_flow_role   = "EC2Access"
 service_role = "Core-Services-Admin-Role"
 master_instance_types = ["c3.xlarge"]
 master_lifecycle      = "ON_DEMAND"
 
 master_ebs_block_device   {
   volumes_per_instance = 1
   volume_type          = "gp2"
   size_in_gb           = 10
 }
 core_instance_types = ["c3.xlarge"]
 core_min_size         = 1
 core_max_size         = 2
 core_desired_capacity = 1
 core_lifecycle        = "ON_DEMAND"
 core_ebs_optimized    = false
 
  core_ebs_block_device   {
   volumes_per_instance = 2
   volume_type          = "gp2"
   size_in_gb           = 40
 }
 
 task_instance_types = ["c3.xlarge"]
 task_min_size         = 5
 task_max_size         = 5
 task_desired_capacity = 5
 task_lifecycle        = "ON_DEMAND"
 task_ebs_optimized    = false
 task_unit             = "instance"
 
 task_ebs_block_device    {
   volumes_per_instance = 2
   volume_type          = "gp2"
   size_in_gb           = 40
 }
 
 instance_weights    {
  instance_type     = "m3.xlarge"
  weighted_capacity = 4 
 }

 keep_job_flow_alive   = true

tags {
    key   = "Creator"
    value = "Terraform"
  }
}
`

const testMRScalerAWSBaseline_Update = `
resource "` + string(commons.MRScalerAWSResourceName) + `" "%v" {
 provider = "%v"

 name               = "scalerName"
 description        = "test update"
 strategy           = "new"
 release_label      = "emr-5.0.3"
 retries            = 1
 region             = "us-east-1"
 availability_zones = ["us-east-1a:subnet-b3b150ec"]

 provisioning_timeout {
   timeout        = 20
   timeout_action = "terminate"
  }

 job_flow_role   = "EC2Access"
 service_role = "Core-Services-Admin-Role"

 master_instance_types = ["c3.xlarge"]
 master_lifecycle      = "ON_DEMAND"

 master_ebs_block_device   {
    volumes_per_instance = 1
    volume_type          = "gp2"
    size_in_gb           = 10
  }
  core_instance_types = ["c3.xlarge"]
  core_min_size         = 1
  core_max_size         = 2
  core_desired_capacity = 1
  core_lifecycle        = "ON_DEMAND"
  core_ebs_optimized    = true

   core_ebs_block_device   {
    volumes_per_instance = 2
    volume_type          = "gp2"
    size_in_gb           = 10
  }

   task_instance_types = ["c3.xlarge"]
  task_min_size         = 1
  task_max_size         = 3
  task_desired_capacity = 2
  task_lifecycle        = "ON_DEMAND"
  task_ebs_optimized    = false
  task_unit             = "instance"

 task_ebs_block_device    {
    volumes_per_instance = 1
    volume_type          = "gp2"
    size_in_gb           = 30
  }

 instance_weights    {
   instance_type     = "m2.xlarge"
   weighted_capacity = 5 
 }

  keep_job_flow_alive   = true

tags {
    key   = "Creator"
    value = "Terraform"
  }

}
`

// endregion

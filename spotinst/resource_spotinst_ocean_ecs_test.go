package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func init() {
	resource.AddTestSweepers("spotinst_ocean_ecs", &resource.Sweeper{
		Name: "spotinst_ocean_ecs",
		F:    testSweepOceanECS,
	})
}

func testSweepOceanECS(region string) error {
	client, err := getProviderClient("aws")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).ocean.CloudProviderAWS()
	input := &aws.ListECSClustersInput{}
	if resp, err := conn.ListECSClusters(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of clusters to sweep")
	} else {
		if len(resp.Clusters) == 0 {
			log.Printf("[INFO] No clusters to sweep")
		}
		for _, cluster := range resp.Clusters {
			if strings.Contains(spotinst.StringValue(cluster.Name), "test-acc-") {
				if _, err := conn.DeleteECSCluster(context.Background(), &aws.DeleteECSClusterInput{ClusterID: cluster.ID}); err != nil {
					return fmt.Errorf("unable to delete cluster %v in sweep", spotinst.StringValue(cluster.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(cluster.ID))
				}
			}
		}
	}
	return nil
}

func createOceanECSResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanECSResourceName), name)
}

func testOceanECSDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OceanECSResourceName) {
			continue
		}
		input := &aws.ReadECSClusterInput{ClusterID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAWS().ReadECSCluster(context.Background(), input)
		if err == nil && resp != nil && resp.Cluster != nil {
			return fmt.Errorf("cluster still exists")
		}
	}
	return nil
}

func testCheckOceanECSAttributes(cluster *aws.ECSCluster, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(cluster.Name) != expectedName {
			return fmt.Errorf("bad content: %v", cluster.Name)
		}
		return nil
	}
}

func testCheckOceanECSExists(cluster *aws.ECSCluster, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &aws.ReadECSClusterInput{ClusterID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAWS().ReadECSCluster(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.Cluster.Name) != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Cluster not found: %+v,\n %+v\n", resp.Cluster, rs.Primary.Attributes)
		}
		*cluster = *resp.Cluster
		return nil
	}
}

type ECSClusterConfigMetadata struct {
	provider             string
	name                 string
	clusterName          string
	instanceWhitelist    string
	instanceBlacklist    string
	launchSpec           string
	variables            string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createOceanECSTerraform(ccm *ECSClusterConfigMetadata) string {
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
		format := testBaselineECSConfig_Update
		template += fmt.Sprintf(format,
			ccm.clusterName,
			ccm.provider,
			ccm.clusterName,
			ccm.name,
			ccm.instanceBlacklist,
			ccm.instanceWhitelist,
			ccm.launchSpec,
			ccm.fieldsToAppend,
		)
	} else {
		format := testBaselineECSConfig_Create
		template += fmt.Sprintf(format,
			ccm.clusterName,
			ccm.provider,
			ccm.clusterName,
			ccm.name,
			ccm.instanceWhitelist,
			ccm.instanceBlacklist,
			ccm.launchSpec,
			ccm.fieldsToAppend,
		)
	}

	if ccm.variables != "" {
		template = ccm.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", ccm.clusterName, template)
	return template
}

// region OceanECS: Baseline
func TestAccSpotinstOceanECS_Baseline(t *testing.T) {
	name := "test-acc-cluster-baseline"
	clusterName := "baseline-cluster-name"
	resourceName := createOceanECSResourceName(clusterName)

	var cluster aws.ECSCluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanECSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					name:        name,
					clusterName: clusterName,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.0", "subnet-0bd585d2c2177c7ee"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.0", "sg-0a8e7b3cd1cfd3d6f"),
					resource.TestCheckResourceAttr(resourceName, "utilize_reserved_instances", "false"),
				),
			},
			{
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					clusterName:          clusterName,
					name:                 name,
					updateBaselineFields: true}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.0", "subnet-0bd585d2c2177c7ee"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.1", "subnet-0faad0b6bb7e99d9f"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.0", "sg-987654"),
					resource.TestCheckResourceAttr(resourceName, "utilize_reserved_instances", "true"),
				),
			},
		},
	})
}

const testBaselineECSConfig_Create = `
resource "` + string(commons.OceanECSResourceName) + `" "%v" {
  provider = "%v"  

  cluster_name = "%v"
  name = "%v"
  region = "us-west-2"

  //max_size = 1
  //min_size = 0
  //desired_capacity = 0

  subnet_ids         = ["subnet-0bd585d2c2177c7ee"]
  security_group_ids = ["sg-0a8e7b3cd1cfd3d6f"]
  utilize_reserved_instances = false


 %v
 %v
 %v
 %v
}
`

const testBaselineECSConfig_Update = `
resource "` + string(commons.OceanECSResourceName) + `" "%v" {
  provider = "%v"

  cluster_name = "%v"
  name = "%v"
  region = "us-west-2"

  //max_size = 1
  //min_size = 0
  //desired_capacity = 0

  subnet_ids = ["subnet-0bd585d2c2177c7ee", "subnet-0faad0b6bb7e99d9f"]
  security_group_ids = ["sg-987654"]
  utilize_reserved_instances = true

 %v
 %v
 %v
 %v
}
`

// endregion

// region OceanECS: Instance Types
func TestAccSpotinstOceanECS_InstanceTypesWhitelist(t *testing.T) {
	name := "test-acc-cluster-instance-types-whitelist"
	clusterName := "whitelist-cluster-name"
	resourceName := createOceanECSResourceName(clusterName)

	var cluster aws.ECSCluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanECSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					name:              name,
					clusterName:       clusterName,
					instanceWhitelist: testInstanceTypesWhitelistECSConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.0", "t1.micro"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.1", "m1.small"),
				),
			},
			{
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					name:              name,
					clusterName:       clusterName,
					instanceWhitelist: testInstanceTypesWhitelistECSConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.0", "t1.micro"),
				),
			},
			{
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					name:              name,
					clusterName:       clusterName,
					instanceWhitelist: testInstanceTypesWhitelistECSConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "0"),
				),
			},
		},
	})
}

const testInstanceTypesWhitelistECSConfig_Create = `
 whitelist = ["t1.micro","m1.small"]
`

const testInstanceTypesWhitelistECSConfig_Update = `
 whitelist = ["t1.micro"]
`

const testInstanceTypesWhitelistECSConfig_EmptyFields = `
`

// endregion

// region OceanECS: Launch Specification
func TestAccSpotinstOceanECS_LaunchSpecification(t *testing.T) {
	name := "test-acc-cluster-launch-spec"
	clusterName := "launch-spec-cluster-name"
	resourceName := createOceanECSResourceName(clusterName)

	var cluster aws.ECSCluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanECSDestroy,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					name:        name,
					clusterName: clusterName,
					launchSpec:  testLaunchSpecECSConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-082b5a644766e0e6f"),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "arn:aws:iam::842422002533:instance-profile/ecsInstanceRole"),
					resource.TestCheckResourceAttr(resourceName, "key_pair", "spotinst-labs-oregon"),
					resource.TestCheckResourceAttr(resourceName, "user_data", "IyEvYmluL2Jhc2gKZWNobyBFQ1NfQ0xVU1RFUj1vcmZyb21FbnZpcm9ubWVudF9CYXRjaF84NTJhNjcwYS1hYTczLTNkNWQtOTU3Ni0xNDdhMjZkNDM0MDEgPj4gL2V0Yy9lY3MvZWNzLmNvbmZpZw=="),
					resource.TestCheckResourceAttr(resourceName, "associate_public_ip_address", "false"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.device_name", "/dev/xvda1"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.delete_on_termination", "true"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.encrypted", "false"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.kms_key_id", "kms-key"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.volume_type", "gp2"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.volume_size", "50"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_options.0.http_put_response_hop_limit", "10"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_options.0.http_tokens", "required"),
					resource.TestCheckResourceAttr(resourceName, "use_as_template_only", "true"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					name:        name,
					clusterName: clusterName,
					launchSpec:  testLaunchSpecECSConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-0f2176987ee50226e"),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "arn:aws:iam::842422002533:instance-profile/ecsInstanceRole"),
					resource.TestCheckResourceAttr(resourceName, "key_pair", ""),
					resource.TestCheckResourceAttr(resourceName, "user_data", "IyEvYmluL2Jhc2gKZWNobyBFQ1NfQ0xVU1RFUj1vcmZyb21FbnZpcm9ubWVudF9CYXRjaF84NTJhNjcwYS1hYTczLTNkNWQtOTU3Ni0xNDdhMjZkNDM0MDEgPj4gL2V0Yy9lY3MvZWNzLmNvbmZpZw=="),
					resource.TestCheckResourceAttr(resourceName, "associate_public_ip_address", "true"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.device_name", "/dev/xvda1"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.delete_on_termination", "true"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.dynamic_volume_size.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.dynamic_volume_size.0.base_size", "50"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.dynamic_volume_size.0.resource", "CPU"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.dynamic_volume_size.0.size_per_resource_unit", "20"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.encrypted", "false"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.kms_key_id", "kms-key"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.volume_type", "gp3"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.throughput", "500"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_options.0.http_put_response_hop_limit", "20"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_options.0.http_tokens", "optional"),
					resource.TestCheckResourceAttr(resourceName, "use_as_template_only", "false"),
				),
			},
		},
	})
}

const testLaunchSpecECSConfig_Create = `
// --- LAUNCH SPECIFICATION --------------

 image_id 					 = "ami-082b5a644766e0e6f"
 iam_instance_profile 		 = "arn:aws:iam::842422002533:instance-profile/ecsInstanceRole"
 key_pair 					 = "spotinst-labs-oregon"
 user_data 					 = "IyEvYmluL2Jhc2gKZWNobyBFQ1NfQ0xVU1RFUj1vcmZyb21FbnZpcm9ubWVudF9CYXRjaF84NTJhNjcwYS1hYTczLTNkNWQtOTU3Ni0xNDdhMjZkNDM0MDEgPj4gL2V0Yy9lY3MvZWNzLmNvbmZpZw=="
 associate_public_ip_address = false
use_as_template_only = true

block_device_mappings {
    device_name = "/dev/xvda1"
    ebs {
      delete_on_termination = "true"
      kms_key_id = "kms-key"
      encrypted = "false"
      volume_type = "gp2"
      volume_size = 50
    }
  }

  instance_metadata_options {
    http_tokens                 = "required"
    http_put_response_hop_limit = 10
  }
// ---------------------------------------
`

const testLaunchSpecECSConfig_Update = `
// --- LAUNCH SPECIFICATION --------------

 image_id 					 = "ami-0f2176987ee50226e"
 iam_instance_profile 		 = "arn:aws:iam::842422002533:instance-profile/ecsInstanceRole"
 key_pair 					 = ""
 user_data					 = "IyEvYmluL2Jhc2gKZWNobyBFQ1NfQ0xVU1RFUj1vcmZyb21FbnZpcm9ubWVudF9CYXRjaF84NTJhNjcwYS1hYTczLTNkNWQtOTU3Ni0xNDdhMjZkNDM0MDEgPj4gL2V0Yy9lY3MvZWNzLmNvbmZpZw=="
 associate_public_ip_address = true
use_as_template_only = false

  block_device_mappings {
        device_name = "/dev/xvda1"
        ebs {
          delete_on_termination = "true"
          kms_key_id = "kms-key"
          encrypted = "false"
          volume_type = "gp3"
          throughput = 500
          dynamic_volume_size {
            base_size = 50
            resource = "CPU"
            size_per_resource_unit = 20
          }
        }
      }

  instance_metadata_options {
	http_tokens = "optional"
	http_put_response_hop_limit = 20
  }
// ---------------------------------------
`

// endregion

// region oceanECS: Autoscaler
func TestAccSpotinstOceanECS_Autoscaler(t *testing.T) {
	name := "test-acc-cluster-auto-scaler"
	clusterName := "auto-scaler-cluster-name"
	resourceName := createOceanECSResourceName(clusterName)

	var cluster aws.ECSCluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanECSDestroy,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					clusterName:    clusterName,
					name:           name,
					fieldsToAppend: testScalerConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.cooldown", "180"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.down.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.down.0.max_scale_down_percentage", "10"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.headroom.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.headroom.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.headroom.0.num_of_units", "3"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.is_auto_config", "true"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_memory_gib", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "2"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "2"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.auto_headroom_percentage", "10"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					clusterName:    clusterName,
					name:           name,
					fieldsToAppend: testScalerConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.cooldown", "240"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.down.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.down.0.max_scale_down_percentage", "20"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.headroom.0.cpu_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.headroom.0.memory_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.headroom.0.num_of_units", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.is_auto_config", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_memory_gib", "2"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.auto_headroom_percentage", "20"),
				),
			},
		},
	})
}

const testScalerConfig_Create = `
// --- AUTOSCALER -----------------
autoscaler {
 cooldown = 180
 headroom {
   cpu_per_unit = 1024
   memory_per_unit = 512
   num_of_units = 3
 } 
 down {
   max_scale_down_percentage = 10
 }
 is_auto_config = true
 is_enabled = true
 resource_limits {
   max_vcpu = 2
   max_memory_gib = 1
 }
 auto_headroom_percentage = 10
}
// --------------------------------
`

const testScalerConfig_Update = `
// --- AUTOSCALER -----------------
autoscaler {
 cooldown = 240
 headroom {
   cpu_per_unit = 512
   memory_per_unit = 1024
   num_of_units = 1
 }
 down {
   max_scale_down_percentage = 20
 }
 is_auto_config = false
 is_enabled = false
 resource_limits {
   max_vcpu = 1
   max_memory_gib = 2
 }
 auto_headroom_percentage = 20
}
// --------------------------------
`

// endregion

// region oceanECS: Strategy
func TestAccSpotinstOceanECS_Strategy(t *testing.T) {
	name := "test-acc-cluster-strategy"
	clusterName := "strategy-cluster-name"
	resourceName := createOceanECSResourceName(clusterName)

	var cluster aws.ECSCluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanECSDestroy,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					clusterName:    clusterName,
					name:           name,
					fieldsToAppend: testStrategy_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "120"),
					resource.TestCheckResourceAttr(resourceName, "spot_percentage", "100"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					clusterName:    clusterName,
					name:           name,
					fieldsToAppend: testStrategy_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "240"),
					resource.TestCheckResourceAttr(resourceName, "spot_percentage", "50"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					clusterName:    clusterName,
					name:           name,
					fieldsToAppend: testStrategy_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "0"),
					resource.TestCheckResourceAttr(resourceName, "spot_percentage", "-1"),
				),
			},
		},
	})
}

const testStrategy_Create = `
// --- STRATEGY -----------------
	draining_timeout = 120
	spot_percentage  = 100
// --------------------------------
`

const testStrategy_Update = `
// --- STRATEGY -----------------
	draining_timeout = 240
	spot_percentage  = 50
// --------------------------------
`

const testStrategy_EmptyFields = `
// --- STRATEGY -----------------
	spot_percentage  = null
// --------------------------------
`

// endregion

// region OceanECS: Scheduling
func TestAccSpotinstOceanECS_Scheduling(t *testing.T) {
	name := "test-acc-cluster-scheduling"
	clusterName := "scheduling-cluster-name"
	resourceName := createOceanECSResourceName(clusterName)

	var cluster aws.ECSCluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanECSDestroy,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					clusterName:    clusterName,
					name:           name,
					fieldsToAppend: testScheduling_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.shutdown_hours.0.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.shutdown_hours.0.time_windows.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.shutdown_hours.0.time_windows.0", "Fri:15:30-Sat:15:30"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.tasks.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.tasks.0.cron_expression", "testcron2"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.tasks.0.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.tasks.0.task_type", "clusterRoll"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					clusterName:    clusterName,
					name:           name,
					fieldsToAppend: testScheduling_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.shutdown_hours.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.shutdown_hours.0.time_windows.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.shutdown_hours.0.time_windows.0", "Fri:15:30-Sat:13:30"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.shutdown_hours.0.time_windows.1", "Sun:15:30-Mon:13:30"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.tasks.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.tasks.0.cron_expression", "testcron"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.tasks.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.tasks.0.task_type", "clusterRoll"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					clusterName:    clusterName,
					name:           name,
					fieldsToAppend: testScheduling_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "0"),
				),
			},
		},
	})
}

const testScheduling_Create = `
 // --- Scheduling --------------------
 scheduled_task {
    shutdown_hours {
      is_enabled = false
      time_windows = ["Fri:15:30-Sat:15:30"]
    }
    tasks {
      is_enabled = false
      cron_expression = "testcron2"
      task_type = "clusterRoll"
    }
  }
 // ---------------------------------
`

const testScheduling_Update = `
 // --- Scheduling --------------------
  scheduled_task   {
    shutdown_hours  {
      is_enabled = true
      time_windows = ["Fri:15:30-Sat:13:30","Sun:15:30-Mon:13:30"]
    }
    tasks  {
      is_enabled = true
      cron_expression = "testcron"
      task_type = "clusterRoll"
    }
  }
 // ---------------------------------
`

const testScheduling_EmptyFields = `
 // --- Scheduling --------------------
 // ---------------------------------
`

// endregion

// region oceanECS: Update Policy
func TestAccSpotinstOceanECS_UpdatePolicy(t *testing.T) {
	name := "test-acc-cluster-update-policy"
	clusterName := "update-policy-cluster-name"
	resourceName := createOceanECSResourceName(clusterName)

	var cluster aws.ECSCluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanECSDestroy,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					name:           name,
					clusterName:    clusterName,
					fieldsToAppend: testUpdatePolicyECSClusterConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "update_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.should_roll", "false"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.conditioned_roll", "true"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_size_percentage", "33"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_min_healthy_percentage", "20"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					name:           name,
					clusterName:    clusterName,
					fieldsToAppend: testUpdatePolicyECSClusterConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "update_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.should_roll", "true"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.conditioned_roll", "false"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_size_percentage", "66"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_min_healthy_percentage", "30"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					name:           name,
					clusterName:    clusterName,
					fieldsToAppend: testUpdatePolicyECSClusterConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "update_policy.#", "0"),
				),
			},
		},
	})
}

const testUpdatePolicyECSClusterConfig_Create = `
// --- UPDATE POLICY ----------------
update_policy {
 should_roll = false
 conditioned_roll = true
 roll_config {
   		batch_size_percentage = 33
		batch_min_healthy_percentage = 20
 }
}
// ----------------------------------
`

const testUpdatePolicyECSClusterConfig_Update = `
// --- UPDATE POLICY ----------------
update_policy {
 should_roll = true
 conditioned_roll = false
 roll_config {
		batch_size_percentage = 66
		batch_min_healthy_percentage = 30
 }
}
// ----------------------------------
`

const testUpdatePolicyECSClusterConfig_EmptyFields = `
// --- UPDATE POLICY ----------------
// ----------------------------------
`

// endregion

// region OceanECS: Optimize Images
func TestAccSpotinstOceanECS_OptimizeImages(t *testing.T) {
	name := "test-acc-cluster-optimize-images"
	clusterName := "optimize-images-cluster-name"
	resourceName := createOceanECSResourceName(clusterName)

	var cluster aws.ECSCluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanECSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					name:           name,
					clusterName:    clusterName,
					fieldsToAppend: testOptimizeImagesECSConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "optimize_images.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "optimize_images.0.perform_at", "timeWindow"),
					resource.TestCheckResourceAttr(resourceName, "optimize_images.0.time_windows.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "optimize_images.0.time_windows.0", "Sun:02:00-Sun:12:00"),
					resource.TestCheckResourceAttr(resourceName, "optimize_images.0.time_windows.1", "Sun:05:00-Sun:16:00"),
					resource.TestCheckResourceAttr(resourceName, "optimize_images.0.should_optimize_ecs_ami", "true"),
				),
			},
			{
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					name:           name,
					clusterName:    clusterName,
					fieldsToAppend: testOptimizeImagesECSConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "optimize_images.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "optimize_images.0.perform_at", "never"),
					resource.TestCheckResourceAttr(resourceName, "optimize_images.0.should_optimize_ecs_ami", "false"),
				),
			},
		},
	})
}

const testOptimizeImagesECSConfig_Create = `
// --- OPTIMIZE IMAGES ----------------
  optimize_images {
   perform_at = "timeWindow"
   time_windows = ["Sun:02:00-Sun:12:00","Sun:05:00-Sun:16:00"]
   should_optimize_ecs_ami = true
// ----OPTIMIZE IMAGES ----------------
 }
`

const testOptimizeImagesECSConfig_Update = `
// --- OPTIMIZE IMAGES ----------------
  optimize_images {
   perform_at = "never"
   should_optimize_ecs_ami = false
// ----OPTIMIZE IMAGES ----------------
 }
`

// region OceanECS: Logging
func TestAccSpotinstOceanECS_Logging(t *testing.T) {
	name := "test-acc-cluster-logging"
	clusterName := "logging-cluster-name"
	resourceName := createOceanECSResourceName(clusterName)

	var cluster aws.ECSCluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanECSDestroy,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					clusterName:    clusterName,
					name:           name,
					fieldsToAppend: testLogging_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "logging.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logging.0.export.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logging.0.export.0.s3.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logging.0.export.0.s3.0.id", "di-5fae075b"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					clusterName:    clusterName,
					name:           name,
					fieldsToAppend: testLogging_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "logging.#", "0"),
				),
			},
		},
	})
}

const testLogging_Create = `
 // --- LOGGING -----------------
  logging {
    export {
      s3 { 
		id = "di-5fae075b"
      }
    }
  }
`

const testLogging_EmptyFields = ``

// endregion

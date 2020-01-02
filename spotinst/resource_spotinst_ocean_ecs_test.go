package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
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
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.0", "subnet-79da021e"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.0", "sg-0a8e7b3cd1cfd3d6f"),
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
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.0", "subnet-79da021e"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.1", "subnet-03b7ed5b"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_group_ids.0", "sg-0e9d5f93224747f51"),
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

  subnet_ids         = ["subnet-79da021e"]
  security_group_ids = ["sg-0a8e7b3cd1cfd3d6f"]

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

  subnet_ids         = ["subnet-79da021e","subnet-03b7ed5b"]
  security_group_ids = ["sg-0e9d5f93224747f51"]

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

// ---------------------------------------
`

const testLaunchSpecECSConfig_Update = `
// --- LAUNCH SPECIFICATION --------------

 image_id 					 = "ami-0f2176987ee50226e"
 iam_instance_profile 		 = "arn:aws:iam::842422002533:instance-profile/ecsInstanceRole"
 key_pair 					 = ""
 user_data					 = "IyEvYmluL2Jhc2gKZWNobyBFQ1NfQ0xVU1RFUj1vcmZyb21FbnZpcm9ubWVudF9CYXRjaF84NTJhNjcwYS1hYTczLTNkNWQtOTU3Ni0xNDdhMjZkNDM0MDEgPj4gL2V0Yy9lY3MvZWNzLmNvbmZpZw=="
 associate_public_ip_address = true

// ---------------------------------------
`

// endregion

// region oceanECS: Autoscaler
func TestAccSpotinstoceanECS_Autoscaler(t *testing.T) {
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
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "1"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanECSTerraform(&ECSClusterConfigMetadata{
					clusterName:    clusterName,
					name:           name,
					fieldsToAppend: testScalerConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanECSExists(&cluster, resourceName),
					testCheckOceanECSAttributes(&cluster, name),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.down.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.headroom.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "0"),
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
}
// --------------------------------
`

const testScalerConfig_EmptyFields = `
// --- AUTOSCALER -----------------
autoscaler {
 is_enabled = false
 is_auto_config = false
 cooldown = 300
}
// --------------------------------
`

// endregion

// region oceanECS: Strategy
func TestAccSpotinstoceanECS_Strategy(t *testing.T) {
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
				),
			},
		},
	})
}

const testStrategy_Create = `
// --- STRATEGY -----------------
	draining_timeout = 120
// --------------------------------
`

const testStrategy_Update = `
// --- AUTOSCALER -----------------
	draining_timeout = 240
// --------------------------------
`

const testStrategy_EmptyFields = `
// --- STRATEGY -----------------

// --------------------------------
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
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_size_percentage", "33"),
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
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_size_percentage", "66"),
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
 roll_config {
   batch_size_percentage = 33
 }
}
// ----------------------------------
`

const testUpdatePolicyECSClusterConfig_Update = `
// --- UPDATE POLICY ----------------
update_policy {
 should_roll = true
 roll_config {
   batch_size_percentage = 66
 }
}
// ----------------------------------
`

const testUpdatePolicyECSClusterConfig_EmptyFields = `
// --- UPDATE POLICY ----------------
// ----------------------------------
`

// endregion

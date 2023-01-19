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
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_aws_launch_configuration"
)

func init() {
	resource.AddTestSweepers("spotinst_ocean_aws", &resource.Sweeper{
		Name: "spotinst_ocean_aws",
		F:    testSweepOceanAWS,
	})
}

func testSweepOceanAWS(region string) error {
	client, err := getProviderClient("aws")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).ocean.CloudProviderAWS()
	input := &aws.ListClustersInput{}
	if resp, err := conn.ListClusters(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of clusters to sweep")
	} else {
		if len(resp.Clusters) == 0 {
			log.Printf("[INFO] No clusters to sweep")
		}
		for _, cluster := range resp.Clusters {
			if strings.Contains(spotinst.StringValue(cluster.Name), "test-acc-") {
				if _, err := conn.DeleteCluster(context.Background(), &aws.DeleteClusterInput{ClusterID: cluster.ID}); err != nil {
					return fmt.Errorf("unable to delete cluster %v in sweep", spotinst.StringValue(cluster.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(cluster.ID))
				}
			}
		}
	}
	return nil
}

func createOceanAWSResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanAWSResourceName), name)
}

func testOceanAWSDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OceanAWSResourceName) {
			continue
		}
		input := &aws.ReadClusterInput{ClusterID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAWS().ReadCluster(context.Background(), input)
		if err == nil && resp != nil && resp.Cluster != nil {
			return fmt.Errorf("cluster still exists")
		}
	}
	return nil
}

func testCheckOceanAWSAttributes(cluster *aws.Cluster, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(cluster.Name) != expectedName {
			return fmt.Errorf("bad content: %v", cluster.Name)
		}
		return nil
	}
}

func testCheckOceanAWSExists(cluster *aws.Cluster, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &aws.ReadClusterInput{ClusterID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAWS().ReadCluster(context.Background(), input)
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

type ClusterConfigMetadata struct {
	provider             string
	clusterName          string
	controllerClusterID  string
	instanceWhitelist    string
	launchConfig         string
	strategy             string
	variables            string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createOceanAWSTerraform(ccm *ClusterConfigMetadata) string {
	if ccm == nil {
		return ""
	}

	if ccm.provider == "" {
		ccm.provider = "aws"
	}

	if ccm.launchConfig == "" {
		ccm.launchConfig = testLaunchConfigAWSConfig_Create
	}

	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`

	if ccm.updateBaselineFields {
		format := testBaselineAWSConfig_Update
		template += fmt.Sprintf(format,
			ccm.clusterName,
			ccm.provider,
			ccm.clusterName,
			ccm.controllerClusterID,
			ccm.instanceWhitelist,
			ccm.launchConfig,
			ccm.strategy,
			ccm.fieldsToAppend,
		)
	} else {
		format := testBaselineAWSConfig_Create
		template += fmt.Sprintf(format,
			ccm.clusterName,
			ccm.provider,
			ccm.clusterName,
			ccm.controllerClusterID,
			ccm.instanceWhitelist,
			ccm.launchConfig,
			ccm.strategy,
			ccm.fieldsToAppend,
		)
	}

	if ccm.variables != "" {
		template = ccm.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", ccm.clusterName, template)
	return template
}

// region OceanAWS: Baseline
func TestAccSpotinstOceanAWS_Baseline(t *testing.T) {
	clusterName := "test-acc-cluster-baseline"
	controllerClusterID := "baseline-controller-id"
	resourceName := createOceanAWSResourceName(clusterName)

	var cluster aws.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "max_size", "1000"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "1"),
				),
			},
			{
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:          clusterName,
					controllerClusterID:  controllerClusterID,
					updateBaselineFields: true}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "max_size", "10"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "1"),
				),
			},
		},
	})
}

const testBaselineAWSConfig_Create = `
resource "` + string(commons.OceanAWSResourceName) + `" "%v" {
  provider = "%v"  
  
  name = "%v"
  controller_id = "%v"
  region = "us-west-2"

 max_size         = 1000
  min_size         = 0
  desired_capacity = 1

  subnet_ids      = ["subnet-0faad0b6bb7e99d9f"]

 %v
 %v
 %v
 %v
}
`

const testBaselineAWSConfig_Update = `
resource "` + string(commons.OceanAWSResourceName) + `" "%v" {
  provider = "%v"

  name = "%v"
  controller_id = "%v"
  region = "us-west-2"

  max_size         = 10
  min_size         = 0
  desired_capacity = 1

  subnet_ids      = ["subnet-0bd585d2c2177c7ee"]

 %v
 %v
 %v
 %v
}
`

// endregion

// region OceanAWS: Instance Types Whitelist
func TestAccSpotinstOceanAWS_InstanceTypesLists(t *testing.T) {
	clusterName := "test-acc-cluster-instance-types-whitelist"
	controllerClusterID := "test-acc-cluster-baseline"
	resourceName := createOceanAWSResourceName(clusterName)

	var cluster aws.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					instanceWhitelist:   testInstanceTypesWhitelistAWSConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.0", "t1.micro"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.1", "m1.small"),
				),
			},
			{
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					instanceWhitelist:   testInstanceTypesWhitelistAWSConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "whitelist.0", "t1.micro"),
				),
			},
			{
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					instanceWhitelist:   testInstanceTypesWhitelistAWSConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "blacklist.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "blacklist.0", "t1.micro"),
					resource.TestCheckResourceAttr(resourceName, "blacklist.1", "m1.small"),
				),
			},
			{
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					instanceWhitelist:   testInstanceTypesBlacklistAWSConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "blacklist.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "blacklist.0", "t1.micro"),
				),
			},
			{
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					instanceWhitelist:   testInstanceTypesBlacklistAWSConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "blacklist.#", "0"),
				),
			},
		},
	})
}

const testInstanceTypesWhitelistAWSConfig_Create = `
  whitelist = ["t1.micro", "m1.small"]
`

const testInstanceTypesWhitelistAWSConfig_Update = `
  whitelist = ["t1.micro"]
`

const testInstanceTypesWhitelistAWSConfig_EmptyFields = `
blacklist = ["t1.micro", "m1.small"] 
`

const testInstanceTypesBlacklistAWSConfig_Update = `
blacklist = ["t1.micro"] 
`

const testInstanceTypesBlacklistAWSConfig_EmptyFields = `

`

// endregion

// region OceanAWS: Launch Configuration
func TestAccSpotinstOceanAWS_LaunchConfiguration(t *testing.T) {
	clusterName := "test-acc-luster-launch-configuration"
	controllerClusterID := "launch-config-cluster-id"
	resourceName := createOceanAWSResourceName(clusterName)

	var cluster aws.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSDestroy,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					launchConfig:        testLaunchConfigAWSConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-05f840082fe2dcac2"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-065c82e9ff8b192a1"),
					resource.TestCheckResourceAttr(resourceName, "associate_public_ip_address", "false"),
					//resource.TestCheckResourceAttr(resourceName, "key_name", "my-key.ssh"),
					resource.TestCheckResourceAttr(resourceName, "user_data", ocean_aws_launch_configuration.Base64StateFunc("echo hello world")),
					//resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "iam-profile"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.key", "fakeKey"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.value", "fakeValue"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.0.arn", "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.0.type", "TARGET_GROUP"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.1.name", "AntonK"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.1.type", "CLASSIC"),
					resource.TestCheckResourceAttr(resourceName, "root_volume_size", "20"),
					resource.TestCheckResourceAttr(resourceName, "monitoring", "true"),
					resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "true"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_options.0.http_put_response_hop_limit", "10"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_options.0.http_tokens", "required"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					launchConfig:        testLaunchConfigAWSConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-05f840082fe2dcac2"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-065c82e9ff8b192a1"),
					resource.TestCheckResourceAttr(resourceName, "associate_public_ip_address", "true"),
					//resource.TestCheckResourceAttr(resourceName, "key_name", "my-key-updated.ssh"),
					resource.TestCheckResourceAttr(resourceName, "user_data", ocean_aws_launch_configuration.Base64StateFunc("echo hello world updated")),
					//resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "iam-profile updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.key", "fakeKeyUpdated"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.value", "fakeValueUpdated"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.0.arn", "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.0.type", "TARGET_GROUP"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.1.name", "AntonK"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.1.type", "CLASSIC"),
					resource.TestCheckResourceAttr(resourceName, "root_volume_size", "24"),
					resource.TestCheckResourceAttr(resourceName, "monitoring", "false"),
					resource.TestCheckResourceAttr(resourceName, "ebs_optimized", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_as_template_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_options.0.http_put_response_hop_limit", "20"),
					resource.TestCheckResourceAttr(resourceName, "instance_metadata_options.0.http_tokens", "optional"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					launchConfig:        testLaunchConfigAWSConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-05f840082fe2dcac2"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-065c82e9ff8b192a1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "0"),
				),
			},
		},
	})
}

const testLaunchConfigAWSConfig_Create = `
 // --- LAUNCH CONFIGURATION --------------
  image_id                    = "ami-05f840082fe2dcac2"
  security_groups             = ["sg-065c82e9ff8b192a1"]
  //key_name                  = "my-key.ssh"
  user_data                   = "echo hello world"
  //iam_instance_profile      = "iam-profile"
  associate_public_ip_address = false
  root_volume_size            = 20
  monitoring                  = true
  ebs_optimized               = true

  instance_metadata_options {
    http_tokens                 = "required"
    http_put_response_hop_limit = 10
  }

  load_balancers {
     arn  = "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456"
      type = "TARGET_GROUP"
    }

	load_balancers {
      name = "AntonK"
      type = "CLASSIC"
    }

  tags {
    key   = "fakeKey"
    value = "fakeValue"
  }
 // ---------------------------------------
`

const testLaunchConfigAWSConfig_Update = `
 // --- LAUNCH CONFIGURATION --------------
  image_id                    = "ami-05f840082fe2dcac2"
  security_groups             = ["sg-065c82e9ff8b192a1"]
  //key_name                  = "my-key-updated.ssh"
  user_data                   = "echo hello world updated"
  //iam_instance_profile      = "iam-profile updated"
  associate_public_ip_address = true
  root_volume_size            = 24
  monitoring                  = false
  ebs_optimized               = false
  use_as_template_only        = false
  instance_metadata_options {
	  http_tokens = "optional"
      http_put_response_hop_limit = 20
  }

  load_balancers {
      arn  = "arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/testTargetGroup/1234567890123456"
      type = "TARGET_GROUP"
    }
	load_balancers {
		name = "AntonK"
		type = "CLASSIC"
	}

  tags {
    key   = "fakeKeyUpdated"
    value = "fakeValueUpdated"
  }
 // ---------------------------------------
`

const testLaunchConfigAWSConfig_EmptyFields = `
 // --- LAUNCH CONFIGURATION --------------
  image_id        = "ami-05f840082fe2dcac2"
  security_groups = ["sg-065c82e9ff8b192a1"]
 // ---------------------------------------
`

// endregion

// region OceanAWS: Strategy
func TestAccSpotinstOceanAWS_Strategy(t *testing.T) {
	clusterName := "test-acc-cluster-strategy"
	controllerClusterID := "strategy-controller-id"
	resourceName := createOceanAWSResourceName(clusterName)

	var cluster aws.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					strategy:            testStrategyConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "fallback_to_ondemand", "true"),
					resource.TestCheckResourceAttr(resourceName, "spot_percentage", "100"),
					resource.TestCheckResourceAttr(resourceName, "utilize_reserved_instances", "false"),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "120"),
					resource.TestCheckResourceAttr(resourceName, "grace_period", "50"),
				),
			},
			{
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					strategy:            testStrategyConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "fallback_to_ondemand", "false"),
					resource.TestCheckResourceAttr(resourceName, "spot_percentage", "50"),
					resource.TestCheckResourceAttr(resourceName, "utilize_reserved_instances", "true"),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "240"),
					resource.TestCheckResourceAttr(resourceName, "grace_period", "100"),
				),
			},
			{
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{

					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					strategy:            testStrategyConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "fallback_to_ondemand", "true"),
					resource.TestCheckResourceAttr(resourceName, "spot_percentage", "-1"),
					resource.TestCheckResourceAttr(resourceName, "utilize_reserved_instances", "true"),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "0"),
				),
			},
		},
	})
}

const testStrategyConfig_Create = `
 // --- STRATEGY --------------------
 fallback_to_ondemand       = true
 spot_percentage            = 100
 utilize_reserved_instances = false
 draining_timeout			= 120
 grace_period = 50
 // ---------------------------------
`

const testStrategyConfig_Update = `
 // --- STRATEGY --------------------
 fallback_to_ondemand       = false
 spot_percentage            = 50
 utilize_reserved_instances = true
 draining_timeout			= 240
 grace_period = 100
 // ---------------------------------
`

const testStrategyConfig_EmptyFields = `
 // --- STRATEGY --------------------
 spot_percentage = null
 // ---------------------------------
`

// endregion

// region OceanAWS: Scheduling
func TestAccSpotinstOceanAWS_Scheduling(t *testing.T) {
	clusterName := "test-acc-cluster-scheduling"
	controllerClusterID := "scheduling-controller-id"
	resourceName := createOceanAWSResourceName(clusterName)

	var cluster aws.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					strategy:            testSchedulingConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
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
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					strategy:            testSchedulingConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.shutdown_hours.0.is_enabled", "false"),
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
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					strategy:            testSchedulingConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "0"),
				),
			},
		},
	})
}

const testSchedulingConfig_Create = `
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

const testSchedulingConfig_Update = `
 // --- Scheduling --------------------
  scheduled_task   {
    shutdown_hours  {
      is_enabled = false
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

const testSchedulingConfig_EmptyFields = `
 // --- Scheduling --------------------
 // ---------------------------------
`

// endregion

// region OceanAWS: Autoscaler
func TestAccSpotinstOceanAWS_Autoscaler(t *testing.T) {
	clusterName := "test-acc-cluster-autoscaler"
	controllerClusterID := "autoscaler-controller-id"
	resourceName := createOceanAWSResourceName(clusterName)

	var cluster aws.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSDestroy,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					fieldsToAppend:      testScalingConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_cooldown", "300"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.0.evaluation_periods", "300"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.0.max_scale_down_percentage", "50.5"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.gpu_per_unit", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.num_of_units", "2"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_auto_config", "true"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.enable_automatic_and_manual_headroom", "true"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.auto_headroom_percentage", "100"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_memory_gib", "20"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.extended_resource_definitions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.extended_resource_definitions.0", "erd-cb74ca43"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					fieldsToAppend:      testScalingConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_cooldown", "600"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.0.evaluation_periods", "600"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.0.max_scale_down_percentage", "10"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.cpu_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.gpu_per_unit", "2"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.memory_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.num_of_units", "4"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_auto_config", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.enable_automatic_and_manual_headroom", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.auto_headroom_percentage", "150"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_memory_gib", "30"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "512"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.extended_resource_definitions.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.extended_resource_definitions.0", "erd-cb74ca43"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.extended_resource_definitions.1", "erd-ced684ab"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					fieldsToAppend:      testScalingConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_cooldown", "300"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_down.0.evaluation_periods", "300"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.num_of_units", "2"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_auto_config", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.enable_automatic_and_manual_headroom", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.auto_headroom_percentage", "0"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_memory_gib", "20"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.extended_resource_definitions.#", "0"),
				),
			},
		},
	})
}

const testScalingConfig_Create = `
 // --- AUTOSCALER -----------------
 autoscaler {
    autoscale_is_enabled     = true
    autoscale_is_auto_config = true
	enable_automatic_and_manual_headroom = true
	auto_headroom_percentage = 100
    autoscale_cooldown       = 300

    autoscale_headroom {
      cpu_per_unit    = 1024
      gpu_per_unit    = 1
      memory_per_unit = 512
      num_of_units    = 2
    }

    autoscale_down {
      evaluation_periods = 300
      max_scale_down_percentage = 50.5
    }

    resource_limits {
      max_vcpu       = 1024
      max_memory_gib = 20
    }

	extended_resource_definitions = ["erd-cb74ca43"]
 }
 // --------------------------------

`

const testScalingConfig_Update = `
 // --- AUTOSCALER -----------------
 autoscaler {
    autoscale_is_enabled     = false
    autoscale_is_auto_config = false
	enable_automatic_and_manual_headroom = false
	auto_headroom_percentage = 150
    autoscale_cooldown       = 600

    autoscale_headroom {
      cpu_per_unit    = 512
      gpu_per_unit    = 2
      memory_per_unit = 1024
      num_of_units    = 4
    }

    autoscale_down {
      evaluation_periods = 600
      max_scale_down_percentage = 10

    }

    resource_limits {
      max_vcpu       = 512
      max_memory_gib = 30
    }

	extended_resource_definitions = ["erd-cb74ca43", "erd-ced684ab"]
 }
 // --------------------------------
`

const testScalingConfig_EmptyFields = `
 // --- AUTOSCALER -----------------
 autoscaler {
    autoscale_is_enabled = false
    autoscale_is_auto_config = false
	enable_automatic_and_manual_headroom = false
	auto_headroom_percentage = 0
    autoscale_cooldown = 300

    autoscale_headroom {
      cpu_per_unit = 1024
      memory_per_unit = 512
      num_of_units = 2
    }

    autoscale_down {
      evaluation_periods = 300
    }

    resource_limits {
      max_vcpu   = 1024
      max_memory_gib = 20
    }

  	extended_resource_definitions = null
 }
 // --------------------------------
`

// endregion

// region OceanAWS: Update Policy

func TestAccSpotinstOceanAWS_UpdatePolicy(t *testing.T) {
	clusterName := "test-acc-cluster-update-policy"
	controllerClusterID := "update-policy-controller-id"
	resourceName := createOceanAWSResourceName(clusterName)

	var cluster aws.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSDestroy,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					fieldsToAppend:      testUpdatePolicyAWSClusterConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "update_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.should_roll", "false"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.conditioned_roll", "true"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_size_percentage", "33"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_min_healthy_percentage", "20"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.respect_pdb", "false"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					fieldsToAppend:      testUpdatePolicyAWSClusterConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "update_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.should_roll", "true"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.conditioned_roll", "false"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_size_percentage", "66"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.batch_min_healthy_percentage", "30"),
					resource.TestCheckResourceAttr(resourceName, "update_policy.0.roll_config.0.respect_pdb", "true"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					fieldsToAppend:      testUpdatePolicyAWSClusterConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "update_policy.#", "0"),
				),
			},
		},
	})
}

const testUpdatePolicyAWSClusterConfig_Create = `
 spot_percentage = 100

 // --- UPDATE POLICY ----------------
  update_policy {
    should_roll = false
	conditioned_roll = true

    roll_config {
      	batch_size_percentage = 33
		batch_min_healthy_percentage = 20
		respect_pdb = false
    }
  }
 // ----------------------------------
`

const testUpdatePolicyAWSClusterConfig_Update = `
 spot_percentage = 50

 // --- UPDATE POLICY ----------------
  update_policy {
    should_roll = true
	conditioned_roll = false

    roll_config {
      	batch_size_percentage = 66
		batch_min_healthy_percentage = 30
		respect_pdb = true
    }
  }
 // ----------------------------------
`

const testUpdatePolicyAWSClusterConfig_EmptyFields = `
 spot_percentage = 0
 // --- UPDATE POLICY ----------------
 // ----------------------------------
`

// endregion

// region OceanAWS: Baseline
func TestAccSpotinstOceanAWS_Logging(t *testing.T) {
	clusterName := "test-acc-cluster-logging"
	controllerClusterID := "logging-controller-id"
	resourceName := createOceanAWSResourceName(clusterName)

	var cluster aws.Cluster
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					fieldsToAppend:      testLoggingAWSConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "logging.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logging.0.export.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logging.0.export.0.s3.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "logging.0.export.0.s3.0.id", "di-5fae075b"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createOceanAWSTerraform(&ClusterConfigMetadata{
					clusterName:         clusterName,
					controllerClusterID: controllerClusterID,
					fieldsToAppend:      testLoggingAWSConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSExists(&cluster, resourceName),
					testCheckOceanAWSAttributes(&cluster, clusterName),
					resource.TestCheckResourceAttr(resourceName, "logging.#", "0"),
				),
			},
		},
	})
}

const testLoggingAWSConfig_Create = `
 // --- LOGGING -----------------
  logging {
    export {
      s3 { 
		id = "di-5fae075b"
      }
    }
  }
`

const testLoggingAWSConfig_EmptyFields = `

`

// endregion

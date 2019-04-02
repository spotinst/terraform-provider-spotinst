package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/ocean_aws_launch_configuration"
	"log"
	"strings"
	"testing"
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

func createOceanAWSTerraform(gcm *ClusterConfigMetadata) string {
	if gcm == nil {
		return ""
	}

	if gcm.provider == "" {
		gcm.provider = "aws"
	}

	if gcm.launchConfig == "" {
		gcm.launchConfig = testLaunchConfigAWSConfig_Create
	}

	//template := ""
	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`

	if gcm.updateBaselineFields {
		format := testBaselineAWSConfig_Update
		template += fmt.Sprintf(format,
			gcm.clusterName,
			gcm.provider,
			gcm.clusterName,
			gcm.controllerClusterID,
			gcm.instanceWhitelist,
			gcm.launchConfig,
			gcm.strategy,
			gcm.fieldsToAppend,
		)
	} else {
		format := testBaselineAWSConfig_Create
		template += fmt.Sprintf(format,
			gcm.clusterName,
			gcm.provider,
			gcm.clusterName,
			gcm.controllerClusterID,
			gcm.instanceWhitelist,
			gcm.launchConfig,
			gcm.strategy,
			gcm.fieldsToAppend,
		)
	}

	if gcm.variables != "" {
		template = gcm.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", gcm.clusterName, template)
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
					resource.TestCheckResourceAttr(resourceName, "min_size", "2"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "2"),
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

  //max_size         = 0
  //min_size         = 0
  //desired_capacity = 0

  subnet_ids      = ["subnet-09d9755d9bdeca3c5"]

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
  min_size         = 2
  desired_capacity = 2

  subnet_ids      = ["subnet-09d9755d9bdeca3c5"]

 %v
 %v
 %v
 %v
}
`

// endregion

// region OceanAWS: Instance Types Whitelist
func TestAccSpotinstOceanAWS_InstanceTypesWhitelist(t *testing.T) {
	clusterName := "test-acc-cluster-instance-types-whitelist"
	controllerClusterID := "whitelist-controller-id"
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
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-79826301"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-042d658b3ee907848"),
					resource.TestCheckResourceAttr(resourceName, "associate_public_ip_address", "false"),
					//resource.TestCheckResourceAttr(resourceName, "key_name", "my-key.ssh"),
					resource.TestCheckResourceAttr(resourceName, "user_data", ocean_aws_launch_configuration.Base64StateFunc("echo hello world")),
					//resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "iam-profile"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.1116605596.key", "fakeKey"),
					resource.TestCheckResourceAttr(resourceName, "tags.1116605596.value", "fakeValue"),

					resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.0.arn", "arn:aws:elasticloadbalancing:us-west-2:842422002533:loadbalancer/app/AntonK/8db573b63a46bfb2"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.0.type", "TARGET_GROUP"),

					resource.TestCheckResourceAttr(resourceName, "load_balancers.1.name", "AntonK"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.1.type", "CLASSIC"),
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
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-79826301"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-042d658b3ee907848"),
					resource.TestCheckResourceAttr(resourceName, "associate_public_ip_address", "true"),
					//resource.TestCheckResourceAttr(resourceName, "key_name", "my-key-updated.ssh"),
					resource.TestCheckResourceAttr(resourceName, "user_data", ocean_aws_launch_configuration.Base64StateFunc("echo hello world updated")),
					//resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "iam-profile updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.3418058476.key", "fakeKeyUpdated"),
					resource.TestCheckResourceAttr(resourceName, "tags.3418058476.value", "fakeValueUpdated"),

					resource.TestCheckResourceAttr(resourceName, "load_balancers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.0.arn", "arn:aws:elasticloadbalancing:us-west-2:842422002533:loadbalancer/app/AntonK/1234567890"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.0.type", "TARGET_GROUP"),

					resource.TestCheckResourceAttr(resourceName, "load_balancers.1.name", "AntonK"),
					resource.TestCheckResourceAttr(resourceName, "load_balancers.1.type", "CLASSIC"),
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
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-79826301"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-042d658b3ee907848"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "0"),
				),
			},
		},
	})
}

const testLaunchConfigAWSConfig_Create = `
 // --- LAUNCH CONFIGURATION --------------
  image_id             = "ami-79826301"
  security_groups      = ["sg-042d658b3ee907848"]
  //key_name             = "my-key.ssh"
  user_data            = "echo hello world"
  //iam_instance_profile = "iam-profile"
  associate_public_ip_address = false
  load_balancers = [
    {
      arn = "arn:aws:elasticloadbalancing:us-west-2:842422002533:loadbalancer/app/AntonK/8db573b63a46bfb2"
      type = "TARGET_GROUP"
    },
    {
      name = "AntonK"
      type = "CLASSIC"
    }
  ]

  tags = [{
    key   = "fakeKey"
    value = "fakeValue"
  }]
 // ---------------------------------------
`

const testLaunchConfigAWSConfig_Update = `
 // --- LAUNCH CONFIGURATION --------------
  image_id             = "ami-79826301"
  security_groups      = ["sg-042d658b3ee907848"]
  //key_name             = "my-key-updated.ssh"
  user_data            = "echo hello world updated"
  //iam_instance_profile = "iam-profile updated"
  associate_public_ip_address = true
  load_balancers = [
    {
      arn = "arn:aws:elasticloadbalancing:us-west-2:842422002533:loadbalancer/app/AntonK/1234567890"
      type = "TARGET_GROUP"
    },
    {
      name = "AntonK"
      type = "CLASSIC"
    }
  ]

  tags = [{
    key   = "fakeKeyUpdated"
    value = "fakeValueUpdated"
  }]
 // ---------------------------------------
`

const testLaunchConfigAWSConfig_EmptyFields = `
 // --- LAUNCH CONFIGURATION --------------
  image_id        = "ami-79826301"
  security_groups = ["sg-042d658b3ee907848"]
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
					resource.TestCheckResourceAttr(resourceName, "spot_percentage", "0"),
					resource.TestCheckResourceAttr(resourceName, "utilize_reserved_instances", "false"),
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
 // ---------------------------------
`

const testStrategyConfig_Update = `
 // --- STRATEGY --------------------
 fallback_to_ondemand       = false
 spot_percentage            = 50
 utilize_reserved_instances = true
 // ---------------------------------
`

const testStrategyConfig_EmptyFields = `
 // --- STRATEGY --------------------
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
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.gpu_per_unit", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.num_of_units", "2"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_auto_config", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_memory_gib", "20"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "1024"),
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
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.cpu_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.gpu_per_unit", "2"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.memory_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_headroom.0.num_of_units", "4"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_auto_config", "true"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.autoscale_is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_memory_gib", "30"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "512"),
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
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_memory_gib", "20"),
					resource.TestCheckResourceAttr(resourceName, "autoscaler.0.resource_limits.0.max_vcpu", "1024"),
				),
			},
		},
	})
}

const testScalingConfig_Create = `
 // --- AUTOSCALER -----------------
 autoscaler = {
    autoscale_is_enabled     = false
    autoscale_is_auto_config = false
    autoscale_cooldown       = 300

    autoscale_headroom = {
      cpu_per_unit    = 1024
      gpu_per_unit    = 1
      memory_per_unit = 512
      num_of_units    = 2
    }

    autoscale_down = {
      evaluation_periods = 300
    }

    resource_limits = {
      max_vcpu       = 1024
      max_memory_gib = 20
    }
 }
 // --------------------------------

`

const testScalingConfig_Update = `
 // --- AUTOSCALER -----------------
 autoscaler = {
    autoscale_is_enabled     = true
    autoscale_is_auto_config = true
    autoscale_cooldown       = 600

    autoscale_headroom = {
      cpu_per_unit    = 512
      gpu_per_unit    = 2
      memory_per_unit = 1024
      num_of_units    = 4
    }

    autoscale_down = {
      evaluation_periods = 600
    }

    resource_limits = {
      max_vcpu       = 512
      max_memory_gib = 30
    }
 }
 // --------------------------------
`

const testScalingConfig_EmptyFields = `
 // --- AUTOSCALER -----------------
 autoscaler = {
    autoscale_is_enabled = false
    autoscale_is_auto_config = false
    autoscale_cooldown = 300

    autoscale_headroom = {
      cpu_per_unit = 1024
      memory_per_unit = 512
      num_of_units = 2
    }

    autoscale_down = {
      evaluation_periods = 300
    }

    resource_limits = {
      max_vcpu   = 1024
      max_memory_gib = 20
    }
 }
 // --------------------------------
`

// endregion

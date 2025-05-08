package spotinst

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/elastigroup_aws_launch_configuration"
)

//func init() {
//	resource.AddTestSweepers("spotinst_ocean_aws_launch_spec", &resource.Sweeper{
//		Name: "spotinst_ocean_aws_launch_spec",
//		F:    testSweepOceanAWSLaunchSpec,
//	})
//}
//
//func testSweepOceanAWSLaunchSpec(region string) error {
//	client, err := getProviderClient("aws")
//	if err != nil {
//		return fmt.Errorf("error getting client: %v", err)
//	}
//
//	conn := client.(*Client).ocean.CloudProviderAWS()
//	input := &aws.ListLaunchSpecsInput{}
//	if resp, err := conn.ListLaunchSpecs(context.Background(), input); err != nil {
//		return fmt.Errorf("error getting list of launch specs to sweep")
//	} else {
//		if len(resp.LaunchSpecs) == 0 {
//			log.Printf("[INFO] No launch specs to sweep")
//		}
//		for _, launchSpec := range resp.LaunchSpecs {
//			if strings.Contains(spotinst.StringValue(launchSpec.<WHAT>), "test-acc-") {
//				if _, err := conn.DeleteLaunchSpec(context.Background(), &aws.DeleteLaunchSpecInput{LaunchSpecID: launchSpec.ID}); err != nil {
//					return fmt.Errorf("unable to delete launch spec %v in sweep", spotinst.StringValue(launchSpec.ID))
//				} else {
//					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(launchSpec.ID))
//				}
//			}
//		}
//	}
//	return nil
//}

func createOceanAWSLaunchSpecResourceOceanID(oceanID string) string {
	return fmt.Sprintf("%v.%v", string(commons.OceanAWSLaunchSpecResourceName), oceanID)
}

func testOceanAWSLaunchSpecDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.OceanAWSLaunchSpecResourceName) {
			continue
		}
		input := &aws.ReadLaunchSpecInput{LaunchSpecID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAWS().ReadLaunchSpec(context.Background(), input)
		if err == nil && resp != nil && resp.LaunchSpec != nil {
			return fmt.Errorf("launchSpec still exists")
		}
	}
	return nil
}

func testCheckOceanAWSLaunchSpecAttributes(launchSpec *aws.LaunchSpec, expectedID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(launchSpec.OceanID) != expectedID {
			return fmt.Errorf("bad content: %v", launchSpec.OceanID)
		}
		return nil
	}
}

func testCheckOceanAWSLaunchSpecExists(launchSpec *aws.LaunchSpec, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &aws.ReadLaunchSpecInput{LaunchSpecID: spotinst.String(rs.Primary.ID)}
		resp, err := client.ocean.CloudProviderAWS().ReadLaunchSpec(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.LaunchSpec.ID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("LaunchSpec not found: %+v,\n %+v\n", resp.LaunchSpec, rs.Primary.Attributes)
		}
		*launchSpec = *resp.LaunchSpec
		return nil
	}
}

type LaunchSpecConfigMetadata struct {
	provider             string
	oceanID              string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createOceanAWSLaunchSpecTerraform(lscm *LaunchSpecConfigMetadata, formatToUse string) string {
	if lscm == nil {
		return ""
	}

	if lscm.provider == "" {
		lscm.provider = "aws"
	}

	template :=
		`provider "aws" {
	 token   = "fake"
	 account = "fake"
	}
	`

	format := formatToUse

	if lscm.updateBaselineFields {
		template += fmt.Sprintf(format,
			lscm.oceanID,
			lscm.provider,
			lscm.oceanID,
			lscm.fieldsToAppend,
		)
	} else {
		template += fmt.Sprintf(format,
			lscm.oceanID,
			lscm.provider,
			lscm.oceanID,
			lscm.fieldsToAppend,
		)
	}

	log.Printf("Terraform LaunchSpec template:\n%v", template)
	return template
}

// region OceanAWSLaunchSpec: Baseline
func TestAccSpotinstOceanAWSLaunchSpec_Baseline(t *testing.T) {
	oceanID := "o-8b34732f"
	resourceName := createOceanAWSLaunchSpecResourceOceanID(oceanID)

	var launchSpec aws.LaunchSpec
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSLaunchSpecDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID: oceanID,
				}, testBaselineOceanAWSLaunchSpecConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "image_id", ""),
					resource.TestCheckResourceAttr(resourceName, "instance_types.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "instance_types.0", "m3.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types.1", "m4.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types.2", "m5.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "preferred_spot_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "preferred_spot_types.0", "m3.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "preferred_od_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "preferred_od_types.0", "m4.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-0041bd3fd6aa2ee3c"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_aws_launch_configuration.Base64StateFunc("hello world")),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "test"),
					resource.TestCheckResourceAttr(resourceName, "labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "labels.0.key", "label key"),
					resource.TestCheckResourceAttr(resourceName, "labels.0.value", "label value"),
					resource.TestCheckResourceAttr(resourceName, "taints.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.key", "taint key"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.value", "taint value"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.effect", "NoSchedule"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.0", "subnet-4333093a"),
					resource.TestCheckResourceAttr(resourceName, "root_volume_size", "20"),
					resource.TestCheckResourceAttr(resourceName, "associate_public_ip_address", "true"),
					resource.TestCheckResourceAttr(resourceName, "restrict_scale_down", "true"),
					resource.TestCheckResourceAttr(resourceName, "reserved_enis", "1"),
					resource.TestCheckResourceAttr(resourceName, "startup_taints.0.key", "startuptaint-key"),
					resource.TestCheckResourceAttr(resourceName, "startup_taints.0.value", "startuptaint-value"),
					resource.TestCheckResourceAttr(resourceName, "startup_taints.0.effect", "NoSchedule"),
				),
			},
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}, testBaselineOceanAWSLaunchSpecConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "image_id", "ami-05a68f290aa68e8f0"),
					resource.TestCheckResourceAttr(resourceName, "name", "launch spec name test update"),
					resource.TestCheckResourceAttr(resourceName, "instance_types.#", "5"),
					resource.TestCheckResourceAttr(resourceName, "instance_types.0", "m3.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types.1", "m4.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types.2", "m4.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types.3", "m5.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "instance_types.4", "m5.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "preferred_spot_types.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "preferred_spot_types.0", "m3.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "preferred_spot_types.1", "m4.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "preferred_od_types.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "preferred_od_types.0", "m4.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "preferred_od_types.1", "m3.2xlarge"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.0", "sg-0041bd3fd6aa2ee3c"),
					resource.TestCheckResourceAttr(resourceName, "security_groups.1", "sg-0195f2ac3a6014a15"),
					resource.TestCheckResourceAttr(resourceName, "user_data", elastigroup_aws_launch_configuration.Base64StateFunc("hello world updated")),
					resource.TestCheckResourceAttr(resourceName, "iam_instance_profile", "updated"),
					resource.TestCheckResourceAttr(resourceName, "labels.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "labels.0.key", "label key updated"),
					resource.TestCheckResourceAttr(resourceName, "labels.0.value", "label value updated"),
					resource.TestCheckResourceAttr(resourceName, "taints.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.key", "taint key updated"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.value", "taint value updated"),
					resource.TestCheckResourceAttr(resourceName, "taints.0.effect", "NoExecute"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.0", "subnet-4333093a"),
					resource.TestCheckResourceAttr(resourceName, "subnet_ids.1", "subnet-8ab89cc1"),
					resource.TestCheckResourceAttr(resourceName, "root_volume_size", "30"),
					resource.TestCheckResourceAttr(resourceName, "associate_public_ip_address", "false"),
					resource.TestCheckResourceAttr(resourceName, "restrict_scale_down", "false"),
					resource.TestCheckResourceAttr(resourceName, "reserved_enis", "2"),
					resource.TestCheckResourceAttr(resourceName, "startup_taints.0.key", "startuptaint-key-updated"),
					resource.TestCheckResourceAttr(resourceName, "startup_taints.0.value", "startuptaint-value-updated"),
					resource.TestCheckResourceAttr(resourceName, "startup_taints.0.effect", "NoExecute"),
				),
			},
		},
	})
}

const testBaselineOceanAWSLaunchSpecConfig_Create = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
  provider = "%v"  

  ocean_id = "%v"
  image_id = ""
  security_groups = ["sg-0041bd3fd6aa2ee3c"]
  user_data = "hello world"
  iam_instance_profile = "test"
  subnet_ids = ["subnet-4333093a"]
  instance_types = ["m3.xlarge", "m4.2xlarge", "m5.2xlarge"]
  preferred_spot_types = ["m3.xlarge"]
  preferred_od_types = ["m4.2xlarge"]
  root_volume_size = 20 
  associate_public_ip_address = true
  restrict_scale_down = true
  reserved_enis = 1

  labels {
    key = "label key"
    value = "label value"
  }

  taints {
    key = "taint key"
    value = "taint value"
    effect = "NoSchedule"
  }

  startup_taints {
	key = "startuptaint-key"
	value = "startuptaint-value"
	effect = "NoSchedule"
  }

 %v
}
`

const testBaselineOceanAWSLaunchSpecConfig_Update = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
  provider = "%v"

  ocean_id = "%v"
  image_id = "ami-05a68f290aa68e8f0"
  user_data = "hello world updated"
  iam_instance_profile = "updated"
  subnet_ids = ["subnet-4333093a", "subnet-8ab89cc1"]
  instance_types = ["m3.2xlarge", "m4.xlarge", "m4.2xlarge", "m5.xlarge", "m5.2xlarge"]
  preferred_spot_types = ["m3.2xlarge","m4.xlarge"]
  preferred_od_types = ["m4.2xlarge", "m3.2xlarge"]
  security_groups = ["sg-0041bd3fd6aa2ee3c","sg-0195f2ac3a6014a15" ]
  root_volume_size = 30
  name = "launch spec name test update"
  associate_public_ip_address = false
  restrict_scale_down = false
  reserved_enis = 2

  labels {
    key = "label key updated"
    value = "label value updated"
  }

  taints {
    key = "taint key updated"
    value = "taint value updated"
    effect = "NoExecute"
  }

  startup_taints {
	key = "startuptaint-key-updated"
	value = "startuptaint-value-updated"
	effect = "NoExecute"
  }

%v
}
`

// endregion

// region OceanAWSLaunchSpec: AutoScale
func TestAccSpotinstOceanAWSLaunchSpec_AutoScale(t *testing.T) {
	oceanID := "o-8b34732f"
	resourceName := createOceanAWSLaunchSpecResourceOceanID(oceanID)

	var launchSpec aws.LaunchSpec
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSLaunchSpecDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID: oceanID,
				}, testAutoScaleOceanAWSLaunchSpecConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.0.gpu_per_unit", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.0.num_of_units", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.0.memory_per_unit", "256"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.1.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.1.gpu_per_unit", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.1.num_of_units", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.1.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms_automatic.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms_automatic.0.auto_headroom_percentage", "10"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.key", "fakeKey"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.value", "fakeVal"),
				),
			},
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}, testAutoScaleOceanAWSLaunchSpecConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.0.cpu_per_unit", "1024"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.0.gpu_per_unit", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.0.num_of_units", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.0.memory_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms_automatic.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms_automatic.0.auto_headroom_percentage", "5"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.key", "updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.0.value", "updated"),
				),
			},
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}, testAutoScaleOceanAWSLaunchSpecConfig_Delete),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "autoscale_headrooms_automatic.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "0"),
				),
			},
		},
	})
}

const testAutoScaleOceanAWSLaunchSpecConfig_Create = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"  
 ocean_id = "%v"

 image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
 name = "launch spec name test"

 autoscale_headrooms {
   cpu_per_unit = 1024
   gpu_per_unit = 1
   memory_per_unit = 512
   num_of_units = 1
 }
   
 autoscale_headrooms {
   cpu_per_unit = 1024
   gpu_per_unit = 1
   memory_per_unit = 256
   num_of_units = 1
 }

 autoscale_headrooms_automatic {
 	auto_headroom_percentage = 10
 }

 tags {
     key   = "fakeKey"
     value = "fakeVal"
  } 
%v
}

`

const testAutoScaleOceanAWSLaunchSpecConfig_Update = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"  
 ocean_id = "%v"
 
 image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
 name = "launch spec name test"

 autoscale_headrooms {
   cpu_per_unit = 1024
   gpu_per_unit = 1
   memory_per_unit = 512
   num_of_units = 1
 }

 autoscale_headrooms_automatic {
 	auto_headroom_percentage = 5
 }

 tags {
     key   = "updated"
     value = "updated"
  } 
%v
}

`

const testAutoScaleOceanAWSLaunchSpecConfig_Delete = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"
 ocean_id = "%v"
 name = "launch spec name test"

 image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
%v
}

`

//endregion

// region OceanAWSLaunchSpec: ElasticIpPool
func TestAccSpotinstOceanAWSLaunchSpec_ElasticIpPool(t *testing.T) {
	oceanID := "o-8b34732f"
	resourceName := createOceanAWSLaunchSpecResourceOceanID(oceanID)

	var launchSpec aws.LaunchSpec
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSLaunchSpecDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID: oceanID,
				}, testLaunchSpecOceanAWSElasticIpPool_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "elastic_ip_pool.#", "1"),
				),
			},
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}, testLaunchSpecOceanAWSElasticIpPool_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "elastic_ip_pool.#", "1"),
				),
			},
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}, testLaunchSpecOceanAWSElasticIpPool_Delete),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "elastic_ip_pool.#", "0"),
				),
			},
		},
	})
}

const testLaunchSpecOceanAWSElasticIpPool_Create = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"  
 ocean_id = "%v"

 image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
 name = "launch spec name test"

 elastic_ip_pool {
   tag_selector {
     tag_key = "create key "
     tag_value = "create value"
   }
 }

%v
}

`

const testLaunchSpecOceanAWSElasticIpPool_Update = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"  
 ocean_id = "%v"
 
 image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
 name = "launch spec name test"

 elastic_ip_pool {
   tag_selector {
     tag_key = "update key "
     tag_value = "update value"
   }
 }

%v
}

`

const testLaunchSpecOceanAWSElasticIpPool_Delete = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"
 ocean_id = "%v"
 name = "launch spec name test"

 image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
%v
}

`

//endregion

// region OceanAWSLaunchSpec: Block Devices
func TestAccSpotinstOceanAWSLaunchSpec_BlockDeviceMappings(t *testing.T) {
	oceanID := "o-8b34732f"
	resourceName := createOceanAWSLaunchSpecResourceOceanID(oceanID)

	var launchSpec aws.LaunchSpec
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSLaunchSpecDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID: oceanID,
				}, testLaunchSpecOceanBlockDeviceMappings_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.device_name", "/dev/xvda1"),
					//resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.delete_on_termination", "true"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.dynamic_volume_size.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.dynamic_volume_size.0.base_size", "50"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.dynamic_volume_size.0.resource", "CPU"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.dynamic_volume_size.0.size_per_resource_unit", "20"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.encrypted", "false"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.kms_key_id", "kms-key"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.volume_type", "gp2"),
					resource.TestCheckResourceAttr(resourceName, "ephemeral_storage.0.ephemeral_storage_device_name", "/dev/xvda1"),
				),
			},
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}, testLaunchSpecOceanBlockDeviceMappings_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.device_name", "/dev/sda1"),
					//resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.delete_on_termination", "true"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.dynamic_volume_size.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.dynamic_volume_size.0.base_size", "50"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.dynamic_volume_size.0.resource", "CPU"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.dynamic_volume_size.0.size_per_resource_unit", "20"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.encrypted", "false"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.kms_key_id", "kms-key"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.volume_type", "gp3"),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.0.throughput", "500"),
					resource.TestCheckResourceAttr(resourceName, "ephemeral_storage.0.ephemeral_storage_device_name", "/dev/sda1"),
				),
			},
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}, testLaunchSpecOceanBlockDeviceMappings_EmptyFields),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "block_device_mappings.#", "0"),
					//resource.TestCheckResourceAttr(resourceName, "block_device_mappings.0.ebs.#", "1"),
				),
			},
		},
	})
}

const testLaunchSpecOceanBlockDeviceMappings_Create = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"  
 ocean_id = "%v"

image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
 name = "launch spec name test"

	block_device_mappings {
        device_name = "/dev/xvda1"
        ebs {
          delete_on_termination = "true"
          kms_key_id = "kms-key"
          encrypted = "false"
          volume_type = "gp2"
		  dynamic_volume_size {
            base_size = 50
            resource = "CPU"
            size_per_resource_unit = 20
          }
        }
}
  ephemeral_storage{
    ephemeral_storage_device_name = "/dev/xvda1"
  }
%v
      }

`

const testLaunchSpecOceanBlockDeviceMappings_Update = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"  
 ocean_id = "%v"

image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
 name = "launch spec name test"

	block_device_mappings {
        device_name = "/dev/sda1"
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

  ephemeral_storage{
    ephemeral_storage_device_name = "/dev/sda1"
  }
%v
      }
`

const testLaunchSpecOceanBlockDeviceMappings_EmptyFields = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
provider = "%v"
ocean_id = "%v"
name = "launch spec name test"

image_id = "ami-05a68f290aa68e8f0"
security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
user_data = "hello world updated"
iam_instance_profile = "updated"
%v
}
`

// endregion

// region OceanAWSLaunchSpec: ResourceLimits
func TestAccSpotinstOceanAWSLaunchSpec_ResourceLimits(t *testing.T) {
	oceanID := "o-8b34732f"
	resourceName := createOceanAWSLaunchSpecResourceOceanID(oceanID)

	var launchSpec aws.LaunchSpec
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSLaunchSpecDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID: oceanID,
				}, testLaunchSpecOceanAWSResourceLimits_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "resource_limits.0.max_instance_count", "5"),
					resource.TestCheckResourceAttr(resourceName, "resource_limits.0.min_instance_count", "0"),
				),
			},
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}, testLaunchSpecOceanAWSResourceLimits_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "resource_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "resource_limits.0.max_instance_count", "4"),
					resource.TestCheckResourceAttr(resourceName, "resource_limits.0.min_instance_count", "0"),
				),
			},
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}, testLaunchSpecOceanAWSResourceLimits_Delete),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "resource_limits.#", "0"),
				),
			},
		},
	})
}

const testLaunchSpecOceanAWSResourceLimits_Create = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"  
 ocean_id = "%v"

 image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
 name = "launch spec name test"

  resource_limits {
    max_instance_count = 5
    min_instance_count = 0
  }

%v
}

`

const testLaunchSpecOceanAWSResourceLimits_Update = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"  
 ocean_id = "%v"
 
 image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
 name = "launch spec name test"

  resource_limits {
    max_instance_count = 4
    min_instance_count = 0
  }

%v
}

`

const testLaunchSpecOceanAWSResourceLimits_Delete = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"
 ocean_id = "%v"
 name = "launch spec name test"

 image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
%v
}

`

//endregion

// region OceanAWSLaunchSpec: Strategy
func TestAccSpotinstOceanAWSLaunchSpec_Strategy(t *testing.T) {
	oceanID := "o-8b34732f"
	resourceName := createOceanAWSLaunchSpecResourceOceanID(oceanID)

	var launchSpec aws.LaunchSpec
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSLaunchSpecDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID: oceanID,
				}, testStrategyOceanAWSLaunchSpecConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.spot_percentage", "70"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.draining_timeout", "360"),
				),
			},
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}, testStrategyOceanAWSLaunchSpecConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.spot_percentage", "30"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.draining_timeout", "420"),
				),
			},
		},
	})
}

const testStrategyOceanAWSLaunchSpecConfig_Create = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"  
 ocean_id = "%v"

 image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
 name = "launch spec name test"

 strategy {
  spot_percentage = 70
  draining_timeout= 360
}

%v
}

`

const testStrategyOceanAWSLaunchSpecConfig_Update = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"  
 ocean_id = "%v"
 
 image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
 name = "launch spec name test"

 strategy {
  spot_percentage = 30
  draining_timeout= 420
}
%v
}

`

//endregion

// region OceanAWSLaunchSpec: Scheduling
func TestAccSpotinstOceanAWSLaunchSpec_Scheduling(t *testing.T) {
	oceanID := "o-8b34732f"
	resourceName := createOceanAWSLaunchSpecResourceOceanID(oceanID)

	var launchSpec aws.LaunchSpec
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "aws") },
		Providers:    TestAccProviders,
		CheckDestroy: testOceanAWSLaunchSpecDestroy,

		Steps: []resource.TestStep{
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID: oceanID,
				}, testSchedulingOceanAWSLaunchSpecConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "scheduling_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_task.0.task_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_task.0.task_headroom.0.cpu_per_unit", "512"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_task.0.task_headroom.0.num_of_units", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_task.0.cron_expression", "0 1 * * *"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_task.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_task.0.task_type", "manualHeadroomUpdate"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_shutdown_hours.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_shutdown_hours.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_shutdown_hours.0.time_windows.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_shutdown_hours.0.time_windows.0", "Sat:08:00-Sat:08:30"),
				),
			},
			{
				Config: createOceanAWSLaunchSpecTerraform(&LaunchSpecConfigMetadata{
					oceanID:              oceanID,
					updateBaselineFields: true}, testSchedulingOceanAWSLaunchSpecConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckOceanAWSLaunchSpecExists(&launchSpec, resourceName),
					testCheckOceanAWSLaunchSpecAttributes(&launchSpec, oceanID),
					resource.TestCheckResourceAttr(resourceName, "scheduling_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_task.0.task_headroom.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_task.0.task_headroom.0.memory_per_unit", "256"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_task.0.task_headroom.0.num_of_units", "2"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_task.0.cron_expression", "0 1 * * *"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_task.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_task.0.task_type", "manualHeadroomUpdate"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_shutdown_hours.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_shutdown_hours.0.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_shutdown_hours.0.time_windows.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_shutdown_hours.0.time_windows.0", "Sat:08:00-Sat:08:30"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_shutdown_hours.0.time_windows.1", "Sun:08:00-Sun:08:30"),
				),
			},
		},
	})
}

const testSchedulingOceanAWSLaunchSpecConfig_Create = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"  
 ocean_id = "%v"

 image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
 name = "launch spec name test"

  scheduling_task {
    is_enabled = true
    cron_expression = "0 1 * * *"
    task_type = "manualHeadroomUpdate"
    task_headroom {
      cpu_per_unit = 512
      num_of_units = 1
    }
  }

  scheduling_shutdown_hours {
    is_enabled = true
    time_windows = ["Sat:08:00-Sat:08:30"]
  }
%v
}

`

const testSchedulingOceanAWSLaunchSpecConfig_Update = `
resource "` + string(commons.OceanAWSLaunchSpecResourceName) + `" "%v" {
 provider = "%v"  
 ocean_id = "%v"
 
 image_id = "ami-05a68f290aa68e8f0"
 security_groups = ["sg-0041bd3fd6aa2ee3c", "sg-0195f2ac3a6014a15"]
 user_data = "hello world updated"
 iam_instance_profile = "updated"
 name = "launch spec name test"

  scheduling_task {
    is_enabled = true
    cron_expression = "0 1 * * *"
    task_type = "manualHeadroomUpdate"
    task_headroom {
      memory_per_unit = 256
      num_of_units = 2
    }
  }

  scheduling_shutdown_hours {
    is_enabled = false
    time_windows = ["Sat:08:00-Sat:08:30", "Sun:08:00-Sun:08:30"]
  }
%v
}

`

//endregion

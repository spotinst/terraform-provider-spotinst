package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func init() {
	resource.AddTestSweepers("spotinst_elastigroup_gcp", &resource.Sweeper{
		Name: "spotinst_elastigroup_gcp",
		F:    testSweepElastigroupGCP,
	})
}

func testSweepElastigroupGCP(region string) error {
	client, err := getProviderClient("gcp")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).elastigroup.CloudProviderGCP()

	input := &gcp.ListGroupsInput{}
	if resp, err := conn.List(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of groups to sweep")
	} else {
		if len(resp.Groups) == 0 {
			log.Printf("[INFO] No groups to sweep")
		}
		for _, group := range resp.Groups {
			if strings.Contains(spotinst.StringValue(group.Name), "test-acc-") {
				if _, err := conn.Delete(context.Background(), &gcp.DeleteGroupInput{GroupID: group.ID}); err != nil {
					return fmt.Errorf("unable to delete group %v in sweep", spotinst.StringValue(group.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(group.ID))
				}
			}
		}
	}
	return nil
}

// createElastigroupGCPResourceName creates a resource name for the test group
func createElastigroupGCPResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.ElastigroupGCPResourceName), name)
}

// testElastigroupGCPDestroy checks whether a group has been destroyed and returns an error if it still exists
func testElastigroupGCPDestroy(s *terraform.State) error {
	client := testAccProviderGCP.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.ElastigroupGCPResourceName) {
			continue
		}
		input := &gcp.ReadGroupInput{GroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.elastigroup.CloudProviderGCP().Read(context.Background(), input)
		if err == nil && resp != nil && resp.Group != nil {
			return fmt.Errorf("group still exists")
		}
	}
	return nil
}

// testCheckElastigroupGCPAttributes checks the correct group is being tests and returns an error if it is not
func testCheckElastigroupGCPAttributes(group *gcp.Group, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(group.Name) != expectedName {
			return fmt.Errorf("bad content: %v", group.Name)
		}
		return nil
	}
}

// testCheckElastigroupGCPExists checks if a group exists and returns an error if the resource isn't found, if the
// group id was not set, or the group does not exist (wasn't created or was unexpectedly deleted, etc)
func testCheckElastigroupGCPExists(group *gcp.Group, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderGCP.Meta().(*Client)
		input := &gcp.ReadGroupInput{GroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.elastigroup.CloudProviderGCP().Read(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(resp.Group.Name) != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Group not found: %+v,\n %+v\n", resp.Group, rs.Primary.Attributes)
		}
		*group = *resp.Group
		return nil
	}
}

// GCPGroupConfigMetadata holds blocks of attributes defined as strings that are used to build a Terraform resource
type GCPGroupConfigMetadata struct {
	variables            string
	provider             string
	groupName            string
	instanceTypes        string
	launchConfig         string
	disk                 string
	strategy             string
	fieldsToAppend       string
	updateBaselineFields bool
}

// createElastigroupGCPTerraform builds a valid Terraform resource as a string.
// This function appends attribute blocks defined as string later in this file.
// These blocks should have fields required for a bare-minimum group to be created.
func createElastigroupGCPTerraform(gcm *GCPGroupConfigMetadata) string {
	if gcm == nil {
		return ""
	}

	if gcm.provider == "" {
		gcm.provider = "gcp"
	}

	if gcm.instanceTypes == "" {
		gcm.instanceTypes = testInstanceTypesGCPGroupConfig_Create
	}

	//if gcm.launchConfig == "" {
	//	gcm.launchConfig = testLaunchConfigurationGCPGroupConfig_Create
	//}

	if gcm.disk == "" {
		gcm.disk = testDiskGCPGroupConfig_Create
	}

	if gcm.strategy == "" {
		gcm.strategy = testStrategyGCPGroupConfig_Create
	}

	template :=
		`provider "gcp" {
	 token   = "fake"
	 account = "fake"
	}
	`
	if gcm.updateBaselineFields {
		format := testBaselineGCPGroupConfig_Update

		template += fmt.Sprintf(format,
			gcm.groupName,
			gcm.provider,
			gcm.groupName,
			gcm.instanceTypes,
			gcm.launchConfig,
			gcm.disk,
			gcm.strategy,
			gcm.fieldsToAppend,
		)
	} else {
		format := testBaselineGCPGroupConfig_Create

		template += fmt.Sprintf(format,
			gcm.groupName,
			gcm.provider,
			gcm.groupName,
			gcm.instanceTypes,
			gcm.launchConfig,
			gcm.disk,
			gcm.strategy,
			gcm.fieldsToAppend,
		)
	}

	if gcm.variables != "" {
		template = gcm.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", gcm.groupName, template)
	return template
}

// region Elastigroup GCP: Baseline
func TestAccSpotinstElastigroupGCP_Baseline(t *testing.T) {
	groupName := "test-acc-eg-gcp-baseline"
	resourceName := createElastigroupGCPResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "gcp") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupGCPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName: groupName,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.0", "us-west1-a"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.1", "us-west1-b"),
					resource.TestCheckResourceAttr(resourceName, "preferred_availability_zones.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "preferred_availability_zones.0", "us-west1-a"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "0"),
				),
			},
			{
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:            groupName,
					updateBaselineFields: true,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "availability_zones.0", "us-central1-b"),
					resource.TestCheckResourceAttr(resourceName, "preferred_availability_zones.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "preferred_availability_zones.0", "us-central1-b"),
					resource.TestCheckResourceAttr(resourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "0"),
				),
			},
		},
	})
}

const testBaselineGCPGroupConfig_Create = `
resource "` + string(commons.ElastigroupGCPResourceName) + `" "%v" {
 provider = "%v"

 name = "%v"
 description = "created by Terraform"
 availability_zones = ["us-west1-a", "us-central1-b"]
 preferred_availability_zones = ["us-west1-a"]

 // --- CAPACITY ------------
 max_size = 0
 min_size = 0
 desired_capacity = 0
 // -------------------------
 
 %v
 %v
 %v
 %v
 %v
}

`

const testBaselineGCPGroupConfig_Update = `
resource "` + string(commons.ElastigroupGCPResourceName) + `" "%v" {
 provider = "%v"

 name = "%v"
 description = "created by Terraform"
 availability_zones = ["us-central1-b"]
 preferred_availability_zones = ["us-central1-b"]
 // --- CAPACITY ------------
 max_size = 0
 min_size = 0
 desired_capacity = 0
 // -------------------------
 
 %v
 %v
 %v
 %v
 %v
}

`

// endregion

// region Elastigroup GCP: Instance Types
func TestAccSpotinstElastigroupGCP_InstanceTypes(t *testing.T) {
	groupName := "test-acc-eg-gcp-instance-types"
	resourceName := createElastigroupGCPResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "gcp") },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupGCPDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:     groupName,
					instanceTypes: testInstanceTypesGCPGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "instance_types_ondemand", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preemptible.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preemptible.0", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preemptible.1", "n1-standard-2"),
				),
			},
			{
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:     groupName,
					instanceTypes: testInstanceTypesGCPGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "instance_types_ondemand", "n1-standard-2"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preemptible.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preemptible.0", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preemptible.1", "n1-standard-2"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preemptible.2", "n1-standard-4"),
				),
			},
		},
	})
}

const testInstanceTypesGCPGroupConfig_Create = `
 // --- INSTANCE TYPES --------------------------------
 instance_types_ondemand = "n1-standard-1"
 instance_types_preemptible = ["n1-standard-1", "n1-standard-2"]
 // ---------------------------------------------------
`

const testInstanceTypesGCPGroupConfig_Update = `
 // --- INSTANCE TYPES --------------------------------
 instance_types_ondemand = "n1-standard-2"
 instance_types_preemptible = ["n1-standard-1", "n1-standard-2", "n1-standard-4"]
 // ---------------------------------------------------
`

// endregion

// region Elastigroup GCP: Launch Configuration
//func TestAccSpotinstElastigroupGCP_LaunchConfiguration(t *testing.T) {
//	groupName := "test-acc-eg-gcp-launch-configuration"
//	resourceName := createElastigroupGCPResourceName(groupName)
//
//	var group gcp.Group
//	resource.Test(t, resource.TestCase{
//		PreCheck:      func() { testAccPreCheck(t, "gcp") },
//		Providers:     TestAccProviders,
//		CheckDestroy:  testElastigroupGCPDestroy,
//		IDRefreshName: resourceName,
//
//		Steps: []resource.TestStep{
//			{
//				ResourceName: resourceName,
//				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
//					groupName:    groupName,
//					launchConfig: testLaunchConfigurationGCPGroupConfig_Create,
//				}),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckElastigroupGCPExists(&group, resourceName),
//					testCheckElastigroupGCPAttributes(&group, groupName),
//					resource.TestCheckResourceAttr(resourceName, "service_account", "265168459660-compute@developer.gserviceaccount.com"),
//					resource.TestCheckResourceAttr(resourceName, "startup_script", "echo hello world"),
//					resource.TestCheckResourceAttr(resourceName, "shutdown_script", elastigroup_gcp_launch_configuration.Base64StateFunc("echo goodbye world")),
//					resource.TestCheckResourceAttr(resourceName, "backend_services.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash_create+".named_ports.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash_create+".service_name", "terraform-bs-dont-delete"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash_create+".location_type", "global"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash_create+".named_ports."+NamedPortsHash_create+".name", "http"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash_create+".named_ports."+NamedPortsHash_create+".ports.#", "2"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash_create+".named_ports."+NamedPortsHash_create+".ports.0", "80"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash_create+".named_ports."+NamedPortsHash_create+".ports.1", "8080"),
//					resource.TestCheckResourceAttr(resourceName, "ip_forwarding", "false"),
//					resource.TestCheckResourceAttr(resourceName, "labels.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "labels."+LabelsHash_create+".key", "test_key"),
//					resource.TestCheckResourceAttr(resourceName, "labels."+LabelsHash_create+".value", "test_value"),
//					resource.TestCheckResourceAttr(resourceName, "metadata.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "metadata."+MetaHash_create+".key", "metadata_key"),
//					resource.TestCheckResourceAttr(resourceName, "metadata."+MetaHash_create+".value", "metadata_value"),
//					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "tags.0", "test_tag"),
//				),
//			},
//			{
//				ResourceName: resourceName,
//				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
//					groupName:    groupName,
//					launchConfig: testLaunchConfigurationGCPGroupConfig_Update,
//				}),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckElastigroupGCPExists(&group, resourceName),
//					testCheckElastigroupGCPAttributes(&group, groupName),
//					resource.TestCheckResourceAttr(resourceName, "service_account", "terraform-acc-test-account@spotinst-labs.iam.gserviceaccount.com"),
//					resource.TestCheckResourceAttr(resourceName, "startup_script", "echo hello world updated"),
//					resource.TestCheckResourceAttr(resourceName, "shutdown_script", elastigroup_gcp_launch_configuration.Base64StateFunc("echo goodbye world updated")),
//					resource.TestCheckResourceAttr(resourceName, "backend_services.#", "2"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash1_update+".service_name", "terraform-bs-dont-delete"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash1_update+".named_ports.#", "1"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash1_update+".named_ports."+NamedPortsHash1_update+".name", "http"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash1_update+".named_ports."+NamedPortsHash1_update+".ports.#", "2"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash1_update+".named_ports."+NamedPortsHash1_update+".ports.0", "40"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash1_update+".named_ports."+NamedPortsHash1_update+".ports.1", "4040"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash2_update+".service_name", "terraform-bs-dont-delete"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash2_update+".location_type", "regional"),
//					resource.TestCheckResourceAttr(resourceName, "backend_services."+BackendSvcHash2_update+".scheme", "EXTERNAL"),
//					resource.TestCheckResourceAttr(resourceName, "ip_forwarding", "true"),
//					resource.TestCheckResourceAttr(resourceName, "labels.#", "2"),
//					resource.TestCheckResourceAttr(resourceName, "labels."+LabelsHash_create+".key", "test_key"),
//					resource.TestCheckResourceAttr(resourceName, "labels."+LabelsHash_create+".value", "test_value"),
//					resource.TestCheckResourceAttr(resourceName, "labels."+LabelsHash_update+".key", "test_key2"),
//					resource.TestCheckResourceAttr(resourceName, "labels."+LabelsHash_update+".value", "test_value2"),
//					resource.TestCheckResourceAttr(resourceName, "metadata.#", "2"),
//					resource.TestCheckResourceAttr(resourceName, "metadata."+MetaHash_create+".key", "metadata_key"),
//					resource.TestCheckResourceAttr(resourceName, "metadata."+MetaHash_create+".value", "metadata_value"),
//					resource.TestCheckResourceAttr(resourceName, "metadata."+MetaHash_update+".key", "metadata_key2"),
//					resource.TestCheckResourceAttr(resourceName, "metadata."+MetaHash_update+".value", "metadata_value2"),
//					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
//					resource.TestCheckResourceAttr(resourceName, "tags.0", "test_tag"),
//					resource.TestCheckResourceAttr(resourceName, "tags.1", "test_tag2"),
//				),
//			},
//			{
//				ResourceName: resourceName,
//				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
//					groupName:    groupName,
//					launchConfig: testLaunchConfigurationGCPGroupConfig_EmptyFields,
//				}),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckElastigroupGCPExists(&group, resourceName),
//					testCheckElastigroupGCPAttributes(&group, groupName),
//					resource.TestCheckResourceAttr(resourceName, "service_account", "cannot set empty service account"),
//					resource.TestCheckResourceAttr(resourceName, "startup_script", "cannot set empty startup script"),
//					resource.TestCheckResourceAttr(resourceName, "shutdown_script", elastigroup_gcp_launch_configuration.Base64StateFunc("cannot set empty shutdown script")),
//					resource.TestCheckResourceAttr(resourceName, "ip_forwarding", "false"),
//					resource.TestCheckResourceAttr(resourceName, "labels.#", "0"),
//					resource.TestCheckResourceAttr(resourceName, "metadata.#", "0"),
//					resource.TestCheckResourceAttr(resourceName, "tags.#", "0"),
//				),
//			},
//		},
//	})
//}
//
//const (
//	LabelsHash_create      = "3430685526"
//	LabelsHash_update      = "3119730257"
//	MetaHash_create        = "1912256051"
//	MetaHash_update        = "284772212"
//	BackendSvcHash_create  = "1781664423"
//	BackendSvcHash1_update = "659180285"
//	BackendSvcHash2_update = "2663756714"
//	NamedPortsHash_create  = "571950593"
//	NamedPortsHash1_update = "981148154"
//	NamedPortsHash2_update = "1016050568"
//)
//
//const testLaunchConfigurationGCPGroupConfig_Create = `
// // --- LAUNCH CONFIGURATION --------------
// service_account = "265168459660-compute@developer.gserviceaccount.com"
// startup_script = "echo hello world"
// shutdown_script = "echo goodbye world"
// ip_forwarding = false
//
// labels {
//     key = "test_key"
//     value = "test_value"
//   }
//
// metadata {
//     key = "metadata_key"
//     value = "metadata_value"
//   }
//
// tags = ["test_tag"]
//
// backend_services {
//    service_name = "terraform-acc-test-backend-service"
//    location_type = "global"
//
//    named_ports {
//      name = "http"
//      ports = [80, 8080]
//    }
//  }
//
// // ---------------------------------------
//`
//
//const testLaunchConfigurationGCPGroupConfig_Update = `
// // --- LAUNCH CONFIGURATION --------------
// service_account = "terraform-acc-test-account@spotinst-labs.iam.gserviceaccount.com"
// startup_script = "echo hello world updated"
// shutdown_script = "echo goodbye world updated"
// ip_forwarding = true
//
// labels {
//     key = "test_key"
//     value = "test_value"
//   }
// labels {
//     key = "test_key2"
//     value = "test_value2"
//   }
//
// metadata {
//     key = "metadata_key"
//     value = "metadata_value"
//   }
//
// metadata {
//     key = "metadata_key2"
//     value = "metadata_value2"
//   }
//
// tags = ["test_tag", "test_tag2"]
//
// backend_services {
//  service_name = "terraform-acc-test-backend-service"
//
//  named_ports {
//    name = "http"
//    ports = [40, 4040]
//  }
// }
//
// backend_services {
//   service_name  = "terraform-acc-test-backend-service-tcp"
//   location_type = "regional"
//   scheme        = "EXTERNAL"
// }
// // ---------------------------------------
//`
//
//const testLaunchConfigurationGCPGroupConfig_EmptyFields = `
// // --- LAUNCH CONFIGURATION --------------
// service_account = "cannot set empty service account"
// startup_script  = "cannot set empty startup script"
// shutdown_script = "cannot set empty shutdown script"
// // ---------------------------------------
//`
//
//// endregion

// region Elastigroup GCP: Disk
func TestAccSpotinstElastigroupGCP_Disk(t *testing.T) {
	groupName := "test-acc-eg-gcp-disk"
	resourceName := createElastigroupGCPResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "gcp") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupGCPDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName: groupName,
					disk:      testDiskGCPGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "disk.0.auto_delete", "false"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.boot", "false"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.device_name", "tf-test-device"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.interface", "SCSI"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.mode", "READ_WRITE"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.source", "fake-source"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.type", "PERSISTENT"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.initialize_params.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.initialize_params.0.disk_size_gb", "20"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.initialize_params.0.disk_type", "pd-standard"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.initialize_params.0.source_image", "https://www.googleapis.com/compute/v1/projects/spotinst-labs/global/images/test-image-1"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName: groupName,
					disk:      testDiskGCPGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "disk.0.auto_delete", "true"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.boot", "true"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.device_name", "tf-test-device-updated"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.interface", "NVM"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.mode", "READ_ONLY"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.source", "fake-source-updated"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.type", "SCRATCH"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.initialize_params.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.initialize_params.0.disk_size_gb", "30"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.initialize_params.0.disk_type", "local-ssd"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.initialize_params.0.source_image", "https://www.googleapis.com/compute/v1/projects/spotinst-labs/global/images/test-image-2"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName: groupName,
					disk:      testDiskGCPGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "disk.0.initialize_params.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "disk.0.initialize_params.0.source_image", "https://www.googleapis.com/compute/v1/projects/spotinst-labs/global/images/test-image-2"),
				),
			},
		},
	})
}

const testDiskGCPGroupConfig_Create = `
 // --- DISK ------------------------
  disk {
    auto_delete = false
    boot = false
    device_name = "tf-test-device"
    interface = "SCSI"
		mode = "READ_WRITE"
		source = "fake-source"
		type = "PERSISTENT"

    initialize_params {
			disk_size_gb = 20
			disk_type = "pd-standard"
			source_image = "https://www.googleapis.com/compute/v1/projects/spotinst-labs/global/images/test-image-1"
		}
  }
 // ---------------------------------
`

const testDiskGCPGroupConfig_Update = `
 // --- DISK ------------------------
  disk {
    auto_delete = true
    boot = true
    device_name = "tf-test-device-updated"
    interface = "NVM"
		mode = "READ_ONLY"
		source = "fake-source-updated"
		type = "SCRATCH"

    initialize_params {
			disk_size_gb = 30
			disk_type = "local-ssd"
      source_image = "https://www.googleapis.com/compute/v1/projects/spotinst-labs/global/images/test-image-2"
		}
  }
 // ---------------------------------
`

const testDiskGCPGroupConfig_EmptyFields = `
 // --- DISK ------------------------
  disk {
    initialize_params {
			source_image = "https://www.googleapis.com/compute/v1/projects/spotinst-labs/global/images/test-image-2"
		}
  }
 // ---------------------------------
`

// endregion

// region Elastigroup GCP: Strategy
func TestAccSpotinstElastigroupGCP_Strategy(t *testing.T) {
	groupName := "test-acc-eg-gcp-strategy"
	resourceName := createElastigroupGCPResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "gcp") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupGCPDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName: groupName,
					strategy:  testStrategyGCPGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "fallback_to_ondemand", "true"),
					resource.TestCheckResourceAttr(resourceName, "preemptible_percentage", "100"),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "300"),
					resource.TestCheckResourceAttr(resourceName, "provisioning_model", "SPOT"),
					resource.TestCheckResourceAttr(resourceName, "should_utilize_commitments", "true"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName: groupName,
					strategy:  testStrategyGCPGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "ondemand_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "draining_timeout", "240"),
					resource.TestCheckResourceAttr(resourceName, "provisioning_model", "PREEMPTIBLE"),
					resource.TestCheckResourceAttr(resourceName, "should_utilize_commitments", "false"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName: groupName,
					strategy:  testStrategyGCPGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "ondemand_count", "1"),
				),
			},
		},
	})
}

const testStrategyGCPGroupConfig_Create = `
 // --- STRATEGY --------------------
 fallback_to_ondemand = true
 preemptible_percentage = 100
 draining_timeout = 300
 provisioning_model = "SPOT"
should_utilize_commitments = true
 // ---------------------------------
`

const testStrategyGCPGroupConfig_Update = `
 // --- STRATEGY --------------------
  ondemand_count = 1
  draining_timeout = 240
  provisioning_model = "PREEMPTIBLE"
should_utilize_commitments = false
 // ---------------------------------
`

const testStrategyGCPGroupConfig_EmptyFields = `
 // --- STRATEGY --------------------
 ondemand_count = 1
 // ---------------------------------
`

// endregion

// region Elastigroup GCP: GPU
func TestAccSpotinstElastigroupGCP_GPU(t *testing.T) {
	groupName := "test-acc-eg-gcp-gpu"
	resourceName := createElastigroupGCPResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "gcp") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupGCPDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testGPUGCPGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "gpu.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "gpu.0.count", "2"),
					resource.TestCheckResourceAttr(resourceName, "gpu.0.type", "nvidia-tesla-p100"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testGPUGCPGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "gpu.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "gpu.0.count", "1"),
					resource.TestCheckResourceAttr(resourceName, "gpu.0.type", "nvidia-tesla-v100"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testGPUGCPGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "gpu.#", "0"),
				),
			},
		},
	})
}

const testGPUGCPGroupConfig_Create = `
// --- GPU ----------------------------------------------
	gpu {
		count = 2
		type = "nvidia-tesla-p100"
	}
// ------------------------------------------------------
`

const testGPUGCPGroupConfig_Update = `
// --- GPU ----------------------------------------------
	gpu {
		count = 1
		type = "nvidia-tesla-v100"
	}
// ------------------------------------------------------
`

const testGPUGCPGroupConfig_EmptyFields = `
// --- GPU ----------------------------------------------
// ------------------------------------------------------
`

// endregion

// region Elastigroup GCP: Health Checks
func TestAccSpotinstElastigroupGCP_HealthChecks(t *testing.T) {
	groupName := "test-acc-eg-gcp-health-checks"
	resourceName := createElastigroupGCPResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "gcp") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupGCPDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testHealthChecksGCPGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "health_check_grace_period", "100"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testHealthChecksGCPGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "health_check_grace_period", "50"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testHealthChecksGCPGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "health_check_grace_period", "0"),
				),
			},
		},
	})
}

const testHealthChecksGCPGroupConfig_Create = `
// --- HEALTH-CHECKS ------------------------------------
 auto_healing = true
 health_check_grace_period = 100
 health_check_type = "K8S_NODE"
 unhealthy_duration = 120
// ------------------------------------------------------
`

const testHealthChecksGCPGroupConfig_Update = `
// --- HEALTH-CHECKS ------------------------------------
 auto_healing = false
 health_check_grace_period = 50
 health_check_type = "K8S_NODE"
 unhealthy_duration = 180 
// ------------------------------------------------------
`

const testHealthChecksGCPGroupConfig_EmptyFields = `
// --- HEALTH-CHECKS ------------------------------------
 health_check_grace_period = 0
 health_check_type = "K8S_NODE"
 unhealthy_duration = 60 
// ------------------------------------------------------
`

// endregion

// region Elastigroup GCP: Network Interfaces
func TestAccSpotinstElastigroupGCP_NetworkInterfaces(t *testing.T) {
	groupName := "test-acc-eg-network-interfaces"
	resourceName := createElastigroupGCPResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "gcp") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupGCPDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testNetworkInterfacesGCPGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "network_interface.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.network", "default"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.0.name", "config1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.0.type", "ONE_TO_ONE_NAT"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.alias_ip_ranges.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.alias_ip_ranges.0.subnetwork_range_name", "range-name-1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.alias_ip_ranges.0.ip_cidr_range", "10.128.0.0/20"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testNetworkInterfacesGCPGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "network_interface.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.network", "updated"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.0.name", "config2"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.access_configs.0.type", "ONE_TO_ONE_NAT"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.alias_ip_ranges.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.alias_ip_ranges.0.subnetwork_range_name", "range-name-2"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.alias_ip_ranges.0.ip_cidr_range", "10.128.0.0/20"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testNetworkInterfacesGCPGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "network_interface.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network_interface.0.network", "default"),
				),
			},
		},
	})
}

const testNetworkInterfacesGCPGroupConfig_Create = `
 // --- NETWORK INTERFACE ------------------
	network_interface { 
		network = "default"
	
		access_configs {
			name = "config1"
			type = "ONE_TO_ONE_NAT"
		}

		alias_ip_ranges {
		 subnetwork_range_name = "range-name-1"
		 ip_cidr_range = "10.128.0.0/20"
		}
  }
 // ----------------------------------------
`

const testNetworkInterfacesGCPGroupConfig_Update = `
 // --- NETWORK INTERFACE ------------------
	network_interface {
		network = "updated"
	
		access_configs {
			name = "config2"
			type = "ONE_TO_ONE_NAT"
		}

		alias_ip_ranges {
		 subnetwork_range_name = "range-name-2"
		 ip_cidr_range 				 = "10.128.0.0/20"
		}
	}
 // ----------------------------------------
`

const testNetworkInterfacesGCPGroupConfig_EmptyFields = `
 // --- NETWORK INTERFACE ------------------
  network_interface {     
    network = "default"
  }
 // ----------------------------------------
`

// endregion

// region Elastigroup GCP: Scaling Up Policies
func TestAccSpotinstElastigroupGCP_ScalingUpPolicies(t *testing.T) {
	groupName := "test-acc-eg-gcp-scaling-up-policy"
	resourceName := createElastigroupGCPResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "gcp") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupGCPDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingUpPolicyGCPGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.policy_name", "policy-name"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.metric_name", "CPUUtilization"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.namespace", "test-namespace"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.source", "spectrum"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.statistic", "count"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.unit", "seconds"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.cooldown", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.dimensions.0.name", "name-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.dimensions.0.value", "value-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.threshold", "10"),
					//resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.operator", "gte"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.evaluation_periods", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.period", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.action_type", "adjustment"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.adjustment", "1"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingUpPolicyGCPGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.policy_name", "policy-name-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.metric_name", "CPUUtilization"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.namespace", "updated-namespace"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.source", "stackdriver"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.statistic", "sum"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.unit", "bytes"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.cooldown", "300"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.dimensions.0.name", "name-1-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.dimensions.0.value", "value-1-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.threshold", "5"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.operator", "lte"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.evaluation_periods", "20"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.period", "300"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.action_type", "adjustment"),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.0.adjustment", "2"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingUpPolicyGCPGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_up_policy.#", "0"),
				),
			},
		},
	})
}

const testScalingUpPolicyGCPGroupConfig_Create = `
// --- SCALE UP POLICY ------------------
	scaling_up_policy {
	 policy_name = "policy-name"
	 metric_name = "CPUUtilization"
	 namespace = "test-namespace"
	 source = "spectrum"
	 statistic = "count"
	 unit = "seconds"
	 cooldown = 60
	
	 dimensions {
			 name = "name-1"
			 value = "value-1"
	 }
	
	 threshold = 10
	 operator = "gte"
	 evaluation_periods = 10
	 period = 60
	
	 // === ADJUSTMENT ===================
	 action_type = "adjustment"
	 adjustment = 1
	 // ==================================
 }
// ----------------------------------------
`

const testScalingUpPolicyGCPGroupConfig_Update = `
// --- SCALE UP POLICY ------------------
	scaling_up_policy {
	 policy_name = "policy-name-update"
	 metric_name = "CPUUtilization"
	 namespace = "updated-namespace"
	 source = "stackdriver"
	 statistic = "sum"
	 unit = "bytes"
	 cooldown = 300
	 threshold = 5
	
	 dimensions {
			 name = "name-1-update"
			 value = "value-1-update"
	 }
	
	 operator = "lte"
	 evaluation_periods = 20
	 period = 300
	
	 // === ADJUSTMENT ===================
	 action_type = "adjustment"
	 adjustment = 2
	 // ==================================
 }
// ----------------------------------------
`

const testScalingUpPolicyGCPGroupConfig_EmptyFields = `
// --- SCALE UP POLICY ------------------
// ----------------------------------------
`

// endregion

// region Elastigroup GCP: Scaling Down Policies
func TestAccSpotinstElastigroupGCP_ScalingDownPolicies(t *testing.T) {
	groupName := "test-acc-eg-gcp-scaling-down-policy"
	resourceName := createElastigroupGCPResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "gcp") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupGCPDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingDownPolicyGCPGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.policy_name", "policy-name"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.metric_name", "CPUUtilization"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.namespace", "test-namespace"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.source", "spectrum"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.statistic", "count"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.unit", "seconds"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.cooldown", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.dimensions.0.name", "name-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.dimensions.0.value", "value-1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.threshold", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.operator", "gte"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.evaluation_periods", "10"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.period", "60"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.action_type", "adjustment"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.adjustment", "1"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingDownPolicyGCPGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.policy_name", "policy-name-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.metric_name", "CPUUtilization"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.namespace", "updated-namespace"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.source", "stackdriver"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.statistic", "sum"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.unit", "bytes"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.cooldown", "300"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.dimensions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.dimensions.0.name", "name-1-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.dimensions.0.value", "value-1-update"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.threshold", "5"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.operator", "lte"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.evaluation_periods", "20"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.period", "300"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.action_type", "adjustment"),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.0.adjustment", "2"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testScalingDownPolicyGCPGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scaling_down_policy.#", "0"),
				),
			},
		},
	})
}

const testScalingDownPolicyGCPGroupConfig_Create = `
// --- SCALE DOWN POLICY ------------------
	scaling_down_policy {
	 policy_name = "policy-name"
	 metric_name = "CPUUtilization"
	 namespace = "test-namespace"
	 source = "spectrum"
	 statistic = "count"
	 unit = "seconds"
	 cooldown = 60
	
	 dimensions {
			 name = "name-1"
			 value = "value-1"
	 }
	
	 threshold = 10
	 operator = "gte"
	 evaluation_periods = 10
	 period = 60
	
	 // === ADJUSTMENT ===================
	 action_type = "adjustment"
	 adjustment = 1
	 // ==================================
 }
// ----------------------------------------
`

const testScalingDownPolicyGCPGroupConfig_Update = `
// --- SCALE DOWN POLICY ------------------
	scaling_down_policy {
	 policy_name = "policy-name-update"
	 metric_name = "CPUUtilization"
	 namespace = "updated-namespace"
	 source = "stackdriver"
	 statistic = "sum"
	 unit = "bytes"
	 cooldown = 300
	 threshold = 5
	
	 dimensions {
			 name = "name-1-update"
			 value = "value-1-update"
	 }
	
	 operator = "lte"
	 evaluation_periods = 20
	 period = 300
	
	 // === ADJUSTMENT ===================
	 action_type = "adjustment"
	 adjustment = 2
	 // ==================================
 }
// ----------------------------------------
`

const testScalingDownPolicyGCPGroupConfig_EmptyFields = `
// --- SCALE DOWN POLICY ------------------
// ----------------------------------------
`

// endregion

// region Elastigroup GCP: Subnets
func TestAccSpotinstElastigroupGCP_Subnets(t *testing.T) {
	groupName := "test-acc-eg-gcp-subnets"
	resourceName := createElastigroupGCPResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "gcp") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupGCPDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testSubnetsGCPGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "subnets.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subnets.0.region", "us-central1"),
					resource.TestCheckResourceAttr(resourceName, "subnets.0.subnet_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subnets.0.subnet_names.0", "us-central1-a"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testSubnetsGCPGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "subnets.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "subnets.0.region", "us-central1"),
					resource.TestCheckResourceAttr(resourceName, "subnets.0.subnet_names.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "subnets.0.subnet_names.0", "us-central1-a"),
					resource.TestCheckResourceAttr(resourceName, "subnets.0.subnet_names.1", "us-central1-b"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testSubnetsGCPGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "subnets.#", "0"),
				),
			},
		},
	})
}

const testSubnetsGCPGroupConfig_Create = `
// --- SUBNETS ------------------------------------------
	subnets {
			region = "us-central1"
			subnet_names = ["us-central1-a"]
  }
// ------------------------------------------------------
`

const testSubnetsGCPGroupConfig_Update = `
// --- SUBNETS ------------------------------------------
	subnets {
			region = "us-central1"
			subnet_names = ["us-central1-a", "us-central1-b"]
  }
// ------------------------------------------------------
`

const testSubnetsGCPGroupConfig_EmptyFields = `
// --- SUBNETS ------------------------------------------
// ------------------------------------------------------
`

// endregion

// region Docker Swarm integration

func TestAccSpotinstElastigroupGCP_IntegrationDockerSwarm(t *testing.T) {
	groupName := "test-acc-eg-integration-docker-swarm"
	resourceName := createElastigroupGCPResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "gcp") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupGCPDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testGCPIntegrationDockerSwarmGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.master_host", "docker-swarm-master-host"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.master_port", "8000"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testGCPIntegrationDockerSwarmGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.master_host", "docker-swarm-master-host-update"),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.0.master_port", "9000"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testGCPIntegrationDockerSwarmGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "integration_docker_swarm.#", "0"),
				),
			},
		},
	})
}

const testGCPIntegrationDockerSwarmGroupConfig_Create = `
 // --- INTEGRATION: DOCKER SWARM -------
 integration_docker_swarm {
    master_host = "docker-swarm-master-host"
    master_port = 8000
 }
 // -------------------------------------
`

const testGCPIntegrationDockerSwarmGroupConfig_Update = `
 // --- INTEGRATION: DOCKER SWARM -------
 integration_docker_swarm {
	master_host = "docker-swarm-master-host-update"
    master_port = 9000
  }
 // -------------------------------------
`

const testGCPIntegrationDockerSwarmGroupConfig_EmptyFields = `
 // --- INTEGRATION: DOCKER SWARM -------
 // -------------------------------------
`

// endregion

// region Elastigroup: Scheduled Tasks
func TestAccSpotinstElastigroupGCP_ScheduledTask(t *testing.T) {
	groupName := "test-acc-eg-gcp-scheduled-task"
	resourceName := createElastigroupGCPResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t, "gcp") },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupGCPDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testGCPScheduledTaskGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.task_type", "setCapacity"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.cron_expression", "* * * * *"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.min_capacity", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.max_capacity", "3"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.target_capacity", "2"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testGCPScheduledTaskGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.task_type", "setCapacity"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.cron_expression", "* * * * *"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.min_capacity", "2"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.max_capacity", "4"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.0.target_capacity", "3"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGCPTerraform(&GCPGroupConfigMetadata{
					groupName:      groupName,
					fieldsToAppend: testGCPScheduledTaskGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGCPExists(&group, resourceName),
					testCheckElastigroupGCPAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "scheduled_task.#", "0"),
				),
			},
		},
	})
}

const testGCPScheduledTaskGroupConfig_Create = `
 // --- SCHEDULED TASK ------------------
  scheduled_task {
		is_enabled = false
    task_type = "setCapacity"
    cron_expression = "* * * * *"
    target_capacity = 2
    min_capacity = 1
    max_capacity = 3
  }
 // -------------------------------------
`

const testGCPScheduledTaskGroupConfig_Update = `
 // --- SCHEDULED TASK ------------------
  scheduled_task {
		is_enabled = true
    task_type = "setCapacity"
    cron_expression = "* * * * *"
    target_capacity = 3
    min_capacity = 2
    max_capacity = 4
  }
 // -------------------------------------
`

const testGCPScheduledTaskGroupConfig_EmptyFields = `
 // --- SCHEDULED TASK ------------------
 // -------------------------------------
`

// endregion

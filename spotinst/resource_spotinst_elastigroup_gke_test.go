package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"log"
	"testing"
)

// createElastigroupGKEResourceName creates a resource name for the test group
func createElastigroupGKEResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.ElastigroupGKEResourceName), name)
}

// testElastigroupGKEDestroy checks whether a group has been destroyed and returns an error if it still exists
func testElastigroupGKEDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.ElastigroupGKEResourceName) {
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

// testCheckElastigroupGKEAttributes checks the correct group is being tests and returns an error if it is not
func testCheckElastigroupGKEAttributes(group *gcp.Group, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(group.Name) != expectedName {
			return fmt.Errorf("bad content: %v", group.Name)
		}
		return nil
	}
}

// testCheckElastigroupGKEExists checks if a group exists and returns an error if the resource isn't found, if the
// group id was not set, or the group does not exist (wasn't created or was unexpectedly deleted, etc)
func testCheckElastigroupGKEExists(group *gcp.Group, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProvider.Meta().(*Client)
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

// GKEGroupConfigMetadata holds blocks of attributes defined as strings that are used to build a Terraform resource
type GKEGroupConfigMetadata struct {
	variables            string
	groupName            string
	instanceTypes        string
	strategy             string
	fieldsToAppend       string
	updateBaselineFields bool
}

// createElastigroupGKETerraform builds a valid Terraform resource as a string.
// This function appends attribute blocks defined as string later in this file.
// These blocks should have fields required for a bare-minimum group to be created.
func createElastigroupGKETerraform(gcm *GKEGroupConfigMetadata) string {
	//os.Setenv("SPOTINST_ACCOUNT", "act-cca49da4")
	//os.Setenv("SPOTINST_TOKEN", "")
	if gcm == nil {
		return ""
	}

	if gcm.instanceTypes == "" {
		gcm.instanceTypes = testInstanceTypesGKEGroupConfig_Create
	}

	if gcm.strategy == "" {
		gcm.strategy = testStrategyGKEGroupConfig_Create
	}

	template := ""
	if gcm.updateBaselineFields {
		format := testBaselineGKEGroupConfig_Update

		template = fmt.Sprintf(format,
			gcm.groupName,
			gcm.groupName,
			gcm.instanceTypes,
			gcm.strategy,
			gcm.fieldsToAppend,
		)
	} else {
		format := testBaselineGKEGroupConfig_Create

		template = fmt.Sprintf(format,
			gcm.groupName,
			gcm.groupName,
			gcm.instanceTypes,
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

// region Elastigroup GKE: Baseline
func TestAccSpotinstElastigroupGKE_Baseline(t *testing.T) {
	groupName := "eg-gke-baseline"
	resourceName := createElastigroupGKEResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupGKEDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupGKETerraform(&GKEGroupConfigMetadata{groupName: groupName}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGKEExists(&group, resourceName),
					testCheckElastigroupGKEAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "0"),
					resource.TestCheckResourceAttr(resourceName, "node_image", "COS"),
				),
			},
			{
				Config: createElastigroupGKETerraform(&GKEGroupConfigMetadata{groupName: groupName, updateBaselineFields: true}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGKEExists(&group, resourceName),
					testCheckElastigroupGKEAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "max_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "min_size", "0"),
					resource.TestCheckResourceAttr(resourceName, "desired_capacity", "0"),
					resource.TestCheckResourceAttr(resourceName, "node_image", "COS"),
				),
			},
		},
	})
}

const testBaselineGKEGroupConfig_Create = `
resource "` + string(commons.ElastigroupGKEResourceName) + `" "%v" {

 name = "%v"
 cluster_id = "terraform-acc-test-cluster"
 cluster_zone_name = "us-central1-a"
 node_image = "COS"

 // --- CAPACITY ------------
 max_size = 0
 min_size = 0
 desired_capacity = 0
 // -------------------------
 
 %v
 %v
 %v
}

`

const testBaselineGKEGroupConfig_Update = `
resource "` + string(commons.ElastigroupGKEResourceName) + `" "%v" {

 name = "%v"
 cluster_id = "terraform-acc-test-cluster"
 cluster_zone_name = "us-central1-a"
 node_image = "COS"

 // --- CAPACITY ------------
 max_size = 0
 min_size = 0
 desired_capacity = 0
 // -------------------------

 %v
 %v
 %v
}

`

// endregion

// region Elastigroup GKE: Instance Types
func TestAccSpotinstElastigroupGKE_InstanceTypes(t *testing.T) {
	groupName := "eg-gke-instance-types"
	resourceName := createElastigroupGKEResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testElastigroupGKEDestroy,

		Steps: []resource.TestStep{
			{
				Config: createElastigroupGKETerraform(&GKEGroupConfigMetadata{
					groupName:     groupName,
					instanceTypes: testInstanceTypesGKEGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGKEExists(&group, resourceName),
					testCheckElastigroupGKEAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "instance_types_ondemand", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preemptible.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preemptible.0", "n1-standard-1"),
					resource.TestCheckResourceAttr(resourceName, "instance_types_preemptible.1", "n1-standard-2"),
				),
			},
			{
				Config: createElastigroupGKETerraform(&GKEGroupConfigMetadata{
					groupName:     groupName,
					instanceTypes: testInstanceTypesGKEGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGKEExists(&group, resourceName),
					testCheckElastigroupGKEAttributes(&group, groupName),
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

const testInstanceTypesGKEGroupConfig_Create = `
 // --- INSTANCE TYPES --------------------------------
 instance_types_ondemand = "n1-standard-1"
 instance_types_preemptible = ["n1-standard-1", "n1-standard-2"]
 // ---------------------------------------------------
`

const testInstanceTypesGKEGroupConfig_Update = `
 // --- INSTANCE TYPES --------------------------------
 instance_types_ondemand = "n1-standard-2"
 instance_types_preemptible = ["n1-standard-1", "n1-standard-2", "n1-standard-4"]
 // ---------------------------------------------------
`

// endregion

// region Elastigroup GKE: Strategy
func TestAccSpotinstElastigroupGKE_Strategy(t *testing.T) {
	groupName := "eg-gcp-strategy"
	resourceName := createElastigroupGKEResourceName(groupName)

	var group gcp.Group
	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		Providers:     TestAccProviders,
		CheckDestroy:  testElastigroupGKEDestroy,
		IDRefreshName: resourceName,

		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config: createElastigroupGKETerraform(&GKEGroupConfigMetadata{
					groupName: groupName,
					strategy:  testStrategyGKEGroupConfig_Create,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGKEExists(&group, resourceName),
					testCheckElastigroupGKEAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "preemptible_percentage", "100"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGKETerraform(&GKEGroupConfigMetadata{
					groupName: groupName,
					strategy:  testStrategyGKEGroupConfig_Update,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGKEExists(&group, resourceName),
					testCheckElastigroupGKEAttributes(&group, groupName),
					resource.TestCheckResourceAttr(resourceName, "preemptible_percentage", "75"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createElastigroupGKETerraform(&GKEGroupConfigMetadata{
					groupName: groupName,
					strategy:  testStrategyGKEGroupConfig_EmptyFields,
				}),
				Check: resource.ComposeTestCheckFunc(
					testCheckElastigroupGKEExists(&group, resourceName),
					testCheckElastigroupGKEAttributes(&group, groupName),
				),
			},
		},
	})
}

const testStrategyGKEGroupConfig_Create = `
 // --- STRATEGY --------------------
 preemptible_percentage = 100
 // ---------------------------------
`

const testStrategyGKEGroupConfig_Update = `
 // --- STRATEGY --------------------
  preemptible_percentage = 75
 // ---------------------------------
`

const testStrategyGKEGroupConfig_EmptyFields = `
 // --- STRATEGY --------------------
 // ---------------------------------
`

// endregion

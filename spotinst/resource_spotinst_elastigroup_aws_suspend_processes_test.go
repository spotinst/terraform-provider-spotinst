package spotinst

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func createSuspendProcessesResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.SuspendProcessesResourceName), name)
}

func testSuspendProcessesDestroy(s *terraform.State) error {
	client := testAccProviderAWS.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.SuspendProcessesResourceName) {
			continue
		}
		input := &aws.DeleteSuspensionsInput{GroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.elastigroup.CloudProviderAWS().DeleteSuspensions(context.Background(), input)
		if err == nil && resp != nil {
			return fmt.Errorf("suspensions still exist")
		}
	}
	return nil
}

func testSuspendProcessesAttributes(spWrapper *commons.SuspendProcessesWrapper, expectedName string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		if spotinst.StringValue(spWrapper.GroupID) != expectedName {
			return fmt.Errorf("bad content: %+v", spWrapper.GroupID)
		}
		return nil
	}
}

func testCheckSuspendProcessesExists(suspendProcesses *aws.SuspendProcesses, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAWS.Meta().(*Client)
		input := &aws.ListSuspensionsInput{GroupID: spotinst.String(rs.Primary.ID)}
		resp, err := client.elastigroup.CloudProviderAWS().ListSuspensions(context.Background(), input)
		if err != nil {
			return err
		}
		if spotinst.StringValue(input.GroupID) != rs.Primary.Attributes["id"] {
			return fmt.Errorf("suspend proceeses not found: %+v,\n %+v\n", resp.SuspendProcesses, rs.Primary.Attributes)
		}
		suspendProcesses = resp.SuspendProcesses
		return nil
	}
}

type SuspendProcessesMetadata struct {
	provider             string
	groupID              string
	variables            string
	fieldsToAppend       string
	updateBaselineFields bool
}

func createSuspendProcessesTerraform(ccm *SuspendProcessesMetadata, formatToUse string) string {
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
		format := testBaselineSuspendProcessesConfig_Update
		template += fmt.Sprintf(format,
			ccm.groupID,
			ccm.provider,
			ccm.fieldsToAppend,
		)
	} else {
		format := testBaselineSuspendProcessesConfig_Create
		template += fmt.Sprintf(format,
			ccm.groupID,
			ccm.provider,
			ccm.fieldsToAppend,
		)
	}

	if ccm.variables != "" {
		template = ccm.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", "suspend_processes_test", template)
	return template
}

// region SuspendProcesses: Baseline
func TestAccSpotinstSuspendProcessesElastigroupAWS_Baseline(t *testing.T) {
	groupId := "sig-05d0a009"
	resourceName := createSuspendProcessesResourceName(groupId)

	var spWrapper commons.SuspendProcessesWrapper
	spWrapper.GroupID = &groupId

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "aws") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testSuspendProcessesDestroy,

		Steps: []resource.TestStep{
			{
				Config: createSuspendProcessesTerraform(&SuspendProcessesMetadata{
					groupID: groupId,
				}, testBaselineSuspendProcessesConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckSuspendProcessesExists(spWrapper.SuspendProcesses, resourceName),
					testSuspendProcessesAttributes(&spWrapper, groupId),
					resource.TestCheckResourceAttr(resourceName, "group_id", "sig-9f6d7870"),
					resource.TestCheckResourceAttr(resourceName, "suspension.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "suspension.0.name", "SCHEDULING"),
					resource.TestCheckResourceAttr(resourceName, "suspension.1.name", "PREVENTIVE_REPLACEMENT"),
					resource.TestCheckResourceAttr(resourceName, "suspension.2.name", "OUT_OF_STRATEGY"),
					resource.TestCheckResourceAttr(resourceName, "suspension.3.name", "REVERT_PREFERRED"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createSuspendProcessesTerraform(&SuspendProcessesMetadata{
					groupID:              groupId,
					updateBaselineFields: true}, testBaselineSuspendProcessesConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckSuspendProcessesExists(spWrapper.SuspendProcesses, resourceName),
					testSuspendProcessesAttributes(&spWrapper, groupId),
					resource.TestCheckResourceAttr(resourceName, "group_id", "sig-9f6d7870"),
					resource.TestCheckResourceAttr(resourceName, "suspension.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "suspension.0.name", "OUT_OF_STRATEGY"),
					resource.TestCheckResourceAttr(resourceName, "suspension.1.name", "REVERT_PREFERRED"),
					resource.TestCheckResourceAttr(resourceName, "suspension.2.name", "AUTO_HEALING"),
				),
			},
		},
	})
}

const testBaselineSuspendProcessesConfig_Create = `
resource "` + string(commons.SuspendProcessesResourceName) + `" "%v" {
  provider = "%v"

	group_id = "sig-9f6d7870"
	suspension  {
    	name = "SCHEDULING"
  	}
  	suspension  {
    	name = "PREVENTIVE_REPLACEMENT"
  	}
	suspension {
    	name = "OUT_OF_STRATEGY"
  	}
  	suspension {
    	name = "REVERT_PREFERRED"
  	}
 %v
}
`

const testBaselineSuspendProcessesConfig_Update = `
resource "` + string(commons.SuspendProcessesResourceName) + `" "%v" {
	provider = "%v"

 	group_id = "sig-9f6d7870"
  	suspension  {
    	name = "OUT_OF_STRATEGY"
	  }
  	suspension  {
    	name = "REVERT_PREFERRED"
  	}
  	suspension {
    	name = "AUTO_HEALING"
  	}
  %v
}
`

// endregion

// region SuspendProcesses: RemoveSuspension
func TestAccSpotinstSuspendProcessesElastigroupAWS_RemoveSuspension(t *testing.T) {
	groupId := "sig-05d0a009"
	resourceName := createSuspendProcessesResourceName(groupId)

	var spWrapper commons.SuspendProcessesWrapper
	spWrapper.GroupID = &groupId

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t, "aws") },
		ProviderFactories: TestAccProviders,
		CheckDestroy:      testSuspendProcessesDestroy,

		Steps: []resource.TestStep{
			{
				Config: createSuspendProcessesTerraform(&SuspendProcessesMetadata{
					groupID: groupId,
				}, testRemoveSuspensionSuspendProcessesConfig_Create),
				Check: resource.ComposeTestCheckFunc(
					testCheckSuspendProcessesExists(spWrapper.SuspendProcesses, resourceName),
					testSuspendProcessesAttributes(&spWrapper, groupId),
					resource.TestCheckResourceAttr(resourceName, "group_id", "sig-9f6d7870"),
					resource.TestCheckResourceAttr(resourceName, "suspension.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "suspension.0.name", "SCHEDULING"),
					resource.TestCheckResourceAttr(resourceName, "suspension.1.name", "PREVENTIVE_REPLACEMENT"),
					resource.TestCheckResourceAttr(resourceName, "suspension.2.name", "OUT_OF_STRATEGY"),
					resource.TestCheckResourceAttr(resourceName, "suspension.3.name", "REVERT_PREFERRED"),
				),
			},
			{
				ResourceName: resourceName,
				Config: createSuspendProcessesTerraform(&SuspendProcessesMetadata{
					groupID:              groupId,
					updateBaselineFields: true}, testRemoveSuspensionSuspendProcessesConfig_Update),
				Check: resource.ComposeTestCheckFunc(
					testCheckSuspendProcessesExists(spWrapper.SuspendProcesses, resourceName),
					testSuspendProcessesAttributes(&spWrapper, groupId),
					resource.TestCheckResourceAttr(resourceName, "group_id", "sig-9f6d7870"),
					resource.TestCheckResourceAttr(resourceName, "suspension.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "suspension.0.name", "OUT_OF_STRATEGY"),
					resource.TestCheckResourceAttr(resourceName, "suspension.1.name", "REVERT_PREFERRED"),
					resource.TestCheckResourceAttr(resourceName, "suspension.2.name", "AUTO_HEALING"),
				),
			},
		},
	})
}

const testRemoveSuspensionSuspendProcessesConfig_Create = `
resource "` + string(commons.SuspendProcessesResourceName) + `" "%v" {
  provider = "%v"

	group_id = "sig-9f6d7870"
	suspension  {
    	name = "OUT_OF_STRATEGY"
  	}
  	suspension  {
    	name = "PREVENTIVE_REPLACEMENT"
  	}
	suspension {
    	name = "SCHEDULING"
  	}
  	suspension {
    	name = "REVERT_PREFERRED"
  	}
 %v
}
`

const testRemoveSuspensionSuspendProcessesConfig_Update = `
resource "` + string(commons.SuspendProcessesResourceName) + `" "%v" {
	provider = "%v"

 	group_id = "sig-9f6d7870"
  	suspension  {
    	name = "SCHEDULING"
  	}
  	suspension  {
    	name = "REVERT_PREFERRED"
  	}
	suspension {
    	name = "AUTO_HEALING"
  	}
  %v
}
`

// endregion

package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"log"
	"strings"
	"testing"
)

func init() {
	resource.AddTestSweepers("resource_spotinst_stateful_node_azure_v3", &resource.Sweeper{
		Name: "spotinst_stateful_node_azure_v3",
		F:    testSweepStatefulNodeAzureV3,
	})
}

func testSweepStatefulNodeAzureV3(region string) error {
	client, err := getProviderClient("azure")
	if err != nil {
		return fmt.Errorf("error getting client: %v", err)
	}

	conn := client.(*Client).statefulNode.CloudProviderAzure()

	input := &azurev3.ListStatefulNodesInput{}
	if resp, err := conn.List(context.Background(), input); err != nil {
		return fmt.Errorf("error getting list of nodes to sweep")
	} else {
		if len(resp.StatefulNodes) == 0 {
			log.Printf("[INFO] No clusters to sweep")
		}
		for _, statefulNode := range resp.StatefulNodes {
			if strings.Contains(spotinst.StringValue(statefulNode.Name), "terraform-acc-tests-") {
				if _, err := conn.Delete(context.Background(), &azurev3.DeleteStatefulNodeInput{ID: statefulNode.ID}); err != nil {
					return fmt.Errorf("unable to delete nodes %v in sweep", spotinst.StringValue(statefulNode.ID))
				} else {
					log.Printf("Sweeper deleted %v\n", spotinst.StringValue(statefulNode.ID))
				}
			}
		}
	}
	return nil
}

func createStatefulNodeAzureV3ResourceName(name string) string {
	return fmt.Sprintf("%v.%v", string(commons.StatefulNodeAzureV3ResourceName), name)
}

func testStatefulNodeAzureV3Destroy(s *terraform.State) error {
	client := testAccProviderAzureV3.Meta().(*Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != string(commons.StatefulNodeAzureV3ResourceName) {
			continue
		}
		input := &azurev3.ReadStatefulNodeInput{ID: spotinst.String(rs.Primary.ID)}
		resp, err := client.statefulNode.CloudProviderAzure().Read(context.Background(), input)
		if err == nil && resp != nil && resp.StatefulNode != nil {
			return fmt.Errorf("stateful Node still exists")
		}
	}
	return nil
}

func testCheckStatefulNodeAzureV3Attributes(statefulNode *azurev3.StatefulNode, expectedName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if spotinst.StringValue(statefulNode.Name) != expectedName {
			return fmt.Errorf("bad content: %v", statefulNode.Name)
		}
		return nil
	}
}

func testCheckStatefulNodeAzureV3Exists(statefulNode *azurev3.StatefulNode, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no resource ID is set")
		}
		client := testAccProviderAzure.Meta().(*Client)
		input := &azurev3.ReadStatefulNodeInput{ID: spotinst.String(rs.Primary.ID)}
		resp, err := client.statefulNode.CloudProviderAzure().Read(context.Background(), input)
		if err != nil {
			return err
		}
		*statefulNode = *resp.StatefulNode
		return nil
	}
}

type AzureV3StatefulNodeConfigMetadata struct {
	statefulNodeName     string
	acdIdentifier        string
	controllerClusterID  string
	provider             string
	launchSpecification  string
	strategy             string
	autoScaler           string
	health               string
	vmSizes              string
	osDisk               string
	image                string
	loadBalancers        string
	network              string
	extensions           string
	login                string
	persistence          string
	secret               string
	scheduling           string
	variables            string
	updateBaselineFields bool
}

func createStatefulNodeAzureV3Terraform(StatefulNodeMeta *AzureV3StatefulNodeConfigMetadata) string {
	if StatefulNodeMeta == nil {
		return ""
	}

	if StatefulNodeMeta.provider == "" {
		StatefulNodeMeta.provider = "azure"
	}

	if StatefulNodeMeta.launchSpecification == "" {
		StatefulNodeMeta.launchSpecification = testLaunchSpecificationStatefulNodeAzureV3Config_Create
	}

	if StatefulNodeMeta.strategy == "" {
		StatefulNodeMeta.strategy = testStrategyStatefulNodeAzureV3Config_Create
	}

	if StatefulNodeMeta.health == "" {
		StatefulNodeMeta.health = testHealthStatefulNodeAzureV3Config_Create
	}

	if StatefulNodeMeta.vmSizes == "" {
		StatefulNodeMeta.vmSizes = testVMSizesStatefulNodeAzureV3Config_Create
	}

	if StatefulNodeMeta.image == "" {
		StatefulNodeMeta.image = testImageStatefulNodeAzureV3Config_Create
	}

	if StatefulNodeMeta.loadBalancers == "" {
		StatefulNodeMeta.loadBalancers = testLoadBalancersStatefulNodeAzureV3Config_Create
	}

	if StatefulNodeMeta.network == "" {
		StatefulNodeMeta.network = testNetworkStatefulNodeAzureV3Config_Create
	}

	if StatefulNodeMeta.extensions == "" {
		StatefulNodeMeta.extensions = testExtensionsStatefulNodeAzureV3Config_Create
	}

	if StatefulNodeMeta.login == "" {
		StatefulNodeMeta.login = testAzureV3LoginStatefulNodeConfig_Create
	}

	if StatefulNodeMeta.persistence == "" {
		StatefulNodeMeta.persistence = testPersistenceStatefulNodeAzureV3Config_Create
	}

	if StatefulNodeMeta.scheduling == "" {
		StatefulNodeMeta.scheduling = testSchedulingStatefulNodeAzureV3Config_Create
	}

	if StatefulNodeMeta.secret == "" {
		StatefulNodeMeta.secret = testSecretsStatefulNodeAzureV3Config_Create
	}

	template :=
		`provider "azure" {
	token   = "fake"
	account = "fake"
	}
	`
	if StatefulNodeMeta.updateBaselineFields {
		format := testBaselineStatefulNodeAzureV3Config_Update
		template += fmt.Sprintf(format,
			StatefulNodeMeta.statefulNodeName,
			StatefulNodeMeta.provider,
			StatefulNodeMeta.login,
			StatefulNodeMeta.launchSpecification,
			StatefulNodeMeta.osDisk,
			StatefulNodeMeta.strategy,
			StatefulNodeMeta.image,
			StatefulNodeMeta.extensions,
			StatefulNodeMeta.network,
			StatefulNodeMeta.health,
			StatefulNodeMeta.loadBalancers,
			StatefulNodeMeta.vmSizes,
			StatefulNodeMeta.persistence,
			StatefulNodeMeta.scheduling,
			StatefulNodeMeta.secret,
		)
	} else {
		format := testBaselineStatefulNodeAzureV3Config_Create
		template += fmt.Sprintf(format,
			StatefulNodeMeta.statefulNodeName,
			StatefulNodeMeta.provider,
			StatefulNodeMeta.login,
			StatefulNodeMeta.launchSpecification,
			StatefulNodeMeta.osDisk,
			StatefulNodeMeta.strategy,
			StatefulNodeMeta.image,
			StatefulNodeMeta.extensions,
			StatefulNodeMeta.network,
			StatefulNodeMeta.health,
			StatefulNodeMeta.loadBalancers,
			StatefulNodeMeta.vmSizes,
			StatefulNodeMeta.persistence,
			StatefulNodeMeta.scheduling,
			StatefulNodeMeta.secret,
		)

	}

	if StatefulNodeMeta.variables != "" {
		template = StatefulNodeMeta.variables + "\n" + template
	}

	log.Printf("Terraform [%v] template:\n%v", StatefulNodeMeta.statefulNodeName, template)
	return template
}

// region Stateful Node Azure: Baseline
func TestAccSpotinstStatefulNodeAzureV3_Baseline(t *testing.T) {
	statefulNodeName := "test-acc-sn-azure-v3-baseline" // what values to insert?
	resourceName := createStatefulNodeAzureV3ResourceName(statefulNodeName)

	var node azurev3.StatefulNode
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testStatefulNodeAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config:             createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{statefulNodeName: statefulNodeName}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckStatefulNodeAzureV3Exists(&node, resourceName),
					testCheckStatefulNodeAzureV3Attributes(&node, statefulNodeName),
					resource.TestCheckResourceAttr(resourceName, "region", "eastus"),
					resource.TestCheckResourceAttr(resourceName, "description", "tamir-test-file-1"),
					resource.TestCheckResourceAttr(resourceName, "os", "Linux"),
				),
			},
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{
					statefulNodeName:     statefulNodeName,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "region", "eastus"),
					resource.TestCheckResourceAttr(resourceName, "description", "tamir-test-file-1-updated"),
					resource.TestCheckResourceAttr(resourceName, "os", "Linux"),
				),
			},
		},
	})
}

const testBaselineStatefulNodeAzureV3Config_Create = `
resource "` + string(commons.StatefulNodeAzureV3ResourceName) + `" "%v" {
  provider = "%v"

 name 				 = "%v"
 os 			     = "Linux"
 region              = "eastus"
 description = "tamir-test-file-1"
 
 %v
 %v
 %v
 %v
 %v
 %v
 %v
 %v
 %v
 %v
 %v
 %v
}
`

const testBaselineStatefulNodeAzureV3Config_Update = `
resource "` + string(commons.StatefulNodeAzureV3ResourceName) + `" "%v" {
  
  provider = "%v"
 name 				 = "%v"
 os 			     = "Linux"
 region              = "eastus"
 description = "tamir-test-file-1-updated"

 %v
 %v
 %v
 %v
 %v
 %v
 %v
 %v
 %v
 %v
 %v
 %v
}
`

//endregion

// region Stateful Node Azure : Login
func TestAccSpotinstStatefulNodeAzureV3_Login(t *testing.T) {
	statefulNodeName := "terraform-tests-do-not-delete"
	resourceName := createStatefulNodeAzureV3ResourceName(statefulNodeName)

	var node azurev3.StatefulNode
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testStatefulNodeAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config:             createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{statefulNodeName: statefulNodeName}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckStatefulNodeAzureV3Exists(&node, resourceName),
					testCheckStatefulNodeAzureV3Attributes(&node, statefulNodeName),
					resource.TestCheckResourceAttr(resourceName, "login.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "login.0.user_name", "azure_v3_terraform"),
					resource.TestCheckResourceAttr(resourceName, "login.0.password", "123456789"),
				),
			},
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{
					statefulNodeName:     statefulNodeName,
					login:                testAzureV3LoginStatefulNodeConfig_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "login.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "login.0.user_name", "azure_v3_terraform"),
					resource.TestCheckResourceAttr(resourceName, "login.0.password", "111111111"),
				),
			},
		},
	})
}

const testAzureV3LoginStatefulNodeConfig_Create = `
  login {
    user_name = "azure_v3_terraform"
	password  = "123456789"
  }
`

const testAzureV3LoginStatefulNodeConfig_Update = `
  login {
    user_name = "azure_v3_terraform"
	password  = "111111111"
  }
`

//endregion

// region Stateful Node Azure : Launch Specification
func TestAccSpotinstStatefulNodeAzureV3_LaunchSpecification(t *testing.T) {
	statefulNodeName := "terraform-tests-do-not-delete"
	resourceName := createStatefulNodeAzureV3ResourceName(statefulNodeName)

	var node azurev3.StatefulNode
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testStatefulNodeAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config: createStatefulNodeAzureV3Terraform(
					&AzureV3StatefulNodeConfigMetadata{
						statefulNodeName: statefulNodeName,
					}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckStatefulNodeAzureV3Exists(&node, resourceName),
					testCheckStatefulNodeAzureV3Attributes(&node, statefulNodeName),
					resource.TestCheckResourceAttr(resourceName, "custom_data", "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"),
					resource.TestCheckResourceAttr(resourceName, "shutdown_script", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.1113301976.tag_key", "tag_key"),
					resource.TestCheckResourceAttr(resourceName, "tags.1113301976.tag_value", "tag_value"),
					resource.TestCheckResourceAttr(resourceName, "managed_service_identities.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_service_identities.name", "foo"),
					resource.TestCheckResourceAttr(resourceName, "managed_service_identities.resource_group_name", "foo2"),
					resource.TestCheckResourceAttr(resourceName, "os_disk.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "os_disk.0.size_gb", "0"),
					resource.TestCheckResourceAttr(resourceName, "os_disk.0.type", "SSD"),
					resource.TestCheckResourceAttr(resourceName, "data_disks.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_disks.0.lun", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_disks.0.size_gb", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_disks.0.type", "SSD"),
					resource.TestCheckResourceAttr(resourceName, "boot_diagnostics.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "boot_diagnostics.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "boot_diagnostics.0.storage_url", "3"),
					resource.TestCheckResourceAttr(resourceName, "boot_diagnostics.0.type", "0"),
				),
			},
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{
					statefulNodeName:     statefulNodeName,
					launchSpecification:  testLaunchSpecificationStatefulNodeAzureV3Config_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "custom_data", "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"),
					resource.TestCheckResourceAttr(resourceName, "shutdown_script", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.1113301976.tag_key", "tag_key"),
					resource.TestCheckResourceAttr(resourceName, "tags.1113301976.tag_value", "tag_value"),
					resource.TestCheckResourceAttr(resourceName, "managed_service_identities.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "managed_service_identities.name", "foo"),
					resource.TestCheckResourceAttr(resourceName, "managed_service_identities.resource_group_name", "foo2"),
					resource.TestCheckResourceAttr(resourceName, "os_disk.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "os_disk.0.size_gb", "0"),
					resource.TestCheckResourceAttr(resourceName, "os_disk.0.type", "HDD"),
					resource.TestCheckResourceAttr(resourceName, "data_disks.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_disks.0.lun", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_disks.0.size_gb", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_disks.0.type", "HDD"),
					resource.TestCheckResourceAttr(resourceName, "boot_diagnostics.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "boot_diagnostics.0.is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "boot_diagnostics.0.storage_url", "3"),
					resource.TestCheckResourceAttr(resourceName, "boot_diagnostics.0.type", "0"),
				),
			},
		},
	})
}

const testLaunchSpecificationStatefulNodeAzureV3Config_Create = `
	custom_data = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"
	shutdown_script = "1"
	tag {
		tag_key = "tag_key"
		tag_value = "tag_value"
	}
	managed_service_identities{
		name = "foo"
		resource_group_name = "foo2"
	}
	os_disk{
		size_gb = "0"
		type = "SSD"
	}
	data_disks{
		lun = "1"
		size_gb = "1"
		type = "SSD"
	}
	boot_diagnostics{
		is_enabled = true
		storage_url = "3"
		type = "0"
	}
`

const testLaunchSpecificationStatefulNodeAzureV3Config_Update = `
	custom_data = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"
	shutdown_script = "1"
	tag {
		tag_key = "tag_key"
		tag_value = "tag_value"
	}
	managed_service_identities{
		name = "foo"
		resource_group_name = "foo2"
	}
	os_disk{
		size_gb = "0"
		type = "HDD"
	}
	data_disks{
		lun = "1"
		size_gb = "1"
		type = "HDD"
	}
	boot_diagnostics{
		is_enabled = false
		storage_url = "3"
		type = "0"
	}
`

//endregion

// region Stateful Node Azure : Persistence
func TestAccSpotinstStatefulNodeAzureV3_Persistence(t *testing.T) {
	statefulNodeName := "terraform-tests-do-not-delete"
	resourceName := createStatefulNodeAzureV3ResourceName(statefulNodeName)

	var node azurev3.StatefulNode
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testStatefulNodeAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config:             createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{statefulNodeName: statefulNodeName}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckStatefulNodeAzureV3Exists(&node, resourceName),
					testCheckStatefulNodeAzureV3Attributes(&node, statefulNodeName),
					resource.TestCheckResourceAttr(resourceName, "should_persist_os_disk", "true"),
					resource.TestCheckResourceAttr(resourceName, "os_disk_persistence_mode", "reattach"),
					resource.TestCheckResourceAttr(resourceName, "should_persist_data_disks", "true"),
					resource.TestCheckResourceAttr(resourceName, "data_disks_persistence_mode", "reattach"),
					resource.TestCheckResourceAttr(resourceName, "should_persist_network", "true"),
					resource.TestCheckResourceAttr(resourceName, "should_persist_vm", "true"),
				),
			},
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{
					statefulNodeName:     statefulNodeName,
					persistence:          testPersistenceStatefulNodeAzureV3Config_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "should_persist_os_disk", "true"),
					resource.TestCheckResourceAttr(resourceName, "os_disk_persistence_mode", "attach"),
					resource.TestCheckResourceAttr(resourceName, "should_persist_data_disks", "false"),
					resource.TestCheckResourceAttr(resourceName, "data_disks_persistence_mode", "reattach"),
					resource.TestCheckResourceAttr(resourceName, "should_persist_network", "true"),
					resource.TestCheckResourceAttr(resourceName, "should_persist_vm", "true"),
				),
			},
		},
	})
}

const testPersistenceStatefulNodeAzureV3Config_Create = `
    persistence {
        should_persist_os_disk = true
		os_disk_persistence_mode = "reattach"
		should_persist_data_disks = true
		data_disks_persistence_mode = "reattach"
		should_persist_network = true
		should_persist_vm = true
	}
`

const testPersistenceStatefulNodeAzureV3Config_Update = `
persistence {
        should_persist_os_disk = true
		os_disk_persistence_mode = "attach"
		should_persist_data_disks = false
		data_disks_persistence_mode = "reattach"
		should_persist_network = true
		should_persist_vm = true
	}
`

//endregion

// region Stateful Node Azure : Strategy
func TestAccSpotinstStatefulNodeAzureV3_Strategy(t *testing.T) {
	statefulNodeName := "test-acc-sn-azure-v3-strategy"
	resourceName := createStatefulNodeAzureV3ResourceName(statefulNodeName)

	var node azurev3.StatefulNode
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testStatefulNodeAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{statefulNodeName: statefulNodeName,
					strategy: testStrategyStatefulNodeAzureV3Config_Update,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckStatefulNodeAzureV3Exists(&node, resourceName),
					testCheckStatefulNodeAzureV3Attributes(&node, statefulNodeName),
					resource.TestCheckResourceAttr(resourceName, "strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.fallback_to_on_demand", "true"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.draining_timeout", "40"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.optimization_windows", "Tue:19:46-Tue:20:46"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.preferred_life_cycle", "3"), //is that needed?
					resource.TestCheckResourceAttr(resourceName, "strategy.0.signals.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.signals.0.type", "vmReady"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.signals.0.timeout", "20"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.revert_to_spot.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.revert_to_spot.0.perform_at", "timeWindow"),
				),
			},
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{
					statefulNodeName:     statefulNodeName,
					strategy:             testStrategyStatefulNodeAzureV3Config_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "strategy.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.fallback_to_on_demand", "true"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.draining_timeout", "40"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.optimization_windows", "Thu:19:47-Thu:20:46"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.preferred_life_cycle", "3"), //is that needed?
					resource.TestCheckResourceAttr(resourceName, "strategy.0.signals.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.signals.0.type", "vmReady"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.signals.0.timeout", "25"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.revert_to_spot.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "strategy.0.revert_to_spot.0.perform_at", "timeWindow"),
				),
			},
		},
	})
}

const testStrategyStatefulNodeAzureV3Config_Create = `
      strategy {
            signal = [
                {
                    type = "vmReady"
                    timeout = "25"
                }
            ]
            fallback_to_on_demand = true
            draining_timeout =  "40"
			preferred_life_cycle = "3"
            revert_to_spot {
                performAt =  "timeWindow"
            }
            optimizationWindows = [
                "Thu:19:47-Thu:20:46"
            ]
        }
`

const testStrategyStatefulNodeAzureV3Config_Update = `
      strategy: {
            signal: [
                {
                    type: "vmReady"
                    timeout: 20
                }
            ]
            fallback_to_on_demand = true
            draining_timeout = 40
			preferred_life_cycle = 3
            revert_to_spot = {
                performAt = "timeWindow"
            }
            optimizationWindows = [
                "Tue:19:46-Tue:20:46"
            ]
        }
`

//endregion

// region Stateful Node Azure : Health
func TestAccSpotinstStatefulNodeAzureV3_Health(t *testing.T) {
	statefulNodeName := "terraform-tests-do-not-delete"
	resourceName := createStatefulNodeAzureV3ResourceName(statefulNodeName)

	var node azurev3.StatefulNode
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testStatefulNodeAzureV3Destroy,
		Steps: []resource.TestStep{
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{statefulNodeName: statefulNodeName,
					health: testHealthStatefulNodeAzureV3Config_Update,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckStatefulNodeAzureV3Exists(&node, resourceName),
					testCheckStatefulNodeAzureV3Attributes(&node, statefulNodeName),
					resource.TestCheckResourceAttr(resourceName, "health.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health.0.grace_period", "10"),
					resource.TestCheckResourceAttr(resourceName, "health.0.unhealthy_duration", "300"),
					resource.TestCheckResourceAttr(resourceName, "health.0.auto_healing", "true"),
					resource.TestCheckResourceAttr(resourceName, "health.0.health_check_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health.0.health_check_types.0", "vmState"),
				),
			},
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{
					statefulNodeName:     statefulNodeName,
					health:               testHealthStatefulNodeAzureV3Config_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "health.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "health.0.grace_period", "20"),
					resource.TestCheckResourceAttr(resourceName, "health.0.unhealthy_duration", "315"),
					resource.TestCheckResourceAttr(resourceName, "health.0.auto_healing", "true"),
					resource.TestCheckResourceAttr(resourceName, "health.0.health_check_types", "1"),
					resource.TestCheckResourceAttr(resourceName, "health.0.health_check_types.0", "vmState"),
				),
			},
		},
	})
}

const testHealthStatefulNodeAzureV3Config_Create = `
		health {
            healthCheckTypes = [
                "vmState"
            ]
            unhealthy_duration = "315"
            grace_period = "20"
            auto_healing = true
        }
`

const testHealthStatefulNodeAzureV3Config_Update = `
		health {
            healthCheckTypes = [
                "vmState"
            ]
            unhealthy_duration: = "300"
            grace_period = "10"
            auto_healing = true
        }
`

//endregion

// region Stateful Node Azure : VMSizes
func TestAccSpotinstStatefulNodeAzureV3_VMSizes(t *testing.T) {
	statefulNodeName := "terraform-tests-do-not-delete"
	resourceName := createStatefulNodeAzureV3ResourceName(statefulNodeName)

	var node azurev3.StatefulNode
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testStatefulNodeAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{statefulNodeName: statefulNodeName,
					vmSizes: testVMSizesStatefulNodeAzureV3Config_Update}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckStatefulNodeAzureV3Exists(&node, resourceName),
					testCheckStatefulNodeAzureV3Attributes(&node, statefulNodeName),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.spot_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.spot_sizes.0", "standard_ds2_v2"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.od_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.od_sizes.0", "standard_ds1_v2"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.preferred_spot_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.preferred_spot_sizes.#", "standard_ds1_v2"),
				),
			},
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{
					statefulNodeName:     statefulNodeName,
					vmSizes:              testVMSizesStatefulNodeAzureV3Config_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.spot_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.spot_sizes.0", "standard_ds3_v2"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.od_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.od_sizes.0", "standard_ds4_v2"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.preferred_spot_sizes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vm_sizes.0.preferred_spot_sizes.#", "standard_ds1_v2"),
				),
			},
		},
	})
}

const testVMSizesStatefulNodeAzureV3Config_Create = `
            vmSizes {
                spotSizes = [
                    "standard_ds2_v2''
                ]
                odSizes = [
                    "standard_ds1_v2"
                ]
                preferredSpotSizes =  [
                    "standard_ds1_v2"
                ]
            }
`

const testVMSizesStatefulNodeAzureV3Config_Update = `
            vmSizes {
                spotSizes = [
                    "standard_ds3_v2"
                ]
                odSizes = [
                    "standard_ds4_v2"
                ]
                preferredSpotSizes =  [
                    "standard_ds1_v2"
                ]
            }
`

//endregion

// region Stateful Node Azure : Scheduling
func TestAccSpotinstStatefulNodeAzureV3_Scheduling(t *testing.T) {
	statefulNodeName := "terraform-tests-do-not-delete"
	resourceName := createStatefulNodeAzureV3ResourceName(statefulNodeName)

	var node azurev3.StatefulNode
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testStatefulNodeAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config:             createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{statefulNodeName: statefulNodeName}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckStatefulNodeAzureV3Exists(&node, resourceName),
					testCheckStatefulNodeAzureV3Attributes(&node, statefulNodeName),
					resource.TestCheckResourceAttr(resourceName, "scheduling_tasks.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_tasks.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_tasks.0.cron_expression", "44 10 * * *"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_tasks.0.type", "pause"),
				),
			},
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{
					statefulNodeName:     statefulNodeName,
					scheduling:           testSchedulingStatefulNodeAzureV3Config_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "scheduling_tasks.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_tasks.0.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_tasks.0.cron_expression", "48 10 * * *"),
					resource.TestCheckResourceAttr(resourceName, "scheduling_tasks.0.type", "resume"),
				),
			},
		},
	})
}

const testSchedulingStatefulNodeAzureV3Config_Create = `
  scheduling_tasks {
    is_enabled = true
	cron_expression = "44 10 * * *"
    type = "pause"
  }
`

const testSchedulingStatefulNodeAzureV3Config_Update = `
  scheduling_tasks {
    is_enabled = true
	cron_expression = "48 10 * * *"
    type = "resume"
  }
`

//endregion

// region Stateful Node Azure : Image
func TestAccSpotinstStatefulNodeAzureV3_Image(t *testing.T) {
	statefulNodeName := "terraform-tests-do-not-delete"
	resourceName := createStatefulNodeAzureV3ResourceName(statefulNodeName)

	var node azurev3.StatefulNode
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testStatefulNodeAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{statefulNodeName: statefulNodeName,
					image: testImageStatefulNodeAzureV3Config_Update}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckStatefulNodeAzureV3Exists(&node, resourceName),
					testCheckStatefulNodeAzureV3Attributes(&node, statefulNodeName),
					resource.TestCheckResourceAttr(resourceName, "image.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.offer", "UbuntuServer"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.publisher", "Canonical"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.sku", "18.04-LTS"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace.0.version", "latest"),
					resource.TestCheckResourceAttr(resourceName, "image.0.custom_image.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.custom_image.0.custom_image_resource_group_name", "foo"),
					resource.TestCheckResourceAttr(resourceName, "image.0.custom_image.0.name", "Canonical"),
					resource.TestCheckResourceAttr(resourceName, "image.0.gallery.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.gallery.0.gallery_resource_group_name", "grc"),
					resource.TestCheckResourceAttr(resourceName, "image.0.gallery.0.gallery_name", "NotCanonically"),
					resource.TestCheckResourceAttr(resourceName, "image.0.gallery.0.image_name", "18.06-LTS"),
					resource.TestCheckResourceAttr(resourceName, "image.0.gallery.0.version_name", "first"),
				),
			},
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{
					statefulNodeName:     statefulNodeName,
					image:                testImageStatefulNodeAzureV3Config_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "image.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace_image.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace_image.0.offer", "UbuntuServer"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace_image.0.publisher", "Pub"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace_image.0.sku", "18.04-LTS"),
					resource.TestCheckResourceAttr(resourceName, "image.0.marketplace_image.0.version", "latestVersion"),
					resource.TestCheckResourceAttr(resourceName, "image.0.custom_image.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.custom_image.0.custom_image_resource_group_name", "foo"),
					resource.TestCheckResourceAttr(resourceName, "image.0.custom_image.0.name", "Canonical"),
					resource.TestCheckResourceAttr(resourceName, "image.0.gallery.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "image.0.gallery.0.gallery_resource_group_name", "grc"),
					resource.TestCheckResourceAttr(resourceName, "image.0.gallery.0.gallery_name", "Canonically"),
					resource.TestCheckResourceAttr(resourceName, "image.0.gallery.0.image_name", "18.06-LTS"),
					resource.TestCheckResourceAttr(resourceName, "image.0.gallery.0.version_name", "latest"),
				),
			},
		},
	})
}

const testImageStatefulNodeAzureV3Config_Create = `
                image {
                    market_space_image {
                        publisher = "Canonical",
                        offer" = "UbuntuServer",
                        sku = "18.04-LTS",
                        version = "latest"
                    }
				}
`

const testImageStatefulNodeAzureV3Config_Update = `
                image {
					gallery{
						gallery_resource_group_name = "grc"
						gallery_name = "Canonically"
						image_name = "18.06-LTS"
						version_name = "latest"
					}
				}
`

//endregion

// region Stateful Node Azure : Load Balancers
func TestAccSpotinstStatefulNodeAzureV3_LoadBalancers(t *testing.T) {
	statefulNodeName := "terraform-tests-do-not-delete"
	resourceName := createStatefulNodeAzureV3ResourceName(statefulNodeName)

	var node azurev3.StatefulNode
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testStatefulNodeAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{statefulNodeName: statefulNodeName,
					loadBalancers: testLoadBalancersStatefulNodeAzureV3Config_Update}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckStatefulNodeAzureV3Exists(&node, resourceName),
					testCheckStatefulNodeAzureV3Attributes(&node, statefulNodeName),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.type", "loadBalancer"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.name", "kubernetes"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.resource_group_name", "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.sku", "Standard"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.backend_pool_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.backend_pool_names.0", "kubernetes"),
				),
			},
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{
					statefulNodeName:     statefulNodeName,
					loadBalancers:        testLoadBalancersStatefulNodeAzureV3Config_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "load_balancer.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.type", "loadBalancer"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.name", "ocean_kubernetes"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.resource_group_name", "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.sku", "NotStandard"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.backend_pool_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "load_balancer.backend_pool_names.0", "kubernetes"),
				),
			},
		},
	})
}

const testLoadBalancersStatefulNodeAzureV3Config_Create = `
  load_balancer {
    backend_pool_names = [
      "kubernetes"
    ]
    sku = "Standard"
    name = "kubernetes"
    resource_group_name = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"
    type = "loadBalancer"
  }
`

const testLoadBalancersStatefulNodeAzureV3Config_Update = `
  load_balancer {
    backend_pool_names = [
      "kubernetes"
    ]
    sku = "NotStandard"
    name = "ocean_kubernetes"
    resource_group_name = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"
    type = "loadBalancer"
  }
`

//endregion

// region Stateful Node Azure : Network
func TestAccSpotinstStatefulNodeAzureV3_Network(t *testing.T) {
	statefulNodeName := "terraform-tests-do-not-delete"
	resourceName := createStatefulNodeAzureV3ResourceName(statefulNodeName)

	var node azurev3.StatefulNode
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testStatefulNodeAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{statefulNodeName: statefulNodeName,
					network: testNetworkStatefulNodeAzureV3Config_Update}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckStatefulNodeAzureV3Exists(&node, resourceName),
					testCheckStatefulNodeAzureV3Attributes(&node, statefulNodeName),
					resource.TestCheckResourceAttr(resourceName, "network.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.subnet_name", "default"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.assign_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.is_primary", "true"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.public_ip_sku", "Standard"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.network_security_group.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.network_security_group.0.name", "core-reliability-network-security-group"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.network_security_group.0.network_resource_group_name", "CoreReliabilityResourceGroup"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.enable_ip_forwarding", "true"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.private_ip_addresses", "1234"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.additional_ip_configurations.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.additional_ip_configurations.0.name", "core-reliability-additional-ip-configurations"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.additional_ip_configurations.0.private_ip_address_version", "12345"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.public_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.public_ips.name", "core-reliability-public-ips-name"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.public_ips.network_resource_group_name", "CoreReliabilityResourceGroup"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.application_security_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.application_security_groups.name", "core-reliability-application-security-groups"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.application_security_groups.network_resource_group_name", "CoreReliabilityResourceGroup"),
					resource.TestCheckResourceAttr(resourceName, "network.0.resource_group_name", "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"),
					resource.TestCheckResourceAttr(resourceName, "network.0.virtual_network_name", "aks-vnet-48068046"),
				),
			},
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{
					statefulNodeName:     statefulNodeName,
					network:              testNetworkStatefulNodeAzureV3Config_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "network.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.subnet_name", "default"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.assign_public_ip", "false"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.is_primary", "true"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.public_ip_sku", "Standard"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.network_security_group.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.network_security_group.0.name", "core-reliability-network-security-group"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.network_security_group.0.network_resource_group_name", "CoreReliabilityResourceGroup2"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.0.enable_ip_forwarding", "true"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.private_ip_addresses", "1234"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.additional_ip_configurations.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.additional_ip_configurations.0.name", "core-reliability-additional-ip-configurations"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.additional_ip_configurations.0.private_ip_address_version", "12345"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.public_ips.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.public_ips.name", "core-reliability-public-ips-name"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.public_ips.network_resource_group_name", "CoreReliabilityResourceGroup"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.application_security_groups.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.application_security_groups.name", "core-reliability-application-security-groups"),
					resource.TestCheckResourceAttr(resourceName, "network.0.network_interface.4170446135.application_security_groups.network_resource_group_name", "CoreReliabilityResourceGroup"),
					resource.TestCheckResourceAttr(resourceName, "network.0.resource_group_name", "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"),
					resource.TestCheckResourceAttr(resourceName, "network.0.virtual_network_name", "aks-vnet-48068046"),
				),
			},
		},
	})
}

const testNetworkStatefulNodeAzureV3Config_Create = `
  network {
    virtual_network_name = "aks-vnet-48068046"
    resource_group_name = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"

    network_interface {
      subnet_name = "default"
      assign_public_ip = true
      is_primary = true
	  public_ip_sku = "Standard"
	  network_security_group{
		name = "core-reliability-network-security-group"
		network_resource_group_name = "CoreReliabilityResourceGroup"
	  }
		enable_ip_forwarding = true
		private_ip_addresses = "1234"
		additional_ip_configurations{
			name = "core-reliability-additional-ip-configurations"
			private_ip_address_version = "12345"
	  	}
		public_ips{
			name = "core-reliability-public-ips-name"
			network_resource_group_name = "CoreReliabilityResourceGroup"
		}
		application_security_groups{
			name = "core-reliability-application-security-groups"
			network_resource_group_name = "CoreReliabilityResourceGroup"
		}
    }
  }
`

const testNetworkStatefulNodeAzureV3Config_Update = `
  network {
    virtual_network_name = "aks-vnet-48068046"
    resource_group_name = "MC_terraform-resource-group-DO-NOT-DELETE_terraform-Kubernetes-cluster_eastus"

    network_interface {
      subnet_name = "default"
      assign_public_ip = false
      is_primary = true
	  public_ip_sku = "Standard"
	  network_security_group{
		name = "core-reliability-network-security-group"
		network_resource_group_name = "CoreReliabilityResourceGroup2"
	  }
		enable_ip_forwarding = true
		private_ip_addresses = "1234"
		additional_ip_configurations{
			name = "core-reliability-additional-ip-configurations"
			private_ip_address_version = "12345"
	  	}
		public_ips{
			name = "core-reliability-public-ips-name"
			network_resource_group_name = "CoreReliabilityResourceGroup"
		}
		application_security_groups{
			name = "core-reliability-application-security-groups"
			network_resource_group_name = "CoreReliabilityResourceGroup"
		}
    }
  }
`

//endregion

// region Stateful Node Azure : Extensions
func TestAccSpotinstStatefulNodeAzureV3_Extensions(t *testing.T) {
	statefulNodeName := "terraform-tests-do-not-delete"
	resourceName := createStatefulNodeAzureV3ResourceName(statefulNodeName)

	var node azurev3.StatefulNode
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testStatefulNodeAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{statefulNodeName: statefulNodeName,
					extensions: testExtensionsStatefulNodeAzureV3Config_Update}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckStatefulNodeAzureV3Exists(&node, resourceName),
					testCheckStatefulNodeAzureV3Attributes(&node, statefulNodeName),
					resource.TestCheckResourceAttr(resourceName, "extension.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.api_version", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.minor_version_auto_upgrade", "true"),
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.name", "terraform-extension"),
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.publisher", "Microsoft.Azure.Extensions"),
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.type", "Linux"),
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.protected_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.protected_settings.0.script", "IyEvYmluL2Jhc2gKZWNobyAibmlyIiA+IC9ob21lL25pci9uaXIudHh0Cg=="), //ToDo check about field script under protected settings
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.public_settings.#", "1"),
				),
			},
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{
					statefulNodeName:     statefulNodeName,
					extensions:           testExtensionsStatefulNodeAzureV3Config_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "extension.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.api_version", "1.0"),
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.minor_version_auto_upgrade", "false"),
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.name", "terraform-extension"),
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.publisher", "Microsoft.Azure.Extensions"),
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.type", "Windows"),
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.protected_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.protected_settings.0.script", "IyEvYmluL2Jhc2gKZWNobyAibmlyIiA+IC9ob21lL25pci9uaXIudHh0Cg=="), //ToDo check about field script under protected settings
					resource.TestCheckResourceAttr(resourceName, "extension.1031128857.public_settings.#", "1"),
				),
			},
		},
	})
}

const testExtensionsStatefulNodeAzureV3Config_Create = `
    extension {
      api_version = "1.0"
      minor_version_auto_upgrade = true
      name = "terraform-extension"
      publisher = "Microsoft.Azure.Extensions"
      type = "Linux"
      protected_settings{
		script = "IyEvYmluL2Jhc2gKZWNobyAibmlyIiA+IC9ob21lL25pci9uaXIudHh0Cg=="
      }
		public_settings{}	
	}
`
const testExtensionsStatefulNodeAzureV3Config_Update = `
    extension {
      api_version = "1.0"
      minor_version_auto_upgrade = false
      name = "terraform-extension"
      publisher = "Microsoft.Azure.Extensions"
      type = "Windows"
      protected_settings{
		script = "IyEvYmluL2Jhc2gKZWNobyAibmlyIiA+IC9ob21lL25pci9uaXIudHh0Cg=="
      }
		public_settings{}	
	}
`

//endregion

// region Stateful Node Azure : Secret
func TestAccSpotinstStatefulNodeAzureV3_Secret(t *testing.T) {
	statefulNodeName := "terraform-tests-do-not-delete"
	resourceName := createStatefulNodeAzureV3ResourceName(statefulNodeName)

	var node azurev3.StatefulNode
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t, "azure") },
		Providers:    TestAccProviders,
		CheckDestroy: testStatefulNodeAzureV3Destroy,

		Steps: []resource.TestStep{
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{statefulNodeName: statefulNodeName,
					extensions: testExtensionsStatefulNodeAzureV3Config_Update}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testCheckStatefulNodeAzureV3Exists(&node, resourceName),
					testCheckStatefulNodeAzureV3Attributes(&node, statefulNodeName),
					resource.TestCheckResourceAttr(resourceName, "secrets.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secrets.1031128857.source_vault.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secrets.1031128857.source_vault.name", "core-reliability-source-vault-name"),
					resource.TestCheckResourceAttr(resourceName, "secrets.1031128857.source_vault.resource_group_name", "CoreReliabilityResourceGroup"),
					resource.TestCheckResourceAttr(resourceName, "secrets.1031128857.vault_certificates.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secrets.1031128857.vault_certificates.certificate_url", "core-reliability-certificate-url"),
					resource.TestCheckResourceAttr(resourceName, "secrets.1031128857.vault_certificates.certificate_store", "core-reliability-certificate-store"),
				),
			},
			{
				Config: createStatefulNodeAzureV3Terraform(&AzureV3StatefulNodeConfigMetadata{
					statefulNodeName:     statefulNodeName,
					secret:               testSecretsStatefulNodeAzureV3Config_Update,
					updateBaselineFields: true,
				}),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "secrets.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secrets.1031128857.source_vault.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secrets.1031128857.source_vault.name", "core-reliability-source-vault-name"),
					resource.TestCheckResourceAttr(resourceName, "secrets.1031128857.source_vault.resource_group_name", "CoreReliabilityResourceGroup3"),
					resource.TestCheckResourceAttr(resourceName, "secrets.1031128857.vault_certificates.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secrets.1031128857.vault_certificates.certificate_url", "core-reliability-certificate-url"),
					resource.TestCheckResourceAttr(resourceName, "secrets.1031128857.vault_certificates.certificate_store", "core-reliability-certificate-store-is"),
				),
			},
		},
	})
}

const testSecretsStatefulNodeAzureV3Config_Create = `
    secrets {
      source_vault {
      resource_group_name = "CoreReliabilityResourceGroup"
      name = "core-reliability-source-vault-name"
	}
      vault_certificates {
      certificate_url = "core-reliability-certificate-url"
      certificate_store = "core-reliability-certificate-store"
	}
}
`
const testSecretsStatefulNodeAzureV3Config_Update = `
    secrets {
      source_vault {
      resource_group_name = "CoreReliabilityResourceGroup3"
      name = "core-reliability-source-vault-name"
	}
      vault_certificates {
      certificate_url = "core-reliability-certificate-url"
      certificate_store = "core-reliability-certificate-store-is"
	}
}
`

//endregion

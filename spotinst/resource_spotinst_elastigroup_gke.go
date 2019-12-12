package spotinst

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_gcp_disk"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_gcp_gpu"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_gcp_instance_types"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_gcp_integrations"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_gcp_launch_configuration"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_gcp_network_interface"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_gcp_scaling_policies"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_gcp_strategy"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_gke"
)

func resourceSpotinstElastigroupGKE() *schema.Resource {
	setupElastigroupGKEResource()

	return &schema.Resource{
		Create: resourceSpotinstElastigroupGKECreate,
		Read:   resourceSpotinstElastigroupGKERead,
		Update: resourceSpotinstElastigroupGKEUpdate,
		Delete: resourceSpotinstElastigroupGKEDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.ElastigroupGKEResource.GetSchemaMap(),
	}
}

// setupElastigroupGKEResource calls the setup function for each of the children blocks.
func setupElastigroupGKEResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	elastigroup_gke.Setup(fieldsMap)
	elastigroup_gcp_disk.Setup(fieldsMap)
	elastigroup_gcp_gpu.Setup(fieldsMap)
	elastigroup_gcp_instance_types.Setup(fieldsMap)
	elastigroup_gcp_integrations.Setup(fieldsMap)
	elastigroup_gcp_launch_configuration.Setup(fieldsMap)
	elastigroup_gcp_network_interface.Setup(fieldsMap)
	elastigroup_gcp_scaling_policies.Setup(fieldsMap)
	elastigroup_gcp_strategy.Setup(fieldsMap)

	commons.ElastigroupGKEResource = commons.NewElastigroupGKEResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//     Import GKE Group
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func importGKEGroup(resourceData *schema.ResourceData, meta interface{}) (*gcp.Group, error) {

	// first build a GCP group from the user's template
	templateGroup, err := commons.ElastigroupGKEResource.OnCreate(resourceData, meta)
	if err != nil {
		return nil, err
	}

	// creates an ImportGroup out of the GCP group
	importedGroup, err := commons.ElastigroupGKEResource.OnImport(templateGroup, resourceData, meta)
	if err != nil {
		return nil, err
	}

	if v, ok := resourceData.Get(string(elastigroup_gke.NodeImage)).(string); ok && v != "" {
		nodeImage := spotinst.String(v)
		importedGroup.SetNodeImage(nodeImage)
	}

	// handle cluster_id scoping due to deprecated cluster_id location
	// this can be updated once the deprecated field is removed
	var clusterID *string
	if v, ok := resourceData.Get(string(elastigroup_gke.ClusterID)).(string); ok && v != "" {
		clusterID = spotinst.String(v)
	} else {
		if templateGroup != nil && templateGroup.Integration != nil && templateGroup.Integration.GKE != nil &&
			templateGroup.Integration.GKE.ClusterID != nil {
			clusterID = templateGroup.Integration.GKE.ClusterID
		} else {
			return nil, fmt.Errorf("cluster_id is required")
		}
	}

	input := &gcp.ImportGKEClusterInput{
		ClusterID:       clusterID,
		ClusterZoneName: spotinst.String(resourceData.Get(string(elastigroup_gke.ClusterZoneName)).(string)),
		DryRun:          spotinst.Bool(true),
		Group:           importedGroup,
	}

	// make te request with the custom group, get back a GCP group with some generated fields
	resp, err := meta.(*Client).elastigroup.CloudProviderGCP().ImportGKECluster(context.Background(), input)

	if err != nil {
		// If the group was not found, return nil so that we can show
		// that the group is gone.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGroupNotFound {
					resourceData.SetId("")
					return nil, err
				}
			}
		}
		// Some other error, report it.
		return nil, fmt.Errorf("GKE:IMPORT failed to read group: %s", err)
	}

	return resp.Group, err
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

// resourceSpotinstElastigroupGKECreate begins the creation request and
// creates an object representing the newly created group or returns an error.
func resourceSpotinstElastigroupGKECreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.ElastigroupGKEResource.GetName())

	// do the import call and get the generated fields
	gkeGroup, err := importGKEGroup(resourceData, meta.(*Client))
	if err != nil {
		return err
	}

	if gkeGroup == nil {
		return fmt.Errorf("[ERROR] Failed to import group. Does the GKE cluster exist?")
	}

	// merge the imported group with the user's template group, giving preference to the user's template
	tempGroup, err := commons.ElastigroupGKEResource.OnMerge(gkeGroup, resourceData, meta)
	if err != nil {
		return err
	}

	// if the user set cluster_id using deprecated method, set the value in the correct spot
	// this can be removed once the deprecated field is removed
	if v, ok := resourceData.Get(string(elastigroup_gke.ClusterID)).(string); ok && v != "" {
		tempGroup.Integration.GKE.ClusterID = spotinst.String(v)
	}

	// call create with the reconciled group
	groupId, err := createGKEGroup(tempGroup, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(groupId))
	log.Printf("===> Elastigroup for GKE created successfully: %s <===", resourceData.Id())
	return resourceSpotinstElastigroupGKERead(resourceData, meta)
}

func createGKEGroup(gkeGroup *gcp.Group, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(gkeGroup); err != nil {
		return nil, err
	} else {
		log.Printf("===> Group create configuration: %s", json)
	}

	input := &gcp.CreateGroupInput{Group: gkeGroup}

	var resp *gcp.CreateGroupOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.elastigroup.CloudProviderGCP().Create(context.Background(), input)
		if err != nil {

			// If there's some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to create GKE group: %s", err)
	}
	return resp.Group.ID, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

// resourceSpotinstElastigroupGKERead creates an object representing an existing elastigroup
// by making a get request using the Spotinst API or returns an error.
func resourceSpotinstElastigroupGKERead(resourceData *schema.ResourceData, meta interface{}) error {
	groupId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.ElastigroupGKEResource.GetName(), groupId)

	input := &gcp.ReadGroupInput{GroupID: spotinst.String(groupId)}
	resp, err := meta.(*Client).elastigroup.CloudProviderGCP().Read(context.Background(), input)
	if err != nil {
		// if the group was not found, return nil to show the group doesn't exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGroupNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// report any other error
		return fmt.Errorf("failed to read group: %s", err)
	}

	// If nothing was found, then return no state.
	groupResponse := resp.Group
	if groupResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.ElastigroupGKEResource.OnRead(groupResponse, resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Elastigroup read successfully: %s <===", groupId)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

// resourceSpotinstElastigroupGKEUpdate updates an existing elastigroup
// and creates an object representing the updated group or returns an error.
func resourceSpotinstElastigroupGKEUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	groupId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.ElastigroupGKEResource.GetName(), groupId)

	shouldUpdate, elastigroup, err := commons.ElastigroupGKEResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		elastigroup.SetID(spotinst.String(groupId))

		if err := updateGKEGroup(elastigroup, resourceData, meta); err != nil {
			return err
		}
	}

	log.Printf("===> Elastigroup updated successfully: %s <===", groupId)
	return resourceSpotinstElastigroupGKERead(resourceData, meta)
}

// updateGKEGroup sends the update request to the Spotinst API and returns an error if the request fails.
func updateGKEGroup(elastigroup *gcp.Group, resourceData *schema.ResourceData, meta interface{}) error {
	// we need to remove the location and clusterID params used when calling Create.
	// The core does not support these when calling Update.
	elastigroup.Integration.SetGKE(&gcp.GKEIntegration{
		AutoUpdate: elastigroup.Integration.GKE.AutoUpdate,
		AutoScale:  elastigroup.Integration.GKE.AutoScale,
	})

	var input = &gcp.UpdateGroupInput{Group: elastigroup}
	groupId := resourceData.Id()

	if json, err := commons.ToJson(elastigroup); err != nil {
		return err
	} else {
		log.Printf("===> Group update configuration: %s", json)
	}

	if _, err := meta.(*Client).elastigroup.CloudProviderGCP().Update(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update group [%v]: %v", groupId, err)
	}

	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

// resourceSpotinstElastigroupGKEDelete deletes a specific elastigroup or returns an error.
func resourceSpotinstElastigroupGKEDelete(resourceData *schema.ResourceData, meta interface{}) error {
	groupId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.ElastigroupGKEResource.GetName(), groupId)

	if err := deleteGKEGroup(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Elastigroup deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

// deleteGKEGroup sends the delete request to the Spotinst API or an error if the request fails.
func deleteGKEGroup(resourceData *schema.ResourceData, meta interface{}) error {
	groupId := resourceData.Id()
	input := &gcp.DeleteGroupInput{GroupID: spotinst.String(groupId)}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Group delete configuration: %s", json)
	}

	if _, err := meta.(*Client).elastigroup.CloudProviderGCP().Delete(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete group: %s", err)
	}
	return nil
}

package spotinst

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_gke_launch_spec"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_gke_launch_spec_scheduling"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_gke_launch_spec_strategy"
)

func resourceSpotinstOceanGKELaunchSpec() *schema.Resource {
	setupOceanGKELaunchSpecResource()

	return &schema.Resource{
		Create: resourceSpotinstOceanGKELaunchSpecCreate,
		Read:   resourceSpotinstOceanGKELaunchSpecRead,
		Update: resourceSpotinstOceanGKELaunchSpecUpdate,
		Delete: resourceSpotinstOceanGKELaunchSpecDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.OceanGKELaunchSpecResource.GetSchemaMap(),
	}
}

func setupOceanGKELaunchSpecResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	ocean_gke_launch_spec.Setup(fieldsMap)
	ocean_gke_launch_spec_strategy.Setup(fieldsMap)
	ocean_gke_launch_spec_scheduling.Setup(fieldsMap)

	commons.OceanGKELaunchSpecResource = commons.NewOceanGKELaunchSpecResource(fieldsMap)
}

func resourceSpotinstOceanGKELaunchSpecCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate), commons.OceanGKELaunchSpecResource.GetName())

	var importedLaunchSpec *gcp.LaunchSpec
	var err error

	if v, ok := resourceData.Get(string(ocean_gke_launch_spec.NodePoolName)).(string); ok && v != "" {
		importedLaunchSpec, err = importGKELaunchSpec(resourceData, meta)

		if err != nil {
			return err
		}

	}

	launchSpec, err := commons.OceanGKELaunchSpecResource.OnCreate(importedLaunchSpec, resourceData, meta.(*Client))

	if err != nil {
		return err
	}

	launchSpecId, err := createGKELaunchSpec(launchSpec, meta.(*Client))

	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(launchSpecId))

	return resourceSpotinstOceanGKELaunchSpecRead(resourceData, meta)
}

func createGKELaunchSpec(launchSpec *gcp.LaunchSpec, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(launchSpec); err != nil {
		return nil, err
	} else {
		log.Printf("===> LaunchSpec GKE create configuration: %s", json)
	}

	input := &gcp.CreateLaunchSpecInput{LaunchSpec: launchSpec}

	if out, err := spotinstClient.ocean.CloudProviderGCP().CreateLaunchSpec(context.Background(), input); err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create launchSpec: %s", err)
	} else {
		return out.LaunchSpec.ID, nil
	}
}

const ErrCodeGKELaunchSpecNotFound = "CANT_GET_OCEAN_LAUNCH_SPEC"

func resourceSpotinstOceanGKELaunchSpecRead(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.OceanGKELaunchSpecResource.GetName(), id)

	input := &gcp.ReadLaunchSpecInput{LaunchSpecID: spotinst.String(id)}
	resp, err := meta.(*Client).ocean.CloudProviderGCP().ReadLaunchSpec(context.Background(), input)

	if err != nil {
		// If the launchSpec was not found, return nil so that we can show
		// that it does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGKELaunchSpecNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return fmt.Errorf("failed to read GKE launchSpec: %s", err)
	}

	// if nothing was found, return no state
	launchSpecResponse := resp.LaunchSpec
	if launchSpecResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanGKELaunchSpecResource.OnRead(launchSpecResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> launchSpec GKE read successfully: %s <===", id)
	return nil
}

func resourceSpotinstOceanGKELaunchSpecUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.OceanGKELaunchSpecResource.GetName(), id)

	shouldUpdate, launchSpec, err := commons.OceanGKELaunchSpecResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		launchSpec.SetId(spotinst.String(id))
		if err := updateGKELaunchSpec(launchSpec, resourceData, meta); err != nil {
			return err
		}
	}
	log.Printf("===> launchSpec GKE updated successfully: %s <===", id)
	return resourceSpotinstOceanGKELaunchSpecRead(resourceData, meta)
}

func updateGKELaunchSpec(launchSpec *gcp.LaunchSpec, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &gcp.UpdateLaunchSpecInput{
		LaunchSpec: launchSpec,
	}

	launchSpecId := resourceData.Id()

	oceanId := resourceData.Get(string(ocean_gke_launch_spec.OceanId))
	var shouldRoll = false
	if updatePolicy, exists := resourceData.GetOkExists(string(ocean_gke_launch_spec.UpdatePolicy)); exists {
		list := updatePolicy.([]interface{})
		if len(list) > 0 && list[0] != nil {
			m := list[0].(map[string]interface{})

			if roll, ok := m[string(ocean_gke_launch_spec.ShouldRoll)].(bool); ok && roll {
				shouldRoll = roll
			}
		}
	}

	if json, err := commons.ToJson(launchSpec); err != nil {
		return err
	} else {
		log.Printf("===> launchSpec GKE update configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderGCP().UpdateLaunchSpec(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update launchSpec GKE [%v]: %v", launchSpecId, err)
	} else if shouldRoll {
		if err := rollOceanGKELaunchSpec(resourceData, meta); err != nil {
			log.Printf("[ERROR] Cluster [%v] roll failed, error: %v", oceanId, err)
			return err
		}
	} else {
		log.Printf("onRoll() -> Field [%v] is false, skipping cluster roll", string(ocean_gke_launch_spec.ShouldRoll))
	}

	return nil
}

func rollOceanGKELaunchSpec(resourceData *schema.ResourceData, meta interface{}) error {
	specID := resourceData.Id()
	clusterID := resourceData.Get(string(ocean_gke_launch_spec.OceanId)).(string)

	updatePolicy, exists := resourceData.GetOkExists(string(ocean_gke_launch_spec.UpdatePolicy))
	if !exists {
		return fmt.Errorf("ocean/gke: missing update policy for cluster %q", clusterID)
	}

	list := updatePolicy.([]interface{})
	if len(list) > 0 && list[0] != nil {
		updateClusterSchema := list[0].(map[string]interface{})

		rollConfig, ok := updateClusterSchema[string(ocean_gke_launch_spec.RollConfig)]
		if !ok || rollConfig == nil {
			return fmt.Errorf("ocean/gke: missing roll configuration, "+
				"skipping roll for cluster %q", clusterID)
		}

		rollSpec, err := expandOceanGKELaunchSpecRollConfig(rollConfig, clusterID, specID)
		if err != nil {
			return fmt.Errorf("ocean/gke: failed expanding roll "+
				"configuration for cluster %q, error: %v", clusterID, err)
		}

		rollJSON, err := commons.ToJson(rollConfig)
		if err != nil {
			return fmt.Errorf("ocean/gke: failed marshaling roll "+
				"configuration for cluster %q, error: %v", clusterID, err)
		}

		log.Printf("onRoll() -> Rolling cluster [%v] with configuration %s", clusterID, rollJSON)
		rollInput := &gcp.CreateRollInput{Roll: rollSpec}
		if _, err = meta.(*Client).ocean.CloudProviderGCP().CreateRoll(context.TODO(), rollInput); err != nil {
			return fmt.Errorf("onRoll() -> Roll failed for cluster [%v], error: %v", clusterID, err)
		}
		log.Printf("onRoll() -> Successfully rolled cluster [%v]", clusterID)
	}

	return nil
}

func resourceSpotinstOceanGKELaunchSpecDelete(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanGKELaunchSpecResource.GetName(), id)

	if err := deleteGKELaunchSpec(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> launchSpec GKE deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteGKELaunchSpec(resourceData *schema.ResourceData, meta interface{}) error {
	launchSpecId := resourceData.Id()
	input := &gcp.DeleteLaunchSpecInput{
		LaunchSpecID: spotinst.String(launchSpecId),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> launchSpec GKE delete configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.CloudProviderGCP().DeleteLaunchSpec(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete launchSpecId: %s", err)
	}
	return nil
}

//region Import Ocean GKE Launch Spec
func importGKELaunchSpec(resourceData *schema.ResourceData, meta interface{}) (*gcp.LaunchSpec, error) {
	input := &gcp.ImportOceanGKELaunchSpecInput{
		OceanId:      spotinst.String(resourceData.Get("ocean_id").(string)),
		NodePoolName: spotinst.String(resourceData.Get("node_pool_name").(string)),
	}

	resp, err := meta.(*Client).ocean.CloudProviderGCP().ImportOceanGKELaunchSpec(context.Background(), input)

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
		return nil, fmt.Errorf("ocean GKE: import failed to read group: %s", err)
	}

	return resp.LaunchSpec, err
}

func expandOceanGKELaunchSpecRollConfig(data interface{}, clusterID, specID string) (*gcp.RollSpec, error) {
	list := data.([]interface{})
	spec := &gcp.RollSpec{
		ClusterID:     spotinst.String(clusterID),
		LaunchSpecIDs: []string{specID},
	}

	if list != nil && list[0] != nil {
		m := list[0].(map[string]interface{})

		if v, ok := m[string(ocean_gke_launch_spec.BatchSizePercentage)].(int); ok {
			spec.BatchSizePercentage = spotinst.Int(v)
		}
	}

	return spec, nil
}

//endregion

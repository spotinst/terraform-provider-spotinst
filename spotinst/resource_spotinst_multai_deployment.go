package spotinst

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/multai"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/multai_deployment"
)

func resourceSpotinstMultaiDeployment() *schema.Resource {
	setupMultaiDeploymentResource()

	return &schema.Resource{
		Create: resourceSpotinstMultaiDeploymentCreate,
		Read:   resourceSpotinstMultaiDeploymentRead,
		Update: resourceSpotinstMultaiDeploymentUpdate,
		Delete: resourceSpotinstMultaiDeploymentDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.MultaiDeploymentResource.GetSchemaMap(),
	}
}

func setupMultaiDeploymentResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	multai_deployment.Setup(fieldsMap)

	commons.MultaiDeploymentResource = commons.NewMultaiDeploymentResource(fieldsMap)
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiDeploymentCreate(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnCreate),
		commons.MultaiDeploymentResource.GetName())

	deployment, err := commons.MultaiDeploymentResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	deploymentId, err := createDeployment(deployment, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(deploymentId))
	log.Printf("===> Deployment created successfully: %s <===", resourceData.Id())

	return resourceSpotinstMultaiDeploymentRead(resourceData, meta)
}

func createDeployment(deployment *multai.Deployment, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(deployment); err != nil {
		return nil, err
	} else {
		log.Printf("===> Deployment create configuration: %s", json)
	}

	input := &multai.CreateDeploymentInput{Deployment: deployment}

	var resp *multai.CreateDeploymentOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.multai.CreateDeployment(context.Background(), input)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create deployment: %s", err)
	}

	return resp.Deployment.ID, nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiDeploymentRead(resourceData *schema.ResourceData, meta interface{}) error {
	deploymentId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.MultaiDeploymentResource.GetName(), deploymentId)

	input := &multai.ReadDeploymentInput{DeploymentID: spotinst.String(deploymentId)}
	resp, err := meta.(*Client).multai.ReadDeployment(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to read deployment: %s", err)
	}

	// If nothing was found, return no state
	deployResponse := resp.Deployment
	if deployResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.MultaiDeploymentResource.OnRead(deployResponse, resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Deployment read successfully: %s <===", deploymentId)
	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiDeploymentUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	deploymentId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.MultaiDeploymentResource.GetName(), deploymentId)

	shouldUpdate, deployment, err := commons.MultaiDeploymentResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		deployment.SetId(spotinst.String(deploymentId))
		if err := updateDeployment(deployment, resourceData, meta); err != nil {
			return err
		}
	}

	log.Printf("===> Deployment updated successfully: %s <===", deploymentId)
	return resourceSpotinstMultaiDeploymentRead(resourceData, meta)
}

func updateDeployment(deployment *multai.Deployment, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &multai.UpdateDeploymentInput{Deployment: deployment}
	deploymentId := resourceData.Id()

	if json, err := commons.ToJson(deployment); err != nil {
		return err
	} else {
		log.Printf("===> Deployment update configuration: %s", json)
	}

	if _, err := meta.(*Client).multai.UpdateDeployment(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update deployment [%v]: %v", deploymentId, err)
	}

	return nil
}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

func resourceSpotinstMultaiDeploymentDelete(resourceData *schema.ResourceData, meta interface{}) error {
	deploymentId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.MultaiDeploymentResource.GetName(), deploymentId)

	if err := deleteDeployment(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> Deployment deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteDeployment(resourceData *schema.ResourceData, meta interface{}) error {
	deploymentId := resourceData.Id()
	input := &multai.DeleteDeploymentInput{DeploymentID: spotinst.String(deploymentId)}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Deployment delete configuration: %s", json)
	}

	if _, err := meta.(*Client).multai.DeleteDeployment(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete deployment: %s", err)
	}
	return nil
}

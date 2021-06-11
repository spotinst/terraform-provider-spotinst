package spotinst

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/healthcheck"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/health_check"
)

func resourceSpotinstHealthCheck() *schema.Resource {
	setupHealthCheckResource()

	return &schema.Resource{
		Create: resourceSpotinstHealthCheckCreate,
		Update: resourceSpotinstHealthCheckUpdate,
		Read:   resourceSpotinstHealthCheckRead,
		Delete: resourceSpotinstHealthCheckDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.HealthCheckResource.GetSchemaMap(),
	}
}

func setupHealthCheckResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)
	health_check.Setup(fieldsMap)

	commons.HealthCheckResource = commons.NewHealthCheckResource(fieldsMap)
}

const ErrCodeHealthCheckNotFound = "HEALTH_CHECK_DOESNT_EXIST"

func resourceSpotinstHealthCheckRead(resourceData *schema.ResourceData, meta interface{}) error {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.HealthCheckResource.GetName(), resourceId)

	input := &healthcheck.ReadHealthCheckInput{HealthCheckID: spotinst.String(resourceId)}
	resp, err := meta.(*Client).healthCheck.Read(context.Background(), input)
	if err != nil {
		// If the HealthCheck was not found, return nil so that we can show
		// that the HealthCheck does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeHealthCheckNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}
		return fmt.Errorf("failed to read health check: %s", err)
	}

	// If nothing was found, then return no state.
	HealthCheckResponse := resp.HealthCheck
	if HealthCheckResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.HealthCheckResource.OnRead(HealthCheckResponse, resourceData, meta); err != nil {
		return err
	}
	log.Printf("===> HealthCheck read successfully: %s <===", resourceId)
	return nil
}

func resourceSpotinstHealthCheckCreate(resourceData *schema.ResourceData, meta interface{}) error {

	log.Printf(string(commons.ResourceOnCreate), commons.HealthCheckResource.GetName())

	healthCheck, err := commons.HealthCheckResource.OnCreate(resourceData, meta)
	if err != nil {
		return err
	}

	healthCheckId, err := createHealthCheck(resourceData, healthCheck, meta.(*Client))
	if err != nil {
		return err
	}

	resourceData.SetId(spotinst.StringValue(healthCheckId))

	log.Printf("===> HealthCheck created successfully: %s <===", resourceData.Id())

	return resourceSpotinstHealthCheckRead(resourceData, meta)

}

func createHealthCheck(resourceData *schema.ResourceData, healthCheck *healthcheck.HealthCheck, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(healthCheck); err != nil {
		return nil, err
	} else {
		log.Printf("===> HealthCheck create configuration: %s", json)
	}
	input := &healthcheck.CreateHealthCheckInput{HealthCheck: healthCheck}

	var resp *healthcheck.CreateHealthCheckOutput = nil
	err := resource.Retry(time.Minute, func() *resource.RetryError {
		r, err := spotinstClient.healthCheck.Create(context.Background(), input)
		if err != nil {
			// Checks whether we should retry the HealthCheck creation.
			if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
				for _, err := range errs {
					if err.Code == "InvalidParameterValue" &&
						strings.Contains(err.Message, "Invalid IAM Instance Profile") {
						return resource.RetryableError(err)
					}
				}
			}
			// Some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create HealthCheck: %s", err)
	}
	return resp.HealthCheck.ID, nil

}

func resourceSpotinstHealthCheckUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate), commons.HealthCheckResource.GetName(), resourceId)

	shouldUpdate, healthCheck, err := commons.HealthCheckResource.OnUpdate(resourceData, meta)
	if err != nil {
		return err
	}

	if shouldUpdate {
		healthCheck.SetId(spotinst.String(resourceId))
		if err := updateHealthCheck(healthCheck, resourceData, meta); err != nil {
			return err
		}
	}
	log.Printf("===> HealthCheck updated successfully: %s <===", resourceId)
	return resourceSpotinstHealthCheckRead(resourceData, meta)
}

func updateHealthCheck(healthCheck *healthcheck.HealthCheck, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &healthcheck.UpdateHealthCheckInput{
		HealthCheck: healthCheck,
	}

	healthCheckId := resourceData.Id()

	if json, err := commons.ToJson(healthCheck); err != nil {
		return err
	} else {
		log.Printf("===> HealthCheck update configuration: %s", json)
	}

	if _, err := meta.(*Client).healthCheck.Update(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update HealthCheck [%v]: %v", healthCheckId, err)
	}
	return nil
}

func resourceSpotinstHealthCheckDelete(resourceData *schema.ResourceData, meta interface{}) error {
	resourceId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete), commons.HealthCheckResource.GetName(), resourceId)

	if err := deleteHealthCheck(resourceData, meta); err != nil {
		return err
	}

	log.Printf("===> HealthCheck deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

func deleteHealthCheck(resourceData *schema.ResourceData, meta interface{}) error {
	healthCheckId := resourceData.Id()
	input := &healthcheck.DeleteHealthCheckInput{
		HealthCheckID: spotinst.String(healthCheckId),
	}
	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> HealthCheck delete configuration: %s", json)
	}

	if _, err := meta.(*Client).healthCheck.Delete(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete HealthCheck: %s", err)
	}
	return nil
}

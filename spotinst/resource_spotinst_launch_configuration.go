package spotinst

import (
	"fmt"
	"math/rand"
	"log"
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/aws_elastigroup"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/launch_configuration"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
)

func resourceSpotinstLaunchConfiguration() *schema.Resource {
	launch_configuration.SetupSpotinstLaunchConfigurationResource()

	return &schema.Resource{
		Create: onSpotinstLaunchConfigurationCreate,
		Read:   onSpotinstLaunchConfigurationRead,
		Update: onSpotinstLaunchConfigurationUpdate,
		Delete: onSpotinstLaunchConfigurationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.LaunchConfigurationRepo.GetSchemaMap(),
	}
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstLaunchConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Deleting group: %s", d.Id())
	input := &aws.DeleteGroupInput{GroupID: spotinst.String(d.Id())}
	if _, err := meta.(*Client).elastigroup.CloudProviderAWS().Delete(context.Background(), input); err != nil {
		return fmt.Errorf("failed to delete group: %s", err)
	}
	d.SetId("")
	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstLaunchConfigurationRead(resourceData *schema.ResourceData, meta interface{}) error {
	// Do nothing, read is being triggered by the Spotinst AWS elastigroup resource
	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstLaunchConfigurationCreate(resourceData *schema.ResourceData, meta interface{}) error {
	if value := commons.ElastigroupRepo.GetField(aws_elastigroup.LaunchConfiguration); value == nil {
		id := fmt.Sprintf("%s/%d", string(aws_elastigroup.LaunchConfiguration), rand.Intn(100))
		resourceData.SetId(id)

		commons.LaunchConfigurationRepo.OnCreate(resourceData,  meta)
	}
	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstLaunchConfigurationUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	if hasChanged, err := commons.LaunchConfigurationRepo.OnUpdate(resourceData, meta); hasChanged && err != nil {
		commons.ElastigroupRepo.GetField(aws_elastigroup.LaunchConfiguration).Touch()
	} else if err != nil {
		return err
	}
	return nil
}
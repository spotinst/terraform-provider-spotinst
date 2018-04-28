package spotinst

import (
	"log"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_launch_configuration"
)

func resourceSpotinstElastigroupLaunchConfiguration() *schema.Resource {
	elastigroup_launch_configuration.SetupSpotinstLaunchConfigurationResource()

	return &schema.Resource{
		Create: onSpotinstElastigroupLaunchConfigurationCreate,
		Read:   onSpotinstElastigroupLaunchConfigurationRead,
		Update: onSpotinstElastigroupLaunchConfigurationUpdate,
		Delete: onSpotinstElastigroupLaunchConfigurationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.ElastigroupLaunchConfigurationResource.GetSchemaMap(),
	}
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstElastigroupLaunchConfigurationDelete(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnDelete),
		commons.ElastigroupLaunchConfigurationResource.GetName(), resourceData.Id())
	resourceData.SetId("")
	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstElastigroupLaunchConfigurationRead(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnRead),
		commons.ElastigroupLaunchConfigurationResource.GetName(), resourceData.Id())

	commons.ElastigroupLaunchConfigurationResource.SetTerraformData(
		&commons.TerraformData{
			ResourceData: resourceData,
			Meta:         meta,
		})

	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstElastigroupLaunchConfigurationCreate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnCreate),
		commons.ElastigroupLaunchConfigurationResource.GetName())

	if id == "" {
		id := commons.GenerateCachedResourceId(string(elastigroup_aws.LaunchConfiguration))
		resourceData.SetId(id)
	}
	commons.ElastigroupLaunchConfigurationResource.OnCreate(resourceData, meta)
	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstElastigroupLaunchConfigurationUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.ElastigroupLaunchConfigurationResource.GetName(), id)

	if hasChanged, err := commons.ElastigroupLaunchConfigurationResource.OnUpdate(resourceData, meta); hasChanged && err != nil {
		commons.ElastigroupResource.GetField(elastigroup_aws.LaunchConfiguration).Touch()
	} else if err != nil {
		return err
	}
	return nil
}
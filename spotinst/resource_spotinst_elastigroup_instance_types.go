package spotinst

import (
	"log"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_instance_types"
)

func resourceSpotinstInstanceTypes() *schema.Resource {
	elastigroup_instance_types.SetupSpotinstInstanceTypesResource()

	return &schema.Resource{
		Create: onSpotinstInstanceTypesCreate,
		Read:   onSpotinstInstanceTypesRead,
		Update: onSpotinstInstanceTypesUpdate,
		Delete: onSpotinstInstanceTypesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.ElastigroupInstanceTypesResource.GetSchemaMap(),
	}
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstInstanceTypesDelete(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnDelete),
		commons.ElastigroupInstanceTypesResource.GetName(), resourceData.Id())
	resourceData.SetId("")
	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstInstanceTypesRead(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnRead),
		commons.ElastigroupInstanceTypesResource.GetName(), resourceData.Id())

	commons.ElastigroupInstanceTypesResource.SetTerraformData(
		&commons.TerraformData{
			ResourceData: resourceData,
			Meta:         meta,
		})

	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstInstanceTypesCreate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnCreate),
		commons.ElastigroupInstanceTypesResource.GetName())

	if id == "" {
		id := commons.GenerateCachedResourceId(string(elastigroup_aws.InstanceTypes))
		resourceData.SetId(id)
	}
	commons.ElastigroupInstanceTypesResource.OnCreate(resourceData, meta)
	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstInstanceTypesUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.ElastigroupInstanceTypesResource.GetName(), id)

	if hasChanged, err := commons.ElastigroupInstanceTypesResource.OnUpdate(resourceData, meta); hasChanged && err != nil {
		commons.ElastigroupResource.GetField(elastigroup_aws.InstanceTypes).Touch()
	} else if err != nil {
		return err
	}
	return nil
}
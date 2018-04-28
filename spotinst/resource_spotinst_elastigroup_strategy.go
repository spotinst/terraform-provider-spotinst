package spotinst

import (
	"log"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_aws"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/elastigroup_strategy"
)

func resourceSpotinstElastigroupStrategy() *schema.Resource {
	elastigroup_strategy.SetupSpotinstElastigroupStrategyResource()

	return &schema.Resource{
		Create: onSpotinstElastigroupStrategyCreate,
		Read:   onSpotinstElastigroupStrategyRead,
		Update: onSpotinstElastigroupStrategyUpdate,
		Delete: onSpotinstElastigroupStrategyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: commons.ElastigroupStrategyResource.GetSchemaMap(),
	}
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Delete
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstElastigroupStrategyDelete(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnDelete),
		commons.ElastigroupStrategyResource.GetName(), resourceData.Id())
	resourceData.SetId("")
	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Read
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstElastigroupStrategyRead(resourceData *schema.ResourceData, meta interface{}) error {
	log.Printf(string(commons.ResourceOnRead),
		commons.ElastigroupStrategyResource.GetName(), resourceData.Id())

	commons.ElastigroupStrategyResource.SetTerraformData(
		&commons.TerraformData{
			ResourceData: resourceData,
			Meta:         meta,
		})

	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Create
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstElastigroupStrategyCreate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnCreate),
		commons.ElastigroupStrategyResource.GetName())

	if id == "" {
		id := commons.GenerateCachedResourceId(string(elastigroup_aws.Strategy))
		resourceData.SetId(id)
	}
	commons.ElastigroupStrategyResource.OnCreate(resourceData, meta)
	return nil
}


//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Update
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func onSpotinstElastigroupStrategyUpdate(resourceData *schema.ResourceData, meta interface{}) error {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.ElastigroupStrategyResource.GetName(), id)

	if hasChanged, err := commons.ElastigroupStrategyResource.OnUpdate(resourceData, meta); hasChanged && err != nil {
		commons.ElastigroupResource.GetField(elastigroup_aws.Strategy).Touch()
	} else if err != nil {
		return err
	}
	return nil
}
package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/spark"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/ocean_spark_virtual_node_group"
)

func resourceSpotinstOceanSparkVirtualNodeGroup() *schema.Resource {
	setupOceanSparkVirtualNodeGroupResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstSparkVirtualNodeGroupCreate,
		DeleteContext: resourceSpotinstSparkClusterVirtualNodeGroupDelete,
		ReadContext:   resourceSpotinstSparkClusterVirtualNodeGroupRead,
		UpdateContext: nil,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: commons.OceanSparkResource.GetSchemaMap(),
	}
}

func resourceSpotinstSparkClusterVirtualNodeGroupRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.OceanSparkVirtualNodeGroupResource.GetName(), id)

	input := &spark.ListVngsInput{ClusterID: spotinst.String(id)}
	resp, err := meta.(*Client).ocean.Spark().ListVirtualNodeGroups(ctx, input)
	if err != nil {
		// If the VNG was not found, return nil so that we can show
		// that the VNG does not exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeResourceDoesNotExist {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return diag.Errorf("failed to read VNG: %s", err)
	}

	// if nothing was found, return no state
	vngsResponse := resp.VirtualNodeGroups
	if vngsResponse == nil {
		resourceData.SetId("")
		return nil
	}

	vng := findVngByID(vngsResponse, id)

	if vng == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.OceanSparkVirtualNodeGroupResource.OnRead(vng, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Cluster read successfully: %s <===", id)
	return nil
}

func findVngByID(vngsResponse []*spark.DedicatedVirtualNodeGroup, id string) *spark.DedicatedVirtualNodeGroup {
	for i := range vngsResponse {
		if spotinst.StringValue(vngsResponse[i].VngID) == id {
			return vngsResponse[i]
		}
	}
	return nil
}

func resourceSpotinstSparkClusterVirtualNodeGroupDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.OceanSparkVirtualNodeGroupResource.GetName(), id)

	if err := detachVng(ctx, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> VNG detached successfully: %s <===", resourceData.Id())
	resourceData.SetId("")

	return nil
}

func resourceSpotinstSparkVirtualNodeGroupCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.OceanSparkVirtualNodeGroupResource.GetName())

	vng, err := commons.OceanSparkVirtualNodeGroupResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	clusterID, err := attachVng(vng, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(clusterID))

	log.Printf("===> VNG attached successfully: %s <===", resourceData.Id())
	return resourceSpotinstSparkClusterRead(ctx, resourceData, meta)
}

func attachVng(vng *spark.DedicatedVirtualNodeGroup, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(vng); err != nil {
		return nil, err
	} else {
		log.Printf("===> Attach VNG configuration: %s", json)
	}

	attachRequest := &spark.AttachVirtualNodeGroupRequest{
		VngID: vng.VngID,
	}

	var resp *spark.AttachVngOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &spark.AttachVngInput{VirtualNodeGroup: attachRequest}
		r, err := spotinstClient.ocean.Spark().AttachVirtualNodeGroup(context.Background(), input)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to attach VNG: %s", err)
	}

	return resp.VirtualNodeGroup.VngID, nil
}

func detachVng(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) error {
	input := &spark.DetachVngInput{
		ClusterID: spotinst.String(resourceData.Get("ocean_spark_cluster_id").(string)),
		VngID:     spotinst.String(resourceData.Get("virtual_node_group_id").(string)),
	}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Detach VNG configuration: %s", json)
	}

	if _, err := meta.(*Client).ocean.Spark().DetachVirtualNodeGroup(ctx, input); err != nil {
		return fmt.Errorf("[ERROR] Failed to detach VNG: %s", err)
	}

	return nil
}

func setupOceanSparkVirtualNodeGroupResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	ocean_spark_virtual_node_group.Setup(fieldsMap)

	commons.OceanSparkVirtualNodeGroupResource = commons.NewOceanSparkVirtualNodeGroupResource(fieldsMap)
}

package elastigroup_gcp_gpu

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[GPU] = commons.NewGenericField(
		commons.ElastigroupGCPGPU,
		GPU,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Count): {
						Type:     schema.TypeInt,
						Required: true,
					},

					string(Type): {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			//Set: hashKV,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(GPU)); ok {
				if gpu, err := expandGPU(v); err != nil {
					return err
				} else {
					elastigroup.Compute.SetGPU(gpu)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			var result *gcp.GPU = nil
			if v, ok := resourceData.GetOk(string(GPU)); ok {
				if gpu, err := expandGPU(v); err != nil {
					return err
				} else {
					result = gpu
				}
			}
			elastigroup.Compute.SetGPU(result)
			return nil
		},
		nil,
	)

}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utilities
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

// expandDisks sets the values from the plan
func expandGPU(data interface{}) (*gcp.GPU, error) {
	list := data.(*schema.Set).List()
	gpu := &gcp.GPU{}
	for _, item := range list {
		m := item.(map[string]interface{})

		if v, ok := m[string(Count)].(int); ok && v >= 0 {
			gpu.SetCount(spotinst.Int(v))
		}
		if v, ok := m[string(Type)].(string); ok && v != "" {
			gpu.SetType(spotinst.String(v))
		}
	}
	return gpu, nil
}

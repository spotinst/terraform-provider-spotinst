package elastigroup_gcp_disk

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

	fieldsMap[Disk] = commons.NewGenericField(
		commons.ElastigroupGCPDisk,
		Disk,
		&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(AutoDelete): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(Boot): {
						Type:     schema.TypeBool,
						Optional: true,
					},

					string(DeviceName): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(InitializeParams): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(DiskSizeGB): {
									Type:     schema.TypeString,
									Optional: true,
								},

								string(DiskType): {
									Type:     schema.TypeString,
									Optional: true,
								},

								string(SourceImage): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},

					string(Interface): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(Mode): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(Source): {
						Type:     schema.TypeString,
						Optional: true,
					},

					string(Type): {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Disk)); ok {
				if networks, err := expandDisks(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetDisks(networks)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(Disk)); ok {
				if networks, err := expandDisks(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetDisks(networks)
				}
			}
			return nil
		},
		nil,
	)

}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Utils
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

// expandDisks sets the values from the plan
func expandDisks(data interface{}) ([]*gcp.Disk, error) {
	list := data.(*schema.Set).List()
	disks := make([]*gcp.Disk, 0, len(list))

	for _, item := range list {
		m := item.(map[string]interface{})
		disk := &gcp.Disk{}

		if v, ok := m[string(AutoDelete)].(bool); ok {
			disk.SetAutoDelete(spotinst.Bool(v))
		}

		if v, ok := m[string(Boot)].(bool); ok {
			disk.SetBoot(spotinst.Bool(v))
		}

		if v, ok := m[string(DeviceName)].(string); ok && v != "" {
			disk.SetDeviceName(spotinst.String(v))
		}

		if v, ok := m[string(Interface)].(string); ok && v != "" {
			disk.SetInterface(spotinst.String(v))
		}

		if v, ok := m[string(InitializeParams)]; ok {
			params, err := expandInitParams(v)
			if err != nil {
				return nil, err
			}

			if params != nil {
				disk.SetInitializeParams(params)
			}
		}

		if v, ok := m[string(Mode)].(string); ok && v != "" {
			disk.SetMode(spotinst.String(v))
		}

		if v, ok := m[string(Source)].(string); ok && v != "" {
			disk.SetSource(spotinst.String(v))
		}

		if v, ok := m[string(Type)].(string); ok && v != "" {
			disk.SetType(spotinst.String(v))
		}
		disks = append(disks, disk)
	}
	return disks, nil
}

// expandInitParams sets the initialization params
func expandInitParams(data interface{}) (*gcp.InitializeParams, error) {
	list := data.(*schema.Set).List()
	param := &gcp.InitializeParams{}
	for _, item := range list {
		m := item.(map[string]interface{})

		if v, ok := m[string(DiskSizeGB)].(int); ok && v >= 0 {
			param.SetDiskSizeGB(spotinst.Int(v))
		}
		if v, ok := m[string(DiskType)].(string); ok && v != "" {
			param.SetDiskType(spotinst.String(v))
		}
		if v, ok := m[string(SourceImage)].(string); ok && v != "" {
			param.SetSourceImage(spotinst.String(v))
		}
	}
	return param, nil
}

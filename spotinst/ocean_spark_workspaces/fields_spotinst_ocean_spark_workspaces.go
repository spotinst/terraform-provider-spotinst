package ocean_spark_workspaces

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/spark"
	"github.com/spotinst/spotinst-sdk-go/spotinst"

	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[Workspaces] = commons.NewGenericField(
		commons.OceanSparkWorkspaces,
		Workspaces,
		block(map[string]*schema.Schema{
			string(Storage): block(map[string]*schema.Schema{
				string(Defaults): block(map[string]*schema.Schema{
					string(StorageClassName): {
						Type:     schema.TypeString,
						Optional: true,
						Description: "The name of the persistent volume storage class to use by default for new " +
							"workspaces. Leave it empty to use the cluster defaults.",
					},
				}),
			}),
		}),
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var result []interface{} = nil
			if cluster.Config != nil && cluster.Config.Workspaces != nil {
				result = flattenWorkspaces(cluster.Config.Workspaces)
			}
			if len(result) > 0 {
				if err := resourceData.Set(string(Workspaces), result); err != nil {
					return fmt.Errorf(commons.FailureFieldReadPattern, string(Workspaces), err)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if value, ok := resourceData.GetOk(string(Workspaces)); ok {
				if workspaces, err := expandWorkspaces(value); err != nil {
					return err
				} else {
					if cluster.Config == nil {
						cluster.Config = &spark.Config{}
					}
					cluster.Config.SetWorkspaces(workspaces)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.SparkClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *spark.WorkspacesConfig = nil
			if v, ok := resourceData.GetOk(string(Workspaces)); ok {
				if workspaces, err := expandWorkspaces(v); err != nil {
					return err
				} else {
					value = workspaces
				}
			}
			if cluster.Config == nil {
				cluster.Config = &spark.Config{}
			}
			cluster.Config.SetWorkspaces(value)
			return nil
		},
		nil,
	)
}

// region Flatten

func flattenWorkspaces(workspaces *spark.WorkspacesConfig) []interface{} {
	if workspaces == nil {
		return nil
	}
	m := make(map[string]interface{})
	if storage := flattenWorkspacesStorage(workspaces); len(storage) > 0 {
		m[string(Storage)] = storage
	}
	return toSlice(m)
}

func flattenWorkspacesStorage(workspaces *spark.WorkspacesConfig) []interface{} {
	if workspaces == nil {
		return nil
	}
	m := make(map[string]interface{})
	if defaults := flattenWorkspacesStorageDefaults(workspaces); len(defaults) > 0 {
		m[string(Defaults)] = defaults
	}
	return toSlice(m)
}

func flattenWorkspacesStorageDefaults(workspaces *spark.WorkspacesConfig) []interface{} {
	if workspaces == nil {
		return nil
	}
	m := make(map[string]interface{})
	if sc := spotinst.StringValue(workspaces.StorageClassOverride); sc != "" {
		m[string(StorageClassName)] = sc
	}
	return toSlice(m)
}

// endregion

// region Expand

func expandWorkspaces(data interface{}) (*spark.WorkspacesConfig, error) {
	workspaces := &spark.WorkspacesConfig{}
	m, ok := toMap(data)
	if !ok {
		return workspaces, nil
	}
	if err := expandWorkspacesStorage(workspaces, m[string(Storage)]); err != nil {
		return workspaces, err
	}
	return workspaces, nil
}

func expandWorkspacesStorage(workspaces *spark.WorkspacesConfig, data interface{}) error {
	m, ok := toMap(data)
	if !ok {
		return nil
	}
	return expandWorkspacesStorageDefaults(workspaces, m[string(Defaults)])
}

func expandWorkspacesStorageDefaults(workspaces *spark.WorkspacesConfig, data interface{}) error {
	m, ok := toMap(data)
	if !ok {
		return nil
	}
	if v, ok := m[string(StorageClassName)].(string); ok {
		val := spotinst.String(v)
		if v == "" {
			// validation expects nil when the intent is to use the default value
			val = nil
		}
		workspaces.SetStorageClassOverride(val)
	}
	return nil
}

// endregion

// region Utilities

func block(content map[string]*schema.Schema) *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		MaxItems: 1,
		Elem:     &schema.Resource{Schema: content},
	}
}

func toMap(data interface{}) (map[string]interface{}, bool) {
	l, ok := data.([]interface{})
	if !ok || len(l) == 0 || l[0] == nil {
		return nil, false
	}
	m, ok := l[0].(map[string]interface{})
	return m, ok
}

func toSlice(m map[string]interface{}) []interface{} {
	if len(m) == 0 {
		return nil
	}
	return []interface{}{m}
}

// endregion

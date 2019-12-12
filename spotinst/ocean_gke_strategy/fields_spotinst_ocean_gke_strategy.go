package ocean_gke_strategy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[DrainingTimeout] = commons.NewGenericField(
		commons.OceanGKEStrategy,
		DrainingTimeout,
		&schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *int = nil

			if cluster.Strategy != nil && cluster.Strategy.DrainingTimeout != nil {
				value = cluster.Strategy.DrainingTimeout
			}
			if value != nil {
				if err := resourceData.Set(string(DrainingTimeout), spotinst.IntValue(value)); err != nil {
					return fmt.Errorf(string(commons.FailureFieldReadPattern), string(DrainingTimeout), err)
				}
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOkExists(string(DrainingTimeout)); ok {
				cluster.Strategy.SetDrainingTimeout(spotinst.Int(v.(int)))
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.GKEClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var dt *int = nil

			if v, ok := resourceData.Get(string(DrainingTimeout)).(int); ok && v > 0 {
				dt = spotinst.Int(v)
			}

			cluster.Strategy.SetDrainingTimeout(dt)

			return nil
		},
		nil,
	)
}

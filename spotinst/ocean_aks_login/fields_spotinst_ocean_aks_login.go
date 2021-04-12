package ocean_aks_login

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	fieldsMap[SSHPublicKey] = commons.NewGenericField(
		commons.OceanAKSLogin,
		SSHPublicKey,
		&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil

			if cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil &&
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.Login != nil &&
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.Login.SSHPublicKey != nil {
				value = cluster.VirtualNodeGroupTemplate.LaunchSpecification.Login.SSHPublicKey
			}
			if err := resourceData.Set(string(SSHPublicKey), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(SSHPublicKey), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()

			if v, ok := resourceData.GetOk(string(SSHPublicKey)); ok {
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.Login.SetSSHPublicKey(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		nil,
	)

	fieldsMap[UserName] = commons.NewGenericField(
		commons.OceanAKSLogin,
		UserName,
		&schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var value *string = nil

			if cluster.VirtualNodeGroupTemplate != nil && cluster.VirtualNodeGroupTemplate.LaunchSpecification != nil &&
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.Login != nil &&
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.Login.UserName != nil {
				value = cluster.VirtualNodeGroupTemplate.LaunchSpecification.Login.UserName
			}
			if err := resourceData.Set(string(UserName), value); err != nil {
				return fmt.Errorf(string(commons.FailureFieldReadPattern), string(UserName), err)
			}

			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			if v, ok := resourceData.GetOk(string(UserName)); ok {
				cluster.VirtualNodeGroupTemplate.LaunchSpecification.Login.SetUserName(spotinst.String(v.(string)))
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			clusterWrapper := resourceObject.(*commons.AKSClusterWrapper)
			cluster := clusterWrapper.GetCluster()
			var userName *string = nil
			if v, ok := resourceData.GetOk(string(UserName)); ok && v != "" {
				userName = spotinst.String(v.(string))
			}
			cluster.VirtualNodeGroupTemplate.LaunchSpecification.Login.SetUserName(userName)
			return nil
		},
		nil,
	)
}

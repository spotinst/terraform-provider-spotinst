## Unreleased

## 1.209.1 (February, 05 2025)
BUG FIXE:
* resource/spotinst_elastigroup_aws: Fixed `instance_types_ondemand` field to accept null.

## 1.209.0 (February 05, 2025)
ENHANCEMENTS:
* resource/spotinst_ocean_spark: Added support for `workspaces` object.

## 1.208.0 (January, 23 2025)
FEATURES:
* **New Resource:** added new resource `resource/spotinst_credentials_azure` to support azure account onboarding to spotinst.

## 1.207.0 (January, 14 2025)
ENHANCEMENTS:
* resource/spotinst_ocean_aks_np: Added support for `vng_template_scheduling`, `logging` and `suspension_hours` object.

## 1.206.0 (January, 10 2025)
ENHANCEMENTS:
* resource/spotinst_ocean_gke_import: Added support for `auto_update` object.

## 1.205.1 (January, 07 2025)
FIXES:
* Upgraded packages go-getter from v1.7.5 to v1.7.6 and go-git/v5 from v5.12.0 to v5.13.0 to solve vulnerabilities.

## 1.205.0 (January, 05 2025)
ENHANCEMENTS:
* resource/spotinst_elastigroup_gcp: added `should_utilize_commitments`, `preferred_availability_zones` and `min_cpu_platform` fields.

## 1.204.0 (January, 2 2025)
FIXES:
* Upgraded packages golang.org/x/crypto from v0.26.0 to v0.31.0, golang.org/x/net from v0.28.0 to v0.33.0 to solve vulnerabilities and upgraded go version to 1.23.

## 1.203.0 (December, 31 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_ecs: added `fallback_to_od` field in `strategy` object.

## 1.202.0 (December, 18 2024)
ENHANCEMENTS:
* resource/spotinst_stateful_node_azure: Added support for `should_revert_to_od` in deallocation config block.

## 1.201.0 (December, 11 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aks_np_virtual_node_group: added support for `shutdown_hours` in `scheduling` block.

## 1.200.0 (December, 06 2024)
ENHANCEMENTS:
* resource/spotinst_stateful_node_azure: Added support for `excluded_vm_sizes` and `spot_size_attributes` fields in `vm_sizes` block.

## 1.199.3 (December, 03 2024)
ENHANCEMENTS:
* Updated crypto/rand package instead of math/rand as part of penetration testing.

## 1.199.2 (November, 26 2024)
BUG FIXES:
* resource/spotinst_ocean_gke_import: Fixed update of attribute `min_size` and `max_size`.

## 1.199.1 (November, 26 2024)
BUG FIXES:
* resource/spotinst_ocean_aws: Fixed `max_vcpu` and `max_memory_gib` fields to accept null.
* resource/spotinst_ocean_ecs: Fixed `max_vcpu` and `max_memory_gib` fields to accept null.

## 1.199.0 (November, 21 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aws: Added `reserved_enis` field to support max pods configuration.
* resource/spotinst_ocean_aws_launch_spec: Added `reserved_enis` field to support max pods configuration.

## 1.198.0 (November, 19 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_gke_import: Added support for `is_aggressive_scale_down_enabled` and `filters` fields.

## 1.197.1 (Nov, 14 2024)
NOTES:
* resource/spotinst_ocean_right_sizing_rule: Fixed document format for readability.

## 1.197.0 (Nov, 13 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_right_sizing_rule: Fixed rightsizing rule attach and detach workloads in subsequent updates.
* resource/spotinst_ocean_right_sizing_rule: Added support for `exclude_preliminary_recommendations`, `recommendation_application_hpa`, and `restart_replicas` attributes.

## 1.196.0 (October, 28 2024)
ENHANCEMENTS:
* resource/spotinst_elastigroup_azure_v3: Added support for `scheduling`, `health`, `load_balancer`, `secrets`, `security` blocks.

## 1.195.1 (Oct 23, 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: Added support for `respect_pdb` field under `roll_config`.

## 1.195.0 (October, 08 2024)
ENHANCEMENTS:
* resource/spotinst_elastigroup_azure_v3: Added support for `extensions` block to support azure extensions.

## 1.194.0 (September, 24 2024)
ENHANCEMENTS:
* resource/spotinst_elastigroup_azure_v3: Added support for `scaling_up_policy` and `scaling_down_policy`.

## 1.193.0 (September, 19 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_gke_import: Added support for `should_utilize_commitments` under strategy.

## 1.192.0 (September, 18 2024)
NOTES:
* Added controller v2 reference in ocean resources.

## 1.191.0 (September, 12 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: Added support for `utilize_commitments` and `utilize_reserved_instances` under strategy.

## 1.190.0 (August, 30 2024)
ENHANCEMENTS:
* resource/spotinst_elastigroup_azure_v3: Added support for `gallery_image` block in `image`.

## 1.189.0 (August, 29 2024)
ENHANCEMENTS:
* resource/spotinst_stateful_node_azure: Added support for `vm_admins` field in `strategy` block.

## 1.188.1 (August, 25 2024)
FIXES:
* Fixed documentation for boolean values to lower case.

## 1.188.0 (August, 23 2024)
ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: Added support for `restrict_single_az`, `auto_healing`, `dynamic_iops` and `dynamic_volume_size` fields.

## 1.187.0 (August, 20 2024)
FIXES:
* Upgraded dependency packages to solve vulnerabilities.

## 1.186.0 (August, 13 2024)
ENHANCEMENTS:
* resource/spotinst_elastigroup_gcp: Added support for `revert_to_preemptible` and `optimization_windows` under `strategy` object.

## 1.185.0 (August, 07 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: Added `draining_timeout` field in `strategy` block.

## 1.184.0 (August, 06 2024)
ENHANCEMENTS:
* resource/spotinst_stateful_node_azure: Added `encryption_at_host` and `confidential_os_disk_encryption` in `security` block.

## 1.183.0 (August, 06 2024)
ENHANCEMENTS:
* resource/spotinst_managed_instance_aws: Added support for `metadata_options`.

## 1.182.1 (August 06, 2024)
FIXES:
* resource/spotinst_ocean_aks_np: Upgraded Kubernetes version in unit tests.

## 1.182.0 (August 04, 2024)
FEATURES:
* **New Resource:** `resource/spotinst_account`
* **New Resource:** `resource/spotinst_credentials_gcp`

## 1.181.1 (July, 18 2024)
BUG FIXES:
* resource/spotinst_elastigroup_aws: Added null check before setting value of `required_gpu_minimum` and `required_gpu_maximum` under `resource_requirements`.

## 1.181.0 (July, 12 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_gke_launch_spec: Added support for `initial_nodes` under `create_options`.

## 1.180.3 (July, 11 2024)
ENHANCEMENTS:
* Upgraded dependency packages to solve vulnerabilities.

## 1.180.2 (July, 05 2024)
BUG FIXES:
* resource/spotinst_ocean_aws: Fixed `scheduled_task` object null issue with terraform refresh.

## 1.180.1 (July, 02 2024)
ENHANCEMENTS:
* Upgraded go-getter package from v1.6.2 to v1.7.5 and other dependency packages.

## 1.180.0 (June, 28 2024)
ENHANCEMENTS:
* resource/spotinst_elastigroup_azure_v3: Modified `od_sizes` and `spot_sizes` to be subfields of `vm_sizes` block.
* resource/spotinst_stateful_node_azure: Modified `od_sizes` and `spot_sizes` to be subfields of `vm_sizes` block.

## 1.179.0 (June, 27 2024)
FEATURES:
* **New Resource:** `resource/spotinst_ocean_right_sizing_rule`
* NOTE: This is internal release. It is not expected to be used by customers.

## 1.178.0 (June, 20 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: Added `preferred_od_types` field to support preferredOnDemandTypes list in vng.

## 1.177.0 (June, 18 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aks: Removed support for `spotinst_ocean_aks` resource as the api's are deprecated.

## 1.176.1 (June, 12 2024)
BUG FIXES:
* resource/spotinst_elastigroup_aws: Fixed conflicts between `instance_types_spot` and `resource_requirements` for group updating. 

## 1.176.0 (June 11, 2024)
ENHANCEMENTS:
* resource/spotinst_stateful_node_azure: Added `spot_account_id` attribute under `gallery` object and `should_deregister_from_lb` field.

## 1.175.0 (May 22, 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aks_np: Added `tasks` object under `scheduling` block supporting schedule roll.

## 1.174.0 (May, 22 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aks_np: added `linux_os_config` block as part of node pool properties of virtual node group template to set maximum number of memory areas a process may have,
* resource/spotinst_ocean_aks_np_virtual_node_group: added `linux_os_config` block as part of node pool properties to set maximum number of memory areas a process may have,

## 1.173.0 (May 21, 2024)
FEATURES: Added the below resources
* **New Resource:** `resource/spotinst_oceancd_verification_provider`
* **New Resource:** `resource/spotinst_oceancd_verification_template`
* **New Resource:** `resource/spotinst_oceancd_strategy`
* **New Resource:** `resource/spotinst_oceancd_rollout_spec`

## 1.172.3 (May, 17 2024)
BUG FIXES:
* resource/spotinst_ocean_aws: Fixed `autoscale_headroom` block to set to null when underlying attributes are not passed in config.

## 1.172.2 (May, 17 2024)
BUG FIXES:
* resource/spotinst_ocean_aws: reverting th fix done for `autoscale_headroom` block.

## 1.172.1 (May, 16 2024)
BUG FIXES:
* resource/spotinst_ocean_aws: Fixed disabling of `autoscale_headroom` object and its attributes in cluster config under `autoscaler`.

## 1.172.0 (May, 08 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aws: added `attach_load_balancer` and `detach_load_balancer` blocks support for attaching and detaching loadBalancers to ocean aws cluster.

## 1.171.4 (May, 06 2024)
BUG FIXES:
* resource/spotinst_ocean_gke_import: Fixed disabling `parameters` under `tasks` in cluster config .

## 1.171.3 (May, 03 2024)
BUG FIXES:
* resource/spotinst_ocean_gke_import: Fixed disabling `tasks` and `shutdown_hours` from cluster config under `scheduled_task`.

## 1.171.2 (May, 02 2024)
BUG FIXES:
* Changed `policy_account_ids` field to Optional in `resource_spotinst_programmatic_user`

## 1.171.1 (April, 25 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aws: modified `image_id` field from optional to required according to api changes
* resource/spotinst_ocean_ecs: modified `image_id` field from optional to required according to api changes

## 1.171.0 (April, 22 2024)
ENHANCEMENTS:
* resource/spotinst_elastigroup_azure: removed support for Elastigroup azure v2 resource.

## 1.170.1 (April, 22 2024)
BUG FIXES:
* Corrected `kubernetes_version` field in test files `resource_spotinst_ocean_aks_np_test` and `resource_spotinst_ocean_aks_np_virtual_node_group_test`

## 1.170.0 (April, 19 2024)
ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: removed support for `MLB`, `MLB_RUNTIME`, `MULTAI_TARGET_SET` from `health_check_type` and `integrations`.
* resource/spotinst_managed_instance_aws: removed support for `MULTAI_TARGET_SET` from `health_check_type` and `integrations`.

## 1.169.1 (April, 10 2024)
BUG FIXES:
* resource/spotinst_ocean_aks_np_virtual_node_group: Fixed `labels` failing to update.
* resource/spotinst_ocean_aws: Fixed drift by adding default value for `health_check_unhealthy_duration_before_replacement`.

## 1.169.0 (April, 08 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aks_np: Added support for `gpu_types` field in `filters` block.
* resource/spotinst_ocean_aks_np_virtual_node_group: Added support for `gpu_types` field in `filters` block.

## 1.168.3 (April, 05 2024)
NOTES:
* Removed redundant declaration of `image_id` in testcases of `resource_spotinst_ocean_ecs_test`.

## 1.168.2 (April, 04 2024)
NOTES:
* Added `image_id` in config for unit tests of `resource_spotinst_ocean_ecs_test`.

## 1.168.1 (April, 02 2024)
BUG FIXES:
* Fix the release Github Actions failure caused by Go version upgrade.

## 1.168.0 (April, 02 2024)
ENHANCEMENTS:
* Upgraded Go version to 1.20 and tfproviderlint package to v0.29.0

## 1.167.1 (March, 29 2024)
BUG FIXES:
* resource/spotinst_ocean_aws_launch_spec: Fixed `ephemeral_storage_device_name` under `ephemeral_storage`, defaults to correct values when not configured.

## 1.167.0 (March, 29 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: Added support for `ephemeral_storage`.

## 1.166.0 (March, 28 2024)
ENHANCEMENTS:
* resource/spotinst_stateful_node_azure: Added support for `caching`, `license_type`, `availability_vs_cost`, `vm_name_prefix`, `od_windows` fields.

## 1.165.1 (March, 27 2024)
BUG FIXES:
* resource/spotinst_elastigroup_gcp: Fixed `instance_types_custom` object, as it was throwing error while creating the EG.

# 1.165.0 (March, 19 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aws: Added support for `health_check_unhealthy_duration_before_replacement` attribute under launchSpecification.

## 1.164.0 (March, 15 2024)
ENHANCEMENTS:
* resource/spotinst_managed_instance_aws: Added support for `snapshot_id`, `encrypted` and `kms_key_id` attributes in `block_device_mappings` block.
* config: added `enabled` field in provider config as an optional parameter. 

## 1.163.0 (March, 15 2024)
BUG FIXES:
* resource/spotinst_elastigroup_aws: Modified `wait_for_roll_timeout` and `wait_for_roll_percentage` fields as optional.

## 1.162.0 (February, 16 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aks_np: Added support to trigger cluster roll on attribute update.
* resource/spotinst_ocean_aks_np_virtual_node_group: Added support to trigger vng roll on attribute update.

# 1.161.0 (February, 15 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aws: Added support for `is_aggressive_scale_down_enabled` attribute under `autoscale_down` block

# 1.160.2 (February, 08 2024)
BUG FIX:
* resource/spotinst_ocean_aws: Fixed `scheduled_task` block to accept null

# 1.160.1 (January, 31 2024)
BUG FIX:
* resource/spotinst_ocean_aws: Updated default value for `batch_min_healthy_percentage` and `batch_size_percentage` and fixed issue to ignore the unchanged difference in `parameters` block

# 1.160.0 (January, 24 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aws: Added support for `parameters` object under `scheduling_tasks` block

## 1.159.0 (January, 17 2024)
NOTES:
* Added random password generation logic for unit test `resource_spotinst_organization_user_test`

## 1.158.0 (January, 05 2024)
BUG FIXES:
* resource/spotinst_managed_instance_aws: Fixed `fallback_to_ondemand` field variable name in `strategy` object.

## 1.157.0 (January, 02 2024)
ENHANCEMENTS:
* resource/spotinst_ocean_aws: Added `conditioned_roll_params` field to customized attribute modification to trigger cluster roll.

## 1.156.0 (December, 19 2023)
NOTES:
* Added unit tests for `spotinst_organization_policy`, `spotinst_organization_user`, `spotinst_organization_user_group` and `spotinst_organization_programmatic_user` resources

## 1.155.0 (December, 11 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aks_np: Added `is_enabled` field support for `automatic` object inside `auto_scaler`

## 1.154.0 (December, 13 2023)
NOTES:
* Added unit tests for `spotinst_ocean_aks_np` and `spotinst_ocean_aks_np_virtual_node_group` resources

## 1.153.1 (December, 12 2023)
NOTES:
* documentation: resource/spotinst_elastigroup_aws: updated document to add `period` and `evaluation_periods` attributes in `target_scaling_policy`.

## 1.153.0 (December, 08 2023)
ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: Added `logging` object support.

## 1.152.0 (December, 07 2023)
NOTES:
* Migration of all unit tests to run on new test account.

## 1.151.1 (December, 06 2023)
NOTES:
* documentation: resource/spotinst_subscription: Added ocean specific events in document for event_type fields.

## 1.151.0 (December, 02 2023)
ENHANCEMENTS:
* resource/spotinst_stateful_node_azure: Added `proximity_placement_groups` object support.
* 

## 1.150.1 (November, 21 2023)
BUG FIXES:
* resource/spotinst_ocean_ecs_launch_spec: Fixed deletion of `images` object.

## 1.150.0 (November, 16 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_ecs_launch_spec: added `images` object to support multi ami

## 1.149.0 (November, 08 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aks_np: removed `isEnabled` and default values
* resource/spotinst_multai_target: removed the support

## 1.148.0 (October 23, 2023)
BUG FIXES:
* resource/spotinst_stateful_node_azure: Fixed `load_balancer_config`, `login` and `os_disk` blocks in `launch_specification` object

## 1.147.0 (October, 18 2023)
BUG FIXES::
* resource/spotinst_ocean_aks_np: Fixed `accelerated_networking` and `disk_performance` fields in `filters` object to accept null
* resource/spotinst_ocean_aks_np_virtual_node_group: Fixed `accelerated_networking` and `disk_performance` fields in `filters` object to accept null

## 1.146.0 (October, 14 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aks_np: Added `accelerated_networking`, `disk_performance`, `min_gpu`, `max_gpu`, `min_nics`, `min_data` and `vm_types` fields in `filters` object
* resource/spotinst_ocean_aks_np_virtual_node_group: Added `accelerated_networking`, `disk_performance`, `min_gpu`, `max_gpu`, `min_nics`, `min_data` and `vm_types` fields in `filters` object

## 1.145.0 (October, 12 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aks_np: Added `pod_subnet_ids` and `vnet_subnet_ids` fields in `node_pool_properties` object
* resource/spotinst_ocean_aks_np_virtual_node_group: Added `pod_subnet_ids` and `vnet_subnet_ids` fields in `node_pool_properties` object

## 1.144.0 (October, 09 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: Added `instancetypes_filters` object under `launchSpec`
BUG FIXES:
* resource/spotinst_ocean_aws: Fix for `autoscale_is_enabled` field update under filters.
NOTES:
* documentation: resource/spotinst_credentials_aws: Moved credentials document under Accounts hierarchy.

## 1.143.0 (October 05, 2023)
FEATURES: added new resources
* **New Resource:** `resource/spotinst_organization_user`
* **New Resource:** `resource/spotinst_organization_policy`
* **New Resource:** `resource/spotinst_organization_programmatic_user`
* **New Resource:** `resource/spotinst_organization_user_group`

## 1.142.0 (September 29, 2023)
FEATURES:
* **New Resource:** `resource/spotinst_account_aws`
* **New Resource:** `resource/spotinst_credentials_aws`

## 1.141.0 (September, 29 2023)
BUG FIXES:
* resource/spotinst_ocean_aks_np: Fix for `os_sku` and `availability_zones` fields

## 1.140.0 (September, 22 2023)
BUG FIXES:
* resource/spotinst_ocean_aws: Fix for ignoring the unnecessary changes shown during terraform plan for the attributes inside `filters` object
* resource/spotinst_ocean_ecs: Fix for ignoring the unnecessary changes shown during terraform plan for the attributes inside `filters` object

## 1.139.0 (September, 15 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aws: Added `dynamic_iops` object in `ebs`

## 1.138.0 (September, 13 2023)
NOTES:
* resource/spotinst_ocean_aks_np: Added basic cluster creation usage example

## 1.137.0 (September, 13 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aks_np: Added `kubernetes_version` attribute support in cluster and vng
BUG FIXES:
* resource/spotinst_ocean_aks_np: Fixed default values in virtual_node_groups_template object

## 1.136.0 (September 07, 2023)
ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: Added `max_scale_down_percentage` field in `kubernetes_integrations` object

## 1.135.0 (September, 06 2023)
BUG FIXES:
* resource/spotinst_ocean_aks_np: Fix for ignoring the unnecessary changes shown during terraform plan for the attributes inside `filters` object
* resource/spotinst_ocean_aks_np_virtual_node_group: Fix for ignoring the unnecessary changes shown during terraform plan for the attributes inside `filters` object

## 1.134.0 (August 18, 2023)
ENHANCEMENTS:
* resource/spotinst_stateful_node_azure: Added `capacity_reservation` block in `strategy` object

## 1.133.0 (August 12, 2023)
BUG FIXES:
* resource/spotinst_stateful_node_azure: Fix for allowing `data_disks`, `os_disk` and `network` blocks to be modified when `persistency` is modified
NOTES:
* documentation: Updated description of `controller_cluster_id`, `os_type`, `series` and added detailed description for `automatic`, `shutdown_hours`, `labels`, `taints` in documentation for `spotinst_ocean_aks_np`

## 1.132.0 (August 07, 2023)
ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: Added `resource_requirements` block in `instance_types` object
* resource/spotinst_ocean_aks_np: Added `exclude_series` in `vm_sizes` object

## 1.131.0 (August 01, 2023)
ENHANCEMENTS:
* resource/spotinst_stateful_node_azure: Added `security` block in `launch_specification` object
  NOTES:
* documentation: Added `delete` usage to the `spotinst_stateful_node_azure` documentation

## 1.130.0 (July 27, 2023)
BUG FIXES:
* resource/spotinst_elastigroup_gcp: Fix for allowing `named_ports` to be configured when `location_type` is regional

## 1.129.0 (July 26, 2023)
BUG FIXES:
* resource/spotinst_stateful_node_azure: Modified `network`, `image` and `login` blocks as optional to support import workflow

## 1.128.0 (July 24, 2023)
ENHANCEMENTS:
* resource/spotinst_stateful_node_aws: Exposing `deallocation_config` to provide an option to the user to choose whether to keep the underlying resources alive or not using `delete` block

## 1.127.0 (July 19, 2023)
ENHANCEMENTS:
* resource/spotinst_stateful_node_azure: Added `vm_name` field in `launch_spec` object

## 1.126.0 (July 17, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aws: added `should_tag_volumes` attribute in `resource_tag_specification` object

## 1.125.1 (July 11, 2023)
NOTES:
* documentation: Modified description of `spot_percentage` in documentation for `spotinst_ocean_aws` and `spotinst_ocean_aws_launch_spec`

## 1.125.0 (Jun 30, 2023)
ENHANCEMENTS:
* resource/spotinst_stateful_node_azure: Added `user_data` field in `launch_spec` object

## 1.124.0 (Jun 28, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_spark: Fixed "Delete cluster waits for the cluster to be deleted if forceDeleted"

## 1.123.0 (Jun 20, 2023)
NOTES:
* documentation: Added documentation for `spotinst_ocean_aks_np` and `spotinst_ocean_aks_np_virtual_node_group`

## 1.122.2 (Jun 08, 2023)
NOTES:
* documentation: resource/spotinst_elastigroup_aws: Fixed Tests - removed `statefulUpdateCapacity` taskType from `scheduled_task` test

## 1.122.1 (Jun 07, 2023)
NOTES:
* documentation: resource/spotinst_ocean_spark: Fixed Tests - removed deprecated `collect_driver_log` property

## 1.122.0 (Jun 05, 2023)
NOTES:
* documentation: resource/spotinst_elastigroup_aws: added `gp3` to the list of supported `volume_type`

## 1.121.0 (Jun 01, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aks_np: Added `vm_sizes` object support
* resource/spotinst_ocean_aks_np_virtual_node_group: Added `vm_sizes` object support

## 1.120.0 (May 27, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aks_np: new fields supported in AKS cluster update

## 1.119.1 (May 25, 2023)
NOTES:
* documentation: Moved `stateful_node_azure` under `Stateful Node` category (documentation change only, terraform resource not modified)

## 1.119.0 (May 23, 2023)
ENHANCEMENTS:
* resource/spotinst_elastigroup_azure_v3: added `tags` object

## 1.118.0 (May 23, 2023)
NOTES:
* documentation: Renamed `managed_instance` to `stateful_node_aws` (documentation change only, terraform resource not modified)
* documentation: Renamed `elastigroup_azure_v3` to `elastigroup_azure` (documentation change only, terraform resource not modified)
* documentation: Removed `elastigroup_azure_v2` from documentation

## 1.117.0 (May 09, 2023)
BUG FIXES:
* resource/spotinst_ocean_gke_launch_spec: fixed `network_interfaces` block for vng import flow

NOTES:
* documentation: resource/spotinst_stateful_node_azure: fixed `public_ip_sku` value

## 1.116.0 (May 06, 2023)
BUG FIXES:
* resource/spotinst_ocean_aks_np: corrected update cluster route

## 1.115.0 (Apr 28, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_spark: added `collectAppLogs` object

## 1.114.0 (Apr 27, 2023)
NOTES:
* documentation: resource/spotinst_ocean_aws_launch_spec: updated example usage for `images`

## 1.113.0 (Apr 20, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_gke_launch_spec: added `network_interfaces` object 
* documentation: resource/spotinst_ocean_aws_launch_spec: added `images` object in the documentation

## 1.112.0 (Apr 13, 2023)
BUG FIXES:
* resource/spotinst_elastigroup_aws: fixed `autoscale_attributes` field in `integration_ecs` to read as an array

## 1.111.0 (Apr 13, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_ecs: added `enable_automatic_and_manual_headroom` field in `autoscaler`

## 1.110.0 (Apr 06, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: added `images` object

## 1.109.0 (Apr 5, 2023)
ENHANCEMENTS:
* resource/spotinst_elastigroup_gcp: added `instance_name_prefix` field in `launch_specification`

## 1.108.0 (Mar 22, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aws: added `block_device_mappings` object in the `launch_specification`

## 1.107.0 (Mar 22, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_gke: added `respect_pdb` field for create and roll cluster `scheduling`

## 1.106.1 (Mar 21, 2023)
BUG FIXES:
* resource/spotinst_ocean_gke: Commented warnings `Please add the imported tags from state file to the tags list`

## 1.106.0 (Mar 14, 2023)
ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: added `instance_metadata_tags` in `metadata_options`

## 1.105.0 (Mar 9, 2023)
FEATURES:
* **New Resource:** `resource/spotinst_ocean_aks_np_virtual_node_group`

## 1.104.0 (Mar 8, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aws: added `associate_ipv6_address`

## 1.103.0 (Mar 3, 2023)
BUG FIXES:
* resource/spotinst_elastigroup_azure_v3: resolved errors with `spot_percentage` and `on_demand_count`

## 1.102.0 (Mar 1, 2023)
FEATURES:
* **New Resource:** `resource/spotinst_ocean_aks_np`

## 1.101.0 (Feb 27, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_spark: added `additional_app_namespaces`
* resource/spotinst_ocean_spark: added support for ocean spark cluster dedicated VNGs.

## 1.100.0 (Feb 15, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aws: added `spread_nodes_by`

## 1.99.0 (Feb 14, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_ecs: added `should_scale_down_non_service_tasks`

## 1.98.0 (Feb 14, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_ecs_launch_spec: added `instance_metadata_options`
* resource/spotinst_ocean_aws_launch_spec: added `instance_metadata_options`

NOTES:
* documentation: resource/spotinst_ocean_spark: Fixed Tests - use the valid
 domain name for custom endpoint address

## 1.97.0 (Jan 26, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_ecs: added `cluster_orientation`

## 1.96.0 (Jan 24, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: added `max_scale_down_percentage`

## 1.95.2 (Jan 20, 2023)
NOTES:
* documentation: resource/elastigroup_aws_suspension: Fixed typos in documentation 
* documentation: resource/ocean_ecs: Fixed documentation for `instance_types`

## 1.95.1 (Jan 19, 2023)
NOTES:
* documentation: resource/ocean_gke_launch_spec: Fixed typos in documentation 
* documentation: resource/ocean_ecs: Added example for import cluster to ocean
* documentation: resource/ocean_aws: Added example for import cluster to ocean

## 1.95.0 (Jan 18, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_ecs: added `filter`
* resource/spotinst_ocean_ecs: updated `isEnaSupported`

## 1.94.0 (Jan 17, 2023)
NOTES:
* documentation: resource/spotinst_elastigroup_gcp: corrected example usage

## 1.93.0 (Jan 17, 2023)
ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: added `immediate_od_recover_threshold`

## 1.92.0 (Jan 17, 2023)
ENHANCEMENTS:
* resource/spotinst_ocean_aws: added `cluster_orientation`

## 1.91.0 (Jan 12, 2023)
ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: added `consider_od_pricing`

## 1.90.0 (Dec 20, 2022)
NOTES:
* documentation: resource/spotinst_ocean_aws_instance_types: update `filters`

## 1.89.0 (Dec 20, 2022)
ENHANCEMENTS:
* resource/spotinst_ocean_spark: added cluster ingress configs

## 1.88.1 (Dec 20, 2022)
NOTES:
* documentation: resource/spotinst_ocean_aws_instance_types: added `filters`

## 1.88.0 (Dec 20, 2022)
ENHANCEMENTS:
* resource/spotinst_ocean_aws_instance_types: added `filters`

## 1.87.1 (Nov 30, 2022)

BUG FIXES:

* resource/spotinst_ocean_aws_launch_spec: resolved errors with `resource_limits`
* resource/spotinst_ocean_gke_launch_spec: resolved errors with `resource_limits`

## 1.87.0 (Nov 02, 2022)
ENHANCEMENTS:
* resource/spotinst_ocean_ecs_launch_spec: added `spot_percentage`

## 1.86.0 (Oct 31, 2022)
ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: added `delete_nodes`

BUG FIXES:
* resource/spotinst_ocean_aws_launch_spec: resolved errors with `max_instance_count` and `min_instance_count`

## 1.85.1 (Oct 19, 2022)

NOTES:
* documentation: add preferred_spot_types in resource/spotinst_ocean_ecs_launcspec
## 1.85.0 (Oct 12, 2022)

NOTES:
* documentation: remove blacklist from resource/spotinst_ocean_ecs

## 1.84.0 (Sep 19, 2022)

BUG FIXES:

* resource/spotinst_elastigroup_aws: resolved errors with `scaling_policies`
* * resource/spotinst_elastigroup_azure: resolved errors with `scaling_policies`
* resource/spotinst_elastigroup_gcp: resolved errors with `scaling_policies`

## 1.83.0 (Sep 13, 2022)

FEATURES:
* **New Resource:** `spotinst_ocean_spark`

ENHANCEMENTS:
* resource/spotinst_ocean_ecs_launch_spec: added `preferred_spot_types`

## 1.82.0 (Sep 06, 2022)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: added `images`

## 1.81.0 (Aug 04, 2022)

ENHANCEMENTS:
* resource/spotinst_ocean_gke_launch_spec: added `tags`
* resource/spotinst_ocean_aks_virtual_node_group: added `max_pods`
* resource/spotinst_ocean_aks: added `max_pods`

## 1.80.0 (July 25, 2022)

ENHANCEMENTS:
* resource/spotinst_ocean_gke_import: added `blacklist`

## 1.79.0 (July 07, 2022)

ENHANCEMENTS:
* resource/spotinst_ocean_aws: added `respect_pdb`

## 1.78.0 (June 30, 2022)

ENHANCEMENTS:
* resource/spotinst_ocean_ecs: added `use_as_template_only`
* resource/spotinst_ocean_gke_import: added `use_as_template_only`

## 1.77.0 (June 23, 2022)

BUG FIXES:

* resource/spotinst_elastigroup_aws: resolved errors with `scaling_policies`

## 1.76.0 (June 01, 2022)

NOTES:
* Upgrade terraform-plugin-sdk version to v2.5.0

## 1.75.0 (May 26, 2022)

ENHANCEMENTS:
* resource/spotinst_ocean_ecs: added support for `logging`
* resource/spotinst_ocean_aks: added support for `zones`
* resource/spotinst_ocean_aks_virtual_node_group: added support for `zones`
* resource/spotinst_ocean_aks_virtual_node_group: added support for `utilize_ephemeral_storage`

## 1.74.0 (May 22, 2022)

FEATURES:
* **New Resource:** `spotinst_stateful_node_azure`

## 1.73.3 (May 09, 2022)

BUG FIXES:

* resource/spotinst_ocean_aws_launch_spec: resolved errors with `max_instance_count`


## 1.73.2 (April 28, 2022)

BUG FIXES:

* resource/spotinst_ocean_gke_launch_spec: resolved errors with spotist-go-sdk dependencies

## 1.73.1 (April 19, 2022)

NOTES:
* documentation: resource/spotinst_data_integration

## 1.73.0 (April 19, 2022)

FEATURES:
**New Resource:** `spotinst_data_integration`

## 1.72.0 (April 12, 2022)

BUG FIXES:
* resource/spotinst_ocean_aws: resolved errors with `auto_headroom_percentage`

## 1.71.0 (April 04, 2022)

ENHANCEMENTS:
* resource/spotinst_ocean_gke_import: added support for `preemptible_percentage`

## 1.70.0 (March 29, 2022)

ENHANCEMENTS:
* resource/spotinst_ocean_aws: added support for `batch_min_healthy_percentage`
* resource/spotinst_ocean_ecs: added support for `batch_min_healthy_percentage`
* resource/spotinst_ocean_gke_import: added support for `batch_min_healthy_percentage`

## 1.69.0 (March 08, 2022)

ENHANCEMENTS:
* resource/spotinst_ocean_gke_import: added support for `shielded_instance_config`
* resource/spotinst_ocean_aws_launch_spec: added support for `auto_headroom_percentage`
* resource/spotinst_ocean_gke_launch_spec: added support for `auto_headroom_percentage`
* resource/spotinst_ocean_aks_virtual_node_group: added support for `auto_headroom_percentage`

FEATURES:
* **New Resource:** `spotinst_ocean_aws_extended_resource_definition`

## 1.68.0 (February 20, 2022)

ENHANCEMENTS:
* resource/spotinst_ocean_gke_import: added support for `provisioning_model`.
* resource/spotinst_ocean_gke_import: added support for `draining_timeout`.

## 1.67.1 (February 14, 2022)

BUG FIXES:
* resource/spotinst_ocean_aws: resolved errors with `extended_resource_definitions`

## 1.67.0 (February 14, 2022)

ENHANCEMENTS:
* resource/spotinst_ocean_aws: added support for `extended_resource_definitions`

## 1.66.0 (February 07, 2022)

ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: added support `shut_down_hours`
* resource/spotinst_elastigroup_aws: added support for `static_target_group`
* resource/spotinst_elastigroup_aws: added support for `default_static_target_group`

## 1.65.0 (January 10, 2022)

ENHANCEMENTS:
* resource/spotinst_ocean_aws: added support for `auto_apply_tags`
* resource/spotinst_ocean_ecs: added support for `auto_apply_tags`
* resource/spotinst_ocean_aws: added support for `enable_automatic_and_manual_headroom`
* resource/spotinst_ocean_gke_import: added support for `enable_automatic_and_manual_headroom`

## 1.64.2 (January 05, 2022)

NOTES:
* documentation: resource/spotinst_ocean_aws: Add usage example for `utilize_commitments`
* documentation: resource/spotinst_ocean_ecs: Add usage example for `utilize_commitments`

BUG FIXES:
* resource/spotinst_elastigroup_aws: resolved errors with `ebs_block_device`

## 1.64.1 (December 14, 2021)

BUG FIXES:
* resource/spotinst_elastigroup_aws: resolved errors with `ebs_block_device`

## 1.64.0 (December 09, 2021)

ENHANCEMENTS:
* resource/spotinst_ocean_aws: added support for `logging`
* resource/spotinst_ocean_gke_import: added support for `conditioned_roll`
* resource/spotinst_ocean_aws: added support for `conditioned_roll`
* resource/spotinst_ocean_ecs: added support for `conditioned_roll`

BUG FIXES:
* resource/spotinst_managed_instance_aws: resolved errors with `persist_block_devices`

## 1.63.0 (November 29, 2021)

ENHANCEMENTS:
* resource/spotinst_ocean_gke_launch_spec: added support for `scheduling`
* resource/spotinst_ocean_aws_launch_spec: added support for `scheduling`
* resource/spotinst_ocean_ecs_launch_spec: added support for `scheduling`

## 1.62.0 (November 01, 2021)

ENHANCEMENTS:
* resource/spotinst_elastigroup_azure_v3: added support for `managed_service_identity`

## 1.61.1 (October 21, 2021)

ENHANCEMENTS:
* resource/spotinst_ocean_gcp_launch_spec: added `update_policy` for managing rolling deployments
* resource/spotinst_elastigroup_gcp: added support for `provisioning_model`
* resource/spotinst_elastigroup_aws: added support for `evaluation_periods`, `period`

## 1.60.0 (September 30, 2021)

ENHANCEMENTS:
* resource/spotinst_ocean_gke_launch_spec: added support for `min_instance_count`
* resource/spotinst_ocean_gke_launch_spec: added support for `name`
* resource/spotinst_ocean_aws_launch_spec: added support for `min_instance_count`
* resource/spotinst_ocean_ecs: added support for `auto_headroom_percentage`

## 1.59.1 (September 19, 2021)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: resolved errors with `preferred_availability_zones`

BUG FIXES:
* resource/spotinst_ocean_aws_launch_spec: resolved errors with `effect`

## 1.58.0 (September 13, 2021)

ENHANCEMENTS:
* resource/spotinst_ocean_gke_import: added `update_policy` for managing rolling deployments
* resource/spotinst_ocean_aws: added support for `launch_spec_ids` in `update_policy`
* resource/spotinst_ocean_aws_launch_spec: added `update_policy` for managing rolling deployments
* resource/spotinst_ocean_aws_launch_spec: added support for `force_delete`
* resource/spotinst_elastigroup_azure_v3: added support for `custom_data`

BUG FIXES:
* resource/spotinst_elastigroup_aws: fix `resource_tag_specification` field statement in wrapper method
* resource/spotinst_managed_instance_aws: fix `resource_tag_specification` field statement in wrapper method

## 1.57.0 (August 22, 2021)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: added support for `multiple_metrics`
* resource/spotinst_elastigroup_aws: added support for `step_adjustments`

## 1.56.1 (August 9, 2021)

BUG FIXES:
* resource/spotinst_mrscaler_aws: fix `instanceGroupType` json tag ([spotinst/spotinst-sdk-go#136](https://github.com/spotinst/spotinst-sdk-go/pull/136))

## 1.56.0 (August 5, 2021)

ENHANCEMENTS:
* resource/spotinst_ocean_ecs: added support for `spot_percentage`
* resource/spotinst_ocean_ecs: added support for `instance_metadata_options`
* resource/spotinst_ocean_aws: added support for `instance_metadata_options`

## 1.55.0 (August 2, 2021)

ENHANCEMENTS:
* resource/spotinst_ocean_gke_launch_spec: added support for: `shielded_instance_config`
* resource/spotinst_ocean_gke_launch_spec: added support for: `enable_secure_boot`
* resource/spotinst_ocean_gke_launch_spec: added support for: `enable_integrity_monitoring`
* resource/spotinst_ocean_gke_launch_spec: added support for: `storage`
* resource/spotinst_ocean_gke_launch_spec: added support for: `local_ssd_count`
* resource/spotinst_ocean_gke_launch_spec: added support for: `resource_limits`
* resource/spotinst_ocean_gke_launch_spec: added support for: `max_instance_count`
* resource/spotinst_ocean_gke_launch_spec: added support for: `service_account`

## 1.54.0 (July 26, 2021)

ENHANCEMENTS:
* resource/spotinst_ocean_gke_import: added support for `root_volume_type`
* resource/spotinst_elastigroup_aws: added support for `resource_tag_specification`
* resource/spotinst_managed_instance_aws: added support for `resource_tag_specification`

## 1.53.1 (July 15, 2021)

BUG FIXES:
* resource/spotinst_*: ensure same state on all retries

## 1.53.0 (July 15, 2021)

ENHANCEMENTS:
* resource/spotinst_ocean_aks: added support for `managed_service_identity` 

## 1.52.0 (July 7, 2021)

ENHANCEMENTS:
* resource/spotinst_ocean_gke_launch_spec: added support for merging between the imported data and the user explicitly inserted data
* resource/spotinst_ocean_gke_launch_spec: added support for: `root_volume_size`
* resource/spotinst_ocean_gke_launch_spec: added support for: `root_volume_type`
* resource/spotinst_ocean_gke_launch_spec: added support for: `instance_types`
* resource/spotinst_mrscaler_aws: added support for: `master_target`

## 1.51.0 (July 4, 2021)

BUG FIXES:
* resource/spotinst_elastigroup_aws: resolved errors with `spot_percentage`

## 1.50.0 (June 21, 2021)

ENHANCEMENTS:
* resource/spotinst_ocean_ecs_launch_spec: added `subnet_ids`

BUG FIXES:
* resource/spotinst_ocean_aws: enable setting `spot_percentage` to 0
* resource/spotinst_elastigroup_azure_v3: resolved errors with `network`

## 1.49.0 (June 13, 2021)

ENHANCEMENTS:
* resource/spotinst_managed_instance_aws: added support for managed instance actions: `pause`, `resume`, `recycle`
* resource/spotinst_managed_instance_aws: added default deletion configuration to managed instances

## 1.48.1 (June 10, 2021)

BUG FIXES:
* resource/spotinst_ocean_aks_virtual_node_group: make headroom fields optional

## 1.48.0 (June 10, 2021)

ENHANCEMENTS:
* resource/spotinst_ocean_aks: added support for `resource_group_name` 
* resource/spotinst_ocean_aks: added support for `custom_data`
* resource/spotinst_ocean_aks: added support for `vm_sizes` 
* resource/spotinst_ocean_aks: added support for `os_disk` 
* resource/spotinst_ocean_aks: added support for `image` 
* resource/spotinst_ocean_aks: added support for `strategy` 
* resource/spotinst_ocean_aks: added support for `health` 
* resource/spotinst_ocean_aks: added support for `network` 
* resource/spotinst_ocean_aks: added support for `extension` 
* resource/spotinst_ocean_aks: added support for `load_balancer` 
* resource/spotinst_ocean_aks: added support for `autoscaler`
* resource/spotinst_ocean_aks: added support for `tag`
		
## 1.47.0 (June 6, 2021)

ENHANCEMENTS:
* resource/spotinst_elastigroup_azure_v3: added support for `application_security_group` ([#196](https://github.com/spotinst/terraform-provider-spotinst/pull/196))

## 1.46.0 (June 3, 2021)

ENHANCEMENTS:
* resource/spotinst_managed_instance_aws: added support for `minimum_instance_lifetime` ([#193](https://github.com/spotinst/terraform-provider-spotinst/pull/193))

## 1.45.1 (June 2, 2021)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: fix(elastigroup/aws): stateful: fix log messages

## 1.45.0 (June 2, 2021)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: added support for stateful instance actions: `pause`, `resume`, `recycle`, `deallocate` ([#192](https://github.com/spotinst/terraform-provider-spotinst/pull/192))

## 1.44.1 (May 28, 2021)

BUG FIXES:
* resource/spotinst_ocean_gke_launch_spec: fix `preemptiblePercentage.onUpdate` ([#191](https://github.com/spotinst/terraform-provider-spotinst/pull/191))

## 1.44.0 (May 24, 2021)

ENHANCEMENTS:
* resource/spotinst_ocean_gke_launch_spec: added support for `strategy` ([#190](https://github.com/spotinst/terraform-provider-spotinst/pull/190))

## 1.43.0 (May 4, 2021)

ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: added support for `initial_noodes`

## 1.42.0 (April 29, 2021)

BUG FIXES:
* resource/spotinst_elastigroup_aws: retry creation for `cant_validate_image` errors

## 1.41.0 (April 28, 2021)

BUG FIXES:
* resource/spotinst_elastigroup_aws: retry creation for `cant_create_group` errors

## 1.40.0 (April 25, 2021)

BUG FIXES:
* resource/spotinst_ocean_ecs: resolved error with update `capacity` fields to 0 

## 1.39.0 (April 20, 2021)

BUG FIXES:

* resource/spotinst_ocean_aks: retry failed import

## 1.38.0 (April 12, 2021)

FEATURES:
* **New Resource:** `spotinst_ocean_aks`
* **New Resource:** `spotinst_ocean_aks_virtual_node_group`

## 1.37.0 (April 5, 2021)

ENHANCEMENTS:
* resource/spotinst_managed_instance_aws: added support for `block_device_mappings`

## 1.36.0 (March 18, 2021)

FEATURES:
* **New Resource:** `spotinst_elastigroup_azure_v3`

BUG FIXES:
* resource/spotinst_ocean_gke_import: resolved error with `desired_capacity` not applying 0 as value
* resource/spotinst_elastigroup_aws: resolved errors with `max_scale_down_percentage`
* resource/spotinst_elastigroup_aws: resolved errors with `key_name`
* resource/spotinst_ocean_strategy: resolved errors with `spot_percentage`

## 1.35.0 (March 1, 2021)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws_launch_configuration: added `cpu_options`.

## 1.34.1 (February 23, 2021)

BUG FIXES:
* resource/spotinst_mrscaler_aws: set additional replica security groups (#156)

## 1.34.0 (February 16, 2021)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws_strategy: added `minimum_instance_lifetime`
* resource/spotinst_ocean_gke_launch_spec: added `restrict_scale_down`
* resource/spotinst_ocean_ecs_launch_spec: added `restrict_scale_down`
* resource/spotinst_ocean_aws_launch_spec: added `restrict_scale_down`

## 1.33.0 (January 18, 2021)

BUG FIXES:
* resource/fields_spotinst_mrscaler_aws_instance_groups_core resolved errors with wrap strategy
* resource/spotinst_elastigroup_aws resolved errors with `batch_num`
* resource/spotinst_elastigroup_aws_block_devices resolved errors with `volume_type`

## 1.32.0 (December 21, 2020)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws_strategy: added `utilize_commitments`
* resource/spotinst_lastigroup_aws_block_devices: added `throughput` under `ebs_block_device`
* resource/spotinst_ocean_aws_launch_spec: added `throughput` under `block_device_mappings`
* resource/spotinst_ocean_aws_strategy: added `utilize_commitments`
* resource/spotinst_ocean_ecs: added `optimize_images`
* resource/spotinst_ocean_ecs_launch_spec: added `throughput` under `block_device_mappings`
* resource/spotinst_ocean_ecs_launch_specification: added `throughput` under `block_device_mappings`
* resource/spotinst_ocean_ecs_strategy: added `utilize_commitments`

BUG FIXES:
* resource/spotinst_ocean_aws_launchspec: resolved errors with `image_id`
* resource/spotinst_elastigroup_aws: resolved errors with `wait_for_roll_percentage`

## 1.31.0 (November 29, 2020)

ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: added `associate_public_ip_address`

## 1.30.0 (November 23, 2020)

ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_configuration: added `use_as_template_only`

## 1.29.0 (November 15, 2020)

ENHANCEMENTS:
* resource/spotinst_ocean_ecs_launch_spec: added `instance_types`
* resource/spotinst_elastigroup_aws_integrations_docker_swarm: added `max_scale_down_percentage`

## 1.28.0 (November 12, 2020)

ENHANCEMENTS:
* resource/spotinst_ocean_ecs_launch_specification: added `block_device_mappings`
* resource/spotinst_ocean_gke_import: added `controller_cluster_id`

DEPRECATIONS:
* resource/spotinst_ocean_gke_import: deprecated `cluster_controller_id`

## 1.27.0 (October 28, 2020)

ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: added `spot_percentage` under `strategy`
* resource/spotinst_elastigroup_aws_launch_configuration: added `metadata_options`.

## 1.26.0 (October 27, 2020)

FEATURES:
* *New Resource*: `elastigroup_aws_suspension`

ENHANCEMENTS:
* resource/spotinst_ocean_ecs_launch_spec: added `block_device_mappings`
* resource/spotinst_elastigroup_aws: added `batch` for `integration_ecs`

BUG FIXES:
* resource/ocean_aws: resolved error with `auto_headroom_percentage` under `auto_scaler`.
* resource/spotinst_elastigroup_aws: resolved error with `target_scaling_policy` under `auto_scaler`.

NOTES:
* documentation: resource/spotinst_mrscaler_aws: fixed usage example and documentation for `retries`.
* documentation: ownership of this repository has been transferred from @terraform-providers to @spotinst

## 1.25.0 (August 26, 2020)

BUG FIXES:
* resource/ocean_aws: resolved error with `spot_percentage` applying 0 automatically.

## 1.24.0 (August 20, 2020)

BUG FIXES:
* resource/ocean_aws: resolved error with `spot_percentage` not applying 0 as value
* resource/ocean_gke_launch_spec_import: `OceanId` and `NodePoolName` are now flagged with ForceNew

## 1.23.0 (August 11, 2020)

ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: added `instance_types`

BUG FIXES:
* resource/spotinst_health_check: fixed backward compatibility in `end_point` and `time_out`

## 1.22.0 (August 05, 2020)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: added support for `CNAME` records in `integration_route53`
* resource/spotinst_managed_instance_aws: added support for `CNAME` records in `integration_route53`

BUG FIXES:
* resource/spotinst_elastigroup_aws: eliminate unmarshalling errors by fixing the type of `maxScaleDownPercentage`
* resource/resource_spotinst_ocean_ecs: eliminate unmarshalling errors by fixing the type of `maxScaleDownPercentage`

## 1.21.0 (August 04, 2020)

ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: added `block_device_mappings`

## 1.20.0 (July 23, 2020)

BUG FIXES:
* resource/ocean_aws_launch_spec: resolved errors with `image_id`
* resource/ocean_aws: resolved error with `	auto_headroom_percentage` field under `autoscaler`

NOTES:
* documentation: resource/spotinst_elastigroup_aws: fixed usage example for `scaling_target_policy`

## 1.19.0 (June 28, 2020)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: added support for `max_capacity_per_scale` under `scaling_target_policy`

## 1.18.0 (June 24, 2020)

BUG FIXES:
* resource/spotinst_elastigroup_aws: resolved errors with `autoscale_constraints` under `integration_nomad`

NOTES:
* documentation: resource/spotinst_elastigroup_aws: added time standard clarification to `start_time` under `scheduled_task`

## 1.17.0 (June 08, 2020)

ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: added `resource_limits`

BUG FIXES:
* resource/spotinst_ocean_aws: resolved errors with `max_scale_down_percentage`
* resource/spotinst_elastigroup_aws: resolved errors with `integration_route53`
* resource/ocean_aws_launch_spec: resolved errors with `name`

## 1.16.0 (May 12, 2020)

ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: added `elastic_ip_pool`

## 1.15.0 (May 06, 2020)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: added `OnFailure` under `roll_config`
* resource/spotinst_ocean_gke_import: added `autoscale`
* resource/spotinst_subscription: update the doc
* resource/spotinst_mrscaler_aws: added `termination_policies`
* resource/spotinst_mrscaler_aws: added `core_unit`
* resource/spotinst_mrscaler_aws: added `task_unit`

BUG FIXES:
* resource/spotinst_ocean_aws: resolved errors with `max_size`
* resource/spotinst_ocean_aws: resolved errors with `autoscale_is_enabled`
* resource/spotinst_elastigroup_aws: resolved `autoscale_cooldown` under `integration_ecs`

## 1.14.3 (April 01, 2020)

ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: added `name`
* resource/spotinst_ocean_aws: added `scheduled_task`
* resource/spotinst_ocean_ecs: added `scheduled_task`
* resource/spotinst_ocean_gke_import: added `scheduled_task`
* resource/spotinst_ocean_aws: added `auto_headroom_percentage`
* resource/spotinst_ocean_aws_launch_spec: added `tags`
* resource/spotinst_ocean_aws: added `grace_period`

BUG FIXES:
* resource/spotinst_mrscaler_aws: resolved errors with `core_min_size`, `core_max_size` , `core_desired_capacity`
* resource/spotinst_elastigroup_aws: resolved errors with `autoscale_scale_down_non_service_tasks`
* resource/spotinst_ocean_aws: resolved errors with `utilize_reserved_instances`
* resource/spotinst_ocean_ecs: resolved errors with `utilize_reserved_instances`

FEATURES:
* *New Resource*: `spotinst_health_check`

NOTES:
* documentation: resource/spotinst_mrscaler_aws: retries are now accurate

## 1.14.2 (January 29, 2020)

BUG FIXES:
* resource/spotinst_elastigroup_aws: resolved errors with roll_config

## 1.14.1 (January 27, 2020)

ENHANCEMENTS:
* resource/spotinst_mrscaler_aws: added `cluster_id` support recreate
* resource/spotinst_managed_instance_aws: update the doc

BUG FIXES:
* resource/spotinst_ocean_aws: resolved errors with `desired_capacity`
* resource/spotinst_ocean_aws: resolved errors with `min_size`
* resource/spotinst_ocean_aws: resolved errors with `max_size`

## 1.14.0 (January 2, 2020)

NOTES:
* This release imports the standalone SDK [hashicorp/terraform-plugin-sdk](https://github.com/hashicorp/terraform-plugin-sdk) v1.4.0.

FEATURES:
* *New Resource*: `spotinst_managed_instance_aws`

ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: added `root_volume_size`
* resource/spotinst_elastigroup_aws_beanstalk: added `scheduled_task`
* resource/spotinst_ocean_aws_launch_spec: added `autoscale_headrooms`
* resource/spotinst_ocean_ecs_launch_spec: added `autoscale_headrooms`
* resource/spotinst_ocean_gke_launch_spec: added `autoscale_headrooms`
* resource/spotinst_ocean_aws_launch_spec: added `subnet_ids`
* resource/spotinst_ocean_aws: added `max_scale_down_percentage`

BUG FIXES:
* resource/spotinst_elastigroup_aws: resolved errors with `integration_ecs`
* resource/spotinst_ocean_aws: resolved errors with `blacklist`
* resource/spotinst_ocean_gke_import: resolved errors with `whitelist`
* resource/spotinst_elastigroup_aws: resolved errors with `ondemand_count`
* resource/spotinst_elastigroup_gcp: resolved errors with `startup_script`
* resource/spotinst_elastigroup_aws: resolved errors with `integration_ecs.roll_config`

## 1.13.5 (October 2, 2019)

ENHANCEMENTS:
* spotinst_ocean_aws_launch_spec: added `security_groups`

NOTES:
* A delay was added during group creation when IAM instance profile linked with the group in order to decrease the retry process of the group creation.

## 1.13.4 (September 11, 2019)

NOTES:
* This release supports Terraform v0.12

FEATURES:
* *New Resource*: `spotinst_ocean_ecs`
* *New Resource*: `spotinst_ocean_ecs_launch_spec`

ENHANCEMENTS:
* spotinst_ocean_gke: added `draining_timeout`
* spotinst_ocean_aws: added `draining_timeout`

FEATURES:
* *New Resource*: `spotinst_ocean_gke_import`
* *New Resource*: `spotinst_ocean_gke_launch_spec`
* *New Resource*: `spotinst_ocean_gke_launch_spec_import`

ENHANCEMENTS:
* spotinst_ocean_gke: moved `backend_services` hierarchy from `launchSpecification` to `compute`

BUG FIXES:
* resource/spotinst_elastigroup_aws: expand `availability_zones` fail to set proper arguments
* resource/spotinst_ocean_aws: fixed wrong headroom GPU per unit and num of units for Ocean AWS AutoScaler

## 1.13.3 (May 13, 2019)

ENHANCEMENTS:
* resource_elastigroup_gcp: added `scheduled_task`
* resource_elastigroup_aws: added `predictive_mode`

## 1.13.2 (May 01, 2019)

NOTES:
* Adding the additonal protocol version (5) for Terraform 0.12 to this release.

## 1.13.1 (May 01, 2019)

ENHANCEMENTS:
* resource/spotinst_ocean_aws: added `update_policy` for managing rolling deployments

## 1.13.0 (April 26, 2019)

ENHANCEMENTS:
* resource/spotinst_ocean_aws_launch_spec: added `iam_instance_profile`
* resource/spotinst_ocean_aws: added `ebs_optimized` and `monitoring`

## 1.12.0 (April 23, 2019)

FEATURES:
* *New Resource*: `spotinst_ocean_gke`

ENHANCEMENTS:
* resource/spotinst_elastigroup_azure: added `custom_data`

NOTES:
* documentation: resource names are now accurate

BUG FIXES:
* resource/spotinst_elastigroup_aws: resolved errors with `max_scale_down_pct`
* resource/spotinst_elastigroup_azure: `custom_data`

## 1.11.0 (April 16, 2019)

FEATURES:
* *New Resource*: `spotinst_ocean_aws_launch_spec`

ENHANCEMENTS:
* resource/elastigroup_azure: added `managed_service_identities`
* resource/elastigroup_gcp: added `shutdown_script`
* resource/elastigroup_gcp: added healthcheck fields
* resource/mrscaler: added `instance_weights`
* resource/ocean_aws: added `root_volume_size` to launch configuration

NOTES:
* migrated to go modules
* added missing docs, example for multai_listener
* resource/multai_listener: added missing tls_config docs
* resource/elastigroup_gcp: fixed typos, examples in docs

## 1.10.0 (April 03, 2019)

ENHANCEMENTS:
* resource/spotinst_elastigroup_gke: Now supports all gcp fields. Added special handling due to parameter import, see notes.
* resource/spotinst_elastigroup_aws: added `max_scale_down_percentage` to `integration_ecs`
* resource/spotinst_elastigroup_aws: `autoscale_scale_down_non_service_tasks` `to integration_ecs`
* resource/spotinst_elastigroup_aws: added `scaling_strategy`

BUG FIXES:
* resource/spotinst_elastigroup_aws: fixes handling of base64-encoded strings
* resource/spotinst_elastigroup_azure: fixes handling of base64-encoded strings
* resource/spotinst_elastigroup_gcp: fixes handling of base64-encoded strings

NOTES:
* added sweepers for acceptance tests. These can be run using the `-sweep` flag, and will destroy any resource with a name beginning with `test-acc-`
* resource/spotinst_elastigroup_gke: Many fields have a diff suppress applied due to this resource's nature (most everything is imported). We will probably support importing and managing GKE clusters using Terraform Modules in the future.

## 1.9.0 (March 27, 2019)

ENHANCEMENTS:
* resource/spotinst_elastigroup_azure: added `additional_ip_configs` to `network`
* resource/spotinst_elastigroup_azure: added kubernetes and Multai to `integrations`
* resource/spotinst_elastigroup_azure: added `scaling policies`

BUG FIXES:
* changed the order that credentials are set. See notes.
* resource/spotinst_elastigroup_azure: `dimensions` changed to properly set `name` and `value` parameters
* resource/spotinst_elastigroup_gcp: `dimensions` changed to properly set `name` and `value` parameters
* resource/spotinst_elastigroup_aws: rolling with `wait_for_roll_percentage` no longer times out after 5 minutes
* resource/spotinst_elastigroup_aws: removed duplicated `wait_for_roll_percentage` and `wait_for_roll_timeout`
* resource/spotinst_mrscaler_aws: `visible_to_all_users` changed to deprecated. Values will not be sent in API calls.

NOTES:
* credentials are now given the following precedence: defined in the template, defined using environment variables, defined in ~/.spotinst/credentials
* spotinst_mrscaler_aws_test: added a delay due to counter AWS rate limiting

## 1.8.0 (February 28, 2019)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: added optional `spotinst_acct_id` to Route53 integration
* resource/spotinst_elastigroup_azure: added `update_policy` to control blue/green deployment options
* resource/spotinst_elastigroup_gcp: added DockerSwarm integration.
* resource/spotinst_elastigroup_gcp: added `location_type` and `scheme` to `backend_services`

BUG FIXES:
* resource/spotinst_elastigroup_aws: `should_roll` now retries on `CANT_ROLL_CAPACITY_BELOW_MINIMUM` error
* resource/spotinst_ocean_aws: `spot_percentage` no longer defaults to `0` when undefined
* resource/spotinst_ocean_aws: `fallback_to_od` now defaults to `true` when undefined
* resource/spotinst_elastigroup_aws: `dimensions` changed to properly set `name` and `value` parameters

## 1.7.0 (February 19, 2019)

FEATURES:
* *New Resource*: `spotinst_mrscaler_aws`
* *New Resource*: `spotinst_multai_balancer`
* *New Resource*: `spotinst_multai_deployment`
* *New Resource*: `spotinst_multai_listener`
* *New Resource*: `spotinst_multai_routing_rule`
* *New Resource*: `spotinst_multai_target`
* *New Resource*: `spotinst_multai_target_set`

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: `network_interface.description` is now Optional (was Required)
* resource/spotinst_elastigroup_aws: `group.description` no longer sends an empty string when undefined
* resource/spotinst_ocean_aws: `headroom` parameters can now be set to 0
* resource/spotinst_ocean_aws: Added `load_balancers` and `load_balancer` objects to ocean parameters
* resource/spotinst_ocean_aws: Added `associate_public_ip_address` to ocean parameters
* resource/spotinst_elastigroup_aws: Added `deployment_preferences` and `managed_actions` to beanstalk integration
* resource/spotinst_elastigroup_aws_beanstalk: Added `deployment_preferences` and `managed_actions` parameters
* added version to user-agent header.

## 1.6.1 (January 31, 2019)

NOTES:
* resource/spotinst_elastigroup_aws: Added `wait_for_roll_timeout` and `wait_for_roll_percentage` to `roll_config` in `update_policy`. Setting both of these fields enables users to wait for a minimum percent of their blue/green deployment to be completed before allowing the plan to continue execution.

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: Added `wait_for_roll_timeout` and `wait_for_roll_percentage` to `roll_config` in `update_policy`.
* resource/spotinst_elastigroup_azure: Added `scheduled_task`

BUG FIXES:
* resource/spotinst_ocean_aws: `min_size`, `max_size`, and `desired_capacity` now defaults to correct values when undefined
* resource/spotinst_elastigroup_aws: configuring `wait_for_capacity` when updating crashed under certain conditions. This has been resolved.

NOTES:
* resource/spotinst_elastigroup_azure: Added a new spotinst_elastigroup_azure resource for creating Spotinst elastigroups using Microsoft Azure
* resource/spotinst_elastigroup_gcp: Added a new spotinst_elastigroup_gcp resource for creating Spotinst elastigroups using Google Cloud
* resource/spotinst_elastigroup_gke: Added a new spotinst_elastigroup_gke resource for creating Spotinst elastigroups using Google Kubernetes Engine
* resource/spotinst_ocean_aws: Added a new spotinst_ocean_aws resource for creating Spotinst Ocean clusters on AWS

FEATURES:
* *New Resource*: `spotinst_elastigroup_azure`
* *New Resource*: `spotinst_elastigroup_gcp`
* *New Resource*: `spotinst_elastigroup_gke`
* *New Resource*: `spotinst_ocean_aws`


## 1.5.0 (December 28, 2018)

NOTES:
* resource/spotinst_elastigroup_aws_beanstalk: Added a new `elastigroup_aws_beanstalk` resource for creating Spotinst elastigroups that are managed by an existing AWS Elastic Beanstalk

FEATURES:
* *New Resource*: `spotinst_elastigroup_aws_beanstalk`
* *New Feature*: spotinst provider version added to the User-Agent header

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws_beanstalk: Added a the ability to transition in and out of maintenance modes by setting `maintenance` mode to `START` or `END`
* resource/spotinst_elastigroup_aws: Added the ability to wait for a minimum number of healthy instances for a certain period of time
* resource/spotinst_elastigroup_aws: Added ability to maintain scaling policy configuration when disabled
* resource/spotinst_elastigroup_aws: Scheduled tasks now support `adjustment` field
* resource/spotinst_elastigroup_aws: Rancher integration now supports `version` field
* resource/spotinst_elastigroup_aws: Use new `wait_for_capacity` field to indicate the minimum number of healthy instances required before continuing plan execution
* resource/spotinst_elastigroup_aws: Use new `wait_for_capacity_timeout` to indicate how long to wait for minimum number of instances to become healthy
* resource/spotinst_elastigroup_aws: Use new `is_enabled` field in scaling policies to indicate if that policy is active
* resource/spotinst_elastigroup_aws: Use new `adjustment` field in `scheduled_tasks` to indicate the number of instances to add or remove when scaling

BUG FIXES:
* resource/spotinst_elastigroup_aws: `user_data` and `shutdown_script` no longer updates to empty string SHA
* resource/spotinst_elastigroup_aws: Fixed an issue of `tags`, `instance_types_spot` and `instance_types_preferred_spot` not being imported properly
* resource/spotinst_elastigroup_aws: Fixed an issue where `associate_public_ip` incorrectly defaulting to `false` when undefined

## 1.4.0 (September 13, 2018)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: Shutdown script is now supported under `shutdown_script`
* resource/spotinst_elastigroup_aws: ECS integration support for `autoscale_is_autoconfig`
* resource/spotinst_elastigroup_aws: Docker Swarm integration as `integration_docker_swarm`

## 1.3.0 (August 13, 2018)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: Added a new Route53 integration as `integration_route53`
* resource/spotinst_elastigroup_aws: Added support for preferred spot instances as `instance_types_preferred_spot`

## 1.2.0 (July 26, 2018)

ENHANCEMENTS:
* resource/spotinst_elastigroup_aws: Added `kms_key_id` support for `ebs_block_device`
* resource/spotinst_elastigroup_aws: Added `autoscale_attributes` support for `integration_ecs`
* resource/spotinst_elastigroup_aws: Added `autoscale_labels` support for `integration_kubernetes`
* resource/spotinst_elastigroup_aws: Added `autoscale_constraints` support for `integration_nomad`

## 1.1.1 (July 09, 2018)

BUG FIXES:
* resource/spotinst_elastigroup_aws: `scheduled_task` & `network_interface` now properly address fields not specified on TF file as nil instead of their default values

## 1.1.0 (July 02, 2018)

NOTES

* resource/spotinst_subscription: Added a new subscription resource for creating Spotinst subscriptions that gets triggered by an elastigroup event type

FEATURES:

* **New Resource:** `spotinst_subscription`

ENHANCEMENTS:

* resource/spotinst_elastigroup_aws: Added a new Gitlab runner integration

BUG FIXES:

* resource/spotinst_elastigroup_aws: Resource now properly create multiple elastigroups using the count parameter and/or using parallelism via terraform apply

## 1.0.0 (June 21, 2018)

BREAKING CHANGES / NOTES

Introduced a new API schema to support the latest Spotinst API additions while using similar AWS terminology.

* resource/spotinst_group_aws: Resource name changed to `spotinst_elastigroup_aws`
* resource/spotinst_elastigroup_aws: Removed `capacity` and flattened its fields on the resource
* resource/spotinst_elastigroup_aws: Changed all previous `capacity` field names to `max_size`, `min_size`, `desired_capacity`, `capacity_unit`
* resource/spotinst_elastigroup_aws: Removed `launch_specification` and flattened its fields on the resource
* resource/spotinst_elastigroup_aws: Removed `persistence` and flattened its fields on the resource
* resource/spotinst_elastigroup_aws: Removed `strategy` and flattened its fields on the resource
* resource/spotinst_elastigroup_aws: Removed `availability_zone` and currently only `availability_zones` field is supported
* resource/spotinst_elastigroup_aws: Removed `load_balancers` and broke it down to the following fields: `elastic_load_balancers`, `target_group_arns`, `multai_target_sets`
* resource/spotinst_elastigroup_aws: Dropped previous `tags` field and changed `tags_kv` name to `tags` which accepts only key/value objects
* resource/spotinst_elastigroup_aws: Introduced a new object `update_policy` for group roll configuration
* resource/spotinst_elastigroup_aws: Field `should_resume_stateful` is now available under `update_policy`
* resource/spotinst_elastigroup_aws: Changed `availability_vs_cost` name to `orientation`
* resource/spotinst_elastigroup_aws: Changed `risk` name to `spot_percentage`
* resource/spotinst_elastigroup_aws: Deprecated `hot_ebs_volume`
* resource/spotinst_elastigroup_aws: Deprecated `launch_specification.load_balancer_names`
* resource/spotinst_elastigroup_aws: Deprecated `elastic_beanstalk_integration`
* resource/spotinst_elastigroup_aws: Renamed `rancher_integration` to `integration_rancher`
* resource/spotinst_elastigroup_aws: Renamed `ec2_container_service_integration` to `integration_ecs`
* resource/spotinst_elastigroup_aws: Renamed `kubernetes_integration` to `integration_kubernetes`
* resource/spotinst_elastigroup_aws: Renamed `nomad_integration` to `integration_nomad`
* resource/spotinst_elastigroup_aws: Renamed `mesosphere_integration` to `integration_mesosphere`
* resource/spotinst_elastigroup_aws: Renamed `multai_runtime_integration` to `integration_multai_runtime`

FEATURES:

* **New Resource:** `spotinst_elastigroup_aws`

ENHANCEMENTS:

* resource/spotinst_elastigroup_aws: All singleton objects e.g. integrations now support proper logs formatting on any change
* resource/spotinst_elastigroup_aws: Added support for vpc zone identifier under field name `subnet_ids` as a list of subnet identifiers Strings and `region` field that represent the AWS region your group will be created in
* resource/spotinst_elastigroup_aws: Added support for `autoscale_is_auto_config` under `integration_kubernetes`
* resource/spotinst_elastigroup_aws: Added support for maintenance window under field name `revert_to_spot`
* resource/spotinst_elastigroup_aws: Kubernetes integration now contain cluster controller support under `integration_mode` and `cluster_identifier`
* resource/spotinst_elastigroup_aws: Flattened previous objects `capacity`, `launch_specification`, `persistence`, `strategy`

## 0.1.0 (June 21, 2017)

NOTES:

* Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)

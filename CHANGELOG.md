## 1.1.2 (Unreleased)

ENHANCEMENTS:

* resource/spotinst_elastigroup_aws: Added kms_key_id support for ebs_block_device
* resource/spotinst_elastigroup_aws: Added autoscale_attributes support for integration_ecs
* resource/spotinst_elastigroup_aws: Added autoscale_labels support for integration_kubernetes
* resource/spotinst_elastigroup_aws: Added autoscale_constraints support for integration_nomad

## 1.1.1 (July 09, 2018)

BUG FIXES:

* resource/spotinst_elastigroup_aws: scheduled_task & network_interface now properly address fields not specified on TF file as nil instead of their default values

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

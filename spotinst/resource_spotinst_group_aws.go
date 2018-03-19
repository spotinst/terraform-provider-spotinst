package spotinst

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
)

func resourceSpotinstAWSGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpotinstAWSGroupCreate,
		Read:   resourceSpotinstAWSGroupRead,
		Update: resourceSpotinstAWSGroupUpdate,
		Delete: resourceSpotinstAWSGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"target_capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},

			"should_resume_stateful": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},

			"private_ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"capacity": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target": &schema.Schema{
							Type:          schema.TypeInt,
							Optional:      true,
							ConflictsWith: []string{"target_capacity"},
						},

						"minimum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"maximum": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"unit": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
				Set: hashAWSGroupCapacity,
			},

			"strategy": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"risk": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
						},

						"ondemand_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"availability_vs_cost": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						"lifetime_period": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"draining_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						"utilize_reserved_instances": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						"fallback_to_ondemand": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						"spin_up_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
				Set: hashAWSGroupStrategy,
			},

			"scheduled_task": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},

						"task_type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"frequency": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"cron_expression": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"scale_target_capacity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"scale_min_capacity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"scale_max_capacity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"batch_size_percentage": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"grace_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"target_capacity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"min_capacity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"max_capacity": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"product": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"instance_types": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ondemand": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"spot": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},

			"persistence": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"persist_root_device": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},

						"persist_block_devices": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},

						"persist_private_ip": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},

						"block_devices_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
				Set: hashAWSGroupPersistence,
			},

			"signal": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"availability_zone": &schema.Schema{
				Type:          schema.TypeSet,
				Optional:      true,
				ConflictsWith: []string{"availability_zones"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"subnet_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"placement_group_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"availability_zones": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"hot_ebs_volume": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"volume_ids": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},

			"load_balancer": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"arn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"balancer_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"target_set_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"zone_awareness": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},

						"auto_weight": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
				Set:           hashAWSGroupLoadBalancer,
				ConflictsWith: []string{"load_balancers"},
			},

			"load_balancers": &schema.Schema{
				Type:          schema.TypeList,
				Optional:      true,
				Elem:          &schema.Schema{Type: schema.TypeString},
				ConflictsWith: []string{"load_balancer"},
			},

			"launch_specification": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load_balancer_names": &schema.Schema{
							Type:       schema.TypeList,
							Optional:   true,
							Elem:       &schema.Schema{Type: schema.TypeString},
							Deprecated: "Attribute `load_balancer_names` is deprecated. Use `load_balancer` instead.",
						},

						"monitoring": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},

						"ebs_optimized": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						"image_id": &schema.Schema{
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"image_id"},
						},

						"tenancy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"key_pair": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"health_check_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"health_check_grace_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"health_check_unhealthy_duration_before_replacement": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"security_group_ids": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},

						"user_data": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							StateFunc: hexStateFunc,
						},

						"shutdown_script": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							StateFunc: hexStateFunc,
						},

						"iam_role": &schema.Schema{
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "Attribute `iam_role` is deprecated. Use `iam_instance_profile` instead.",
						},

						"iam_instance_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"elastic_ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"tags": &schema.Schema{
				Type:          schema.TypeMap,
				Optional:      true,
				ConflictsWith: []string{"tags_kv"},
			},

			"tags_kv": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"value": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Set: hashKV,
			},

			"instance_type_weights": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},

			"ebs_block_device": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete_on_termination": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						"device_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"encrypted": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						"iops": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"snapshot_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"volume_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"volume_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
				Set: hashAWSGroupEBSBlockDevice,
			},

			"ephemeral_block_device": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"virtual_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"network_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"device_index": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"secondary_private_ip_address_count": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"associate_public_ip_address": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},

						"delete_on_termination": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						"security_group_ids": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},

						"network_interface_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"private_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"subnet_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"scaling_up_policy": upDownScalingPolicySchema(),

			"scaling_down_policy": upDownScalingPolicySchema(),

			"scaling_target_policy": targetScalingPolicySchema(),

			"rancher_integration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"master_host": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"access_key": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"secret_key": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"elastic_beanstalk_integration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"environment_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"ec2_container_service_integration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"autoscale_is_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},

						"autoscale_cooldown": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"autoscale_headroom": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_per_unit": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},

									"memory_per_unit": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},

									"num_of_units": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"autoscale_down": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"evaluation_periods": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"kubernetes_integration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"integration_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"cluster_identifier": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"api_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"token": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"autoscale_is_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},

						"autoscale_cooldown": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"autoscale_headroom": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_per_unit": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},

									"memory_per_unit": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},

									"num_of_units": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"autoscale_down": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"evaluation_periods": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"nomad_integration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"master_host": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"master_port": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"autoscale_is_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},

						"autoscale_cooldown": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},

						"acl_token": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"autoscale_headroom": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_per_unit": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},

									"memory_per_unit": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},

									"num_of_units": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"autoscale_down": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"evaluation_periods": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"autoscale_constraints": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": &schema.Schema{
										Type:      schema.TypeString,
										Required:  true,
										StateFunc: attrStateFunc,
									},

									"value": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
							Set: hashKV,
						},
					},
				},
			},

			"mesosphere_integration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api_server": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"multai_integration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployment_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"codedeploy_integration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cleanup_on_failure": &schema.Schema{
							Type:     schema.TypeBool,
							Required: true,
						},

						"terminate_instance_on_failure": &schema.Schema{
							Type:     schema.TypeBool,
							Required: true,
						},

						"deployment_groups": &schema.Schema{
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"application_name": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},

									"deployment_group_name": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},

			"roll_config": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"should_roll": &schema.Schema{
							Type:     schema.TypeBool,
							Required: true,
						},

						"batch_size_percentage": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},

						"grace_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Default:  -1,
						},

						"health_check_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func baseScalingPolicySchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"policy_name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},

				"metric_name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},

				"namespace": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},

				"source": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},

				"statistic": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},

				"unit": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},

				"cooldown": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Computed: true,
				},

				"dimensions": &schema.Schema{
					Type:     schema.TypeMap,
					Optional: true,
				},
			},
		},
	}
}

func upDownScalingPolicySchema() *schema.Schema {
	o := baseScalingPolicySchema()
	s := o.Elem.(*schema.Resource).Schema

	s["threshold"] = &schema.Schema{
		Type:     schema.TypeFloat,
		Required: true,
	}

	s["adjustment"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}

	s["adjustment_expression"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	s["min_target_capacity"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}

	s["max_target_capacity"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}

	s["operator"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	}

	s["evaluation_periods"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	}

	s["period"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	}

	s["minimum"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}

	s["maximum"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}

	s["target"] = &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
	}

	s["action_type"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}

	return o
}

func targetScalingPolicySchema() *schema.Schema {
	o := baseScalingPolicySchema()
	s := o.Elem.(*schema.Resource).Schema

	s["target"] = &schema.Schema{
		Type:     schema.TypeFloat,
		Required: true,
	}

	return o
}

//region CRUD methods

func resourceSpotinstAWSGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	newAWSGroup, err := buildAWSGroupOpts(d, meta)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Group create configuration: %s", stringutil.Stringify(newAWSGroup))
	input := &aws.CreateGroupInput{Group: newAWSGroup}
	resp, err := client.elastigroup.CloudProviderAWS().Create(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to create group: %s", err)
	}
	d.SetId(spotinst.StringValue(resp.Group.ID))
	log.Printf("[INFO] AWSGroup created successfully: %s", d.Id())
	return resourceSpotinstAWSGroupRead(d, meta)
}

// ErrCodeGroupNotFound for service response error code "GROUP_DOESNT_EXIST".
const ErrCodeGroupNotFound = "GROUP_DOESNT_EXIST"

func resourceSpotinstAWSGroupRead(d *schema.ResourceData, meta interface{}) error {
	input := &aws.ReadGroupInput{GroupID: spotinst.String(d.Id())}
	resp, err := meta.(*Client).elastigroup.CloudProviderAWS().Read(context.Background(), input)
	if err != nil {
		// If the group was not found, return nil so that we can show
		// that the group is gone.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGroupNotFound {
					d.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return fmt.Errorf("failed to read group: %s", err)
	}

	// If nothing was found, then return no state.
	if resp.Group == nil {
		d.SetId("")
		return nil
	}

	g := resp.Group
	d.Set("name", g.Name)
	d.Set("description", g.Description)
	d.Set("product", g.Compute.Product)
	d.Set("elastic_ips", g.Compute.ElasticIPs)
	d.Set("private_ips", g.Compute.PrivateIPs)

	// Set capacity.

	if g.Capacity != nil {
		targetCapacitySetInCapacity := true
		if v, ok := d.GetOk("target_capacity"); ok && v != nil {
			targetCapacitySetInCapacity = false
		}
		if err := d.Set("capacity", flattenAWSGroupCapacity(g.Capacity, targetCapacitySetInCapacity)); err != nil {
			return fmt.Errorf("failed to set capacity onfiguration: %#v", err)
		}
	}

	// Set Target Capacity.
	if g.Capacity.Target != nil {
		if v, ok := d.GetOk("target_capacity"); ok && v != nil {
			d.Set("target_capacity", g.Capacity.Target)
		}
	}

	if g.Strategy != nil {
		if err := d.Set("strategy", flattenAWSGroupStrategy(g.Strategy)); err != nil {
			return fmt.Errorf("failed to set strategy configuration: %#v", err)
		}

		// Set signals.
		if g.Strategy.Signals != nil {
			if err := d.Set("signal", flattenAWSGroupSignals(g.Strategy.Signals)); err != nil {
				return fmt.Errorf("failed to set signals configuration: %#v", err)
			}
		} else {
			d.Set("signal", []*aws.ScalingPolicy{})
		}

		if g.Strategy.Persistence != nil {
			if err := d.Set("persistence", flattenAWSGroupPersistence(g.Strategy.Persistence)); err != nil {
				return fmt.Errorf("failed to set persistence configuration: %#v", err)
			}
		}
	}

	if g.Scheduling != nil {
		// Set scheduled tasks.
		if g.Scheduling.Tasks != nil {
			if err := d.Set("scheduled_task", flattenAWSGroupScheduledTasks(g.Scheduling.Tasks)); err != nil {
				return fmt.Errorf("failed to set scheduled tasks configuration: %#v", err)
			}
		} else {
			d.Set("scheduled_task", []*aws.Task{})
		}

	}

	// Set launch specification.
	if g.Compute.LaunchSpecification != nil {
		// Check if image ID is set in launch spec
		imageIDSetInLaunchSpec := true
		if v, ok := d.GetOk("image_id"); ok && v != nil {
			imageIDSetInLaunchSpec = false
		}
		if err := d.Set("launch_specification", flattenAWSGroupLaunchSpecification(g.Compute.LaunchSpecification, imageIDSetInLaunchSpec)); err != nil {
			return fmt.Errorf("failed to set launch specification configuration: %#v", err)
		}
	}

	// Set image ID.
	if g.Compute.LaunchSpecification.ImageID != nil {
		if v, ok := d.GetOk("image_id"); ok && v != nil {
			d.Set("image_id", g.Compute.LaunchSpecification.ImageID)
		}
	}

	// Set load balancers.
	if g.Compute.LaunchSpecification.LoadBalancersConfig != nil {
		if v, ok := d.GetOk("load_balancer"); ok && v != nil {
			if err := d.Set("load_balancer", flattenAWSGroupLoadBalancers(g.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers)); err != nil {
				return fmt.Errorf("failed to set load balancers configuration: %#v", err)
			}
		}
	} else {
		d.Set("load_balancer", []*aws.LoadBalancer{})
	}

	// Set EBS volume pool.
	if g.Compute.EBSVolumePool != nil {
		if err := d.Set("hot_ebs_volume", flattenAWSGroupEBSVolumePool(g.Compute.EBSVolumePool)); err != nil {
			return fmt.Errorf("failed to set EBS volume pool configuration: %#v", err)
		}
	} else {
		d.Set("hot_ebs_volume", []*aws.EBSVolume{})
	}

	// Set network interfaces.
	if g.Compute.LaunchSpecification.NetworkInterfaces != nil {
		if err := d.Set("network_interface", flattenAWSGroupNetworkInterfaces(g.Compute.LaunchSpecification.NetworkInterfaces)); err != nil {
			return fmt.Errorf("failed to set network interfaces configuration: %#v", err)
		}
	} else {
		d.Set("network_interface", []*aws.NetworkInterface{})
	}

	// Set block devices.
	if g.Compute.LaunchSpecification.BlockDeviceMappings != nil {
		if err := d.Set("ebs_block_device", flattenAWSGroupEBSBlockDevices(g.Compute.LaunchSpecification.BlockDeviceMappings)); err != nil {
			return fmt.Errorf("failed to set EBS block devices configuration: %#v", err)
		}
		if err := d.Set("ephemeral_block_device", flattenAWSGroupEphemeralBlockDevices(g.Compute.LaunchSpecification.BlockDeviceMappings)); err != nil {
			return fmt.Errorf("failed to set Ephemeral block devices configuration: %#v", err)
		}
	} else {
		d.Set("ebs_block_device", []*aws.BlockDeviceMapping{})
		d.Set("ephemeral_block_device", []*aws.BlockDeviceMapping{})
	}

	if g.Integration != nil {
		// Set Rancher integration.
		if g.Integration.Rancher != nil {
			if err := d.Set("rancher_integration", flattenAWSGroupRancherIntegration(g.Integration.Rancher)); err != nil {
				return fmt.Errorf("failed to set Rancher configuration: %#v", err)
			}
		} else {
			d.Set("rancher_integration", []*aws.RancherIntegration{})
		}

		// Set Elastic Beanstalk integration.
		if g.Integration.ElasticBeanstalk != nil {
			if err := d.Set("elastic_beanstalk_integration", flattenAWSGroupElasticBeanstalkIntegration(g.Integration.ElasticBeanstalk)); err != nil {
				return fmt.Errorf("failed to set Elastic Beanstalk configuration: %#v", err)
			}
		} else {
			d.Set("elastic_beanstalk_integration", []*aws.ElasticBeanstalkIntegration{})
		}

		// Set Mesosphere integration.
		if g.Integration.Mesosphere != nil {
			if err := d.Set("mesosphere_integration", flattenAWSGroupMesosphereIntegration(g.Integration.Mesosphere)); err != nil {
				return fmt.Errorf("failed to set Mesosphere configuration: %#v", err)
			}
		} else {
			d.Set("mesosphere_integration", []*aws.MesosphereIntegration{})
		}

		// Set Multai integration.
		if g.Integration.Multai != nil {
			if err := d.Set("multai_integration", flattenAWSGroupMultaiIntegration(g.Integration.Multai)); err != nil {
				return fmt.Errorf("failed to set Multai configuration: %#v", err)
			}
		} else {
			d.Set("multai_integration", []*aws.MultaiIntegration{})
		}

		// Set CodeDeploy integration.
		if g.Integration.CodeDeploy != nil {
			if err := d.Set("codedeploy_integration", flattenAWSGroupCodeDeployIntegration(g.Integration.CodeDeploy)); err != nil {
				return fmt.Errorf("failed to set CodeDeploy configuration: %#v", err)
			}
		} else {
			d.Set("codedeploy_integration", []*aws.CodeDeployIntegration{})
		}
	}

	return nil
}

func resourceSpotinstAWSGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	group := &aws.Group{}
	group.SetId(spotinst.String(d.Id()))
	update := false
	nullify := true

	if d.HasChange("name") {
		group.SetName(spotinst.String(d.Get("name").(string)))
		update = true
	}

	if d.HasChange("description") {
		group.SetDescription(spotinst.String(d.Get("description").(string)))
		update = true
	}

	if d.HasChange("capacity") {
		if v, ok := d.GetOk("capacity"); ok {
			_, exists := d.GetOkExists("target_capacity")
			if capacity, err := expandAWSGroupCapacity(v, nullify, !exists, true); err != nil {
				return err
			} else {
				group.SetCapacity(capacity)
				update = true
			}
		}
	}

	if d.HasChange("private_ips") {
		if v, ok := d.GetOk("private_ips"); ok {
			if privateIPs, err := expandAWSGroupPrivateIPs(v); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				group.Compute.SetPrivateIPs(privateIPs)
				update = true
			}
		} else {
			if group.Compute == nil {
				group.SetCompute(&aws.Compute{})
			}
			group.Compute.SetPrivateIPs(nil)
			update = true
		}
	}

	if d.HasChange("target_capacity") {
		if v, ok := d.GetOk("target_capacity"); ok {
			if group.Capacity == nil {
				newCapacity := &aws.Capacity{}
				group.SetCapacity(newCapacity)
			}
			group.Capacity.SetTarget(spotinst.Int(v.(int)))
			update = true
		}
	}

	if d.HasChange("strategy") {
		if v, ok := d.GetOk("strategy"); ok {
			if strategy, err := expandAWSGroupStrategy(v, nullify); err != nil {
				return err
			} else {
				group.SetStrategy(strategy)
				if v, ok := d.GetOk("signal"); ok {
					if signals, err := expandAWSGroupSignals(v, nullify); err != nil {
						return err
					} else {
						group.Strategy.SetSignals(signals)
					}
				}
				update = true
			}
		}
	}

	if d.HasChange("launch_specification") {
		if v, ok := d.GetOk("launch_specification"); ok {
			lc, err := expandAWSGroupLaunchSpecification(v, true)
			if err != nil {
				return err
			}
			if group.Compute == nil {
				group.SetCompute(&aws.Compute{})
			}
			group.Compute.SetLaunchSpecification(lc)
			update = true
		}
	}

	if d.HasChange("image_id") {
		if d.Get("image_id") != nil && d.Get("image_id") != "" {
			if group.Compute == nil {
				group.SetCompute(&aws.Compute{})
			}
			if group.Compute.LaunchSpecification == nil {
				group.Compute.SetLaunchSpecification(&aws.LaunchSpecification{})
			}
			group.Compute.LaunchSpecification.SetImageId(spotinst.String(d.Get("image_id").(string)))
			update = true
		}
	}

	if d.HasChange("load_balancer") {
		if v, ok := d.GetOk("load_balancer"); ok {
			if lbs, err := expandAWSGroupLoadBalancer(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&aws.LaunchSpecification{})
				}
				if group.Compute.LaunchSpecification.LoadBalancersConfig == nil {
					group.Compute.LaunchSpecification.SetLoadBalancersConfig(&aws.LoadBalancersConfig{})
					group.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(lbs)
					update = true
				}
			}
		} else {
			if group.Compute == nil {
				group.SetCompute(&aws.Compute{})
			}
			if group.Compute.LaunchSpecification == nil {
				group.Compute.SetLaunchSpecification(&aws.LaunchSpecification{})
			}
			if group.Compute.LaunchSpecification.LoadBalancersConfig == nil {
				group.Compute.LaunchSpecification.SetLoadBalancersConfig(&aws.LoadBalancersConfig{})
			}
			group.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(nil)
			update = true
		}
	}

	if d.HasChange("load_balancers") {
		if v, ok := d.GetOk("load_balancers"); ok {
			if lbs, err := expandAWSGroupLoadBalancers(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&aws.LaunchSpecification{})
				}
				if group.Compute.LaunchSpecification.LoadBalancersConfig == nil {
					group.Compute.LaunchSpecification.SetLoadBalancersConfig(&aws.LoadBalancersConfig{})
					group.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(lbs)
					update = true
				}
			}
		} else {
			if group.Compute == nil {
				group.SetCompute(&aws.Compute{})
			}
			if group.Compute.LaunchSpecification == nil {
				group.Compute.SetLaunchSpecification(&aws.LaunchSpecification{})
			}
			if group.Compute.LaunchSpecification.LoadBalancersConfig == nil {
				group.Compute.LaunchSpecification.SetLoadBalancersConfig(&aws.LoadBalancersConfig{})
			}
			group.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(nil)
			update = true
		}
	}

	var blockDevicesExpanded bool

	if d.HasChange("ebs_block_device") {
		if v, ok := d.GetOk("ebs_block_device"); ok {
			if devices, err := expandAWSGroupEBSBlockDevices(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&aws.LaunchSpecification{})
				}
				if len(group.Compute.LaunchSpecification.BlockDeviceMappings) > 0 {
					group.Compute.LaunchSpecification.SetBlockDeviceMappings(append(group.Compute.LaunchSpecification.BlockDeviceMappings, devices...))
				} else {
					if v, ok := d.GetOk("ephemeral_block_device"); ok {
						if ephemeral, err := expandAWSGroupEphemeralBlockDevices(v, nullify); err != nil {
							return err
						} else {
							devices = append(devices, ephemeral...)
							blockDevicesExpanded = true
						}
					}
					group.Compute.LaunchSpecification.SetBlockDeviceMappings(devices)
				}
				update = true
			}
		} else {
			if group.Compute == nil {
				group.SetCompute(&aws.Compute{})
			}
			if group.Compute.LaunchSpecification == nil {
				group.Compute.SetLaunchSpecification(&aws.LaunchSpecification{})
			}
			group.Compute.LaunchSpecification.SetBlockDeviceMappings(nil)
			update = true
		}
	}

	if d.HasChange("ephemeral_block_device") && !blockDevicesExpanded {
		if v, ok := d.GetOk("ephemeral_block_device"); ok {
			if devices, err := expandAWSGroupEphemeralBlockDevices(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&aws.LaunchSpecification{})
				}
				if len(group.Compute.LaunchSpecification.BlockDeviceMappings) > 0 {
					group.Compute.LaunchSpecification.SetBlockDeviceMappings(append(group.Compute.LaunchSpecification.BlockDeviceMappings, devices...))
				} else {
					if v, ok := d.GetOk("ebs_block_device"); ok {
						if ebs, err := expandAWSGroupEBSBlockDevices(v, nullify); err != nil {
							return err
						} else {
							devices = append(devices, ebs...)
						}
					}
					group.Compute.LaunchSpecification.SetBlockDeviceMappings(devices)
				}
				update = true
			}
		}
	}

	if d.HasChange("network_interface") {
		if v, ok := d.GetOk("network_interface"); ok {
			if interfaces, err := expandAWSGroupNetworkInterfaces(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&aws.LaunchSpecification{})
				}
				group.Compute.LaunchSpecification.SetNetworkInterfaces(interfaces)
				update = true
			}
		} else {
			if group.Compute == nil {
				group.SetCompute(&aws.Compute{})
			}
			if group.Compute.LaunchSpecification == nil {
				group.Compute.SetLaunchSpecification(&aws.LaunchSpecification{})
			}
			group.Compute.LaunchSpecification.SetNetworkInterfaces(nil)
			update = true
		}
	}

	if d.HasChange("availability_zone") {
		if v, ok := d.GetOk("availability_zone"); ok {
			if zones, err := expandAWSGroupAvailabilityZones(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				group.Compute.SetAvailabilityZones(zones)
				update = true
			}
		}
	}

	if d.HasChange("availability_zones") {
		if v, ok := d.GetOk("availability_zones"); ok {
			if zones, err := expandAWSGroupAvailabilityZonesSlice(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				group.Compute.SetAvailabilityZones(zones)
				update = true
			}
		}
	}

	if d.HasChange("hot_ebs_volume") {
		if v, ok := d.GetOk("hot_ebs_volume"); ok {
			if ebsVolumePool, err := expandAWSGroupEBSVolumePool(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				group.Compute.SetEBSVolumePool(ebsVolumePool)
				update = true
			}
		} else {
			group.Compute.SetEBSVolumePool(nil)
			update = true
		}
	}

	if d.HasChange("signal") {
		if v, ok := d.GetOk("signal"); ok {
			if signals, err := expandAWSGroupSignals(v, nullify); err != nil {
				return err
			} else {
				if group.Strategy == nil {
					group.SetStrategy(&aws.Strategy{})
				}
				group.Strategy.SetSignals(signals)
				update = true
			}
		} else {
			if group.Strategy == nil {
				group.SetStrategy(&aws.Strategy{})
			}
			group.Strategy.SetSignals(nil)
			update = true
		}
	}

	if d.HasChange("instance_types") {
		if v, ok := d.GetOk("instance_types"); ok {
			if types, err := expandAWSGroupInstanceTypes(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				group.Compute.SetInstanceTypes(types)
				update = true
			}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			if tags, err := expandAWSGroupTags(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&aws.LaunchSpecification{})
				}
				group.Compute.LaunchSpecification.SetTags(tags)
				update = true
			}
		} else {
			if _, ok := d.GetOk("tags_kv"); !ok {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&aws.LaunchSpecification{})
				}
				group.Compute.LaunchSpecification.SetTags(nil)
				update = true
			}
		}
	}

	if d.HasChange("tags_kv") {
		if v, ok := d.GetOk("tags_kv"); ok {
			if tags, err := expandAWSGroupTagsKV(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&aws.LaunchSpecification{})
				}
				group.Compute.LaunchSpecification.SetTags(tags)
				update = true
			}
		} else {
			if group.Compute == nil {
				group.SetCompute(&aws.Compute{})
			}
			if group.Compute.LaunchSpecification == nil {
				group.Compute.SetLaunchSpecification(&aws.LaunchSpecification{})
			}
			group.Compute.LaunchSpecification.SetTags(nil)
			update = true
		}
	}

	if d.HasChange("instance_type_weights") {
		if v, ok := d.GetOk("instance_type_weights"); ok {
			if weights, err := expandAWSGroupInstanceTypeWeights(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				if group.Compute.InstanceTypes == nil {
					group.Compute.SetInstanceTypes(&aws.InstanceTypes{})
				}
				group.Compute.InstanceTypes.SetWeights(weights)
			}
		} else {
			if group.Compute == nil {
				group.SetCompute(&aws.Compute{})
			}
			if group.Compute.InstanceTypes == nil {
				group.Compute.SetInstanceTypes(&aws.InstanceTypes{})
			}
			group.Compute.InstanceTypes.SetWeights(nil)
		}
		update = true
	}

	if d.HasChange("elastic_ips") {
		if v, ok := d.GetOk("elastic_ips"); ok {
			if eips, err := expandAWSGroupElasticIPs(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&aws.Compute{})
				}
				group.Compute.SetElasticIPs(eips)
				update = true
			}
		} else {
			if group.Compute == nil {
				group.SetCompute(&aws.Compute{})
			}
			group.Compute.SetElasticIPs(nil)
			update = true
		}
	}

	if d.HasChange("scheduled_task") {
		if v, ok := d.GetOk("scheduled_task"); ok {
			if tasks, err := expandAWSGroupScheduledTasks(v, nullify); err != nil {
				return err
			} else {
				if group.Scheduling == nil {
					group.SetScheduling(&aws.Scheduling{})
				}
				group.Scheduling.SetTasks(tasks)
				update = true
			}
		} else {
			if group.Scheduling == nil {
				group.SetScheduling(&aws.Scheduling{})
			}
			group.Scheduling.SetTasks(nil)
			update = true
		}
	}

	if d.HasChange("persistence") {
		if v, ok := d.GetOk("persistence"); ok {
			if persistence, err := expandAWSGroupPersistence(v, nullify); err != nil {
				return err
			} else {
				if group.Strategy == nil {
					group.SetStrategy(&aws.Strategy{})
				}
				if group.Strategy.Persistence == nil {
					group.Strategy.SetPersistence(&aws.Persistence{})
				}
				group.Strategy.SetPersistence(persistence)
				update = true
			}
		} else {
			if group.Strategy == nil {
				group.SetStrategy(&aws.Strategy{})
			}
			group.Strategy.SetPersistence(nil)
			update = true
		}
	}

	if d.HasChange("scaling_up_policy") {
		if v, ok := d.GetOk("scaling_up_policy"); ok {
			if policies, err := expandAWSGroupScalingPolicies(v, nullify); err != nil {
				return err
			} else {
				if group.Scaling == nil {
					group.SetScaling(&aws.Scaling{})
				}
				group.Scaling.SetUp(policies)
				update = true
			}
		} else {
			if group.Scaling == nil {
				group.SetScaling(&aws.Scaling{})
			}
			group.Scaling.SetUp(nil)
			update = true
		}
	}

	if d.HasChange("scaling_down_policy") {
		if v, ok := d.GetOk("scaling_down_policy"); ok {
			if policies, err := expandAWSGroupScalingPolicies(v, nullify); err != nil {
				return err
			} else {
				if group.Scaling == nil {
					group.SetScaling(&aws.Scaling{})
				}
				group.Scaling.SetDown(policies)
				update = true
			}
		} else {
			if group.Scaling == nil {
				group.SetScaling(&aws.Scaling{})
			}
			group.Scaling.SetDown(nil)
			update = true
		}
	}

	if d.HasChange("scaling_target_policy") {
		if v, ok := d.GetOk("scaling_target_policy"); ok {
			if policies, err := expandAWSGroupScalingPolicies(v, nullify); err != nil {
				return err
			} else {
				if group.Scaling == nil {
					group.SetScaling(&aws.Scaling{})
				}
				group.Scaling.SetTarget(policies)
				update = true
			}
		} else {
			if group.Scaling == nil {
				group.SetScaling(&aws.Scaling{})
			}
			group.Scaling.SetTarget(nil)
			update = true
		}
	}

	if d.HasChange("rancher_integration") {
		if v, ok := d.GetOk("rancher_integration"); ok {
			if integration, err := expandAWSGroupRancherIntegration(v, nullify); err != nil {
				return err
			} else {
				if group.Integration == nil {
					group.SetIntegration(&aws.Integration{})
				}
				group.Integration.SetRancher(integration)
				update = true
			}
		} else {
			if group.Integration == nil {
				group.SetIntegration(&aws.Integration{})
			}
			group.Integration.SetRancher(nil)
			update = true
		}
	}

	if d.HasChange("elastic_beanstalk_integration") {
		if v, ok := d.GetOk("elastic_beanstalk_integration"); ok {
			if integration, err := expandAWSGroupElasticBeanstalkIntegration(v, nullify); err != nil {
				return err
			} else {
				if group.Integration == nil {
					group.SetIntegration(&aws.Integration{})
				}
				group.Integration.SetElasticBeanstalk(integration)
				update = true
			}
		} else {
			if group.Integration == nil {
				group.SetIntegration(&aws.Integration{})
			}
			group.Integration.SetElasticBeanstalk(nil)
			update = true
		}
	}

	if d.HasChange("ec2_container_service_integration") {
		if v, ok := d.GetOk("ec2_container_service_integration"); ok {
			if integration, err := expandAWSGroupEC2ContainerServiceIntegration(v, nullify); err != nil {
				return err
			} else {
				if group.Integration == nil {
					group.SetIntegration(&aws.Integration{})
				}
				group.Integration.SetEC2ContainerService(integration)
				update = true
			}
		} else {
			if group.Integration == nil {
				group.SetIntegration(&aws.Integration{})
			}
			group.Integration.SetEC2ContainerService(nil)
			update = true
		}
	}

	if d.HasChange("kubernetes_integration") {
		if v, ok := d.GetOk("kubernetes_integration"); ok {
			if integration, err := expandAWSGroupKubernetesIntegration(v, nullify); err != nil {
				return err
			} else {
				if group.Integration == nil {
					group.SetIntegration(&aws.Integration{})
				}
				group.Integration.SetKubernetes(integration)
				update = true
			}
		} else {
			if group.Integration == nil {
				group.SetIntegration(&aws.Integration{})
			}
			group.Integration.SetKubernetes(nil)
			update = true
		}
	}

	if d.HasChange("nomad_integration") {
		if v, ok := d.GetOk("nomad_integration"); ok {
			if integration, err := expandAWSGroupNomadIntegration(v, nullify); err != nil {
				return err
			} else {
				if group.Integration == nil {
					group.SetIntegration(&aws.Integration{})
				}
				group.Integration.SetNomad(integration)
				update = true
			}
		} else {
			if group.Integration == nil {
				group.SetIntegration(&aws.Integration{})
			}
			group.Integration.SetNomad(nil)
			update = true
		}
	}

	if d.HasChange("mesosphere_integration") {
		if v, ok := d.GetOk("mesosphere_integration"); ok {
			if integration, err := expandAWSGroupMesosphereIntegration(v, nullify); err != nil {
				return err
			} else {
				if group.Integration == nil {
					group.SetIntegration(&aws.Integration{})
				}
				group.Integration.SetMesosphere(integration)
				update = true
			}
		} else {
			if group.Integration == nil {
				group.SetIntegration(&aws.Integration{})
			}
			group.Integration.SetMesosphere(nil)
			update = true
		}
	}

	if d.HasChange("multai_integration") {
		if v, ok := d.GetOk("multai_integration"); ok {
			if integration, err := expandAWSGroupMultaiIntegration(v, nullify); err != nil {
				return err
			} else {
				if group.Integration == nil {
					group.SetIntegration(&aws.Integration{})
				}
				group.Integration.SetMultai(integration)
				update = true
			}
		} else {
			if group.Integration == nil {
				group.SetIntegration(&aws.Integration{})
			}
			group.Integration.SetMultai(nil)
			update = true
		}
	}

	if d.HasChange("codedeploy_integration") {
		if v, ok := d.GetOk("codedeploy_integration"); ok {
			if integration, err := expandAWSGroupCodeDeployIntegration(v, nullify); err != nil {
				return err
			} else {
				if group.Integration == nil {
					group.SetIntegration(&aws.Integration{})
				}
				group.Integration.SetCodeDeploy(integration)
				update = true
			}
		} else {
			if group.Integration == nil {
				group.SetIntegration(&aws.Integration{})
			}
			group.Integration.SetCodeDeploy(nil)
			update = true
		}
	}

	if update {
		var shouldResumeStateful bool
		var input *aws.UpdateGroupInput

		if _, exist := d.GetOkExists("should_resume_stateful"); exist {
			log.Print("[DEBUG] Resuming paused stateful instances on group if any exist")
			shouldResumeStateful = d.Get("should_resume_stateful").(bool)
		}

		input = &aws.UpdateGroupInput{
			Group:                group,
			ShouldResumeStateful: spotinst.Bool(shouldResumeStateful),
		}

		log.Printf("[DEBUG] Group update configuration: %s", stringutil.Stringify(group))

		if _, err := client.elastigroup.CloudProviderAWS().Update(context.Background(), input); err != nil {
			return fmt.Errorf("failed to update group %s: %s", d.Id(), err)
		} else {
			// On Update Success, roll if required.
			if rc, ok := d.GetOk("roll_config"); ok {
				list := rc.(*schema.Set).List()
				m := list[0].(map[string]interface{})
				if sr, ok := m["should_roll"].(bool); ok && sr != false {
					log.Printf("[DEBUG] User has chosen to roll this group: %s", d.Id())
					if roll, err := expandAWSGroupRollConfig(rc, d.Id()); err != nil {
						log.Printf("[ERROR] Failed to expand roll configuration for group %s: %s", d.Id(), err)
						return err
					} else {
						log.Printf("[DEBUG] Sending roll request to the Spotinst API...")
						if _, err := client.elastigroup.CloudProviderAWS().Roll(context.Background(), roll); err != nil {
							log.Printf("[ERROR] Failed to roll group: %s", err)
						}
					}
				} else {
					log.Printf("[DEBUG] User has chosen not to roll this group: %s", d.Id())
				}
			}

		}
	}

	return resourceSpotinstAWSGroupRead(d, meta)
}

func resourceSpotinstAWSGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	log.Printf("[INFO] Deleting group: %s", d.Id())
	input := &aws.DeleteGroupInput{GroupID: spotinst.String(d.Id())}
	if _, err := client.elastigroup.CloudProviderAWS().Delete(context.Background(), input); err != nil {
		return fmt.Errorf("failed to delete group: %s", err)
	}
	d.SetId("")
	return nil
}

//endregion

//region Flatten methods

func flattenAWSGroupCapacity(capacity *aws.Capacity, targetCapacitySetInCapacity bool) []interface{} {
	result := make(map[string]interface{})

	if targetCapacitySetInCapacity {
		log.Print("[DEBUG] Populating target in capacity block")
		result["target"] = spotinst.IntValue(capacity.Target)
	}

	result["minimum"] = spotinst.IntValue(capacity.Minimum)
	result["maximum"] = spotinst.IntValue(capacity.Maximum)
	result["unit"] = spotinst.StringValue(capacity.Unit)
	return []interface{}{result}
}

func flattenAWSGroupStrategy(strategy *aws.Strategy) []interface{} {
	result := make(map[string]interface{})
	result["risk"] = spotinst.Float64Value(strategy.Risk)
	result["ondemand_count"] = spotinst.IntValue(strategy.OnDemandCount)
	result["availability_vs_cost"] = spotinst.StringValue(strategy.AvailabilityVsCost)
	result["lifetime_period"] = spotinst.StringValue(strategy.LifetimePeriod)
	result["draining_timeout"] = spotinst.IntValue(strategy.DrainingTimeout)
	result["utilize_reserved_instances"] = spotinst.BoolValue(strategy.UtilizeReservedInstances)
	result["fallback_to_ondemand"] = spotinst.BoolValue(strategy.FallbackToOnDemand)
	result["spin_up_time"] = spotinst.IntValue(strategy.SpinUpTime)
	return []interface{}{result}
}

func flattenAWSGroupLaunchSpecification(lspec *aws.LaunchSpecification, includeImageID bool) []interface{} {
	result := make(map[string]interface{})
	result["health_check_grace_period"] = spotinst.IntValue(lspec.HealthCheckGracePeriod)
	result["health_check_unhealthy_duration_before_replacement"] = spotinst.IntValue(lspec.HealthCheckUnhealthyDurationBeforeReplacement)
	result["health_check_type"] = spotinst.StringValue(lspec.HealthCheckType)
	if includeImageID {
		result["image_id"] = spotinst.StringValue(lspec.ImageID)
	}
	result["tenancy"] = spotinst.StringValue(lspec.Tenancy)
	result["key_pair"] = spotinst.StringValue(lspec.KeyPair)
	if lspec.UserData != nil && spotinst.StringValue(lspec.UserData) != "" {
		decodedUserData, _ := base64.StdEncoding.DecodeString(spotinst.StringValue(lspec.UserData))
		result["user_data"] = string(decodedUserData)
	} else {
		result["user_data"] = ""
	}
	if lspec.ShutdownScript != nil && spotinst.StringValue(lspec.ShutdownScript) != "" {
		decodedShutdownScript, _ := base64.StdEncoding.DecodeString(spotinst.StringValue(lspec.ShutdownScript))
		result["shutdown_script"] = string(decodedShutdownScript)
	} else {
		result["shutdown_script"] = ""
	}
	result["monitoring"] = spotinst.BoolValue(lspec.Monitoring)
	result["ebs_optimized"] = spotinst.BoolValue(lspec.EBSOptimized)
	result["load_balancer_names"] = lspec.LoadBalancerNames
	result["security_group_ids"] = lspec.SecurityGroupIDs
	if lspec.IAMInstanceProfile != nil {
		if lspec.IAMInstanceProfile.Arn != nil {
			result["iam_instance_profile"] = spotinst.StringValue(lspec.IAMInstanceProfile.Arn)
		} else {
			result["iam_instance_profile"] = spotinst.StringValue(lspec.IAMInstanceProfile.Name)
		}
	}
	return []interface{}{result}
}

func flattenAWSGroupLoadBalancers(balancers []*aws.LoadBalancer) []interface{} {
	if balancers == nil {
		log.Print("[ERROR] Cannot expand AWS group load balancers due to <nil> value")
		// Do not fail the terraform process
		return nil
	}
	result := make([]interface{}, 0, len(balancers))
	for _, b := range balancers {
		if b == nil {
			log.Print("[ERROR] Empty load balancer value, skipping creation")
			continue
		}
		m := make(map[string]interface{})
		m["type"] = strings.ToLower(spotinst.StringValue(b.Type))

		if b.Name != nil {
			m["name"] = spotinst.StringValue(b.Name)
		}
		if b.Arn != nil {
			m["arn"] = spotinst.StringValue(b.Arn)
		}
		if b.BalancerID != nil {
			m["balancer_id"] = spotinst.StringValue(b.BalancerID)
		}
		if b.TargetSetID != nil {
			m["target_set_id"] = spotinst.StringValue(b.TargetSetID)
		}
		if b.ZoneAwareness != nil {
			m["zone_awareness"] = spotinst.BoolValue(b.ZoneAwareness)
		}
		if b.AutoWeight != nil {
			m["auto_weight"] = spotinst.BoolValue(b.AutoWeight)
		}

		result = append(result, m)
	}
	return result
}

func flattenAWSGroupEBSVolumePool(volumes []*aws.EBSVolume) []interface{} {
	result := make([]interface{}, 0, len(volumes))
	for _, v := range volumes {
		m := make(map[string]interface{})
		m["device_name"] = spotinst.StringValue(v.DeviceName)
		m["volume_ids"] = v.VolumeIDs
		result = append(result, m)
	}
	return result
}

func flattenAWSGroupSignals(signals []*aws.Signal) []interface{} {
	result := make([]interface{}, 0, len(signals))
	for _, s := range signals {
		m := make(map[string]interface{})
		m["name"] = strings.ToLower(spotinst.StringValue(s.Name))
		m["timeout"] = spotinst.IntValue(s.Timeout)
		result = append(result, m)
	}
	return result
}

func flattenAWSGroupScheduledTasks(tasks []*aws.Task) []interface{} {
	result := make([]interface{}, 0, len(tasks))
	for _, t := range tasks {
		m := make(map[string]interface{})
		m["is_enabled"] = spotinst.BoolValue(t.IsEnabled)
		m["task_type"] = spotinst.StringValue(t.Type)
		m["cron_expression"] = spotinst.StringValue(t.CronExpression)
		m["frequency"] = spotinst.StringValue(t.Frequency)
		m["scale_target_capacity"] = spotinst.IntValue(t.ScaleTargetCapacity)
		m["scale_min_capacity"] = spotinst.IntValue(t.ScaleMinCapacity)
		m["scale_max_capacity"] = spotinst.IntValue(t.ScaleMaxCapacity)
		m["batch_size_percentage"] = spotinst.IntValue(t.BatchSizePercentage)
		m["grace_period"] = spotinst.IntValue(t.GracePeriod)
		m["target_capacity"] = spotinst.IntValue(t.TargetCapacity)
		m["min_capacity"] = spotinst.IntValue(t.MinCapacity)
		m["max_capacity"] = spotinst.IntValue(t.MaxCapacity)
		result = append(result, m)
	}
	return result
}

func flattenAWSGroupPersistence(persistence *aws.Persistence) []interface{} {
	result := make(map[string]interface{})
	result["persist_block_devices"] = spotinst.BoolValue(persistence.ShouldPersistBlockDevices)
	result["persist_private_ip"] = spotinst.BoolValue(persistence.ShouldPersistPrivateIP)
	result["persist_root_device"] = spotinst.BoolValue(persistence.ShouldPersistRootDevice)
	result["block_devices_mode"] = spotinst.StringValue(persistence.BlockDevicesMode)
	return []interface{}{result}
}

func flattenAWSGroupNetworkInterfaces(ifaces []*aws.NetworkInterface) []interface{} {
	result := make([]interface{}, 0, len(ifaces))
	for _, iface := range ifaces {
		m := make(map[string]interface{})
		m["associate_public_ip_address"] = spotinst.BoolValue(iface.AssociatePublicIPAddress)
		m["delete_on_termination"] = spotinst.BoolValue(iface.DeleteOnTermination)
		m["description"] = spotinst.StringValue(iface.Description)
		m["device_index"] = spotinst.IntValue(iface.DeviceIndex)
		m["network_interface_id"] = spotinst.StringValue(iface.ID)
		m["private_ip_address"] = spotinst.StringValue(iface.PrivateIPAddress)
		m["secondary_private_ip_address_count"] = spotinst.IntValue(iface.SecondaryPrivateIPAddressCount)
		m["subnet_id"] = spotinst.StringValue(iface.SubnetID)
		m["security_group_ids"] = iface.SecurityGroupsIDs
		result = append(result, m)
	}
	return result
}

func flattenAWSGroupEBSBlockDevices(devices []*aws.BlockDeviceMapping) []interface{} {
	result := make([]interface{}, 0, len(devices))
	for _, dev := range devices {
		if dev.EBS != nil {
			m := make(map[string]interface{})
			m["device_name"] = spotinst.StringValue(dev.DeviceName)
			m["delete_on_termination"] = spotinst.BoolValue(dev.EBS.DeleteOnTermination)
			m["encrypted"] = spotinst.BoolValue(dev.EBS.Encrypted)
			m["iops"] = spotinst.IntValue(dev.EBS.IOPS)
			m["snapshot_id"] = spotinst.StringValue(dev.EBS.SnapshotID)
			m["volume_type"] = spotinst.StringValue(dev.EBS.VolumeType)
			m["volume_size"] = spotinst.IntValue(dev.EBS.VolumeSize)
			result = append(result, m)
		}
	}
	return result
}

func flattenAWSGroupEphemeralBlockDevices(devices []*aws.BlockDeviceMapping) []interface{} {
	result := make([]interface{}, 0, len(devices))
	for _, dev := range devices {
		if dev.EBS == nil {
			m := make(map[string]interface{})
			m["device_name"] = spotinst.StringValue(dev.DeviceName)
			m["virtual_name"] = spotinst.StringValue(dev.VirtualName)
			result = append(result, m)
		}
	}
	return result
}

func flattenAWSGroupRancherIntegration(integration *aws.RancherIntegration) []interface{} {
	result := make(map[string]interface{})
	result["master_host"] = spotinst.StringValue(integration.MasterHost)
	result["access_key"] = spotinst.StringValue(integration.AccessKey)
	result["secret_key"] = spotinst.StringValue(integration.SecretKey)
	return []interface{}{result}
}

func flattenAWSGroupElasticBeanstalkIntegration(integration *aws.ElasticBeanstalkIntegration) []interface{} {
	result := make(map[string]interface{})
	result["environment_id"] = spotinst.StringValue(integration.EnvironmentID)
	return []interface{}{result}
}

func flattenAWSGroupKubernetesIntegration(integration *aws.KubernetesIntegration) []interface{} {
	result := make(map[string]interface{})
	result["api_server"] = spotinst.StringValue(integration.Server)
	result["token"] = spotinst.StringValue(integration.Token)
	return []interface{}{result}
}

func flattenAWSGroupMesosphereIntegration(integration *aws.MesosphereIntegration) []interface{} {
	result := make(map[string]interface{})
	result["api_server"] = spotinst.StringValue(integration.Server)
	return []interface{}{result}
}

func flattenAWSGroupMultaiIntegration(integration *aws.MultaiIntegration) []interface{} {
	result := make(map[string]interface{})
	result["deployment_id"] = spotinst.StringValue(integration.DeploymentID)
	return []interface{}{result}
}

func flattenAWSGroupCodeDeployIntegration(integration *aws.CodeDeployIntegration) []interface{} {
	result := make(map[string]interface{})
	result["cleanup_on_failure"] = spotinst.BoolValue(integration.CleanUpOnFailure)
	result["terminate_instance_on_failure"] = spotinst.BoolValue(integration.TerminateInstanceOnFailure)

	deploymentGroups := make([]interface{}, len(integration.DeploymentGroups))
	for i, dg := range integration.DeploymentGroups {
		m := make(map[string]interface{})
		m["application_name"] = spotinst.StringValue(dg.ApplicationName)
		m["deployment_group_name"] = spotinst.StringValue(dg.DeploymentGroupName)
		deploymentGroups[i] = m
	}

	return []interface{}{result}
}

//endregion

//region Build method

// buildAWSGroupOpts builds the Spotinst AWS Group options.
func buildAWSGroupOpts(d *schema.ResourceData, meta interface{}) (*aws.Group, error) {
	group := &aws.Group{
		Scaling:     &aws.Scaling{},
		Scheduling:  &aws.Scheduling{},
		Integration: &aws.Integration{},
		Compute: &aws.Compute{
			LaunchSpecification: &aws.LaunchSpecification{},
		},
	}
	nullify := false

	group.SetName(spotinst.String(d.Get("name").(string)))
	group.SetDescription(spotinst.String(d.Get("description").(string)))
	group.Compute.SetProduct(spotinst.String(d.Get("product").(string)))

	if tfPrivateIPs, ok := d.GetOk("private_ips"); ok {
		if privateIPsArr, err := expandAWSGroupPrivateIPs(tfPrivateIPs); err != nil {
			return nil, err
		} else {
			group.Compute.SetPrivateIPs(privateIPsArr)
		}
	}

	if v, ok := d.GetOk("capacity"); ok {
		_, exists := d.GetOkExists("target_capacity")
		if capacity, err := expandAWSGroupCapacity(v, nullify, !exists, false); err != nil {
			return nil, err
		} else {
			group.SetCapacity(capacity)
		}
	}

	if _, exists := d.GetOkExists("target_capacity"); exists {
		group.Capacity.SetTarget(spotinst.Int(d.Get("target_capacity").(int)))
	}

	if v, ok := d.GetOk("strategy"); ok {
		if strategy, err := expandAWSGroupStrategy(v, nullify); err != nil {
			return nil, err
		} else {
			group.SetStrategy(strategy)
		}
	}

	if v, ok := d.GetOk("persistence"); ok {
		if persistence, err := expandAWSGroupPersistence(v, nullify); err != nil {
			return nil, err
		} else {
			group.Strategy.SetPersistence(persistence)
		}
	}

	if v, ok := d.GetOk("scaling_up_policy"); ok {
		if policies, err := expandAWSGroupScalingPolicies(v, nullify); err != nil {
			return nil, err
		} else {
			group.Scaling.SetUp(policies)
		}
	}

	if v, ok := d.GetOk("scaling_down_policy"); ok {
		if policies, err := expandAWSGroupScalingPolicies(v, nullify); err != nil {
			return nil, err
		} else {
			group.Scaling.SetDown(policies)
		}
	}

	if v, ok := d.GetOk("scaling_target_policy"); ok {
		if policies, err := expandAWSGroupScalingPolicies(v, nullify); err != nil {
			return nil, err
		} else {
			group.Scaling.SetTarget(policies)
		}
	}

	if v, ok := d.GetOk("scheduled_task"); ok {
		if tasks, err := expandAWSGroupScheduledTasks(v, nullify); err != nil {
			return nil, err
		} else {
			group.Scheduling.SetTasks(tasks)
		}
	}

	if v, ok := d.GetOk("instance_types"); ok {
		if types, err := expandAWSGroupInstanceTypes(v, nullify); err != nil {
			return nil, err
		} else {
			group.Compute.SetInstanceTypes(types)
		}
	}

	if v, ok := d.GetOk("elastic_ips"); ok {
		if eips, err := expandAWSGroupElasticIPs(v, nullify); err != nil {
			return nil, err
		} else {
			group.Compute.SetElasticIPs(eips)
		}
	}

	if v, ok := d.GetOk("availability_zone"); ok {
		if zones, err := expandAWSGroupAvailabilityZones(v, nullify); err != nil {
			return nil, err
		} else {
			group.Compute.SetAvailabilityZones(zones)
		}
	}

	if v, ok := d.GetOk("availability_zones"); ok {
		if zones, err := expandAWSGroupAvailabilityZonesSlice(v, nullify); err != nil {
			return nil, err
		} else {
			group.Compute.SetAvailabilityZones(zones)
		}
	}

	if v, ok := d.GetOk("hot_ebs_volume"); ok {
		if ebsVolumePool, err := expandAWSGroupEBSVolumePool(v, nullify); err != nil {
			return nil, err
		} else {
			group.Compute.SetEBSVolumePool(ebsVolumePool)
		}
	}

	if v, ok := d.GetOk("signal"); ok {
		if signals, err := expandAWSGroupSignals(v, nullify); err != nil {
			return nil, err
		} else {
			group.Strategy.SetSignals(signals)
		}
	}

	if v, ok := d.GetOk("launch_specification"); ok {
		if lc, err := expandAWSGroupLaunchSpecification(v, nullify); err != nil {
			return nil, err
		} else {
			group.Compute.SetLaunchSpecification(lc)
		}
	}

	if v, ok := d.GetOk("image_id"); ok {
		group.Compute.LaunchSpecification.SetImageId(spotinst.String(v.(string)))
	}

	if v, ok := d.GetOk("load_balancer"); ok {
		if lbs, err := expandAWSGroupLoadBalancer(v, nullify); err != nil {
			return nil, err
		} else {
			if group.Compute.LaunchSpecification.LoadBalancersConfig == nil {
				group.Compute.LaunchSpecification.LoadBalancersConfig = &aws.LoadBalancersConfig{}
			}
			group.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(lbs)
		}
	}

	if v, ok := d.GetOk("load_balancers"); ok {
		if lbs, err := expandAWSGroupLoadBalancers(v, nullify); err != nil {
			return nil, err
		} else {
			if group.Compute.LaunchSpecification.LoadBalancersConfig == nil {
				group.Compute.LaunchSpecification.LoadBalancersConfig = &aws.LoadBalancersConfig{}
			}
			group.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(lbs)
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		if tags, err := expandAWSGroupTags(v, nullify); err != nil {
			return nil, err
		} else {
			group.Compute.LaunchSpecification.SetTags(tags)
		}
	}

	if v, ok := d.GetOk("tags_kv"); ok {
		if tags, err := expandAWSGroupTagsKV(v, nullify); err != nil {
			return nil, err
		} else {
			group.Compute.LaunchSpecification.SetTags(tags)
		}
	}

	if v, ok := d.GetOkExists("instance_type_weights"); ok && v != "" {
		if weights, err := expandAWSGroupInstanceTypeWeights(v, nullify); err != nil {
			return nil, err
		} else {
			group.Compute.InstanceTypes.SetWeights(weights)
			log.Printf("[DEBUG] Group instance type weights configuration: %s", stringutil.Stringify(weights))
		}
	}

	if v, ok := d.GetOk("network_interface"); ok {
		if interfaces, err := expandAWSGroupNetworkInterfaces(v, nullify); err != nil {
			return nil, err
		} else {
			group.Compute.LaunchSpecification.SetNetworkInterfaces(interfaces)
		}
	}

	if v, ok := d.GetOk("ebs_block_device"); ok {
		if devices, err := expandAWSGroupEBSBlockDevices(v, nullify); err != nil {
			return nil, err
		} else {
			group.Compute.LaunchSpecification.SetBlockDeviceMappings(devices)
		}
	}

	if v, ok := d.GetOk("ephemeral_block_device"); ok {
		if devices, err := expandAWSGroupEphemeralBlockDevices(v, nullify); err != nil {
			return nil, err
		} else {
			if len(group.Compute.LaunchSpecification.BlockDeviceMappings) > 0 {
				all := append(group.Compute.LaunchSpecification.BlockDeviceMappings, devices...)
				group.Compute.LaunchSpecification.SetBlockDeviceMappings(all)
			} else {
				group.Compute.LaunchSpecification.SetBlockDeviceMappings(devices)
			}
		}
	}

	if v, ok := d.GetOk("rancher_integration"); ok {
		if integration, err := expandAWSGroupRancherIntegration(v, nullify); err != nil {
			return nil, err
		} else {
			group.Integration.SetRancher(integration)
		}
	}

	if v, ok := d.GetOk("elastic_beanstalk_integration"); ok {
		if integration, err := expandAWSGroupElasticBeanstalkIntegration(v, nullify); err != nil {
			return nil, err
		} else {
			group.Integration.SetElasticBeanstalk(integration)
		}
	}

	if v, ok := d.GetOk("ec2_container_service_integration"); ok {
		if integration, err := expandAWSGroupEC2ContainerServiceIntegration(v, nullify); err != nil {
			return nil, err
		} else {
			group.Integration.SetEC2ContainerService(integration)
		}
	}

	if v, ok := d.GetOk("kubernetes_integration"); ok {
		if integration, err := expandAWSGroupKubernetesIntegration(v, nullify); err != nil {
			return nil, err
		} else {
			group.Integration.SetKubernetes(integration)
		}
	}

	if v, ok := d.GetOk("nomad_integration"); ok {
		if integration, err := expandAWSGroupNomadIntegration(v, nullify); err != nil {
			return nil, err
		} else {
			group.Integration.SetNomad(integration)
		}
	}

	if v, ok := d.GetOk("mesosphere_integration"); ok {
		if integration, err := expandAWSGroupMesosphereIntegration(v, nullify); err != nil {
			return nil, err
		} else {
			group.Integration.SetMesosphere(integration)
		}
	}

	if v, ok := d.GetOk("multai_integration"); ok {
		if integration, err := expandAWSGroupMultaiIntegration(v, nullify); err != nil {
			return nil, err
		} else {
			group.Integration.SetMultai(integration)
		}
	}

	if v, ok := d.GetOk("codedeploy_integration"); ok {
		if integration, err := expandAWSGroupCodeDeployIntegration(v, nullify); err != nil {
			return nil, err
		} else {
			group.Integration.SetCodeDeploy(integration)
		}
	}

	return group, nil
}

//endregion

//region Expand methods

// expandAWSGroupCapacity expands the Capacity block.
func expandAWSGroupCapacity(data interface{}, nullify, withTarget, isUpdate bool) (*aws.Capacity, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	capacity := &aws.Capacity{}

	if v, ok := m["minimum"].(int); ok && v >= 0 {
		capacity.SetMinimum(spotinst.Int(v))
	}

	if v, ok := m["maximum"].(int); ok && v >= 0 {
		capacity.SetMaximum(spotinst.Int(v))
	}

	if withTarget {
		if v, ok := m["target"].(int); ok && v >= 0 {
			capacity.SetTarget(spotinst.Int(v))
		}
	}

	if !isUpdate {
		if v, ok := m["unit"].(string); ok && v != "" {
			capacity.SetUnit(spotinst.String(v))
		}
	}

	log.Printf("[DEBUG] Group capacity configuration: %s", stringutil.Stringify(capacity))
	return capacity, nil
}

// expandAWSGroupStrategy expands the Strategy block.
func expandAWSGroupStrategy(data interface{}, nullify bool) (*aws.Strategy, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	strategy := &aws.Strategy{}

	if v, ok := m["risk"].(float64); ok && v >= 0 {
		strategy.SetRisk(spotinst.Float64(v))
	}

	if v, ok := m["ondemand_count"].(int); ok && v > 0 {
		strategy.SetOnDemandCount(spotinst.Int(v))
	}

	if v, ok := m["availability_vs_cost"].(string); ok && v != "" {
		strategy.SetAvailabilityVsCost(spotinst.String(v))
	}

	if v, ok := m["lifetime_period"].(string); ok && v != "" {
		strategy.SetLifetimePeriod(spotinst.String(v))
	} else if nullify {
		strategy.SetLifetimePeriod(nil)
	}

	if v, ok := m["draining_timeout"].(int); ok && v > 0 {
		strategy.SetDrainingTimeout(spotinst.Int(v))
	}

	if v, ok := m["utilize_reserved_instances"].(bool); ok && v {
		strategy.SetUtilizeReservedInstances(spotinst.Bool(v))
	} else if nullify {
		strategy.SetUtilizeReservedInstances(nil)
	}

	if v, ok := m["fallback_to_ondemand"].(bool); ok && v {
		strategy.SetFallbackToOnDemand(spotinst.Bool(v))
	} else if nullify {
		strategy.SetFallbackToOnDemand(nil)
	}

	if v, ok := m["spin_up_time"].(int); ok && v > 0 {
		strategy.SetSpinUpTime(spotinst.Int(v))
	}

	log.Printf("[DEBUG] Group strategy configuration: %s", stringutil.Stringify(strategy))
	return strategy, nil
}

// expandAWSGroupScalingPolicies expands the Scaling Policy block.
func expandAWSGroupScalingPolicies(data interface{}, nullify bool) ([]*aws.ScalingPolicy, error) {
	list := data.(*schema.Set).List()
	policies := make([]*aws.ScalingPolicy, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		policy := &aws.ScalingPolicy{}

		if v, ok := m["policy_name"].(string); ok && v != "" {
			policy.SetPolicyName(spotinst.String(v))
		}

		if v, ok := m["metric_name"].(string); ok && v != "" {
			policy.SetMetricName(spotinst.String(v))
		}

		if v, ok := m["namespace"].(string); ok && v != "" {
			policy.SetNamespace(spotinst.String(v))
		}

		if v, ok := m["source"].(string); ok && v != "" {
			policy.SetSource(spotinst.String(v))
		}

		if v, ok := m["statistic"].(string); ok && v != "" {
			policy.SetStatistic(spotinst.String(v))
		}

		if v, ok := m["unit"].(string); ok && v != "" {
			policy.SetUnit(spotinst.String(v))
		}

		if v, ok := m["threshold"].(float64); ok && v > 0 {
			policy.SetThreshold(spotinst.Float64(v))
		}

		if v, ok := m["operator"].(string); ok && v != "" {
			policy.SetOperator(spotinst.String(v))
		}

		if v, ok := m["period"].(int); ok && v > 0 {
			policy.SetPeriod(spotinst.Int(v))
		}

		if v, ok := m["evaluation_periods"].(int); ok && v > 0 {
			policy.SetEvaluationPeriods(spotinst.Int(v))
		}

		if v, ok := m["cooldown"].(int); ok && v > 0 {
			policy.SetCooldown(spotinst.Int(v))
		}

		if v, ok := m["dimensions"]; ok {
			dimensions := expandAWSGroupScalingPolicyDimensions(v.(map[string]interface{}))
			if len(dimensions) > 0 {
				policy.SetDimensions(dimensions)
			}
		}

		if v, ok := m["action_type"].(string); ok && v != "" {
			action := &aws.Action{}
			action.SetType(spotinst.String(v))

			if v, ok := m["adjustment"].(int); ok && v > 0 {
				action.SetAdjustment(spotinst.String(strconv.Itoa(v)))
			} else if v, ok := m["adjustment_expression"].(string); ok && v != "" {
				action.SetAdjustment(spotinst.String(v))
			}

			if v, ok := m["min_target_capacity"].(int); ok && v > 0 {
				action.SetMinTargetCapacity(spotinst.Int(v))
			}

			if v, ok := m["max_target_capacity"].(int); ok && v > 0 {
				action.SetMaxTargetCapacity(spotinst.Int(v))
			}

			if v, ok := m["minimum"].(int); ok && v > 0 {
				action.SetMinimum(spotinst.Int(v))
			}

			if v, ok := m["maximum"].(int); ok && v > 0 {
				action.SetMaximum(spotinst.Int(v))
			}

			if v, ok := m["target"].(int); ok && v > 0 {
				action.SetTarget(spotinst.Int(v))
			}

			policy.SetAction(action)
		} else {
			if v, ok := m["adjustment"].(int); ok && v > 0 {
				policy.SetAdjustment(spotinst.Int(v))
			}

			if v, ok := m["min_target_capacity"].(int); ok && v > 0 {
				policy.SetMinTargetCapacity(spotinst.Int(v))
			}

			if v, ok := m["max_target_capacity"].(int); ok && v > 0 {
				policy.SetMaxTargetCapacity(spotinst.Int(v))
			}
		}

		// Target scaling policy?
		if policy.Threshold == nil {
			if v, ok := m["target"].(float64); ok && v >= 0 {
				policy.SetTarget(spotinst.Float64(v))
			}
		}

		if policy.Namespace != nil {
			log.Printf("[DEBUG] Group scaling policy configuration: %s", stringutil.Stringify(policy))
			policies = append(policies, policy)
		}
	}

	return policies, nil
}

func expandAWSGroupScalingPolicyDimensions(list map[string]interface{}) []*aws.Dimension {
	dimensions := make([]*aws.Dimension, 0, len(list))
	for name, val := range list {
		dimension := &aws.Dimension{}
		dimension.SetName(spotinst.String(name))
		dimension.SetValue(spotinst.String(val.(string)))
		log.Printf("[DEBUG] Group scaling policy dimension: %s", stringutil.Stringify(dimension))
		dimensions = append(dimensions, dimension)
	}
	return dimensions
}

const taskTypeStatefulUpdateCapacity = "statefulUpdateCapacity"

// expandAWSGroupScheduledTasks expands the Scheduled Task block.
func expandAWSGroupScheduledTasks(data interface{}, nullify bool) ([]*aws.Task, error) {
	list := data.(*schema.Set).List()
	tasks := make([]*aws.Task, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		task := &aws.Task{}

		if v, ok := m["is_enabled"].(bool); ok {
			task.SetIsEnabled(spotinst.Bool(v))
		}

		if v, ok := m["task_type"].(string); ok && v != "" {
			task.SetType(spotinst.String(v))
		}

		if v, ok := m["frequency"].(string); ok && v != "" {
			task.SetFrequency(spotinst.String(v))
		}

		if v, ok := m["cron_expression"].(string); ok && v != "" {
			task.SetCronExpression(spotinst.String(v))
		}

		if v, ok := m["batch_size_percentage"].(int); ok && v > 0 {
			task.SetBatchSizePercentage(spotinst.Int(v))
		}

		if v, ok := m["grace_period"].(int); ok && v > 0 {
			task.SetGracePeriod(spotinst.Int(v))
		}

		if spotinst.StringValue(task.Type) != taskTypeStatefulUpdateCapacity {
			if v, ok := m["scale_target_capacity"].(int); ok && v >= 0 {
				task.SetScaleTargetCapacity(spotinst.Int(v))
			}

			if v, ok := m["scale_min_capacity"].(int); ok && v >= 0 {
				task.SetScaleMinCapacity(spotinst.Int(v))
			}

			if v, ok := m["scale_max_capacity"].(int); ok && v >= 0 {
				task.SetScaleMaxCapacity(spotinst.Int(v))
			}
		}

		if spotinst.StringValue(task.Type) == taskTypeStatefulUpdateCapacity {
			if v, ok := m["target_capacity"].(int); ok && v >= 0 {
				task.SetTargetCapacity(spotinst.Int(v))
			}

			if v, ok := m["min_capacity"].(int); ok && v >= 0 {
				task.SetMinCapacity(spotinst.Int(v))
			}

			if v, ok := m["max_capacity"].(int); ok && v >= 0 {
				task.SetMaxCapacity(spotinst.Int(v))
			}
		}

		log.Printf("[DEBUG] Group scheduled task configuration: %s", stringutil.Stringify(task))
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// expandAWSGroupPersistence expands the Persistence block.
func expandAWSGroupPersistence(data interface{}, nullify bool) (*aws.Persistence, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	persistence := &aws.Persistence{}

	if v, ok := m["persist_private_ip"].(bool); ok {
		persistence.SetShouldPersistPrivateIP(spotinst.Bool(v))
	} else if nullify {
		persistence.SetShouldPersistPrivateIP(nil)
	}

	if v, ok := m["persist_root_device"].(bool); ok {
		persistence.SetShouldPersistRootDevice(spotinst.Bool(v))
	} else if nullify {
		persistence.SetShouldPersistRootDevice(nil)
	}

	if v, ok := m["persist_block_devices"].(bool); ok {
		persistence.SetShouldPersistBlockDevices(spotinst.Bool(v))
	} else if nullify {
		persistence.SetShouldPersistBlockDevices(nil)
	}

	if v, ok := m["block_devices_mode"].(string); ok {
		persistence.SetBlockDevicesMode(spotinst.String(v))
	} else if nullify {
		persistence.SetBlockDevicesMode(nil)
	}

	log.Printf("[DEBUG] Group persistence configuration: %s", stringutil.Stringify(persistence))
	return persistence, nil
}

// expandAWSGroupAvailabilityZones expands the Availability Zone block.
func expandAWSGroupAvailabilityZones(data interface{}, nullify bool) ([]*aws.AvailabilityZone, error) {
	list := data.(*schema.Set).List()
	zones := make([]*aws.AvailabilityZone, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		zone := &aws.AvailabilityZone{}

		if v, ok := m["name"].(string); ok && v != "" {
			zone.SetName(spotinst.String(v))
		}

		if v, ok := m["subnet_id"].(string); ok && v != "" {
			zone.SetSubnetId(spotinst.String(v))
		}

		if v, ok := m["placement_group_name"].(string); ok && v != "" {
			zone.SetPlacementGroupName(spotinst.String(v))
		}

		log.Printf("[DEBUG] Group availability zone configuration: %s", stringutil.Stringify(zone))
		zones = append(zones, zone)
	}

	return zones, nil
}

// expandAWSGroupAvailabilityZonesSlice expands the Availability Zone block when provided as a slice.
func expandAWSGroupAvailabilityZonesSlice(data interface{}, nullify bool) ([]*aws.AvailabilityZone, error) {
	list := data.([]interface{})
	zones := make([]*aws.AvailabilityZone, 0, len(list))
	for _, str := range list {
		if s, ok := str.(string); ok {
			parts := strings.Split(s, ":")
			zone := &aws.AvailabilityZone{}
			if len(parts) >= 1 && parts[0] != "" {
				zone.SetName(spotinst.String(parts[0]))
			}
			if len(parts) == 2 && parts[1] != "" {
				zone.SetSubnetId(spotinst.String(parts[1]))
			}
			if len(parts) == 3 && parts[2] != "" {
				zone.SetPlacementGroupName(spotinst.String(parts[2]))
			}
			log.Printf("[DEBUG] Group availability zone configuration: %s", stringutil.Stringify(zone))
			zones = append(zones, zone)
		}
	}

	return zones, nil
}

// expandAWSGroupEBSVolumePool expands the EBS Volume Pool block.
func expandAWSGroupEBSVolumePool(data interface{}, nullify bool) ([]*aws.EBSVolume, error) {
	list := data.(*schema.Set).List()
	volumes := make([]*aws.EBSVolume, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		volume := &aws.EBSVolume{}

		if v, ok := m["device_name"].(string); ok && len(v) > 0 {
			volume.SetDeviceName(spotinst.String(v))
		}

		if v, ok := m["volume_ids"].([]interface{}); ok {
			ids := make([]string, 0, len(v))
			for _, id := range v {
				if v, ok := id.(string); ok && len(v) > 0 {
					ids = append(ids, v)
				}
			}
			volume.SetVolumeIDs(ids)
		}

		if volume.DeviceName != nil && len(volume.VolumeIDs) > 0 {
			log.Printf("[DEBUG] Group EBS volume (pool) configuration: %s", stringutil.Stringify(volume))
			volumes = append(volumes, volume)
		}
	}

	return volumes, nil
}

// expandAWSGroupSignals expands the Signal block.
func expandAWSGroupSignals(data interface{}, nullify bool) ([]*aws.Signal, error) {
	list := data.(*schema.Set).List()
	signals := make([]*aws.Signal, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		signal := &aws.Signal{}

		if v, ok := m["name"].(string); ok && v != "" {
			signal.SetName(spotinst.String(strings.ToUpper(v)))
		}

		if v, ok := m["timeout"].(int); ok && v > 0 {
			signal.SetTimeout(spotinst.Int(v))
		}

		log.Printf("[DEBUG] Group signal configuration: %s", stringutil.Stringify(signal))
		signals = append(signals, signal)
	}

	return signals, nil
}

// expandAWSGroupInstanceTypes expands the Instance Types block.
func expandAWSGroupInstanceTypes(data interface{}, nullify bool) (*aws.InstanceTypes, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	types := &aws.InstanceTypes{}
	if v, ok := m["ondemand"].(string); ok && v != "" {
		types.SetOnDemand(spotinst.String(v))
	}
	if v, ok := m["spot"].([]interface{}); ok {
		it := make([]string, len(v))
		for i, j := range v {
			it[i] = j.(string)
		}
		types.SetSpot(it)
	}

	log.Printf("[DEBUG] Group instance types configuration: %s", stringutil.Stringify(types))
	return types, nil
}

// expandAWSGroupNetworkInterfaces expands the Elastic Network Interface block.
func expandAWSGroupNetworkInterfaces(data interface{}, nullify bool) ([]*aws.NetworkInterface, error) {
	list := data.(*schema.Set).List()
	interfaces := make([]*aws.NetworkInterface, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		iface := &aws.NetworkInterface{}

		if v, ok := m["network_interface_id"].(string); ok && v != "" {
			iface.SetId(spotinst.String(v))
		}

		if v, ok := m["description"].(string); ok && v != "" {
			iface.SetDescription(spotinst.String(v))
		}

		if v, ok := m["device_index"].(int); ok && v >= 0 {
			iface.SetDeviceIndex(spotinst.Int(v))
		}

		if v, ok := m["secondary_private_ip_address_count"].(int); ok && v > 0 {
			iface.SetSecondaryPrivateIPAddressCount(spotinst.Int(v))
		}

		if v, ok := m["associate_public_ip_address"].(bool); ok {
			iface.SetAssociatePublicIPAddress(spotinst.Bool(v))
		}

		if v, ok := m["delete_on_termination"].(bool); ok {
			iface.SetDeleteOnTermination(spotinst.Bool(v))
		}

		if v, ok := m["private_ip_address"].(string); ok && v != "" {
			iface.SetPrivateIPAddress(spotinst.String(v))
		}

		if v, ok := m["subnet_id"].(string); ok && v != "" {
			iface.SetSubnetId(spotinst.String(v))
		}

		if v, ok := m["security_group_ids"].([]interface{}); ok {
			ids := make([]string, len(v))
			for i, j := range v {
				ids[i] = j.(string)
			}
			iface.SetSecurityGroupsIDs(ids)
		}

		log.Printf("[DEBUG] Group network interface configuration: %s", stringutil.Stringify(iface))
		interfaces = append(interfaces, iface)
	}

	return interfaces, nil
}

// expandAWSGroupEphemeralBlockDevice expands the Ephemeral Block Device block.
func expandAWSGroupEphemeralBlockDevices(data interface{}, nullify bool) ([]*aws.BlockDeviceMapping, error) {
	list := data.(*schema.Set).List()
	devices := make([]*aws.BlockDeviceMapping, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		device := &aws.BlockDeviceMapping{}

		if v, ok := m["device_name"].(string); ok && v != "" {
			device.SetDeviceName(spotinst.String(v))
		}

		if v, ok := m["virtual_name"].(string); ok && v != "" {
			device.SetVirtualName(spotinst.String(v))
		}

		log.Printf("[DEBUG] Group ephemeral block device configuration: %s", stringutil.Stringify(device))
		devices = append(devices, device)
	}

	return devices, nil
}

// expandAWSGroupEBSBlockDevices expands the EBS Block Device block.
func expandAWSGroupEBSBlockDevices(data interface{}, nullify bool) ([]*aws.BlockDeviceMapping, error) {
	list := data.(*schema.Set).List()
	devices := make([]*aws.BlockDeviceMapping, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		device := &aws.BlockDeviceMapping{EBS: &aws.EBS{}}

		if v, ok := m["device_name"].(string); ok && v != "" {
			device.SetDeviceName(spotinst.String(v))
		}

		if v, ok := m["delete_on_termination"].(bool); ok {
			device.EBS.SetDeleteOnTermination(spotinst.Bool(v))
		}

		if v, ok := m["encrypted"].(bool); ok && v != false {
			device.EBS.SetEncrypted(spotinst.Bool(v))
		}

		if v, ok := m["snapshot_id"].(string); ok && v != "" {
			device.EBS.SetSnapshotId(spotinst.String(v))
		}

		if v, ok := m["volume_type"].(string); ok && v != "" {
			device.EBS.SetVolumeType(spotinst.String(v))
		}

		if v, ok := m["volume_size"].(int); ok && v > 0 {
			device.EBS.SetVolumeSize(spotinst.Int(v))
		}

		if v, ok := m["iops"].(int); ok && v > 0 {
			device.EBS.SetIOPS(spotinst.Int(v))
		}

		log.Printf("[DEBUG] Group elastic block device configuration: %s", stringutil.Stringify(device))
		devices = append(devices, device)
	}

	return devices, nil
}

// iprofArnRE is a regular expression for matching IAM instance profile ARNs.
var iprofArnRE = regexp.MustCompile(`arn:aws:iam::\d{12}:instance-profile/?[a-zA-Z_0-9+=,.@\-_/]+`)

// expandAWSGroupLaunchSpecification expands the launch Specification block.
func expandAWSGroupLaunchSpecification(data interface{}, nullify bool) (*aws.LaunchSpecification, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	lc := &aws.LaunchSpecification{}

	if v, ok := m["monitoring"].(bool); ok {
		lc.SetMonitoring(spotinst.Bool(v))
	}

	if v, ok := m["ebs_optimized"].(bool); ok {
		lc.SetEBSOptimized(spotinst.Bool(v))
	}

	if v, ok := m["image_id"].(string); ok && v != "" {
		lc.SetImageId(spotinst.String(v))
	}

	if v, ok := m["tenancy"].(string); ok && v != "" {
		lc.SetTenancy(spotinst.String(v))
	} else if nullify {
		lc.SetTenancy(nil)
	}

	if v, ok := m["key_pair"].(string); ok && v != "" {
		lc.SetKeyPair(spotinst.String(v))
	}

	if v, ok := m["health_check_type"].(string); ok && v != "" {
		lc.SetHealthCheckType(spotinst.String(v))
	} else if nullify {
		lc.SetHealthCheckType(nil)
	}

	if v, ok := m["health_check_grace_period"].(int); ok && v > 0 {
		lc.SetHealthCheckGracePeriod(spotinst.Int(v))
	} else if nullify {
		lc.SetHealthCheckGracePeriod(nil)
	}

	if v, ok := m["health_check_unhealthy_duration_before_replacement"].(int); ok && v > 0 {
		lc.SetHealthCheckUnhealthyDurationBeforeReplacement(spotinst.Int(v))
	} else if nullify {
		lc.SetHealthCheckUnhealthyDurationBeforeReplacement(nil)
	}

	if v, ok := m["iam_instance_profile"].(string); ok && v != "" {
		iprof := &aws.IAMInstanceProfile{}
		if iprofArnRE.MatchString(v) {
			iprof.SetArn(spotinst.String(v))
		} else {
			iprof.SetName(spotinst.String(v))
		}
		lc.SetIAMInstanceProfile(iprof)
	} else if nullify {
		lc.SetIAMInstanceProfile(nil)
	}

	if v, ok := m["user_data"].(string); ok && v != "" {
		lc.SetUserData(spotinst.String(base64Encode(v)))
	} else if nullify {
		lc.SetUserData(nil)
	}

	if v, ok := m["shutdown_script"].(string); ok && v != "" {
		lc.SetShutdownScript(spotinst.String(base64Encode(v)))
	} else if nullify {
		lc.SetShutdownScript(nil)
	}

	if v, ok := m["security_group_ids"].([]interface{}); ok {
		ids := make([]string, len(v))
		for i, j := range v {
			ids[i] = j.(string)
		}
		lc.SetSecurityGroupIDs(ids)
	}

	if v, ok := m["load_balancer_names"].([]interface{}); ok {
		var names []string
		for _, j := range v {
			if name, ok := j.(string); ok && name != "" {
				names = append(names, name)
			}
		}
		lc.SetLoadBalancerNames(names)
	} else if nullify {
		lc.SetLoadBalancerNames(nil)
	}

	log.Printf("[DEBUG] Group launch specification configuration: %s", stringutil.Stringify(lc))
	return lc, nil
}

const (
	loadBalanacerTypeClassic         = "CLASSIC"
	loadBalanacerTypeTargetGroup     = "TARGET_GROUP"
	loadBalanacerTypeMultaiTargetSet = "MULTAI_TARGET_SET"
)

// expandAWSGroupLoadBalancer expands the Load Balancer block.
func expandAWSGroupLoadBalancer(data interface{}, nullify bool) ([]*aws.LoadBalancer, error) {
	if data == nil {
		log.Print("[ERROR] Cannot expand AWS group load balancers due to <nil> value")
		// Do not fail the terraform process
		return nil, nil
	}
	list := data.(*schema.Set).List()
	lbs := make([]*aws.LoadBalancer, 0, len(list))
	for _, item := range list {
		if item == nil {
			log.Print("[ERROR] Empty load balancer value, skipping creation")
			continue
		}
		m := item.(map[string]interface{})
		lb := &aws.LoadBalancer{}

		if v, ok := m["type"].(string); ok && v != "" {
			lb.SetType(spotinst.String(strings.ToUpper(v)))
		}

		if v, ok := m["name"].(string); ok && v != "" {
			lb.SetName(spotinst.String(v))
		}

		if spotinst.StringValue(lb.Type) == loadBalanacerTypeTargetGroup {
			if v, ok := m["arn"].(string); ok && v != "" {
				lb.SetArn(spotinst.String(v))
			}
		}

		if spotinst.StringValue(lb.Type) == loadBalanacerTypeMultaiTargetSet {
			if v, ok := m["balancer_id"].(string); ok && v != "" {
				lb.SetBalancerId(spotinst.String(v))
			}

			if v, ok := m["target_set_id"].(string); ok && v != "" {
				lb.SetTargetSetId(spotinst.String(v))
			}

			if v, ok := m["zone_awareness"].(bool); ok {
				lb.SetZoneAwareness(spotinst.Bool(v))
			}

			if v, ok := m["auto_weight"].(bool); ok {
				lb.SetAutoWeight(spotinst.Bool(v))
			}
		}

		log.Printf("[DEBUG] Group load balancer configuration: %s", stringutil.Stringify(lb))
		lbs = append(lbs, lb)
	}

	return lbs, nil
}

// expandAWSGroupLoadBalancers expands the Load Balancer block.
func expandAWSGroupLoadBalancers(data interface{}, nullify bool) ([]*aws.LoadBalancer, error) {
	if data == nil {
		log.Print("[ERROR] Cannot expand AWS group load balancers due to <nil> value")
		// Do not fail the terraform process
		return nil, nil
	}
	list := data.([]interface{})
	lbs := make([]*aws.LoadBalancer, 0, len(list))
	for _, item := range list {
		if item == nil {
			log.Print("[ERROR] Empty load balancer value, skipping creation")
			continue
		}
		m := item.(string)
		lb := &aws.LoadBalancer{}

		fields := strings.Split(m, ",")
		for _, field := range fields {
			kv := strings.Split(field, "=")
			if len(kv) == 2 {
				key := kv[0]
				val := spotinst.String(kv[1])
				switch key {
				case "type":
					lb.SetType(val)
				case "name":
					lb.SetName(val)
				case "arn":
					lb.SetArn(val)
				case "balancer_id":
					lb.SetBalancerId(val)
				case "target_set_id":
					lb.SetTargetSetId(val)
				case "auto_weight":
					if kv[1] == "true" {
						lb.SetAutoWeight(spotinst.Bool(true))
					}
				case "zone_awareness":
					if kv[1] == "true" {
						lb.SetZoneAwareness(spotinst.Bool(true))
					}
				}
			}
		}

		log.Printf("[DEBUG] Group load balancer configuration: %s", stringutil.Stringify(lb))
		lbs = append(lbs, lb)
	}

	return lbs, nil
}

// expandAWSGroupRancherIntegration expands the Rancher Integration block.
func expandAWSGroupRancherIntegration(data interface{}, nullify bool) (*aws.RancherIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &aws.RancherIntegration{}

	if v, ok := m["master_host"].(string); ok && v != "" {
		i.SetMasterHost(spotinst.String(v))
	}

	if v, ok := m["access_key"].(string); ok && v != "" {
		i.SetAccessKey(spotinst.String(v))
	}

	if v, ok := m["secret_key"].(string); ok && v != "" {
		i.SetSecretKey(spotinst.String(v))
	}

	log.Printf("[DEBUG] Group Rancher integration configuration: %s", stringutil.Stringify(i))
	return i, nil
}

// expandAWSGroupElasticBeanstalkIntegration expands the Elastic Beanstalk Integration block.
func expandAWSGroupElasticBeanstalkIntegration(data interface{}, nullify bool) (*aws.ElasticBeanstalkIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &aws.ElasticBeanstalkIntegration{}

	if v, ok := m["environment_id"].(string); ok && v != "" {
		i.SetEnvironmentId(spotinst.String(v))
	}

	log.Printf("[DEBUG] Group Elastic Beanstalk integration configuration:  %s", stringutil.Stringify(i))
	return i, nil
}

// expandAWSGroupEC2ContainerServiceIntegration expands the EC2 Container Service Integration block.
func expandAWSGroupEC2ContainerServiceIntegration(data interface{}, nullify bool) (*aws.EC2ContainerServiceIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &aws.EC2ContainerServiceIntegration{}

	if v, ok := m["cluster_name"].(string); ok && v != "" {
		i.SetClusterName(spotinst.String(v))
	}

	if v, ok := m["autoscale_is_enabled"].(bool); ok {
		if i.AutoScale == nil {
			i.SetAutoScale(&aws.AutoScale{})
		}
		i.AutoScale.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m["autoscale_cooldown"].(int); ok && v > 0 {
		if i.AutoScale == nil {
			i.SetAutoScale(&aws.AutoScale{})
		}
		i.AutoScale.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m["autoscale_headroom"]; ok {
		headroom, err := expandAWSGroupAutoScaleHeadroom(v, nullify)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			if i.AutoScale == nil {
				i.SetAutoScale(&aws.AutoScale{})
			}
			i.AutoScale.SetHeadroom(headroom)
		}
	}

	if v, ok := m["autoscale_down"]; ok {
		down, err := expandAWSGroupAutoScaleDown(v, nullify)
		if err != nil {
			return nil, err
		}
		if down != nil {
			if i.AutoScale == nil {
				i.SetAutoScale(&aws.AutoScale{})
			}
			i.AutoScale.SetDown(down)
		}
	}

	log.Printf("[DEBUG] Group ECS integration configuration: %s", stringutil.Stringify(i))
	return i, nil
}

// expandAWSGroupKubernetesIntegration expands the Kubernetes Integration block.
func expandAWSGroupKubernetesIntegration(data interface{}, nullify bool) (*aws.KubernetesIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &aws.KubernetesIntegration{}

	if v, ok := m["integration_mode"].(string); ok && v != "" {
		i.SetIntegrationMode(spotinst.String(v))
	}

	if v, ok := m["cluster_identifier"].(string); ok && v != "" {
		i.SetClusterIdentifier(spotinst.String(v))
	}

	if v, ok := m["api_server"].(string); ok && v != "" {
		i.SetServer(spotinst.String(v))
	}

	if v, ok := m["token"].(string); ok && v != "" {
		i.SetToken(spotinst.String(v))
	}

	if v, ok := m["autoscale_is_enabled"].(bool); ok {
		if i.AutoScale == nil {
			i.SetAutoScale(&aws.AutoScale{})
		}
		i.AutoScale.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m["autoscale_cooldown"].(int); ok && v > 0 {
		if i.AutoScale == nil {
			i.SetAutoScale(&aws.AutoScale{})
		}
		i.AutoScale.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m["autoscale_headroom"]; ok {
		headroom, err := expandAWSGroupAutoScaleHeadroom(v, nullify)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			if i.AutoScale == nil {
				i.SetAutoScale(&aws.AutoScale{})
			}
			i.AutoScale.SetHeadroom(headroom)
		}
	}

	if v, ok := m["autoscale_down"]; ok {
		down, err := expandAWSGroupAutoScaleDown(v, nullify)
		if err != nil {
			return nil, err
		}
		if down != nil {
			if i.AutoScale == nil {
				i.SetAutoScale(&aws.AutoScale{})
			}
			i.AutoScale.SetDown(down)
		}
	}

	log.Printf("[DEBUG] Group Kubernetes integration configuration: %s", stringutil.Stringify(i))
	return i, nil
}

// expandAWSGroupNomadIntegration expands the Nomad Integration block.
func expandAWSGroupNomadIntegration(data interface{}, nullify bool) (*aws.NomadIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &aws.NomadIntegration{}

	if v, ok := m["master_host"].(string); ok && v != "" {
		i.SetMasterHost(spotinst.String(v))
	}

	if v, ok := m["master_port"].(int); ok && v > 0 {
		i.SetMasterPort(spotinst.Int(v))
	}

	if v, ok := m["acl_token"].(string); ok && v != "" {
		i.SetAclToken(spotinst.String(v))
	} else if nullify {
		i.SetAclToken(nil)
	}

	if v, ok := m["autoscale_is_enabled"].(bool); ok {
		if i.AutoScale == nil {
			i.SetAutoScale(&aws.AutoScale{})
		}
		i.AutoScale.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m["autoscale_cooldown"].(int); ok && v > 0 {
		if i.AutoScale == nil {
			i.SetAutoScale(&aws.AutoScale{})
		}
		i.AutoScale.SetCooldown(spotinst.Int(v))
	}

	if v, ok := m["autoscale_headroom"]; ok {
		headroom, err := expandAWSGroupAutoScaleHeadroom(v, nullify)
		if err != nil {
			return nil, err
		}
		if headroom != nil {
			if i.AutoScale == nil {
				i.SetAutoScale(&aws.AutoScale{})
			}
			i.AutoScale.SetHeadroom(headroom)
		}
	}

	if v, ok := m["autoscale_down"]; ok {
		down, err := expandAWSGroupAutoScaleDown(v, nullify)
		if err != nil {
			return nil, err
		}
		if down != nil {
			if i.AutoScale == nil {
				i.SetAutoScale(&aws.AutoScale{})
			}
			i.AutoScale.SetDown(down)
		}
	}

	if v, ok := m["autoscale_constraints"]; ok {
		consts, err := expandAWSGroupAutoScaleConstraints(v, nullify)
		if err != nil {
			return nil, err
		}
		if consts != nil {
			if i.AutoScale == nil {
				i.SetAutoScale(&aws.AutoScale{})
			}
			i.AutoScale.SetConstraints(consts)
		}
	}

	log.Printf("[DEBUG] Group Nomad integration configuration: %s", stringutil.Stringify(i))
	return i, nil
}

func expandAWSGroupAutoScaleHeadroom(data interface{}, nullify bool) (*aws.AutoScaleHeadroom, error) {
	if list := data.(*schema.Set).List(); len(list) > 0 {
		m := list[0].(map[string]interface{})
		i := &aws.AutoScaleHeadroom{}

		if v, ok := m["cpu_per_unit"].(int); ok && v > 0 {
			i.SetCPUPerUnit(spotinst.Int(v))
		}

		if v, ok := m["memory_per_unit"].(int); ok && v > 0 {
			i.SetMemoryPerUnit(spotinst.Int(v))
		}

		if v, ok := m["num_of_units"].(int); ok && v > 0 {
			i.SetNumOfUnits(spotinst.Int(v))
		}

		return i, nil
	}

	return nil, nil
}

func expandAWSGroupAutoScaleDown(data interface{}, nullify bool) (*aws.AutoScaleDown, error) {
	if list := data.(*schema.Set).List(); len(list) > 0 {
		m := list[0].(map[string]interface{})
		i := &aws.AutoScaleDown{}

		if v, ok := m["evaluation_periods"].(int); ok && v > 0 {
			i.SetEvaluationPeriods(spotinst.Int(v))
		}

		return i, nil
	}

	return nil, nil
}

func expandAWSGroupAutoScaleConstraints(data interface{}, nullify bool) ([]*aws.AutoScaleConstraint, error) {
	list := data.(*schema.Set).List()
	out := make([]*aws.AutoScaleConstraint, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr["key"]; !ok {
			return nil, errors.New("invalid constraint attributes: key missing")
		}

		if _, ok := attr["value"]; !ok {
			return nil, errors.New("invalid constraint attributes: value missing")
		}
		c := &aws.AutoScaleConstraint{
			Key:   spotinst.String(fmt.Sprintf("${%s}", attr["key"].(string))),
			Value: spotinst.String(attr["value"].(string)),
		}
		log.Printf("[DEBUG] Group constraint configuration: %s", stringutil.Stringify(c))
		out = append(out, c)
	}
	return out, nil
}

// expandAWSGroupMesosphereIntegration expands the Mesosphere Integration block.
func expandAWSGroupMesosphereIntegration(data interface{}, nullify bool) (*aws.MesosphereIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &aws.MesosphereIntegration{}

	if v, ok := m["api_server"].(string); ok && v != "" {
		i.SetServer(spotinst.String(v))
	}

	log.Printf("[DEBUG] Group Mesosphere integration configuration: %s", stringutil.Stringify(i))
	return i, nil
}

// expandAWSGroupMultaiIntegration expands the Multai Integration block.
func expandAWSGroupMultaiIntegration(data interface{}, nullify bool) (*aws.MultaiIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &aws.MultaiIntegration{}

	if v, ok := m["deployment_id"].(string); ok && v != "" {
		i.SetDeploymentId(spotinst.String(v))
	}

	log.Printf("[DEBUG] Group Multai integration configuration: %s", stringutil.Stringify(i))
	return i, nil
}

// expandAWSGroupCodeDeployIntegration expands the CodeDeploy Integration block.
func expandAWSGroupCodeDeployIntegration(data interface{}, nullify bool) (*aws.CodeDeployIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &aws.CodeDeployIntegration{}

	if v, ok := m["cleanup_on_failure"].(bool); ok {
		i.SetCleanUpOnFailure(spotinst.Bool(v))
	}

	if v, ok := m["terminate_instance_on_failure"].(bool); ok {
		i.SetTerminateInstanceOnFailure(spotinst.Bool(v))
	}

	if v, ok := m["deployment_groups"]; ok {
		deploymentGroups, err := expandAWSGroupCodeDeployIntegrationDeploymentGroups(v, nullify)
		if err != nil {
			return nil, err
		}
		i.SetDeploymentGroups(deploymentGroups)
	}

	log.Printf("[DEBUG] Group CodeDeploy integration configuration: %s", stringutil.Stringify(i))
	return i, nil
}

func expandAWSGroupCodeDeployIntegrationDeploymentGroups(data interface{}, nullify bool) ([]*aws.DeploymentGroup, error) {
	list := data.(*schema.Set).List()
	deploymentGroups := make([]*aws.DeploymentGroup, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr["application_name"]; !ok {
			return nil, errors.New("invalid deployment group attributes: application_name missing")
		}

		if _, ok := attr["deployment_group_name"]; !ok {
			return nil, errors.New("invalid deployment group attributes: deployment_group_name missing")
		}
		deploymentGroup := &aws.DeploymentGroup{
			ApplicationName:     spotinst.String(attr["application_name"].(string)),
			DeploymentGroupName: spotinst.String(attr["deployment_group_name"].(string)),
		}
		deploymentGroups = append(deploymentGroups, deploymentGroup)
	}
	return deploymentGroups, nil
}

// expandAWSGroupElasticIPs expands the Elastic IPs block.
func expandAWSGroupPrivateIPs(data interface{}) ([]string, error) {
	list := data.([]interface{})
	result := make([]string, 0, len(list))
	for _, str := range list {
		if privateIP, ok := str.(string); ok {
			log.Printf("[DEBUG] Group private IP configuration: %s", stringutil.Stringify(privateIP))
			result = append(result, privateIP)
		}
	}
	return result, nil
}

// expandAWSGroupElasticIPs expands the Elastic IPs block.
func expandAWSGroupElasticIPs(data interface{}, nullify bool) ([]string, error) {
	list := data.([]interface{})
	eips := make([]string, 0, len(list))
	for _, str := range list {
		if eip, ok := str.(string); ok {
			log.Printf("[DEBUG] Group elastic IP configuration: %s", stringutil.Stringify(eip))
			eips = append(eips, eip)
		}
	}
	return eips, nil
}

// expandAWSGroupTags expands the Tags block.
func expandAWSGroupTags(data interface{}, nullify bool) ([]*aws.Tag, error) {
	list := data.(map[string]interface{})
	tags := make([]*aws.Tag, 0, len(list))
	for k, v := range list {
		tag := &aws.Tag{}
		tag.SetKey(spotinst.String(k))
		tag.SetValue(spotinst.String(v.(string)))
		log.Printf("[DEBUG] Group tag configuration: %s", stringutil.Stringify(tag))
		tags = append(tags, tag)
	}
	return tags, nil
}

// expandAWSGroupTagsKV expands the Tags KV block.
func expandAWSGroupTagsKV(data interface{}, nullify bool) ([]*aws.Tag, error) {
	list := data.(*schema.Set).List()
	tags := make([]*aws.Tag, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr["key"]; !ok {
			return nil, errors.New("invalid tag attributes: key missing")
		}

		if _, ok := attr["value"]; !ok {
			return nil, errors.New("invalid tag attributes: value missing")
		}
		tag := &aws.Tag{
			Key:   spotinst.String(attr["key"].(string)),
			Value: spotinst.String(attr["value"].(string)),
		}
		log.Printf("[DEBUG] Group tag configuration: %s", stringutil.Stringify(tag))
		tags = append(tags, tag)
	}
	return tags, nil
}

// expandAWSGroupInstanceTypeWeights expands the Instance type weights.
func expandAWSGroupInstanceTypeWeights(data interface{}, nullify bool) ([]*aws.InstanceTypeWeight, error) {
	list := data.(*schema.Set).List()
	weights := make([]*aws.InstanceTypeWeight, 0, len(list))
	for _, v := range list {
		attr, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := attr["instance_type"]; !ok {
			return nil, errors.New("invalid instance type weight: instance_type missing")
		}

		if _, ok := attr["weight"]; !ok {
			return nil, errors.New("invalid instance type weight: weight missing")
		}
		instance_weight := &aws.InstanceTypeWeight{}
		instance_weight.SetInstanceType(spotinst.String(attr["instance_type"].(string)))
		instance_weight.SetWeight(spotinst.Int(attr["weight"].(int)))
		log.Printf("[DEBUG] Group instance type weight configuration: %s", stringutil.Stringify(instance_weight))
		weights = append(weights, instance_weight)
	}
	return weights, nil
}

// expandAWSGroupRollConfig expands the Group Roll Configuration block.
func expandAWSGroupRollConfig(data interface{}, groupID string) (*aws.RollGroupInput, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &aws.RollGroupInput{GroupID: spotinst.String(groupID)}

	if v, ok := m["batch_size_percentage"].(int); ok { // Required value
		i.BatchSizePercentage = spotinst.Int(v)
	}

	if v, ok := m["grace_period"].(int); ok && v != -1 { // Default value set to -1
		i.GracePeriod = spotinst.Int(v)
	}

	if v, ok := m["health_check_type"].(string); ok && v != "" { // Default value ""
		i.HealthCheckType = spotinst.String(v)
	}

	log.Printf("[DEBUG] Group roll configuration: %s", stringutil.Stringify(i))
	return i, nil
}

//endregion

//region Hash methods

func hashAWSGroupCapacity(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%d-", m["target"].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m["minimum"].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m["maximum"].(int)))
	return hashcode.String(buf.String())
}

func hashAWSGroupStrategy(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%f-", m["risk"].(float64)))
	buf.WriteString(fmt.Sprintf("%d-", m["draining_timeout"].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m["ondemand_count"].(int)))
	buf.WriteString(fmt.Sprintf("%t-", m["utilize_reserved_instances"].(bool)))
	buf.WriteString(fmt.Sprintf("%t-", m["fallback_to_ondemand"].(bool)))
	buf.WriteString(fmt.Sprintf("%d-", m["spin_up_time"].(int)))
	buf.WriteString(fmt.Sprintf("%s-", m["availability_vs_cost"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["lifetime_period"].(string)))
	return hashcode.String(buf.String())
}

func hashAWSGroupPersistence(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%t-", m["persist_root_device"].(bool)))
	buf.WriteString(fmt.Sprintf("%t-", m["persist_block_devices"].(bool)))
	buf.WriteString(fmt.Sprintf("%t-", m["persist_private_ip"].(bool)))
	buf.WriteString(fmt.Sprintf("%s-", m["block_devices_mode"].(string)))
	return hashcode.String(buf.String())
}

func hashAWSGroupLoadBalancer(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", strings.ToLower(m["type"].(string))))
	if v, ok := m["name"].(string); ok && len(v) > 0 {
		buf.WriteString(fmt.Sprintf("%s-", m["name"].(string)))
	}
	if v, ok := m["arn"].(string); ok && len(v) > 0 {
		buf.WriteString(fmt.Sprintf("%s-", v))
	}
	if v, ok := m["balancer_id"].(string); ok && len(v) > 0 {
		buf.WriteString(fmt.Sprintf("%s-", v))
	}
	if v, ok := m["target_set_id"].(string); ok && len(v) > 0 {
		buf.WriteString(fmt.Sprintf("%s-", v))
	}
	if v, ok := m["zone_awareness"].(bool); ok {
		buf.WriteString(fmt.Sprintf("%t-", v))
	}
	if v, ok := m["auto_weight"].(bool); ok {
		buf.WriteString(fmt.Sprintf("%t-", v))
	}
	return hashcode.String(buf.String())
}

func hashAWSGroupEBSBlockDevice(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m["device_name"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["snapshot_id"].(string)))
	buf.WriteString(fmt.Sprintf("%d-", m["volume_size"].(int)))
	buf.WriteString(fmt.Sprintf("%t-", m["delete_on_termination"].(bool)))
	buf.WriteString(fmt.Sprintf("%t-", m["encrypted"].(bool)))
	buf.WriteString(fmt.Sprintf("%d-", m["iops"].(int)))
	return hashcode.String(buf.String())
}

func hashKV(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m["key"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["value"].(string)))
	return hashcode.String(buf.String())
}

//endregion

//region Helper methods

func hexStateFunc(v interface{}) string {
	switch s := v.(type) {
	case string:
		hash := sha1.Sum([]byte(s))
		return hex.EncodeToString(hash[:])
	default:
		return ""
	}
}

func attrStateFunc(v interface{}) string {
	switch s := v.(type) {
	case string:
		return fmt.Sprintf("${%s}", s)
	default:
		return ""
	}
}

// base64Encode encodes data if the input isn't already encoded using
// base64.StdEncoding.EncodeToString. If the input is already base64 encoded,
// return the original input unchanged.
func base64Encode(data string) string {
	// Check whether the data is already Base64 encoded; don't double-encode
	if isBase64Encoded(data) {
		return data
	}
	// data has not been encoded encode and return
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func isBase64Encoded(data string) bool {
	_, err := base64.StdEncoding.DecodeString(data)
	return err == nil
}

//endregion

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
	"github.com/spotinst/spotinst-sdk-go/spotinst"
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

			"capacity": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
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
							Optional: true,
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"arn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},

						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Set: hashAWSGroupLoadBalancer,
			},

			"launch_specification": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load_balancer_names": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
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
				Set: hashAWSGroupTagKV,
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

			"scaling_up_policy": scalingPolicySchema(),

			"scaling_down_policy": scalingPolicySchema(),

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
					},
				},
			},

			"kubernetes_integration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api_server": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"token": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
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

func scalingPolicySchema() *schema.Schema {
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

				"statistic": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},

				"unit": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},

				"threshold": &schema.Schema{
					Type:     schema.TypeFloat,
					Required: true,
				},

				"adjustment": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},

				"adjustment_expression": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},

				"min_target_capacity": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},

				"max_target_capacity": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},

				"namespace": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},

				"operator": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},

				"evaluation_periods": &schema.Schema{
					Type:     schema.TypeInt,
					Required: true,
				},

				"period": &schema.Schema{
					Type:     schema.TypeInt,
					Required: true,
				},

				"cooldown": &schema.Schema{
					Type:     schema.TypeInt,
					Required: true,
				},

				"dimensions": &schema.Schema{
					Type:     schema.TypeMap,
					Optional: true,
				},

				"minimum": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},

				"maximum": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},

				"target": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},

				"action_type": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
		Set: hashAWSGroupScalingPolicy,
	}
}

//region CRUD methods

func resourceSpotinstAWSGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*spotinst.Client)
	newAWSGroup, err := buildAWSGroupOpts(d, meta)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] AWSGroup create configuration: %s", stringutil.Stringify(newAWSGroup))
	input := &spotinst.CreateAWSGroupInput{Group: newAWSGroup}
	resp, err := client.GroupService.CloudProviderAWS().Create(context.Background(), input)
	if err != nil {
		return fmt.Errorf("Error creating group: %s", err)
	}
	d.SetId(spotinst.StringValue(resp.Group.ID))
	log.Printf("[INFO] AWSGroup created successfully: %s", d.Id())
	return resourceSpotinstAWSGroupRead(d, meta)
}

func resourceSpotinstAWSGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*spotinst.Client)
	input := &spotinst.ReadAWSGroupInput{GroupID: spotinst.String(d.Id())}
	resp, err := client.GroupService.CloudProviderAWS().Read(context.Background(), input)
	if err != nil {
		return fmt.Errorf("Error retrieving group: %s", err)
	}
	if g := resp.Group; g != nil {
		d.Set("name", g.Name)
		d.Set("description", g.Description)
		d.Set("product", g.Compute.Product)
		d.Set("elastic_ips", g.Compute.ElasticIPs)

		// Set capacity.
		if g.Capacity != nil {
			if err := d.Set("capacity", flattenAWSGroupCapacity(g.Capacity)); err != nil {
				return fmt.Errorf("Error setting capacity onfiguration: %#v", err)
			}
		}

		if g.Strategy != nil {
			if err := d.Set("strategy", flattenAWSGroupStrategy(g.Strategy)); err != nil {
				return fmt.Errorf("Error setting strategy configuration: %#v", err)
			}

			// Set signals.
			if g.Strategy.Signals != nil {
				if err := d.Set("signal", flattenAWSGroupSignals(g.Strategy.Signals)); err != nil {
					return fmt.Errorf("Error setting signals configuration: %#v", err)
				}
			} else {
				d.Set("signal", []*spotinst.AWSGroupScalingPolicy{})
			}

			if g.Strategy.Persistence != nil {
				if err := d.Set("persistence", flattenAWSGroupPersistence(g.Strategy.Persistence)); err != nil {
					return fmt.Errorf("Error setting persistence configuration: %#v", err)
				}
			}
		}

		if g.Scaling != nil {
			// Set scaling up policies.
			if g.Scaling.Up != nil {
				if err := d.Set("scaling_up_policy", flattenAWSGroupScalingPolicies(g.Scaling.Up)); err != nil {
					return fmt.Errorf("Error setting scaling up policies configuration: %#v", err)
				}
			} else {
				d.Set("scaling_up_policy", []*spotinst.AWSGroupScalingPolicy{})
			}

			// Set scaling down policies.
			if g.Scaling.Down != nil {
				if err := d.Set("scaling_down_policy", flattenAWSGroupScalingPolicies(g.Scaling.Down)); err != nil {
					return fmt.Errorf("Error setting scaling down policies configuration: %#v", err)
				}
			} else {
				d.Set("scaling_down_policy", []*spotinst.AWSGroupScalingPolicy{})
			}

		}

		if g.Scheduling != nil {
			// Set scheduled tasks.
			if g.Scheduling.Tasks != nil {
				if err := d.Set("scheduled_task", flattenAWSGroupScheduledTasks(g.Scheduling.Tasks)); err != nil {
					return fmt.Errorf("Error setting scheduled tasks configuration: %#v", err)
				}
			} else {
				d.Set("scheduled_task", []*spotinst.AWSGroupScheduledTask{})
			}

		}

		// Set launch specification.
		if g.Compute.LaunchSpecification != nil {
			// Check if image ID is set in launch spec
			imageIDSetInLaunchSpec := true
			if v, ok := d.GetOk("image_id"); ok && v != "" {
				imageIDSetInLaunchSpec = false
			}
			if err := d.Set("launch_specification", flattenAWSGroupLaunchSpecification(g.Compute.LaunchSpecification, imageIDSetInLaunchSpec)); err != nil {
				return fmt.Errorf("Error setting launch specification configuration: %#v", err)
			}
		}

		// Set image ID.
		if g.Compute.LaunchSpecification.ImageID != nil {
			if d.Get("image_id") != nil && d.Get("image_id") != "" {
				d.Set("image_id", g.Compute.LaunchSpecification.ImageID)
			}
		}

		// Set load balancers.
		if g.Compute.LaunchSpecification.LoadBalancersConfig != nil {
			if err := d.Set("load_balancer", flattenAWSGroupLoadBalancers(g.Compute.LaunchSpecification.LoadBalancersConfig.LoadBalancers)); err != nil {
				return fmt.Errorf("Error setting load balancers configuration: %#v", err)
			}
		} else {
			d.Set("load_balancer", []*spotinst.AWSGroupComputeLoadBalancer{})
		}

		// Set EBS volume pool.
		if g.Compute.EBSVolumePool != nil {
			if err := d.Set("hot_ebs_volume", flattenAWSGroupEBSVolumePool(g.Compute.EBSVolumePool)); err != nil {
				return fmt.Errorf("Error setting EBS volume pool configuration: %#v", err)
			}
		} else {
			d.Set("hot_ebs_volume", []*spotinst.AWSGroupComputeEBSVolume{})
		}

		// Set network interfaces.
		if g.Compute.LaunchSpecification.NetworkInterfaces != nil {
			if err := d.Set("network_interface", flattenAWSGroupNetworkInterfaces(g.Compute.LaunchSpecification.NetworkInterfaces)); err != nil {
				return fmt.Errorf("Error setting network interfaces configuration: %#v", err)
			}
		} else {
			d.Set("network_interface", []*spotinst.AWSGroupComputeNetworkInterface{})
		}

		// Set block devices.
		if g.Compute.LaunchSpecification.BlockDevices != nil {
			if err := d.Set("ebs_block_device", flattenAWSGroupEBSBlockDevices(g.Compute.LaunchSpecification.BlockDevices)); err != nil {
				return fmt.Errorf("Error setting EBS block devices configuration: %#v", err)
			}
			if err := d.Set("ephemeral_block_device", flattenAWSGroupEphemeralBlockDevices(g.Compute.LaunchSpecification.BlockDevices)); err != nil {
				return fmt.Errorf("Error setting Ephemeral block devices configuration: %#v", err)
			}
		} else {
			d.Set("ebs_block_device", []*spotinst.AWSGroupComputeBlockDevice{})
			d.Set("ephemeral_block_device", []*spotinst.AWSGroupComputeBlockDevice{})
		}

		if g.Integration != nil {
			// Set Rancher integration.
			if g.Integration.Rancher != nil {
				if err := d.Set("rancher_integration", flattenAWSGroupRancherIntegration(g.Integration.Rancher)); err != nil {
					return fmt.Errorf("Error setting Rancher configuration: %#v", err)
				}
			} else {
				d.Set("rancher_integration", []*spotinst.AWSGroupRancherIntegration{})
			}

			// Set Elastic Beanstalk integration.
			if g.Integration.ElasticBeanstalk != nil {
				if err := d.Set("elastic_beanstalk_integration", flattenAWSGroupElasticBeanstalkIntegration(g.Integration.ElasticBeanstalk)); err != nil {
					return fmt.Errorf("Error setting Elastic Beanstalk configuration: %#v", err)
				}
			} else {
				d.Set("elastic_beanstalk_integration", []*spotinst.AWSGroupElasticBeanstalkIntegration{})
			}

			// Set EC2 Container Service integration.
			if g.Integration.EC2ContainerService != nil {
				if err := d.Set("ec2_container_service_integration", flattenAWSGroupEC2ContainerServiceIntegration(g.Integration.EC2ContainerService)); err != nil {
					return fmt.Errorf("Error setting EC2 Container Service configuration: %#v", err)
				}
			} else {
				d.Set("ec2_container_service_integration", []*spotinst.AWSGroupEC2ContainerServiceIntegration{})
			}

			// Set Kubernetes integration.
			if g.Integration.Kubernetes != nil {
				if err := d.Set("kubernetes_integration", flattenAWSGroupKubernetesIntegration(g.Integration.Kubernetes)); err != nil {
					return fmt.Errorf("Error setting Kubernetes configuration: %#v", err)
				}
			} else {
				d.Set("kubernetes_integration", []*spotinst.AWSGroupKubernetesIntegration{})
			}

			// Set Mesosphere integration.
			if g.Integration.Mesosphere != nil {
				if err := d.Set("mesosphere_integration", flattenAWSGroupMesosphereIntegration(g.Integration.Mesosphere)); err != nil {
					return fmt.Errorf("Error setting Mesosphere configuration: %#v", err)
				}
			} else {
				d.Set("mesosphere_integration", []*spotinst.AWSGroupMesosphereIntegration{})
			}

			// Set CodeDeploy integration.
			if g.Integration.CodeDeploy != nil {
				if err := d.Set("codedeploy_integration", flattenAWSGroupCodeDeployIntegration(g.Integration.CodeDeploy)); err != nil {
					return fmt.Errorf("Error setting CodeDeploy configuration: %#v", err)
				}
			} else {
				d.Set("codedeploy_integration", []*spotinst.AWSGroupCodeDeployIntegration{})
			}
		}

	} else {
		d.SetId("")
	}
	return nil
}

func resourceSpotinstAWSGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*spotinst.Client)
	group := &spotinst.AWSGroup{}
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
			if capacity, err := expandAWSGroupCapacity(v, nullify); err != nil {
				return err
			} else {
				group.SetCapacity(capacity)
				update = true
			}
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
				group.SetCompute(&spotinst.AWSGroupCompute{})
			}
			group.Compute.SetLaunchSpecification(lc)
			update = true
		}
	}

	if d.HasChange("image_id") {
		if d.Get("image_id") != nil && d.Get("image_id") != "" {
			if group.Compute == nil {
				group.SetCompute(&spotinst.AWSGroupCompute{})
			}
			if group.Compute.LaunchSpecification == nil {
				group.Compute.SetLaunchSpecification(&spotinst.AWSGroupComputeLaunchSpecification{})
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
					group.SetCompute(&spotinst.AWSGroupCompute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&spotinst.AWSGroupComputeLaunchSpecification{})
				}
				if group.Compute.LaunchSpecification.LoadBalancersConfig == nil {
					group.Compute.LaunchSpecification.SetLoadBalancersConfig(&spotinst.AWSGroupComputeLoadBalancersConfig{})
					group.Compute.LaunchSpecification.LoadBalancersConfig.SetLoadBalancers(lbs)
					update = true
				}
			}
		} else {
			if group.Compute == nil {
				group.SetCompute(&spotinst.AWSGroupCompute{})
			}
			if group.Compute.LaunchSpecification == nil {
				group.Compute.SetLaunchSpecification(&spotinst.AWSGroupComputeLaunchSpecification{})
			}
			if group.Compute.LaunchSpecification.LoadBalancersConfig == nil {
				group.Compute.LaunchSpecification.SetLoadBalancersConfig(&spotinst.AWSGroupComputeLoadBalancersConfig{})
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
					group.SetCompute(&spotinst.AWSGroupCompute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&spotinst.AWSGroupComputeLaunchSpecification{})
				}
				if len(group.Compute.LaunchSpecification.BlockDevices) > 0 {
					group.Compute.LaunchSpecification.SetBlockDevices(append(group.Compute.LaunchSpecification.BlockDevices, devices...))
				} else {
					if v, ok := d.GetOk("ephemeral_block_device"); ok {
						if ephemeral, err := expandAWSGroupEphemeralBlockDevices(v, nullify); err != nil {
							return err
						} else {
							devices = append(devices, ephemeral...)
							blockDevicesExpanded = true
						}
					}
					group.Compute.LaunchSpecification.SetBlockDevices(devices)
				}
				update = true
			}
		} else {
			if group.Compute == nil {
				group.SetCompute(&spotinst.AWSGroupCompute{})
			}
			if group.Compute.LaunchSpecification == nil {
				group.Compute.SetLaunchSpecification(&spotinst.AWSGroupComputeLaunchSpecification{})
			}
			group.Compute.LaunchSpecification.SetBlockDevices(nil)
			update = true
		}
	}

	if d.HasChange("ephemeral_block_device") && !blockDevicesExpanded {
		if v, ok := d.GetOk("ephemeral_block_device"); ok {
			if devices, err := expandAWSGroupEphemeralBlockDevices(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&spotinst.AWSGroupCompute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&spotinst.AWSGroupComputeLaunchSpecification{})
				}
				if len(group.Compute.LaunchSpecification.BlockDevices) > 0 {
					group.Compute.LaunchSpecification.SetBlockDevices(append(group.Compute.LaunchSpecification.BlockDevices, devices...))
				} else {
					if v, ok := d.GetOk("ebs_block_device"); ok {
						if ebs, err := expandAWSGroupEBSBlockDevices(v, nullify); err != nil {
							return err
						} else {
							devices = append(devices, ebs...)
						}
					}
					group.Compute.LaunchSpecification.SetBlockDevices(devices)
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
					group.SetCompute(&spotinst.AWSGroupCompute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&spotinst.AWSGroupComputeLaunchSpecification{})
				}
				group.Compute.LaunchSpecification.SetNetworkInterfaces(interfaces)
				update = true
			}
		} else {
			if group.Compute == nil {
				group.SetCompute(&spotinst.AWSGroupCompute{})
			}
			if group.Compute.LaunchSpecification == nil {
				group.Compute.SetLaunchSpecification(&spotinst.AWSGroupComputeLaunchSpecification{})
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
					group.SetCompute(&spotinst.AWSGroupCompute{})
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
					group.SetCompute(&spotinst.AWSGroupCompute{})
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
					group.SetCompute(&spotinst.AWSGroupCompute{})
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
					group.SetStrategy(&spotinst.AWSGroupStrategy{})
				}
				group.Strategy.SetSignals(signals)
				update = true
			}
		} else {
			if group.Strategy == nil {
				group.SetStrategy(&spotinst.AWSGroupStrategy{})
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
					group.SetCompute(&spotinst.AWSGroupCompute{})
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
					group.SetCompute(&spotinst.AWSGroupCompute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&spotinst.AWSGroupComputeLaunchSpecification{})
				}
				group.Compute.LaunchSpecification.SetTags(tags)
				update = true
			}
		} else {
			if _, ok := d.GetOk("tags_kv"); !ok {
				if group.Compute == nil {
					group.SetCompute(&spotinst.AWSGroupCompute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&spotinst.AWSGroupComputeLaunchSpecification{})
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
					group.SetCompute(&spotinst.AWSGroupCompute{})
				}
				if group.Compute.LaunchSpecification == nil {
					group.Compute.SetLaunchSpecification(&spotinst.AWSGroupComputeLaunchSpecification{})
				}
				group.Compute.LaunchSpecification.SetTags(tags)
				update = true
			}
		} else {
			if group.Compute == nil {
				group.SetCompute(&spotinst.AWSGroupCompute{})
			}
			if group.Compute.LaunchSpecification == nil {
				group.Compute.SetLaunchSpecification(&spotinst.AWSGroupComputeLaunchSpecification{})
			}
			group.Compute.LaunchSpecification.SetTags(nil)
			update = true
		}
	}

	if d.HasChange("elastic_ips") {
		if v, ok := d.GetOk("elastic_ips"); ok {
			if eips, err := expandAWSGroupElasticIPs(v, nullify); err != nil {
				return err
			} else {
				if group.Compute == nil {
					group.SetCompute(&spotinst.AWSGroupCompute{})
				}
				group.Compute.SetElasticIPs(eips)
				update = true
			}
		} else {
			if group.Compute == nil {
				group.SetCompute(&spotinst.AWSGroupCompute{})
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
					group.SetScheduling(&spotinst.AWSGroupScheduling{})
				}
				group.Scheduling.SetTasks(tasks)
				update = true
			}
		} else {
			if group.Scheduling == nil {
				group.SetScheduling(&spotinst.AWSGroupScheduling{})
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
					group.SetStrategy(&spotinst.AWSGroupStrategy{})
				}
				if group.Strategy.Persistence == nil {
					group.Strategy.SetPersistence(&spotinst.AWSGroupPersistence{})
				}
				group.Strategy.SetPersistence(persistence)
				update = true
			}
		} else {
			if group.Strategy == nil {
				group.SetStrategy(&spotinst.AWSGroupStrategy{})
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
					group.SetScaling(&spotinst.AWSGroupScaling{})
				}
				group.Scaling.SetUp(policies)
				update = true
			}
		} else {
			if group.Scaling == nil {
				group.SetScaling(&spotinst.AWSGroupScaling{})
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
					group.SetScaling(&spotinst.AWSGroupScaling{})
				}
				group.Scaling.SetDown(policies)
				update = true
			}
		} else {
			if group.Scaling == nil {
				group.SetScaling(&spotinst.AWSGroupScaling{})
			}
			group.Scaling.SetDown(nil)
			update = true
		}
	}

	if d.HasChange("rancher_integration") {
		if v, ok := d.GetOk("rancher_integration"); ok {
			if integration, err := expandAWSGroupRancherIntegration(v, nullify); err != nil {
				return err
			} else {
				if group.Integration == nil {
					group.SetIntegration(&spotinst.AWSGroupIntegration{})
				}
				group.Integration.SetRancher(integration)
				update = true
			}
		} else {
			if group.Integration == nil {
				group.SetIntegration(&spotinst.AWSGroupIntegration{})
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
					group.SetIntegration(&spotinst.AWSGroupIntegration{})
				}
				group.Integration.SetElasticBeanstalk(integration)
				update = true
			}
		} else {
			if group.Integration == nil {
				group.SetIntegration(&spotinst.AWSGroupIntegration{})
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
					group.SetIntegration(&spotinst.AWSGroupIntegration{})
				}
				group.Integration.SetEC2ContainerService(integration)
				update = true
			}
		} else {
			if group.Integration == nil {
				group.SetIntegration(&spotinst.AWSGroupIntegration{})
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
					group.SetIntegration(&spotinst.AWSGroupIntegration{})
				}
				group.Integration.SetKubernetes(integration)
				update = true
			}
		} else {
			if group.Integration == nil {
				group.SetIntegration(&spotinst.AWSGroupIntegration{})
			}
			group.Integration.SetKubernetes(nil)
			update = true
		}
	}

	if d.HasChange("mesosphere_integration") {
		if v, ok := d.GetOk("mesosphere_integration"); ok {
			if integration, err := expandAWSGroupMesosphereIntegration(v, nullify); err != nil {
				return err
			} else {
				if group.Integration == nil {
					group.SetIntegration(&spotinst.AWSGroupIntegration{})
				}
				group.Integration.SetMesosphere(integration)
				update = true
			}
		} else {
			if group.Integration == nil {
				group.SetIntegration(&spotinst.AWSGroupIntegration{})
			}
			group.Integration.SetMesosphere(nil)
			update = true
		}
	}

	if d.HasChange("codedeploy_integration") {
		if v, ok := d.GetOk("codedeploy_integration"); ok {
			if integration, err := expandAWSGroupCodeDeployIntegration(v, nullify); err != nil {
				return err
			} else {
				if group.Integration == nil {
					group.SetIntegration(&spotinst.AWSGroupIntegration{})
				}
				group.Integration.SetCodeDeploy(integration)
				update = true
			}
		} else {
			if group.Integration == nil {
				group.SetIntegration(&spotinst.AWSGroupIntegration{})
			}
			group.Integration.SetCodeDeploy(nil)
			update = true
		}
	}

	if update {
		log.Printf("[DEBUG] AWSGroup update configuration: %s", stringutil.Stringify(group))
		input := &spotinst.UpdateAWSGroupInput{Group: group}
		if _, err := client.GroupService.CloudProviderAWS().Update(context.Background(), input); err != nil {
			return fmt.Errorf("Error updating group %s: %s", d.Id(), err)
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
						if _, err := client.GroupService.CloudProviderAWS().Roll(context.Background(), roll); err != nil {
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
	client := meta.(*spotinst.Client)
	log.Printf("[INFO] Deleting group: %s", d.Id())
	input := &spotinst.DeleteAWSGroupInput{GroupID: spotinst.String(d.Id())}
	if _, err := client.GroupService.CloudProviderAWS().Delete(context.Background(), input); err != nil {
		return fmt.Errorf("Error deleting group: %s", err)
	}
	d.SetId("")
	return nil
}

//endregion

//region Flatten methods

func flattenAWSGroupCapacity(capacity *spotinst.AWSGroupCapacity) []interface{} {
	result := make(map[string]interface{})
	result["target"] = spotinst.IntValue(capacity.Target)
	result["minimum"] = spotinst.IntValue(capacity.Minimum)
	result["maximum"] = spotinst.IntValue(capacity.Maximum)
	result["unit"] = spotinst.StringValue(capacity.Unit)
	return []interface{}{result}
}

func flattenAWSGroupStrategy(strategy *spotinst.AWSGroupStrategy) []interface{} {
	result := make(map[string]interface{})
	result["risk"] = spotinst.Float64Value(strategy.Risk)
	result["ondemand_count"] = spotinst.IntValue(strategy.OnDemandCount)
	result["availability_vs_cost"] = spotinst.StringValue(strategy.AvailabilityVsCost)
	result["draining_timeout"] = spotinst.IntValue(strategy.DrainingTimeout)
	result["utilize_reserved_instances"] = spotinst.BoolValue(strategy.UtilizeReservedInstances)
	result["fallback_to_ondemand"] = spotinst.BoolValue(strategy.FallbackToOnDemand)
	result["spin_up_time"] = spotinst.IntValue(strategy.SpinUpTime)
	return []interface{}{result}
}

func flattenAWSGroupLaunchSpecification(lspec *spotinst.AWSGroupComputeLaunchSpecification, includeImageID bool) []interface{} {
	result := make(map[string]interface{})
	result["health_check_grace_period"] = spotinst.IntValue(lspec.HealthCheckGracePeriod)
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

func flattenAWSGroupLoadBalancers(balancers []*spotinst.AWSGroupComputeLoadBalancer) []interface{} {
	result := make([]interface{}, 0, len(balancers))
	for _, b := range balancers {
		m := make(map[string]interface{})
		m["name"] = spotinst.StringValue(b.Name)
		m["arn"] = spotinst.StringValue(b.Arn)
		m["type"] = strings.ToLower(spotinst.StringValue(b.Type))
		result = append(result, m)
	}
	return result
}

func flattenAWSGroupEBSVolumePool(volumes []*spotinst.AWSGroupComputeEBSVolume) []interface{} {
	result := make([]interface{}, 0, len(volumes))
	for _, v := range volumes {
		m := make(map[string]interface{})
		m["device_name"] = spotinst.StringValue(v.DeviceName)
		m["volume_ids"] = v.VolumeIDs
		result = append(result, m)
	}
	return result
}

func flattenAWSGroupSignals(signals []*spotinst.AWSGroupStrategySignal) []interface{} {
	result := make([]interface{}, 0, len(signals))
	for _, s := range signals {
		m := make(map[string]interface{})
		m["name"] = strings.ToLower(spotinst.StringValue(s.Name))
		m["timeout"] = spotinst.IntValue(s.Timeout)
		result = append(result, m)
	}
	return result
}

func flattenAWSGroupScheduledTasks(tasks []*spotinst.AWSGroupScheduledTask) []interface{} {
	result := make([]interface{}, 0, len(tasks))
	for _, t := range tasks {
		m := make(map[string]interface{})
		m["is_enabled"] = spotinst.BoolValue(t.IsEnabled)
		m["task_type"] = spotinst.StringValue(t.TaskType)
		m["cron_expression"] = spotinst.StringValue(t.CronExpression)
		m["frequency"] = spotinst.StringValue(t.Frequency)
		m["scale_target_capacity"] = spotinst.IntValue(t.ScaleTargetCapacity)
		m["scale_min_capacity"] = spotinst.IntValue(t.ScaleMinCapacity)
		m["scale_max_capacity"] = spotinst.IntValue(t.ScaleMaxCapacity)
		m["batch_size_percentage"] = spotinst.IntValue(t.BatchSizePercentage)
		m["grace_period"] = spotinst.IntValue(t.GracePeriod)
		result = append(result, m)
	}
	return result
}

func flattenAWSGroupPersistence(persistence *spotinst.AWSGroupPersistence) []interface{} {
	result := make(map[string]interface{})
	result["persist_block_devices"] = spotinst.BoolValue(persistence.ShouldPersistBlockDevices)
	result["persist_private_ip"] = spotinst.BoolValue(persistence.ShouldPersistPrivateIp)
	result["persist_root_device"] = spotinst.BoolValue(persistence.ShouldPersistRootDevice)
	return []interface{}{result}
}

func flattenAWSGroupScalingPolicies(policies []*spotinst.AWSGroupScalingPolicy) []interface{} {
	result := make([]interface{}, 0, len(policies))
	for _, p := range policies {
		m := make(map[string]interface{})

		if p.Action != nil && p.Action.Type != nil {
			m["action_type"] = spotinst.StringValue(p.Action.Type)

			if adjInt, err := strconv.Atoi(spotinst.StringValue(p.Action.Adjustment)); err == nil {
				m["adjustment"] = adjInt
			} else {
				m["adjustment_expression"] = spotinst.StringValue(p.Action.Adjustment)
			}

			if mintcInt, err := strconv.Atoi(spotinst.StringValue(p.Action.MinTargetCapacity)); err == nil {
				m["min_target_capacity"] = mintcInt
			}

			if maxtcInt, err := strconv.Atoi(spotinst.StringValue(p.Action.MaxTargetCapacity)); err == nil {
				m["max_target_capacity"] = maxtcInt
			}

			m["minimum"] = spotinst.StringValue(p.Action.Minimum)
			m["maximum"] = spotinst.StringValue(p.Action.Maximum)
			m["target"] = spotinst.StringValue(p.Action.Target)
		} else {
			m["adjustment"] = spotinst.IntValue(p.Adjustment)
			m["min_target_capacity"] = spotinst.IntValue(p.MinTargetCapacity)
			m["max_target_capacity"] = spotinst.IntValue(p.MaxTargetCapacity)
		}

		m["cooldown"] = spotinst.IntValue(p.Cooldown)
		m["evaluation_periods"] = spotinst.IntValue(p.EvaluationPeriods)
		m["metric_name"] = spotinst.StringValue(p.MetricName)
		m["namespace"] = spotinst.StringValue(p.Namespace)
		m["operator"] = spotinst.StringValue(p.Operator)
		m["period"] = spotinst.IntValue(p.Period)
		m["policy_name"] = spotinst.StringValue(p.PolicyName)
		m["statistic"] = spotinst.StringValue(p.Statistic)
		m["threshold"] = spotinst.Float64Value(p.Threshold)
		m["unit"] = spotinst.StringValue(p.Unit)
		if len(p.Dimensions) > 0 {
			flatDims := make(map[string]interface{})
			for _, d := range p.Dimensions {
				flatDims[spotinst.StringValue(d.Name)] = *d.Value
			}
			m["dimensions"] = flatDims
		}
		result = append(result, m)
	}
	return result
}

func flattenAWSGroupNetworkInterfaces(ifaces []*spotinst.AWSGroupComputeNetworkInterface) []interface{} {
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

func flattenAWSGroupEBSBlockDevices(devices []*spotinst.AWSGroupComputeBlockDevice) []interface{} {
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

func flattenAWSGroupEphemeralBlockDevices(devices []*spotinst.AWSGroupComputeBlockDevice) []interface{} {
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

func flattenAWSGroupRancherIntegration(integration *spotinst.AWSGroupRancherIntegration) []interface{} {
	result := make(map[string]interface{})
	result["master_host"] = spotinst.StringValue(integration.MasterHost)
	result["access_key"] = spotinst.StringValue(integration.AccessKey)
	result["secret_key"] = spotinst.StringValue(integration.SecretKey)
	return []interface{}{result}
}

func flattenAWSGroupElasticBeanstalkIntegration(integration *spotinst.AWSGroupElasticBeanstalkIntegration) []interface{} {
	result := make(map[string]interface{})
	result["environment_id"] = spotinst.StringValue(integration.EnvironmentID)
	return []interface{}{result}
}

func flattenAWSGroupEC2ContainerServiceIntegration(integration *spotinst.AWSGroupEC2ContainerServiceIntegration) []interface{} {
	result := make(map[string]interface{})
	result["cluster_name"] = spotinst.StringValue(integration.ClusterName)
	if integration.AutoScale != nil {
		result["autoscale_is_enabled"] = spotinst.BoolValue(integration.AutoScale.IsEnabled)
		result["autoscale_cooldown"] = spotinst.IntValue(integration.AutoScale.Cooldown)
	}
	return []interface{}{result}
}

func flattenAWSGroupKubernetesIntegration(integration *spotinst.AWSGroupKubernetesIntegration) []interface{} {
	result := make(map[string]interface{})
	result["api_server"] = spotinst.StringValue(integration.Server)
	result["token"] = spotinst.StringValue(integration.Token)
	return []interface{}{result}
}

func flattenAWSGroupMesosphereIntegration(integration *spotinst.AWSGroupMesosphereIntegration) []interface{} {
	result := make(map[string]interface{})
	result["api_server"] = spotinst.StringValue(integration.Server)
	return []interface{}{result}
}

func flattenAWSGroupCodeDeployIntegration(integration *spotinst.AWSGroupCodeDeployIntegration) []interface{} {
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

/* buildAWSGroupOpts builds the Spotinst AWS Group options.*/
func buildAWSGroupOpts(d *schema.ResourceData, meta interface{}) (*spotinst.AWSGroup, error) {
	group := &spotinst.AWSGroup{
		Scaling:     &spotinst.AWSGroupScaling{},
		Scheduling:  &spotinst.AWSGroupScheduling{},
		Integration: &spotinst.AWSGroupIntegration{},
		Compute: &spotinst.AWSGroupCompute{
			LaunchSpecification: &spotinst.AWSGroupComputeLaunchSpecification{},
		},
	}
	nullify := false

	group.SetName(spotinst.String(d.Get("name").(string)))
	group.SetDescription(spotinst.String(d.Get("description").(string)))
	group.Compute.SetProduct(spotinst.String(d.Get("product").(string)))

	if v, ok := d.GetOk("capacity"); ok {
		if capacity, err := expandAWSGroupCapacity(v, nullify); err != nil {
			return nil, err
		} else {
			group.SetCapacity(capacity)
		}
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
				group.Compute.LaunchSpecification.LoadBalancersConfig = &spotinst.AWSGroupComputeLoadBalancersConfig{}
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
			group.Compute.LaunchSpecification.SetBlockDevices(devices)
		}
	}

	if v, ok := d.GetOk("ephemeral_block_device"); ok {
		if devices, err := expandAWSGroupEphemeralBlockDevices(v, nullify); err != nil {
			return nil, err
		} else {
			if len(group.Compute.LaunchSpecification.BlockDevices) > 0 {
				all := append(group.Compute.LaunchSpecification.BlockDevices, devices...)
				group.Compute.LaunchSpecification.SetBlockDevices(all)
			} else {
				group.Compute.LaunchSpecification.SetBlockDevices(devices)
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

	if v, ok := d.GetOk("mesosphere_integration"); ok {
		if integration, err := expandAWSGroupMesosphereIntegration(v, nullify); err != nil {
			return nil, err
		} else {
			group.Integration.SetMesosphere(integration)
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

/* expandAWSGroupCapacity expands the Capacity block.*/
func expandAWSGroupCapacity(data interface{}, nullify bool) (*spotinst.AWSGroupCapacity, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	capacity := &spotinst.AWSGroupCapacity{}

	if v, ok := m["minimum"].(int); ok && v >= 0 {
		capacity.SetMinimum(spotinst.Int(v))
	}

	if v, ok := m["maximum"].(int); ok && v >= 0 {
		capacity.SetMaximum(spotinst.Int(v))
	}

	if v, ok := m["target"].(int); ok && v >= 0 {
		capacity.SetTarget(spotinst.Int(v))
	}

	if v, ok := m["unit"].(string); ok && v != "" {
		capacity.SetUnit(spotinst.String(v))
	}

	log.Printf("[DEBUG] AWSGroup capacity configuration: %s", stringutil.Stringify(capacity))
	return capacity, nil
}

// expandAWSGroupStrategy expands the Strategy block.
func expandAWSGroupStrategy(data interface{}, nullify bool) (*spotinst.AWSGroupStrategy, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	strategy := &spotinst.AWSGroupStrategy{}

	if v, ok := m["risk"].(float64); ok && v >= 0 {
		strategy.SetRisk(spotinst.Float64(v))
	}

	if v, ok := m["ondemand_count"].(int); ok && v > 0 {
		strategy.SetOnDemandCount(spotinst.Int(v))
	}

	if v, ok := m["availability_vs_cost"].(string); ok && v != "" {
		strategy.SetAvailabilityVsCost(spotinst.String(v))
	}

	if v, ok := m["draining_timeout"].(int); ok && v > 0 {
		strategy.SetDrainingTimeout(spotinst.Int(v))
	}

	if v, ok := m["utilize_reserved_instances"].(bool); ok {
		strategy.SetUtilizeReservedInstances(spotinst.Bool(v))
	} else if nullify {
		strategy.SetUtilizeReservedInstances(nil)
	}

	if v, ok := m["fallback_to_ondemand"].(bool); ok {
		strategy.SetFallbackToOnDemand(spotinst.Bool(v))
	} else if nullify {
		strategy.SetFallbackToOnDemand(nil)
	}

	if v, ok := m["spin_up_time"].(int); ok && v > 0 {
		strategy.SetSpinUpTime(spotinst.Int(v))
	}

	log.Printf("[DEBUG] AWSGroup strategy configuration: %s", stringutil.Stringify(strategy))
	return strategy, nil
}

// expandAWSGroupScalingPolicies expands the Scaling Policy block.
func expandAWSGroupScalingPolicies(data interface{}, nullify bool) ([]*spotinst.AWSGroupScalingPolicy, error) {
	list := data.(*schema.Set).List()
	policies := make([]*spotinst.AWSGroupScalingPolicy, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		policy := &spotinst.AWSGroupScalingPolicy{}

		if v, ok := m["policy_name"].(string); ok && v != "" {
			policy.SetPolicyName(spotinst.String(v))
		}

		if v, ok := m["metric_name"].(string); ok && v != "" {
			policy.SetMetricName(spotinst.String(v))
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

		if v, ok := m["namespace"].(string); ok && v != "" {
			policy.SetNamespace(spotinst.String(v))
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
			action := &spotinst.AWSGroupScalingPolicyAction{}
			action.SetType(spotinst.String(v))

			if v, ok := m["adjustment"].(int); ok && v > 0 {
				vStr := strconv.Itoa(v)
				action.SetAdjustment(spotinst.String(vStr))
			} else if v, ok := m["adjustment_expression"].(string); ok && v != "" {
				action.SetAdjustment(spotinst.String(v))
			}

			if v, ok := m["min_target_capacity"].(int); ok && v > 0 {
				vStr := strconv.Itoa(v)
				action.SetMinTargetCapacity(spotinst.String(vStr))
			}

			if v, ok := m["max_target_capacity"].(int); ok && v > 0 {
				vStr := strconv.Itoa(v)
				action.SetMaxTargetCapacity(spotinst.String(vStr))
			}

			if v, ok := m["minimum"].(string); ok && v != "" {
				action.SetMinimum(spotinst.String(v))
			}

			if v, ok := m["maximum"].(string); ok && v != "" {
				action.SetMaximum(spotinst.String(v))
			}

			if v, ok := m["target"].(string); ok && v != "" {
				action.SetTarget(spotinst.String(v))
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

		if v, ok := m["namespace"].(string); ok && v != "" {
			log.Printf("[DEBUG] AWSGroup scaling policy configuration: %s", stringutil.Stringify(policy))
			policies = append(policies, policy)
		}
	}

	return policies, nil
}

func expandAWSGroupScalingPolicyDimensions(list map[string]interface{}) []*spotinst.AWSGroupScalingPolicyDimension {
	dimensions := make([]*spotinst.AWSGroupScalingPolicyDimension, 0, len(list))
	for name, val := range list {
		dimension := &spotinst.AWSGroupScalingPolicyDimension{}
		dimension.SetName(spotinst.String(name))
		dimension.SetValue(spotinst.String(val.(string)))
		log.Printf("[DEBUG] AWSGroup scaling policy dimension: %s", stringutil.Stringify(dimension))
		dimensions = append(dimensions, dimension)
	}
	return dimensions
}

// expandAWSGroupScheduledTasks expands the Scheduled Task block.
func expandAWSGroupScheduledTasks(data interface{}, nullify bool) ([]*spotinst.AWSGroupScheduledTask, error) {
	list := data.(*schema.Set).List()
	tasks := make([]*spotinst.AWSGroupScheduledTask, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		task := &spotinst.AWSGroupScheduledTask{}

		if v, ok := m["is_enabled"].(bool); ok {
			task.SetIsEnabled(spotinst.Bool(v))
		}

		if v, ok := m["task_type"].(string); ok && v != "" {
			task.SetTaskType(spotinst.String(v))
		}

		if v, ok := m["frequency"].(string); ok && v != "" {
			task.SetFrequency(spotinst.String(v))
		}

		if v, ok := m["cron_expression"].(string); ok && v != "" {
			task.SetCronExpression(spotinst.String(v))
		}

		if v, ok := m["scale_target_capacity"].(int); ok && v >= 0 {
			task.SetScaleTargetCapacity(spotinst.Int(v))
		}

		if v, ok := m["scale_min_capacity"].(int); ok && v >= 0 {
			task.SetScaleMinCapacity(spotinst.Int(v))
		}

		if v, ok := m["scale_max_capacity"].(int); ok && v >= 0 {
			task.SetScaleMaxCapacity(spotinst.Int(v))
		}

		if v, ok := m["batch_size_percentage"].(int); ok && v > 0 {
			task.SetBatchSizePercentage(spotinst.Int(v))
		}

		if v, ok := m["grace_period"].(int); ok && v > 0 {
			task.SetGracePeriod(spotinst.Int(v))
		}

		log.Printf("[DEBUG] AWSGroup scheduled task configuration: %s", stringutil.Stringify(task))
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// expandAWSGroupPersistence expands the Persistence block.
func expandAWSGroupPersistence(data interface{}, nullify bool) (*spotinst.AWSGroupPersistence, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	persistence := &spotinst.AWSGroupPersistence{}

	if v, ok := m["persist_private_ip"].(bool); ok {
		persistence.SetShouldPersistPrivateIp(spotinst.Bool(v))
	} else if nullify {
		persistence.SetShouldPersistPrivateIp(nil)
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

	log.Printf("[DEBUG] AWSGroup persistence configuration: %s", stringutil.Stringify(persistence))
	return persistence, nil
}

// expandAWSGroupAvailabilityZones expands the Availability Zone block.
func expandAWSGroupAvailabilityZones(data interface{}, nullify bool) ([]*spotinst.AWSGroupComputeAvailabilityZone, error) {
	list := data.(*schema.Set).List()
	zones := make([]*spotinst.AWSGroupComputeAvailabilityZone, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		zone := &spotinst.AWSGroupComputeAvailabilityZone{}

		if v, ok := m["name"].(string); ok && v != "" {
			zone.SetName(spotinst.String(v))
		}

		if v, ok := m["subnet_id"].(string); ok && v != "" {
			zone.SetSubnetId(spotinst.String(v))
		}

		log.Printf("[DEBUG] AWSGroup availability zone configuration: %s", stringutil.Stringify(zone))
		zones = append(zones, zone)
	}

	return zones, nil
}

// expandAWSGroupAvailabilityZonesSlice expands the Availability Zone block when provided as a slice.
func expandAWSGroupAvailabilityZonesSlice(data interface{}, nullify bool) ([]*spotinst.AWSGroupComputeAvailabilityZone, error) {
	list := data.([]interface{})
	zones := make([]*spotinst.AWSGroupComputeAvailabilityZone, 0, len(list))
	for _, str := range list {
		if s, ok := str.(string); ok {
			parts := strings.Split(s, ":")
			zone := &spotinst.AWSGroupComputeAvailabilityZone{}
			if len(parts) >= 1 && parts[0] != "" {
				zone.SetName(spotinst.String(parts[0]))
			}
			if len(parts) == 2 && parts[1] != "" {
				zone.SetSubnetId(spotinst.String(parts[1]))
			}
			log.Printf("[DEBUG] AWSGroup availability zone configuration: %s", stringutil.Stringify(zone))
			zones = append(zones, zone)
		}
	}

	return zones, nil
}

// expandAWSGroupEBSVolumePool expands the EBS Volume Pool block.
func expandAWSGroupEBSVolumePool(data interface{}, nullify bool) ([]*spotinst.AWSGroupComputeEBSVolume, error) {
	list := data.(*schema.Set).List()
	volumes := make([]*spotinst.AWSGroupComputeEBSVolume, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		volume := &spotinst.AWSGroupComputeEBSVolume{}

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
			log.Printf("[DEBUG] AWSGroup EBS volume (pool) configuration: %s", stringutil.Stringify(volume))
			volumes = append(volumes, volume)
		}
	}

	return volumes, nil
}

// expandAWSGroupSignals expands the Signal block.
func expandAWSGroupSignals(data interface{}, nullify bool) ([]*spotinst.AWSGroupStrategySignal, error) {
	list := data.(*schema.Set).List()
	signals := make([]*spotinst.AWSGroupStrategySignal, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		signal := &spotinst.AWSGroupStrategySignal{}

		if v, ok := m["name"].(string); ok && v != "" {
			signal.SetName(spotinst.String(strings.ToUpper(v)))
		}

		if v, ok := m["timeout"].(int); ok && v > 0 {
			signal.SetTimeout(spotinst.Int(v))
		}

		log.Printf("[DEBUG] AWSGroup signal configuration: %s", stringutil.Stringify(signal))
		signals = append(signals, signal)
	}

	return signals, nil
}

// expandAWSGroupInstanceTypes expands the Instance Types block.
func expandAWSGroupInstanceTypes(data interface{}, nullify bool) (*spotinst.AWSGroupComputeInstanceType, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	types := &spotinst.AWSGroupComputeInstanceType{}
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

	log.Printf("[DEBUG] AWSGroup instance types configuration: %s", stringutil.Stringify(types))
	return types, nil
}

// expandAWSGroupNetworkInterfaces expands the Elastic Network Interface block.
func expandAWSGroupNetworkInterfaces(data interface{}, nullify bool) ([]*spotinst.AWSGroupComputeNetworkInterface, error) {
	list := data.(*schema.Set).List()
	interfaces := make([]*spotinst.AWSGroupComputeNetworkInterface, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		iface := &spotinst.AWSGroupComputeNetworkInterface{}

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

		log.Printf("[DEBUG] AWSGroup network interface configuration: %s", stringutil.Stringify(iface))
		interfaces = append(interfaces, iface)
	}

	return interfaces, nil
}

// expandAWSGroupEphemeralBlockDevice expands the Ephemeral Block Device block.
func expandAWSGroupEphemeralBlockDevices(data interface{}, nullify bool) ([]*spotinst.AWSGroupComputeBlockDevice, error) {
	list := data.(*schema.Set).List()
	devices := make([]*spotinst.AWSGroupComputeBlockDevice, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		device := &spotinst.AWSGroupComputeBlockDevice{}

		if v, ok := m["device_name"].(string); ok && v != "" {
			device.SetDeviceName(spotinst.String(v))
		}

		if v, ok := m["virtual_name"].(string); ok && v != "" {
			device.SetVirtualName(spotinst.String(v))
		}

		log.Printf("[DEBUG] AWSGroup ephemeral block device configuration: %s", stringutil.Stringify(device))
		devices = append(devices, device)
	}

	return devices, nil
}

// expandAWSGroupEBSBlockDevices expands the EBS Block Device block.
func expandAWSGroupEBSBlockDevices(data interface{}, nullify bool) ([]*spotinst.AWSGroupComputeBlockDevice, error) {
	list := data.(*schema.Set).List()
	devices := make([]*spotinst.AWSGroupComputeBlockDevice, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		device := &spotinst.AWSGroupComputeBlockDevice{EBS: &spotinst.AWSGroupComputeEBS{}}

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

		log.Printf("[DEBUG] AWSGroup elastic block device configuration: %s", stringutil.Stringify(device))
		devices = append(devices, device)
	}

	return devices, nil
}

// iprofArnRE is a regular expression for matching IAM instance profile ARNs.
var iprofArnRE = regexp.MustCompile(`arn:aws:iam::\d{12}:instance-profile/?[a-zA-Z_0-9+=,.@\-_/]+`)

// expandAWSGroupLaunchSpecification expands the launch Specification block.
func expandAWSGroupLaunchSpecification(data interface{}, nullify bool) (*spotinst.AWSGroupComputeLaunchSpecification, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	lc := &spotinst.AWSGroupComputeLaunchSpecification{}

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
		iprof := &spotinst.AWSGroupComputeIAMInstanceProfile{}
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

	log.Printf("[DEBUG] AWSGroup launch specification configuration: %s", stringutil.Stringify(lc))
	return lc, nil
}

// expandAWSGroupLoadBalancer expands the Load Balancer block.
func expandAWSGroupLoadBalancer(data interface{}, nullify bool) ([]*spotinst.AWSGroupComputeLoadBalancer, error) {
	list := data.(*schema.Set).List()
	lbs := make([]*spotinst.AWSGroupComputeLoadBalancer, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		lb := &spotinst.AWSGroupComputeLoadBalancer{}

		if v, ok := m["name"].(string); ok && v != "" {
			lb.SetName(spotinst.String(v))
		}

		if v, ok := m["arn"].(string); ok && v != "" {
			lb.SetArn(spotinst.String(v))
		}

		if v, ok := m["type"].(string); ok && v != "" {
			lb.SetType(spotinst.String(strings.ToUpper(v)))
		}

		log.Printf("[DEBUG] AWSGroup load balancer configuration: %s", stringutil.Stringify(lb))
		lbs = append(lbs, lb)
	}

	return lbs, nil
}

// expandAWSGroupRancherIntegration expands the Rancher Integration block.
func expandAWSGroupRancherIntegration(data interface{}, nullify bool) (*spotinst.AWSGroupRancherIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &spotinst.AWSGroupRancherIntegration{}

	if v, ok := m["master_host"].(string); ok && v != "" {
		i.SetMasterHost(spotinst.String(v))
	}

	if v, ok := m["access_key"].(string); ok && v != "" {
		i.SetAccessKey(spotinst.String(v))
	}

	if v, ok := m["secret_key"].(string); ok && v != "" {
		i.SetSecretKey(spotinst.String(v))
	}

	log.Printf("[DEBUG] AWSGroup Rancher integration configuration: %s", stringutil.Stringify(i))
	return i, nil
}

// expandAWSGroupElasticBeanstalkIntegration expands the Elastic Beanstalk Integration block.
func expandAWSGroupElasticBeanstalkIntegration(data interface{}, nullify bool) (*spotinst.AWSGroupElasticBeanstalkIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &spotinst.AWSGroupElasticBeanstalkIntegration{}

	if v, ok := m["environment_id"].(string); ok && v != "" {
		i.SetEnvironmentId(spotinst.String(v))
	}

	log.Printf("[DEBUG] AWSGroup Elastic Beanstalk integration configuration:  %s", stringutil.Stringify(i))
	return i, nil
}

// expandAWSGroupEC2ContainerServiceIntegration expands the EC2 Container Service Integration block.
func expandAWSGroupEC2ContainerServiceIntegration(data interface{}, nullify bool) (*spotinst.AWSGroupEC2ContainerServiceIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &spotinst.AWSGroupEC2ContainerServiceIntegration{}

	if v, ok := m["cluster_name"].(string); ok && v != "" {
		i.SetClusterName(spotinst.String(v))
	}

	if v, ok := m["autoscale_is_enabled"].(bool); ok {
		if i.AutoScale == nil {
			i.SetAutoScale(&spotinst.AWSGroupEC2ContainerServiceIntegrationAutoScale{})
		}
		i.AutoScale.SetIsEnabled(spotinst.Bool(v))
	}

	if v, ok := m["autoscale_cooldown"].(int); ok && v > 0 {
		if i.AutoScale == nil {
			i.SetAutoScale(&spotinst.AWSGroupEC2ContainerServiceIntegrationAutoScale{})
		}
		i.AutoScale.SetCooldown(spotinst.Int(v))
	}

	log.Printf("[DEBUG] AWSGroup ECS integration configuration:  %s", stringutil.Stringify(i))
	return i, nil
}

// expandAWSGroupKubernetesIntegration expands the Kubernetes Integration block.
func expandAWSGroupKubernetesIntegration(data interface{}, nullify bool) (*spotinst.AWSGroupKubernetesIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &spotinst.AWSGroupKubernetesIntegration{}

	if v, ok := m["api_server"].(string); ok && v != "" {
		i.SetServer(spotinst.String(v))
	}

	if v, ok := m["token"].(string); ok && v != "" {
		i.SetToken(spotinst.String(v))
	}

	log.Printf("[DEBUG] AWSGroup Kubernetes integration configuration:  %s", stringutil.Stringify(i))
	return i, nil
}

// expandAWSGroupMesosphereIntegration expands the Mesosphere Integration block.
func expandAWSGroupMesosphereIntegration(data interface{}, nullify bool) (*spotinst.AWSGroupMesosphereIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &spotinst.AWSGroupMesosphereIntegration{}

	if v, ok := m["api_server"].(string); ok && v != "" {
		i.SetServer(spotinst.String(v))
	}

	log.Printf("[DEBUG] AWSGroup Mesosphere integration configuration: %s", stringutil.Stringify(i))
	return i, nil
}

// expandAWSGroupCodeDeployIntegration expands the CodeDeploy Integration block.
func expandAWSGroupCodeDeployIntegration(data interface{}, nullify bool) (*spotinst.AWSGroupCodeDeployIntegration, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &spotinst.AWSGroupCodeDeployIntegration{}

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

	log.Printf("[DEBUG] AWSGroup CodeDeploy integration configuration: %s", stringutil.Stringify(i))
	return i, nil
}

func expandAWSGroupCodeDeployIntegrationDeploymentGroups(data interface{}, nullify bool) ([]*spotinst.AWSGroupCodeDeployIntegrationDeploymentGroup, error) {
	list := data.(*schema.Set).List()
	deploymentGroups := make([]*spotinst.AWSGroupCodeDeployIntegrationDeploymentGroup, 0, len(list))
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
		deploymentGroup := &spotinst.AWSGroupCodeDeployIntegrationDeploymentGroup{
			ApplicationName:     spotinst.String(attr["application_name"].(string)),
			DeploymentGroupName: spotinst.String(attr["deployment_group_name"].(string)),
		}
		deploymentGroups = append(deploymentGroups, deploymentGroup)
	}
	return deploymentGroups, nil
}

// expandAWSGroupElasticIPs expands the Elastic IPs block.
func expandAWSGroupElasticIPs(data interface{}, nullify bool) ([]string, error) {
	list := data.([]interface{})
	eips := make([]string, 0, len(list))
	for _, str := range list {
		if eip, ok := str.(string); ok {
			log.Printf("[DEBUG] AWSGroup elastic IP configuration: %s", stringutil.Stringify(eip))
			eips = append(eips, eip)
		}
	}
	return eips, nil
}

// expandAWSGroupTags expands the Tags block.
func expandAWSGroupTags(data interface{}, nullify bool) ([]*spotinst.AWSGroupComputeTag, error) {
	list := data.(map[string]interface{})
	tags := make([]*spotinst.AWSGroupComputeTag, 0, len(list))
	for k, v := range list {
		tag := &spotinst.AWSGroupComputeTag{}
		tag.SetKey(spotinst.String(k))
		tag.SetValue(spotinst.String(v.(string)))
		log.Printf("[DEBUG] AWSGroup tag configuration: %s", stringutil.Stringify(tag))
		tags = append(tags, tag)
	}
	return tags, nil
}

// expandAWSGroupTagsKV expands the Tags KV block.
func expandAWSGroupTagsKV(data interface{}, nullify bool) ([]*spotinst.AWSGroupComputeTag, error) {
	list := data.(*schema.Set).List()
	tags := make([]*spotinst.AWSGroupComputeTag, 0, len(list))
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
		tag := &spotinst.AWSGroupComputeTag{
			Key:   spotinst.String(attr["key"].(string)),
			Value: spotinst.String(attr["value"].(string)),
		}
		log.Printf("[DEBUG] AWSGroup tag configuration: %s", stringutil.Stringify(tag))
		tags = append(tags, tag)
	}
	return tags, nil
}

// expandAWSGroupRollConfig expands the Group Roll Configuration block.
func expandAWSGroupRollConfig(data interface{}, groupID string) (*spotinst.RollAWSGroupInput, error) {
	list := data.(*schema.Set).List()
	m := list[0].(map[string]interface{})
	i := &spotinst.RollAWSGroupInput{GroupID: spotinst.String(groupID)}

	if v, ok := m["batch_size_percentage"].(int); ok { // Required value
		i.BatchSizePercentage = spotinst.Int(v)
	}

	if v, ok := m["grace_period"].(int); ok && v != -1 { // Default value set to -1
		i.GracePeriod = spotinst.Int(v)
	}

	if v, ok := m["health_check_type"].(string); ok && v != "" { // Default value ""
		i.HealthCheckType = spotinst.String(v)
	}

	log.Printf("[DEBUG] AWSGroup roll configuration: %s", stringutil.Stringify(i))
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
	return hashcode.String(buf.String())
}

func hashAWSGroupPersistence(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%t-", m["persist_root_device"].(bool)))
	buf.WriteString(fmt.Sprintf("%t-", m["persist_block_devices"].(bool)))
	buf.WriteString(fmt.Sprintf("%t-", m["persist_private_ip"].(bool)))
	return hashcode.String(buf.String())
}

func hashAWSGroupLoadBalancer(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-", m["name"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["type"].(string)))
	if v, ok := m["arn"].(string); ok && len(v) > 0 {
		buf.WriteString(fmt.Sprintf("%s-", v))
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

func hashAWSGroupScalingPolicy(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	buf.WriteString(fmt.Sprintf("%d-", m["cooldown"].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m["evaluation_periods"].(int)))
	buf.WriteString(fmt.Sprintf("%s-", m["metric_name"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["namespace"].(string)))
	buf.WriteString(fmt.Sprintf("%d-", m["period"].(int)))
	buf.WriteString(fmt.Sprintf("%s-", m["policy_name"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["statistic"].(string)))
	buf.WriteString(fmt.Sprintf("%f-", m["threshold"].(float64)))
	buf.WriteString(fmt.Sprintf("%s-", m["unit"].(string)))
	buf.WriteString(fmt.Sprintf("%d-", m["min_target_capacity"].(int)))
	buf.WriteString(fmt.Sprintf("%d-", m["max_target_capacity"].(int)))
	buf.WriteString(fmt.Sprintf("%s-", m["action_type"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["target"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["minimum"].(string)))
	buf.WriteString(fmt.Sprintf("%s-", m["maximum"].(string)))

	if v, ok := m["adjustment_expression"].(string); ok && v != "" {
		buf.WriteString(fmt.Sprintf("%s-", m["adjustment_expression"].(string)))
	} else {
		buf.WriteString(fmt.Sprintf("%d-", m["adjustment"].(int)))
	}

	if d, ok := m["dimensions"]; ok {
		if len(d.(map[string]interface{})) > 0 {
			e := d.(map[string]interface{})
			for k, v := range e {
				buf.WriteString(fmt.Sprintf("%s:%s-", k, v.(string)))
			}
		}
	}

	return hashcode.String(buf.String())
}

func hashAWSGroupTagKV(v interface{}) int {
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

package spotinst

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/mrscaler"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
)

// NOTE:
// Currently the values set to default -1 are ones which have to be allowed to
// be optional as they are either optional or not always set. These are all
// target min and max values. Each of these must be validated to be != -1 upon
// creation and update. IOPS has default 0 as it cannot be otherwise anyway.

//region Resource

func resourceSpotinstAWSMrScaler() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpotinstAWSMrScalerCreate,
		Read:   resourceSpotinstAWSMrScalerRead,
		Update: resourceSpotinstAWSMrScalerUpdate,
		Delete: resourceSpotinstAWSMrScalerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"description": {
				Type:     schema.TypeString,
				Required: true,
			},

			"region": {
				Type:     schema.TypeString,
				Required: true,
			},

			"strategy": {
				Type:     schema.TypeString,
				Required: true,
			},

			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"master_instance_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"master_target": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},

			"master_lifecycle": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"master_ebs_optimized": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"master_ebs_block_device": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volumes_per_instance": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1,
						},

						"volume_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						"size_in_gb": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"iops": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
					},
				},
			},

			"core_instance_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"core_target": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},

			"core_minimum": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},

			"core_maximum": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},

			"core_lifecycle": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"core_ebs_optimized": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"core_ebs_block_device": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volumes_per_instance": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1,
						},

						"volume_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						"size_in_gb": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"iops": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
					},
				},
			},

			"core_scaling_up_policy": mrScalerScalingPolicySchema(),

			"core_scaling_down_policy": mrScalerScalingPolicySchema(),

			"task_instance_types": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"task_target": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"task_minimum": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"task_maximum": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"task_lifecycle": {
				Type:     schema.TypeString,
				Required: true,
			},

			"task_ebs_optimized": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"task_ebs_block_device": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volumes_per_instance": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1,
						},

						"volume_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						"size_in_gb": {
							Type:     schema.TypeInt,
							Required: true,
						},

						"iops": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
					},
				},
			},

			"task_scaling_up_policy": mrScalerScalingPolicySchema(),

			"task_scaling_down_policy": mrScalerScalingPolicySchema(),

			"configurations_file": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bucket": {
							Type:     schema.TypeString,
							Required: true,
						},

						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"availability_zone": {
				Type:          schema.TypeSet,
				Optional:      true,
				ConflictsWith: []string{"availability_zones"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"availability_zones": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},

						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func mrScalerScalingPolicySchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"policy_name": {
					Type:     schema.TypeString,
					Required: true,
				},

				"metric_name": {
					Type:     schema.TypeString,
					Required: true,
				},

				"statistic": {
					Type:     schema.TypeString,
					Required: true,
				},

				"unit": {
					Type:     schema.TypeString,
					Required: true,
				},

				"threshold": {
					Type:     schema.TypeFloat,
					Required: true,
				},

				"adjustment": {
					Type:     schema.TypeString,
					Optional: true,
					Default:  -1,
				},

				"min_target_capacity": {
					Type:     schema.TypeString,
					Optional: true,
				},

				"max_target_capacity": {
					Type:     schema.TypeString,
					Optional: true,
				},

				"namespace": {
					Type:     schema.TypeString,
					Required: true,
				},

				"operator": {
					Type:     schema.TypeString,
					Required: true,
				},

				"evaluation_periods": {
					Type:     schema.TypeInt,
					Required: true,
				},

				"period": {
					Type:     schema.TypeInt,
					Required: true,
				},

				"cooldown": {
					Type:     schema.TypeInt,
					Required: true,
				},

				"dimensions": {
					Type:     schema.TypeMap,
					Optional: true,
				},

				"minimum": {
					Type:     schema.TypeString,
					Optional: true,
				},

				"maximum": {
					Type:     schema.TypeString,
					Optional: true,
				},

				"target": {
					Type:     schema.TypeString,
					Optional: true,
				},

				"action_type": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}

//endregion

//region CRUD methods

func resourceSpotinstAWSMrScalerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	newAWSMrScaler, err := buildAWSMrScalerOpts(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] MR Scaler create configuration: %s", stringutil.Stringify(newAWSMrScaler))
	input := &mrscaler.CreateScalerInput{Scaler: newAWSMrScaler}
	resp, err := client.mrscaler.Create(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to create mr scaler: %s", err)
	}
	d.SetId(spotinst.StringValue(resp.Scaler.ID))
	log.Printf("[INFO] AWSMrScaler created successfully: %s", d.Id())
	return resourceSpotinstAWSMrScalerRead(d, meta)
}

func resourceSpotinstAWSMrScalerRead(in *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	input := &mrscaler.ReadScalerInput{ScalerID: spotinst.String(in.Id())}
	resp, err := client.mrscaler.Read(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to read mr scaler: %s", err)
	}
	if ext := resp.Scaler; ext != nil {
		in.Set("name", ext.Name)
		in.Set("description", ext.Description)
		in.Set("region", ext.Region)

		if ext.Strategy != nil {
			if ext.Strategy.Wrapping != nil {
				in.Set("clusterId", ext.Strategy.Wrapping.SourceClusterID)
			} else {
				in.Set("clusterId", ext.Strategy.Cloning.OriginClusterID)
			}
		}

		if ext.Compute != nil {
			if ext.Compute.InstanceGroups != nil {
				// Set Master Instance Group
				if ext.Compute.InstanceGroups.MasterGroup != nil {
					// Set Lifecycle
					if ext.Compute.InstanceGroups.MasterGroup.LifeCycle != nil {
						if err := in.Set("master_lifecycle", ext.Compute.InstanceGroups.MasterGroup.LifeCycle); err != nil {
							return fmt.Errorf("failed to set master lifecycle configuration: %#v", err)
						}
					}

					// Set Target
					if ext.Compute.InstanceGroups.MasterGroup.Target != nil {
						if err := in.Set("master_target", ext.Compute.InstanceGroups.MasterGroup.Target); err != nil {
							return fmt.Errorf("failed to set master target configuration: %#v", err)
						}
					}

					// Set Instance Types
					if ext.Compute.InstanceGroups.MasterGroup.InstanceTypes != nil {
						if err := in.Set("master_instance_types", ext.Compute.InstanceGroups.MasterGroup.InstanceTypes); err != nil {
							return fmt.Errorf("failed to set master instance types configuration: %#v", err)
						}
					}

					// Set EBS Configuration
					if ext.Compute.InstanceGroups.MasterGroup.EBSConfiguration != nil {
						// Set EBS Optimized
						if ext.Compute.InstanceGroups.MasterGroup.EBSConfiguration.Optimized != nil {
							if err := in.Set("master_ebs_optimized", ext.Compute.InstanceGroups.MasterGroup.EBSConfiguration.Optimized); err != nil {
								return fmt.Errorf("failed to set master ebs optimized configuration: %#v", err)
							}
						}

						// Set block devices.
						if ext.Compute.InstanceGroups.MasterGroup.EBSConfiguration.BlockDeviceConfigs != nil {
							if ext.Compute.InstanceGroups.MasterGroup.EBSConfiguration.BlockDeviceConfigs != nil {
								if err := in.Set("master_ebs_block_device", flattenAWSMrScalerBlockDevices(ext.Compute.InstanceGroups.MasterGroup.EBSConfiguration.BlockDeviceConfigs)); err != nil {
									return fmt.Errorf("failed to set master block devices configuration: %#v", err)
								}
							}
						}
					}
				}

				// Set Core Instance Group
				if ext.Compute.InstanceGroups.CoreGroup != nil {
					// Set Lifecycle
					if ext.Compute.InstanceGroups.CoreGroup.LifeCycle != nil {
						if err := in.Set("core_lifecycle", ext.Compute.InstanceGroups.CoreGroup.LifeCycle); err != nil {
							return fmt.Errorf("failed to set core lifecycle configuration: %#v", err)
						}
					}

					// Set Target
					if ext.Compute.InstanceGroups.CoreGroup.Target != nil {
						if err := in.Set("core_target", ext.Compute.InstanceGroups.CoreGroup.Target); err != nil {
							return fmt.Errorf("failed to set core capacity target configuration: %#v", err)
						}
					}

					// Set Capacity
					if ext.Compute.InstanceGroups.CoreGroup.Capacity != nil {
						// Set Target
						if ext.Compute.InstanceGroups.CoreGroup.Capacity.Target != nil {
							if err := in.Set("core_target", ext.Compute.InstanceGroups.CoreGroup.Capacity.Target); err != nil {
								return fmt.Errorf("failed to set core target configuration: %#v", err)
							}
						}

						// Set Minimum
						if ext.Compute.InstanceGroups.CoreGroup.Capacity.Minimum != nil {
							if err := in.Set("core_minimum", ext.Compute.InstanceGroups.CoreGroup.Capacity.Minimum); err != nil {
								return fmt.Errorf("failed to set core minimum configuration: %#v", err)
							}
						}

						// Set Maximum
						if ext.Compute.InstanceGroups.CoreGroup.Capacity.Maximum != nil {
							if err := in.Set("core_maximum", ext.Compute.InstanceGroups.CoreGroup.Capacity.Maximum); err != nil {
								return fmt.Errorf("failed to set core maximum configuration: %#v", err)
							}
						}
					}

					// Set Instance Types
					if ext.Compute.InstanceGroups.CoreGroup.InstanceTypes != nil {
						if err := in.Set("core_instance_types", ext.Compute.InstanceGroups.CoreGroup.InstanceTypes); err != nil {
							return fmt.Errorf("failed to set core instance types configuration: %#v", err)
						}
					}

					// Set EBS Configuration
					if ext.Compute.InstanceGroups.CoreGroup.EBSConfiguration != nil {
						// Set EBS Optimized
						if ext.Compute.InstanceGroups.CoreGroup.EBSConfiguration.Optimized != nil {
							if err := in.Set("core_ebs_optimized", ext.Compute.InstanceGroups.CoreGroup.EBSConfiguration.Optimized); err != nil {
								return fmt.Errorf("failed to set core ebs optimized configuration: %#v", err)
							}
						}

						// Set block devices.
						if ext.Compute.InstanceGroups.CoreGroup.EBSConfiguration.BlockDeviceConfigs != nil {
							if ext.Compute.InstanceGroups.CoreGroup.EBSConfiguration.BlockDeviceConfigs != nil {
								if err := in.Set("core_ebs_block_device", flattenAWSMrScalerBlockDevices(ext.Compute.InstanceGroups.CoreGroup.EBSConfiguration.BlockDeviceConfigs)); err != nil {
									return fmt.Errorf("failed to set core block devices configuration: %#v", err)
								}
							}
						}
					}
				}

				// Set Task Instance Group
				if ext.Compute.InstanceGroups.TaskGroup != nil {
					// Set Lifecycle
					if ext.Compute.InstanceGroups.TaskGroup.LifeCycle != nil {
						if err := in.Set("task_lifecycle", ext.Compute.InstanceGroups.TaskGroup.LifeCycle); err != nil {
							return fmt.Errorf("failed to set task lifecycle configuration: %#v", err)
						}
					}

					// Set Capacity
					if ext.Compute.InstanceGroups.TaskGroup.Capacity != nil {
						// Set Target
						if ext.Compute.InstanceGroups.TaskGroup.Capacity.Target != nil {
							if err := in.Set("task_target", ext.Compute.InstanceGroups.TaskGroup.Capacity.Target); err != nil {
								return fmt.Errorf("failed to set task capacity target configuration: %#v", err)
							}
						}

						// Set Minimum
						if ext.Compute.InstanceGroups.TaskGroup.Capacity.Minimum != nil {
							if err := in.Set("task_minimum", ext.Compute.InstanceGroups.TaskGroup.Capacity.Minimum); err != nil {
								return fmt.Errorf("failed to set task minimum configuration: %#v", err)
							}
						}

						// Set Maximum
						if ext.Compute.InstanceGroups.TaskGroup.Capacity.Maximum != nil {
							if err := in.Set("task_maximum", ext.Compute.InstanceGroups.TaskGroup.Capacity.Maximum); err != nil {
								return fmt.Errorf("failed to set task maximum configuration: %#v", err)
							}
						}
					}

					// Set Instance Types
					if ext.Compute.InstanceGroups.TaskGroup.InstanceTypes != nil {
						if err := in.Set("task_instance_types", ext.Compute.InstanceGroups.TaskGroup.InstanceTypes); err != nil {
							return fmt.Errorf("failed to set task instance types configuration: %#v", err)
						}
					}

					// Set EBS Configuration
					if ext.Compute.InstanceGroups.TaskGroup.EBSConfiguration != nil {
						// Set EBS Optimized
						if ext.Compute.InstanceGroups.TaskGroup.EBSConfiguration.Optimized != nil {
							if err := in.Set("task_ebs_optimized", ext.Compute.InstanceGroups.TaskGroup.EBSConfiguration.Optimized); err != nil {
								return fmt.Errorf("failed to set task ebs optimized configuration: %#v", err)
							}
						}

						// Set block devices.
						if ext.Compute.InstanceGroups.TaskGroup.EBSConfiguration.BlockDeviceConfigs != nil {
							if ext.Compute.InstanceGroups.TaskGroup.EBSConfiguration.BlockDeviceConfigs != nil {
								if err := in.Set("task_ebs_block_device", flattenAWSMrScalerBlockDevices(ext.Compute.InstanceGroups.TaskGroup.EBSConfiguration.BlockDeviceConfigs)); err != nil {
									return fmt.Errorf("failed to set task block devices configuration: %#v", err)
								}
							}
						}
					}
				}
			}

			// Set Configuration file
			if ext.Compute.Configurations != nil {
				if ext.Compute.Configurations.File != nil {
					if err := in.Set("configurations_file", flattenAWSMrScalerConfigurationsFile(ext.Compute.Configurations.File)); err != nil {
						return fmt.Errorf("failed to set configurations file configuration: %#v", err)
					}
				}
			}

			// Set Availability Zones
			if ext.Compute.AvailabilityZones != nil {
				// TODO - implement
			}

			// Set Tags
			if ext.Compute.Tags != nil {
				// TODO - implement
			}
		}

		// Set Task Scaling
		if ext.Scaling != nil {
			// Set Down scaling policies
			if ext.Scaling.Down != nil {
				if err := in.Set("task_scaling_down_policy", flattenAWSMrScalerScalingPolicies(ext.Scaling.Down)); err != nil {
					return fmt.Errorf("failed to set task down scaling policy configuration: %#v", err)
				}
			}

			// Set Up scaling policies
			if ext.Scaling.Up != nil {
				if err := in.Set("task_scaling_up_policy", flattenAWSMrScalerScalingPolicies(ext.Scaling.Up)); err != nil {
					return fmt.Errorf("failed to set task up scaling policy configuration: %#v", err)
				}
			}
		}

		// Set Core Scaling
		if ext.CoreScaling != nil {
			// Set Down scaling policies
			if ext.CoreScaling.Down != nil {
				if err := in.Set("core_scaling_down_policy", flattenAWSMrScalerScalingPolicies(ext.CoreScaling.Down)); err != nil {
					return fmt.Errorf("failed to set core down scaling policy configuration: %#v", err)
				}
			}

			// Set Up scaling policies
			if ext.CoreScaling.Up != nil {
				if err := in.Set("core_scaling_up_policy", flattenAWSMrScalerScalingPolicies(ext.CoreScaling.Up)); err != nil {
					return fmt.Errorf("failed to set core up scaling policy configuration: %#v", err)
				}
			}
		}

	} else {
		in.SetId("")
	}
	return nil
}

func resourceSpotinstAWSMrScalerUpdate(loc *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	ext := &mrscaler.Scaler{}
	ext.SetId(spotinst.String(loc.Id()))

	disallowedFieldUpdate := false
	disallowedField := ""

	if loc.HasChange("configurations_file") {
		disallowedField = "configurations_file"
		disallowedFieldUpdate = true
	}

	if loc.HasChange("availability_zones") {
		disallowedField = "availability_zones"
		disallowedFieldUpdate = true
	}

	if loc.HasChange("task_lifecycle") {
		disallowedField = "task_lifecycle"
		disallowedFieldUpdate = true
	}

	if loc.HasChange("tags") {
		disallowedField = "tags"
		disallowedFieldUpdate = true
	}

	if loc.HasChange("core_lifecycle") {
		disallowedField = "core_lifecycle"
		disallowedFieldUpdate = true
	}

	if loc.HasChange("core_instance_types") {
		disallowedField = "core_instance_types"
		disallowedFieldUpdate = true
	}

	if loc.HasChange("master_ebs_block_device") {
		disallowedField = "master_ebs_block_device"
		disallowedFieldUpdate = true
	}

	if loc.HasChange("master_ebs_optimized") {
		disallowedField = "master_ebs_optimized"
		disallowedFieldUpdate = true
	}

	if loc.HasChange("master_lifecycle") {
		disallowedField = "master_lifecycle"
		disallowedFieldUpdate = true
	}

	if loc.HasChange("master_target") {
		disallowedField = "master_target"
		disallowedFieldUpdate = true
	}

	if loc.HasChange("master_instance_types") {
		disallowedField = "master_instance_types"
		disallowedFieldUpdate = true
	}

	if loc.HasChange("cluster_id") {
		disallowedField = "cluster_id"
		disallowedFieldUpdate = true
	}

	if loc.HasChange("strategy") {
		disallowedField = "strategy"
		disallowedFieldUpdate = true
	}

	if loc.HasChange("region") {
		disallowedField = "region"
		disallowedFieldUpdate = true
	}

	if disallowedFieldUpdate == true {
		return fmt.Errorf("field %s is immutable - revert to previous value to proceed", disallowedField)
	}

	update := false

	if loc.HasChange("name") {
		ext.SetName(spotinst.String(loc.Get("name").(string)))
		update = true
	}

	if loc.HasChange("description") {
		ext.SetDescription(spotinst.String(loc.Get("description").(string)))
		update = true
	}

	if loc.HasChange("core_minimum") || loc.HasChange("core_maximum") || loc.HasChange("core_target") {
		buildEmptyCoreCapacity(ext)

		// Required properties (despite not being required in schema)
		ext.Compute.InstanceGroups.CoreGroup.Capacity.SetMinimum(spotinst.Int(loc.Get("core_minimum").(int)))
		ext.Compute.InstanceGroups.CoreGroup.Capacity.SetMaximum(spotinst.Int(loc.Get("core_maximum").(int)))
		ext.Compute.InstanceGroups.CoreGroup.Capacity.SetTarget(spotinst.Int(loc.Get("core_target").(int)))

		update = true
	}

	if loc.HasChange("task_minimum") || loc.HasChange("task_maximum") || loc.HasChange("task_target") {
		buildEmptyTaskCapacity(ext)

		// Required properties
		ext.Compute.InstanceGroups.TaskGroup.Capacity.SetMinimum(spotinst.Int(loc.Get("task_minimum").(int)))
		ext.Compute.InstanceGroups.TaskGroup.Capacity.SetMaximum(spotinst.Int(loc.Get("task_maximum").(int)))
		ext.Compute.InstanceGroups.TaskGroup.Capacity.SetTarget(spotinst.Int(loc.Get("task_target").(int)))

		update = true
	}

	if loc.HasChange("task_instance_types") {
		buildEmptyTaskGroup(ext)

		if v, ok := loc.GetOk("task_instance_types"); ok {
			types := expandInstanceTypesList(v.([]interface{}))
			ext.Compute.InstanceGroups.TaskGroup.SetInstanceTypes(types)
		} else {
			ext.Compute.InstanceGroups.TaskGroup.SetInstanceTypes(nil)
		}
		update = true
	}

	if loc.HasChange("core_ebs_optimized") {
		buildEmptyCoreEBSConfiguration(ext)

		// Has default boolean value
		ext.Compute.InstanceGroups.CoreGroup.EBSConfiguration.SetOptimized(spotinst.Bool(loc.Get("core_ebs_optimized").(bool)))
		update = true
	}

	if loc.HasChange("task_ebs_optimized") {
		buildEmptyTaskEBSConfiguration(ext)

		// Has default boolean value
		ext.Compute.InstanceGroups.TaskGroup.EBSConfiguration.SetOptimized(spotinst.Bool(loc.Get("task_ebs_optimized").(bool)))
		update = true
	}

	if loc.HasChange("task_ebs_block_device") {
		buildEmptyTaskEBSConfiguration(ext)

		if devices, err := expandAWSMrScalerDevices(loc.Get("task_ebs_block_device")); err != nil {
			return err
		} else {
			log.Printf("[DEBUG] Mr Scaler update task ebs devices configuration: %s", stringutil.Stringify(devices))
			ext.Compute.InstanceGroups.TaskGroup.EBSConfiguration.SetBlockDeviceConfigs(devices)
		}
		update = true
	}

	if loc.HasChange("core_ebs_block_device") {
		buildEmptyCoreEBSConfiguration(ext)

		if devices, err := expandAWSMrScalerDevices(loc.Get("core_ebs_block_device")); err != nil {
			return err
		} else {
			log.Printf("[DEBUG] Mr Scaler update core ebs devices configuration: %s", stringutil.Stringify(devices))
			ext.Compute.InstanceGroups.CoreGroup.EBSConfiguration.SetBlockDeviceConfigs(devices)
		}

		update = true
	}

	if loc.HasChange("core_scaling_up_policy") {
		buildEmptyCoreScaling(ext)

		if v := loc.Get("core_scaling_up_policy"); v != nil {
			if policies, err := expandAWSMrScalerScalingPolicies(loc.Get("core_scaling_up_policy")); err != nil {
				return err
			} else {
				ext.CoreScaling.SetUp(policies)
			}
		} else {
			// Partial update in lists
			ext.CoreScaling.SetUp([]*mrscaler.ScalingPolicy{})
		}

		update = true
	}

	if loc.HasChange("core_scaling_down_policy") {
		buildEmptyCoreScaling(ext)

		if v := loc.Get("core_scaling_down_policy"); v != nil {
			if policies, err := expandAWSMrScalerScalingPolicies(loc.Get("core_scaling_down_policy")); err != nil {
				return err
			} else {
				ext.CoreScaling.SetDown(policies)
			}
		} else {
			// Partial update in lists
			ext.CoreScaling.SetDown([]*mrscaler.ScalingPolicy{})
		}

		update = true
	}

	if loc.HasChange("task_scaling_up_policy") {
		buildEmptyTaskScaling(ext)

		if v := loc.Get("task_scaling_up_policy"); v != nil {
			if policies, err := expandAWSMrScalerScalingPolicies(loc.Get("task_scaling_up_policy")); err != nil {
				return err
			} else {
				ext.Scaling.SetUp(policies)
			}
		} else {
			// Partial update in lists
			ext.Scaling.SetUp([]*mrscaler.ScalingPolicy{})
		}
		update = true
	}

	if loc.HasChange("task_scaling_down_policy") {
		buildEmptyTaskScaling(ext)

		if v := loc.Get("task_scaling_down_policy"); v != nil {
			if policies, err := expandAWSMrScalerScalingPolicies(loc.Get("task_scaling_down_policy")); err != nil {
				return err
			} else {
				ext.Scaling.SetDown(policies)
			}
		} else {
			// Partial update in lists
			ext.Scaling.SetDown([]*mrscaler.ScalingPolicy{})
		}
		update = true
	}

	if update {
		log.Printf("[DEBUG] Mr Scaler update configuration: %s", stringutil.Stringify(ext))
		input := &mrscaler.UpdateScalerInput{Scaler: ext}
		if _, err := client.mrscaler.Update(context.Background(), input); err != nil {
			return fmt.Errorf("failed to update mr scaler %s: %s", loc.Id(), err)
		}
	}

	return resourceSpotinstAWSMrScalerRead(loc, meta)
}

func resourceSpotinstAWSMrScalerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	log.Printf("[INFO] Deleting mr scaler: %s", d.Id())
	input := &mrscaler.DeleteScalerInput{ScalerID: spotinst.String(d.Id())}
	if _, err := client.mrscaler.Delete(context.Background(), input); err != nil {
		return fmt.Errorf("failed to delete mr scaler: %s", err)
	}
	d.SetId("")
	return nil
}

//endregion

//region Flatten methods
func flattenAWSMrScalerBlockDevices(blockDevices []*mrscaler.BlockDeviceConfig) []interface{} {
	result := make([]interface{}, 0, len(blockDevices))
	for _, b := range blockDevices {
		m := make(map[string]interface{})

		m["volumes_per_instance"] = spotinst.IntValue(b.VolumesPerInstance)

		if b.VolumeSpecification != nil {
			m["iops"] = spotinst.IntValue(b.VolumeSpecification.IOPS)
			m["size_in_gb"] = spotinst.IntValue(b.VolumeSpecification.SizeInGB)
			m["volume_type"] = spotinst.StringValue(b.VolumeSpecification.VolumeType)
		}

		result = append(result, m)
	}
	return result
}

func flattenAWSMrScalerScalingPolicies(policies []*mrscaler.ScalingPolicy) []interface{} {
	result := make([]interface{}, 0, len(policies))
	for _, p := range policies {
		m := make(map[string]interface{})

		if p.Action != nil && p.Action.Type != nil {
			m["action_type"] = spotinst.StringValue(p.Action.Type)
			m["adjustment"] = spotinst.StringValue(p.Action.Adjustment)
			m["min_target_capacity"] = spotinst.StringValue(p.Action.MaxTargetCapacity)
			m["max_target_capacity"] = spotinst.StringValue(p.Action.MaxTargetCapacity)
			m["minimum"] = spotinst.StringValue(p.Action.Minimum)
			m["maximum"] = spotinst.StringValue(p.Action.Maximum)
			m["target"] = spotinst.StringValue(p.Action.Target)
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

func flattenAWSMrScalerConfigurationsFile(file *mrscaler.ConfigurationFile) []interface{} {
	result := make(map[string]interface{})
	result["key"] = spotinst.StringValue(file.Key)
	result["bucket"] = spotinst.StringValue(file.Bucket)
	return []interface{}{result}
}

//endregion

//region Build Empty Methods

func buildEmptyConfigurations(scaler *mrscaler.Scaler) {
	buildEmptyCompute(scaler)
	if scaler.Compute.Configurations == nil {
		scaler.Compute.SetConfigurations(&mrscaler.Configurations{})
	}
}

func buildEmptyCoreScaling(ext *mrscaler.Scaler) {
	if ext.CoreScaling == nil {
		ext.SetCoreScaling(&mrscaler.Scaling{})
	}
}

func buildEmptyTaskScaling(ext *mrscaler.Scaler) {
	if ext.Scaling == nil {
		ext.SetScaling(&mrscaler.Scaling{})
	}
}

func buildEmptyCoreCapacity(ext *mrscaler.Scaler) {
	buildEmptyCoreGroup(ext)

	if ext.Compute.InstanceGroups.CoreGroup.Capacity == nil {
		ext.Compute.InstanceGroups.CoreGroup.SetCapacity(&mrscaler.InstanceGroupCapacity{})
	}
}

func buildEmptyCoreEBSConfiguration(ext *mrscaler.Scaler) {
	buildEmptyCoreGroup(ext)

	if ext.Compute.InstanceGroups.CoreGroup.EBSConfiguration == nil {
		ext.Compute.InstanceGroups.CoreGroup.SetEBSConfiguration(&mrscaler.EBSConfiguration{})
	}
}

func buildEmptyTaskEBSConfiguration(ext *mrscaler.Scaler) {
	buildEmptyTaskGroup(ext)

	if ext.Compute.InstanceGroups.TaskGroup.EBSConfiguration == nil {
		ext.Compute.InstanceGroups.TaskGroup.SetEBSConfiguration(&mrscaler.EBSConfiguration{})
	}
}

func buildEmptyCoreGroup(ext *mrscaler.Scaler) {
	buildEmptyInstanceGroups(ext)
	if ext.Compute.InstanceGroups.CoreGroup == nil {
		ext.Compute.InstanceGroups.SetCoreGroup(&mrscaler.InstanceGroup{})
	}
}

func buildEmptyInstanceGroups(ext *mrscaler.Scaler) {

	buildEmptyCompute(ext)

	if ext.Compute.InstanceGroups == nil {
		ext.Compute.SetInstanceGroups(&mrscaler.InstanceGroups{})
	}
}

func buildEmptyCompute(ext *mrscaler.Scaler) {
	if ext.Compute == nil {
		ext.SetCompute(&mrscaler.Compute{})
	}
}

func buildEmptyTaskCapacity(ext *mrscaler.Scaler) {
	buildEmptyTaskGroup(ext)

	if ext.Compute.InstanceGroups.TaskGroup.Capacity == nil {
		ext.Compute.InstanceGroups.TaskGroup.SetCapacity(&mrscaler.InstanceGroupCapacity{})
	}
}

func buildEmptyTaskGroup(ext *mrscaler.Scaler) {
	buildEmptyInstanceGroups(ext)

	if ext.Compute.InstanceGroups.TaskGroup == nil {
		ext.Compute.InstanceGroups.SetTaskGroup(&mrscaler.InstanceGroup{})
	}
}

//endregion

//region Build methods
func buildAWSMrScalerOpts(d *schema.ResourceData) (*mrscaler.Scaler, error) {
	scaler := &mrscaler.Scaler{
		Strategy: &mrscaler.Strategy{},
		Compute: &mrscaler.Compute{
			InstanceGroups: &mrscaler.InstanceGroups{
				MasterGroup: &mrscaler.InstanceGroup{},
				CoreGroup:   &mrscaler.InstanceGroup{},
				TaskGroup:   &mrscaler.InstanceGroup{},
			},
		},
		Scaling:     &mrscaler.Scaling{},
		CoreScaling: &mrscaler.Scaling{},
	}

	scaler.SetName(spotinst.String(d.Get("name").(string)))

	scaler.SetDescription(spotinst.String(d.Get("description").(string)))

	scaler.SetRegion(spotinst.String(d.Get("region").(string)))

	strategy := handleStrategy(d, scaler)
	buildAWSMrScalerInstanceGroups(d, scaler, strategy)

	if v, ok := d.GetOk("tags"); ok {
		buildEmptyCompute(scaler)

		list := v.(*schema.Set).List()
		tags := make([]*mrscaler.Tag, 0, len(list))
		for _, v := range list {
			attr, ok := v.(map[string]interface{})
			if !ok {
				continue
			}
			if _, ok := attr["key"]; !ok {
				return scaler, errors.New("invalid tag attributes: key missing")
			}

			if _, ok := attr["value"]; !ok {
				return scaler, errors.New("invalid tag attributes: value missing")
			}
			tag := &mrscaler.Tag{
				Key:   spotinst.String(attr["key"].(string)),
				Value: spotinst.String(attr["value"].(string)),
			}

			log.Printf("[DEBUG] MR Scaler tag configuration: %s", stringutil.Stringify(tag))
			tags = append(tags, tag)
		}

		scaler.Compute.SetTags(tags)
	}

	if zones, ok := d.GetOk("availability_zones"); ok {
		buildEmptyCompute(scaler)
		list := zones.([]interface{})
		zones := make([]*mrscaler.AvailabilityZone, 0, len(list))
		for _, str := range list {
			if s, ok := str.(string); ok {
				parts := strings.Split(s, ":")
				zone := &mrscaler.AvailabilityZone{}
				if len(parts) >= 1 && parts[0] != "" {
					zone.SetName(spotinst.String(parts[0]))
				}
				if len(parts) == 2 && parts[1] != "" {
					zone.SetSubnetId(spotinst.String(parts[1]))
				}
				log.Printf("[DEBUG] MrScaler availability zone configuration: %s", stringutil.Stringify(zone))
				zones = append(zones, zone)
			}
		}

		scaler.Compute.SetAvailabilityZones(zones)

	}

	if cfgFile, ok := d.GetOk("configurations_file"); ok {
		buildEmptyConfigurations(scaler)
		cf := &mrscaler.ConfigurationFile{}

		list := cfgFile.(*schema.Set).List()
		m := list[0].(map[string]interface{})

		if v, ok := m["key"].(string); ok {
			cf.SetKey(spotinst.String(v))
		}

		if v, ok := m["bucket"].(string); ok {
			cf.SetBucket(spotinst.String(v))
		}

		scaler.Compute.Configurations.SetFile(cf)
	}

	if v, ok := d.GetOk("core_scaling_up_policy"); ok {
		if policies, err := expandAWSMrScalerScalingPolicies(v); err != nil {
			return nil, err
		} else {
			buildEmptyCoreScaling(scaler)
			scaler.CoreScaling.SetUp(policies)
		}
	}

	if v, ok := d.GetOk("core_scaling_down_policy"); ok {
		if policies, err := expandAWSMrScalerScalingPolicies(v); err != nil {
			return nil, err
		} else {
			buildEmptyCoreScaling(scaler)
			scaler.CoreScaling.SetDown(policies)
		}
	}

	if v, ok := d.GetOk("task_scaling_up_policy"); ok {
		if policies, err := expandAWSMrScalerScalingPolicies(v); err != nil {
			return nil, err
		} else {
			buildEmptyTaskScaling(scaler)
			scaler.Scaling.SetUp(policies)
		}
	}

	if v, ok := d.GetOk("task_scaling_down_policy"); ok {
		if policies, err := expandAWSMrScalerScalingPolicies(v); err != nil {
			return nil, err
		} else {
			buildEmptyTaskScaling(scaler)
			scaler.Scaling.SetDown(policies)
		}
	}

	return scaler, nil
}

func handleStrategy(d *schema.ResourceData, scaler *mrscaler.Scaler) string {
	retVal := ""

	if strategy, ok := d.Get("strategy").(string); ok && strategy != "" {
		retVal = strategy
		log.Printf("[DEBUG] Strategy chosen: %s", stringutil.Stringify(strategy))
		s := &mrscaler.Strategy{}

		if clusterID, ok := d.Get("cluster_id").(string); ok && clusterID != "" {
			if strategy == "clone" {
				clone := &mrscaler.Cloning{}
				clone.SetOriginClusterId(spotinst.String(clusterID))
				s.SetCloning(clone)
				scaler.SetStrategy(s)
			} else if strategy == "wrap" {
				wrap := &mrscaler.Wrapping{}
				wrap.SetSourceClusterId(spotinst.String(clusterID))
				s.SetWrapping(wrap)
				scaler.SetStrategy(s)
			}
		}
	}

	log.Printf("[DEBUG] Strategy chosen: %s", stringutil.Stringify(retVal))
	return retVal
}

func buildAWSMrScalerInstanceGroups(d *schema.ResourceData, scaler *mrscaler.Scaler, strategy string) error {
	buildEmptyCompute(scaler)
	instanceGroups := &mrscaler.InstanceGroups{}

	// On Clone Strategy - set master and core group
	if strategy == "clone" {
		masterGroup, err := buildInstanceGroup(d, mrscaler.InstanceGroupTypeMaster.String())
		if err != nil {
			return errors.New("could not expand master instance group")
		}
		instanceGroups.SetMasterGroup(masterGroup)

		coreGroup, err := buildInstanceGroup(d, mrscaler.InstanceGroupTypeCore.String())
		if err != nil {
			return errors.New("could not expand core instance group")
		}
		instanceGroups.SetCoreGroup(coreGroup)
	}

	// On either strategy - set task group
	taskGroup, err := buildInstanceGroup(d, mrscaler.InstanceGroupTypeTask.String())
	if err != nil {
		return errors.New("could not expand task instance group")
	}
	instanceGroups.SetTaskGroup(taskGroup)

	scaler.Compute.SetInstanceGroups(instanceGroups)
	return err
}

func buildInstanceGroup(d *schema.ResourceData, instanceGroupType string) (*mrscaler.InstanceGroup, error) {
	instanceGroup := &mrscaler.InstanceGroup{}
	instanceTypes := d.Get(instanceGroupType + "_instance_types").([]interface{})
	types := make([]string, 0, len(instanceTypes))

	for _, str := range instanceTypes {
		if typ, ok := str.(string); ok {
			log.Printf("[DEBUG] Instance group Instance type: %s", stringutil.Stringify(typ))
			types = append(types, typ)
		}
	}

	log.Printf("[DEBUG] " + reflect.TypeOf(types).String())
	log.Printf("[DEBUG] Instance types: %s", stringutil.Stringify(instanceTypes))
	instanceGroup.SetInstanceTypes(types)
	instanceGroup.SetLifeCycle(spotinst.String(d.Get(instanceGroupType + "_lifecycle").(string)))

	if instanceGroupType == mrscaler.InstanceGroupTypeMaster.String() {
		if target, ok := d.Get(instanceGroupType + "_target").(int); ok && target != -1 {
			log.Printf("[DEBUG] Setting Maximum in %s", instanceGroupType)
			instanceGroup.SetTarget(spotinst.Int(target))
		}
	} else {
		buildGroupCapacity(d, instanceGroupType, instanceGroup)
	}

	if v, ok := d.GetOk(instanceGroupType + "_ebs_block_device"); ok {
		if devices, err := expandAWSMrScalerDevices(v); err != nil {
			return nil, err
		} else {
			ebsConfig := &mrscaler.EBSConfiguration{}
			ebsConfig.SetBlockDeviceConfigs(devices)
			if optimized, ok := d.Get(instanceGroupType + "_ebs_optimized").(bool); ok {
				ebsConfig.SetOptimized(spotinst.Bool(optimized))
			}
			instanceGroup.SetEBSConfiguration(ebsConfig)
		}
	}

	return instanceGroup, nil
}

func buildGroupCapacity(d *schema.ResourceData, instanceGroupType string, instanceGroup *mrscaler.InstanceGroup) {
	capacity := &mrscaler.InstanceGroupCapacity{}

	if max, ok := d.Get(instanceGroupType + "_maximum").(int); ok && max != -1 {
		log.Printf("[DEBUG] Setting Maximum in %s", instanceGroupType)
		capacity.SetMaximum(spotinst.Int(max))
	}
	if min, ok := d.Get(instanceGroupType + "_minimum").(int); ok && min != -1 {
		log.Printf("[DEBUG] Setting Minimum in %s", instanceGroupType)
		capacity.SetMinimum(spotinst.Int(min))
	}
	if target, ok := d.Get(instanceGroupType + "_target").(int); ok && target != -1 {
		log.Printf("[DEBUG] Setting Target in %s", instanceGroupType)
		capacity.SetTarget(spotinst.Int(target))
	}

	instanceGroup.SetCapacity(capacity)
}

//endregion

//region Expanders
func expandInstanceTypesList(instanceTypes []interface{}) []string {
	types := make([]string, 0, len(instanceTypes))
	for _, str := range instanceTypes {
		if typ, ok := str.(string); ok {
			types = append(types, typ)
		}
	}
	return types
}

func expandAWSMrScalerDevices(data interface{}) ([]*mrscaler.BlockDeviceConfig, error) {
	list := data.(*schema.Set).List()
	devices := make([]*mrscaler.BlockDeviceConfig, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		device := &mrscaler.BlockDeviceConfig{VolumeSpecification: &mrscaler.VolumeSpecification{}}

		if v, ok := m["volumes_per_instance"].(int); ok {
			device.SetVolumesPerInstance(spotinst.Int(v))
		}

		if v, ok := m["size_in_gb"].(int); ok {
			device.VolumeSpecification.SetSizeInGB(spotinst.Int(v))
		}

		if v, ok := m["iops"].(int); ok && v > 0 {
			device.VolumeSpecification.SetIOPS(spotinst.Int(v))
		}

		if v, ok := m["volume_type"].(string); ok {
			device.VolumeSpecification.SetVolumeType(spotinst.String(v))
		}

		log.Printf("[DEBUG] Group elastic block device configuration: %s", stringutil.Stringify(device))
		devices = append(devices, device)
	}

	return devices, nil
}

func expandAWSMrScalerScalingPolicies(data interface{}) ([]*mrscaler.ScalingPolicy, error) {
	list := data.(*schema.Set).List()
	policies := make([]*mrscaler.ScalingPolicy, 0, len(list))
	for _, item := range list {
		m := item.(map[string]interface{})
		policy := &mrscaler.ScalingPolicy{}

		if v, ok := m["policy_name"].(string); ok {
			policy.SetPolicyName(spotinst.String(v))
		}

		if v, ok := m["metric_name"].(string); ok {
			policy.SetMetricName(spotinst.String(v))
		}

		if v, ok := m["statistic"].(string); ok {
			policy.SetStatistic(spotinst.String(v))
		}

		if v, ok := m["unit"].(string); ok {
			policy.SetUnit(spotinst.String(v))
		}

		if v, ok := m["threshold"].(float64); ok {
			policy.SetThreshold(spotinst.Float64(v))
		}

		if v, ok := m["namespace"].(string); ok {
			policy.SetNamespace(spotinst.String(v))
		}

		if v, ok := m["operator"].(string); ok {
			policy.SetOperator(spotinst.String(v))
		}

		if v, ok := m["period"].(int); ok {
			policy.SetPeriod(spotinst.Int(v))
		}

		if v, ok := m["evaluation_periods"].(int); ok {
			policy.SetEvaluationPeriods(spotinst.Int(v))
		}

		if v, ok := m["cooldown"].(int); ok {
			policy.SetCooldown(spotinst.Int(v))
		}

		if v, ok := m["dimensions"]; ok {
			dimensions := expandAWSMrScalerDimensions(v.(map[string]interface{}))
			if len(dimensions) > 0 {
				policy.SetDimensions(dimensions)
			}
		}

		if v, ok := m["action_type"].(string); ok {
			action := &mrscaler.Action{}
			action.SetType(spotinst.String(v))

			if v, ok := m["adjustment"].(string); ok && v != "" {
				action.SetAdjustment(spotinst.String(v))
			}

			if v, ok := m["min_target_capacity"].(string); ok && v != "" {
				action.SetMinTargetCapacity(spotinst.String(v))
			}

			if v, ok := m["max_target_capacity"].(string); ok && v != "" {
				action.SetMaxTargetCapacity(spotinst.String(v))
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
		}

		if v, ok := m["namespace"].(string); ok && v != "" {
			log.Printf("[DEBUG] Mr Scaler scaling policy configuration: %s", stringutil.Stringify(policy))
			policies = append(policies, policy)
		}
	}

	return policies, nil
}

func expandAWSMrScalerDimensions(list map[string]interface{}) []*mrscaler.Dimension {
	dimensions := make([]*mrscaler.Dimension, 0, len(list))
	for name, val := range list {
		dimension := &mrscaler.Dimension{}
		dimension.SetName(spotinst.String(name))
		dimension.SetValue(spotinst.String(val.(string)))
		log.Printf("[DEBUG] AWS Mr Scaler scaling policy dimension: %s", stringutil.Stringify(dimension))
		dimensions = append(dimensions, dimension)
	}
	return dimensions
}

//endregion

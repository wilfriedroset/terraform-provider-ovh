package ovh

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers"
)

func resourceDedicatedServerOlaAggregation() *schema.Resource {
	return &schema.Resource{
		Create: resourceDedicatedServerOlaAggregationCreate,
		Update: resourceDedicatedServerOlaAggregationUpdate,
		Read:   resourceDedicatedServerOlaAggregationRead,
		Delete: resourceDedicatedServerOlaAggregationDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(45 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The internal name of your dedicated server.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The resulting VirtualNetworkInterface name.",
			},
			"virtual_network_interfaces": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem:     schema.Schema{Type: schema.TypeString},
			},

			//Computed
			"comment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Details of this task",
			},
			"done_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Completion date",
			},
			"function": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Function name",
			},
			"last_update": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last update",
			},
			"need_scheduler": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether this task need to be scheduled or not",
			},
			"note": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Extra information about this task",
			},
			"planned_intervention_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the planned intervention for this task",
			},
			"tags": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Task result tags output",
				Elem:        schema.Schema{Type: schema.TypeString},
			},
			"task_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "the id of the task",
			},
			"ticket_reference": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Reference of the thread on the support tool",
			},
			"start_date": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Task Creation date",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Task status",
			},
		},
	}
}

func resourceDedicatedServerOlaAggregationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)

	endpoint := fmt.Sprintf(
		"/dedicated/server/%s/ola/aggregation",
		url.PathEscape(serviceName),
	)
	opts := (&DedicatedServerOlaAggregationCreateOpts{}).FromResource(d)
	task := &DedicatedServerTask{}

	if err := config.OVHClient.Post(endpoint, opts, task); err != nil {
		return fmt.Errorf("Error calling POST %s:\n\t %q", endpoint, err)
	}

	if err := waitForDedicatedServerTask(serviceName, task, config.OVHClient); err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%d", task.Id))

	return dedicatedServerOlaAggregationRead(d, meta)
}

func dedicatedServerOlaAggregationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf(
			"Could not parse install task id %s,%s:\n\t %q",
			serviceName,
			d.Id(),
			err,
		)
	}

	task, err := getDedicatedServerTask(serviceName, id, config.OVHClient)
	if err != nil {
		return helpers.CheckDeleted(d, err, fmt.Sprintf(
			"dedicated server task %s/%s",
			serviceName,
			d.Id(),
		))
	}

	d.Set("function", task.Function)
	d.Set("comment", task.Comment)
	d.Set("status", task.Status)
	d.Set("last_update", task.LastUpdate.Format(time.RFC3339))
	d.Set("done_date", task.DoneDate.Format(time.RFC3339))
	d.Set("start_date", task.StartDate.Format(time.RFC3339))

	return nil
}

func resourceDedicatedServerOlaAggregationUpdate(d *schema.ResourceData, meta interface{}) error {
	// nothing to do on update
	return resourceDedicatedServerOlaAggregationRead(d, meta)
}

func resourceDedicatedServerOlaAggregationRead(d *schema.ResourceData, meta interface{}) error {
	// Nothing to do on READ
	//
	// IMPORTANT: This resource doesn't represent a real resource
	// but instead a task on a dedicated server. OVH may clean its tasks database after a while
	// so that the API may return a 404 on a task id. If we hit a 404 on a READ, then
	// terraform will understand that it has to recreate the resource, and consequently
	// will trigger new install task on the dedicated server.
	// This is something we must avoid!
	//
	return nil
}

func resourceDedicatedServerOlaAggregationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)

	endpoint := fmt.Sprintf(
		"/dedicated/server/%s/ola/reset",
		url.PathEscape(serviceName),
	)
	opts := (&DedicatedServerOlaAggregationDeleteOpts{}).FromResource(d)
	task := &DedicatedServerTask{}

	for _, v := range opts.VirtualNetworkInterface {
		singleOpts := DedicatedServerOlaAggregationSingleDeleteOpts{
			VirtualNetworkInterface: v,
		}
		if err := config.OVHClient.Post(endpoint, singleOpts, task); err != nil {
			return fmt.Errorf("Error calling POST %s:\n\t %q", endpoint, err)
		}

		if err := waitForDedicatedServerTask(serviceName, task, config.OVHClient); err != nil {
			return err
		}
	}

	// we cant delete the task through the API, just forget about its Id
	d.SetId("")
	return nil
}

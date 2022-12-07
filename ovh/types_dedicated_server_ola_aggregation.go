package ovh

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DedicatedServerOlaAggregationCreateOpts struct {
	Name                    string      `json:"name"`
	VirtualNetworkInterface []uuid.UUID `json:"virtualNetworkInterfaces"`
}

func (opts *DedicatedServerOlaAggregationCreateOpts) FromResource(d *schema.ResourceData) *DedicatedServerOlaAggregationCreateOpts {
	opts.Name = d.Get("name").(string)
	virtualNetworkInterface := d.Get("virtual_network_interfaces").([]interface{})
	// Convert virtualNetworkInterface from []interface{} to []string
	opts.VirtualNetworkInterface = make([]uuid.UUID, len(virtualNetworkInterface))
	for i, v := range virtualNetworkInterface {
		virtualNetworkInterface[i] = uuid.MustParse(fmt.Sprint(v))
	}

	return opts
}

type DedicatedServerOlaAggregationSingleDeleteOpts struct {
	VirtualNetworkInterface uuid.UUID `json:"virtualNetworkInterface"`
}

type DedicatedServerOlaAggregationDeleteOpts struct {
	VirtualNetworkInterface []uuid.UUID `json:"virtualNetworkInterfaces"`
}

func (opts *DedicatedServerOlaAggregationDeleteOpts) FromResource(d *schema.ResourceData) *DedicatedServerOlaAggregationDeleteOpts {
	virtualNetworkInterface := d.Get("virtual_network_interfaces").([]interface{})
	// Convert virtualNetworkInterface from []interface{} to []string
	opts.VirtualNetworkInterface = make([]uuid.UUID, len(virtualNetworkInterface))
	for i, v := range virtualNetworkInterface {
		virtualNetworkInterface[i] = uuid.MustParse(fmt.Sprint(v))
	}

	return opts
}

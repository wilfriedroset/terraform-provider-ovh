package ovh

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDedicatedServerOlaAggregation_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckCredentials(t)
			testAccPreCheckDedicatedServer(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDedicatedServerOlaAggregation_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ovh_dedicated_server_aggregate.aggregation", "name", os.Getenv("OVH_DEDICATED_SERVER")),
					resource.TestCheckResourceAttr(
						"ovh_dedicated_server_aggregate.aggregation", "server_name", "bound0"),
					resource.TestCheckResourceAttr(
						"ovh_dedicated_server_aggregate.aggregation", "virtual_network_interfaces", `"3ebdf237-808c-4309-b4a3-0a96d4524ab8", "5f522cd6-4011-4889-8dfd-e55d8ae2e098", "928228c1-c574-4e67-9430-549888cfcd2b"`),
				),
			},
		},
	})
}

func testAccDedicatedServerOlaAggregation_basic() string {
	serviceName := os.Getenv("OVH_DEDICATED_SERVER")
	testName := "bound0"
	virtualNetworkInterfaces := `"3ebdf237-808c-4309-b4a3-0a96d4524ab8", "5f522cd6-4011-4889-8dfd-e55d8ae2e098", "928228c1-c574-4e67-9430-549888cfcd2b"`

	return fmt.Sprintf(
		testAccDedicatedServerOlaAggregation_Basic,
		serviceName,
		testName,
		virtualNetworkInterfaces,
	)
}

const testAccDedicatedServerOlaAggregation_Basic = `
data ovh_dedicated_server_boots "harddisk" {
  service_name 					= "%s"
  name 							= "%s"
  virtual_network_interfaces    = [%s]
}

resource ovh_dedicated_server_aggregate "aggregation" {
  service_name 					= data.ovh_dedicated_server_boots.harddisk.service_name
  name      					= data.ovh_dedicated_server_boots.harddisk.name
  virtual_network_interfaces   	= data.ovh_dedicated_server_boots.harddisk.virtual_network_interfaces
}
`

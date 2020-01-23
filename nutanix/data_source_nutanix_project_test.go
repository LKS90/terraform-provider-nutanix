package nutanix

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

// func TestAccNutanixProjectDataSource_basic(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:  func() { testAccPreCheck(t) },
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccProjectDataSourceConfig(randIntBetween(1, 10)),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(
// 						"data.nutanix_project.test", "prefix_length", "24"),
// 					resource.TestCheckResourceAttr(
// 						"data.nutanix_project.test", "project_type", "VLAN"),
// 					resource.TestCheckResourceAttrSet("data.nutanix_project.test", "cluster_name"),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccNutanixProjectDataSource_name(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:  func() { testAccPreCheck(t) },
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccProjectDataSourceConfigName(randIntBetween(11, 20)),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(
// 						"data.nutanix_project.test", "prefix_length", "24"),
// 					resource.TestCheckResourceAttr(
// 						"data.nutanix_project.test", "project_type", "VLAN"),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccNutanixProjectDataSource_conflicts(t *testing.T) {
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:  func() { testAccPreCheck(t) },
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config:      testAccProjectDataSourceBadConfig(),
// 				ExpectError: regexp.MustCompile("conflicts with"),
// 			},
// 		},
// 	})
// }

// func testAccProjectDataSourceConfig(r int) string {
// 	return fmt.Sprintf(`
// data "nutanix_clusters" "clusters" {}

// locals {
// 	cluster1 = [
// 	for cluster in data.nutanix_clusters.clusters.entities :
// 	cluster.metadata.uuid if cluster.service_list[0] != "PRISM_CENTRAL"
// 	][0]
// }

// resource "nutanix_project" "test" {
// 	name = "dou_vlan0_test_%d"
// 	cluster_uuid = local.cluster1

// 	vlan_id = %d
// 	project_type = "VLAN"

// 	prefix_length = 24
// 	default_gateway_ip = "192.168.0.1"
// 	project_ip = "192.168.0.0"
// 	#ip_config_pool_list_ranges = ["192.168.0.5", "192.168.0.100"]

// 	dhcp_options = {
// 		boot_file_name   = "bootfile"
// 		domain_name      = "nutanix"
// 		tftp_server_name = "10.250.140.200"
// 	}
	
// 	dhcp_domain_name_server_list = ["8.8.8.8", "4.2.2.2"]
// 	dhcp_domain_search_list      = ["terraform.nutanix.com", "terraform.unit.test.com"]
// }

data "nutanix_project" "test" {
	project_id = nutanix_project.test.id
}
`, r, r)
}

func testAccProjectDataSourceConfigName(r int) string {
	return fmt.Sprintf(`
data "nutanix_clusters" "clusters" {}

locals {
	cluster1 = [
	for cluster in data.nutanix_clusters.clusters.entities :
	cluster.metadata.uuid if cluster.service_list[0] != "PRISM_CENTRAL"
	][0]
}

// resource "nutanix_project" "test" {
// 	name = "dou_vlan0_test_%d"
// 	cluster_uuid = local.cluster1
// 	vlan_id = %d
// 	project_type = "VLAN"

// 	prefix_length = 24
// 	default_gateway_ip = "192.168.0.1"
// 	project_ip = "192.168.0.0"
// 	ip_config_pool_list_ranges = ["192.168.0.10 192.168.0.100"]

// 	dhcp_options = {
// 		boot_file_name   = "bootfile"
// 		domain_name      = "nutanix"
// 		tftp_server_name = "10.250.140.200"
// 	}
	
// 	dhcp_domain_name_server_list = ["8.8.8.8", "4.2.2.2"]
// 	dhcp_domain_search_list      = ["terraform.nutanix.com", "terraform.unit.test.com"]
// }

data "nutanix_project" "test" {
	project_name = nutanix_project.test.name
}
`, r, r)
}

func testAccProjectDataSourceBadConfig() string {
	return `
data "nutanix_project" "test" {
	project_id   = "test-project-id"
	project_name = "test-project-name"
}
`
}

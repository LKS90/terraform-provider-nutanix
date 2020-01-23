package nutanix

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	v3 "github.com/terraform-providers/terraform-provider-nutanix/client/v3"
	"github.com/terraform-providers/terraform-provider-nutanix/utils"
)

func dataSourceNutanixProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNutanixProjectRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"project_name"},
			},
			"project_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"project_id"},
			},
			"api_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"last_update_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"creation_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"spec_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"spec_hash": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"categories": categoriesSchema(),
			"owner_reference": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"project_reference": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"availability_zone_reference": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"message_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reason": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"details": {
							Type:     schema.TypeMap,
							Computed: true,
						},
					},
				},
			},
			"cluster_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cluster_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vswitch_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_gateway_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"prefix_length": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"project_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_server_address": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fqdn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dhcp_server_address_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ip_config_pool_list_ranges": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"dhcp_options": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"boot_file_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tftp_server_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dhcp_domain_name_server_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"dhcp_domain_search_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"network_function_chain_reference": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func findProjectByUUID(conn *v3.Client, uuid string) (*v3.ProjectIntentResponse, error) {
	return conn.V3.GetProject(uuid)
}

func findProjectByName(conn *v3.Client, name string) (*v3.ProjectIntentResponse, error) {
	resp, err := conn.V3.ListAllProject()
	if err != nil {
		return nil, err
	}

	entities := resp.Entities

	found := make([]*v3.ProjectIntentResponse, 0)
	for _, v := range entities {
		if *v.Spec.Name == name {
			found = append(found, v)
		}
	}

	if len(found) > 1 {
		return nil, fmt.Errorf("your query returned more than one result. Please use project_id argument instead")
	}

	if len(found) == 0 {
		return nil, fmt.Errorf("project with the given name, not found")
	}

	return found[0], nil

}

func dataSourceNutanixProjectRead(d *schema.ResourceData, meta interface{}) error {
	// Get client connection
	conn := meta.(*Client).API

	projectID, iok := d.GetOk("project_id")
	projectName, nok := d.GetOk("project_name")

	if !iok && !nok {
		return fmt.Errorf("please provide one of project_id or project_name attributes")
	}

	var reqErr error
	var resp *v3.ProjectIntentResponse

	if iok {
		resp, reqErr = findProjectByUUID(conn, projectID.(string))
	} else {
		resp, reqErr = findProjectByName(conn, projectName.(string))
	}

	if reqErr != nil {
		return reqErr
	}

	m, c := setRSEntityMetadata(resp.Metadata)

	if err := d.Set("metadata", m); err != nil {
		return err
	}
	if err := d.Set("categories", c); err != nil {
		return err
	}
	if err := d.Set("project_reference", flattenReferenceValues(resp.Metadata.ProjectReference)); err != nil {
		return err
	}
	if err := d.Set("owner_reference", flattenReferenceValues(resp.Metadata.OwnerReference)); err != nil {
		return err
	}
	if err := d.Set("availability_zone_reference", flattenReferenceValues(resp.Status.AvailabilityZoneReference)); err != nil {
		return err
	}
	if err := flattenClusterReference(resp.Status.ClusterReference, d); err != nil {
		return err
	}

	dgIP := ""
	sIP := ""
	pl := int64(0)
	port := int64(0)
	dhcpSA := make(map[string]interface{})
	dOptions := make(map[string]interface{})
	ipcpl := make([]string, 0)
	dnsList := make([]string, 0)
	dsList := make([]string, 0)

	if resp.Status.Resources.IPConfig != nil {
		dgIP = utils.StringValue(resp.Status.Resources.IPConfig.DefaultGatewayIP)
		pl = utils.Int64Value(resp.Status.Resources.IPConfig.PrefixLength)
		sIP = utils.StringValue(resp.Status.Resources.IPConfig.ProjectIP)

		if resp.Status.Resources.IPConfig.DHCPServerAddress != nil {
			dhcpSA["ip"] = utils.StringValue(resp.Status.Resources.IPConfig.DHCPServerAddress.IP)
			dhcpSA["fqdn"] = utils.StringValue(resp.Status.Resources.IPConfig.DHCPServerAddress.FQDN)
			dhcpSA["ipv6"] = utils.StringValue(resp.Status.Resources.IPConfig.DHCPServerAddress.IPV6)
			port = utils.Int64Value(resp.Status.Resources.IPConfig.DHCPServerAddress.Port)
		}

		if resp.Status.Resources.IPConfig.PoolList != nil {
			pl := resp.Status.Resources.IPConfig.PoolList
			poolList := make([]string, len(pl))
			for k, v := range pl {
				poolList[k] = utils.StringValue(v.Range)
			}
			ipcpl = poolList
		}
		if resp.Status.Resources.IPConfig.DHCPOptions != nil {
			dOptions["boot_file_name"] = utils.StringValue(resp.Status.Resources.IPConfig.DHCPOptions.BootFileName)
			dOptions["domain_name"] = utils.StringValue(resp.Status.Resources.IPConfig.DHCPOptions.DomainName)
			dOptions["tftp_server_name"] = utils.StringValue(resp.Status.Resources.IPConfig.DHCPOptions.TFTPServerName)

			if resp.Status.Resources.IPConfig.DHCPOptions.DomainNameServerList != nil {
				dnsList = utils.StringValueSlice(resp.Status.Resources.IPConfig.DHCPOptions.DomainNameServerList)
			}
			if resp.Status.Resources.IPConfig.DHCPOptions.DomainSearchList != nil {
				dsList = utils.StringValueSlice(resp.Status.Resources.IPConfig.DHCPOptions.DomainSearchList)
			}
		}
	}

	if err := d.Set("dhcp_server_address", dhcpSA); err != nil {
		return err
	}
	if err := d.Set("ip_config_pool_list_ranges", ipcpl); err != nil {
		return err
	}
	if err := d.Set("dhcp_options", dOptions); err != nil {
		return err
	}
	if err := d.Set("dhcp_domain_name_server_list", dnsList); err != nil {
		return err
	}
	if err := d.Set("dhcp_domain_search_list", dsList); err != nil {
		return err
	}

	d.Set("cluster_reference_name", utils.StringValue(resp.Status.ClusterReference.Name))
	d.Set("api_version", utils.StringValue(resp.APIVersion))
	d.Set("name", utils.StringValue(resp.Status.Name))
	d.Set("description", utils.StringValue(resp.Status.Description))
	d.Set("state", utils.StringValue(resp.Status.State))
	d.Set("vswitch_name", utils.StringValue(resp.Status.Resources.VswitchName))
	d.Set("project_type", utils.StringValue(resp.Status.Resources.ProjectType))
	d.Set("default_gateway_ip", dgIP)
	d.Set("prefix_length", pl)
	d.Set("project_ip", sIP)
	d.Set("dhcp_server_address_port", port)
	d.Set("vlan_id", utils.Int64Value(resp.Status.Resources.VlanID))
	d.Set("network_function_chain_reference", flattenReferenceValues(resp.Status.Resources.NetworkFunctionChainReference))

	d.SetId(utils.StringValue(resp.Metadata.UUID))

	return nil
}

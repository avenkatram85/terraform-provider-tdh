---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tdh_network_policy Resource - tdh"
subcategory: ""
description: |-
  Represents a policy on TDH.
---

# tdh_network_policy (Resource)

Represents a policy on TDH.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the policy
- `network_spec` (Attributes) Network config to allow access to service resource. (see [below for nested schema](#nestedatt--network_spec))
- `service_type` (String) Type of policy to manage. Supported values is:  `NETWORK`.

### Read-Only

- `id` (String) Auto-generated ID of the policy after creation, and can be used to import it from TDH to terraform state.
- `resource_ids` (Set of String) IDs of service resources/instances being managed by the policy.

<a id="nestedatt--network_spec"></a>
### Nested Schema for `network_spec`

Required:

- `cidr` (String) CIDR value to allow access from. Ex: `10.45.66.80/30`
- `network_port_ids` (Set of String) IDs of network ports to open up for access. Please make use of datasource `tdh_network_ports` to get IDs of ports available for services.


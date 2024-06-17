---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tdh_data_planes Data Source - tdh"
subcategory: ""
description: |-
  Used to fetch all Data planes
---

# tdh_data_planes (Data Source)

Used to fetch all Data planes

## Example Usage

```terraform
data "tdh_dataplanes" "all" {
}
output "resp" {
  value = data.tdh_dataplanes.all
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `list` (Attributes List) (see [below for nested schema](#nestedatt--list))

### Read-Only

- `id` (String) The testing framework requires an id attribute to be present in every data source and resource

<a id="nestedatt--list"></a>
### Nested Schema for `list`

Read-Only:

- `account` (Attributes) (see [below for nested schema](#nestedatt--list--account))
- `auto_upgrade` (Boolean) Whether auto-upgrade is enabled on this data plane or not
- `backup_storage_policy` (String) Backup Storage Policy
- `certificate` (Attributes) (see [below for nested schema](#nestedatt--list--certificate))
- `created` (String) Creation time of this data plane.
- `data_plane_name` (String) Data plane Name
- `data_plane_on_control_plane` (Boolean) Whether Data plane is deployed on same cluster as TDH Control plane
- `data_plane_release_id` (String) Helm Version Id
- `data_plane_release_name` (String) Helm Version
- `default_policy_name` (String) Default Policy Name
- `id` (String) ID of the data plane.
- `modified` (String) Modified time of this data plane.
- `name` (String) Name of the TKC.
- `provider` (String) Provider Name
- `region` (String) Data plane Region
- `services` (Set of String) Services available on the data plane
- `shared` (Boolean) Whether this account is shared between multiple Organisations or not.
- `status` (String) Status of the data plane
- `storage_policies` (Set of String) Storage Policies of the data plane
- `tags` (Set of String) Tags set on this data plane.
- `version` (String) Data plane Version

<a id="nestedatt--list--account"></a>
### Nested Schema for `list.account`

Read-Only:

- `id` (String) ID of the cloud account.
- `name` (String) Name of the cloud account.


<a id="nestedatt--list--certificate"></a>
### Nested Schema for `list.certificate`

Read-Only:

- `id` (String) ID of the certificate.
- `name` (String) Name of the certificate.


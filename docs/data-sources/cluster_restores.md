---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tdh_cluster_restores Data Source - tdh"
subcategory: ""
description: |-
  Used to fetch all the backups available for the service type on TDH.
---

# tdh_cluster_restores (Data Source)

Used to fetch all the backups available for the service type on TDH.



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) Restore of the ID
- `list` (Attributes List) List of restores. (see [below for nested schema](#nestedatt--list))

<a id="nestedatt--list"></a>
### Nested Schema for `list`

Required:

- `service_type` (String) Service Type of the cluster.

Optional:

- `id` (String) Restore ID

Read-Only:

- `backup_id` (String) Backup ID
- `backup_name` (String) Backup name
- `dataplane_id` (String) Dataplane ID
- `name` (String) Name of the Backup.
- `target_instance` (String) Target instance
- `target_instance_name` (String) Target Instance Name


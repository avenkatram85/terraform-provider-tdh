---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tdh_cluster_backup Resource - tdh"
subcategory: ""
description: |-
  This is used to create backup (and restore a backup) of a database service cluster like POSTGRES, MYSQL, REDIS.
  Note: To restore a backup, either create a backup or import by ID.
---

# tdh_cluster_backup (Resource)

This is used to create backup (and restore a backup) of a database service cluster like `POSTGRES`, `MYSQL`, `REDIS`.
**Note:** To restore a backup, either create a backup or import by ID.

## Example Usage

```terraform
# to create cluster backup
resource "tdh_cluster_backup" "example" {
  name        = "tf-first-backup"
  description = "my postgres cluster backup using TDH provider"
  cluster_id  = "CLUSTER_ID" # ID of the cluster to backup, use datasource "tdh_clusters" to see available clusters

  // non editable fields
  lifecycle {
    ignore_changes = [name]
  }
}

# to restore a backup, either use same created resource or import a backup using backup ID. Then initialize the restore config like so:
resource "tdh_cluster_backup" "example" {
  name       = "created-backup" # if importing, name has to match imported backup name
  cluster_id = "CLUSTER_ID"     # if importing, ID has to match with existing state

  restore = {
    cluster_name       = "tf-restore-6"
    # this will be the name of cluster that will be created with this restore. Not Applicable for "REDIS"
    storage_policy     = "tdh-k8s-cluster-policy"
    # can get using datasource "tdh_storage_policies". Not Applicable for "REDIS"
    network_policy_ids = [
      # can get using datasource "tdh_network_policies"
      # Not Applicable for "REDIS"
      "6ad0cf49-81be-48e3-bab4-2a13b9de0c95"
    ]
  }
  // non editable fields
  lifecycle {
    ignore_changes = [name]
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_id` (String) ID of the cluster.
- `name` (String) Name of the backup.

### Optional

- `description` (String) Description for the create backup.
- `metadata` (Attributes) The metadata of the backup. (see [below for nested schema](#nestedatt--metadata))
- `restore` (Attributes) Use it to restore this backup.
**Note:** Just declare it as empty block in case of `REDIS` cluster backup since in case of Redis, restore happens on same cluster i.e. the cluster has to be present and there will be some downtime. (see [below for nested schema](#nestedatt--restore))

### Read-Only

- `backup_trigger_type` (String) The type of trigger for the backup.
- `cluster_name` (String) Name of the cluster.
- `cluster_version` (String) Version of the cluster.
- `data_plane_id` (String) The ID of the data plane.
- `generated_name` (String) Name that was automatically generated for this backup.
- `id` (String) ID of the backup.
- `org_id` (String) ID of the org this backup belongs to.
- `provider_name` (String) The provider of the cluster.
- `region` (String) The region of the cluster.
- `service_type` (String) Service Type of the cluster.
- `size` (String) Size of the cluster.
- `status` (String) Status of the backup.
- `time_completed` (String) Time when the backup was completed.
- `time_started` (String) Time when the backup was initiated.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Read-Only:

- `backup_location` (String) Backup Location
- `cluster_name` (String) Name of the cluster.
- `cluster_size` (String) Size of the Instance Type.
- `databases` (Set of String) List of databases part of backup.
- `extensions` (Attributes Set) List of extensions part of backup. *(Specific to service `POSTGRES`)* (see [below for nested schema](#nestedatt--metadata--extensions))

<a id="nestedatt--metadata--extensions"></a>
### Nested Schema for `metadata.extensions`

Read-Only:

- `name` (String) Name of the extension.
- `version` (String) Version of the extension.



<a id="nestedatt--restore"></a>
### Nested Schema for `restore`

Optional:

- `cluster_name` (String) Name of the target instance.
- `network_policy_ids` (Set of String) List of network policy IDs for network configuration on the instance.
- `storage_policy` (String) Name of the storage policy.
- `tags` (Set of String) List of tags to set on the instance.

## Import

Import is supported using the following syntax:

```shell
# Cluster Backup can be imported by specifying the UUID.
terraform import tdh_cluster_backup.example d3c49288-7b17-4e78-a6af-257b49e34e53
```

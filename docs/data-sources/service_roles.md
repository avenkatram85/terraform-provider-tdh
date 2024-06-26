---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tdh_service_roles Data Source - tdh"
subcategory: ""
description: |-
  Used to fetch all roles applicable for services on TDH.
---

# tdh_service_roles (Data Source)

Used to fetch all roles applicable for services on TDH.

## Example Usage

```terraform
data "tdh_service_roles" "pg" {
  type = "POSTGRES"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `type` (String) Type of the service on TDH.

### Read-Only

- `id` (String) The testing framework requires an id attribute to be present in every data source and resource
- `list` (Attributes List) List of service roles. (see [below for nested schema](#nestedatt--list))

<a id="nestedatt--list"></a>
### Nested Schema for `list`

Read-Only:

- `description` (String) Description of the role.
- `name` (String) Name of the role.
- `permission_id` (String) ID of the permission for the role.
- `role_id` (String) ID of the role.
- `type` (String) Type of the role.



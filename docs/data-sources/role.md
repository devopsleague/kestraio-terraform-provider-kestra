---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kestra_role Data Source - terraform-provider-kestra"
subcategory: ""
description: |-
  Use this data source to access information about an existing Kestra Role.
  -> This resource is only available on the Enterprise Edition https://kestra.io/enterprise
---

# kestra_role (Data Source)

Use this data source to access information about an existing Kestra Role.

-> This resource is only available on the [Enterprise Edition](https://kestra.io/enterprise)

## Example Usage

```terraform
data "kestra_role" "example" {
  role_id = "3kcvnr27ZcdHXD2AUvIe7z"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `role_id` (String) The role.

### Optional

- `is_default` (Boolean) The role is the default one at user creation. Only one role can be default. Latest create/update to true will be keep as default. Defaults to `false`.
- `namespace` (String) The linked namespace.

### Read-Only

- `description` (String) The role description.
- `id` (String) The ID of this resource.
- `name` (String) The role name.
- `permissions` (Set of Object) The role permissions. (see [below for nested schema](#nestedatt--permissions))
- `tenant_id` (String) The tenant id.

<a id="nestedatt--permissions"></a>
### Nested Schema for `permissions`

Read-Only:

- `permissions` (List of String)
- `type` (String)

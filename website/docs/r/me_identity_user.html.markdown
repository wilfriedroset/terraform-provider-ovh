---
layout: "ovh"
page_title: "OVH: ovh_me_identity_user"
sidebar_current: "docs-ovh-resource-me-identity-user"
description: |-
  Creates an identity user.
---

# ovh_me_identity_user

Creates an identity user.

## Example Usage

```hcl
resource "ovh_me_identity_user" "my_user" {
  description = "Some custom description"
  email       = "my_login@example.com"
  group       = "DEFAULT"
  login       = "my_login"
  password    = "super-s3cr3t!password"
}
```

## Argument Reference

* `description` - User description.
* `email` - User's email.
* `group` - User's group.
* `login` - User's login suffix.
* `password` - User's password.

## Attributes Reference

* `creation` - Creation date of this user.
* `last_update` - Last update of this user.
* `password_last_update` - When the user changed his password for the last time.
* `status` - Current user's status.

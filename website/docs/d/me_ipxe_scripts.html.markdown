---
layout: "ovh"
page_title: "OVH: me_ipxe_scripts"
sidebar_current: "docs-ovh-datasource-ipxe-scripts"
description: |-
  Get the list of the IPXE Scripts of the account.
---

# ovh_me_ipxe_scripts (Data Source)

Use this data source to retrieve a list of the names of the account's IPXE Scripts.

## Example Usage

```hcl
data "ovh_me_ipxe_scripts" "scripts" {}
```

## Argument Reference

This datasource takes no arguments.

## Attributes Reference

* `result` - The list of the names of all the IPXE Scripts.

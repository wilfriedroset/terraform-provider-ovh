---
layout: "ovh"
page_title: "OVH: cloud_project_kube_nodepool"
sidebar_current: "docs-ovh-resource-cloud-project-kube-nodepool-x"
description: |-
  Creates a nodepool in a kubernetes managed cluster.
---

# ovh_cloud_project_kube_nodepool

Creates a nodepool in a OVHcloud Managed Kubernetes Service cluster.

## Example Usage

```hcl
resource "ovh_cloud_project_kube_nodepool" "pool" {
  service_name  = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  kube_id       = "xxxxxxxx-2bf9-xxxx-xxxx-xxxxxxxxxxxx"
  name          = "my-pool"
  flavor_name   = "b2-7"
  desired_nodes = 3
  max_nodes     = 3
  min_nodes     = 3
  template {
    metadata {
      annotations = {
        k1 = "v1"
        k2 = "v2"
      }
      finalizers = ["F1", "F2"]
      labels = {
        k3 = "v3"
        k4 = "v4"
      }
    }
    spec {
      unschedulable = false
      taints = [
        {
          effect = "PreferNoSchedule"
          key    = "k"
          value  = "v"
        }
      ]
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `service_name` - (Optional) The id of the public cloud project. If omitted,
    the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.

* `kube_id` - The id of the managed kubernetes cluster.

* `name` - (Optional) The name of the nodepool.
   Changing this value recreates the resource.
   Warning: "_" char is not allowed!

* `flavor_name` - a valid OVHcloud public cloud flavor ID in which the nodes will be started.
   Ex: "b2-7". Changing this value recreates the resource.
   You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/

* `desired_nodes` - number of nodes to start.

* `max_nodes` - maximum number of nodes allowed in the pool.
   Setting `desired_nodes` over this value will raise an error.

* `min_nodes` - minimum number of nodes allowed in the pool.
   Setting `desired_nodes` under this value will raise an error.

* `monthly_billed` - (Optional) should the nodes be billed on a monthly basis. Default to `false`.

* `anti_affinity` - (Optional) should the pool use the anti-affinity feature. Default to `false`.

* `autoscale` - (Optional) Enable auto-scaling for the pool. Default to `false`.

* `template ` - (Optional) Managed Kubernetes nodepool template, which is a complex object constituted by two main nested objects:
  * metadata (Optional) Metadata of each nodes in the pool
    * annotations (Optional) Annotations to apply to each nodes
    * finalizers (Optional) Finalizers to apply to each nodes
    * labels (Optional) Labels to apply to each nodes
  * spec (Optional) Spec of each nodes in the pool
    * taints (Optional) Taints to apply to each nodes
    * unschedulable (Optional) If true, set nodes as un-schedulable

## Attributes Reference

In addition, the following attributes are exported:

* `available_nodes` - Number of nodes which are actually ready in the pool
* `created_at` - Creation date
* `current_nodes` - Number of nodes present in the pool
* `desired_nodes` - Number of nodes you desire in the pool
* `flavor` - Flavor name
* `project_id` - Project id
* `size_status` - Status describing the state between number of nodes wanted and available ones
* `status` - Current status
* `up_to_date_nodes` - Number of nodes with latest version installed in the pool
* `updated_at` - Last update date

## Import

OVHcloud Managed Kubernetes Service cluster node pool can be imported using the `service_name`, the `id` of the cluster, and the `id` of the nodepool separated by "/" E.g.,

```bash
$ terraform import ovh_cloud_project_kube_nodepool.pool service_name/kube_id/poolid
```

---
layout: "rancher2"
page_title: "Rancher2: rancher2_node_template_ubicloud"
sidebar_current: "docs-rancher2-resource-node_template_ubicloud"
description: |-
  Provides a Rancher v2 UbiCloud Node Template resource. This can be used to create UbiCloud Node Template for rancher v2 rke clusters and retrieve their information.
---

# rancher2\_node\_template\_ubicloud

Provides a Rancher v2 UbiCloud Node Template resource. This can be used to create UbiCloud Node Template for rancher v2 rke clusters and retrieve their information.

## Example Usage

```hcl
# Create a new rancher2 UbiCloud Node Template
resource "rancher2_node_template_ubicloud" "foo" {
    description = "My node template description"
    engine_env = {
        foo = "env"
    }
    engine_insecure_registry = []
    engine_install_url = "https://releases.rancher.com/install-docker/17.03.sh"
    engine_label = {}
    engine_opt = {}
    engine_registry_mirror = []
    engine_storage_driver = "overlay"
    name = "my-node-template"
    ubicloud_config = {
        active_timeout = "200"
        auth_url = "https://authurl.foo.com/v3"
        availability_zone = "us-east-1"
        cacert = "ca-cert"
        config_drive = false
        domain_name = "default"
        endpoint_type = "endpoint-type"
        flavor_name = "ug2"
        floatingip_pool = "floatingip-pool"
        image_name = "image-name"
        insecure = false
        ip_version = "4"
        keypair_name = "keypair-name"
        net_name = "netname"
        nodepool_anti_affinity = true
        nova_network = false
        private_key_file = "/tmp/my_key.pem"
        region = "east"
        sec_groups = "security-group"
        ssh_port = "22"
        ssh_user = "root"
        tenant_name = "tenant-name"
        user_data_file = "user-data-file"
        username = "user"
    }
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Optional) Node template description
* `engine_env` - (Optional) Docker engine environment variables
* `engine_insecure_registry` - (Optional) Docker engine insecure registries
* `engine_install_url` - (Required) Docker install url
* `engine_label` - (Optional) Docker engine labels
* `engine_opt` - (Optional) Docker engine options
* `engine_registry_mirror` - (Optional) Docker engine insecure registry domains
* `engine_storage_driver` - (Optional) Docker engine storage driver
* `name` - (Required) Name of the template.
* `ubicloud_config` - (Required) Nested attributes containing the UbiCloud instances attributes
    * `active-timeout`- OpenStack active timeout (default: 200)
    * `auth-url` - (Required) OpenStack authentication URL
    * `availability-zone` - (Required) OpenStack availability zone
    * `cacert` - (Optional) CA certificate bundle to verify against
    * `config-drive` - (Optional) Enables the OpenStack config drive for the instance (Default: false)
    * `domain-id` - (Required*) OpenStack domain ID (identity v3 only),
    * `domain-name` - (Required*) OpenStack domain name (identity v3 only)
    * `endpoint-type` - (Optional) OpenStack endpoint type (adminURL, internalURL or publicURL)
    * `flavor-id` - (Required*) OpenStack flavor id to use for the instance
    * `flavor-name` - (Required*) OpenStack flavor name to use for the instance
    * `floatingip-pool` - (Optional) OpenStack floating IP pool to get an IP from to assign to the instance
    * `image-id` - (Required*) OpenStack image id to use for the instance
    * `image-name` - (Required*) OpenStack image name to use for the instance
    * `insecure` - (Optional) Disable TLS credential checking.
    * `ip-version` - (Optional) OpenStack version of IP address assigned for the machine (default: 4)
    * `keypair-name` - (Optional) OpenStack keypair to use to SSH to the instance
    * `net-id` - (Required*) OpenStack network id the machine will be connected on
    * `net-name` - (Required*) OpenStack network name the machine will be connected on
    * `nodepool-anti-affinity` - (Optional) Enables anti-affinity for node in the same nodepool (Default: true)
    * `nova-network` - (Optional) Use the nova networking services instead of neutron.
    * `password` - (Required) OpenStack password
    * `private-key-file` - (Optional) Private keyfile to use for SSH (absolute path)
    * `region` - (Required) OpenStack region name
    * `sec-groups` - (Optional) OpenStack comma separated security groups for the machine
    * `ssh-port` - (Optional) OpenStack SSH port * (default: 22),
    * `ssh-user` - (Optional) OpenStack SSH user * (default: root),
    * `tenant-id` - (Required*) OpenStack tenant id
    * `tenant-name` - (Required*) OpenStack tenant name
    * `user-data-file` - (Optional) File containing an openstack userdata script
    * `username` - (Required) OpenStack username

> **Note**: `Required*` denotes that either the _name or _id is required but you cannot use both.

[UKS node template documentation](https://docs.uks.ubisoft.org/guide/node-templates/openstack/)

## Attributes Reference

The following attributes are exported:

* `id` - (Computed) The ID of the resource.
* `driver` - (Computed) Always `ubicloud`

## Import

UbiCloud Node Template can be imported using the rancher Node Template ID

```
$ terraform import rancher2_node_template_ubicloud.foo <node_template_id>
```

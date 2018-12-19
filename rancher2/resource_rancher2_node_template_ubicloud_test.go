package rancher2

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	testAccRancher2NodeTemplateType   = "rancher2_node_template_ubicloud"
	testAccRancher2NodeTemplateConfig = `
resource "rancher2_node_template_ubicloud" "foo" {
	description = "node template description"
	engine_env = {
		foo = "env"
	}
	engine_insecure_registry = [
		"http://docker.com"
	]
	engine_install_url = "https://docker.com/install.sh"
	engine_label = {
		foo = "label"
	}
	engine_opt = {
		foo = "opt"
	}
	engine_registry_mirror = [
		"https://mirror"
	]
	engine_storage_driver = "overlay"
	name = "node-template-foo"
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
`
	// TODO: checks for conflicting _id vs _name

	testAccRancher2NodeTemplateUpdateConfig = `
resource "rancher2_node_template_ubicloud" "foo" {
	description = "node template description-updated"
	engine_env = {
		foo = "env-updated"
	}
	engine_insecure_registry = [
		"http://docker.com/updated"
	]
	engine_install_url = "https://docker.com/install.sh.updated"
	engine_label = {
		foo = "label-updated"
	}
	engine_opt = {
		foo = "opt-updated"
	}
	engine_registry_mirror = [
		"https://mirror/updated"
	]
	engine_storage_driver = "overlay-updated"
	name = "node-template-foo-updated"
	ubicloud_config = {
		active_timeout = "200-updated"
		auth_url = "https://authurl.foo.com/v3/updated"
		availability_zone = "us-east-1-updated"
		cacert = "ca-cert-updated"
		config_drive = true
		domain_name = "default-updated"
		endpoint_type = "endpoint-type-updated"
		flavor_name = "ug2-updated"
		floatingip_pool = "floatingip-pool-updated"
		image_name = "image-name-updated"
		insecure = true
		ip_version = "4-updated"
		keypair_name = "keypair-name-updated"
		net_name = "netname-updated"
		nodepool_anti_affinity = false
		nova_network = true
		private_key_file = "/tmp/my_key.pem-updated"
		region = "east-updated"
		sec_groups = "security-group-updated"
		ssh_port = "23"
		ssh_user = "root-updated"
		tenant_name = "tenant-name-updated"
		user_data_file = "user-data-file-updated"
		username = "user-updated"
	}
}
`
)

func TestAccRancher2NodeTemplate_basic(t *testing.T) {
	var nodeTemplate *NodeTemplate
	name := testAccRancher2NodeTemplateType + ".foo"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRancher2NodeTemplateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRancher2NodeTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRancher2NodeTemplateExists(name, nodeTemplate),
					resource.TestCheckResourceAttr(name, "description", "node template description"),
					resource.TestCheckResourceAttr(name, "driver", UbiCloudDriver),
					resource.TestCheckResourceAttr(name, "engine_env.foo", "env"),
					resource.TestCheckResourceAttr(name, "engine_insecure_registry.0", "http://docker.com"),
					resource.TestCheckResourceAttr(name, "engine_install_url", "https://docker.com/install.sh"),
					resource.TestCheckResourceAttr(name, "engine_label.foo", "label"),
					resource.TestCheckResourceAttr(name, "engine_opt.foo", "opt"),
					resource.TestCheckResourceAttr(name, "engine_registry_mirror.0", "https://mirror"),
					resource.TestCheckResourceAttr(name, "engine_storage_driver", "overlay"),
					resource.TestCheckResourceAttr(name, "name", "node-template-foo"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.active_timeout", "200"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.auth_url", "https://authurl.foo.com/v3"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.availability_zone", "us-east-1"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.cacert", "ca-cert"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.config_drive", "false"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.domain_name", "default"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.endpoint_type", "endpoint-type"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.flavor_name", "ug2"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.floatingip_pool", "floatingip-pool"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.image_name", "image-name"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.insecure", "false"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.ip_version", "4"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.keypair_name", "keypair-name"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.net_name", "netname"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.nodepool_anti_affinity", "true"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.nova_network", "false"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.private_key_file", "/tmp/my_key.pem"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.region", "east"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.sec_groups", "security-group"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.ssh_port", "22"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.ssh_user", "root"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.tenant_name", "tenant-name"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.user_data_file", "user-data-file"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.username", "user"),
				),
			},
			resource.TestStep{
				Config: testAccRancher2NodeTemplateUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRancher2NodeTemplateExists(name, nodeTemplate),
					resource.TestCheckResourceAttr(name, "description", "node template description-updated"),
					resource.TestCheckResourceAttr(name, "driver", UbiCloudDriver),
					resource.TestCheckResourceAttr(name, "engine_env.foo", "env-updated"),
					resource.TestCheckResourceAttr(name, "engine_insecure_registry.0", "http://docker.com/updated"),
					resource.TestCheckResourceAttr(name, "engine_install_url", "https://docker.com/install.sh.updated"),
					resource.TestCheckResourceAttr(name, "engine_label.foo", "label-updated"),
					resource.TestCheckResourceAttr(name, "engine_opt.foo", "opt-updated"),
					resource.TestCheckResourceAttr(name, "engine_registry_mirror.0", "https://mirror/updated"),
					resource.TestCheckResourceAttr(name, "engine_storage_driver", "overlay-updated"),
					resource.TestCheckResourceAttr(name, "name", "node-template-foo-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.active_timeout", "200-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.auth_url", "https://authurl.foo.com/v3/updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.availability_zone", "us-east-1-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.cacert", "ca-cert-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.config_drive", "true"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.domain_name", "default-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.endpoint_type", "endpoint-type-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.flavor_name", "ug2-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.floatingip_pool", "floatingip-pool-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.image_name", "image-name-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.insecure", "true"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.ip_version", "4-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.keypair_name", "keypair-name-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.net_name", "netname-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.nodepool_anti_affinity", "false"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.nova_network", "true"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.private_key_file", "/tmp/my_key.pem-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.region", "east-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.sec_groups", "security-group-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.ssh_port", "23"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.ssh_user", "root-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.tenant_name", "tenant-name-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.user_data_file", "user-data-file-updated"),
					resource.TestCheckResourceAttr(name, "ubicloud_config.0.username", "user-updated"),
				),
			},
		},
	})
}

// func TestAccRancher2NodeTemplate_disappears(t *testing.T) {
// 	var nodeTemplate *managementClient.NodeTemplate

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckRancher2NodeTemplateDestroy,
// 		Steps: []resource.TestStep{
// 			resource.TestStep{
// 				Config: testAccRancher2NodeTemplateConfig,
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckRancher2NodeTemplateExists(testAccRancher2NodeTemplateType+".foo", nodeTemplate),
// 					testAccRancher2NodeTemplateDisappears(nodeTemplate),
// 				),
// 				ExpectNonEmptyPlan: true,
// 			},
// 		},
// 	})
// }

// func testAccRancher2NodeTemplateDisappears(nodeTemplate *managementClient.NodeTemplate) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		for _, rs := range s.RootModule().Resources {
// 			if rs.Type != testAccRancher2NodeTemplateType {
// 				continue
// 			}
// 			client, err := testAccProvider.Meta().(*Config).ManagementClient()
// 			if err != nil {
// 				return err
// 			}

// 			nodeTemplate, err := client.NodeTemplate.ByID(rs.Primary.ID)
// 			if err != nil {
// 				if IsNotFound(err) {
// 					return nil
// 				}
// 				return err
// 			}

// 			err = client.NodeTemplate.Delete(nodeTemplate)
// 			if err != nil {
// 				return fmt.Errorf("Error removing Node Driver: %s", err)
// 			}
// 		}
// 		return nil
// 	}
// }

func testAccCheckRancher2NodeTemplateExists(n string, nodeTemplate *NodeTemplate) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Node Template ID is set")
		}

		client, err := testAccProvider.Meta().(*Config).ManagementClient()
		if err != nil {
			return err
		}

		nodeTemplateClient := NewNodeTemplateClient(client)
		foundNodeTemplate, err := nodeTemplateClient.ByID(rs.Primary.ID)
		if err != nil {
			if IsNotFound(err) {
				return fmt.Errorf("Node Driver not found")
			}
			return err
		}

		nodeTemplate = foundNodeTemplate

		return nil
	}
}

func testAccCheckRancher2NodeTemplateDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != testAccRancher2NodeTemplateType {
			continue
		}
		client, err := testAccProvider.Meta().(*Config).ManagementClient()
		if err != nil {
			return err
		}

		nodeTemplateClient := NewNodeTemplateClient(client)
		_, err = nodeTemplateClient.ByID(rs.Primary.ID)
		if err != nil {
			if IsNotFound(err) {
				return nil
			}
			return err
		}
		return fmt.Errorf("Node Driver still exists")
	}
	return nil
}

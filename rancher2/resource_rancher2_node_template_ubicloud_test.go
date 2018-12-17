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
    active = false
    builtin = false
    checksum = "0x0"
    description = "Foo description"
    external_id = "foo_external"
    name = "foo"
    ui_url = "local://ui"
    url = "local://"
    whitelist_domains = ["*.foo.com"]
}
`

// 	testAccRancher2NodeTemplateUpdateConfig = `
// resource "rancher2_node_template_ubicloud" "foo" {
//     active = true,
//     builtin = false
//     checksum = "0x1"
//     description= "Foo description - updated"
//     external_id = "external"
//     name = "foo"
//     ui_url = "local://ui/updated"
//     url = "local://updated"
//     whitelist_domains = ["*.foo.com", "updated.foo.com"]
// }
// `
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
					resource.TestCheckResourceAttr(name, "active", "false"),
					resource.TestCheckResourceAttr(name, "builtin", "false"),
					resource.TestCheckResourceAttr(name, "checksum", "0x0"),
					resource.TestCheckResourceAttr(name, "description", "Foo description"),
					resource.TestCheckResourceAttr(name, "external_id", "foo_external"),
					resource.TestCheckResourceAttr(name, "name", "foo"),
					resource.TestCheckResourceAttr(name, "ui_url", "local://ui"),
					resource.TestCheckResourceAttr(name, "url", "local://"),
					resource.TestCheckResourceAttr(name, "whitelist_domains.0", "*.foo.com"),
				),
			},
			// resource.TestStep{
			// 	Config: testAccRancher2NodeTemplateUpdateConfig,
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccCheckRancher2NodeTemplateExists(name, nodeTemplate),
			// 		resource.TestCheckResourceAttr(name, "active", "true"),
			// 		resource.TestCheckResourceAttr(name, "builtin", "false"),
			// 		resource.TestCheckResourceAttr(name, "checksum", "0x1"),
			// 		resource.TestCheckResourceAttr(name, "description", "Foo description - updated"),
			// 		resource.TestCheckResourceAttr(name, "external_id", "external"),
			// 		resource.TestCheckResourceAttr(name, "name", "foo"),
			// 		resource.TestCheckResourceAttr(name, "ui_url", "local://ui/updated"),
			// 		resource.TestCheckResourceAttr(name, "url", "local://updated"),
			// 		resource.TestCheckResourceAttr(name, "whitelist_domains.0", "*.foo.com"),
			// 		resource.TestCheckResourceAttr(name, "whitelist_domains.1", "updated.foo.com"),
			// 	),
			// },
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
			return fmt.Errorf("No Node Driver ID is set")
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

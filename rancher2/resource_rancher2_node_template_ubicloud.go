package rancher2

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

// Schema

func ubiCloudFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"active_timeout": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "200",
		},
		"auth_url": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"availability_zone": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"cacert": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"config_drive": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"domain_id": &schema.Schema{
			Type: schema.TypeString,
			// ConflictsWith: []string{"domain_name"},
			Optional: true,
		},
		"domain_name": &schema.Schema{
			Type: schema.TypeString,
			// ConflictsWith: []string{"domain_id"},
			Optional: true,
		},
		"endpoint_type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"flavor_id": &schema.Schema{
			Type: schema.TypeString,
			// ConflictsWith: []string{"flavor_name"},
			Optional: true,
		},
		"flavor_name": &schema.Schema{
			Type: schema.TypeString,
			// ConflictsWith: []string{"flavor_id"},
			Optional: true,
		},
		"floatingip_pool": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"image_id": &schema.Schema{
			Type: schema.TypeString,
			// ConflictsWith: []string{"image_name"},
			Optional: true,
		},
		"image_name": &schema.Schema{
			Type: schema.TypeString,
			// ConflictsWith: []string{"image_id"},
			Optional: true,
		},
		"insecure": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"ip_version": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "4",
		},
		"keypair_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"net_id": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"net_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"nodepool_anti_affinity": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"nova_network": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"password": &schema.Schema{
			Type:      schema.TypeString,
			Required:  true,
			Sensitive: true,
		},
		"private_key_file": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"region": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"sec_groups": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"ssh_port": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "22",
		},
		"ssh_user": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "root",
		},
		"tenant_id": &schema.Schema{
			Type: schema.TypeString,
			// ConflictsWith: []string{"tenant_name"},
			Optional: true,
		},
		"tenant_name": &schema.Schema{
			Type: schema.TypeString,
			// ConflictsWith: []string{"tenant_id"},
			Optional: true,
		},
		"user_data_file": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"username": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
	}
	return s
}

func nodeTemplateUbiCloudFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{

		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"driver": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		// TODO default to proper HTTP_PROXY env vars
		"engine_env": &schema.Schema{
			Type:     schema.TypeMap,
			Optional: true,
		},
		"engine_insecure_registry": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"engine_install_url": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			// Default:  "https://releases.rancher.com/install-docker/17.03.sh",
		},
		"engine_label": &schema.Schema{
			Type:     schema.TypeMap,
			Optional: true,
		},
		"engine_opt": &schema.Schema{
			Type:     schema.TypeMap,
			Optional: true,
		},
		"engine_registry_mirror": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"engine_storage_driver": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"ubicloud_config": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: ubiCloudFields(),
			},
		},
	}

	return s
}

// Flatteners
func flattenUbiCloudConfig(in *UbiCloudConfig) []interface{} {
	config := make(map[string]interface{})

	config["active_timeout"] = in.ActiveTimeout
	config["auth_url"] = in.AuthURL
	config["availability_zone"] = in.AvailabilityZone
	config["cacert"] = in.CaCert
	config["config_drive"] = in.ConfigDrive
	config["domain_id"] = in.DomainID
	config["domain_name"] = in.DomainName
	config["endpoint_type"] = in.EndpointType
	config["flavor_id"] = in.FlavorID
	config["flavor_name"] = in.FlavorName
	config["floatingip_pool"] = in.FloatingIPPool
	config["image_id"] = in.ImageID
	config["image_name"] = in.ImageName
	config["insecure"] = in.Insecure
	config["ip_version"] = in.IPVersion
	config["keypair_name"] = in.KeypairName
	config["net_id"] = in.NetID
	config["net_name"] = in.NetName
	config["nodepool_anti_affinity"] = in.NodePoolAntiAffinity
	config["nova_network"] = in.NovaNetwork
	config["password"] = in.Password
	config["private_key_file"] = in.PrivateKeyFile
	config["region"] = in.Region
	config["sec_groups"] = in.SecGroups
	config["ssh_port"] = in.SSHPort
	config["ssh_user"] = in.SSHUser
	config["tenant_id"] = in.TenantID
	config["tenant_name"] = in.TenantName
	config["user_data_file"] = in.UserDataFile
	config["username"] = in.Username

	out := make([]interface{}, 0, 0)
	return append(out, config)
}

func flattenNodeTemplateUbiCloud(d *schema.ResourceData, in *NodeTemplate) error {
	if in == nil {
		return nil
	}

	d.SetId(in.ID)

	err := d.Set("driver", UbiCloudDriver)
	if err != nil {
		return err
	}

	err = d.Set("description", in.Description)
	if err != nil {
		return err
	}

	err = d.Set("engine_env", toMapInterface(in.EngineEnv))
	if err != nil {
		return err
	}

	err = d.Set("engine_insecure_registry", toArrayInterface(in.EngineInsecureRegistry))
	if err != nil {
		return err
	}

	err = d.Set("engine_install_url", in.EngineInstallURL)
	if err != nil {
		return err
	}

	err = d.Set("engine_label", toMapInterface(in.EngineLabel))
	if err != nil {
		return err
	}

	err = d.Set("engine_opt", toMapInterface(in.EngineOpt))
	if err != nil {
		return err
	}

	err = d.Set("engine_registry_mirror", toArrayInterface(in.EngineRegistryMirror))
	if err != nil {
		return err
	}

	err = d.Set("engine_storage_driver", in.EngineStorageDriver)
	if err != nil {
		return err
	}

	err = d.Set("name", in.Name)
	if err != nil {
		return err
	}

	return d.Set("ubicloud_config", flattenUbiCloudConfig(in.UbiCloudConfig))
}

// Expanders

func expandUbiCloudConfig(c []interface{}) *UbiCloudConfig {
	obj := &UbiCloudConfig{}
	if len(c) == 0 || c[0] == nil {
		return obj
	}
	in := c[0].(map[string]interface{})

	if v, ok := in["active_timeout"].(string); ok && len(v) > 0 {
		obj.ActiveTimeout = v
	}
	if v, ok := in["auth_url"].(string); ok && len(v) > 0 {
		obj.AuthURL = v
	}
	if v, ok := in["availability_zone"].(string); ok && len(v) > 0 {
		obj.AvailabilityZone = v
	}
	if v, ok := in["cacert"].(string); ok && len(v) > 0 {
		obj.CaCert = v
	}
	if v, ok := in["config_drive"].(bool); ok {
		obj.ConfigDrive = v
	}
	if v, ok := in["domain_id"].(string); ok && len(v) > 0 {
		obj.DomainID = v
	}
	if v, ok := in["domain_name"].(string); ok && len(v) > 0 {
		obj.DomainName = v
	}
	if v, ok := in["endpoint_type"].(string); ok && len(v) > 0 {
		obj.EndpointType = v
	}
	if v, ok := in["flavor_id"].(string); ok && len(v) > 0 {
		obj.FlavorID = v
	}
	if v, ok := in["flavor_name"].(string); ok && len(v) > 0 {
		obj.FlavorName = v
	}
	if v, ok := in["floatingip_pool"].(string); ok && len(v) > 0 {
		obj.FloatingIPPool = v
	}
	if v, ok := in["ip_version"].(string); ok && len(v) > 0 {
		obj.IPVersion = v
	}
	if v, ok := in["image_id"].(string); ok && len(v) > 0 {
		obj.ImageID = v
	}
	if v, ok := in["image_name"].(string); ok && len(v) > 0 {
		obj.ImageName = v
	}
	if v, ok := in["insecure"].(bool); ok {
		obj.Insecure = v
	}
	if v, ok := in["ip_version"].(string); ok && len(v) > 0 {
		obj.IPVersion = v
	}
	if v, ok := in["keypair_name"].(string); ok && len(v) > 0 {
		obj.KeypairName = v
	}
	if v, ok := in["net_id"].(string); ok && len(v) > 0 {
		obj.NetID = v
	}
	if v, ok := in["net_name"].(string); ok && len(v) > 0 {
		obj.NetName = v
	}
	if v, ok := in["nodepool_anti_affinity"].(bool); ok {
		obj.NodePoolAntiAffinity = v
	}
	if v, ok := in["nova_network"].(bool); ok {
		obj.NovaNetwork = v
	}
	if v, ok := in["password"].(string); ok && len(v) > 0 {
		obj.Password = v
	}
	if v, ok := in["private_key_file"].(string); ok && len(v) > 0 {
		obj.PrivateKeyFile = v
	}
	if v, ok := in["region"].(string); ok && len(v) > 0 {
		obj.Region = v
	}
	if v, ok := in["sec_groups"].(string); ok && len(v) > 0 {
		obj.SecGroups = v
	}
	if v, ok := in["ssh_port"].(string); ok && len(v) > 0 {
		obj.SSHPort = v
	}
	if v, ok := in["ssh_user"].(string); ok && len(v) > 0 {
		obj.SSHUser = v
	}
	if v, ok := in["tenant_id"].(string); ok && len(v) > 0 {
		obj.TenantID = v
	}
	if v, ok := in["tenant_name"].(string); ok && len(v) > 0 {
		obj.TenantName = v
	}
	if v, ok := in["user_data_file"].(string); ok && len(v) > 0 {
		obj.UserDataFile = v
	}
	if v, ok := in["username"].(string); ok && len(v) > 0 {
		obj.Username = v
	}
	return obj
}

func expandNodeTemplateUbiCloud(in *schema.ResourceData) *NodeTemplate {
	if in == nil {
		return nil
	}
	obj := &NodeTemplate{
		Driver: UbiCloudDriver,
	}
	if v := in.Id(); len(v) > 0 {
		obj.ID = v
	}

	obj.Description = in.Get("description").(string)
	if v, ok := in.Get("engine_env").(map[string]interface{}); ok && len(v) > 0 {
		obj.EngineEnv = toMapString(v)
	}
	if v, ok := in.Get("engine_insecure_registry").([]interface{}); ok && len(v) > 0 {
		obj.EngineInsecureRegistry = toArrayString(v)
	}
	obj.EngineInstallURL = in.Get("engine_install_url").(string)
	if v, ok := in.Get("engine_label").(map[string]interface{}); ok && len(v) > 0 {
		obj.EngineLabel = toMapString(v)
	}
	if v, ok := in.Get("engine_opt").(map[string]interface{}); ok && len(v) > 0 {
		obj.EngineOpt = toMapString(v)
	}
	if v, ok := in.Get("engine_registry_mirror").([]interface{}); ok && len(v) > 0 {
		obj.EngineRegistryMirror = toArrayString(v)
	}
	obj.EngineStorageDriver = in.Get("engine_storage_driver").(string)
	obj.Name = in.Get("name").(string)

	obj.UbiCloudConfig = expandUbiCloudConfig(in.Get("ubicloud_config").([]interface{}))

	return obj
}

func resourceRancher2NodeTemplateUbiCloud() *schema.Resource {
	return &schema.Resource{
		Create: resourceRancher2NodeTemplateUbiCloudCreate,
		Delete: resourceRancher2NodeTemplateUbiCloudDelete,
		Importer: &schema.ResourceImporter{
			State: resourceRancher2NodeTemplateUbiCloudImport,
		},
		Read:   resourceRancher2NodeTemplateUbiCloudRead,
		Update: resourceRancher2NodeTemplateUbiCloudUpdate,

		Schema: nodeTemplateUbiCloudFields(),
	}
}

func resourceRancher2NodeTemplateUbiCloudCreate(d *schema.ResourceData, meta interface{}) error {
	nodeTemplate := expandNodeTemplateUbiCloud(d)
	log.Printf("[INFO] Creating Node Template %s", nodeTemplate.Name)
	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return err
	}

	nodeTemplateClient := NewNodeTemplateClient(client)
	newNodeTemplate, err := nodeTemplateClient.Create(nodeTemplate)
	if err != nil {
		return err
	}

	d.SetId(newNodeTemplate.ID)

	return resourceRancher2NodeTemplateUbiCloudRead(d, meta)
}

func resourceRancher2NodeTemplateUbiCloudRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Refreshing Node Template ID %s", d.Id())
	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return err
	}

	nodeTemplateClient := NewNodeTemplateClient(client)

	nodeTemplate, err := nodeTemplateClient.ByID(d.Id())
	if err != nil {
		if IsNotFound(err) {
			log.Printf("[INFO] Node Template ID %s not found", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	return flattenNodeTemplateUbiCloud(d, nodeTemplate)
}

func resourceRancher2NodeTemplateUbiCloudUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Updating Node Template ID %s", d.Id())
	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return err
	}
	nodeTemplateClient := NewNodeTemplateClient(client)

	nodeTemplate, err := nodeTemplateClient.ByID(d.Id())
	if err != nil {
		return err
	}

	updateUbiCloudConfig := func(c []interface{}) map[string]interface{} {
		config := make(map[string]interface{})
		if len(c) == 0 || c[0] == nil {
			return config
		}
		in := c[0].(map[string]interface{})

		config[UbiCloudConfigActiveTimeout] = in["active_timeout"].(string)
		config[UbiCloudConfigAuthURL] = in["auth_url"].(string)
		config[UbiCloudConfigAvailabilityZone] = in["availability_zone"].(string)
		config[UbiCloudConfigCaCert] = in["cacert"].(string)
		config[UbiCloudConfigConfigDrive] = in["config_drive"].(bool)
		config[UbiCloudConfigDomainID] = in["domain_id"].(string)
		config[UbiCloudConfigDomainName] = in["domain_name"].(string)
		config[UbiCloudConfigEndpointType] = in["endpoint_type"].(string)
		config[UbiCloudConfigFlavorID] = in["flavor_id"].(string)
		config[UbiCloudConfigFlavorName] = in["flavor_name"].(string)
		config[UbiCloudConfigFloatingIPPool] = in["floatingip_pool"].(string)
		config[UbiCloudConfigImageID] = in["image_id"].(string)
		config[UbiCloudConfigImageName] = in["image_name"].(string)
		config[UbiCloudConfigInsecure] = in["insecure"].(bool)
		config[UbiCloudConfigIPVersion] = in["ip_version"].(string)
		config[UbiCloudConfigKeypairName] = in["keypair_name"].(string)
		config[UbiCloudConfigNetID] = in["net_id"].(string)
		config[UbiCloudConfigNetName] = in["net_name"].(string)
		config[UbiCloudConfigNodePoolAntiAffinity] = in["nodepool_anti_affinity"].(bool)
		config[UbiCloudConfigNovaNetwork] = in["nova_network"].(bool)
		config[UbiCloudConfigPassword] = in["password"].(string)
		config[UbiCloudConfigPrivateKeyFile] = in["private_key_file"].(string)
		config[UbiCloudConfigRegion] = in["region"].(string)
		config[UbiCloudConfigSecGroups] = in["sec_groups"].(string)
		config[UbiCloudConfigSSHPort] = in["ssh_port"].(string)
		config[UbiCloudConfigSSHUser] = in["ssh_user"].(string)
		config[UbiCloudConfigTenantID] = in["tenant_id"].(string)
		config[UbiCloudConfigTenantName] = in["tenant_name"].(string)
		config[UbiCloudConfigUserDataFile] = in["user_data_file"].(string)
		config[UbiCloudConfigUsername] = in["username"].(string)

		return config
	}

	update := map[string]interface{}{
		NodeTemplateFieldDescription:            d.Get("description").(string),
		NodeTemplateFieldEngineEnv:              toMapString(d.Get("engine_env").(map[string]interface{})),
		NodeTemplateFieldEngineInsecureRegistry: toArrayString(d.Get("engine_insecure_registry").([]interface{})),
		NodeTemplateFieldEngineInstallURL:       d.Get("engine_install_url").(string),
		NodeTemplateFieldEngineLabel:            toMapString(d.Get("engine_label").(map[string]interface{})),
		NodeTemplateFieldEngineOpt:              toMapString(d.Get("engine_opt").(map[string]interface{})),
		NodeTemplateFieldEngineRegistryMirror:   toArrayString(d.Get("engine_registry_mirror").([]interface{})),
		NodeTemplateFieldEngineStorageDriver:    d.Get("engine_storage_driver").(string),
		NodeTemplateFieldName:                   d.Get("name").(string),
		NodeTemplateFieldUbiCloudConfig:         updateUbiCloudConfig(d.Get("ubicloud_config").([]interface{})),
	}

	_, err = nodeTemplateClient.Update(nodeTemplate, update)
	if err != nil {
		return err
	}

	return resourceRancher2NodeTemplateUbiCloudRead(d, meta)
}

func resourceRancher2NodeTemplateUbiCloudDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Deleting Node Template ID %s", d.Id())
	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return err
	}
	nodeTemplateClient := NewNodeTemplateClient(client)

	nodeTemplate, err := nodeTemplateClient.ByID(d.Id())
	if err != nil {
		if IsNotFound(err) {
			log.Printf("[INFO] Node Template ID %s not found", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	err = nodeTemplateClient.Delete(nodeTemplate)
	if err != nil {
		return fmt.Errorf("Error removing Node Template: %s", err)
	}
	d.SetId("")
	return nil
}

func resourceRancher2NodeTemplateUbiCloudImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[INFO] Importing Node Template ID %s", d.Id())
	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return []*schema.ResourceData{}, err
	}
	nodeTemplateClient := NewNodeTemplateClient(client)

	nodeTemplate, err := nodeTemplateClient.ByID(d.Id())
	if err != nil {
		return []*schema.ResourceData{}, err
	}

	err = flattenNodeTemplateUbiCloud(d, nodeTemplate)
	if err != nil {
		return []*schema.ResourceData{}, err
	}

	return []*schema.ResourceData{d}, nil
}

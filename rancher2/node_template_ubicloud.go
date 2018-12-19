package rancher2

import (
	"github.com/rancher/norman/types"
	managementClient "github.com/rancher/types/client/management/v3"
)

const (
	UbiCloudDriver = "ubicloud"

	UbiCloudConfigActiveTimeout        = "activeTimeout"
	UbiCloudConfigAuthURL              = "authUrl"
	UbiCloudConfigAvailabilityZone     = "availabilityZone"
	UbiCloudConfigCaCert               = "cacert"
	UbiCloudConfigConfigDrive          = "configDrive"
	UbiCloudConfigDomainID             = "domainId"
	UbiCloudConfigDomainName           = "domainName"
	UbiCloudConfigEndpointType         = "endpointType"
	UbiCloudConfigFlavorID             = "flavorId"
	UbiCloudConfigFlavorName           = "flavorName"
	UbiCloudConfigFloatingIPPool       = "floatingipPool"
	UbiCloudConfigImageID              = "imageId"
	UbiCloudConfigImageName            = "imageName"
	UbiCloudConfigInsecure             = "insecure"
	UbiCloudConfigIPVersion            = "ipVersion"
	UbiCloudConfigKeypairName          = "keypairName"
	UbiCloudConfigNetID                = "netId"
	UbiCloudConfigNetName              = "netName"
	UbiCloudConfigNodePoolAntiAffinity = "nodepoolAntiAffinity"
	UbiCloudConfigNovaNetwork          = "novaNetwork"
	UbiCloudConfigPrivateKeyFile       = "privateKeyFile"
	UbiCloudConfigRegion               = "region"
	UbiCloudConfigSecGroups            = "secGroups"
	UbiCloudConfigSSHPort              = "sshPort"
	UbiCloudConfigSSHUser              = "sshUser"
	UbiCloudConfigTenantID             = "tenantId"
	UbiCloudConfigTenantName           = "tenantName"
	UbiCloudConfigUserDataFile         = "userDataFile"
	UbiCloudConfigUsername             = "username"
)

type UbiCloudConfig struct {
	ActiveTimeout        string `json:"activeTimeout,omitempty" yaml:"activeTimeout,omitempty"`
	AuthURL              string `json:"authUrl,omitempty" yaml:"authUrl,omitempty"`
	AvailabilityZone     string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`
	CaCert               string `json:"cacert,omitempty" yaml:"cacert,omitempty"`
	ConfigDrive          bool   `json:"configDrive,omitempty" yaml:"configDrive,omitempty"`
	DomainID             string `json:"domainId,omitempty" yaml:"domainId,omitempty"`
	DomainName           string `json:"domainName,omitempty" yaml:"domainName,omitempty"`
	EndpointType         string `json:"endpointType,omitempty" yaml:"endpointType,omitempty"`
	FlavorID             string `json:"flavorId,omitempty" yaml:"flavorId,omitempty"`
	FlavorName           string `json:"flavorName,omitempty" yaml:"flavorName,omitempty"`
	FloatingIPPool       string `json:"floatingipPool,omitempty" yaml:"floatingipPool,omitempty"`
	ImageID              string `json:"imageId,omitempty" yaml:"imageId,omitempty"`
	ImageName            string `json:"imageName,omitempty" yaml:"imageName,omitempty"`
	Insecure             bool   `json:"insecure,omitempty" yaml:"insecure,omitempty"`
	IPVersion            string `json:"ipVersion,omitempty" yaml:"ipVersion,omitempty"`
	KeypairName          string `json:"keypairName,omitempty" yaml:"keypairName,omitempty"`
	NetID                string `json:"netId,omitempty" yaml:"netId,omitempty"`
	NetName              string `json:"netName,omitempty" yaml:"netName,omitempty"`
	NodePoolAntiAffinity bool   `json:"nodepoolAntiAffinity,omitempty" yaml:"nodepoolAntiAffinity,omitempty"`
	NovaNetwork          bool   `json:"novaNetwork,omitempty" yaml:"novaNetwork,omitempty"`
	PrivateKeyFile       string `json:"privateKeyFile,omitempty" yaml:"privateKeyFile,omitempty"`
	Region               string `json:"region,omitempty" yaml:"region,omitempty"`
	SecGroups            string `json:"secGroups,omitempty" yaml:"secGroups,omitempty"`
	SSHPort              string `json:"sshPort,omitempty" yaml:"sshPort,omitempty"`
	SSHUser              string `json:"sshUser,omitempty" yaml:"sshUser,omitempty"`
	TenantID             string `json:"tenantId,omitempty" yaml:"tenantId,omitempty"`
	TenantName           string `json:"tenantName,omitempty" yaml:"tenantName,omitempty"`
	UserDataFile         string `json:"userDataFile,omitempty" yaml:"userDataFile,omitempty"`
	Username             string `json:"username,omitempty" yaml:"username,omitempty"`
}

const (
	NodeTemplateType                          = "nodeTemplate"
	NodeTemplateFieldAnnotations              = "annotations"
	NodeTemplateFieldAuthCertificateAuthority = "authCertificateAuthority"
	NodeTemplateFieldAuthKey                  = "authKey"
	NodeTemplateFieldCreated                  = "created"
	NodeTemplateFieldCreatorID                = "creatorId"
	NodeTemplateFieldDescription              = "description"
	NodeTemplateFieldDockerVersion            = "dockerVersion"
	NodeTemplateFieldDriver                   = "driver"
	NodeTemplateFieldEngineEnv                = "engineEnv"
	NodeTemplateFieldEngineInsecureRegistry   = "engineInsecureRegistry"
	NodeTemplateFieldEngineInstallURL         = "engineInstallURL"
	NodeTemplateFieldEngineLabel              = "engineLabel"
	NodeTemplateFieldEngineOpt                = "engineOpt"
	NodeTemplateFieldEngineRegistryMirror     = "engineRegistryMirror"
	NodeTemplateFieldEngineStorageDriver      = "engineStorageDriver"
	NodeTemplateFieldLabels                   = "labels"
	NodeTemplateFieldName                     = "name"
	NodeTemplateFieldOwnerReferences          = "ownerReferences"
	NodeTemplateFieldRemoved                  = "removed"
	NodeTemplateFieldState                    = "state"
	NodeTemplateFieldStatus                   = "status"
	NodeTemplateFieldTransitioning            = "transitioning"
	NodeTemplateFieldTransitioningMessage     = "transitioningMessage"
	NodeTemplateFieldUbiCloudConfig           = "ubicloudConfig"
	NodeTemplateFieldUUID                     = "uuid"
	NodeTemplateFieldUseInternalIPAddress     = "useInternalIpAddress"
)

type NodeTemplate struct {
	types.Resource
	Annotations              map[string]string                    `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AuthCertificateAuthority string                               `json:"authCertificateAuthority,omitempty" yaml:"authCertificateAuthority,omitempty"`
	AuthKey                  string                               `json:"authKey,omitempty" yaml:"authKey,omitempty"`
	Created                  string                               `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                string                               `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Description              string                               `json:"description,omitempty" yaml:"description,omitempty"`
	DockerVersion            string                               `json:"dockerVersion,omitempty" yaml:"dockerVersion,omitempty"`
	Driver                   string                               `json:"driver,omitempty" yaml:"driver,omitempty"`
	EngineEnv                map[string]string                    `json:"engineEnv,omitempty" yaml:"engineEnv,omitempty"`
	EngineInsecureRegistry   []string                             `json:"engineInsecureRegistry,omitempty" yaml:"engineInsecureRegistry,omitempty"`
	EngineInstallURL         string                               `json:"engineInstallURL,omitempty" yaml:"engineInstallURL,omitempty"`
	EngineLabel              map[string]string                    `json:"engineLabel,omitempty" yaml:"engineLabel,omitempty"`
	EngineOpt                map[string]string                    `json:"engineOpt,omitempty" yaml:"engineOpt,omitempty"`
	EngineRegistryMirror     []string                             `json:"engineRegistryMirror,omitempty" yaml:"engineRegistryMirror,omitempty"`
	EngineStorageDriver      string                               `json:"engineStorageDriver,omitempty" yaml:"engineStorageDriver,omitempty"`
	Labels                   map[string]string                    `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                     string                               `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences          []managementClient.OwnerReference    `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed                  string                               `json:"removed,omitempty" yaml:"removed,omitempty"`
	State                    string                               `json:"state,omitempty" yaml:"state,omitempty"`
	Status                   *managementClient.NodeTemplateStatus `json:"status,omitempty" yaml:"status,omitempty"`
	Transitioning            string                               `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage     string                               `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                     string                               `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	UseInternalIPAddress     bool                                 `json:"useInternalIpAddress,omitempty" yaml:"useInternalIpAddress,omitempty"`
	UbiCloudConfig           *UbiCloudConfig                      `json:"ubicloudConfig" yaml:"ubicloudConfig"`
}

type NodeTemplateCollection struct {
	types.Collection
	Data   []NodeTemplate `json:"data,omitempty"`
	client *NodeTemplateClient
}

type NodeTemplateClient struct {
	apiClient *managementClient.Client
}

type NodeTemplateOperations interface {
	List(opts *types.ListOpts) (*NodeTemplateCollection, error)
	Create(opts *NodeTemplate) (*NodeTemplate, error)
	Update(existing *NodeTemplate, updates interface{}) (*NodeTemplate, error)
	Replace(existing *NodeTemplate) (*NodeTemplate, error)
	ByID(id string) (*NodeTemplate, error)
	Delete(container *NodeTemplate) error
}

func NewNodeTemplateClient(apiClient *managementClient.Client) *NodeTemplateClient {
	return &NodeTemplateClient{
		apiClient: apiClient,
	}
}

func (c *NodeTemplateClient) Create(container *NodeTemplate) (*NodeTemplate, error) {
	resp := &NodeTemplate{}
	err := c.apiClient.Ops.DoCreate(NodeTemplateType, container, resp)
	return resp, err
}

func (c *NodeTemplateClient) Update(existing *NodeTemplate, updates interface{}) (*NodeTemplate, error) {
	resp := &NodeTemplate{}
	err := c.apiClient.Ops.DoUpdate(NodeTemplateType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *NodeTemplateClient) Replace(obj *NodeTemplate) (*NodeTemplate, error) {
	resp := &NodeTemplate{}
	err := c.apiClient.Ops.DoReplace(NodeTemplateType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *NodeTemplateClient) List(opts *types.ListOpts) (*NodeTemplateCollection, error) {
	resp := &NodeTemplateCollection{}
	err := c.apiClient.Ops.DoList(NodeTemplateType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *NodeTemplateCollection) Next() (*NodeTemplateCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &NodeTemplateCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *NodeTemplateClient) ByID(id string) (*NodeTemplate, error) {
	resp := &NodeTemplate{}
	err := c.apiClient.Ops.DoByID(NodeTemplateType, id, resp)
	return resp, err
}

func (c *NodeTemplateClient) Delete(container *NodeTemplate) error {
	return c.apiClient.Ops.DoResourceDelete(NodeTemplateType, &container.Resource)
}

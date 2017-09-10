package rm

import (
	"github.com/containers/image/signature"
	"github.com/containers/image/types"
	"github.com/mangirdaz/regman/registry"
)

const (
	dockerRegistryAPI  = "https://registry-1.docker.io"
	dockerRegistryName = "docker.io"
)

// RegistryConfig struct contains registry configuration and type
type RegistryConfig struct {
	URL      string
	Username string
	Password string
	Type     string
}

// NamespaceMap - structure containing map of namespaces if image need to be retagged.
// In example images from namespace/image get get renamed to dest registry namspace-org/image by setting this struct in the config
type NamespaceMap struct {
	Source      string
	Destination string
}

// Config - main config structure for Registry Manager.
type Config struct {
	SourceRegistry              RegistryConfig
	DestinationRegistry         RegistryConfig
	InsecurePolicy              bool
	RegistriesDirPath           string
	DockerCertPath              string
	DockerInsecureSkipTLSVerify bool
	NamespaceMap                []NamespaceMap
	ImageList                   []string
}

// RegMan - Registry Manager structure with private objects
type RegMan struct {
	Config   Config
	srcCtx   *types.SystemContext
	dstCtx   *types.SystemContext
	policy   *signature.PolicyContext
	registry *registry.Registry
}

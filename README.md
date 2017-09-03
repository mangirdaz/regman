# Registry Manager - regman

Registry Manager libary is cut version of skopeo tools. Libary is far from finish, but it can be already used to do simple copy/sync images from one registry to other.  

## Basics

At the moment libary has only one interface implementation based on `github.com/containers` 

Package API:

```
	Copy(img string) error
```

## Example

Usage example:
Populate configuration object with your registry details:
```
import(
rm "github.com/mangirdaz/regman"
containers "github.com/mangirdaz/regman/containers"
)

config:= containers.Config{}

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
```

Iniciate new method:
```
rm: = containers.NewInstance(c Config)
```

Call copy method with image:
```
rm.Copy ("namespace/image:tag")
```

# Compile:

Because containers libary uses C libaries you might need to install some libs :)

```sh
Fedora$ sudo dnf install gpgme-devel libassuan-devel btrfs-progs-devel device-mapper-devel
macOS$ brew install gpgme
```

# Vendoring

For vendoring we use `vndr` package
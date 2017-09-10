package rm

import (
	"github.com/containers/image/signature"
	"github.com/containers/image/types"
	"github.com/mangirdaz/regman/registry"
	log "github.com/sirupsen/logrus"
)

func getPolicyContext(c Config) (*signature.PolicyContext, error) {
	var policy *signature.Policy // This could be cached across calls, if we had an application context.
	var err error
	if c.InsecurePolicy {
		policy = &signature.Policy{Default: []signature.PolicyRequirement{signature.NewPRInsecureAcceptAnything()}}
	} else {
		policy, err = signature.DefaultPolicy(nil)
	}
	if err != nil {
		return nil, err
	}

	return signature.NewPolicyContext(policy)
}

// NewInstance will return new instance of Registry Manager
func NewInstance(c Config) RegMan {

	srcAuth := &types.DockerAuthConfig{
		Password: c.SourceRegistry.Password,
		Username: c.SourceRegistry.Username,
	}

	dstAuth := &types.DockerAuthConfig{
		Password: c.DestinationRegistry.Password,
		Username: c.DestinationRegistry.Username,
	}

	srcCtx := &types.SystemContext{
		RegistriesDirPath:           c.RegistriesDirPath,
		DockerCertPath:              c.DockerCertPath,
		DockerInsecureSkipTLSVerify: c.DockerInsecureSkipTLSVerify,
		DockerAuthConfig:            srcAuth,
	}
	dstCtx := &types.SystemContext{
		RegistriesDirPath:           c.RegistriesDirPath,
		DockerCertPath:              c.DockerCertPath,
		DockerInsecureSkipTLSVerify: c.DockerInsecureSkipTLSVerify,
		DockerAuthConfig:            dstAuth,
	}

	policyContext, err := getPolicyContext(c)
	if err != nil {
		log.Errorf("Error loading trust policy: %v", err)
	}
	//defer policyContext.Destroy()

	registry, err := registry.New(c.SourceRegistry.URL, c.SourceRegistry.Username, c.SourceRegistry.Password)
	if err != nil {
		log.Error(err)
	}

	regman := RegMan{
		Config:   c,
		srcCtx:   srcCtx,
		dstCtx:   dstCtx,
		policy:   policyContext,
		registry: registry,
	}

	return regman
}

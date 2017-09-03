package rm

import (
	"fmt"
	"os"
	"strings"

	"github.com/containers/image/copy"
	"github.com/containers/image/transports/alltransports"
	log "github.com/sirupsen/logrus"
)

//TODO: add channel for concurant copies

// Copy function will copy image as per src and dst registry configuration
func (a RegMan) Copy(img string) error {

	sourceURL, err := getRegistryURL(a.Config.SourceRegistry.Type, a.Config.SourceRegistry.URL, img)
	if err != nil {
		log.Error(err)
	}

	destinationURL, err := getRegistryURL(a.Config.DestinationRegistry.Type, a.Config.DestinationRegistry.URL, getDestinationImage(img, a.Config.NamespaceMap))

	srcRef, err := alltransports.ParseImageName(sourceURL)
	if err != nil {
		log.Errorf("Invalid source name %s: %v", sourceURL, err)
	}
	destRef, err := alltransports.ParseImageName(destinationURL)
	if err != nil {
		log.Errorf("Invalid destination name %s: %v", destinationURL, err)
	}

	log.Debugf("Copy %s to %s", sourceURL, destinationURL)

	//signBy := context.String("sign-by")
	removeSignatures := true

	log.Debug("Start copy")
	err = copy.Image(a.policy, destRef, srcRef, &copy.Options{
		RemoveSignatures: removeSignatures,
		ReportWriter:     os.Stdout,
		SourceCtx:        a.srcCtx,
		DestinationCtx:   a.dstCtx,
	})
	if err != nil {
		log.Error(err)
	}
	return err
}

func getDestinationImage(image string, namespaceMap []NamespaceMap) string {
	//TODO: add logic if image does not have namespace (library tipe)
	imageRaw := strings.Split(image, "/")
	imageNew := fmt.Sprintf("%s/%s", getMapNamespace(imageRaw[0], namespaceMap), imageRaw[1])
	return imageNew
}

func getMapNamespace(source string, namespaceMap []NamespaceMap) string {
	//TODO add check if namespaces exist
	for _, element := range namespaceMap {
		if element.Source == source {
			return element.Destination
		}
	}
	return source
}

func getRegistryURL(regtype, url, image string) (string, error) {
	//TODO: add type, url, image validation
	return fmt.Sprintf("%s://%s/%s", regtype, url, image), nil
}

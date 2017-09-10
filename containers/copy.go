package rm

import (
	"fmt"
	"os"

	"github.com/containers/image/copy"
	"github.com/containers/image/transports/alltransports"
	log "github.com/sirupsen/logrus"
)

//TODO: add channel for concurant copies

// Copy function will copy image as per src and dst registry configuration
func (a RegMan) Copy(imgSrc, imgDst string) error {

	sourceURL, err := getRegistryURL(a.Config.SourceRegistry.Type, a.Config.SourceRegistry.URL, imgSrc)
	if err != nil {
		log.Error(err)
	}

	destinationURL, err := getRegistryURL(a.Config.DestinationRegistry.Type, a.Config.DestinationRegistry.URL, imgDst)

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

// GetImages - returns images in src registry namespace
func GetImages(namespace string) (string, error) {
	return "", nil
}

// GetTags - return tags of particular image
func GetTags(namespace, image string) (string, error) {
	return "", nil
}

// GetDigest - return unique ID/DIGEST of the image
func GetDigest(namespace string) (string, error) {
	return "", nil
}

func getRegistryURL(regtype, url, image string) (string, error) {
	//TODO: add type, url, image validation
	return fmt.Sprintf("%s://%s/%s", regtype, url, image), nil
}

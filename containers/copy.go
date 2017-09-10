package rm

import (
	"fmt"
	"os"
	"regexp"

	"github.com/containers/image/copy"
	"github.com/containers/image/transports/alltransports"
	log "github.com/sirupsen/logrus"
)

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

func getRegistryURL(regtype, url, image string) (string, error) {
	//TODO: add type, url, image validation
	var re = regexp.MustCompile("https|http")
	s := re.ReplaceAllString(url, "")
	log.Debug(s)
	return fmt.Sprintf("%s%s/%s", regtype, s, image), nil
}

package rm

import (
	log "github.com/sirupsen/logrus"
)

// GetImages - returns images in src registry namespace
func (a RegMan) GetImages() ([]string, error) {
	log.Debug("Get Images")
	list, err := a.registry.Repositories()
	return list, err
}

// GetTags - return tags of particular image
func (a RegMan) GetTags(image string) ([]string, error) {
	log.Debug("Get Tags")
	list, err := a.registry.Tags(image)
	return list, err
}

// GetDigest - return unique ID/DIGEST of the image
//TODO
func (a RegMan) GetDigest(namespacem, image, tag string) (string, error) {
	return "", nil
}

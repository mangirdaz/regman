package rm

// Registry - generic registry communication interface
type Registry interface {
	//requires both src and dst registries to be set
	Copy(srcImg, dstImg string) error
	//will query only src registry
	GetImages(namespace string) ([]string, error)
	GetTags(namespace, image string) ([]string, error)
	GetDigest(namespace, image, tag string) (string, error)
}

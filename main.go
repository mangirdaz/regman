package rm

// Registry - generic registry communication interface
type Registry interface {
	//requires both src and dst registries to be set
	Copy(srcImg, dstImg string) error
	//will query only src registry
	GetImages() ([]string, error)
	GetTags(image string) ([]string, error)
	//TODO
	GetDigest(namespace, image, tag string) (string, error)
}

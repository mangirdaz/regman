package rm

// Registry - generic registry communication interface
type Registry interface {
	Copy(img string) error
}

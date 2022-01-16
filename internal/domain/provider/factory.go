package provider

type Factory interface {
	Get(t Type) (Provider, error)
}

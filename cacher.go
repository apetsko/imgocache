package imgocache

// Cacher interface realize GET SET and DELETE methods
type Cacher interface {
	Get(string) (interface{}, bool)
	Set(string, interface{})
	Delete(string)
}

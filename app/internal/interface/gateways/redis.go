package gateways

type Redis interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

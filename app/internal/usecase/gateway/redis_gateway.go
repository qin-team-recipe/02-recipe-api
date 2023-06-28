package gateway

type RedisGateway interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

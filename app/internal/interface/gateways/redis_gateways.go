package gateways

type RedisGateways struct {
	Redis Redis
}

func (r *RedisGateways) Set(key string, value interface{}) error {
	return r.Redis.Set(key, value)
}

func (r *RedisGateways) Get(key string) (interface{}, error) {
	return r.Redis.Get(key)
}

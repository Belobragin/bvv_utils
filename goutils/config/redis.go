package config

type RedisI interface {
	GetRedisAddr() string
	GetRedisPassw() string
	GetRedisDb() int
}

type RedisConfig struct {
	Address  string `conf:"env:REDIS_ADDRESS"`
	RedisPsw string `default:, conf:"env:REDIS_PSW"`
	RedisDb  int    `default:0, conf:"env:REDIS_DB"`
}

func (r *RedisConfig) GetRedisAddr() string {
	return r.Address
}
func (r *RedisConfig) GetRedisPassw() string {
	return r.RedisPsw
}
func (r *RedisConfig) GetRedisDb() int {
	return r.RedisDb
}

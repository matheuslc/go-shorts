package config

type RedisOptions struct {
	RedisAddress  string `env:"REDIS_ADDRESS"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	Database      int    `env:"REDIS_DATABASE"`
}

type Environment struct {
	Redis RedisOptions
}

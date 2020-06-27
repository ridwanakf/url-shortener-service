package config

type Redis struct {
	Address   string `yaml:"address" env:"REDIS_URL"`
	Timeout   int    `yaml:"timeout" env:"REDIS_TIMEOUT"`
	MaxIdle   int    `yaml:"max_idle" env:"REDIS_MAX_IDLE"`
	MaxActive int    `yaml:"max_active" env:"REDIS_MAX_ACTIVE"`
}

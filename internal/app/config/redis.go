package config

type Redis struct {
	Address string `yaml:"address"`
	Timeout int    `yaml:"timeout"`
	MaxIdle int    `yaml:"max_idle"`
}

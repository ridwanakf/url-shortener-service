package config

type Config struct {
	DB     Database `yaml:"database"`
	Params Params   `yaml:"params"`
}

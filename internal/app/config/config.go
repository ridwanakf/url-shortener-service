package config

type Config struct {
	Server Server   `yaml:"server"`
	Params Params   `yaml:"params"`
	DB     Database `yaml:"database"`
	Redis  Redis    `yaml:"redis"`
}

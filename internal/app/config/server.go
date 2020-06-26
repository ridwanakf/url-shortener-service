package config

type Server struct {
	Port                string `yaml:"port"`
	Debug               bool   `yaml:"debug"`
	ReadTimeoutSeconds  int    `yaml:"read_timeout_seconds"`
	WriteTimeoutSeconds int    `yaml:"write_timeout_seconds"`
}

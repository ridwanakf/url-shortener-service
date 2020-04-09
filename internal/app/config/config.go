package config

type Config struct {
	DB   Database `yaml:"database"`
	Flag Flag     `yaml:"flag"`
}

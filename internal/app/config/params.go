package config

type Params struct {
	Port           string `yaml:"port"`
	ShortUrlLength int    `yaml:"short_url_length"`
	ExpireDuration int    `yaml:"expire_duration"` //days
}

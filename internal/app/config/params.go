package config

type Params struct {
	ShortUrlLength int    `yaml:"short_url_length"`
	ExpireDuration int    `yaml:"expire_duration"` //days
}

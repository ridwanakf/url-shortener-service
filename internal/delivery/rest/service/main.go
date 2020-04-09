package service

import "github.com/ridwanakf/url-shortener-service/internal/app"

type Services struct {
	*ShortenerService
	*DefaultService
}

func GetServices(app *app.UrlShortenerApp) *Services {
	return &Services{
		ShortenerService: NewShortenerService(app),
		DefaultService:   NewDefaultService(),
	}
}

package service

import "github.com/ridwanakf/url-shortener-service/internal/app"

var NotFound = struct {
	Message string `json:"message"`
}{
	Message: "404 Not Found!",
}

type Services struct {
	*ShortenerService
}

func GetServices(app *app.UrlShortenerApp) *Services {
	return &Services{
		ShortenerService: NewShortenerService(app),
	}
}

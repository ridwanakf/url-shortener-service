package main

import (
	"log"

	"github.com/ridwanakf/url-shortener-service/internal/app"
	"github.com/ridwanakf/url-shortener-service/internal/delivery/rest"
)

func main() {
	// init app
	UrlShortenerApp, err := app.NewUrlShortenerApp()
	if err != nil {
		log.Fatalf("marshal error %+v", err)
	}
	defer func() {
		errs := UrlShortenerApp.Close()
		for e := range errs {
			log.Println(e)
		}
	}()

	rest.Start(UrlShortenerApp)
}

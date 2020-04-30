package utils

import (
	"net/http"

	"github.com/ridwanakf/url-shortener-service/internal/entity"
)

func ConvertShortURL(r *http.Request, shortURL string) string {
	return r.Host + "/" + shortURL
}

func ConvertBatchShortURL(r *http.Request, urls []entity.URL) []entity.URL {
	for idx := range urls {
		urls[idx].ShortURL = ConvertShortURL(r, urls[idx].ShortURL)
	}

	return urls
}

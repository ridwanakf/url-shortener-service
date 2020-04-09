package internal

import "github.com/ridwanakf/url-shortener-service/internal/entity"

// ShortenerUC contains logic for all URL Shortener Usecase
//go:generate mockgen -destination=usecase/url_shortener_mock.go -package=usecase github.com/ridwanakf/url-shortener-service/internal ShortenerUC
type ShortenerUC interface {
	GetAllURL() ([]entity.URL, error)
	CreateNewShortURL(longURL string) (string, error)                        //return shortURL
	CreateNewCustomShortURL(shortURL string, longURL string) (string, error) //return custom shortURL if success
	UpdateShortURL(shortURL string, longURL string) (bool, error)            //return true if success
	GetLongURL(shortURL string) (string, error)
	DeleteURL(shortURL string) (bool, error)
}

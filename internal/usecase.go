package internal

import "github.com/ridwanakf/url-shortener-service/internal/entity"

// ShortenerUC contains logic for all URL Shortener Usecase
//go:generate mockgen -destination=usecase/url_shortener_usecase_mock.go -package=usecase github.com/ridwanakf/url-shortener-service/internal ShortenerUC
type ShortenerUC interface {
	GetAllURL() ([]entity.URL, error)
	CreateNewShortURL(longURL string) (entity.URL, error)
	CreateNewCustomShortURL(shortURL string, longURL string) (entity.URL, error)
	UpdateShortURL(shortURL string, longURL string) error
	GetLongURL(shortURL string) (string, error)
	DeleteURL(shortURL string) error

	GenerateShortURL(length int) string
	IsValidURL(url string) (string, error)
}

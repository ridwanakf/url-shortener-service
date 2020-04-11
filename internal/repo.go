package internal

import "github.com/ridwanakf/url-shortener-service/internal/entity"

// ShortenerDB contains repo for all URL Shortener DB
//go:generate mockgen -destination=repo/db/url_shortener_db_mock.go -package=db github.com/ridwanakf/url-shortener-service/internal ShortenerDB
type ShortenerDB interface {
	GetAllURL() ([]entity.URL, error)
	GetLongURL(shortURL string) (string, error)
	CreateNewShortURL(url entity.URL) error
	UpdateShortURL(shortURL string, longURL string) error
	DeleteURL(shortURL string) error

	IsShortURLExist(shortURL string) bool
	HasShortURLExpired(shortURL string) bool
}

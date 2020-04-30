package internal

import (
	"github.com/ridwanakf/url-shortener-service/internal/entity"
)

// ShortenerDBRepo contains repo for all URL Shortener DB
//go:generate mockgen -destination=repo/db/url_shortener_db_mock.go -package=db github.com/ridwanakf/url-shortener-service/internal ShortenerDBRepo
type ShortenerDBRepo interface {
	GetAllURL() ([]entity.URL, error)
	GetLongURL(shortURL string) (string, error)
	CreateNewShortURL(url entity.URL) error
	UpdateShortURL(shortURL string, longURL string) error
	DeleteURL(shortURL string) error

	IsShortURLExist(shortURL string) bool
	HasShortURLExpired(shortURL string) bool
}

// ShortenerCacheRepo contains repo for all URL Shortener Cache
//go:generate mockgen -destination=repo/redis_cache/url_shortener_redis_mock.go -package=redis_cache github.com/ridwanakf/url-shortener-service/internal ShortenerCacheRepo
type ShortenerCacheRepo interface {
	GetLongURL(shortURL string) (string, error)
	IsShortURLExist(shortURL string) (bool, error)
	DeleteURL(shortURL string) (bool, error)
}

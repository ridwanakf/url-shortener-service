package internal

import (
	"github.com/ridwanakf/url-shortener-service/internal/entity"
)

// ShortenerDBRepo contains repo for all URL Shortener DB
//go:generate mockgen -destination=repo/db/url_shortener_db_mock.go -package=db github.com/ridwanakf/url-shortener-service/internal ShortenerDBRepo
type ShortenerDBRepo interface {
	GetAllURL(userID string) ([]entity.URL, error)
	GetURL(shortURL string) (entity.URL, error)
	CreateNewShortURL(url entity.URL) error
	UpdateShortURL(shortURL string, longURL string) error
	DeleteURL(shortURL string) error

	IsShortURLExist(shortURL string) bool
	HasShortURLExpired(shortURL string) bool
}

// ShortenerCacheRepo contains repo for all URL Shortener Cache
//go:generate mockgen -destination=repo/redis_cache/url_shortener_redis_mock.go -package=redis_cache github.com/ridwanakf/url-shortener-service/internal ShortenerCacheRepo
type ShortenerCacheRepo interface {
	GetAllURL(userID string) ([]entity.URL, error)
	GetURL(shortURL string) (string, error)
	SetURL(url entity.URL) error
	DeleteURL(shortURL string) (int64, error)
	IsSingleURLExist(shortURL string) (bool, error)
	HasShortURLExpired(shortURL string) (bool, error)
	IsCollectionURLExist(userID string) (bool, error)
}

package redis_cache

import bridge "github.com/ridwanakf/url-shortener-service/internal"

type ShortenerCacheRepo struct {
	rd  bridge.Redis
	jsn bridge.Json
}

func NewShortenerCacheRepo(rd bridge.Redis, jsn bridge.Json) *ShortenerCacheRepo {
	return &ShortenerCacheRepo{
		rd:  rd,
		jsn: jsn,
	}
}

func (c *ShortenerCacheRepo) GetLongURL(shortURL string) (string, error) {
	panic("implement me!")
}

func (c *ShortenerCacheRepo) IsShortURLExist(shortURL string) (bool, error) {
	panic("implement me!")
}

func (c *ShortenerCacheRepo) DeleteURL(shortURL string) (bool, error) {
	panic("implement me!")
}

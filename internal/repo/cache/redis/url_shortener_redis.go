package redis

import (
	"encoding/json"
	"fmt"
	"time"

	bridge "github.com/ridwanakf/bridges"
	"github.com/ridwanakf/url-shortener-service/internal/entity"
)

const (
	SingleURLKey     = "url:single"
	CollectionURLKey = "url:all"
)

type ShortenerCacheRepo struct {
	rd bridge.Redis
}

func NewShortenerCacheRepo(rd bridge.Redis) *ShortenerCacheRepo {
	return &ShortenerCacheRepo{
		rd: rd,
	}
}

func (c *ShortenerCacheRepo) GetAllURL(userID string) ([]entity.URL, error) {
	var res []entity.URL

	urlsByte, err := c.rd.Get(c.getCollectionURLKey(userID))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(urlsByte, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c *ShortenerCacheRepo) GetURL(shortURL string) (string, error) {
	var (
		url entity.URL
		res string
	)

	urlByte, err := c.rd.Get(c.getSingleURLKey(shortURL))
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(urlByte, &url)
	if err != nil {
		return res, err
	}

	res = url.LongURL

	return res, nil
}

func (c *ShortenerCacheRepo) SetURL(url entity.URL) error {
	key := c.getSingleURLKey(url.ShortURL)

	urlByte, err := json.Marshal(url)
	if err != nil {
		return err
	}

	err = c.rd.Set(key, string(urlByte))
	if err != nil {
		return err
	}

	return nil
}

func (c *ShortenerCacheRepo) DeleteURL(shortURL string) (int64, error) {
	return c.rd.Del(c.getSingleURLKey(shortURL))
}

func (c *ShortenerCacheRepo) IsSingleURLExist(shortURL string) (bool, error) {
	return c.rd.Exists(c.getSingleURLKey(shortURL))
}

func (c *ShortenerCacheRepo) HasShortURLExpired(shortURL string) (bool, error) {
	var url entity.URL

	res := true

	urlByte, err := c.rd.Get(c.getSingleURLKey(shortURL))
	if err != nil {
		return true, err
	}

	err = json.Unmarshal(urlByte, &url)
	if err != nil {
		return true, err
	}

	expire := url.ExpireAt

	if expire.Unix() > time.Now().UTC().Unix() {
		res = false
	}

	return res, nil
}

func (c *ShortenerCacheRepo) IsCollectionURLExist(userID string) (bool, error) {
	return c.rd.Exists(c.getCollectionURLKey(userID))
}

func (c *ShortenerCacheRepo) getSingleURLKey(shortURL string) string {
	return fmt.Sprintf("%s:%s", SingleURLKey, shortURL)
}

func (c *ShortenerCacheRepo) getCollectionURLKey(userID string) string {
	return fmt.Sprintf("%s:%s", CollectionURLKey, userID)
}

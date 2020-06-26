package usecase

import (
	"log"
	"math/rand"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/ridwanakf/url-shortener-service/internal"
	"github.com/ridwanakf/url-shortener-service/internal/entity"
)

type ShortenerUsecase struct {
	db             internal.ShortenerDBRepo
	cache          internal.ShortenerCacheRepo
	shortUrlLength int
	expireDuration int
}

func NewShortenerUsecase(db internal.ShortenerDBRepo, cache internal.ShortenerCacheRepo, shortUrlLength int, expireDuration int) *ShortenerUsecase {
	return &ShortenerUsecase{
		db:             db,
		cache:          cache,
		shortUrlLength: shortUrlLength,
		expireDuration: expireDuration,
	}
}

func (u *ShortenerUsecase) GetAllURL(userID string) ([]entity.URL, error) {
	var urls []entity.URL

	urls, err := u.db.GetAllURL(userID) //TODO: implement user id when OAUTH has been implemented
	if err != nil {
		return []entity.URL{}, err
	}

	return urls, nil
}

func (u *ShortenerUsecase) CreateNewShortURL(longURL string) (entity.URL, error) {
	longURL, err := u.IsValidURL(longURL)
	if err != nil {
		return entity.URL{}, err
	}

	shortURL := u.GenerateShortURL(u.shortUrlLength)
	for {
		if !u.db.IsShortURLExist(shortURL) {
			break
		} else {
			shortURL = u.GenerateShortURL(u.shortUrlLength)
		}
	}

	URL := entity.URL{
		ShortURL:  shortURL,
		LongURL:   longURL,
		CreatedAt: time.Now().UTC(),
		ExpireAt:  time.Now().UTC().Add(time.Hour * 24 * time.Duration(u.expireDuration)),
		CreatedBy: "", //TODO: using ID if auth is implemented
	}

	if err := u.db.CreateNewShortURL(URL); err != nil {
		return entity.URL{}, err
	}

	return URL, nil
}

func (u *ShortenerUsecase) CreateNewCustomShortURL(shortURL string, longURL string) (entity.URL, error) {
	longURL, err := u.IsValidURL(longURL)
	if err != nil {
		return entity.URL{}, err
	}

	if u.db.IsShortURLExist(shortURL) {
		return entity.URL{}, errors.New("URL has already existed")
	}

	URL := entity.URL{
		ShortURL:  shortURL,
		LongURL:   longURL,
		CreatedAt: time.Now().UTC(),
		ExpireAt:  time.Now().UTC().Add(time.Hour * 24 * time.Duration(u.expireDuration)),
		CreatedBy: "", //TODO: using ID if auth is implemented
	}

	if err := u.db.CreateNewShortURL(URL); err != nil {
		return entity.URL{}, err
	}

	return URL, nil
}

func (u *ShortenerUsecase) UpdateShortURL(shortURL string, longURL string) error {
	if !u.db.IsShortURLExist(shortURL) {
		return errors.New("URL does not exist")
	}

	if err := u.db.UpdateShortURL(shortURL, longURL); err != nil {
		return err
	}

	// Delete previous record from cache if exist
	isExistInCache, err := u.cache.IsSingleURLExist(shortURL)
	if err != nil {
		log.Printf("[ShortenerUsecase][UpdateShortURL] : %+v\n", err)
	}

	if isExistInCache {
		_, err := u.cache.DeleteURL(shortURL)
		if err != nil {
			log.Printf("[ShortenerUsecase][UpdateShortURL] : %+v\n", err)
		}
	}

	return nil
}

func (u *ShortenerUsecase) GetLongURL(shortURL string) (string, error) {
	// Check from cache
	isExistInCache, err := u.cache.IsSingleURLExist(shortURL)
	if err != nil {
		log.Printf("[ShortenerUsecase][GetLongURL] : %+v\n", err)
	}

	if isExistInCache {
		hasExpired, err := u.cache.HasShortURLExpired(shortURL)
		if err != nil {
			log.Printf("[ShortenerUsecase][GetLongURL] : %+v\n", err)
		}

		if hasExpired {
			_, err := u.cache.DeleteURL(shortURL)
			if err != nil {
				log.Printf("[ShortenerUsecase][GetLongURL] : %+v\n", err)
			}
		} else {
			longURL, err := u.cache.GetURL(shortURL)
			if err != nil {
				log.Printf("[ShortenerUsecase][GetLongURL] : %+v\n", err)
			}
			return longURL, err
		}
	}

	// Check from db
	if !u.db.IsShortURLExist(shortURL) {
		return "", errors.New("URL does not exist")
	}

	if u.db.HasShortURLExpired(shortURL) {
		if err := u.db.DeleteURL(shortURL); err != nil {
			return "", errors.New("URL has expired! but failed to delete it")
		}
		return "", errors.New("URL has expired!")
	}

	urlEntity, err := u.db.GetURL(shortURL)
	if err != nil {
		return "", err
	}

	// Set cache
	err = u.cache.SetURL(urlEntity)
	if err != nil {
		log.Printf("[ShortenerUsecase][GetLongURL] : %+v\n", err)
	}

	return urlEntity.LongURL, nil
}

func (u *ShortenerUsecase) DeleteURL(shortURL string) error {
	if !u.db.IsShortURLExist(shortURL) {
		return errors.New("URL does not exist")
	}

	if err := u.db.DeleteURL(shortURL); err != nil {
		return err
	}

	// Delete from cache if exist
	isExistInCache, err := u.cache.IsSingleURLExist(shortURL)
	if err != nil {
		log.Printf("[ShortenerUsecase][DeleteURL] : %+v\n", err)
	}

	if isExistInCache {
		_, err := u.cache.DeleteURL(shortURL)
		if err != nil {
			log.Printf("[ShortenerUsecase][DeleteURL] : %+v\n", err)
		}
	}

	return nil
}

func (u *ShortenerUsecase) GenerateShortURL(length int) string {
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func (u *ShortenerUsecase) IsValidURL(input string) (string, error) {
	uri, err := url.Parse(input)
	if err != nil {
		return "", err
	}

	switch uri.Scheme {
	case "http":
	case "https":
	case "":
		uri.Scheme = "http"
	default:
		return "", errors.New("Invalid scheme")
	}

	return uri.String(), nil
}

package usecase

import (
	"github.com/pkg/errors"
	"github.com/ridwanakf/url-shortener-service/internal"
	"github.com/ridwanakf/url-shortener-service/internal/entity"
	"math/rand"
	"time"
)

type ShortenerUsecase struct {
	db             internal.ShortenerDB
	shortUrlLength int
	expireDuration int
}

func NewShortenerUsecase(db internal.ShortenerDB, shortUrlLength int, expireDuration int) *ShortenerUsecase {
	return &ShortenerUsecase{
		db:             db,
		shortUrlLength: shortUrlLength,
		expireDuration: expireDuration,
	}
}

func (u *ShortenerUsecase) GetAllURL() ([]entity.URL, error) {
	var urls []entity.URL

	urls, err := u.db.GetAllURL()
	if err != nil {
		return []entity.URL{}, err
	}

	return urls, nil
}

func (u *ShortenerUsecase) CreateNewShortURL(longURL string) (entity.URL, error) {
	shortURL := u.GenerateShortURL(u.shortUrlLength) //TODO: move shorturl length to config
	for {
		if !u.db.IsShortURLExist(shortURL) {
			break
		} else {
			shortURL = u.GenerateShortURL(u.shortUrlLength)
		}
	}

	url := entity.URL{
		ShortURL:  shortURL,
		LongURL:   longURL,
		CreatedAt: time.Now(),
		ExpireAt:  time.Now().Add(time.Hour * 12 * time.Duration(u.expireDuration)),
		CreatedBy: "", //TODO: using ID if auth is implemented
	}

	if err := u.db.CreateNewShortURL(url); err != nil {
		return entity.URL{}, err
	}

	return url, nil
}

func (u *ShortenerUsecase) CreateNewCustomShortURL(shortURL string, longURL string) (entity.URL, error) {
	if u.db.IsShortURLExist(shortURL) {
		return entity.URL{}, errors.New("URL has already existed")
	}

	url := entity.URL{
		ShortURL:  shortURL,
		LongURL:   longURL,
		CreatedAt: time.Now(),
		ExpireAt:  time.Now().Add(time.Hour * 12 * time.Duration(u.expireDuration)),
		CreatedBy: "", //TODO: using ID if auth is implemented
	}

	if err := u.db.CreateNewShortURL(url); err != nil {
		return entity.URL{}, err
	}

	return url, nil
}

func (u *ShortenerUsecase) UpdateShortURL(shortURL string, longURL string) error {
	if !u.db.IsShortURLExist(shortURL) {
		return errors.New("URL does not exist")
	}

	if err := u.db.UpdateShortURL(shortURL, longURL); err != nil {
		return err
	}

	return nil
}

func (u *ShortenerUsecase) GetLongURL(shortURL string) (string, error) {
	if !u.db.IsShortURLExist(shortURL) {
		return "", errors.New("URL does not exist")
	}

	if u.db.HasShortURLExpired(shortURL) {
		if err := u.DeleteURL(shortURL); err != nil {
			return "", errors.New("URL has expired! but failed to delete it")
		}
		return "", errors.New("URL has expired!")
	}

	longURL, err := u.db.GetLongURL(shortURL)
	if err != nil {
		return "", err
	}

	return longURL, nil
}

func (u *ShortenerUsecase) DeleteURL(shortURL string) error {
	if !u.db.IsShortURLExist(shortURL) {
		return errors.New("URL does not exist")
	}

	if err := u.db.DeleteURL(shortURL); err != nil {
		return err
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

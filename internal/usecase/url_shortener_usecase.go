package usecase

import (
	"github.com/ridwanakf/url-shortener-service/internal"
	"github.com/ridwanakf/url-shortener-service/internal/entity"
)

type ShortenerUsecase struct {
	db internal.ShortenerDB
}

func NewShortenerUsecase(db internal.ShortenerDB) *ShortenerUsecase {
	return &ShortenerUsecase{
		db: db,
	}
}

func (u *ShortenerUsecase) GetAllURL() ([]entity.URL, error) {
	panic("implement me!")
}

func (u *ShortenerUsecase) CreateNewShortURL(longURL string) (entity.URL, error) {
	panic("implement me!")
}

func (u *ShortenerUsecase) CreateNewCustomShortURL(shortURL string, longURL string) (entity.URL, error) {
	panic("implement me!")
}

func (u *ShortenerUsecase) UpdateShortURL(shortURL string, longURL string) error {
	panic("implement me!")
}

func (u *ShortenerUsecase) GetLongURL(shortURL string) (string, error) {
	panic("implement me!")
}

func (u *ShortenerUsecase) DeleteURL(shortURL string) error {
	panic("implement me!")
}

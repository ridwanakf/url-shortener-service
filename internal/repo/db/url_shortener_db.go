package db

import (
	"database/sql"
	"github.com/ridwanakf/url-shortener-service/internal/entity"
)

type ShortenerDBRepo struct {
	db *sql.DB
}

func NewShortenerDBRepo(db *sql.DB) *ShortenerDBRepo {
	return &ShortenerDBRepo{
		db: db,
	}
}

func (d *ShortenerDBRepo) GetAllURL() ([]entity.URL, error) {
	panic("implement me!")
}

func (d *ShortenerDBRepo) CreateNewShortURL(url entity.URL) error {
	panic("implement me!")
}

func (d *ShortenerDBRepo) UpdateShortURL(shortURL string, longURL string) error {
	panic("implement me!")
}

func (d *ShortenerDBRepo) GetLongURL(shortURL string) (string, error) {
	panic("implement me!")
}

func (d *ShortenerDBRepo) DeleteURL(shortURL string) error {
	panic("implement me!")
}

func (d *ShortenerDBRepo) IsShortURLExist(shortURL string) bool {
	panic("implement me!")
}
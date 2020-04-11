package db

import (
	"database/sql"
	"github.com/ridwanakf/url-shortener-service/internal/entity"
)

type ShortenerDB struct {
	db *sql.DB
}

func NewShortenerDB(db *sql.DB) *ShortenerDB {
	return &ShortenerDB{
		db: db,
	}
}

func (d *ShortenerDB) GetAllURL() ([]entity.URL, error) {
	panic("implement me!")
}

func (d *ShortenerDB) CreateNewShortURL(url entity.URL) error {
	panic("implement me!")
}

func (d *ShortenerDB) UpdateShortURL(shortURL string, longURL string) error {
	panic("implement me!")
}

func (d *ShortenerDB) GetLongURL(shortURL string) (string, error) {
	panic("implement me!")
}

func (d *ShortenerDB) DeleteURL(shortURL string) error {
	panic("implement me!")
}

func (d *ShortenerDB) IsShortURLExist(shortURL string) bool {
	panic("implement me!")
}
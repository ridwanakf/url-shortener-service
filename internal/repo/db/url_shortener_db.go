package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/ridwanakf/url-shortener-service/internal/entity"
	"time"
)

type ShortenerDBRepo struct {
	db *sqlx.DB
}

func NewShortenerDBRepo(db *sqlx.DB) *ShortenerDBRepo {
	return &ShortenerDBRepo{
		db: db,
	}
}

func (r *ShortenerDBRepo) GetAllURL() ([]entity.URL, error) {
	var res []entity.URL

	err := r.db.Select(&res, SQLGetAllURLQuery)
	if err != nil {
		return []entity.URL{}, err
	}

	return res, nil
}

func (r *ShortenerDBRepo) GetLongURL(shortURL string) (string, error) {
	var res string

	err := r.db.Get(&res, SQLGetLongURLQuery, shortURL)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (r *ShortenerDBRepo) CreateNewShortURL(url entity.URL) error {
	_, err := r.db.Exec(SQLCreateNewEntryQuery, url.ShortURL, url.LongURL, url.CreatedAt, url.ExpireAt, url.CreatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *ShortenerDBRepo) UpdateShortURL(shortURL string, longURL string) error {
	_, err := r.db.Exec(SQLUpdateShortURLQuery, longURL, shortURL)
	if err != nil {
		return err
	}

	return nil
}

func (r *ShortenerDBRepo) DeleteURL(shortURL string) error {
	_, err := r.db.Exec(SQLDeleteEntryQuery, shortURL)
	if err != nil {
		return err
	}

	return nil
}

func (r *ShortenerDBRepo) IsShortURLExist(shortURL string) bool {
	var isExist bool

	err := r.db.Get(&isExist, SQLIsEntryExistQuery, shortURL)
	if err != nil {
		panic(err)
	}

	return isExist
}

func (r *ShortenerDBRepo) HasShortURLExpired(shortURL string) bool {
	var hasExpired bool

	err := r.db.Get(&hasExpired, SQLHasEntryExpiredQuery, time.Now(), shortURL)
	if err != nil {
		panic(err)
	}

	return hasExpired
}

package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/ridwanakf/url-shortener-service/internal"
	db2 "github.com/ridwanakf/url-shortener-service/internal/repo/db"
)

type Repos struct {
	ShortenerDB internal.ShortenerDBRepo
}

func newRepos(db *sqlx.DB) (*Repos, error) {
	r := &Repos{
		ShortenerDB: db2.NewShortenerDBRepo(db),
	}

	return r, nil
}

func (r *Repos) Close() []error {
	var errs []error
	return errs
}

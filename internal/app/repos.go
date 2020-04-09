package app

import (
	"database/sql"
	"github.com/ridwanakf/url-shortener-service/internal"
	db2 "github.com/ridwanakf/url-shortener-service/internal/repo/db"
)

type Repos struct {
	ShortenerDB internal.ShortenerDB
}

func newRepos(db *sql.DB) (*Repos, error) {
	r := &Repos{
		ShortenerDB: db2.NewShortenerDBRepo(db),
	}

	return r, nil
}

func (r *Repos) Close() []error {
	var errs []error
	return errs
}

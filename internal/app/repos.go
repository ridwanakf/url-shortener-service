package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/ridwanakf/url-shortener-service/internal"
	"github.com/ridwanakf/url-shortener-service/internal/repo/cache/redis"
	"github.com/ridwanakf/url-shortener-service/internal/repo/db/postgres"
)

type Repos struct {
	ShortenerDB    internal.ShortenerDBRepo
	ShortenerCache internal.ShortenerCacheRepo
}

func newRepos(bridges *Bridges, db *sqlx.DB) (*Repos, error) {
	r := &Repos{
		ShortenerDB:    postgres.NewShortenerDBRepo(db),
		ShortenerCache: redis.NewShortenerCacheRepo(bridges.Redis),
	}

	return r, nil
}

func (r *Repos) Close() []error {
	var errs []error
	return errs
}

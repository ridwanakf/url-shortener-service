package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/ridwanakf/url-shortener-service/internal"
	db2 "github.com/ridwanakf/url-shortener-service/internal/repo/db"
	"github.com/ridwanakf/url-shortener-service/internal/repo/redis_cache"
)

type Repos struct {
	ShortenerDB    internal.ShortenerDBRepo
	ShortenerCache internal.ShortenerCacheRepo
}

func newRepos(bridges *Bridges, db *sqlx.DB) (*Repos, error) {
	r := &Repos{
		ShortenerDB:    db2.NewShortenerDBRepo(db),
		ShortenerCache: redis_cache.NewShortenerCacheRepo(bridges.Redis, bridges.Json),
	}

	return r, nil
}

func (r *Repos) Close() []error {
	var errs []error
	return errs
}

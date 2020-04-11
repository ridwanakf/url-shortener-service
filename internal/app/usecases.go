package app

import (
	"github.com/ridwanakf/url-shortener-service/internal"
	"github.com/ridwanakf/url-shortener-service/internal/app/config"
	"github.com/ridwanakf/url-shortener-service/internal/usecase"
)

type Usecases struct {
	ShortenerUC internal.ShortenerUC
}

func newUsecases(repos *Repos, cfg *config.Config) *Usecases {
	return &Usecases{
		ShortenerUC: usecase.NewShortenerUsecase(
			repos.ShortenerDB,
			cfg.Params.ShortUrlLength,
			cfg.Params.ExpireDuration),
	}
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}

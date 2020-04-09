package app

import (
	"github.com/ridwanakf/url-shortener-service/internal"
	"github.com/ridwanakf/url-shortener-service/internal/usecase"
)

type Usecases struct {
	ShortenerUC internal.ShortenerUC
}

func newUsecases(repos *Repos) *Usecases {
	return &Usecases{
		ShortenerUC: usecase.NewShortenerUsecase(repos.ShortenerDB),
	}
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}

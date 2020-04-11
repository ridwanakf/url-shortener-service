package app

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/ridwanakf/url-shortener-service/constant"
	"github.com/ridwanakf/url-shortener-service/internal/app/config"
	"gopkg.in/yaml.v2"
)

type UrlShortenerApp struct {
	Repos    *Repos
	UseCases *Usecases
	Cfg      config.Config
}

func NewUrlShortenerApp() (*UrlShortenerApp, error) {
	cfg, err := readConfig(constant.ConfigProjectFilepath)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting config")
	}

	db, err := initDB(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "error connect db")
	}

	app := new(UrlShortenerApp)

	app.Cfg = cfg

	app.Repos, err = newRepos(db)
	if err != nil {
		return nil, errors.Wrap(err, "errors invoking newRepos")
	}

	app.UseCases = newUsecases(app.Repos)

	return app, nil
}

func (a *UrlShortenerApp) Close() []error {
	var errs []error

	errs = append(errs, a.Repos.Close()...)
	errs = append(errs, a.UseCases.Close()...)

	return errs
}

func readConfig(cfgPath string) (config.Config, error) {
	f, err := os.Open(cfgPath)
	if err != nil {
		return config.Config{}, errors.Wrapf(err, "config file not found")
	}
	defer f.Close()

	var cfg config.Config

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return config.Config{}, errors.Wrapf(err, "error reading config")
	}

	return cfg, nil
}

func initDB(cfg config.Config) (*sqlx.DB, error) {
	dbAddress := os.Getenv("DATABASE_URL")
	if dbAddress == "" {
		dbAddress = cfg.DB.Address
	}

	// Connect SQL DB
	db, err := sqlx.Connect(cfg.DB.Driver, dbAddress)
	if err != nil {
		return nil, err
	}

	// Set db params
	db.SetMaxIdleConns(cfg.DB.MaxConns)
	db.SetMaxOpenConns(cfg.DB.MaxIdleConns)

	return db, nil
}

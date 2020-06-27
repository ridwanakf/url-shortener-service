package app

import (
	"log"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/ridwanakf/url-shortener-service/constant"
	"github.com/ridwanakf/url-shortener-service/internal/app/config"
	"gopkg.in/yaml.v2"
)

type UrlShortenerApp struct {
	Bridges  *Bridges
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

	app.Bridges, err = newBridges(&cfg)

	app.Repos, err = newRepos(app.Bridges, db)
	if err != nil {
		return nil, errors.Wrap(err, "errors invoking newRepos")
	}

	app.UseCases = newUsecases(app.Repos, &cfg)

	return app, nil
}

func (a *UrlShortenerApp) Close() []error {
	var errs []error

	errs = append(errs, a.Repos.Close()...)
	errs = append(errs, a.UseCases.Close()...)
	errs = append(errs, a.Bridges.Close()...)

	return errs
}

func readConfig(cfgPath string) (config.Config, error) {
	f, err := os.Open(cfgPath)
	if err != nil {
		return config.Config{}, errors.Wrapf(err, "config file not found")
	}
	defer f.Close()

	var cfg config.Config

	// Read from config file
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return config.Config{}, errors.Wrapf(err, "error reading config from file")
	}

	// Replace vars that exist in ENV
	if err := env.Parse(&cfg); err != nil {
		return config.Config{}, errors.Wrapf(err, "error reading config from ENV")
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

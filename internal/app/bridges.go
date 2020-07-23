package app

import (
	bridge "github.com/ridwanakf/bridges"
	"github.com/ridwanakf/bridges/redis"
	"github.com/ridwanakf/url-shortener-service/internal/app/config"
)

type Bridges struct {
	Redis bridge.Redis
}

func newBridges(cfg *config.Config) (*Bridges, error) {
	rd := redis.NewRedigo(redisConfigConverter(cfg.Redis))
	return &Bridges{
		Redis: rd,
	}, nil
}

func redisConfigConverter(config config.Redis) redis.ConfigOptions {
	return redis.ConfigOptions{
		Address:   config.Address,
		Timeout:   config.Timeout,
		MaxIdle:   config.MaxIdle,
		MaxActive: config.MaxActive,
	}
}

func (a *Bridges) Close() []error {
	var errs []error

	errs = append(errs, a.Redis.Close())

	return errs
}

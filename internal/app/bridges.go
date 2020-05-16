package app

import (
	"os"

	bridge "github.com/ridwanakf/go-bridges"
	"github.com/ridwanakf/go-bridges/json"
	"github.com/ridwanakf/go-bridges/redis"
	"github.com/ridwanakf/go-bridges/redisjson"
	"github.com/ridwanakf/url-shortener-service/internal/app/config"
)

type Bridges struct {
	Json      bridge.Json
	Redis     bridge.Redis
	RedisJson bridge.RedisJson
}

func newBridges(cfg *config.Config) (*Bridges, error) {
	js := json.NewJsoniter()

	redisAddress := os.Getenv("REDIS_URL")
	if redisAddress != "" {
		cfg.Redis.Address = redisAddress
	}

	rd := redis.NewRedigo(redisConfigConverter(cfg.Redis))

	return &Bridges{
		Json:      js,
		RedisJson: redisjson.NewRedisJson(rd, js),
		Redis:     rd,
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
	errs = append(errs, a.RedisJson.Close())

	return errs
}

package app

import (
	bridge "github.com/ridwanakf/url-shortener-service/internal"
	"github.com/ridwanakf/url-shortener-service/internal/app/config"
	"github.com/ridwanakf/url-shortener-service/internal/bridge/json"
	"github.com/ridwanakf/url-shortener-service/internal/bridge/redis"
	"github.com/ridwanakf/url-shortener-service/internal/bridge/redisjson"
	"os"
)

type Bridges struct {
	Json      bridge.Json
	Redis     bridge.Redis
	RedisJson bridge.RedisJson
}

func newBridges(cfg *config.Config) (*Bridges, error) {
	js := json.NewJsoniter()

	redisAddress := os.Getenv("REDIS_URL")
	if redisAddress == "" {
		redisAddress = cfg.Redis.Address
	}
	rd := redis.NewRedigo(redisAddress, cfg.Redis)

	return &Bridges{
		Json:      js,
		RedisJson: redisjson.NewRedisJson(rd, js),
		Redis:     rd,
	}, nil
}

func (a *Bridges) Close() []error {
	var errs []error

	errs = append(errs, a.Redis.Close())
	errs = append(errs, a.RedisJson.Close())

	return errs
}

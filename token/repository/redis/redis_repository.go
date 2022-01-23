package redis

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/madjiebimaa/nakafam/domain"
)

type redisTokenRepository struct {
	rdb *redis.Client
}

func NewRedisTokenRepository(rdb *redis.Client) domain.TokenRepository {
	return &redisTokenRepository{
		rdb,
	}
}

func (r *redisTokenRepository) Get(ctx context.Context, key string) (string, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
		return "", domain.ErrInternalServerError
	}

	return val, nil
}

func (r *redisTokenRepository) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	if err := r.rdb.Set(ctx, key, val, expiration).Err(); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (r *redisTokenRepository) Del(ctx context.Context, key string) error {
	if err := r.rdb.Del(ctx, key).Err(); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

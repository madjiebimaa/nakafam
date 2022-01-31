package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type configDB struct {
	host string
	port string
}

func NewConfigDB(
	host string,
	port string,
) *configDB {
	return &configDB{
		host,
		port,
	}
}

func (c *configDB) Init(ctx context.Context) *redis.Client {
	if c.host == "" || c.port == "" {
		log.Fatal("not configure the environment variables")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     c.host + ":" + c.port,
		Password: "",
		DB:       0,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatal(err)
	}

	return rdb
}

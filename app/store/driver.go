package store

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions/redis"
)

type configStore struct {
	host string
	port string
}

func NewConfigStore(
	host string,
	port string,
) *configStore {
	return &configStore{
		host,
		port,
	}
}

func (c *configStore) Init() redis.Store {
	store, err := redis.NewStore(10, "tcp", c.host+":"+c.port, "", []byte(os.Getenv("TOKEN_SECRET_SESSION")))
	if err != nil {
		log.Fatal(err)
	}

	return store
}

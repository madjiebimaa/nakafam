package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/nakafam/app/config"
	"github.com/madjiebimaa/nakafam/app/mongo"
	"github.com/madjiebimaa/nakafam/app/redis"
)

func main() {
	config.LoadEnv()
	r := gin.Default()

	ctx := context.TODO()
	mongoConfig := mongo.NewConfigDB(os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))
	mn := mongoConfig.Init(ctx)
	defer func() {
		if err := mn.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	redisConfig := redis.NewConfigDB(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	rdb := redisConfig.Init(ctx)
	defer func() {
		if err := rdb.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := r.Run(os.Getenv("SERVER_ADDRESS")); err != nil {
		log.Fatal(err)
	}
}

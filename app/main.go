package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/nakafam/app/config"
	"github.com/madjiebimaa/nakafam/app/mongo"
	"github.com/madjiebimaa/nakafam/app/redis"
	"github.com/madjiebimaa/nakafam/app/store"
	"github.com/madjiebimaa/nakafam/constant"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	config.LoadEnv()
	r := gin.Default()

	ctx := context.Background()
	mongoConfig := mongo.NewConfigDB(os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))
	cl := mongoConfig.Init(ctx)
	defer func() {

		if err := cl.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	if err := cl.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	redisConfig := redis.NewConfigDB(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	rdb := redisConfig.Init(ctx)
	defer func() {
		if err := rdb.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	storeConfig := store.NewConfigStore(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	store := storeConfig.Init()

	r.Use(sessions.Sessions(constant.SESSION_NAME, store))

	// db := cl.Database(os.Getenv("DATABASE_NAME"))
	// collNakama := db.Collection(os.Getenv("COLLECTION_NAKAMA"))
	// collFamily := db.Collection(os.Getenv("COLLECTION_FAMILY"))

	// timeoutContextEnv, _ := strconv.Atoi(os.Getenv("TIMEOUT_CONTEXT"))
	// timeoutContext := time.Duration(timeoutContextEnv) * time.Second

	if err := r.Run(os.Getenv("SERVER_ADDRESS")); err != nil {
		log.Fatal(err)
	}
}

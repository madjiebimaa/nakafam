package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/nakafam/app/config"
	"github.com/madjiebimaa/nakafam/app/mongo"
	"github.com/madjiebimaa/nakafam/app/redis"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	config.LoadEnv()
	r := gin.Default()

	ctx := context.Background()
	mongoConfig := mongo.NewConfigDB(os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))
	mn := mongoConfig.Init(ctx)
	defer func() {

		if err := mn.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	if err := mn.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	redisConfig := redis.NewConfigDB(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	rdb := redisConfig.Init(ctx)
	defer func() {
		if err := rdb.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// nakafamDB := mn.Database(constant.DATABASE_NAME)
	// nakamaColl := nakafamDB.Collection(constant.NAKAMA_COLLECTION)
	// familyColl := nakafamDB.Collection(constant.FAMILY_COLLECTION)

	// timeoutContextEnv, _ := strconv.Atoi(os.Getenv("TIMEOUT_CONTEXT"))
	// timeoutContext := time.Duration(timeoutContextEnv) * time.Second

	if err := r.Run(os.Getenv("SERVER_ADDRESS")); err != nil {
		log.Fatal(err)
	}
}

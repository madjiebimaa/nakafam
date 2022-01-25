package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/nakafam/app/config"
	"github.com/madjiebimaa/nakafam/app/mongo"
	"github.com/madjiebimaa/nakafam/app/redis"
	"github.com/madjiebimaa/nakafam/app/route"
	"github.com/madjiebimaa/nakafam/app/store"
	"github.com/madjiebimaa/nakafam/constant"
	_familyHttpDelivery "github.com/madjiebimaa/nakafam/family/delivery/http"
	_familyRepo "github.com/madjiebimaa/nakafam/family/repository/mongo"
	_familyUCase "github.com/madjiebimaa/nakafam/family/usecase"
	_nakamaHttpDelivery "github.com/madjiebimaa/nakafam/nakama/delivery/http"
	_nakamaRepo "github.com/madjiebimaa/nakafam/nakama/repository/mongo"
	_nakamaUCase "github.com/madjiebimaa/nakafam/nakama/usecase"
	_tokenRepo "github.com/madjiebimaa/nakafam/token/repository/redis"
	_userHttpDelivery "github.com/madjiebimaa/nakafam/user/delivery/http"
	_userRepo "github.com/madjiebimaa/nakafam/user/repository/mongo"
	_userUCase "github.com/madjiebimaa/nakafam/user/usecase"
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

	db := cl.Database(os.Getenv("DATABASE_NAME"))
	collUser := db.Collection(os.Getenv("COLLECTION_USER"))
	collNakama := db.Collection(os.Getenv("COLLECTION_NAKAMA"))
	collFamily := db.Collection(os.Getenv("COLLECTION_FAMILY"))

	timeoutContextEnv, _ := strconv.Atoi(os.Getenv("TIMEOUT_CONTEXT"))
	timeoutContext := time.Duration(timeoutContextEnv) * time.Second

	userRepo := _userRepo.NewMongoUserRepository(collUser)
	nakamaRepo := _nakamaRepo.NewMongoNakamaRepository(collNakama)
	familyRepo := _familyRepo.NewMongoFamilyRepository(collFamily)
	tokenRepo := _tokenRepo.NewRedisTokenRepository(rdb)

	userUCase := _userUCase.NewUserUseCase(userRepo, tokenRepo, timeoutContext)
	nakamaUCase := _nakamaUCase.NewNakamaUseCase(nakamaRepo, userRepo, familyRepo, timeoutContext)
	familyUCase := _familyUCase.NewFamilyUseCase(familyRepo, nakamaRepo, timeoutContext)

	userHttpDelivery := _userHttpDelivery.NewUserHandler(userUCase)
	nakamaHttpDelivery := _nakamaHttpDelivery.NewNakamaDelivery(nakamaUCase)
	familyHttpDelivery := _familyHttpDelivery.NewFamilyDelivery(familyUCase)

	ro := route.NewRoutes(userHttpDelivery, nakamaHttpDelivery, familyHttpDelivery)
	ro.Init(r)

	if err := r.Run(os.Getenv("SERVER_ADDRESS")); err != nil {
		log.Fatal(err)
	}
}

package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type configDB struct {
	host     string
	port     string
	user     string
	password string
}

func NewConfigDB(
	host string,
	port string,
	user string,
	password string,
) *configDB {
	return &configDB{
		host,
		port,
		user,
		password,
	}
}

func (c *configDB) Init(ctx context.Context) *mongo.Client {
	if c.host == "" || c.port == "" || c.user == "" || c.password == "" {
		log.Fatal("not configure the environment variables")
	}

	credential := options.Credential{
		Username: c.user,
		Password: c.password,
	}

	cl, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+c.host+":"+c.port).SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}

	return cl
}

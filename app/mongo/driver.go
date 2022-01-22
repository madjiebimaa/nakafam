package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (c *configDB) Init(ctx context.Context) *mongo.Client {
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+c.host+":"+c.port))
	if err != nil {
		log.Fatal(err)
	}

	return cl
}

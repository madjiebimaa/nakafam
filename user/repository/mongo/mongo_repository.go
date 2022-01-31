package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/madjiebimaa/nakafam/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepository struct {
	coll *mongo.Collection
}

func NewMongoUserRepository(coll *mongo.Collection) domain.UserRepository {
	return &mongoUserRepository{
		coll,
	}
}

func (m *mongoUserRepository) Register(ctx context.Context, user *domain.User) error {
	if _, err :=
		m.coll.InsertOne(ctx, user); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (m *mongoUserRepository) GetByID(ctx context.Context, id primitive.ObjectID) (domain.User, error) {
	var user domain.User
	filter := bson.D{{Key: "_id", Value: id}}
	err := m.coll.FindOne(ctx, filter).Decode(&user)
	if err == nil {
		return user, nil
	} else if err == mongo.ErrNoDocuments {
		return domain.User{}, nil
	} else {
		fmt.Println(err)
		return domain.User{}, domain.ErrInternalServerError
	}
}

func (m *mongoUserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	filter := bson.D{{Key: "email", Value: email}}
	err := m.coll.FindOne(ctx, filter).Decode(&user)
	if err == nil {
		return user, nil
	} else if err == mongo.ErrNoDocuments {
		return domain.User{}, nil
	} else {
		fmt.Println(err)
		return domain.User{}, domain.ErrInternalServerError
	}
}

func (m *mongoUserRepository) ToLeaderRole(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}
	updater := bson.D{
		{Key: "$set", Value: bson.D{{Key: "role", Value: "leader"}}},
	}
	if _, err := m.coll.UpdateOne(ctx, filter, updater); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

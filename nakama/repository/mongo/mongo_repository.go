package mongo

import (
	"context"
	"log"

	"github.com/madjiebimaa/nakafam/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoNakamaRepository struct {
	coll *mongo.Collection
}

func NewMongoNakamaRepository(coll *mongo.Collection) domain.NakamaRepository {
	return &mongoNakamaRepository{
		coll,
	}
}

func (m *mongoNakamaRepository) Create(ctx context.Context, nakama *domain.Nakama) error {
	if _, err := m.coll.InsertOne(ctx, nakama); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (m *mongoNakamaRepository) Update(ctx context.Context, nakama *domain.Nakama) error {
	filter := bson.D{{Key: "_id", Value: nakama.ID}}
	updater := bson.D{
		{Key: "$set", Value: bson.D{{Key: "name", Value: nakama.Name}}},
		{Key: "$set", Value: bson.D{{Key: "profile_image", Value: nakama.ProfileImage}}},
		{Key: "$set", Value: bson.D{{Key: "description", Value: nakama.Description}}},
		{Key: "$set", Value: bson.D{{Key: "social_media", Value: nakama.SocialMedia}}},
		{Key: "$set", Value: bson.D{{Key: "updated_at", Value: nakama.UpdatedAt}}},
	}

	if _, err := m.coll.UpdateOne(ctx, filter, updater); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (m *mongoNakamaRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}
	if _, err := m.coll.DeleteOne(ctx, filter); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (m *mongoNakamaRepository) GetByID(ctx context.Context, id primitive.ObjectID) (domain.Nakama, error) {
	var nakama domain.Nakama
	filter := bson.D{{Key: "_id", Value: id}}
	if err := m.coll.FindOne(ctx, filter).Decode(&nakama); err != nil {
		log.Fatal(err)
		return domain.Nakama{}, domain.ErrInternalServerError
	}

	return nakama, nil
}

func (m *mongoNakamaRepository) GetAll(ctx context.Context) ([]domain.Nakama, error) {
	opts := options.Find().SetSort(bson.D{{Key: "name", Value: -1}})
	cur, err := m.coll.Find(ctx, bson.D{{}}, opts)
	if err != nil {
		log.Fatal(err)
		return nil, domain.ErrInternalServerError
	}
	defer cur.Close(ctx)

	var nakamas []domain.Nakama
	if err := cur.All(ctx, &nakamas); err != nil {
		log.Fatal(err)
		return nil, domain.ErrInternalServerError
	}

	if cur.Err() != nil {
		log.Fatal(cur.Err())
		return nil, domain.ErrInternalServerError
	}

	return nakamas, nil
}

func (m *mongoNakamaRepository) RegisterToFamily(ctx context.Context, id primitive.ObjectID, familyID primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}
	updater := bson.D{
		{Key: "$set", Value: bson.D{{Key: "family_ud", Value: familyID}}},
	}

	if _, err := m.coll.UpdateOne(ctx, filter, updater); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

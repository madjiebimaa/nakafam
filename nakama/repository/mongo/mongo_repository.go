package mongo

import (
	"context"
	"log"

	"github.com/madjiebimaa/nakafam/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoNakamaRepository struct {
	db *mongo.Collection
}

func NewMongoNakamaRepository(db *mongo.Collection) domain.NakamaRepository {
	return &mongoNakamaRepository{
		db,
	}
}

func (m *mongoNakamaRepository) Create(ctx context.Context, nakama *domain.Nakama) error {
	if _, err := m.db.InsertOne(ctx, nakama); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (m *mongoNakamaRepository) Update(ctx context.Context, nakama *domain.Nakama) error {
	if _, err := m.db.UpdateByID(ctx, nakama.ID, nakama); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (m *mongoNakamaRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	if _, err := m.db.DeleteOne(ctx, domain.Nakama{ID: id}); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (m *mongoNakamaRepository) GetByID(ctx context.Context, id primitive.ObjectID) (domain.Nakama, error) {
	var nakama domain.Nakama
	if err := m.db.FindOne(ctx, domain.Nakama{ID: id}).Decode(&nakama); err != nil {
		log.Fatal(err)
		return nakama, domain.ErrInternalServerError
	}

	return nakama, nil
}

func (m *mongoNakamaRepository) GetByName(ctx context.Context, name string) (domain.Nakama, error) {
	var nakama domain.Nakama
	if err := m.db.FindOne(ctx, domain.Nakama{Name: name}).Decode(&nakama); err != nil {
		log.Fatal(err)
		return nakama, domain.ErrInternalServerError
	}

	return nakama, nil
}

func (m *mongoNakamaRepository) GetAll(ctx context.Context) ([]domain.Nakama, error) {
	cur, err := m.db.Find(ctx, bson.M{})
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

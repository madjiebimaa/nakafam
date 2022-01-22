package mongo

import (
	"context"
	"log"

	"github.com/madjiebimaa/nakafam/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Collection
}

func NewMongoRepository(db *mongo.Collection) domain.NakamaRepository {
	return &mongoRepository{
		db,
	}
}

func (m *mongoRepository) Create(ctx context.Context, nakama *domain.Nakama) error {
	if _, err := m.db.InsertOne(ctx, nakama); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (m *mongoRepository) Update(ctx context.Context, nakama *domain.Nakama) error {
	if _, err := m.db.UpdateByID(ctx, nakama.ID, domain.Nakama{
		SocialMedia:  nakama.SocialMedia,
		ProfileImage: nakama.ProfileImage,
		Description:  nakama.Description,
		UpdatedAt:    nakama.UpdatedAt,
	}); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (m *mongoRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	if _, err := m.db.DeleteOne(ctx, domain.Nakama{ID: id}); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (m *mongoRepository) GetByID(ctx context.Context, id primitive.ObjectID) (domain.Nakama, error) {
	var nakama domain.Nakama
	if err := m.db.FindOne(ctx, domain.Nakama{ID: id}).Decode(&nakama); err != nil {
		log.Fatal(err)
		return nakama, domain.ErrInternalServerError
	}

	return nakama, nil
}

func (m *mongoRepository) GetByName(ctx context.Context, name string) (domain.Nakama, error) {
	var nakama domain.Nakama
	if err := m.db.FindOne(ctx, domain.Nakama{Name: name}).Decode(&nakama); err != nil {
		log.Fatal(err)
		return nakama, domain.ErrInternalServerError
	}

	return nakama, nil
}

func (m *mongoRepository) GetAll(ctx context.Context) ([]domain.Nakama, error) {
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

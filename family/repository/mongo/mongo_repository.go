package mongo

import (
	"context"
	"log"

	"github.com/madjiebimaa/nakafam/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoFamilyRepository struct {
	db *mongo.Collection
}

func NewMongoFamilyRepository(db *mongo.Collection) domain.FamilyRepository {
	return &mongoFamilyRepository{
		db,
	}
}

func (m *mongoFamilyRepository) Create(ctx context.Context, family *domain.Family) error {
	if _, err := m.db.InsertOne(ctx, family); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (m *mongoFamilyRepository) Update(ctx context.Context, family *domain.Family) error {
	if _, err := m.db.UpdateByID(ctx, family.ID, family); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (m *mongoFamilyRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	if _, err := m.db.DeleteOne(ctx, domain.Family{ID: id}); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (m *mongoFamilyRepository) GetByID(ctx context.Context, id primitive.ObjectID) (domain.Family, error) {
	var family domain.Family
	if err := m.db.FindOne(ctx, domain.Family{ID: id}).Decode(&family); err != nil {
		log.Fatal(err)
		return family, domain.ErrInternalServerError
	}

	if family.ID == primitive.NilObjectID {
		return family, domain.ErrNotFound
	}

	return family, nil
}

func (m *mongoFamilyRepository) GetByName(ctx context.Context, name string) (domain.Family, error) {
	var family domain.Family
	if err := m.db.FindOne(ctx, domain.Family{Name: name}).Decode(&family); err != nil {
		log.Fatal(err)
		return family, domain.ErrInternalServerError
	}

	if family.ID == primitive.NilObjectID {
		return family, domain.ErrNotFound
	}

	return family, nil
}

func (m *mongoFamilyRepository) GetAll(ctx context.Context) ([]domain.Family, error) {
	cur, err := m.db.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil, domain.ErrInternalServerError
	}
	defer cur.Close(ctx)

	var families []domain.Family
	if err := cur.All(ctx, &families); err != nil {
		log.Fatal(err)
		return nil, domain.ErrInternalServerError
	}

	if families == nil {
		return families, domain.ErrNotFound
	}

	if err := cur.Err(); err != nil {
		log.Fatal(cur.Err())
		return nil, domain.ErrInternalServerError
	}

	return families, nil
}

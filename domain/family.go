package domain

import (
	"context"
	"time"

	"github.com/madjiebimaa/nakafam/family/delivery/http/requests"
	"github.com/madjiebimaa/nakafam/family/delivery/http/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Family struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Password    string             `json:"password"`
	FamilyImage string             `json:"family_image" bson:"family_image"`
	Nakamas     []Nakama           `json:"nakamas" bson:"nakamas"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type FamilyRepository interface {
	Create(ctx context.Context, family *Family) error
	Update(ctx context.Context, family *Family) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Family, error)
	GetByName(ctx context.Context, name string) (Family, error)
	GetAll(ctx context.Context) ([]Family, error)
}

type FamilyUseCase interface {
	Create(c context.Context, req *requests.FamilyCreate) error
	Update(c context.Context, req *requests.FamilyUpdate) error
	Delete(c context.Context, id primitive.ObjectID) error
	GetByID(c context.Context, id primitive.ObjectID) (responses.FamilyBase, error)
	GetByName(c context.Context, name string) (responses.FamilyBase, error)
	GetAll(c context.Context) ([]responses.FamilyBase, error)
}

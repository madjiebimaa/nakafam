package domain

import (
	"context"
	"time"

	_familyReq "github.com/madjiebimaa/nakafam/family/delivery/http/requests"
	_familyRes "github.com/madjiebimaa/nakafam/family/delivery/http/responses"
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
	Create(c context.Context, req *_familyReq.FamilyCreate) error
	Update(c context.Context, req *_familyReq.FamilyUpdate) error
	Delete(c context.Context, id primitive.ObjectID) error
	GetByID(c context.Context, id primitive.ObjectID) (_familyRes.FamilyBase, error)
	GetByName(c context.Context, name string) (_familyRes.FamilyBase, error)
	GetAll(c context.Context) ([]_familyRes.FamilyBase, error)
}

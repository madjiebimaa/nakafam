package domain

import (
	"context"
	"time"

	"github.com/madjiebimaa/nakafam/nakama/delivery/http/requests"
	"github.com/madjiebimaa/nakafam/nakama/delivery/http/responses"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SocialMedia struct {
	Blogs     string `json:"blogs"`
	Portfolio string `json:"portfolio"`
	Email     string `json:"email"`
	Github    string `json:"github"`
	Linkedin  string `json:"linkedin"`
	Twitter   string `json:"twitter"`
	Discord   string `json:"discord"`
	Youtube   string `json:"youtube"`
	Instagram string `json:"instagram"`
}

type Nakama struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	FamilyID     primitive.ObjectID `json:"family" bson:"family_id"`
	Name         string             `json:"name" bson:"name"`
	UserName     string             `json:"username" bson:"username"`
	ProfileImage string             `json:"profile_image" bson:"profile_image"`
	Description  string             `json:"description" bson:"description"`
	SocialMedia  *SocialMedia       `json:"social_media" bson:"social_media"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

type NakamaRepository interface {
	Create(ctx context.Context, nakama *Nakama) error
	Update(ctx context.Context, nakama *Nakama) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Nakama, error)
	GetByName(ctx context.Context, name string) (Nakama, error)
	GetByFamilyID(ctx context.Context, familyID primitive.ObjectID) ([]Nakama, error)
	RegisterToFamily(ctx context.Context, id primitive.ObjectID, familyID primitive.ObjectID) error
}

type NakamaUseCase interface {
	Create(c context.Context, req *requests.NakamaCreate) error
	Update(c context.Context, req *requests.NakamaUpdate) error
	Delete(c context.Context, id primitive.ObjectID) error
	GetByID(c context.Context, id primitive.ObjectID) (responses.NakamaBase, error)
	GetByName(c context.Context, name string) (responses.NakamaBase, error)
	GetByFamilyID(c context.Context, familyID primitive.ObjectID) ([]responses.NakamaBase, error)
	RegisterToFamily(c context.Context, req *requests.NakamaRegisterToFamily) error
}

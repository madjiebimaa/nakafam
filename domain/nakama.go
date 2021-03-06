package domain

import (
	"context"
	"time"

	_nakamaReq "github.com/madjiebimaa/nakafam/nakama/delivery/http/requests"
	_nakamaRes "github.com/madjiebimaa/nakafam/nakama/delivery/http/responses"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SocialMedia struct {
	Blogs     string `json:"blogs" bson:"blogs"`
	Portfolio string `json:"portfolio" bson:"portfolio"`
	Email     string `json:"email" bson:"email"`
	Github    string `json:"github" bson:"github"`
	Linkedin  string `json:"linkedin" bson:"linkedin"`
	Twitter   string `json:"twitter" bson:"twitter"`
	Discord   string `json:"discord" bson:"discord"`
	Youtube   string `json:"youtube" bson:"youtube"`
	Instagram string `json:"instagram" bson:"instagram"`
}

type Nakama struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
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
	GetAll(ctx context.Context) ([]Nakama, error)
	RegisterToFamily(ctx context.Context, id primitive.ObjectID, familyID primitive.ObjectID) error
}

type NakamaUseCase interface {
	Create(c context.Context, req *_nakamaReq.NakamaCreate) (Nakama, error)
	Update(c context.Context, req *_nakamaReq.NakamaUpdate) error
	Delete(c context.Context, req *_nakamaReq.NakamaDelete) error
	GetByID(c context.Context, id primitive.ObjectID) (_nakamaRes.NakamaBase, error)
	GetAll(c context.Context) ([]_nakamaRes.NakamaBase, error)
	RegisterToFamily(c context.Context, req *_nakamaReq.NakamaRegisterToFamily) error
}

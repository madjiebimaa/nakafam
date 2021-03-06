package domain

import (
	"context"
	"time"

	_userReq "github.com/madjiebimaa/nakafam/user/delivery/http/requests"
	_userRes "github.com/madjiebimaa/nakafam/user/delivery/http/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Role      string             `json:"role" bson:"role"` // default is staff but can request as a leader if using email confirmation
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type UserRepository interface {
	Register(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id primitive.ObjectID) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	ToLeaderRole(ctx context.Context, id primitive.ObjectID) error
}

type UserUseCase interface {
	Register(c context.Context, req *_userReq.UserRegisterOrLogin) error
	Login(c context.Context, req *_userReq.UserRegisterOrLogin) (_userRes.UserBase, error)
	UpgradeRole(c context.Context, id primitive.ObjectID) (_userRes.UserUpgradeRole, error)
	ToLeaderRole(c context.Context, req *_userReq.UserToLeaderRole) error
	Me(c context.Context, id primitive.ObjectID) (_userRes.UserBase, error)
}

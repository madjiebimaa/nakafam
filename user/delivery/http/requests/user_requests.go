package requests

import (
	"github.com/madjiebimaa/nakafam/nakama/delivery/http/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRegisterOrLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterAsLeader struct {
	Token string `json:"token"`
}

type UserCreateNakama struct {
	UserID       primitive.ObjectID         `json:"user_id"`
	Name         string                     `json:"name"`
	UserName     string                     `json:"username"`
	ProfileImage string                     `json:"profile_image"`
	Description  string                     `json:"description"`
	SocialMedia  *responses.SocialMediaBase `json:"social_media"`
}

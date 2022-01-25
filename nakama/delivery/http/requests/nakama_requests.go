package requests

import (
	_nakamaRes "github.com/madjiebimaa/nakafam/nakama/delivery/http/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NakamaCreate struct {
	UserID       primitive.ObjectID          `json:"user_id"`
	Name         string                      `json:"name"`
	UserName     string                      `json:"username"`
	ProfileImage string                      `json:"profile_image"`
	Description  string                      `json:"description"`
	SocialMedia  *_nakamaRes.SocialMediaBase `json:"social_media"`
}

type NakamaUpdate struct {
	NakamaID     primitive.ObjectID          `json:"nakama_id"`
	Name         string                      `json:"name"`
	ProfileImage string                      `json:"profile_image"`
	Description  string                      `json:"description"`
	SocialMedia  *_nakamaRes.SocialMediaBase `json:"social_media"`
}

type NakamaRegisterToFamily struct {
	NakamaID primitive.ObjectID `json:"nakama_id"`
	FamilyID primitive.ObjectID `json:"family_id"`
	Password string             `json:"password"`
}

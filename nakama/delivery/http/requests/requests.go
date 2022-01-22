package requests

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SocialMediaBase struct {
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

type NakamaCreate struct {
	Name         string           `json:"name"`
	UserName     string           `json:"username"`
	ProfileImage string           `json:"profile_image"`
	Description  string           `json:"description"`
	SocialMedia  *SocialMediaBase `json:"social_media"`
}

type NakamaUpdate struct {
	ID           primitive.ObjectID `json:"id"`
	FamilyID     primitive.ObjectID `json:"family_id"`
	Name         string             `json:"name"`
	UserName     string             `json:"username"`
	ProfileImage string             `json:"profile_image"`
	Description  string             `json:"description"`
	SocialMedia  *SocialMediaBase   `json:"social_media"`
}

type NakamaRegisterToFamily struct {
	ID       primitive.ObjectID `json:"id"`
	FamilyID primitive.ObjectID `json:"family_id"`
	Password string             `json:"password"`
}

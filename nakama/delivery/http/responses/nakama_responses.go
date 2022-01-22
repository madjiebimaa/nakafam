package responses

import (
	"time"

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

type NakamaBase struct {
	ID           primitive.ObjectID `json:"id"`
	FamilyID     primitive.ObjectID `json:"family"`
	Name         string             `json:"name"`
	UserName     string             `json:"username"`
	ProfileImage string             `json:"profile_image"`
	Description  string             `json:"description"`
	SocialMedia  *SocialMediaBase   `json:"social_media"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}

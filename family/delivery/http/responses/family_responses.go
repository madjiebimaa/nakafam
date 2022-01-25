package responses

import (
	"time"

	_nakamaRes "github.com/madjiebimaa/nakafam/nakama/delivery/http/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FamilyBase struct {
	ID          primitive.ObjectID      `json:"id" `
	Name        string                  `json:"name" `
	Password    string                  `json:"password"`
	FamilyImage string                  `json:"family_image" `
	Nakamas     []_nakamaRes.NakamaBase `json:"nakamas" `
	CreatedAt   time.Time               `json:"created_at" `
	UpdatedAt   time.Time               `json:"updated_at" `
}

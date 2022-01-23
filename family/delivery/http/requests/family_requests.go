package requests

import (
	"github.com/madjiebimaa/nakafam/nakama/delivery/http/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FamilyCreate struct {
	NakamaID    primitive.ObjectID     `json:"nakama_id"`
	Name        string                 `json:"name"`
	Password    string                 `json:"password"`
	FamilyImage string                 `json:"family_image"`
	Nakamas     []responses.NakamaBase `json:"nakamas"`
}

type FamilyUpdate struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Password    string             `json:"password"`
	FamilyImage string             `json:"family_image"`
}
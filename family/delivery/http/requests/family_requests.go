package requests

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FamilyCreate struct {
	NakamaID    primitive.ObjectID `json:"nakama_id"`
	Name        string             `json:"name"`
	Password    string             `json:"password"`
	FamilyImage string             `json:"family_image"`
}

type FamilyUpdate struct {
	FamilyID    primitive.ObjectID `json:"family_id"`
	NakamaID    primitive.ObjectID `json:"nakama_id"`
	Name        string             `json:"name"`
	FamilyImage string             `json:"family_image"`
}

type FamilyDelete struct {
	FamilyID primitive.ObjectID `json:"family_id"`
	NakamaID primitive.ObjectID `json:"nakama_id"`
}

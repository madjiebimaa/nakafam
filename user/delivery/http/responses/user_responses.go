package responses

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBase struct {
	ID        primitive.ObjectID `json:"id"`
	Email     string             `json:"email"`
	Role      string             `json:"role"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type UserUpgradeRole struct {
	URL string `json:"url"`
}

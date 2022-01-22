package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"` // default is staff but can request as a leader if using email confirmation
	CreatedAt time.Time `json:"created_at"`
}

type UserRepository interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (User, error)
	UpdateRepo(ctx context.Context, id primitive.ObjectID, role string) error
}

type UserUseCase interface {
}

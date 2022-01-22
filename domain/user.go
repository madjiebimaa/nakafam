package domain

import "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type UserRepository interface {
}

type UserUseCase interface {
}

package domain

import (
	"context"
	"time"
)

type Family struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Nakamas   []Nakama  `json:"nakamas"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FamilyRepository interface {
	Create(ctx context.Context, family *Family) error
	Update(ctx context.Context, family *Family) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (Family, error)
	GetByName(ctx context.Context, name string) (Family, error)
	GetAll(ctx context.Context) ([]Family, error)
}

type FamilyUseCase interface {
	Create(ctx context.Context, family *Family) error
	Update(ctx context.Context, family *Family) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (Family, error)
	GetByName(ctx context.Context, name string) (Family, error)
	GetAll(ctx context.Context) ([]Family, error)
}

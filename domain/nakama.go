package domain

import (
	"context"
	"time"
)

type SocialMedia struct {
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

type Nakama struct {
	ID           string      `json:"_id"`
	Name         string      `json:"name"`
	UserName     string      `json:"username"`
	ProfileImage string      `json:"profile_image"`
	Description  string      `json:"description"`
	SocialMedia  SocialMedia `json:"social_media"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

type NakamaRepository interface {
	Create(ctx context.Context, nakama *Nakama) error
	Update(ctx context.Context, nakama *Nakama) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (Nakama, error)
	GetByName(ctx context.Context, name string) (Nakama, error)
	GetAll(ctx context.Context) ([]Nakama, error)
}

type NakamaUseCase interface {
	Create(ctx context.Context, nakama *Nakama) error
	Update(ctx context.Context, nakama *Nakama) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (Nakama, error)
	GetByName(ctx context.Context, name string) (Nakama, error)
	GetAll(ctx context.Context) ([]Nakama, error)
}

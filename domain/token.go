package domain

import (
	"context"
	"time"
)

type Token struct {
	ActiveToken  string `json:"active_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenRepository interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Del(ctx context.Context, key string) error
}

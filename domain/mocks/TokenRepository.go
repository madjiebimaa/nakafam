package mocks

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type TokenRepository struct {
	mock.Mock
}

func (t *TokenRepository) Get(ctx context.Context, key string) (string, error) {
	ret := t.Called(ctx, key)

	var r0 string
	if ref, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = ref(ctx, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = ref(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (t *TokenRepository) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	ret := t.Called(ctx, key, val, expiration)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = ref(ctx, key, val, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (t *TokenRepository) Del(ctx context.Context, key string) error {
	ret := t.Called(ctx, key)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = ref(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

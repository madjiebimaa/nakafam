package mocks

import (
	"context"

	"github.com/madjiebimaa/nakafam/domain"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	mock.Mock
}

func (u *UserRepository) Register(ctx context.Context, user *domain.User) error {
	ret := u.Called(ctx, user)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, *domain.User) error); ok {
		r0 = ref(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (u *UserRepository) GetByID(ctx context.Context, id primitive.ObjectID) (domain.User, error) {
	ret := u.Called(ctx, id)

	var r0 domain.User
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) domain.User); ok {
		r0 = ref(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = ref(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	ret := u.Called(ctx, email)

	var r0 domain.User
	if ref, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = ref(ctx, email)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = ref(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (u *UserRepository) ToLeaderRole(ctx context.Context, id primitive.ObjectID) error {
	ret := u.Called(ctx, id)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = ref(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

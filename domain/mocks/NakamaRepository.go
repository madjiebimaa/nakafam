package mocks

import (
	"context"

	"github.com/madjiebimaa/nakafam/domain"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NakamaRepository struct {
	mock.Mock
}

func (n *NakamaRepository) Create(ctx context.Context, nakama *domain.Nakama) error {
	ret := n.Called(ctx, nakama)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, *domain.Nakama) error); ok {
		r0 = ref(ctx, nakama)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (n *NakamaRepository) Update(ctx context.Context, nakama *domain.Nakama) error {
	ret := n.Called(ctx, nakama)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, *domain.Nakama) error); ok {
		r0 = ref(ctx, nakama)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (n *NakamaRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	ret := n.Called(ctx, id)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = ref(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (n *NakamaRepository) GetByID(ctx context.Context, id primitive.ObjectID) (domain.Nakama, error) {
	ret := n.Called(ctx, id)

	var r0 domain.Nakama
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) domain.Nakama); ok {
		r0 = ref(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Nakama)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = ref(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (n *NakamaRepository) GetAll(ctx context.Context) ([]domain.Nakama, error) {
	ret := n.Called(ctx)

	var r0 []domain.Nakama
	if ref, ok := ret.Get(0).(func(context.Context) []domain.Nakama); ok {
		r0 = ref(ctx)
	} else {
		r0 = ret.Get(0).([]domain.Nakama)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = ref(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (n *NakamaRepository) RegisterToFamily(ctx context.Context, id primitive.ObjectID, familyID primitive.ObjectID) error {
	ret := n.Called(ctx, id, familyID)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, primitive.ObjectID) error); ok {
		r0 = ref(ctx, id, familyID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

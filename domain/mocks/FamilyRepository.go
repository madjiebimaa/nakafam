package mocks

import (
	"context"

	"github.com/madjiebimaa/nakafam/domain"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FamilyRepository struct {
	mock.Mock
}

func (f *FamilyRepository) Create(ctx context.Context, family *domain.Family) error {
	ret := f.Called(ctx, family)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, *domain.Family) error); ok {
		r0 = ref(ctx, family)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (f *FamilyRepository) Update(ctx context.Context, family *domain.Family) error {
	ret := f.Called(ctx, family)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, *domain.Family) error); ok {
		r0 = ref(ctx, family)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (f *FamilyRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	ret := f.Called(ctx, id)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = ref(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (f *FamilyRepository) GetByID(ctx context.Context, id primitive.ObjectID) (domain.Family, error) {
	ret := f.Called(ctx, id)

	var r0 domain.Family
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) domain.Family); ok {
		r0 = ref(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Family)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = ref(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (f *FamilyRepository) GetAll(ctx context.Context) ([]domain.Family, error) {
	ret := f.Called(ctx)

	var r0 []domain.Family
	if ref, ok := ret.Get(0).(func(context.Context) []domain.Family); ok {
		r0 = ref(ctx)
	} else {
		r0 = ret.Get(0).([]domain.Family)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = ref(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

package mocks

import (
	"context"

	_familyReq "github.com/madjiebimaa/nakafam/family/delivery/http/requests"
	_familyRes "github.com/madjiebimaa/nakafam/family/delivery/http/responses"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FamilyUCase struct {
	mock.Mock
}

func (f *FamilyUCase) Create(ctx context.Context, req *_familyReq.FamilyCreate) error {
	ret := f.Called(ctx, req)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, *_familyReq.FamilyCreate) error); ok {
		r0 = ref(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (f *FamilyUCase) Update(ctx context.Context, req *_familyReq.FamilyUpdate) error {
	ret := f.Called(ctx, req)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, *_familyReq.FamilyUpdate) error); ok {
		r0 = ref(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (f *FamilyUCase) Delete(ctx context.Context, id primitive.ObjectID) error {
	ret := f.Called(ctx, id)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = ref(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (f *FamilyUCase) GetByID(ctx context.Context, id primitive.ObjectID) (_familyRes.FamilyBase, error) {
	ret := f.Called(ctx, id)

	var r0 _familyRes.FamilyBase
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) _familyRes.FamilyBase); ok {
		r0 = ref(ctx, id)
	} else {
		r0 = ret.Get(0).(_familyRes.FamilyBase)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = ref(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (f *FamilyUCase) GetAll(ctx context.Context) ([]_familyRes.FamilyBase, error) {
	ret := f.Called(ctx)

	var r0 []_familyRes.FamilyBase
	if ref, ok := ret.Get(0).(func(context.Context) []_familyRes.FamilyBase); ok {
		r0 = ref(ctx)
	} else {
		r0 = ret.Get(0).([]_familyRes.FamilyBase)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = ref(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

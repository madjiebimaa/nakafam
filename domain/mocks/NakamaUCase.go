package mocks

import (
	"context"

	_nakamaReq "github.com/madjiebimaa/nakafam/nakama/delivery/http/requests"
	_nakamaRes "github.com/madjiebimaa/nakafam/nakama/delivery/http/responses"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NakamaUCase struct {
	mock.Mock
}

func (n *NakamaUCase) Create(ctx context.Context, req *_nakamaReq.NakamaCreate) error {
	ret := n.Called(ctx, req)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, *_nakamaReq.NakamaCreate) error); ok {
		r0 = ref(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (n *NakamaUCase) Update(ctx context.Context, req *_nakamaReq.NakamaUpdate) error {
	ret := n.Called(ctx, req)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, *_nakamaReq.NakamaUpdate) error); ok {
		r0 = ref(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (n *NakamaUCase) Delete(ctx context.Context, id primitive.ObjectID) error {
	ret := n.Called(ctx, id)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = ref(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (n *NakamaUCase) GetByID(ctx context.Context, id primitive.ObjectID) (_nakamaRes.NakamaBase, error) {
	ret := n.Called(ctx, id)

	var r0 _nakamaRes.NakamaBase
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) _nakamaRes.NakamaBase); ok {
		r0 = ref(ctx, id)
	} else {
		r0 = ret.Get(0).(_nakamaRes.NakamaBase)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = ref(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (n *NakamaUCase) GetAll(ctx context.Context) ([]_nakamaRes.NakamaBase, error) {
	ret := n.Called(ctx)

	var r0 []_nakamaRes.NakamaBase
	if ref, ok := ret.Get(0).(func(context.Context) []_nakamaRes.NakamaBase); ok {
		r0 = ref(ctx)
	} else {
		r0 = ret.Get(0).([]_nakamaRes.NakamaBase)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = ref(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (n *NakamaUCase) RegisterToFamily(ctx context.Context, req *_nakamaReq.NakamaRegisterToFamily) error {
	ret := n.Called(ctx, req)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, *_nakamaReq.NakamaRegisterToFamily) error); ok {
		r0 = ref(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

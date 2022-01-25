package mocks

import (
	"context"

	_userReq "github.com/madjiebimaa/nakafam/user/delivery/http/requests"
	_userRes "github.com/madjiebimaa/nakafam/user/delivery/http/responses"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUCase struct {
	mock.Mock
}

func (u *UserUCase) Register(ctx context.Context, req *_userReq.UserRegisterOrLogin) error {
	ret := u.Called(ctx, req)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, *_userReq.UserRegisterOrLogin) error); ok {
		r0 = ref(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (u *UserUCase) Login(ctx context.Context, req *_userReq.UserRegisterOrLogin) (_userRes.UserBase, error) {
	ret := u.Called(ctx, req)

	var r0 _userRes.UserBase
	if ref, ok := ret.Get(0).(func(context.Context, *_userReq.UserRegisterOrLogin) _userRes.UserBase); ok {
		r0 = ref(ctx, req)
	} else {
		r0 = ret.Get(0).(_userRes.UserBase)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context, *_userReq.UserRegisterOrLogin) error); ok {
		r1 = ref(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (u *UserUCase) UpgradeRole(ctx context.Context, id primitive.ObjectID) (string, error) {
	ret := u.Called(ctx, id)

	var r0 string
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) string); ok {
		r0 = ref(ctx, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = ref(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (u *UserUCase) ToLeaderRole(ctx context.Context, token string) error {
	ret := u.Called(ctx, token)

	var r0 error
	if ref, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = ref(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (u *UserUCase) Me(ctx context.Context, id primitive.ObjectID) (_userRes.UserBase, error) {
	ret := u.Called(ctx, id)

	var r0 _userRes.UserBase
	if ref, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) _userRes.UserBase); ok {
		r0 = ref(ctx, id)
	} else {
		r0 = ret.Get(0).(_userRes.UserBase)
	}

	var r1 error
	if ref, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = ref(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

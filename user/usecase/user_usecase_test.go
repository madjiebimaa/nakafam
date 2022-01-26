package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/madjiebimaa/nakafam/constant"
	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/domain/fakes"
	"github.com/madjiebimaa/nakafam/domain/mocks"
	"github.com/madjiebimaa/nakafam/user/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRegister(t *testing.T) {
	fakeReq := fakes.FakeUserRegisterOrLoginRequest()
	fakeUser := fakes.FakeUser()

	userRepo := new(mocks.UserRepository)
	tokenRepo := new(mocks.TokenRepository)
	userUCase := usecase.NewUserUseCase(userRepo, tokenRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		userRepo.On("GetByEmail", mock.Anything, fakeReq.Email).Return(domain.User{}, nil).Once()
		userRepo.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUCase.Register(context.TODO(), &fakeReq)
		assert.NoError(t, err)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail check email in repository", func(t *testing.T) {
		userRepo.On("GetByEmail", mock.Anything, fakeReq.Email).Return(domain.User{}, domain.ErrInternalServerError).Once()

		err := userUCase.Register(context.TODO(), &fakeReq)
		assert.Error(t, err)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail email already exist", func(t *testing.T) {
		userRepo.On("GetByEmail", mock.Anything, fakeReq.Email).Return(fakeUser, nil).Once()

		err := userUCase.Register(context.TODO(), &fakeReq)
		assert.Error(t, err)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail register user into repository", func(t *testing.T) {
		userRepo.On("GetByEmail", mock.Anything, fakeReq.Email).Return(domain.User{}, nil).Once()
		userRepo.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(domain.ErrInternalServerError).Once()

		err := userUCase.Register(context.TODO(), &fakeReq)
		assert.Error(t, err)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

}

func TestLogin(t *testing.T) {
	fakeReq := fakes.FakeUserRegisterOrLoginRequest()
	fakeUser := fakes.FakeUser()

	userRepo := new(mocks.UserRepository)
	tokenRepo := new(mocks.TokenRepository)
	userUCase := usecase.NewUserUseCase(userRepo, tokenRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		userRepo.On("GetByEmail", mock.Anything, fakeReq.Email).Return(fakeUser, nil).Once()

		user, err := userUCase.Login(context.TODO(), &fakeReq)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail check email in repository", func(t *testing.T) {
		userRepo.On("GetByEmail", mock.Anything, fakeReq.Email).Return(domain.User{}, domain.ErrInternalServerError).Once()

		user, err := userUCase.Login(context.TODO(), &fakeReq)
		assert.Error(t, err)
		assert.Equal(t, user.ID, primitive.NilObjectID)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail user with that email is not exist", func(t *testing.T) {
		userRepo.On("GetByEmail", mock.Anything, fakeReq.Email).Return(domain.User{}, nil).Once()

		user, err := userUCase.Login(context.TODO(), &fakeReq)
		assert.Error(t, err)
		assert.Equal(t, user.ID, primitive.NilObjectID)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail invalid credentials", func(t *testing.T) {
		fakeReq.Password = "test"
		userRepo.On("GetByEmail", mock.Anything, fakeReq.Email).Return(fakeUser, nil).Once()

		user, err := userUCase.Login(context.TODO(), &fakeReq)
		assert.Error(t, err)
		assert.Equal(t, user.ID, primitive.NilObjectID)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})
}

func TestUpgradeRole(t *testing.T) {
	fakeUser := fakes.FakeUser()
	id := fakeUser.ID

	userRepo := new(mocks.UserRepository)
	tokenRepo := new(mocks.TokenRepository)
	userUCase := usecase.NewUserUseCase(userRepo, tokenRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		userRepo.On("GetByID", mock.Anything, id).Return(fakeUser, nil).Once()
		tokenRepo.On("Set", mock.Anything, mock.AnythingOfType("string"), fakeUser.ID, mock.AnythingOfType("time.Duration")).Return(nil).Once()

		res, err := userUCase.UpgradeRole(context.TODO(), id)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail find user in repository", func(t *testing.T) {
		userRepo.On("GetByID", mock.Anything, id).Return(domain.User{}, domain.ErrInternalServerError).Once()

		res, err := userUCase.UpgradeRole(context.TODO(), id)
		assert.Error(t, err)
		assert.Equal(t, res.URL, "")
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail user not found by id", func(t *testing.T) {
		userRepo.On("GetByID", mock.Anything, id).Return(domain.User{}, nil).Once()

		res, err := userUCase.UpgradeRole(context.TODO(), id)
		assert.Error(t, err)
		assert.Equal(t, res.URL, "")
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail store user id in repository", func(t *testing.T) {
		userRepo.On("GetByID", mock.Anything, id).Return(fakeUser, nil).Once()
		tokenRepo.On("Set", mock.Anything, mock.AnythingOfType("string"), fakeUser.ID, mock.AnythingOfType("time.Duration")).Return(domain.ErrInternalServerError).Once()

		res, err := userUCase.UpgradeRole(context.TODO(), id)
		assert.Error(t, err)
		assert.Equal(t, res.URL, "")
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})
}

func TestToLeaderRole(t *testing.T) {
	val := "61f0ed7d6af9e743099437e5"
	id, err := primitive.ObjectIDFromHex(val)
	assert.NoError(t, err)

	fakeReq := fakes.FakeUserToLeaderRoleRequest()
	fakeUser := fakes.FakeUser()
	fakeUser.ID = id

	key := constant.TOKEN_REGISTER_LEADER_PREFIX + fakeReq.Token

	userRepo := new(mocks.UserRepository)
	tokenRepo := new(mocks.TokenRepository)
	userUCase := usecase.NewUserUseCase(userRepo, tokenRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		tokenRepo.On("Get", mock.Anything, key).Return(val, nil).Once()
		userRepo.On("GetByID", mock.Anything, fakeUser.ID).Return(fakeUser, nil).Once()
		userRepo.On("ToLeaderRole", mock.Anything, fakeUser.ID).Return(nil).Once()

		err = userUCase.ToLeaderRole(context.TODO(), &fakeReq)
		assert.NoError(t, err)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail not found value in repository", func(t *testing.T) {
		tokenRepo.On("Get", mock.Anything, key).Return("", redis.Nil).Once()

		err = userUCase.ToLeaderRole(context.TODO(), &fakeReq)
		assert.Error(t, err)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail find value in repository", func(t *testing.T) {
		tokenRepo.On("Get", mock.Anything, key).Return("", domain.ErrInternalServerError).Once()

		err = userUCase.ToLeaderRole(context.TODO(), &fakeReq)
		assert.Error(t, err)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail value is not object id", func(t *testing.T) {
		tokenRepo.On("Get", mock.Anything, key).Return("test", nil).Once()

		err = userUCase.ToLeaderRole(context.TODO(), &fakeReq)
		assert.Error(t, err)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail find user in repository", func(t *testing.T) {
		tokenRepo.On("Get", mock.Anything, key).Return(val, nil).Once()
		userRepo.On("GetByID", mock.Anything, fakeUser.ID).Return(domain.User{}, domain.ErrInternalServerError).Once()

		err = userUCase.ToLeaderRole(context.TODO(), &fakeReq)
		assert.Error(t, err)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail user not found", func(t *testing.T) {
		tokenRepo.On("Get", mock.Anything, key).Return(val, nil).Once()
		userRepo.On("GetByID", mock.Anything, fakeUser.ID).Return(domain.User{}, nil).Once()

		err = userUCase.ToLeaderRole(context.TODO(), &fakeReq)
		assert.Error(t, err)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail invalid credentials", func(t *testing.T) {
		fakeReq.Password = "test"
		tokenRepo.On("Get", mock.Anything, key).Return(val, nil).Once()
		userRepo.On("GetByID", mock.Anything, fakeUser.ID).Return(fakeUser, nil).Once()

		err = userUCase.ToLeaderRole(context.TODO(), &fakeReq)
		assert.Error(t, err)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})
}

func TestMe(t *testing.T) {
	fakeUser := fakes.FakeUser()
	id := fakeUser.ID

	userRepo := new(mocks.UserRepository)
	tokenRepo := new(mocks.TokenRepository)
	userUCase := usecase.NewUserUseCase(userRepo, tokenRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		userRepo.On("GetByID", mock.Anything, id).Return(fakeUser, nil).Once()

		res, err := userUCase.Me(context.TODO(), id)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail check id in repository", func(t *testing.T) {
		userRepo.On("GetByID", mock.Anything, id).Return(domain.User{}, domain.ErrInternalServerError).Once()

		res, err := userUCase.Me(context.TODO(), id)
		assert.Error(t, err)
		assert.Equal(t, res.ID, primitive.NilObjectID)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})

	t.Run("fail user not found", func(t *testing.T) {
		userRepo.On("GetByID", mock.Anything, id).Return(domain.User{}, nil).Once()

		res, err := userUCase.Me(context.TODO(), id)
		assert.Error(t, err)
		assert.Equal(t, res.ID, primitive.NilObjectID)
		userRepo.AssertExpectations(t)
		tokenRepo.AssertExpectations(t)
	})
}

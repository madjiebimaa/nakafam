package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/domain/fakes"
	"github.com/madjiebimaa/nakafam/domain/mocks"
	"github.com/madjiebimaa/nakafam/nakama/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {
	fakeReq := fakes.FakeNakamaCreateRequest()
	fakeUser := fakes.FakeUser()

	nakamaRepo := new(mocks.NakamaRepository)
	userRepo := new(mocks.UserRepository)
	familyRepo := new(mocks.FamilyRepository)
	nakamaUCase := usecase.NewNakamaUseCase(nakamaRepo, userRepo, familyRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		userRepo.On("GetByID", mock.Anything, fakeReq.UserID).Return(fakeUser, nil).Once()
		nakamaRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Nakama")).Return(nil).Once()

		err := nakamaUCase.Create(context.TODO(), &fakeReq)
		assert.NoError(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail find user in repository", func(t *testing.T) {
		userRepo.On("GetByID", mock.Anything, fakeReq.UserID).Return(domain.User{}, domain.ErrInternalServerError).Once()

		err := nakamaUCase.Create(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail user not found", func(t *testing.T) {
		userRepo.On("GetByID", mock.Anything, fakeReq.UserID).Return(domain.User{}, nil).Once()

		err := nakamaUCase.Create(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail create user in repository", func(t *testing.T) {
		userRepo.On("GetByID", mock.Anything, fakeReq.UserID).Return(fakeUser, nil).Once()
		nakamaRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Nakama")).Return(domain.ErrInternalServerError).Once()

		err := nakamaUCase.Create(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	fakeReq := fakes.FakeNakamaUpdateRequest()
	fakeNakama := fakes.FakeNakama()
	userID := fakeReq.UserID

	nakamaRepo := new(mocks.NakamaRepository)
	userRepo := new(mocks.UserRepository)
	familyRepo := new(mocks.FamilyRepository)
	nakamaUCase := usecase.NewNakamaUseCase(nakamaRepo, userRepo, familyRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(fakeNakama, nil).Once()
		nakamaRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Nakama")).Return(nil).Once()

		err := nakamaUCase.Update(context.TODO(), &fakeReq)
		assert.NoError(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail find nakama in repository", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(domain.Nakama{}, domain.ErrInternalServerError).Once()

		err := nakamaUCase.Update(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail nakama not found", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(domain.Nakama{}, nil).Once()

		err := nakamaUCase.Update(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail user is not a author of the nakama", func(t *testing.T) {
		fakeReq.UserID = primitive.NewObjectID()
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(fakeNakama, nil).Once()

		err := nakamaUCase.Update(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail update nakama in repository", func(t *testing.T) {
		fakeReq.UserID = userID
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(fakeNakama, nil).Once()
		nakamaRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Nakama")).Return(domain.ErrInternalServerError).Once()

		err := nakamaUCase.Update(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	fakeReq := fakes.FakeNakamaDeleteRequest()
	fakeNakama := fakes.FakeNakama()
	userID := fakeReq.UserID

	nakamaRepo := new(mocks.NakamaRepository)
	userRepo := new(mocks.UserRepository)
	familyRepo := new(mocks.FamilyRepository)
	nakamaUCase := usecase.NewNakamaUseCase(nakamaRepo, userRepo, familyRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(fakeNakama, nil).Once()
		nakamaRepo.On("Delete", mock.Anything, fakeNakama.ID).Return(nil).Once()

		err := nakamaUCase.Delete(context.TODO(), &fakeReq)
		assert.NoError(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail find nakama in repository", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(domain.Nakama{}, domain.ErrInternalServerError).Once()

		err := nakamaUCase.Delete(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail nakama not found", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(domain.Nakama{}, nil).Once()

		err := nakamaUCase.Delete(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail user is not author if the nakama", func(t *testing.T) {
		fakeReq.UserID = primitive.NewObjectID()
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(fakeNakama, nil).Once()

		err := nakamaUCase.Delete(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail user is not author if the nakama", func(t *testing.T) {
		fakeReq.UserID = userID
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(fakeNakama, nil).Once()
		nakamaRepo.On("Delete", mock.Anything, fakeNakama.ID).Return(domain.ErrInternalServerError).Once()

		err := nakamaUCase.Delete(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	fakeNakama := fakes.FakeNakama()

	nakamaRepo := new(mocks.NakamaRepository)
	userRepo := new(mocks.UserRepository)
	familyRepo := new(mocks.FamilyRepository)
	nakamaUCase := usecase.NewNakamaUseCase(nakamaRepo, userRepo, familyRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeNakama.ID).Return(fakeNakama, nil).Once()

		nakama, err := nakamaUCase.GetByID(context.TODO(), fakeNakama.ID)
		assert.NoError(t, err)
		assert.NotNil(t, nakama)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail find nakama in repository", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeNakama.ID).Return(domain.Nakama{}, domain.ErrInternalServerError).Once()

		nakama, err := nakamaUCase.GetByID(context.TODO(), fakeNakama.ID)
		assert.Error(t, err)
		assert.Equal(t, nakama.ID, primitive.NilObjectID)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail nakama not found", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeNakama.ID).Return(domain.Nakama{}, nil).Once()

		nakama, err := nakamaUCase.GetByID(context.TODO(), fakeNakama.ID)
		assert.Error(t, err)
		assert.Equal(t, nakama.ID, primitive.NilObjectID)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	fakeNakamas := []domain.Nakama{fakes.FakeNakama(), fakes.FakeNakama()}

	nakamaRepo := new(mocks.NakamaRepository)
	userRepo := new(mocks.UserRepository)
	familyRepo := new(mocks.FamilyRepository)
	nakamaUCase := usecase.NewNakamaUseCase(nakamaRepo, userRepo, familyRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		nakamaRepo.On("GetAll", mock.Anything).Return(fakeNakamas, nil).Once()

		nakamas, err := nakamaUCase.GetAll(context.TODO())
		assert.NoError(t, err)
		assert.NotNil(t, nakamas)
		assert.Equal(t, 2, len(fakeNakamas))
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail find nakamas in repository", func(t *testing.T) {
		nakamaRepo.On("GetAll", mock.Anything).Return([]domain.Nakama{}, domain.ErrInternalServerError).Once()

		nakamas, err := nakamaUCase.GetAll(context.TODO())
		assert.Error(t, err)
		assert.Equal(t, 0, len(nakamas))
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail nakamas not found", func(t *testing.T) {
		nakamaRepo.On("GetAll", mock.Anything).Return([]domain.Nakama{}, nil).Once()

		nakamas, err := nakamaUCase.GetAll(context.TODO())
		assert.Error(t, err)
		assert.Equal(t, 0, len(nakamas))
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})
}

func TestRegisterToFamily(t *testing.T) {
	fakeReq := fakes.FakeNakamaRegisterFamilyRequest()
	fakeNakama := fakes.FakeNakama()
	fakeFamily := fakes.FakeFamily()
	password := fakeReq.Password

	nakamaRepo := new(mocks.NakamaRepository)
	userRepo := new(mocks.UserRepository)
	familyRepo := new(mocks.FamilyRepository)
	nakamaUCase := usecase.NewNakamaUseCase(nakamaRepo, userRepo, familyRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(fakeNakama, nil).Once()
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(fakeFamily, nil).Once()
		nakamaRepo.On("RegisterToFamily", mock.Anything, fakeNakama.ID, fakeFamily.ID).Return(nil).Once()

		err := nakamaUCase.RegisterToFamily(context.TODO(), &fakeReq)
		assert.NoError(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail find nakama in repository", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(domain.Nakama{}, domain.ErrInternalServerError).Once()

		err := nakamaUCase.RegisterToFamily(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail nakama not found", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(domain.Nakama{}, nil).Once()

		err := nakamaUCase.RegisterToFamily(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail find family in repository", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(fakeNakama, nil).Once()
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(domain.Family{}, domain.ErrInternalServerError).Once()

		err := nakamaUCase.RegisterToFamily(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail family not found", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(fakeNakama, nil).Once()
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(domain.Family{}, nil).Once()

		err := nakamaUCase.RegisterToFamily(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail invalid credentials", func(t *testing.T) {
		fakeReq.Password = "testing"

		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(fakeNakama, nil).Once()
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(fakeFamily, nil).Once()

		err := nakamaUCase.RegisterToFamily(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})

	t.Run("fail register to family in repository", func(t *testing.T) {
		fakeReq.Password = password
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(fakeNakama, nil).Once()
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(fakeFamily, nil).Once()
		nakamaRepo.On("RegisterToFamily", mock.Anything, fakeNakama.ID, fakeFamily.ID).Return(domain.ErrInternalServerError).Once()

		err := nakamaUCase.RegisterToFamily(context.TODO(), &fakeReq)
		assert.Error(t, err)
		nakamaRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
		familyRepo.AssertExpectations(t)
	})
}

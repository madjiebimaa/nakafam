package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/domain/fakes"
	"github.com/madjiebimaa/nakafam/domain/mocks"
	"github.com/madjiebimaa/nakafam/family/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {
	fakeReq := fakes.FakeFamilyCreateRequest()
	fakeNakama := fakes.FakeNakama()

	familyRepo := new(mocks.FamilyRepository)
	nakamaRepo := new(mocks.NakamaRepository)
	familyUCase := usecase.NewFamilyUseCase(familyRepo, nakamaRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(fakeNakama, nil).Once()
		familyRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Family")).Return(nil).Once()

		err := familyUCase.Create(context.TODO(), &fakeReq)
		assert.NoError(t, err)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail find family in repository", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(domain.Nakama{}, domain.ErrInternalServerError).Once()

		err := familyUCase.Create(context.TODO(), &fakeReq)
		assert.Error(t, err)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail family not found", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(domain.Nakama{}, nil).Once()

		err := familyUCase.Create(context.TODO(), &fakeReq)
		assert.Error(t, err)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail create family in repository", func(t *testing.T) {
		nakamaRepo.On("GetByID", mock.Anything, fakeReq.NakamaID).Return(fakeNakama, nil).Once()
		familyRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Family")).Return(domain.ErrInternalServerError).Once()

		err := familyUCase.Create(context.TODO(), &fakeReq)
		assert.Error(t, err)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	fakeReq := fakes.FakeFamilyUpdateRequest()
	nID := fakeReq.NakamaID
	fakeFamily := fakes.FakeFamily()

	familyRepo := new(mocks.FamilyRepository)
	nakamaRepo := new(mocks.NakamaRepository)
	familyUCase := usecase.NewFamilyUseCase(familyRepo, nakamaRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(fakeFamily, nil).Once()
		familyRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Family")).Return(nil).Once()

		err := familyUCase.Update(context.TODO(), &fakeReq)
		assert.NoError(t, err)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail find family in repository", func(t *testing.T) {
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(domain.Family{}, domain.ErrInternalServerError).Once()

		err := familyUCase.Update(context.TODO(), &fakeReq)
		assert.Error(t, err)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail family not found", func(t *testing.T) {
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(domain.Family{}, nil).Once()

		err := familyUCase.Update(context.TODO(), &fakeReq)
		assert.Error(t, err)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail nakama is not a leader of the family", func(t *testing.T) {
		fakeReq.NakamaID = primitive.NewObjectID()
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(fakeFamily, nil).Once()

		err := familyUCase.Update(context.TODO(), &fakeReq)
		assert.Error(t, err)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail update family in repository", func(t *testing.T) {
		fakeReq.NakamaID = nID
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(fakeFamily, nil).Once()
		familyRepo.On("Update", mock.Anything, mock.AnythingOfType("*domain.Family")).Return(domain.ErrInternalServerError).Once()

		err := familyUCase.Update(context.TODO(), &fakeReq)
		assert.Error(t, err)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	fakeReq := fakes.FakeFamilyDeleteRequest()
	nID := fakeReq.NakamaID
	fakeFamily := fakes.FakeFamily()

	familyRepo := new(mocks.FamilyRepository)
	nakamaRepo := new(mocks.NakamaRepository)
	familyUCase := usecase.NewFamilyUseCase(familyRepo, nakamaRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(fakeFamily, nil).Once()
		familyRepo.On("Delete", mock.Anything, fakeFamily.ID).Return(nil).Once()

		err := familyUCase.Delete(context.TODO(), &fakeReq)
		assert.NoError(t, err)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail find family in repository", func(t *testing.T) {
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(domain.Family{}, domain.ErrInternalServerError).Once()

		err := familyUCase.Delete(context.TODO(), &fakeReq)
		assert.Error(t, err)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail family not found", func(t *testing.T) {
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(domain.Family{}, nil).Once()

		err := familyUCase.Delete(context.TODO(), &fakeReq)
		assert.Error(t, err)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail nakama is not a leader of the family", func(t *testing.T) {
		fakeReq.NakamaID = primitive.NewObjectID()
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(fakeFamily, nil).Once()

		err := familyUCase.Delete(context.TODO(), &fakeReq)
		assert.Error(t, err)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail delete family in repository", func(t *testing.T) {
		fakeReq.NakamaID = nID
		familyRepo.On("GetByID", mock.Anything, fakeReq.FamilyID).Return(fakeFamily, nil).Once()
		familyRepo.On("Delete", mock.Anything, fakeFamily.ID).Return(domain.ErrInternalServerError).Once()

		err := familyUCase.Delete(context.TODO(), &fakeReq)
		assert.Error(t, err)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	fakeFamily := fakes.FakeFamily()
	fID := fakeFamily.ID

	familyRepo := new(mocks.FamilyRepository)
	nakamaRepo := new(mocks.NakamaRepository)
	familyUCase := usecase.NewFamilyUseCase(familyRepo, nakamaRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		familyRepo.On("GetByID", mock.Anything, fID).Return(fakeFamily, nil).Once()

		family, err := familyUCase.GetByID(context.TODO(), fID)
		assert.NoError(t, err)
		assert.NotNil(t, family)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail find family in repository", func(t *testing.T) {
		familyRepo.On("GetByID", mock.Anything, fID).Return(domain.Family{}, domain.ErrInternalServerError).Once()

		family, err := familyUCase.GetByID(context.TODO(), fID)
		assert.Error(t, err)
		assert.Equal(t, primitive.NilObjectID, family.ID)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail family not found", func(t *testing.T) {
		familyRepo.On("GetByID", mock.Anything, fID).Return(domain.Family{}, nil).Once()

		family, err := familyUCase.GetByID(context.TODO(), fID)
		assert.Error(t, err)
		assert.Equal(t, primitive.NilObjectID, family.ID)
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	fakeFamilies := []domain.Family{fakes.FakeFamily(), fakes.FakeFamily()}

	familyRepo := new(mocks.FamilyRepository)
	nakamaRepo := new(mocks.NakamaRepository)
	familyUCase := usecase.NewFamilyUseCase(familyRepo, nakamaRepo, 2*time.Second)

	t.Run("success", func(t *testing.T) {
		familyRepo.On("GetAll", mock.Anything).Return(fakeFamilies, nil).Once()

		families, err := familyUCase.GetAll(context.TODO())
		assert.NoError(t, err)
		assert.NotNil(t, families)
		assert.Equal(t, 2, len(families))
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail find families in repository", func(t *testing.T) {
		familyRepo.On("GetAll", mock.Anything).Return([]domain.Family{}, domain.ErrInternalServerError).Once()

		families, err := familyUCase.GetAll(context.TODO())
		assert.Error(t, err)
		assert.Equal(t, 0, len(families))
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})

	t.Run("fail families not found", func(t *testing.T) {
		familyRepo.On("GetAll", mock.Anything).Return([]domain.Family{}, nil).Once()

		families, err := familyUCase.GetAll(context.TODO())
		assert.Error(t, err)
		assert.Equal(t, 0, len(families))
		familyRepo.AssertExpectations(t)
		nakamaRepo.AssertExpectations(t)
	})
}

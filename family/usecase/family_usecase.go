package usecase

import (
	"context"
	"log"
	"time"

	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/family/delivery/http/requests"
	"github.com/madjiebimaa/nakafam/family/delivery/http/responses"
	"github.com/madjiebimaa/nakafam/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type familyUseCase struct {
	familyRepo     domain.FamilyRepository
	nakamaRepo     domain.NakamaRepository
	contextTimeout time.Duration
}

func NewFamilyUseCase(
	familyRepo domain.FamilyRepository,
	nakamaRepo domain.NakamaRepository,
	contextTimeout time.Duration,
) domain.FamilyUseCase {
	return &familyUseCase{
		familyRepo,
		nakamaRepo,
		contextTimeout,
	}
}

func (f *familyUseCase) Create(c context.Context, req *requests.FamilyCreate) error {
	ctx, cancel := context.WithTimeout(c, f.contextTimeout)
	defer cancel()

	nakama, err := f.nakamaRepo.GetByID(ctx, req.NakamaID)
	if err != nil {
		return err
	}

	if nakama.ID == primitive.NilObjectID {
		return domain.ErrNotFound
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	now := time.Now()
	family := domain.Family{
		ID:          primitive.NewObjectID(),
		Name:        req.Name,
		Password:    string(hashedPass),
		FamilyImage: req.FamilyImage,
		Nakamas:     []domain.Nakama{nakama},
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := f.familyRepo.Create(ctx, &family); err != nil {
		return err
	}

	return nil
}

func (f *familyUseCase) Update(c context.Context, req *requests.FamilyUpdate) error {
	ctx, cancel := context.WithTimeout(c, f.contextTimeout)
	defer cancel()

	family, err := f.familyRepo.GetByID(ctx, req.FamilyID)
	if err != nil {
		return nil
	}

	if family.ID == primitive.NilObjectID {
		return domain.ErrNotFound
	}

	if req.Name != "" {
		family.Name = req.Name
	}

	if req.FamilyImage != "" {
		family.FamilyImage = req.FamilyImage
	}

	family.UpdatedAt = time.Now()
	if err := f.familyRepo.Update(ctx, &family); err != nil {
		return err
	}

	return nil
}

func (f *familyUseCase) Delete(c context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, f.contextTimeout)
	defer cancel()

	family, err := f.familyRepo.GetByID(ctx, id)
	if err != nil {
		return nil
	}

	if family.ID == primitive.NilObjectID {
		return domain.ErrNotFound
	}

	if err := f.familyRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (f *familyUseCase) GetByID(c context.Context, id primitive.ObjectID) (responses.FamilyBase, error) {
	ctx, cancel := context.WithTimeout(c, f.contextTimeout)
	defer cancel()

	family, err := f.familyRepo.GetByID(ctx, id)
	if err != nil {
		return responses.FamilyBase{}, err
	}

	if family.ID == primitive.NilObjectID {
		return responses.FamilyBase{}, domain.ErrNotFound
	}

	res := helpers.ToFamilyBase(&family)
	return res, nil
}

func (f *familyUseCase) GetByName(c context.Context, name string) (responses.FamilyBase, error) {
	ctx, cancel := context.WithTimeout(c, f.contextTimeout)
	defer cancel()

	family, err := f.familyRepo.GetByName(ctx, name)
	if err != nil {
		return responses.FamilyBase{}, err
	}

	if family.ID == primitive.NilObjectID {
		return responses.FamilyBase{}, domain.ErrNotFound
	}

	res := helpers.ToFamilyBase(&family)
	return res, nil
}

func (f *familyUseCase) GetAll(c context.Context) ([]responses.FamilyBase, error) {
	ctx, cancel := context.WithTimeout(c, f.contextTimeout)
	defer cancel()

	families, err := f.familyRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if families == nil {
		return nil, domain.ErrNotFound
	}

	res := helpers.ToFamiliesBase(families)
	return res, nil
}

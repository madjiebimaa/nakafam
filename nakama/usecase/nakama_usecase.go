package usecase

import (
	"context"
	"log"
	"time"

	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/helpers"
	"github.com/madjiebimaa/nakafam/nakama/delivery/http/requests"
	"github.com/madjiebimaa/nakafam/nakama/delivery/http/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type nakamaUseCase struct {
	nakamRepo      domain.NakamaRepository
	familyRepo     domain.FamilyRepository
	contextTimeout time.Duration
}

func NewNakamaUseCase() domain.NakamaUseCase {
	return &nakamaUseCase{}
}

func (n *nakamaUseCase) Create(c context.Context, req *requests.NakamaCreate) error {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	now := time.Now()
	nakama := domain.Nakama{
		ID:           primitive.NewObjectID(),
		Name:         req.Name,
		UserName:     req.UserName,
		ProfileImage: req.ProfileImage,
		Description:  req.Description,
		SocialMedia:  (*domain.SocialMedia)(req.SocialMedia),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := n.nakamRepo.Create(ctx, &nakama); err != nil {
		return err
	}

	return nil
}

func (n *nakamaUseCase) Update(c context.Context, req *requests.NakamaUpdate) error {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	now := time.Now()
	nakama := domain.Nakama{
		ID:           req.ID,
		FamilyID:     req.FamilyID,
		Name:         req.Name,
		UserName:     req.UserName,
		ProfileImage: req.ProfileImage,
		Description:  req.Description,
		SocialMedia:  (*domain.SocialMedia)(req.SocialMedia),
		UpdatedAt:    now,
	}

	if err := n.nakamRepo.Update(ctx, &nakama); err != nil {
		return err
	}

	return nil
}

func (n *nakamaUseCase) Delete(c context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	if err := n.nakamRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (n *nakamaUseCase) GetByID(c context.Context, id primitive.ObjectID) (responses.NakamaBase, error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	nakama, err := n.nakamRepo.GetByID(ctx, id)
	if err != nil {
		return responses.NakamaBase{}, err
	}

	res := helpers.ToNakamaBase(&nakama)

	return res, nil
}

func (n *nakamaUseCase) GetByName(c context.Context, name string) (responses.NakamaBase, error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	nakama, err := n.nakamRepo.GetByName(ctx, name)
	if err != nil {
		return responses.NakamaBase{}, err
	}

	res := helpers.ToNakamaBase(&nakama)

	return res, nil
}

func (n *nakamaUseCase) GetByFamilyID(c context.Context, familyID primitive.ObjectID) ([]responses.NakamaBase, error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	nakamas, err := n.nakamRepo.GetByFamilyID(ctx, familyID)
	if err != nil {
		return nil, err
	}

	res := helpers.ToNakamasBase(nakamas)

	return res, nil
}

func (n *nakamaUseCase) RegisterToFamily(c context.Context, req *requests.NakamaRegisterToFamily) error {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	family, err := n.familyRepo.GetByID(ctx, req.FamilyID)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(family.Password), []byte(req.Password)); err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	nakama, err := n.nakamRepo.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if err := n.nakamRepo.RegisterToFamily(ctx, nakama.ID, family.ID); err != nil {
		return err
	}

	return nil
}

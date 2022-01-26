package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/helpers"

	_nakamaReq "github.com/madjiebimaa/nakafam/nakama/delivery/http/requests"
	_nakamaRes "github.com/madjiebimaa/nakafam/nakama/delivery/http/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type nakamaUseCase struct {
	nakamaRepo     domain.NakamaRepository
	userRepo       domain.UserRepository
	familyRepo     domain.FamilyRepository
	contextTimeout time.Duration
}

func NewNakamaUseCase(
	nakamaRepo domain.NakamaRepository,
	userRepo domain.UserRepository,
	familyRepo domain.FamilyRepository,
	contextTimeout time.Duration,
) domain.NakamaUseCase {
	return &nakamaUseCase{
		nakamaRepo,
		userRepo,
		familyRepo,
		contextTimeout,
	}
}

func (n *nakamaUseCase) Create(c context.Context, req *_nakamaReq.NakamaCreate) error {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	user, err := n.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	if user.ID == primitive.NilObjectID {
		return domain.ErrNotFound
	}

	now := time.Now()
	nakama := domain.Nakama{
		ID:           primitive.NewObjectID(),
		UserID:       user.ID,
		Name:         req.Name,
		UserName:     req.UserName,
		ProfileImage: req.ProfileImage,
		Description:  req.Description,
		SocialMedia:  (*domain.SocialMedia)(req.SocialMedia),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := n.nakamaRepo.Create(ctx, &nakama); err != nil {
		return err
	}

	return nil
}

func (n *nakamaUseCase) Update(c context.Context, req *_nakamaReq.NakamaUpdate) error {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	nakama, err := n.nakamaRepo.GetByID(ctx, req.NakamaID)
	if err != nil {
		return err
	}

	if nakama.ID == primitive.NilObjectID {
		return domain.ErrNotFound
	}

	if req.Name != "" {
		nakama.Name = req.Name
	}

	if req.ProfileImage != "" {
		nakama.ProfileImage = req.ProfileImage
	}

	if req.Description != "" {
		nakama.Description = req.Description
	}

	if req.SocialMedia != nil {
		nakama.SocialMedia = (*domain.SocialMedia)(req.SocialMedia)
	}

	if err := helpers.IsAuthorOfNakama(req.UserID, &nakama); err != nil {
		return err
	}

	if err := n.nakamaRepo.Update(ctx, &nakama); err != nil {
		return err
	}

	return nil
}

func (n *nakamaUseCase) Delete(c context.Context, req *_nakamaReq.NakamaDelete) error {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	nakama, err := n.nakamaRepo.GetByID(ctx, req.NakamaID)
	if err != nil {
		return err
	}

	if nakama.ID == primitive.NilObjectID {
		return domain.ErrNotFound
	}

	if err := helpers.IsAuthorOfNakama(req.UserID, &nakama); err != nil {
		return err
	}

	if err := n.nakamaRepo.Delete(ctx, nakama.ID); err != nil {
		return err
	}

	return nil
}

func (n *nakamaUseCase) GetByID(c context.Context, id primitive.ObjectID) (_nakamaRes.NakamaBase, error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	nakama, err := n.nakamaRepo.GetByID(ctx, id)
	if err != nil {
		return _nakamaRes.NakamaBase{}, err
	}

	if nakama.ID == primitive.NilObjectID {
		return _nakamaRes.NakamaBase{}, domain.ErrNotFound
	}

	res := helpers.ToNakamaBase(&nakama)
	return res, nil
}

func (n *nakamaUseCase) GetAll(c context.Context) ([]_nakamaRes.NakamaBase, error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	nakamas, err := n.nakamaRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(nakamas) == 0 {
		return nil, domain.ErrNotFound
	}

	res := helpers.ToNakamasBase(nakamas)
	return res, nil
}

func (n *nakamaUseCase) RegisterToFamily(c context.Context, req *_nakamaReq.NakamaRegisterToFamily) error {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	nakama, err := n.nakamaRepo.GetByID(ctx, req.NakamaID)
	if err != nil {
		return err
	}

	if nakama.ID == primitive.NilObjectID {
		return domain.ErrNotFound
	}

	family, err := n.familyRepo.GetByID(ctx, req.FamilyID)
	if err != nil {
		return err
	}

	if family.ID == primitive.NilObjectID {
		return domain.ErrNotFound
	}

	fmt.Println("password: ", req.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(family.Password), []byte(req.Password)); err != nil {
		return domain.ErrUnAuthorized
	}

	if err := n.nakamaRepo.RegisterToFamily(ctx, nakama.ID, family.ID); err != nil {
		return err
	}

	return nil
}

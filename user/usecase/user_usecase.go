package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/madjiebimaa/nakafam/constant"
	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/helpers"
	"github.com/madjiebimaa/nakafam/user/delivery/http/requests"
	"github.com/madjiebimaa/nakafam/user/delivery/http/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo       domain.UserRepository
	nakamaRepo     domain.NakamaRepository
	tokenRepo      domain.TokenRepository
	contextTimeout time.Duration
}

func NewUserUseCase(
	userRepo domain.UserRepository,
	nakamaRepo domain.NakamaRepository,
	tokenRepo domain.TokenRepository,
	contextTimeout time.Duration,
) domain.UserUseCase {
	return &userUseCase{
		userRepo,
		nakamaRepo,
		tokenRepo,
		contextTimeout,
	}
}

func (u *userUseCase) Register(c context.Context, req *requests.UserRegisterOrLogin) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	isUserExist, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	if isUserExist.ID != primitive.NilObjectID {
		return domain.ErrConflict
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		log.Fatal(err)
		return domain.ErrInternalServerError
	}

	now := time.Now()
	user := domain.User{
		ID:        primitive.NewObjectID(),
		Email:     req.Email,
		Password:  string(hashedPass),
		Role:      "staff",
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := u.userRepo.Register(ctx, &user); err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) Login(c context.Context, req *requests.UserRegisterOrLogin) (responses.UserBase, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return responses.UserBase{}, err
	}

	if user.ID == primitive.NilObjectID {
		return responses.UserBase{}, domain.ErrNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return responses.UserBase{}, domain.ErrUnAuthorized
	}

	res := helpers.ToUserBase(&user)
	return res, nil
}

func (u *userUseCase) UpgradeRole(c context.Context, id primitive.ObjectID) (string, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return "", err
	}

	if user.ID == primitive.NilObjectID {
		return "", domain.ErrNotFound
	}

	token := uuid.NewString()
	key := constant.TOKEN_REGISTER_LEADER_PREFIX + token
	if err := u.tokenRepo.Set(ctx, key, user.ID, 3*24*time.Hour); err != nil {
		return "", err
	}

	url := fmt.Sprintf("http://localhost:3000/api/users/upgrade-role/%s", token)
	return url, nil
}

func (u *userUseCase) ToLeaderRole(c context.Context, token string) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	key := constant.TOKEN_REGISTER_LEADER_PREFIX + token
	val, err := u.tokenRepo.Get(ctx, key)
	if err == redis.Nil {
		return domain.ErrNotFound
	}

	if err != nil {
		return err
	}

	userID, err := primitive.ObjectIDFromHex(val)
	if err != nil {
		return domain.ErrBadParamInput
	}

	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if user.ID == primitive.NilObjectID {
		return domain.ErrNotFound
	}

	if err := u.userRepo.ToLeaderRole(ctx, user.ID); err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) Me(c context.Context, id primitive.ObjectID) (responses.UserBase, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return responses.UserBase{}, err
	}

	if user.ID == primitive.NilObjectID {
		return responses.UserBase{}, domain.ErrNotFound
	}

	res := helpers.ToUserBase(&user)
	return res, nil
}

func (u *userUseCase) CreateNakama(c context.Context, req *requests.UserCreateNakama) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	if user.ID == primitive.NilObjectID {
		return domain.ErrNotFound
	}

	now := time.Now()
	nakama := domain.Nakama{
		ID:           primitive.NewObjectID(),
		UserID:       req.UserID,
		FamilyID:     primitive.NilObjectID,
		Name:         req.Name,
		UserName:     req.UserName,
		ProfileImage: req.ProfileImage,
		Description:  req.Description,
		SocialMedia:  (*domain.SocialMedia)(req.SocialMedia),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := u.nakamaRepo.Create(ctx, &nakama); err != nil {
		return err
	}

	return nil
}

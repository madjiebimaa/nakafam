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
	_userReq "github.com/madjiebimaa/nakafam/user/delivery/http/requests"
	_userRes "github.com/madjiebimaa/nakafam/user/delivery/http/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo       domain.UserRepository
	tokenRepo      domain.TokenRepository
	contextTimeout time.Duration
}

func NewUserUseCase(
	userRepo domain.UserRepository,
	tokenRepo domain.TokenRepository,
	contextTimeout time.Duration,
) domain.UserUseCase {
	return &userUseCase{
		userRepo,
		tokenRepo,
		contextTimeout,
	}
}

func (u *userUseCase) Register(c context.Context, req *_userReq.UserRegisterOrLogin) error {
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

func (u *userUseCase) Login(c context.Context, req *_userReq.UserRegisterOrLogin) (_userRes.UserBase, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return _userRes.UserBase{}, err
	}

	if user.ID == primitive.NilObjectID {
		return _userRes.UserBase{}, domain.ErrNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return _userRes.UserBase{}, domain.ErrUnAuthorized
	}

	res := helpers.ToUserBase(&user)
	return res, nil
}

func (u *userUseCase) UpgradeRole(c context.Context, id primitive.ObjectID) (_userRes.UserUpgradeRole, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return _userRes.UserUpgradeRole{}, err
	}

	if user.ID == primitive.NilObjectID {
		return _userRes.UserUpgradeRole{}, domain.ErrNotFound
	}

	token := uuid.NewString()
	key := constant.TOKEN_REGISTER_LEADER_PREFIX + token
	if err := u.tokenRepo.Set(ctx, key, user.ID, 3*24*time.Hour); err != nil {
		return _userRes.UserUpgradeRole{}, err
	}

	url := fmt.Sprintf("http://localhost:3000/api/users/upgrade-role/%s", token)
	res := _userRes.UserUpgradeRole{
		URL: url,
	}
	return res, nil
}

func (u *userUseCase) ToLeaderRole(c context.Context, req *_userReq.UserToLeaderRole) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	key := constant.TOKEN_REGISTER_LEADER_PREFIX + req.Token
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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return domain.ErrUnAuthorized
	}

	if err := u.userRepo.ToLeaderRole(ctx, user.ID); err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) Me(c context.Context, id primitive.ObjectID) (_userRes.UserBase, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return _userRes.UserBase{}, err
	}

	if user.ID == primitive.NilObjectID {
		return _userRes.UserBase{}, domain.ErrNotFound
	}

	res := helpers.ToUserBase(&user)
	return res, nil
}

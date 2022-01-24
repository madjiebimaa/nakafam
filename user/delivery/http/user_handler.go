package http

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/helpers"
	"github.com/madjiebimaa/nakafam/user/delivery/http/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	userUCase domain.UserUseCase
}

func NewUserHandler(
	userUCase domain.UserUseCase,
) *UserHandler {
	return &UserHandler{
		userUCase,
	}
}

func (u *UserHandler) Register(c *gin.Context) {
	var req requests.UserRegisterOrLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	ctx := c.Request.Context()
	if err := u.userUCase.Register(ctx, &req); err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusNoContent, nil)
}

func (u *UserHandler) Login(c *gin.Context) {
	var req requests.UserRegisterOrLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	ctx := c.Request.Context()
	user, err := u.userUCase.Login(ctx, &req)
	if err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	// set userID to session means user have logged in
	// similar in JWT that create access token
	sess := sessions.Default(c)
	sess.Set("user_id", user.ID)
	if err := sess.Save(); err != nil {
		helpers.FailResponse(c, http.StatusInternalServerError, "session", err)
	}

	helpers.SuccessResponse(c, http.StatusNoContent, nil)
}

func (u *UserHandler) UpgradeRole(c *gin.Context) {
	id, _ := c.Get("user_id")
	userID := id.(primitive.ObjectID)
	ctx := c.Request.Context()
	res, err := u.userUCase.UpgradeRole(ctx, userID)
	if err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusOK, res)
}

func (u *UserHandler) ToLeaderRole(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		helpers.FailResponse(c, http.StatusBadRequest, "token", domain.ErrBadParamInput)
	}

	ctx := c.Request.Context()
	if err := u.userUCase.ToLeaderRole(ctx, token); err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusNoContent, nil)
}

func (u *UserHandler) Me(c *gin.Context) {
	id, _ := c.Get("user_id")
	userID := id.(primitive.ObjectID)
	ctx := c.Request.Context()
	user, err := u.userUCase.Me(ctx, userID)
	if err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusOK, user)
}

func (u *UserHandler) Logout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Clear()
	if err := sess.Save(); err != nil {
		helpers.FailResponse(c, http.StatusUnauthorized, "session", err)
	}

	helpers.SuccessResponse(c, http.StatusNoContent, nil)
}

func (u *UserHandler) CreateNakama(c *gin.Context) {
	var req requests.UserCreateNakama
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	id, _ := c.Get("user_id")
	userID := id.(primitive.ObjectID)
	req.UserID = userID
	ctx := c.Request.Context()
	if err := u.userUCase.CreateNakama(ctx, &req); err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusNoContent, nil)
}

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

type UserDelivery struct {
	userUCase domain.UserUseCase
}

func NewUserDelivery(
	userUCase domain.UserUseCase,
) *UserDelivery {
	return &UserDelivery{
		userUCase,
	}
}

func (u *UserDelivery) Register(c *gin.Context) {
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

func (u *UserDelivery) Login(c *gin.Context) {
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
	sess.Set("userID", user.ID)
	if err := sess.Save(); err != nil {
		helpers.FailResponse(c, http.StatusInternalServerError, "session", err)
	}

	helpers.SuccessResponse(c, http.StatusNoContent, nil)
}

func (u *UserDelivery) RegisterAsLeader(c *gin.Context) {

}

func (u *UserDelivery) Me(c *gin.Context) {
	sess := sessions.Default(c)
	val := sess.Get("userID")
	if val == nil {
		helpers.FailResponse(c, http.StatusUnauthorized, "session", domain.ErrUnAuthorized)
	}

	userID := val.(primitive.ObjectID)
	ctx := c.Request.Context()
	user, err := u.userUCase.Me(ctx, userID)
	if err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusOK, user)
}

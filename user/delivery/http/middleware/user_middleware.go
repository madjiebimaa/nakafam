package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userMiddleware struct{}

func NewUserMiddleware() *userMiddleware {
	return &userMiddleware{}
}

func (u *userMiddleware) IsAuth(c *gin.Context) {
	sess := sessions.Default(c)
	sessionID := sess.Get("user_id")
	if sessionID == nil {
		helpers.FailResponse(c, http.StatusUnauthorized, "session", domain.ErrUnAuthorized)
	}

	userID := sessionID.(primitive.ObjectID)
	c.Set("user_id", userID)
	c.Next()
}

func (u *userMiddleware) IsStaff(c *gin.Context) {
	sess := sessions.Default(c)
	sessionID := sess.Get("user_role")
	if sessionID == nil {
		helpers.FailResponse(c, http.StatusUnauthorized, "session", domain.ErrUnAuthorized)
	}

	role := sessionID.(string)
	if role != "staff" {
		helpers.FailResponse(c, http.StatusUnauthorized, "session", domain.ErrUnAuthorized)
	}

	c.Next()
}

func (u *userMiddleware) IsLeader(c *gin.Context) {
	sess := sessions.Default(c)
	sessionID := sess.Get("user_role")
	if sessionID == nil {
		helpers.FailResponse(c, http.StatusUnauthorized, "session", domain.ErrUnAuthorized)
	}

	role := sessionID.(string)
	if role != "leader" {
		helpers.FailResponse(c, http.StatusUnauthorized, "session", domain.ErrUnAuthorized)
	}

	c.Next()
}

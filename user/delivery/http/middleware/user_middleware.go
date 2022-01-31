package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/helpers"
)

type userMiddleware struct{}

func NewUserMiddleware() *userMiddleware {
	return &userMiddleware{}
}

func (u *userMiddleware) IsAuth(c *gin.Context) {
	sess := sessions.Default(c)
	uID := sess.Get("userID")
	if uID == nil {
		helpers.FailResponse(c, http.StatusUnauthorized, "session", domain.ErrUnAuthorized)
	}
	userID := uID.(string)

	uRole := sess.Get("userRole")
	if uRole == nil {
		helpers.FailResponse(c, http.StatusUnauthorized, "session", domain.ErrUnAuthorized)
	}
	userRole := uRole.(string)

	nID := sess.Get("nakamaID")
	if nID == nil {
		helpers.FailResponse(c, http.StatusUnauthorized, "session", domain.ErrUnAuthorized)
	}
	nakamaID := nID.(string)

	c.Set("userID", userID)
	c.Set("userRole", userRole)
	c.Set("nakamaID", nakamaID)
	c.Next()
}

func (u *userMiddleware) IsLeader(c *gin.Context) {
	uRole, _ := c.Get("userRole")
	if uRole == nil {
		helpers.FailResponse(c, http.StatusUnauthorized, "session", domain.ErrUnAuthorized)
	}

	userRole := uRole.(string)
	if userRole != "leader" {
		helpers.FailResponse(c, http.StatusUnauthorized, "session", domain.ErrUnAuthorized)
	}

	c.Next()
}

package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sessions.Default(c)
		sessionID := sess.Get("user_id")
		if sessionID == nil {
			helpers.FailResponse(c, http.StatusUnauthorized, "session", domain.ErrUnAuthorized)
		}

		userID := sessionID.(primitive.ObjectID)
		c.Set("user_id", userID)
		c.Next()
	}
}

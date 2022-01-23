package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/nakafam/domain"
	"github.com/sirupsen/logrus"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type BaseResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   *FieldError `json:"error"`
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrBadParamInput:
		return http.StatusBadRequest
	case domain.ErrUnAuthorized:
		return http.StatusUnauthorized
	case domain.ErrRedirect:
		return http.StatusTemporaryRedirect
	case domain.ErrExpired:
		return http.StatusGone
	default:
		return http.StatusInternalServerError
	}
}

func SuccessResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, BaseResponse{
		Success: true,
		Data:    data,
		Error:   nil,
	})
}

func FailResponse(c *gin.Context, status int, errField string, errMessage error) {
	c.JSON(status, BaseResponse{
		Success: false,
		Data:    nil,
		Error: &FieldError{
			Field:   errField,
			Message: errMessage.Error(),
		},
	})
}

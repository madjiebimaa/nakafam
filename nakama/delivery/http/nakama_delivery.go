package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/helpers"
	"github.com/madjiebimaa/nakafam/nakama/delivery/http/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NakamaDelivery struct {
	nakamaUCase domain.NakamaUseCase
}

func NewNakamaDelivery(nakamaUCase domain.NakamaUseCase) *NakamaDelivery {
	return &NakamaDelivery{nakamaUCase}
}

func (n *NakamaDelivery) Create(c *gin.Context) {
	var req requests.NakamaCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	id, _ := c.Get("user_id")
	userID := id.(primitive.ObjectID)
	req.UserID = userID
	ctx := c.Request.Context()
	if err := n.nakamaUCase.Create(ctx, &req); err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusNoContent, nil)
}

func (n *NakamaDelivery) Update(c *gin.Context) {
	var req requests.NakamaUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	// TODO: not implemented

	helpers.SuccessResponse(c, http.StatusNoContent, nil)
}

func (n *NakamaDelivery) Delete(c *gin.Context) {
	// TODO: most likely not have to implemented
}

func (n *NakamaDelivery) GetByID(c *gin.Context) {
	id := c.Param("nakama_id")
	if id == "" {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	nakamaID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	ctx := c.Request.Context()
	res, err := n.nakamaUCase.GetByID(ctx, nakamaID)
	if err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusOK, res)
}

func (n *NakamaDelivery) GetByName(c *gin.Context) {
	// TODO: most likely not have to implemented

}

func (n *NakamaDelivery) GetByFamilyID(c *gin.Context) {
	id := c.Param("nakama_id")
	if id == "" {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	familyID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	ctx := c.Request.Context()
	res, err := n.nakamaUCase.GetByFamilyID(ctx, familyID)
	if err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusOK, res)
}

func (n *NakamaDelivery) RegisterToFamily(c *gin.Context) {
	var req requests.NakamaRegisterToFamily
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	id := c.Param("family_id")
	if id == "" {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	familyID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}
	req.FamilyID = familyID

	ctx := c.Request.Context()
	if err := n.nakamaUCase.RegisterToFamily(ctx, &req); err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusNoContent, nil)
}

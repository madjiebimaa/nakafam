package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/helpers"
	_nakamaReq "github.com/madjiebimaa/nakafam/nakama/delivery/http/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NakamaDelivery struct {
	nakamaUCase domain.NakamaUseCase
}

func NewNakamaDelivery(nakamaUCase domain.NakamaUseCase) *NakamaDelivery {
	return &NakamaDelivery{nakamaUCase}
}

func (n *NakamaDelivery) Create(c *gin.Context) {
	var req _nakamaReq.NakamaCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	val, _ := c.Get("userID")
	if val == nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	userID := val.(primitive.ObjectID)
	req.UserID = userID
	ctx := c.Request.Context()
	if err := n.nakamaUCase.Create(ctx, &req); err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusCreated, nil)
}

func (n *NakamaDelivery) Update(c *gin.Context) {
	var req _nakamaReq.NakamaUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	id := c.Param("nakamaID")
	if id == "" {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	nakamaID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	req.NakamaID = nakamaID
	ctx := c.Request.Context()
	if err := n.nakamaUCase.Update(ctx, &req); err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusNoContent, nil)
}

func (n *NakamaDelivery) Delete(c *gin.Context) {
	id := c.Param("nakamaID")
	if id == "" {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	nakamaID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	uID, _ := c.Get("userID")
	if uID == nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	userID := uID.(primitive.ObjectID)
	req := _nakamaReq.NakamaDelete{
		NakamaID: nakamaID,
		UserID:   userID,
	}

	ctx := c.Request.Context()
	if err := n.nakamaUCase.Delete(ctx, &req); err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusNoContent, nil)
}

func (n *NakamaDelivery) GetByID(c *gin.Context) {
	id := c.Param("nakamaID")
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

func (n *NakamaDelivery) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := n.nakamaUCase.GetAll(ctx)
	if err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusOK, res)
}

func (n *NakamaDelivery) RegisterToFamily(c *gin.Context) {
	var req _nakamaReq.NakamaRegisterToFamily
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	nID := c.Param("nakamaID")
	if nID == "" {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}
	nakamaID, err := primitive.ObjectIDFromHex(nID)
	if err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}
	req.NakamaID = nakamaID

	fID := c.Param("familyID")
	if fID == "" {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}
	familyID, err := primitive.ObjectIDFromHex(fID)
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

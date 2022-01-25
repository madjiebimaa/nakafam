package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/nakafam/domain"
	_familyReq "github.com/madjiebimaa/nakafam/family/delivery/http/requests"
	"github.com/madjiebimaa/nakafam/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FamilyDelivery struct {
	familyUCase domain.FamilyUseCase
}

func NewFamilyDelivery(familyUCase domain.FamilyUseCase) *FamilyDelivery {
	return &FamilyDelivery{familyUCase}
}

func (f *FamilyDelivery) Create(c *gin.Context) {
	var req _familyReq.FamilyCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	// TODO: how to get this kinda things?
	val, _ := c.Get("nakamaID")
	if val == "" {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	nakamaID := val.(primitive.ObjectID)
	req.NakamaID = nakamaID
	ctx := c.Request.Context()
	if err := f.familyUCase.Create(ctx, &req); err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusCreated, nil)
}

func (f *FamilyDelivery) Update(c *gin.Context) {
	var req _familyReq.FamilyUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	val := c.Param("familyID")
	if val == "" {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	familyID, err := primitive.ObjectIDFromHex(val)
	if err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	// TODO: how to get Nakama ID

	req.FamilyID = familyID
	ctx := c.Request.Context()
	if err := f.familyUCase.Update(ctx, &req); err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusNoContent, nil)
}

func (f *FamilyDelivery) Delete(c *gin.Context) {
	id := c.Param("familyID")
	if id == "" {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	familyID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	// TODO: how to get Nakama ID
	req := _familyReq.FamilyDelete{
		FamilyID: familyID,
	}

	ctx := c.Request.Context()
	if err := f.familyUCase.Delete(ctx, &req); err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusNoContent, nil)
}

func (f *FamilyDelivery) GetByID(c *gin.Context) {
	id := c.Param("familyID")
	if id == "" {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	familyID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		helpers.FailResponse(c, http.StatusBadRequest, "input value", domain.ErrBadParamInput)
	}

	ctx := c.Request.Context()
	res, err := f.familyUCase.GetByID(ctx, familyID)
	if err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusOK, res)
}

func (f *FamilyDelivery) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := f.familyUCase.GetAll(ctx)
	if err != nil {
		helpers.FailResponse(c, helpers.GetStatusCode(err), "service", err)
	}

	helpers.SuccessResponse(c, http.StatusOK, res)
}

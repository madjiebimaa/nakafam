package http

import (
	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/nakafam/domain"
)

type FamilyDelivery struct {
	familyUCase domain.FamilyUseCase
}

func NewFamilyDelivery(familyUCase domain.FamilyUseCase) *FamilyDelivery {
	return &FamilyDelivery{familyUCase}
}

func (f *FamilyDelivery) GetByID(c *gin.Context) {}

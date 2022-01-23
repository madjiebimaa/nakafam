package http

import "github.com/madjiebimaa/nakafam/domain"

type FamilyDelivery struct {
	familyUCase domain.FamilyUseCase
}

func NewFamilyDelivery(familyUCase domain.FamilyUseCase) *FamilyDelivery {
	return &FamilyDelivery{familyUCase}
}

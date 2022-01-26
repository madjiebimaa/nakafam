package fakes

import (
	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/family/delivery/http/requests"
)

func FakeFamily() domain.Family {
	return domain.Family{
		ID:          familyID,
		LeaderID:    nakamaID,
		Name:        familyName,
		Password:    string(hashedPass),
		FamilyImage: familyImage,
		Nakamas:     []domain.Nakama{FakeNakama()},
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func FakeFamilyCreateRequest() requests.FamilyCreate {
	return requests.FamilyCreate{
		NakamaID:    nakamaID,
		Name:        familyName,
		Password:    password,
		FamilyImage: familyImage,
	}
}

func FakeFamilyUpdateRequest() requests.FamilyUpdate {
	return requests.FamilyUpdate{
		FamilyID:    familyID,
		NakamaID:    nakamaID,
		Name:        familyName,
		FamilyImage: familyImage,
	}
}

func FakeFamilyDeleteRequest() requests.FamilyDelete {
	return requests.FamilyDelete{
		FamilyID: familyID,
		NakamaID: nakamaID,
	}
}

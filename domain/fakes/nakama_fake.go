package fakes

import (
	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/nakama/delivery/http/requests"
	"github.com/madjiebimaa/nakafam/nakama/delivery/http/responses"
)

func FakeNakama() domain.Nakama {
	return domain.Nakama{
		ID:           nakamaID,
		UserID:       userID,
		Name:         nakamaName,
		UserName:     nakamaUserName,
		ProfileImage: nakamaProfileImage,
		Description:  description,
		SocialMedia:  &socialMedia,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}

func FakeNakamaCreateRequest() requests.NakamaCreate {
	return requests.NakamaCreate{
		UserID:       userID,
		Name:         nakamaName,
		UserName:     nakamaUserName,
		ProfileImage: nakamaProfileImage,
		Description:  description,
		SocialMedia:  (*responses.SocialMediaBase)(&socialMedia),
	}
}

func FakeNakamaUpdateRequest() requests.NakamaUpdate {
	return requests.NakamaUpdate{
		NakamaID:     nakamaID,
		UserID:       userID,
		Name:         nakamaName,
		ProfileImage: nakamaProfileImage,
		Description:  description,
		SocialMedia:  (*responses.SocialMediaBase)(&socialMedia),
	}
}

func FakeNakamaDeleteRequest() requests.NakamaDelete {
	return requests.NakamaDelete{
		NakamaID: nakamaID,
		UserID:   userID,
	}
}

func FakeNakamaRegisterFamilyRequest() requests.NakamaRegisterToFamily {
	return requests.NakamaRegisterToFamily{
		NakamaID: nakamaID,
		FamilyID: familyID,
		Password: password,
	}
}

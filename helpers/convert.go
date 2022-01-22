package helpers

import (
	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/nakama/delivery/http/responses"
)

func ToNakamaBase(nakama *domain.Nakama) responses.NakamaBase {
	return responses.NakamaBase{
		ID:           nakama.ID,
		FamilyID:     nakama.FamilyID,
		Name:         nakama.Name,
		UserName:     nakama.UserName,
		ProfileImage: nakama.ProfileImage,
		Description:  nakama.Description,
		SocialMedia:  (*responses.SocialMediaBase)(nakama.SocialMedia),
		CreatedAt:    nakama.CreatedAt,
		UpdatedAt:    nakama.UpdatedAt,
	}
}

func ToNakamasBase(nakamas []domain.Nakama) []responses.NakamaBase {
	var nakamasBase []responses.NakamaBase
	for _, nakama := range nakamas {
		nakamasBase = append(nakamasBase, ToNakamaBase(&nakama))
	}

	return nakamasBase
}

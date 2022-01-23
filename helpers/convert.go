package helpers

import (
	"github.com/madjiebimaa/nakafam/domain"
	familyRes "github.com/madjiebimaa/nakafam/family/delivery/http/responses"
	nakamaRes "github.com/madjiebimaa/nakafam/nakama/delivery/http/responses"
	userRes "github.com/madjiebimaa/nakafam/user/delivery/http/responses"
)

func ToNakamaBase(nakama *domain.Nakama) nakamaRes.NakamaBase {
	return nakamaRes.NakamaBase{
		ID:           nakama.ID,
		FamilyID:     nakama.FamilyID,
		Name:         nakama.Name,
		UserName:     nakama.UserName,
		ProfileImage: nakama.ProfileImage,
		Description:  nakama.Description,
		SocialMedia:  (*nakamaRes.SocialMediaBase)(nakama.SocialMedia),
		CreatedAt:    nakama.CreatedAt,
		UpdatedAt:    nakama.UpdatedAt,
	}
}

func ToNakamasBase(nakamas []domain.Nakama) []nakamaRes.NakamaBase {
	var nakamasBase []nakamaRes.NakamaBase
	for _, nakama := range nakamas {
		nakamasBase = append(nakamasBase, ToNakamaBase(&nakama))
	}

	return nakamasBase
}

func ToFamilyBase(family *domain.Family) familyRes.FamilyBase {
	return familyRes.FamilyBase{
		ID:          family.ID,
		Name:        family.Name,
		Password:    family.Password,
		FamilyImage: family.FamilyImage,
		Nakamas:     ToNakamasBase(family.Nakamas),
		CreatedAt:   family.CreatedAt,
		UpdatedAt:   family.UpdatedAt,
	}
}

func ToFamiliesBase(families []domain.Family) []familyRes.FamilyBase {
	var familiesBase []familyRes.FamilyBase
	for _, family := range families {
		familiesBase = append(familiesBase, ToFamilyBase(&family))
	}

	return familiesBase
}

func ToUserBase(user *domain.User) userRes.UserBase {
	return userRes.UserBase{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

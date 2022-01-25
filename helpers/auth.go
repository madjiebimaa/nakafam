package helpers

import (
	"github.com/madjiebimaa/nakafam/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsAuthorOfNakama(userID primitive.ObjectID, nakama *domain.Nakama) error {
	if userID != nakama.UserID {
		return domain.ErrUnAuthorized
	}

	return nil
}

func IsLeaderOfFamily(nakamaID primitive.ObjectID, family *domain.Family) error {
	if nakamaID != family.LeaderID {
		return domain.ErrUnAuthorized
	}

	return nil
}

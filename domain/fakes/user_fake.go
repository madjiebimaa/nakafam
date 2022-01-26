package fakes

import (
	"github.com/madjiebimaa/nakafam/domain"
	"github.com/madjiebimaa/nakafam/user/delivery/http/requests"
)

func FakeUser() domain.User {
	return domain.User{
		ID:        userID,
		Email:     email,
		Password:  string(hashedPass),
		Role:      role,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func FakeUserRegisterOrLoginRequest() requests.UserRegisterOrLogin {
	return requests.UserRegisterOrLogin{
		Email:    email,
		Password: password,
	}
}

func FakeUserToLeaderRoleRequest() requests.UserToLeaderRole {
	return requests.UserToLeaderRole{
		Token:    token,
		Password: password,
	}
}

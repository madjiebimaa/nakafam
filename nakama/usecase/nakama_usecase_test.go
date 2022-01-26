package usecase_test

import (
	"testing"
)

// func TestCreate(t *testing.T) {
// 	fakeReq := requests.NakamaCreate{
// 		UserID:       primitive.NewObjectID(),
// 		Name:         "testing",
// 		UserName:     "testing",
// 		ProfileImage: "http://images.com/testing",
// 		Description:  "testing",
// 		SocialMedia: &responses.SocialMediaBase{
// 			Blogs: "http://my-blogs.com",
// 		},
// 	}

// 	fakeUser := domain.User{
// 		ID:    primitive.NewObjectID(),
// 		Email: "testing",
// 	}

// 	nakamaRepo := new(mocks.NakamaRepository)
// 	userRepo := new(mocks.UserRepository)
// 	familyRepo := new(mocks.FamilyRepository)
// 	nakamaUCase := usecase.NewNakamaUseCase(nakamaRepo, userRepo, familyRepo, 2*time.Second)

// 	t.Run("success", func(t *testing.T) {

// 		err := nakamaUCase.Create(context.TODO(), &fakeReq)
// 		assert.NoError(t, err)
// 		nakamaRepo.AssertExpectations(t)
// 		userRepo.AssertExpectations(t)
// 		familyRepo.AssertExpectations(t)
// 	})
// }

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}

func TestGetByID(t *testing.T) {

}

func TestGetAll(t *testing.T) {

}

func TestRegisterToFamily(t *testing.T) {

}

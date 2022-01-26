package fakes

import (
	"time"

	"github.com/google/uuid"
	"github.com/madjiebimaa/nakafam/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	// GENERAL
	role          = "staff"
	now           = time.Now()
	createdAt     = now
	updatedAt     = now
	token         = uuid.NewString()
	password      = "testing-password"
	hashedPass, _ = bcrypt.GenerateFromPassword([]byte(password), 12)
	description   = "testing-descprion"
	socialMedia   = domain.SocialMedia{
		Blogs:     "https://testing.com/blogs",
		Portfolio: "https://testing.com/protfolio",
		Email:     "testing-nakama@gmail.com",
		Github:    "https://testing.com/gtihub",
		Linkedin:  "https://testing.com/linkedin",
		Twitter:   "https://testing.com/twitter",
		Discord:   "https://testing.com/discord",
		Youtube:   "https://testing.com/youtube",
		Instagram: "https://testing.com/instagram",
	}

	// USER
	userID = primitive.NewObjectID()
	email  = "testing-user@gmail.com"

	// NAKAMA
	nakamaID           = primitive.NewObjectID()
	nakamaName         = "testing-nakama-name"
	nakamaUserName     = "testing-nakama-username"
	nakamaProfileImage = "https://testing.com/profile-image"

	// FAMILY
	familyID    = primitive.NewObjectID()
	familyName  = "testing-family-name"
	familyImage = "https://testing.com/family-image"
)

package route

import (
	"github.com/gin-gonic/gin"
	_familyHttpDelivery "github.com/madjiebimaa/nakafam/family/delivery/http"
	_nakamaHttpDelivery "github.com/madjiebimaa/nakafam/nakama/delivery/http"
	_userHttpDelivery "github.com/madjiebimaa/nakafam/user/delivery/http"
)

type Routes struct {
	userHttpDeliver    *_userHttpDelivery.UserHandler
	nakamaHttpDelivery *_nakamaHttpDelivery.NakamaDelivery
	familyHttpDelivery *_familyHttpDelivery.FamilyDelivery
}

func NewRoutes(
	userHttpDeliver *_userHttpDelivery.UserHandler,
	nakamaHttpDelivery *_nakamaHttpDelivery.NakamaDelivery,
	familyHttpDelivery *_familyHttpDelivery.FamilyDelivery,
) *Routes {
	return &Routes{
		userHttpDeliver,
		nakamaHttpDelivery,
		familyHttpDelivery,
	}
}

func (ro *Routes) Init(r *gin.Engine) {
	// cor := cors.Config{
	// 	AllowAllOrigins:  true,
	// 	AllowMethods:     []string{"GET, POST, PATCH, DELETE"},
	// 	AllowCredentials: true,
	// }

	// logger := gin.Logger()
	// recovery := gin.Recovery()
	api := r.Group("/api")

	// userMid := _userMid.NewUserMiddleware()

	// api.Use(cors.New(cor), logger, recovery)
	{
		api.POST("/users/register", ro.userHttpDeliver.Register)
		api.POST("/users/login", ro.userHttpDeliver.Login)
		api.PATCH("/users/upgrade-role/:token", ro.userHttpDeliver.ToLeaderRole)

		api.GET("/nakamas/:nakamaID", ro.nakamaHttpDelivery.GetByID)
		api.GET("/nakamas", ro.nakamaHttpDelivery.GetAll)

		api.GET("/families/:familyID", ro.familyHttpDelivery.GetByID)
		api.GET("/families", ro.familyHttpDelivery.GetAll)
	}

	// auth := api.Group("", userMid.IsAuth)
	// users := auth.Group("")
	// {
	// 	users.POST("/users/logout", ro.userHttpDeliver.Logout)
	// 	users.GET("/users/upgrade-role", ro.userHttpDeliver.UpgradeRole)
	// 	users.GET("/users/me", ro.userHttpDeliver.Me)
	// }

	// nakamas := auth.Group("")
	// {
	// 	nakamas.POST("/nakamas", ro.nakamaHttpDelivery.Create)
	// 	nakamas.PATCH("/nakamas/:nakamaID", ro.nakamaHttpDelivery.Update)
	// 	nakamas.DELETE("/nakamas/:nakamaID", ro.nakamaHttpDelivery.Delete)
	// 	nakamas.PATCH("/nakamas/:nakamaID/families/:familyID", ro.nakamaHttpDelivery.RegisterToFamily)
	// }

	// familiesLeader := auth.Group("", userMid.IsLeader)
	// {
	// 	familiesLeader.POST("/families", ro.familyHttpDelivery.Create)
	// 	familiesLeader.PATCH("/families/:familyID", ro.familyHttpDelivery.Update)
	// 	familiesLeader.DELETE("/families/:familyID", ro.familyHttpDelivery.Delete)
	// }
}

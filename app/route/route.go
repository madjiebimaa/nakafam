package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_familyHttpDelivery "github.com/madjiebimaa/nakafam/family/delivery/http"
	_nakamaHttpDelivery "github.com/madjiebimaa/nakafam/nakama/delivery/http"
	_userHttpDelivery "github.com/madjiebimaa/nakafam/user/delivery/http"
	_userMid "github.com/madjiebimaa/nakafam/user/delivery/http/middleware"
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
	cor := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET, POST, PATCH, DELETE"},
		AllowCredentials: true,
	}

	logger := gin.Logger()
	recovery := gin.Recovery()
	api := r.Group("/api")
	api.Use(cors.New(cor), logger, recovery)
	{
		api.POST("/users/register", ro.userHttpDeliver.Register)
		api.POST("/users/login", ro.userHttpDeliver.Login)
	}

	userMid := _userMid.NewUserMiddleware()

	auth := api.Group("", userMid.IsAuth)
	users := auth.Group("")
	{
		users.GET("/users/upgrade-role", ro.userHttpDeliver.UpgradeRole)
		users.PATCH("/users/upgrade-role/:token", ro.userHttpDeliver.ToLeaderRole)
		users.GET("/users/me", ro.userHttpDeliver.Me)
		users.POST("/users/logout", ro.userHttpDeliver.Logout)
		users.POST("/users/nakamas", ro.userHttpDeliver.CreateNakama)
	}

	nakamas := auth.Group("")
	{
		nakamas.GET("/nakamas/:nakama_id", ro.nakamaHttpDelivery.GetByID)
		nakamas.PATCH("/nakamas/:nakama_id", ro.nakamaHttpDelivery.Update)
		nakamas.DELETE("/nakamas/:nakama_id", ro.nakamaHttpDelivery.Delete)
		nakamas.GET("/nakamas", ro.nakamaHttpDelivery.GetAll)
	}

	nakamasLeader := auth.Group("", userMid.IsLeader)
	{
		// TODO: not implemented yet
		nakamasLeader.POST("/nakamas/families", ro.nakamaHttpDelivery.GetAll)
		nakamasLeader.PATCH("/nakamas/families", ro.nakamaHttpDelivery.GetAll)
	}

	families := auth.Group("")
	{
		// TODO: not implemented yet
		families.GET("/families/:family_id", ro.familyHttpDelivery.GetByID)
	}

	familiesLeader := auth.Group("", userMid.IsLeader)
	{
		// TODO: not implemented yet
		familiesLeader.DELETE("/families/:family_id", ro.familyHttpDelivery.GetByID)
	}
}

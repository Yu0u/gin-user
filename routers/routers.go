package routers

import (
	"github.com/gin-gonic/gin"
	v1 "study07/api/v1"
	"study07/middleware"
	"study07/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		auth.GET("user/info",v1.GetInfo)
	}

	router := r.Group("api/v1")
	{
		router.POST("register",v1.Register)
		router.POST("login",v1.Login)
	}

	_ = r.Run(utils.HttpPort)
}

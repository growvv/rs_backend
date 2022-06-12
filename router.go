package main

import (
	"github.com/gin-gonic/gin"
	"github.com/growvv/rs_demo/controller"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")
	apiRouter := r.Group("/rs")

	// basic apis
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)

}

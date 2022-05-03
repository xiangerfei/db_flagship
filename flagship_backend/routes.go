package main

import (
	"github.com/gin-gonic/gin"
	"xiangerfer.com/db_flagship/controller"
	"xiangerfer.com/db_flagship/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine{
	r.POST("/api/v1/auth/register", controller.Register)
	r.POST("/api/v1/auth/login", controller.Login)
	r.GET("/api/v1/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
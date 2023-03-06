package main

import (
	"github.com/gin-gonic/gin"
	"ginDemo/controller"
	"ginDemo/middleware"
)

func CollectRoute (r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	r.POST("/api/article/upload", controller.UploadFile)
	r.GET("/api/article/download", controller.DownloadFile)
	return r
}
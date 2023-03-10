package main

import (
	"ginDemo/controller"
	"ginDemo/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	r.POST("/api/auth/rename", middleware.AuthMiddleware(), controller.ReName)
	r.POST("/api/auth/repwd", middleware.AuthMiddleware(), controller.RePassword)

	r.POST("/api/article/upload", controller.UploadFile)
	r.GET("/api/article/download", controller.DownloadFile)
	return r
}

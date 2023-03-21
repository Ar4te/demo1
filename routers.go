package main

import (
	"ginDemo/controller"
	"ginDemo/middleware"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	lmt := tollbooth.NewLimiter(1, nil)
	lmt.SetMessage("服务繁忙，请稍后再试。。。")
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	r.POST("/api/auth/rename", middleware.AuthMiddleware(), controller.ReName)
	r.POST("/api/auth/rePwd", middleware.AuthMiddleware(), controller.RePassword)
	r.GET("/api/auth/testid", controller.TestId)

	r.POST("/api/article/upload", controller.UploadFile)
	r.GET("/api/article/download", controller.DownloadFile)
	r.GET("/api/article/getAll", middleware.LimitHandler(lmt), middleware.AuthMiddleware(), controller.GetAllArticle)
	r.POST("/api/article/updateFile", middleware.AuthMiddleware(), controller.UpdateFile)
	r.DELETE("/api/article/delFile", middleware.AuthMiddleware(), controller.DeleteFile)
	return r
}

package router

import (
	"ginDemo/controller"
	"ginDemo/middleware"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	lmt := tollbooth.NewLimiter(1, nil)
	lmt.SetMessage("服务繁忙，请稍后再试。。。")

	//user
	{
		r.POST("/api/auth/register", controller.Register)
		r.POST("/api/auth/login", controller.Login)
		r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
		r.POST("/api/auth/rename", middleware.AuthMiddleware(), controller.ReName)
		r.POST("/api/auth/rePwd", middleware.AuthMiddleware(), controller.RePassword)
		r.GET("/api/auth/testId", controller.TestId)
	}
	// article
	{
		r.POST("/api/article/upload", controller.UploadFile)
		r.GET("/api/article/download", controller.DownloadFile)
		r.GET("/api/article/getAll", middleware.LimitHandler(lmt), middleware.AuthMiddleware(), controller.GetAllArticle)
		r.POST("/api/article/updateFile", middleware.AuthMiddleware(), controller.UpdateFile)
		r.DELETE("/api/article/delFile", middleware.AuthMiddleware(), controller.DeleteFile)
	}
	// test
	{
		r.GET("/api/test/somJSON", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "someJSON", "status": 200})
		})
		r.GET("/api/test/someStruct", func(c *gin.Context) {
			var msg struct {
				Name    string
				Message string
				Number  int
			}
			msg.Name = "root"
			msg.Message = "message"
			msg.Number = 123
			c.JSON(200, msg)
		})
		r.GET("/api/test/someXML", func(c *gin.Context) {
			c.XML(200, gin.H{"message": "abc"})
		})
		r.GET("/api/test/someYAML", func(c *gin.Context) {
			c.YAML(200, gin.H{"name": "zhangsan"})
		})
		r.GET("/api/test/someProtoBuf", func(c *gin.Context) {
			reps := []int64{int64(1), int64(2)}
			label := "label"
			data := &protoexample.Test{
				Label: &label,
				Reps:  reps,
			}
			c.ProtoBuf(200, data)
		})
	}

	// topic
	{
		r.POST("/api/topic/pubtopic", controller.PubTopic)
	}
	return r
}

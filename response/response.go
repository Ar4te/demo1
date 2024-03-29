package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string, success bool) {
	c.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg, "success": success})
}

func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 200, data, msg, true)
}

func Fail(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 200, data, msg, false)
}

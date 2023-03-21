package middleware

import (
	"ginDemo/response"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
)

func LimitHandler(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			response.Fail(c, gin.H{"code": httpError.StatusCode, "msgContentType": lmt.GetMessageContentType()}, httpError.Message)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

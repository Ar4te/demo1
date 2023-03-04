package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"ginDemo/common"
	"ginDemo/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")

		// validate token format
		// strings.HasPrefix(), 以什么开头
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code":401,"msg":"权限不足"})
			// 抛弃这次请求
			c.Abort()
			return
		}

		// 从第7位后截取token有效部分
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code":401,"msg":"权限不足"})
			// 抛弃这次请求
			c.Abort()
			return
		}

		// 验证通过后获取claim中的UserId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户不存在
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code":401,"msg":"权限不足"})
			// 抛弃这次请求
			c.Abort()
			return
		}

		// 用户存在，将user信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
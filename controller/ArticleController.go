package controller

import (
	"github.com/gin-gonic/gin"
	"ginDemo/response"
	"ginDemo/common"
	"ginDemo/model"
	"net/http"
)

func UploadFile (c *gin.Context) {
	DB := common.GetDB()
	name := c.PostForm("name")
	userId := c.PostForm("userId")
	fileName := c.PostForm("fileName")

	var user model.User
	DB.First(&user, userId)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"code":401,"msg":"权限不足"})
		return
	}

	newArticle := model.Article{
		Name: name,
		UserId: userId,
		FileName: fileName,
	}

	DB.Create(&newArticle)

	response.Success(c, gin.H{"name":name, "userId":userId, "fileName": fileName}, "")
}
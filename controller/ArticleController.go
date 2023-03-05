package controller

import (
	"github.com/gin-gonic/gin"
	"ginDemo/response"
	"ginDemo/common"
	"ginDemo/model"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/base64"
)

func UploadFile (c *gin.Context) {
	DB := common.GetDB()
	name := c.PostForm("name")
	userId := c.PostForm("userId")
	fileName := c.PostForm("fileName")
	file, err := c.FormFile("file")
	if err != nil {
		msg := "get form err: " + err.Error()
		c.JSON(http.StatusBadRequest, gin.H{"code":400,"msg":msg})
		return
	}

	// 将*multipart.FileHeader转成*os.File
	_file, _ := file.Open()
	defer _file.Close()

	// 读取文件存为[]uint8
	fileData, _ := ioutil.ReadAll(_file)

	// 将[]uint8转成base64
	bs64 := base64.StdEncoding.EncodeToString(fileData)

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
		FileStream: bs64,
	}

	DB.Create(&newArticle)

	response.Success(c, gin.H{"name":name, "userId":userId, "fileName": fileName}, "上传成功")
}